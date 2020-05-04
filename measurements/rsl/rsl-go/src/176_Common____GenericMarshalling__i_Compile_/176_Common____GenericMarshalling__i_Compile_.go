// Package _176_Common____GenericMarshalling__i_Compile
// Dafny module _176_Common____GenericMarshalling__i_Compile compiled into Go

package _176_Common____GenericMarshalling__i_Compile

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

type Dummy__ struct{}








// Definition of data type G
type G struct {
  Data_G_
}

func (_this G) Get() Data_G_ {
  return _this.Data_G_
}

type Data_G_ interface {
  isG()
}

type CompanionStruct_G_ struct {}
var Companion_G_ = CompanionStruct_G_{}

type G_GUint64 struct {
}

func (G_GUint64) isG() {}

func (CompanionStruct_G_) Create_GUint64_() G {
  return G{G_GUint64{}}
}

func (_this G) Is_GUint64() bool {
  _, ok := _this.Get().(G_GUint64)
return ok
}

type G_GArray struct {
  Elt _176_Common____GenericMarshalling__i_Compile.G
}

func (G_GArray) isG() {}

func (CompanionStruct_G_) Create_GArray_(Elt _176_Common____GenericMarshalling__i_Compile.G) G {
  return G{G_GArray{Elt}}
}

func (_this G) Is_GArray() bool {
  _, ok := _this.Get().(G_GArray)
return ok
}

type G_GTuple struct {
  T _dafny.Seq
}

func (G_GTuple) isG() {}

func (CompanionStruct_G_) Create_GTuple_(T _dafny.Seq) G {
  return G{G_GTuple{T}}
}

func (_this G) Is_GTuple() bool {
  _, ok := _this.Get().(G_GTuple)
return ok
}

type G_GByteArray struct {
}

func (G_GByteArray) isG() {}

func (CompanionStruct_G_) Create_GByteArray_() G {
  return G{G_GByteArray{}}
}

func (_this G) Is_GByteArray() bool {
  _, ok := _this.Get().(G_GByteArray)
return ok
}

type G_GTaggedUnion struct {
  Cases _dafny.Seq
}

func (G_GTaggedUnion) isG() {}

func (CompanionStruct_G_) Create_GTaggedUnion_(Cases _dafny.Seq) G {
  return G{G_GTaggedUnion{Cases}}
}

func (_this G) Is_GTaggedUnion() bool {
  _, ok := _this.Get().(G_GTaggedUnion)
return ok
}

func (_this G) Dtor_elt() _176_Common____GenericMarshalling__i_Compile.G {
  return _this.Get().(G_GArray).Elt
}

func (_this G) Dtor_t() _dafny.Seq {
  return _this.Get().(G_GTuple).T
}

func (_this G) Dtor_cases() _dafny.Seq {
  return _this.Get().(G_GTaggedUnion).Cases
}

