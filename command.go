package smerch

import "github.com/spf13/cobra"

type Command struct {
	GreetCMD *cobra.Command
}

func NewCommand() *Command {
	return &Command{
		GreetCMD: nil,
	}
}

func (c *Command) Init() {
	c.GreetCMD = &cobra.Command{
		Use: "greet [name]",
	}
}
