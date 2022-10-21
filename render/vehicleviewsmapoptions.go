// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package render

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type VehicleViewsMapOption struct {
	f func(*vehicleViewsMapOptionImpl)
	s string
}

func (o VehicleViewsMapOption) String() string { return o.s }

type VehicleViewsMapOptions interface {
	Latitude() float64
	HasLatitude() bool
	Longitude() float64
	HasLongitude() bool
	ShowCenters() bool
	HasShowCenters() bool
	SleepSecs() int
	HasSleepSecs() bool
	Zoom() int
	HasZoom() bool
}

func VehicleViewsMapLatitude(latitude float64) VehicleViewsMapOption {
	return VehicleViewsMapOption{func(opts *vehicleViewsMapOptionImpl) {
		opts.has_latitude = true
		opts.latitude = latitude
	}, fmt.Sprintf("render.VehicleViewsMapLatitude(float64 %+v)}", latitude)}
}
func VehicleViewsMapLatitudeFlag(latitude *float64) VehicleViewsMapOption {
	return VehicleViewsMapOption{func(opts *vehicleViewsMapOptionImpl) {
		if latitude == nil {
			return
		}
		opts.has_latitude = true
		opts.latitude = *latitude
	}, fmt.Sprintf("render.VehicleViewsMapLatitude(float64 %+v)}", latitude)}
}

func VehicleViewsMapLongitude(longitude float64) VehicleViewsMapOption {
	return VehicleViewsMapOption{func(opts *vehicleViewsMapOptionImpl) {
		opts.has_longitude = true
		opts.longitude = longitude
	}, fmt.Sprintf("render.VehicleViewsMapLongitude(float64 %+v)}", longitude)}
}
func VehicleViewsMapLongitudeFlag(longitude *float64) VehicleViewsMapOption {
	return VehicleViewsMapOption{func(opts *vehicleViewsMapOptionImpl) {
		if longitude == nil {
			return
		}
		opts.has_longitude = true
		opts.longitude = *longitude
	}, fmt.Sprintf("render.VehicleViewsMapLongitude(float64 %+v)}", longitude)}
}

func VehicleViewsMapShowCenters(showCenters bool) VehicleViewsMapOption {
	return VehicleViewsMapOption{func(opts *vehicleViewsMapOptionImpl) {
		opts.has_showCenters = true
		opts.showCenters = showCenters
	}, fmt.Sprintf("render.VehicleViewsMapShowCenters(bool %+v)}", showCenters)}
}
func VehicleViewsMapShowCentersFlag(showCenters *bool) VehicleViewsMapOption {
	return VehicleViewsMapOption{func(opts *vehicleViewsMapOptionImpl) {
		if showCenters == nil {
			return
		}
		opts.has_showCenters = true
		opts.showCenters = *showCenters
	}, fmt.Sprintf("render.VehicleViewsMapShowCenters(bool %+v)}", showCenters)}
}

func VehicleViewsMapSleepSecs(sleepSecs int) VehicleViewsMapOption {
	return VehicleViewsMapOption{func(opts *vehicleViewsMapOptionImpl) {
		opts.has_sleepSecs = true
		opts.sleepSecs = sleepSecs
	}, fmt.Sprintf("render.VehicleViewsMapSleepSecs(int %+v)}", sleepSecs)}
}
func VehicleViewsMapSleepSecsFlag(sleepSecs *int) VehicleViewsMapOption {
	return VehicleViewsMapOption{func(opts *vehicleViewsMapOptionImpl) {
		if sleepSecs == nil {
			return
		}
		opts.has_sleepSecs = true
		opts.sleepSecs = *sleepSecs
	}, fmt.Sprintf("render.VehicleViewsMapSleepSecs(int %+v)}", sleepSecs)}
}

func VehicleViewsMapZoom(zoom int) VehicleViewsMapOption {
	return VehicleViewsMapOption{func(opts *vehicleViewsMapOptionImpl) {
		opts.has_zoom = true
		opts.zoom = zoom
	}, fmt.Sprintf("render.VehicleViewsMapZoom(int %+v)}", zoom)}
}
func VehicleViewsMapZoomFlag(zoom *int) VehicleViewsMapOption {
	return VehicleViewsMapOption{func(opts *vehicleViewsMapOptionImpl) {
		if zoom == nil {
			return
		}
		opts.has_zoom = true
		opts.zoom = *zoom
	}, fmt.Sprintf("render.VehicleViewsMapZoom(int %+v)}", zoom)}
}

type vehicleViewsMapOptionImpl struct {
	latitude        float64
	has_latitude    bool
	longitude       float64
	has_longitude   bool
	showCenters     bool
	has_showCenters bool
	sleepSecs       int
	has_sleepSecs   bool
	zoom            int
	has_zoom        bool
}

func (v *vehicleViewsMapOptionImpl) Latitude() float64    { return or.Float64(v.latitude, 40.7701286) }
func (v *vehicleViewsMapOptionImpl) HasLatitude() bool    { return v.has_latitude }
func (v *vehicleViewsMapOptionImpl) Longitude() float64   { return or.Float64(v.longitude, -73.9829762) }
func (v *vehicleViewsMapOptionImpl) HasLongitude() bool   { return v.has_longitude }
func (v *vehicleViewsMapOptionImpl) ShowCenters() bool    { return v.showCenters }
func (v *vehicleViewsMapOptionImpl) HasShowCenters() bool { return v.has_showCenters }
func (v *vehicleViewsMapOptionImpl) SleepSecs() int       { return or.Int(v.sleepSecs, 0) }
func (v *vehicleViewsMapOptionImpl) HasSleepSecs() bool   { return v.has_sleepSecs }
func (v *vehicleViewsMapOptionImpl) Zoom() int            { return or.Int(v.zoom, 14) }
func (v *vehicleViewsMapOptionImpl) HasZoom() bool        { return v.has_zoom }

type VehicleViewsMapParams struct {
	Latitude    float64 `json:"latitude" default:"40.7701286"`
	Longitude   float64 `json:"longitude" default:"-73.9829762"`
	ShowCenters bool    `json:"show_centers"`
	SleepSecs   int     `json:"sleep_secs" default:"0"`
	Token       string  `json:"token" required:"true"`
	Zoom        int     `json:"zoom" default:"14"`
}

func (o VehicleViewsMapParams) Options() []VehicleViewsMapOption {
	return []VehicleViewsMapOption{
		VehicleViewsMapLatitude(o.Latitude),
		VehicleViewsMapLongitude(o.Longitude),
		VehicleViewsMapShowCenters(o.ShowCenters),
		VehicleViewsMapSleepSecs(o.SleepSecs),
		VehicleViewsMapZoom(o.Zoom),
	}
}

func makeVehicleViewsMapOptionImpl(opts ...VehicleViewsMapOption) *vehicleViewsMapOptionImpl {
	res := &vehicleViewsMapOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeVehicleViewsMapOptions(opts ...VehicleViewsMapOption) VehicleViewsMapOptions {
	return makeVehicleViewsMapOptionImpl(opts...)
}
