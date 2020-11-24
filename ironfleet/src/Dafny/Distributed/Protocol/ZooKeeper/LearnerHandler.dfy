include "Types.dfy"
include "ZKDatabase.dfy"
include "ZKEnvironment.dfy"


module ZooKeeper_LearnerHandler {
import opened ZooKeeper_Types
import opened ZooKeeper_ZKDatabase
import opened ZooKeeper_Environment


datatype LearnerHandlerState = 
                | LH_HANDSHAKE_A | LH_HANDSHAKE_B 
                | LH_PREP_SYNC | LH_SYNC
                | LH_PROCESS_ACK
                | LH_RUNNING | LH_ERROR

datatype SyncMode = SNAP | DIFF | TRUNC


// config[my_id] is my own endpoint, config[follower_id] is the follower endpoint
datatype LearnerHandler = LearnerHandler(
    my_id: nat,
    follower_id: nat,
    config: Config,
    state: LearnerHandlerState,

    // Local state
    queuedPackets: seq<ZKMessage>,
    newEpoch: int,
    peerLastZxid: Zxid
)


/* Global Leader variables shared by all LearnerHandler threads */
datatype LeaderGlobals = LeaderGlobals(
    zkdb: ZKDatabase,

    // Synchronization globals
    leaderEpoch: int,
    connectingFollowers: set<nat>,
    electingFollowers: set<nat>,
    ackSet: set<nat>
)


/*****************************************************************************************
*                                      Actions                                           *
******************************************************************************************/ 

predicate LearnerHandlerInit(s:LearnerHandler, my_id:nat, follower_id:nat, config:Config) {
    && s.my_id == my_id
    && s.follower_id == follower_id
    && s.config == config
    && s.state == LH_HANDSHAKE_A

    && s.queuedPackets == []
    && s.newEpoch == -1
    && s.peerLastZxid == NullZxid
}

predicate LearnerHandlerNext(s:LearnerHandler, s':LearnerHandler, g:LeaderGlobals, g':LeaderGlobals, ios:seq<ZKIo>) {
    match s.state 
        case LH_HANDSHAKE_A => GetEpochToPropose(s, s', g, g', ios)
        case LH_HANDSHAKE_B => WaitForEpochAck(s, s', g, g', ios)
        case LH_PREP_SYNC => PrepareSync(s, s', g, g', ios)
        case LH_SYNC => DoSync(s, s', g, g', ios)
        case LH_PROCESS_ACK => ProcessAck(s, s', g, g', ios)
        case LH_RUNNING => LearnerHandlerStutter(s, s', ios) && g' == g
        case LH_ERROR => LearnerHandlerStutter(s, s', ios) && g' == g
}

predicate LearnerHandlerStutter(s:LearnerHandler, s':LearnerHandler, ios:seq<ZKIo>) {
    && |ios| == 0
    && s' == s
}

/* Process_FI */
predicate GetEpochToPropose(s:LearnerHandler, s':LearnerHandler, g:LeaderGlobals, g':LeaderGlobals, ios:seq<ZKIo>) 
    requires s.state == LH_HANDSHAKE_A
{
    if IsVerifiedQuorum(s.follower_id, |s.config|, g.connectingFollowers) 
    then ( // Send Leader.LEADERINFO message to follower, and proceed to LH_HANDSHAKE_B state
        && g' == g
        && |ios| == 1
        && s' == s.(state := LH_HANDSHAKE_B, newEpoch := g.leaderEpoch)
        && ios[0].LIoOpSend?
        && (var outbound_packet := ios[0].s;
            && 0 <= s.follower_id < |s.config|
            && outbound_packet.dst == s.config[s.follower_id]
            && outbound_packet.msg == LeaderInfo(s.my_id, Zxid(g.leaderEpoch, 0))
        )       
    ) else ( // Add sender to my connectingFollowers set, and continue waiting for quorum
        // if |ios| == 0 then LearnerHandlerStutter(s, s', ios) && g' == g  // Case where follower has not sent anything
        // else 
            && |ios| == 1
             && ios[0].LIoOpReceive?
             && ios[0].r.src in s.config
             && ios[0].r.msg.FollowerInfo?
             && 0 <= ios[0].r.msg.sid < |s.config| 
             && ios[0].r.msg.sid == s.follower_id
             && s.config[ios[0].r.msg.sid] == ios[0].r.src
             && s' == s.(
                 follower_id := ios[0].r.msg.sid,
                 peerLastZxid := ios[0].r.msg.latestZxid
             ) 
             && g' == g.(
                    leaderEpoch := (if ios[0].r.msg.latestZxid.epoch >= g.leaderEpoch then ios[0].r.msg.latestZxid.epoch + 1 else g.leaderEpoch),
                    connectingFollowers := g.connectingFollowers + {ios[0].r.msg.sid}
                ) 
    )
}


predicate WaitForEpochAck(s:LearnerHandler, s':LearnerHandler, g:LeaderGlobals, g':LeaderGlobals, ios:seq<ZKIo>) 
    requires s.state == LH_HANDSHAKE_B
{
    if IsVerifiedQuorum(s.follower_id, |s.config|, g.electingFollowers) 
    then (
        // Proceed to state sync
        && g' == g
        && |ios| == 0   
        && s' == s.(state := LH_PREP_SYNC)
    ) else (
        // Add sender to my electingFollowers set, store peerLastZxid, and continue waiting for quorum
        // if |ios| == 0 then LearnerHandlerStutter(s, s', ios) && g' == g // Case where follower has not sent anything
        // else 
            && |ios| == 1
             && ios[0].LIoOpReceive?
             && ios[0].r.src in s.config
             && ios[0].r.msg.AckEpoch?
             && 0 <= s.follower_id < |s.config| 
             && s.config[s.follower_id] == ios[0].r.src
             && ios[0].r.msg.sid == s.follower_id
             && s' == s.(peerLastZxid := ios[0].r.msg.lastLoggedZxid)
             && g' == g.(
                    electingFollowers := g.electingFollowers + {ios[0].r.msg.sid}
            )
    )
}