func (_this G) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case G_GUint64: {
      return "_176_Common____GenericMarshalling__i_Compile.G.GUint64"
    }
    case G_GArray: {
      return "_176_Common____GenericMarshalling__i_Compile.G.GArray" + "(" + _dafny.String(data.Elt) + ")"
    }
    case G_GTuple: {
      return "_176_Common____GenericMarshalling__i_Compile.G.GTuple" + "(" + _dafny.String(data.T) + ")"
    }
    case G_GByteArray: {
      return "_176_Common____GenericMarshalling__i_Compile.G.GByteArray"
    }
    case G_GTaggedUnion: {
      return "_176_Common____GenericMarshalling__i_Compile.G.GTaggedUnion" + "(" + _dafny.String(data.Cases) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this G) Equals(other G) bool {
  switch data1 := _this.Get().(type) {
    case G_GUint64: {
      _, ok := other.Get().(G_GUint64)
return ok
    }
    case G_GArray: {
      data2, ok := other.Get().(G_GArray)
return ok && data1.Elt.Equals(data2.Elt)
    }
    case G_GTuple: {
      data2, ok := other.Get().(G_GTuple)
return ok && data1.T.Equals(data2.T)
    }
    case G_GByteArray: {
      _, ok := other.Get().(G_GByteArray)
return ok
    }
    case G_GTaggedUnion: {
      data2, ok := other.Get().(G_GTaggedUnion)
return ok && data1.Cases.Equals(data2.Cases)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this G) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(G)
return ok && _this.Equals(typed)
}
func Type_G_() _dafny.Type {
  return type_G_{}
}

type type_G_ struct {
}

func (_this type_G_) Default() interface{} {
  return G{G_GUint64{}}
}

func (_this type_G_) String() string {
  return "_176_Common____GenericMarshalling__i_Compile.G"
}
// End of data type G

// Definition of data type V
type V struct {
  Data_V_
}

func (_this V) Get() Data_V_ {
  return _this.Data_V_
}

type Data_V_ interface {
  isV()
}

type CompanionStruct_V_ struct {}
var Companion_V_ = CompanionStruct_V_{}

type V_VUint64 struct {
  U uint64
}

func (V_VUint64) isV() {}

func (CompanionStruct_V_) Create_VUint64_(U uint64) V {
  return V{V_VUint64{U}}
}

func (_this V) Is_VUint64() bool {
  _, ok := _this.Get().(V_VUint64)
return ok
}

type V_VArray struct {
  A _dafny.Seq
}

func (V_VArray) isV() {}

func (CompanionStruct_V_) Create_VArray_(A _dafny.Seq) V {
  return V{V_VArray{A}}
}

func (_this V) Is_VArray() bool {
  _, ok := _this.Get().(V_VArray)
return ok
}

type V_VTuple struct {
  T _dafny.Seq
}

func (V_VTuple) isV() {}

func (CompanionStruct_V_) Create_VTuple_(T _dafny.Seq) V {
  return V{V_VTuple{T}}
}

func (_this V) Is_VTuple() bool {
  _, ok := _this.Get().(V_VTuple)
return ok
}

type V_VByteArray struct {
  B _dafny.Seq
}

func (V_VByteArray) isV() {}

func (CompanionStruct_V_) Create_VByteArray_(B _dafny.Seq) V {
  return V{V_VByteArray{B}}
}

func (_this V) Is_VByteArray() bool {
  _, ok := _this.Get().(V_VByteArray)
return ok
}

type V_VCase struct {
  C uint64
Val _176_Common____GenericMarshalling__i_Compile.V
}

func (V_VCase) isV() {}

func (CompanionStruct_V_) Create_VCase_(C uint64, Val _176_Common____GenericMarshalling__i_Compile.V) V {
  return V{V_VCase{C,Val}}
}

func (_this V) Is_VCase() bool {
  _, ok := _this.Get().(V_VCase)
return ok
}

func (_this V) Dtor_u() uint64 {
  return _this.Get().(V_VUint64).U
}

func (_this V) Dtor_a() _dafny.Seq {
  return _this.Get().(V_VArray).A
}

func (_this V) Dtor_t() _dafny.Seq {
  return _this.Get().(V_VTuple).T
}

func (_this V) Dtor_b() _dafny.Seq {
  return _this.Get().(V_VByteArray).B
}

func (_this V) Dtor_c() uint64 {
  return _this.Get().(V_VCase).C
}

func (_this V) Dtor_val() _176_Common____GenericMarshalling__i_Compile.V {
  return _this.Get().(V_VCase).Val
}

func (_this V) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case V_VUint64: {
      return "_176_Common____GenericMarshalling__i_Compile.V.VUint64" + "(" + _dafny.String(data.U) + ")"
    }
    case V_VArray: {
      return "_176_Common____GenericMarshalling__i_Compile.V.VArray" + "(" + _dafny.String(data.A) + ")"
    }
    case V_VTuple: {
      return "_176_Common____GenericMarshalling__i_Compile.V.VTuple" + "(" + _dafny.String(data.T) + ")"
    }
    case V_VByteArray: {
      return "_176_Common____GenericMarshalling__i_Compile.V.VByteArray" + "(" + _dafny.String(data.B) + ")"
    }
    case V_VCase: {
      return "_176_Common____GenericMarshalling__i_Compile.V.VCase" + "(" + _dafny.String(data.C) + ", " + _dafny.String(data.Val) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this V) Equals(other V) bool {
  switch data1 := _this.Get().(type) {
    case V_VUint64: {
      data2, ok := other.Get().(V_VUint64)
return ok && data1.U == data2.U
    }
    case V_VArray: {
      data2, ok := other.Get().(V_VArray)
return ok && data1.A.Equals(data2.A)
    }
    case V_VTuple: {
      data2, ok := other.Get().(V_VTuple)
return ok && data1.T.Equals(data2.T)
    }
    case V_VByteArray: {
      data2, ok := other.Get().(V_VByteArray)
return ok && data1.B.Equals(data2.B)
    }
    case V_VCase: {
      data2, ok := other.Get().(V_VCase)
return ok && data1.C == data2.C && data1.Val.Equals(data2.Val)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this V) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(V)
return ok && _this.Equals(typed)
}
func Type_V_() _dafny.Type {
  return type_V_{}
}

type type_V_ struct {
}

func (_this type_V_) Default() interface{} {
  return V{V_VUint64{0}}
}

func (_this type_V_) String() string {
  return "_176_Common____GenericMarshalling__i_Compile.V"
}
// End of data type V

// Definition of data type ContentsTraceStep
type ContentsTraceStep struct {
  Data_ContentsTraceStep_
}

func (_this ContentsTraceStep) Get() Data_ContentsTraceStep_ {
  return _this.Data_ContentsTraceStep_
}

type Data_ContentsTraceStep_ interface {
  isContentsTraceStep()
}

type CompanionStruct_ContentsTraceStep_ struct {}
var Companion_ContentsTraceStep_ = CompanionStruct_ContentsTraceStep_{}

type ContentsTraceStep_ContentsTraceStep struct {
  Data _dafny.Seq
Val _135_Logic____Option__i_Compile.Option
}

func (ContentsTraceStep_ContentsTraceStep) isContentsTraceStep() {}

func (CompanionStruct_ContentsTraceStep_) Create_ContentsTraceStep_(Data _dafny.Seq, Val _135_Logic____Option__i_Compile.Option) ContentsTraceStep {
  return ContentsTraceStep{ContentsTraceStep_ContentsTraceStep{Data,Val}}
}

func (_this ContentsTraceStep) Is_ContentsTraceStep() bool {
  _, ok := _this.Get().(ContentsTraceStep_ContentsTraceStep)
return ok
}

func (_this ContentsTraceStep) Dtor_data() _dafny.Seq {
  return _this.Get().(ContentsTraceStep_ContentsTraceStep).Data
}

func (_this ContentsTraceStep) Dtor_val() _135_Logic____Option__i_Compile.Option {
  return _this.Get().(ContentsTraceStep_ContentsTraceStep).Val
}

func (_this ContentsTraceStep) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case ContentsTraceStep_ContentsTraceStep: {
      return "_176_Common____GenericMarshalling__i_Compile.ContentsTraceStep.ContentsTraceStep" + "(" + _dafny.String(data.Data) + ", " + _dafny.String(data.Val) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this ContentsTraceStep) Equals(other ContentsTraceStep) bool {
  switch data1 := _this.Get().(type) {
    case ContentsTraceStep_ContentsTraceStep: {
      data2, ok := other.Get().(ContentsTraceStep_ContentsTraceStep)
return ok && data1.Data.Equals(data2.Data) && data1.Val.Equals(data2.Val)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this ContentsTraceStep) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(ContentsTraceStep)
return ok && _this.Equals(typed)
}
func Type_ContentsTraceStep_() _dafny.Type {
  return type_ContentsTraceStep_{}
}

type type_ContentsTraceStep_ struct {
}

func (_this type_ContentsTraceStep_) Default() interface{} {
  return ContentsTraceStep{ContentsTraceStep_ContentsTraceStep{_dafny.EmptySeq, _135_Logic____Option__i_Compile.Type_Option_().Default().(_135_Logic____Option__i_Compile.Option)}}
}

func (_this type_ContentsTraceStep_) String() string {
  return "_176_Common____GenericMarshalling__i_Compile.ContentsTraceStep"
}
// End of data type ContentsTraceStep

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
  return "_176_Common____GenericMarshalling__i_Compile.Default__"
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
  return "_176_Common____GenericMarshalling__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) Parse__Uint64(data _dafny.Seq) _System.Tuple2 {
  if ((uint64((data).CardinalityInt())) >= (_138_Native____NativeTypes__i_Compile.Companion_Default___.Uint64Size())) {
    return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_Some{V{V_VUint64{_170_Common____Util__i_Compile.Companion_Default___.SeqByteToUint64((data).Subseq(_dafny.NilInt, _138_Native____NativeTypes__i_Compile.Companion_Default___.Uint64Size()))}}}}, (data).Subseq(_138_Native____NativeTypes__i_Compile.Companion_Default___.Uint64Size(), _dafny.NilInt))
  } else  {
    return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
  }
}
func (_this *CompanionStruct_Default___) ParseUint64(data *_dafny.Array, index uint64) (bool, _176_Common____GenericMarshalling__i_Compile.V, uint64) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var success bool = false
  var _ = success
