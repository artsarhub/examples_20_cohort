package debug

import (
	"fmt"
)

func StacktraceReading() {
	fmt.Println("Пример многомодульной программы")
	res, err := Divide(4, 0)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
	fmt.Println(res) // Вызовет деление на ноль
}

func Divide(a, b int) (int, error) {
	if b <= 0 {
		return 0, fmt.Errorf("b = %d", b)
	}
	return a / b, nil // Здесь возникнет ошибка времени выполнения
}
