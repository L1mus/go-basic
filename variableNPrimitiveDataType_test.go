package gobasic

import (
	"fmt"
	"testing"
)

func TestDeclarationvariable(t *testing.T) {

	// declaration variable
	// 1. Explicit
	// var
	var item string = "Box"

	// const
	const phi = 3.14

	// 2. Inference, compiler mendeteksi otomatis type berdasarkan value
	var isAvailable = true

	// short declaration, harus di deklarasikan di dalama function jika tidak akan expected declaration
	firstName := "Jenna Ortega"

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
		KB ByteSize = 1 << (10 * iota)
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

	fmt.Println("declaration variable")
	fmt.Println(item)
	fmt.Println(isAvailable)
	fmt.Println(firstName)
	fmt.Println(qty)
	fmt.Println(word)
	fmt.Println(heigh)
	fmt.Println(name)
	fmt.Printf("x=%d,y=%d,z=%d", x, y, z)
	fmt.Println("\ndeclaration variable IOTA")
	fmt.Printf("\n%.0f", KB)
	fmt.Printf("\n%.0f", MB)
	fmt.Printf("\n%.0f", GB)
	fmt.Println("\ntype data primitive")
	fmt.Println(tipeInteger8)
	fmt.Println(tipeInteger16)
	fmt.Println(tipeInteger32)
	fmt.Println(tipeInteger64)
	fmt.Println(tipeuInteger8)
	fmt.Println(tipeuInteger16)
	fmt.Println(tipeuInteger32)
	fmt.Println(tipeuInteger64)
	fmt.Println(tipeFloat32)
	fmt.Println(tipeFloat64)
	fmt.Println(tipeComplex64)
	fmt.Println(tipeComplex128)
	fmt.Println(tipeAliasInt32)
	fmt.Println(tipeAliasUint8)
	fmt.Println(str)
	fmt.Println(oneWord)
	fmt.Println(boolean)
	fmt.Println(isValid)

}
