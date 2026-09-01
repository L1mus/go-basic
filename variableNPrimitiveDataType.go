package gobasic

import (
	"fmt"
)

// declaration variable
// 1. Explicit
// var
var item string = "Box"

// const
const phi = 3.14

// 2. Inference, compiler mendeteksi otomatis type berdasarkan value
var isAvailable = true

// short declaration, harus di deklarasikan di dalama function jika tidak akan expected declaration
func yourName() {
	firstName := "Jenna Ortega"
	fmt.Println(firstName)
}

// 3. ZeroValue , primitive number 0/0.0, string "", bool false. non-primitive nil/null
var qty int32

// 4. Multi variable declaration
var (
	word  = "hello"
	name  = "limus"
	heigh = 26
)

var x, y, z int8 = 1, 2, 3

// IOTA. untuk kebutuhan enumerisasi multiple variable declaration
type ByteSize float64

const (
	_           = iota
	kb ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// PRIMITIVE
// NUMBER
// INTEGER
var tipeInteger8 int8
var tipeInteger16 int16
var tipeInteger32 int32
var tipeInteger64 int64

// UINTEGER
var tipeuInteger8 uint8
var tipeuInteger16 uint16
var tipeuInteger32 uint32
var tipeuInteger64 uint64

// FLOAT
var tipeFloat32 float32
var tipeFloat64 float64

// COMPLEX

var tipeComplex64 complex64
var tipeComplex128 complex128

// BUILDIN COMPLEX FN imag(),real(),abs()

// ALIAS
var tipeAliasUint8 byte
var tipeAliasInt32 rune

// STRING
var str string
var oneWord = "string"

// BOOLEAN
var boolean bool
var isValid = true
