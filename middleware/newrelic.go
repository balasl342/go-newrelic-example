package middleware

import (
	"net/http"
	"relicapp/config"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelicApp() (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppName(config.AppName),
		newrelic.ConfigLicense(config.LicenseKey),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
}

func WrapHandleFunc(app *newrelic.Application, pattern string, handler func(http.ResponseWriter, *http.Request)) (string, func(http.ResponseWriter, *http.Request)) {
	return newrelic.WrapHandleFunc(app, pattern, handler)
}
