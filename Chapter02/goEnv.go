package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Ты используешь ", runtime.Compiler, " ")
	fmt.Println("на ", runtime.GOARCH, " компе")
	fmt.Println("Используемая GO версия: ", runtime.Version())
	fmt.Println("Количество процов: ", runtime.NumCPU())
	fmt.Println("Количество горутин: ", runtime.NumGoroutine())
}
