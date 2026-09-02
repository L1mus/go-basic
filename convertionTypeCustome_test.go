package gobasic

import (
	"log"
	"testing"
)

func TestConvertionTypeCustome(t *testing.T) {
	conveeerstion()
}

func conveeerstion() {
	type MyType int

	var initialType int8 = 8

	convert := MyType(initialType)

	log.Printf("\nConvert to type custome MyType : %d", convert)

	convertToInt16 := int16(convert)

	log.Printf("\nConvert to type custome int16 : %d", convertToInt16)

}
