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
    state: FollowerState 
)

datatype FollowerState = F_HANDSHAKE_A | F_HANDSHAKE_B | F_SYNC | F_RUNNING | F_ERROR

predicate FollowerInit(s:Follower, my_id:nat, leader_id:nat, config:Config, zkdb: ZKDatabase)
{
    && s.my_id == my_id
    && s.leader_id == leader_id
    && s.config == config
    && s.zkdb == zkdb
    && s.accepted_epoch == -1
    && s.state == F_HANDSHAKE_A
}

predicate FollowerNext(s:Follower, s':Follower, ios:seq<ZKIo>) {
    match s.state 
        case F_HANDSHAKE_A => SendMyInfo(s, s', ios)        // SendFI
        case F_HANDSHAKE_B => AcceptNewEpoch(s, s', ios)    // ProcLI
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
        && outbound_packet.msg == FollowerInfo(s.my_id, getLastLoggedZxid(s.zkdb))
    )
}


/* State transition from F_HANDSHAKE_B -> F_SYNC 
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
            && s'.state == F_SYNC
            && s'.accepted_epoch == ios[0].r.msg.newZxid.epoch
            && |ios| == 2
            && ios[1].LIoOpSend?
            && ios[1].s.dst == s.config[s.leader_id]
            && ios[1].s.msg.AckEpoch?
            && ios[1].s.msg.sid == s.my_id
            && ios[1].s.msg.lastLoggedZxid == getLastLoggedZxid(s.zkdb)
            && if ios[0].r.msg.newZxid.epoch ==  s.accepted_epoch
                then ios[1].s.msg.lastAcceptedEpoch == -1
                else ios[1].s.msg.lastAcceptedEpoch == ios[0].r.msg.newZxid.epoch
        )
    )
}


/* State transition from F_SYNC -> F_RUNNING
* Note that F_SYNC has cycles */
predicate SyncWithLeader(s:Follower, s':Follower, ios:seq<ZKIo>) {
    && |ios| >= 1
    && ios[0].LIoOpReceive?
    && ios[0].r.src in s.config
    && ios[0].r.msg.sid == s.leader_id
    && match ios[0].r.msg
        // Ignore these 
        case FollowerInfo(sid, latestZxid) => FollowerStutter(s, s', ios)
        case LeaderInfo(sid, newZxid) => FollowerStutter(s, s', ios)
        case AckEpoch(sid, lastLoggedZxid, lastAcceptedEpoch) => FollowerStutter(s, s', ios)
        case Ack(sid, ackZxid) => FollowerStutter(s, s', ios)

        // Sync messages
        case SyncDIFF(sid, lastProcessedZxid) => 
            && |ios| == 1
            && FollowerStutter(s, s', ios)
        case SyncSNAP(sid, leaderDb, lastProcessedZxid) =>
            && |ios| == 1
            && ClearAndLoadDbSnapshot(s, s', leaderDb)
        case SyncTRUNC(sid, lastProcessedZxid) => 
            && |ios| == 1
            && s' == s.(zkdb := s'.zkdb)
            && truncDatabase(s.zkdb, s'.zkdb, lastProcessedZxid)

        // Process new transactions
        case Commit(sid, txn) => 
            && |ios| == 1
            && ProcessTxn(s, s', txn)
        case NewLeader(sid, newLeaderZxid) =>
            && s' == s.(zkdb := s'.zkdb)
            && takeSnapshot(s.zkdb, s'.zkdb)
            && |ios| == 2
            && ios[1].LIoOpSend?
            && ios[1].s.dst == ios[0].r.src
            && ios[1].s.msg == Ack(s.my_id, newLeaderZxid)

        // Terminating condition to move to next state
        case UpToDate(sid) => 
            // Send Ack with new epoch, and move to running state
            && |ios| == 2
            && ios[1].LIoOpSend?
            && ios[1].s.dst == ios[0].r.src
            && ios[1].s.msg == Ack(s.my_id, Zxid(s.accepted_epoch, 0))
            && s' == s.(state := F_RUNNING)
}


predicate ClearAndLoadDbSnapshot(s:Follower, s':Follower, snapshot:ZKDatabase){
    s' == s.(zkdb := snapshot)
}

predicate ProcessTxn(s:Follower, s':Follower, txn:Zxid) {
    // Not what actually happens in ZooKeeper, but I'm simplifying this step to simply
    // append to the log
    && s' == s.(zkdb := s'.zkdb)
    && commitToLog(s.zkdb, s'.zkdb, txn)
}
}