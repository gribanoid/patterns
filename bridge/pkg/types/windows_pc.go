package types

import (
	"patterns/bridge/pkg/base"
)

type WindowsPC struct {
	scanner base.Scanner
}

func (pc *WindowsPC) Scan() {
	print("scan pc windows system ")
	pc.scanner.ScanFile()
}
func (pc *WindowsPC) SetScanner(scanner base.Scanner) base.PC {
	pc.scanner = scanner
	return pc
}
