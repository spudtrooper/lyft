package api

import (
	goutilerrors "github.com/spudtrooper/goutil/errors"
	"github.com/spudtrooper/goutil/parallel"
)

type AllRideHistoryBatchInfo struct {
	Rides []RideHistoryInfo `json:"rides"`
}

// TODO: genopts should transitively extend options
//go:generate genopts --params --function AllRideHistoryBatch --extends AllRideHistory,RideHistory,Base
func (c *Client) AllRideHistoryBatch(optss ...AllRideHistoryBatchOption) (*AllRideHistoryBatchInfo, error) {
	opts := MakeAllRideHistoryBatchOptions(optss...)

	data, errs := c.AllRideHistory(opts.ToAllRideHistoryOptions()...)
	var rides []RideHistoryInfo
	errBuilder := goutilerrors.MakeErrorCollector()
	parallel.WaitFor(func() {
		for d := range data {
			rides = append(rides, d)
		}
	}, func() {
		for e := range errs {
			errBuilder.Add(e)
		}
	})

	if err := errBuilder.Build(); err != nil {
		return nil, err
	}
	res := &AllRideHistoryBatchInfo{
		Rides: rides,
	}
	return res, nil
}
