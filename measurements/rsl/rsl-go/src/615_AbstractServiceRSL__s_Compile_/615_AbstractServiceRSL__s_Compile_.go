// Package _615_AbstractServiceRSL__s_Compile
// Dafny module _615_AbstractServiceRSL__s_Compile compiled into Go

package _615_AbstractServiceRSL__s_Compile

import (
  _dafny "dafny"
_System "System_"
_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
_5_Temporal____Temporal__s_Compile "5_Temporal____Temporal__s_Compile_"
_7_Environment__s_Compile "7_Environment__s_Compile_"
_9_Native____Io__s_Compile "9_Native____Io__s_Compile_"
_26_Collections____Seqs__s_Compile "26_Collections____Seqs__s_Compile_"
_30_Collections____Seqs__i_Compile "30_Collections____Seqs__i_Compile_"
_32_Math____mul__nonlinear__i_Compile "32_Math____mul__nonlinear__i_Compile_"
_34_Math____mul__auto__proofs__i_Compile "34_Math____mul__auto__proofs__i_Compile_"
_36_Math____mul__auto__i_Compile "36_Math____mul__auto__i_Compile_"
_40_Math____mul__i_Compile "40_Math____mul__i_Compile_"
_42_Math____div__nonlinear__i_Compile "42_Math____div__nonlinear__i_Compile_"
_44_Math____mod__auto__proofs__i_Compile "44_Math____mod__auto__proofs__i_Compile_"
_46_Math____mod__auto__i_Compile "46_Math____mod__auto__i_Compile_"
_50_AppStateMachine__i_Compile "50_AppStateMachine__i_Compile_"
_54_Concrete__NodeIdentity__i_Compile "54_Concrete__NodeIdentity__i_Compile_"
_56_LiveRSL____Types__i_Compile "56_LiveRSL____Types__i_Compile_"
_58_Collections____Sets__i_Compile "58_Collections____Sets__i_Compile_"
_61_LiveRSL____Configuration__i_Compile "61_LiveRSL____Configuration__i_Compile_"
_65_LiveRSL____Message__i_Compile "65_LiveRSL____Message__i_Compile_"
_68_LiveRSL____Environment__i_Compile "68_LiveRSL____Environment__i_Compile_"
_71_LiveRSL____ClockReading__i_Compile "71_LiveRSL____ClockReading__i_Compile_"
_74_Common____UpperBound__s_Compile "74_Common____UpperBound__s_Compile_"
_76_LiveRSL____Parameters__i_Compile "76_LiveRSL____Parameters__i_Compile_"
_78_LiveRSL____Constants__i_Compile "78_LiveRSL____Constants__i_Compile_"
_85_LiveRSL____Broadcast__i_Compile "85_LiveRSL____Broadcast__i_Compile_"
_91_Collections____CountMatches__i_Compile "91_Collections____CountMatches__i_Compile_"
_93_LiveRSL____Acceptor__i_Compile "93_LiveRSL____Acceptor__i_Compile_"
_99_LiveRSL____Election__i_Compile "99_LiveRSL____Election__i_Compile_"
_101_LiveRSL____Proposer__i_Compile "101_LiveRSL____Proposer__i_Compile_"
_115_LiveRSL____StateMachine__i_Compile "115_LiveRSL____StateMachine__i_Compile_"
_118_Collections____Maps__i_Compile "118_Collections____Maps__i_Compile_"
_120_LiveRSL____Executor__i_Compile "120_LiveRSL____Executor__i_Compile_"
_123_LiveRSL____Learner__i_Compile "123_LiveRSL____Learner__i_Compile_"
_126_LiveRSL____Replica__i_Compile "126_LiveRSL____Replica__i_Compile_"
_135_Logic____Option__i_Compile "135_Logic____Option__i_Compile_"
_138_Native____NativeTypes__i_Compile "138_Native____NativeTypes__i_Compile_"
_141_Libraries____base__s_Compile "141_Libraries____base__s_Compile_"
_143_Math____power2__s_Compile "143_Math____power2__s_Compile_"
_145_Math____power__s_Compile "145_Math____power__s_Compile_"
_149_Math____power__i_Compile "149_Math____power__i_Compile_"
_153_Math____div__def__i_Compile "153_Math____div__def__i_Compile_"
_157_Math____div__boogie__i_Compile "157_Math____div__boogie__i_Compile_"
_162_Math____div__auto__proofs__i_Compile "162_Math____div__auto__proofs__i_Compile_"
_164_Math____div__auto__i_Compile "164_Math____div__auto__i_Compile_"
_166_Math____div__i_Compile "166_Math____div__i_Compile_"
_168_Math____power2__i_Compile "168_Math____power2__i_Compile_"
_170_Common____Util__i_Compile "170_Common____Util__i_Compile_"
_174_Common____MarshallInt__i_Compile "174_Common____MarshallInt__i_Compile_"
_176_Common____GenericMarshalling__i_Compile "176_Common____GenericMarshalling__i_Compile_"
_180_Common____UdpClient__i_Compile "180_Common____UdpClient__i_Compile_"
_182_Common____SeqIsUniqueDef__i_Compile "182_Common____SeqIsUniqueDef__i_Compile_"
_185_Common____SeqIsUnique__i_Compile "185_Common____SeqIsUnique__i_Compile_"
_191_GenericRefinement__i_Compile "191_GenericRefinement__i_Compile_"
_194_Common____NodeIdentity__i_Compile "194_Common____NodeIdentity__i_Compile_"
_197_LiveRSL____AppInterface__i_Compile "197_LiveRSL____AppInterface__i_Compile_"
_214_LiveRSL____CTypes__i_Compile "214_LiveRSL____CTypes__i_Compile_"
_217_LiveRSL____CMessage__i_Compile "217_LiveRSL____CMessage__i_Compile_"
_225_LiveRSL____CMessageRefinements__i_Compile "225_LiveRSL____CMessageRefinements__i_Compile_"
_228_LiveRSL____PacketParsing__i_Compile "228_LiveRSL____PacketParsing__i_Compile_"
_234_Common____UpperBound__i_Compile "234_Common____UpperBound__i_Compile_"
_236_LiveRSL____ParametersState__i_Compile "236_LiveRSL____ParametersState__i_Compile_"
_238_LiveRSL____CPaxosConfiguration__i_Compile "238_LiveRSL____CPaxosConfiguration__i_Compile_"
_240_LiveRSL____ConstantsState__i_Compile "240_LiveRSL____ConstantsState__i_Compile_"
_243_LiveRSL____PaxosWorldState__i_Compile "243_LiveRSL____PaxosWorldState__i_Compile_"
_245_LiveRSL____ReplicaConstantsState__i_Compile "245_LiveRSL____ReplicaConstantsState__i_Compile_"
_251_LiveRSL____ElectionState__i_Compile "251_LiveRSL____ElectionState__i_Compile_"
_254_LiveRSL____ProposerState__i_Compile "254_LiveRSL____ProposerState__i_Compile_"
_265_LiveRSL____COperationNumberSort__i_Compile "265_LiveRSL____COperationNumberSort__i_Compile_"
_267_LiveRSL____CLastCheckpointedMap__i_Compile "267_LiveRSL____CLastCheckpointedMap__i_Compile_"
_269_LiveRSL____AcceptorState__i_Compile "269_LiveRSL____AcceptorState__i_Compile_"
_276_LiveRSL____ExecutorState__i_Compile "276_LiveRSL____ExecutorState__i_Compile_"
_278_LiveRSL____LearnerState__i_Compile "278_LiveRSL____LearnerState__i_Compile_"
_283_LiveRSL____CClockReading__i_Compile "283_LiveRSL____CClockReading__i_Compile_"
_285_LiveRSL____ReplicaState__i_Compile "285_LiveRSL____ReplicaState__i_Compile_"
_294_LiveRSL____MinCQuorumSize__i_Compile "294_LiveRSL____MinCQuorumSize__i_Compile_"
_297_LiveRSL____ElectionModel__i_Compile "297_LiveRSL____ElectionModel__i_Compile_"
_301_Impl____LiveRSL____Broadcast__i_Compile "301_Impl____LiveRSL____Broadcast__i_Compile_"
_305_LiveRSL____ProposerLemmas__i_Compile "305_LiveRSL____ProposerLemmas__i_Compile_"
_308_LiveRSL____ProposerModel__i_Compile "308_LiveRSL____ProposerModel__i_Compile_"
_312_LiveRSL____AcceptorModel__i_Compile "312_LiveRSL____AcceptorModel__i_Compile_"
_316_LiveRSL____LearnerModel__i_Compile "316_LiveRSL____LearnerModel__i_Compile_"
_321_LiveRSL____ExecutorModel__i_Compile "321_LiveRSL____ExecutorModel__i_Compile_"
_323_LiveRSL____ReplicaModel__Part1__i_Compile "323_LiveRSL____ReplicaModel__Part1__i_Compile_"
_331_LiveRSL____ReplicaModel__Part2__i_Compile "331_LiveRSL____ReplicaModel__Part2__i_Compile_"
_339_LiveRSL____ReplicaModel__Part3__i_Compile "339_LiveRSL____ReplicaModel__Part3__i_Compile_"
_347_LiveRSL____ReplicaModel__Part4__i_Compile "347_LiveRSL____ReplicaModel__Part4__i_Compile_"
_355_LiveRSL____ReplicaModel__Part5__i_Compile "355_LiveRSL____ReplicaModel__Part5__i_Compile_"
_357_LiveRSL____ReplicaModel__i_Compile "357_LiveRSL____ReplicaModel__i_Compile_"
_368_LiveRSL____UdpRSL__i_Compile "368_LiveRSL____UdpRSL__i_Compile_"
_372_LiveRSL____QRelations__i_Compile "372_LiveRSL____QRelations__i_Compile_"
_374_LiveRSL____ReplicaImplLemmas__i_Compile "374_LiveRSL____ReplicaImplLemmas__i_Compile_"
_383_LiveRSL____ReplicaImplClass__i_Compile "383_LiveRSL____ReplicaImplClass__i_Compile_"
_405_LiveRSL____ReplicaImplDelivery__i_Compile "405_LiveRSL____ReplicaImplDelivery__i_Compile_"
_409_LiveRSL____ReplicaImplReadClock__i_Compile "409_LiveRSL____ReplicaImplReadClock__i_Compile_"
_420_LiveRSL____ReplicaImplProcessPacketNoClock__i_Compile "420_LiveRSL____ReplicaImplProcessPacketNoClock__i_Compile_"
_426_LiveRSL____Unsendable__i_Compile "426_LiveRSL____Unsendable__i_Compile_"
_428_LiveRSL____ReplicaImplProcessPacketX__i_Compile "428_LiveRSL____ReplicaImplProcessPacketX__i_Compile_"
_439_LiveRSL____ReplicaImplNoReceiveNoClock__i_Compile "439_LiveRSL____ReplicaImplNoReceiveNoClock__i_Compile_"
_450_LiveRSL____ReplicaImplNoReceiveClock__i_Compile "450_LiveRSL____ReplicaImplNoReceiveClock__i_Compile_"
_454_LiveRSL____ReplicaImplMain__i_Compile "454_LiveRSL____ReplicaImplMain__i_Compile_"
_460_CmdLineParser__i_Compile "460_CmdLineParser__i_Compile_"
_463_PaxosCmdLineParser__i_Compile "463_PaxosCmdLineParser__i_Compile_"
_466_Host__i_Compile "466_Host__i_Compile_"
_469_RSL__DistributedSystem__i_Compile "469_RSL__DistributedSystem__i_Compile_"
_472_Collections____Maps2__i_Compile "472_Collections____Maps2__i_Compile_"
_477_LiveRSL____DistributedSystem__i_Compile "477_LiveRSL____DistributedSystem__i_Compile_"
_479_DirectRefinement____StateMachine__i_Compile "479_DirectRefinement____StateMachine__i_Compile_"
_483_DirectRefinement____HandleRequestBatch__i_Compile "483_DirectRefinement____HandleRequestBatch__i_Compile_"
_489_Temporal____Heuristics__i_Compile "489_Temporal____Heuristics__i_Compile_"
_493_Temporal____Rules__i_Compile "493_Temporal____Rules__i_Compile_"
_496_CommonProof____Assumptions__i_Compile "496_CommonProof____Assumptions__i_Compile_"
_498_CommonProof____Constants__i_Compile "498_CommonProof____Constants__i_Compile_"
_503_CommonProof____Actions__i_Compile "503_CommonProof____Actions__i_Compile_"
_506_CommonProof____PacketSending__i_Compile "506_CommonProof____PacketSending__i_Compile_"
_517_CommonProof____Environment__i_Compile "517_CommonProof____Environment__i_Compile_"
_542_CommonProof____MaxBallotISent1a__i_Compile "542_CommonProof____MaxBallotISent1a__i_Compile_"
_547_CommonProof____Received1b__i_Compile "547_CommonProof____Received1b__i_Compile_"
_549_CommonProof____Message2a__i_Compile "549_CommonProof____Message2a__i_Compile_"
_551_CommonProof____Message2b__i_Compile "551_CommonProof____Message2b__i_Compile_"
_554_CommonProof____LearnerState__i_Compile "554_CommonProof____LearnerState__i_Compile_"
_558_CommonProof____Quorum__i_Compile "558_CommonProof____Quorum__i_Compile_"
_573_CommonProof____MaxBallot__i_Compile "573_CommonProof____MaxBallot__i_Compile_"
_575_CommonProof____Message1b__i_Compile "575_CommonProof____Message1b__i_Compile_"
_577_CommonProof____Chosen__i_Compile "577_CommonProof____Chosen__i_Compile_"
_582_DirectRefinement____Chosen__i_Compile "582_DirectRefinement____Chosen__i_Compile_"
_592_DirectRefinement____Execution__i_Compile "592_DirectRefinement____Execution__i_Compile_"
_602_CommonProof____Requests__i_Compile "602_CommonProof____Requests__i_Compile_"
_605_DirectRefinement____Requests__i_Compile "605_DirectRefinement____Requests__i_Compile_"
_610_DirectRefinement____Refinement__i_Compile "610_DirectRefinement____Refinement__i_Compile_"
)
var _ _dafny.Dummy__
var _ _System.Dummy__
var _ _0_Native____NativeTypes__s_Compile.Dummy__
var _ _2_Collections____Maps2__s_Compile.Dummy__
var _ _5_Temporal____Temporal__s_Compile.Dummy__
var _ _7_Environment__s_Compile.Dummy__
var _ _9_Native____Io__s_Compile.Dummy__
var _ _26_Collections____Seqs__s_Compile.Dummy__
var _ _30_Collections____Seqs__i_Compile.Dummy__
var _ _32_Math____mul__nonlinear__i_Compile.Dummy__
var _ _34_Math____mul__auto__proofs__i_Compile.Dummy__
var _ _36_Math____mul__auto__i_Compile.Dummy__
var _ _40_Math____mul__i_Compile.Dummy__
var _ _42_Math____div__nonlinear__i_Compile.Dummy__
var _ _44_Math____mod__auto__proofs__i_Compile.Dummy__
var _ _46_Math____mod__auto__i_Compile.Dummy__
var _ _50_AppStateMachine__i_Compile.Dummy__
var _ _54_Concrete__NodeIdentity__i_Compile.Dummy__
var _ _56_LiveRSL____Types__i_Compile.Dummy__
var _ _58_Collections____Sets__i_Compile.Dummy__
var _ _61_LiveRSL____Configuration__i_Compile.Dummy__
var _ _65_LiveRSL____Message__i_Compile.Dummy__
var _ _68_LiveRSL____Environment__i_Compile.Dummy__
var _ _71_LiveRSL____ClockReading__i_Compile.Dummy__
var _ _74_Common____UpperBound__s_Compile.Dummy__
var _ _76_LiveRSL____Parameters__i_Compile.Dummy__
var _ _78_LiveRSL____Constants__i_Compile.Dummy__
var _ _85_LiveRSL____Broadcast__i_Compile.Dummy__
var _ _91_Collections____CountMatches__i_Compile.Dummy__
var _ _93_LiveRSL____Acceptor__i_Compile.Dummy__
var _ _99_LiveRSL____Election__i_Compile.Dummy__
var _ _101_LiveRSL____Proposer__i_Compile.Dummy__
var _ _115_LiveRSL____StateMachine__i_Compile.Dummy__
var _ _118_Collections____Maps__i_Compile.Dummy__
var _ _120_LiveRSL____Executor__i_Compile.Dummy__
var _ _123_LiveRSL____Learner__i_Compile.Dummy__
var _ _126_LiveRSL____Replica__i_Compile.Dummy__
var _ _135_Logic____Option__i_Compile.Dummy__
var _ _138_Native____NativeTypes__i_Compile.Dummy__
var _ _141_Libraries____base__s_Compile.Dummy__
var _ _143_Math____power2__s_Compile.Dummy__
var _ _145_Math____power__s_Compile.Dummy__
var _ _149_Math____power__i_Compile.Dummy__
var _ _153_Math____div__def__i_Compile.Dummy__
var _ _157_Math____div__boogie__i_Compile.Dummy__
var _ _162_Math____div__auto__proofs__i_Compile.Dummy__
var _ _164_Math____div__auto__i_Compile.Dummy__
var _ _166_Math____div__i_Compile.Dummy__
var _ _168_Math____power2__i_Compile.Dummy__
var _ _170_Common____Util__i_Compile.Dummy__
var _ _174_Common____MarshallInt__i_Compile.Dummy__
var _ _176_Common____GenericMarshalling__i_Compile.Dummy__
var _ _180_Common____UdpClient__i_Compile.Dummy__
var _ _182_Common____SeqIsUniqueDef__i_Compile.Dummy__
var _ _185_Common____SeqIsUnique__i_Compile.Dummy__
var _ _191_GenericRefinement__i_Compile.Dummy__
var _ _194_Common____NodeIdentity__i_Compile.Dummy__
var _ _197_LiveRSL____AppInterface__i_Compile.Dummy__
var _ _214_LiveRSL____CTypes__i_Compile.Dummy__
var _ _217_LiveRSL____CMessage__i_Compile.Dummy__
var _ _225_LiveRSL____CMessageRefinements__i_Compile.Dummy__
var _ _228_LiveRSL____PacketParsing__i_Compile.Dummy__
var _ _234_Common____UpperBound__i_Compile.Dummy__
var _ _236_LiveRSL____ParametersState__i_Compile.Dummy__
var _ _238_LiveRSL____CPaxosConfiguration__i_Compile.Dummy__
var _ _240_LiveRSL____ConstantsState__i_Compile.Dummy__
var _ _243_LiveRSL____PaxosWorldState__i_Compile.Dummy__
var _ _245_LiveRSL____ReplicaConstantsState__i_Compile.Dummy__
var _ _251_LiveRSL____ElectionState__i_Compile.Dummy__
var _ _254_LiveRSL____ProposerState__i_Compile.Dummy__
var _ _265_LiveRSL____COperationNumberSort__i_Compile.Dummy__
var _ _267_LiveRSL____CLastCheckpointedMap__i_Compile.Dummy__
var _ _269_LiveRSL____AcceptorState__i_Compile.Dummy__
var _ _276_LiveRSL____ExecutorState__i_Compile.Dummy__
var _ _278_LiveRSL____LearnerState__i_Compile.Dummy__
var _ _283_LiveRSL____CClockReading__i_Compile.Dummy__
var _ _285_LiveRSL____ReplicaState__i_Compile.Dummy__
var _ _294_LiveRSL____MinCQuorumSize__i_Compile.Dummy__
var _ _297_LiveRSL____ElectionModel__i_Compile.Dummy__
var _ _301_Impl____LiveRSL____Broadcast__i_Compile.Dummy__
var _ _305_LiveRSL____ProposerLemmas__i_Compile.Dummy__
var _ _308_LiveRSL____ProposerModel__i_Compile.Dummy__
var _ _312_LiveRSL____AcceptorModel__i_Compile.Dummy__
var _ _316_LiveRSL____LearnerModel__i_Compile.Dummy__
var _ _321_LiveRSL____ExecutorModel__i_Compile.Dummy__
var _ _323_LiveRSL____ReplicaModel__Part1__i_Compile.Dummy__
var _ _331_LiveRSL____ReplicaModel__Part2__i_Compile.Dummy__
var _ _339_LiveRSL____ReplicaModel__Part3__i_Compile.Dummy__
var _ _347_LiveRSL____ReplicaModel__Part4__i_Compile.Dummy__
var _ _355_LiveRSL____ReplicaModel__Part5__i_Compile.Dummy__
var _ _357_LiveRSL____ReplicaModel__i_Compile.Dummy__
var _ _368_LiveRSL____UdpRSL__i_Compile.Dummy__
var _ _372_LiveRSL____QRelations__i_Compile.Dummy__
var _ _374_LiveRSL____ReplicaImplLemmas__i_Compile.Dummy__
var _ _383_LiveRSL____ReplicaImplClass__i_Compile.Dummy__
var _ _405_LiveRSL____ReplicaImplDelivery__i_Compile.Dummy__
var _ _409_LiveRSL____ReplicaImplReadClock__i_Compile.Dummy__
var _ _420_LiveRSL____ReplicaImplProcessPacketNoClock__i_Compile.Dummy__
var _ _426_LiveRSL____Unsendable__i_Compile.Dummy__
var _ _428_LiveRSL____ReplicaImplProcessPacketX__i_Compile.Dummy__
var _ _439_LiveRSL____ReplicaImplNoReceiveNoClock__i_Compile.Dummy__
var _ _450_LiveRSL____ReplicaImplNoReceiveClock__i_Compile.Dummy__
var _ _454_LiveRSL____ReplicaImplMain__i_Compile.Dummy__
var _ _460_CmdLineParser__i_Compile.Dummy__
var _ _463_PaxosCmdLineParser__i_Compile.Dummy__
var _ _466_Host__i_Compile.Dummy__
var _ _469_RSL__DistributedSystem__i_Compile.Dummy__
var _ _472_Collections____Maps2__i_Compile.Dummy__
var _ _477_LiveRSL____DistributedSystem__i_Compile.Dummy__
var _ _479_DirectRefinement____StateMachine__i_Compile.Dummy__
var _ _483_DirectRefinement____HandleRequestBatch__i_Compile.Dummy__
var _ _489_Temporal____Heuristics__i_Compile.Dummy__
var _ _493_Temporal____Rules__i_Compile.Dummy__
var _ _496_CommonProof____Assumptions__i_Compile.Dummy__
var _ _498_CommonProof____Constants__i_Compile.Dummy__
var _ _503_CommonProof____Actions__i_Compile.Dummy__
var _ _506_CommonProof____PacketSending__i_Compile.Dummy__
var _ _517_CommonProof____Environment__i_Compile.Dummy__
var _ _542_CommonProof____MaxBallotISent1a__i_Compile.Dummy__
var _ _547_CommonProof____Received1b__i_Compile.Dummy__
var _ _549_CommonProof____Message2a__i_Compile.Dummy__
var _ _551_CommonProof____Message2b__i_Compile.Dummy__
var _ _554_CommonProof____LearnerState__i_Compile.Dummy__
var _ _558_CommonProof____Quorum__i_Compile.Dummy__
var _ _573_CommonProof____MaxBallot__i_Compile.Dummy__
var _ _575_CommonProof____Message1b__i_Compile.Dummy__
var _ _577_CommonProof____Chosen__i_Compile.Dummy__
var _ _582_DirectRefinement____Chosen__i_Compile.Dummy__
var _ _592_DirectRefinement____Execution__i_Compile.Dummy__
var _ _602_CommonProof____Requests__i_Compile.Dummy__
var _ _605_DirectRefinement____Requests__i_Compile.Dummy__
var _ _610_DirectRefinement____Refinement__i_Compile.Dummy__

