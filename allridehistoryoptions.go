// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package lyft

import "fmt"

type AllRideHistoryOption struct {
	f func(*allRideHistoryOptionImpl)
	s string
}

func (o AllRideHistoryOption) String() string { return o.s }

type AllRideHistoryOptions interface {
	Debug() bool
	HasDebug() bool
}

func AllRideHistoryDebug(debug bool) AllRideHistoryOption {
	return AllRideHistoryOption{func(opts *allRideHistoryOptionImpl) {
		opts.has_debug = true
		opts.debug = debug
	}, fmt.Sprintf("lyft.AllRideHistoryDebug(bool %+v)}", debug)}
}
func AllRideHistoryDebugFlag(debug *bool) AllRideHistoryOption {
	return AllRideHistoryOption{func(opts *allRideHistoryOptionImpl) {
		if debug == nil {
			return
		}
		opts.has_debug = true
		opts.debug = *debug
	}, fmt.Sprintf("lyft.AllRideHistoryDebug(bool %+v)}", debug)}
}

type allRideHistoryOptionImpl struct {
	debug     bool
	has_debug bool
}

func (a *allRideHistoryOptionImpl) Debug() bool    { return a.debug }
func (a *allRideHistoryOptionImpl) HasDebug() bool { return a.has_debug }

type AllRideHistoryParams struct {
	Debug bool `json:"debug"`
}

func (o AllRideHistoryParams) Options() []AllRideHistoryOption {
	return []AllRideHistoryOption{
		AllRideHistoryDebug(o.Debug),
	}
}

func makeAllRideHistoryOptionImpl(opts ...AllRideHistoryOption) *allRideHistoryOptionImpl {
	res := &allRideHistoryOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeAllRideHistoryOptions(opts ...AllRideHistoryOption) AllRideHistoryOptions {
	return makeAllRideHistoryOptionImpl(opts...)
}
