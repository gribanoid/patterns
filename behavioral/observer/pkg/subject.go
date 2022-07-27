package pkg

type Subject interface {
	Subscribe(Consumer)
	UnSubscribe(Consumer)
	Notify()
}
