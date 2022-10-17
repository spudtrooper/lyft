package render

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"sort"

	"github.com/spudtrooper/lyft/api"
	"github.com/spudtrooper/minimalcli/handler"
)

//go:embed tmpl/nearbyDrivers.html
var nearbyDriversTmpl string

func NearbyDrivers(input any) ([]byte, handler.RendererConfig, error) {
	params := input.(*api.NearbyDriversInfo)

	config := handler.RendererConfig{
		IsFragment: true,
	}

	type vehicleView struct {
		ID       string
		ImageURL string
		Lat, Lng float64
		IsNearby bool
		JSON     string
	}
	var vehicleViews []vehicleView

	for _, d := range params.NearbyDrivers {
		id := d.ID
		var jsonObj = struct {
			NearbyDriver                       api.NearbyDriversInfoNearbyDriver                       `json:"nearby_drivers"`
			NearbyDriverByStableOfferProductID api.NearbyDriversInfoNearbyDriverByStableOfferProductID `json:"nearby_drivers_by_stable_offer_product_id"`
		}{
			NearbyDriver: d,
		}
		var imageURL string
		var isNearby bool
		var lat, lng float64
		for _, nb := range params.NearbyDriversByStableOfferProductID {
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
				}
			}
		}
		jsonBytes, err := json.Marshal(jsonObj)
		if err != nil {
			return nil, config, err
		}
		jsonStr := string(jsonBytes)
		vehicleViews = append(vehicleViews, vehicleView{
			ID:       id,
			ImageURL: imageURL,
			Lat:      lat,
			Lng:      lng,
			IsNearby: isNearby,
			JSON:     jsonStr,
		})
	}

	sort.Slice(vehicleViews, func(i, j int) bool { return vehicleViews[i].ID < vehicleViews[j].ID })
	var data = struct {
		VehicleViews []vehicleView
	}{
		VehicleViews: vehicleViews,
	}
	var buf bytes.Buffer
	if err := renderTemplate(&buf, nearbyDriversTmpl, "NearbyDrivers", data); err != nil {
		return nil, config, err
	}
	return buf.Bytes(), config, nil
}
