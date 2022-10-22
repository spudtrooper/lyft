package api

import (
	"strings"

	"github.com/spudtrooper/goutil/request"
)

type OfferingsInfoOffer struct {
	ID             string `json:"id"`
	OfferProductID string `json:"offer_product_id"`
	OfferToken     string `json:"offer_token"`
	CostEstimate   struct {
		PriceQuoteID          string  `json:"price_quote_id"`
		CostToken             string  `json:"cost_token"`
		RideType              string  `json:"ride_type"`
		EstimatedCostCentsMax string  `json:"estimated_cost_cents_max"`
		EstimatedCostCentsMin string  `json:"estimated_cost_cents_min"`
		UpfrontCostCents      string  `json:"upfront_cost_cents"`
		Currency              string  `json:"currency"`
		ApplePayPreAuthCents  string  `json:"apple_pay_pre_auth_cents"`
		PrimetimePercentage   string  `json:"primetime_percentage"`
		PrimetimeMultiplier   float64 `json:"primetime_multiplier"`
		PriceExplanationText  string  `json:"price_explanation_text"`
		// ApplicableCoupons     []struct {
		// 	ID                string `json:"id"`
		// 	Name              string `json:"name"`
		// 	DiscountType      string `json:"discount_type"`
		// 	Currency          string `json:"currency"`
		// 	DiscountAmountMin string `json:"discount_amount_min"`
		// 	DiscountAmountMax string `json:"discount_amount_max"`
		// 	ExplanationText   string `json:"explanation_text"`
		// 	IsMembership      bool   `json:"is_membership"`
		// 	DetailText        string `json:"detail_text"`
		// 	LineItems         []struct {
		// 		Name   string `json:"name"`
		// 		Amount struct {
		// 			Currency string `json:"currency"`
		// 			Amount   string `json:"amount"`
		// 			Exponent string `json:"exponent"`
		// 		} `json:"amount"`
		// 	} `json:"line_items"`
		// } `json:"applicable_coupons"`
		IsScheduledRide          bool   `json:"is_scheduled_ride"`
		CostTokenExpiryTime      string `json:"cost_token_expiry_time"`
		PriceChanged             bool   `json:"price_changed"`
		EstimatedDurationSeconds string `json:"estimated_duration_seconds"`
		LineItems                []struct {
			Description string `json:"description"`
			Amount      string `json:"amount"`
			Currency    string `json:"currency"`
		} `json:"line_items"`
		// StopsAdjustmentThreshold struct {
		// 	Origin struct {
		// 		Lat           float64 `json:"lat"`
		// 		Lng           float64 `json:"lng"`
		// 		RangeInMeters int     `json:"range_in_meters"`
		// 	} `json:"origin"`
		// } `json:"stops_adjustment_threshold"`
		EstimatedDistanceInMiles float64 `json:"estimated_distance_in_miles"`
	} `json:"cost_estimate,omitempty"`
	// RideTypeDetails struct {
	// 	RideType       string `json:"ride_type"`
	// 	Seats          int    `json:"seats"`
	// 	PricingDetails struct {
	// 		BaseChargeV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"base_charge_v2"`
	// 		CostMinimumV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_minimum_v2"`
	// 		CostPerDistanceUnit struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_per_distance_unit"`
	// 		DistanceUnits   string `json:"distance_units"`
	// 		CostPerMinuteV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_per_minute_v2"`
	// 		ServiceFeeDescription string `json:"service_fee_description"`
	// 	} `json:"pricing_details"`
	// 	Features          []string `json:"features"`
	// 	DisplayProperties struct {
	// 		Name                 string `json:"name"`
	// 		Description          string `json:"description"`
	// 		MapMarkerImage       string `json:"map_marker_image"`
	// 		SmallImage           string `json:"small_image"`
	// 		LargeBackgroundImage string `json:"large_background_image"`
	// 		LargeForegroundImage string `json:"large_foreground_image"`
	// 		HiddenByDefault      bool   `json:"hidden_by_default"`
	// 	} `json:"display_properties"`
	// 	ScheduledPricingDetails struct {
	// 		BaseChargeV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"base_charge_v2"`
	// 		CostMinimumV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_minimum_v2"`
	// 		CostPerDistanceUnit struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_per_distance_unit"`
	// 		DistanceUnits   string `json:"distance_units"`
	// 		CostPerMinuteV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_per_minute_v2"`
	// 		ServiceFeeDescription string `json:"service_fee_description"`
	// 	} `json:"scheduled_pricing_details"`
	// } `json:"ride_type_details,omitempty"`
	// RidePickupDetails struct {
	// } `json:"ride_pickup_details,omitempty"`
	// RideTravelDetails struct {
	// 	PickupEstimate struct {
	// 		TimeRange struct {
	// 			TimestampMs string `json:"timestamp_ms"`
	// 		} `json:"time_range"`
	// 		DurationRange struct {
	// 			DurationMs string `json:"duration_ms"`
	// 		} `json:"duration_range"`
	// 		Timezone string `json:"timezone"`
	// 	} `json:"pickup_estimate"`
	// 	DropoffEstimate struct {
	// 		TimeRange struct {
	// 			TimestampMs string `json:"timestamp_ms"`
	// 		} `json:"time_range"`
	// 		DurationRange struct {
	// 			DurationMs string `json:"duration_ms"`
	// 		} `json:"duration_range"`
	// 		Timezone string `json:"timezone"`
	// 	} `json:"dropoff_estimate"`
	// } `json:"ride_travel_details,omitempty"`
	// AvailabilityDetails struct {
	// 	Display int `json:"display"`
	// 	Latest  int `json:"latest"`
	// } `json:"availability_details"`
	// CostEstimate0 struct {
	// 	PriceQuoteID             string  `json:"price_quote_id"`
	// 	CostToken                string  `json:"cost_token"`
	// 	RideType                 string  `json:"ride_type"`
	// 	EstimatedCostCentsMax    string  `json:"estimated_cost_cents_max"`
	// 	EstimatedCostCentsMin    string  `json:"estimated_cost_cents_min"`
	// 	UpfrontCostCents         string  `json:"upfront_cost_cents"`
	// 	Currency                 string  `json:"currency"`
	// 	ApplePayPreAuthCents     string  `json:"apple_pay_pre_auth_cents"`
	// 	PrimetimePercentage      string  `json:"primetime_percentage"`
	// 	PrimetimeMultiplier      float64 `json:"primetime_multiplier"`
	// 	IsScheduledRide          bool    `json:"is_scheduled_ride"`
	// 	CostTokenExpiryTime      string  `json:"cost_token_expiry_time"`
	// 	PriceChanged             bool    `json:"price_changed"`
	// 	EstimatedDurationSeconds string  `json:"estimated_duration_seconds"`
	// 	LineItems                []struct {
	// 		Description string `json:"description"`
	// 		Amount      string `json:"amount"`
	// 		Currency    string `json:"currency"`
	// 	} `json:"line_items"`
	// 	StopsAdjustmentThreshold struct {
	// 		Origin struct {
	// 			Lat           float64 `json:"lat"`
	// 			Lng           float64 `json:"lng"`
	// 			RangeInMeters int     `json:"range_in_meters"`
	// 		} `json:"origin"`
	// 	} `json:"stops_adjustment_threshold"`
	// 	EstimatedDistanceInMiles float64 `json:"estimated_distance_in_miles"`
	// } `json:"cost_estimate,omitempty"`
	// RideTypeDetails0 struct {
	// 	RideType       string `json:"ride_type"`
	// 	Seats          int    `json:"seats"`
	// 	PricingDetails struct {
	// 		BaseChargeV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"base_charge_v2"`
	// 		CostMinimumV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_minimum_v2"`
	// 		CostPerDistanceUnit struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_per_distance_unit"`
	// 		DistanceUnits   string `json:"distance_units"`
	// 		CostPerMinuteV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_per_minute_v2"`
	// 		ServiceFeeDescription string `json:"service_fee_description"`
	// 	} `json:"pricing_details"`
	// 	Features          []string `json:"features"`
	// 	DisplayProperties struct {
	// 		Name                 string `json:"name"`
	// 		Description          string `json:"description"`
	// 		MapMarkerImage       string `json:"map_marker_image"`
	// 		SmallImage           string `json:"small_image"`
	// 		LargeBackgroundImage string `json:"large_background_image"`
	// 		LargeForegroundImage string `json:"large_foreground_image"`
	// 		HiddenByDefault      bool   `json:"hidden_by_default"`
	// 	} `json:"display_properties"`
	// } `json:"ride_type_details,omitempty"`
	// RideTypeDetails1 struct {
	// 	RideType       string `json:"ride_type"`
	// 	Seats          int    `json:"seats"`
	// 	PricingDetails struct {
	// 		BaseChargeV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"base_charge_v2"`
	// 		CostMinimumV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_minimum_v2"`
	// 		CostPerDistanceUnit struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_per_distance_unit"`
	// 		DistanceUnits   string `json:"distance_units"`
	// 		CostPerMinuteV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_per_minute_v2"`
	// 		ServiceFeeDescription string `json:"service_fee_description"`
	// 	} `json:"pricing_details"`
	// 	Features          []string `json:"features"`
	// 	DisplayProperties struct {
	// 		Name                 string `json:"name"`
	// 		Description          string `json:"description"`
	// 		MapMarkerImage       string `json:"map_marker_image"`
	// 		SmallImage           string `json:"small_image"`
	// 		LargeBackgroundImage string `json:"large_background_image"`
	// 		LargeForegroundImage string `json:"large_foreground_image"`
	// 		HiddenByDefault      bool   `json:"hidden_by_default"`
	// 	} `json:"display_properties"`
	// } `json:"ride_type_details,omitempty"`
	// RidePickupDetails0 struct {
	// 	MaxMatchingWindowMs string `json:"max_matching_window_ms"`
	// } `json:"ride_pickup_details,omitempty"`
	// PaxSavings struct {
	// 	ComparisonProductDisplayName string `json:"comparison_product_display_name"`
	// 	CurrentSavings               struct {
	// 		Currency string `json:"currency"`
	// 		Amount   string `json:"amount"`
	// 		Exponent string `json:"exponent"`
	// 	} `json:"current_savings"`
	// } `json:"pax_savings,omitempty"`
	// CostEstimate1 struct {
	// 	PriceQuoteID          string  `json:"price_quote_id"`
	// 	CostToken             string  `json:"cost_token"`
	// 	RideType              string  `json:"ride_type"`
	// 	EstimatedCostCentsMax string  `json:"estimated_cost_cents_max"`
	// 	EstimatedCostCentsMin string  `json:"estimated_cost_cents_min"`
	// 	UpfrontCostCents      string  `json:"upfront_cost_cents"`
	// 	Currency              string  `json:"currency"`
	// 	ApplePayPreAuthCents  string  `json:"apple_pay_pre_auth_cents"`
	// 	PrimetimePercentage   string  `json:"primetime_percentage"`
	// 	PrimetimeMultiplier   float64 `json:"primetime_multiplier"`
	// 	ApplicableCoupons     []struct {
	// 		ID                string `json:"id"`
	// 		Name              string `json:"name"`
	// 		DiscountType      string `json:"discount_type"`
	// 		Currency          string `json:"currency"`
	// 		DiscountAmountMin string `json:"discount_amount_min"`
	// 		DiscountAmountMax string `json:"discount_amount_max"`
	// 		ExplanationText   string `json:"explanation_text"`
	// 		IsMembership      bool   `json:"is_membership"`
	// 		DetailText        string `json:"detail_text"`
	// 		LineItems         []struct {
	// 			Name   string `json:"name"`
	// 			Amount struct {
	// 				Currency string `json:"currency"`
	// 				Amount   string `json:"amount"`
	// 				Exponent string `json:"exponent"`
	// 			} `json:"amount"`
	// 		} `json:"line_items"`
	// 	} `json:"applicable_coupons"`
	// 	IsScheduledRide          bool   `json:"is_scheduled_ride"`
	// 	CostTokenExpiryTime      string `json:"cost_token_expiry_time"`
	// 	PriceChanged             bool   `json:"price_changed"`
	// 	EstimatedDurationSeconds string `json:"estimated_duration_seconds"`
	// 	LineItems                []struct {
	// 		Description string `json:"description"`
	// 		Amount      string `json:"amount"`
	// 		Currency    string `json:"currency"`
	// 	} `json:"line_items"`
	// 	StopsAdjustmentThreshold struct {
	// 		Origin struct {
	// 			Lat           float64 `json:"lat"`
	// 			Lng           float64 `json:"lng"`
	// 			RangeInMeters int     `json:"range_in_meters"`
	// 		} `json:"origin"`
	// 	} `json:"stops_adjustment_threshold"`
	// 	EstimatedDistanceInMiles float64 `json:"estimated_distance_in_miles"`
	// } `json:"cost_estimate,omitempty"`
	// CostEstimate2 struct {
	// 	RideType              string `json:"ride_type"`
	// 	EstimatedCostCentsMax string `json:"estimated_cost_cents_max"`
	// 	EstimatedCostCentsMin string `json:"estimated_cost_cents_min"`
	// 	Currency              string `json:"currency"`
	// 	ApplePayPreAuthCents  string `json:"apple_pay_pre_auth_cents"`
	// 	ApplePayCountry       string `json:"apple_pay_country"`
	// 	ComparisonCents       string `json:"comparison_cents"`
	// 	ApplicableCoupons     []struct {
	// 		ID              string `json:"id"`
	// 		Currency        string `json:"currency"`
	// 		ExplanationText string `json:"explanation_text"`
	// 		IsMembership    bool   `json:"is_membership"`
	// 		LineItems       []struct {
	// 			Name string `json:"name"`
	// 		} `json:"line_items"`
	// 	} `json:"applicable_coupons"`
	// 	IsScheduledRide          bool    `json:"is_scheduled_ride"`
	// 	PriceChanged             bool    `json:"price_changed"`
	// 	EstimatedDurationSeconds string  `json:"estimated_duration_seconds"`
	// 	PriceSummary             string  `json:"price_summary"`
	// 	PriceDescription         string  `json:"price_description"`
	// 	EstimatedDistanceInMiles float64 `json:"estimated_distance_in_miles"`
	// } `json:"cost_estimate,omitempty"`
	// RideTypeDetails2 struct {
	// 	RideType       string `json:"ride_type"`
	// 	PricingDetails struct {
	// 		BaseChargeV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"base_charge_v2"`
	// 		CostMinimumV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_minimum_v2"`
	// 		CostPerDistanceUnit struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_per_distance_unit"`
	// 		DistanceUnits   string `json:"distance_units"`
	// 		CostPerMinuteV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_per_minute_v2"`
	// 		AdditionalCharge struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"additional_charge"`
	// 		BaseChargeThresholdMinutesV3       string `json:"base_charge_threshold_minutes_v3"`
	// 		AdditionalChargeThresholdMinutesV3 string `json:"additional_charge_threshold_minutes_v3"`
	// 	} `json:"pricing_details"`
	// 	Features          []string `json:"features"`
	// 	DisplayProperties struct {
	// 		Name                 string `json:"name"`
	// 		Description          string `json:"description"`
	// 		MapMarkerImage       string `json:"map_marker_image"`
	// 		SmallImage           string `json:"small_image"`
	// 		LargeBackgroundImage string `json:"large_background_image"`
	// 		LargeForegroundImage string `json:"large_foreground_image"`
	// 		HiddenByDefault      bool   `json:"hidden_by_default"`
	// 	} `json:"display_properties"`
	// 	ScheduledPricingDetails struct {
	// 	} `json:"scheduled_pricing_details"`
	// } `json:"ride_type_details,omitempty"`
	// RideableDetails struct {
	// 	RideableID     string `json:"rideable_id"`
	// 	StartStationID string `json:"start_station_id"`
	// 	EndStationID   string `json:"end_station_id"`
	// } `json:"rideable_details,omitempty"`
	// CostEstimate3 struct {
	// 	RideType              string `json:"ride_type"`
	// 	EstimatedCostCentsMax string `json:"estimated_cost_cents_max"`
	// 	EstimatedCostCentsMin string `json:"estimated_cost_cents_min"`
	// 	Currency              string `json:"currency"`
	// 	ApplePayPreAuthCents  string `json:"apple_pay_pre_auth_cents"`
	// 	ApplePayCountry       string `json:"apple_pay_country"`
	// 	ComparisonCents       string `json:"comparison_cents"`
	// 	ApplicableCoupons     []struct {
	// 		ID              string `json:"id"`
	// 		Currency        string `json:"currency"`
	// 		ExplanationText string `json:"explanation_text"`
	// 		IsMembership    bool   `json:"is_membership"`
	// 		LineItems       []struct {
	// 			Name string `json:"name"`
	// 		} `json:"line_items"`
	// 	} `json:"applicable_coupons"`
	// 	IsScheduledRide          bool    `json:"is_scheduled_ride"`
	// 	PriceChanged             bool    `json:"price_changed"`
	// 	EstimatedDurationSeconds string  `json:"estimated_duration_seconds"`
	// 	PriceSummary             string  `json:"price_summary"`
	// 	PriceDescription         string  `json:"price_description"`
	// 	EstimatedDistanceInMiles float64 `json:"estimated_distance_in_miles"`
	// } `json:"cost_estimate,omitempty"`
	// RideTypeDetails3 struct {
	// 	RideType       string `json:"ride_type"`
	// 	PricingDetails struct {
	// 		BaseChargeV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"base_charge_v2"`
	// 		CostMinimumV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_minimum_v2"`
	// 		CostPerDistanceUnit struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_per_distance_unit"`
	// 		DistanceUnits   string `json:"distance_units"`
	// 		CostPerMinuteV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_per_minute_v2"`
	// 		AdditionalCharge struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"additional_charge"`
	// 		BaseChargeThresholdMinutesV3       string `json:"base_charge_threshold_minutes_v3"`
	// 		AdditionalChargeThresholdMinutesV3 string `json:"additional_charge_threshold_minutes_v3"`
	// 	} `json:"pricing_details"`
	// 	Features          []string `json:"features"`
	// 	DisplayProperties struct {
	// 		Name                 string `json:"name"`
	// 		Description          string `json:"description"`
	// 		MapMarkerImage       string `json:"map_marker_image"`
	// 		SmallImage           string `json:"small_image"`
	// 		LargeBackgroundImage string `json:"large_background_image"`
	// 		LargeForegroundImage string `json:"large_foreground_image"`
	// 		HiddenByDefault      bool   `json:"hidden_by_default"`
	// 	} `json:"display_properties"`
	// 	ScheduledPricingDetails struct {
	// 	} `json:"scheduled_pricing_details"`
	// } `json:"ride_type_details,omitempty"`
	// RideableDetails0 struct {
	// 	StartStationID string `json:"start_station_id"`
	// 	EndStationID   string `json:"end_station_id"`
	// } `json:"rideable_details,omitempty"`
	// CostEstimate4 struct {
	// 	UpfrontCostCents     string `json:"upfront_cost_cents"`
	// 	ApplePayPreAuthCents string `json:"apple_pay_pre_auth_cents"`
	// 	ApplePayCountry      string `json:"apple_pay_country"`
	// 	LineItems            []struct {
	// 		Description string `json:"description"`
	// 		Amount      string `json:"amount"`
	// 		Currency    string `json:"currency"`
	// 	} `json:"line_items"`
	// } `json:"cost_estimate,omitempty"`
	// RideTypeDetails4 struct {
	// 	RideType       string `json:"ride_type"`
	// 	PricingDetails struct {
	// 		BaseChargeV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"base_charge_v2"`
	// 		CostMinimumV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_minimum_v2"`
	// 		CostPerDistanceUnit struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_per_distance_unit"`
	// 		DistanceUnits   string `json:"distance_units"`
	// 		CostPerMinuteV2 struct {
	// 			Currency string `json:"currency"`
	// 			Amount   string `json:"amount"`
	// 			Exponent string `json:"exponent"`
	// 		} `json:"cost_per_minute_v2"`
	// 		ServiceFeeDescription string `json:"service_fee_description"`
	// 	} `json:"pricing_details"`
	// 	Features          []string `json:"features"`
	// 	DisplayProperties struct {
	// 		Name                 string `json:"name"`
	// 		Description          string `json:"description"`
	// 		MapMarkerImage       string `json:"map_marker_image"`
	// 		SmallImage           string `json:"small_image"`
	// 		LargeBackgroundImage string `json:"large_background_image"`
	// 		LargeForegroundImage string `json:"large_foreground_image"`
	// 		HiddenByDefault      bool   `json:"hidden_by_default"`
	// 	} `json:"display_properties"`
	// } `json:"ride_type_details,omitempty"`
}

