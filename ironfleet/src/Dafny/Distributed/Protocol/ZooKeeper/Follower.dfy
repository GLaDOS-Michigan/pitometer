include "Types.dfy"
include "ZKDatabase.dfy"
include "ZKEnvironment.dfy"


module ZooKeeper_Follower {
import opened ZooKeeper_Types
import opened ZooKeeper_ZKDatabase
import opened ZooKeeper_Environment


// config[my_id] is my own endpoint, config[leader_id] is the leader endpoint
datatype Follower = Follower(
    my_id: nat,
    leader_id: nat,
    config: Config,
    zkdb: ZKDatabase,
    accepted_epoch: int,
    state: FollowerState,

    serialLI: int,
    serialSync: int,
    serialNL: int
)

datatype FollowerState = F_HANDSHAKE_A | F_HANDSHAKE_B | F_PRESYNC | F_SYNC | F_RUNNING | F_ERROR

predicate FollowerInit(s:Follower, my_id:nat, leader_id:nat, config:Config, zkdb: ZKDatabase)
{
    && s.my_id == my_id
    && s.leader_id == leader_id
    && s.config == config
    && s.zkdb == zkdb
    && s.accepted_epoch == -1
    && s.state == F_HANDSHAKE_A

    && s.serialLI == -1
    && s.serialSync == -1
    && s.serialNL == -1
}

predicate FollowerNext(s:Follower, s':Follower, ios:seq<ZKIo>) {
    match s.state 
        case F_HANDSHAKE_A => SendMyInfo(s, s', ios)        // SendFI
        case F_HANDSHAKE_B => AcceptNewEpoch(s, s', ios)    // ProcLI
        case F_PRESYNC => PreSyncWithLeader(s, s', ios)     // ProcSync | ProcSyncSnap
        case F_SYNC => SyncWithLeader(s, s', ios)           // ProcSync
        case F_RUNNING => FollowerStutter(s, s', ios)       
        case F_ERROR => FollowerStutter(s, s', ios)
}

predicate FollowerStutter(s:Follower, s':Follower, ios:seq<ZKIo>) {
    && |ios| == 0
    && s' == s
}


/* State transition from F_HANDSHAKE_A -> F_HANDSHAKE_B
* Send_FI */
predicate SendMyInfo(s:Follower, s':Follower, ios:seq<ZKIo>) {
    && s' == s.(state := F_HANDSHAKE_B)
    && |ios| == 1
    && 0 <= s.leader_id < |s.config|
    && ios[0].LIoOpSend?
    && (var outbound_packet := ios[0].s;
        && outbound_packet.dst == s.config[s.leader_id]
        && outbound_packet.sender_index == s.my_id
        && outbound_packet.msg == FollowerInfo(s.my_id, getLastLoggedZxid(s.zkdb))
    )
}


/* State transition from F_HANDSHAKE_B -> F_PRESYNC 
* Process_LI */
predicate AcceptNewEpoch(s:Follower, s':Follower, ios:seq<ZKIo>) {
    // state and epoch are the only properties that change in this transition
    && s' == s.(state := s'.state, accepted_epoch := s'.accepted_epoch)  
    && |ios| >= 1
    && 0 <= s.leader_id < |s.config|
    && ios[0].LIoOpReceive?
    && ios[0].r.src in s.config
    && ios[0].r.msg.LeaderInfo?
    && ios[0].r.msg.sid == s.leader_id
    && (if ios[0].r.msg.newZxid.epoch <  s.accepted_epoch 
        then && s'.state == F_ERROR
             && |ios| == 1
        else (
            && s'.serialLI == ios[0].r.msg.serial
            && s'.state == F_PRESYNC
            && s'.accepted_epoch == ios[0].r.msg.newZxid.epoch
            && |ios| == 2
            && ios[1].LIoOpSend?
            && ios[1].s.dst == s.config[s.leader_id]
            && ios[1].s.sender_index == s.my_id
            && ios[1].s.msg.AckEpoch?
            && ios[1].s.msg.sid == s.my_id
            && ios[1].s.msg.serial == s'.serialLI
            && ios[1].s.msg.lastLoggedZxid == getLastLoggedZxid(s.zkdb)
            && if ios[0].r.msg.newZxid.epoch ==  s.accepted_epoch
                then ios[1].s.msg.lastAcceptedEpoch == -1
                else ios[1].s.msg.lastAcceptedEpoch == ios[0].r.msg.newZxid.epoch
        )
    )
}


/* State transition from F_PRESYNC -> F_SYNC */
predicate PreSyncWithLeader(s:Follower, s':Follower, ios:seq<ZKIo>) 
    requires s.state == F_PRESYNC
{
    && |ios| >= 1
    && ios[0].LIoOpReceive?
    && ios[0].r.src in s.config
    && ios[0].r.msg.sid == s.leader_id
    && match ios[0].r.msg
        // Ignore these 
        case FollowerInfo(sid, latestZxid) => false
        case LeaderInfo(sid, sn, newZxid) => false
        case AckEpoch(sid, serial, lastLoggedZxid, lastAcceptedEpoch) => false
        case Ack(sid, serial, ackZxid) => false
        case Commit(sid, serial, txn) => false
        case NewLeader(sid, serial, newLeaderZxid) => false
        case UpToDate(sid) => false

        // Sync messages
        case SyncDIFF(sid, serial, lastProcessedZxid) => 
            && s.serialSync == -1  
            && |ios| == 1
            && s' == s.(state := F_SYNC, serialSync:= serial)
        case SyncSNAP(sid, serial, leaderDb, lastProcessedZxid) =>
            && s.serialSync == -1  
            && |ios| == 1
            && s' == s.(zkdb := s'.zkdb, state := F_SYNC, serialSync:= serial)
            && ClearAndLoadDbSnapshot(s, s', leaderDb)
        case SyncTRUNC(sid, serial, lastProcessedZxid) => 
            // This case is unexpected
            && |ios| == 1
            && s' == s
            // && s.serialSync == -1  
            // && |ios| == 1
            // && s' == s.(zkdb := s'.zkdb, state := F_SYNC, serialSync:= serial)
            // && truncDatabase(s.zkdb, s'.zkdb, lastProcessedZxid)
}



/* State transition from F_SYNC -> F_RUNNING
* Note that F_SYNC has cycles */
predicate SyncWithLeader(s:Follower, s':Follower, ios:seq<ZKIo>) 
    requires s.state == F_SYNC
{
    && |ios| >= 1
    && ios[0].LIoOpReceive?
    && ios[0].r.src in s.config
    && ios[0].r.msg.sid == s.leader_id
    && match ios[0].r.msg
        // Ignore these 
        case FollowerInfo(sid, latestZxid) => false
        case LeaderInfo(sid, sn, newZxid) => false
        case AckEpoch(sid, serial, lastLoggedZxid, lastAcceptedEpoch) => false
        case Ack(sid, serial, ackZxid) => false
        case SyncDIFF(sid, serial, lastProcessedZxid) => false
        case SyncSNAP(sid, serial, leaderDb, lastProcessedZxid) => false
        case SyncTRUNC(sid, serial, lastProcessedZxid) => false

        // Process new transactions
        case Commit(sid, serial, txn) => 
            // This case is unexpected
            && |ios| == 1
            && s' == s
            // && |ios| == 1
            // && s' == s.(zkdb := s'.zkdb, serialSync:= serial)
            // && commitToLog(s.zkdb, s'.zkdb, txn)
        case NewLeader(sid, serial, newLeaderZxid) =>
            && s.serialNL == -1   
            && s' == s.(zkdb := s'.zkdb, serialNL:= serial)
            && takeSnapshot(s.zkdb, s'.zkdb)
            && |ios| == 2
            && ios[1].LIoOpSend?
            && ios[1].s.dst == ios[0].r.src
            && ios[1].s.sender_index == s.my_id
            && ios[1].s.msg == Ack(s.my_id, s'.serialNL, newLeaderZxid)

        // Terminating condition to move to next state
        case UpToDate(sid) => 
            // Send Ack with new epoch, and move to running state
            && |ios| == 1
            // && ios[1].LIoOpSend?
            // && ios[1].s.dst == ios[0].r.src
            // && ios[1].s.sender_index == s.my_id
            // && ios[1].s.msg == Ack(s.my_id, Zxid(s.accepted_epoch, 0))
            && s' == s.(state := F_RUNNING)
}


predicate ClearAndLoadDbSnapshot(s:Follower, s':Follower, snapshot:ZKDatabase){
    && snapshot.initialized
    && isValidZKDatabase(snapshot)
    && s'.zkdb == snapshot
}
}