var v _176_Common____GenericMarshalling__i_Compile.V = Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = v
var rest__index uint64 = 0
  var _ = rest__index
  { }
  if (((uint64((data).LenInt(0))) >= (uint64(8))) && ((index) <= ((uint64((data).LenInt(0))) - (func () uint64 { return  (uint64(8)) })()))) {
    var _4541_result uint64
    var _ = _4541_result
    _4541_result = ((((((((uint64((*(data).Index(((index) + (uint64(0))))).(uint8))) * (uint64(72057594037927936))) + ((uint64((*(data).Index(((index) + (uint64(1))))).(uint8))) * (uint64(281474976710656)))) + ((uint64((*(data).Index(((index) + (uint64(2))))).(uint8))) * (uint64(1099511627776)))) + ((uint64((*(data).Index(((index) + (uint64(3))))).(uint8))) * (uint64(4294967296)))) + ((uint64((*(data).Index(((index) + (uint64(4))))).(uint8))) * (uint64(16777216)))) + ((uint64((*(data).Index(((index) + (uint64(5))))).(uint8))) * (uint64(65536)))) + ((uint64((*(data).Index(((index) + (uint64(6))))).(uint8))) * (uint64(256)))) + (uint64((*(data).Index(((index) + (uint64(7))))).(uint8)))
    success = true
    v = V{V_VUint64{_4541_result}}
    rest__index = (index) + (_138_Native____NativeTypes__i_Compile.Companion_Default___.Uint64Size())
  } else {
    success = false
    rest__index = uint64((data).LenInt(0))
  }
  return success,v,rest__index
}
func (_this *CompanionStruct_Default___) Parse__Array__contents(data _dafny.Seq, eltType _176_Common____GenericMarshalling__i_Compile.G, len_ uint64) _System.Tuple2 {
  if ((len_) == (uint64(0))) {
    return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_Some{_dafny.SeqOf()}}, data)
  } else  {
    var _let_tmp_rhs0 _System.Tuple2 = Companion_Default___.Parse__Val(data, eltType)
    var _ = _let_tmp_rhs0
var _4542_val _135_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs0).IndexInt(0)).(_135_Logic____Option__i_Compile.Option)
    var _ = _4542_val
var _4543_rest1 _dafny.Seq = (*(_let_tmp_rhs0).IndexInt(1)).(_dafny.Seq)
    var _ = _4543_rest1
var _let_tmp_rhs1 _System.Tuple2 = Companion_Default___.Parse__Array__contents(_4543_rest1, eltType, (len_) - (func () uint64 { return  (uint64(1)) })())
    var _ = _let_tmp_rhs1
var _4544_others _135_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs1).IndexInt(0)).(_135_Logic____Option__i_Compile.Option)
    var _ = _4544_others
var _4545_rest2 _dafny.Seq = (*(_let_tmp_rhs1).IndexInt(1)).(_dafny.Seq)
    var _ = _4545_rest2
if ((!((_4542_val).Is_None())) && (!((_4544_others).Is_None()))) {
      return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_Some{(_dafny.SeqOf((_4542_val).Dtor_v().(_176_Common____GenericMarshalling__i_Compile.V))).Concat((_4544_others).Dtor_v().(_dafny.Seq))}}, _4545_rest2)
    } else  {
      return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
    }
  }
}
func (_this *CompanionStruct_Default___) ParseArrayContents(data *_dafny.Array, index uint64, eltType _176_Common____GenericMarshalling__i_Compile.G, len_ uint64) (bool, _dafny.Seq, uint64) {
  var success bool = false
  var _ = success
var v _dafny.Seq = _dafny.EmptySeq
  var _ = v
var rest__index uint64 = 0
  var _ = rest__index
  { }
  var _4546_vArr *_dafny.Array
  var _ = _4546_vArr
  var _nw3 = _dafny.NewArrayWithValue(Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V), len_)
  var _ = _nw3
  _4546_vArr = _nw3
  { }
  success = true
  var _4547_i uint64
  var _ = _4547_i
  _4547_i = uint64(0)
  var _4548_next__val__index uint64
  var _ = _4548_next__val__index
  _4548_next__val__index = index
  { }
  for (_4547_i) < (len_) {
    var _4549_some1 bool
    var _ = _4549_some1
var _4550_val _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4550_val
var _4551_rest1 uint64
    var _ = _4551_rest1
var _out0 bool
    var _ = _out0
var _out1 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out1
var _out2 uint64
    var _ = _out2
_out0,_out1,_out2 = Companion_Default___.ParseVal(data, _4548_next__val__index, eltType)
_4549_some1 = _out0
_4550_val = _out1
_4551_rest1 = _out2
    { }
    { }
    { }
    if (!(_4549_some1)) {
      success = false
      rest__index = uint64((data).LenInt(0))
      { }
      return success,v,rest__index
    }
    { }
    *((_4546_vArr).Index(_dafny.IntOfAny((_4547_i)))) = _4550_val
    _4548_next__val__index = _4551_rest1
    _4547_i = (_4547_i) + (uint64(1))
  }
  success = true
  rest__index = _4548_next__val__index
  v = (_4546_vArr).RangeToSeq(_dafny.NilInt, _dafny.NilInt)
  { }
  return success,v,rest__index
}
func (_this *CompanionStruct_Default___) Parse__Array(data _dafny.Seq, eltType _176_Common____GenericMarshalling__i_Compile.G) _System.Tuple2 {
  var _let_tmp_rhs2 _System.Tuple2 = Companion_Default___.Parse__Uint64(data)
  var _ = _let_tmp_rhs2
var _4552_len _135_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs2).IndexInt(0)).(_135_Logic____Option__i_Compile.Option)
  var _ = _4552_len
var _4553_rest _dafny.Seq = (*(_let_tmp_rhs2).IndexInt(1)).(_dafny.Seq)
  var _ = _4553_rest
if (!((_4552_len).Is_None())) {
    var _let_tmp_rhs3 _System.Tuple2 = Companion_Default___.Parse__Array__contents(_4553_rest, eltType, ((_4552_len).Dtor_v().(_176_Common____GenericMarshalling__i_Compile.V)).Dtor_u())
    var _ = _let_tmp_rhs3
var _4554_contents _135_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs3).IndexInt(0)).(_135_Logic____Option__i_Compile.Option)
    var _ = _4554_contents
var _4555_remainder _dafny.Seq = (*(_let_tmp_rhs3).IndexInt(1)).(_dafny.Seq)
    var _ = _4555_remainder
