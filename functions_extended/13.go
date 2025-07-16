package functions_extended

import "fmt"

func Example13() {
	var funcs []func()
	for i := 0; i < 3; i++ {
		i := i
		funcs = append(funcs, func() {
			fmt.Printf("%d\n", i)
		})
	}

	/*
		Все функции выведут 3, поскольку замыкание захватывает одну и ту же переменную `i`
	*/
	for _, f := range funcs {
		f()
	}
}
