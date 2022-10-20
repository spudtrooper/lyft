// DO NOT EDIT MANUALLY: Generated from https://github.com/spudtrooper/genopts
package api

import (
	"fmt"

	"github.com/spudtrooper/goutil/or"
)

type OfferingsOption struct {
	f func(*offeringsOptionImpl)
	s string
}

func (o OfferingsOption) String() string { return o.s }

type OfferingsOptions interface {
	DestinationLatitudeE6() int
	HasDestinationLatitudeE6() bool
	DestinationLongitudeE6() int
	HasDestinationLongitudeE6() bool
	LastOffersID() string
	HasLastOffersID() bool
	OfferSelectorSessionID() string
	HasOfferSelectorSessionID() bool
	OriginLatitudeE6() int
	HasOriginLatitudeE6() bool
	OriginLongitudeE6() int
	HasOriginLongitudeE6() bool
	RequestSource() int
	HasRequestSource() bool
	SupportsAccordionOfferCell() bool
	HasSupportsAccordionOfferCell() bool
	SupportsOfferSelector() bool
	HasSupportsOfferSelector() bool
	SupportsSelectableOfferCell() bool
	HasSupportsSelectableOfferCell() bool
	Token() string
	HasToken() bool
	ToBaseOptions() []BaseOption
}

func OfferingsDestinationLatitudeE6(destinationLatitudeE6 int) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		opts.has_destinationLatitudeE6 = true
		opts.destinationLatitudeE6 = destinationLatitudeE6
	}, fmt.Sprintf("api.OfferingsDestinationLatitudeE6(int %+v)}", destinationLatitudeE6)}
}
func OfferingsDestinationLatitudeE6Flag(destinationLatitudeE6 *int) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		if destinationLatitudeE6 == nil {
			return
		}
		opts.has_destinationLatitudeE6 = true
		opts.destinationLatitudeE6 = *destinationLatitudeE6
	}, fmt.Sprintf("api.OfferingsDestinationLatitudeE6(int %+v)}", destinationLatitudeE6)}
}

func OfferingsDestinationLongitudeE6(destinationLongitudeE6 int) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		opts.has_destinationLongitudeE6 = true
		opts.destinationLongitudeE6 = destinationLongitudeE6
	}, fmt.Sprintf("api.OfferingsDestinationLongitudeE6(int %+v)}", destinationLongitudeE6)}
}
func OfferingsDestinationLongitudeE6Flag(destinationLongitudeE6 *int) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		if destinationLongitudeE6 == nil {
			return
		}
		opts.has_destinationLongitudeE6 = true
		opts.destinationLongitudeE6 = *destinationLongitudeE6
	}, fmt.Sprintf("api.OfferingsDestinationLongitudeE6(int %+v)}", destinationLongitudeE6)}
}

func OfferingsLastOffersID(lastOffersID string) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		opts.has_lastOffersID = true
		opts.lastOffersID = lastOffersID
	}, fmt.Sprintf("api.OfferingsLastOffersID(string %+v)}", lastOffersID)}
}
func OfferingsLastOffersIDFlag(lastOffersID *string) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		if lastOffersID == nil {
			return
		}
		opts.has_lastOffersID = true
		opts.lastOffersID = *lastOffersID
	}, fmt.Sprintf("api.OfferingsLastOffersID(string %+v)}", lastOffersID)}
}

func OfferingsOfferSelectorSessionID(offerSelectorSessionID string) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		opts.has_offerSelectorSessionID = true
		opts.offerSelectorSessionID = offerSelectorSessionID
	}, fmt.Sprintf("api.OfferingsOfferSelectorSessionID(string %+v)}", offerSelectorSessionID)}
}
func OfferingsOfferSelectorSessionIDFlag(offerSelectorSessionID *string) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		if offerSelectorSessionID == nil {
			return
		}
		opts.has_offerSelectorSessionID = true
		opts.offerSelectorSessionID = *offerSelectorSessionID
	}, fmt.Sprintf("api.OfferingsOfferSelectorSessionID(string %+v)}", offerSelectorSessionID)}
}

func OfferingsOriginLatitudeE6(originLatitudeE6 int) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		opts.has_originLatitudeE6 = true
		opts.originLatitudeE6 = originLatitudeE6
	}, fmt.Sprintf("api.OfferingsOriginLatitudeE6(int %+v)}", originLatitudeE6)}
}
func OfferingsOriginLatitudeE6Flag(originLatitudeE6 *int) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		if originLatitudeE6 == nil {
			return
		}
		opts.has_originLatitudeE6 = true
		opts.originLatitudeE6 = *originLatitudeE6
	}, fmt.Sprintf("api.OfferingsOriginLatitudeE6(int %+v)}", originLatitudeE6)}
}

