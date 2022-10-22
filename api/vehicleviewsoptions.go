// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type VehicleViewsOption struct {
	f func(*vehicleViewsOptionImpl)
	s string
}

func (o VehicleViewsOption) String() string { return o.s }

type VehicleViewsOptions interface {
	DestinationLatitudeE6() int
	HasDestinationLatitudeE6() bool
	DestinationLongitudeE6() int
	HasDestinationLongitudeE6() bool
	OrginPlaceID() string
	HasOrginPlaceID() bool
	OriginLatitudeE6() int
	HasOriginLatitudeE6() bool
	OriginLongitudeE6() int
	HasOriginLongitudeE6() bool
	Token() string
	HasToken() bool
	UsingCommuterPayment() bool
	HasUsingCommuterPayment() bool
	ToBaseOptions() []BaseOption
	ToNearbyDriversOptions() []NearbyDriversOption
}

func VehicleViewsDestinationLatitudeE6(destinationLatitudeE6 int) VehicleViewsOption {
	return VehicleViewsOption{func(opts *vehicleViewsOptionImpl) {
		opts.has_destinationLatitudeE6 = true
		opts.destinationLatitudeE6 = destinationLatitudeE6
	}, fmt.Sprintf("api.VehicleViewsDestinationLatitudeE6(int %+v)", destinationLatitudeE6)}
}
func VehicleViewsDestinationLatitudeE6Flag(destinationLatitudeE6 *int) VehicleViewsOption {
	return VehicleViewsOption{func(opts *vehicleViewsOptionImpl) {
		if destinationLatitudeE6 == nil {
			return
		}
		opts.has_destinationLatitudeE6 = true
		opts.destinationLatitudeE6 = *destinationLatitudeE6
	}, fmt.Sprintf("api.VehicleViewsDestinationLatitudeE6(int %+v)", destinationLatitudeE6)}
}

func VehicleViewsDestinationLongitudeE6(destinationLongitudeE6 int) VehicleViewsOption {
	return VehicleViewsOption{func(opts *vehicleViewsOptionImpl) {
		opts.has_destinationLongitudeE6 = true
		opts.destinationLongitudeE6 = destinationLongitudeE6
	}, fmt.Sprintf("api.VehicleViewsDestinationLongitudeE6(int %+v)", destinationLongitudeE6)}
}
func VehicleViewsDestinationLongitudeE6Flag(destinationLongitudeE6 *int) VehicleViewsOption {
	return VehicleViewsOption{func(opts *vehicleViewsOptionImpl) {
		if destinationLongitudeE6 == nil {
			return
		}
		opts.has_destinationLongitudeE6 = true
		opts.destinationLongitudeE6 = *destinationLongitudeE6
	}, fmt.Sprintf("api.VehicleViewsDestinationLongitudeE6(int %+v)", destinationLongitudeE6)}
}

func VehicleViewsOrginPlaceID(orginPlaceID string) VehicleViewsOption {
	return VehicleViewsOption{func(opts *vehicleViewsOptionImpl) {
		opts.has_orginPlaceID = true
		opts.orginPlaceID = orginPlaceID
	}, fmt.Sprintf("api.VehicleViewsOrginPlaceID(string %+v)", orginPlaceID)}
}
func VehicleViewsOrginPlaceIDFlag(orginPlaceID *string) VehicleViewsOption {
	return VehicleViewsOption{func(opts *vehicleViewsOptionImpl) {
		if orginPlaceID == nil {
			return
		}
		opts.has_orginPlaceID = true
		opts.orginPlaceID = *orginPlaceID
	}, fmt.Sprintf("api.VehicleViewsOrginPlaceID(string %+v)", orginPlaceID)}
}

