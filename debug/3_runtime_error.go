package debug

import "fmt"

func RuntimeError() {
	arr := []int{1, 2, 3}
	const fifthElement = 5
	fmt.Println(arr[fifthElement]) // Доступ к несуществующему индексу вызовет панику
}
