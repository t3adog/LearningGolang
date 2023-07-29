package main

import "fmt"

func main() {
	d1()
	d2()
	fmt.Println()
	d3()
	fmt.Println()
	d4()
	fmt.Println()
}

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
}

func d4() {
	for i := 3; i > 0; i-- {
		defer print(i)
	}
}

func print(i int) {
	fmt.Print(i, " ")
}
