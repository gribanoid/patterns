package pkg

import (
	"errors"
	"time"
)

type Shop struct {
	Name     string
	Products []Product
}

func (shop Shop) Sell(user User, product string) error {
	println("[Магазин] Запрос к пользователю, для получения остатка по карте")
	time.Sleep(time.Millisecond * 500)
	err := user.Card.CheckBalance()
	if err != nil {
		return err
	}
	time.Sleep(time.Millisecond * 500)
	for _, prod := range shop.Products {
		if prod.Name != product {
			continue
		}
		if prod.Price > user.GetBalance() {
			return errors.New("[Магазин] у пользователя недостаточно средств")
		}
		println("товар", product, "куплен")
	}
	return nil
}
