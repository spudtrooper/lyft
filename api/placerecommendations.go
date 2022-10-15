package api

import (
	"log"

	"github.com/spudtrooper/goutil/request"
)

type placeRecommendationInfoPayload struct {
	RecommendedPlaces []struct {
		ConversionScore              float64 `json:"conversion_score"`
		DisplayName                  string  `json:"display_name"`
		DistanceToSearchOriginMeters float64 `json:"distance_to_search_origin_meters"`
		LocationV2                   struct {
			LocationID       string `json:"location_id"`
			LocationMetadata struct {
				StaticMetadata struct {
					IsUnencumbered bool `json:"is_unencumbered"`
					Spot           struct {
						Name       string `json:"name"`
						Navigation struct {
							GmapsExternalNavigation struct {
								Location struct {
									LatMicrodegrees int `json:"lat_microdegrees"`
									LngMicrodegrees int `json:"lng_microdegrees"`
								} `json:"location"`
							} `json:"gmaps_external_navigation"`
							NavigationMethod struct {
								Location struct {
									LatMicrodegrees int `json:"lat_microdegrees"`
									LngMicrodegrees int `json:"lng_microdegrees"`
								} `json:"location"`
							} `json:"navigation_method"`
							WazeExternalNavigation struct {
								Location struct {
									LatMicrodegrees int `json:"lat_microdegrees"`
									LngMicrodegrees int `json:"lng_microdegrees"`
								} `json:"location"`
							} `json:"waze_external_navigation"`
						} `json:"navigation"`
						Place struct {
							ID       string `json:"id"`
							Provider int    `json:"provider"`
						} `json:"place"`
						RoutableAddress string `json:"routable_address"`
					} `json:"spot"`
				} `json:"static_metadata"`
			} `json:"location_metadata"`
			PortableLocationWithFeatures struct {
				PortableLocation struct {
					CompressedSegmentPosition struct {
						Fraction float64 `json:"fraction"`
						Segment  struct {
							DirectionalSegmentID int64 `json:"directional_segment_id"`
							FromNode             struct {
								LatMicrodegrees int   `json:"lat_microdegrees"`
								LngMicrodegrees int   `json:"lng_microdegrees"`
								NodeID          int64 `json:"node_id"`
							} `json:"from_node"`
							LengthMeters float64 `json:"length_meters"`
							MapProfile   struct {
								MapProvider int    `json:"map_provider"`
								MapVersion  string `json:"map_version"`
							} `json:"map_profile"`
							ToNode struct {
								LatMicrodegrees int   `json:"lat_microdegrees"`
								LngMicrodegrees int   `json:"lng_microdegrees"`
								NodeID          int64 `json:"node_id"`
							} `json:"to_node"`
						} `json:"segment"`
					} `json:"compressed_segment_position"`
					Location struct {
						LatMicrodegrees int `json:"lat_microdegrees"`
						LngMicrodegrees int `json:"lng_microdegrees"`
					} `json:"location"`
				} `json:"portable_location"`
			} `json:"portable_location_with_features"`
			Source      string `json:"source"`
			TimestampMs int64  `json:"timestamp_ms"`
		} `json:"location_v2"`
		Place              placeRecommendationInfoPlace `json:"place,omitempty"`
		RecommendationType string                       `json:"recommendation_type"`
		// Place0             placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place1             placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place2             placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place3             placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place4             placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place5             placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place6             placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place7             placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place8             placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place9             placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place10            placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place11            placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place12            placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place13            placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place14            placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place15            placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place16            placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place17            placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place18            placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place19            placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place20            placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place21            placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place22            placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place23            placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place24            placeRecommendationInfoPlace `json:"place,omitempty"`
		// Place25            placeRecommendationInfoPlace `json:"place,omitempty"`
	} `json:"recommended_places"`
}

type placeRecommendationInfoPlace struct {
	Address             string  `json:"address,omitempty"`
	IsVenue             bool    `json:"is_venue,omitempty"`
	Lat                 float64 `json:"lat,omitempty"`
	Lng                 float64 `json:"lng,omitempty"`
	LocationInputSource string  `json:"location_input_source,omitempty"`
	PlaceID             string  `json:"place_id,omitempty"`
	PlaceName           string  `json:"place_name,omitempty"`
	RoutableAddress     string  `json:"routable_address,omitempty"`
	SpotPlaceProvider   int     `json:"spot_place_provider,omitempty"`
}

