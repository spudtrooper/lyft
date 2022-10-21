package render

import (
	_ "embed"
)

//go:generate genopts --params --function AllVehicleViewsMap --required "token string" --extends VehicleViewsMap multiples:int:1 deltaE6:int:130
