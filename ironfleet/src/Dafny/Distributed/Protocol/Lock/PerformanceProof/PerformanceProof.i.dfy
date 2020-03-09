include "../../../Services/Lock/LockTaggedDistributedSystem.i.dfy"
  include "TaggedGLS.i.dfy"

module PerformanceProof_i {
import opened LockTaggedDistributedSystem_i
  import opened TaggedGLS_i

predicate SingleGLSPerformanceAssumption(tgls:TaggedGLS_State)
{
  tgls.tls.t_environment.nextStep.LEnvStepHostIos? ==> tgls.tls.t_environment.nextStep.actor in tgls.tls.t_servers
}

predicate GLSPerformanceAssumption(tglb:seq<TaggedGLS_State>)
{
  forall tgls :: tgls in tglb ==> SingleGLSPerformanceAssumption(tgls)
}

predicate SingleGLSPerformanceGuarantee(gls:TaggedGLS_State)
{
  forall pkt :: pkt in gls.tls.t_environment.sentPackets &&
    pkt.msg.v == Locked(|gls.tls.config|) ==> pkt.msg.pr == PerformanceReport(0, 0)
}

predicate GLSPerformanceGuarantee(tglb:seq<TaggedGLS_State>)
{
  forall tgls :: tgls in tglb ==> SingleGLSPerformanceGuarantee(tgls)
}

lemma PerformanceGuaranteeHolds(config:Config, tglb:seq<TaggedGLS_State>)
  requires ValidTaggedGLSBehavior(tglb, config)
  requires GLSPerformanceAssumption(tglb)
  ensures GLSPerformanceGuarantee(tglb)
{
  
}

}
