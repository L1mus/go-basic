package gobasic

import (
	"log"
	"strconv"
	"testing"
)

func TestConvertionBasic(t *testing.T) {
	// number to number
	int8ToInt64 := int64(20)
	log.Printf("\nconvertion ke int64: %v", int8ToInt64)

	int32Tofloat32 := float32(3000)
	log.Printf("\nconvertion ke FLOAT32: %v", int32Tofloat32)

	//string to number
	//strconv
	number, _ := strconv.Atoi("20")
	log.Printf("\nconvertion ke Number dari string: %d", number)
	numberFormatint32, _ := strconv.ParseInt("50", 10, 8)
	log.Printf("\nconvertion ke Number dari string dengan lebih spesifik: %d", numberFormatint32)
	numberFormatfloat32, _ := strconv.ParseFloat("50", 32)
	log.Printf("\nconvertion ke Number float dari string dengan lebih spesifik: %0.2f", numberFormatfloat32)
	//number to string
	str := strconv.Itoa(20)
	log.Printf("\nconvertion ke String %s", str)

	// string to byte slice
	word := "Hello World!"
	wordOfByte := []byte{72, 101, 108, 108, 111}
	arrString := []byte(word)
	log.Printf("\nconvertion ke slice of byte ke bentuk ASCII Format: %v", arrString)
	//byteslice to string
	strByte := string(wordOfByte)
	log.Printf("\nconvertion ke string: %s", strByte)

	// string to slice of rune
	r := []rune(word)
	log.Printf("\nconvertion ke slice of rune: %v", r)

	// slice of rune to string
	sRune := string(r)
	log.Printf("\nconvertion ke string: %s", sRune)
}
