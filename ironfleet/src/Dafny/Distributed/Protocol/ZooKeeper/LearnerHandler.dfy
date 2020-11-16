include "Types.dfy"
include "ZKDatabase.dfy"
include "ZKEnvironment.dfy"


module ZooKeeper_LearnerHandler {
import opened ZooKeeper_Types
import opened ZooKeeper_ZKDatabase
import opened ZooKeeper_Environment


datatype LearnerHandlerState = LH_HANDSHAKE_A | LH_HANDSHAKE_B | LH_DECIDE_SYNC | LH_SYNC | LH_RUNNING | LH_ERROR
/*
* LH_HANDSHAKE_A: Receive FOLLOWERINFO, wait for quorum, then send new epoch
* LH_HANDSHAKE_B: Wait for a quorum of ACKEPOCH response. This ends the handshake with follower
* LH_DECIDE_SYNC: 
*/


// config[my_id] is my own endpoint, config[follower_id] is the follower endpoint
datatype LearnerHandler = LearnerHandler(
    my_id: nat,
    follower_id: nat,
    config: Config,
    zkdb: ZKDatabase,  //This is the leader db, and is not specified by LearnerHandler
    state: LearnerHandlerState,
    globals: LeaderGlobals,

    // Local state
    newEpoch: int
)


/* Global Leader variables shared by all LearnerHandler threads */
datatype LeaderGlobals = LearnerHandler(
    zkdb: ZKDatabase,

    // Handshake globals
    leaderEpoch: int,
    connectingFollowers: set<nat>
)


/*****************************************************************************************
*                                      Actions                                           *
******************************************************************************************/ 

predicate LearnerHandlerInit(s:LearnerHandler, my_id:nat, follower_id:nat, config:Config, zkdb: ZKDatabase, globals: LeaderGlobals) {
    && s.my_id == my_id
    && s.follower_id == s.follower_id  // follower id is initially unknown
    && s.config == config
    && s.zkdb == zkdb
    && s.state == LH_HANDSHAKE_A
    && s.globals == globals

    && s.newEpoch == -1
}

predicate LearnerHandlerNext(s:LearnerHandler, s':LearnerHandler, ios:seq<ZKIo>) {
    match s.state 
        case LH_HANDSHAKE_A => GetEpochToPropose(s, s', ios)
        case LH_HANDSHAKE_B => false
        case LH_DECIDE_SYNC => false
        case LH_SYNC => false
        case LH_RUNNING => LearnerHandlerStutter(s, s')
        case LH_ERROR => LearnerHandlerStutter(s, s')
}

predicate LearnerHandlerStutter(s:LearnerHandler, s':LearnerHandler) {
    s' == s
}

predicate GetEpochToPropose(s:LearnerHandler, s':LearnerHandler, ios:seq<ZKIo>) {
    if IsVerifiedQuorum(s.my_id, |s.config|, s.globals.connectingFollowers) 
    then (
        // Send Leader.LEADERINFO message to follower, and proceed to LH_HANDSHAKE_B state
        && |ios| == 1
        && s' == s.(state := LH_HANDSHAKE_B, newEpoch := s.globals.leaderEpoch)
        && ios[0].LIoOpSend?
        && (var outbound_packet := ios[0].s;
            && 0 <= s.follower_id < |s.config|
            && outbound_packet.dst == s.config[s.follower_id]
            && outbound_packet.msg == LeaderInfo(s.my_id, Zxid(s.globals.leaderEpoch, 0))
        )       
    ) else (
        // Add sender to my connectingFollowers set, and continue waiting for quorum
        if |ios| == 0 then LearnerHandlerStutter(s, s')  // Case where follower has not sent anything
        else && |ios| == 1
             && ios[0].LIoOpReceive?
             && ios[0].r.src in s.config
             && ios[0].r.msg.FollowerInfo?
             && 0 <= ios[0].r.msg.sid < |s.config| 
             && s.config[ios[0].r.msg.sid] == ios[0].r.src
             && s' == s.(
                 follower_id := ios[0].r.msg.sid,
                 globals := s.globals.(
                    leaderEpoch := (if ios[0].r.msg.latestZxid.epoch >= s.globals.leaderEpoch then ios[0].r.msg.latestZxid.epoch + 1 else s.globals.leaderEpoch),
                    connectingFollowers := s.globals.connectingFollowers + {ios[0].r.msg.sid}
                )
             )
    )
}


/*****************************************************************************************
*                                    Helper defs                                         *
******************************************************************************************/ 

predicate IsVerifiedQuorum(my_id:nat, n:int, quorum: set<nat>) {
    && my_id in quorum
    && |quorum| >= (n/2) + 1
}
}