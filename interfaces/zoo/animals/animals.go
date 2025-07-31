package animals

type AnimalType int

const (
	TypeCat AnimalType = iota
	TypeDog
)

type AnimalStruct struct {
	name       string
	fullness   int
	animalType AnimalType
}

func (c AnimalStruct) GetFullness() int {
	return c.fullness
}

func (c AnimalStruct) GetType() AnimalType {
	return c.animalType
}
