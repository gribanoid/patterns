package main

import (
	"patterns/behavioral/mediator/pkg"
	"time"
)

func main() {
	stationManger := pkg.NewStationManager()
	passangerBus := &pkg.Passenger{
		Dispatcher: stationManger,
	}
	cargo := &pkg.Cargo{
		Dispatcher: stationManger,
	}
	passangerBus.Arrive()
	time.Sleep(time.Second)
	cargo.Arrive()
	time.Sleep(time.Second)
	passangerBus.Go()

}
