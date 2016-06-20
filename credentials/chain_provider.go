package credentials

import (
	"errors"
	"log"
)

var (
	// errNoValidProvidersFoundInChain Is returned when there are no valid
	// providers in the ChainProvider.
	errNoValidProvidersFoundInChain = errors.New("no valid credential providers found in chain")
)

// A ChainProvider will search for a provider which returns credentials
// and cache that provider until Retrieve is called again.
//
// The ChainProvider provides a way of chaining multiple providers together
// which will pick the first available using priority order of the Providers
// in the list.
//
// If none of the Providers retrieve valid credentials Value, ChainProvider's
// Retrieve() will return an error.
//
// If a Provider is found which returns valid credentials Value ChainProvider
// will cache that Provider until Retrieve is called again.
type ChainProvider struct {
	Providers []Provider
}

// Logger can be set to help diagnose problems acquiring credentials
var Logger *log.Logger

// NewChainCredentials returns a pointer to a new Credentials object
// wrapping a chain of providers.
func NewChainCredentials(providers ...Provider) *Credentials {
	return NewCredentials(&ChainProvider{
		Providers: append([]Provider{}, providers...),
	})
}

// Retrieve returns the credentials value or error if no provider returned
// without error.
func (c *ChainProvider) Retrieve() (Value, error) {
	var errs []error
	var providerNames []string
	for _, p := range c.Providers {
		creds, err := p.Retrieve()
		if err == nil {
			return creds, nil
		}
		errs = append(errs, err)
		providerNames = append(providerNames, creds.ProviderName)
	}

	if Logger != nil {
		for i, err := range errs {
			Logger.Printf("credential check %d failed: %s: %v", i+1, providerNames[i], err)
		}
	}

	return Value{}, errNoValidProvidersFoundInChain
}
