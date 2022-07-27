package pkg

import "fmt"

type ProxyDatabase struct {
	Users map[string]bool
	DB    *Database
}

func (pDB *ProxyDatabase) GetData(user string) ([]string, error) {
	v, ok := pDB.Users[user]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	if !v {
		return nil, fmt.Errorf("access is denied")
	}
	return pDB.DB.GetData(user)
}
