package api

import (
	"log"
	"strings"

	"github.com/spudtrooper/goutil/request"
)

type nearbyDriversInfoPayloadNearbyDriver struct {
	ID        string `json:"id"`
	Locations []struct {
		Bearing      float64 `json:"bearing"`
		Lat          float64 `json:"lat"`
		Lng          float64 `json:"lng"`
		RecordedAtMs int64   `json:"recorded_at_ms"`
	} `json:"locations"`
}

type nearbyDriversInfoPayload struct {
	DefaultNearbyDrivers struct {
		MapMarkerImage struct {
			Sources []struct {
				Media struct {
					UserInterfaceStyle int `json:"user_interface_style"`
				} `json:"media"`
				URL string `json:"url"`
			} `json:"sources"`
		} `json:"map_marker_image"`
		NearbyDriverProducts []struct {
			DriverID string `json:"driver_id"`
		} `json:"nearby_driver_products"`
	} `json:"default_nearby_drivers"`
	NearbyDrivers                       map[string]nearbyDriversInfoPayloadNearbyDriver `json:"nearby_drivers"`
	NearbyDriversByStableOfferProductID struct {
		Access struct {
			MapMarkerImage struct {
				Sources []struct {
					Media struct {
						UserInterfaceStyle int `json:"user_interface_style"`
					} `json:"media"`
					URL string `json:"url"`
				} `json:"sources"`
			} `json:"map_marker_image"`
			NearbyDriverProducts []struct {
				DriverID string `json:"driver_id"`
			} `json:"nearby_driver_products"`
		} `json:"access"`
		Lux struct {
			MapMarkerImage struct {
				Sources []struct {
					Media struct {
						UserInterfaceStyle int `json:"user_interface_style"`
					} `json:"media"`
					URL string `json:"url"`
				} `json:"sources"`
			} `json:"map_marker_image"`
			NearbyDriverProducts []struct {
				DriverID string `json:"driver_id"`
			} `json:"nearby_driver_products"`
		} `json:"lux"`
		Luxsuv struct {
			MapMarkerImage struct {
				Sources []struct {
					Media struct {
						UserInterfaceStyle int `json:"user_interface_style"`
					} `json:"media"`
					URL string `json:"url"`
				} `json:"sources"`
			} `json:"map_marker_image"`
			NearbyDriverProducts []struct {
				DriverID string `json:"driver_id"`
			} `json:"nearby_driver_products"`
		} `json:"luxsuv"`
		Plus struct {
			MapMarkerImage struct {
				Sources []struct {
					Media struct {
						UserInterfaceStyle int `json:"user_interface_style"`
					} `json:"media"`
					URL string `json:"url"`
				} `json:"sources"`
			} `json:"map_marker_image"`
			NearbyDriverProducts []struct {
				DriverID string `json:"driver_id"`
			} `json:"nearby_driver_products"`
		} `json:"plus"`
		Promo struct {
			MapMarkerImage struct {
				Sources []struct {
					Media struct {
						UserInterfaceStyle int `json:"user_interface_style"`
					} `json:"media"`
					URL string `json:"url"`
				} `json:"sources"`
			} `json:"map_marker_image"`
			NearbyDriverProducts []interface{} `json:"nearby_driver_products"`
		} `json:"promo"`
		Standard struct {
			MapMarkerImage struct {
				Sources []struct {
					Media struct {
						UserInterfaceStyle int `json:"user_interface_style"`
					} `json:"media"`
					URL string `json:"url"`
				} `json:"sources"`
			} `json:"map_marker_image"`
			NearbyDriverProducts []struct {
				DriverID string `json:"driver_id"`
			} `json:"nearby_driver_products"`
		} `json:"standard"`
		StandardSaver struct {
			MapMarkerImage struct {
				Sources []struct {
					Media struct {
						UserInterfaceStyle int `json:"user_interface_style"`
					} `json:"media"`
					URL string `json:"url"`
				} `json:"sources"`
			} `json:"map_marker_image"`
			NearbyDriverProducts []struct {
				DriverID string `json:"driver_id"`
			} `json:"nearby_driver_products"`
		} `json:"standard_saver"`
	} `json:"nearby_drivers_by_stable_offer_product_id"`
}

