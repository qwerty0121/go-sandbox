// 課題3: 構造体とポインタメソッド

package main

import "fmt"

// Account は、銀行口座を表す構造体です。
type Account struct {
	// 残高
	balance int
}

// Deposit は、指定された金額を口座に入金するメソッドです。
func (a *Account) Deposit(amount int) {
	if amount <= 0 {
		return
	}

	a.balance += amount
}

// Withdraw は、指定された金額を口座から出金するメソッドです。
func (a *Account) Withdraw(amount int) bool {
	if amount <= 0 {
		return false
	}

	if a.balance < amount {
		// 残高不足の場合
		return false
	}

	a.balance -= amount
	return true
}

func main() {
	// ケース1
	account1 := Account{balance: 100}
	account1.Deposit(50)
	fmt.Printf("account1's balance = %v\n", account1.balance)

	// ケース2
	account2 := Account{balance: 150}
	result2 := account2.Withdraw(50)
	fmt.Printf("account2 withdraw result = %v\n", result2)
	fmt.Printf("account2's balance = %v\n", account2.balance)

	// ケース3
	account3 := Account{balance: 100}
	result3 := account3.Withdraw(150)
	fmt.Printf("account3 withdraw result = %v\n", result3)
	fmt.Printf("account3's balance = %v\n", account3.balance)
}
