package main

import (
	"context"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/lyft/cli"
)

func main() {
	check.Err(cli.Main(context.Background()))
}
