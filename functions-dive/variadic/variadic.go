package variadic

import "fmt"

func main() {
	//variadic function
	// numbers := []int{1, 10, 15}
	sum := sumup(1, 10, 15, 40, -5)
	fmt.Println(sum)
}

// func sumup(numbers []int) int {
	func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}
	return sum
}