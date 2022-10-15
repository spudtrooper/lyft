// Package handlers is a bridge between the API and handlers to create a CLI/API server.
package handlers

import (
	"context"
	_ "embed"

	"github.com/spudtrooper/lyft/api"
	"github.com/spudtrooper/minimalcli/handler"
)

//go:generate minimalcli gsl --input handlers.go --uri_root "github.com/spudtrooper/lyft/blob/main/handlers" --output handlers.go.json
//go:embed handlers.go.json
var SourceLocations []byte

func CreateHandlers(client *api.Client) []handler.Handler {
	b := handler.NewHandlerBuilder()

	b.NewHandler("PlaceRecommendations",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.PlaceRecommendationsParams)
			return client.PlaceRecommendations(p.Options()...)
		},
		api.PlaceRecommendationsParams{},
	)

	b.NewHandler("NearbyDrivers",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.NearbyDriversParams)
			return client.NearbyDrivers(p.Options()...)
		},
		api.NearbyDriversParams{},
	)

	b.NewHandler("RideHistory",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.RideHistoryParams)
			return client.RideHistory(p.Options()...)
		},
		api.RideHistoryParams{},
	)

	return b.Build()
}
