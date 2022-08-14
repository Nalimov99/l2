package main

import (
	"errors"
	"fmt"
)

/*
Паттерн: Facade

Суть паттерна: предоставить унифицированный интерфейс вместо набора
интерфейсов некоторой подсистемы. Фасад определяет интерфейс более
высокого уровня, который упрощает использование подсистемы.
Благодаря такому подхому, отдельные компоненты системы могут быть разработаны
изолировано, затем интегрированы вмсете.

Недостатки: Высокая связность модулей системы.
*/

//========================
// подсистема Account
type account struct {
	name string
}

func newAccount(name string) *account {
	return &account{
		name,
	}
}

func (a *account) checkAccount(name string) error {
	if a.name != name {
		return errors.New("account Name is incorrect")
	}

	return nil
}

//========================
// подсистема securityCode
type securityCode struct {
	code int
}

func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code,
	}
}

func (s *securityCode) checkSecurityCode(code int) error {
	if s.code != code {
		return errors.New("security Code is incorrect")
	}

	return nil
}

//========================
// подсистема wallet
type wallet struct {
	balance int
}

func newWallet(balance int) *wallet {
	return &wallet{
		balance,
	}
}

func (w *wallet) addToBalance(amount int) {
	w.balance += amount
}

func (w *wallet) writeOff(amount int) error {
	if amount > w.balance {
		return errors.New("balance is not sufficient")
	}

	w.balance -= amount
	return nil
}

//========================
// Фасад
type WalletFacade struct {
	account      *account
	securityCode *securityCode
	wallet       *wallet
}

func NewWalletFacade(accountName string, securityCode int, initialBalance int) *WalletFacade {
	return &WalletFacade{
		account:      newAccount(accountName),
		securityCode: newSecurityCode(securityCode),
		wallet:       newWallet(initialBalance),
	}
}

func (wf *WalletFacade) checkValidity(name string, secusecurityCode int) error {
	if err := wf.account.checkAccount(name); err != nil {
		return err
	}

	if err := wf.securityCode.checkSecurityCode(secusecurityCode); err != nil {
		return err
	}

	return nil
}

func (wf *WalletFacade) AddToBalance(name string, securityCode int, amount int) {
	if err := wf.checkValidity(name, securityCode); err != nil {
		fmt.Println(err)
		return
	}

	wf.wallet.addToBalance(amount)
	fmt.Printf("Added: %d; Account: %s; Current balance: %d\n", amount, name, wf.wallet.balance)
}

func (wf *WalletFacade) WriteOff(name string, securityCode int, amount int) {
	if err := wf.checkValidity(name, securityCode); err != nil {
		fmt.Println(err)
		return
	}

	if err := wf.wallet.writeOff(amount); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Write off: %d; Account: %s; Current balance: %d\n", amount, name, wf.wallet.balance)
}

func main() {
	wf := NewWalletFacade("Ilia", 123, 0)

	wf.AddToBalance("Ilia", 123, 100)
	wf.WriteOff("Ilia", 123, 50)

	wf.WriteOff("Ilia", 123, 100)
	wf.WriteOff("Ilia", 12, 50)
	wf.WriteOff("I", 123, 10)
}
