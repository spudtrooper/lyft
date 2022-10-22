// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type VenuesOption struct {
	f func(*venuesOptionImpl)
	s string
}

func (o VenuesOption) String() string { return o.s }

type VenuesOptions interface {
	Latitude() float64
	HasLatitude() bool
	Longitude() float64
	HasLongitude() bool
	NoIncludeLocationV2Info() bool
	HasNoIncludeLocationV2Info() bool
	RadiusMiles() int
	HasRadiusMiles() bool
	Token() string
	HasToken() bool
	ToBaseOptions() []BaseOption
}

func VenuesLatitude(latitude float64) VenuesOption {
	return VenuesOption{func(opts *venuesOptionImpl) {
		opts.has_latitude = true
		opts.latitude = latitude
	}, fmt.Sprintf("api.VenuesLatitude(float64 %+v)}", latitude)}
}
func VenuesLatitudeFlag(latitude *float64) VenuesOption {
	return VenuesOption{func(opts *venuesOptionImpl) {
		if latitude == nil {
			return
		}
		opts.has_latitude = true
		opts.latitude = *latitude
	}, fmt.Sprintf("api.VenuesLatitude(float64 %+v)}", latitude)}
}

func VenuesLongitude(longitude float64) VenuesOption {
	return VenuesOption{func(opts *venuesOptionImpl) {
		opts.has_longitude = true
		opts.longitude = longitude
	}, fmt.Sprintf("api.VenuesLongitude(float64 %+v)}", longitude)}
}
func VenuesLongitudeFlag(longitude *float64) VenuesOption {
	return VenuesOption{func(opts *venuesOptionImpl) {
		if longitude == nil {
			return
		}
		opts.has_longitude = true
		opts.longitude = *longitude
	}, fmt.Sprintf("api.VenuesLongitude(float64 %+v)}", longitude)}
}

func VenuesNoIncludeLocationV2Info(noIncludeLocationV2Info bool) VenuesOption {
	return VenuesOption{func(opts *venuesOptionImpl) {
		opts.has_noIncludeLocationV2Info = true
		opts.noIncludeLocationV2Info = noIncludeLocationV2Info
	}, fmt.Sprintf("api.VenuesNoIncludeLocationV2Info(bool %+v)}", noIncludeLocationV2Info)}
}
func VenuesNoIncludeLocationV2InfoFlag(noIncludeLocationV2Info *bool) VenuesOption {
	return VenuesOption{func(opts *venuesOptionImpl) {
		if noIncludeLocationV2Info == nil {
			return
		}
		opts.has_noIncludeLocationV2Info = true
		opts.noIncludeLocationV2Info = *noIncludeLocationV2Info
	}, fmt.Sprintf("api.VenuesNoIncludeLocationV2Info(bool %+v)}", noIncludeLocationV2Info)}
}

func VenuesRadiusMiles(radiusMiles int) VenuesOption {
	return VenuesOption{func(opts *venuesOptionImpl) {
		opts.has_radiusMiles = true
		opts.radiusMiles = radiusMiles
	}, fmt.Sprintf("api.VenuesRadiusMiles(int %+v)}", radiusMiles)}
}
func VenuesRadiusMilesFlag(radiusMiles *int) VenuesOption {
	return VenuesOption{func(opts *venuesOptionImpl) {
		if radiusMiles == nil {
			return
		}
		opts.has_radiusMiles = true
		opts.radiusMiles = *radiusMiles
	}, fmt.Sprintf("api.VenuesRadiusMiles(int %+v)}", radiusMiles)}
}

func VenuesToken(token string) VenuesOption {
	return VenuesOption{func(opts *venuesOptionImpl) {
		opts.has_token = true
		opts.token = token
	}, fmt.Sprintf("api.VenuesToken(string %+v)}", token)}
}
func VenuesTokenFlag(token *string) VenuesOption {
	return VenuesOption{func(opts *venuesOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}, fmt.Sprintf("api.VenuesToken(string %+v)}", token)}
}

type venuesOptionImpl struct {
	latitude                    float64
	has_latitude                bool
	longitude                   float64
	has_longitude               bool
	noIncludeLocationV2Info     bool
	has_noIncludeLocationV2Info bool
	radiusMiles                 int
	has_radiusMiles             bool
	token                       string
	has_token                   bool
}

func (v *venuesOptionImpl) Latitude() float64                { return or.Float64(v.latitude, 40.770114) }
func (v *venuesOptionImpl) HasLatitude() bool                { return v.has_latitude }
func (v *venuesOptionImpl) Longitude() float64               { return or.Float64(v.longitude, -73.98302) }
func (v *venuesOptionImpl) HasLongitude() bool               { return v.has_longitude }
func (v *venuesOptionImpl) NoIncludeLocationV2Info() bool    { return v.noIncludeLocationV2Info }
func (v *venuesOptionImpl) HasNoIncludeLocationV2Info() bool { return v.has_noIncludeLocationV2Info }
func (v *venuesOptionImpl) RadiusMiles() int                 { return or.Int(v.radiusMiles, 1) }
func (v *venuesOptionImpl) HasRadiusMiles() bool             { return v.has_radiusMiles }
func (v *venuesOptionImpl) Token() string                    { return v.token }
func (v *venuesOptionImpl) HasToken() bool                   { return v.has_token }

type VenuesParams struct {
	Latitude                float64 `json:"latitude" default:"40.770114"`
	Longitude               float64 `json:"longitude" default:"-73.98302"`
	NoIncludeLocationV2Info bool    `json:"no_include_location_v_2_info"`
	RadiusMiles             int     `json:"radius_miles" default:"1"`
	Token                   string  `json:"token"`
}

func (o VenuesParams) Options() []VenuesOption {
	return []VenuesOption{
		VenuesLatitude(o.Latitude),
		VenuesLongitude(o.Longitude),
		VenuesNoIncludeLocationV2Info(o.NoIncludeLocationV2Info),
		VenuesRadiusMiles(o.RadiusMiles),
		VenuesToken(o.Token),
	}
}

// ToBaseOptions converts VenuesOption to an array of BaseOption
func (o *venuesOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
	}
}

func makeVenuesOptionImpl(opts ...VenuesOption) *venuesOptionImpl {
	res := &venuesOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeVenuesOptions(opts ...VenuesOption) VenuesOptions {
	return makeVenuesOptionImpl(opts...)
}
