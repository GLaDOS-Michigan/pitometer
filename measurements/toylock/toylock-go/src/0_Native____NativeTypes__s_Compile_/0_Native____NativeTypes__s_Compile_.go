// Package _0_Native____NativeTypes__s_Compile
// Dafny module _0_Native____NativeTypes__s_Compile compiled into Go

package _0_Native____NativeTypes__s_Compile

import (
  _dafny "dafny"
_System "System_"
)
var _ _dafny.Dummy__
var _ _System.Dummy__

type Dummy__ struct{}


// Definition of class Sbyte
type Sbyte struct {
}

func New_Sbyte_() *Sbyte {
  _this := Sbyte{}

  return &_this
}

type CompanionStruct_Sbyte_ struct {
}
var Companion_Sbyte_ = CompanionStruct_Sbyte_ {
}

func (*Sbyte) String() string {
  return "_0_Native____NativeTypes__s_Compile.Sbyte"
}
func (_this *CompanionStruct_Sbyte_) IntegerRange(lo _dafny.Int, hi _dafny.Int) _dafny.Iterator {
  iter := _dafny.IntegerRange(lo, hi)
return func() (interface{}, bool) {
    next, ok := iter()
if !ok { return int8(0), false }
return next.(_dafny.Int).Int8(), true
  }
}
// End of class Sbyte
func Type_Sbyte_() _dafny.Type {
  return type_Sbyte_{}
}

type type_Sbyte_ struct {
}

func (_this type_Sbyte_) Default() interface{} {
  return 0
}

func (_this type_Sbyte_) String() string {
  return "_0_Native____NativeTypes__s_Compile.Sbyte"
}

// Definition of class Byte
type Byte struct {
}

func New_Byte_() *Byte {
  _this := Byte{}

  return &_this
}

type CompanionStruct_Byte_ struct {
}
var Companion_Byte_ = CompanionStruct_Byte_ {
}

func (*Byte) String() string {
  return "_0_Native____NativeTypes__s_Compile.Byte"
}
func (_this *CompanionStruct_Byte_) IntegerRange(lo _dafny.Int, hi _dafny.Int) _dafny.Iterator {
  iter := _dafny.IntegerRange(lo, hi)
return func() (interface{}, bool) {
    next, ok := iter()
if !ok { return uint8(0), false }
return next.(_dafny.Int).Uint8(), true
  }
}
// End of class Byte
func Type_Byte_() _dafny.Type {
  return type_Byte_{}
}

type type_Byte_ struct {
}

func (_this type_Byte_) Default() interface{} {
  return 0
}

func (_this type_Byte_) String() string {
  return "_0_Native____NativeTypes__s_Compile.Byte"
}

// Definition of class Int16
type Int16 struct {
}

func New_Int16_() *Int16 {
  _this := Int16{}

  return &_this
}

type CompanionStruct_Int16_ struct {
}
var Companion_Int16_ = CompanionStruct_Int16_ {
}

func (*Int16) String() string {
  return "_0_Native____NativeTypes__s_Compile.Int16"
}
func (_this *CompanionStruct_Int16_) IntegerRange(lo _dafny.Int, hi _dafny.Int) _dafny.Iterator {
  iter := _dafny.IntegerRange(lo, hi)
return func() (interface{}, bool) {
    next, ok := iter()
if !ok { return int16(0), false }
return next.(_dafny.Int).Int16(), true
  }
}
// End of class Int16
func Type_Int16_() _dafny.Type {
  return type_Int16_{}
}

type type_Int16_ struct {
}

func (_this type_Int16_) Default() interface{} {
  return 0
}

func (_this type_Int16_) String() string {
  return "_0_Native____NativeTypes__s_Compile.Int16"
}

// Definition of class Uint16
type Uint16 struct {
}

func New_Uint16_() *Uint16 {
  _this := Uint16{}

  return &_this
}

type CompanionStruct_Uint16_ struct {
}
var Companion_Uint16_ = CompanionStruct_Uint16_ {
}

func (*Uint16) String() string {
  return "_0_Native____NativeTypes__s_Compile.Uint16"
}
func (_this *CompanionStruct_Uint16_) IntegerRange(lo _dafny.Int, hi _dafny.Int) _dafny.Iterator {
  iter := _dafny.IntegerRange(lo, hi)
return func() (interface{}, bool) {
    next, ok := iter()
if !ok { return uint16(0), false }
return next.(_dafny.Int).Uint16(), true
  }
}
// End of class Uint16
func Type_Uint16_() _dafny.Type {
  return type_Uint16_{}
}

