// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type AllNearbyDriversOption struct {
	f func(*allNearbyDriversOptionImpl)
	s string
}

func (o AllNearbyDriversOption) String() string { return o.s }

type AllNearbyDriversOptions interface {
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
	ToBaseOptions() []BaseOption
	ToNearbyDriversOptions() []NearbyDriversOption
}

func AllNearbyDriversDebug(debug bool) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		opts.has_debug = true
		opts.debug = debug
	}, fmt.Sprintf("api.AllNearbyDriversDebug(bool %+v)}", debug)}
}
func AllNearbyDriversDebugFlag(debug *bool) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		if debug == nil {
			return
		}
		opts.has_debug = true
		opts.debug = *debug
	}, fmt.Sprintf("api.AllNearbyDriversDebug(bool %+v)}", debug)}
}

func AllNearbyDriversDeltaE6(deltaE6 int) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		opts.has_deltaE6 = true
		opts.deltaE6 = deltaE6
	}, fmt.Sprintf("api.AllNearbyDriversDeltaE6(int %+v)}", deltaE6)}
}
func AllNearbyDriversDeltaE6Flag(deltaE6 *int) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		if deltaE6 == nil {
			return
		}
		opts.has_deltaE6 = true
		opts.deltaE6 = *deltaE6
	}, fmt.Sprintf("api.AllNearbyDriversDeltaE6(int %+v)}", deltaE6)}
}

func AllNearbyDriversDestinationLatitudeE6(destinationLatitudeE6 int) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		opts.has_destinationLatitudeE6 = true
		opts.destinationLatitudeE6 = destinationLatitudeE6
	}, fmt.Sprintf("api.AllNearbyDriversDestinationLatitudeE6(int %+v)}", destinationLatitudeE6)}
}
func AllNearbyDriversDestinationLatitudeE6Flag(destinationLatitudeE6 *int) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		if destinationLatitudeE6 == nil {
			return
		}
		opts.has_destinationLatitudeE6 = true
		opts.destinationLatitudeE6 = *destinationLatitudeE6
	}, fmt.Sprintf("api.AllNearbyDriversDestinationLatitudeE6(int %+v)}", destinationLatitudeE6)}
}

func AllNearbyDriversDestinationLongitudeE6(destinationLongitudeE6 int) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		opts.has_destinationLongitudeE6 = true
		opts.destinationLongitudeE6 = destinationLongitudeE6
	}, fmt.Sprintf("api.AllNearbyDriversDestinationLongitudeE6(int %+v)}", destinationLongitudeE6)}
}
func AllNearbyDriversDestinationLongitudeE6Flag(destinationLongitudeE6 *int) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		if destinationLongitudeE6 == nil {
			return
		}
		opts.has_destinationLongitudeE6 = true
		opts.destinationLongitudeE6 = *destinationLongitudeE6
	}, fmt.Sprintf("api.AllNearbyDriversDestinationLongitudeE6(int %+v)}", destinationLongitudeE6)}
}

func AllNearbyDriversMultiples(multiples int) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		opts.has_multiples = true
		opts.multiples = multiples
	}, fmt.Sprintf("api.AllNearbyDriversMultiples(int %+v)}", multiples)}
}
func AllNearbyDriversMultiplesFlag(multiples *int) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		if multiples == nil {
			return
		}
		opts.has_multiples = true
		opts.multiples = *multiples
	}, fmt.Sprintf("api.AllNearbyDriversMultiples(int %+v)}", multiples)}
}

func AllNearbyDriversOrginPlaceID(orginPlaceID string) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		opts.has_orginPlaceID = true
		opts.orginPlaceID = orginPlaceID
	}, fmt.Sprintf("api.AllNearbyDriversOrginPlaceID(string %+v)}", orginPlaceID)}
}
func AllNearbyDriversOrginPlaceIDFlag(orginPlaceID *string) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		if orginPlaceID == nil {
			return
		}
		opts.has_orginPlaceID = true
		opts.orginPlaceID = *orginPlaceID
	}, fmt.Sprintf("api.AllNearbyDriversOrginPlaceID(string %+v)}", orginPlaceID)}
}

