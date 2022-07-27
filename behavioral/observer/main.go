package main

import (
	"fmt"
	"patterns/behavioral/observer/pkg"
)

func main() {
	sub1 := &pkg.Subscriber{Name: "Sub1"}
	sub2 := &pkg.Subscriber{Name: "Sub2"}
	sub3 := &pkg.Subscriber{Name: "Sub3"}
	channel := pkg.Publisher{
		Name:      "Publisher channel",
		Consumers: map[string]pkg.Consumer{},
	}
	channel.Subscribe(sub1).Subscribe(sub2).Subscribe(sub3).Notify()
	fmt.Println("............")
	channel.UnSubscribe(sub1).UnSubscribe(sub2).Notify()
}