predicate PrepareSync(s:LearnerHandler, s':LearnerHandler, g:LeaderGlobals, g':LeaderGlobals, ios:seq<ZKIo>) 
    requires s.state == LH_PREP_SYNC
{
    if ! (g.zkdb.initialized && isValidZKDatabase(g.zkdb)) then |ios| == 0 && s' == s.(state := LH_ERROR) && g' == g
    else
    var proposals := getInMemorySuffix(g.zkdb);
    && |ios| == 0  // no I/O in this step
    && s' == s.(state := s'.state)  //syncMode to be modified accordingly
    && g' == g
    && if |proposals| > 0 
        then (
            if ZxidLt(g.zkdb.maxCommittedLog, s.peerLastZxid) 
            then    && s'.state == LH_SYNC
                    && s'.queuedPackets == [SyncTRUNC(s.my_id, g.zkdb.maxCommittedLog)]
            else if ZxidLt(s.peerLastZxid, g.zkdb.minCommittedLog) 
            then && s'.state == LH_SYNC
                 && s'.queuedPackets == [SyncSNAP(s.my_id, g.zkdb, getLastLoggedZxid(g.zkdb))]
            else  // peerLastZxid is in the range of my proposals list
                if s.peerLastZxid !in proposals then s' == s.(state := LH_ERROR) 
                else
                && s'.state == LH_SYNC
                && s'.queuedPackets == [SyncDIFF(s.my_id, s.peerLastZxid)] + PrepareDiffCommits(s.my_id, proposals, s.peerLastZxid)
       ) else (
            // Sync empty diff
            && s'.state == LH_SYNC
            && s'.queuedPackets == [SyncDIFF(s.my_id, s.peerLastZxid)]
       )
}


predicate DoSync(s:LearnerHandler, s':LearnerHandler, g:LeaderGlobals, g':LeaderGlobals, ios:seq<ZKIo>) 
    requires s.state == LH_SYNC
{
    && g' == g
    && |ios| == 1 
    && ios[0].LIoOpSend?
    && 0 <= s.follower_id < |s.config| 
    && ios[0].s.dst == s.config[s.follower_id]
    && if |s.queuedPackets| == 0 
        then // Done with sync. Send NewLeader msg
            && s' == s.(state := LH_PROCESS_ACK)
            && ios[0].s.msg == NewLeader(s.my_id, getLastLoggedZxid(g.zkdb))
        else // Send next item in queuedPackets. Remain in LH_SYNC state
            && s' == s.(queuedPackets := s.queuedPackets[1..])
            && ios[0].s.msg == s.queuedPackets[0]
}


predicate ProcessAck(s:LearnerHandler, s':LearnerHandler, g:LeaderGlobals, g':LeaderGlobals, ios:seq<ZKIo>) 
    requires s.state == LH_PROCESS_ACK
{
    // Leader is in charge of checking AckSet and starting the actual local DB
    && 0 <= s.follower_id < |s.config| 
    && if g.zkdb.isRunning  // Leader is in charge of starting db once AckSet is a good quorum
    then (
        // Proceed to Running state, send UPTODATE to follower
        && s' == s.(state := LH_RUNNING)
        && g' == g
        && |ios| == 1
        && ios[0].LIoOpSend?
        && ios[0].s.dst == s.config[s.follower_id]
        && ios[0].s.msg == UpToDate(s.my_id)
    ) else (
        // Add sender to my ackSet set, store peerLastZxid, and continue waiting for quorum
        // if |ios| == 0 then LearnerHandlerStutter(s, s', ios)  && g' == g // Case where follower has not sent anything
        // else 
            && |ios| == 1
             && ios[0].LIoOpReceive?
             && ios[0].r.src in s.config
             && ios[0].r.msg.Ack?
             && s.config[s.follower_id] == ios[0].r.src
             && s' == s
             && g' == g.(ackSet := g.ackSet + {ios[0].r.msg.sid})
    )
}



/*****************************************************************************************
*                                    Helper defs                                         *
******************************************************************************************/ 

predicate IsVerifiedQuorum(my_id:nat, n:int, quorum: set<nat>) {
    && my_id in quorum
    && |quorum| >= (n/2) + 1
}

function PrepareDiffCommits(leaderId:nat, proposals:seq<Zxid>, peerLastZxid:Zxid) : seq<ZKMessage> 
    decreases proposals
{
    if |proposals| == 0 then []
    else if peerLastZxid in proposals then PrepareDiffCommits(leaderId, proposals[1..], peerLastZxid)
    else [Commit(leaderId, proposals[0])] + PrepareDiffCommits(leaderId, proposals[1..], peerLastZxid)
}
}