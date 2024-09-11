package main

import (
	"fmt"
	"math"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getRequiredInput("Enter revenue: ", revenue)
	expenses = getRequiredInput("Enter expenses: ", expenses)
	taxRate = getRequiredInput("Enter tax rate: ", taxRate)

	ebt, profit, profitRatio := calculateEarnings(revenue, expenses, taxRate)

	fmt.Println("Earnings before tax: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Profit ratio: ", math.Round(profitRatio*100), "%")
}

func getRequiredInput(text string, value float64) float64 {
	fmt.Print(text)
	fmt.Scan(&value)

	return value
}

func calculateEarnings(revenue float64, expenses float64, taxRate float64) (ebt float64, profit float64, profitRatio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	profitRatio = profit / revenue

	return ebt, profit, profitRatio
}
