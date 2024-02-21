package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	expectedReturnRate := 5.5
	var investmentAmount, years float64 = 1000 , 10

	fmt.Println("Please enter a investment value: ")
	fmt.Scan(&investmentAmount)

	fmt.Println("How long should you be waiting for? ")
	fmt.Scan(&years)

	fmt.Println("Whats your expected rate? ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue);
	fmt.Println(futureRealValue);
}