type Dummy__ struct{}




// Definition of data type AppRequest
type AppRequest struct {
  Data_AppRequest_
}

func (_this AppRequest) Get() Data_AppRequest_ {
  return _this.Data_AppRequest_
}

type Data_AppRequest_ interface {
  isAppRequest()
}

type CompanionStruct_AppRequest_ struct {}
var Companion_AppRequest_ = CompanionStruct_AppRequest_{}

type AppRequest_AppRequest struct {
  Client _9_Native____Io__s_Compile.EndPoint
Seqno _dafny.Int
Request _50_AppStateMachine__i_Compile.AppMessage_k
}

func (AppRequest_AppRequest) isAppRequest() {}

func (CompanionStruct_AppRequest_) Create_AppRequest_(Client _9_Native____Io__s_Compile.EndPoint, Seqno _dafny.Int, Request _50_AppStateMachine__i_Compile.AppMessage_k) AppRequest {
  return AppRequest{AppRequest_AppRequest{Client,Seqno,Request}}
}

func (_this AppRequest) Is_AppRequest() bool {
  _, ok := _this.Get().(AppRequest_AppRequest)
return ok
}

func (_this AppRequest) Dtor_client() _9_Native____Io__s_Compile.EndPoint {
  return _this.Get().(AppRequest_AppRequest).Client
}

