package main

import "fmt"

func main() {
	var b Budget

	b.Tatal = 1000.00
	b.Owner = "Vasya"
	b.Expences = map[string]float32{}

	b.AddExpence("Transport", 300)

	fmt.Println(b)
}

type Budget struct {
	Tatal    float32
	Expences map[string]float32
	Owner    string
}

func (b *Budget) AddExpence(cathegory string, amount float32) {
	b.Expences[cathegory] = amount
	b.Tatal -= amount
}
