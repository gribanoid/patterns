package types

import (
	"patterns/bridge/pkg/base"
)

type LinuxPC struct {
	scanner base.Scanner
}

func (pc *LinuxPC) Scan() {
	print("scan pc linux system ")
	pc.scanner.ScanFile()
}
func (pc *LinuxPC) SetScanner(scanner base.Scanner) base.PC {
	pc.scanner = scanner
	return pc
}