type type_Uint16_ struct {
}

func (_this type_Uint16_) Default() interface{} {
  return 0
}

func (_this type_Uint16_) String() string {
  return "_0_Native____NativeTypes__s_Compile.Uint16"
}

// Definition of class Int32
type Int32 struct {
}

func New_Int32_() *Int32 {
  _this := Int32{}

  return &_this
}

type CompanionStruct_Int32_ struct {
}
var Companion_Int32_ = CompanionStruct_Int32_ {
}

func (*Int32) String() string {
  return "_0_Native____NativeTypes__s_Compile.Int32"
}
func (_this *CompanionStruct_Int32_) IntegerRange(lo _dafny.Int, hi _dafny.Int) _dafny.Iterator {
  iter := _dafny.IntegerRange(lo, hi)
return func() (interface{}, bool) {
    next, ok := iter()
if !ok { return int32(0), false }
return next.(_dafny.Int).Int32(), true
  }
}
// End of class Int32
func Type_Int32_() _dafny.Type {
  return type_Int32_{}
}

type type_Int32_ struct {
}

func (_this type_Int32_) Default() interface{} {
  return 0
}

func (_this type_Int32_) String() string {
  return "_0_Native____NativeTypes__s_Compile.Int32"
}

// Definition of class Uint32
type Uint32 struct {
}

func New_Uint32_() *Uint32 {
  _this := Uint32{}

  return &_this
}

type CompanionStruct_Uint32_ struct {
}
var Companion_Uint32_ = CompanionStruct_Uint32_ {
}

func (*Uint32) String() string {
  return "_0_Native____NativeTypes__s_Compile.Uint32"
}
func (_this *CompanionStruct_Uint32_) IntegerRange(lo _dafny.Int, hi _dafny.Int) _dafny.Iterator {
  iter := _dafny.IntegerRange(lo, hi)
return func() (interface{}, bool) {
    next, ok := iter()
if !ok { return uint32(0), false }
return next.(_dafny.Int).Uint32(), true
  }
}
// End of class Uint32
func Type_Uint32_() _dafny.Type {
  return type_Uint32_{}
}

type type_Uint32_ struct {
}

func (_this type_Uint32_) Default() interface{} {
  return 0
}

func (_this type_Uint32_) String() string {
  return "_0_Native____NativeTypes__s_Compile.Uint32"
}

// Definition of class Int64
type Int64 struct {
}

func New_Int64_() *Int64 {
  _this := Int64{}

  return &_this
}

type CompanionStruct_Int64_ struct {
}
var Companion_Int64_ = CompanionStruct_Int64_ {
}

func (*Int64) String() string {
  return "_0_Native____NativeTypes__s_Compile.Int64"
}
func (_this *CompanionStruct_Int64_) IntegerRange(lo _dafny.Int, hi _dafny.Int) _dafny.Iterator {
  iter := _dafny.IntegerRange(lo, hi)
return func() (interface{}, bool) {
    next, ok := iter()
if !ok { return int64(0), false }
return next.(_dafny.Int).Int64(), true
  }
}
// End of class Int64
func Type_Int64_() _dafny.Type {
  return type_Int64_{}
}

type type_Int64_ struct {
}

func (_this type_Int64_) Default() interface{} {
  return 0
}

func (_this type_Int64_) String() string {
  return "_0_Native____NativeTypes__s_Compile.Int64"
}

// Definition of class Uint64
type Uint64 struct {
}

func New_Uint64_() *Uint64 {
  _this := Uint64{}

  return &_this
}

type CompanionStruct_Uint64_ struct {
}
var Companion_Uint64_ = CompanionStruct_Uint64_ {
}

func (*Uint64) String() string {
  return "_0_Native____NativeTypes__s_Compile.Uint64"
}
func (_this *CompanionStruct_Uint64_) IntegerRange(lo _dafny.Int, hi _dafny.Int) _dafny.Iterator {
  iter := _dafny.IntegerRange(lo, hi)
return func() (interface{}, bool) {
    next, ok := iter()
if !ok { return uint64(0), false }
return next.(_dafny.Int).Uint64(), true
  }
}
// End of class Uint64
func Type_Uint64_() _dafny.Type {
  return type_Uint64_{}
}

type type_Uint64_ struct {
}

func (_this type_Uint64_) Default() interface{} {
  return 0
}

