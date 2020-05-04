// Package _460_CmdLineParser__i_Compile
// Dafny module _460_CmdLineParser__i_Compile compiled into Go

package _460_CmdLineParser__i_Compile

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
  return "_460_CmdLineParser__i_Compile.Default__"
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
  return "_460_CmdLineParser__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) Ascii__to__int(short uint16) _System.Tuple2 {
  if (((uint16(48)) <= (short)) && ((short) <= (uint16(57)))) {
    return _dafny.TupleOf(true, uint8((short) - (func () uint16 { return  (uint16(48)) })()))
  } else  {
    return _dafny.TupleOf(false, uint8(0))
  }
}
func (_this *CompanionStruct_Default___) Power10(e _dafny.Int) _dafny.Int {
  var val _dafny.Int = _dafny.Zero
  var _ = val
  { }
  if (_dafny.AreEqual(e, _dafny.Zero)) {
    val = _dafny.IntOfInt64(1)
return val
  } else {
    var _5282_tmp _dafny.Int
    var _ = _5282_tmp
var _out519 _dafny.Int
    var _ = _out519
_out519 = Companion_Default___.Power10((e).Minus(_dafny.IntOfInt64(1)))
_5282_tmp = _out519
    val = (_dafny.IntOfInt64(10)).Times(_5282_tmp)
return val
  }
  return val
}
func (_this *CompanionStruct_Default___) Shorts__to__bytes(shorts _dafny.Seq) _System.Tuple2 {
  if (((shorts).Cardinality()).Cmp(_dafny.Zero) == 0) {
    return _dafny.TupleOf(true, _dafny.SeqOf())
  } else  {
    var _5283_tuple _System.Tuple2 = Companion_Default___.Shorts__to__bytes((shorts).Subseq(_dafny.IntOfInt64(1), _dafny.NilInt))
    var _ = _5283_tuple
var _5284_ok bool = (*((_5283_tuple)).IndexInt(0)).(bool)
    var _ = _5284_ok
var _5285_rest _dafny.Seq = (*((_5283_tuple)).IndexInt(1)).(_dafny.Seq)
    var _ = _5285_rest
var _5286_tuple_k _System.Tuple2 = Companion_Default___.Ascii__to__int((shorts).Index(_dafny.Zero).(uint16))
    var _ = _5286_tuple_k
var _5287_ok_k bool = (*((_5286_tuple_k)).IndexInt(0)).(bool)
    var _ = _5287_ok_k
var _5288_a__byte uint8 = (*((_5286_tuple_k)).IndexInt(1)).(uint8)
    var _ = _5288_a__byte
if ((_5284_ok) && (_5287_ok_k)) {
      return _dafny.TupleOf(true, (_dafny.SeqOf(_5288_a__byte)).Concat(_5285_rest))
    } else  {
      return _dafny.TupleOf(false, _dafny.SeqOf())
    }
  }
}
func (_this *CompanionStruct_Default___) Bytes__to__decimal(bytes _dafny.Seq) _dafny.Int {
  if (((bytes).Cardinality()).Cmp(_dafny.Zero) == 0) {
    return _dafny.Zero
  } else  {
    return (_dafny.IntOfUint8((bytes).Index(((bytes).Cardinality()).Minus(_dafny.IntOfInt64(1))).(uint8))).Plus((_dafny.IntOfInt64(10)).Times(Companion_Default___.Bytes__to__decimal((bytes).Subseq(_dafny.Zero, ((bytes).Cardinality()).Minus(_dafny.IntOfInt64(1))))))
  }
}
func (_this *CompanionStruct_Default___) Shorts__to__nat(shorts _dafny.Seq) _System.Tuple2 {
  if (((shorts).Cardinality()).Cmp(_dafny.Zero) == 0) {
    return _dafny.TupleOf(false, _dafny.Zero)
  } else  {
    var _5289_tuple _System.Tuple2 = Companion_Default___.Shorts__to__bytes(shorts)
    var _ = _5289_tuple
var _5290_ok bool = (*((_5289_tuple)).IndexInt(0)).(bool)
    var _ = _5290_ok
var _5291_bytes _dafny.Seq = (*((_5289_tuple)).IndexInt(1)).(_dafny.Seq)
    var _ = _5291_bytes
if (!(_5290_ok)) {
      return _dafny.TupleOf(false, _dafny.Zero)
    } else  {
      return _dafny.TupleOf(true, Companion_Default___.Bytes__to__decimal(_5291_bytes))
    }
  }
}
func (_this *CompanionStruct_Default___) Shorts__to__byte(shorts _dafny.Seq) _System.Tuple2 {
  var _5292_tuple _System.Tuple2 = Companion_Default___.Shorts__to__nat(shorts)
  var _ = _5292_tuple
var _5293_ok bool = (*((_5292_tuple)).IndexInt(0)).(bool)
  var _ = _5293_ok
var _5294_val _dafny.Int = (*((_5292_tuple)).IndexInt(1)).(_dafny.Int)
  var _ = _5294_val
if (((_dafny.Zero).Cmp(_5294_val) <= 0) && ((_5294_val).Cmp(_dafny.IntOfInt64(256)) < 0)) {
    return _dafny.TupleOf(true, (_5294_val).Uint8())
  } else  {
    return _dafny.TupleOf(false, uint8(0))
  }
}
func (_this *CompanionStruct_Default___) Shorts__to__uint16(shorts _dafny.Seq) _System.Tuple2 {
  var _5295_tuple _System.Tuple2 = Companion_Default___.Shorts__to__nat(shorts)
  var _ = _5295_tuple
var _5296_ok bool = (*((_5295_tuple)).IndexInt(0)).(bool)
  var _ = _5296_ok
var _5297_val _dafny.Int = (*((_5295_tuple)).IndexInt(1)).(_dafny.Int)
  var _ = _5297_val
if (((_dafny.Zero).Cmp(_5297_val) <= 0) && ((_5297_val).Cmp(_dafny.IntOfInt64(65536)) < 0)) {
    return _dafny.TupleOf(true, (_5297_val).Uint16())
  } else  {
    return _dafny.TupleOf(false, uint16(0))
  }
}
func (_this *CompanionStruct_Default___) Shorts__to__uint32(shorts _dafny.Seq) _System.Tuple2 {
  var _5298_tuple _System.Tuple2 = Companion_Default___.Shorts__to__nat(shorts)
  var _ = _5298_tuple
var _5299_ok bool = (*((_5298_tuple)).IndexInt(0)).(bool)
  var _ = _5299_ok
var _5300_val _dafny.Int = (*((_5298_tuple)).IndexInt(1)).(_dafny.Int)
  var _ = _5300_val
if (((_dafny.Zero).Cmp(_5300_val) <= 0) && ((_5300_val).Cmp(_dafny.IntOfInt64(4294967296)) < 0)) {
    return _dafny.TupleOf(true, (_5300_val).Uint32())
  } else  {
    return _dafny.TupleOf(false, uint32(0))
  }
}
func (_this *CompanionStruct_Default___) Is__ascii__period(short uint16) bool {
  return (short) == (uint16(46))
}
func (_this *CompanionStruct_Default___) Parse__ip__addr__helper(ip__shorts _dafny.Seq, current__octet__shorts _dafny.Seq) _System.Tuple2 {
  if (((ip__shorts).Cardinality()).Cmp(_dafny.Zero) == 0) {
    var _5301_tuple _System.Tuple2 = Companion_Default___.Shorts__to__byte(current__octet__shorts)
    var _ = _5301_tuple
var _5302_okay bool = (*((_5301_tuple)).IndexInt(0)).(bool)
    var _ = _5302_okay
var _5303_b uint8 = (*((_5301_tuple)).IndexInt(1)).(uint8)
    var _ = _5303_b
if (!(_5302_okay)) {
      return _dafny.TupleOf(false, _dafny.SeqOf())
    } else  {
      return _dafny.TupleOf(true, _dafny.SeqOf(_5303_b))
    }
  } else  {
    if (Companion_Default___.Is__ascii__period((ip__shorts).Index(_dafny.Zero).(uint16))) {
      var _5304_tuple _System.Tuple2 = Companion_Default___.Shorts__to__byte(current__octet__shorts)
      var _ = _5304_tuple
var _5305_okay bool = (*((_5304_tuple)).IndexInt(0)).(bool)
      var _ = _5305_okay
var _5306_b uint8 = (*((_5304_tuple)).IndexInt(1)).(uint8)
      var _ = _5306_b
if (!(_5305_okay)) {
        return _dafny.TupleOf(false, _dafny.SeqOf())
      } else  {
        var _5307_tuple_k _System.Tuple2 = Companion_Default___.Parse__ip__addr__helper((ip__shorts).Subseq(_dafny.IntOfInt64(1), _dafny.NilInt), _dafny.SeqOf())
        var _ = _5307_tuple_k
var _5308_ok bool = (*((_5307_tuple_k)).IndexInt(0)).(bool)
        var _ = _5308_ok
var _5309_ip__bytes _dafny.Seq = (*((_5307_tuple_k)).IndexInt(1)).(_dafny.Seq)
        var _ = _5309_ip__bytes
if (!(_5308_ok)) {
          return _dafny.TupleOf(false, _dafny.SeqOf())
        } else  {
          return _dafny.TupleOf(true, (_dafny.SeqOf(_5306_b)).Concat(_5309_ip__bytes))
        }
      }
    } else  {
      return Companion_Default___.Parse__ip__addr__helper((ip__shorts).Subseq(_dafny.IntOfInt64(1), _dafny.NilInt), (current__octet__shorts).Concat(_dafny.SeqOf((ip__shorts).Index(_dafny.Zero).(uint16))))
    }
  }
}
func (_this *CompanionStruct_Default___) Parse__ip__addr(ip__shorts _dafny.Seq) _System.Tuple2 {
  var _5310_tuple _System.Tuple2 = Companion_Default___.Parse__ip__addr__helper(ip__shorts, _dafny.SeqOf())
  var _ = _5310_tuple
var _5311_ok bool = (*((_5310_tuple)).IndexInt(0)).(bool)
  var _ = _5311_ok
var _5312_ip__bytes _dafny.Seq = (*((_5310_tuple)).IndexInt(1)).(_dafny.Seq)
  var _ = _5312_ip__bytes
if ((_5311_ok) && (((_5312_ip__bytes).Cardinality()).Cmp(_dafny.IntOfInt64(4)) == 0)) {
    return _dafny.TupleOf(true, _5312_ip__bytes)
  } else  {
    return _dafny.TupleOf(false, _dafny.SeqOf())
  }
}
func (_this *CompanionStruct_Default___) Parse__end__point(ip__shorts _dafny.Seq, port__shorts _dafny.Seq) _System.Tuple2 {
  var _5313_tuple _System.Tuple2 = Companion_Default___.Parse__ip__addr(ip__shorts)
  var _ = _5313_tuple
var _5314_okay bool = (*((_5313_tuple)).IndexInt(0)).(bool)
  var _ = _5314_okay
var _5315_ip__bytes _dafny.Seq = (*((_5313_tuple)).IndexInt(1)).(_dafny.Seq)
  var _ = _5315_ip__bytes
if (!(_5314_okay)) {
    return _dafny.TupleOf(false, _9_Native____Io__s_Compile.EndPoint{_9_Native____Io__s_Compile.EndPoint_EndPoint{_dafny.SeqOf(uint8(0), uint8(0), uint8(0), uint8(0)), uint16(0)}})
  } else  {
    var _5316_tuple_k _System.Tuple2 = Companion_Default___.Shorts__to__uint16(port__shorts)
    var _ = _5316_tuple_k
var _5317_okay_k bool = (*((_5316_tuple_k)).IndexInt(0)).(bool)
    var _ = _5317_okay_k
var _5318_port uint16 = (*((_5316_tuple_k)).IndexInt(1)).(uint16)
    var _ = _5318_port
if (!(_5317_okay_k)) {
      return _dafny.TupleOf(false, _9_Native____Io__s_Compile.EndPoint{_9_Native____Io__s_Compile.EndPoint_EndPoint{_dafny.SeqOf(uint8(0), uint8(0), uint8(0), uint8(0)), uint16(0)}})
    } else  {
      return _dafny.TupleOf(true, _9_Native____Io__s_Compile.EndPoint{_9_Native____Io__s_Compile.EndPoint_EndPoint{_5315_ip__bytes, _5318_port}})
    }
  }
}
func (_this *CompanionStruct_Default___) Test__unique_k(endpoints _dafny.Seq) bool {
  goto TAIL_CALL_START
TAIL_CALL_START:
var unique bool = false
  var _ = unique
  unique = true
  var _5319_i _dafny.Int
  var _ = _5319_i
  _5319_i = _dafny.Zero
  for (_5319_i).Cmp((endpoints).Cardinality()) < 0 {
    var _5320_j _dafny.Int
    var _ = _5320_j
    _5320_j = _dafny.Zero
    for (_5320_j).Cmp((endpoints).Cardinality()) < 0 {
      if (((_5319_i).Cmp(_5320_j) != 0) && (((endpoints).Index(_5319_i).(_9_Native____Io__s_Compile.EndPoint)).Equals((endpoints).Index(_5320_j).(_9_Native____Io__s_Compile.EndPoint)))) {
        unique = false
        { }
        return unique
      }
      _5320_j = (_5320_j).Plus(_dafny.IntOfInt64(1))
    }
    _5319_i = (_5319_i).Plus(_dafny.IntOfInt64(1))
  }
  { }
  return unique
}
func (_this *CompanionStruct_Default___) Parse__end__points(args _dafny.Seq) _System.Tuple2 {
  if (((args).Cardinality()).Cmp(_dafny.Zero) == 0) {
    return _dafny.TupleOf(true, _dafny.SeqOf())
  } else  {
    var _let_tmp_rhs10 _System.Tuple2 = Companion_Default___.Parse__end__point((args).Index(_dafny.Zero).(_dafny.Seq), (args).Index(_dafny.IntOfInt64(1)).(_dafny.Seq))
    var _ = _let_tmp_rhs10
var _5321_ok1 bool = (*(_let_tmp_rhs10).IndexInt(0)).(bool)
    var _ = _5321_ok1
var _5322_ep _9_Native____Io__s_Compile.EndPoint = (*(_let_tmp_rhs10).IndexInt(1)).(_9_Native____Io__s_Compile.EndPoint)
    var _ = _5322_ep
var _let_tmp_rhs11 _System.Tuple2 = Companion_Default___.Parse__end__points((args).Subseq(_dafny.IntOfInt64(2), _dafny.NilInt))
    var _ = _let_tmp_rhs11
var _5323_ok2 bool = (*(_let_tmp_rhs11).IndexInt(0)).(bool)
    var _ = _5323_ok2
var _5324_rest _dafny.Seq = (*(_let_tmp_rhs11).IndexInt(1)).(_dafny.Seq)
    var _ = _5324_rest
if (!((_5321_ok1) && (_5323_ok2))) {
      return _dafny.TupleOf(false, _dafny.SeqOf())
    } else  {
      return _dafny.TupleOf(true, (_dafny.SeqOf(_5322_ep)).Concat(_5324_rest))
    }
  }
}
func (_this *CompanionStruct_Default___) Collect__cmd__line__args() _dafny.Seq {
  goto TAIL_CALL_START
TAIL_CALL_START:
var args _dafny.Seq = _dafny.EmptySeq
  var _ = args
  var _5325_num__args uint32
  var _ = _5325_num__args
var _out520 uint32
  var _ = _out520
_out520 = _9_Native____Io__s_Compile.Companion_HostConstants_.NumCommandLineArgs()
_5325_num__args = _out520
  var _5326_i uint32
  var _ = _5326_i
  _5326_i = uint32(0)
  args = _dafny.SeqOf()
  for (_5326_i) < (_5325_num__args) {
    var _5327_arg *_dafny.Array
    var _ = _5327_arg
var _out521 *_dafny.Array
    var _ = _out521
_out521 = _9_Native____Io__s_Compile.Companion_HostConstants_.GetCommandLineArg(uint64(_5326_i))
_5327_arg = _out521
    args = (args).Concat(_dafny.SeqOf((_5327_arg).RangeToSeq(_dafny.NilInt, _dafny.NilInt)))
    _5326_i = (_5326_i) + (uint32(1))
  }
  return args
}
// End of class Default__
