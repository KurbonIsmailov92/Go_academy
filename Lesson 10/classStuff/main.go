package main

import (
	"fmt"
)

func main() {
	var firstTeam, secondTeam, totalScoreOfFirstTeam, totalScoreOfSecondTeam uint

	for i := 0; i < 4; i++ {
		fmt.Scan(&firstTeam, &secondTeam)
		
		totalScoreOfFirstTeam += firstTeam
		totalScoreOfSecondTeam += secondTeam
	}
	
	if totalScoreOfFirstTeam > totalScoreOfSecondTeam {
		fmt.Println("1")
	}else if totalScoreOfFirstTeam < totalScoreOfSecondTeam {
		fmt.Println("2")
	}else{
		fmt.Println("DRAW")
	}

}

//Belka i oreshki
// var n, m, k uint
// fmt.Scan(&n, &m, &k)
// if n*m >= k {
// 	fmt.Println("YES")
// } else {
// 	fmt.Println("NO")
// }

//Dva bandita
// var a, b uint
// 	fmt.Scan(&a, &b)
// 	a, b = b-1, a-1
// 	fmt.Println(a, b)

// //maxminsalary
// var a, b, c uint
// 	var minS, maxS uint
// 	fmt.Scan(&a, &b, &c)

// 	if a >= b && a >= c {
// 		maxS = a
// 	}else if b >= a && b >= c {
// 		maxS = b
// 	}else{
// 		maxS = c
// 	}

// 	if a <= b && a <= c {
// 		minS = a
// 	}else if b <= a && b <= c {
// 		minS = b
// 	} else {
// 		minS = c
// 	}

// 	fmt.Println(maxS - minS)

//eniya
// var n, a, b int
// 	fmt.Scan(&n, &a, &b)
// 	result := n * a * b * 2
// 	fmt.Println(result)
