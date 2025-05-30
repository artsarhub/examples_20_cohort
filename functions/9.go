package functions

func NamedReturn(a, b int) (res int) {
	defer func() {
		res = 0
	}()
	res = a + b
	// Вернётся 0 потому что в данном случае возвращаемое значение будет вычисляться после return, но до выхода из функции через defer
	return res
}
