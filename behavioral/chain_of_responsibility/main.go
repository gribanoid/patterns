package main

import (
	"patterns/behavioral/chain_of_responsibility/pkg"
)

func main() {
	device := &pkg.Device{Name: "Device-1"}
	updateSVC := &pkg.UpdateDataService{Name: "Update-1"}
	dataSVC := &pkg.DataService{}
	device.SetNext(updateSVC)
	updateSVC.SetNext(dataSVC)
	data := &pkg.Data{}
	device.Execute(data)
}