if (!((_4554_contents).Is_None())) {
      return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_Some{V{V_VArray{(_4554_contents).Dtor_v().(_dafny.Seq)}}}}, _4555_remainder)
    } else  {
      return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
    }
  } else  {
    return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
  }
}
func (_this *CompanionStruct_Default___) ParseArray(data *_dafny.Array, index uint64, eltType _176_Common____GenericMarshalling__i_Compile.G) (bool, _176_Common____GenericMarshalling__i_Compile.V, uint64) {
  var success bool = false
  var _ = success
var v _176_Common____GenericMarshalling__i_Compile.V = Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = v
var rest__index uint64 = 0
  var _ = rest__index
  var _4556_some1 bool
  var _ = _4556_some1
var _4557_len _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4557_len
var _4558_rest uint64
  var _ = _4558_rest
var _out3 bool
  var _ = _out3
var _out4 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out4
var _out5 uint64
  var _ = _out5
_out3,_out4,_out5 = Companion_Default___.ParseUint64(data, index)
_4556_some1 = _out3
_4557_len = _out4
_4558_rest = _out5
  if (_4556_some1) {
    var _4559_some2 bool
    var _ = _4559_some2
var _4560_contents _dafny.Seq
    var _ = _4560_contents
var _4561_remainder uint64
    var _ = _4561_remainder
var _out6 bool
    var _ = _out6
var _out7 _dafny.Seq
    var _ = _out7
var _out8 uint64
    var _ = _out8
_out6,_out7,_out8 = Companion_Default___.ParseArrayContents(data, _4558_rest, eltType, (_4557_len).Dtor_u())
_4559_some2 = _out6
_4560_contents = _out7
_4561_remainder = _out8
    if (_4559_some2) {
      success = true
      v = V{V_VArray{_4560_contents}}
      rest__index = _4561_remainder
    } else {
      success = false
      rest__index = uint64((data).LenInt(0))
    }
  } else {
    success = false
    rest__index = uint64((data).LenInt(0))
  }
  return success,v,rest__index
}
func (_this *CompanionStruct_Default___) Parse__Tuple__contents(data _dafny.Seq, eltTypes _dafny.Seq) _System.Tuple2 {
  if ((eltTypes).Equals(_dafny.SeqOf())) {
    return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_Some{_dafny.SeqOf()}}, data)
  } else  {
    var _let_tmp_rhs4 _System.Tuple2 = Companion_Default___.Parse__Val(data, (eltTypes).Index(uint64(0)).(_176_Common____GenericMarshalling__i_Compile.G))
    var _ = _let_tmp_rhs4
var _4562_val _135_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs4).IndexInt(0)).(_135_Logic____Option__i_Compile.Option)
    var _ = _4562_val
var _4563_rest1 _dafny.Seq = (*(_let_tmp_rhs4).IndexInt(1)).(_dafny.Seq)
    var _ = _4563_rest1
var _let_tmp_rhs5 _System.Tuple2 = Companion_Default___.Parse__Tuple__contents(_4563_rest1, (eltTypes).Subseq(uint64(1), _dafny.NilInt))
    var _ = _let_tmp_rhs5
var _4564_contents _135_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs5).IndexInt(0)).(_135_Logic____Option__i_Compile.Option)
    var _ = _4564_contents
var _4565_rest2 _dafny.Seq = (*(_let_tmp_rhs5).IndexInt(1)).(_dafny.Seq)
    var _ = _4565_rest2
if ((!((_4562_val).Is_None())) && (!((_4564_contents).Is_None()))) {
      return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_Some{(_dafny.SeqOf((_4562_val).Dtor_v().(_176_Common____GenericMarshalling__i_Compile.V))).Concat((_4564_contents).Dtor_v().(_dafny.Seq))}}, _4565_rest2)
    } else  {
      return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
    }
  }
}
func (_this *CompanionStruct_Default___) ParseTupleContents(data *_dafny.Array, index uint64, eltTypes _dafny.Seq) (bool, _dafny.Seq, uint64) {
  var success bool = false
  var _ = success
var v _dafny.Seq = _dafny.EmptySeq
  var _ = v
var rest__index uint64 = 0
  var _ = rest__index
  { }
  var _4566_vArr *_dafny.Array
  var _ = _4566_vArr
  var _nw4 = _dafny.NewArrayWithValue(Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V), uint64((eltTypes).CardinalityInt()))
  var _ = _nw4
  _4566_vArr = _nw4
  { }
  success = true
  var _4567_i uint64
  var _ = _4567_i
  _4567_i = uint64(0)
  var _4568_next__val__index uint64
  var _ = _4568_next__val__index
  _4568_next__val__index = index
  { }
  for (_4567_i) < (uint64((eltTypes).CardinalityInt())) {
    var _4569_some1 bool
    var _ = _4569_some1
var _4570_val _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4570_val
var _4571_rest1 uint64
    var _ = _4571_rest1
var _out9 bool
    var _ = _out9
var _out10 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out10
var _out11 uint64
    var _ = _out11
_out9,_out10,_out11 = Companion_Default___.ParseVal(data, _4568_next__val__index, (eltTypes).Index(_4567_i).(_176_Common____GenericMarshalling__i_Compile.G))
_4569_some1 = _out9
_4570_val = _out10
_4571_rest1 = _out11
    { }
    { }
    { }
    if (!(_4569_some1)) {
      success = false
      rest__index = uint64((data).LenInt(0))
      { }
      return success,v,rest__index
    }
    { }
    *((_4566_vArr).Index(_dafny.IntOfAny((_4567_i)))) = _4570_val
    _4568_next__val__index = _4571_rest1
    _4567_i = (_4567_i) + (uint64(1))
  }
  success = true
  rest__index = _4568_next__val__index
  v = (_4566_vArr).RangeToSeq(_dafny.NilInt, _dafny.NilInt)
  { }
  return success,v,rest__index
}
func (_this *CompanionStruct_Default___) Parse__Tuple(data _dafny.Seq, eltTypes _dafny.Seq) _System.Tuple2 {
  var _let_tmp_rhs6 _System.Tuple2 = Companion_Default___.Parse__Tuple__contents(data, eltTypes)
  var _ = _let_tmp_rhs6
var _4572_contents _135_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs6).IndexInt(0)).(_135_Logic____Option__i_Compile.Option)
  var _ = _4572_contents
var _4573_rest _dafny.Seq = (*(_let_tmp_rhs6).IndexInt(1)).(_dafny.Seq)
  var _ = _4573_rest
