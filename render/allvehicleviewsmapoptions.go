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
	Latitude() float64
	HasLatitude() bool
	Longitude() float64
	HasLongitude() bool
	Multiple() int
	HasMultiple() bool
	SleepSecs() int
	HasSleepSecs() bool
	Zoom() int
	HasZoom() bool
	ToVehicleViewsMapOptions() []VehicleViewsMapOption
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

func AllVehicleViewsMapMultiple(multiple int) AllVehicleViewsMapOption {
	return AllVehicleViewsMapOption{func(opts *allVehicleViewsMapOptionImpl) {
		opts.has_multiple = true
		opts.multiple = multiple
	}, fmt.Sprintf("render.AllVehicleViewsMapMultiple(int %+v)}", multiple)}
}
func AllVehicleViewsMapMultipleFlag(multiple *int) AllVehicleViewsMapOption {
	return AllVehicleViewsMapOption{func(opts *allVehicleViewsMapOptionImpl) {
		if multiple == nil {
			return
		}
		opts.has_multiple = true
		opts.multiple = *multiple
	}, fmt.Sprintf("render.AllVehicleViewsMapMultiple(int %+v)}", multiple)}
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
	latitude      float64
	has_latitude  bool
	longitude     float64
	has_longitude bool
	multiple      int
	has_multiple  bool
	sleepSecs     int
	has_sleepSecs bool
	zoom          int
	has_zoom      bool
}

func (a *allVehicleViewsMapOptionImpl) Latitude() float64 { return or.Float64(a.latitude, 40.7701286) }
func (a *allVehicleViewsMapOptionImpl) HasLatitude() bool { return a.has_latitude }
func (a *allVehicleViewsMapOptionImpl) Longitude() float64 {
	return or.Float64(a.longitude, -73.9829762)
}
func (a *allVehicleViewsMapOptionImpl) HasLongitude() bool { return a.has_longitude }
func (a *allVehicleViewsMapOptionImpl) Multiple() int      { return or.Int(a.multiple, 1) }
func (a *allVehicleViewsMapOptionImpl) HasMultiple() bool  { return a.has_multiple }
func (a *allVehicleViewsMapOptionImpl) SleepSecs() int     { return or.Int(a.sleepSecs, 0) }
func (a *allVehicleViewsMapOptionImpl) HasSleepSecs() bool { return a.has_sleepSecs }
func (a *allVehicleViewsMapOptionImpl) Zoom() int          { return or.Int(a.zoom, 14) }
func (a *allVehicleViewsMapOptionImpl) HasZoom() bool      { return a.has_zoom }

type AllVehicleViewsMapParams struct {
	Latitude  float64 `json:"latitude" default:"40.7701286"`
	Longitude float64 `json:"longitude" default:"-73.9829762"`
	Multiple  int     `json:"multiple" default:"1"`
	SleepSecs int     `json:"sleep_secs" default:"0"`
	Token     string  `json:"token" required:"true"`
	Zoom      int     `json:"zoom" default:"14"`
}

func (o AllVehicleViewsMapParams) Options() []AllVehicleViewsMapOption {
	return []AllVehicleViewsMapOption{
		AllVehicleViewsMapLatitude(o.Latitude),
		AllVehicleViewsMapLongitude(o.Longitude),
		AllVehicleViewsMapMultiple(o.Multiple),
		AllVehicleViewsMapSleepSecs(o.SleepSecs),
		AllVehicleViewsMapZoom(o.Zoom),
	}
}

// ToVehicleViewsMapOptions converts AllVehicleViewsMapOption to an array of VehicleViewsMapOption
func (o *allVehicleViewsMapOptionImpl) ToVehicleViewsMapOptions() []VehicleViewsMapOption {
	return []VehicleViewsMapOption{
		VehicleViewsMapLatitude(o.Latitude()),
		VehicleViewsMapLongitude(o.Longitude()),
		VehicleViewsMapZoom(o.Zoom()),
		VehicleViewsMapSleepSecs(o.SleepSecs()),
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