package animals

import (
	"fmt"

	"examples_20_cohort/interfaces/logger"
)

type Cat struct {
	AnimalStruct
	tailLength int
	log        logger.Logger
}

func NewCat(name string, fullness int, log logger.Logger) Cat {
	cat := Cat{}
	cat.name = name
	cat.fullness = fullness
	cat.animalType = TypeCat
	cat.log = log
	return cat
}

func (c *Cat) Feed() {
	c.fullness++
	c.log.Info(fmt.Sprintf("So yummy! My fullness is %d.", c.fullness))
}

func (c *Cat) SayHello() {
	c.log.Info(fmt.Sprintf("Hi! My name is %s. I'm cat.", c.name))
	c.fullness -= 2
}

func (c *Cat) GetTailLength() int {
	return c.tailLength
}
