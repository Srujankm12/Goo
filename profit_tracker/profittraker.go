package main

import "fmt"

func main() {

	var revenue float64
	var Expenses float64
	var Tax float64

	fmt.Print("Enter the revenue :")
	fmt.Scan(&revenue)

	fmt.Print("Enter the expenses :")
	fmt.Scan(&Expenses)

	fmt.Print("Enter the Taxrate :")
	fmt.Scan(&Tax)

	ebt := revenue - Expenses
	Profit := ebt * (1 - Tax/100)
	ratio := ebt / Profit

	fmt.Println("the earning before tax is : ", ebt)
	fmt.Println("The profit is ", Profit)
	fmt.Println("The ratio of ebt and profit is", ratio)
}
