package render

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"fmt"
	"sort"
	"time"

	"github.com/spudtrooper/lyft/api"
	"github.com/spudtrooper/minimalcli/handler"
)

//go:embed tmpl/ride-history.html
var rideHistoryTmpl string

func RideHistory(input any) ([]byte, handler.RendererConfig, error) {
	params := input.(*api.RideHistoryInfo)

	config := handler.RendererConfig{
		IsFragment: true,
	}

	type row struct {
		ID              string
		DriverFirstName string
		DriverPhotoURL  string
		VehicleImageURL string
		PickUp          string
		DropOff         string
		JSON            string
		Distance        string
		Total           string
	}
	var rows []row
	for _, d := range params.Data {
		var jsonObj = struct {
			Data api.RideHistoryInfoData
		}{
			Data: d,
		}
		jsonBytes, err := json.Marshal(jsonObj)
		if err != nil {
			return nil, config, err
		}
		jsonStr := string(jsonBytes)
		pickUp := time.Unix(int64(d.PickupTimestamp), 0).String()
		dropOff := time.Unix(int64(d.DropoffTimestamp), 0).String()
		distance := fmt.Sprintf("%0.1f %s", d.RideDistance.Value, d.RideDistance.Unit)
		total := fmt.Sprintf("%0.2f %s", float64(d.TotalMoney.Amount)/100.0, d.TotalMoney.Currency)
		rows = append(rows, row{
			ID:              d.RideID,
			DriverFirstName: d.DriverFirstName,
			DriverPhotoURL:  d.DriverPhotoURL,
			VehicleImageURL: d.VehicleImageURL,
			JSON:            jsonStr,
			PickUp:          pickUp,
			DropOff:         dropOff,
			Distance:        distance,
			Total:           total,
		})
	}

	sort.Slice(rows, func(i, j int) bool { return rows[i].ID < rows[j].ID })
	var data = struct {
		VehicleViews []row
	}{
		VehicleViews: rows,
	}
	var buf bytes.Buffer
	if err := renderTemplate(&buf, rideHistoryTmpl, "RideHistory", data); err != nil {
		return nil, config, err
	}
	return buf.Bytes(), config, nil
}
