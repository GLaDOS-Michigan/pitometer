// Package _107_Common____GenericMarshalling__i_Compile
// Dafny module _107_Common____GenericMarshalling__i_Compile compiled into Go

package _107_Common____GenericMarshalling__i_Compile

import (
	_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
	_101_Common____Util__i_Compile "101_Common____Util__i_Compile_"
	_105_Common____MarshallInt__i_Compile "105_Common____MarshallInt__i_Compile_"
	_26_Collections____Seqs__s_Compile "26_Collections____Seqs__s_Compile_"
	_29_Collections____Sets__i_Compile "29_Collections____Sets__i_Compile_"
	_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
	_33_Types__i_Compile "33_Types__i_Compile_"
	_36_Protocol__Node__i_Compile "36_Protocol__Node__i_Compile_"
	_39_Message__i_Compile "39_Message__i_Compile_"
	_42_Common____UdpClient__i_Compile "42_Common____UdpClient__i_Compile_"
	_44_Logic____Option__i_Compile "44_Logic____Option__i_Compile_"
	_47_Collections____Maps__i_Compile "47_Collections____Maps__i_Compile_"
	_50_Collections____Seqs__i_Compile "50_Collections____Seqs__i_Compile_"
	_54_Native____NativeTypes__i_Compile "54_Native____NativeTypes__i_Compile_"
	_57_Libraries____base__s_Compile "57_Libraries____base__s_Compile_"
	_59_Math____power2__s_Compile "59_Math____power2__s_Compile_"
	_5_Temporal____Temporal__s_Compile "5_Temporal____Temporal__s_Compile_"
	_61_Math____power__s_Compile "61_Math____power__s_Compile_"
	_64_Math____mul__nonlinear__i_Compile "64_Math____mul__nonlinear__i_Compile_"
	_67_Math____mul__auto__proofs__i_Compile "67_Math____mul__auto__proofs__i_Compile_"
	_69_Math____mul__auto__i_Compile "69_Math____mul__auto__i_Compile_"
	_71_Math____mul__i_Compile "71_Math____mul__i_Compile_"
	_73_Math____power__i_Compile "73_Math____power__i_Compile_"
	_77_Math____div__def__i_Compile "77_Math____div__def__i_Compile_"
	_7_Environment__s_Compile "7_Environment__s_Compile_"
	_81_Math____div__boogie__i_Compile "81_Math____div__boogie__i_Compile_"
	_83_Math____div__nonlinear__i_Compile "83_Math____div__nonlinear__i_Compile_"
	_88_Math____mod__auto__proofs__i_Compile "88_Math____mod__auto__proofs__i_Compile_"
	_90_Math____mod__auto__i_Compile "90_Math____mod__auto__i_Compile_"
	_93_Math____div__auto__proofs__i_Compile "93_Math____div__auto__proofs__i_Compile_"
	_95_Math____div__auto__i_Compile "95_Math____div__auto__i_Compile_"
	_97_Math____div__i_Compile "97_Math____div__i_Compile_"
	_99_Math____power2__i_Compile "99_Math____power2__i_Compile_"
	_9_Native____Io__s_Compile "9_Native____Io__s_Compile_"
	_System "System_"
	_dafny "dafny"
)

var _ _dafny.Dummy__
var _ _System.Dummy__
var _ _0_Native____NativeTypes__s_Compile.Dummy__
var _ _2_Collections____Maps2__s_Compile.Dummy__
var _ _5_Temporal____Temporal__s_Compile.Dummy__
var _ _7_Environment__s_Compile.Dummy__
var _ _9_Native____Io__s_Compile.Dummy__
var _ _26_Collections____Seqs__s_Compile.Dummy__
var _ _29_Collections____Sets__i_Compile.Dummy__
var _ _33_Types__i_Compile.Dummy__
var _ _36_Protocol__Node__i_Compile.Dummy__
var _ _39_Message__i_Compile.Dummy__
var _ _42_Common____UdpClient__i_Compile.Dummy__
var _ _44_Logic____Option__i_Compile.Dummy__
var _ _47_Collections____Maps__i_Compile.Dummy__
var _ _50_Collections____Seqs__i_Compile.Dummy__
var _ _54_Native____NativeTypes__i_Compile.Dummy__
var _ _57_Libraries____base__s_Compile.Dummy__
var _ _59_Math____power2__s_Compile.Dummy__
var _ _61_Math____power__s_Compile.Dummy__
var _ _64_Math____mul__nonlinear__i_Compile.Dummy__
var _ _67_Math____mul__auto__proofs__i_Compile.Dummy__
var _ _69_Math____mul__auto__i_Compile.Dummy__
var _ _71_Math____mul__i_Compile.Dummy__
var _ _73_Math____power__i_Compile.Dummy__
var _ _77_Math____div__def__i_Compile.Dummy__
var _ _81_Math____div__boogie__i_Compile.Dummy__
var _ _83_Math____div__nonlinear__i_Compile.Dummy__
var _ _88_Math____mod__auto__proofs__i_Compile.Dummy__
var _ _90_Math____mod__auto__i_Compile.Dummy__
var _ _93_Math____div__auto__proofs__i_Compile.Dummy__
var _ _95_Math____div__auto__i_Compile.Dummy__
var _ _97_Math____div__i_Compile.Dummy__
var _ _99_Math____power2__i_Compile.Dummy__
var _ _101_Common____Util__i_Compile.Dummy__
var _ _105_Common____MarshallInt__i_Compile.Dummy__

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

type CompanionStruct_G_ struct{}

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
	Elt G
}

func (G_GArray) isG() {}

