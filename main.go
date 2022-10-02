package main

import (
	_ "go.uber.org/automaxprocs"

	"github.com/mariner-group/marinerd/cmd"
)

func main() {
	cmd.Execute()
}
