package main;

import (
	"fmt"
)

//func euler10(number int) {
//
//	list := findPrimeNumbers(2, number)
//	fmt.Println("Done")
//	sum := 0
//	for i :=0; i<len(list);i=i+1{
//		sum = sum + list[i]
//	}
//	fmt.Println(sum)
//}

func euler10_eratosthenes(number int) {
	type value struct {
		number     int
		crossedOff bool
	}

	rangeOfNumbers := make([]value, 0)
	primeList := make([]int, 0)

	for i := 2; i <= number; i++ {
		rangeOfNumbers = append(rangeOfNumbers, value{i, false})
	}


	for i := 0; i < number - 1; i++ {
		if (rangeOfNumbers[i].crossedOff == false) {
			primeList = append(primeList, rangeOfNumbers[i].number)
			for j := i; j < number - 1; j = j + rangeOfNumbers[i].number {
				rangeOfNumbers[j].crossedOff = true
			}
		}

	}
	result :=0
	for i:=0;i<len(primeList);i++{
		result = result + primeList[i]
	}
	fmt.Println(result)

}

