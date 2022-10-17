// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"
	"time"

	"github.com/spudtrooper/goutil/or"
)

type AllRideHistoryBatchOption struct {
	f func(*allRideHistoryBatchOptionImpl)
	s string
}

func (o AllRideHistoryBatchOption) String() string { return o.s }

type AllRideHistoryBatchOptions interface {
	Debug() bool
	HasDebug() bool
	TotalLimit() int
	HasTotalLimit() bool
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
	ToAllRideHistoryOptions() []AllRideHistoryOption
	ToRideHistoryOptions() []RideHistoryOption
	ToBaseOptions() []BaseOption
}

func AllRideHistoryBatchDebug(debug bool) AllRideHistoryBatchOption {
	return AllRideHistoryBatchOption{func(opts *allRideHistoryBatchOptionImpl) {
		opts.has_debug = true
		opts.debug = debug
	}, fmt.Sprintf("api.AllRideHistoryBatchDebug(bool %+v)}", debug)}
}
func AllRideHistoryBatchDebugFlag(debug *bool) AllRideHistoryBatchOption {
	return AllRideHistoryBatchOption{func(opts *allRideHistoryBatchOptionImpl) {
		if debug == nil {
			return
		}
		opts.has_debug = true
		opts.debug = *debug
	}, fmt.Sprintf("api.AllRideHistoryBatchDebug(bool %+v)}", debug)}
}

func AllRideHistoryBatchTotalLimit(totalLimit int) AllRideHistoryBatchOption {
	return AllRideHistoryBatchOption{func(opts *allRideHistoryBatchOptionImpl) {
		opts.has_totalLimit = true
		opts.totalLimit = totalLimit
	}, fmt.Sprintf("api.AllRideHistoryBatchTotalLimit(int %+v)}", totalLimit)}
}
func AllRideHistoryBatchTotalLimitFlag(totalLimit *int) AllRideHistoryBatchOption {
	return AllRideHistoryBatchOption{func(opts *allRideHistoryBatchOptionImpl) {
		if totalLimit == nil {
			return
		}
		opts.has_totalLimit = true
		opts.totalLimit = *totalLimit
	}, fmt.Sprintf("api.AllRideHistoryBatchTotalLimit(int %+v)}", totalLimit)}
}

func AllRideHistoryBatchLimit(limit int) AllRideHistoryBatchOption {
	return AllRideHistoryBatchOption{func(opts *allRideHistoryBatchOptionImpl) {
		opts.has_limit = true
		opts.limit = limit
	}, fmt.Sprintf("api.AllRideHistoryBatchLimit(int %+v)}", limit)}
}
func AllRideHistoryBatchLimitFlag(limit *int) AllRideHistoryBatchOption {
	return AllRideHistoryBatchOption{func(opts *allRideHistoryBatchOptionImpl) {
		if limit == nil {
			return
		}
		opts.has_limit = true
		opts.limit = *limit
	}, fmt.Sprintf("api.AllRideHistoryBatchLimit(int %+v)}", limit)}
}

func AllRideHistoryBatchSource(source string) AllRideHistoryBatchOption {
	return AllRideHistoryBatchOption{func(opts *allRideHistoryBatchOptionImpl) {
		opts.has_source = true
		opts.source = source
	}, fmt.Sprintf("api.AllRideHistoryBatchSource(string %+v)}", source)}
}
func AllRideHistoryBatchSourceFlag(source *string) AllRideHistoryBatchOption {
	return AllRideHistoryBatchOption{func(opts *allRideHistoryBatchOptionImpl) {
		if source == nil {
			return
		}
		opts.has_source = true
		opts.source = *source
	}, fmt.Sprintf("api.AllRideHistoryBatchSource(string %+v)}", source)}
}

func AllRideHistoryBatchStartTime(startTime time.Time) AllRideHistoryBatchOption {
	return AllRideHistoryBatchOption{func(opts *allRideHistoryBatchOptionImpl) {
		opts.has_startTime = true
		opts.startTime = startTime
	}, fmt.Sprintf("api.AllRideHistoryBatchStartTime(time.Time %+v)}", startTime)}
}
func AllRideHistoryBatchStartTimeFlag(startTime *time.Time) AllRideHistoryBatchOption {
	return AllRideHistoryBatchOption{func(opts *allRideHistoryBatchOptionImpl) {
		if startTime == nil {
			return
		}
		opts.has_startTime = true
		opts.startTime = *startTime
	}, fmt.Sprintf("api.AllRideHistoryBatchStartTime(time.Time %+v)}", startTime)}
}

func AllRideHistoryBatchTimeout(timeout time.Duration) AllRideHistoryBatchOption {
	return AllRideHistoryBatchOption{func(opts *allRideHistoryBatchOptionImpl) {
		opts.has_timeout = true
		opts.timeout = timeout
	}, fmt.Sprintf("api.AllRideHistoryBatchTimeout(time.Duration %+v)}", timeout)}
}
func AllRideHistoryBatchTimeoutFlag(timeout *time.Duration) AllRideHistoryBatchOption {
	return AllRideHistoryBatchOption{func(opts *allRideHistoryBatchOptionImpl) {
		if timeout == nil {
			return
		}
		opts.has_timeout = true
		opts.timeout = *timeout
	}, fmt.Sprintf("api.AllRideHistoryBatchTimeout(time.Duration %+v)}", timeout)}
}

