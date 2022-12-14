package api

type VehicleViewsInfo struct {
	VehicleViews      []VehicleView
	NearbyDriversInfo NearbyDriversInfo
	Centers           []PointE6
}

//go:generate genopts --params --function VehicleViews --extends Base,NearbyDrivers
func (c *Client) VehicleViews(optss ...VehicleViewsOption) (*VehicleViewsInfo, error) {
	opts := MakeVehicleViewsOptions(optss...)

	info, err := c.NearbyDrivers(opts.ToNearbyDriversOptions()...)
	if err != nil {
		return nil, err
	}

	vehicleViews, err := info.VehicleViews()
	if err != nil {
		return nil, err
	}

	res := &VehicleViewsInfo{
		VehicleViews:      vehicleViews,
		NearbyDriversInfo: *info,
		Centers:           info.Centers,
	}

	return res, nil
}