type PlaceRecommendationInfoLocationV2 struct {
	LocationID       string
	LocationMetadata struct {
		StaticMetadata struct {
			IsUnencumbered bool
			Spot           struct {
				Name       string
				Navigation struct {
					GmapsExternalNavigation struct {
						Location struct {
							LatMicrodegrees int
							LngMicrodegrees int
						}
					}
					NavigationMethod struct {
						Location struct {
							LatMicrodegrees int
							LngMicrodegrees int
						}
					}
					WazeExternalNavigation struct {
						Location struct {
							LatMicrodegrees int
							LngMicrodegrees int
						}
					}
				}
				Place struct {
					ID       string
					Provider int
				}
				RoutableAddress string
			}
		}
	}
	PortableLocationWithFeatures struct {
		PortableLocation struct {
			CompressedSegmentPosition struct {
				Fraction float64
				Segment  struct {
					DirectionalSegmentID int64
					FromNode             struct {
						LatMicrodegrees int
						LngMicrodegrees int
						NodeID          int64
					}
					LengthMeters float64
					MapProfile   struct {
						MapProvider int
						MapVersion  string
					}
					ToNode struct {
						LatMicrodegrees int
						LngMicrodegrees int
						NodeID          int64
					}
				}
			}
			Location struct {
				LatMicrodegrees int
				LngMicrodegrees int
			}
		}
	}
	Source      string
	TimestampMs int64
}

type PlaceRecommendationInfoPlace struct {
	Address             string
	IsVenue             bool
	Lat                 float64
	Lng                 float64
	LocationInputSource string
	PlaceID             string
	PlaceName           string
	RoutableAddress     string
	SpotPlaceProvider   int
}

type PlaceRecommendationInfoRecommendedPlace struct {
	ConversionScore              float64
	DisplayName                  string
	DistanceToSearchOriginMeters float64
	Place                        PlaceRecommendationInfoPlace
	RecommendationType           string
	Location                     PlaceRecommendationInfoLocationV2
}

type PlaceRecommendationInfo struct {
	RecommendedPlaces []PlaceRecommendationInfoRecommendedPlace
}

func convertPlaceRecommendationInfoPayload(p placeRecommendationInfoPayload) *PlaceRecommendationInfo {
	var recommendedPlaces []PlaceRecommendationInfoRecommendedPlace
	for _, rp := range p.RecommendedPlaces {
		recommendedPlaces = append(recommendedPlaces, PlaceRecommendationInfoRecommendedPlace{
			ConversionScore:              rp.ConversionScore,
			DisplayName:                  rp.DisplayName,
			DistanceToSearchOriginMeters: rp.DistanceToSearchOriginMeters,
			Place:                        PlaceRecommendationInfoPlace(rp.Place),
			RecommendationType:           rp.RecommendationType,
			Location:                     PlaceRecommendationInfoLocationV2(rp.LocationV2),
		})
	}
	return &PlaceRecommendationInfo{
		RecommendedPlaces: recommendedPlaces,
	}
}

//go:generate genopts --params --function PlaceRecommendations --extends Base lat:float64:40.7683 lng:float64:-73.9802 accuracy:int
func (c *Client) PlaceRecommendations(optss ...PlaceRecommendationsOption) (*PlaceRecommendationInfo, error) {
	opts := MakePlaceRecommendationsOptions(optss...)

	uri := request.MakeURL("https://ride.lyft.com/v1/place-recommendations",
		request.Param{"lat", opts.Lat()},
		request.Param{"lng", opts.Lng()},
		request.Param{"accuracy", opts.Accuracy()},
	)
	headers := c.makeHeaders(true, opts.ToBaseOptions()...)

	var payload placeRecommendationInfoPayload

	if _, err := request.Get(uri, &payload, request.RequestExtraHeaders(headers)); err != nil {
		return nil, err
	}

	if opts.DebugPayload() {
		log.Printf("payload: %+v", payload)
	}

	return convertPlaceRecommendationInfoPayload(payload), nil
}
