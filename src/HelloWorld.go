package main

import "fmt"

func main() {
	euler1(1000);
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

