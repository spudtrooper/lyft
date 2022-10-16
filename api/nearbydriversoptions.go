// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"github.com/spudtrooper/goutil/or"
)

type NearbyDriversOption func(*nearbyDriversOptionImpl)

type NearbyDriversOptions interface {
	OriginLatitudeE6() int
	HasOriginLatitudeE6() bool
	OriginLongitudeE6() int
	HasOriginLongitudeE6() bool
	DestinationLatitudeE6() int
	HasDestinationLatitudeE6() bool
	DestinationLongitudeE6() int
	HasDestinationLongitudeE6() bool
	OrginPlaceID() string
	HasOrginPlaceID() bool
	UsingCommuterPayment() bool
	HasUsingCommuterPayment() bool
	Token() string
	HasToken() bool
	ToBaseOptions() []BaseOption
}

func NearbyDriversOriginLatitudeE6(originLatitudeE6 int) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		opts.has_originLatitudeE6 = true
		opts.originLatitudeE6 = originLatitudeE6
	}
}
func NearbyDriversOriginLatitudeE6Flag(originLatitudeE6 *int) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		if originLatitudeE6 == nil {
			return
		}
		opts.has_originLatitudeE6 = true
		opts.originLatitudeE6 = *originLatitudeE6
	}
}

func NearbyDriversOriginLongitudeE6(originLongitudeE6 int) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		opts.has_originLongitudeE6 = true
		opts.originLongitudeE6 = originLongitudeE6
	}
}
func NearbyDriversOriginLongitudeE6Flag(originLongitudeE6 *int) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		if originLongitudeE6 == nil {
			return
		}
		opts.has_originLongitudeE6 = true
		opts.originLongitudeE6 = *originLongitudeE6
	}
}

func NearbyDriversDestinationLatitudeE6(destinationLatitudeE6 int) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		opts.has_destinationLatitudeE6 = true
		opts.destinationLatitudeE6 = destinationLatitudeE6
	}
}
func NearbyDriversDestinationLatitudeE6Flag(destinationLatitudeE6 *int) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		if destinationLatitudeE6 == nil {
			return
		}
		opts.has_destinationLatitudeE6 = true
		opts.destinationLatitudeE6 = *destinationLatitudeE6
	}
}

func NearbyDriversDestinationLongitudeE6(destinationLongitudeE6 int) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		opts.has_destinationLongitudeE6 = true
		opts.destinationLongitudeE6 = destinationLongitudeE6
	}
}
func NearbyDriversDestinationLongitudeE6Flag(destinationLongitudeE6 *int) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		if destinationLongitudeE6 == nil {
			return
		}
		opts.has_destinationLongitudeE6 = true
		opts.destinationLongitudeE6 = *destinationLongitudeE6
	}
}

func NearbyDriversOrginPlaceID(orginPlaceID string) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		opts.has_orginPlaceID = true
		opts.orginPlaceID = orginPlaceID
	}
}
func NearbyDriversOrginPlaceIDFlag(orginPlaceID *string) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		if orginPlaceID == nil {
			return
		}
		opts.has_orginPlaceID = true
		opts.orginPlaceID = *orginPlaceID
	}
}

func NearbyDriversUsingCommuterPayment(usingCommuterPayment bool) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		opts.has_usingCommuterPayment = true
		opts.usingCommuterPayment = usingCommuterPayment
	}
}
func NearbyDriversUsingCommuterPaymentFlag(usingCommuterPayment *bool) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		if usingCommuterPayment == nil {
			return
		}
		opts.has_usingCommuterPayment = true
		opts.usingCommuterPayment = *usingCommuterPayment
	}
}

func NearbyDriversToken(token string) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		opts.has_token = true
		opts.token = token
	}
}
func NearbyDriversTokenFlag(token *string) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}
}

type nearbyDriversOptionImpl struct {
	originLatitudeE6           int
	has_originLatitudeE6       bool
	originLongitudeE6          int
	has_originLongitudeE6      bool
	destinationLatitudeE6      int
	has_destinationLatitudeE6  bool
	destinationLongitudeE6     int
	has_destinationLongitudeE6 bool
	orginPlaceID               string
	has_orginPlaceID           bool
	usingCommuterPayment       bool
	has_usingCommuterPayment   bool
	token                      string
	has_token                  bool
}

func (n *nearbyDriversOptionImpl) OriginLatitudeE6() int     { return or.Int(n.originLatitudeE6, 40770034) }
func (n *nearbyDriversOptionImpl) HasOriginLatitudeE6() bool { return n.has_originLatitudeE6 }
func (n *nearbyDriversOptionImpl) OriginLongitudeE6() int {
	return or.Int(n.originLongitudeE6, -73982912)
}
func (n *nearbyDriversOptionImpl) HasOriginLongitudeE6() bool     { return n.has_originLongitudeE6 }
func (n *nearbyDriversOptionImpl) DestinationLatitudeE6() int     { return n.destinationLatitudeE6 }
func (n *nearbyDriversOptionImpl) HasDestinationLatitudeE6() bool { return n.has_destinationLatitudeE6 }
func (n *nearbyDriversOptionImpl) DestinationLongitudeE6() int    { return n.destinationLongitudeE6 }
func (n *nearbyDriversOptionImpl) HasDestinationLongitudeE6() bool {
	return n.has_destinationLongitudeE6
}
func (n *nearbyDriversOptionImpl) OrginPlaceID() string {
	return or.String(n.orginPlaceID, "lyft:address:3eaa5572-4d37-4a39-92ed-c61906139955")
}
func (n *nearbyDriversOptionImpl) HasOrginPlaceID() bool         { return n.has_orginPlaceID }
func (n *nearbyDriversOptionImpl) UsingCommuterPayment() bool    { return n.usingCommuterPayment }
func (n *nearbyDriversOptionImpl) HasUsingCommuterPayment() bool { return n.has_usingCommuterPayment }
func (n *nearbyDriversOptionImpl) Token() string                 { return n.token }
func (n *nearbyDriversOptionImpl) HasToken() bool                { return n.has_token }

type NearbyDriversParams struct {
	OriginLatitudeE6       int    `json:"origin_latitude_e_6" default:"40770034"`
	OriginLongitudeE6      int    `json:"origin_longitude_e_6" default:"-73982912"`
	DestinationLatitudeE6  int    `json:"destination_latitude_e_6"`
	DestinationLongitudeE6 int    `json:"destination_longitude_e_6"`
	OrginPlaceID           string `json:"orgin_place_id" default:"\"lyft:address:3eaa5572-4d37-4a39-92ed-c61906139955\""`
	UsingCommuterPayment   bool   `json:"using_commuter_payment"`
	Token                  string `json:"token"`
}

func (o NearbyDriversParams) Options() []NearbyDriversOption {
	return []NearbyDriversOption{
		NearbyDriversOriginLatitudeE6(o.OriginLatitudeE6),
		NearbyDriversOriginLongitudeE6(o.OriginLongitudeE6),
		NearbyDriversDestinationLatitudeE6(o.DestinationLatitudeE6),
		NearbyDriversDestinationLongitudeE6(o.DestinationLongitudeE6),
		NearbyDriversOrginPlaceID(o.OrginPlaceID),
		NearbyDriversUsingCommuterPayment(o.UsingCommuterPayment),
		NearbyDriversToken(o.Token),
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
		opt(res)
	}
	return res
}

func MakeNearbyDriversOptions(opts ...NearbyDriversOption) NearbyDriversOptions {
	return makeNearbyDriversOptionImpl(opts...)
}