func (_this type_Uint64_) String() string {
  return "_0_Native____NativeTypes__s_Compile.Uint64"
}

// Definition of class Nat8
type Nat8 struct {
}

func New_Nat8_() *Nat8 {
  _this := Nat8{}

  return &_this
}

type CompanionStruct_Nat8_ struct {
}
var Companion_Nat8_ = CompanionStruct_Nat8_ {
}

func (*Nat8) String() string {
  return "_0_Native____NativeTypes__s_Compile.Nat8"
}
func (_this *CompanionStruct_Nat8_) IntegerRange(lo _dafny.Int, hi _dafny.Int) _dafny.Iterator {
  iter := _dafny.IntegerRange(lo, hi)
return func() (interface{}, bool) {
    next, ok := iter()
if !ok { return int8(0), false }
return next.(_dafny.Int).Int8(), true
  }
}
// End of class Nat8
func Type_Nat8_() _dafny.Type {
  return type_Nat8_{}
}

type type_Nat8_ struct {
}

func (_this type_Nat8_) Default() interface{} {
  return 0
}

func (_this type_Nat8_) String() string {
  return "_0_Native____NativeTypes__s_Compile.Nat8"
}

// Definition of class Nat16
type Nat16 struct {
}

func New_Nat16_() *Nat16 {
  _this := Nat16{}

  return &_this
}

type CompanionStruct_Nat16_ struct {
}
var Companion_Nat16_ = CompanionStruct_Nat16_ {
}

func (*Nat16) String() string {
  return "_0_Native____NativeTypes__s_Compile.Nat16"
}
func (_this *CompanionStruct_Nat16_) IntegerRange(lo _dafny.Int, hi _dafny.Int) _dafny.Iterator {
  iter := _dafny.IntegerRange(lo, hi)
return func() (interface{}, bool) {
    next, ok := iter()
if !ok { return int16(0), false }
return next.(_dafny.Int).Int16(), true
  }
}
// End of class Nat16
func Type_Nat16_() _dafny.Type {
  return type_Nat16_{}
}

type type_Nat16_ struct {
}

func (_this type_Nat16_) Default() interface{} {
  return 0
}

func (_this type_Nat16_) String() string {
  return "_0_Native____NativeTypes__s_Compile.Nat16"
}

// Definition of class Nat32
type Nat32 struct {
}

func New_Nat32_() *Nat32 {
  _this := Nat32{}

  return &_this
}

type CompanionStruct_Nat32_ struct {
}
var Companion_Nat32_ = CompanionStruct_Nat32_ {
}

func (*Nat32) String() string {
  return "_0_Native____NativeTypes__s_Compile.Nat32"
}
func (_this *CompanionStruct_Nat32_) IntegerRange(lo _dafny.Int, hi _dafny.Int) _dafny.Iterator {
  iter := _dafny.IntegerRange(lo, hi)
return func() (interface{}, bool) {
    next, ok := iter()
if !ok { return int32(0), false }
return next.(_dafny.Int).Int32(), true
  }
}
// End of class Nat32
func Type_Nat32_() _dafny.Type {
  return type_Nat32_{}
}

type type_Nat32_ struct {
}

func (_this type_Nat32_) Default() interface{} {
  return 0
}

func (_this type_Nat32_) String() string {
  return "_0_Native____NativeTypes__s_Compile.Nat32"
}

// Definition of class Nat64
type Nat64 struct {
}

func New_Nat64_() *Nat64 {
  _this := Nat64{}

  return &_this
}

type CompanionStruct_Nat64_ struct {
}
var Companion_Nat64_ = CompanionStruct_Nat64_ {
}

func (*Nat64) String() string {
  return "_0_Native____NativeTypes__s_Compile.Nat64"
}
func (_this *CompanionStruct_Nat64_) IntegerRange(lo _dafny.Int, hi _dafny.Int) _dafny.Iterator {
  iter := _dafny.IntegerRange(lo, hi)
return func() (interface{}, bool) {
    next, ok := iter()
if !ok { return int64(0), false }
return next.(_dafny.Int).Int64(), true
  }
}
// End of class Nat64
func Type_Nat64_() _dafny.Type {
  return type_Nat64_{}
}

type type_Nat64_ struct {
}

func (_this type_Nat64_) Default() interface{} {
  return 0
}

func (_this type_Nat64_) String() string {
  return "_0_Native____NativeTypes__s_Compile.Nat64"
}

