package animals

import (
	"fmt"

	"examples_20_cohort/interfaces/logger"
)

type Dog struct {
	AnimalStruct
	log logger.Logger
}

func NewDog(name string, fullness int, log logger.Logger) Dog {
	dog := Dog{}
	dog.name = name
	dog.fullness = fullness
	dog.animalType = TypeDog
	dog.log = log
	return dog
}

func (d *Dog) Feed() {
	d.fullness++
	d.log.Info(fmt.Sprintf("So yummy! My fullness is %d.", d.fullness))
}

func (d *Dog) SayHello() {
	d.log.Info(fmt.Sprintf("Hi! My name is %s. I'm dog.", d.name))
	d.fullness -= 5
}
