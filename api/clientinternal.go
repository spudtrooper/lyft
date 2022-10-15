package api

import (
	"github.com/spudtrooper/goutil/or"
	"github.com/spudtrooper/goutil/request"
)

func (c *Client) makeHeaders(auth bool, optss ...BaseOption) map[string]string {
	opts := MakeBaseOptions(optss...)

	headers := map[string]string{
		"authority":          `ride.lyft.com`,
		"accept":             `application/json`,
		"accept-language":    `en-US,en;q=0.9`,
		"cache-control":      `no-cache`,
		"dnt":                `1`,
		"downlink":           `10`,
		"dpr":                `2`,
		"pragma":             `no-cache`,
		"referer":            `https://ride.lyft.com/?origin=40.770034%2C-73.982912&originName=24+W+61st+St`,
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
	}
	if auth {
		token := or.String(opts.Token(), c.token)
		headers["cookie"] = request.CreateCookie([][2]string{
			{"lyftAccessToken", token},
		})
	}

	return headers
}
