include "../TimestampedRslSystem.i.dfy"
include "../../CommonProof/Constants.i.dfy"
include "definitions.i.dfy"

module Common_Assumptions {
import opened TimestampedRslSystem_i
import opened CommonProof__Constants_i
import opened Common_Definitions



/*****************************************************************************************
*                                     Assumptions                                        *
*****************************************************************************************/


predicate CommonAssumptions(ts:TimestampedRslState) {
    && RslConsistency(ts)
    && NoPacketDuplication(ts)
    && NoExternalPackets(ts)
    && NoExternalSteps(ts)
    && BoundedQueueingAssumption(ts)
    && |ts.t_replicas| > 2  // For failure to be meaningful
    && ts.t_replicas[0].v.replica.proposer.constants.all.params.max_batch_size > 0
    && ts.t_replicas[1].v.replica.proposer.constants.all.params.max_batch_size > 0
    && ts.constants.params.max_batch_size == 1
}

predicate NoExternalPackets(ts:TimestampedRslState) {
    forall pkt | pkt in ts.undeliveredPackets ::
        && pkt in ts.t_environment.sentPackets
        && pkt.src in ts.constants.config.replica_ids
        && pkt.dst in ts.constants.config.replica_ids
}

predicate NoExternalSteps(ts:TimestampedRslState) {
    !(exists eid, ios :: ts.t_environment.nextStep == LEnvStepHostIos(eid, ios, ExternalStep()))
}

// Bounded queueing assumption for simplified RSL
predicate BoundedQueueingAssumption(ts:TimestampedRslState) 
    requires RslConsistency(ts)
{
    forall idx, ios :: (
        && 0 <= idx < |ts.constants.config.replica_ids|
        && ts.t_environment.nextStep == LEnvStepHostIos(ts.constants.config.replica_ids[idx], ios, RslStep(ts.t_replicas[idx].v.nextActionIndex))
        ==>
        (forall io | io in ts.t_environment.nextStep.ios && io.LIoOpReceive? ::
            // this means that max(replica.ts, msg.ts) <= msg.ts + MaxQueueTime
            ts.t_replicas[idx].ts <= io.r.msg.ts + MaxQueueTime
        )
    )
}

/* Assume that the leader does not receive leftover 2b packets from before leader election */
predicate NewLeaderDoesNotReceiveOld2a2b(ts:TimestampedRslState) 
    requires |ts.t_replicas| > 2 
{
    var nextStep := ts.t_environment.nextStep;
    nextStep.LEnvStepHostIos? ==>
    && (forall io | io in nextStep.ios && io.LIoOpReceive? && io.r.msg.v.RslMessage_2b? :: io.r.msg.v.bal_2b != Ballot(1, 0))
    && (forall io | io in nextStep.ios && io.LIoOpReceive? && io.r.msg.v.RslMessage_2a? :: io.r.msg.v.bal_2a != Ballot(1, 0))
}

predicate NewLeaderDoesNotProposeFurtherOps(ts:TimestampedRslState, opn:OperationNumber) 
    requires |ts.t_replicas| > 2 
{
    ts.t_replicas[1].v.replica.proposer.constants.all.params.max_integer_val == UpperBoundFinite(opn + 1)
}

}
