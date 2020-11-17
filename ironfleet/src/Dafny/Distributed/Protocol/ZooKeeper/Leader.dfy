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
    match s.state
        case L_STARTING => LeaderStartStep(s, s', ios)
        case L_RUNNING => LeaderRunStep(s, s', ios)
}

predicate LeaderStutter(s:Leader, s':Leader, ios:seq<ZKIo>) {
    s' == s && |ios| == 0
}

/* Wait for `ackSet` to form a quorum, then start db and transition to Running state */
predicate LeaderStartStep(s:Leader, s':Leader, ios:seq<ZKIo>) 
    requires s.state == L_STARTING
{
    false
}

/* Leader already running, so db is in running state. Simply 
* move next handler by one step */
predicate LeaderRunStep(s:Leader, s':Leader, ios:seq<ZKIo>) 
    requires s.state == L_RUNNING
{
    && 0 <= s.nextHandlerToStep < |s.handlers|
    && var handler := s.handlers[s.nextHandlerToStep];
    && s' == s.(globals := s'.globals, handlers := s'.handlers, nextHandlerToStep := IncNextHandlerToStep(s.nextHandlerToStep, |s.handlers|))
    && |s'.handlers| == |s.handlers|
    && forall i | 0 <= i < |s.handlers| ::
        if i == s.nextHandlerToStep 
        then LearnerHandlerNext(s.handlers[i], s'.handlers[i], s.globals, s'.globals, ios)
        else s'.handlers[i] == s.handlers[i]
}



/*****************************************************************************************
*                                    Helper defs                                         *
******************************************************************************************/ 


predicate InitHandlers(handlers:seq<LearnerHandler>, my_id: nat, config: Config) {
    && |handlers| == |config| - 1
    && forall i | 1 <= i < |config| :: LearnerHandlerInit(handlers[i-1], my_id, i, config)
}

function IncNextHandlerToStep(i: int, n: int) : int {
    (i + 1) % n
}
}