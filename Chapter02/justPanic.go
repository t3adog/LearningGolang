package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		panic("!!!МНЕ НЕДОСТАТОЧНО АРГУМЕНТОВ!!!")
	}

	fmt.Println("Аргументов достаточно, спасибо.")
}