type OfferingsInfoSelectorCellDisplay struct {
	SelectableOfferCell struct {
		SelectedStandardCell struct {
			HeroImage struct {
				URL   string `json:"url"`
				Alpha int    `json:"alpha"`
			} `json:"hero_image"`
			ThreeLineCenterContentArea struct {
				PrimaryContent struct {
					Blocks []struct {
						Text struct {
							Strings []struct {
								Content string `json:"content"`
								Bold    bool   `json:"bold"`
								Color   int    `json:"color"`
							} `json:"strings"`
							Style int `json:"style"`
							Alpha int `json:"alpha"`
						} `json:"text"`
					} `json:"blocks"`
				} `json:"primary_content"`
				SecondaryContent struct {
					Blocks []struct {
						Text struct {
							Strings []struct {
								Content string `json:"content"`
								Color   int    `json:"color"`
							} `json:"strings"`
							Style int `json:"style"`
							Alpha int `json:"alpha"`
						} `json:"text"`
						InlineIcon struct {
							Icon struct {
								CoreIconV1 int `json:"core_icon_v1"`
							} `json:"icon"`
							Color int `json:"color"`
						} `json:"inline_icon"`
					} `json:"blocks"`
				} `json:"secondary_content"`
				TertiaryContent struct {
					Blocks []struct {
						Text struct {
							Strings []struct {
								Content string `json:"content"`
								Color   int    `json:"color"`
							} `json:"strings"`
							Style int `json:"style"`
							Alpha int `json:"alpha"`
						} `json:"text"`
					} `json:"blocks"`
				} `json:"tertiary_content"`
			} `json:"three_line_center_content_area"`
			ThreeLineEndContentArea struct {
				PrimaryContent struct {
					Blocks []struct {
						Text struct {
							Strings []struct {
								Content string `json:"content"`
								Color   int    `json:"color"`
							} `json:"strings"`
							Style int `json:"style"`
							Alpha int `json:"alpha"`
						} `json:"text"`
						InlineIcon struct {
							Icon struct {
								NewIconV1 int `json:"new_icon_v1"`
							} `json:"icon"`
						} `json:"inline_icon"`
					} `json:"blocks"`
				} `json:"primary_content"`
				SecondaryContent struct {
					Blocks []struct {
						Text struct {
							Strings []struct {
								Content       string `json:"content"`
								Strikethrough bool   `json:"strikethrough"`
								Color         int    `json:"color"`
							} `json:"strings"`
							Style int `json:"style"`
							Alpha int `json:"alpha"`
						} `json:"text"`
					} `json:"blocks"`
				} `json:"secondary_content"`
				TertiaryContent struct {
					Blocks []struct {
						Text struct {
							Strings []struct {
								Content string `json:"content"`
								Color   int    `json:"color"`
							} `json:"strings"`
							Style int `json:"style"`
							Alpha int `json:"alpha"`
						} `json:"text"`
					} `json:"blocks"`
				} `json:"tertiary_content"`
			} `json:"three_line_end_content_area"`
			Accessibility struct {
				Description string `json:"description"`
			} `json:"accessibility"`
		} `json:"selected_standard_cell"`
		DeselectedStandardCell struct {
			HeroImage struct {
				URL   string `json:"url"`
				Alpha int    `json:"alpha"`
			} `json:"hero_image"`
			ThreeLineCenterContentArea struct {
				PrimaryContent struct {
					Blocks []struct {
						Text struct {
							Strings []struct {
								Content string `json:"content"`
								Bold    bool   `json:"bold"`
								Color   int    `json:"color"`
							} `json:"strings"`
							Style int `json:"style"`
							Alpha int `json:"alpha"`
						} `json:"text"`
					} `json:"blocks"`
				} `json:"primary_content"`
				SecondaryContent struct {
					Blocks []struct {
						Text struct {
							Strings []struct {
								Content string `json:"content"`
								Color   int    `json:"color"`
							} `json:"strings"`
							Style int `json:"style"`
							Alpha int `json:"alpha"`
						} `json:"text"`
						InlineIcon struct {
							Icon struct {
								CoreIconV1 int `json:"core_icon_v1"`
							} `json:"icon"`
							Color int `json:"color"`
						} `json:"inline_icon"`
					} `json:"blocks"`
				} `json:"secondary_content"`
			} `json:"three_line_center_content_area"`
			ThreeLineEndContentArea struct {
				PrimaryContent struct {
					Blocks []struct {
						Text struct {
							Strings []struct {
								Content string `json:"content"`
								Color   int    `json:"color"`
							} `json:"strings"`
							Style int `json:"style"`
							Alpha int `json:"alpha"`
						} `json:"text"`
						InlineIcon struct {
							Icon struct {
								NewIconV1 int `json:"new_icon_v1"`
							} `json:"icon"`
						} `json:"inline_icon"`
					} `json:"blocks"`
				} `json:"primary_content"`
				SecondaryContent struct {
					Blocks []struct {
						Text struct {
							Strings []struct {
								Content       string `json:"content"`
								Strikethrough bool   `json:"strikethrough"`
								Color         int    `json:"color"`
							} `json:"strings"`
							Style int `json:"style"`
							Alpha int `json:"alpha"`
						} `json:"text"`
					} `json:"blocks"`
				} `json:"secondary_content"`
				TertiaryContent struct {
					Blocks []struct {
						Text struct {
							Strings []struct {
								Content string `json:"content"`
								Color   int    `json:"color"`
							} `json:"strings"`
							Style int `json:"style"`
							Alpha int `json:"alpha"`
						} `json:"text"`
					} `json:"blocks"`
				} `json:"tertiary_content"`
				HideSecondaryContentWhenOfferSelectorCollapsed bool `json:"hide_secondary_content_when_offer_selector_collapsed"`
			} `json:"three_line_end_content_area"`
			Accessibility struct {
				Description string `json:"description"`
			} `json:"accessibility"`
		} `json:"deselected_standard_cell"`
		BundleKey    string `json:"bundle_key"`
		ActionButton struct {
			ButtonPrimary struct {
				Strings []struct {
					Content string `json:"content"`
				} `json:"strings"`
				Style int `json:"style"`
				Alpha int `json:"alpha"`
			} `json:"button_primary"`
		} `json:"action_button"`
	}
}

