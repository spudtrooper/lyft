package api

type AllVehicleViewsInfo struct {
	VehicleViews         []VehicleView
	AllNearbyDriversInfo AllNearbyDriversInfo
}

//go:generate genopts --params --function AllVehicleViews --extends Base,NearbyDrivers,AllNearbyDrivers
func (c *Client) AllVehicleViews(optss ...AllVehicleViewsOption) (*AllVehicleViewsInfo, error) {
	opts := MakeAllVehicleViewsOptions(optss...)

	info, err := c.AllNearbyDrivers(opts.ToAllNearbyDriversOptions()...)
	if err != nil {
		return nil, err
	}

	vehicleViews, err := info.VehicleViews()
	if err != nil {
		return nil, err
	}

	res := &AllVehicleViewsInfo{
		VehicleViews:         vehicleViews,
		AllNearbyDriversInfo: *info,
	}

	return res, nil
}
