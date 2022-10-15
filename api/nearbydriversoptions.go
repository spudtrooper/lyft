// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"github.com/spudtrooper/goutil/or"
)

type NearbyDriversOption func(*nearbyDriversOptionImpl)

type NearbyDriversOptions interface {
	LatitudeE6() int
	HasLatitudeE6() bool
	LongitudeE6() int
	HasLongitudeE6() bool
	OrginPlaceID() string
	HasOrginPlaceID() bool
	UsingCommuterPayment() bool
	HasUsingCommuterPayment() bool
	Token() string
	HasToken() bool
	ToBaseOptions() []BaseOption
}

func NearbyDriversLatitudeE6(latitudeE6 int) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		opts.has_latitudeE6 = true
		opts.latitudeE6 = latitudeE6
	}
}
func NearbyDriversLatitudeE6Flag(latitudeE6 *int) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		if latitudeE6 == nil {
			return
		}
		opts.has_latitudeE6 = true
		opts.latitudeE6 = *latitudeE6
	}
}

func NearbyDriversLongitudeE6(longitudeE6 int) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		opts.has_longitudeE6 = true
		opts.longitudeE6 = longitudeE6
	}
}
func NearbyDriversLongitudeE6Flag(longitudeE6 *int) NearbyDriversOption {
	return func(opts *nearbyDriversOptionImpl) {
		if longitudeE6 == nil {
			return
		}
		opts.has_longitudeE6 = true
		opts.longitudeE6 = *longitudeE6
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
	latitudeE6               int
	has_latitudeE6           bool
	longitudeE6              int
	has_longitudeE6          bool
	orginPlaceID             string
	has_orginPlaceID         bool
	usingCommuterPayment     bool
	has_usingCommuterPayment bool
	token                    string
	has_token                bool
}

func (n *nearbyDriversOptionImpl) LatitudeE6() int      { return or.Int(n.latitudeE6, 40770034) }
func (n *nearbyDriversOptionImpl) HasLatitudeE6() bool  { return n.has_latitudeE6 }
func (n *nearbyDriversOptionImpl) LongitudeE6() int     { return or.Int(n.longitudeE6, -73982912) }
func (n *nearbyDriversOptionImpl) HasLongitudeE6() bool { return n.has_longitudeE6 }
func (n *nearbyDriversOptionImpl) OrginPlaceID() string {
	return or.String(n.orginPlaceID, "lyft:address:3eaa5572-4d37-4a39-92ed-c61906139955")
}
func (n *nearbyDriversOptionImpl) HasOrginPlaceID() bool         { return n.has_orginPlaceID }
func (n *nearbyDriversOptionImpl) UsingCommuterPayment() bool    { return n.usingCommuterPayment }
func (n *nearbyDriversOptionImpl) HasUsingCommuterPayment() bool { return n.has_usingCommuterPayment }
func (n *nearbyDriversOptionImpl) Token() string                 { return n.token }
func (n *nearbyDriversOptionImpl) HasToken() bool                { return n.has_token }

type NearbyDriversParams struct {
	LatitudeE6           int    `json:"latitude_e_6" default:"40770034"`
	LongitudeE6          int    `json:"longitude_e_6" default:"-73982912"`
	OrginPlaceID         string `json:"orgin_place_id" default:"\"lyft:address:3eaa5572-4d37-4a39-92ed-c61906139955\""`
	UsingCommuterPayment bool   `json:"using_commuter_payment"`
	Token                string `json:"token"`
}

func (o NearbyDriversParams) Options() []NearbyDriversOption {
	return []NearbyDriversOption{
		NearbyDriversLatitudeE6(o.LatitudeE6),
		NearbyDriversLongitudeE6(o.LongitudeE6),
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
