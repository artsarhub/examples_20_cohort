package zoo

import "examples_20_cohort/interfaces/zoo/animals"

//go:generate mockgen -package=mock -destination=./mock/animal_mock.go yp-examples/interfaces/zoo Animal
type Animal interface {
	Feed()
	SayHello()
	GetFullness() int
	GetType() animals.AnimalType
}