func (_this AppRequest) Dtor_seqno() _dafny.Int {
  return _this.Get().(AppRequest_AppRequest).Seqno
}

func (_this AppRequest) Dtor_request() _50_AppStateMachine__i_Compile.AppMessage_k {
  return _this.Get().(AppRequest_AppRequest).Request
}

func (_this AppRequest) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case AppRequest_AppRequest: {
      return "_615_AbstractServiceRSL__s_Compile.AppRequest.AppRequest" + "(" + _dafny.String(data.Client) + ", " + _dafny.String(data.Seqno) + ", " + _dafny.String(data.Request) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this AppRequest) Equals(other AppRequest) bool {
  switch data1 := _this.Get().(type) {
    case AppRequest_AppRequest: {
      data2, ok := other.Get().(AppRequest_AppRequest)
return ok && data1.Client.Equals(data2.Client) && data1.Seqno.Cmp(data2.Seqno) == 0 && data1.Request.Equals(data2.Request)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this AppRequest) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(AppRequest)
return ok && _this.Equals(typed)
}
func Type_AppRequest_() _dafny.Type {
  return type_AppRequest_{}
}

type type_AppRequest_ struct {
}

func (_this type_AppRequest_) Default() interface{} {
  return AppRequest{AppRequest_AppRequest{_9_Native____Io__s_Compile.Type_EndPoint_().Default().(_9_Native____Io__s_Compile.EndPoint), _dafny.Zero, _50_AppStateMachine__i_Compile.Type_AppMessage_k_().Default().(_50_AppStateMachine__i_Compile.AppMessage_k)}}
}

func (_this type_AppRequest_) String() string {
  return "_615_AbstractServiceRSL__s_Compile.AppRequest"
}
// End of data type AppRequest

// Definition of data type AppReply
type AppReply struct {
  Data_AppReply_
}

func (_this AppReply) Get() Data_AppReply_ {
  return _this.Data_AppReply_
}

type Data_AppReply_ interface {
  isAppReply()
}

type CompanionStruct_AppReply_ struct {}
var Companion_AppReply_ = CompanionStruct_AppReply_{}

type AppReply_AppReply struct {
  Client _9_Native____Io__s_Compile.EndPoint
Seqno _dafny.Int
Reply _50_AppStateMachine__i_Compile.AppMessage_k
}

func (AppReply_AppReply) isAppReply() {}

func (CompanionStruct_AppReply_) Create_AppReply_(Client _9_Native____Io__s_Compile.EndPoint, Seqno _dafny.Int, Reply _50_AppStateMachine__i_Compile.AppMessage_k) AppReply {
  return AppReply{AppReply_AppReply{Client,Seqno,Reply}}
}

func (_this AppReply) Is_AppReply() bool {
  _, ok := _this.Get().(AppReply_AppReply)
return ok
}

func (_this AppReply) Dtor_client() _9_Native____Io__s_Compile.EndPoint {
  return _this.Get().(AppReply_AppReply).Client
}

func (_this AppReply) Dtor_seqno() _dafny.Int {
  return _this.Get().(AppReply_AppReply).Seqno
}

func (_this AppReply) Dtor_reply() _50_AppStateMachine__i_Compile.AppMessage_k {
  return _this.Get().(AppReply_AppReply).Reply
}

func (_this AppReply) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case AppReply_AppReply: {
      return "_615_AbstractServiceRSL__s_Compile.AppReply.AppReply" + "(" + _dafny.String(data.Client) + ", " + _dafny.String(data.Seqno) + ", " + _dafny.String(data.Reply) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this AppReply) Equals(other AppReply) bool {
  switch data1 := _this.Get().(type) {
    case AppReply_AppReply: {
      data2, ok := other.Get().(AppReply_AppReply)
return ok && data1.Client.Equals(data2.Client) && data1.Seqno.Cmp(data2.Seqno) == 0 && data1.Reply.Equals(data2.Reply)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this AppReply) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(AppReply)
return ok && _this.Equals(typed)
}
func Type_AppReply_() _dafny.Type {
  return type_AppReply_{}
}

type type_AppReply_ struct {
}

func (_this type_AppReply_) Default() interface{} {
  return AppReply{AppReply_AppReply{_9_Native____Io__s_Compile.Type_EndPoint_().Default().(_9_Native____Io__s_Compile.EndPoint), _dafny.Zero, _50_AppStateMachine__i_Compile.Type_AppMessage_k_().Default().(_50_AppStateMachine__i_Compile.AppMessage_k)}}
}

func (_this type_AppReply_) String() string {
  return "_615_AbstractServiceRSL__s_Compile.AppReply"
}
// End of data type AppReply

// Definition of data type ServiceState_k
type ServiceState_k struct {
  Data_ServiceState_k_
}

func (_this ServiceState_k) Get() Data_ServiceState_k_ {
  return _this.Data_ServiceState_k_
}

type Data_ServiceState_k_ interface {
  isServiceState_k()
}

type CompanionStruct_ServiceState_k_ struct {}
var Companion_ServiceState_k_ = CompanionStruct_ServiceState_k_{}

type ServiceState_k_ServiceState_k struct {
  ServerAddresses _dafny.Set
App uint64
Requests _dafny.Set
Replies _dafny.Set
}

func (ServiceState_k_ServiceState_k) isServiceState_k() {}

func (CompanionStruct_ServiceState_k_) Create_ServiceState_k_(ServerAddresses _dafny.Set, App uint64, Requests _dafny.Set, Replies _dafny.Set) ServiceState_k {
  return ServiceState_k{ServiceState_k_ServiceState_k{ServerAddresses,App,Requests,Replies}}
}

func (_this ServiceState_k) Is_ServiceState_k() bool {
  _, ok := _this.Get().(ServiceState_k_ServiceState_k)
return ok
}

func (_this ServiceState_k) Dtor_serverAddresses() _dafny.Set {
  return _this.Get().(ServiceState_k_ServiceState_k).ServerAddresses
}

func (_this ServiceState_k) Dtor_app() uint64 {
  return _this.Get().(ServiceState_k_ServiceState_k).App
}

func (_this ServiceState_k) Dtor_requests() _dafny.Set {
  return _this.Get().(ServiceState_k_ServiceState_k).Requests
}

func (_this ServiceState_k) Dtor_replies() _dafny.Set {
  return _this.Get().(ServiceState_k_ServiceState_k).Replies
}

func (_this ServiceState_k) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case ServiceState_k_ServiceState_k: {
      return "_615_AbstractServiceRSL__s_Compile.ServiceState'.ServiceState'" + "(" + _dafny.String(data.ServerAddresses) + ", " + _dafny.String(data.App) + ", " + _dafny.String(data.Requests) + ", " + _dafny.String(data.Replies) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this ServiceState_k) Equals(other ServiceState_k) bool {
  switch data1 := _this.Get().(type) {
    case ServiceState_k_ServiceState_k: {
      data2, ok := other.Get().(ServiceState_k_ServiceState_k)
return ok && data1.ServerAddresses.Equals(data2.ServerAddresses) && data1.App == data2.App && data1.Requests.Equals(data2.Requests) && data1.Replies.Equals(data2.Replies)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this ServiceState_k) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(ServiceState_k)
return ok && _this.Equals(typed)
}
func Type_ServiceState_k_() _dafny.Type {
  return type_ServiceState_k_{}
}

type type_ServiceState_k_ struct {
}

func (_this type_ServiceState_k_) Default() interface{} {
  return ServiceState_k{ServiceState_k_ServiceState_k{_dafny.EmptySet, 0, _dafny.EmptySet, _dafny.EmptySet}}
}

func (_this type_ServiceState_k_) String() string {
  return "_615_AbstractServiceRSL__s_Compile.ServiceState_k"
}
// End of data type ServiceState_k





