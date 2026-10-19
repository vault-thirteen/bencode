package bencode

import (
	"github.com/vault-thirteen/auxie/cio"
)

func PrintTestNumber(idx int) {
	cio.PrintMessage(cio.FormatMessage("Test #%v.", idx+1))
}
