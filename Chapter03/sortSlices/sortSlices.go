package main

import (
	"fmt"
	"sort"
)

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	mySlice := make([]aStructure, 0)
	mySlice = append(mySlice, aStructure{"Мишаня", 180, 90})
	mySlice = append(mySlice, aStructure{"Борян", 134, 45})
	mySlice = append(mySlice, aStructure{"Ритусик", 155, 45})
	mySlice = append(mySlice, aStructure{"Епифанцев", 144, 50})
	mySlice = append(mySlice, aStructure{"Афинчик", 134, 40})

	fmt.Println("0:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height < mySlice[j].height
	})
	fmt.Println("<:", mySlice)

	sort.Slice(mySlice, func(i, j int) bool {
		return mySlice[i].height > mySlice[j].height
	})
	fmt.Println(">:", mySlice)
}
