// Package _355_LiveRSL____ReplicaModel__Part5__i_Compile
// Dafny module _355_LiveRSL____ReplicaModel__Part5__i_Compile compiled into Go

package _355_LiveRSL____ReplicaModel__Part5__i_Compile

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
  return "_355_LiveRSL____ReplicaModel__Part5__i_Compile.Default__"
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
  return "_355_LiveRSL____ReplicaModel__Part5__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) ReplicaNextProcessAppStateSupplyIgnore(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
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
func (_this *CompanionStruct_Default___) ReplicaNextProcessAppStateSupplyActual(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets, *_9_Native____Io__s_Compile.MutableMap) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
var reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap = (*_9_Native____Io__s_Compile.MutableMap)(nil)
  var _ = reply__cache__mutable
  var _5154_newLearner _278_LiveRSL____LearnerState__i_Compile.CLearnerState
  var _ = _5154_newLearner
var _out329 _278_LiveRSL____LearnerState__i_Compile.CLearnerState
  var _ = _out329
_out329 = _316_LiveRSL____LearnerModel__i_Compile.Companion_Default___.LearnerModel__ForgetOperationsBefore((replica).Dtor_learner(), ((inp).Dtor_msg()).Dtor_opn__state__supply())
_5154_newLearner = _out329
  var _5155_newExecutor _276_LiveRSL____ExecutorState__i_Compile.ExecutorState = _276_LiveRSL____ExecutorState__i_Compile.Type_ExecutorState_().Default().(_276_LiveRSL____ExecutorState__i_Compile.ExecutorState)
  var _ = _5155_newExecutor
  var _out330 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState
  var _ = _out330
var _out331 *_9_Native____Io__s_Compile.MutableMap
  var _ = _out331
_out330,_out331 = _321_LiveRSL____ExecutorModel__i_Compile.Companion_Default___.ExecutorProcessAppStateSupply((replica).Dtor_executor(), inp)
_5155_newExecutor = _out330
reply__cache__mutable = _out331
  replica_k = func (_pat_let143_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
    return func (_5156_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
      return func (_pat_let144_0 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
        return func (_5157_dt__update_hexecutor_h0 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
          return func (_pat_let145_0 _278_LiveRSL____LearnerState__i_Compile.CLearnerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
            return func (_5158_dt__update_hlearner_h0 _278_LiveRSL____LearnerState__i_Compile.CLearnerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
              return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5156_dt__update__tmp_h0).Dtor_constants(), (_5156_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), (_5156_dt__update__tmp_h0).Dtor_proposer(), (_5156_dt__update__tmp_h0).Dtor_acceptor(), _5158_dt__update_hlearner_h0, _5157_dt__update_hexecutor_h0}}
            }(_pat_let145_0)
          }(_5154_newLearner)
        }(_pat_let144_0)
      }(_5155_newExecutor)
    }(_pat_let143_0)
  }(replica)
  packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}}}
  return replica_k,packets__sent,reply__cache__mutable
}
func (_this *CompanionStruct_Default___) Replica__Next__Process__AppStateSupply(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, inp _217_LiveRSL____CMessage__i_Compile.CPacket) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets, bool, *_9_Native____Io__s_Compile.MutableMap) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
var replicaChanged bool = false
  var _ = replicaChanged
var reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap = (*_9_Native____Io__s_Compile.MutableMap)(nil)
  var _ = reply__cache__mutable
  var _5159_empty__Mutable__Map *_9_Native____Io__s_Compile.MutableMap
  var _ = _5159_empty__Mutable__Map
var _out332 *_9_Native____Io__s_Compile.MutableMap
  var _ = _out332
_out332 = _9_Native____Io__s_Compile.Companion_MutableMap_.EmptyMap()
_5159_empty__Mutable__Map = _out332
  reply__cache__mutable = _5159_empty__Mutable__Map
  var _5160_start__time uint64
  var _ = _5160_start__time
var _out333 uint64
  var _ = _out333
