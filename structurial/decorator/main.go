package main

import (
	"fmt"
	"patterns/structurial/decorator/pkg"
)

var (
	base   = pkg.BasePC{}
	home   = pkg.HomePC{CPU: 4, Wrapper: base}
	server = pkg.ServerPC{CPU: 16, Memory: 256, Wrapper: base}
)

func main() {
	fmt.Printf("Base price [%.2f]\n", base.GetPrice())
	fmt.Printf("Home price [%.2f]\n", home.GetPrice())
	fmt.Printf("Server price [%.2f]\n", server.GetPrice())
}
