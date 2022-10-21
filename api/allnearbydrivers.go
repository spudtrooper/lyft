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
}

//go:generate genopts --params --function AllNearbyDrivers --extends Base,NearbyDrivers deltaE6:int:130 multiples:int:1 threads:int:5 debug
func (c *Client) AllNearbyDrivers(optss ...AllNearbyDriversOption) (*AllNearbyDriversInfo, error) {
	opts := MakeAllNearbyDriversOptions(optss...)

	type grid struct {
		dlat, dlng int
	}
	mul := opts.Multiples()
	lat := opts.OriginLatitudeE6()
	lng := opts.OriginLongitudeE6()
	del := opts.DeltaE6()
	threads := opts.Threads()
	debug := opts.Debug()

	grids := make(chan grid)
	go func() {
		for i := -mul; i <= mul; i++ {
			for j := -mul; j <= mul; j++ {
				dlat, dlng := lat+i*del, lng+j*del
				grids <- grid{
					dlat: dlat,
					dlng: dlng,
				}
			}
		}
		close(grids)
	}()

	infosCh := make(chan NearbyDriversInfo)
	errs := make(chan error)

	go func() {
		var wg sync.WaitGroup
		for i := 0; i < threads; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for g := range grids {
					if debug {
						log.Printf("grid: %+v", g)
					}
					info, err := c.NearbyDrivers(opts.ToNearbyDriversOptions()...)
					if err != nil {
						if debug {
							log.Printf("err: %+v", err)
						}
						errs <- err
						continue
					}
					if debug {
						log.Printf("info: %+v", info)
					}
					infosCh <- *info
				}
				log.Printf("done")
			}()
		}
		wg.Wait()
		close(infosCh)
		close(errs)
	}()

	var infos []NearbyDriversInfo
	var err error
	parallel.WaitFor(func() {
		for info := range infosCh {
			infos = append(infos, info)
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
	}

	return res, nil
}
