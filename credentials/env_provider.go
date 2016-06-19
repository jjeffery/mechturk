package credentials

import (
	"errors"
	"os"
)

// EnvProviderName provides a name of Env provider
const EnvProviderName = "EnvProvider"

// A EnvProvider retrieves credentials from the environment variables of the
// running process. Environment credentials never expire.
//
// Environment variables used:
//
// * Access Key ID:     AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY
// * Secret Access Key: AWS_SECRET_ACCESS_KEY or AWS_SECRET_KEY
type EnvProvider struct{}

// NewEnvCredentials returns a pointer to a new Credentials object
// wrapping the environment variable provider.
func NewEnvCredentials() *Credentials {
	return NewCredentials(&EnvProvider{})
}

// Retrieve retrieves the keys from the environment.
func (e *EnvProvider) Retrieve() (Value, error) {
	id := os.Getenv("AWS_ACCESS_KEY_ID")
	if id == "" {
		id = os.Getenv("AWS_ACCESS_KEY")
	}

	secret := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if secret == "" {
		secret = os.Getenv("AWS_SECRET_KEY")
	}

	if id == "" {
		return Value{ProviderName: EnvProviderName}, errors.New("AWS_ACCESS_KEY_ID or AWS_ACCESS_KEY not found in environment")
	}

	if secret == "" {
		return Value{ProviderName: EnvProviderName}, errors.New("AWS_SECRET_ACCESS_KEY or AWS_SECRET_KEY not found in environment")
	}

	return Value{
		AccessKeyID:     id,
		SecretAccessKey: secret,
		ProviderName:    EnvProviderName,
	}, nil
}
