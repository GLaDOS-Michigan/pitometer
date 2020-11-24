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
    && s.globals == LeaderGlobals(zkdb, -1, {my_id}, {my_id}, {my_id})  // Leader is defacto part of all quorums
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
    if IsVerifiedQuorum(s.my_id, |s.config|, s.globals.ackSet) 
    then && |ios| == 0
         && s' == s.(state := L_RUNNING, globals := s'.globals)
         && s'.globals == s.globals.(zkdb := s'.globals.zkdb)
         && s'.globals.zkdb == s.globals.zkdb.(isRunning := true)
    else 
        && 0 <= s.nextHandlerToStep < |s.handlers|
        && s' == s.(globals := s'.globals, handlers := s'.handlers, nextHandlerToStep := IncNextHandlerToStep(s.nextHandlerToStep, |s.handlers|))
        && StepSingleHandler(s, s', ios)
}

/* Leader already running, so db is in running state. Simply 
* move next handler by one step */
predicate LeaderRunStep(s:Leader, s':Leader, ios:seq<ZKIo>) 
    requires s.state == L_RUNNING
{
    && 0 <= s.nextHandlerToStep < |s.handlers|
    && s' == s.(globals := s'.globals, handlers := s'.handlers, nextHandlerToStep := IncNextHandlerToStep(s.nextHandlerToStep, |s.handlers|))
    && StepSingleHandler(s, s', ios)
}



/*****************************************************************************************
*                                    Helper defs                                         *
******************************************************************************************/ 


predicate InitHandlers(handlers:seq<LearnerHandler>, my_id: nat, config: Config) {
    && |handlers| == |config| - 1
    && forall i | 1 <= i < |config| :: LearnerHandlerInit(handlers[i-1], my_id, i, config)
}

function IncNextHandlerToStep(i: int, n: int) : int {
    if n == 0 then 0 else (i + 1) % n
}

predicate StepSingleHandler(s:Leader, s':Leader, ios:seq<ZKIo>) {
    && |s'.handlers| == |s.handlers|
    && 0 <= s.nextHandlerToStep < |s.handlers|
    && var follower_id := s.handlers[s.nextHandlerToStep].follower_id;
    && (forall io | io in ios && io.LIoOpReceive? :: io.r.sender_index == follower_id)  // received packets are bound for the right thread
    && LearnerHandlerNext(s.handlers[s.nextHandlerToStep], s'.handlers[s.nextHandlerToStep], s.globals, s'.globals, ios)
    && (forall i | 0 <= i < |s.handlers| && i != s.nextHandlerToStep :: s'.handlers[i] == s.handlers[i])
}
}