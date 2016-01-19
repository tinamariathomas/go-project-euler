package main

import "fmt"

func main() {
	simple_loop();
}

func simple_loop() {
	for i := 1; i<=10;i++{
		fmt.Println(i);
	}
}