func AllRideHistoryBatchToken(token string) AllRideHistoryBatchOption {
	return AllRideHistoryBatchOption{func(opts *allRideHistoryBatchOptionImpl) {
		opts.has_token = true
		opts.token = token
	}, fmt.Sprintf("api.AllRideHistoryBatchToken(string %+v)}", token)}
}
func AllRideHistoryBatchTokenFlag(token *string) AllRideHistoryBatchOption {
	return AllRideHistoryBatchOption{func(opts *allRideHistoryBatchOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}, fmt.Sprintf("api.AllRideHistoryBatchToken(string %+v)}", token)}
}

type allRideHistoryBatchOptionImpl struct {
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

func (a *allRideHistoryBatchOptionImpl) Debug() bool         { return a.debug }
func (a *allRideHistoryBatchOptionImpl) HasDebug() bool      { return a.has_debug }
func (a *allRideHistoryBatchOptionImpl) TotalLimit() int     { return a.totalLimit }
func (a *allRideHistoryBatchOptionImpl) HasTotalLimit() bool { return a.has_totalLimit }
func (a *allRideHistoryBatchOptionImpl) Limit() int          { return or.Int(a.limit, 10) }
func (a *allRideHistoryBatchOptionImpl) HasLimit() bool      { return a.has_limit }
func (a *allRideHistoryBatchOptionImpl) Source() string {
	return or.String(a.source, "ride_history_list")
}
func (a *allRideHistoryBatchOptionImpl) HasSource() bool        { return a.has_source }
func (a *allRideHistoryBatchOptionImpl) StartTime() time.Time   { return a.startTime }
func (a *allRideHistoryBatchOptionImpl) HasStartTime() bool     { return a.has_startTime }
func (a *allRideHistoryBatchOptionImpl) Timeout() time.Duration { return a.timeout }
func (a *allRideHistoryBatchOptionImpl) HasTimeout() bool       { return a.has_timeout }
func (a *allRideHistoryBatchOptionImpl) Token() string          { return a.token }
func (a *allRideHistoryBatchOptionImpl) HasToken() bool         { return a.has_token }

type AllRideHistoryBatchParams struct {
	Debug      bool          `json:"debug"`
	TotalLimit int           `json:"total_limit"`
	Limit      int           `json:"limit" default:"10"`
	Source     string        `json:"source" default:"\"ride_history_list\""`
	StartTime  time.Time     `json:"start_time"`
	Timeout    time.Duration `json:"timeout"`
	Token      string        `json:"token"`
}

func (o AllRideHistoryBatchParams) Options() []AllRideHistoryBatchOption {
	return []AllRideHistoryBatchOption{
		AllRideHistoryBatchDebug(o.Debug),
		AllRideHistoryBatchTotalLimit(o.TotalLimit),
		AllRideHistoryBatchLimit(o.Limit),
		AllRideHistoryBatchSource(o.Source),
		AllRideHistoryBatchStartTime(o.StartTime),
		AllRideHistoryBatchTimeout(o.Timeout),
		AllRideHistoryBatchToken(o.Token),
	}
}

// ToAllRideHistoryOptions converts AllRideHistoryBatchOption to an array of AllRideHistoryOption
func (o *allRideHistoryBatchOptionImpl) ToAllRideHistoryOptions() []AllRideHistoryOption {
	return []AllRideHistoryOption{
		AllRideHistoryLimit(o.Limit()),
		AllRideHistorySource(o.Source()),
		AllRideHistoryStartTime(o.StartTime()),
		AllRideHistoryTimeout(o.Timeout()),
		AllRideHistoryToken(o.Token()),
		AllRideHistoryDebug(o.Debug()),
		AllRideHistoryTotalLimit(o.TotalLimit()),
	}
}

// ToRideHistoryOptions converts AllRideHistoryBatchOption to an array of RideHistoryOption
func (o *allRideHistoryBatchOptionImpl) ToRideHistoryOptions() []RideHistoryOption {
	return []RideHistoryOption{
		RideHistorySource(o.Source()),
		RideHistoryStartTime(o.StartTime()),
		RideHistoryTimeout(o.Timeout()),
		RideHistoryToken(o.Token()),
		RideHistoryLimit(o.Limit()),
	}
}

// ToBaseOptions converts AllRideHistoryBatchOption to an array of BaseOption
func (o *allRideHistoryBatchOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
	}
}

func makeAllRideHistoryBatchOptionImpl(opts ...AllRideHistoryBatchOption) *allRideHistoryBatchOptionImpl {
	res := &allRideHistoryBatchOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeAllRideHistoryBatchOptions(opts ...AllRideHistoryBatchOption) AllRideHistoryBatchOptions {
	return makeAllRideHistoryBatchOptionImpl(opts...)
}
