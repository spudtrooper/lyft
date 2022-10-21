package render

import (
	_ "embed"
)

//go:embed tmpl/vehicle-views-map.html
var VehicleViewsMapTmpl string

//go:generate genopts --params --function VehicleViewsMap --required "token string" latitude:float64:40.7701286 longitude:float64:-73.9829762 zoom:int:14 sleepSecs:int:0
