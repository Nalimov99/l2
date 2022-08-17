package main

import "fmt"

/*
Паттерн: Visitor

Суть паттерна: открепить алгоритм от структуры того объекта, которым он оперирует.
Практический результат такого открепления – способность добавлять новые операции к имеющимся структурам объектов,
не модифицируя эти структуры.
Плюсы: Упрощает добавление функционала, не расширя структуры
*/

// ================
// Структуры для которых будет определен visitor
type Employee interface {
	GetFullName() string
	Accept(Visitor)
}

type Developer struct {
	FirstName, LastName string
	Income              int
}

func (d *Developer) GetFullName() string {
	return d.FirstName + " " + d.LastName
}

func (d *Developer) Accept(visitor Visitor) {
	visitor.VisitDeveloper(d)
}

type PM struct {
	FirstName, LastName string
	Income              int
}

func (p *PM) GetFullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *PM) Accept(visitor Visitor) {
	visitor.VisitPM(p)
}

// ================
// Visitor
type Visitor interface {
	VisitDeveloper(*Developer)
	VisitPM(*PM)
}

type CalculateIncome struct {
	BonusRate int
}

func (c *CalculateIncome) VisitDeveloper(d *Developer) {
	fmt.Printf("Developer: %s; Income: %d\n", d.GetFullName(), d.Income*c.BonusRate)
}

func (c *CalculateIncome) VisitPM(p *PM) {
	fmt.Printf("PM: %s; Income: %d\n", p.GetFullName(), p.Income*c.BonusRate)
}

type Position struct{}

func (p Position) VisitDeveloper(d *Developer) {
	fmt.Println("Developer")
}

func (p Position) VisitPM(pm *PM) {
	fmt.Println("PM")
}

func main() {
	developer := Developer{
		FirstName: "Ilia",
		LastName:  "Nalimov",
		Income:    10,
	}

	pm := PM{
		FirstName: "Ivan",
		LastName:  "Ivanov",
		Income:    12,
	}

	developer.Accept(&CalculateIncome{10})
	pm.Accept(&CalculateIncome{8})

	developer.Accept(Position{})
	pm.Accept(Position{})
}
