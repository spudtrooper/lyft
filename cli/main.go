package cli

import (
	"context"

	"github.com/spudtrooper/lyft/api"
	"github.com/spudtrooper/lyft/handlers"
	"github.com/spudtrooper/minimalcli/handler"
	flag "github.com/spudtrooper/minimalcli/handler"
)

func Main(ctx context.Context) error {
	flag.Float64("lat", 0, "latitude")
	flag.Float64("lng", 0, "longitude")
	flag.Int("accuracy", 0, "accuracy")
	flag.Int("limit", 0, "limit")
	flag.Duration("timeout", 0, "timeout")
	flag.Bool("debug", false, "debug")

	client, err := api.NewClientFromFlags()
	if err != nil {
		return err
	}

	return handler.RunCLI(ctx, handlers.CreateHandlers("api", client)...)
}
