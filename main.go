package main

import (
	_0_SRP "fe-2025-1/00_SRP"
	_1_OCP "fe-2025-1/01_OCP"
	_2_LSP "fe-2025-1/02_LSP"
	_3_ISP "fe-2025-1/03_ISP"
	_4_DIP "fe-2025-1/04_DIP"
)

func main() {
	_0_SRP.SrpI()
	_1_OCP.OcpI()
	_2_LSP.LspI()
	_3_ISP.IspI()
	_4_DIP.DipI()

	_0_SRP.SrpC()
	_1_OCP.OcpC()
	_2_LSP.LspC()
	_3_ISP.IspC()
	_4_DIP.DipC()
}
