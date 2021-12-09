include "../TimestampedRslSystem.i.dfy"
include "../../CommonProof/Constants.i.dfy"

module Common_Assumptions {
import opened TimestampedRslSystem_i
import opened CommonProof__Constants_i



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

predicate RslConsistency(ts:TimestampedRslState)
{
    ConstantsAllConsistentInv(UntimestampRslState(ts))
        && WellFormedLConfiguration(ts.constants.config)
}

}
