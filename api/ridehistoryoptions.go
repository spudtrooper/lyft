// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"time"

	"github.com/spudtrooper/goutil/or"
)

type RideHistoryOption func(*rideHistoryOptionImpl)

type RideHistoryOptions interface {
	Limit() int
	HasLimit() bool
	Source() string
	HasSource() bool
	StartTime() time.Time
	HasStartTime() bool
	Token() string
	HasToken() bool
	DebugBody() bool
	HasDebugBody() bool
	DebugPayload() bool
	HasDebugPayload() bool
	ToBaseOptions() []BaseOption
}

func RideHistoryLimit(limit int) RideHistoryOption {
	return func(opts *rideHistoryOptionImpl) {
		opts.has_limit = true
		opts.limit = limit
	}
}
func RideHistoryLimitFlag(limit *int) RideHistoryOption {
	return func(opts *rideHistoryOptionImpl) {
		if limit == nil {
			return
		}
		opts.has_limit = true
		opts.limit = *limit
	}
}

func RideHistorySource(source string) RideHistoryOption {
	return func(opts *rideHistoryOptionImpl) {
		opts.has_source = true
		opts.source = source
	}
}
func RideHistorySourceFlag(source *string) RideHistoryOption {
	return func(opts *rideHistoryOptionImpl) {
		if source == nil {
			return
		}
		opts.has_source = true
		opts.source = *source
	}
}

func RideHistoryStartTime(startTime time.Time) RideHistoryOption {
	return func(opts *rideHistoryOptionImpl) {
		opts.has_startTime = true
		opts.startTime = startTime
	}
}
func RideHistoryStartTimeFlag(startTime *time.Time) RideHistoryOption {
	return func(opts *rideHistoryOptionImpl) {
		if startTime == nil {
			return
		}
		opts.has_startTime = true
		opts.startTime = *startTime
	}
}

func RideHistoryToken(token string) RideHistoryOption {
	return func(opts *rideHistoryOptionImpl) {
		opts.has_token = true
		opts.token = token
	}
}
func RideHistoryTokenFlag(token *string) RideHistoryOption {
	return func(opts *rideHistoryOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}
}

func RideHistoryDebugBody(debugBody bool) RideHistoryOption {
	return func(opts *rideHistoryOptionImpl) {
		opts.has_debugBody = true
		opts.debugBody = debugBody
	}
}
func RideHistoryDebugBodyFlag(debugBody *bool) RideHistoryOption {
	return func(opts *rideHistoryOptionImpl) {
		if debugBody == nil {
			return
		}
		opts.has_debugBody = true
		opts.debugBody = *debugBody
	}
}

func RideHistoryDebugPayload(debugPayload bool) RideHistoryOption {
	return func(opts *rideHistoryOptionImpl) {
		opts.has_debugPayload = true
		opts.debugPayload = debugPayload
	}
}
func RideHistoryDebugPayloadFlag(debugPayload *bool) RideHistoryOption {
	return func(opts *rideHistoryOptionImpl) {
		if debugPayload == nil {
			return
		}
		opts.has_debugPayload = true
		opts.debugPayload = *debugPayload
	}
}

type rideHistoryOptionImpl struct {
	limit            int
	has_limit        bool
	source           string
	has_source       bool
	startTime        time.Time
	has_startTime    bool
	token            string
	has_token        bool
	debugBody        bool
	has_debugBody    bool
	debugPayload     bool
	has_debugPayload bool
}

func (r *rideHistoryOptionImpl) Limit() int            { return or.Int(r.limit, 10) }
func (r *rideHistoryOptionImpl) HasLimit() bool        { return r.has_limit }
func (r *rideHistoryOptionImpl) Source() string        { return or.String(r.source, "ride_history_list") }
func (r *rideHistoryOptionImpl) HasSource() bool       { return r.has_source }
func (r *rideHistoryOptionImpl) StartTime() time.Time  { return r.startTime }
func (r *rideHistoryOptionImpl) HasStartTime() bool    { return r.has_startTime }
func (r *rideHistoryOptionImpl) Token() string         { return r.token }
func (r *rideHistoryOptionImpl) HasToken() bool        { return r.has_token }
func (r *rideHistoryOptionImpl) DebugBody() bool       { return r.debugBody }
func (r *rideHistoryOptionImpl) HasDebugBody() bool    { return r.has_debugBody }
func (r *rideHistoryOptionImpl) DebugPayload() bool    { return r.debugPayload }
func (r *rideHistoryOptionImpl) HasDebugPayload() bool { return r.has_debugPayload }

type RideHistoryParams struct {
	Limit        int       `json:"limit" default:"10"`
	Source       string    `json:"source" default:"\"ride_history_list\""`
	StartTime    time.Time `json:"start_time"`
	Token        string    `json:"token"`
	DebugBody    bool      `json:"debug_body"`
	DebugPayload bool      `json:"debug_payload"`
}

func (o RideHistoryParams) Options() []RideHistoryOption {
	return []RideHistoryOption{
		RideHistoryLimit(o.Limit),
		RideHistorySource(o.Source),
		RideHistoryStartTime(o.StartTime),
		RideHistoryToken(o.Token),
		RideHistoryDebugBody(o.DebugBody),
		RideHistoryDebugPayload(o.DebugPayload),
	}
}

// ToBaseOptions converts RideHistoryOption to an array of BaseOption
func (o *rideHistoryOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
		BaseDebugBody(o.DebugBody()),
		BaseDebugPayload(o.DebugPayload()),
	}
}

func makeRideHistoryOptionImpl(opts ...RideHistoryOption) *rideHistoryOptionImpl {
	res := &rideHistoryOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeRideHistoryOptions(opts ...RideHistoryOption) RideHistoryOptions {
	return makeRideHistoryOptionImpl(opts...)
}
