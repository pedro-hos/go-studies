package main

import (
	"fmt"
	"go-studies/go-as-oop/pets"
	"time"
)

func main() {
	lastSlept := time.Now()
	//lastSlept := time.Now().Add(time.Duration(-5) * time.Hour)

	var animals []pets.Pet

	animals = append(animals, pets.NewDog("Luke", "Black and White", "mixed", lastSlept))
	animals = append(animals, pets.NewCat("Duke", "Black", "mixed"))

	for _, pet := range animals {
		if pet.IsHungry() {
			fmt.Println(pet.Feed("steak"))
		} else {
			fmt.Println("Your animal is not hungry, waiting ...")
			time.Sleep(5 * time.Second)
			fmt.Println(pet.Feed("kibble"))
		}

		fmt.Println(pet.GiveAttention("playing fetch"))
	}
}
