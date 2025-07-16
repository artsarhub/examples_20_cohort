package functions_extended

import "fmt"

// Возвращает функцию, которая добавляет заданное значение
func createAdder(value int) func(int) int {
	return func(x int) int {
		return x + value
	}
}

func Example3() {
	addFive := createAdder(5)
	v := addFive(10)
	fmt.Printf("Добавление %d к 10: %d\n", 5, v) // Output: Добавление 5 к 10: 15
}
