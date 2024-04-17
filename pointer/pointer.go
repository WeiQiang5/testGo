package pointer

import (
	"errors"
	"fmt"
)

type Bitcoin int

// 这里添加个具体的名称类型，使balance具有描述性
type Wallet struct {
	balance Bitcoin
}

var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

// 这里需要注意如果不加*符号的话，是调用的副本
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}
func (w *Wallet) WithDraw(amount Bitcoin) error {
	fmt.Printf("amount => %d\n", amount)
	fmt.Printf("amount => %d\n", w.balance)
	if amount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// %s会输出String函数
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
