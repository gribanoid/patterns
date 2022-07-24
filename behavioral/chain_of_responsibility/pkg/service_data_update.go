package pkg

import "fmt"

type UpdateDataService struct {
	Name string
	Next Service
}

func (upd *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Printf("Data from device [%s] already update.\n", upd.Name)
		upd.Next.Execute(data)
		return
	}
	fmt.Printf("Update data from device [%s].\n", upd.Name)
	data.UpdateSource = true
	upd.Next.Execute(data)
}

func (upd *UpdateDataService) SetNext(service Service) {
	upd.Next = service
}
