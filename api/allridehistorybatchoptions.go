// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"time"

	"github.com/spudtrooper/goutil/or"
)

type AllRideHistoryBatchOption func(*allRideHistoryBatchOptionImpl)

type AllRideHistoryBatchOptions interface {
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
	ToAllRideHistoryOptions() []AllRideHistoryOption
	ToBaseOptions() []BaseOption
	ToRideHistoryOptions() []RideHistoryOption
}

func AllRideHistoryBatchDebug(debug bool) AllRideHistoryBatchOption {
	return func(opts *allRideHistoryBatchOptionImpl) {
		opts.has_debug = true
		opts.debug = debug
	}
}
func AllRideHistoryBatchDebugFlag(debug *bool) AllRideHistoryBatchOption {
	return func(opts *allRideHistoryBatchOptionImpl) {
		if debug == nil {
			return
		}
		opts.has_debug = true
		opts.debug = *debug
	}
}

func AllRideHistoryBatchToken(token string) AllRideHistoryBatchOption {
	return func(opts *allRideHistoryBatchOptionImpl) {
		opts.has_token = true
		opts.token = token
	}
}
func AllRideHistoryBatchTokenFlag(token *string) AllRideHistoryBatchOption {
	return func(opts *allRideHistoryBatchOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}
}

func AllRideHistoryBatchLimit(limit int) AllRideHistoryBatchOption {
	return func(opts *allRideHistoryBatchOptionImpl) {
		opts.has_limit = true
		opts.limit = limit
	}
}
func AllRideHistoryBatchLimitFlag(limit *int) AllRideHistoryBatchOption {
	return func(opts *allRideHistoryBatchOptionImpl) {
		if limit == nil {
			return
		}
		opts.has_limit = true
		opts.limit = *limit
	}
}

func AllRideHistoryBatchSource(source string) AllRideHistoryBatchOption {
	return func(opts *allRideHistoryBatchOptionImpl) {
		opts.has_source = true
		opts.source = source
	}
}
func AllRideHistoryBatchSourceFlag(source *string) AllRideHistoryBatchOption {
	return func(opts *allRideHistoryBatchOptionImpl) {
		if source == nil {
			return
		}
		opts.has_source = true
		opts.source = *source
	}
}

func AllRideHistoryBatchStartTime(startTime time.Time) AllRideHistoryBatchOption {
	return func(opts *allRideHistoryBatchOptionImpl) {
		opts.has_startTime = true
		opts.startTime = startTime
	}
}
func AllRideHistoryBatchStartTimeFlag(startTime *time.Time) AllRideHistoryBatchOption {
	return func(opts *allRideHistoryBatchOptionImpl) {
		if startTime == nil {
			return
		}
		opts.has_startTime = true
		opts.startTime = *startTime
	}
}

func AllRideHistoryBatchTimeout(timeout time.Duration) AllRideHistoryBatchOption {
	return func(opts *allRideHistoryBatchOptionImpl) {
		opts.has_timeout = true
		opts.timeout = timeout
	}
}
func AllRideHistoryBatchTimeoutFlag(timeout *time.Duration) AllRideHistoryBatchOption {
	return func(opts *allRideHistoryBatchOptionImpl) {
		if timeout == nil {
			return
		}
		opts.has_timeout = true
		opts.timeout = *timeout
	}
}

type allRideHistoryBatchOptionImpl struct {
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

func (a *allRideHistoryBatchOptionImpl) Debug() bool    { return a.debug }
func (a *allRideHistoryBatchOptionImpl) HasDebug() bool { return a.has_debug }
func (a *allRideHistoryBatchOptionImpl) Token() string  { return a.token }
func (a *allRideHistoryBatchOptionImpl) HasToken() bool { return a.has_token }
func (a *allRideHistoryBatchOptionImpl) Limit() int     { return or.Int(a.limit, 10) }
func (a *allRideHistoryBatchOptionImpl) HasLimit() bool { return a.has_limit }
func (a *allRideHistoryBatchOptionImpl) Source() string {
	return or.String(a.source, "ride_history_list")
}
func (a *allRideHistoryBatchOptionImpl) HasSource() bool        { return a.has_source }
func (a *allRideHistoryBatchOptionImpl) StartTime() time.Time   { return a.startTime }
func (a *allRideHistoryBatchOptionImpl) HasStartTime() bool     { return a.has_startTime }
func (a *allRideHistoryBatchOptionImpl) Timeout() time.Duration { return a.timeout }
func (a *allRideHistoryBatchOptionImpl) HasTimeout() bool       { return a.has_timeout }

type AllRideHistoryBatchParams struct {
	Debug     bool          `json:"debug"`
	Token     string        `json:"token"`
	Limit     int           `json:"limit" default:"10"`
	Source    string        `json:"source" default:"\"ride_history_list\""`
	StartTime time.Time     `json:"start_time"`
	Timeout   time.Duration `json:"timeout"`
}

func (o AllRideHistoryBatchParams) Options() []AllRideHistoryBatchOption {
	return []AllRideHistoryBatchOption{
		AllRideHistoryBatchDebug(o.Debug),
		AllRideHistoryBatchToken(o.Token),
		AllRideHistoryBatchLimit(o.Limit),
		AllRideHistoryBatchSource(o.Source),
		AllRideHistoryBatchStartTime(o.StartTime),
		AllRideHistoryBatchTimeout(o.Timeout),
	}
}

// ToAllRideHistoryOptions converts AllRideHistoryBatchOption to an array of AllRideHistoryOption
func (o *allRideHistoryBatchOptionImpl) ToAllRideHistoryOptions() []AllRideHistoryOption {
	return []AllRideHistoryOption{
		AllRideHistoryDebug(o.Debug()),
	}
}

// ToBaseOptions converts AllRideHistoryBatchOption to an array of BaseOption
func (o *allRideHistoryBatchOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
	}
}

// ToRideHistoryOptions converts AllRideHistoryBatchOption to an array of RideHistoryOption
func (o *allRideHistoryBatchOptionImpl) ToRideHistoryOptions() []RideHistoryOption {
	return []RideHistoryOption{
		RideHistoryLimit(o.Limit()),
		RideHistorySource(o.Source()),
		RideHistoryStartTime(o.StartTime()),
		RideHistoryTimeout(o.Timeout()),
	}
}

func makeAllRideHistoryBatchOptionImpl(opts ...AllRideHistoryBatchOption) *allRideHistoryBatchOptionImpl {
	res := &allRideHistoryBatchOptionImpl{}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

func MakeAllRideHistoryBatchOptions(opts ...AllRideHistoryBatchOption) AllRideHistoryBatchOptions {
	return makeAllRideHistoryBatchOptionImpl(opts...)
}
