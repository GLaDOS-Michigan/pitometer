// Package _331_LiveRSL____ReplicaModel__Part2__i_Compile
// Dafny module _331_LiveRSL____ReplicaModel__Part2__i_Compile compiled into Go

package _331_LiveRSL____ReplicaModel__Part2__i_Compile

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

type Dummy__ struct{}








// Definition of class Default__
type Default__ struct {
  dummy byte
}

func New_Default___() *Default__ {
  _this := Default__{}

  return &_this
}

type CompanionStruct_Default___ struct {
}
var Companion_Default___ = CompanionStruct_Default___ {
}

func (_this *Default__) Equals(other *Default__) bool {
  return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
  other, ok := x.(*Default__)
return ok && _this.Equals(other)
}

func (*Default__) String() string {
  return "_331_LiveRSL____ReplicaModel__Part2__i_Compile.Default__"
}

func Type_Default___() _dafny.Type {
  return type_Default___{}
}

type type_Default___ struct {
}

func (_this type_Default___) Default() interface{} {
  return (*Default__)(nil)
}

func (_this type_Default___) String() string {
  return "_331_LiveRSL____ReplicaModel__Part2__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) ProposerSrcNotPresent(proposer _254_LiveRSL____ProposerState__i_Compile.ProposerState, packet _217_LiveRSL____CMessage__i_Compile.CPacket) bool {
  goto TAIL_CALL_START
TAIL_CALL_START:
var b bool = false
  var _ = b
  b = _dafny.Quantifier(((proposer).Dtor_received__1b__packets()).Elements(), true, func (_5079_other__packet _217_LiveRSL____CMessage__i_Compile.CPacket) bool {
    return !(((proposer).Dtor_received__1b__packets()).Contains(_5079_other__packet)) || (!((_5079_other__packet).Dtor_src()).Equals((packet).Dtor_src()))
  })
  return b
}
func (_this *CompanionStruct_Default___) ReplicaNextProcess1bIgnore(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  replica_k = replica
  packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}}}
  return replica_k,packets__sent
}
func (_this *CompanionStruct_Default___) ReplicaNextProcess1bAlreadyHave1bFromSource(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  { }
  { }
  replica_k = replica
  packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}}}
  return replica_k,packets__sent
}
func (_this *CompanionStruct_Default___) ReplicaNextProcess1bActual(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  { }
  { }
  packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}}}
  { }
  { }
  var _5080_newProposer _254_LiveRSL____ProposerState__i_Compile.ProposerState
  var _ = _5080_newProposer
var _out273 _254_LiveRSL____ProposerState__i_Compile.ProposerState
  var _ = _out273
_out273 = _308_LiveRSL____ProposerModel__i_Compile.Companion_Default___.ProposerProcess1b((replica).Dtor_proposer(), inp)
_5080_newProposer = _out273
  var _5081_newAcceptor _269_LiveRSL____AcceptorState__i_Compile.AcceptorState
  var _ = _5081_newAcceptor
var _out274 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState
  var _ = _out274