func VehicleViewsOriginLatitudeE6(originLatitudeE6 int) VehicleViewsOption {
	return VehicleViewsOption{func(opts *vehicleViewsOptionImpl) {
		opts.has_originLatitudeE6 = true
		opts.originLatitudeE6 = originLatitudeE6
	}, fmt.Sprintf("api.VehicleViewsOriginLatitudeE6(int %+v)", originLatitudeE6)}
}
func VehicleViewsOriginLatitudeE6Flag(originLatitudeE6 *int) VehicleViewsOption {
	return VehicleViewsOption{func(opts *vehicleViewsOptionImpl) {
		if originLatitudeE6 == nil {
			return
		}
		opts.has_originLatitudeE6 = true
		opts.originLatitudeE6 = *originLatitudeE6
	}, fmt.Sprintf("api.VehicleViewsOriginLatitudeE6(int %+v)", originLatitudeE6)}
}

func VehicleViewsOriginLongitudeE6(originLongitudeE6 int) VehicleViewsOption {
	return VehicleViewsOption{func(opts *vehicleViewsOptionImpl) {
		opts.has_originLongitudeE6 = true
		opts.originLongitudeE6 = originLongitudeE6
	}, fmt.Sprintf("api.VehicleViewsOriginLongitudeE6(int %+v)", originLongitudeE6)}
}
func VehicleViewsOriginLongitudeE6Flag(originLongitudeE6 *int) VehicleViewsOption {
	return VehicleViewsOption{func(opts *vehicleViewsOptionImpl) {
		if originLongitudeE6 == nil {
			return
		}
		opts.has_originLongitudeE6 = true
		opts.originLongitudeE6 = *originLongitudeE6
	}, fmt.Sprintf("api.VehicleViewsOriginLongitudeE6(int %+v)", originLongitudeE6)}
}

func VehicleViewsToken(token string) VehicleViewsOption {
	return VehicleViewsOption{func(opts *vehicleViewsOptionImpl) {
		opts.has_token = true
		opts.token = token
	}, fmt.Sprintf("api.VehicleViewsToken(string %+v)", token)}
}
func VehicleViewsTokenFlag(token *string) VehicleViewsOption {
	return VehicleViewsOption{func(opts *vehicleViewsOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}, fmt.Sprintf("api.VehicleViewsToken(string %+v)", token)}
}

func VehicleViewsUsingCommuterPayment(usingCommuterPayment bool) VehicleViewsOption {
	return VehicleViewsOption{func(opts *vehicleViewsOptionImpl) {
		opts.has_usingCommuterPayment = true
		opts.usingCommuterPayment = usingCommuterPayment
	}, fmt.Sprintf("api.VehicleViewsUsingCommuterPayment(bool %+v)", usingCommuterPayment)}
}
func VehicleViewsUsingCommuterPaymentFlag(usingCommuterPayment *bool) VehicleViewsOption {
	return VehicleViewsOption{func(opts *vehicleViewsOptionImpl) {
		if usingCommuterPayment == nil {
			return
		}
		opts.has_usingCommuterPayment = true
		opts.usingCommuterPayment = *usingCommuterPayment
	}, fmt.Sprintf("api.VehicleViewsUsingCommuterPayment(bool %+v)", usingCommuterPayment)}
}

type vehicleViewsOptionImpl struct {
	destinationLatitudeE6      int
	has_destinationLatitudeE6  bool
	destinationLongitudeE6     int
	has_destinationLongitudeE6 bool
	orginPlaceID               string
	has_orginPlaceID           bool
	originLatitudeE6           int
	has_originLatitudeE6       bool
	originLongitudeE6          int
	has_originLongitudeE6      bool
	token                      string
	has_token                  bool
	usingCommuterPayment       bool
	has_usingCommuterPayment   bool
}

