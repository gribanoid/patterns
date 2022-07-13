package types

import "patterns/bridge/pkg/base"

type MacPC struct {
	scanner base.Scanner
}

func (pc *MacPC) Scan() {
	print("scan pc mac system ")
	pc.scanner.ScanFile()
}

func (pc *MacPC) SetScanner(scanner base.Scanner) base.PC {
	pc.scanner = scanner
	return pc
}