func AllNearbyDriversOriginLatitudeE6(originLatitudeE6 int) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		opts.has_originLatitudeE6 = true
		opts.originLatitudeE6 = originLatitudeE6
	}, fmt.Sprintf("api.AllNearbyDriversOriginLatitudeE6(int %+v)}", originLatitudeE6)}
}
func AllNearbyDriversOriginLatitudeE6Flag(originLatitudeE6 *int) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		if originLatitudeE6 == nil {
			return
		}
		opts.has_originLatitudeE6 = true
		opts.originLatitudeE6 = *originLatitudeE6
	}, fmt.Sprintf("api.AllNearbyDriversOriginLatitudeE6(int %+v)}", originLatitudeE6)}
}

func AllNearbyDriversOriginLongitudeE6(originLongitudeE6 int) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		opts.has_originLongitudeE6 = true
		opts.originLongitudeE6 = originLongitudeE6
	}, fmt.Sprintf("api.AllNearbyDriversOriginLongitudeE6(int %+v)}", originLongitudeE6)}
}
func AllNearbyDriversOriginLongitudeE6Flag(originLongitudeE6 *int) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		if originLongitudeE6 == nil {
			return
		}
		opts.has_originLongitudeE6 = true
		opts.originLongitudeE6 = *originLongitudeE6
	}, fmt.Sprintf("api.AllNearbyDriversOriginLongitudeE6(int %+v)}", originLongitudeE6)}
}

func AllNearbyDriversThreads(threads int) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		opts.has_threads = true
		opts.threads = threads
	}, fmt.Sprintf("api.AllNearbyDriversThreads(int %+v)}", threads)}
}
func AllNearbyDriversThreadsFlag(threads *int) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		if threads == nil {
			return
		}
		opts.has_threads = true
		opts.threads = *threads
	}, fmt.Sprintf("api.AllNearbyDriversThreads(int %+v)}", threads)}
}

func AllNearbyDriversToken(token string) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		opts.has_token = true
		opts.token = token
	}, fmt.Sprintf("api.AllNearbyDriversToken(string %+v)}", token)}
}
func AllNearbyDriversTokenFlag(token *string) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}, fmt.Sprintf("api.AllNearbyDriversToken(string %+v)}", token)}
}

func AllNearbyDriversUsingCommuterPayment(usingCommuterPayment bool) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		opts.has_usingCommuterPayment = true
		opts.usingCommuterPayment = usingCommuterPayment
	}, fmt.Sprintf("api.AllNearbyDriversUsingCommuterPayment(bool %+v)}", usingCommuterPayment)}
}
func AllNearbyDriversUsingCommuterPaymentFlag(usingCommuterPayment *bool) AllNearbyDriversOption {
	return AllNearbyDriversOption{func(opts *allNearbyDriversOptionImpl) {
		if usingCommuterPayment == nil {
			return
		}
		opts.has_usingCommuterPayment = true
		opts.usingCommuterPayment = *usingCommuterPayment
	}, fmt.Sprintf("api.AllNearbyDriversUsingCommuterPayment(bool %+v)}", usingCommuterPayment)}
}

