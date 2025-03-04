package client_code

import (
	"testing"

	datadog "github.com/GlobalFreightSolutions/logrus-datadog-hook"
)

func TestEndpointIsPubliclyAccessible(t *testing.T) {
	// Arrange
	var endpoint datadog.Endpoint = "http://some-injected-endpoint.com"

	// Act
	options := &datadog.Options{
		DatadogEndpoint: &endpoint,
	}

	// Assert
	if *options.DatadogEndpoint != endpoint {
		t.Error("Endpoint is not publicly accessible")
	}
}
