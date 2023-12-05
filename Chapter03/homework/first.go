package main

import "fmt"

type DayOfWeek int

const (
	Monday DayOfWeek = iota
	Tusday
)

func main() {

	fmt.Println(Monday)
	fmt.Println(Tusday)
}
