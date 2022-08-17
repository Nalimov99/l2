package main

import "fmt"

/*
Паттерн: Strategy

Суть паттерна: вынести набор алгоритом в отдельные структуры и сделать их взаимозаменямыми
Плюсы: Инкапсуляция методов
Минусы: Для каждого метода придеться создавать свой собсвтенный тип данных
*/

type Executor interface {
	Execute(int, int) int
}

type Operation struct {
	Executor
}

type Add struct{}

func (a *Add) Execute(x, y int) int {
	return x + y
}

type Multi struct{}

func (m *Multi) Execute(x, y int) int {
	return x * y
}

func main() {
	fmt.Println("Adding:", Operation{Executor: &Add{}}.Execute(1, 2))
	fmt.Println("Multi:", Operation{Executor: &Multi{}}.Execute(2, 2))
}
