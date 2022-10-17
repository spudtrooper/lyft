package api

import (
	"time"

	"github.com/spudtrooper/goutil/or"
	"github.com/spudtrooper/goutil/request"
)

type RideHistoryInfoData struct {
	CancelRole              int    `json:"cancel_role"`
	Distance                int    `json:"distance,omitempty"`
	DriverFirstName         string `json:"driver_first_name"`
	DriverPhotoURL          string `json:"driver_photo_url"`
	DropoffTimestamp        int    `json:"dropoff_timestamp,omitempty"`
	IsBusinessRide          bool   `json:"is_business_ride"`
	IsSharedUserPaymentRide bool   `json:"is_shared_user_payment_ride"`
	PickupTimestamp         int    `json:"pickup_timestamp"`
	RequestTimestamp        int    `json:"request_timestamp"`
	RideDistance            struct {
		Unit  string  `json:"unit"`
		Value float64 `json:"value"`
	} `json:"ride_distance"`
	RideID        string `json:"ride_id"`
	RideState     string `json:"ride_state"`
	RideType      string `json:"ride_type"`
	RideTypeLabel string `json:"ride_type_label"`
	Timezone      string `json:"timezone"`
	TotalMoney    struct {
		Amount   int    `json:"amount"`
		Currency string `json:"currency"`
		Exponent int    `json:"exponent"`
	} `json:"total_money"`
	VehicleImageURL string `json:"vehicle_image_url"`
}

type RideHistoryInfo struct {
	Data        []RideHistoryInfoData `json:"data"`
	HasMore     bool                  `json:"has_more"`
	Limit       int                   `json:"limit"`
	Skip        int                   `json:"skip"`
	StartTimeMs int64                 `json:"start_time_ms"`
}

//go:generate genopts --params --function RideHistory --extends Base limit:int:10 source:string:ride_history_list startTime:time.Time timeout:time.Duration
func (c *Client) RideHistory(optss ...RideHistoryOption) (*RideHistoryInfo, error) {
	opts := MakeRideHistoryOptions(optss...)

	startTime := or.Time(opts.StartTime(), time.Now())

	uri := request.MakeURL("https://ride.lyft.com/v1/ridehistory",
		request.Param{"source", opts.Source()},
		request.Param{"limit", opts.Limit()},
		request.Param{"start_time_ms", startTime.UnixMilli()},
	)

	headers := c.makeHeaders(true, opts.ToBaseOptions()...)

	var res RideHistoryInfo
	if _, err := request.Get(uri, &res,
		request.RequestExtraHeaders(headers),
		request.RequestTimeout(opts.Timeout())); err != nil {
		return nil, err
	}
	return &res, nil
}
