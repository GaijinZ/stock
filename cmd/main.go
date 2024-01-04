package main

import "stock_exchange/order"

func main() {
	orderBook := order.Stocks{}

	orderBook.Add(order.Stock{ID: 001, Order: "Buy", Price: 20.0, Quantity: 100})
	orderBook.Add(order.Stock{ID: 002, Order: "Sell", Price: 25.0, Quantity: 200})
	orderBook.Add(order.Stock{ID: 003, Order: "Buy", Price: 23.0, Quantity: 50})
	orderBook.Add(order.Stock{ID: 004, Order: "Buy", Price: 23.0, Quantity: 70})
	orderBook.Remove(003)
	orderBook.Add(order.Stock{ID: 005, Order: "Sell", Price: 28.0, Quantity: 100})

	orderBook.DisplayAll()
}
