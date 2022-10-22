// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type NearbyDriversOption struct {
	f func(*nearbyDriversOptionImpl)
	s string
}

func (o NearbyDriversOption) String() string { return o.s }

type NearbyDriversOptions interface {
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
}

func NearbyDriversDestinationLatitudeE6(destinationLatitudeE6 int) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		opts.has_destinationLatitudeE6 = true
		opts.destinationLatitudeE6 = destinationLatitudeE6
	}, fmt.Sprintf("api.NearbyDriversDestinationLatitudeE6(int %+v)", destinationLatitudeE6)}
}
func NearbyDriversDestinationLatitudeE6Flag(destinationLatitudeE6 *int) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		if destinationLatitudeE6 == nil {
			return
		}
		opts.has_destinationLatitudeE6 = true
		opts.destinationLatitudeE6 = *destinationLatitudeE6
	}, fmt.Sprintf("api.NearbyDriversDestinationLatitudeE6(int %+v)", destinationLatitudeE6)}
}

func NearbyDriversDestinationLongitudeE6(destinationLongitudeE6 int) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		opts.has_destinationLongitudeE6 = true
		opts.destinationLongitudeE6 = destinationLongitudeE6
	}, fmt.Sprintf("api.NearbyDriversDestinationLongitudeE6(int %+v)", destinationLongitudeE6)}
}
func NearbyDriversDestinationLongitudeE6Flag(destinationLongitudeE6 *int) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		if destinationLongitudeE6 == nil {
			return
		}
		opts.has_destinationLongitudeE6 = true
		opts.destinationLongitudeE6 = *destinationLongitudeE6
	}, fmt.Sprintf("api.NearbyDriversDestinationLongitudeE6(int %+v)", destinationLongitudeE6)}
}

func NearbyDriversOrginPlaceID(orginPlaceID string) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		opts.has_orginPlaceID = true
		opts.orginPlaceID = orginPlaceID
	}, fmt.Sprintf("api.NearbyDriversOrginPlaceID(string %+v)", orginPlaceID)}
}
func NearbyDriversOrginPlaceIDFlag(orginPlaceID *string) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		if orginPlaceID == nil {
			return
		}
		opts.has_orginPlaceID = true
		opts.orginPlaceID = *orginPlaceID
	}, fmt.Sprintf("api.NearbyDriversOrginPlaceID(string %+v)", orginPlaceID)}
}

func NearbyDriversOriginLatitudeE6(originLatitudeE6 int) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		opts.has_originLatitudeE6 = true
		opts.originLatitudeE6 = originLatitudeE6
	}, fmt.Sprintf("api.NearbyDriversOriginLatitudeE6(int %+v)", originLatitudeE6)}
}
func NearbyDriversOriginLatitudeE6Flag(originLatitudeE6 *int) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		if originLatitudeE6 == nil {
			return
		}
		opts.has_originLatitudeE6 = true
		opts.originLatitudeE6 = *originLatitudeE6
	}, fmt.Sprintf("api.NearbyDriversOriginLatitudeE6(int %+v)", originLatitudeE6)}
}

func NearbyDriversOriginLongitudeE6(originLongitudeE6 int) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		opts.has_originLongitudeE6 = true
		opts.originLongitudeE6 = originLongitudeE6
	}, fmt.Sprintf("api.NearbyDriversOriginLongitudeE6(int %+v)", originLongitudeE6)}
}
func NearbyDriversOriginLongitudeE6Flag(originLongitudeE6 *int) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		if originLongitudeE6 == nil {
			return
		}
		opts.has_originLongitudeE6 = true
		opts.originLongitudeE6 = *originLongitudeE6
	}, fmt.Sprintf("api.NearbyDriversOriginLongitudeE6(int %+v)", originLongitudeE6)}
}

func NearbyDriversToken(token string) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		opts.has_token = true
		opts.token = token
	}, fmt.Sprintf("api.NearbyDriversToken(string %+v)", token)}
}
func NearbyDriversTokenFlag(token *string) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}, fmt.Sprintf("api.NearbyDriversToken(string %+v)", token)}
}

