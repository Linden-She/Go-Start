/**
 @author: lin.she
 @date: 2021/6/22
 @note:在 Go 中，如果一个符号（例如变量、类型、函数等）是以小写符号开头，那么它在 定义它的包之外 就是私有的。
**/
package example

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()
	fmt.Println("address of balance in test is:", &wallet.balance)
	want := 10
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
}
