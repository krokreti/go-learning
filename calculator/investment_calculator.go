package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	expectedReturnRate := 5.5
	var investmentAmount, years float64 = 1000 , 10

	// fmt.Println("Please enter a investment value: ")
	outputText("Please enter a investment value:")
	fmt.Scan(&investmentAmount)

	fmt.Println("How long should you be waiting for? ")
	fmt.Scan(&years)

	fmt.Println("Whats your expected rate? ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years);

	//fmt.Println("Future value: ", futureValue);
	// fmt.Println("Future value (inflation): ", futureRealValue)
	fmt.Printf("Future Value: %.2f\nFuture value (inflation): %.2f\n", futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	futureValue :=investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, realFutureValue
}