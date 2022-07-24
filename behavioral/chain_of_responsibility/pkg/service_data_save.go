package pkg

import "fmt"

type DataService struct {
	Next Service
}

func (ds *DataService) Execute(data *Data) {
	if !data.UpdateSource {
		fmt.Println("Data not update.")
		return
	}
	fmt.Println("Data save.")
}

func (ds *DataService) SetNext(service Service) {
	ds.Next = service
}
