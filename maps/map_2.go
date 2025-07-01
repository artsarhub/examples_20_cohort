package maps

import "fmt"

func MapExample2() {
	m := map[int]int{1: 1, 2: 2}
	w := 5

	w = changeIntValue(w)
	fmt.Println(w)

	changeMapValue(m, 1)
	fmt.Println(m)

	//tmp := &m[1]
	//fmt.Println(tmp)
}

func changeIntValue(v int) int {
	v = 10 + v
	return v
}

func changeMapValue(v map[int]int, key int) {
	v[key] += 10
}
