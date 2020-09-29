include "../../../Common/Framework/Environment.s.dfy"
include "TimestampedGLS.i.dfy"

module Invariants_i {
    import opened TimestampedGLS_i
    import opened Environment_s


/*****************************************************************************************
/                           ValidBehavior translation lemma                              *
*****************************************************************************************/

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


/*****************************************************************************************
/                                      EpochInvariant                                    *
*****************************************************************************************/


predicate {:opaque} EpochInvariant(config:Config, tgls:TimestampedGLS_State) 
    requires ConfigInvariant(config, tgls);
{   && (forall ep | ep in config :: tgls.tls.t_servers[ep].v.epoch >= 0)
    && (forall pkt  
        | && pkt in tgls.tls.t_environment.sentPackets 
          && pkt.src in config
          && pkt.dst in config
          :: 
            pkt.msg.v.Locked? || pkt.msg.v.Transfer?)
    && (forall pkt  
        | && pkt in tgls.tls.t_environment.sentPackets 
          && pkt.msg.v.Transfer?
          && pkt.src in config
          && pkt.dst in config
        :: 
             pkt.msg.v.transfer_epoch > 0
        )
    && (forall pkt  
        | && pkt in tgls.tls.t_environment.sentPackets 
          && pkt.msg.v.Locked?
          && pkt.src in config
          && pkt.dst in config
        :: 
            tgls.tls.t_servers[pkt.dst].v.epoch > 0)
}


lemma lemma_EpochInvariant(config:Config, tglb:seq<TimestampedGLS_State>) 
    requires |config| > 1;
    requires ValidTimestampedGLSBehavior(tglb, config);
    requires forall i | 0 <= i < |tglb| :: ConfigInvariant(config, tglb[i]);
    ensures forall i | 0 <= i < |tglb| :: EpochInvariant(config, tglb[i]);
{
    // TONY TODO
    assume false;
}



/*****************************************************************************************
/                                 ConfigurationInvariant                                 *
*****************************************************************************************/


/* Invariants about servers in gls and the config */
predicate ConfigInvariant(config:ConcreteConfiguration, tgls:TimestampedGLS_State) 
{
    && (forall ep :: ep in tgls.tls.t_servers <==> ep in config)
    && (forall ep | ep in tgls.tls.t_servers :: tgls.tls.t_servers[ep].v.config == config)
}

lemma lemma_ConfigInvariant(config:ConcreteConfiguration, tglb:seq<TimestampedGLS_State>) 
    requires |config| > 1;
    requires ValidTimestampedGLSBehavior(tglb, config)
    ensures forall i | 0 <= i < |tglb| :: ConfigInvariant(config, tglb[i]);
{
    lemma_ValidBehavior(config, tglb);
    var ls := UntagLS_State(tglb[0].tls);
    assert LS_Init(ls, config);
    assert forall e :: e in config <==> e in ls.servers;
    assert forall ep | ep in ls.servers :: ls.servers[ep].config == config;
    assert ConfigInvariant(config, tglb[0]);

    var i := 0;
    while i < |tglb| - 1
        decreases |tglb| - i;
        invariant 0 <= i < |tglb|;
        invariant forall k | 0 <= k <= i :: ConfigInvariant(config, tglb[k]);
    {
        var tgls, tgls' := tglb[i], tglb[i+1];
        assert TGLS_Next(tgls, tgls');
        assert forall ep :: ep in tgls'.tls.t_servers <==> ep in config;
        assert forall ep | ep in tgls'.tls.t_servers :: tgls'.tls.t_servers[ep].v.config == config;
        assert ConfigInvariant(config, tgls');
        i := i + 1;    
    }
}


/*****************************************************************************************
/                                 HistoryLengthInvariant                                 *
*****************************************************************************************/


predicate HistoryLengthInvariant(config:ConcreteConfiguration, tgls:TimestampedGLS_State)
    requires |tgls.history| > 0;
    requires forall ep :: ep in tgls.tls.t_servers <==> ep in config;
    requires forall ep | ep in tgls.tls.t_servers :: tgls.tls.t_servers[ep].v.config == config;
{
    && AtMostOneLockHolder(config, tgls)
    && NoLockHolderImpliesOneTransferMessageInFlight(config, tgls)
    && AllTransferMessageInFlightHaveDstLastInHistory(config, tgls)
    && AllTransferMessageInFlightHaveEpochEqualsHistoryLength(config, tgls)
    && SomeLockHolderImpliesNoTransferMessageInFlight(config, tgls)
    && SomeLockHolderImpliesHolderIsLastInHistory(config, tgls)
    && SomeLockHolderImpliesEpochIsHistoryLength(config, tgls)
}



predicate AtMostOneLockHolder(config:ConcreteConfiguration, gls:TimestampedGLS_State) 
    requires forall ep :: ep in gls.tls.t_servers <==> ep in config;
    requires forall ep | ep in gls.tls.t_servers :: gls.tls.t_servers[ep].v.config == config;
{
    forall ep1, ep2 | ep1 in config && ep2 in config :: gls.tls.t_servers[ep1].v.held && gls.tls.t_servers[ep2].v.held ==> ep1 == ep2
}   

predicate NoLockHolder(config:ConcreteConfiguration, gls:TimestampedGLS_State) 
    requires forall ep :: ep in gls.tls.t_servers <==> ep in config;
    requires forall ep | ep in gls.tls.t_servers :: gls.tls.t_servers[ep].v.config == config;
{
    forall ep | ep in config :: !gls.tls.t_servers[ep].v.held
}


predicate IsInFlightTransferMessage(config:ConcreteConfiguration, gls:TimestampedGLS_State, p:LPacket<EndPoint, LockMessage>) 
    requires forall ep :: ep in gls.tls.t_servers <==> ep in config;
    requires forall ep | ep in gls.tls.t_servers :: gls.tls.t_servers[ep].v.config == config;
{
    && p in UntagSentPkts(gls.tls.t_environment.sentPackets)
    && p.dst in config 
    && p.src in config 
    && p.msg.Transfer?
    && p.msg.transfer_epoch > gls.tls.t_servers[p.dst].v.epoch
}

predicate OneTransferMessageIsInFlight(config:ConcreteConfiguration, gls:TimestampedGLS_State) 
    requires forall ep :: ep in gls.tls.t_servers <==> ep in config;
    requires forall ep | ep in gls.tls.t_servers :: gls.tls.t_servers[ep].v.config == config;
{
    forall p1, p2 | IsInFlightTransferMessage(config, gls, p1) && IsInFlightTransferMessage(config, gls, p2) :: p1 == p2
}

predicate ZeroTransferMessageIsInFlight(config:ConcreteConfiguration, gls:TimestampedGLS_State) 
    requires forall ep :: ep in gls.tls.t_servers <==> ep in config;
    requires forall ep | ep in gls.tls.t_servers :: gls.tls.t_servers[ep].v.config == config;
{
    forall p | p in UntagSentPkts(gls.tls.t_environment.sentPackets)  :: ! IsInFlightTransferMessage(config, gls, p)
}

predicate NoLockHolderImpliesOneTransferMessageInFlight(config:ConcreteConfiguration, gls:TimestampedGLS_State) 
    requires forall ep :: ep in gls.tls.t_servers <==> ep in config;
    requires forall ep | ep in gls.tls.t_servers :: gls.tls.t_servers[ep].v.config == config;
{
    NoLockHolder(config, gls) ==> OneTransferMessageIsInFlight(config, gls)
}   

predicate AllTransferMessageInFlightHaveDstLastInHistory(config:ConcreteConfiguration, gls:TimestampedGLS_State)
    requires |gls.history| > 0;
    requires forall ep :: ep in gls.tls.t_servers <==> ep in config;
    requires forall ep | ep in gls.tls.t_servers :: gls.tls.t_servers[ep].v.config == config;
{
    forall p | IsInFlightTransferMessage(config, gls, p) :: p.dst == gls.history[|gls.history|-1]
}   

predicate AllTransferMessageInFlightHaveEpochEqualsHistoryLength(config:ConcreteConfiguration, gls:TimestampedGLS_State)
    requires |gls.history| > 0;
    requires forall ep :: ep in gls.tls.t_servers <==> ep in config;
    requires forall ep | ep in gls.tls.t_servers :: gls.tls.t_servers[ep].v.config == config;
{
    forall p | IsInFlightTransferMessage(config, gls, p) :: p.msg.transfer_epoch == |gls.history|
}   

predicate SomeLockHolderImpliesNoTransferMessageInFlight(config:ConcreteConfiguration, gls:TimestampedGLS_State) 
    requires forall ep :: ep in gls.tls.t_servers <==> ep in config;
    requires forall ep | ep in gls.tls.t_servers :: gls.tls.t_servers[ep].v.config == config;
{
    !NoLockHolder(config, gls) ==> ZeroTransferMessageIsInFlight(config, gls)
}   

predicate SomeLockHolderImpliesEpochIsHistoryLength(config:ConcreteConfiguration, gls:TimestampedGLS_State) 
    requires forall ep :: ep in gls.tls.t_servers <==> ep in config;
    requires forall ep | ep in gls.tls.t_servers :: gls.tls.t_servers[ep].v.config == config;
{
    forall h | h in config && gls.tls.t_servers[h].v.held :: gls.tls.t_servers[h].v.epoch == |gls.history|
}   

predicate SomeLockHolderImpliesHolderIsLastInHistory(config:ConcreteConfiguration, gls:TimestampedGLS_State)
    requires |gls.history| > 0;
    requires forall ep :: ep in gls.tls.t_servers <==> ep in config;
    requires forall ep | ep in gls.tls.t_servers :: gls.tls.t_servers[ep].v.config == config;
{
    forall ep | ep in config && gls.tls.t_servers[ep].v.held :: gls.history[|gls.history|-1] == ep
}


/* Proof of HistoryLengthInvariant */
lemma lemma_History_Length_Invariant(config:ConcreteConfiguration, tglb:seq<TimestampedGLS_State>)
    requires |tglb| > 0;
    requires  TGLS_Init(tglb[0], config);
    requires forall i {:trigger TGLS_Next(tglb[i], tglb[i+1])} :: 0 <= i < |tglb| - 1 ==>  TGLS_Next(tglb[i], tglb[i+1]);
    requires forall i | 0 <= i < |tglb| :: (forall ep :: ep in tglb[i].tls.t_servers <==> ep in config);
    requires forall i | 0 <= i < |tglb| :: (forall ep | ep in tglb[i].tls.t_servers :: tglb[i].tls.t_servers[ep].v.config == config);
    ensures forall i | 0 <= i < |tglb| :: SomeLockHolderImpliesEpochIsHistoryLength(config, tglb[i]);
    ensures forall i | 0 <= i < |tglb| :: |tglb[i].history| > 0;
    ensures forall i | 0 <= i < |tglb| :: HistoryLengthInvariant(config, tglb[i]);
{
    // Base Case
    assert tglb[0].history == [config[0]];
    assert tglb[0].tls.t_servers[config[0]].v.epoch == 1;
    assert HistoryLengthInvariant(config, tglb[0]);

    // Inductive Case
    var i := 1;
    while i < |tglb| 
        decreases |tglb| - i;
        invariant 1 <= i <= |tglb|;
        invariant forall k | 0 <= k < |tglb|-1 :: TGLS_Next(tglb[k], tglb[k+1]);
        invariant forall k | 0 <= k < i :: |tglb[k].history| > 0;
        invariant forall k | 0 <= k < i :: HistoryLengthInvariant(config, tglb[k]);
    {
        var k := i-1;
        var tgls, tgls' := tglb[k], tglb[k+1];
        assert HistoryLengthInvariant(config, tgls);

        if !NoLockHolder(config, tgls) {
            // If someone holds the lock in state gls. 
            // Then I know that there are no Transfer messages in flight in gls
            lemma_History_Length_Invariant_Induction_A(config, tgls, tgls');
        } else {
            // If no one holds the lock in state gls. 
            // Then I know that there is only one Transfer message in flight in gls
            assert NoLockHolder(config, tgls);
            lemma_History_Length_Invariant_Induction_B(config, tgls, tgls');
        }
        assert HistoryLengthInvariant(config, tgls');
        forall k | 0 <= k <= i 
        ensures HistoryLengthInvariant(config, tglb[k])
        {}
        forall k | 0 <= k <= i 
        ensures |tglb[k].history| > 0
        {}
        i := i + 1; 
    }
}

/* First branch of lemma_History_Length_Invariant proof, 
* where !NoLockHolder(config, gls); */
lemma lemma_History_Length_Invariant_Induction_A(config:ConcreteConfiguration, tgls:TimestampedGLS_State, tgls':TimestampedGLS_State)
    requires forall ep :: ep in tgls.tls.t_servers <==> ep in config;
    requires forall ep | ep in tgls.tls.t_servers :: tgls.tls.t_servers[ep].v.config == config;
    requires TGLS_Next(tgls, tgls');
    requires |tgls.history| > 0
    requires HistoryLengthInvariant(config, tgls);
    requires !NoLockHolder(config, tgls);
    ensures |tgls'.history| > 0
    ensures HistoryLengthInvariant(config, tgls');
{
    assert ZeroTransferMessageIsInFlight(config, tgls);
    var gls, gls' := UntagGLS_State(tgls), UntagGLS_State(tgls');
    assert LS_Next(gls.ls, gls'.ls);

    if (&& gls.ls.environment.nextStep.LEnvStepHostIos? 
        && gls.ls.environment.nextStep.actor in gls.ls.servers) {
        // If gls->gls' is a node Grant or Accept step
        var id, ios := gls.ls.environment.nextStep.actor, gls.ls.environment.nextStep.ios;
        assert NodeNext(gls.ls.servers[id], gls'.ls.servers[id], ios);

        if NodeGrant(gls.ls.servers[id], gls'.ls.servers[id], ios) {
            assert HistoryLengthInvariant(config, tgls');  // Dafny magic!
        } else {
            assert NodeAccept(gls.ls.servers[id], gls'.ls.servers[id], ios);
            if ios[0].LIoOpTimeoutReceive? {
                assert gls'.ls.servers == gls.ls.servers;
                assert gls'.ls.environment.sentPackets == gls.ls.environment.sentPackets;
            } else {
                if (&& !gls.ls.servers[id].held
                    && ios[0].r.src in gls.ls.servers[id].config
                    && ios[0].r.msg.Transfer? 
                    && ios[0].r.msg.transfer_epoch > gls.ls.servers[id].epoch
                ) {
                    assert ios[0].LIoOpReceive?;
                    assert IsValidLIoOp(ios[0], id, gls.ls.environment);
                    var m := ios[0].r;
                    assert m in gls.ls.environment.sentPackets;
                    assert m.dst == id; 
                    assert false; // contradicts ZeroTransferMessageIsInFlight(config, gls)
                } else {
                    assert gls'.ls.servers == gls.ls.servers;
                    assert gls'.ls.environment.sentPackets == gls.ls.environment.sentPackets;
                }
            }
        }
        assert HistoryLengthInvariant(config, tgls');
    } else {
        // If gls->gls' is purely an environment step
        assert gls.ls.servers == gls'.ls.servers && gls.history == gls'.history;
        if !gls.ls.environment.nextStep.LEnvStepHostIos? {
            match gls.ls.environment.nextStep {
                case LEnvStepHostIos(actor, ios, nodestep) => assert false;
                case LEnvStepDeliverPacket(p) => assert LEnvironment_Stutter(gls.ls.environment, gls'.ls.environment);
                case LEnvStepAdvanceTime => assert LEnvironment_AdvanceTime(gls.ls.environment, gls'.ls.environment);
                case LEnvStepStutter => assert LEnvironment_Stutter(gls.ls.environment, gls'.ls.environment);
            }
            assert gls'.ls.environment.sentPackets == gls.ls.environment.sentPackets;
        } else {
            assert gls.ls.environment.nextStep.actor !in gls.ls.servers;
        }
        assert HistoryLengthInvariant(config, tgls');
    }
}


/* Second branch of lemma_History_Length_Invariant proof, 
* where NoLockHolder(config, gls); */
lemma lemma_History_Length_Invariant_Induction_B(config:ConcreteConfiguration, tgls:TimestampedGLS_State, tgls':TimestampedGLS_State)
    requires forall ep :: ep in tgls.tls.t_servers <==> ep in config;
    requires forall ep | ep in tgls.tls.t_servers :: tgls.tls.t_servers[ep].v.config == config;
    requires TGLS_Next(tgls, tgls');
    requires |tgls.history| > 0
    requires HistoryLengthInvariant(config, tgls);
    requires NoLockHolder(config, tgls);
    ensures |tgls'.history| > 0
    ensures HistoryLengthInvariant(config, tgls');
{
    assert OneTransferMessageIsInFlight(config, tgls);
    var gls, gls' := UntagGLS_State(tgls), UntagGLS_State(tgls');
    assert LS_Next(gls.ls, gls'.ls);
    if (&& gls.ls.environment.nextStep.LEnvStepHostIos? 
        && gls.ls.environment.nextStep.actor in gls.ls.servers) {
        // If gls->gls' is a node Grant or Accept step
        var id, ios := gls.ls.environment.nextStep.actor, gls.ls.environment.nextStep.ios;
        assert id in config;
        assert NodeNext(gls.ls.servers[id], gls'.ls.servers[id], ios);

        if NodeGrant(gls.ls.servers[id], gls'.ls.servers[id], ios) {
            if gls.ls.servers[id].held && gls.ls.servers[id].epoch < 0xFFFF_FFFF_FFFF_FFFF {
                assert false;  // contradics NoLockHolder(config, gls);
            } else {
                assert gls'.ls.servers == gls.ls.servers;
                assert gls'.ls.environment.sentPackets == gls.ls.environment.sentPackets;
            }
        } else {
            assert NodeAccept(gls.ls.servers[id], gls'.ls.servers[id], ios);
            if ios[0].LIoOpTimeoutReceive? {
                assert gls'.ls.servers == gls.ls.servers;
                assert gls'.ls.environment.sentPackets == gls.ls.environment.sentPackets;
            } else {
                if (!gls.ls.servers[id].held 
                    && ios[0].r.src in gls.ls.servers[id].config
                    && ios[0].r.msg.Transfer? 
                    && ios[0].r.msg.transfer_epoch > gls.ls.servers[id].epoch)
                {
                    assert gls'.ls.servers[id].held;
                    assert !NoLockHolder(config, tgls');
                    assert forall id' | id' in config && id' != id :: !gls.ls.servers[id'].held;
                    forall p | p in gls'.ls.environment.sentPackets  
                    ensures !IsInFlightTransferMessage(config, tgls', p)
                    {
                        if p.dst in config && p.src in config && p.msg.Transfer? {
                            assert p in gls.ls.environment.sentPackets;
                            if p.dst == id {
                                if p.msg.transfer_epoch > gls.ls.servers[id].epoch {
                                    var receivedMessage := ios[0].r;
                                    assert IsValidLIoOp(ios[0], id, gls.ls.environment);
                                    assert receivedMessage.dst == id;
                                    assert IsInFlightTransferMessage(config, tgls, receivedMessage);
                                    assert IsInFlightTransferMessage(config, tgls, p);
                                    assert p == receivedMessage;
                                    assert p.msg.transfer_epoch == gls'.ls.servers[id].epoch;
                                    assert !IsInFlightTransferMessage(config, tgls', p);
                                } else {
                                    assert !IsInFlightTransferMessage(config, tgls', p);
                                }
                            } else {
                                assert gls'.ls.servers[p.dst].epoch == gls.ls.servers[p.dst].epoch;
                                assert gls.ls.servers[p.dst].epoch >= p.msg.transfer_epoch;
                            }
                        }
                    }
                    assert HistoryLengthInvariant(config, tgls');  // Dafny magic!
                } else {
                    assert gls'.ls.servers == gls.ls.servers;
                    assert gls'.ls.environment.sentPackets == gls.ls.environment.sentPackets;
                }
            }
        }
        assert HistoryLengthInvariant(config, tgls');
    } else {
        // If gls->gls' is purely an environment step
        // (not gls.ls.environment.nextStep.LEnvStepHostIos?) or (not gls.ls.environment.nextStep.actor in gls.ls.servers)
        assert gls.ls.servers == gls'.ls.servers && gls.history == gls'.history;
        if !gls.ls.environment.nextStep.LEnvStepHostIos? {
            match gls.ls.environment.nextStep {
                case LEnvStepHostIos(actor, ios, nodestep) => assert false;
                case LEnvStepDeliverPacket(p) => assert LEnvironment_Stutter(gls.ls.environment, gls'.ls.environment);
                case LEnvStepAdvanceTime => assert LEnvironment_AdvanceTime(gls.ls.environment, gls'.ls.environment);
                case LEnvStepStutter => assert LEnvironment_Stutter(gls.ls.environment, gls'.ls.environment);
            }
            assert gls'.ls.environment.sentPackets == gls.ls.environment.sentPackets;
        } else {
            assert gls.ls.environment.nextStep.actor !in gls.ls.servers;
            if gls.ls.environment.nextStep.LEnvStepHostIos? {  
                // Any message sent in this step cannot be classified as InFlight 
                // because their source == actor is not in config
                assert OneTransferMessageIsInFlight(config, tgls);
                var id, ios := gls.ls.environment.nextStep.actor, gls.ls.environment.nextStep.ios;
                forall io | io in ios && io.LIoOpSend?
                ensures !IsInFlightTransferMessage(config, tgls', io.s)
                {}
                assert gls'.ls.environment.sentPackets == gls.ls.environment.sentPackets + (set io | io in ios && io.LIoOpSend? :: io.s);
                assert forall p | p in gls'.ls.environment.sentPackets && IsInFlightTransferMessage(config, tgls', p) :: p in gls.ls.environment.sentPackets;
                assert NoLockHolder(config, tgls');
                assert OneTransferMessageIsInFlight(config, tgls');
            } else {
                assert gls'.ls.environment.sentPackets == gls.ls.environment.sentPackets;
            }
        }
        assert HistoryLengthInvariant(config, tgls');
    }
}
}
