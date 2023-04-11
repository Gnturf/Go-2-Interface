package repository

import (
	"fmt"

	"github.com/NaylaDeLis/Go-2-Interface/entity"
)

type PersonHumanRepository struct {
	human entity.PersonEntity
}

func CreateNewHumanEntity(humanPerson entity.PersonEntity) PrintGreeting {
	return &PersonHumanRepository{human: humanPerson}
}

func (repo *PersonHumanRepository) PrintWithHi() {
	fmt.Println("Hi " + repo.human.Name + " (Human)")
}

func (repo *PersonHumanRepository) PrintWithSorry() {
	fmt.Println("Sorry " + repo.human.Name + " (Human)")
}

func (repo *PersonHumanRepository) PrintWithGoodBye() {
	fmt.Println("Good Bye " + repo.human.Name + " (Human)")
}
