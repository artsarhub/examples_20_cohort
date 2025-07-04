package maps

import "fmt"

func MapExample1() {
	// Инициализация
	m := make(map[int]int)
	//var m map[int]int
	//m := map[int]int{}

	// Вставка
	m[2] = 0
	fmt.Println(m)

	// Поиск
	if v, ok := m[29]; ok {
		fmt.Println(v)
	}
	v, ok := m[29]
	if ok {
		fmt.Println("Value found", v)
	} else {
		fmt.Println("Value not found")
	}
	vv := m[2]
	fmt.Println(vv)

	// Удаление
	delete(m, 2)
	fmt.Println(m)
}