type allNearbyDriversOptionImpl struct {
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

func (a *allNearbyDriversOptionImpl) Debug() bool                { return a.debug }
func (a *allNearbyDriversOptionImpl) HasDebug() bool             { return a.has_debug }
func (a *allNearbyDriversOptionImpl) DeltaE6() int               { return or.Int(a.deltaE6, 130) }
func (a *allNearbyDriversOptionImpl) HasDeltaE6() bool           { return a.has_deltaE6 }
func (a *allNearbyDriversOptionImpl) DestinationLatitudeE6() int { return a.destinationLatitudeE6 }
func (a *allNearbyDriversOptionImpl) HasDestinationLatitudeE6() bool {
	return a.has_destinationLatitudeE6
}
func (a *allNearbyDriversOptionImpl) DestinationLongitudeE6() int { return a.destinationLongitudeE6 }
func (a *allNearbyDriversOptionImpl) HasDestinationLongitudeE6() bool {
	return a.has_destinationLongitudeE6
}
func (a *allNearbyDriversOptionImpl) Multiples() int     { return or.Int(a.multiples, 1) }
func (a *allNearbyDriversOptionImpl) HasMultiples() bool { return a.has_multiples }
func (a *allNearbyDriversOptionImpl) OrginPlaceID() string {
	return or.String(a.orginPlaceID, "lyft:address:3eaa5572-4d37-4a39-92ed-c61906139955")
}
func (a *allNearbyDriversOptionImpl) HasOrginPlaceID() bool { return a.has_orginPlaceID }
func (a *allNearbyDriversOptionImpl) OriginLatitudeE6() int {
	return or.Int(a.originLatitudeE6, 40770034)
}
func (a *allNearbyDriversOptionImpl) HasOriginLatitudeE6() bool { return a.has_originLatitudeE6 }
func (a *allNearbyDriversOptionImpl) OriginLongitudeE6() int {
	return or.Int(a.originLongitudeE6, -73982912)
}
func (a *allNearbyDriversOptionImpl) HasOriginLongitudeE6() bool { return a.has_originLongitudeE6 }
func (a *allNearbyDriversOptionImpl) Threads() int               { return or.Int(a.threads, 5) }
func (a *allNearbyDriversOptionImpl) HasThreads() bool           { return a.has_threads }
func (a *allNearbyDriversOptionImpl) Token() string              { return a.token }
func (a *allNearbyDriversOptionImpl) HasToken() bool             { return a.has_token }
func (a *allNearbyDriversOptionImpl) UsingCommuterPayment() bool { return a.usingCommuterPayment }
func (a *allNearbyDriversOptionImpl) HasUsingCommuterPayment() bool {
	return a.has_usingCommuterPayment
}

type AllNearbyDriversParams struct {
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

func (o AllNearbyDriversParams) Options() []AllNearbyDriversOption {
	return []AllNearbyDriversOption{
		AllNearbyDriversDebug(o.Debug),
		AllNearbyDriversDeltaE6(o.DeltaE6),
		AllNearbyDriversDestinationLatitudeE6(o.DestinationLatitudeE6),
		AllNearbyDriversDestinationLongitudeE6(o.DestinationLongitudeE6),
		AllNearbyDriversMultiples(o.Multiples),
		AllNearbyDriversOrginPlaceID(o.OrginPlaceID),
		AllNearbyDriversOriginLatitudeE6(o.OriginLatitudeE6),
		AllNearbyDriversOriginLongitudeE6(o.OriginLongitudeE6),
		AllNearbyDriversThreads(o.Threads),
		AllNearbyDriversToken(o.Token),
		AllNearbyDriversUsingCommuterPayment(o.UsingCommuterPayment),
	}
}

// ToBaseOptions converts AllNearbyDriversOption to an array of BaseOption
func (o *allNearbyDriversOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
	}
}

// ToNearbyDriversOptions converts AllNearbyDriversOption to an array of NearbyDriversOption
func (o *allNearbyDriversOptionImpl) ToNearbyDriversOptions() []NearbyDriversOption {
	return []NearbyDriversOption{
		NearbyDriversOriginLongitudeE6(o.OriginLongitudeE6()),
		NearbyDriversDestinationLatitudeE6(o.DestinationLatitudeE6()),
		NearbyDriversDestinationLongitudeE6(o.DestinationLongitudeE6()),
		NearbyDriversOrginPlaceID(o.OrginPlaceID()),
		NearbyDriversUsingCommuterPayment(o.UsingCommuterPayment()),
		NearbyDriversToken(o.Token()),
		NearbyDriversOriginLatitudeE6(o.OriginLatitudeE6()),
	}
}

func makeAllNearbyDriversOptionImpl(opts ...AllNearbyDriversOption) *allNearbyDriversOptionImpl {
	res := &allNearbyDriversOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeAllNearbyDriversOptions(opts ...AllNearbyDriversOption) AllNearbyDriversOptions {
	return makeAllNearbyDriversOptionImpl(opts...)
}
