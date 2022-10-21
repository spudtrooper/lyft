// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package render

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type AllVehicleViewsMapOption struct {
	f func(*allVehicleViewsMapOptionImpl)
	s string
}

func (o AllVehicleViewsMapOption) String() string { return o.s }

type AllVehicleViewsMapOptions interface {
	DeltaE6() int
	HasDeltaE6() bool
	Latitude() float64
	HasLatitude() bool
	Longitude() float64
	HasLongitude() bool
	Multiples() int
	HasMultiples() bool
	ShowCenters() bool
	HasShowCenters() bool
	SleepSecs() int
	HasSleepSecs() bool
	Zoom() int
	HasZoom() bool
	ToVehicleViewsMapOptions() []VehicleViewsMapOption
}

func AllVehicleViewsMapDeltaE6(deltaE6 int) AllVehicleViewsMapOption {
	return AllVehicleViewsMapOption{func(opts *allVehicleViewsMapOptionImpl) {
		opts.has_deltaE6 = true
		opts.deltaE6 = deltaE6
	}, fmt.Sprintf("render.AllVehicleViewsMapDeltaE6(int %+v)}", deltaE6)}
}
func AllVehicleViewsMapDeltaE6Flag(deltaE6 *int) AllVehicleViewsMapOption {
	return AllVehicleViewsMapOption{func(opts *allVehicleViewsMapOptionImpl) {
		if deltaE6 == nil {
			return
		}
		opts.has_deltaE6 = true
		opts.deltaE6 = *deltaE6
	}, fmt.Sprintf("render.AllVehicleViewsMapDeltaE6(int %+v)}", deltaE6)}
}

func AllVehicleViewsMapLatitude(latitude float64) AllVehicleViewsMapOption {
	return AllVehicleViewsMapOption{func(opts *allVehicleViewsMapOptionImpl) {
		opts.has_latitude = true
		opts.latitude = latitude
	}, fmt.Sprintf("render.AllVehicleViewsMapLatitude(float64 %+v)}", latitude)}
}
func AllVehicleViewsMapLatitudeFlag(latitude *float64) AllVehicleViewsMapOption {
	return AllVehicleViewsMapOption{func(opts *allVehicleViewsMapOptionImpl) {
		if latitude == nil {
			return
		}
		opts.has_latitude = true
		opts.latitude = *latitude
	}, fmt.Sprintf("render.AllVehicleViewsMapLatitude(float64 %+v)}", latitude)}
}

func AllVehicleViewsMapLongitude(longitude float64) AllVehicleViewsMapOption {
	return AllVehicleViewsMapOption{func(opts *allVehicleViewsMapOptionImpl) {
		opts.has_longitude = true
		opts.longitude = longitude
	}, fmt.Sprintf("render.AllVehicleViewsMapLongitude(float64 %+v)}", longitude)}
}
func AllVehicleViewsMapLongitudeFlag(longitude *float64) AllVehicleViewsMapOption {
	return AllVehicleViewsMapOption{func(opts *allVehicleViewsMapOptionImpl) {
		if longitude == nil {
			return
		}
		opts.has_longitude = true
		opts.longitude = *longitude
	}, fmt.Sprintf("render.AllVehicleViewsMapLongitude(float64 %+v)}", longitude)}
}

func AllVehicleViewsMapMultiples(multiples int) AllVehicleViewsMapOption {
	return AllVehicleViewsMapOption{func(opts *allVehicleViewsMapOptionImpl) {
		opts.has_multiples = true
		opts.multiples = multiples
	}, fmt.Sprintf("render.AllVehicleViewsMapMultiples(int %+v)}", multiples)}
}
func AllVehicleViewsMapMultiplesFlag(multiples *int) AllVehicleViewsMapOption {
	return AllVehicleViewsMapOption{func(opts *allVehicleViewsMapOptionImpl) {
		if multiples == nil {
			return
		}
		opts.has_multiples = true
		opts.multiples = *multiples
	}, fmt.Sprintf("render.AllVehicleViewsMapMultiples(int %+v)}", multiples)}
}

func AllVehicleViewsMapShowCenters(showCenters bool) AllVehicleViewsMapOption {
	return AllVehicleViewsMapOption{func(opts *allVehicleViewsMapOptionImpl) {
		opts.has_showCenters = true
		opts.showCenters = showCenters
	}, fmt.Sprintf("render.AllVehicleViewsMapShowCenters(bool %+v)}", showCenters)}
}
func AllVehicleViewsMapShowCentersFlag(showCenters *bool) AllVehicleViewsMapOption {
	return AllVehicleViewsMapOption{func(opts *allVehicleViewsMapOptionImpl) {
		if showCenters == nil {
			return
		}
		opts.has_showCenters = true
		opts.showCenters = *showCenters
	}, fmt.Sprintf("render.AllVehicleViewsMapShowCenters(bool %+v)}", showCenters)}
}

func AllVehicleViewsMapSleepSecs(sleepSecs int) AllVehicleViewsMapOption {
	return AllVehicleViewsMapOption{func(opts *allVehicleViewsMapOptionImpl) {
		opts.has_sleepSecs = true
		opts.sleepSecs = sleepSecs
	}, fmt.Sprintf("render.AllVehicleViewsMapSleepSecs(int %+v)}", sleepSecs)}
}
func AllVehicleViewsMapSleepSecsFlag(sleepSecs *int) AllVehicleViewsMapOption {
	return AllVehicleViewsMapOption{func(opts *allVehicleViewsMapOptionImpl) {
		if sleepSecs == nil {
			return
		}
		opts.has_sleepSecs = true
		opts.sleepSecs = *sleepSecs
	}, fmt.Sprintf("render.AllVehicleViewsMapSleepSecs(int %+v)}", sleepSecs)}
}

