package mturk

import (
	"log"
	"github.com/jjeffery/mturk/credentials"
)

// A Config provides configuration for AWS Mechanical Turk clients.
type Config struct {
	// The credentials object to use when signing requests. Defaults to
	// a chain of credential providers to search for credentials in environment
	// variables and the shared credential file.
	Credentials *credentials.Credentials

	// Endpoint for accessing Mechanical Turk.
	Endpoint *string

	// Sandbox indicates whether to use the AWS Mechanical Turk Sandbox.
	// If Endpoint is set, then it overrides this setting.
	Sandbox *bool

	// Logger for diagnostics
	Logger *log.Logger
}

// getEndpoint returns the URL for sending requests to.
func (c *Config) getEndpoint() string {
	if c.Endpoint != nil {
		return *c.Endpoint
	}
	sandbox := c.Sandbox != nil && *c.Sandbox
	if sandbox {
		return "https://mechanicalturk.sandbox.amazonaws.com/?Service=AWSMechanicalTurkRequester"
	}
	return "https://mechanicalturk.amazonaws.com/?Service=AWSMechanicalTurkRequester"
}

func (c *Config) getCredentials() *credentials.Credentials {
	if c.Credentials != nil {
		return c.Credentials
	}
	return credentials.Default
}
