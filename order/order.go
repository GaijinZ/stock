package order

import (
	"fmt"
	"sort"
)

type Stocks struct {
	Stock []Stock
}

type Stock struct {
	ID       int
	Order    string
	Price    float64
	Quantity int
}

func (s *Stocks) Add(stock Stock) {
	s.Stock = append(s.Stock, stock)
}

func (s *Stocks) Remove(stockID int) {
	var updatedOrders []Stock

	for _, stock := range s.Stock {
		if stock.ID != stockID {
			updatedOrders = append(updatedOrders, stock)
		}
	}
	s.Stock = updatedOrders

}

func (s *Stocks) DisplayAll() {
	buyOrders := make(map[float64]int)
	sellOrders := make(map[float64]int)

	for _, stock := range s.Stock {
		switch stock.Order {
		case "Buy":
			buyOrders[stock.Price] += stock.Quantity
		case "Sell":
			sellOrders[stock.Price] += stock.Quantity
		}
	}

	fmt.Println("Buy Orders:")
	for _, price := range sortAscending(getKeys(buyOrders)) {
		fmt.Printf("Price: %.2f, Quantity: %d\n", price, buyOrders[price])
	}

	fmt.Println("\nSell Orders:")
	for _, price := range sortDescending(getKeys(sellOrders)) {
		fmt.Printf("Price: %.2f, Quantity: %d\n", price, sellOrders[price])
	}
}

func getKeys(m map[float64]int) []float64 {
	var keys []float64

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

func sortDescending(keys []float64) []float64 {
	sort.Sort(sort.Reverse(sort.Float64Slice(keys)))
	return keys
}

func sortAscending(keys []float64) []float64 {
	sort.Sort(sort.Float64Slice(keys))
	return keys
}
