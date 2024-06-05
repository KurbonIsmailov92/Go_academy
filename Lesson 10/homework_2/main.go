package main

import "fmt"

// Сотрудники завода по производству золотого песка из воздуха решили поправить свое финансовое положение. 
// Они пробрались на склад завода, где хранился золотой песок трех видов. 
// Один килограмм золотого песка первого вида они смогли бы продать за A1 рублей, второго вида – за A2 рублей,
// а третьего вида – за A3 рублей. Так получилось, что у сотрудников оказалось с собой только три емкости: 
// первая была рассчитана на B1 килограмм груза, вторая на B2 килограмм, а третья на B3 килограмм. 
// Им надо было заполнить полностью все емкости таким образом, чтобы получить как можно больше денег за весь песок. 
// При заполнении емкостей нельзя смешивать песок разных видов, то есть, в одну емкость помещать более одного вида 
// песка, и заполнять емкости песком так, чтобы один вид песка находился более чем в одной емкости.

// Требуется написать программу, которая определяет, за какую сумму предприимчивые сотрудники смогут продать 
// весь песок в случае наилучшего для себя заполнения емкостей песком.

func main() {
	var firstSort, secondSort, thirdSort, firstJar, secondJar, thirdJar, bestWeyToGetRich uint

	fmt.Scan(&firstSort, &secondSort, &thirdSort, &firstJar, &secondJar, &thirdJar)

	if firstSort < secondSort {
		firstSort, secondSort = secondSort, firstSort
	}

	if secondSort < thirdSort {
		secondSort, thirdSort = thirdSort, secondSort
	}

	if firstSort < secondSort {
		firstSort, secondSort = secondSort, firstSort
	}

	if firstJar < secondJar {
		firstJar, secondJar = secondJar, firstJar
	}

	if secondJar < thirdJar {
		secondJar, thirdJar = thirdJar, secondJar
	}

	if firstJar < secondJar {
		firstJar, secondJar = secondJar, firstJar
	}

	bestWeyToGetRich = firstJar*firstSort + secondJar*secondSort + thirdJar*thirdSort

	fmt.Println(bestWeyToGetRich)

}
