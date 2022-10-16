// Package handlers is a bridge between the API and handlers to create a CLI/API server.
package handlers

import (
	"context"
	_ "embed"

	"github.com/spudtrooper/lyft/api"
	"github.com/spudtrooper/lyft/render"
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
		handler.NewHandlerExtraRequiredFields([]string{"token"}),
	)

	b.NewHandler("NearbyDrivers",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.NearbyDriversParams)
			return client.NearbyDrivers(p.Options()...)
		},
		api.NearbyDriversParams{},
		handler.NewHandlerExtraRequiredFields([]string{"token"}),
		handler.NewHandlerRenderer(render.Status),
	)

	b.NewHandler("RideHistory",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.RideHistoryParams)
			return client.RideHistory(p.Options()...)
		},
		api.RideHistoryParams{},
		handler.NewHandlerExtraRequiredFields([]string{"token"}),
	)

	b.NewHandler("AllRideHistory",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.AllRideHistoryBatchParams)
			return client.AllRideHistoryBatch(p.Options()...)
		},
		api.AllRideHistoryBatchParams{},
		handler.NewHandlerExtraRequiredFields([]string{"token"}),
	)

	b.NewHandler("Offerings",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.OfferingsParams)
			return client.Offerings(p.Options()...)
		},
		api.OfferingsParams{},
		handler.NewHandlerExtraRequiredFields([]string{"token"}),
	)

	return b.Build()
}
