package main

import "fmt"

func main() {
	euler2(4000000);
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
