package pkg

import "fmt"

type Passenger struct {
	Dispatcher
}

func (g *Passenger) PermitArrive() {
	fmt.Println("Пассажиры, займите метса...")
	g.Arrive()
}

func (g *Passenger) Go() {
	fmt.Println("Пассажиры, отправление!")
	g.Dispatcher.NotifyAboutGo()
}

func (g *Passenger) Arrive() {
	if !g.CanArrive(g) {
		fmt.Println("Пассажиры, отправление задерживается, платформа занята...")
		return
	}
	fmt.Println("Пассажиры, займите метса...")
}