type NearbyDriversInfo struct {
	DefaultNearbyDrivers struct {
		MapMarkerImage struct {
			Sources []struct {
				Media struct {
					UserInterfaceStyle int
				}
				URL string
			}
		}
		NearbyDriverProducts []struct {
			DriverID string
		}
	}
	NearbyDrivers                       map[string]nearbyDriversInfoPayloadNearbyDriver
	NearbyDriversByStableOfferProductID struct {
		Access struct {
			MapMarkerImage struct {
				Sources []struct {
					Media struct {
						UserInterfaceStyle int
					}
					URL string
				}
			}
			NearbyDriverProducts []struct {
				DriverID string
			}
		}
		Lux struct {
			MapMarkerImage struct {
				Sources []struct {
					Media struct {
						UserInterfaceStyle int
					}
					URL string
				}
			}
			NearbyDriverProducts []struct {
				DriverID string
			}
		}
		Luxsuv struct {
			MapMarkerImage struct {
				Sources []struct {
					Media struct {
						UserInterfaceStyle int
					}
					URL string
				}
			}
			NearbyDriverProducts []struct {
				DriverID string
			}
		}
		Plus struct {
			MapMarkerImage struct {
				Sources []struct {
					Media struct {
						UserInterfaceStyle int
					}
					URL string
				}
			}
			NearbyDriverProducts []struct {
				DriverID string
			}
		}
		Promo struct {
			MapMarkerImage struct {
				Sources []struct {
					Media struct {
						UserInterfaceStyle int
					}
					URL string
				}
			}
			NearbyDriverProducts []interface{}
		}
		Standard struct {
			MapMarkerImage struct {
				Sources []struct {
					Media struct {
						UserInterfaceStyle int
					}
					URL string
				}
			}
			NearbyDriverProducts []struct {
				DriverID string
			}
		}
		StandardSaver struct {
			MapMarkerImage struct {
				Sources []struct {
					Media struct {
						UserInterfaceStyle int
					}
					URL string
				}
			}
			NearbyDriverProducts []struct {
				DriverID string
			}
		}
	}
}

//go:generate genopts --params --function NearbyDrivers --extends Base latitudeE6:int:40770034 longitudeE6:int:-73982912 orginPlaceID:string:lyft:address:3eaa5572-4d37-4a39-92ed-c61906139955 usingCommuterPayment
func (c *Client) NearbyDrivers(optss ...NearbyDriversOption) (*NearbyDriversInfo, error) {
	opts := MakeNearbyDriversOptions(optss...)

	uri := request.MakeURL("https://ride.lyft.com/v1/nearby-drivers")

	headers := c.makeHeaders(true, opts.ToBaseOptions()...)

	type origin struct {
		LatitudeE6  int `json:"latitude_e6"`
		LongitudeE6 int `json:"longitude_e6"`
	}
	type body struct {
		Origin               origin `json:"origin"`
		OriginPlaceID        string `json:"origin_place_id"`
		UsingCommuterPayment bool   `json:"using_commuter_payment"`
	}

	b, err := request.JSONMarshal(body{
		Origin: origin{
			LatitudeE6:  opts.LatitudeE6(),
			LongitudeE6: opts.LongitudeE6(),
		},
		OriginPlaceID:        opts.OrginPlaceID(),
		UsingCommuterPayment: opts.UsingCommuterPayment(),
	})
	if err != nil {
		return nil, err
	}

	var payload nearbyDriversInfoPayload

	if _, err := request.Post(uri, &payload, strings.NewReader(string(b)), request.RequestExtraHeaders(headers)); err != nil {
		return nil, err
	}

	if opts.DebugPayload() {
		log.Printf("payload: %+v", payload)
	}

	res := NearbyDriversInfo(payload)
	return &res, nil
}