func NearbyDriversUsingCommuterPayment(usingCommuterPayment bool) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		opts.has_usingCommuterPayment = true
		opts.usingCommuterPayment = usingCommuterPayment
	}, fmt.Sprintf("api.NearbyDriversUsingCommuterPayment(bool %+v)", usingCommuterPayment)}
}
func NearbyDriversUsingCommuterPaymentFlag(usingCommuterPayment *bool) NearbyDriversOption {
	return NearbyDriversOption{func(opts *nearbyDriversOptionImpl) {
		if usingCommuterPayment == nil {
			return
		}
		opts.has_usingCommuterPayment = true
		opts.usingCommuterPayment = *usingCommuterPayment
	}, fmt.Sprintf("api.NearbyDriversUsingCommuterPayment(bool %+v)", usingCommuterPayment)}
}

type nearbyDriversOptionImpl struct {
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

func (n *nearbyDriversOptionImpl) DestinationLatitudeE6() int     { return n.destinationLatitudeE6 }
func (n *nearbyDriversOptionImpl) HasDestinationLatitudeE6() bool { return n.has_destinationLatitudeE6 }
func (n *nearbyDriversOptionImpl) DestinationLongitudeE6() int    { return n.destinationLongitudeE6 }
func (n *nearbyDriversOptionImpl) HasDestinationLongitudeE6() bool {
	return n.has_destinationLongitudeE6
}
func (n *nearbyDriversOptionImpl) OrginPlaceID() string {
	return or.String(n.orginPlaceID, "lyft:address:3eaa5572-4d37-4a39-92ed-c61906139955")
}
func (n *nearbyDriversOptionImpl) HasOrginPlaceID() bool     { return n.has_orginPlaceID }
func (n *nearbyDriversOptionImpl) OriginLatitudeE6() int     { return or.Int(n.originLatitudeE6, 40770034) }
func (n *nearbyDriversOptionImpl) HasOriginLatitudeE6() bool { return n.has_originLatitudeE6 }
func (n *nearbyDriversOptionImpl) OriginLongitudeE6() int {
	return or.Int(n.originLongitudeE6, -73982912)
}
func (n *nearbyDriversOptionImpl) HasOriginLongitudeE6() bool    { return n.has_originLongitudeE6 }
func (n *nearbyDriversOptionImpl) Token() string                 { return n.token }
func (n *nearbyDriversOptionImpl) HasToken() bool                { return n.has_token }
func (n *nearbyDriversOptionImpl) UsingCommuterPayment() bool    { return n.usingCommuterPayment }
func (n *nearbyDriversOptionImpl) HasUsingCommuterPayment() bool { return n.has_usingCommuterPayment }

type NearbyDriversParams struct {
	DestinationLatitudeE6  int    `json:"destination_latitude_e_6"`
	DestinationLongitudeE6 int    `json:"destination_longitude_e_6"`
	OrginPlaceID           string `json:"orgin_place_id" default:"\"lyft:address:3eaa5572-4d37-4a39-92ed-c61906139955\""`
	OriginLatitudeE6       int    `json:"origin_latitude_e_6" default:"40770034"`
	OriginLongitudeE6      int    `json:"origin_longitude_e_6" default:"-73982912"`
	Token                  string `json:"token"`
	UsingCommuterPayment   bool   `json:"using_commuter_payment"`
}

func (o NearbyDriversParams) Options() []NearbyDriversOption {
	return []NearbyDriversOption{
		NearbyDriversDestinationLatitudeE6(o.DestinationLatitudeE6),
		NearbyDriversDestinationLongitudeE6(o.DestinationLongitudeE6),
		NearbyDriversOrginPlaceID(o.OrginPlaceID),
		NearbyDriversOriginLatitudeE6(o.OriginLatitudeE6),
		NearbyDriversOriginLongitudeE6(o.OriginLongitudeE6),
		NearbyDriversToken(o.Token),
		NearbyDriversUsingCommuterPayment(o.UsingCommuterPayment),
	}
}

// ToBaseOptions converts NearbyDriversOption to an array of BaseOption
func (o *nearbyDriversOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
	}
}

func makeNearbyDriversOptionImpl(opts ...NearbyDriversOption) *nearbyDriversOptionImpl {
	res := &nearbyDriversOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeNearbyDriversOptions(opts ...NearbyDriversOption) NearbyDriversOptions {
	return makeNearbyDriversOptionImpl(opts...)
}
