package main

import "patterns/bridge/pkg/types"

var (
	/*scanners*/
	hpScanner    = types.HP{}
	epsonScanner = types.Epson{}
	/*computers*/
	linuxPC   = types.LinuxPC{}
	windowsPC = types.WindowsPC{}
	macPC     = types.MacPC{}
)

func main() {
	linuxPC.SetScanner(hpScanner).Scan()
	linuxPC.SetScanner(epsonScanner).Scan()
	windowsPC.SetScanner(hpScanner).Scan()
	macPC.SetScanner(epsonScanner).Scan()
}
