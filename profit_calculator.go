package main

import (
	"fmt"
	"math"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getRequiredInput("Enter revenue: ")
	expenses = getRequiredInput("Enter expenses: ")
	taxRate = getRequiredInput("Enter tax rate: ")

	ebt, profit, profitRatio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println("Earnings before tax: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Profit ratio: ", math.Round(profitRatio*100), "%")
}

func getRequiredInput(text string) float64 {
	var value float64

	fmt.Print(text)
	fmt.Scan(&value)

	return value
}

func calculateFinancials(revenue, expenses, taxRate float64) (ebt, profit, profitRatio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	profitRatio = profit / revenue

	return ebt, profit, profitRatio
}
