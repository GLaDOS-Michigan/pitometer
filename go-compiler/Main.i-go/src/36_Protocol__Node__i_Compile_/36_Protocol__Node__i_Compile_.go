// Package _36_Protocol__Node__i_Compile
// Dafny module _36_Protocol__Node__i_Compile compiled into Go

package _36_Protocol__Node__i_Compile

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

type Dummy__ struct{}





// Definition of data type Node
type Node struct {
  Data_Node_
}

func (_this Node) Get() Data_Node_ {
  return _this.Data_Node_
}

type Data_Node_ interface {
  isNode()
}

type CompanionStruct_Node_ struct {}
var Companion_Node_ = CompanionStruct_Node_{}

type Node_Node struct {
  Held bool
Epoch _dafny.Int
My__index _dafny.Int
Config _dafny.Seq
}

func (Node_Node) isNode() {}

func (CompanionStruct_Node_) Create_Node_(Held bool, Epoch _dafny.Int, My__index _dafny.Int, Config _dafny.Seq) Node {
  return Node{Node_Node{Held,Epoch,My__index,Config}}
}

func (_this Node) Is_Node() bool {
  _, ok := _this.Get().(Node_Node)
return ok
}

func (_this Node) Dtor_held() bool {
  return _this.Get().(Node_Node).Held
}

func (_this Node) Dtor_epoch() _dafny.Int {
  return _this.Get().(Node_Node).Epoch
}

func (_this Node) Dtor_my__index() _dafny.Int {
  return _this.Get().(Node_Node).My__index
}

func (_this Node) Dtor_config() _dafny.Seq {
  return _this.Get().(Node_Node).Config
}

func (_this Node) String() string {
  switch data := _this.Get().(type) {
    case nil: return "null"
case Node_Node: {
      return "_36_Protocol__Node__i_Compile.Node.Node" + "(" + _dafny.String(data.Held) + ", " + _dafny.String(data.Epoch) + ", " + _dafny.String(data.My__index) + ", " + _dafny.String(data.Config) + ")"
    }
    default: {
      return "<unexpected>"
    }
  }
}

func (_this Node) Equals(other Node) bool {
  switch data1 := _this.Get().(type) {
    case Node_Node: {
      data2, ok := other.Get().(Node_Node)
return ok && data1.Held == data2.Held && data1.Epoch.Cmp(data2.Epoch) == 0 && data1.My__index.Cmp(data2.My__index) == 0 && data1.Config.Equals(data2.Config)
    }
    default: {
      return false; // unexpected
    }
  }
}

func (_this Node) EqualsGeneric(other interface{}) bool {
  typed, ok := other.(Node)
return ok && _this.Equals(typed)
}
func Type_Node_() _dafny.Type {
  return type_Node_{}
}

type type_Node_ struct {
}

func (_this type_Node_) Default() interface{} {
  return Node{Node_Node{false, _dafny.Zero, _dafny.Zero, _dafny.EmptySeq}}
}

func (_this type_Node_) String() string {
  return "_36_Protocol__Node__i_Compile.Node"
}
// End of data type Node

