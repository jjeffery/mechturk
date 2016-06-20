package credentials

import (
	"github.com/pkg/errors"
)

// StaticProviderName provides a name of Static provider
const StaticProviderName = "StaticProvider"

var (
	// errStaticCredentialsEmpty is emitted when static credentials are empty.
	errStaticCredentialsEmpty = errors.New("static credentials are empty")
)

// A StaticProvider is a set of credentials which are set programmatically,
// and will never expire.
type StaticProvider struct {
	Value
}

// NewStaticCredentials returns a pointer to a new Credentials object
// wrapping a static credentials value provider.
func NewStaticCredentials(id, secret string) *Credentials {
	return NewCredentials(&StaticProvider{Value: Value{
		AccessKeyID:     id,
		SecretAccessKey: secret,
	}})
}

// Retrieve returns the credentials or error if the credentials are invalid.
func (s *StaticProvider) Retrieve() (Value, error) {
	if s.AccessKeyID == "" || s.SecretAccessKey == "" {
		return Value{ProviderName: StaticProviderName}, errStaticCredentialsEmpty
	}

	s.Value.ProviderName = StaticProviderName
	return s.Value, nil
}
