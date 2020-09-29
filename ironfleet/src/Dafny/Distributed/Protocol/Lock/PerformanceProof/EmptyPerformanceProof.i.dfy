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

predicate PerformanceInductiveInvariant(config:ConcreteConfiguration, tgls:TimestampedGLS_State)
    requires |tgls.history| > 0;
{
    && ConfigInvariant(config, tgls)
    && HistoryLengthInvariant(config, tgls)
    && TransferInvariant(tgls)
    && LockedInvariant(tgls)
    && NeverHeldInvariant(tgls)
}


predicate TransferInvariant(tgls:TimestampedGLS_State) {
    forall pkt  
        | && pkt in tgls.tls.t_environment.sentPackets 
          && pkt.msg.v.Transfer?
          && 0 < pkt.msg.v.transfer_epoch <= |tgls.tls.config|  // first round 
        ::
          TimeEq(pkt.msg.ts, PerfBoundLockInNetwork(pkt.msg.v.transfer_epoch))
} 


predicate LockedInvariant(tgls:TimestampedGLS_State) {
    forall ep 
        | && ep in tgls.tls.t_servers
          && tgls.tls.t_servers[ep].v.held
          && 0 < tgls.tls.t_servers[ep].v.epoch <= |tgls.tls.config|  // first round
        ::
          TimeEq(tgls.tls.t_servers[ep].ts, PerfBoundLockHeld(tgls.tls.t_servers[ep].v.epoch))
} 


predicate NeverHeldInvariant(tgls:TimestampedGLS_State) {
    forall ep 
        | && ep in tgls.tls.t_servers
          && tgls.tls.t_servers[ep].v.epoch == 0 
        ::
          TimeEq(tgls.tls.t_servers[ep].ts, TimeZero());
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
        invariant forall h | 0 <= h < i :: PerformanceInductiveInvariant(config, tglb[h]);
    {
        var tls, tls' := tglb[i-1].tls, tglb[i].tls;
        if tls.t_environment.nextStep.LEnvStepHostIos? && tls.t_environment.nextStep.actor in tls.t_servers {
            PerformanceGuaranteeHolds_Induction_IOStep(config, tglb[i-1], tglb[i]);
        } else {
            PerformanceGuaranteeHolds_Induction_NonIOStep(config, tglb[i-1], tglb[i]);
        }
        forall h | 0 <= h <= i 
        ensures PerformanceInductiveInvariant(config, tglb[h])
        {}
        i := i + 1;
    }
    PerformanceInductiveInvariant_Implies_GLSPerformanceGuarantee(config, tglb);
}