_out333 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5160_start__time = _out333
  if ((((((((replica).Dtor_executor()).Dtor_constants()).Dtor_all()).Dtor_config()).Dtor_replica__ids()).Contains((inp).Dtor_src())) && (((((inp).Dtor_msg()).Dtor_opn__state__supply()).Dtor_n()) > ((((replica).Dtor_executor()).Dtor_ops__complete()).Dtor_n()))) {
    var _out334 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
    var _ = _out334
var _out335 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
    var _ = _out335
var _out336 *_9_Native____Io__s_Compile.MutableMap
    var _ = _out336
_out334,_out335,_out336 = Companion_Default___.ReplicaNextProcessAppStateSupplyActual(replica, inp)
replica_k = _out334
packets__sent = _out335
reply__cache__mutable = _out336
    var _5161_end__time uint64
    var _ = _5161_end__time
var _out337 uint64
    var _ = _out337
_out337 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5161_end__time = _out337
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Process_AppStateSupply_work"), _5160_start__time, _5161_end__time)
    replicaChanged = true
  } else {
    var _out338 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
    var _ = _out338
var _out339 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
    var _ = _out339
_out338,_out339 = Companion_Default___.ReplicaNextProcessAppStateSupplyIgnore(replica, inp)
replica_k = _out338
packets__sent = _out339
    var _5162_end__time uint64
    var _ = _5162_end__time
var _out340 uint64
    var _ = _out340
_out340 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5162_end__time = _out340
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Process_AppStateSupply_nada"), _5160_start__time, _5162_end__time)
    replicaChanged = false
  }
  return replica_k,packets__sent,replicaChanged,reply__cache__mutable
}
func (_this *CompanionStruct_Default___) ReplicaNextSpontaneousMaybeExecuteIgnore(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  replica_k = replica
  packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_OutboundPacket{_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_None{}}}}
  return replica_k,packets__sent
}
func (_this *CompanionStruct_Default___) ReplicaNextSpontaneousMaybeExecuteActual(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet, reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  var _5163_val _dafny.Seq
  var _ = _5163_val
  _5163_val = (((replica).Dtor_executor()).Dtor_next__op__to__execute()).Dtor_v()
  var _5164_newLearner _278_LiveRSL____LearnerState__i_Compile.CLearnerState
  var _ = _5164_newLearner
var _out341 _278_LiveRSL____LearnerState__i_Compile.CLearnerState
  var _ = _out341
_out341 = _316_LiveRSL____LearnerModel__i_Compile.Companion_Default___.LearnerModel__ForgetDecision((replica).Dtor_learner(), ((replica).Dtor_executor()).Dtor_ops__complete())
_5164_newLearner = _out341
  { }
  var _5165_newExecutor _276_LiveRSL____ExecutorState__i_Compile.ExecutorState
  var _ = _5165_newExecutor
var _5166_packets _217_LiveRSL____CMessage__i_Compile.OutboundPackets
  var _ = _5166_packets
var _out342 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState
  var _ = _out342
var _out343 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
  var _ = _out343
_out342,_out343 = _321_LiveRSL____ExecutorModel__i_Compile.Companion_Default___.ExecutorExecute((replica).Dtor_executor(), reply__cache__mutable)
_5165_newExecutor = _out342
_5166_packets = _out343
  { }
  var _5167_newProposer _254_LiveRSL____ProposerState__i_Compile.ProposerState
  var _ = _5167_newProposer
var _out344 _254_LiveRSL____ProposerState__i_Compile.ProposerState
  var _ = _out344
_out344 = _308_LiveRSL____ProposerModel__i_Compile.Companion_Default___.ProposerResetViewTimerDueToExecution((replica).Dtor_proposer(), _5163_val, cur__req__set, prev__req__set)
_5167_newProposer = _out344
  { }
  { }
  { }
  replica_k = func (_pat_let146_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
    return func (_5168_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
      return func (_pat_let147_0 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
        return func (_5169_dt__update_hexecutor_h0 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
          return func (_pat_let148_0 _278_LiveRSL____LearnerState__i_Compile.CLearnerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
            return func (_5170_dt__update_hlearner_h0 _278_LiveRSL____LearnerState__i_Compile.CLearnerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
              return func (_pat_let149_0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
                return func (_5171_dt__update_hproposer_h0 _254_LiveRSL____ProposerState__i_Compile.ProposerState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
                  return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5168_dt__update__tmp_h0).Dtor_constants(), (_5168_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), _5171_dt__update_hproposer_h0, (_5168_dt__update__tmp_h0).Dtor_acceptor(), _5170_dt__update_hlearner_h0, _5169_dt__update_hexecutor_h0}}
                }(_pat_let149_0)
              }(_5167_newProposer)
            }(_pat_let148_0)
          }(_5164_newLearner)
        }(_pat_let147_0)
      }(_5165_newExecutor)
    }(_pat_let146_0)
  }(replica)
  packets__sent = _5166_packets
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  return replica_k,packets__sent
}
func (_this *CompanionStruct_Default___) Replica__Next__Spontaneous__MaybeExecute(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, cur__req__set *_9_Native____Io__s_Compile.MutableSet, prev__req__set *_9_Native____Io__s_Compile.MutableSet, reply__cache__mutable *_9_Native____Io__s_Compile.MutableMap) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  if (((((replica).Dtor_executor()).Dtor_next__op__to__execute()).Is_COutstandingOpKnown()) && (((((replica).Dtor_executor()).Dtor_ops__complete()).Dtor_n()) < ((((((replica).Dtor_executor()).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val()))) {
    var _out345 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
    var _ = _out345
var _out346 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
    var _ = _out346
_out345,_out346 = Companion_Default___.ReplicaNextSpontaneousMaybeExecuteActual(replica, cur__req__set, prev__req__set, reply__cache__mutable)
replica_k = _out345
packets__sent = _out346
  } else {
    var _out347 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
    var _ = _out347
var _out348 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
    var _ = _out348
_out347,_out348 = Companion_Default___.ReplicaNextSpontaneousMaybeExecuteIgnore(replica)
replica_k = _out347
packets__sent = _out348
  }
  return replica_k,packets__sent
}
func (_this *CompanionStruct_Default___) ReplicaNextReadClockMaybeSendHeartbeatSkip(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, clock _283_LiveRSL____CClockReading__i_Compile.CClockReading) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
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
func (_this *CompanionStruct_Default___) ReplicaNextReadClockMaybeSendHeartbeatActual(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, clock _283_LiveRSL____CClockReading__i_Compile.CClockReading) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  var _5172_heartbeat uint64
  var _ = _5172_heartbeat
var _out349 uint64
  var _ = _out349
_out349 = _234_Common____UpperBound__i_Compile.Companion_Default___.UpperBoundedAdditionImpl((clock).Dtor_t(), ((((replica).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_heartbeat__period(), ((((replica).Dtor_constants()).Dtor_all()).Dtor_params()).Dtor_max__integer__val())
_5172_heartbeat = _out349
  replica_k = func (_pat_let150_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
    return func (_5173_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
      return func (_pat_let151_0 uint64) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
        return func (_5174_dt__update_hnextHeartbeatTime_h0 uint64) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
          return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5173_dt__update__tmp_h0).Dtor_constants(), _5174_dt__update_hnextHeartbeatTime_h0, (_5173_dt__update__tmp_h0).Dtor_proposer(), (_5173_dt__update__tmp_h0).Dtor_acceptor(), (_5173_dt__update__tmp_h0).Dtor_learner(), (_5173_dt__update__tmp_h0).Dtor_executor()}}
        }(_pat_let151_0)
      }(_5172_heartbeat)
    }(_pat_let150_0)
  }(replica)
  var _5175_flag bool
  var _ = _5175_flag
  _5175_flag = ((((replica).Dtor_proposer()).Dtor_election__state()).Dtor_current__view__suspectors()).Contains(((replica).Dtor_constants()).Dtor_my__index())
  var _5176_msg _217_LiveRSL____CMessage__i_Compile.CMessage
  var _ = _5176_msg
  _5176_msg = _217_LiveRSL____CMessage__i_Compile.CMessage{_217_LiveRSL____CMessage__i_Compile.CMessage_CMessage__Heartbeat{(((replica).Dtor_proposer()).Dtor_election__state()).Dtor_current__view(), _5175_flag, ((replica).Dtor_executor()).Dtor_ops__complete()}}
  var _5177_packets _217_LiveRSL____CMessage__i_Compile.CBroadcast
  var _ = _5177_packets
