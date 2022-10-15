// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"github.com/spudtrooper/goutil/or"
)

type PlaceRecommendationsOption func(*placeRecommendationsOptionImpl)

type PlaceRecommendationsOptions interface {
	Lat() float64
	HasLat() bool
	Lng() float64
	HasLng() bool
	Accuracy() int
	HasAccuracy() bool
	Token() string
	HasToken() bool
	DebugBody() bool
	HasDebugBody() bool
	DebugPayload() bool
	HasDebugPayload() bool
	ToBaseOptions() []BaseOption
}

func PlaceRecommendationsLat(lat float64) PlaceRecommendationsOption {
	return func(opts *placeRecommendationsOptionImpl) {
		opts.has_lat = true
		opts.lat = lat
	}
}
func PlaceRecommendationsLatFlag(lat *float64) PlaceRecommendationsOption {
	return func(opts *placeRecommendationsOptionImpl) {
		if lat == nil {
			return
		}
		opts.has_lat = true
		opts.lat = *lat
	}
}

func PlaceRecommendationsLng(lng float64) PlaceRecommendationsOption {
	return func(opts *placeRecommendationsOptionImpl) {
		opts.has_lng = true
		opts.lng = lng
	}
}
func PlaceRecommendationsLngFlag(lng *float64) PlaceRecommendationsOption {
	return func(opts *placeRecommendationsOptionImpl) {
		if lng == nil {
			return
		}
		opts.has_lng = true
		opts.lng = *lng
	}
}

func PlaceRecommendationsAccuracy(accuracy int) PlaceRecommendationsOption {
	return func(opts *placeRecommendationsOptionImpl) {
		opts.has_accuracy = true
		opts.accuracy = accuracy
	}
}
func PlaceRecommendationsAccuracyFlag(accuracy *int) PlaceRecommendationsOption {
	return func(opts *placeRecommendationsOptionImpl) {
		if accuracy == nil {
			return
		}
		opts.has_accuracy = true
		opts.accuracy = *accuracy
	}
}

func PlaceRecommendationsToken(token string) PlaceRecommendationsOption {
	return func(opts *placeRecommendationsOptionImpl) {
		opts.has_token = true
		opts.token = token
	}
}
func PlaceRecommendationsTokenFlag(token *string) PlaceRecommendationsOption {
	return func(opts *placeRecommendationsOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}
}

func PlaceRecommendationsDebugBody(debugBody bool) PlaceRecommendationsOption {
	return func(opts *placeRecommendationsOptionImpl) {
		opts.has_debugBody = true
		opts.debugBody = debugBody
	}
}
func PlaceRecommendationsDebugBodyFlag(debugBody *bool) PlaceRecommendationsOption {
	return func(opts *placeRecommendationsOptionImpl) {
		if debugBody == nil {
			return
		}
		opts.has_debugBody = true
		opts.debugBody = *debugBody
	}
}

func PlaceRecommendationsDebugPayload(debugPayload bool) PlaceRecommendationsOption {
	return func(opts *placeRecommendationsOptionImpl) {
		opts.has_debugPayload = true
		opts.debugPayload = debugPayload
	}
}
func PlaceRecommendationsDebugPayloadFlag(debugPayload *bool) PlaceRecommendationsOption {
	return func(opts *placeRecommendationsOptionImpl) {
		if debugPayload == nil {
			return
		}
		opts.has_debugPayload = true
		opts.debugPayload = *debugPayload
	}
}

type placeRecommendationsOptionImpl struct {
	lat              float64
	has_lat          bool
	lng              float64
	has_lng          bool
	accuracy         int
	has_accuracy     bool
	token            string
	has_token        bool
	debugBody        bool
	has_debugBody    bool
	debugPayload     bool
	has_debugPayload bool
}

func (p *placeRecommendationsOptionImpl) Lat() float64          { return or.Float64(p.lat, 40.7683) }
func (p *placeRecommendationsOptionImpl) HasLat() bool          { return p.has_lat }
func (p *placeRecommendationsOptionImpl) Lng() float64          { return or.Float64(p.lng, -73.9802) }
func (p *placeRecommendationsOptionImpl) HasLng() bool          { return p.has_lng }
func (p *placeRecommendationsOptionImpl) Accuracy() int         { return p.accuracy }
func (p *placeRecommendationsOptionImpl) HasAccuracy() bool     { return p.has_accuracy }
func (p *placeRecommendationsOptionImpl) Token() string         { return p.token }
func (p *placeRecommendationsOptionImpl) HasToken() bool        { return p.has_token }
func (p *placeRecommendationsOptionImpl) DebugBody() bool       { return p.debugBody }
func (p *placeRecommendationsOptionImpl) HasDebugBody() bool    { return p.has_debugBody }
func (p *placeRecommendationsOptionImpl) DebugPayload() bool    { return p.debugPayload }
func (p *placeRecommendationsOptionImpl) HasDebugPayload() bool { return p.has_debugPayload }

type PlaceRecommendationsParams struct {
	Lat          float64 `json:"lat" default:"40.7683"`
	Lng          float64 `json:"lng" default:"-73.9802"`
	Accuracy     int     `json:"accuracy"`
	Token        string  `json:"token"`
	DebugBody    bool    `json:"debug_body"`
	DebugPayload bool    `json:"debug_payload"`
}

func (o PlaceRecommendationsParams) Options() []PlaceRecommendationsOption {
	return []PlaceRecommendationsOption{
		PlaceRecommendationsLat(o.Lat),
		PlaceRecommendationsLng(o.Lng),
		PlaceRecommendationsAccuracy(o.Accuracy),
		PlaceRecommendationsToken(o.Token),
		PlaceRecommendationsDebugBody(o.DebugBody),
		PlaceRecommendationsDebugPayload(o.DebugPayload),
	}
}

// ToBaseOptions converts PlaceRecommendationsOption to an array of BaseOption
func (o *placeRecommendationsOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
		BaseDebugBody(o.DebugBody()),
		BaseDebugPayload(o.DebugPayload()),
	}
}

func makePlaceRecommendationsOptionImpl(opts ...PlaceRecommendationsOption) *placeRecommendationsOptionImpl {
	res := &placeRecommendationsOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakePlaceRecommendationsOptions(opts ...PlaceRecommendationsOption) PlaceRecommendationsOptions {
	return makePlaceRecommendationsOptionImpl(opts...)
}