func AllVehicleViewsMapZoom(zoom int) AllVehicleViewsMapOption {
	return AllVehicleViewsMapOption{func(opts *allVehicleViewsMapOptionImpl) {
		opts.has_zoom = true
		opts.zoom = zoom
	}, fmt.Sprintf("render.AllVehicleViewsMapZoom(int %+v)}", zoom)}
}
func AllVehicleViewsMapZoomFlag(zoom *int) AllVehicleViewsMapOption {
	return AllVehicleViewsMapOption{func(opts *allVehicleViewsMapOptionImpl) {
		if zoom == nil {
			return
		}
		opts.has_zoom = true
		opts.zoom = *zoom
	}, fmt.Sprintf("render.AllVehicleViewsMapZoom(int %+v)}", zoom)}
}

type allVehicleViewsMapOptionImpl struct {
	deltaE6         int
	has_deltaE6     bool
	latitude        float64
	has_latitude    bool
	longitude       float64
	has_longitude   bool
	multiples       int
	has_multiples   bool
	showCenters     bool
	has_showCenters bool
	sleepSecs       int
	has_sleepSecs   bool
	zoom            int
	has_zoom        bool
}

func (a *allVehicleViewsMapOptionImpl) DeltaE6() int      { return or.Int(a.deltaE6, 9000) }
func (a *allVehicleViewsMapOptionImpl) HasDeltaE6() bool  { return a.has_deltaE6 }
func (a *allVehicleViewsMapOptionImpl) Latitude() float64 { return or.Float64(a.latitude, 40.7701286) }
func (a *allVehicleViewsMapOptionImpl) HasLatitude() bool { return a.has_latitude }
func (a *allVehicleViewsMapOptionImpl) Longitude() float64 {
	return or.Float64(a.longitude, -73.9829762)
}
func (a *allVehicleViewsMapOptionImpl) HasLongitude() bool   { return a.has_longitude }
func (a *allVehicleViewsMapOptionImpl) Multiples() int       { return or.Int(a.multiples, 1) }
func (a *allVehicleViewsMapOptionImpl) HasMultiples() bool   { return a.has_multiples }
func (a *allVehicleViewsMapOptionImpl) ShowCenters() bool    { return a.showCenters }
func (a *allVehicleViewsMapOptionImpl) HasShowCenters() bool { return a.has_showCenters }
func (a *allVehicleViewsMapOptionImpl) SleepSecs() int       { return or.Int(a.sleepSecs, 0) }
func (a *allVehicleViewsMapOptionImpl) HasSleepSecs() bool   { return a.has_sleepSecs }
func (a *allVehicleViewsMapOptionImpl) Zoom() int            { return or.Int(a.zoom, 14) }
func (a *allVehicleViewsMapOptionImpl) HasZoom() bool        { return a.has_zoom }

type AllVehicleViewsMapParams struct {
	DeltaE6     int     `json:"delta_e_6" default:"9000"`
	Latitude    float64 `json:"latitude" default:"40.7701286"`
	Longitude   float64 `json:"longitude" default:"-73.9829762"`
	Multiples   int     `json:"multiples" default:"1"`
	ShowCenters bool    `json:"show_centers"`
	SleepSecs   int     `json:"sleep_secs" default:"0"`
	Token       string  `json:"token" required:"true"`
	Zoom        int     `json:"zoom" default:"14"`
}

func (o AllVehicleViewsMapParams) Options() []AllVehicleViewsMapOption {
	return []AllVehicleViewsMapOption{
		AllVehicleViewsMapDeltaE6(o.DeltaE6),
		AllVehicleViewsMapLatitude(o.Latitude),
		AllVehicleViewsMapLongitude(o.Longitude),
		AllVehicleViewsMapMultiples(o.Multiples),
		AllVehicleViewsMapShowCenters(o.ShowCenters),
		AllVehicleViewsMapSleepSecs(o.SleepSecs),
		AllVehicleViewsMapZoom(o.Zoom),
	}
}

// ToVehicleViewsMapOptions converts AllVehicleViewsMapOption to an array of VehicleViewsMapOption
func (o *allVehicleViewsMapOptionImpl) ToVehicleViewsMapOptions() []VehicleViewsMapOption {
	return []VehicleViewsMapOption{
		VehicleViewsMapLatitude(o.Latitude()),
		VehicleViewsMapLongitude(o.Longitude()),
		VehicleViewsMapShowCenters(o.ShowCenters()),
		VehicleViewsMapSleepSecs(o.SleepSecs()),
		VehicleViewsMapZoom(o.Zoom()),
	}
}

func makeAllVehicleViewsMapOptionImpl(opts ...AllVehicleViewsMapOption) *allVehicleViewsMapOptionImpl {
	res := &allVehicleViewsMapOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeAllVehicleViewsMapOptions(opts ...AllVehicleViewsMapOption) AllVehicleViewsMapOptions {
	return makeAllVehicleViewsMapOptionImpl(opts...)
}
