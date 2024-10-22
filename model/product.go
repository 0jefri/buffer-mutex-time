package model

import (
	"fmt"
	"sync"
)

type Product struct {
	Name string
	Qty  int
}

func Reduce(n int, wg *sync.WaitGroup) {
	// var produk Product
	product := &Product{
		Name: "Baju",
		Qty:  10,
	}
	defer wg.Done()

	if product.Qty >= n {
		product.Qty -= n
		fmt.Printf("Product %s berkurang %d sisanya %d\n", product.Name, n, product.Qty)
	} else {
		fmt.Println("product tidak cukup untuk dikurangi")
	}
}