func OfferingsOriginLongitudeE6(originLongitudeE6 int) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		opts.has_originLongitudeE6 = true
		opts.originLongitudeE6 = originLongitudeE6
	}, fmt.Sprintf("api.OfferingsOriginLongitudeE6(int %+v)}", originLongitudeE6)}
}
func OfferingsOriginLongitudeE6Flag(originLongitudeE6 *int) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		if originLongitudeE6 == nil {
			return
		}
		opts.has_originLongitudeE6 = true
		opts.originLongitudeE6 = *originLongitudeE6
	}, fmt.Sprintf("api.OfferingsOriginLongitudeE6(int %+v)}", originLongitudeE6)}
}

func OfferingsRequestSource(requestSource int) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		opts.has_requestSource = true
		opts.requestSource = requestSource
	}, fmt.Sprintf("api.OfferingsRequestSource(int %+v)}", requestSource)}
}
func OfferingsRequestSourceFlag(requestSource *int) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		if requestSource == nil {
			return
		}
		opts.has_requestSource = true
		opts.requestSource = *requestSource
	}, fmt.Sprintf("api.OfferingsRequestSource(int %+v)}", requestSource)}
}

func OfferingsSupportsAccordionOfferCell(supportsAccordionOfferCell bool) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		opts.has_supportsAccordionOfferCell = true
		opts.supportsAccordionOfferCell = supportsAccordionOfferCell
	}, fmt.Sprintf("api.OfferingsSupportsAccordionOfferCell(bool %+v)}", supportsAccordionOfferCell)}
}
func OfferingsSupportsAccordionOfferCellFlag(supportsAccordionOfferCell *bool) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		if supportsAccordionOfferCell == nil {
			return
		}
		opts.has_supportsAccordionOfferCell = true
		opts.supportsAccordionOfferCell = *supportsAccordionOfferCell
	}, fmt.Sprintf("api.OfferingsSupportsAccordionOfferCell(bool %+v)}", supportsAccordionOfferCell)}
}

func OfferingsSupportsOfferSelector(supportsOfferSelector bool) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		opts.has_supportsOfferSelector = true
		opts.supportsOfferSelector = supportsOfferSelector
	}, fmt.Sprintf("api.OfferingsSupportsOfferSelector(bool %+v)}", supportsOfferSelector)}
}
func OfferingsSupportsOfferSelectorFlag(supportsOfferSelector *bool) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		if supportsOfferSelector == nil {
			return
		}
		opts.has_supportsOfferSelector = true
		opts.supportsOfferSelector = *supportsOfferSelector
	}, fmt.Sprintf("api.OfferingsSupportsOfferSelector(bool %+v)}", supportsOfferSelector)}
}

func OfferingsSupportsSelectableOfferCell(supportsSelectableOfferCell bool) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		opts.has_supportsSelectableOfferCell = true
		opts.supportsSelectableOfferCell = supportsSelectableOfferCell
	}, fmt.Sprintf("api.OfferingsSupportsSelectableOfferCell(bool %+v)}", supportsSelectableOfferCell)}
}
func OfferingsSupportsSelectableOfferCellFlag(supportsSelectableOfferCell *bool) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		if supportsSelectableOfferCell == nil {
			return
		}
		opts.has_supportsSelectableOfferCell = true
		opts.supportsSelectableOfferCell = *supportsSelectableOfferCell
	}, fmt.Sprintf("api.OfferingsSupportsSelectableOfferCell(bool %+v)}", supportsSelectableOfferCell)}
}

func OfferingsToken(token string) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		opts.has_token = true
		opts.token = token
	}, fmt.Sprintf("api.OfferingsToken(string %+v)}", token)}
}
func OfferingsTokenFlag(token *string) OfferingsOption {
	return OfferingsOption{func(opts *offeringsOptionImpl) {
		if token == nil {
			return
		}
		opts.has_token = true
		opts.token = *token
	}, fmt.Sprintf("api.OfferingsToken(string %+v)}", token)}
}

type offeringsOptionImpl struct {
	destinationLatitudeE6           int
	has_destinationLatitudeE6       bool
	destinationLongitudeE6          int
	has_destinationLongitudeE6      bool
	lastOffersID                    string
	has_lastOffersID                bool
	offerSelectorSessionID          string
	has_offerSelectorSessionID      bool
	originLatitudeE6                int
	has_originLatitudeE6            bool
	originLongitudeE6               int
	has_originLongitudeE6           bool
	requestSource                   int
	has_requestSource               bool
	supportsAccordionOfferCell      bool
	has_supportsAccordionOfferCell  bool
	supportsOfferSelector           bool
	has_supportsOfferSelector       bool
	supportsSelectableOfferCell     bool
	has_supportsSelectableOfferCell bool
	token                           string
	has_token                       bool
}

