package main

import (
	"examples_20_cohort/interfaces/logger"
	"examples_20_cohort/interfaces/zoo"
)

func main() {

	// Functions
	/*
		//functions.SayHello()
		//functions.SayHelloToUser("Petya")
		//fmt.Println(functions.GetSum(1, 2))
		//fmt.Println(functions.GetSumOfAnyArgs())
		//fmt.Println(functions.ApplyOperation(10, 2, func(a, b int) int { return a - b }))
		//fmt.Println(functions.ApplyOperation(10, 2, functions.GetSum))
		//fmt.Println(functions.Validate("test", functions.ValidateLogin))
		//fmt.Println(functions.Validate("1234567890", functions.ValidatePassword))
		//fmt.Println(functions.Factorial(3))
		//functions.RunAnonymousFunction()
		//multiplication, division := functions.GetManyReturns(10, 2)
		//fmt.Println(multiplication, division)
		fmt.Println(functions.NamedReturn(1, 2))
	*/

	// Maps
	/*
		//maps.MapExample1()
		//maps.MapExample2()
		//maps.MapExample3()
		//maps.SliceVsMapExample1()
		maps.SliceVsMapExample2()
	*/

	// Functions extended
	/*
		//functions_extended.Example1()
		//functions_extended.Example2()
		//functions_extended.Example3()
		//functions_extended.Example4()
		//functions_extended.Example5()
		//functions_extended.Example6()
		//functions_extended.Example7()
		//functions_extended.Example8()
		//functions_extended.Example9()
		//functions_extended.Example10()
		//functions_extended.Example11()
		//functions_extended.Example12()
		functions_extended.Example13()
	*/

	// Debug
	/*
		//debug.MissingQuote()
		//debug.LogicError()
		//debug.RuntimeError()
		//debug.Logging()
		//debug.LoggingLevels()
		debug.StacktraceReading()
	*/

	// Pointers
	/*
		//pointers.Example1()
		//pointers.Example2()
		//pointers.Example3()
		//pointers.Example4()
		//test_mem.RunTestMem()
		pointers.TreeExample()
	*/

	// Interfaces
	///*
	log := logger.NewConsoleLogger()
	//log := logger.NewFileLogger("logs/service.log")

	//log.Info("test info message")
	//log.Warn("test warn message")
	//log.Error("test error message")

	zoo.EmulateZoo(log)
	//*/
}
