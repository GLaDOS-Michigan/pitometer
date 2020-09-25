include "../../../Common/Framework/Environment.s.dfy"
include "../../Lock/RefinementProof/DistributedSystem.i.dfy"

module Invariants_i {
    import opened DistributedSystem_i
    import opened Environment_s


/*****************************************************************************************
/                                 HistoryLengthInvariant                                 *
*****************************************************************************************/


predicate HistoryLengthInvariant(config:ConcreteConfiguration, gls:GLS_State)
    requires |gls.history| > 0;
    requires forall ep :: ep in gls.ls.servers <==> ep in config;
    requires forall ep | ep in gls.ls.servers :: gls.ls.servers[ep].config == config;
{
    && AtMostOneLockHolder(config, gls)
    && NoLockHolderImpliesOneTransferMessageInFlight(config, gls)
    && AllTransferMessageInFlightHaveDstLastInHistory(config, gls)
    && AllTransferMessageInFlightHaveEpochEqualsHistoryLength(config, gls)
    && SomeLockHolderImpliesNoTransferMessageInFlight(config, gls)
    && SomeLockHolderImpliesHolderIsLastInHistory(config, gls)
    && SomeLockHolderImpliesEpochIsHistoryLength(config, gls)
}



predicate AtMostOneLockHolder(config:ConcreteConfiguration, gls:GLS_State) 
    requires forall ep :: ep in gls.ls.servers <==> ep in config;
    requires forall ep | ep in gls.ls.servers :: gls.ls.servers[ep].config == config;
{
    forall ep1, ep2 | ep1 in config && ep2 in config :: gls.ls.servers[ep1].held && gls.ls.servers[ep2].held ==> ep1 == ep2
}   

predicate NoLockHolder(config:ConcreteConfiguration, gls:GLS_State) 
    requires forall ep :: ep in gls.ls.servers <==> ep in config;
    requires forall ep | ep in gls.ls.servers :: gls.ls.servers[ep].config == config;
{
    forall ep | ep in config :: !gls.ls.servers[ep].held
}


predicate IsInFlightTransferMessage(config:ConcreteConfiguration, gls:GLS_State, p:LPacket<EndPoint, LockMessage>) 
    requires forall ep :: ep in gls.ls.servers <==> ep in config;
    requires forall ep | ep in gls.ls.servers :: gls.ls.servers[ep].config == config;
{
    && p in gls.ls.environment.sentPackets  
    && p.dst in config 
    && p.src in config 
    && p.msg.Transfer?
    && p.msg.transfer_epoch > gls.ls.servers[p.dst].epoch
}

predicate OneTransferMessageIsInFlight(config:ConcreteConfiguration, gls:GLS_State) 
    requires forall ep :: ep in gls.ls.servers <==> ep in config;
    requires forall ep | ep in gls.ls.servers :: gls.ls.servers[ep].config == config;
{
    forall p1, p2 | IsInFlightTransferMessage(config, gls, p1) && IsInFlightTransferMessage(config, gls, p2) :: p1 == p2
}

predicate ZeroTransferMessageIsInFlight(config:ConcreteConfiguration, gls:GLS_State) 
    requires forall ep :: ep in gls.ls.servers <==> ep in config;
    requires forall ep | ep in gls.ls.servers :: gls.ls.servers[ep].config == config;
{
    forall p | p in gls.ls.environment.sentPackets  :: ! IsInFlightTransferMessage(config, gls, p)
}

predicate NoLockHolderImpliesOneTransferMessageInFlight(config:ConcreteConfiguration, gls:GLS_State) 
    requires forall ep :: ep in gls.ls.servers <==> ep in config;
    requires forall ep | ep in gls.ls.servers :: gls.ls.servers[ep].config == config;
{
    NoLockHolder(config, gls) ==> OneTransferMessageIsInFlight(config, gls)
}   

predicate AllTransferMessageInFlightHaveDstLastInHistory(config:ConcreteConfiguration, gls:GLS_State)
    requires |gls.history| > 0;
    requires forall ep :: ep in gls.ls.servers <==> ep in config;
    requires forall ep | ep in gls.ls.servers :: gls.ls.servers[ep].config == config;
{
    forall p | IsInFlightTransferMessage(config, gls, p) :: p.dst == gls.history[|gls.history|-1]
}   

predicate AllTransferMessageInFlightHaveEpochEqualsHistoryLength(config:ConcreteConfiguration, gls:GLS_State)
    requires |gls.history| > 0;
    requires forall ep :: ep in gls.ls.servers <==> ep in config;
    requires forall ep | ep in gls.ls.servers :: gls.ls.servers[ep].config == config;
{
    forall p | IsInFlightTransferMessage(config, gls, p) :: p.msg.transfer_epoch == |gls.history|
}   

predicate SomeLockHolderImpliesNoTransferMessageInFlight(config:ConcreteConfiguration, gls:GLS_State) 
    requires forall ep :: ep in gls.ls.servers <==> ep in config;
    requires forall ep | ep in gls.ls.servers :: gls.ls.servers[ep].config == config;
{
    !NoLockHolder(config, gls) ==> ZeroTransferMessageIsInFlight(config, gls)
}   

predicate SomeLockHolderImpliesEpochIsHistoryLength(config:ConcreteConfiguration, gls:GLS_State) 
    requires forall ep :: ep in gls.ls.servers <==> ep in config;
    requires forall ep | ep in gls.ls.servers :: gls.ls.servers[ep].config == config;
{
    forall h | h in config && gls.ls.servers[h].held :: gls.ls.servers[h].epoch == |gls.history|
}   

predicate SomeLockHolderImpliesHolderIsLastInHistory(config:ConcreteConfiguration, gls:GLS_State)
    requires |gls.history| > 0;
    requires forall ep :: ep in gls.ls.servers <==> ep in config;
    requires forall ep | ep in gls.ls.servers :: gls.ls.servers[ep].config == config;
{
    forall ep | ep in config && gls.ls.servers[ep].held :: gls.history[|gls.history|-1] == ep
}

lemma lemma_History_Length_Invariant(config:ConcreteConfiguration, glb:seq<GLS_State>)
    requires |glb| > 0;
    requires  GLS_Init(glb[0], config);
    requires forall i {:trigger GLS_Next(glb[i], glb[i+1])} :: 0 <= i < |glb| - 1 ==>  GLS_Next(glb[i], glb[i+1]);
    requires forall i | 0 <= i < |glb| :: (forall ep :: ep in glb[i].ls.servers <==> ep in config);
    requires forall i | 0 <= i < |glb| :: (forall ep | ep in glb[i].ls.servers :: glb[i].ls.servers[ep].config == config);
    ensures forall i | 0 <= i < |glb| :: SomeLockHolderImpliesEpochIsHistoryLength(config, glb[i]);
    ensures forall i | 0 <= i < |glb| :: |glb[i].history| > 0;
    ensures forall i | 0 <= i < |glb| :: HistoryLengthInvariant(config, glb[i]);
{
    // Base Case
    assert glb[0].history == [config[0]];
    assert glb[0].ls.servers[config[0]].epoch == 1;
    assert HistoryLengthInvariant(config, glb[0]);

    // Inductive Case
    var i := 1;
    while i < |glb| 
        decreases |glb| - i;
        invariant 1 <= i <= |glb|;
        invariant forall k | 0 <= k < |glb|-1 :: GLS_Next(glb[k], glb[k+1]);
        invariant forall k | 0 <= k < i :: |glb[k].history| > 0;
        invariant forall k | 0 <= k < i :: HistoryLengthInvariant(config, glb[k]);
    {
        var k := i-1;
        var gls, gls' := glb[k], glb[k+1];
        assert HistoryLengthInvariant(config, gls);

        if !NoLockHolder(config, gls) {
            // If someone holds the lock in state gls. 
            // Then I know that there are no Transfer messages in flight in gls
            lemma_History_Length_Invariant_Induction_A(config, gls, gls');
        } else {
            // If no one holds the lock in state gls. 
            // Then I know that there is only one Transfer message in flight in gls
            assert NoLockHolder(config, gls);
            lemma_History_Length_Invariant_Induction_B(config, gls, gls');
        }
        i := i + 1; 
    }
}

/* First branch of lemma_History_Length_Invariant proof, 
* where !NoLockHolder(config, gls); */
lemma lemma_History_Length_Invariant_Induction_A(config:ConcreteConfiguration, gls:GLS_State, gls':GLS_State)
    requires forall ep :: ep in gls.ls.servers <==> ep in config;
    requires forall ep | ep in gls.ls.servers :: gls.ls.servers[ep].config == config;
    requires GLS_Next(gls, gls');
    requires |gls.history| > 0
    requires HistoryLengthInvariant(config, gls);
    requires !NoLockHolder(config, gls);
    ensures |gls'.history| > 0
    ensures HistoryLengthInvariant(config, gls');
{
    assert ZeroTransferMessageIsInFlight(config, gls);
    assert LS_Next(gls.ls, gls'.ls);

    if (&& gls.ls.environment.nextStep.LEnvStepHostIos? 
        && gls.ls.environment.nextStep.actor in gls.ls.servers) {
        // If gls->gls' is a node Grant or Accept step
        var id, ios := gls.ls.environment.nextStep.actor, gls.ls.environment.nextStep.ios;
        assert NodeNext(gls.ls.servers[id], gls'.ls.servers[id], ios);

        if NodeGrant(gls.ls.servers[id], gls'.ls.servers[id], ios) {
            assert HistoryLengthInvariant(config, gls');  // Dafny magic!
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
        assert HistoryLengthInvariant(config, gls');
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
        assert HistoryLengthInvariant(config, gls');
    }
}


/* Second branch of lemma_History_Length_Invariant proof, 
* where NoLockHolder(config, gls); */
lemma lemma_History_Length_Invariant_Induction_B(config:ConcreteConfiguration, gls:GLS_State, gls':GLS_State)
    requires forall ep :: ep in gls.ls.servers <==> ep in config;
    requires forall ep | ep in gls.ls.servers :: gls.ls.servers[ep].config == config;
    requires GLS_Next(gls, gls');
    requires |gls.history| > 0
    requires HistoryLengthInvariant(config, gls);
    requires NoLockHolder(config, gls);
    ensures |gls'.history| > 0
    ensures HistoryLengthInvariant(config, gls');
{
    assert OneTransferMessageIsInFlight(config, gls);
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
                    assert !NoLockHolder(config, gls');
                    assert forall id' | id' in config && id' != id :: !gls.ls.servers[id'].held;
                    forall p | p in gls'.ls.environment.sentPackets  
                    ensures !IsInFlightTransferMessage(config, gls', p)
                    {
                        if p.dst in config && p.src in config && p.msg.Transfer? {
                            assert p in gls.ls.environment.sentPackets;
                            if p.dst == id {
                                if p.msg.transfer_epoch > gls.ls.servers[id].epoch {
                                    var receivedMessage := ios[0].r;
                                    assert IsValidLIoOp(ios[0], id, gls.ls.environment);
                                    assert receivedMessage.dst == id;
                                    assert IsInFlightTransferMessage(config, gls, receivedMessage);
                                    assert IsInFlightTransferMessage(config, gls, p);
                                    assert p == receivedMessage;
                                    assert p.msg.transfer_epoch == gls'.ls.servers[id].epoch;
                                    assert !IsInFlightTransferMessage(config, gls', p);
                                } else {
                                    assert !IsInFlightTransferMessage(config, gls', p);
                                }
                            } else {
                                assert gls'.ls.servers[p.dst].epoch == gls.ls.servers[p.dst].epoch;
                                assert gls.ls.servers[p.dst].epoch >= p.msg.transfer_epoch;
                            }
                        }
                    }
                    assert HistoryLengthInvariant(config, gls');  // Dafny magic!
                } else {
                    assert gls'.ls.servers == gls.ls.servers;
                    assert gls'.ls.environment.sentPackets == gls.ls.environment.sentPackets;
                }
            }
        }
        assert HistoryLengthInvariant(config, gls');
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
                assert OneTransferMessageIsInFlight(config, gls);
                var id, ios := gls.ls.environment.nextStep.actor, gls.ls.environment.nextStep.ios;
                forall io | io in ios && io.LIoOpSend?
                ensures !IsInFlightTransferMessage(config, gls', io.s)
                {}
                assert gls'.ls.environment.sentPackets == gls.ls.environment.sentPackets + (set io | io in ios && io.LIoOpSend? :: io.s);
                assert forall p | p in gls'.ls.environment.sentPackets && IsInFlightTransferMessage(config, gls', p) :: p in gls.ls.environment.sentPackets;
                assert NoLockHolder(config, gls');
                assert OneTransferMessageIsInFlight(config, gls');
            } else {
                assert gls'.ls.environment.sentPackets == gls.ls.environment.sentPackets;
            }
        }
        assert HistoryLengthInvariant(config, gls');
    }
}
}
