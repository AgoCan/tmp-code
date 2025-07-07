package main

import (
	"log"

	"cheetah/cmd/cheetah/app"
)

func main() {
	cmd := app.NewServerCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
