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

    handlers: map<int,LearnerHandler>
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
         && s'.handlers ==  s.handlers
    else 
        && s' == s.(globals := s'.globals, handlers := s'.handlers)
        && StepSingleHandler(s, s', ios)
}

/* Leader already running, so db is in running state. Simply 
* move next handler by one step */
predicate LeaderRunStep(s:Leader, s':Leader, ios:seq<ZKIo>) 
    requires s.state == L_RUNNING
{
    // To make the proof easier, I am making the program halt immediately here
    LeaderStutter(s, s', ios)
}



/*****************************************************************************************
*                                  Handler Actions                                       *
******************************************************************************************/ 


predicate InitHandlers(handlers:map<int,LearnerHandler>, my_id: nat, config: Config) {
    && |handlers| == |config| - 1
    && (forall i | 0 <= i < |config| && i != my_id :: 
        && i in handlers
        && LearnerHandlerInit(handlers[i], my_id, i, config)
    )
}


predicate StepSingleHandler(s:Leader, s':Leader, ios:seq<ZKIo>) {
    || StepSingleHandler_Rcv(s, s', ios)
    || StepSingleHandler_NoRcv(s, s', ios)
}

/* Step the handler that is receiving io */
predicate StepSingleHandler_Rcv(s:Leader, s':Leader, ios:seq<ZKIo>) {
    && |ios| >= 1 
    && ios[0].LIoOpReceive?
    && var follower_id := ios[0].r.sender_index;  // received packets are bound for the right thread
    && LearnerHandlerNext(s, s', follower_id, ios)
}

/* Spontaneously step any handler that does not receive from network */
predicate StepSingleHandler_NoRcv(s:Leader, s':Leader, ios:seq<ZKIo>) {
    && (forall io | io in ios :: !io.LIoOpReceive?)
    && exists follower_id :: LearnerHandlerNext(s, s', follower_id, ios)
}

predicate LearnerHandlerNext(s:Leader, s':Leader, id:int, ios:seq<ZKIo>) {
    && id in s.handlers
    && s'.handlers.Keys == s.handlers.Keys
    && s'.handlers == s.handlers[id := s'.handlers[id]]
    && ZooKeeper_LearnerHandler.LearnerHandlerNext(s.handlers[id], s'.handlers[id], s.globals, s'.globals, ios)
}
}