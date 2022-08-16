package main

import "fmt"

/*
Паттерн: Chain of Responsibility

Суть паттерна: Эффективно и компактно реализовать механизм обработки потока
событий/запросов/сообщений в системах с потенциально большим количеством обработчиков.
Недостатки: Возможность появление сложносоставных цепей.
*/

// Department знает какие методы должен реализовывать
// каждый элемент цепи
type Department interface {
	Execute(*Patient)
	SetNext(Department)
}

type Reception struct {
	next Department
}

func (r *Reception) Execute(p *Patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.Execute(p)
		return
	}

	fmt.Println("Reception registering patient:", p.Name)
	p.registrationDone = true
	r.next.Execute(p)
}

func (r *Reception) SetNext(next Department) {
	r.next = next
}

type Doctor struct {
	next Department
}

func (d *Doctor) Execute(p *Patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.Execute(p)
		return
	}
	fmt.Println("Doctor prescribes medicine to a patient")
	p.doctorCheckUpDone = true
	d.next.Execute(p)
}

func (d *Doctor) SetNext(next Department) {
	d.next = next
}

type Medical struct {
	next Department
}

func (m *Medical) Execute(p *Patient) {
	if p.medicineDone {
		fmt.Println("Medicine already given to patient")
		m.next.Execute(p)
		return
	}
	fmt.Println("Medical giving medicine to patient")
	p.medicineDone = true
	m.next.Execute(p)
}

func (m *Medical) SetNext(next Department) {
	m.next = next
}

type Cashier struct {
	next Department
}

func (c *Cashier) Execute(p *Patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Println("Cashier getting money from patient patient")
}

func (c *Cashier) SetNext(next Department) {
	c.next = next
}

// Patient структура которая запускает цепочку событий
// Приватные поля структуры описывают на каких стадиях находиться событие
type Patient struct {
	Name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func main() {
	cashier := &Cashier{}

	medical := &Medical{}
	medical.SetNext(cashier)

	doctor := &Doctor{}
	doctor.SetNext(medical)

	reception := &Reception{}
	reception.SetNext(doctor)
	patient := &Patient{Name: "Ilia Nalimov"}

	reception.Execute(patient)
}
