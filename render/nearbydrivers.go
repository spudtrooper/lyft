package render

import (
	"bytes"
	_ "embed"
	"sort"

	"github.com/spudtrooper/lyft/api"
	"github.com/spudtrooper/minimalcli/handler"
)

//go:embed tmpl/nearby-drivers.html
var nearbyDriversTmpl string

func NearbyDrivers(input any) ([]byte, handler.RendererConfig, error) {
	params := input.(*api.NearbyDriversInfo)

	config := handler.RendererConfig{
		IsFragment: true,
	}
	vehicleViews, err := params.VehicleViews()
	if err != nil {
		return nil, config, err
	}
	sort.Slice(vehicleViews, func(i, j int) bool { return vehicleViews[i].ID < vehicleViews[j].ID })
	var data = struct {
		VehicleViews []api.VehicleView
	}{
		VehicleViews: vehicleViews,
	}
	var buf bytes.Buffer
	if err := renderTemplate(&buf, nearbyDriversTmpl, "NearbyDrivers", data); err != nil {
		return nil, config, err
	}
	return buf.Bytes(), config, nil
}
