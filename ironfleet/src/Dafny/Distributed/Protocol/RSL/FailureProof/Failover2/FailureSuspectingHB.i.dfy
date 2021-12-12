include "../../DistributedSystem.i.dfy"

include "../../../../Common/Collections/Maps2.i.dfy"
include "../../Constants.i.dfy"
include "../../Environment.i.dfy"
include "../../Replica.i.dfy"
include "../Common/assumptions.i.dfy"

include "../TimestampedRslSystem.i.dfy"
include "FailureHelpers.i.dfy"

include "../../CommonProof/Constants.i.dfy"
// include "../CommonProof/Requests.i.dfy"

module FailureSuspectingHB_i {
import opened TimestampedRslSystem_i
import opened FailureHelpers_i
import opened Common_Assumptions

// FIXME: probably require RslConsistency or some such
predicate SusHBSpec(s:TimestampedRslState)
{
  forall pkt ::
    && pkt in s.t_environment.sentPackets
    && pkt.msg.v.RslMessage_Heartbeat?
    && pkt.msg.v.suspicious == true
    ==>
  (exists pkt' ::
    && pkt' in s.t_environment.sentPackets
    && pkt'.msg.v.RslMessage_Heartbeat?
    && pkt'.src == pkt.src
    && pkt'.dst == pkt.dst
    && pkt.msg.v.suspicious == true
    && TimeLe(pkt.msg.ts, TBFirstSuspectingHB())
  )
}

predicate BecomeSuspectorStepSpec(s:TimestampedRslState, s':TimestampedRslState)
{
  forall j :: 0 <= j < |s.t_replicas| ==>
  (s.t_replicas[j].v.replica.constants.my_index !in s.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors)
  ==>
  (s'.t_replicas[j].v.replica.constants.my_index in s'.t_replicas[j].v.replica.proposer.election_state.current_view_suspectors)
  ==>
  TimeLe(s'.t_replicas[j].ts, TBBecomeSuspector())
}

// TODO: this
lemma Part1TopLevel()
  requires true
  ensures true
{

}

}
