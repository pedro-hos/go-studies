package pets

import (
	"fmt"
	"strings"
	"time"
)

type Dog struct {
	Name      string
	Color     string
	Breed     string
	lastSlept time.Time
	Animal
}

func (d Dog) needsSleep() bool {
	return time.Since(d.lastSlept) > 4*time.Hour
}

func (d Dog) sleep() {
	d.lastSlept = time.Now()
}

func (d Dog) GiveAttention(activity string) string {

	if d.needsSleep() {
		d.sleep()
		return "Your dog is tired and needs to rest"
	}

	response := ""
	switch strings.ToUpper(activity) {
	case "PET":
		response = "wags tail"
	case "PLAYING FETCH":
		response = "return the ball and jump waiting for you to throw it again"
	default:
		response = "barks"
	}

	return fmt.Sprintf("%s loves attention, %s will cause him %s", d.Name, activity, response)
}

func NewDog(name string, color string, breed string, lastSlept time.Time) Pet {
	return Dog{
		Name:      name,
		Color:     color,
		Breed:     breed,
		lastSlept: lastSlept,
		Animal: Animal{
			lastAte: time.Now(),
		},
	}
}
