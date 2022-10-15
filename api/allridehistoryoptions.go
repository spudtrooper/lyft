// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"time"

	"github.com/spudtrooper/goutil/or"
)

type AllRideHistoryOption func(*allRideHistoryOptionImpl)

type AllRideHistoryOptions interface {
	Debug() bool
	HasDebug() bool
	Token() string
	HasToken() bool
	Limit() int
	HasLimit() bool
	Source() string
	HasSource() bool
	StartTime() time.Time
	HasStartTime() bool
	Timeout() time.Duration
	HasTimeout() bool
	ToBaseOptions() []BaseOption
	ToRideHistoryOptions() []RideHistoryOption
}

func AllRideHistoryDebug(debug bool) AllRideHistoryOption {
	return func(opts *allRideHistoryOptionImpl) {
		opts.has_debug = true
		opts.debug = debug
	}
}
func AllRideHistoryDebugFlag(debug *bool) AllRideHistoryOption {
	return func(opts *allRideHistoryOptionImpl) {
		if debug == nil {
			return
		}
		opts.has_debug = true
		opts.debug = *debug
	}
}

func AllRideHistoryToken(token string) AllRideHistoryOption {
	return func(opts *allRideHistoryOptionImpl) {
		opts.has_token = true
		opts.token = token
	}
}
func AllRideHistoryTokenFlag(token *string) AllRideHistoryOption {
	return func(opts *allRideHistoryOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}
}

func AllRideHistoryLimit(limit int) AllRideHistoryOption {
	return func(opts *allRideHistoryOptionImpl) {
		opts.has_limit = true
		opts.limit = limit
	}
}
func AllRideHistoryLimitFlag(limit *int) AllRideHistoryOption {
	return func(opts *allRideHistoryOptionImpl) {
		if limit == nil {
			return
		}
		opts.has_limit = true
		opts.limit = *limit
	}
}

func AllRideHistorySource(source string) AllRideHistoryOption {
	return func(opts *allRideHistoryOptionImpl) {
		opts.has_source = true
		opts.source = source
	}
}
func AllRideHistorySourceFlag(source *string) AllRideHistoryOption {
	return func(opts *allRideHistoryOptionImpl) {
		if source == nil {
			return
		}
		opts.has_source = true
		opts.source = *source
	}
}

func AllRideHistoryStartTime(startTime time.Time) AllRideHistoryOption {
	return func(opts *allRideHistoryOptionImpl) {
		opts.has_startTime = true
		opts.startTime = startTime
	}
}
func AllRideHistoryStartTimeFlag(startTime *time.Time) AllRideHistoryOption {
	return func(opts *allRideHistoryOptionImpl) {
		if startTime == nil {
			return
		}
		opts.has_startTime = true
		opts.startTime = *startTime
	}
}

func AllRideHistoryTimeout(timeout time.Duration) AllRideHistoryOption {
	return func(opts *allRideHistoryOptionImpl) {
		opts.has_timeout = true
		opts.timeout = timeout
	}
}
func AllRideHistoryTimeoutFlag(timeout *time.Duration) AllRideHistoryOption {
	return func(opts *allRideHistoryOptionImpl) {
		if timeout == nil {
			return
		}
		opts.has_timeout = true
		opts.timeout = *timeout
	}
}

type allRideHistoryOptionImpl struct {
	debug         bool
	has_debug     bool
	token         string
	has_token     bool
	limit         int
	has_limit     bool
	source        string
	has_source    bool
	startTime     time.Time
	has_startTime bool
	timeout       time.Duration
	has_timeout   bool
}

func (a *allRideHistoryOptionImpl) Debug() bool            { return a.debug }
func (a *allRideHistoryOptionImpl) HasDebug() bool         { return a.has_debug }
func (a *allRideHistoryOptionImpl) Token() string          { return a.token }
func (a *allRideHistoryOptionImpl) HasToken() bool         { return a.has_token }
func (a *allRideHistoryOptionImpl) Limit() int             { return or.Int(a.limit, 10) }
func (a *allRideHistoryOptionImpl) HasLimit() bool         { return a.has_limit }
func (a *allRideHistoryOptionImpl) Source() string         { return or.String(a.source, "ride_history_list") }
func (a *allRideHistoryOptionImpl) HasSource() bool        { return a.has_source }
func (a *allRideHistoryOptionImpl) StartTime() time.Time   { return a.startTime }
func (a *allRideHistoryOptionImpl) HasStartTime() bool     { return a.has_startTime }
func (a *allRideHistoryOptionImpl) Timeout() time.Duration { return a.timeout }
func (a *allRideHistoryOptionImpl) HasTimeout() bool       { return a.has_timeout }

type AllRideHistoryParams struct {
	Debug     bool          `json:"debug"`
	Token     string        `json:"token"`
	Limit     int           `json:"limit" default:"10"`
	Source    string        `json:"source" default:"\"ride_history_list\""`
	StartTime time.Time     `json:"start_time"`
	Timeout   time.Duration `json:"timeout"`
}

func (o AllRideHistoryParams) Options() []AllRideHistoryOption {
	return []AllRideHistoryOption{
		AllRideHistoryDebug(o.Debug),
		AllRideHistoryToken(o.Token),
		AllRideHistoryLimit(o.Limit),
		AllRideHistorySource(o.Source),
		AllRideHistoryStartTime(o.StartTime),
		AllRideHistoryTimeout(o.Timeout),
	}
}

// ToBaseOptions converts AllRideHistoryOption to an array of BaseOption
func (o *allRideHistoryOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
	}
}

// ToRideHistoryOptions converts AllRideHistoryOption to an array of RideHistoryOption
func (o *allRideHistoryOptionImpl) ToRideHistoryOptions() []RideHistoryOption {
	return []RideHistoryOption{
		RideHistoryLimit(o.Limit()),
		RideHistorySource(o.Source()),
		RideHistoryStartTime(o.StartTime()),
		RideHistoryTimeout(o.Timeout()),
	}
}

func makeAllRideHistoryOptionImpl(opts ...AllRideHistoryOption) *allRideHistoryOptionImpl {
	res := &allRideHistoryOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeAllRideHistoryOptions(opts ...AllRideHistoryOption) AllRideHistoryOptions {
	return makeAllRideHistoryOptionImpl(opts...)
}
