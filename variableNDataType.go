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
