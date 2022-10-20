// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"
	"time"

	"github.com/spudtrooper/goutil/or"
)

type AllRideHistoryOption struct {
	f func(*allRideHistoryOptionImpl)
	s string
}

func (o AllRideHistoryOption) String() string { return o.s }

type AllRideHistoryOptions interface {
	Debug() bool
	HasDebug() bool
	Limit() int
	HasLimit() bool
	Source() string
	HasSource() bool
	StartTime() time.Time
	HasStartTime() bool
	Timeout() time.Duration
	HasTimeout() bool
	Token() string
	HasToken() bool
	TotalLimit() int
	HasTotalLimit() bool
	ToRideHistoryOptions() []RideHistoryOption
	ToBaseOptions() []BaseOption
}

func AllRideHistoryDebug(debug bool) AllRideHistoryOption {
	return AllRideHistoryOption{func(opts *allRideHistoryOptionImpl) {
		opts.has_debug = true
		opts.debug = debug
	}, fmt.Sprintf("api.AllRideHistoryDebug(bool %+v)}", debug)}
}
func AllRideHistoryDebugFlag(debug *bool) AllRideHistoryOption {
	return AllRideHistoryOption{func(opts *allRideHistoryOptionImpl) {
		if debug == nil {
			return
		}
		opts.has_debug = true
		opts.debug = *debug
	}, fmt.Sprintf("api.AllRideHistoryDebug(bool %+v)}", debug)}
}

func AllRideHistoryLimit(limit int) AllRideHistoryOption {
	return AllRideHistoryOption{func(opts *allRideHistoryOptionImpl) {
		opts.has_limit = true
		opts.limit = limit
	}, fmt.Sprintf("api.AllRideHistoryLimit(int %+v)}", limit)}
}
func AllRideHistoryLimitFlag(limit *int) AllRideHistoryOption {
	return AllRideHistoryOption{func(opts *allRideHistoryOptionImpl) {
		if limit == nil {
			return
		}
		opts.has_limit = true
		opts.limit = *limit
	}, fmt.Sprintf("api.AllRideHistoryLimit(int %+v)}", limit)}
}

func AllRideHistorySource(source string) AllRideHistoryOption {
	return AllRideHistoryOption{func(opts *allRideHistoryOptionImpl) {
		opts.has_source = true
		opts.source = source
	}, fmt.Sprintf("api.AllRideHistorySource(string %+v)}", source)}
}
func AllRideHistorySourceFlag(source *string) AllRideHistoryOption {
	return AllRideHistoryOption{func(opts *allRideHistoryOptionImpl) {
		if source == nil {
			return
		}
		opts.has_source = true
		opts.source = *source
	}, fmt.Sprintf("api.AllRideHistorySource(string %+v)}", source)}
}

func AllRideHistoryStartTime(startTime time.Time) AllRideHistoryOption {
	return AllRideHistoryOption{func(opts *allRideHistoryOptionImpl) {
		opts.has_startTime = true
		opts.startTime = startTime
	}, fmt.Sprintf("api.AllRideHistoryStartTime(time.Time %+v)}", startTime)}
}
func AllRideHistoryStartTimeFlag(startTime *time.Time) AllRideHistoryOption {
	return AllRideHistoryOption{func(opts *allRideHistoryOptionImpl) {
		if startTime == nil {
			return
		}
		opts.has_startTime = true
		opts.startTime = *startTime
	}, fmt.Sprintf("api.AllRideHistoryStartTime(time.Time %+v)}", startTime)}
}

func AllRideHistoryTimeout(timeout time.Duration) AllRideHistoryOption {
	return AllRideHistoryOption{func(opts *allRideHistoryOptionImpl) {
		opts.has_timeout = true
		opts.timeout = timeout
	}, fmt.Sprintf("api.AllRideHistoryTimeout(time.Duration %+v)}", timeout)}
}
func AllRideHistoryTimeoutFlag(timeout *time.Duration) AllRideHistoryOption {
	return AllRideHistoryOption{func(opts *allRideHistoryOptionImpl) {
		if timeout == nil {
			return
		}
		opts.has_timeout = true
		opts.timeout = *timeout
	}, fmt.Sprintf("api.AllRideHistoryTimeout(time.Duration %+v)}", timeout)}
}

