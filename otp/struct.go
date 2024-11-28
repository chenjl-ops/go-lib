package otp

import "net/url"

type Key struct {
	orig string
	url  *url.URL
}
