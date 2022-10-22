package main

import (
	"flag"
	"log"
	"net/url"

	"github.com/spudtrooper/goutil/check"
	"github.com/spudtrooper/goutil/json"
	"github.com/spudtrooper/goutil/request"
)

func main() {
	flag.Parse()
	// Options
	printData := true
	printCookies := true
	printPayload := true

	// Data
	uri := request.MakeURL("https://ride.lyft.com/v3/venues",
		request.MakeParam("type[]", `pickup`),
		request.MakeParam("type[]", `prohibited`),
		request.MakeParam("lat", 40.770114),
		request.MakeParam("lng", -73.983024),
		request.MakeParam("radius_miles", 1),
		request.MakeParam("include_location_v2_info", true),
	)
	cookie := [][2]string{
		{"sessId", `c81373c5-189b-4db8-81d3-e095c82045e1L1665798805`},
		{"_gcl_au", `1.1.964926049.1665798806`},
		{"stickyLyftBrowserId", `AJCyF-JkAFM7vH5VDR7TEOBC`},
		{"_fbp", `fb.1.1665798808967.903911973`},
		{"_tt_enable_cookie", `1`},
		{"_ttp", `d540a25a-7094-482c-a3d9-289a2197ed48`},
		{"__stripe_mid", `f766b515-868f-4f72-9d3c-13f5938ac07b9f9021`},
		{"_scid", `940a6637-8391-43d5-a5ad-c366a813fb87`},
		{"lyftAccessToken", `xxRWGEwhDiNREi70cJqfj0JSUQg9k7iKxvGz8DrIxImxTSzOcBkck+A+ofzUMyJTnczGKFqYtJVrZ8lBz1AxFUDpQAF07v6sWTXi/3xsgN0AWlNt6UCo7dw=`},
		{"_gid", `GA1.2.808968682.1666448719`},
		{"_clck", `juh9kb|1|f5x|0`},
		{"_sctr", `1|1666411200000`},
		{"_uetsid", `5adccf10521511edba15dbd53290e29f`},
		{"_uetvid", `2a415b404c2c11edbaa9157c03722033`},
		{"OptanonAlertBoxClosed", `2022-10-22T14:25:21.280Z`},
		{"_clsk", `1u8l3b5|1666448721449|1|1|i.clarity.ms/collect`},
		{"_ga", `GA1.2.1739230232.1665798806`},
		{"_ga_LQ1KHS36LD", `GS1.1.1666448718.9.1.1666448723.55.0.0`},
		{"__stripe_sid", `3d36b90d-2deb-4c79-a9cc-5bd5f057e81408891e`},
		{"_gat", `1`},
		{"OptanonConsent", url.QueryEscape(`isIABGlobal=false&datestamp=Sat Oct 22 2022 10:25:21 GMT-0400 (Eastern Daylight Time)&version=6.18.0&hosts=&consentId=7a20dde4-fde3-4888-8c13-dab5567ad08b&interactionCount=1&landingPath=NotLandingPage&groups=C0001:1,C0003:1,BG1:1,C0002:1,C0004:1&AwaitingReconsent=false&geolocation=US;NY`)},
	}
	headers := map[string]string{
		"authority":          `ride.lyft.com`,
		"accept":             `application/json`,
		"accept-language":    `en-US,en;q=0.9`,
		"cache-control":      `no-cache`,
		"dnt":                `1`,
		"downlink":           `10`,
		"dpr":                `2`,
		"pragma":             `no-cache`,
		"referer":            `https://ride.lyft.com/ridetype?origin=40.770114%2C-73.983024&destination=40.74201%2C-74.004745&offerProductId=standard`,
		"sec-ch-ua":          `"Chromium";v="106", "Google Chrome";v="106", "Not;A=Brand";v="99"`,
		"sec-ch-ua-mobile":   `?0`,
		"sec-ch-ua-platform": `"macOS"`,
		"sec-fetch-dest":     `empty`,
		"sec-fetch-mode":     `cors`,
		"sec-fetch-site":     `same-origin`,
		"user-agent":         `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36`,
		"viewport-width":     `1680`,
		"x-device-density":   `2`,
		"x-locale-language":  `en-US`,

		"cookie": request.CreateCookie(cookie),
	}

	// Make the request
	var payload interface{}
	var res *request.Response
	var err error

	res, err = request.Get(uri, &payload, request.RequestExtraHeaders(headers))

	if printData {
		log.Printf("data: %s", string(res.Data))
	}
	if printCookies {
		log.Printf("cookies: %v", res.Cookies)
	}
	if printPayload {
		log.Printf("payload: %s", json.MustColorMarshal(payload))
	}
	check.Err(err)
}
