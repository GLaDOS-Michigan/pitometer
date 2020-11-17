include "Types.dfy"
include "ZKDatabase.dfy"
include "ZKEnvironment.dfy"
include "LearnerHandler.dfy"


module ZooKeeper_Leader {
import opened ZooKeeper_Types
import opened ZooKeeper_ZKDatabase
import opened ZooKeeper_Environment
import opened ZooKeeper_LearnerHandler

/* Leader enters running state once `ackSet` forms a quorum */


datatype LeaderState = L_STARTING | L_RUNNING


// config[my_id] is my own endpoint, config[follower_id] is the follower endpoint
datatype Leader = Leader(
    my_id: nat,
    config: Config,
    state: LeaderState,
    globals: LeaderGlobals,

    handlers: seq<LearnerHandler>,
    nextHandlerToStep: int
)

/* Global Leader variables shared by all LearnerHandler threads 
datatype LeaderGlobals = LearnerHandler(
    zkdb: ZKDatabase,

    // Synchronization globals
    leaderEpoch: int,
    connectingFollowers: set<nat>,
    electingFollowers: set<nat>,
    ackSet: set<nat>
)
*/


/*****************************************************************************************
*                                      Actions                                           *
******************************************************************************************/ 

predicate LeaderInit(s:Leader, my_id:nat, config:Config, zkdb: ZKDatabase) {
    && s.my_id == my_id
    && s.config == config
    && s.state == L_STARTING
    && s.globals == LeaderGlobals(zkdb, -1, {}, {}, {})
    && InitHandlers(s.handlers, my_id, config)
    && s.nextHandlerToStep == 0
}

predicate LeaderNext(s:Leader, s':Leader, ios:seq<ZKIo>) {
    false
}

predicate LeaderStutter(s:Leader, s':Leader, ios:seq<ZKIo>) {
    false
}



/*****************************************************************************************
*                                    Helper defs                                         *
******************************************************************************************/ 


predicate InitHandlers(handlers:seq<LearnerHandler>, my_id: nat, config: Config) {
    && |handlers| == |config| - 1
    && forall i | 1 <= i < |config| :: LearnerHandlerInit(handlers[i-1], my_id, i, config)
}
}