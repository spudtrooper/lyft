package main

import (
	"flag"
	"log"
	"net/url"
	"strings"

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
	uri := request.MakeURL("https://ride.lyft.com/v1/offerings")
	cookie := [][2]string{
		{"sessId", `c81373c5-189b-4db8-81d3-e095c82045e1L1665798805`},
		{"_gcl_au", `1.1.964926049.1665798806`},
		{"_gid", `GA1.2.894392121.1665798806`},
		{"stickyLyftBrowserId", `AJCyF-JkAFM7vH5VDR7TEOBC`},
		{"_fbp", `fb.1.1665798808967.903911973`},
		{"_tt_enable_cookie", `1`},
		{"_ttp", `d540a25a-7094-482c-a3d9-289a2197ed48`},
		{"lyftAccessToken", `ebwQ5gwOlTHvNx1uuWRrmCUOIePqlAirzPUI+EwdVsHoPNdJ/WFdb4LuvXkwZUnFEPBy/Kj204GtcIXON7YN9/Nu9PwSb0XPwuz7u2Sq2m0ZF3Qs2BvWxwo=`},
		{"__stripe_mid", `f766b515-868f-4f72-9d3c-13f5938ac07b9f9021`},
		{"_scid", `940a6637-8391-43d5-a5ad-c366a813fb87`},
		{"_sctr", `1|1665806400000`},
		{"__stripe_sid", `52193719-4ef9-4e70-a6b8-77a239c40dbc066a9d`},
		{"OptanonAlertBoxClosed", `2022-10-16T00:04:22.094Z`},
		{"_uetsid", `2a4141f04c2c11edbbeda94c2a8f6fe6`},
		{"_uetvid", `2a415b404c2c11edbaa9157c03722033`},
		{"_clck", `juh9kb|1|f5r|0`},
		{"_clsk", `ems90q|1665878662371|1|1|d.clarity.ms/collect`},
		{"_ga_LQ1KHS36LD", `GS1.1.1665878662.3.1.1665878664.58.0.0`},
		{"_ga", `GA1.2.1739230232.1665798806`},
		{"OptanonConsent", url.QueryEscape(`isIABGlobal=false&datestamp=Sat Oct 15 2022 20:04:22 GMT-0400 (Eastern Daylight Time)&version=6.18.0&hosts=&consentId=7a20dde4-fde3-4888-8c13-dab5567ad08b&interactionCount=1&landingPath=NotLandingPage&groups=C0001:1,C0003:1,BG1:1,C0002:1,C0004:1&AwaitingReconsent=false&geolocation=US;NY`)},
	}
	headers := map[string]string{
		"authority":          `ride.lyft.com`,
		"accept":             `application/json`,
		"accept-language":    `en-US,en;q=0.9`,
		"cache-control":      `no-cache`,
		"content-type":       `application/json`,
		"dnt":                `1`,
		"downlink":           `10`,
		"dpr":                `2`,
		"origin":             `https://ride.lyft.com`,
		"pragma":             `no-cache`,
		"referer":            `https://ride.lyft.com/ridetype?tab=All&destination=40.75616%2C-73.97161&offerProductId=standard`,
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
	body := `{"origin":{"latitude_e6":40770010,"longitude_e6":-73982839},"destination":{"latitude_e6":40756160,"longitude_e6":-73971610},"template_rendering_specification":{"supported_cell_types":{"supports_selectable_offer_cell":true,"supports_accordion_offer_cell":true},"supported_content_block_text_styles":[6,7,8,9,10,11,12,13],"supported_decision_tree_domains":{"supports_offer_selector":false}},"last_offers_id":"9ceeb6ef-7023-44ea-9821-acb063780607","offer_selector_session_id":"45825ecc-6344-45fc-a433-3dafcbcb8665","request_source":1,"waypoints":[]}`

	{
		type SupportedCellTypes struct {
			SupportsSelectableOfferCell bool `json:"supports_selectable_offer_cell"`
			SupportsAccordionOfferCell  bool `json:"supports_accordion_offer_cell"`
		}
		type SupportedDecisionTreeDomains struct {
			SupportsOfferSelector bool `json:"supports_offer_selector"`
		}
		type TemplateRenderingSpecification struct {
			SupportedCellTypes              SupportedCellTypes           `json:"supported_cell_types"`
			SupportedContentBlockTextStyles []int                        `json:"supported_content_block_text_styles"`
			SupportedDecisionTreeDomains    SupportedDecisionTreeDomains `json:"supported_decision_tree_domains"`
		}
		type Destination struct {
			LatitudeE6  int `json:"latitude_e6"`
			LongitudeE6 int `json:"longitude_e6"`
		}
		type Origin struct {
			LatitudeE6  int `json:"latitude_e6"`
			LongitudeE6 int `json:"longitude_e6"`
		}
		type Body struct {
			Origin                         Origin                         `json:"origin"`
			Destination                    Destination                    `json:"destination"`
			TemplateRenderingSpecification TemplateRenderingSpecification `json:"template_rendering_specification"`
			LastOffersID                   string                         `json:"last_offers_id"`
			OfferSelectorSessionID         string                         `json:"offer_selector_session_id"`
			RequestSource                  int                            `json:"request_source"`
			Waypoints                      []interface{}                  `json:"waypoints"`
		}

		bodyObject := Body{
			Origin: Origin{
				LatitudeE6:  40770010,
				LongitudeE6: -73982839,
			}, Destination: Destination{
				LatitudeE6:  40756160,
				LongitudeE6: -73971610,
			},
			TemplateRenderingSpecification: TemplateRenderingSpecification{
				SupportedCellTypes: SupportedCellTypes{
					SupportsSelectableOfferCell: true,
					SupportsAccordionOfferCell:  true,
				},
				SupportedContentBlockTextStyles: []int{6, 7, 8, 9, 10, 11, 12, 13},
				SupportedDecisionTreeDomains: SupportedDecisionTreeDomains{
					SupportsOfferSelector: false,
				},
			},
			LastOffersID:           "9ceeb6ef-7023-44ea-9821-acb063780607",
			OfferSelectorSessionID: "45825ecc-6344-45fc-a433-3dafcbcb8665",
			RequestSource:          1,
		}
		body = string(request.MustJSONMarshal(bodyObject))
	}

	// Make the request
	var payload interface{}
	var res *request.Response
	var err error

	printData = false
	printCookies = false

	log.Printf("    body: %+v", body)
	panic("asdf")

	res, err = request.Post(uri, &payload, strings.NewReader(body), request.RequestExtraHeaders(headers))

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
