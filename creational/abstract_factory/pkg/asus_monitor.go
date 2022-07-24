package pkg

import "fmt"

type AsusMonitor struct {
	Size int
}

func (pc *AsusMonitor) PrintDetails() {
	fmt.Printf("asus monitor size:%d\n", pc.Size)
}
