package functions_extended

import "fmt"

func logStartAndEnd(name string) func() func() {
	fmt.Printf("Начало функции %s\n", name)
	return func() func() {
		fmt.Println("Another func")
		return func() {
			fmt.Printf("Конец функции %s\n", name)
		}
	}
}

func Example9() {
	defer logStartAndEnd("main")()() // Логирование выполнения функции main

	fmt.Println("Основная работа функции")
	// Output:
	// Начало функции main
	// Основная работа функции
	// Конец функции main
}
