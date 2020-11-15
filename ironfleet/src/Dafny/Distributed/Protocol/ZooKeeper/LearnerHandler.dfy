include "Types.dfy"
include "ZKDatabase.dfy"
include "ZKEnvironment.dfy"


module ZooKeeper_LearnerHandler {
import opened ZooKeeper_Types
import opened ZooKeeper_ZKDatabase
import opened ZooKeeper_Environment


// config[my_id] is my own endpoint, config[follower_id] is the follower endpoint
datatype LearnerHandler = LearnerHandler(
    my_id: nat,
    follower_id: nat,
    config: Config,
    zkdb: ZKDatabase,  //This is the leader db, and is not specified by LearnerHandler
    state: LearnerHandlerState 
)

datatype LearnerHandlerState = LH_HANDSHAKE_A | LH_HANDSHAKE_B | LH_DECIDE_SYNC | LH_SYNC | LH_RUNNING | LH_ERROR


predicate LearnerHandlerInit(s:LearnerHandler, my_id:nat, follower_id:nat, config:Config, zkdb: ZKDatabase) {
    && s.my_id == my_id
    && s.follower_id == follower_id
    && s.config == config
    && s.zkdb == zkdb
    && s.state == LH_HANDSHAKE_A
}

predicate LearnerHandlerNext(s:LearnerHandler, s':LearnerHandler, ios:seq<ZKIo>) {
    match s.state 
        case LH_HANDSHAKE_A => false
        case LH_HANDSHAKE_B => false
        case LH_DECIDE_SYNC => false
        case LH_SYNC => false
        case LH_RUNNING => LearnerHandlerStutter(s, s')
        case LH_ERROR => LearnerHandlerStutter(s, s')
}

predicate LearnerHandlerStutter(s:LearnerHandler, s':LearnerHandler) {
    s' == s
}

}