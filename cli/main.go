package cli

import (
	"context"

	"github.com/spudtrooper/lyft/api"
	"github.com/spudtrooper/lyft/handlers"
	"github.com/spudtrooper/minimalcli/handler"
	flag "github.com/spudtrooper/minimalcli/handler"
)

func Main(ctx context.Context) error {
	flag.Bool("debug_payload", false, "debug payload")
	flag.Float64("lat", 0, "latitude")
	flag.Float64("lng", 0, "longitude")
	flag.Int("accuracy", 0, "accuracy")

	client, err := api.NewClientFromFlags()
	if err != nil {
		return err
	}

	return handler.RunCLI(ctx, handlers.CreateHandlers(client)...)
}
