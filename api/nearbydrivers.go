package api

import (
	"strings"

	"github.com/spudtrooper/goutil/request"
)

type NearbyDriversInfoNearbyDriver struct {
	ID        string `json:"id"`
	Locations []struct {
		Bearing      float64 `json:"bearing"`
		Lat          float64 `json:"lat"`
		Lng          float64 `json:"lng"`
		RecordedAtMs int64   `json:"recorded_at_ms"`
	} `json:"locations"`
}

type NearbyDriversInfoNearbyDriverByStableOfferProductID struct {
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
	DefaultNearbyDrivers                NearbyDriversInfoNearbyDriverByStableOfferProductID
	NearbyDrivers                       map[string]NearbyDriversInfoNearbyDriver                       `json:"nearby_drivers"`
	NearbyDriversByStableOfferProductID map[string]NearbyDriversInfoNearbyDriverByStableOfferProductID `json:"nearby_drivers_by_stable_offer_product_id"`
}

//go:generate genopts --params --function NearbyDrivers --extends Base originLatitudeE6:int:40770034 originLongitudeE6:int:-73982912 destinationLatitudeE6:int destinationLongitudeE6:int orginPlaceID:string:lyft:address:3eaa5572-4d37-4a39-92ed-c61906139955 usingCommuterPayment
func (c *Client) NearbyDrivers(optss ...NearbyDriversOption) (*NearbyDriversInfo, error) {
	opts := MakeNearbyDriversOptions(optss...)

	const uri = "https://ride.lyft.com/v1/nearby-drivers"

	headers := c.makeHeaders(true, opts.ToBaseOptions()...)

	type loc struct {
		LatitudeE6  int `json:"latitude_e6"`
		LongitudeE6 int `json:"longitude_e6"`
	}
	type body struct {
		Origin               *loc   `json:"origin"`
		Destination          *loc   `json:"destination,omitempty"`
		OriginPlaceID        string `json:"origin_place_id"`
		UsingCommuterPayment bool   `json:"using_commuter_payment"`
	}
	params := body{
		Origin: &loc{
			LatitudeE6:  opts.OriginLatitudeE6(),
			LongitudeE6: opts.OriginLongitudeE6(),
		},
		OriginPlaceID:        opts.OrginPlaceID(),
		UsingCommuterPayment: opts.UsingCommuterPayment(),
	}
	if opts.DestinationLatitudeE6() != 0 && opts.DestinationLongitudeE6() != 0 {
		params.Destination = &loc{
			LatitudeE6:  opts.DestinationLatitudeE6(),
			LongitudeE6: opts.DestinationLongitudeE6(),
		}
	}

	b, err := request.JSONMarshal(params)
	if err != nil {
		return nil, err
	}

	var res NearbyDriversInfo
	if _, err := request.Post(uri, &res, strings.NewReader(string(b)), request.RequestExtraHeaders(headers)); err != nil {
		return nil, err
	}
	return &res, nil
}
