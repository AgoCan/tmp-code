package command

import (
	"fmt"
)

type Command struct {
}

func New() *Command {
	return &Command{}
}

func (s *Command) Run() {
	fmt.Println("run command")
}
