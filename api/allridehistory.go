package api

import (
	"log"
	"time"
)

//go:generate genopts --params --function AllRideHistory --extends RideHistory,Base debug
func (c *Client) AllRideHistory(optss ...AllRideHistoryOption) (chan rideHistoryInfo, chan error) {
	opts := MakeAllRideHistoryOptions(optss...)

	data, errs := make(chan rideHistoryInfo), make(chan error)
	var startTimeMS int
	go func() {
		for {
			os := opts.ToRideHistoryOptions()
			if startTimeMS != 0 {
				startTime := time.Unix(int64(startTimeMS/1000), 0)
				if opts.Debug() {
					log.Printf("using start_time_ms: %d startTime:%+v", startTimeMS, startTime)
				}
				os = append(os, RideHistoryStartTime(startTime))
			}
			res, err := c.RideHistory(os...)
			if err != nil {
				errs <- err
				continue
			}

			data <- *res

			if opts.Debug() {
				log.Printf("Got %d rides, hasMore:%t", len(res.Data), res.HasMore)
			}

			if !res.HasMore {
				break
			}
			for _, t := range res.Data {
				millis := t.RequestTimestamp * 1000
				if startTimeMS == 0 || millis < startTimeMS {
					startTimeMS = millis
				}
			}
		}
		close(data)
		close(errs)
	}()

	return data, errs
}
