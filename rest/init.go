package rest

import (
	"gopkg.in/resty.v0"
)

type ConfigMap map[string]string

type ConfStructMap map[string]interface{}

func init() {
	resty.SetRedirectPolicy(resty.FlexibleRedirectPolicy(20))
}