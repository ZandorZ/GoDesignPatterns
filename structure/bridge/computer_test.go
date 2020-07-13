package bridge

import "testing"

func TestMacEpson(t *testing.T) {

	epsonPrinter := &epson{}
	hpPrinter := &hp{}
	macComputer := &mac{}
	winComputer := &windows{}

	macComputer.setPrinter(epsonPrinter)

	msg := macComputer.print()
	expected := "Print request for Mac / Printing by a EPSON Printer"
	if msg != expected {
		t.Errorf("Invalid printing.\nExpected: %s, \nGot: %s ", expected, msg)
	}

	macComputer.setPrinter(hpPrinter)
	msg = macComputer.print()
	expected = "Print request for Mac / Printing by a HP Printer"
	if msg != expected {
		t.Errorf("Invalid printing.\nExpected: %s, \nGot: %s ", expected, msg)
	}

	winComputer.setPrinter(epsonPrinter)
	msg = winComputer.print()
	expected = "Print request for Windows / Printing by a EPSON Printer"
	if msg != expected {
		t.Errorf("Invalid printing.\nExpected: %s, \nGot: %s ", expected, msg)
	}

	winComputer.setPrinter(hpPrinter)
	msg = winComputer.print()
	expected = "Print request for Windows / Printing by a HP Printer"
	if msg != expected {
		t.Errorf("Invalid printing.\nExpected: %s, \nGot: %s ", expected, msg)
	}

}
