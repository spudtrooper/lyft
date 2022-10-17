package render

import (
	"github.com/spudtrooper/lyft/api"
	"github.com/spudtrooper/minimalcli/handler"
)

func AllRideHistoryBatch(input any) ([]byte, handler.RendererConfig, error) {
	params := input.(*api.AllRideHistoryBatchInfo)

	var data []api.RideHistoryInfoData
	for _, r := range params.Rides {
		data = append(data, r.Data...)
	}

	return rideHistoryFromRides(data)
}
