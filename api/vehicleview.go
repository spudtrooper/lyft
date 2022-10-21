package api

import (
	"encoding/json"
)

type VehicleView struct {
	ID         string
	ImageURL   string
	Lat, Lng   float64
	IsNearby   bool
	JSON       string
	DriverType string
}

func toVehicleViews(
	nearbyDrivers map[string]NearbyDriversInfoNearbyDriver,
	nearbyDriversByStableOfferProductID map[string]NearbyDriversInfoNearbyDriverByStableOfferProductID) ([]VehicleView, error) {
	var vehicleViews []VehicleView
	for _, d := range nearbyDrivers {
		id := d.ID
		var jsonObj = struct {
			NearbyDriver                       NearbyDriversInfoNearbyDriver                       `json:"nearby_drivers"`
			NearbyDriverByStableOfferProductID NearbyDriversInfoNearbyDriverByStableOfferProductID `json:"nearby_drivers_by_stable_offer_product_id"`
		}{
			NearbyDriver: d,
		}
		var imageURL string
		var isNearby bool
		var lat, lng float64
		var driverType string
		for typ, nb := range nearbyDriversByStableOfferProductID {
			for _, p := range nb.NearbyDriverProducts {
				if id == p.DriverID {
					jsonObj.NearbyDriverByStableOfferProductID = nb
					isNearby = true
					if len(d.Locations) > 0 {
						lat = d.Locations[0].Lat
						lng = d.Locations[0].Lng
					}
					if len(nb.MapMarkerImage.Sources) > 0 {
						imageURL = nb.MapMarkerImage.Sources[0].URL
					}
					driverType = typ
				}
			}
		}
		jsonBytes, err := json.Marshal(jsonObj)
		if err != nil {
			return nil, err
		}
		jsonStr := string(jsonBytes)
		vehicleViews = append(vehicleViews, VehicleView{
			ID:         id,
			ImageURL:   imageURL,
			Lat:        lat,
			Lng:        lng,
			IsNearby:   isNearby,
			JSON:       jsonStr,
			DriverType: driverType,
		})
	}
	return vehicleViews, nil
}
