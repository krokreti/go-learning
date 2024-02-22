package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

const inflationRate = 2.5

const balanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(balanceFile)
	if err != nil {
		return 1000, errors.New("failed to read file")
	}
	//nao da pra converter um array de bytes para float diretamente
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to store balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(balanceFile,[]byte(balanceText), 0644)
}

func main() {
	expectedReturnRate := 5.5
	var investmentAmount, years float64 = 1000 , 10

	var storedValue, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("------")
		panic("Can't continue sorry.")
	}

	fmt.Println(storedValue)

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

	writeBalanceToFile(futureValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	futureValue :=investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	realFutureValue := futureValue / math.Pow(1+inflationRate/100, years)
	return futureValue, realFutureValue
}