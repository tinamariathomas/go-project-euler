package main;

import "fmt"
import "math"


func euler5(start int, end int) {
	listOfPrimes := findPrimeNumbers(start, end)
	mapOfPrimes := make(map[int]int)
	for i := 0; i < len(listOfPrimes); i++ {
		mapOfPrimes[listOfPrimes[i]] = 0
	}
	for k := range mapOfPrimes {
		primeNumber := k;
		for i := start; i <= end; i++ {
			number := i
			count := 0
			for number > 1 && number % primeNumber == 0 {
				count = count + 1
				number = number / primeNumber
			}
			if (mapOfPrimes[primeNumber] < count) {
				mapOfPrimes[primeNumber] = count
			}
		}
	}

	var result float64 = 1
	for k := range mapOfPrimes {
		result = result * math.Pow(float64(k),float64(mapOfPrimes[k]))
	}

	fmt.Println(int(result))

}

func findPrimeNumbers(start int, end int) []int {
	listOfPrimeNumbers := make([]int, 0)
	listOfPrimeNumbers = append(listOfPrimeNumbers, 2)
	for i := 3; i <= end; i = i + 2 {
		if (isPrime(i)) {
			listOfPrimeNumbers = append(listOfPrimeNumbers, i)
		}
	}
	return listOfPrimeNumbers
}

func isPrime(number int) bool {
	var flag bool = true
	for i := 3; i < number / 2; i = i + 2 {
		if number % i == 0 {
			flag = false
			break
		}
	}
	return flag
}





func euler1(limit int) {
	sumOfValues := 0

	for i := 1; i < limit; i++ {
		if i % 5 == 0 || i % 3 == 0 {
			sumOfValues = sumOfValues + i
		}
	}
	fmt.Println(sumOfValues)
}


func euler2(limit int) {
	var x, y int = 1, 2;
	var z int;
	var sum int = 2;

	for {
		z = x + y;
		if z > limit {
			break
		}
		if (z % 2 == 0) {
			sum = sum + z;
		}
		x = y
		y = z
	}

	fmt.Println(sum)

}
