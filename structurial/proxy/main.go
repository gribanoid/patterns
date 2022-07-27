package main

import (
	"fmt"
	"patterns/structurial/proxy/pkg"
)

var (
	admin = "admin"
	user  = "user"
	users = map[string]bool{
		admin: true,
		user:  false,
	}
)

func main() {
	proxy := pkg.ProxyDatabase{
		Users: users,
		DB:    &pkg.Database{},
	}
	adminData, err := proxy.GetData(admin)
	fmt.Printf("From [%s] Data:[%v] Err:[%v]\n", admin, adminData, err)
	userData, err := proxy.GetData(user)
	fmt.Printf("From [%s] Data:[%v] Err:[%v]\n", user, userData, err)
}