if (!((_4572_contents).Is_None())) {
    return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_Some{V{V_VTuple{(_4572_contents).Dtor_v().(_dafny.Seq)}}}}, _4573_rest)
  } else  {
    return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
  }
}
func (_this *CompanionStruct_Default___) ParseTuple(data *_dafny.Array, index uint64, eltTypes _dafny.Seq) (bool, _176_Common____GenericMarshalling__i_Compile.V, uint64) {
  var success bool = false
  var _ = success
var v _176_Common____GenericMarshalling__i_Compile.V = Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = v
var rest__index uint64 = 0
  var _ = rest__index
  var _4574_some bool
  var _ = _4574_some
var _4575_contents _dafny.Seq
  var _ = _4575_contents
var _4576_rest uint64
  var _ = _4576_rest
var _out12 bool
  var _ = _out12
var _out13 _dafny.Seq
  var _ = _out13
var _out14 uint64
  var _ = _out14
_out12,_out13,_out14 = Companion_Default___.ParseTupleContents(data, index, eltTypes)
_4574_some = _out12
_4575_contents = _out13
_4576_rest = _out14
  if (_4574_some) {
    success = true
    v = V{V_VTuple{_4575_contents}}
    rest__index = _4576_rest
  } else {
    success = false
    rest__index = uint64((data).LenInt(0))
  }
  return success,v,rest__index
}
func (_this *CompanionStruct_Default___) Parse__ByteArray(data _dafny.Seq) _System.Tuple2 {
  var _let_tmp_rhs7 _System.Tuple2 = Companion_Default___.Parse__Uint64(data)
  var _ = _let_tmp_rhs7
var _4577_len _135_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs7).IndexInt(0)).(_135_Logic____Option__i_Compile.Option)
  var _ = _4577_len
var _4578_rest _dafny.Seq = (*(_let_tmp_rhs7).IndexInt(1)).(_dafny.Seq)
  var _ = _4578_rest
if ((!((_4577_len).Is_None())) && ((((_4577_len).Dtor_v().(_176_Common____GenericMarshalling__i_Compile.V)).Dtor_u()) <= (uint64((_4578_rest).CardinalityInt())))) {
    return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_Some{V{V_VByteArray{(_4578_rest).Subseq(uint64(0), ((_4577_len).Dtor_v().(_176_Common____GenericMarshalling__i_Compile.V)).Dtor_u())}}}}, (_4578_rest).Subseq(((_4577_len).Dtor_v().(_176_Common____GenericMarshalling__i_Compile.V)).Dtor_u(), _dafny.NilInt))
  } else  {
    return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
  }
}
func (_this *CompanionStruct_Default___) ParseByteArray(data *_dafny.Array, index uint64) (bool, _176_Common____GenericMarshalling__i_Compile.V, uint64) {
  goto TAIL_CALL_START
TAIL_CALL_START:
var success bool = false
  var _ = success
var v _176_Common____GenericMarshalling__i_Compile.V = Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = v
var rest__index uint64 = 0
  var _ = rest__index
  var _4579_some bool
  var _ = _4579_some
var _4580_len _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4580_len
var _4581_rest uint64
  var _ = _4581_rest
var _out15 bool
  var _ = _out15
var _out16 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out16
var _out17 uint64
  var _ = _out17
_out15,_out16,_out17 = Companion_Default___.ParseUint64(data, index)
_4579_some = _out15
_4580_len = _out16
_4581_rest = _out17
  if ((_4579_some) && (((_4580_len).Dtor_u()) <= ((uint64((data).LenInt(0))) - (func () uint64 { return  (_4581_rest) })()))) {
    var _4582_rest__seq _dafny.Seq
    var _ = _4582_rest__seq
    _4582_rest__seq = (data).RangeToSeq(_4581_rest, _dafny.NilInt)
    { }
    { }
    success = true
    v = V{V_VByteArray{(data).RangeToSeq(_4581_rest, (_4581_rest) + ((_4580_len).Dtor_u()))}}
    rest__index = (_4581_rest) + ((_4580_len).Dtor_u())
  } else {
    success = false
    rest__index = uint64((data).LenInt(0))
  }
  return success,v,rest__index
}
func (_this *CompanionStruct_Default___) Parse__Case(data _dafny.Seq, cases _dafny.Seq) _System.Tuple2 {
  var _let_tmp_rhs8 _System.Tuple2 = Companion_Default___.Parse__Uint64(data)
  var _ = _let_tmp_rhs8
var _4583_caseID _135_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs8).IndexInt(0)).(_135_Logic____Option__i_Compile.Option)
  var _ = _4583_caseID
var _4584_rest1 _dafny.Seq = (*(_let_tmp_rhs8).IndexInt(1)).(_dafny.Seq)
  var _ = _4584_rest1
if ((!((_4583_caseID).Is_None())) && ((((_4583_caseID).Dtor_v().(_176_Common____GenericMarshalling__i_Compile.V)).Dtor_u()) < (uint64((cases).CardinalityInt())))) {
    var _let_tmp_rhs9 _System.Tuple2 = Companion_Default___.Parse__Val(_4584_rest1, (cases).Index(((_4583_caseID).Dtor_v().(_176_Common____GenericMarshalling__i_Compile.V)).Dtor_u()).(_176_Common____GenericMarshalling__i_Compile.G))
    var _ = _let_tmp_rhs9
var _4585_val _135_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs9).IndexInt(0)).(_135_Logic____Option__i_Compile.Option)
    var _ = _4585_val
var _4586_rest2 _dafny.Seq = (*(_let_tmp_rhs9).IndexInt(1)).(_dafny.Seq)
    var _ = _4586_rest2
