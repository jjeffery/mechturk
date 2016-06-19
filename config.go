package mturk

import (
	"github.com/jjeffery/mturk/credentials"
	"log"
)

// DefaultConfig is the default configuration used if not
// explicitly specified.
var DefaultConfig = &Config{}

// A Config provides configuration for AWS Mechanical Turk clients.
type Config struct {
	// The credentials object to use when signing requests. Defaults to
	// a chain of credential providers to search for credentials in environment
	// variables and the shared credential file.
	Credentials *credentials.Credentials

	// Sandbox indicates whether to use the AWS Mechanical Turk Sandbox.
	// If Endpoint is set, then it overrides this setting.
	Sandbox *bool

	// Endpoint for accessing Mechanical Turk. If set, overrides
	// the Sandbox setting.
	Endpoint *string

	// Logger for diagnostics
	Logger *log.Logger
}

// Clone returns a deep copy of the Config.
func (c *Config) Clone() *Config {
	if c == nil {
		return &Config{}
	}
	clone := *c
	return &clone
}

// Merge returns a deep copy of the config merged with
// the other config.
func (c *Config) Merge(other *Config) *Config {
	merged := c.Clone()
	if other != nil {
		if other.Credentials != nil {
			merged.Credentials = other.Credentials
		}
		if other.Endpoint != nil {
			merged.Endpoint = other.Endpoint
		}
		if other.Sandbox != nil {
			merged.Sandbox = other.Sandbox
		}
		if other.Logger != nil {
			merged.Logger = other.Logger
		}
	}
	return merged
}

// WithCredentials returns a copy of the config with the
// credentials set.
func (c *Config) WithCredentials(cred *credentials.Credentials) *Config {
	c2 := c.Clone()
	c2.Credentials = cred
	return c2
}

// WithSandbox returns a copy of the config with the sandbox
// setting.
func (c *Config) WithSandbox(sandbox bool) *Config {
	c2 := c.Clone()
	c2.Sandbox = &sandbox
	return c2
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
