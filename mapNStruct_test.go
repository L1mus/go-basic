package gobasic

import (
	"fmt"
	"log"
	"testing"
)

func TestMapNStruct(t *testing.T) {
	mapType()
	structType()
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

type Product struct {
	ID          int
	Name        string
	Price       float64
	Stock       int
	IsAvailable bool
}

func NewProduct(id int, name string, price float64, stock int, available bool) *Product {
	return &Product{
		ID:          id,
		Name:        name,
		Price:       price,
		Stock:       stock,
		IsAvailable: available,
	}
}

func (p *Product) UpdateAvailability() {
	if p.Stock > 0 {
		p.IsAvailable = true
	} else {
		p.IsAvailable = false
	}
}

func (p *Product) BuyItem() {
	if p.Stock > 0 {
		p.Stock--
	} else {
		p.Stock = 0
	}

}

func structType() {

	item1 := NewProduct(1, "Sugar", 5.00, 10, true)
	item2 := NewProduct(2, "Garlic", 0.40, 100, true)
	item3 := NewProduct(3, "Carrot", 0.70, 50, true)
	item4 := NewProduct(4, "Fluor", 0.30, 1, true)

	products := make([]Product, 0)
	products = append(products, *item1, *item2, *item3, *item4)

	for _, v := range products {
		fmt.Printf(`
		Nama Produk : %s
		harga : $%0.2f
		stock : %d
		status ketersediaan : %v
		`, v.Name, v.Price, v.Stock, v.IsAvailable)
	}

	simulationStockZero(&products[3])
	simulationStockZero(&products[3])

	var totalInventaris float64

	for _, v := range products {
		totalInventaris += v.Price * float64(v.Stock)
	}

	fmt.Printf("\njumlah product : %d \nTotal Inventaris semua product adalah : $%0.2f\n", len(products), totalInventaris)
}

func simulationStockZero(item *Product) {
	fmt.Printf("\nBefore buy : %v", item)

	//check availablelity
	available := item.Stock > 0

	if !available {
		fmt.Println("\nItem out of stock")
		return
	}

	//buy simulation
	item.BuyItem()

	item.UpdateAvailability()

	fmt.Printf("\nAfter buy : %v", item)

	//asumsiton succes buy
	fmt.Println("\nThanks for buying, have a nice day")

}
