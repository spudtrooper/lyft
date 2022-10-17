package api

import (
	"log"
	"time"
)

//go:generate genopts --params --function AllRideHistory --extends RideHistory,Base debug totalLimit:int
func (c *Client) AllRideHistory(optss ...AllRideHistoryOption) (chan RideHistoryInfo, chan error) {
	opts := MakeAllRideHistoryOptions(optss...)

	log.Printf("optss: %+v", optss)

	data, errs := make(chan RideHistoryInfo), make(chan error)
	var startTimeMS int
	go func() {
		for i := 0; ; i++ {
			if opts.TotalLimit() > 0 && i >= opts.TotalLimit() {
				log.Printf("break because at limit")
				break
			}
			os := opts.ToRideHistoryOptions()
			if startTimeMS != 0 {
				startTime := time.Unix(int64(startTimeMS/1000), 0)
				if opts.Debug() {
					log.Printf("using start_time_ms: %d startTime:%+v", startTimeMS, startTime)
				}
				os = append(os, RideHistoryStartTime(startTime))
			} else {
				if opts.Debug() {
					log.Printf("start_time_ms is 0")
				}
			}
			res, err := c.RideHistory(os...)
			if err != nil {
				errs <- err
				break
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