if (!((_4585_val).Is_None())) {
      return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_Some{V{V_VCase{((_4583_caseID).Dtor_v().(_176_Common____GenericMarshalling__i_Compile.V)).Dtor_u(), (_4585_val).Dtor_v().(_176_Common____GenericMarshalling__i_Compile.V)}}}}, _4586_rest2)
    } else  {
      return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
    }
  } else  {
    return _dafny.TupleOf(_135_Logic____Option__i_Compile.Option{_135_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
  }
}
func (_this *CompanionStruct_Default___) ParseCase(data *_dafny.Array, index uint64, cases _dafny.Seq) (bool, _176_Common____GenericMarshalling__i_Compile.V, uint64) {
  var success bool = false
  var _ = success
var v _176_Common____GenericMarshalling__i_Compile.V = Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = v
var rest__index uint64 = 0
  var _ = rest__index
  var _4587_some1 bool
  var _ = _4587_some1
var _4588_caseID _176_Common____GenericMarshalling__i_Compile.V
  var _ = _4588_caseID
var _4589_rest1 uint64
  var _ = _4589_rest1
var _out18 bool
  var _ = _out18
var _out19 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out19
var _out20 uint64
  var _ = _out20
_out18,_out19,_out20 = Companion_Default___.ParseUint64(data, index)
_4587_some1 = _out18
_4588_caseID = _out19
_4589_rest1 = _out20
  if ((_4587_some1) && (((_4588_caseID).Dtor_u()) < (uint64((cases).CardinalityInt())))) {
    var _4590_some2 bool
    var _ = _4590_some2
var _4591_val _176_Common____GenericMarshalling__i_Compile.V
    var _ = _4591_val
var _4592_rest2 uint64
    var _ = _4592_rest2
var _out21 bool
    var _ = _out21
var _out22 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out22
var _out23 uint64
    var _ = _out23
_out21,_out22,_out23 = Companion_Default___.ParseVal(data, _4589_rest1, (cases).Index((_4588_caseID).Dtor_u()).(_176_Common____GenericMarshalling__i_Compile.G))
_4590_some2 = _out21
_4591_val = _out22
_4592_rest2 = _out23
    if (_4590_some2) {
      success = true
      v = V{V_VCase{(_4588_caseID).Dtor_u(), _4591_val}}
      rest__index = _4592_rest2
    } else {
      success = false
      rest__index = uint64((data).LenInt(0))
    }
  } else {
    success = false
    rest__index = uint64((data).LenInt(0))
  }
  return success,v,rest__index
}
func (_this *CompanionStruct_Default___) Parse__Val(data _dafny.Seq, grammar _176_Common____GenericMarshalling__i_Compile.G) _System.Tuple2 {
  var _source0 _176_Common____GenericMarshalling__i_Compile.G = grammar
  var _ = _source0
if (_source0.Is_GUint64()) {
    return Companion_Default___.Parse__Uint64(data)
  } else if (_source0.Is_GArray()) {
    var _4593_elt _176_Common____GenericMarshalling__i_Compile.G = _source0.Get().(_176_Common____GenericMarshalling__i_Compile.G_GArray).Elt
    var _ = _4593_elt
return Companion_Default___.Parse__Array(data, _4593_elt)
  } else if (_source0.Is_GTuple()) {
    var _4594_t _dafny.Seq = _source0.Get().(_176_Common____GenericMarshalling__i_Compile.G_GTuple).T
    var _ = _4594_t
return Companion_Default___.Parse__Tuple(data, _4594_t)
  } else if (_source0.Is_GByteArray()) {
    return Companion_Default___.Parse__ByteArray(data)
  } else {
    var _4595_cases _dafny.Seq = _source0.Get().(_176_Common____GenericMarshalling__i_Compile.G_GTaggedUnion).Cases
    var _ = _4595_cases
return Companion_Default___.Parse__Case(data, _4595_cases)
  }
}
func (_this *CompanionStruct_Default___) ParseVal(data *_dafny.Array, index uint64, grammar _176_Common____GenericMarshalling__i_Compile.G) (bool, _176_Common____GenericMarshalling__i_Compile.V, uint64) {
  var success bool = false
  var _ = success
var v _176_Common____GenericMarshalling__i_Compile.V = Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = v
var rest__index uint64 = 0
  var _ = rest__index
  { }
  var _source1 _176_Common____GenericMarshalling__i_Compile.G = grammar
  var _ = _source1
if (_source1.Is_GUint64()) {
    var _out24 bool
    var _ = _out24
var _out25 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out25
var _out26 uint64
    var _ = _out26
_out24,_out25,_out26 = Companion_Default___.ParseUint64(data, index)
success = _out24
v = _out25
rest__index = _out26
  } else if (_source1.Is_GArray()) {
    var _4596_elt _176_Common____GenericMarshalling__i_Compile.G = _source1.Get().(_176_Common____GenericMarshalling__i_Compile.G_GArray).Elt
    var _ = _4596_elt
    var _out27 bool
    var _ = _out27
var _out28 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out28
var _out29 uint64
    var _ = _out29
_out27,_out28,_out29 = Companion_Default___.ParseArray(data, index, _4596_elt)
success = _out27
v = _out28
rest__index = _out29
  } else if (_source1.Is_GTuple()) {
    var _4597_t _dafny.Seq = _source1.Get().(_176_Common____GenericMarshalling__i_Compile.G_GTuple).T
    var _ = _4597_t
    var _out30 bool
    var _ = _out30
var _out31 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out31
var _out32 uint64
    var _ = _out32
_out30,_out31,_out32 = Companion_Default___.ParseTuple(data, index, _4597_t)
success = _out30
v = _out31
rest__index = _out32
  } else if (_source1.Is_GByteArray()) {
    var _out33 bool
    var _ = _out33
var _out34 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out34
var _out35 uint64
    var _ = _out35
_out33,_out34,_out35 = Companion_Default___.ParseByteArray(data, index)
success = _out33
v = _out34
rest__index = _out35
  } else {
    var _4598_cases _dafny.Seq = _source1.Get().(_176_Common____GenericMarshalling__i_Compile.G_GTaggedUnion).Cases
    var _ = _4598_cases
    var _out36 bool
    var _ = _out36
var _out37 _176_Common____GenericMarshalling__i_Compile.V
    var _ = _out37
var _out38 uint64
    var _ = _out38
_out36,_out37,_out38 = Companion_Default___.ParseCase(data, index, _4598_cases)
success = _out36
v = _out37
rest__index = _out38
  }
  return success,v,rest__index
}
func (_this *CompanionStruct_Default___) Demarshall(data *_dafny.Array, grammar _176_Common____GenericMarshalling__i_Compile.G) (bool, _176_Common____GenericMarshalling__i_Compile.V) {
  var success bool = false
  var _ = success
var v _176_Common____GenericMarshalling__i_Compile.V = Type_V_().Default().(_176_Common____GenericMarshalling__i_Compile.V)
  var _ = v
  var _4599_rest uint64 = 0
  var _ = _4599_rest
  var _out39 bool
  var _ = _out39
var _out40 _176_Common____GenericMarshalling__i_Compile.V
  var _ = _out40
var _out41 uint64
  var _ = _out41
_out39,_out40,_out41 = Companion_Default___.ParseVal(data, uint64(0), grammar)
success = _out39
v = _out40
_4599_rest = _out41
  if ((success) && ((_4599_rest) == (uint64((data).LenInt(0))))) {
    { }
    { }
    { }
  } else {
    success = false
    { }
  }
  return success,v
}
func (_this *CompanionStruct_Default___) ComputeSeqSum(s _dafny.Seq) uint64 {
  var size uint64 = 0
  var _ = size
  { }
  if ((uint64((s).CardinalityInt())) == (uint64(0))) {
    size = uint64(0)
  } else {
    var _4600_v__size uint64
    var _ = _4600_v__size
var _out42 uint64
    var _ = _out42
_out42 = Companion_Default___.ComputeSizeOf((s).Index(uint64(0)).(_176_Common____GenericMarshalling__i_Compile.V))
_4600_v__size = _out42
    var _4601_rest__size uint64
    var _ = _4601_rest__size
var _out43 uint64
    var _ = _out43
_out43 = Companion_Default___.ComputeSeqSum((s).Subseq(uint64(1), _dafny.NilInt))
_4601_rest__size = _out43
    size = (_4600_v__size) + (_4601_rest__size)
  }
  return size
}
func (_this *CompanionStruct_Default___) ComputeSizeOf(val _176_Common____GenericMarshalling__i_Compile.V) uint64 {
  var size uint64 = 0
  var _ = size
  var _source2 _176_Common____GenericMarshalling__i_Compile.V = val
  var _ = _source2
if (_source2.Is_VUint64()) {
    var _4602___v3 uint64 = _source2.Get().(_176_Common____GenericMarshalling__i_Compile.V_VUint64).U
    var _ = _4602___v3
    size = uint64(8)
  } else if (_source2.Is_VArray()) {
    var _4603_a _dafny.Seq = _source2.Get().(_176_Common____GenericMarshalling__i_Compile.V_VArray).A
    var _ = _4603_a
    var _4604_v uint64
    var _ = _4604_v
var _out44 uint64
    var _ = _out44
_out44 = Companion_Default___.ComputeSeqSum(_4603_a)
_4604_v = _out44
    if ((_4604_v) == (uint64(0))) {
      size = uint64(8)
    } else {
      size = (uint64(8)) + (_4604_v)
    }
  } else if (_source2.Is_VTuple()) {
    var _4605_t _dafny.Seq = _source2.Get().(_176_Common____GenericMarshalling__i_Compile.V_VTuple).T
    var _ = _4605_t
    var _out45 uint64
    var _ = _out45
_out45 = Companion_Default___.ComputeSeqSum(_4605_t)
size = _out45
  } else if (_source2.Is_VByteArray()) {
    var _4606_b _dafny.Seq = _source2.Get().(_176_Common____GenericMarshalling__i_Compile.V_VByteArray).B
    var _ = _4606_b
    size = (uint64(8)) + (uint64((_4606_b).CardinalityInt()))
  } else {
    var _4607_c uint64 = _source2.Get().(_176_Common____GenericMarshalling__i_Compile.V_VCase).C
    var _ = _4607_c
var _4608_v _176_Common____GenericMarshalling__i_Compile.V = _source2.Get().(_176_Common____GenericMarshalling__i_Compile.V_VCase).Val
    var _ = _4608_v
    var _4609_vs uint64
    var _ = _4609_vs
var _out46 uint64
    var _ = _out46
_out46 = Companion_Default___.ComputeSizeOf(_4608_v)
_4609_vs = _out46
    size = (uint64(8)) + (_4609_vs)
  }
  return size
}
func (_this *CompanionStruct_Default___) MarshallUint64(n uint64, data *_dafny.Array, index uint64) {
  goto TAIL_CALL_START
TAIL_CALL_START:
  var _4610_tuple _System.Tuple2
  var _ = _4610_tuple
  _4610_tuple = Companion_Default___.Parse__Uint64((data).RangeToSeq(index, _dafny.NilInt))
  _174_Common____MarshallInt__i_Compile.Companion_Default___.MarshallUint64__guts(n, data, index)
}
func (_this *CompanionStruct_Default___) MarshallArrayContents(contents _dafny.Seq, eltType _176_Common____GenericMarshalling__i_Compile.G, data *_dafny.Array, index uint64) uint64 {
  var size uint64 = 0
  var _ = size
  var _4611_i uint64
  var _ = _4611_i
  _4611_i = uint64(0)
  var _4612_cur__index uint64
  var _ = _4612_cur__index
  _4612_cur__index = index
  { }
  { }
  { }
  { }
  for (_4611_i) < (uint64((contents).CardinalityInt())) {
    { }
    { }
    var _4613_item__size uint64
    var _ = _4613_item__size
var _out47 uint64
    var _ = _out47
_out47 = Companion_Default___.MarshallVal((contents).Index(_4611_i).(_176_Common____GenericMarshalling__i_Compile.V), eltType, data, _4612_cur__index)
_4613_item__size = _out47
    { }
    { }
    { }
    { }
    { }
    { }
    _4612_cur__index = (_4612_cur__index) + (_4613_item__size)
    _4611_i = (_4611_i) + (uint64(1))
    { }
    { }
    { }
    { }
    { }
    { }
  }
  { }
  { }
  { }
  { }
  size = (_4612_cur__index) - (func () uint64 { return  (index) })()
  return size
}
func (_this *CompanionStruct_Default___) MarshallArray(val _176_Common____GenericMarshalling__i_Compile.V, grammar _176_Common____GenericMarshalling__i_Compile.G, data *_dafny.Array, index uint64) uint64 {
  var size uint64 = 0
  var _ = size
  { }
  Companion_Default___.MarshallUint64(uint64(((val).Dtor_a()).CardinalityInt()), data, index)
  { }
  { }
  { }
  { }
  var _4614_contents__size uint64
  var _ = _4614_contents__size
var _out48 uint64
  var _ = _out48
_out48 = Companion_Default___.MarshallArrayContents((val).Dtor_a(), (grammar).Dtor_elt(), data, (index) + (_138_Native____NativeTypes__i_Compile.Companion_Default___.Uint64Size()))
_4614_contents__size = _out48
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  size = (uint64(8)) + (_4614_contents__size)
  return size
}
func (_this *CompanionStruct_Default___) MarshallTupleContents(contents _dafny.Seq, eltTypes _dafny.Seq, data *_dafny.Array, index uint64) uint64 {
  var size uint64 = 0
  var _ = size
  var _4615_i uint64
  var _ = _4615_i
  _4615_i = uint64(0)
  var _4616_cur__index uint64
  var _ = _4616_cur__index
  _4616_cur__index = index
  { }
  { }
  { }
  { }
  for (_4615_i) < (uint64((contents).CardinalityInt())) {
    { }
    { }
    { }
    { }
    { }
    var _4617_item__size uint64
    var _ = _4617_item__size
var _out49 uint64
    var _ = _out49
_out49 = Companion_Default___.MarshallVal((contents).Index(_4615_i).(_176_Common____GenericMarshalling__i_Compile.V), (eltTypes).Index(_4615_i).(_176_Common____GenericMarshalling__i_Compile.G), data, _4616_cur__index)
_4617_item__size = _out49
    { }
    { }
    { }
    { }
    { }
    { }
    _4616_cur__index = (_4616_cur__index) + (_4617_item__size)
    _4615_i = (_4615_i) + (uint64(1))
    { }
    { }
    { }
    { }
    { }
  }
  { }
  { }
  { }
  { }
  size = (_4616_cur__index) - (func () uint64 { return  (index) })()
  return size
}
func (_this *CompanionStruct_Default___) MarshallTuple(val _176_Common____GenericMarshalling__i_Compile.V, grammar _176_Common____GenericMarshalling__i_Compile.G, data *_dafny.Array, index uint64) uint64 {
  var size uint64 = 0
  var _ = size
  var _out50 uint64
  var _ = _out50
_out50 = Companion_Default___.MarshallTupleContents((val).Dtor_t(), (grammar).Dtor_t(), data, index)
size = _out50
  { }
  return size
}
func (_this *CompanionStruct_Default___) MarshallBytes(bytes _dafny.Seq, data *_dafny.Array, index uint64) {
  goto TAIL_CALL_START
TAIL_CALL_START:
  _9_Native____Io__s_Compile.Companion_Arrays_.CopySeqIntoArray(bytes, uint64(0), data, index, uint64((bytes).CardinalityInt()))
}
func (_this *CompanionStruct_Default___) MarshallByteArray(val _176_Common____GenericMarshalling__i_Compile.V, grammar _176_Common____GenericMarshalling__i_Compile.G, data *_dafny.Array, index uint64) uint64 {
  var size uint64 = 0
  var _ = size
  Companion_Default___.MarshallUint64(uint64(((val).Dtor_b()).CardinalityInt()), data, index)
  { }
  Companion_Default___.MarshallBytes((val).Dtor_b(), data, (index) + (uint64(8)))
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  size = (uint64(8)) + (uint64(((val).Dtor_b()).CardinalityInt()))
  return size
}
func (_this *CompanionStruct_Default___) MarshallCase(val _176_Common____GenericMarshalling__i_Compile.V, grammar _176_Common____GenericMarshalling__i_Compile.G, data *_dafny.Array, index uint64) uint64 {
  var size uint64 = 0
  var _ = size
  Companion_Default___.MarshallUint64((val).Dtor_c(), data, index)
  { }
  { }
  { }
  { }
  { }
  { }
  var _4618_val__size uint64
  var _ = _4618_val__size
var _out51 uint64
  var _ = _out51
_out51 = Companion_Default___.MarshallVal((val).Dtor_val(), ((grammar).Dtor_cases()).Index((val).Dtor_c()).(_176_Common____GenericMarshalling__i_Compile.G), data, (index) + (uint64(8)))
_4618_val__size = _out51
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  { }
  size = (uint64(8)) + (_4618_val__size)
  return size
}
func (_this *CompanionStruct_Default___) MarshallVUint64(val _176_Common____GenericMarshalling__i_Compile.V, grammar _176_Common____GenericMarshalling__i_Compile.G, data *_dafny.Array, index uint64) uint64 {
  goto TAIL_CALL_START
TAIL_CALL_START:
var size uint64 = 0
  var _ = size
  Companion_Default___.MarshallUint64((val).Dtor_u(), data, index)
  { }
  size = uint64(8)
return size
  return size
}
func (_this *CompanionStruct_Default___) MarshallVal(val _176_Common____GenericMarshalling__i_Compile.V, grammar _176_Common____GenericMarshalling__i_Compile.G, data *_dafny.Array, index uint64) uint64 {
  var size uint64 = 0
  var _ = size
  var _source3 _176_Common____GenericMarshalling__i_Compile.V = val
  var _ = _source3
if (_source3.Is_VUint64()) {
    var _4619___v4 uint64 = _source3.Get().(_176_Common____GenericMarshalling__i_Compile.V_VUint64).U
    var _ = _4619___v4
    var _out52 uint64
    var _ = _out52
_out52 = Companion_Default___.MarshallVUint64(val, grammar, data, index)
size = _out52
  } else if (_source3.Is_VArray()) {
    var _4620___v5 _dafny.Seq = _source3.Get().(_176_Common____GenericMarshalling__i_Compile.V_VArray).A
    var _ = _4620___v5
    var _out53 uint64
    var _ = _out53
_out53 = Companion_Default___.MarshallArray(val, grammar, data, index)
size = _out53
  } else if (_source3.Is_VTuple()) {
    var _4621___v6 _dafny.Seq = _source3.Get().(_176_Common____GenericMarshalling__i_Compile.V_VTuple).T
    var _ = _4621___v6
    var _out54 uint64
    var _ = _out54
_out54 = Companion_Default___.MarshallTuple(val, grammar, data, index)
size = _out54
  } else if (_source3.Is_VByteArray()) {
    var _4622___v7 _dafny.Seq = _source3.Get().(_176_Common____GenericMarshalling__i_Compile.V_VByteArray).B
    var _ = _4622___v7
    var _out55 uint64
    var _ = _out55
_out55 = Companion_Default___.MarshallByteArray(val, grammar, data, index)
size = _out55
  } else {
    var _4623___v8 uint64 = _source3.Get().(_176_Common____GenericMarshalling__i_Compile.V_VCase).C
    var _ = _4623___v8
var _4624___v9 _176_Common____GenericMarshalling__i_Compile.V = _source3.Get().(_176_Common____GenericMarshalling__i_Compile.V_VCase).Val
    var _ = _4624___v9
    var _out56 uint64
    var _ = _out56
_out56 = Companion_Default___.MarshallCase(val, grammar, data, index)
size = _out56
  }
  return size
}
func (_this *CompanionStruct_Default___) Marshall(val _176_Common____GenericMarshalling__i_Compile.V, grammar _176_Common____GenericMarshalling__i_Compile.G) *_dafny.Array {
  goto TAIL_CALL_START
TAIL_CALL_START:
var data *_dafny.Array = _dafny.NewArrayWithValue(0, _dafny.IntOf(0))
  var _ = data
  var _4625_size uint64
  var _ = _4625_size
var _out57 uint64
  var _ = _out57
_out57 = Companion_Default___.ComputeSizeOf(val)
_4625_size = _out57
  var _nw5 = _dafny.NewArrayWithValue(0, _4625_size)
  var _ = _nw5
  data = _nw5
  var _4626_computed__size uint64
  var _ = _4626_computed__size
var _out58 uint64
  var _ = _out58
_out58 = Companion_Default___.MarshallVal(val, grammar, data, uint64(0))
_4626_computed__size = _out58
  { }
  { }
  return data
}
// End of class Default__
