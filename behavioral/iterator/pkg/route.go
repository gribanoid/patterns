package pkg

import "sync"

type Route struct {
	Name       string
	TravelTime int
}

type Iterator interface {
	HasNext() bool
	GetNext() *Route
}

type Routes struct {
	index  int
	Routes []Route
}

func (r *Routes) HasNext() bool {
	return r.index < len(r.Routes)
}

func (r *Routes) GetNext() *Route {
	mu := new(sync.Mutex)
	defer func() {
		mu.Lock()
		r.index++
		mu.Unlock()
	}()
	if r.HasNext() {
		return &r.Routes[r.index]
	}
	return nil
}