var _out350 _217_LiveRSL____CMessage__i_Compile.CBroadcast
  var _ = _out350
_out350 = _301_Impl____LiveRSL____Broadcast__i_Compile.Companion_Default___.BuildBroadcastToEveryone((((replica).Dtor_constants()).Dtor_all()).Dtor_config(), ((replica).Dtor_constants()).Dtor_my__index(), _5176_msg)
_5177_packets = _out350
  { }
  packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_5177_packets}}
  return replica_k,packets__sent
}
func (_this *CompanionStruct_Default___) Replica__Next__ReadClock__MaybeSendHeartbeat(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, clock _283_LiveRSL____CClockReading__i_Compile.CClockReading) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  var _5178_start__time uint64
  var _ = _5178_start__time
var _out351 uint64
  var _ = _out351
_out351 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5178_start__time = _out351
  if (((clock).Dtor_t()) >= ((replica).Dtor_nextHeartbeatTime())) {
    var _out352 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
    var _ = _out352
var _out353 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
    var _ = _out353
_out352,_out353 = Companion_Default___.ReplicaNextReadClockMaybeSendHeartbeatActual(replica, clock)
replica_k = _out352
packets__sent = _out353
    var _5179_end__time uint64
    var _ = _5179_end__time
var _out354 uint64
    var _ = _out354
_out354 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5179_end__time = _out354
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_ReadClock_MaybeSendHeartbeat_work"), _5178_start__time, _5179_end__time)
  } else {
    var _out355 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
    var _ = _out355
var _out356 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
    var _ = _out356
_out355,_out356 = Companion_Default___.ReplicaNextReadClockMaybeSendHeartbeatSkip(replica, clock)
replica_k = _out355
packets__sent = _out356
    var _5180_end__time uint64
    var _ = _5180_end__time
var _out357 uint64
    var _ = _out357
_out357 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5180_end__time = _out357
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_ReadClock_MaybeSendHeartbeat_nada"), _5178_start__time, _5180_end__time)
  }
  return replica_k,packets__sent
}
func (_this *CompanionStruct_Default___) ReplicaNextSpontaneousMaybeMakeDecisionSkip(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  replica_k = replica
  packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}}}
  { }
  { }
  { }
  { }
  { }
  return replica_k,packets__sent
}
func (_this *CompanionStruct_Default___) ReplicaNextSpontaneousMaybeMakeDecisionActual(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  { }
  { }
  { }
  var _5181_opn _214_LiveRSL____CTypes__i_Compile.COperationNumber
  var _ = _5181_opn
  _5181_opn = ((replica).Dtor_executor()).Dtor_ops__complete()
  { }
  var _5182_candValue _dafny.Seq
  var _ = _5182_candValue
  _5182_candValue = ((((replica).Dtor_learner()).Dtor_unexecuted__ops()).Get(_5181_opn).(_278_LiveRSL____LearnerState__i_Compile.CLearnerTuple)).Dtor_candidate__learned__value()
  { }
  { }
  var _5183_newExecutor _276_LiveRSL____ExecutorState__i_Compile.ExecutorState
  var _ = _5183_newExecutor
var _out358 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState
  var _ = _out358
_out358 = _321_LiveRSL____ExecutorModel__i_Compile.Companion_Default___.ExecutorGetDecision((replica).Dtor_executor(), ((replica).Dtor_learner()).Dtor_max__ballot__seen(), _5181_opn, _5182_candValue)
_5183_newExecutor = _out358
  replica_k = func (_pat_let152_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
    return func (_5184_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
      return func (_pat_let153_0 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
        return func (_5185_dt__update_hexecutor_h0 _276_LiveRSL____ExecutorState__i_Compile.ExecutorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
          return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5184_dt__update__tmp_h0).Dtor_constants(), (_5184_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), (_5184_dt__update__tmp_h0).Dtor_proposer(), (_5184_dt__update__tmp_h0).Dtor_acceptor(), (_5184_dt__update__tmp_h0).Dtor_learner(), _5185_dt__update_hexecutor_h0}}
        }(_pat_let153_0)
      }(_5183_newExecutor)
    }(_pat_let152_0)
  }(replica)
  packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}}}
  { }
  { }
  { }
  { }
  { }
  { }
  return replica_k,packets__sent
}
func (_this *CompanionStruct_Default___) Replica__Next__Spontaneous__MaybeMakeDecision(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  var _5186_start__time uint64
  var _ = _5186_start__time
var _out359 uint64
  var _ = _out359
_out359 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5186_start__time = _out359
  var _5187_opn _214_LiveRSL____CTypes__i_Compile.COperationNumber
  var _ = _5187_opn
  _5187_opn = ((replica).Dtor_executor()).Dtor_ops__complete()
  var _5188_minCQS uint64
  var _ = _5188_minCQS
var _out360 uint64
  var _ = _out360
_out360 = _294_LiveRSL____MinCQuorumSize__i_Compile.Companion_Default___.MinCQuorumSize(((((replica).Dtor_learner()).Dtor_rcs()).Dtor_all()).Dtor_config())
_5188_minCQS = _out360
  if ((((((replica).Dtor_executor()).Dtor_next__op__to__execute()).Is_COutstandingOpUnknown()) && ((((replica).Dtor_learner()).Dtor_unexecuted__ops()).Contains(_5187_opn))) && (((((((replica).Dtor_learner()).Dtor_unexecuted__ops()).Get(_5187_opn).(_278_LiveRSL____LearnerState__i_Compile.CLearnerTuple)).Dtor_received__2b__message__senders()).Cardinality()).Cmp(_dafny.IntOfUint64(_5188_minCQS)) >= 0)) {
    var _out361 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
    var _ = _out361
var _out362 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
    var _ = _out362
_out361,_out362 = Companion_Default___.ReplicaNextSpontaneousMaybeMakeDecisionActual(replica)
replica_k = _out361
packets__sent = _out362
    var _5189_end__time uint64
    var _ = _5189_end__time
var _out363 uint64
    var _ = _out363
_out363 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5189_end__time = _out363
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Spontaneous_MaybeMakeDecision_work"), _5186_start__time, _5189_end__time)
  } else {
    var _out364 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
    var _ = _out364
var _out365 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
    var _ = _out365
_out364,_out365 = Companion_Default___.ReplicaNextSpontaneousMaybeMakeDecisionSkip(replica)
replica_k = _out364
packets__sent = _out365
    var _5190_end__time uint64
    var _ = _5190_end__time
var _out366 uint64
    var _ = _out366
_out366 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5190_end__time = _out366
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Spontaneous_MaybeMakeDecision_nada"), _5186_start__time, _5190_end__time)
  }
  return replica_k,packets__sent
}
func (_this *CompanionStruct_Default___) ReplicaNextSpontaneousTruncateLogBasedOnCheckpointsSkip(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, newLogTruncationPoint _214_LiveRSL____CTypes__i_Compile.COperationNumber) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
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
func (_this *CompanionStruct_Default___) ReplicaNextSpontaneousTruncateLogBasedOnCheckpointsActual(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState, newLogTruncationPoint _214_LiveRSL____CTypes__i_Compile.COperationNumber) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  { }
  var _5191_newAcceptor _269_LiveRSL____AcceptorState__i_Compile.AcceptorState
  var _ = _5191_newAcceptor
