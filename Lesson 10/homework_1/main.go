package main

import "fmt"

// Три толстяка решили поспорить: кто из них самый тяжелый. После взвешивания оказалось, что их масса соответственно M1, M2 и M3 килограмм.
// Считается, что масса толстяка должна быть не менее 94 и не более 727 килограмм.

// Помогите определить массу самого тяжелого из них, либо выяснить, что была допущена ошибка при взвешивании.

func main() {

	var firstMass, secondMass, thirdMass, winner, minWeight, maxWeight uint

	minWeight = 94
	maxWeight = 727

	fmt.Scan(&firstMass, &secondMass, &thirdMass)

	if ((firstMass < minWeight) || (secondMass < minWeight) || (thirdMass < minWeight)) ||
		((firstMass > maxWeight) || (secondMass > maxWeight) || (thirdMass > maxWeight)) {

		fmt.Println("Error")
		return
	}

	if firstMass > 10_000 || secondMass > 10_000 || thirdMass > 10_000 {

		fmt.Println("Error")
		return
	}

	winner = thirdMass

	if firstMass > winner {
		winner = firstMass
	} else if secondMass > winner {
		winner = secondMass
	}

	fmt.Println(winner)

}
