package zoo

import (
	"fmt"

	"examples_20_cohort/interfaces/logger"
	"examples_20_cohort/interfaces/zoo/animals"
)

func EmulateZoo(log logger.Logger) {
	var animalsArr []Animal

	for _, name := range []string{"Bars", "Murzik"} {
		newCat := animals.NewCat(name, 100, log)
		animalsArr = append(animalsArr, &newCat)
	}
	for _, name := range []string{"Sharik", "Rex"} {
		newDog := animals.NewDog(name, 100, log)
		animalsArr = append(animalsArr, &newDog)
	}

	for _, animal := range animalsArr {
		switch animal.(type) {
		case *animals.Cat:
			log.Info("Is's a cat")
		case *animals.Dog:
			log.Info("Is's a dog")
		default:
			log.Info("Unknown type")
		}
		animal.SayHello()
		log.Warn(fmt.Sprintf("My fullness is %d for now.", animal.GetFullness()))
		animal.Feed()
		log.Error("")
	}
}
