package gobasic

import (
	"fmt"
	"log"
	"testing"
)

func TestMapNStruct(t *testing.T) {
	mapType()
}

func mapType() {
	//error dereference to map with value nil cause not initialitation the value to variable map
	// var mapNil map[string]string
	// mapNil["test"] = "result"

	mapping := make(map[string]int64)

	log.Printf("\nempty map: %v", mapping)

	mapping["box"] = 56
	mapping["item"] = 4

	log.Printf("\nMap after insert value:\n%v", mapping)

	//Comma-Ok idiom type map
	var stocks map[string]float64
	sym := "TTWO"
	price := stocks[sym]
	fmt.Printf("\n\n1. %s -> $%.2f\n", sym, price)

	if price, ok := stocks[sym]; ok { // using comma ok to check if key exists
		fmt.Printf("2. %s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("2. %s not found\n", sym)
	}

	stocks = make(map[string]float64) // It's importatnt to initialize the map before adding values. Else it's a panic situation
	stocks[sym] = 123.4
	stocks["APPL"] = 345.6

	if price, ok := stocks[sym]; ok {
		fmt.Printf("3. %s -> $%.2f\n", sym, price)
	} else {
		fmt.Printf("3. %s not found\n", sym)
	}

}
