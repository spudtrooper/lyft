package render

import (
	"bytes"
	_ "embed"
	"sort"

	"github.com/spudtrooper/lyft/api"
)

//go:embed tmpl/nearbyDrivers.html
var nearbyDriversTmpl string

func NearbyDrivers(input any) ([]byte, error) {
	params := input.(*api.NearbyDriversInfo)

	type vehicleView struct {
		ID       string
		ImageURL string
		Lat, Lng float64
		IsNearby bool
	}
	var vehicleViews []vehicleView

	for _, d := range params.NearbyDrivers {
		id := d.ID
		var imageURL string
		var isNearby bool
		var lat, lng float64
		for _, nb := range params.NearbyDriversByStableOfferProductID {
			for _, p := range nb.NearbyDriverProducts {
				if id == p.DriverID {
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
		vehicleViews = append(vehicleViews, vehicleView{
			ID:       id,
			ImageURL: imageURL,
			Lat:      lat,
			Lng:      lng,
			IsNearby: isNearby,
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
		return nil, err
	}
	return buf.Bytes(), nil
}
