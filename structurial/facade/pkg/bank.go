package pkg

import (
	"errors"
	"time"
)

type Bank struct {
	Name  string
	Cards []Card
}

func (bank *Bank) CheckBalance(cardNumber string) error {
	println("[Банк] Получение по карте", cardNumber)
	time.Sleep(time.Millisecond * 300)
	for _, card := range bank.Cards {
		if card.Name != cardNumber {
			continue
		}
		if card.Balance <= 0 {
			return errors.New("[Банк] Недостаточно средств")
		}
	}
	println("[Банк] Остаток больше нуля")
	return nil
}
