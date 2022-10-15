// Package handlers is a bridge between the API and handlers to create a CLI/API server.
package handlers

import (
	"context"
	_ "embed"

	goutilerrors "github.com/spudtrooper/goutil/errors"
	"github.com/spudtrooper/goutil/parallel"
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
		handler.NewHandlerExtraRequiredFields([]string{"token"}),
	)

	b.NewHandler("NearbyDrivers",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.NearbyDriversParams)
			return client.NearbyDrivers(p.Options()...)
		},
		api.NearbyDriversParams{},
		handler.NewHandlerExtraRequiredFields([]string{"token"}),
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
			p := ip.(api.AllRideHistoryParams)
			data, errs := client.AllRideHistory(p.Options()...)
			var res []any
			errBuilder := goutilerrors.MakeErrorCollector()
			parallel.WaitFor(func() {
				for d := range data {
					res = append(res, d)
				}
			}, func() {
				for e := range errs {
					errBuilder.Add(e)
				}
			})

			if err := errBuilder.Build(); err != nil {
				return nil, err
			}
			return res, nil
		},
		api.AllRideHistoryParams{},
		handler.NewHandlerExtraRequiredFields([]string{"token"}),
	)

	return b.Build()
}