lemma PerformanceGuaranteeHolds_Induction_IOStep(config:Config, tgls:TimestampedGLS_State, tgls':TimestampedGLS_State)
    // Standard pre-conditions
    requires |tgls.history| > 0 && |tgls'.history| > 0;
    requires PerformanceInductiveInvariant(config, tgls);
    requires ConfigInvariant(config, tgls');
    requires HistoryLengthInvariant(config, tgls');
    requires SingleGLSPerformanceAssumption(tgls) && SingleGLSPerformanceAssumption(tgls');
    requires TGLS_Next(tgls, tgls');
    // Branch condition
    requires tgls.tls.t_environment.nextStep.LEnvStepHostIos? && tgls.tls.t_environment.nextStep.actor in tgls.tls.t_servers;
    // Post-conditions
    ensures PerformanceInductiveInvariant(config, tgls');
{
    var tls, tls' := tgls.tls, tgls'.tls;
    var id, ios, hstep := tls.t_environment.nextStep.actor, tls.t_environment.nextStep.ios, tls.t_environment.nextStep.nodeStep;

    assert TLS_NextOneServer(tls, tls', id, ios, hstep);
    if |ios| > 0 && ios[0].LIoOpReceive? {
        PerformanceGuaranteeHolds_Induction_IOStep_Accept(config, tgls, tgls');
    } else {
        PerformanceGuaranteeHolds_Induction_IOStep_Grant(config, tgls, tgls');
    }
    assert TransferInvariant(tgls');
    assert LockedInvariant(tgls');
}


lemma PerformanceGuaranteeHolds_Induction_IOStep_Accept(config:Config, tgls:TimestampedGLS_State, tgls':TimestampedGLS_State)
    // Standard pre-conditions
    requires |tgls.history| > 0 && |tgls'.history| > 0;
    requires PerformanceInductiveInvariant(config, tgls);
    requires ConfigInvariant(config, tgls');
    requires HistoryLengthInvariant(config, tgls');
    requires SingleGLSPerformanceAssumption(tgls) && SingleGLSPerformanceAssumption(tgls');
    requires TGLS_Next(tgls, tgls');
    // Branch condition
    requires tgls.tls.t_environment.nextStep.LEnvStepHostIos? && tgls.tls.t_environment.nextStep.actor in tgls.tls.t_servers;
    requires |tls.t_environment.nextStep.ios| > 0 && tls.t_environment.nextStep.ios[0].LIoOpReceive?
    // Post-conditions
    ensures PerformanceInductiveInvariant(config, tgls');
{
    var tls, tls' := tgls.tls, tgls'.tls;
    var e, e' := tls.t_environment, tls'.t_environment;
    var id, ios, hstep := e.nextStep.actor, e.nextStep.ios, e.nextStep.nodeStep;
    assert TLS_NextOneServer(tls, tls', id, ios, hstep);

    assert tls'.t_servers[id].ts == TLS_RecvPerfUpdate(tls.t_servers[id].ts, ios[0].r.msg.ts, hstep);
    assert hstep == AcceptStep;
    var ls, ls', untagged_ios := UntagLS_State(tls), UntagLS_State(tls'), UntagLIoOpSeq(ios);
    assert NodeAccept(ls.servers[id], ls'.servers[id], untagged_ios);
    if !ls.servers[id].held 
        && untagged_ios[0].r.src in ls.servers[id].config
        && untagged_ios[0].r.msg.Transfer? 
        && untagged_ios[0].r.msg.transfer_epoch > ls.servers[id].epoch 
    {
        /* I do not hold the lock. Hence, the updated counter is the max of the 
        * message counter and my local counter. The message counter has the desired 
        * counter by TransferInvariant. If my local epoch is 0, then my local counter
        * is 0 by NeverHeldInvariant, and I get the desired local update that 
        * satisfies LockedInvariant.
        * If my local epoch is x != 0, Then my epoch must by x + n|config| for n >= 0. 
        * (LEMMA) Then LockedInvariant is trivially satisfied. */ 
        assume false;
    } else {
        /* By assumption, I do not hold lock. Hence, TransferInvariant and 
        * LockedInvariant are trivially satisfied for the next step because I am not 
        * emitting any Transfer packets, nor am I going to hold the lock. Hence, I 
        * am tasked with satisfying NeverHeldInvariant in the next step.
        * If my ep != 0, invariant is trivially satisfied.
        * Else, if receiving a Transfer message, any Transfer message destined for me 
        * must be in flight, so contradiction (LEMMA). I could also be receiving a 
        * Locked message. In this case my epoch cannot be 0 (LEMMA), 
        * so contradiction again. */
        assume false;
    }
}


lemma PerformanceGuaranteeHolds_Induction_IOStep_Grant(config:Config, tgls:TimestampedGLS_State, tgls':TimestampedGLS_State)
    // Standard pre-conditions
    requires |tgls.history| > 0 && |tgls'.history| > 0;
    requires PerformanceInductiveInvariant(config, tgls);
    requires ConfigInvariant(config, tgls');
    requires HistoryLengthInvariant(config, tgls');
    requires SingleGLSPerformanceAssumption(tgls) && SingleGLSPerformanceAssumption(tgls');
    requires TGLS_Next(tgls, tgls');
    // Branch condition
    requires tgls.tls.t_environment.nextStep.LEnvStepHostIos? && tgls.tls.t_environment.nextStep.actor in tgls.tls.t_servers;
    requires !(|tls.t_environment.nextStep.ios| > 0 && tls.t_environment.nextStep.ios[0].LIoOpReceive?)
    // Post-conditions
    ensures PerformanceInductiveInvariant(config, tgls');
{
    var tls, tls' := tgls.tls, tgls'.tls;
    var e, e' := tls.t_environment, tls'.t_environment;
    var id, ios, hstep := e.nextStep.actor, e.nextStep.ios, e.nextStep.nodeStep;
    assert TLS_NextOneServer(tls, tls', id, ios, hstep);

    assert tls'.t_servers[id].ts == TLS_NoRecvPerfUpdate(tls.t_servers[id].ts, hstep);
    assert hstep == GrantStep;
    var ls, ls', untagged_ios := UntagLS_State(tls), UntagLS_State(tls'), UntagLIoOpSeq(ios);
    assert NodeGrant(ls.servers[id], ls'.servers[id], untagged_ios);
    if ls.servers[id].held && ls.servers[id].epoch < 0xFFFF_FFFF_FFFF_FFFF {
        /* Locked invariant easily implies Transfer invariant in the next step.
        * ep != 0, so NeverHeldInvariant trivially satisfied. */
        assume false;
    } else {
        /* If my epoch >= 0xFFFF_FFFF_FFFF_FFFF, by HistoryLengthInvariant, |history|
        * >= 0xFFFF_FFFF_FFFF_FFFF, violating SingleGLSPerformanceAssumption.
        * Else, I don't hold lock. This violates SingleGLSPerformanceAssumption. */
        assume false;
    }
}


lemma PerformanceGuaranteeHolds_Induction_NonIOStep(config:Config, tgls:TimestampedGLS_State, tgls':TimestampedGLS_State)
    // Standard pre-conditions
    requires |tgls.history| > 0 && |tgls'.history| > 0;
    requires PerformanceInductiveInvariant(config, tgls);
    requires ConfigInvariant(config, tgls');
    requires HistoryLengthInvariant(config, tgls');
    requires SingleGLSPerformanceAssumption(tgls) && SingleGLSPerformanceAssumption(tgls');
    requires TGLS_Next(tgls, tgls');
    // Branch condition
    requires !(tgls.tls.t_environment.nextStep.LEnvStepHostIos? && tgls.tls.t_environment.nextStep.actor in tgls.tls.t_servers);
    // Post-conditions
    ensures PerformanceInductiveInvariant(config, tgls');
{
    var tls, tls' := tgls.tls, tgls'.tls;
    var e, e' := tls.t_environment, tls'.t_environment;
    assert LEnvironment_Next(e, e');
    match e.nextStep {
        case LEnvStepHostIos(actor, ios, nodeStep) => {
            assert SingleGLSPerformanceAssumption(tgls);
            assert e.nextStep.actor in tls.t_servers;  // by assumption
            assert false;
        }
        case LEnvStepDeliverPacket(p) => assert e'.sentPackets == e.sentPackets;
        case LEnvStepAdvanceTime => assert e'.sentPackets == e.sentPackets;
        case LEnvStepStutter => assert e'.sentPackets == e.sentPackets;
    }
    assert tls'.t_servers == tls.t_servers;
    assert e'.sentPackets == e.sentPackets;
    assert TransferInvariant(tgls');
    assert LockedInvariant(tgls');
}


lemma PerformanceInductiveInvariant_Implies_GLSPerformanceGuarantee(config:Config, tglb:seq<TimestampedGLS_State>)
    requires |config| > 1;
    requires forall i | 0 <= i < |tglb| :: |tglb[i].history| > 0;
    requires ValidTimestampedGLSBehavior(tglb, config);
    requires GLSPerformanceAssumption(tglb);
    requires forall i | 0 <= i < |tglb| :: PerformanceInductiveInvariant(config, tglb[i]);
    ensures GLSPerformanceGuarantee(tglb);
{
    // TODO TONY
    assume false;
}
}