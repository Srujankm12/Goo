package main

import (
	"fmt"
	"math"
)
const infationrate = 2.5
	var investment float64
	var returnintrest float64
	var years float64

func main() {
	

	outpuText("Enter the investment rate : ")
	fmt.Scan(&investment)

	outpuText("Enter the Expected Return Rate : ")
	fmt.Scan(&returnintrest)

	outpuText("Enter the years :")
	fmt.Scan(&years)

	futureyear := investment * math.Pow(1+returnintrest/100, years)

	futureinflation := futureyear / math.Pow(1+infationrate/100, years)

	formeattedFv := fmt.Sprintf("The futureyear amount will be : %.1f\n", futureyear)

	formattedRFV := fmt.Sprintf("Future inflation amount will be : %.1f\n", futureinflation)

	fmt.Print(formeattedFv, formattedRFV)

}
func outpuText(text string) {
	fmt.Print(text)
}

func calculatefuturevalues(investment , returnintrest , years float64 ){

	fv := investment * math.Pow(1+returnintrest/100, years),
	
	rfv := futureyear / math.Pow(1+infationrate/100, years)
	return fv , rfv


}

