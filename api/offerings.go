package api

import (
	"strings"

	"github.com/spudtrooper/goutil/request"
)

// TODO: Make this something.
type OfferingsInfo interface{}

//go:generate genopts --params --function Offerings --extends Base originLatitudeE6:int:40770010 originLongitudeE6:int:-73982839 destinationLatitudeE6:int:40756160 destinationLongitudeE6:int:-73971610 supportsSelectableOfferCell:bool:true supportsAccordionOfferCell:bool:true supportsOfferSelector:bool:true lastOffersID:string:9ceeb6ef-7023-44ea-9821-acb063780607 offerSelectorSessionID:string:45825ecc-6344-45fc-a433-3dafcbcb8665 requestSource:int:1
func (c *Client) Offerings(optss ...OfferingsOption) (OfferingsInfo, error) {
	opts := MakeOfferingsOptions(optss...)

	const uri = "https://ride.lyft.com/v1/offerings"

	headers := c.makeHeaders(true, opts.ToBaseOptions()...)
	supportedContentBlockTextStyles := []int{6, 7, 8, 9, 10, 11, 12, 13}

	type supportedCellTypes struct {
		SupportsSelectableOfferCell bool `json:"supports_selectable_offer_cell"`
		SupportsAccordionOfferCell  bool `json:"supports_accordion_offer_cell"`
	}
	type supportedDecisionTreeDomains struct {
		SupportsOfferSelector bool `json:"supports_offer_selector"`
	}
	type templateRenderingSpecification struct {
		SupportedCellTypes              supportedCellTypes           `json:"supported_cell_types"`
		SupportedContentBlockTextStyles []int                        `json:"supported_content_block_text_styles"`
		SupportedDecisionTreeDomains    supportedDecisionTreeDomains `json:"supported_decision_tree_domains"`
	}
	type destination struct {
		LatitudeE6  int `json:"latitude_e6"`
		LongitudeE6 int `json:"longitude_e6"`
	}
	type origin struct {
		LatitudeE6  int `json:"latitude_e6"`
		LongitudeE6 int `json:"longitude_e6"`
	}
	type body struct {
		Origin                         origin                         `json:"origin"`
		Destination                    destination                    `json:"destination"`
		TemplateRenderingSpecification templateRenderingSpecification `json:"template_rendering_specification"`
		LastOffersID                   string                         `json:"last_offers_id"`
		OfferSelectorSessionID         string                         `json:"offer_selector_session_id"`
		RequestSource                  int                            `json:"request_source"`
		// Waypoints                      []interface{}                  `json:"waypoints"`
	}

	// TODO: Not handling default bools correctly hence the commented out (e.g. SupportsSelectableOfferCell) bools below.
	// The problem is that we are setting them to false, but really we shouldn't set them if nothing is passed in.

	b, err := request.JSONMarshal(body{
		Origin: origin{
			LatitudeE6:  opts.OriginLatitudeE6(),
			LongitudeE6: opts.OriginLongitudeE6(),
		}, Destination: destination{
			LatitudeE6:  opts.DestinationLatitudeE6(),
			LongitudeE6: opts.DestinationLongitudeE6(),
		},
		TemplateRenderingSpecification: templateRenderingSpecification{
			SupportedCellTypes: supportedCellTypes{
				SupportsSelectableOfferCell: true, //opts.SupportsSelectableOfferCell(),
				SupportsAccordionOfferCell:  true, //opts.SupportsAccordionOfferCell(),
			},
			SupportedContentBlockTextStyles: supportedContentBlockTextStyles,
			SupportedDecisionTreeDomains: supportedDecisionTreeDomains{
				SupportsOfferSelector: true, //opts.SupportsOfferSelector(),
			},
		},
		LastOffersID:           opts.LastOffersID(),
		OfferSelectorSessionID: opts.OfferSelectorSessionID(),
		RequestSource:          opts.RequestSource(),
	})

	if err != nil {
		return nil, err
	}

	var res OfferingsInfo
	if _, err := request.Post(uri, &res, strings.NewReader(string(b)), request.RequestExtraHeaders(headers)); err != nil {
		return nil, err
	}
	return res, nil
}
