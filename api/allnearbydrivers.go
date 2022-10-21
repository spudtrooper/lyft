package api

import (
	"log"
	"sync"

	"github.com/spudtrooper/goutil/errors"
	"github.com/spudtrooper/goutil/parallel"
)

type AllNearbyDriversInfo struct {
	DefaultNearbyDrivers                []NearbyDriversInfoNearbyDriverByStableOfferProductID
	NearbyDrivers                       map[string]NearbyDriversInfoNearbyDriver                       `json:"nearby_drivers"`
	NearbyDriversByStableOfferProductID map[string]NearbyDriversInfoNearbyDriverByStableOfferProductID `json:"nearby_drivers_by_stable_offer_product_id"`
	Centers                             []PointE6
}

type PointE6 struct {
	Latitude, Longitude int
}

func (n AllNearbyDriversInfo) VehicleViews() ([]VehicleView, error) {
	return toVehicleViews(n.NearbyDrivers, n.NearbyDriversByStableOfferProductID)
}

//go:generate genopts --params --function AllNearbyDrivers --extends Base,NearbyDrivers deltaE6:int:9000 multiples:int:1 threads:int:5 debug
func (c *Client) AllNearbyDrivers(optss ...AllNearbyDriversOption) (*AllNearbyDriversInfo, error) {
	opts := MakeAllNearbyDriversOptions(optss...)

	mul := opts.Multiples()
	lat := opts.OriginLatitudeE6()
	lng := opts.OriginLongitudeE6()
	del := opts.DeltaE6()
	threads := opts.Threads()
	debug := opts.Debug()

	points := make(chan PointE6)
	go func() {
		for i := -mul; i <= mul; i++ {
			for j := -mul; j <= mul; j++ {
				lat, lng := lat+i*del, lng+j*del
				points <- PointE6{
					Latitude:  lat,
					Longitude: lng,
				}
			}
		}
		close(points)
	}()

	type s struct {
		info  NearbyDriversInfo
		point PointE6
	}

	ss := make(chan s)
	errs := make(chan error)

	go func() {
		var wg sync.WaitGroup
		for i := 0; i < threads; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for p := range points {
					os := opts.ToNearbyDriversOptions()
					os = append(os,
						NearbyDriversOriginLatitudeE6(p.Latitude),
						NearbyDriversOriginLongitudeE6(p.Longitude))
					info, err := c.NearbyDrivers(os...)
					if err != nil {
						if debug {
							log.Printf("err: %+v", err)
						}
						errs <- err
						continue
					}
					ss <- s{
						info:  *info,
						point: p,
					}
				}
			}()
		}
		wg.Wait()
		close(ss)
		close(errs)
	}()

	var infos []NearbyDriversInfo
	var centers []PointE6
	var err error
	parallel.WaitFor(func() {
		for s := range ss {
			infos = append(infos, s.info)
			centers = append(centers, s.point)
		}
	}, func() {
		err = errors.FromChannel(errs)
	})

	if err != nil {
		return nil, err
	}

	type nearbyDriversInfoNearbyDriver struct {
		id        string
		locations map[string]NearbyDriversInfoNearbyDriverLocation
	}
	nearbyDrivers := map[string]nearbyDriversInfoNearbyDriver{}
	type nearbyDriversInfoNearbyDriverByStableOfferProductID struct {
		userInterfaceStylesBySourceURL map[string]int
		driverIDs                      map[string]bool
	}
	nearbyDriversByStableOfferProductID := map[string]nearbyDriversInfoNearbyDriverByStableOfferProductID{}

	for _, info := range infos {
		for id, d := range info.NearbyDrivers {
			nearbyDriver, ok := nearbyDrivers[id]
			if !ok {
				nearbyDriver = nearbyDriversInfoNearbyDriver{
					id:        id,
					locations: map[string]NearbyDriversInfoNearbyDriverLocation{},
				}
			}
			for _, l := range d.Locations {
				nearbyDriver.locations[l.hash()] = l
			}
			nearbyDrivers[id] = nearbyDriver
		}
		for id, d := range info.NearbyDriversByStableOfferProductID {
			nearbyDriverByStableOfferProductID, ok := nearbyDriversByStableOfferProductID[id]
			if !ok {
				nearbyDriverByStableOfferProductID = nearbyDriversInfoNearbyDriverByStableOfferProductID{
					userInterfaceStylesBySourceURL: map[string]int{},
					driverIDs:                      map[string]bool{},
				}
			}
			for _, s := range d.MapMarkerImage.Sources {
				nearbyDriverByStableOfferProductID.userInterfaceStylesBySourceURL[s.URL] = s.Media.UserInterfaceStyle
			}
			for _, p := range d.NearbyDriverProducts {
				nearbyDriverByStableOfferProductID.driverIDs[p.DriverID] = true
			}
			nearbyDriversByStableOfferProductID[id] = nearbyDriverByStableOfferProductID
		}
	}

	resNearbyDrivers := map[string]NearbyDriversInfoNearbyDriver{}
	for id, nd := range nearbyDrivers {
		var locs []NearbyDriversInfoNearbyDriverLocation
		for _, l := range nd.locations {
			locs = append(locs, l)
		}
		resNearbyDrivers[id] = NearbyDriversInfoNearbyDriver{
			ID:        nd.id,
			Locations: locs,
		}
	}

	resNearbyDriversByStableOfferProductID := map[string]NearbyDriversInfoNearbyDriverByStableOfferProductID{}
	for id, nd := range nearbyDriversByStableOfferProductID {
		var sources []NearbyDriversInfoNearbyDriverByStableOfferProductIDSource
		for url, style := range nd.userInterfaceStylesBySourceURL {
			sources = append(sources, NearbyDriversInfoNearbyDriverByStableOfferProductIDSource{
				URL: url,
				Media: NearbyDriversInfoNearbyDriverByStableOfferProductIDSourceMedia{
					UserInterfaceStyle: style,
				},
			})
		}
		var nearbyDriverProducts []NearbyDriversInfoNearbyDriverByStableOfferProductIDNearbyDriverProduct
		for driverID := range nd.driverIDs {
			nearbyDriverProducts = append(nearbyDriverProducts, NearbyDriversInfoNearbyDriverByStableOfferProductIDNearbyDriverProduct{
				DriverID: driverID,
			})
		}
		resNearbyDriversByStableOfferProductID[id] = NearbyDriversInfoNearbyDriverByStableOfferProductID{
			MapMarkerImage: NearbyDriversInfoNearbyDriverByStableOfferProductIDMapMarkerImage{
				Sources: sources,
			},
			NearbyDriverProducts: nearbyDriverProducts,
		}
	}

	res := &AllNearbyDriversInfo{
		NearbyDrivers:                       resNearbyDrivers,
		NearbyDriversByStableOfferProductID: resNearbyDriversByStableOfferProductID,
		Centers:                             centers,
	}

	return res, nil
}
