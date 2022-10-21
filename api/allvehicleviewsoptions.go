// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type AllVehicleViewsOption struct {
	f func(*allVehicleViewsOptionImpl)
	s string
}

func (o AllVehicleViewsOption) String() string { return o.s }

type AllVehicleViewsOptions interface {
	Debug() bool
	HasDebug() bool
	DeltaE6() int
	HasDeltaE6() bool
	DestinationLatitudeE6() int
	HasDestinationLatitudeE6() bool
	DestinationLongitudeE6() int
	HasDestinationLongitudeE6() bool
	Multiples() int
	HasMultiples() bool
	OrginPlaceID() string
	HasOrginPlaceID() bool
	OriginLatitudeE6() int
	HasOriginLatitudeE6() bool
	OriginLongitudeE6() int
	HasOriginLongitudeE6() bool
	Threads() int
	HasThreads() bool
	Token() string
	HasToken() bool
	UsingCommuterPayment() bool
	HasUsingCommuterPayment() bool
	ToAllNearbyDriversOptions() []AllNearbyDriversOption
	ToBaseOptions() []BaseOption
	ToNearbyDriversOptions() []NearbyDriversOption
}

func AllVehicleViewsDebug(debug bool) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		opts.has_debug = true
		opts.debug = debug
	}, fmt.Sprintf("api.AllVehicleViewsDebug(bool %+v)}", debug)}
}
func AllVehicleViewsDebugFlag(debug *bool) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		if debug == nil {
			return
		}
		opts.has_debug = true
		opts.debug = *debug
	}, fmt.Sprintf("api.AllVehicleViewsDebug(bool %+v)}", debug)}
}

func AllVehicleViewsDeltaE6(deltaE6 int) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		opts.has_deltaE6 = true
		opts.deltaE6 = deltaE6
	}, fmt.Sprintf("api.AllVehicleViewsDeltaE6(int %+v)}", deltaE6)}
}
func AllVehicleViewsDeltaE6Flag(deltaE6 *int) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		if deltaE6 == nil {
			return
		}
		opts.has_deltaE6 = true
		opts.deltaE6 = *deltaE6
	}, fmt.Sprintf("api.AllVehicleViewsDeltaE6(int %+v)}", deltaE6)}
}

func AllVehicleViewsDestinationLatitudeE6(destinationLatitudeE6 int) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		opts.has_destinationLatitudeE6 = true
		opts.destinationLatitudeE6 = destinationLatitudeE6
	}, fmt.Sprintf("api.AllVehicleViewsDestinationLatitudeE6(int %+v)}", destinationLatitudeE6)}
}
func AllVehicleViewsDestinationLatitudeE6Flag(destinationLatitudeE6 *int) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		if destinationLatitudeE6 == nil {
			return
		}
		opts.has_destinationLatitudeE6 = true
		opts.destinationLatitudeE6 = *destinationLatitudeE6
	}, fmt.Sprintf("api.AllVehicleViewsDestinationLatitudeE6(int %+v)}", destinationLatitudeE6)}
}

func AllVehicleViewsDestinationLongitudeE6(destinationLongitudeE6 int) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		opts.has_destinationLongitudeE6 = true
		opts.destinationLongitudeE6 = destinationLongitudeE6
	}, fmt.Sprintf("api.AllVehicleViewsDestinationLongitudeE6(int %+v)}", destinationLongitudeE6)}
}
func AllVehicleViewsDestinationLongitudeE6Flag(destinationLongitudeE6 *int) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		if destinationLongitudeE6 == nil {
			return
		}
		opts.has_destinationLongitudeE6 = true
		opts.destinationLongitudeE6 = *destinationLongitudeE6
	}, fmt.Sprintf("api.AllVehicleViewsDestinationLongitudeE6(int %+v)}", destinationLongitudeE6)}
}

