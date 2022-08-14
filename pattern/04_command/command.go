package main

import "fmt"

/*
Паттерн: Command

Суть паттерна: Использование единного интерфейса для описание всех типов операций,
которые можно производить с системой.
*/

type Button struct {
	Command Command
}

func (b *Button) Press() {
	b.Command.Execute()
}

type Command interface {
	Execute()
}

type OnCommand struct {
	Device Device
}

func (c *OnCommand) Execute() {
	c.Device.on()
}

type OffComand struct {
	Device Device
}

func (c *OffComand) Execute() {
	c.Device.off()
}

type Device interface {
	on()
	off()
}

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("tv off")
}

func main() {
	tv := Tv{}
	onCommand := OnCommand{
		Device: &tv,
	}
	offComand := OffComand{
		Device: &tv,
	}

	onCommand.Execute()
	offComand.Execute()

	onButton := Button{
		Command: &onCommand,
	}
	offButton := Button{
		Command: &offComand,
	}

	onButton.Press()
	offButton.Press()
}
