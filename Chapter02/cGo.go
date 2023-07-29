package main

// #include <stdio.h>
// void callC() {
//printf("Calling C code!\n");
// }
import "C"
import "fmt"

func main() {
	fmt.Println("Вывод Go!")
	C.callC()
	fmt.Println("И снова вывод Go!")
}