func AllVehicleViewsMultiples(multiples int) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		opts.has_multiples = true
		opts.multiples = multiples
	}, fmt.Sprintf("api.AllVehicleViewsMultiples(int %+v)}", multiples)}
}
func AllVehicleViewsMultiplesFlag(multiples *int) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		if multiples == nil {
			return
		}
		opts.has_multiples = true
		opts.multiples = *multiples
	}, fmt.Sprintf("api.AllVehicleViewsMultiples(int %+v)}", multiples)}
}

func AllVehicleViewsOrginPlaceID(orginPlaceID string) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		opts.has_orginPlaceID = true
		opts.orginPlaceID = orginPlaceID
	}, fmt.Sprintf("api.AllVehicleViewsOrginPlaceID(string %+v)}", orginPlaceID)}
}
func AllVehicleViewsOrginPlaceIDFlag(orginPlaceID *string) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		if orginPlaceID == nil {
			return
		}
		opts.has_orginPlaceID = true
		opts.orginPlaceID = *orginPlaceID
	}, fmt.Sprintf("api.AllVehicleViewsOrginPlaceID(string %+v)}", orginPlaceID)}
}

func AllVehicleViewsOriginLatitudeE6(originLatitudeE6 int) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		opts.has_originLatitudeE6 = true
		opts.originLatitudeE6 = originLatitudeE6
	}, fmt.Sprintf("api.AllVehicleViewsOriginLatitudeE6(int %+v)}", originLatitudeE6)}
}
func AllVehicleViewsOriginLatitudeE6Flag(originLatitudeE6 *int) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		if originLatitudeE6 == nil {
			return
		}
		opts.has_originLatitudeE6 = true
		opts.originLatitudeE6 = *originLatitudeE6
	}, fmt.Sprintf("api.AllVehicleViewsOriginLatitudeE6(int %+v)}", originLatitudeE6)}
}

func AllVehicleViewsOriginLongitudeE6(originLongitudeE6 int) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		opts.has_originLongitudeE6 = true
		opts.originLongitudeE6 = originLongitudeE6
	}, fmt.Sprintf("api.AllVehicleViewsOriginLongitudeE6(int %+v)}", originLongitudeE6)}
}
func AllVehicleViewsOriginLongitudeE6Flag(originLongitudeE6 *int) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		if originLongitudeE6 == nil {
			return
		}
		opts.has_originLongitudeE6 = true
		opts.originLongitudeE6 = *originLongitudeE6
	}, fmt.Sprintf("api.AllVehicleViewsOriginLongitudeE6(int %+v)}", originLongitudeE6)}
}

func AllVehicleViewsThreads(threads int) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		opts.has_threads = true
		opts.threads = threads
	}, fmt.Sprintf("api.AllVehicleViewsThreads(int %+v)}", threads)}
}
func AllVehicleViewsThreadsFlag(threads *int) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		if threads == nil {
			return
		}
		opts.has_threads = true
		opts.threads = *threads
	}, fmt.Sprintf("api.AllVehicleViewsThreads(int %+v)}", threads)}
}

func AllVehicleViewsToken(token string) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		opts.has_token = true
		opts.token = token
	}, fmt.Sprintf("api.AllVehicleViewsToken(string %+v)}", token)}
}
func AllVehicleViewsTokenFlag(token *string) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}, fmt.Sprintf("api.AllVehicleViewsToken(string %+v)}", token)}
}

func AllVehicleViewsUsingCommuterPayment(usingCommuterPayment bool) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		opts.has_usingCommuterPayment = true
		opts.usingCommuterPayment = usingCommuterPayment
	}, fmt.Sprintf("api.AllVehicleViewsUsingCommuterPayment(bool %+v)}", usingCommuterPayment)}
}
func AllVehicleViewsUsingCommuterPaymentFlag(usingCommuterPayment *bool) AllVehicleViewsOption {
	return AllVehicleViewsOption{func(opts *allVehicleViewsOptionImpl) {
		if usingCommuterPayment == nil {
			return
		}
		opts.has_usingCommuterPayment = true
		opts.usingCommuterPayment = *usingCommuterPayment
	}, fmt.Sprintf("api.AllVehicleViewsUsingCommuterPayment(bool %+v)}", usingCommuterPayment)}
}

