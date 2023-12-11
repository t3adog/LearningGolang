package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	validateArgs(args)

	sum := 0.0
	for i := 1; i < len(args); i++ {
		n, err := strconv.ParseFloat(args[i], 64)
		if err == nil {
			sum = sum + n
		}
	}

	avg := sum / float64((len(args) - 1))

	fmt.Println("Avg:", avg)
}

func validateArgs(args []string) {
	if len(args) == 1 {
		fmt.Println("Нужно больше аргументов")
		os.Exit(1)
	}
}
