package main

import (
	"fmt"
)

func calculateProfit(sellingPrice int, purchasePrice int, quantity int) int {
	pricePerItem := sellingPrice - purchasePrice
	totalProfit := pricePerItem * quantity

	return totalProfit
}

func main() {
	productName1 := "Peak Milk"
	milkProfit := calculateProfit(1500, 1200, 4)

	productName2 := "Milo"
	miloProfit := calculateProfit(1200, 1000, 5)

	productName3 := "indomie"
	indomieProfit := calculateProfit(800, 600, 10)

	aggregate := "Total"
	combinedTotal := milkProfit + miloProfit + indomieProfit
	fmt.Println(productName1, "Profit", milkProfit)
	fmt.Println(productName2, "Profit", miloProfit)
	fmt.Println(productName3, "Profit", indomieProfit)
	fmt.Println(aggregate, "profit", combinedTotal)
}
