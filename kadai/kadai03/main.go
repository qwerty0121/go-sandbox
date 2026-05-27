// 課題3: 構造体とポインタメソッド

package main

import "fmt"

// Account は、銀行口座を表す構造体です。
type Account struct {
	// 残高
	balance int
}

// NewAccount は、Accountのコンストラクタです。
// 引数 initialBalance には口座の初期残高を指定します。
func NewAccount(initialBalance int) *Account {
	if initialBalance < 0 {
		// 初期残高は常に0以上となるようにする
		initialBalance = 0
	}

	return &Account{balance: initialBalance}
}

// Deposit は、指定された金額を口座に入金するメソッドです。
// 引数 amount には入金する金額を指定します。
// amount が正でない場合は何もしません。
func (a *Account) Deposit(amount int) {
	if amount <= 0 {
		return
	}

	a.balance += amount
}

// Withdraw は、指定された金額を口座から出金するメソッドです。
// 引数 amount には出金する金額を指定します。
// 正常に出金できた場合は true を返します。
// amount が正でない、または残高不足である場合は false を返します。
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

// Balance は、口座の残高を返します。
func (a *Account) Balance() int {
	return a.balance
}

func main() {
	// ケース1
	account1 := NewAccount(100)
	account1.Deposit(50)
	fmt.Printf("account1's balance = %v\n", account1.Balance())

	// ケース2
	account2 := NewAccount(150)
	result2 := account2.Withdraw(50)
	fmt.Printf("account2 withdraw result = %v\n", result2)
	fmt.Printf("account2's balance = %v\n", account2.Balance())

	// ケース3
	account3 := NewAccount(100)
	result3 := account3.Withdraw(150)
	fmt.Printf("account3 withdraw result = %v\n", result3)
	fmt.Printf("account3's balance = %v\n", account3.Balance())
}
