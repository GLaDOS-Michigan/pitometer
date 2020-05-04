// Package _148_Common____SeqIsUnique__i_Compile
// Dafny module _148_Common____SeqIsUnique__i_Compile compiled into Go

package _148_Common____SeqIsUnique__i_Compile

import (
  _dafny "dafny"
_System "System_"
_0_Native____NativeTypes__s_Compile "0_Native____NativeTypes__s_Compile_"
_2_Collections____Maps2__s_Compile "2_Collections____Maps2__s_Compile_"
_5_Temporal____Temporal__s_Compile "5_Temporal____Temporal__s_Compile_"
_7_Environment__s_Compile "7_Environment__s_Compile_"
_9_Native____Io__s_Compile "9_Native____Io__s_Compile_"
_26_Collections____Seqs__s_Compile "26_Collections____Seqs__s_Compile_"
_29_Collections____Sets__i_Compile "29_Collections____Sets__i_Compile_"
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
_61_Math____power__s_Compile "61_Math____power__s_Compile_"
_64_Math____mul__nonlinear__i_Compile "64_Math____mul__nonlinear__i_Compile_"
_67_Math____mul__auto__proofs__i_Compile "67_Math____mul__auto__proofs__i_Compile_"
_69_Math____mul__auto__i_Compile "69_Math____mul__auto__i_Compile_"
_71_Math____mul__i_Compile "71_Math____mul__i_Compile_"
_73_Math____power__i_Compile "73_Math____power__i_Compile_"
_77_Math____div__def__i_Compile "77_Math____div__def__i_Compile_"
_81_Math____div__boogie__i_Compile "81_Math____div__boogie__i_Compile_"
_83_Math____div__nonlinear__i_Compile "83_Math____div__nonlinear__i_Compile_"
_88_Math____mod__auto__proofs__i_Compile "88_Math____mod__auto__proofs__i_Compile_"
_90_Math____mod__auto__i_Compile "90_Math____mod__auto__i_Compile_"
_93_Math____div__auto__proofs__i_Compile "93_Math____div__auto__proofs__i_Compile_"
_95_Math____div__auto__i_Compile "95_Math____div__auto__i_Compile_"
_97_Math____div__i_Compile "97_Math____div__i_Compile_"
_99_Math____power2__i_Compile "99_Math____power2__i_Compile_"
_101_Common____Util__i_Compile "101_Common____Util__i_Compile_"
_105_Common____MarshallInt__i_Compile "105_Common____MarshallInt__i_Compile_"
_107_Common____GenericMarshalling__i_Compile "107_Common____GenericMarshalling__i_Compile_"
_111_PacketParsing__i_Compile "111_PacketParsing__i_Compile_"
_113_Common____SeqIsUniqueDef__i_Compile "113_Common____SeqIsUniqueDef__i_Compile_"
_115_Impl__Node__i_Compile "115_Impl__Node__i_Compile_"
_119_UdpLock__i_Compile "119_UdpLock__i_Compile_"
_121_NodeImpl__i_Compile "121_NodeImpl__i_Compile_"
_127_CmdLineParser__i_Compile "127_CmdLineParser__i_Compile_"
_129_LockCmdLineParser__i_Compile "129_LockCmdLineParser__i_Compile_"
_131_Host__i_Compile "131_Host__i_Compile_"
_133_Lock__DistributedSystem__i_Compile "133_Lock__DistributedSystem__i_Compile_"
_138_Concrete__NodeIdentity__i_Compile "138_Concrete__NodeIdentity__i_Compile_"
_143_AbstractServiceLock__s_Compile "143_AbstractServiceLock__s_Compile_"
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
var _ _107_Common____GenericMarshalling__i_Compile.Dummy__
var _ _111_PacketParsing__i_Compile.Dummy__
var _ _113_Common____SeqIsUniqueDef__i_Compile.Dummy__
var _ _115_Impl__Node__i_Compile.Dummy__
var _ _119_UdpLock__i_Compile.Dummy__
var _ _121_NodeImpl__i_Compile.Dummy__
var _ _127_CmdLineParser__i_Compile.Dummy__
var _ _129_LockCmdLineParser__i_Compile.Dummy__
var _ _131_Host__i_Compile.Dummy__
var _ _133_Lock__DistributedSystem__i_Compile.Dummy__
var _ _138_Concrete__NodeIdentity__i_Compile.Dummy__
var _ _143_AbstractServiceLock__s_Compile.Dummy__

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
  return "_148_Common____SeqIsUnique__i_Compile.Default__"
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
  return "_148_Common____SeqIsUnique__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) SeqToSetConstruct(xs _dafny.Seq) _dafny.Set {
  goto TAIL_CALL_START
TAIL_CALL_START:
var s _dafny.Set = _dafny.EmptySet
  var _ = s
  { }
  s = _dafny.SetOf()
  var _1750_i _dafny.Int
  var _ = _1750_i
  _1750_i = _dafny.Zero
  for (_1750_i).Cmp((xs).Cardinality()) < 0 {
    s = (s).Union(_dafny.SetOf((xs).Index(_1750_i).(interface{})))
    _1750_i = (_1750_i).Plus(_dafny.IntOfInt64(1))
  }
  return s
}
func (_this *CompanionStruct_Default___) SetToUniqueSeqConstruct(Type_X_ _dafny.Type, s _dafny.Set) _dafny.Seq {
  var xs _dafny.Seq = _dafny.EmptySeq
  var _ = xs
  var _1751_arr *_dafny.Array
  var _ = _1751_arr
  var _nw8 = _dafny.NewArrayWithValue(Type_X_.Default(), (s).Cardinality())
  var _ = _nw8
  _1751_arr = _nw8
  var _1752_s1 _dafny.Set
  var _ = _1752_s1
  _1752_s1 = s
  { }
  { }
  { }
  for ((_1752_s1).Cardinality()).Cmp(_dafny.Zero) != 0 {
    { }
    { }
    var _1753_x interface{}
    var _ = _1753_x
    for _iter2 := _dafny.Iterate((_1752_s1).Elements());; {
      _val2, _ok2 := _iter2()
if !_ok2 { break }
_assign_such_that_0 := _val2.(interface{})
_1753_x = _assign_such_that_0
if ((_1752_s1).Contains(_1753_x)) {
        goto L_ASSIGN_SUCH_THAT_0
      }
    }
    panic("assign-such-that search produced no value (line 85)")
  L_ASSIGN_SUCH_THAT_0:
    { }
    { }
    var _index7 = ((s).Cardinality()).Minus((_1752_s1).Cardinality())
    var _ = _index7
    *((_1751_arr).Index(_dafny.IntOfAny(_index7))) = _1753_x
    _1752_s1 = (_1752_s1).Difference(_dafny.SetOf(_1753_x))
    { }
    { }
    { }
  }
  xs = (_1751_arr).RangeToSeq(_dafny.NilInt, _dafny.NilInt)
  { }
  return xs
}
func (_this *CompanionStruct_Default___) SubsequenceConstruct(Type_X_ _dafny.Type, xs _dafny.Seq, f func (interface{}) bool) _dafny.Seq {
  var xs_k _dafny.Seq = _dafny.EmptySeq
  var _ = xs_k
  { }
  var _1754_arr *_dafny.Array
  var _ = _1754_arr
  var _nw9 = _dafny.NewArrayWithValue(Type_X_.Default(), (xs).Cardinality())
  var _ = _nw9
  _1754_arr = _nw9
  var _1755_i _dafny.Int
  var _ = _1755_i
  _1755_i = _dafny.Zero
  var _1756_j _dafny.Int
  var _ = _1756_j
  _1756_j = _dafny.Zero
  for (_1755_i).Cmp((xs).Cardinality()) < 0 {
    { }
    { }
    if ((f)((xs).Index(_1755_i).(interface{}))) {
      { }
      *((_1754_arr).Index(_dafny.IntOfAny((_1756_j)))) = (xs).Index(_1755_i).(interface{})
      _1756_j = (_1756_j).Plus(_dafny.IntOfInt64(1))
      { }
    }
    _1755_i = (_1755_i).Plus(_dafny.IntOfInt64(1))
    { }
  }
  xs_k = (_1754_arr).RangeToSeq(_dafny.NilInt, _1756_j)
  return xs_k
}
func (_this *CompanionStruct_Default___) UniqueSubsequenceConstruct(Type_X_ _dafny.Type, xs _dafny.Seq, f func (interface{}) bool) _dafny.Seq {
  goto TAIL_CALL_START
TAIL_CALL_START:
var xs_k _dafny.Seq = _dafny.EmptySeq
  var _ = xs_k
  var _1757_s _dafny.Set
  var _ = _1757_s
  _1757_s = func () _dafny.Set {
    var _coll2 = _dafny.NewBuilder()
    var _ = _coll2
for _iter3 := _dafny.Iterate((xs).Elements());; {
      _val3, _ok3 := _iter3()
if !_ok3 { break }
_compr_1 := _val3.(interface{})
_1758_x := _compr_1
if (((xs).Contains(_1758_x)) && ((f)(_1758_x))) {
        _coll2.Add(_1758_x)
      }
    }
    return _coll2.ToSet()
  }()
  var _out106 _dafny.Seq
  var _ = _out106
_out106 = Companion_Default___.SetToUniqueSeqConstruct(Type_X_, _1757_s)
xs_k = _out106
  return xs_k
}
func (_this *CompanionStruct_Default___) AppendToUniqueSeq(xs _dafny.Seq, x interface{}) _dafny.Seq {
  var _1759_xs_k _dafny.Seq = (xs).Concat(_dafny.SeqOf(x))
  var _ = _1759_xs_k
return _1759_xs_k
}
func (_this *CompanionStruct_Default___) AppendToUniqueSeqMaybe(xs _dafny.Seq, x interface{}) _dafny.Seq {
  if ((xs).Contains(x)) {
    return xs
  } else  {
    var _1760_xs_k _dafny.Seq = (xs).Concat(_dafny.SeqOf(x))
    var _ = _1760_xs_k
return _1760_xs_k
  }
}
// End of class Default__
