include "Phase2Proof.i.dfy"

module Rs2Phase1Proof_Tony_i {
import opened RslPhase2Proof_i


/*
predicate BehaviorPerformanceAssumption(con:LConstants, tb:seq<TimestampedRslState>, req_time:Timestamp, opn:OperationNumber) {
    && Pre2aInvariant(tb[0], req_time, opn)
    && (forall tgls :: tgls in tb ==> RslAssumption(tgls))
}

predicate BehaviorPerformanceGuarantee(con:LConstants, tb:seq<TimestampedRslState>, req_time:Timestamp) {
    forall tgls :: tgls in tb ==> PerformanceGuarantee(tgls, req_time)
}


predicate RslPerfInvariant(s:TimestampedRslState, req_time:Timestamp, opn:OperationNumber, t2a:Timestamp) {
    && RslConsistency(s)
    && (
        || (exists progresses :: Phase2AcceptedLeaderInvariant(s, req_time, opn, t2a, progresses))
        || (exists progresses :: Phase2UnacceptedLeaderInvariant(s, req_time, opn, t2a, progresses))
        || Phase2CompletedInvariant(s, req_time, opn)
    )
}
*/


lemma PerformanceGuaranteeHolds(con:LConstants, tb:seq<TimestampedRslState>, req_time:Timestamp) 
    requires ValidTimestampedRSLBehavior(con, tb)
    requires BehaviorPerformanceAssumption(con, tb)
    ensures BehaviorPerformanceGuarantee(con, tb, req_time)
{
    assert PerformanceGuarantee(tb[0], req_time);
    var i := 0;
    while i < |tb| 
        decreases |tb| - i
        invariant 0 <= i <= |tb|
        invariant forall k | 0 <= k < i :: !Phase2CompletedInvariant(tb[k], req_time, opn) ==> RslPerfInvariant(tb[k], req_time)
    {
        PerfInvariantMaintained(tlb[i], tlb[i+1]);
        i := i + 1;
    }
}


lemma PerfInvariantMaintained(s:TimestampedRslState, s':TimestampedRslState, req_time:Timestamp, opn:OperationNumber, t2a:Timestamp)
    requires RslAssumption(s) && RslConsistency(s)
    requires RslAssumption(s') && RslConsistency(s')
    requires TimestampedRslNext(s, s')
    requires RslPerfInvariant(s, req_time, opn, t2a)
    requires !Phase2CompletedInvariant(s, req_time, opn)
    ensures RslPerfInvariant(s', req_time, opn, t2a)
{   
    if (exists progresses :: Phase2AcceptedLeaderInvariant(s, req_time, opn, t2a, progresses)) {
        var progresses :| Phase2AcceptedLeaderInvariant(s, req_time, opn, t2a, progresses);
        if (exists idx, ios :: TimestampedRslNextOneReplica(s, s', idx, ios)){
            var j :| TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);
            if j == 0 {
                if s.t_environment.nextStep.nodeStep == RslStep(6) {
                    var p' := lemma_leader_execute_P2GoesToP2(s, s', j, req_time, opn, t2a, progresses);
                } 
                else if s.t_environment.nextStep.nodeStep == RslStep(0) {
                    var p' := lemma_leader_receive_P2GoesToP2(s, s', j, req_time, opn, t2a, progresses);
                }
                else {
                    var p' := lemma_leader_notreceiveOrExecute_P2GoesToP2(s, s', j, req_time, opn, t2a, progresses);
                }
            } else {
                if s.t_environment.nextStep.nodeStep == RslStep(0) {
                    var p' := lemma_notleader_receive_P2GoesToP2(s, s', j, req_time, opn, t2a, progresses);
                } else {
                    var p' := lemma_notleader_notreceive_P2GoesToP2(s, s', j, req_time, opn, t2a, progresses);
                }
            }
        } else {
            assert Phase2AcceptedLeaderInvariant(s', req_time, opn, t2a, progresses);
        }
    } else {
        var progresses :| Phase2UnacceptedLeaderInvariant(s, req_time, opn, t2a, progresses);
        if (exists idx, ios :: TimestampedRslNextOneReplica(s, s', idx, ios)){
            var j :| TimestampedRslNextOneReplica(s, s', j, s.t_environment.nextStep.ios);
            if j == 0 {
                if s.t_environment.nextStep.nodeStep == RslStep(0) {
                    var p' := lemma_leader_receive_P2UnacceptedGoesToP2(s, s', j, req_time, opn, t2a, progresses);
                } else {
                    var p' := lemma_leader_notreceive_P2UnacceptedGoesToP2Unaccepted(s, s', j, req_time, opn, t2a, progresses);
                }
            } else {
                if s.t_environment.nextStep.nodeStep == RslStep(0) {
                    var p' := lemma_notleader_receive_P2UnacceptedGoesToP2Unaccepted(s, s', j, req_time, opn, t2a, progresses);
                } else {
                    var p' := lemma_notleader_notreceive_P2UnacceptedGoesToP2Unaccepted(s, s', j, req_time, opn, t2a, progresses);
                }
            }
        } else {
            assert Phase2UnacceptedLeaderInvariant(s', req_time, opn, t2a, progresses);
        }
    }
}
}