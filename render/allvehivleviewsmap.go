package render

import (
	_ "embed"
)

//go:generate genopts --params --function AllVehicleViewsMap --required "token string" --extends VehicleViewsMap multiple:int:1
