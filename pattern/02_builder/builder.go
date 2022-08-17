package main

import "fmt"

/*
Паттерн: Builder

Суть паттерна: поэтапное создание сложного объекта с помощью четко определенной последовательности действий.

Плюсы: Пошаговое создание продукта.
Недостатки: Builder и создаваемый им продукт жестко связаны между собой,
поэтому при изменении продукта скорее всего придется соотвествующим образом изменять и Builder.
*/

type Color string

const (
	RED   Color = "red"
	BLACK Color = "black"
)

// ================
// Продукт билдера
type Car interface {
	Drive()
}

type car struct {
	speed int
	color Color
}

func (c *car) Drive() {
	fmt.Printf("%s car driving at speed %d\n", c.color, c.speed)
}

// ================
// Билдер
type CarBuilder interface {
	SetSpeed(int) CarBuilder
	Paint(Color) CarBuilder
	Build() Car
}

type carBuilder struct {
	speed int
	color Color
}

func NewCarBuilder() CarBuilder {
	return &carBuilder{}
}

func (cb *carBuilder) SetSpeed(speed int) CarBuilder {
	cb.speed = speed
	return cb
}

func (cb *carBuilder) Paint(color Color) CarBuilder {
	cb.color = color
	return cb
}

func (cb *carBuilder) Build() Car {
	return &car{
		speed: cb.speed,
		color: cb.color,
	}
}

func main() {
	NewCarBuilder().
		SetSpeed(100).
		Paint(BLACK).
		Build().
		Drive()
}
