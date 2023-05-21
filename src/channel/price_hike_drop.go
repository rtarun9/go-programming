package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Product struct {
	Name  string
	Price int
}

func captureHike(product *Product, hikeChan chan int) {
	hikedPrice := product.Price * 3
	hikeChan <- hikedPrice
}

func captureReduction(product *Product, reductionChan chan int) {
	reducedPrice := product.Price - 50
	reductionChan <- reducedPrice
}

func main() {
	rand.Seed(time.Now().UnixNano())

	products := []Product{
		{Name: "Product A", Price: 100},
		{Name: "Product B", Price: 200},
		{Name: "Product C", Price: 300},
		{Name: "Product D", Price: 400},
		{Name: "Product E", Price: 500},
	}

	hikeChan := make(chan int)
	reductionChan := make(chan int)

	for i := 0; i < len(products); i++ {
		go captureHike(&products[i], hikeChan)
		go captureReduction(&products[i], reductionChan)
	}

	for i := 0; i < len(products); i++ {
		select {
		case hikedPrice := <-hikeChan:
			fmt.Printf("%s - Price Hiked: %d\n", products[i].Name, hikedPrice)
		case reducedPrice := <-reductionChan:
			fmt.Printf("%s - Price Reduced: %d\n", products[i].Name, reducedPrice)
		}
	}
}
