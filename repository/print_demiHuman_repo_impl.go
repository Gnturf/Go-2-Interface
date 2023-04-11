package repository

import (
	"fmt"

	"github.com/NaylaDeLis/Go-2-Interface/entity"
)

type personDemiHumanRepository struct {
	demiHuman entity.DemiHumanEntity
}

func CreateNewDemiHumanEntity(PersonDemiHuman entity.DemiHumanEntity) PrintGreeting {
	return &personDemiHumanRepository{demiHuman: PersonDemiHuman}
}

func (repo *personDemiHumanRepository) PrintWithHi() {
	fmt.Println("Hi " + repo.demiHuman.Name + " (" + repo.demiHuman.Race + ")")
}

func (repo *personDemiHumanRepository) PrintWithSorry() {
	fmt.Println("Sorry " + repo.demiHuman.Name + " (" + repo.demiHuman.Race + ")")
}

func (repo *personDemiHumanRepository) PrintWithGoodBye() {
	fmt.Println("Good Bye " + repo.demiHuman.Name + " (" + repo.demiHuman.Race + ")")
}