var _out367 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState
  var _ = _out367
_out367 = _312_LiveRSL____AcceptorModel__i_Compile.Companion_Default___.NextAcceptorState__TruncateLog((replica).Dtor_acceptor(), newLogTruncationPoint)
_5191_newAcceptor = _out367
  replica_k = func (_pat_let154_0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
    return func (_5192_dt__update__tmp_h0 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
      return func (_pat_let155_0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
        return func (_5193_dt__update_hacceptor_h0 _269_LiveRSL____AcceptorState__i_Compile.AcceptorState) _285_LiveRSL____ReplicaState__i_Compile.ReplicaState {
          return _285_LiveRSL____ReplicaState__i_Compile.ReplicaState{_285_LiveRSL____ReplicaState__i_Compile.ReplicaState_ReplicaState{(_5192_dt__update__tmp_h0).Dtor_constants(), (_5192_dt__update__tmp_h0).Dtor_nextHeartbeatTime(), (_5192_dt__update__tmp_h0).Dtor_proposer(), _5193_dt__update_hacceptor_h0, (_5192_dt__update__tmp_h0).Dtor_learner(), (_5192_dt__update__tmp_h0).Dtor_executor()}}
        }(_pat_let155_0)
      }(_5191_newAcceptor)
    }(_pat_let154_0)
  }(replica)
  packets__sent = _217_LiveRSL____CMessage__i_Compile.OutboundPackets{_217_LiveRSL____CMessage__i_Compile.OutboundPackets_Broadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast{_217_LiveRSL____CMessage__i_Compile.CBroadcast_CBroadcastNop{}}}}
  return replica_k,packets__sent
}
func (_this *CompanionStruct_Default___) Replica__Next__Spontaneous__TruncateLogBasedOnCheckpoints(replica _285_LiveRSL____ReplicaState__i_Compile.ReplicaState) (_285_LiveRSL____ReplicaState__i_Compile.ReplicaState, _217_LiveRSL____CMessage__i_Compile.OutboundPackets) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var replica_k _285_LiveRSL____ReplicaState__i_Compile.ReplicaState = _285_LiveRSL____ReplicaState__i_Compile.Type_ReplicaState_().Default().(_285_LiveRSL____ReplicaState__i_Compile.ReplicaState)
  var _ = replica_k
