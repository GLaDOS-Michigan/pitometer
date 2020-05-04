// Package _185_Common____SeqIsUnique__i_Compile
// Dafny module _185_Common____SeqIsUnique__i_Compile compiled into Go

package _185_Common____SeqIsUnique__i_Compile

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
  return "_185_Common____SeqIsUnique__i_Compile.Default__"
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
  return "_185_Common____SeqIsUnique__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) SeqToSetConstruct(xs _dafny.Seq) _dafny.Set {
  goto TAIL_CALL_START
TAIL_CALL_START:
var s _dafny.Set = _dafny.EmptySet
  var _ = s
  { }
  s = _dafny.SetOf()
  var _4627_i _dafny.Int
  var _ = _4627_i
  _4627_i = _dafny.Zero
  for (_4627_i).Cmp((xs).Cardinality()) < 0 {
    s = (s).Union(_dafny.SetOf((xs).Index(_4627_i).(interface{})))
    _4627_i = (_4627_i).Plus(_dafny.IntOfInt64(1))
  }
  return s
}
func (_this *CompanionStruct_Default___) SetToUniqueSeqConstruct(Type_X_ _dafny.Type, s _dafny.Set) _dafny.Seq {
  var xs _dafny.Seq = _dafny.EmptySeq
  var _ = xs
  var _4628_arr *_dafny.Array
  var _ = _4628_arr
  var _nw6 = _dafny.NewArrayWithValue(Type_X_.Default(), (s).Cardinality())
  var _ = _nw6
  _4628_arr = _nw6
  var _4629_s1 _dafny.Set
  var _ = _4629_s1
  _4629_s1 = s
  { }
  { }
  { }
  for ((_4629_s1).Cardinality()).Cmp(_dafny.Zero) != 0 {
    { }
    { }
    var _4630_x interface{}
    var _ = _4630_x
    for _iter2 := _dafny.Iterate((_4629_s1).Elements());; {
      _val2, _ok2 := _iter2()
if !_ok2 { break }
_assign_such_that_0 := _val2.(interface{})
_4630_x = _assign_such_that_0
if ((_4629_s1).Contains(_4630_x)) {
        goto L_ASSIGN_SUCH_THAT_0
      }
    }
    panic("assign-such-that search produced no value (line 85)")
  L_ASSIGN_SUCH_THAT_0:
    { }
    { }
    var _index7 = ((s).Cardinality()).Minus((_4629_s1).Cardinality())
    var _ = _index7
    *((_4628_arr).Index(_dafny.IntOfAny(_index7))) = _4630_x
    _4629_s1 = (_4629_s1).Difference(_dafny.SetOf(_4630_x))
    { }
    { }
    { }
  }
  xs = (_4628_arr).RangeToSeq(_dafny.NilInt, _dafny.NilInt)
  { }
  return xs
}
func (_this *CompanionStruct_Default___) SubsequenceConstruct(Type_X_ _dafny.Type, xs _dafny.Seq, f func (interface{}) bool) _dafny.Seq {
  var xs_k _dafny.Seq = _dafny.EmptySeq
  var _ = xs_k
  { }
  var _4631_arr *_dafny.Array
  var _ = _4631_arr
  var _nw7 = _dafny.NewArrayWithValue(Type_X_.Default(), (xs).Cardinality())
  var _ = _nw7
  _4631_arr = _nw7
  var _4632_i _dafny.Int
  var _ = _4632_i
  _4632_i = _dafny.Zero
  var _4633_j _dafny.Int
  var _ = _4633_j
  _4633_j = _dafny.Zero
  for (_4632_i).Cmp((xs).Cardinality()) < 0 {
    { }
    { }
    if ((f)((xs).Index(_4632_i).(interface{}))) {
      { }
      *((_4631_arr).Index(_dafny.IntOfAny((_4633_j)))) = (xs).Index(_4632_i).(interface{})
      _4633_j = (_4633_j).Plus(_dafny.IntOfInt64(1))
      { }
    }
    _4632_i = (_4632_i).Plus(_dafny.IntOfInt64(1))
    { }
  }
  xs_k = (_4631_arr).RangeToSeq(_dafny.NilInt, _4633_j)
  return xs_k
}
func (_this *CompanionStruct_Default___) UniqueSubsequenceConstruct(Type_X_ _dafny.Type, xs _dafny.Seq, f func (interface{}) bool) _dafny.Seq {
  goto TAIL_CALL_START
TAIL_CALL_START:
var xs_k _dafny.Seq = _dafny.EmptySeq
  var _ = xs_k
  var _4634_s _dafny.Set
  var _ = _4634_s
  _4634_s = func () _dafny.Set {
    var _coll2 = _dafny.NewBuilder()
    var _ = _coll2
for _iter3 := _dafny.Iterate((xs).Elements());; {
      _val3, _ok3 := _iter3()
if !_ok3 { break }
_compr_1 := _val3.(interface{})
_4635_x := _compr_1
if (((xs).Contains(_4635_x)) && ((f)(_4635_x))) {
        _coll2.Add(_4635_x)
      }
    }
    return _coll2.ToSet()
  }()
  var _out59 _dafny.Seq
  var _ = _out59
_out59 = Companion_Default___.SetToUniqueSeqConstruct(Type_X_, _4634_s)
xs_k = _out59
  return xs_k
}
func (_this *CompanionStruct_Default___) AppendToUniqueSeq(xs _dafny.Seq, x interface{}) _dafny.Seq {
  var _4636_xs_k _dafny.Seq = (xs).Concat(_dafny.SeqOf(x))
  var _ = _4636_xs_k
return _4636_xs_k
}
func (_this *CompanionStruct_Default___) AppendToUniqueSeqMaybe(xs _dafny.Seq, x interface{}) _dafny.Seq {
  if ((xs).Contains(x)) {
    return xs
  } else  {
    var _4637_xs_k _dafny.Seq = (xs).Concat(_dafny.SeqOf(x))
    var _ = _4637_xs_k
return _4637_xs_k
  }
}
// End of class Default__
