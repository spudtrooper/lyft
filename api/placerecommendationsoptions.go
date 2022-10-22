// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type PlaceRecommendationsOption struct {
	f func(*placeRecommendationsOptionImpl)
	s string
}

func (o PlaceRecommendationsOption) String() string { return o.s }

type PlaceRecommendationsOptions interface {
	Accuracy() int
	HasAccuracy() bool
	Lat() float64
	HasLat() bool
	Lng() float64
	HasLng() bool
	Token() string
	HasToken() bool
	ToBaseOptions() []BaseOption
}

func PlaceRecommendationsAccuracy(accuracy int) PlaceRecommendationsOption {
	return PlaceRecommendationsOption{func(opts *placeRecommendationsOptionImpl) {
		opts.has_accuracy = true
		opts.accuracy = accuracy
	}, fmt.Sprintf("api.PlaceRecommendationsAccuracy(int %+v)", accuracy)}
}
func PlaceRecommendationsAccuracyFlag(accuracy *int) PlaceRecommendationsOption {
	return PlaceRecommendationsOption{func(opts *placeRecommendationsOptionImpl) {
		if accuracy == nil {
			return
		}
		opts.has_accuracy = true
		opts.accuracy = *accuracy
	}, fmt.Sprintf("api.PlaceRecommendationsAccuracy(int %+v)", accuracy)}
}

func PlaceRecommendationsLat(lat float64) PlaceRecommendationsOption {
	return PlaceRecommendationsOption{func(opts *placeRecommendationsOptionImpl) {
		opts.has_lat = true
		opts.lat = lat
	}, fmt.Sprintf("api.PlaceRecommendationsLat(float64 %+v)", lat)}
}
func PlaceRecommendationsLatFlag(lat *float64) PlaceRecommendationsOption {
	return PlaceRecommendationsOption{func(opts *placeRecommendationsOptionImpl) {
		if lat == nil {
			return
		}
		opts.has_lat = true
		opts.lat = *lat
	}, fmt.Sprintf("api.PlaceRecommendationsLat(float64 %+v)", lat)}
}

func PlaceRecommendationsLng(lng float64) PlaceRecommendationsOption {
	return PlaceRecommendationsOption{func(opts *placeRecommendationsOptionImpl) {
		opts.has_lng = true
		opts.lng = lng
	}, fmt.Sprintf("api.PlaceRecommendationsLng(float64 %+v)", lng)}
}
func PlaceRecommendationsLngFlag(lng *float64) PlaceRecommendationsOption {
	return PlaceRecommendationsOption{func(opts *placeRecommendationsOptionImpl) {
		if lng == nil {
			return
		}
		opts.has_lng = true
		opts.lng = *lng
	}, fmt.Sprintf("api.PlaceRecommendationsLng(float64 %+v)", lng)}
}

func PlaceRecommendationsToken(token string) PlaceRecommendationsOption {
	return PlaceRecommendationsOption{func(opts *placeRecommendationsOptionImpl) {
		opts.has_token = true
		opts.token = token
	}, fmt.Sprintf("api.PlaceRecommendationsToken(string %+v)", token)}
}
func PlaceRecommendationsTokenFlag(token *string) PlaceRecommendationsOption {
	return PlaceRecommendationsOption{func(opts *placeRecommendationsOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}, fmt.Sprintf("api.PlaceRecommendationsToken(string %+v)", token)}
}

type placeRecommendationsOptionImpl struct {
	accuracy     int
	has_accuracy bool
	lat          float64
	has_lat      bool
	lng          float64
	has_lng      bool
	token        string
	has_token    bool
}

func (p *placeRecommendationsOptionImpl) Accuracy() int     { return p.accuracy }
func (p *placeRecommendationsOptionImpl) HasAccuracy() bool { return p.has_accuracy }
func (p *placeRecommendationsOptionImpl) Lat() float64      { return or.Float64(p.lat, 40.7683) }
func (p *placeRecommendationsOptionImpl) HasLat() bool      { return p.has_lat }
func (p *placeRecommendationsOptionImpl) Lng() float64      { return or.Float64(p.lng, -73.9802) }
func (p *placeRecommendationsOptionImpl) HasLng() bool      { return p.has_lng }
func (p *placeRecommendationsOptionImpl) Token() string     { return p.token }
func (p *placeRecommendationsOptionImpl) HasToken() bool    { return p.has_token }

type PlaceRecommendationsParams struct {
	Accuracy int     `json:"accuracy"`
	Lat      float64 `json:"lat" default:"40.7683"`
	Lng      float64 `json:"lng" default:"-73.9802"`
	Token    string  `json:"token"`
}

func (o PlaceRecommendationsParams) Options() []PlaceRecommendationsOption {
	return []PlaceRecommendationsOption{
		PlaceRecommendationsAccuracy(o.Accuracy),
		PlaceRecommendationsLat(o.Lat),
		PlaceRecommendationsLng(o.Lng),
		PlaceRecommendationsToken(o.Token),
	}
}

// ToBaseOptions converts PlaceRecommendationsOption to an array of BaseOption
func (o *placeRecommendationsOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
	}
}

func makePlaceRecommendationsOptionImpl(opts ...PlaceRecommendationsOption) *placeRecommendationsOptionImpl {
	res := &placeRecommendationsOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakePlaceRecommendationsOptions(opts ...PlaceRecommendationsOption) PlaceRecommendationsOptions {
	return makePlaceRecommendationsOptionImpl(opts...)
}