_out274 = _312_LiveRSL____AcceptorModel__i_Compile.Companion_Default___.NextAcceptorState__TruncateLog((replica).Dtor_acceptor(), ((inp).Dtor_msg()).Dtor_log__truncation__point())
_5081_newAcceptor = _out274
  replica_k = func (_pat_let119_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
    return func (_5082_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
      return func (_pat_let120_0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
        return func (_5083_dt__update_hacceptor_h0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
          return func (_pat_let121_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
            return func (_5084_dt__update_hproposer_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
              return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5082_dt__update__tmp_h0).Dtor_constants(), (_5082_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), _5084_dt__update_hproposer_h0, _5083_dt__update_hacceptor_h0, (_5082_dt__update__tmp_h0).Dtor_learner(), (_5082_dt__update__tmp_h0).Dtor_executor()}}
            }(_pat_let121_0)
          }(_5080_newProposer)
        }(_pat_let120_0)
      }(_5081_newAcceptor)
    }(_pat_let119_0)
  }(replica)
  { }
  return replica_k,packets__sent
}
func (_this *CompanionStruct_Default___) Replica__Next__Process__1b(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  var _5085_start__time uint64
  var _ = _5085_start__time
var _out275 uint64
  var _ = _out275
_out275 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5085_start__time = _out275
  if (((!((((((replica).Dtor_proposer()).Dtor_constants()).Dtor_all()).Dtor_config()).Dtor_replica__ids()).Contains((inp).Dtor_src())) || (!(((inp).Dtor_msg()).Dtor_bal__1b()).Equals(((replica).Dtor_proposer()).Dtor_max__ballot__i__sent__1a()))) || ((((replica).Dtor_proposer()).Dtor_current__state()) != (uint8(1))/* dircomp */)) {
    var _out276 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
    var _ = _out276
var _out277 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
    var _ = _out277
_out276,_out277 = Companion_Default___.ReplicaNextProcess1bIgnore(replica, inp)
replica_k = _out276
packets__sent = _out277
    var _5086_end__time uint64
    var _ = _5086_end__time
var _out278 uint64
    var _ = _out278
_out278 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5086_end__time = _out278
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Process_1b_discard"), _5085_start__time, _5086_end__time)
  } else {
    var _5087_srcNotPresent bool
    var _ = _5087_srcNotPresent
var _out279 bool
    var _ = _out279
_out279 = Companion_Default___.ProposerSrcNotPresent((replica).Dtor_proposer(), inp)
_5087_srcNotPresent = _out279
    if (_5087_srcNotPresent) {
      var _out280 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
      var _ = _out280
var _out281 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
      var _ = _out281
_out280,_out281 = Companion_Default___.ReplicaNextProcess1bActual(replica, inp)
replica_k = _out280
packets__sent = _out281
      var _5088_end__time uint64
      var _ = _5088_end__time
var _out282 uint64
      var _ = _out282
_out282 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5088_end__time = _out282
      _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Process_1b_use"), _5085_start__time, _5088_end__time)
    } else {
      var _out283 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
      var _ = _out283
var _out284 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
      var _ = _out284
_out283,_out284 = Companion_Default___.ReplicaNextProcess1bAlreadyHave1bFromSource(replica, inp)
replica_k = _out283
packets__sent = _out284
      var _5089_end__time uint64
      var _ = _5089_end__time
var _out285 uint64
      var _ = _out285
_out285 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5089_end__time = _out285
      _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Process_1b_discard"), _5085_start__time, _5089_end__time)
    }
  }
  return replica_k,packets__sent
}
func (_this *CompanionStruct_Default___) Replica__Next__Process__StartingPhase2(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  var _5090_start__time uint64
  var _ = _5090_start__time
var _out286 uint64
  var _ = _out286
_out286 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5090_start__time = _out286
  var _5091_newExecutor _276_LiveRSL____ExecutorState__i_Compile.ExecutorState
  var _ = _5091_newExecutor
var _5092_packets _217_LiveRSL____CMessage__i_Compile.CBroadcast
  var _ = _5092_packets
var _out287 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState
  var _ = _out287
var _out288 _217_LiveRSL____CMessage__i_Compile.CBroadcast
  var _ = _out288
_out287,_out288 = _321_LiveRSL____ExecutorModel__i_Compile.Companion_Default___.ExecutorProcessStartingPhase2((replica).Dtor_executor(), inp)
_5091_newExecutor = _out287
_5092_packets = _out288
  replica_k = func (_pat_let122_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
    return func (_5093_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
      return func (_pat_let123_0 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
        return func (_5094_dt__update_hexecutor_h0 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
          return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5093_dt__update__tmp_h0).Dtor_constants(), (_5093_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), (_5093_dt__update__tmp_h0).Dtor_proposer(), (_5093_dt__update__tmp_h0).Dtor_acceptor(), (_5093_dt__update__tmp_h0).Dtor_learner(), _5094_dt__update_hexecutor_h0}}
        }(_pat_let123_0)
      }(_5091_newExecutor)
    }(_pat_let122_0)
  }(replica)
  packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_5092_packets}}
  var _5095_end__time uint64
  var _ = _5095_end__time
var _out289 uint64
  var _ = _out289
