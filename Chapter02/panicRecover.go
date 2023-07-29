package main

import "fmt"

func a() {
	fmt.Println("Внутри a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Восстановление внутри а()!")
		}
	}()
	fmt.Println("Что насчет вызова b()?")
	b()
	fmt.Println("Вызов b() завершен")
	fmt.Println("Выход из a()")
}

func b() {
	fmt.Println("Мы внутри b()")
	panic("Атаческая паника в b()")
	fmt.Println("Выходим из b()")
}

func main() {
	a()
	fmt.Println("main() окончен")
}
