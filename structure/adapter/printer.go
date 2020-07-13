package adapter

import (
	"fmt"
)

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct{}

func (m *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s", s)
	println(newMsg)
	return
}

//---------------------------------------------

type NewPrinter interface {
	PrintStored() string
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter == nil {
		newMsg = p.Msg
	} else {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	}
	return
}
