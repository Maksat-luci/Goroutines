package main

import (
	"fmt"
	"time"
)

// Paypal структура
type Paypal struct {
	money int64
}

//TakeMoney Метод для снятия денег с PayPal
func (p *Paypal) TakeMoney(user *User) {
	fmt.Println("user", user.Name, "get money", user.takeMoney)
	p.SetMoney(-user.takeMoney)
}

//SetMoney Метод пополнения счета Paypal
func (p *Paypal) SetMoney(money int64) {
	p.money += money
}

//GetMoney Метод получения счета Paypal
func (p *Paypal) GetMoney() int64 {
	return p.money
}

// Конструктор
func newPaypal() *Paypal {
	return &Paypal{
		money: 100,
	}
}
// User абстрактная структура для реализации этого таска
type User struct {
	Name      string
	takeMoney int64 //сколько денег будет брать user
}

func humanParalel(user *User, p *Paypal) {
	for {
		p.TakeMoney(user)
		time.Sleep(time.Millisecond * 300)
	}
}

func main() {
	paypal := newPaypal()
	users := []*User{{"Denis", 20}, {"Richard", 10}}
	stop := make(chan struct{})

	go humanParalel(users[0], paypal)
	go humanParalel(users[1], paypal)
	go func() {
		for {
			// Раз в 200 секунд мы проверяем закончился ли счет PayPal
			if paypal.GetMoney() <= 0 {
				stop <- struct{}{}
			}
			time.Sleep(time.Millisecond * 200)
		}
	}()
	<-stop
	time.Sleep(time.Second * 2)
	fmt.Println("\n\nPaypal money", paypal.GetMoney())

}
