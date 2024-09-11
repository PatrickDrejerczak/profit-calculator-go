package main

import (
	"fmt"
	"math"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	profitRatio := profit / revenue

	fmt.Println("Earnings before tax: ", earningsBeforeTax)
	fmt.Println("Profit: ", profit)
	fmt.Println("Profit ratio: ", math.Round(profitRatio*100), "%")
}
