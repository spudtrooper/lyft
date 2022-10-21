// Package handlers is a bridge between the API and handlers to create a CLI/API server.
package handlers

import (
	"bytes"
	"context"
	_ "embed"
	"html/template"
	"io"
	"log"
	"strings"

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
	)

	b.NewHandler("AllNearbyDrivers",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.AllNearbyDriversParams)
			return client.AllNearbyDrivers(p.Options()...)
		},
		api.AllNearbyDriversParams{},
		handler.NewHandlerExtraRequiredFields([]string{"token"}),
	)

	b.NewHandler("VehicleViews",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.VehicleViewsParams)
			return client.VehicleViews(p.Options()...)
		},
		api.VehicleViewsParams{},
		handler.NewHandlerExtraRequiredFields([]string{"token"}),
		handler.NewHandlerRenderer(render.VehicleViews),
	)

	var vehicleViewsData = struct {
		Route string
	}{
		Route: "vehicleviews",
	}
	var vehicleViewsBuf bytes.Buffer
	if err := renderTemplate(&vehicleViewsBuf, render.VehicleViewsMapTmpl, "VehicleViewsMap", vehicleViewsData); err != nil {
		// TODO: return error instead
		log.Fatalf("vehicleViews error: %v", err)
	}

	b.NewStaticHandler("VehicleViewsMap",
		vehicleViewsBuf.Bytes(),
		render.VehicleViewsMapParams{},
		handler.NewHandlerRendererConfig(handler.RendererConfig{IsFragment: false}),
	)

	b.NewHandler("AllVehicleViews",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.AllVehicleViewsParams)
			return client.AllVehicleViews(p.Options()...)
		},
		api.AllVehicleViewsParams{},
		handler.NewHandlerExtraRequiredFields([]string{"token"}),
		handler.NewHandlerRenderer(render.AllVehicleViews),
	)

	var allVehicleViewsData = struct {
		Route string
	}{
		Route: "allvehicleviews",
	}
	var allVehicleViewsBuf bytes.Buffer
	if err := renderTemplate(&allVehicleViewsBuf, render.VehicleViewsMapTmpl, "AllVehicleViewsMap", allVehicleViewsData); err != nil {
		// TODO: return error instead
		log.Fatalf("allVehicleViewsBuf error: %v", err)
	}

	b.NewStaticHandler("AllVehicleViewsMap",
		allVehicleViewsBuf.Bytes(),
		render.AllVehicleViewsMapParams{},
		handler.NewHandlerRendererConfig(handler.RendererConfig{IsFragment: false}),
	)

	b.NewHandler("RideHistory",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.RideHistoryParams)
			return client.RideHistory(p.Options()...)
		},
		api.RideHistoryParams{},
		handler.NewHandlerExtraRequiredFields([]string{"token"}),
		handler.NewHandlerRenderer(render.RideHistory),
	)

	b.NewHandler("AllRideHistory",
		func(ctx context.Context, ip any) (any, error) {
			p := ip.(api.AllRideHistoryBatchParams)
			return client.AllRideHistoryBatch(p.Options()...)
		},
		api.AllRideHistoryBatchParams{},
		handler.NewHandlerExtraRequiredFields([]string{"token"}),
		handler.NewHandlerRenderer(render.AllRideHistoryBatch),
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

func renderTemplate(buf io.Writer, t string, name string, data interface{}) error {
	tmpl, err := template.New(name).Parse(strings.TrimSpace(t))
	if err != nil {
		return err
	}
	if err := tmpl.Execute(buf, data); err != nil {
		return err
	}
	return nil
}
