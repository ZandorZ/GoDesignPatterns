package chain

import (
	"fmt"
	"io"
	"strings"
)

// ChainLogger ...
type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (f *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)

	if f.NextChain != nil {
		f.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (sl *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n", s)

		if sl.NextChain != nil {
			sl.NextChain.Next(s)
		}
		return
	}

	fmt.Printf("Finishing in second logging\n\n")
}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (w *WriterLogger) Next(s string) {
	if w.Writer != nil {
		w.Writer.Write([]byte("WriterLogger: " + s))
	}

	if w.NextChain != nil {
		w.NextChain.Next(s)
	}
}

type ClosureChain struct {
	NextChain ChainLogger
	Clousure  func(string)
}

func (c *ClosureChain) Next(s string) {
	if c.Clousure != nil {
		c.Clousure(s)
	}

	if c.NextChain != nil {
		c.NextChain.Next(s)
	}

}