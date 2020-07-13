package main

import "fmt"

// Command ...
type Command interface {
	Execute()
}

type ConsoleOutput struct {
	message string
}

func (c *ConsoleOutput) Execute() {
	fmt.Println(c.message)
}

func CreateCommand(s string) Command {
	fmt.Println("Creating command")
	return &ConsoleOutput{
		message: s,
	}
}

type CommandQueue struct {
	queue []Command
}

func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)
	if len(p.queue) == 3 {
		for _, command := range p.queue {
			command.Execute()
		}
		p.queue = make([]Command, 3)
	}
}

func main() {
	queue := CommandQueue{}
	queue.AddCommand(CreateCommand("First command"))
	queue.AddCommand(CreateCommand("Second command"))
	queue.AddCommand(CreateCommand("Third command"))
	queue.AddCommand(CreateCommand("Fourth command"))
	queue.AddCommand(CreateCommand("Fifth command"))
}