func (v *vehicleViewsOptionImpl) DestinationLatitudeE6() int     { return v.destinationLatitudeE6 }
func (v *vehicleViewsOptionImpl) HasDestinationLatitudeE6() bool { return v.has_destinationLatitudeE6 }
func (v *vehicleViewsOptionImpl) DestinationLongitudeE6() int    { return v.destinationLongitudeE6 }
func (v *vehicleViewsOptionImpl) HasDestinationLongitudeE6() bool {
	return v.has_destinationLongitudeE6
}
func (v *vehicleViewsOptionImpl) OrginPlaceID() string {
	return or.String(v.orginPlaceID, "lyft:address:3eaa5572-4d37-4a39-92ed-c61906139955")
}
func (v *vehicleViewsOptionImpl) HasOrginPlaceID() bool     { return v.has_orginPlaceID }
func (v *vehicleViewsOptionImpl) OriginLatitudeE6() int     { return or.Int(v.originLatitudeE6, 40770034) }
func (v *vehicleViewsOptionImpl) HasOriginLatitudeE6() bool { return v.has_originLatitudeE6 }
func (v *vehicleViewsOptionImpl) OriginLongitudeE6() int {
	return or.Int(v.originLongitudeE6, -73982912)
}
func (v *vehicleViewsOptionImpl) HasOriginLongitudeE6() bool    { return v.has_originLongitudeE6 }
func (v *vehicleViewsOptionImpl) Token() string                 { return v.token }
func (v *vehicleViewsOptionImpl) HasToken() bool                { return v.has_token }
func (v *vehicleViewsOptionImpl) UsingCommuterPayment() bool    { return v.usingCommuterPayment }
func (v *vehicleViewsOptionImpl) HasUsingCommuterPayment() bool { return v.has_usingCommuterPayment }

type VehicleViewsParams struct {
	DestinationLatitudeE6  int    `json:"destination_latitude_e_6"`
	DestinationLongitudeE6 int    `json:"destination_longitude_e_6"`
	OrginPlaceID           string `json:"orgin_place_id" default:"\"lyft:address:3eaa5572-4d37-4a39-92ed-c61906139955\""`
	OriginLatitudeE6       int    `json:"origin_latitude_e_6" default:"40770034"`
	OriginLongitudeE6      int    `json:"origin_longitude_e_6" default:"-73982912"`
	Token                  string `json:"token"`
	UsingCommuterPayment   bool   `json:"using_commuter_payment"`
}

func (o VehicleViewsParams) Options() []VehicleViewsOption {
	return []VehicleViewsOption{
		VehicleViewsDestinationLatitudeE6(o.DestinationLatitudeE6),
		VehicleViewsDestinationLongitudeE6(o.DestinationLongitudeE6),
		VehicleViewsOrginPlaceID(o.OrginPlaceID),
		VehicleViewsOriginLatitudeE6(o.OriginLatitudeE6),
		VehicleViewsOriginLongitudeE6(o.OriginLongitudeE6),
		VehicleViewsToken(o.Token),
		VehicleViewsUsingCommuterPayment(o.UsingCommuterPayment),
	}
}

// ToBaseOptions converts VehicleViewsOption to an array of BaseOption
func (o *vehicleViewsOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
	}
}

// ToNearbyDriversOptions converts VehicleViewsOption to an array of NearbyDriversOption
func (o *vehicleViewsOptionImpl) ToNearbyDriversOptions() []NearbyDriversOption {
	return []NearbyDriversOption{
		NearbyDriversDestinationLatitudeE6(o.DestinationLatitudeE6()),
		NearbyDriversDestinationLongitudeE6(o.DestinationLongitudeE6()),
		NearbyDriversOrginPlaceID(o.OrginPlaceID()),
		NearbyDriversOriginLatitudeE6(o.OriginLatitudeE6()),
		NearbyDriversOriginLongitudeE6(o.OriginLongitudeE6()),
		NearbyDriversToken(o.Token()),
		NearbyDriversUsingCommuterPayment(o.UsingCommuterPayment()),
	}
}

func makeVehicleViewsOptionImpl(opts ...VehicleViewsOption) *vehicleViewsOptionImpl {
	res := &vehicleViewsOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeVehicleViewsOptions(opts ...VehicleViewsOption) VehicleViewsOptions {
	return makeVehicleViewsOptionImpl(opts...)
}
