// Реализовать паттерн «адаптер» на любом примере.

package main

import (
	"log"
)

// PaymentSystem структура платежной системы
type PaymentSystem struct{}

// Pay метод оплаты, который мы будем потом адаптировать под наши нужды
func (p *PaymentSystem) Pay(amount float64) error {
	// Какой-то процесс ...
	return nil
}

// PaymenterSystem интерфейс платежной системы
type PaymenterSystem interface {
	Pay(amount float64) error
}

// PaymentAdapter структура адаптера, которая содержит
// интерфейс PaymenterSystem, чтобы можно было использовать PaymentSystem
type PaymentAdapter struct {
	paymenterSystem PaymenterSystem
}

// ProcessPayment наш метод в котором будем использовать метод Pay
func (p *PaymentAdapter) ProcessPayment(amount float64, cardNumber, cardExpirationMonth, cardExpirationYear string) error {
	if err := p.paymenterSystem.Pay(amount); err != nil { // Успешно используем paymenterSystem
		return err
	}
	// Какой-то процесс ...
	return nil
}

func main() {
	paymentSystem := &PaymentSystem{}                                 // Инициализация PaymentSystem
	paymentAdapter := &PaymentAdapter{paymenterSystem: paymentSystem} // Внедрение его в paymentAdapter
	// Используем свой метод
	err := paymentAdapter.ProcessPayment(100, "1234 5678 9012 3456", "12", "2023")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Оплата завершена успешно")
}
