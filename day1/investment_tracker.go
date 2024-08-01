package main

import (
	"fmt"
	"math"
)

func main() {
	const infationrate = 2.5
	var investment float64
	var returnintrest float64
	var years float64

	fmt.Print("Enter the investment rate : ")
	fmt.Scan(&investment)

	fmt.Print("Enter the Expected Return Rate : ")
	fmt.Scan(&returnintrest)

	fmt.Print("Enter the years :")
	fmt.Scan(&years)

	var futureyear = investment * math.Pow(1+returnintrest/100, years)

	futureinflation := futureyear / math.Pow(1+infationrate/100, years)

	fmt.Println("The futureyear amount will be :", futureyear)

	fmt.Println("Future inflation amount will be :", futureinflation)

}
