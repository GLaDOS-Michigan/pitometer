// Package _127_CmdLineParser__i_Compile
// Dafny module _127_CmdLineParser__i_Compile compiled into Go

package _127_CmdLineParser__i_Compile

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
  return "_127_CmdLineParser__i_Compile.Default__"
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
  return "_127_CmdLineParser__i_Compile.Default__"
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
    var _1693_tmp _dafny.Int
    var _ = _1693_tmp
var _out93 _dafny.Int
    var _ = _out93
_out93 = Companion_Default___.Power10((e).Minus(_dafny.IntOfInt64(1)))
_1693_tmp = _out93
    val = (_dafny.IntOfInt64(10)).Times(_1693_tmp)
return val
  }
  return val
}
func (_this *CompanionStruct_Default___) Shorts__to__bytes(shorts _dafny.Seq) _System.Tuple2 {
  if (((shorts).Cardinality()).Cmp(_dafny.Zero) == 0) {
    return _dafny.TupleOf(true, _dafny.SeqOf())
  } else  {
    var _1694_tuple _System.Tuple2 = Companion_Default___.Shorts__to__bytes((shorts).Subseq(_dafny.IntOfInt64(1), _dafny.NilInt))
    var _ = _1694_tuple
var _1695_ok bool = (*((_1694_tuple)).IndexInt(0)).(bool)
    var _ = _1695_ok
var _1696_rest _dafny.Seq = (*((_1694_tuple)).IndexInt(1)).(_dafny.Seq)
    var _ = _1696_rest
var _1697_tuple_k _System.Tuple2 = Companion_Default___.Ascii__to__int((shorts).Index(_dafny.Zero).(uint16))
    var _ = _1697_tuple_k
var _1698_ok_k bool = (*((_1697_tuple_k)).IndexInt(0)).(bool)
    var _ = _1698_ok_k
var _1699_a__byte uint8 = (*((_1697_tuple_k)).IndexInt(1)).(uint8)
    var _ = _1699_a__byte
if ((_1695_ok) && (_1698_ok_k)) {
      return _dafny.TupleOf(true, (_dafny.SeqOf(_1699_a__byte)).Concat(_1696_rest))
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
    var _1700_tuple _System.Tuple2 = Companion_Default___.Shorts__to__bytes(shorts)
    var _ = _1700_tuple
var _1701_ok bool = (*((_1700_tuple)).IndexInt(0)).(bool)
    var _ = _1701_ok
var _1702_bytes _dafny.Seq = (*((_1700_tuple)).IndexInt(1)).(_dafny.Seq)
    var _ = _1702_bytes
if (!(_1701_ok)) {
      return _dafny.TupleOf(false, _dafny.Zero)
    } else  {
      return _dafny.TupleOf(true, Companion_Default___.Bytes__to__decimal(_1702_bytes))
    }
  }
}
func (_this *CompanionStruct_Default___) Shorts__to__byte(shorts _dafny.Seq) _System.Tuple2 {
  var _1703_tuple _System.Tuple2 = Companion_Default___.Shorts__to__nat(shorts)
  var _ = _1703_tuple
var _1704_ok bool = (*((_1703_tuple)).IndexInt(0)).(bool)
  var _ = _1704_ok
var _1705_val _dafny.Int = (*((_1703_tuple)).IndexInt(1)).(_dafny.Int)
  var _ = _1705_val
if (((_dafny.Zero).Cmp(_1705_val) <= 0) && ((_1705_val).Cmp(_dafny.IntOfInt64(256)) < 0)) {
    return _dafny.TupleOf(true, (_1705_val).Uint8())
  } else  {
    return _dafny.TupleOf(false, uint8(0))
  }
}
func (_this *CompanionStruct_Default___) Shorts__to__uint16(shorts _dafny.Seq) _System.Tuple2 {
  var _1706_tuple _System.Tuple2 = Companion_Default___.Shorts__to__nat(shorts)
  var _ = _1706_tuple
var _1707_ok bool = (*((_1706_tuple)).IndexInt(0)).(bool)
  var _ = _1707_ok
var _1708_val _dafny.Int = (*((_1706_tuple)).IndexInt(1)).(_dafny.Int)
  var _ = _1708_val
if (((_dafny.Zero).Cmp(_1708_val) <= 0) && ((_1708_val).Cmp(_dafny.IntOfInt64(65536)) < 0)) {
    return _dafny.TupleOf(true, (_1708_val).Uint16())
  } else  {
    return _dafny.TupleOf(false, uint16(0))
  }
}
func (_this *CompanionStruct_Default___) Shorts__to__uint32(shorts _dafny.Seq) _System.Tuple2 {
  var _1709_tuple _System.Tuple2 = Companion_Default___.Shorts__to__nat(shorts)
  var _ = _1709_tuple
var _1710_ok bool = (*((_1709_tuple)).IndexInt(0)).(bool)
  var _ = _1710_ok
var _1711_val _dafny.Int = (*((_1709_tuple)).IndexInt(1)).(_dafny.Int)
  var _ = _1711_val
if (((_dafny.Zero).Cmp(_1711_val) <= 0) && ((_1711_val).Cmp(_dafny.IntOfInt64(4294967296)) < 0)) {
    return _dafny.TupleOf(true, (_1711_val).Uint32())
  } else  {
    return _dafny.TupleOf(false, uint32(0))
  }
}
func (_this *CompanionStruct_Default___) Is__ascii__period(short uint16) bool {
  return (short) == (uint16(46))
}
func (_this *CompanionStruct_Default___) Parse__ip__addr__helper(ip__shorts _dafny.Seq, current__octet__shorts _dafny.Seq) _System.Tuple2 {
  if (((ip__shorts).Cardinality()).Cmp(_dafny.Zero) == 0) {
    var _1712_tuple _System.Tuple2 = Companion_Default___.Shorts__to__byte(current__octet__shorts)
    var _ = _1712_tuple
var _1713_okay bool = (*((_1712_tuple)).IndexInt(0)).(bool)
    var _ = _1713_okay
var _1714_b uint8 = (*((_1712_tuple)).IndexInt(1)).(uint8)
    var _ = _1714_b
if (!(_1713_okay)) {
      return _dafny.TupleOf(false, _dafny.SeqOf())
    } else  {
      return _dafny.TupleOf(true, _dafny.SeqOf(_1714_b))
    }
  } else  {
    if (Companion_Default___.Is__ascii__period((ip__shorts).Index(_dafny.Zero).(uint16))) {
      var _1715_tuple _System.Tuple2 = Companion_Default___.Shorts__to__byte(current__octet__shorts)
      var _ = _1715_tuple
var _1716_okay bool = (*((_1715_tuple)).IndexInt(0)).(bool)
      var _ = _1716_okay
var _1717_b uint8 = (*((_1715_tuple)).IndexInt(1)).(uint8)
      var _ = _1717_b
if (!(_1716_okay)) {
        return _dafny.TupleOf(false, _dafny.SeqOf())
      } else  {
        var _1718_tuple_k _System.Tuple2 = Companion_Default___.Parse__ip__addr__helper((ip__shorts).Subseq(_dafny.IntOfInt64(1), _dafny.NilInt), _dafny.SeqOf())
        var _ = _1718_tuple_k
var _1719_ok bool = (*((_1718_tuple_k)).IndexInt(0)).(bool)
        var _ = _1719_ok
var _1720_ip__bytes _dafny.Seq = (*((_1718_tuple_k)).IndexInt(1)).(_dafny.Seq)
        var _ = _1720_ip__bytes
if (!(_1719_ok)) {
          return _dafny.TupleOf(false, _dafny.SeqOf())
        } else  {
          return _dafny.TupleOf(true, (_dafny.SeqOf(_1717_b)).Concat(_1720_ip__bytes))
        }
      }
    } else  {
      return Companion_Default___.Parse__ip__addr__helper((ip__shorts).Subseq(_dafny.IntOfInt64(1), _dafny.NilInt), (current__octet__shorts).Concat(_dafny.SeqOf((ip__shorts).Index(_dafny.Zero).(uint16))))
    }
  }
}
func (_this *CompanionStruct_Default___) Parse__ip__addr(ip__shorts _dafny.Seq) _System.Tuple2 {
  var _1721_tuple _System.Tuple2 = Companion_Default___.Parse__ip__addr__helper(ip__shorts, _dafny.SeqOf())
  var _ = _1721_tuple
var _1722_ok bool = (*((_1721_tuple)).IndexInt(0)).(bool)
  var _ = _1722_ok
var _1723_ip__bytes _dafny.Seq = (*((_1721_tuple)).IndexInt(1)).(_dafny.Seq)
  var _ = _1723_ip__bytes
if ((_1722_ok) && (((_1723_ip__bytes).Cardinality()).Cmp(_dafny.IntOfInt64(4)) == 0)) {
    return _dafny.TupleOf(true, _1723_ip__bytes)
  } else  {
    return _dafny.TupleOf(false, _dafny.SeqOf())
  }
}
func (_this *CompanionStruct_Default___) Parse__end__point(ip__shorts _dafny.Seq, port__shorts _dafny.Seq) _System.Tuple2 {
  var _1724_tuple _System.Tuple2 = Companion_Default___.Parse__ip__addr(ip__shorts)
  var _ = _1724_tuple
var _1725_okay bool = (*((_1724_tuple)).IndexInt(0)).(bool)
  var _ = _1725_okay
var _1726_ip__bytes _dafny.Seq = (*((_1724_tuple)).IndexInt(1)).(_dafny.Seq)
  var _ = _1726_ip__bytes
if (!(_1725_okay)) {
    return _dafny.TupleOf(false, _9_Native____Io__s_Compile.EndPoint{_9_Native____Io__s_Compile.EndPoint_EndPoint{_dafny.SeqOf(uint8(0), uint8(0), uint8(0), uint8(0)), uint16(0)}})
  } else  {
    var _1727_tuple_k _System.Tuple2 = Companion_Default___.Shorts__to__uint16(port__shorts)
    var _ = _1727_tuple_k
var _1728_okay_k bool = (*((_1727_tuple_k)).IndexInt(0)).(bool)
    var _ = _1728_okay_k
var _1729_port uint16 = (*((_1727_tuple_k)).IndexInt(1)).(uint16)
    var _ = _1729_port
if (!(_1728_okay_k)) {
      return _dafny.TupleOf(false, _9_Native____Io__s_Compile.EndPoint{_9_Native____Io__s_Compile.EndPoint_EndPoint{_dafny.SeqOf(uint8(0), uint8(0), uint8(0), uint8(0)), uint16(0)}})
    } else  {
      return _dafny.TupleOf(true, _9_Native____Io__s_Compile.EndPoint{_9_Native____Io__s_Compile.EndPoint_EndPoint{_1726_ip__bytes, _1729_port}})
    }
  }
}
func (_this *CompanionStruct_Default___) Test__unique_k(endpoints _dafny.Seq) bool {
  goto TAIL_CALL_START
TAIL_CALL_START:
var unique bool = false
  var _ = unique
  unique = true
  var _1730_i _dafny.Int
  var _ = _1730_i
  _1730_i = _dafny.Zero
  for (_1730_i).Cmp((endpoints).Cardinality()) < 0 {
    var _1731_j _dafny.Int
    var _ = _1731_j
    _1731_j = _dafny.Zero
    for (_1731_j).Cmp((endpoints).Cardinality()) < 0 {
      if (((_1730_i).Cmp(_1731_j) != 0) && (((endpoints).Index(_1730_i).(_9_Native____Io__s_Compile.EndPoint)).Equals((endpoints).Index(_1731_j).(_9_Native____Io__s_Compile.EndPoint)))) {
        unique = false
        { }
        return unique
      }
      _1731_j = (_1731_j).Plus(_dafny.IntOfInt64(1))
    }
    _1730_i = (_1730_i).Plus(_dafny.IntOfInt64(1))
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
var _1732_ok1 bool = (*(_let_tmp_rhs10).IndexInt(0)).(bool)
    var _ = _1732_ok1
var _1733_ep _9_Native____Io__s_Compile.EndPoint = (*(_let_tmp_rhs10).IndexInt(1)).(_9_Native____Io__s_Compile.EndPoint)
    var _ = _1733_ep
var _let_tmp_rhs11 _System.Tuple2 = Companion_Default___.Parse__end__points((args).Subseq(_dafny.IntOfInt64(2), _dafny.NilInt))
    var _ = _let_tmp_rhs11
var _1734_ok2 bool = (*(_let_tmp_rhs11).IndexInt(0)).(bool)
    var _ = _1734_ok2
var _1735_rest _dafny.Seq = (*(_let_tmp_rhs11).IndexInt(1)).(_dafny.Seq)
    var _ = _1735_rest
if (!((_1732_ok1) && (_1734_ok2))) {
      return _dafny.TupleOf(false, _dafny.SeqOf())
    } else  {
      return _dafny.TupleOf(true, (_dafny.SeqOf(_1733_ep)).Concat(_1735_rest))
    }
  }
}
func (_this *CompanionStruct_Default___) Collect__cmd__line__args() _dafny.Seq {
  goto TAIL_CALL_START
TAIL_CALL_START:
var args _dafny.Seq = _dafny.EmptySeq
  var _ = args
  var _1736_num__args uint32
  var _ = _1736_num__args
var _out94 uint32
  var _ = _out94
_out94 = _9_Native____Io__s_Compile.Companion_HostConstants_.NumCommandLineArgs()
_1736_num__args = _out94
  var _1737_i uint32
  var _ = _1737_i
  _1737_i = uint32(0)
  args = _dafny.SeqOf()
  for (_1737_i) < (_1736_num__args) {
    var _1738_arg *_dafny.Array
    var _ = _1738_arg
var _out95 *_dafny.Array
    var _ = _out95
_out95 = _9_Native____Io__s_Compile.Companion_HostConstants_.GetCommandLineArg(uint64(_1737_i))
_1738_arg = _out95
    args = (args).Concat(_dafny.SeqOf((_1738_arg).RangeToSeq(_dafny.NilInt, _dafny.NilInt)))
    _1737_i = (_1737_i) + (uint32(1))
  }
  return args
}
// End of class Default__
