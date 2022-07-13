package base

type PC interface {
	Scan()
	SetScanner(scanner Scanner) PC
}