func AllRideHistoryToken(token string) AllRideHistoryOption {
	return AllRideHistoryOption{func(opts *allRideHistoryOptionImpl) {
		opts.has_token = true
		opts.token = token
	}, fmt.Sprintf("api.AllRideHistoryToken(string %+v)}", token)}
}
func AllRideHistoryTokenFlag(token *string) AllRideHistoryOption {
	return AllRideHistoryOption{func(opts *allRideHistoryOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}, fmt.Sprintf("api.AllRideHistoryToken(string %+v)}", token)}
}

func AllRideHistoryTotalLimit(totalLimit int) AllRideHistoryOption {
	return AllRideHistoryOption{func(opts *allRideHistoryOptionImpl) {
		opts.has_totalLimit = true
		opts.totalLimit = totalLimit
	}, fmt.Sprintf("api.AllRideHistoryTotalLimit(int %+v)}", totalLimit)}
}
func AllRideHistoryTotalLimitFlag(totalLimit *int) AllRideHistoryOption {
	return AllRideHistoryOption{func(opts *allRideHistoryOptionImpl) {
		if totalLimit == nil {
			return
		}
		opts.has_totalLimit = true
		opts.totalLimit = *totalLimit
	}, fmt.Sprintf("api.AllRideHistoryTotalLimit(int %+v)}", totalLimit)}
}

type allRideHistoryOptionImpl struct {
	debug          bool
	has_debug      bool
	totalLimit     int
	has_totalLimit bool
	limit          int
	has_limit      bool
	source         string
	has_source     bool
	startTime      time.Time
	has_startTime  bool
	timeout        time.Duration
	has_timeout    bool
	token          string
	has_token      bool
}

func (a *allRideHistoryOptionImpl) Debug() bool            { return a.debug }
func (a *allRideHistoryOptionImpl) HasDebug() bool         { return a.has_debug }
func (a *allRideHistoryOptionImpl) Limit() int             { return or.Int(a.limit, 10) }
func (a *allRideHistoryOptionImpl) HasLimit() bool         { return a.has_limit }
func (a *allRideHistoryOptionImpl) Source() string         { return or.String(a.source, "ride_history_list") }
func (a *allRideHistoryOptionImpl) HasSource() bool        { return a.has_source }
func (a *allRideHistoryOptionImpl) StartTime() time.Time   { return a.startTime }
func (a *allRideHistoryOptionImpl) HasStartTime() bool     { return a.has_startTime }
func (a *allRideHistoryOptionImpl) Timeout() time.Duration { return a.timeout }
func (a *allRideHistoryOptionImpl) HasTimeout() bool       { return a.has_timeout }
func (a *allRideHistoryOptionImpl) Token() string          { return a.token }
func (a *allRideHistoryOptionImpl) HasToken() bool         { return a.has_token }
func (a *allRideHistoryOptionImpl) TotalLimit() int        { return a.totalLimit }
func (a *allRideHistoryOptionImpl) HasTotalLimit() bool    { return a.has_totalLimit }

type AllRideHistoryParams struct {
	Debug      bool          `json:"debug"`
	Limit      int           `json:"limit" default:"10"`
	Source     string        `json:"source" default:"\"ride_history_list\""`
	StartTime  time.Time     `json:"start_time"`
	Timeout    time.Duration `json:"timeout"`
	Token      string        `json:"token"`
	TotalLimit int           `json:"total_limit"`
}

func (o AllRideHistoryParams) Options() []AllRideHistoryOption {
	return []AllRideHistoryOption{
		AllRideHistoryDebug(o.Debug),
		AllRideHistoryLimit(o.Limit),
		AllRideHistorySource(o.Source),
		AllRideHistoryStartTime(o.StartTime),
		AllRideHistoryTimeout(o.Timeout),
		AllRideHistoryToken(o.Token),
		AllRideHistoryTotalLimit(o.TotalLimit),
	}
}

// ToRideHistoryOptions converts AllRideHistoryOption to an array of RideHistoryOption
func (o *allRideHistoryOptionImpl) ToRideHistoryOptions() []RideHistoryOption {
	return []RideHistoryOption{
		RideHistoryLimit(o.Limit()),
		RideHistorySource(o.Source()),
		RideHistoryStartTime(o.StartTime()),
		RideHistoryTimeout(o.Timeout()),
		RideHistoryToken(o.Token()),
	}
}

// ToBaseOptions converts AllRideHistoryOption to an array of BaseOption
func (o *allRideHistoryOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
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
