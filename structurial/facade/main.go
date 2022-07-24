package main

import "patterns/structurial/facade/pkg"

var (
	bank = pkg.Bank{
		Name:  "БАНК",
		Cards: []pkg.Card{},
	}
	card1 = pkg.Card{
		Name:    "card-1",
		Balance: 200,
		Bank:    &bank,
	}
	card2 = pkg.Card{
		Name:    "card-2",
		Balance: 5,
		Bank:    &bank,
	}
	user1 = pkg.User{
		Name: "Покупатель-1",
		Card: &card1,
	}
	user2 = pkg.User{
		Name: "Покупатель-2",
		Card: &card2,
	}
	prod = pkg.Product{
		Name:  "Сыр",
		Price: 150,
	}
	shop = pkg.Shop{
		Name: "Highpixel",
		Products: []pkg.Product{
			prod,
		},
	}
)

func main() {
	println("[БАНК] Выпуск карт")
	bank.Cards = append(bank.Cards, card1, card2)
	println(user1.Name)
	err := shop.Sell(user1, prod.Name)
	if err != nil {
		println(err)
		return
	}
	println(user2.Name)
	err = shop.Sell(user2, prod.Name)
	if err != nil {
		println(err.Error())
		return
	}
}
