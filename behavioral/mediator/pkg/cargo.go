package pkg

import "fmt"

type Cargo struct {
	Dispatcher
}

func (g *Cargo) PermitArrive() {
	fmt.Println("Погрузка грузовика...")
	g.Arrive()
}

func (g *Cargo) Go() {
	fmt.Println("Грузовик успешно погружен!")
	g.Dispatcher.NotifyAboutGo()
}

func (g *Cargo) Arrive() {
	if !g.CanArrive(g) {
		fmt.Println("Погрузка грузовика запрещена, на платформе пассажиры!")
		return
	}
	fmt.Println("Грузовик отправлен")
}
