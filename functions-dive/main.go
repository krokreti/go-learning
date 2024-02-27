package main

import "fmt"

func main() {
	//variadic function
	numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15, 40, -5)
	//pull out elements like ...javascript
	anotherSum := sumup(1, numbers...)
	fmt.Println(sum)
	fmt.Println(anotherSum)
}

// func sumup(numbers []int) int {
	func sumup(startingValue int,numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}