func (o *offeringsOptionImpl) DestinationLatitudeE6() int {
	return or.Int(o.destinationLatitudeE6, 40756160)
}
func (o *offeringsOptionImpl) HasDestinationLatitudeE6() bool { return o.has_destinationLatitudeE6 }
func (o *offeringsOptionImpl) DestinationLongitudeE6() int {
	return or.Int(o.destinationLongitudeE6, -73971610)
}
func (o *offeringsOptionImpl) HasDestinationLongitudeE6() bool { return o.has_destinationLongitudeE6 }
func (o *offeringsOptionImpl) LastOffersID() string {
	return or.String(o.lastOffersID, "9ceeb6ef-7023-44ea-9821-acb063780607")
}
func (o *offeringsOptionImpl) HasLastOffersID() bool { return o.has_lastOffersID }
func (o *offeringsOptionImpl) OfferSelectorSessionID() string {
	return or.String(o.offerSelectorSessionID, "45825ecc-6344-45fc-a433-3dafcbcb8665")
}
func (o *offeringsOptionImpl) HasOfferSelectorSessionID() bool { return o.has_offerSelectorSessionID }
func (o *offeringsOptionImpl) OriginLatitudeE6() int           { return or.Int(o.originLatitudeE6, 40770010) }
func (o *offeringsOptionImpl) HasOriginLatitudeE6() bool       { return o.has_originLatitudeE6 }
func (o *offeringsOptionImpl) OriginLongitudeE6() int          { return or.Int(o.originLongitudeE6, -73982839) }
func (o *offeringsOptionImpl) HasOriginLongitudeE6() bool      { return o.has_originLongitudeE6 }
func (o *offeringsOptionImpl) RequestSource() int              { return or.Int(o.requestSource, 1) }
func (o *offeringsOptionImpl) HasRequestSource() bool          { return o.has_requestSource }
func (o *offeringsOptionImpl) SupportsAccordionOfferCell() bool {
	if o.HasSupportsAccordionOfferCell() {
		return o.supportsAccordionOfferCell
	}
	return true
}
func (o *offeringsOptionImpl) HasSupportsAccordionOfferCell() bool {
	return o.has_supportsAccordionOfferCell
}
func (o *offeringsOptionImpl) SupportsOfferSelector() bool {
	if o.HasSupportsOfferSelector() {
		return o.supportsOfferSelector
	}
	return true
}
func (o *offeringsOptionImpl) HasSupportsOfferSelector() bool { return o.has_supportsOfferSelector }
func (o *offeringsOptionImpl) SupportsSelectableOfferCell() bool {
	if o.HasSupportsSelectableOfferCell() {
		return o.supportsSelectableOfferCell
	}
	return true
}
func (o *offeringsOptionImpl) HasSupportsSelectableOfferCell() bool {
	return o.has_supportsSelectableOfferCell
}
func (o *offeringsOptionImpl) Token() string  { return o.token }
func (o *offeringsOptionImpl) HasToken() bool { return o.has_token }

type OfferingsParams struct {
	DestinationLatitudeE6       int    `json:"destination_latitude_e_6" default:"40756160"`
	DestinationLongitudeE6      int    `json:"destination_longitude_e_6" default:"-73971610"`
	LastOffersID                string `json:"last_offers_id" default:"\"9ceeb6ef-7023-44ea-9821-acb063780607\""`
	OfferSelectorSessionID      string `json:"offer_selector_session_id" default:"\"45825ecc-6344-45fc-a433-3dafcbcb8665\""`
	OriginLatitudeE6            int    `json:"origin_latitude_e_6" default:"40770010"`
	OriginLongitudeE6           int    `json:"origin_longitude_e_6" default:"-73982839"`
	RequestSource               int    `json:"request_source" default:"1"`
	SupportsAccordionOfferCell  bool   `json:"supports_accordion_offer_cell" default:"true"`
	SupportsOfferSelector       bool   `json:"supports_offer_selector" default:"true"`
	SupportsSelectableOfferCell bool   `json:"supports_selectable_offer_cell" default:"true"`
	Token                       string `json:"token"`
}

func (o OfferingsParams) Options() []OfferingsOption {
	return []OfferingsOption{
		OfferingsDestinationLatitudeE6(o.DestinationLatitudeE6),
		OfferingsDestinationLongitudeE6(o.DestinationLongitudeE6),
		OfferingsLastOffersID(o.LastOffersID),
		OfferingsOfferSelectorSessionID(o.OfferSelectorSessionID),
		OfferingsOriginLatitudeE6(o.OriginLatitudeE6),
		OfferingsOriginLongitudeE6(o.OriginLongitudeE6),
		OfferingsRequestSource(o.RequestSource),
		OfferingsSupportsAccordionOfferCell(o.SupportsAccordionOfferCell),
		OfferingsSupportsOfferSelector(o.SupportsOfferSelector),
		OfferingsSupportsSelectableOfferCell(o.SupportsSelectableOfferCell),
		OfferingsToken(o.Token),
	}
}

// ToBaseOptions converts OfferingsOption to an array of BaseOption
func (o *offeringsOptionImpl) ToBaseOptions() []BaseOption {
	return []BaseOption{
		BaseToken(o.Token()),
	}
}

func makeOfferingsOptionImpl(opts ...OfferingsOption) *offeringsOptionImpl {
	res := &offeringsOptionImpl{}
	for _, opt := range opts {
		opt.f(res)
	}
	return res
}

func MakeOfferingsOptions(opts ...OfferingsOption) OfferingsOptions {
	return makeOfferingsOptionImpl(opts...)
}
