package main

import "fmt"

type myStruct struct {
	Name    string
	Surname string
	Height  int32
}

func createStruct(n, s string, h int32) *myStruct {
	if h > 300 {
		h = 0
	}
	return &myStruct{n, s, h}
}

func retStructure(n, s string, h int32) myStruct {
	if h > 300 {
		h = 0
	}
	return myStruct{n, s, h}
}

func main() {
	s1 := createStruct("Mihails", "Tsoukalos", 123)
	s2 := retStructure("Mihails", "Tsoukalos", 123)

	fmt.Println((*s1).Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)

	f := fmt.Println

	f("123")
}
