package render

import (
	"github.com/spudtrooper/lyft/api"
	"github.com/spudtrooper/minimalcli/handler"
)

func RideHistory(input any) ([]byte, handler.RendererConfig, error) {
	params := input.(*api.RideHistoryInfo)

	return rideHistoryFromRides(params.Data)
}
