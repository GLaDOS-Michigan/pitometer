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
/                                      EpochInvariant                                    *
*****************************************************************************************/

predicate EpochInvariant(config:Config, tgls:TimestampedGLS_State) 
    requires ConfigInvariant(config, tgls);
{  
    && EpochInvariant_Packets(config, tgls)
    && EpochInvariant_Nodes(config, tgls)
}

predicate IsValidPacket(config:Config, tgls:TimestampedGLS_State, pkt:LPacket<EndPoint, LockMessage>) {
    && pkt in UntagSentPkts(tgls.tls.t_environment.sentPackets)
    && pkt.src in config
    && pkt.dst in config
}

predicate EpochInvariant_Packets(config:Config, tgls:TimestampedGLS_State) 
    requires ConfigInvariant(config, tgls);
{  
    reveal_ConfigInvariant();
    // All packets have valid source 
    && (forall pkt | pkt in tgls.tls.t_environment.sentPackets :: pkt.src in config)
    // All valid packets cannot have Invalid message
    && (forall pkt  | IsValidPacket(config, tgls, pkt)
          :: pkt.msg.Locked? || pkt.msg.Transfer?)
    // All valid transfer packets have epoch > 1 and epoch congurent to dest index mod |config|
    && (forall pkt  
        | && IsValidPacket(config, tgls, pkt)
          && pkt.msg.Transfer?
        :: 
          && pkt.msg.transfer_epoch > 1
          && CongrentModM(tgls.tls.t_servers[pkt.dst].v.my_index+1, pkt.msg.transfer_epoch, |config|)
        )
    // All in-flight packets have epoch == source epoch
    && (forall pkt | IsInFlightTransferMessage(config, tgls, pkt) 
           :: pkt.msg.transfer_epoch == tgls.tls.t_servers[pkt.dst].v.epoch
    )
    // All valid lock packets have dest epoch > 0.
    && (forall pkt  
        | && IsValidPacket(config, tgls, pkt)
          && pkt.msg.Locked?
        :: 
          tgls.tls.t_servers[pkt.dst].v.epoch > 0)
}

predicate EpochInvariant_Nodes(config:Config, tgls:TimestampedGLS_State) 
    requires ConfigInvariant(config, tgls);
{  
    reveal_ConfigInvariant();
    // All nodes have non-negative epoch 0 or epoch congurent to index mod |config|
    && (forall ep | ep in config :: tgls.tls.t_servers[ep].v.epoch >= 0)
    && (forall ep | ep in config && tgls.tls.t_servers[ep].v.epoch != 0 :: CongrentModM(tgls.tls.t_servers[ep].v.my_index+1, tgls.tls.t_servers[ep].v.epoch, |config|))
}

predicate {:opaque} CongrentModM(a:int, b:int, m:int) 
    requires m > 0;
{
    && b >= a 
    && a == b % m
}


