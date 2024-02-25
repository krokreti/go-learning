package lists

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := [4]float64{1, 2, 3, 4}
	products := []Product{{title: "tv", id: "1", price: 12.99}, {title: "house", id: "2", price: 15.9}}

	fmt.Println(products)
	//slices
	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)
	//cant use negative indexes
	fmt.Println(prices[:3])
	fmt.Println(prices[1:])
	fmt.Print(len(prices[1:]), cap(prices))

	newPrices := []float64{1, 2}
	newPrices = append(newPrices, 5.99)

}