type allVehicleViewsOptionImpl struct {
	debug                      bool
	has_debug                  bool
	deltaE6                    int
	has_deltaE6                bool
	destinationLatitudeE6      int
	has_destinationLatitudeE6  bool
	destinationLongitudeE6     int
	has_destinationLongitudeE6 bool
	multiples                  int
	has_multiples              bool
	orginPlaceID               string
	has_orginPlaceID           bool
	originLatitudeE6           int
	has_originLatitudeE6       bool
	originLongitudeE6          int
	has_originLongitudeE6      bool
	threads                    int
	has_threads                bool
	token                      string
	has_token                  bool
	usingCommuterPayment       bool
	has_usingCommuterPayment   bool
}

func (a *allVehicleViewsOptionImpl) Debug() bool                { return a.debug }
func (a *allVehicleViewsOptionImpl) HasDebug() bool             { return a.has_debug }
func (a *allVehicleViewsOptionImpl) DeltaE6() int               { return or.Int(a.deltaE6, 130) }
func (a *allVehicleViewsOptionImpl) HasDeltaE6() bool           { return a.has_deltaE6 }
func (a *allVehicleViewsOptionImpl) DestinationLatitudeE6() int { return a.destinationLatitudeE6 }
func (a *allVehicleViewsOptionImpl) HasDestinationLatitudeE6() bool {
	return a.has_destinationLatitudeE6
}
func (a *allVehicleViewsOptionImpl) DestinationLongitudeE6() int { return a.destinationLongitudeE6 }
func (a *allVehicleViewsOptionImpl) HasDestinationLongitudeE6() bool {
	return a.has_destinationLongitudeE6
}
func (a *allVehicleViewsOptionImpl) Multiples() int     { return or.Int(a.multiples, 1) }
func (a *allVehicleViewsOptionImpl) HasMultiples() bool { return a.has_multiples }
func (a *allVehicleViewsOptionImpl) OrginPlaceID() string {
	return or.String(a.orginPlaceID, "lyft:address:3eaa5572-4d37-4a39-92ed-c61906139955")
}
func (a *allVehicleViewsOptionImpl) HasOrginPlaceID() bool { return a.has_orginPlaceID }
func (a *allVehicleViewsOptionImpl) OriginLatitudeE6() int {
	return or.Int(a.originLatitudeE6, 40770034)
}
func (a *allVehicleViewsOptionImpl) HasOriginLatitudeE6() bool { return a.has_originLatitudeE6 }
func (a *allVehicleViewsOptionImpl) OriginLongitudeE6() int {
	return or.Int(a.originLongitudeE6, -73982912)
}
func (a *allVehicleViewsOptionImpl) HasOriginLongitudeE6() bool    { return a.has_originLongitudeE6 }
func (a *allVehicleViewsOptionImpl) Threads() int                  { return or.Int(a.threads, 5) }
func (a *allVehicleViewsOptionImpl) HasThreads() bool              { return a.has_threads }
func (a *allVehicleViewsOptionImpl) Token() string                 { return a.token }
func (a *allVehicleViewsOptionImpl) HasToken() bool                { return a.has_token }
func (a *allVehicleViewsOptionImpl) UsingCommuterPayment() bool    { return a.usingCommuterPayment }
func (a *allVehicleViewsOptionImpl) HasUsingCommuterPayment() bool { return a.has_usingCommuterPayment }

