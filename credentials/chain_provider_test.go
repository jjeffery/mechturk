package credentials

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type secondStubProvider struct {
	creds   Value
	expired bool
	err     error
}

func (s *secondStubProvider) Retrieve() (Value, error) {
	s.creds.ProviderName = "secondStubProvider"
	return s.creds, s.err
}

func TestChainProviderWithNames(t *testing.T) {
	p := &ChainProvider{
		Providers: []Provider{
			&stubProvider{err: errors.New("first provider error")},
			&stubProvider{err: errors.New("second provider error")},
			&secondStubProvider{
				creds: Value{
					AccessKeyID:     "AKIF",
					SecretAccessKey: "NOSECRET",
				},
			},
			&stubProvider{
				creds: Value{
					AccessKeyID:     "AKID",
					SecretAccessKey: "SECRET",
				},
			},
		},
	}

	creds, err := p.Retrieve()
	assert.Nil(t, err, "Expect no error")
	assert.Equal(t, "secondStubProvider", creds.ProviderName, "Expect provider name to match")

	// Also check credentials
	assert.Equal(t, "AKIF", creds.AccessKeyID, "Expect access key ID to match")
	assert.Equal(t, "NOSECRET", creds.SecretAccessKey, "Expect secret access key to match")
}

func TestChainProviderGet(t *testing.T) {
	p := &ChainProvider{
		Providers: []Provider{
			&stubProvider{err: errors.New("first provider error")},
			&stubProvider{err: errors.New("second provider error")},
			&stubProvider{
				creds: Value{
					AccessKeyID:     "AKID",
					SecretAccessKey: "SECRET",
				},
			},
		},
	}

	creds, err := p.Retrieve()
	assert.Nil(t, err, "Expect no error")
	assert.Equal(t, "AKID", creds.AccessKeyID, "Expect access key ID to match")
	assert.Equal(t, "SECRET", creds.SecretAccessKey, "Expect secret access key to match")
}

func TestChainProviderWithNoProvider(t *testing.T) {
	p := &ChainProvider{
		Providers: []Provider{},
	}

	_, err := p.Retrieve()
	assert.Equal(t,
		errNoValidProvidersFoundInChain,
		err,
		"Expect no providers error returned")
}

func TestChainProviderWithNoValidProvider(t *testing.T) {
	errs := []error{
		errors.New("first provider error"),
		errors.New("second provider error"),
	}
	p := &ChainProvider{
		Providers: []Provider{
			&stubProvider{err: errs[0]},
			&stubProvider{err: errs[1]},
		},
	}

	_, err := p.Retrieve()

	assert.Equal(t,
		errNoValidProvidersFoundInChain,
		err,
		"Expect no providers error returned")
}
