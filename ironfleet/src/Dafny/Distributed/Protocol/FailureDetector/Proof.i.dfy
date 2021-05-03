include "TimestampedDistributedSystem.i.dfy"

module FailureDetector_Proof {
import opened FailureDetector_TimestampedDistributedSystem

ghost const FailureTime:Timestamp
ghost const TimeoutInterval:Timestamp

predicate BoundedQueueingAssumption(s:TFD_State)
{
    && s.t_environment.nextStep.LEnvStepHostIos?
    ==>
  var ios := s.t_environment.nextStep.ios;
    (forall io :: io in s.t_environment.nextStep.ios && io.LIoOpReceive? && io.r.dst in s.t_servers ==>

      // this means that max(replica.ts, msg.ts) <= msg.ts + MaxQueueTime
      s.t_servers[io.r.dst].ts <= io.r.msg.ts + Q
    )
}

//
predicate Assumption(s:TFD_State)
{
  BoundedQueueingAssumption(s)

  && (forall i :: i in s.t_servers ==>
    TimeLe(s.t_servers[i].ts, FailureTime)
    && (s.t_servers[i].v.N? ==> s.t_servers[i].v.n.state == RUNNING)
    && (s.t_servers[i].v.D? ==> s.t_servers[i].v.d.timeoutInterval == TimeoutInterval)
    // FIXME: assume that RUNNING ==> ts <= FailureTime()
   )
}

predicate ClockAssumption(s:TFD_State)
{
  true
}

// inductive on its own
predicate BasicInvariant(s:TFD_State)
{
  && |s.t_servers| == 2
    && s.config.nodeEp in s.t_servers
    && s.config.detectorEp in s.t_servers
    && s.t_servers[s.config.detectorEp].v.D?
    && s.t_servers[s.config.nodeEp].v.N?
    && s.t_servers[s.config.detectorEp].v.D?
}

function LastHBDeliveryTime() : Timestamp
{
  FailureTime + MbeSend + Delay
}

function FailureDetectionTime() : Timestamp
{
  FailureTime + MbeSend + Delay + Q + TimeoutInterval
}

predicate Guarantee(s:TFD_State)
{
  && (forall pkt :: pkt in s.t_environment.sentPackets && pkt.msg.v.Alert? ==>
    TimeLe(pkt.msg.ts, LastHBDeliveryTime())
    )
}

datatype FailState =
  | NotFailed
  | Failed(pkt:TimestampedLPacket<EndPoint,FDMessage>)

predicate Invariant_aux(s:TFD_State, fs:FailState)
  requires BasicInvariant(s)
{
  && (forall pkt :: pkt in s.t_environment.sentPackets && pkt.msg.v.Heartbeat? ==>
    TimeLe(pkt.msg.ts, LastHBDeliveryTime())
    )
    && (
    if fs.Failed? then
    (fs.pkt in s.t_environment.sentPackets && fs.pkt.msg.v.Alert?
       && TimeLe(fs.pkt.msg.ts, FailureDetectionTime()))
    else
      (forall pkt :: pkt in s.t_environment.sentPackets ==> pkt.msg.v.Alert? ==> false)
      && s.t_servers[s.config.detectorEp].v.d.lastHeartbeatTime <= (LastHBDeliveryTime() + Q)
    )
}

lemma InvInductiveDetector_0(s:TFD_State, s':TFD_State, fs:FailState)
  requires Assumption(s)
  requires Assumption(s')
  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.config.detectorEp;
  requires s.t_environment.nextStep.nodeStep == DetectorStep(0)
  requires TFD_NextOneServer(s, s', s.t_environment.nextStep.actor, s.t_environment.nextStep.ios);
  requires BasicInvariant(s)
  requires BasicInvariant(s')

  requires Invariant_aux(s, fs)
  ensures Invariant_aux(s', fs)
{
  var ios := s.t_environment.nextStep.ios;
  if ios[0].LIoOpReceive? && ios[0].r.msg.v.Heartbeat? {
    var actor := s.config.detectorEp;
    var hstep := DetectorStep(0);
    assert ios[1].t <= s'.t_servers[s.config.detectorEp].ts;
    assert TimeLe(s'.t_servers[actor].ts, FD_RecvPerfUpdate(s.t_servers[actor].ts, ios[0].r.msg.ts, hstep));
    // assert FD_RecvPerfUpdate(s.t_servers[s.config.detectorEp].ts, ios[0].r.msg.ts, DetectorStep(0)) <= (LastHBDeliveryTime() + Q);
    assert s'.t_servers[s.config.detectorEp].ts <= (LastHBDeliveryTime()); // FIXME: this should fail...
    assert ios[1].t <= (LastHBDeliveryTime() + Q);
  } else {

  }
}

lemma InvInductiveDetector_1(s:TFD_State, s':TFD_State, fs:FailState) returns (fs':FailState)
  requires Assumption(s)
  requires Assumption(s')
  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == s.config.detectorEp;
  requires s.t_environment.nextStep.nodeStep == DetectorStep(1)
  requires TFD_NextOneServer(s, s', s.t_environment.nextStep.actor, s.t_environment.nextStep.ios);
  requires BasicInvariant(s)
  requires BasicInvariant(s')

  requires Invariant_aux(s, fs)
  ensures Invariant_aux(s', fs')
{
  var ios := s.t_environment.nextStep.ios;
  if fs.Failed? {
    fs' := fs;
  } else {
    var sentPackets := ExtractSentPacketsFromIos(UntagLIoOpSeq(ios));
    if |sentPackets| > 0 {
      var ios_u := UntagLIoOpSeq(ios);
      assert SpontaneousIos(ios_u, 1);
      assert sentPackets[0] in sentPackets; // XXX: Here to trigger ensures of ExtractSentPacketsFromIos
      assert 1 < |ios_u|;
      assert ios_u[1].LIoOpSend?;
      assert ios[1].LIoOpSend?;
      fs' := Failed(ios[1].s);
    } else {
      fs' := NotFailed;
      assert (forall pkt :: pkt in s'.t_environment.sentPackets ==> pkt.msg.v.Alert? ==> false);
      assert s.t_servers[s'.config.detectorEp].v.d.lastHeartbeatTime == s'.t_servers[s'.config.detectorEp].v.d.lastHeartbeatTime;
      //assert s'.t_servers[s'.config.detectorEp].v.d.lastHeartbeatTime <= LastHBDeliveryTime();
    }
  }
}


lemma {:verify false} InvInductive(s:TFD_State, s':TFD_State, actor:EndPoint, fs:FailState) returns (fs':FailState)
  requires Assumption(s)
  requires Assumption(s')
  requires s.t_environment.nextStep.LEnvStepHostIos?;
  requires s.t_environment.nextStep.actor == actor
  requires TFD_NextOneServer(s, s', s.t_environment.nextStep.actor, s.t_environment.nextStep.ios);
  requires BasicInvariant(s)
  requires BasicInvariant(s')

  requires Invariant_aux(s, fs)
  ensures Invariant_aux(s', fs')
{
  if fs.Failed? {
    fs' := fs;
    if actor == s.config.nodeEp {
      assert
        && (forall pkt :: pkt in s'.t_environment.sentPackets && pkt.msg.v.Heartbeat? ==>
        TimeLe(pkt.msg.ts, LastHBDeliveryTime())
        );
      assume false;
    } else {
      assume actor == s.config.detectorEp;
    }
  } else {
    assume false;
  }
}

}
