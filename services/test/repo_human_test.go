package test

import (
	"testing"

	"github.com/NaylaDeLis/Go-2-Interface/entity"
	"github.com/NaylaDeLis/Go-2-Interface/repository"
)

func TestPrintWithHi(t *testing.T) {
	person_1 := entity.PersonEntity{
		Id:   1,
		Name: "Guntur Firmansyah",
		Age:  20,
	}

	person_1_repo := repository.CreateNewHumanEntity(person_1)
	person_1_repo.PrintWithHi()
}

func TestPrintWithSorry(t *testing.T) {
	person_1 := entity.PersonEntity{
		Id:   1,
		Name: "Guntur Firmansyah",
		Age:  20,
	}

	person_1_repo := repository.CreateNewHumanEntity(person_1)
	person_1_repo.PrintWithSorry()
}

func TestPrintWithGoodBye(t *testing.T) {
	person_1 := entity.PersonEntity{
		Id:   1,
		Name: "Guntur Firmansyah",
		Age:  20,
	}

	person_1_repo := repository.CreateNewHumanEntity(person_1)
	person_1_repo.PrintWithGoodBye()
}
