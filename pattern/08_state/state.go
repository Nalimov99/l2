package main

import "fmt"

/*
Паттерн: State

Суть паттерна: изменять поведение структуры во время выполнение программы взависиомсти от состояние структуры.
Плюсы: Условные операторы заменяются полиморфизмом
Минусы: Сложно реализовывать методы не релеватные для некоторых состояних.
*/

type State interface {
	EmptyCart()
	Items()
	Money()
}

type Cart struct {
	CartEmpty     State
	CartWithItem  State
	CartWithMoney State

	currentState State
	item         bool
	money        bool
}

func NewCart() *Cart {
	cart := &Cart{}
	emptyCart := &EmtyCartState{cart}
	items := &ItemsState{cart}
	money := &MoneyState{cart}

	cart.currentState = emptyCart
	cart.CartWithItem = items
	cart.CartWithMoney = money
	cart.CartEmpty = emptyCart

	return cart
}

func (c *Cart) EmptyCart() {
	c.currentState.EmptyCart()
}

func (c *Cart) Items() {
	c.currentState.Items()
}

func (c *Cart) Money() {
	c.currentState.Money()
}

// Стейт пустой корзины
type EmtyCartState struct {
	Cart *Cart
}

func (e *EmtyCartState) EmptyCart() {
	fmt.Println("Корзина итак пуста")
}

func (e *EmtyCartState) Items() {
	e.Cart.item = true
	e.Cart.currentState = e.Cart.CartWithItem
	fmt.Println("Добавили товар")
}

func (e *EmtyCartState) Money() {
	fmt.Println("Корзина пуста, деньги внести нельзя")
}

// Стейт корзины с вещами
type ItemsState struct {
	Cart *Cart
}

func (i *ItemsState) Items() {
	fmt.Println("В корзине уже есть вещи")
}

func (i *ItemsState) EmptyCart() {
	i.Cart.item = false
	i.Cart.currentState = i.Cart.CartEmpty
	fmt.Println("Корзина отчищена")
}

func (i *ItemsState) Money() {
	i.Cart.money = true
	i.Cart.currentState = i.Cart.CartWithMoney
	fmt.Println("Покупка совершена")
}

// Стейт корзины с деньгами
type MoneyState struct {
	Cart *Cart
}

func (m *MoneyState) EmptyCart() {
	m.Cart.item = false
	m.Cart.money = false
	m.Cart.currentState = m.Cart.CartEmpty
	fmt.Println("Корзина успешно отчищена после покупки")
}

func (m *MoneyState) Items() {
	fmt.Println("Добавить новую вещь можно после отчистки корзины")
}

func (m *MoneyState) Money() {
	fmt.Println("Деньги можно внести только после того как добавили вещь")
}

func main() {
	cart := NewCart()

	fmt.Println("Стейт пустой корзины:")
	cart.EmptyCart()
	cart.Money()
	cart.Items()

	fmt.Println("\nСтейт корзины с вещами:")
	cart.Items()
	cart.EmptyCart()
	cart.Items()
	cart.Money()

	fmt.Println("\nСтейт оплаченной корзины:")
	cart.Items()
	cart.Money()
	cart.EmptyCart()

	fmt.Println("\nНовая покупа:")
	cart.Items()
	cart.Money()
	cart.EmptyCart()
}
