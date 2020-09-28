include "../../../Services/Lock/LockTimestampedDistributedSystem.i.dfy"
include "TimestampedGLS.i.dfy"
include "Definitions.i.dfy"
include "../../../../Libraries/Math/mod_auto.i.dfy"
include "Invariants.i.dfy"


module PerformanceProof_i {
    import opened LockTimestampedDistributedSystem_i
    import opened TimestampedGLS_i
    import opened PerformanceProof__Definitions_i
    import opened Math__mod_auto_i
    import opened Invariants_i


/*****************************************************************************************
*                                 Performance Invariants                                 *
*****************************************************************************************/

predicate PerformanceInvariant(config:ConcreteConfiguration, tgls:TimestampedGLS_State)
    requires |tgls.history| > 0;
    requires forall ep :: ep in tgls.tls.t_servers <==> ep in config;
    requires forall ep | ep in tgls.tls.t_servers :: tgls.tls.t_servers[ep].v.config == config;
{
    && ConfigInvariant(config, tgls)
    && HistoryLengthInvariant(config, tgls)
    && TransferInvariant(tgls)
    && LockedInvariant(tgls)
}

predicate TransferInvariant(tgls:TimestampedGLS_State) {
    forall pkt  
        | && pkt in tgls.tls.t_environment.sentPackets 
          && pkt.msg.v.Transfer?
        :: && pkt.msg.v.transfer_epoch > 0
           &&TimeEq(pkt.msg.ts, PerfBoundLockInNetwork(pkt.msg.v.transfer_epoch))
} 


predicate LockedInvariant(tgls:TimestampedGLS_State) {
    forall ep 
        | && ep in tgls.tls.t_servers
          && tgls.tls.t_servers[ep].v.held
        :: && tgls.tls.t_servers[ep].v.epoch > 0
           && TimeEq(tgls.tls.t_servers[ep].ts, PerfBoundLockHeld(tgls.tls.t_servers[ep].v.epoch))
} 



/*****************************************************************************************
*                                 Performance Theorem                                    *
*****************************************************************************************/

predicate SingleGLSPerformanceAssumption(tgls:TimestampedGLS_State)
{
    // The only nodes that take steps are in the ring
    && (tgls.tls.t_environment.nextStep.LEnvStepHostIos? ==> tgls.tls.t_environment.nextStep.actor in tgls.tls.t_servers)
    // The size of the history/epochs never reaches UINT64_MAX
    && (|tgls.history| < 0xFFFF_FFFF_FFFF_FFFF)
    // No timeouts. This means that NodeAccept never takes the LIoOpTimeoutReceive branch
    && (var nextStep := tgls.tls.t_environment.nextStep; nextStep.LEnvStepHostIos? ==>
    (forall io :: io in nextStep.ios ==> !io.LIoOpTimeoutReceive?))
    // TODO: This should be part of the state machine definition, and not an assumption
    // The only nodes that can do NodeGrant() currently hold the lock
    && (var nextStep := tgls.tls.t_environment.nextStep; nextStep.LEnvStepHostIos? ==> (nextStep.nodeStep == GrantStep <==> tgls.tls.t_servers[nextStep.actor].v.held))
}


predicate GLSPerformanceAssumption(tglb:seq<TimestampedGLS_State>)
{
    forall tgls | tgls in tglb :: SingleGLSPerformanceAssumption(tgls)
}


predicate SingleGLSPerformanceGuarantee(gls:TimestampedGLS_State)
    requires |gls.tls.config| > 1;
{   
    /* We know that all packets in the ether are valid packets because we assume in 
    * SingleGLSPerformanceAssumption that nodes that take I/O steps must be in the config */
    forall pkt  
        | && pkt in gls.tls.t_environment.sentPackets 
          && pkt.msg.v == Transfer(|gls.tls.config|) 
        :: TimeEq(pkt.msg.ts, PerfBoundLockInNetwork(|gls.tls.config|))
}


predicate GLSPerformanceGuarantee(tglb:seq<TimestampedGLS_State>)
{
    && (forall tgls | tgls in tglb :: |tgls.tls.config| > 1)
    && (forall tgls | tgls in tglb :: SingleGLSPerformanceGuarantee(tgls))
}

/* Main Performance Theorem */
lemma PerformanceGuaranteeHolds(config:Config, tglb:seq<TimestampedGLS_State>)
    requires |config| > 1;
    requires ValidTimestampedGLSBehavior(tglb, config)
    requires GLSPerformanceAssumption(tglb)
    ensures GLSPerformanceGuarantee(tglb);
{
    lemma_ValidBehavior(config, tglb);
    lemma_ConfigInvariant(config, tglb);
    lemma_History_Length_Invariant(config, tglb);
    
    var i := 1;
    while i < |tglb| 
        decreases |tglb| - i;
        invariant 0 <= i <= |tglb|;
        invariant forall h | 0 <= h < i :: TransferInvariant(tglb[h]);
        invariant forall h | 0 <= h < i :: LockedInvariant(tglb[h]);
    {
        assume false;
        i := i + 1;
    }
}


lemma lemma_ValidBehavior(config:Config, tglb:seq<TimestampedGLS_State>) 
    requires |config| > 1;
    requires ValidTimestampedGLSBehavior(tglb, config);
    ensures forall i | 0 <= i < |tglb| - 1 :: TGLS_Next(tglb[i], tglb[i+1]);
{
    forall i | 0 <= i < |tglb| - 1
    ensures TGLS_Next(tglb[i], tglb[i+1]) {
        var h' := i + 1;
        var h := h' - 1;
        assert 0 <= h < h' < |tglb|;
        assert TGLS_Next(tglb[h], tglb[h']);
    }
}

}