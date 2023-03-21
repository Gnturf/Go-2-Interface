package main

import (
	"fmt"
	"strconv"
)

type Info interface {
	getInfo() string
}

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

type Animal struct {
	AnimalType string
	Age        int
}

func (h Human) getInfo() string {
	return h.FirstName + " " + h.LastName + " at Age " + strconv.Itoa(h.Age) + " y.o"
}

func (a Animal) getInfo() string {
	return a.AnimalType + " at Age " + strconv.Itoa(a.Age) + " y.o"
}

func PrintInfo(data Info) {
	fmt.Println(data.getInfo())
}

func main() {
	person1 := Human{FirstName: "Joko", LastName: "Widodo", Age: 60}
	animal1 := Animal{AnimalType: "Panda", Age: 4}

	PrintInfo(person1)
	PrintInfo(animal1)
}
