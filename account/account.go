package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Account interface {
	Withdraw(uint)
	Deposit(uint)
	Balance() int
}
type Bank struct {
	account Account
}

func NewBank(account Account) *Bank {
	return &Bank{account: account}
}
func (bank *Bank) Withdraw(amount uint, actor_name string) {
	fmt.Println("[-]", amount, actor_name)
	bank.account.Withdraw(amount)
}
func (bank *Bank) Deposit(amount uint, actor_name string) {
	fmt.Println("[+]", amount, actor_name)
	bank.account.Deposit(amount)
}
func (bank *Bank) Balance() int {
	return bank.account.Balance()
}

type SimpleAccount struct {
	balance int
}

func NewSimpleAccount(balance int) *SimpleAccount {
	return &SimpleAccount{balance: balance}
}
func (acc *SimpleAccount) Deposit(amount uint) {
	acc.setBalance(acc.balance + int(amount))
}
func (acc *SimpleAccount) Withdraw(amount uint) {
	if acc.balance >= int(amount) {
		acc.setBalance(acc.balance - int(amount))
	} else {
		panic("杰克穷死")
	}
}
func (acc *SimpleAccount) Balance() int {
	return acc.balance
}
func (acc *SimpleAccount) setBalance(balance int) {
	acc.add_some_latency() //增加一个延时函数，方便演示
	acc.balance = balance
}
func (acc *SimpleAccount) add_some_latency() {
	<-time.After(time.Duration(rand.Intn(100)) * time.Millisecond)
}

type LockingAccount struct {
	lock    sync.Mutex
	account *SimpleAccount
}

//封装一下 SimpleAccount
func NewLockingAccount(balance int) *LockingAccount {
	return &LockingAccount{account: NewSimpleAccount(balance)}
}
func (acc *LockingAccount) Deposit(amount uint) {
	acc.lock.Lock()
	defer acc.lock.Unlock()
	acc.account.Deposit(amount)
}
func (acc *LockingAccount) Withdraw(amount uint) {
	acc.lock.Lock()
	defer acc.lock.Unlock()
	acc.account.Withdraw(amount)
}
func (acc *LockingAccount) Balance() int {
	acc.lock.Lock()
	defer acc.lock.Unlock()
	return acc.account.Balance()
}

type ConcurrentAccount struct {
	account     *SimpleAccount
	deposits    chan uint
	withdrawals chan uint
	balances    chan chan int
}

func NewConcurrentAccount(amount int) *ConcurrentAccount {
	acc := &ConcurrentAccount{
		account:     &SimpleAccount{balance: amount},
		deposits:    make(chan uint),
		withdrawals: make(chan uint),
		balances:    make(chan chan int),
	}
	acc.listen()

	return acc
}
func (acc *ConcurrentAccount) Balance() int {
	ch := make(chan int)
	acc.balances <- ch
	return <-ch
}
func (acc *ConcurrentAccount) Deposit(amount uint) {
	acc.deposits <- amount
}
func (acc *ConcurrentAccount) Withdraw(amount uint) {
	acc.withdrawals <- amount
}
func (acc *ConcurrentAccount) listen() {
	go func() {
		for {
			select {
			case amnt := <-acc.deposits:
				acc.account.Deposit(amnt)
			case amnt := <-acc.withdrawals:
				acc.account.Withdraw(amnt)
			case ch := <-acc.balances:
				ch <- acc.account.Balance()
			}
		}
	}()
}

func main() {
	balance := 80
	b := NewBank(NewConcurrentAccount(balance))

	fmt.Println("初始化余额", b.Balance())

	done := make(chan bool)

	go func() { b.Withdraw(30, "马伊琍"); done <- true }()
	go func() { b.Withdraw(10, "姚笛"); done <- true }()

	//等待 goroutine 执行完成
	<-done
	<-done

	fmt.Println("-----------------")
	fmt.Println("剩余余额", b.Balance())
}
