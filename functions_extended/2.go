package functions_extended

import "fmt"

// Функция, которая принимает другую функцию как аргумент
func applyOperation(a, b int, operation func(int, int) int) int {
	fmt.Println(a)
	fmt.Println(b)
	return operation(a, b)
}

func Example2() {
	// Обычная функция сложения
	add := func(a, b int) int {
		return a + b
	}

	result := applyOperation(5, 3, add)
	fmt.Println("Результат операции:", result) // Output: Результат операции: 8
}
