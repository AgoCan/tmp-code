package options

import (
	"cheetah/internal/command"
)

type AppOptions struct {
	Name string
	Path string
	Type string
}

func NewAppOptions() *AppOptions {
	o := &AppOptions{}
	return o
}

func (o *AppOptions) NewServer() (*command.Command, error) {
	s := command.New()

	return s, nil
}