lemma lemma_EpochInvariant(config:Config, tglb:seq<TimestampedGLS_State>) 
    requires |config| > 1;
    requires ValidTimestampedGLSBehavior(tglb, config);
    requires GLSPerformanceAssumption(tglb);
    requires forall k | 0 <= k < |tglb| :: ConfigInvariant(config, tglb[k]);
    ensures forall k | 0 <= k < |tglb| :: EpochInvariant(config, tglb[k]);
{
    lemma_mod_auto(|config|);
    lemma_ValidBehavior(config, tglb);
    reveal_ConfigInvariant();
    reveal_CongrentModM();
    forall ep | ep in config && tglb[0].tls.t_servers[ep].v.epoch != 0 
    ensures CongrentModM(tglb[0].tls.t_servers[ep].v.my_index + 1, tglb[0].tls.t_servers[ep].v.epoch, |config|)
    {
        var index := tglb[0].tls.t_servers[ep].v.my_index;
        if index == 0 {
            assert tglb[0].tls.t_servers[ep].v.epoch == 1;
            assert CongrentModM(1, 1, |config|);
        } else {
            assert tglb[0].tls.t_servers[ep].v.epoch == 0;
        }
    }

    var i := 1;
    while i < |tglb| 
        decreases |tglb| - i;
        invariant 0 <= i <= |tglb|;
        invariant forall k | 0 <= k < i :: EpochInvariant(config, tglb[k]);
    {   
        var tgls, tgls' := tglb[i-1], tglb[i];

        if tgls.tls.t_environment.nextStep.LEnvStepHostIos? && tgls.tls.t_environment.nextStep.actor in tgls.tls.t_servers {
            lemma_EpochInvariant_IOStep(config, tgls, tgls');
        } else {
            lemma_EpochInvariant_NonIOStep(config, tgls, tgls');
        }
        assert EpochInvariant(config, tglb[i]);
        forall k | 0 <= k <= i 
        ensures EpochInvariant(config, tglb[k]) 
        {}
        i := i + 1;
    }
}


lemma lemma_EpochInvariant_IOStep(config:Config, tgls:TimestampedGLS_State, tgls':TimestampedGLS_State) 
    // Standard pre-conditions
    requires ConfigInvariant(config, tgls) && ConfigInvariant(config, tgls');
    requires SingleGLSPerformanceAssumption(tgls) && SingleGLSPerformanceAssumption(tgls');
    requires TGLS_Next(tgls, tgls');
    requires EpochInvariant(config, tgls);
    // Branch condition
    requires tgls.tls.t_environment.nextStep.LEnvStepHostIos? && tgls.tls.t_environment.nextStep.actor in tgls.tls.t_servers;
    // Post-conditions
    ensures EpochInvariant(config, tgls');
{
    reveal_ConfigInvariant();
    reveal_CongrentModM();
    var tls, tls' := tgls.tls, tgls'.tls;
    var ls, ls' := UntagLS_State(tls), UntagLS_State(tls');
    var e, e' := ls.environment, ls'.environment;
    
    var id, ios, step := e.nextStep.actor, e.nextStep.ios, e.nextStep.nodeStep;
    assert LS_NextOneServer(ls, ls', id, ios, step);
    if NodeGrant(ls.servers[id], ls'.servers[id], ios) {
        /* Node Grant step */
        assume false;
    } else {
        /* Node Accept step */
        if !ls.servers[id].held 
           && ios[0].r.src in ls.servers[id].config
           && ios[0].r.msg.Transfer? 
           && ios[0].r.msg.transfer_epoch > ls.servers[id].epoch 
        {
            var pkt := ios[0].r;
            assert CongrentModM(ls.servers[pkt.dst].my_index+1, pkt.msg.transfer_epoch, |config|);
            assert ios[1].s.msg.locked_epoch == pkt.msg.transfer_epoch > 0;
            assert e'.sentPackets == e.sentPackets + {ios[1].s};
        } else {
            assert e'.sentPackets == e.sentPackets;
            assert ls'.servers == ls.servers;
        }
    }
}


lemma lemma_EpochInvariant_NonIOStep(config:Config, tgls:TimestampedGLS_State, tgls':TimestampedGLS_State) 
    // Standard pre-conditions
    requires ConfigInvariant(config, tgls) && ConfigInvariant(config, tgls');
    requires SingleGLSPerformanceAssumption(tgls) && SingleGLSPerformanceAssumption(tgls');
    requires TGLS_Next(tgls, tgls');
    requires EpochInvariant(config, tgls);
    // Branch condition
    requires !(tgls.tls.t_environment.nextStep.LEnvStepHostIos? && tgls.tls.t_environment.nextStep.actor in tgls.tls.t_servers);
    // Post-conditions
    ensures EpochInvariant(config, tgls');
{
    reveal_ConfigInvariant();
    var tls, tls' := tgls.tls, tgls'.tls;
    var ls, ls' := UntagLS_State(tls), UntagLS_State(tls');
    var e, e' := ls.environment, ls'.environment;

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
    assert forall ep | ep in config :: ls'.servers[ep].epoch == ls'.servers[ep].epoch;
}


/*****************************************************************************************
*                                 Performance Invariants                                 *
*****************************************************************************************/

predicate CommonInvariants(config:ConcreteConfiguration, tgls:TimestampedGLS_State)
    requires |tgls.history| > 0;
{
    reveal_ConfigInvariant();
    && ConfigInvariant(config, tgls)
    && HistoryLengthInvariant(config, tgls)
    && EpochInvariant(config, tgls)
}

predicate PerformanceInductiveInvariant(config:ConcreteConfiguration, tgls:TimestampedGLS_State)
    requires |tgls.history| > 0;
{
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
          && 0 < tgls.tls.t_servers[ep].v.epoch < |tgls.tls.config|  // first round
        ::
          TimeEq(tgls.tls.t_servers[ep].ts, PerfBoundLockHeld(tgls.tls.t_servers[ep].v.epoch))
} 


predicate NeverHeldInvariant(tgls:TimestampedGLS_State) {
    forall ep 
        | && ep in tgls.tls.t_servers
          && tgls.tls.t_servers[ep].v.epoch == 0 
        ::
          TimeEq(tgls.tls.t_servers[ep].ts, TimeZero())
}



/*****************************************************************************************
*                                 Performance Theorem                                    *
*****************************************************************************************/

predicate SingleGLSPerformanceAssumption(tgls:TimestampedGLS_State)
{
    // The only nodes that take steps are in the ring
    && (tgls.tls.t_environment.nextStep.LEnvStepHostIos? 
        ==> 
        tgls.tls.t_environment.nextStep.actor in tgls.tls.t_servers)
    // The size of the history/epochs never reaches UINT64_MAX
    && (|tgls.history| < 0xFFFF_FFFF_FFFF_FFFF)
    // No timeouts. This means that NodeAccept never takes the LIoOpTimeoutReceive branch
    && (tgls.tls.t_environment.nextStep.LEnvStepHostIos? 
        ==>
        (forall io | io in tgls.tls.t_environment.nextStep.ios :: !io.LIoOpTimeoutReceive?))
    // TODO: This should be part of the state machine definition, and not an assumption
    // The only nodes that can do NodeGrant() currently hold the lock
    && (tgls.tls.t_environment.nextStep.LEnvStepHostIos? 
        ==> 
        (tgls.tls.t_environment.nextStep.nodeStep == GrantStep <==> tgls.tls.t_servers[tgls.tls.t_environment.nextStep.actor].v.held))
}


predicate GLSPerformanceAssumption(tglb:seq<TimestampedGLS_State>)
{
    forall i | 0 <= i < |tglb| :: SingleGLSPerformanceAssumption(tglb[i])
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
    reveal_ConfigInvariant();
    lemma_ValidBehavior(config, tglb);
    lemma_ConfigInvariant(config, tglb);
    lemma_History_Length_Invariant(config, tglb);
    lemma_EpochInvariant(config, tglb);
    
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
    requires CommonInvariants(config, tgls) && CommonInvariants(config, tgls');
    requires PerformanceInductiveInvariant(config, tgls);
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
    requires CommonInvariants(config, tgls) && CommonInvariants(config, tgls');
    requires PerformanceInductiveInvariant(config, tgls);
    requires SingleGLSPerformanceAssumption(tgls) && SingleGLSPerformanceAssumption(tgls');
    requires TGLS_Next(tgls, tgls');
    // Branch condition
    requires tgls.tls.t_environment.nextStep.LEnvStepHostIos? && tgls.tls.t_environment.nextStep.actor in tgls.tls.t_servers;
    requires |tgls.tls.t_environment.nextStep.ios| > 0 && tgls.tls.t_environment.nextStep.ios[0].LIoOpReceive?
    // Post-conditions
    ensures PerformanceInductiveInvariant(config, tgls');
{
    var tls, tls' := tgls.tls, tgls'.tls;
    var e, e' := tls.t_environment, tls'.t_environment;
    var id, ios, hstep := e.nextStep.actor, e.nextStep.ios, e.nextStep.nodeStep;
    var ls, ls', untagged_ios := UntagLS_State(tls), UntagLS_State(tls'), UntagLIoOpSeq(ios);

    assert ios[0].r in e.sentPackets;
    assert IsValidLIoOp(ios[0], id, e);
    assert untagged_ios[0].r.dst == id;

    assert hstep == AcceptStep;
    assert NodeAccept(ls.servers[id], ls'.servers[id], untagged_ios);
    if !ls.servers[id].held 
        && untagged_ios[0].r.src in ls.servers[id].config
        && untagged_ios[0].r.msg.Transfer? 
        && untagged_ios[0].r.msg.transfer_epoch > ls.servers[id].epoch 
    {
        var pkt := ios[0].r;
        if 0 < pkt.msg.v.transfer_epoch <= |tgls.tls.config| {
            /* I do not hold the lock. Hence, the updated counter is the max of the 
            * message counter and my local counter. The message counter has the desired 
            * counter by TransferInvariant. If my local epoch is 0, then my local counter
            * is 0 by NeverHeldInvariant, and I get the desired local update that 
            * satisfies LockedInvariant.
            * If my local epoch != 0, Then my epoch must be >= my_index. Since transfer_epoch
            * > my epoch, transfer_epoch >= my_index + |config| (EpochInvariant).
            * Then LockedInvariant is trivially satisfied. */ 
            if ls.servers[id].epoch == 0 {
                assert tls.t_servers[id].ts == TimeZero();
                assert TimeEq(pkt.msg.ts, PerfBoundLockInNetwork(pkt.msg.v.transfer_epoch));
                Accept_j_helper();
                assert TimeEq(tls'.t_servers[id].ts, PerfBoundLockHeld(tls'.t_servers[id].v.epoch));
            } else {
                lemma_mod_auto(|config|);
                reveal_CongrentModM();
                assert pkt.msg.v.transfer_epoch > ls.servers[id].my_index;
                assert pkt.msg.v.transfer_epoch >= ls.servers[id].my_index + |config|;
                assert tls'.t_servers[id].v.epoch == pkt.msg.v.transfer_epoch > |config|;
            }
        } else {
            assert tls'.t_servers[id].v.epoch == pkt.msg.v.transfer_epoch > |config|;
        }
        assert LockedInvariant(tgls');
    } else {
        /* By assumption, I do not hold lock. Hence, TransferInvariant and 
        * LockedInvariant are trivially satisfied for the next step because I am not 
        * emitting any Transfer packets, nor am I going to hold the lock. Hence, I 
        * am tasked with satisfying NeverHeldInvariant in the next step.
        * If my ep != 0, invariant is trivially satisfied.
        * Else, if receiving a Transfer message, any Transfer message destined for me 
        * must be in flight, so contradiction (EpochInvariant). 
        * I could also be receiving a Locked message. In this case my epoch cannot 
        * be 0 (EpochInvariant), so contradiction again. */
        assert !ls.servers[id].held && !ls'.servers[id].held;
        if ls.servers[id].epoch == 0 {
            if untagged_ios[0].r.msg.Transfer? {
                assert untagged_ios[0].r.msg.transfer_epoch <= ls.servers[id].epoch == 0;
                assert EpochInvariant(config, tgls);
                assert false;
            } else {
                assert untagged_ios[0].r.msg.Locked?;
                assert EpochInvariant(config, tgls);
                assert false;
            }
        }
    }
}


lemma PerformanceGuaranteeHolds_Induction_IOStep_Grant(config:Config, tgls:TimestampedGLS_State, tgls':TimestampedGLS_State)
    // Standard pre-conditions
    requires |tgls.history| > 0 && |tgls'.history| > 0;
    requires CommonInvariants(config, tgls) && CommonInvariants(config, tgls');
    requires PerformanceInductiveInvariant(config, tgls);
    requires SingleGLSPerformanceAssumption(tgls) && SingleGLSPerformanceAssumption(tgls');
    requires TGLS_Next(tgls, tgls');
    // Branch condition
    requires tgls.tls.t_environment.nextStep.LEnvStepHostIos? && tgls.tls.t_environment.nextStep.actor in tgls.tls.t_servers;
    requires !(|tgls.tls.t_environment.nextStep.ios| > 0 && tgls.tls.t_environment.nextStep.ios[0].LIoOpReceive?)
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
        assert e'.sentPackets == e.sentPackets + {ios[0].s};
        if ls.servers[id].epoch < |tgls.tls.config| {
            assert ls'.servers[id].epoch == ls.servers[id].epoch == |tgls.history| > 0;
            var node_pr := PerfBoundLockHeld(tls.t_servers[id].v.epoch);
            assert tls.t_servers[id].ts == node_pr;
            var node_pr_next := TimeAdd2(node_pr, StepToTimeDelta(hstep));
            Grant_j_helper();
            assert TimeEq(node_pr_next, PerfBoundLockInNetwork(tls.t_servers[id].v.epoch+1));
            assert ios[0].s.msg.ts == node_pr_next;
        } else {
            assert ls'.servers[id].epoch > 0;
            assert ios[0].s.msg.v.transfer_epoch > |tgls.tls.config|;
        }
    } else {
        /* I must be holding the lock, as stipulated by SingleGLSPerformanceAssumption. 
        * Hence, my epoch >= 0xFFFF_FFFF_FFFF_FFFF. By HistoryLengthInvariant, |history|
        * >= 0xFFFF_FFFF_FFFF_FFFF, violating SingleGLSPerformanceAssumption. */
        assert ls.servers[id].held;
        assert ls.servers[id].epoch >= 0xFFFF_FFFF_FFFF_FFFF;
        assert |tgls.history| >= 0xFFFF_FFFF_FFFF_FFFF;
        assert false;
    }
}


lemma PerformanceGuaranteeHolds_Induction_NonIOStep(config:Config, tgls:TimestampedGLS_State, tgls':TimestampedGLS_State)
    // Standard pre-conditions
    requires |tgls.history| > 0 && |tgls'.history| > 0;
    requires CommonInvariants(config, tgls) && CommonInvariants(config, tgls');
    requires PerformanceInductiveInvariant(config, tgls);
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
    requires forall i | 0 <= i < |tglb| :: CommonInvariants(config, tglb[i]);
    requires forall i | 0 <= i < |tglb| :: PerformanceInductiveInvariant(config, tglb[i]);
    ensures GLSPerformanceGuarantee(tglb);
{
    forall i | 0 <= i < |tglb| 
    ensures SingleGLSPerformanceGuarantee(tglb[i])
    {
        assert TransferInvariant(tglb[i]);
    }
}
}