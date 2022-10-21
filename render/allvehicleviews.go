package render

import (
	_ "embed"

	"github.com/spudtrooper/lyft/api"
	"github.com/spudtrooper/minimalcli/handler"
)

func AllVehicleViews(input any) ([]byte, handler.RendererConfig, error) {
	params := input.(*api.AllVehicleViewsInfo)
	return renderVehicleViews(params.VehicleViews)
}
