package render

import (
	_ "embed"

	"github.com/spudtrooper/lyft/api"
	"github.com/spudtrooper/minimalcli/handler"
)

func VehicleViews(input any) ([]byte, handler.RendererConfig, error) {
	params := input.(*api.VehicleViewsInfo)
	return renderVehicleViews(params.VehicleViews)
}