type AllVehicleViewsParams struct {
	Debug                  bool   `json:"debug"`
	DeltaE6                int    `json:"delta_e_6" default:"130"`
	DestinationLatitudeE6  int    `json:"destination_latitude_e_6"`
	DestinationLongitudeE6 int    `json:"destination_longitude_e_6"`
	Multiples              int    `json:"multiples" default:"1"`
	OrginPlaceID           string `json:"orgin_place_id" default:"\"lyft:address:3eaa5572-4d37-4a39-92ed-c61906139955\""`
	OriginLatitudeE6       int    `json:"origin_latitude_e_6" default:"40770034"`
	OriginLongitudeE6      int    `json:"origin_longitude_e_6" default:"-73982912"`
	Threads                int    `json:"threads" default:"5"`
	Token                  string `json:"token"`
	UsingCommuterPayment   bool   `json:"using_commuter_payment"`
}

func (o AllVehicleViewsParams) Options() []AllVehicleViewsOption {
	return []AllVehicleViewsOption{
		AllVehicleViewsDebug(o.Debug),
		AllVehicleViewsDeltaE6(o.DeltaE6),
		AllVehicleViewsDestinationLatitudeE6(o.DestinationLatitudeE6),
		AllVehicleViewsDestinationLongitudeE6(o.DestinationLongitudeE6),
		AllVehicleViewsMultiples(o.Multiples),
		AllVehicleViewsOrginPlaceID(o.OrginPlaceID),
		AllVehicleViewsOriginLatitudeE6(o.OriginLatitudeE6),
		AllVehicleViewsOriginLongitudeE6(o.OriginLongitudeE6),
		AllVehicleViewsThreads(o.Threads),
		AllVehicleViewsToken(o.Token),
		AllVehicleViewsUsingCommuterPayment(o.UsingCommuterPayment),
	}
}

// ToAllNearbyDriversOptions converts AllVehicleViewsOption to an array of AllNearbyDriversOption
func (o *allVehicleViewsOptionImpl) ToAllNearbyDriversOptions() []AllNearbyDriversOption {
	return []AllNearbyDriversOption{
		AllNearbyDriversDeltaE6(o.DeltaE6()),
		AllNearbyDriversMultiples(o.Multiples()),
		AllNearbyDriversThreads(o.Threads()),
		AllNearbyDriversDebug(o.Debug()),
		AllNearbyDriversToken(o.Token()),
		AllNearbyDriversOriginLatitudeE6(o.OriginLatitudeE6()),
		AllNearbyDriversOriginLongitudeE6(o.OriginLongitudeE6()),
		AllNearbyDriversDestinationLatitudeE6(o.DestinationLatitudeE6()),
		AllNearbyDriversDestinationLongitudeE6(o.DestinationLongitudeE6()),
		AllNearbyDriversOrginPlaceID(o.OrginPlaceID()),
		AllNearbyDriversUsingCommuterPayment(o.UsingCommuterPayment()),
	}
}

// ToBaseOptions converts AllVehicleViewsOption to an array of BaseOption
func (o *allVehicleViewsOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
	}
}

// ToNearbyDriversOptions converts AllVehicleViewsOption to an array of NearbyDriversOption
func (o *allVehicleViewsOptionImpl) ToNearbyDriversOptions() []NearbyDriversOption {
	return []NearbyDriversOption{
		NearbyDriversUsingCommuterPayment(o.UsingCommuterPayment()),
		NearbyDriversToken(o.Token()),
		NearbyDriversOriginLatitudeE6(o.OriginLatitudeE6()),
		NearbyDriversOriginLongitudeE6(o.OriginLongitudeE6()),
		NearbyDriversDestinationLatitudeE6(o.DestinationLatitudeE6()),
		NearbyDriversDestinationLongitudeE6(o.DestinationLongitudeE6()),
		NearbyDriversOrginPlaceID(o.OrginPlaceID()),
	}
}

func makeAllVehicleViewsOptionImpl(opts ...AllVehicleViewsOption) *allVehicleViewsOptionImpl {
	res := &allVehicleViewsOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeAllVehicleViewsOptions(opts ...AllVehicleViewsOption) AllVehicleViewsOptions {
	return makeAllVehicleViewsOptionImpl(opts...)
}