_out289 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5095_end__time = _out289
  _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Process_StartingPhase2"), _5090_start__time, _5095_end__time)
  return replica_k,packets__sent
}
func (_this *CompanionStruct_Default___) ReplicaNextProcess2aIgnore(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  replica_k = replica
  packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}}}
  return replica_k,packets__sent
}
func (_this *CompanionStruct_Default___) ReplicaNextProcess2aActual(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  { }
  var _5096_maxLogLengthMinus1 uint64
  var _ = _5096_maxLogLengthMinus1
  _5096_maxLogLengthMinus1 = ((((((replica).Dtor_acceptor()).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__log__length()) - (func () uint64 { return  (uint64(1)) })()
  var _5097_newLogTruncationPoint uint64
  var _ = _5097_newLogTruncationPoint
  _5097_newLogTruncationPoint = (((replica).Dtor_acceptor()).Dtor_log__truncation__point()).Dtor_n()
  var _5098_newAcceptor _269_LiveRSL____AcceptorState__i_Compile.AcceptorState
  var _ = _5098_newAcceptor
var _5099_packets _217_LiveRSL____CMessage__i_Compile.CBroadcast
  var _ = _5099_packets
var _out290 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState
  var _ = _out290
var _out291 _217_LiveRSL____CMessage__i_Compile.CBroadcast
  var _ = _out291
_out290,_out291 = _312_LiveRSL____AcceptorModel__i_Compile.Companion_Default___.NextAcceptorState__Phase2((replica).Dtor_acceptor(), (inp).Dtor_msg(), (inp).Dtor_src())
_5098_newAcceptor = _out290
_5099_packets = _out291
  replica_k = func (_pat_let124_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
    return func (_5100_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
      return func (_pat_let125_0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
        return func (_5101_dt__update_hacceptor_h0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
          return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5100_dt__update__tmp_h0).Dtor_constants(), (_5100_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), (_5100_dt__update__tmp_h0).Dtor_proposer(), _5101_dt__update_hacceptor_h0, (_5100_dt__update__tmp_h0).Dtor_learner(), (_5100_dt__update__tmp_h0).Dtor_executor()}}
        }(_pat_let125_0)
      }(_5098_newAcceptor)
    }(_pat_let124_0)
  }(replica)
  packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_5099_packets}}
  return replica_k,packets__sent
}
func (_this *CompanionStruct_Default___) Replica__Next__Process__2a(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  var _5102_start__time uint64
  var _ = _5102_start__time
var _out292 uint64
  var _ = _out292
_out292 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5102_start__time = _out292
  var _5103_ballot__leq bool
  var _ = _5103_ballot__leq
var _out293 bool
  var _ = _out293
_out293 = _214_LiveRSL____CTypes__i_Compile.Companion_Default___.CBalLeq(((replica).Dtor_acceptor()).Dtor_maxBallot(), ((inp).Dtor_msg()).Dtor_bal__2a())
_5103_ballot__leq = _out293
  if (((((((((replica).Dtor_acceptor()).Dtor_constants()).Dtor_all()).Dtor_config()).Dtor_replica__ids()).Contains((inp).Dtor_src())) && (_5103_ballot__leq)) && (((((inp).Dtor_msg()).Dtor_opn__2a()).Dtor_n()) <= ((((((replica).Dtor_acceptor()).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val()))) {
    var _out294 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
    var _ = _out294
var _out295 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
    var _ = _out295
_out294,_out295 = Companion_Default___.ReplicaNextProcess2aActual(replica, inp)
replica_k = _out294
packets__sent = _out295
    var _5104_end__time uint64
    var _ = _5104_end__time
var _out296 uint64
    var _ = _out296
_out296 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5104_end__time = _out296
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Process_2a_use"), _5102_start__time, _5104_end__time)
  } else {
    var _out297 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
    var _ = _out297
var _out298 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
    var _ = _out298
_out297,_out298 = Companion_Default___.ReplicaNextProcess2aIgnore(replica, inp)
replica_k = _out297
packets__sent = _out298
    var _5105_end__time uint64
    var _ = _5105_end__time
var _out299 uint64
    var _ = _out299
_out299 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5105_end__time = _out299
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Process_2a_discard"), _5102_start__time, _5105_end__time)
  }
  return replica_k,packets__sent
}
// End of class Default__
