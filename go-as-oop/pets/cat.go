package pets

import (
	"fmt"
	"time"
)

type Cat struct {
	Name  string
	Color string
	Breed string
	Animal
}

func (c Cat) GiveAttention(activity string) string {
	return fmt.Sprintf("The cat is ignoring your attemps to %s", activity)
}

func NewCat(name, color, breed string) Pet {
	return Cat{
		Name:  name,
		Color: color,
		Breed: breed,
		Animal: Animal{
			lastAte: time.Now(),
		},
	}
}
