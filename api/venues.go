package api

import (
	"github.com/spudtrooper/goutil/request"
)

// TODO: Make this something.
type VenuesInfo interface{}

//go:generate genopts --params --function Venues --extends Base latitude:float64:40.770114 longitude:float64:-73.98302 radiusMiles:int:1 noIncludeLocationV2Info
func (c *Client) Venues(optss ...VenuesOption) (VenuesInfo, error) {
	opts := MakeVenuesOptions(optss...)

	uri := request.MakeURL("https://ride.lyft.com/v3/venues",
		request.MakeParam("type[]", `pickup`),
		request.MakeParam("type[]", `prohibited`),
		request.MakeParam("lat", opts.Latitude()),
		request.MakeParam("lng", opts.Longitude()),
		request.MakeParam("radius_miles", opts.RadiusMiles()),
		request.MakeParam("include_location_v2_info", !opts.NoIncludeLocationV2Info()),
	)

	var res VenuesInfo
	if _, err := request.Get(uri, &res); err != nil {
		return nil, err
	}
	return res, nil
}
