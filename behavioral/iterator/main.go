package main

import (
	"fmt"
	"patterns/behavioral/iterator/pkg"
)

var routes = pkg.Routes{
	Routes: []pkg.Route{
		{Name: "Маршрут-1", TravelTime: 110},
		{Name: "Маршрут-2", TravelTime: 50},
		{Name: "Маршрут-3", TravelTime: 60},
		{Name: "Маршрут-4", TravelTime: 40},
	},
}

func main() {
	for routes.HasNext() {
		route := routes.GetNext()
		fmt.Printf("R:[%s] Time:[%d]\n", route.Name, route.TravelTime)
	}
}
