// Package day7 — указатели: передача по значению и по указателю.
package day7

import "errors"

var (
	ErrInvalidAmount     = errors.New("сумма должна быть положительной")
	ErrInsufficientFunds = errors.New("недостаточно средств")
)

// Swap меняет местами значения, на которые указывают a и b.
func Swap(a, b *int) {
	// TODO: ваше решение
}

// User — пользователь.
type User struct {
	Name string
	Age  int
}

// ResetAgeCopy получает копию структуры: изменения не видны снаружи.
// Реализовывать не нужно, это пример для сравнения с ResetAge.
func ResetAgeCopy(u User) {
	u.Age = 0
}

// ResetAge обнуляет возраст через указатель. При nil ничего не делает.
func ResetAge(u *User) {
	// TODO: ваше решение
}

// Account — банковский счёт, баланс в копейках.
type Account struct {
	Owner   string
	Balance int64
}

// NewAccount создаёт счёт с нулевым балансом и возвращает указатель на него.
func NewAccount(owner string) *Account {
	// TODO: ваше решение
	return nil
}

// Deposit пополняет счёт. amount <= 0 — ErrInvalidAmount.
func (a *Account) Deposit(amount int64) error {
	// TODO: ваше решение
	return nil
}

// Withdraw снимает деньги. amount <= 0 — ErrInvalidAmount, больше баланса — ErrInsufficientFunds.
func (a *Account) Withdraw(amount int64) error {
	// TODO: ваше решение
	return nil
}

// Transfer переводит amount со счёта from на счёт to.
// При ошибке ни один баланс не меняется.
func Transfer(from, to *Account, amount int64) error {
	// TODO: ваше решение
	return nil
}
