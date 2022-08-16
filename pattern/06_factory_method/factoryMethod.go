package main

import "fmt"

/*
Паттерн: Factory method
Суть паттерна: определить интерфейс для создание экземпляров некоторой
структуры, но непосредственное решение о том, экземпляром какой структуры создавать
делегировать на данную структуру
*/

type CarModel string

const (
	AudiModel CarModel = "Audi"
	BMWModel  CarModel = "BMW"
)

type Driver interface {
	Drive()
}

type Car struct {
	name  string
	speed int
}

func (c *Car) Drive() {
	fmt.Printf("%s driving at speed %d\n", c.name, c.speed)
}

func GetCar(model CarModel) Driver {
	switch model {
	case AudiModel:
		return &Car{
			name:  "Audi",
			speed: 50,
		}
	case BMWModel:
		return &Car{
			name:  "BMW",
			speed: 50,
		}
	default:
		return nil
	}
}

type Audi struct {
	Car
}

type BMW struct {
	Car
}

func main() {
	bmw := GetCar(BMWModel)
	audi := GetCar(AudiModel)

	bmw.Drive()
	audi.Drive()
}