type OfferingsInfo struct {
	OffersID   string               `json:"offers_id"`
	Offers     []OfferingsInfoOffer `json:"offers"`
	Categories []struct {
		DisplayName     string   `json:"display_name"`
		OfferProductIds []string `json:"offer_product_ids"`
		IsRecommended   bool     `json:"is_recommended"`
		ID              struct {
			ID string `json:"id"`
		} `json:"id"`
	} `json:"categories"`
	PreselectedOfferProductID    string                                      `json:"preselected_offer_product_id"`
	OfferSelectorSessionID       string                                      `json:"offer_selector_session_id"`
	OfferSelectorCellDisplaysMap map[string]OfferingsInfoSelectorCellDisplay `json:"offer_selector_cell_displays_map"`
	DesiredClientPollingRateSec  string                                      `json:"desired_client_polling_rate_sec"`
	PanelConstraint              struct {
		NumItemsDisplayed int `json:"num_items_displayed"`
		Peek              int `json:"peek"`
	} `json:"panel_constraint"`
}

//go:generate genopts --params --function Offerings --extends Base originLatitudeE6:int:40770010 originLongitudeE6:int:-73982839 destinationLatitudeE6:int:40756160 destinationLongitudeE6:int:-73971610 noSupportsSelectableOfferCell noSupportsAccordionOfferCell noSupportsOfferSelector lastOffersID:string:9ceeb6ef-7023-44ea-9821-acb063780607 offerSelectorSessionID:string:45825ecc-6344-45fc-a433-3dafcbcb8665 requestSource:int:1
func (c *Client) Offerings(optss ...OfferingsOption) (*OfferingsInfo, error) {
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
				SupportsSelectableOfferCell: !opts.NoSupportsSelectableOfferCell(),
				SupportsAccordionOfferCell:  !opts.NoSupportsAccordionOfferCell(),
			},
			SupportedContentBlockTextStyles: supportedContentBlockTextStyles,
			SupportedDecisionTreeDomains: supportedDecisionTreeDomains{
				SupportsOfferSelector: !opts.NoSupportsOfferSelector(),
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
	return &res, nil
}
