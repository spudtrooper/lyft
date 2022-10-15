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
	uri := request.MakeURL("https://ride.lyft.com/v1/ridehistory",
		request.Param{"source", `ride_history_list`},
		request.Param{"limit", 10},
		request.Param{"start_time_ms", 1665839945301},
	)
	cookie := [][2]string{
		{"sessId", `c81373c5-189b-4db8-81d3-e095c82045e1L1665798805`},
		{"_gcl_au", `1.1.964926049.1665798806`},
		{"_gid", `GA1.2.894392121.1665798806`},
		{"stickyLyftBrowserId", `AJCyF-JkAFM7vH5VDR7TEOBC`},
		{"_fbp", `fb.1.1665798808967.903911973`},
		{"_tt_enable_cookie", `1`},
		{"_ttp", `d540a25a-7094-482c-a3d9-289a2197ed48`},
		{"_clck", `juh9kb|1|f5q|0`},
		{"lyftAccessToken", `ebwQ5gwOlTHvNx1uuWRrmCUOIePqlAirzPUI+EwdVsHoPNdJ/WFdb4LuvXkwZUnFEPBy/Kj204GtcIXON7YN9/Nu9PwSb0XPwuz7u2Sq2m0ZF3Qs2BvWxwo=`},
		{"__stripe_mid", `f766b515-868f-4f72-9d3c-13f5938ac07b9f9021`},
		{"_uetsid", `2a4141f04c2c11edbbeda94c2a8f6fe6`},
		{"_uetvid", `2a415b404c2c11edbaa9157c03722033`},
		{"OptanonAlertBoxClosed", `2022-10-15T01:54:08.817Z`},
		{"_clsk", `niehx0|1665798848929|5|0|d.clarity.ms/collect`},
		{"_ga", `GA1.2.1739230232.1665798806`},
		{"_ga_LQ1KHS36LD", `GS1.1.1665798806.1.1.1665798850.16.0.0`},
		{"_gat", `1`},
		{"__stripe_sid", `30ca5893-9152-46d9-a105-dc6712a51f3ca536d6`},
		{"OptanonConsent", url.QueryEscape(`isIABGlobal=false&datestamp=Fri Oct 14 2022 21:54:08 GMT-0400 (Eastern Daylight Time)&version=6.18.0&hosts=&consentId=7a20dde4-fde3-4888-8c13-dab5567ad08b&interactionCount=1&landingPath=NotLandingPage&groups=C0001:1,C0003:1,BG1:1,C0002:1,C0004:1&AwaitingReconsent=false&geolocation=US;NY`)},
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
		"referer":            `https://ride.lyft.com/profile/rides?tab=All`,
		"sec-ch-ua":          `"Google Chrome";v="105", "Not)A;Brand";v="8", "Chromium";v="105"`,
		"sec-ch-ua-mobile":   `?0`,
		"sec-ch-ua-platform": `"macOS"`,
		"sec-fetch-dest":     `empty`,
		"sec-fetch-mode":     `cors`,
		"sec-fetch-site":     `same-origin`,
		"user-agent":         `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36`,
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
