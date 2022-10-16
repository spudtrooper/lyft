package api

import (
	"strings"

	"github.com/spudtrooper/goutil/request"
)

type nearbyDriversInfoNearbyDriver struct {
	ID        string `json:"id"`
	Locations []struct {
		Bearing      float64 `json:"bearing"`
		Lat          float64 `json:"lat"`
		Lng          float64 `json:"lng"`
		RecordedAtMs int64   `json:"recorded_at_ms"`
	} `json:"locations"`
}

type nearbyDriversInfoNearbyDriverByStableOfferProductID struct {
	MapMarkerImage struct {
		Sources []struct {
			Media struct {
				UserInterfaceStyle int `json:"user_interface_style"`
			} `json:"media"`
			URL string `json:"url"`
		} `json:"sources"`
	} `json:"map_marker_image"`
	NearbyDriverProducts []struct {
		DriverID string `json:"driver_id"`
	} `json:"nearby_driver_products"`
}

type NearbyDriversInfo struct {
	DefaultNearbyDrivers                nearbyDriversInfoNearbyDriverByStableOfferProductID
	NearbyDrivers                       map[string]nearbyDriversInfoNearbyDriver                       `json:"nearby_drivers"`
	NearbyDriversByStableOfferProductID map[string]nearbyDriversInfoNearbyDriverByStableOfferProductID `json:"nearby_drivers_by_stable_offer_product_id"`
}

//go:generate genopts --params --function NearbyDrivers --extends Base latitudeE6:int:40770034 longitudeE6:int:-73982912 orginPlaceID:string:lyft:address:3eaa5572-4d37-4a39-92ed-c61906139955 usingCommuterPayment
func (c *Client) NearbyDrivers(optss ...NearbyDriversOption) (*NearbyDriversInfo, error) {
	opts := MakeNearbyDriversOptions(optss...)

	const uri = "https://ride.lyft.com/v1/nearby-drivers"

	headers := c.makeHeaders(true, opts.ToBaseOptions()...)

	type origin struct {
		LatitudeE6  int `json:"latitude_e6"`
		LongitudeE6 int `json:"longitude_e6"`
	}
	type body struct {
		Origin               origin `json:"origin"`
		OriginPlaceID        string `json:"origin_place_id"`
		UsingCommuterPayment bool   `json:"using_commuter_payment"`
	}

	b, err := request.JSONMarshal(body{
		Origin: origin{
			LatitudeE6:  opts.LatitudeE6(),
			LongitudeE6: opts.LongitudeE6(),
		},
		OriginPlaceID:        opts.OrginPlaceID(),
		UsingCommuterPayment: opts.UsingCommuterPayment(),
	})
	if err != nil {
		return nil, err
	}

	var res NearbyDriversInfo
	if _, err := request.Post(uri, &res, strings.NewReader(string(b)), request.RequestExtraHeaders(headers)); err != nil {
		return nil, err
	}
	return &res, nil
}
