package render

import (
	"bytes"
	_ "embed"
	"sort"

	"github.com/spudtrooper/lyft/api"
	"github.com/spudtrooper/minimalcli/handler"
)

//go:embed tmpl/vehicle-views.html
var viewhicleViewsTmpl string

func renderVehicleViews(vehicleViews []api.VehicleView) ([]byte, handler.RendererConfig, error) {
	config := handler.RendererConfig{
		IsFragment: true,
	}
	sort.Slice(vehicleViews, func(i, j int) bool { return vehicleViews[i].ID < vehicleViews[j].ID })
	var data = struct {
		VehicleViews []api.VehicleView
	}{
		VehicleViews: vehicleViews,
	}
	var buf bytes.Buffer
	if err := renderTemplate(&buf, viewhicleViewsTmpl, "VehicleViews", data); err != nil {
		return nil, config, err
	}
	return buf.Bytes(), config, nil
}
