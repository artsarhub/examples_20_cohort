package functions_extended

import "fmt"

func Example12() {
	fmt.Println("Начало работы")
	safeFunction()
	fmt.Println("Продолжение после recover")
}

func safeFunction() {
	defer func() {
		fmt.Println("DEFER")
		if r := recover(); r != nil {
			fmt.Println("Recover from:", r)
		}
	}()

	fmt.Println("До вызова panic")
	panic("что-то пошло не так")
	fmt.Println("Этот код не будет выполнен")
}