var packets__sent _217_LiveRSL____CMessage__i_Compile.OutboundPackets = _217_LiveRSL____CMessage__i_Compile.Type_OutboundPackets_().Default().(_217_LiveRSL____CMessage__i_Compile.OutboundPackets)
  var _ = packets__sent
  var _5194_start__time uint64
  var _ = _5194_start__time
var _out368 uint64
  var _ = _out368
_out368 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5194_start__time = _out368
  var _5195_minCQS uint64
  var _ = _5195_minCQS
var _out369 uint64
  var _ = _out369
_out369 = _294_LiveRSL____MinCQuorumSize__i_Compile.Companion_Default___.MinCQuorumSize(((((replica).Dtor_acceptor()).Dtor_constants()).Dtor_all()).Dtor_config())
_5195_minCQS = _out369
  var _5196_newLogTruncationPoint _214_LiveRSL____CTypes__i_Compile.COperationNumber
  var _ = _5196_newLogTruncationPoint
var _out370 _214_LiveRSL____CTypes__i_Compile.COperationNumber
  var _ = _out370
_out370 = _312_LiveRSL____AcceptorModel__i_Compile.Companion_Default___.AcceptorModel__GetNthHighestValueAmongReportedCheckpoints((replica).Dtor_acceptor(), _5195_minCQS)
_5196_newLogTruncationPoint = _out370
  if (((_5196_newLogTruncationPoint).Dtor_n()) > ((((replica).Dtor_acceptor()).Dtor_log__truncation__point()).Dtor_n())) {
    var _out371 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
    var _ = _out371
var _out372 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
    var _ = _out372
_out371,_out372 = Companion_Default___.ReplicaNextSpontaneousTruncateLogBasedOnCheckpointsActual(replica, _5196_newLogTruncationPoint)
replica_k = _out371
packets__sent = _out372
    var _5197_end__time uint64
    var _ = _5197_end__time
var _out373 uint64
    var _ = _out373
_out373 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5197_end__time = _out373
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Spontaneous_TruncateLogBasedOnCheckpoints_work"), _5194_start__time, _5197_end__time)
  } else {
    var _out374 _285_LiveRSL____ReplicaState__i_Compile.ReplicaState
    var _ = _out374
var _out375 _217_LiveRSL____CMessage__i_Compile.OutboundPackets
    var _ = _out375
_out374,_out375 = Companion_Default___.ReplicaNextSpontaneousTruncateLogBasedOnCheckpointsSkip(replica, _5196_newLogTruncationPoint)
replica_k = _out374
packets__sent = _out375
    var _5198_end__time uint64
    var _ = _5198_end__time
var _out376 uint64
    var _ = _out376
_out376 = _9_Native____Io__s_Compile.Companion_Time_.GetDebugTimeTicks()
_5198_end__time = _out376
    _170_Common____Util__i_Compile.Companion_Default___.RecordTimingSeq(_dafny.SeqOfString("Replica_Next_Spontaneous_TruncateLogBasedOnCheckpoints_nada"), _5194_start__time, _5198_end__time)
  }
  return replica_k,packets__sent
}
// End of class Default__
