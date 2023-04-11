package test

import (
	"testing"

	"github.com/NaylaDeLis/Go-2-Interface/entity"
	"github.com/NaylaDeLis/Go-2-Interface/repository"
)

func TestPrintWithHiDemiHuman(t *testing.T) {
	demiHuman_1 := entity.DemiHumanEntity{
		Id:   1,
		Name: "Lystia",
		Age:  100,
		Race: "Fairies",
	}

	demiHuman_1_repo := repository.CreateNewDemiHumanEntity(demiHuman_1)
	demiHuman_1_repo.PrintWithHi()
}

func TestPrintWithSorryDemiHuman(t *testing.T) {
	demiHuman_1 := entity.DemiHumanEntity{
		Id:   1,
		Name: "Lystia",
		Age:  100,
		Race: "Fairies",
	}

	demiHuman_1_repo := repository.CreateNewDemiHumanEntity(demiHuman_1)
	demiHuman_1_repo.PrintWithSorry()
}

func TestPrintWithGoodByeDemiHuman(t *testing.T) {
	demiHuman_1 := entity.DemiHumanEntity{
		Id:   1,
		Name: "Lystia",
		Age:  100,
		Race: "Fairies",
	}

	demiHuman_1_repo := repository.CreateNewDemiHumanEntity(demiHuman_1)
	demiHuman_1_repo.PrintWithGoodBye()
}