func (CompanionStruct_G_) Create_GArray_(Elt G) G {
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

func (_this G) Dtor_elt() G {
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
	case nil:
		return "null"
	case G_GUint64:
		{
			return "G.GUint64"
		}
	case G_GArray:
		{
			return "G.GArray" + "(" + _dafny.String(data.Elt) + ")"
		}
	case G_GTuple:
		{
			return "G.GTuple" + "(" + _dafny.String(data.T) + ")"
		}
	case G_GByteArray:
		{
			return "G.GByteArray"
		}
	case G_GTaggedUnion:
		{
			return "G.GTaggedUnion" + "(" + _dafny.String(data.Cases) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this G) Equals(other G) bool {
	switch data1 := _this.Get().(type) {
	case G_GUint64:
		{
			_, ok := other.Get().(G_GUint64)
			return ok
		}
	case G_GArray:
		{
			data2, ok := other.Get().(G_GArray)
			return ok && data1.Elt.Equals(data2.Elt)
		}
	case G_GTuple:
		{
			data2, ok := other.Get().(G_GTuple)
			return ok && data1.T.Equals(data2.T)
		}
	case G_GByteArray:
		{
			_, ok := other.Get().(G_GByteArray)
			return ok
		}
	case G_GTaggedUnion:
		{
			data2, ok := other.Get().(G_GTaggedUnion)
			return ok && data1.Cases.Equals(data2.Cases)
		}
	default:
		{
			return false // unexpected
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
	return "G"
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

type CompanionStruct_V_ struct{}

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
	C   uint64
	Val V
}

func (V_VCase) isV() {}

func (CompanionStruct_V_) Create_VCase_(C uint64, Val V) V {
	return V{V_VCase{C, Val}}
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

func (_this V) Dtor_val() V {
	return _this.Get().(V_VCase).Val
}

func (_this V) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case V_VUint64:
		{
			return "V.VUint64" + "(" + _dafny.String(data.U) + ")"
		}
	case V_VArray:
		{
			return "V.VArray" + "(" + _dafny.String(data.A) + ")"
		}
	case V_VTuple:
		{
			return "V.VTuple" + "(" + _dafny.String(data.T) + ")"
		}
	case V_VByteArray:
		{
			return "V.VByteArray" + "(" + _dafny.String(data.B) + ")"
		}
	case V_VCase:
		{
			return "V.VCase" + "(" + _dafny.String(data.C) + ", " + _dafny.String(data.Val) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this V) Equals(other V) bool {
	switch data1 := _this.Get().(type) {
	case V_VUint64:
		{
			data2, ok := other.Get().(V_VUint64)
			return ok && data1.U == data2.U
		}
	case V_VArray:
		{
			data2, ok := other.Get().(V_VArray)
			return ok && data1.A.Equals(data2.A)
		}
	case V_VTuple:
		{
			data2, ok := other.Get().(V_VTuple)
			return ok && data1.T.Equals(data2.T)
		}
	case V_VByteArray:
		{
			data2, ok := other.Get().(V_VByteArray)
			return ok && data1.B.Equals(data2.B)
		}
	case V_VCase:
		{
			data2, ok := other.Get().(V_VCase)
			return ok && data1.C == data2.C && data1.Val.Equals(data2.Val)
		}
	default:
		{
			return false // unexpected
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
	return "V"
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

type CompanionStruct_ContentsTraceStep_ struct{}

var Companion_ContentsTraceStep_ = CompanionStruct_ContentsTraceStep_{}

type ContentsTraceStep_ContentsTraceStep struct {
	Data _dafny.Seq
	Val  _44_Logic____Option__i_Compile.Option
}

func (ContentsTraceStep_ContentsTraceStep) isContentsTraceStep() {}

func (CompanionStruct_ContentsTraceStep_) Create_ContentsTraceStep_(Data _dafny.Seq, Val _44_Logic____Option__i_Compile.Option) ContentsTraceStep {
	return ContentsTraceStep{ContentsTraceStep_ContentsTraceStep{Data, Val}}
}

func (_this ContentsTraceStep) Is_ContentsTraceStep() bool {
	_, ok := _this.Get().(ContentsTraceStep_ContentsTraceStep)
	return ok
}

func (_this ContentsTraceStep) Dtor_data() _dafny.Seq {
	return _this.Get().(ContentsTraceStep_ContentsTraceStep).Data
}

func (_this ContentsTraceStep) Dtor_val() _44_Logic____Option__i_Compile.Option {
	return _this.Get().(ContentsTraceStep_ContentsTraceStep).Val
}

func (_this ContentsTraceStep) String() string {
	switch data := _this.Get().(type) {
	case nil:
		return "null"
	case ContentsTraceStep_ContentsTraceStep:
		{
			return "_107_Common____GenericMarshalling__i_Compile.ContentsTraceStep.ContentsTraceStep" + "(" + _dafny.String(data.Data) + ", " + _dafny.String(data.Val) + ")"
		}
	default:
		{
			return "<unexpected>"
		}
	}
}

func (_this ContentsTraceStep) Equals(other ContentsTraceStep) bool {
	switch data1 := _this.Get().(type) {
	case ContentsTraceStep_ContentsTraceStep:
		{
			data2, ok := other.Get().(ContentsTraceStep_ContentsTraceStep)
			return ok && data1.Data.Equals(data2.Data) && data1.Val.Equals(data2.Val)
		}
	default:
		{
			return false // unexpected
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
	return ContentsTraceStep{ContentsTraceStep_ContentsTraceStep{_dafny.EmptySeq, _44_Logic____Option__i_Compile.Type_Option_().Default().(_44_Logic____Option__i_Compile.Option)}}
}

func (_this type_ContentsTraceStep_) String() string {
	return "_107_Common____GenericMarshalling__i_Compile.ContentsTraceStep"
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

var Companion_Default___ = CompanionStruct_Default___{}

func (_this *Default__) Equals(other *Default__) bool {
	return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Default__)
	return ok && _this.Equals(other)
}

func (*Default__) String() string {
	return "_107_Common____GenericMarshalling__i_Compile.Default__"
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
	return "_107_Common____GenericMarshalling__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) Parse__Uint64(data _dafny.Seq) _System.Tuple2 {
	if (((data).CardinalityInt()).Uint64()) >= (_54_Native____NativeTypes__i_Compile.Companion_Default___.Uint64Size()) {
		return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_Some{V{V_VUint64{_101_Common____Util__i_Compile.Companion_Default___.SeqByteToUint64((data).Subseq(_dafny.NilInt, _dafny.IntOfAny(_54_Native____NativeTypes__i_Compile.Companion_Default___.Uint64Size())))}}}}, (data).Subseq(_dafny.IntOfAny(_54_Native____NativeTypes__i_Compile.Companion_Default___.Uint64Size()), _dafny.NilInt))
	} else {
		return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
	}
}
func (_this *CompanionStruct_Default___) ParseUint64(data *_dafny.Array, index uint64) (bool, V, uint64) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var success bool = false
	var _ = success
	var v V = Type_V_().Default().(V)
	var _ = v
	var rest__index uint64 = 0
	var _ = rest__index
	{
	}
	if ((uint64((data).LenInt(0))) >= (uint64(8))) && ((index) <= ((uint64((data).LenInt(0))) - (func() uint64 { return (uint64(8)) })())) {
		var _1578_result uint64
		var _ = _1578_result
		_1578_result = ((((((((uint64((*(data).Index((_dafny.IntOfAny((index) + (uint64(0)))))).(uint8))) * (uint64(72057594037927936))) + ((uint64((*(data).Index(_dafny.IntOfAny((index) + (uint64(1))))).(uint8))) * (uint64(281474976710656)))) + ((uint64((*(data).Index(_dafny.IntOfAny((index) + (uint64(2))))).(uint8))) * (uint64(1099511627776)))) + ((uint64((*(data).Index(_dafny.IntOfAny((index) + (uint64(3))))).(uint8))) * (uint64(4294967296)))) + ((uint64((*(data).Index(_dafny.IntOfAny((index) + (uint64(4))))).(uint8))) * (uint64(16777216)))) + ((uint64((*(data).Index(_dafny.IntOfAny((index) + (uint64(5))))).(uint8))) * (uint64(65536)))) + ((uint64((*(data).Index(_dafny.IntOfAny((index) + (uint64(6))))).(uint8))) * (uint64(256)))) + (uint64((*(data).Index(_dafny.IntOfAny((index) + (uint64(7))))).(uint8)))
		success = true
		v = V{V_VUint64{_1578_result}}
		rest__index = (index) + (_54_Native____NativeTypes__i_Compile.Companion_Default___.Uint64Size())
	} else {
		success = false
		rest__index = uint64((data).LenInt(0))
	}
	return success, v, rest__index
}
func (_this *CompanionStruct_Default___) Parse__Array__contents(data _dafny.Seq, eltType G, len_ uint64) _System.Tuple2 {
	if (len_) == (uint64(0)) {
		return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_Some{_dafny.SeqOf()}}, data)
	} else {
		var _let_tmp_rhs0 _System.Tuple2 = Companion_Default___.Parse__Val(data, eltType)
		var _ = _let_tmp_rhs0
		var _1579_val _44_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs0).IndexInt(0)).(_44_Logic____Option__i_Compile.Option)
		var _ = _1579_val
		var _1580_rest1 _dafny.Seq = (*(_let_tmp_rhs0).IndexInt(1)).(_dafny.Seq)
		var _ = _1580_rest1
		var _let_tmp_rhs1 _System.Tuple2 = Companion_Default___.Parse__Array__contents(_1580_rest1, eltType, (len_)-(func() uint64 { return (uint64(1)) })())
		var _ = _let_tmp_rhs1
		var _1581_others _44_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs1).IndexInt(0)).(_44_Logic____Option__i_Compile.Option)
		var _ = _1581_others
		var _1582_rest2 _dafny.Seq = (*(_let_tmp_rhs1).IndexInt(1)).(_dafny.Seq)
		var _ = _1582_rest2
		if (!((_1579_val).Is_None())) && (!((_1581_others).Is_None())) {
			return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_Some{(_dafny.SeqOf((_1579_val).Dtor_v().(V))).Concat((_1581_others).Dtor_v().(_dafny.Seq))}}, _1582_rest2)
		} else {
			return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
		}
	}
}
func (_this *CompanionStruct_Default___) ParseArrayContents(data *_dafny.Array, index uint64, eltType G, len_ uint64) (bool, _dafny.Seq, uint64) {
	var success bool = false
	var _ = success
	var v _dafny.Seq = _dafny.EmptySeq
	var _ = v
	var rest__index uint64 = 0
	var _ = rest__index
	{
	}
	var _1583_vArr *_dafny.Array
	var _ = _1583_vArr
	var _nw3 = _dafny.NewArrayWithValue(Type_V_().Default().(V), _dafny.IntOfAny(len_))
	var _ = _nw3
	_1583_vArr = _nw3
	{
	}
	success = true
	var _1584_i uint64
	var _ = _1584_i
	_1584_i = uint64(0)
	var _1585_next__val__index uint64
	var _ = _1585_next__val__index
	_1585_next__val__index = index
	{
	}
	for (_1584_i) < (len_) {
		var _1586_some1 bool
		var _ = _1586_some1
		var _1587_val V
		var _ = _1587_val
		var _1588_rest1 uint64
		var _ = _1588_rest1
		var _out0 bool
		var _ = _out0
		var _out1 V
		var _ = _out1
		var _out2 uint64
		var _ = _out2
		_out0, _out1, _out2 = Companion_Default___.ParseVal(data, _1585_next__val__index, eltType)
		_1586_some1 = _out0
		_1587_val = _out1
		_1588_rest1 = _out2
		{
		}
		{
		}
		{
		}
		if !(_1586_some1) {
			success = false
			rest__index = uint64((data).LenInt(0))
			{
			}
			return success, v, rest__index
		}
		{
		}
		*((_1583_vArr).Index(_dafny.IntOfAny((_1584_i)))) = _1587_val
		_1585_next__val__index = _1588_rest1
		_1584_i = (_1584_i) + (uint64(1))
	}
	success = true
	rest__index = _1585_next__val__index
	v = (_1583_vArr).RangeToSeq(_dafny.NilInt, _dafny.NilInt)
	{
	}
	return success, v, rest__index
}
func (_this *CompanionStruct_Default___) Parse__Array(data _dafny.Seq, eltType G) _System.Tuple2 {
	var _let_tmp_rhs2 _System.Tuple2 = Companion_Default___.Parse__Uint64(data)
	var _ = _let_tmp_rhs2
	var _1589_len _44_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs2).IndexInt(0)).(_44_Logic____Option__i_Compile.Option)
	var _ = _1589_len
	var _1590_rest _dafny.Seq = (*(_let_tmp_rhs2).IndexInt(1)).(_dafny.Seq)
	var _ = _1590_rest
	if !((_1589_len).Is_None()) {
		var _let_tmp_rhs3 _System.Tuple2 = Companion_Default___.Parse__Array__contents(_1590_rest, eltType, ((_1589_len).Dtor_v().(V)).Dtor_u())
		var _ = _let_tmp_rhs3
		var _1591_contents _44_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs3).IndexInt(0)).(_44_Logic____Option__i_Compile.Option)
		var _ = _1591_contents
		var _1592_remainder _dafny.Seq = (*(_let_tmp_rhs3).IndexInt(1)).(_dafny.Seq)
		var _ = _1592_remainder
		if !((_1591_contents).Is_None()) {
			return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_Some{V{V_VArray{(_1591_contents).Dtor_v().(_dafny.Seq)}}}}, _1592_remainder)
		} else {
			return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
		}
	} else {
		return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
	}
}
func (_this *CompanionStruct_Default___) ParseArray(data *_dafny.Array, index uint64, eltType G) (bool, V, uint64) {
	var success bool = false
	var _ = success
	var v V = Type_V_().Default().(V)
	var _ = v
	var rest__index uint64 = 0
	var _ = rest__index
	var _1593_some1 bool
	var _ = _1593_some1
	var _1594_len V
	var _ = _1594_len
	var _1595_rest uint64
	var _ = _1595_rest
	var _out3 bool
	var _ = _out3
	var _out4 V
	var _ = _out4
	var _out5 uint64
	var _ = _out5
	_out3, _out4, _out5 = Companion_Default___.ParseUint64(data, index)
	_1593_some1 = _out3
	_1594_len = _out4
	_1595_rest = _out5
	if _1593_some1 {
		var _1596_some2 bool
		var _ = _1596_some2
		var _1597_contents _dafny.Seq
		var _ = _1597_contents
		var _1598_remainder uint64
		var _ = _1598_remainder
		var _out6 bool
		var _ = _out6
		var _out7 _dafny.Seq
		var _ = _out7
		var _out8 uint64
		var _ = _out8
		_out6, _out7, _out8 = Companion_Default___.ParseArrayContents(data, _1595_rest, eltType, (_1594_len).Dtor_u())
		_1596_some2 = _out6
		_1597_contents = _out7
		_1598_remainder = _out8
		if _1596_some2 {
			success = true
			v = V{V_VArray{_1597_contents}}
			rest__index = _1598_remainder
		} else {
			success = false
			rest__index = uint64((data).LenInt(0))
		}
	} else {
		success = false
		rest__index = uint64((data).LenInt(0))
	}
	return success, v, rest__index
}
func (_this *CompanionStruct_Default___) Parse__Tuple__contents(data _dafny.Seq, eltTypes _dafny.Seq) _System.Tuple2 {
	if (eltTypes).Equals(_dafny.SeqOf()) {
		return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_Some{_dafny.SeqOf()}}, data)
	} else {
		var _let_tmp_rhs4 _System.Tuple2 = Companion_Default___.Parse__Val(data, (eltTypes).IndexUint64(uint64(0)).(G))
		var _ = _let_tmp_rhs4
		var _1599_val _44_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs4).IndexInt(0)).(_44_Logic____Option__i_Compile.Option)
		var _ = _1599_val
		var _1600_rest1 _dafny.Seq = (*(_let_tmp_rhs4).IndexInt(1)).(_dafny.Seq)
		var _ = _1600_rest1
		var _let_tmp_rhs5 _System.Tuple2 = Companion_Default___.Parse__Tuple__contents(_1600_rest1, (eltTypes).Subseq(_dafny.IntOfAny(uint64(1)), _dafny.NilInt))
		var _ = _let_tmp_rhs5
		var _1601_contents _44_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs5).IndexInt(0)).(_44_Logic____Option__i_Compile.Option)
		var _ = _1601_contents
		var _1602_rest2 _dafny.Seq = (*(_let_tmp_rhs5).IndexInt(1)).(_dafny.Seq)
		var _ = _1602_rest2
		if (!((_1599_val).Is_None())) && (!((_1601_contents).Is_None())) {
			return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_Some{(_dafny.SeqOf((_1599_val).Dtor_v().(V))).Concat((_1601_contents).Dtor_v().(_dafny.Seq))}}, _1602_rest2)
		} else {
			return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
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
	{
	}
	var _1603_vArr *_dafny.Array
	var _ = _1603_vArr
	var _nw4 = _dafny.NewArrayWithValue(Type_V_().Default().(V), (eltTypes).CardinalityInt())
	var _ = _nw4
	_1603_vArr = _nw4
	{
	}
	success = true
	var _1604_i uint64
	var _ = _1604_i
	_1604_i = uint64(0)
	var _1605_next__val__index uint64
	var _ = _1605_next__val__index
	_1605_next__val__index = index
	{
	}
	for (_1604_i) < (uint64((eltTypes).CardinalityInt64())) {
		var _1606_some1 bool
		var _ = _1606_some1
		var _1607_val V
		var _ = _1607_val
		var _1608_rest1 uint64
		var _ = _1608_rest1
		var _out9 bool
		var _ = _out9
		var _out10 V
		var _ = _out10
		var _out11 uint64
		var _ = _out11
		_out9, _out10, _out11 = Companion_Default___.ParseVal(data, _1605_next__val__index, (eltTypes).Index(_dafny.IntOfAny(_1604_i)).(G))
		_1606_some1 = _out9
		_1607_val = _out10
		_1608_rest1 = _out11
		{
		}
		{
		}
		{
		}
		if !(_1606_some1) {
			success = false
			rest__index = uint64((data).LenInt(0))
			{
			}
			return success, v, rest__index
		}
		{
		}
		*((_1603_vArr).Index(_dafny.IntOfAny((_1604_i)))) = _1607_val
		_1605_next__val__index = _1608_rest1
		_1604_i = (_1604_i) + (uint64(1))
	}
	success = true
	rest__index = _1605_next__val__index
	v = (_1603_vArr).RangeToSeq(_dafny.NilInt, _dafny.NilInt)
	{
	}
	return success, v, rest__index
}
func (_this *CompanionStruct_Default___) Parse__Tuple(data _dafny.Seq, eltTypes _dafny.Seq) _System.Tuple2 {
	var _let_tmp_rhs6 _System.Tuple2 = Companion_Default___.Parse__Tuple__contents(data, eltTypes)
	var _ = _let_tmp_rhs6
	var _1609_contents _44_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs6).IndexInt(0)).(_44_Logic____Option__i_Compile.Option)
	var _ = _1609_contents
	var _1610_rest _dafny.Seq = (*(_let_tmp_rhs6).IndexInt(1)).(_dafny.Seq)
	var _ = _1610_rest
	if !((_1609_contents).Is_None()) {
		return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_Some{V{V_VTuple{(_1609_contents).Dtor_v().(_dafny.Seq)}}}}, _1610_rest)
	} else {
		return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
	}
}
func (_this *CompanionStruct_Default___) ParseTuple(data *_dafny.Array, index uint64, eltTypes _dafny.Seq) (bool, V, uint64) {
	var success bool = false
	var _ = success
	var v V = Type_V_().Default().(V)
	var _ = v
	var rest__index uint64 = 0
	var _ = rest__index
	var _1611_some bool
	var _ = _1611_some
	var _1612_contents _dafny.Seq
	var _ = _1612_contents
	var _1613_rest uint64
	var _ = _1613_rest
	var _out12 bool
	var _ = _out12
	var _out13 _dafny.Seq
	var _ = _out13
	var _out14 uint64
	var _ = _out14
	_out12, _out13, _out14 = Companion_Default___.ParseTupleContents(data, index, eltTypes)
	_1611_some = _out12
	_1612_contents = _out13
	_1613_rest = _out14
	if _1611_some {
		success = true
		v = V{V_VTuple{_1612_contents}}
		rest__index = _1613_rest
	} else {
		success = false
		rest__index = uint64((data).LenInt(0))
	}
	return success, v, rest__index
}
func (_this *CompanionStruct_Default___) Parse__ByteArray(data _dafny.Seq) _System.Tuple2 {
	var _let_tmp_rhs7 _System.Tuple2 = Companion_Default___.Parse__Uint64(data)
	var _ = _let_tmp_rhs7
	var _1614_len _44_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs7).IndexInt(0)).(_44_Logic____Option__i_Compile.Option)
	var _ = _1614_len
	var _1615_rest _dafny.Seq = (*(_let_tmp_rhs7).IndexInt(1)).(_dafny.Seq)
	var _ = _1615_rest
	if (!((_1614_len).Is_None())) && ((((_1614_len).Dtor_v().(V)).Dtor_u()) <= (uint64((_1615_rest).CardinalityInt64()))) {
		return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_Some{V{V_VByteArray{(_1615_rest).Subseq(_dafny.IntOfAny(uint64(0)), _dafny.IntOfAny(((_1614_len).Dtor_v().(V)).Dtor_u()))}}}}, (_1615_rest).Subseq(_dafny.IntOfAny(((_1614_len).Dtor_v().(V)).Dtor_u()), _dafny.NilInt))
	} else {
		return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
	}
}
func (_this *CompanionStruct_Default___) ParseByteArray(data *_dafny.Array, index uint64) (bool, V, uint64) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var success bool = false
	var _ = success
	var v V = Type_V_().Default().(V)
	var _ = v
	var rest__index uint64 = 0
	var _ = rest__index
	var _1616_some bool
	var _ = _1616_some
	var _1617_len V
	var _ = _1617_len
	var _1618_rest uint64
	var _ = _1618_rest
	var _out15 bool
	var _ = _out15
	var _out16 V
	var _ = _out16
	var _out17 uint64
	var _ = _out17
	_out15, _out16, _out17 = Companion_Default___.ParseUint64(data, index)
	_1616_some = _out15
	_1617_len = _out16
	_1618_rest = _out17
	if (_1616_some) && (((_1617_len).Dtor_u()) <= ((uint64((data).LenInt(0))) - (func() uint64 { return (_1618_rest) })())) {
		var _1619_rest__seq _dafny.Seq
		var _ = _1619_rest__seq
		_1619_rest__seq = (data).RangeToSeq(_dafny.IntOfAny(_1618_rest), _dafny.NilInt)
		{
		}
		{
		}
		success = true
		v = V{V_VByteArray{(data).RangeToSeq(_dafny.IntOfAny(_1618_rest), _dafny.IntOfAny((_1618_rest)+((_1617_len).Dtor_u())))}}
		rest__index = (_1618_rest) + ((_1617_len).Dtor_u())
	} else {
		success = false
		rest__index = uint64((data).LenInt(0))
	}
	return success, v, rest__index
}
func (_this *CompanionStruct_Default___) Parse__Case(data _dafny.Seq, cases _dafny.Seq) _System.Tuple2 {
	var _let_tmp_rhs8 _System.Tuple2 = Companion_Default___.Parse__Uint64(data)
	var _ = _let_tmp_rhs8
	var _1620_caseID _44_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs8).IndexInt(0)).(_44_Logic____Option__i_Compile.Option)
	var _ = _1620_caseID
	var _1621_rest1 _dafny.Seq = (*(_let_tmp_rhs8).IndexInt(1)).(_dafny.Seq)
	var _ = _1621_rest1
	if (!((_1620_caseID).Is_None())) && ((((_1620_caseID).Dtor_v().(V)).Dtor_u()) < (uint64((cases).CardinalityInt64()))) {
		var _let_tmp_rhs9 _System.Tuple2 = Companion_Default___.Parse__Val(_1621_rest1, (cases).Index(_dafny.IntOfAny(((_1620_caseID).Dtor_v().(V)).Dtor_u())).(G))
		var _ = _let_tmp_rhs9
		var _1622_val _44_Logic____Option__i_Compile.Option = (*(_let_tmp_rhs9).IndexInt(0)).(_44_Logic____Option__i_Compile.Option)
		var _ = _1622_val
		var _1623_rest2 _dafny.Seq = (*(_let_tmp_rhs9).IndexInt(1)).(_dafny.Seq)
		var _ = _1623_rest2
		if !((_1622_val).Is_None()) {
			return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_Some{V{V_VCase{((_1620_caseID).Dtor_v().(V)).Dtor_u(), (_1622_val).Dtor_v().(V)}}}}, _1623_rest2)
		} else {
			return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
		}
	} else {
		return _dafny.TupleOf(_44_Logic____Option__i_Compile.Option{_44_Logic____Option__i_Compile.Option_None{}}, _dafny.SeqOf())
	}
}
func (_this *CompanionStruct_Default___) ParseCase(data *_dafny.Array, index uint64, cases _dafny.Seq) (bool, V, uint64) {
	var success bool = false
	var _ = success
	var v V = Type_V_().Default().(V)
	var _ = v
	var rest__index uint64 = 0
	var _ = rest__index
	var _1624_some1 bool
	var _ = _1624_some1
	var _1625_caseID V
	var _ = _1625_caseID
	var _1626_rest1 uint64
	var _ = _1626_rest1
	var _out18 bool
	var _ = _out18
	var _out19 V
	var _ = _out19
	var _out20 uint64
	var _ = _out20
	_out18, _out19, _out20 = Companion_Default___.ParseUint64(data, index)
	_1624_some1 = _out18
	_1625_caseID = _out19
	_1626_rest1 = _out20
	if (_1624_some1) && (((_1625_caseID).Dtor_u()) < (uint64((cases).CardinalityInt64()))) {
		var _1627_some2 bool
		var _ = _1627_some2
		var _1628_val V
		var _ = _1628_val
		var _1629_rest2 uint64
		var _ = _1629_rest2
		var _out21 bool
		var _ = _out21
		var _out22 V
		var _ = _out22
		var _out23 uint64
		var _ = _out23
		_out21, _out22, _out23 = Companion_Default___.ParseVal(data, _1626_rest1, (cases).IndexUint64((_1625_caseID).Dtor_u()).(G))
		_1627_some2 = _out21
		_1628_val = _out22
		_1629_rest2 = _out23
		if _1627_some2 {
			success = true
			v = V{V_VCase{(_1625_caseID).Dtor_u(), _1628_val}}
			rest__index = _1629_rest2
		} else {
			success = false
			rest__index = uint64((data).LenInt(0))
		}
	} else {
		success = false
		rest__index = uint64((data).LenInt(0))
	}
	return success, v, rest__index
}
func (_this *CompanionStruct_Default___) Parse__Val(data _dafny.Seq, grammar G) _System.Tuple2 {
	var _source0 G = grammar
	var _ = _source0
	if _source0.Is_GUint64() {
		return Companion_Default___.Parse__Uint64(data)
	} else if _source0.Is_GArray() {
		var _1630_elt G = _source0.Get().(G_GArray).Elt
		var _ = _1630_elt
		return Companion_Default___.Parse__Array(data, _1630_elt)
	} else if _source0.Is_GTuple() {
		var _1631_t _dafny.Seq = _source0.Get().(G_GTuple).T
		var _ = _1631_t
		return Companion_Default___.Parse__Tuple(data, _1631_t)
	} else if _source0.Is_GByteArray() {
		return Companion_Default___.Parse__ByteArray(data)
	} else {
		var _1632_cases _dafny.Seq = _source0.Get().(G_GTaggedUnion).Cases
		var _ = _1632_cases
		return Companion_Default___.Parse__Case(data, _1632_cases)
	}
}
func (_this *CompanionStruct_Default___) ParseVal(data *_dafny.Array, index uint64, grammar G) (bool, V, uint64) {
	var success bool = false
	var _ = success
	var v V = Type_V_().Default().(V)
	var _ = v
	var rest__index uint64 = 0
	var _ = rest__index
	{
	}
	var _source1 G = grammar
	var _ = _source1
	if _source1.Is_GUint64() {
		var _out24 bool
		var _ = _out24
		var _out25 V
		var _ = _out25
		var _out26 uint64
		var _ = _out26
		_out24, _out25, _out26 = Companion_Default___.ParseUint64(data, index)
		success = _out24
		v = _out25
		rest__index = _out26
	} else if _source1.Is_GArray() {
		var _1633_elt G = _source1.Get().(G_GArray).Elt
		var _ = _1633_elt
		var _out27 bool
		var _ = _out27
		var _out28 V
		var _ = _out28
		var _out29 uint64
		var _ = _out29
		_out27, _out28, _out29 = Companion_Default___.ParseArray(data, index, _1633_elt)
		success = _out27
		v = _out28
		rest__index = _out29
	} else if _source1.Is_GTuple() {
		var _1634_t _dafny.Seq = _source1.Get().(G_GTuple).T
		var _ = _1634_t
		var _out30 bool
		var _ = _out30
		var _out31 V
		var _ = _out31
		var _out32 uint64
		var _ = _out32
		_out30, _out31, _out32 = Companion_Default___.ParseTuple(data, index, _1634_t)
		success = _out30
		v = _out31
		rest__index = _out32
	} else if _source1.Is_GByteArray() {
		var _out33 bool
		var _ = _out33
		var _out34 V
		var _ = _out34
		var _out35 uint64
		var _ = _out35
		_out33, _out34, _out35 = Companion_Default___.ParseByteArray(data, index)
		success = _out33
		v = _out34
		rest__index = _out35
	} else {
		var _1635_cases _dafny.Seq = _source1.Get().(G_GTaggedUnion).Cases
		var _ = _1635_cases
		var _out36 bool
		var _ = _out36
		var _out37 V
		var _ = _out37
		var _out38 uint64
		var _ = _out38
		_out36, _out37, _out38 = Companion_Default___.ParseCase(data, index, _1635_cases)
		success = _out36
		v = _out37
		rest__index = _out38
	}
	return success, v, rest__index
}
func (_this *CompanionStruct_Default___) Demarshall(data *_dafny.Array, grammar G) (bool, V) {
	// fmt.Printf("TONY DEBUG: demarshalling %v\n", data.String())
	var success bool = false
	var _ = success
	var v V = Type_V_().Default().(V)
	var _ = v
	var _1636_rest uint64 = 0
	var _ = _1636_rest
	var _out39 bool
	var _ = _out39
	var _out40 V
	var _ = _out40
	var _out41 uint64
	var _ = _out41
	_out39, _out40, _out41 = Companion_Default___.ParseVal(data, uint64(0), grammar)
	success = _out39
	v = _out40
	_1636_rest = _out41
	if (success) && ((_1636_rest) == (uint64((data).LenInt(0)))) {
		{
		}
		{
		}
		{
		}
	} else {
		success = false
		{
		}
	}
	return success, v
}
func (_this *CompanionStruct_Default___) ComputeSeqSum(s _dafny.Seq) uint64 {
	var size uint64 = 0
	var _ = size
	{
	}
	if (uint64((s).CardinalityInt64())) == (uint64(0)) {
		size = uint64(0)
	} else {
		var _1637_v__size uint64
		var _ = _1637_v__size
		var _out42 uint64
		var _ = _out42
		_out42 = Companion_Default___.ComputeSizeOf((s).IndexUint64(uint64(0)).(V))
		_1637_v__size = _out42
		var _1638_rest__size uint64
		var _ = _1638_rest__size
		var _out43 uint64
		var _ = _out43
		_out43 = Companion_Default___.ComputeSeqSum((s).Subseq(_dafny.IntOfAny(uint64(1)), _dafny.NilInt))
		_1638_rest__size = _out43
		size = (_1637_v__size) + (_1638_rest__size)
	}
	return size
}
func (_this *CompanionStruct_Default___) ComputeSizeOf(val V) uint64 {
	var size uint64 = 0
	var _ = size
	var _source2 V = val
	var _ = _source2
	if _source2.Is_VUint64() {
		var _1639___v3 uint64 = _source2.Get().(V_VUint64).U
		var _ = _1639___v3
		size = uint64(8)
	} else if _source2.Is_VArray() {
		var _1640_a _dafny.Seq = _source2.Get().(V_VArray).A
		var _ = _1640_a
		var _1641_v uint64
		var _ = _1641_v
		var _out44 uint64
		var _ = _out44
		_out44 = Companion_Default___.ComputeSeqSum(_1640_a)
		_1641_v = _out44
		if (_1641_v) == (uint64(0)) {
			size = uint64(8)
		} else {
			size = (uint64(8)) + (_1641_v)
		}
	} else if _source2.Is_VTuple() {
		var _1642_t _dafny.Seq = _source2.Get().(V_VTuple).T
		var _ = _1642_t
		var _out45 uint64
		var _ = _out45
		_out45 = Companion_Default___.ComputeSeqSum(_1642_t)
		size = _out45
	} else if _source2.Is_VByteArray() {
		var _1643_b _dafny.Seq = _source2.Get().(V_VByteArray).B
		var _ = _1643_b
		size = (uint64(8)) + (uint64((_1643_b).CardinalityInt64()))
	} else {
		var _1644_c uint64 = _source2.Get().(V_VCase).C
		var _ = _1644_c
		var _1645_v V = _source2.Get().(V_VCase).Val
		var _ = _1645_v
		var _1646_vs uint64
		var _ = _1646_vs
		var _out46 uint64
		var _ = _out46
		_out46 = Companion_Default___.ComputeSizeOf(_1645_v)
		_1646_vs = _out46
		size = (uint64(8)) + (_1646_vs)
	}
	return size
}
func (_this *CompanionStruct_Default___) MarshallUint64(n uint64, data *_dafny.Array, index uint64) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var _1647_tuple _System.Tuple2
	var _ = _1647_tuple
	_1647_tuple = Companion_Default___.Parse__Uint64((data).RangeToSeq(_dafny.IntOfAny(index), _dafny.NilInt))
	_105_Common____MarshallInt__i_Compile.Companion_Default___.MarshallUint64__guts(n, data, index)
}
func (_this *CompanionStruct_Default___) MarshallArrayContents(contents _dafny.Seq, eltType G, data *_dafny.Array, index uint64) uint64 {
	var size uint64 = 0
	var _ = size
	var _1648_i uint64
	var _ = _1648_i
	_1648_i = uint64(0)
	var _1649_cur__index uint64
	var _ = _1649_cur__index
	_1649_cur__index = index
	{
	}
	{
	}
	{
	}
	{
	}
	for (_1648_i) < (uint64((contents).CardinalityInt64())) {
		{
		}
		{
		}
		var _1650_item__size uint64
		var _ = _1650_item__size
		var _out47 uint64
		var _ = _out47
		_out47 = Companion_Default___.MarshallVal((contents).IndexUint64(_1648_i).(V), eltType, data, _1649_cur__index)
		_1650_item__size = _out47
		{
		}
		{
		}
		{
		}
		{
		}
		{
		}
		{
		}
		_1649_cur__index = (_1649_cur__index) + (_1650_item__size)
		_1648_i = (_1648_i) + (uint64(1))
		{
		}
		{
		}
		{
		}
		{
		}
		{
		}
		{
		}
	}
	{
	}
	{
	}
	{
	}
	{
	}
	size = (_1649_cur__index) - (func() uint64 { return (index) })()
	return size
}
func (_this *CompanionStruct_Default___) MarshallArray(val V, grammar G, data *_dafny.Array, index uint64) uint64 {
	var size uint64 = 0
	var _ = size
	{
	}
	Companion_Default___.MarshallUint64(uint64(((val).Dtor_a()).CardinalityInt64()), data, index)
	{
	}
	{
	}
	{
	}
	{
	}
	var _1651_contents__size uint64
	var _ = _1651_contents__size
	var _out48 uint64
	var _ = _out48
	_out48 = Companion_Default___.MarshallArrayContents((val).Dtor_a(), (grammar).Dtor_elt(), data, (index)+(_54_Native____NativeTypes__i_Compile.Companion_Default___.Uint64Size()))
	_1651_contents__size = _out48
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	size = (uint64(8)) + (_1651_contents__size)
	return size
}
func (_this *CompanionStruct_Default___) MarshallTupleContents(contents _dafny.Seq, eltTypes _dafny.Seq, data *_dafny.Array, index uint64) uint64 {
	var size uint64 = 0
	var _ = size
	var _1652_i uint64
	var _ = _1652_i
	_1652_i = uint64(0)
	var _1653_cur__index uint64
	var _ = _1653_cur__index
	_1653_cur__index = index
	{
	}
	{
	}
	{
	}
	{
	}
	for (_1652_i) < (uint64((contents).CardinalityInt64())) {
		{
		}
		{
		}
		{
		}
		{
		}
		{
		}
		var _1654_item__size uint64
		var _ = _1654_item__size
		var _out49 uint64
		var _ = _out49
		_out49 = Companion_Default___.MarshallVal((contents).IndexUint64(_1652_i).(V), (eltTypes).IndexUint64(_1652_i).(G), data, _1653_cur__index)
		_1654_item__size = _out49
		{
		}
		{
		}
		{
		}
		{
		}
		{
		}
		{
		}
		_1653_cur__index = (_1653_cur__index) + (_1654_item__size)
		_1652_i = (_1652_i) + (uint64(1))
		{
		}
		{
		}
		{
		}
		{
		}
		{
		}
	}
	{
	}
	{
	}
	{
	}
	{
	}
	size = (_1653_cur__index) - (func() uint64 { return (index) })()
	return size
}
func (_this *CompanionStruct_Default___) MarshallTuple(val V, grammar G, data *_dafny.Array, index uint64) uint64 {
	var size uint64 = 0
	var _ = size
	var _out50 uint64
	var _ = _out50
	_out50 = Companion_Default___.MarshallTupleContents((val).Dtor_t(), (grammar).Dtor_t(), data, index)
	size = _out50
	{
	}
	return size
}
func (_this *CompanionStruct_Default___) MarshallBytes(bytes _dafny.Seq, data *_dafny.Array, index uint64) {
	goto TAIL_CALL_START
TAIL_CALL_START:
	// _9_Native____Io__s_Compile.Companion_Arrays_.CopySeqIntoArray(bytes, uint64(0), data, index, uint64((bytes).CardinalityInt64()))
}
func (_this *CompanionStruct_Default___) MarshallByteArray(val V, grammar G, data *_dafny.Array, index uint64) uint64 {
	var size uint64 = 0
	var _ = size
	Companion_Default___.MarshallUint64(uint64(((val).Dtor_b()).CardinalityInt64()), data, index)
	{
	}
	Companion_Default___.MarshallBytes((val).Dtor_b(), data, (index)+(uint64(8)))
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	size = (uint64(8)) + (uint64(((val).Dtor_b()).CardinalityInt64()))
	return size
}
func (_this *CompanionStruct_Default___) MarshallCase(val V, grammar G, data *_dafny.Array, index uint64) uint64 {
	var size uint64 = 0
	var _ = size
	Companion_Default___.MarshallUint64((val).Dtor_c(), data, index)
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	var _1655_val__size uint64
	var _ = _1655_val__size
	var _out51 uint64
	var _ = _out51
	_out51 = Companion_Default___.MarshallVal((val).Dtor_val(), ((grammar).Dtor_cases()).IndexUint64((val).Dtor_c()).(G), data, (index)+(uint64(8)))
	_1655_val__size = _out51
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	{
	}
	size = (uint64(8)) + (_1655_val__size)
	return size
}
func (_this *CompanionStruct_Default___) MarshallVUint64(val V, grammar G, data *_dafny.Array, index uint64) uint64 {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var size uint64 = 0
	var _ = size
	Companion_Default___.MarshallUint64((val).Dtor_u(), data, index)
	{
	}
	size = uint64(8)
	return size
	return size
}
func (_this *CompanionStruct_Default___) MarshallVal(val V, grammar G, data *_dafny.Array, index uint64) uint64 {
	var size uint64 = 0
	var _ = size
	var _source3 V = val
	var _ = _source3
	if _source3.Is_VUint64() {
		var _1656___v4 uint64 = _source3.Get().(V_VUint64).U
		var _ = _1656___v4
		var _out52 uint64
		var _ = _out52
		_out52 = Companion_Default___.MarshallVUint64(val, grammar, data, index)
		size = _out52
	} else if _source3.Is_VArray() {
		var _1657___v5 _dafny.Seq = _source3.Get().(V_VArray).A
		var _ = _1657___v5
		var _out53 uint64
		var _ = _out53
		_out53 = Companion_Default___.MarshallArray(val, grammar, data, index)
		size = _out53
	} else if _source3.Is_VTuple() {
		var _1658___v6 _dafny.Seq = _source3.Get().(V_VTuple).T
		var _ = _1658___v6
		var _out54 uint64
		var _ = _out54
		_out54 = Companion_Default___.MarshallTuple(val, grammar, data, index)
		size = _out54
	} else if _source3.Is_VByteArray() {
		var _1659___v7 _dafny.Seq = _source3.Get().(V_VByteArray).B
		var _ = _1659___v7
		var _out55 uint64
		var _ = _out55
		_out55 = Companion_Default___.MarshallByteArray(val, grammar, data, index)
		size = _out55
	} else {
		var _1660___v8 uint64 = _source3.Get().(V_VCase).C
		var _ = _1660___v8
		var _1661___v9 V = _source3.Get().(V_VCase).Val
		var _ = _1661___v9
		var _out56 uint64
		var _ = _out56
		_out56 = Companion_Default___.MarshallCase(val, grammar, data, index)
		size = _out56
	}
	return size
}
func (_this *CompanionStruct_Default___) Marshall(val V, grammar G) *_dafny.Array {
	goto TAIL_CALL_START
TAIL_CALL_START:
	var data *_dafny.Array = _dafny.NewArrayWithValue(0, _dafny.IntOf(0))
	var _ = data
	var _1662_size uint64
	var _ = _1662_size
	var _out57 uint64
	var _ = _out57
	_out57 = Companion_Default___.ComputeSizeOf(val)
	_1662_size = _out57
	var _nw5 = _dafny.NewArrayWithValue(0, _dafny.IntOfAny(_1662_size))
	var _ = _nw5
	data = _nw5
	var _1663_computed__size uint64
	var _ = _1663_computed__size
	var _out58 uint64
	var _ = _out58
	_out58 = Companion_Default___.MarshallVal(val, grammar, data, uint64(0))
	_1663_computed__size = _out58
	{
	}
	{
	}
	return data
}

// End of class Default__
