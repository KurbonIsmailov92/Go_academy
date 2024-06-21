package main

import "fmt"

/*
Задача 3
Реализация "Банковский счет" с методами Deposit, Withdraw и GetBalance:
Определите структуру BankAccount с полями Balance и Owner.
Реализуйте методы:
Deposit(amount): увеличивает баланс счета на заданную сумму.
Withdraw(amount) error: уменьшает баланс счета на заданную сумму, возвращая ошибку, если недостаточно средств.
GetBalance() float64: возвращает текущий баланс счета.
Протестируйте работу банковского счета с помощью этих методов.
*/

func main() {
	var acc BankAccount
	acc.BankAccountCreator()
	acc.Deposit()
	fmt.Println("Баланс:", acc.GetBalance())
	acc.Withdraw()
	fmt.Println("Баланс:", acc.GetBalance())
	acc.Deposit()
	fmt.Println("Баланс:", acc.GetBalance())
	acc.Withdraw()
	fmt.Println("Баланс:", acc.GetBalance())
}

type BankAccount struct {
	Balance float64
	Owner   string
}

func (b *BankAccount) BankAccountCreator() {
	fmt.Println("Введите имя собственника счета")
	fmt.Scan(&b.Owner)
	fmt.Println("Счет создан!")
}

func (b *BankAccount) Deposit() {
	var amount float64

	fmt.Println("Введите сумму прихода:")
	fmt.Scan(&amount)

	b.Balance += amount
}

func (b *BankAccount) Withdraw() {
	var amount float64

	fmt.Println("Введите сумму расхода:")
	fmt.Scan(&amount)

	if b.Balance < amount {
		fmt.Printf("У вас недостаточно средств на счету. Баланс равен: %.2f Сумма расхода: %.2f\n", b.Balance, amount)
		return
	}

	b.Balance -= amount
}

func (b *BankAccount) GetBalance() (balance float64) {
	return b.Balance
}
