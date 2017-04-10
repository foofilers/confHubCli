package rest

import "gopkg.in/resty.v0"

func init() {
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(20))
}