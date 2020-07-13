package bridge

import "fmt"

type computer interface {
	print() string
	setPrinter(printer)
}

type printer interface {
	printFile(msg string) string
}

//--------------------

type mac struct {
	printer printer
}

func (m *mac) print() string {
	return m.printer.printFile("Print request for Mac")
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}

//------------------------

type windows struct {
	printer printer
}

func (w *windows) print() string {
	return w.printer.printFile("Print request for Windows")
}

func (w *windows) setPrinter(p printer) {
	w.printer = p
}

//---------------------------

type epson struct {
}

func (p *epson) printFile(msg string) string {
	return fmt.Sprintf("%s / Printing by a EPSON Printer", msg)
}

//--------------------------

type hp struct {
}

func (p *hp) printFile(msg string) string {
	return fmt.Sprintf("%s / Printing by a HP Printer", msg)
}
