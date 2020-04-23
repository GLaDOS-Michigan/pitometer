// Package _47_Collections____Maps__i_Compile
// Dafny module _47_Collections____Maps__i_Compile compiled into Go

package _47_Collections____Maps__i_Compile

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
  return "_47_Collections____Maps__i_Compile.Default__"
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
  return "_47_Collections____Maps__i_Compile.Default__"
}
func (_this *CompanionStruct_Default___) Domain(m _dafny.Map) _dafny.Set {
  return func () _dafny.Set {
    var _coll0 = _dafny.NewBuilder()
    var _ = _coll0
for _iter0 := _dafny.Iterate((m).Keys().Elements());; {
      _val0, _ok0 := _iter0()
if !_ok0 { break }
_compr_0 := _val0.(interface{})
_1565_s := _compr_0
if ((m).Contains(_1565_s)) {
        _coll0.Add(_1565_s)
      }
    }
    return _coll0.ToSet()
  }()
}
func (_this *CompanionStruct_Default___) RemoveElt(m _dafny.Map, elt interface{}) _dafny.Map {
  var _1566_m_k _dafny.Map = func () _dafny.Map {
    var _coll1 = _dafny.NewMapBuilder()
    var _ = _coll1
for _iter1 := _dafny.Iterate((m).Keys().Elements());; {
      _val1, _ok1 := _iter1()
if !_ok1 { break }
_1567_elt_k := _val1.(interface{})
if (((m).Contains(_1567_elt_k)) && (!_dafny.AreEqual(_1567_elt_k, elt))) {
        _coll1.Add(_1567_elt_k,(m).Get(_1567_elt_k).(interface{}))
      }
    }
    return _coll1.ToMap()
  }()
  var _ = _1566_m_k
return _1566_m_k
}
// End of class Default__
