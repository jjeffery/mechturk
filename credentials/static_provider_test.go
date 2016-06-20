package credentials

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStaticProviderGet(t *testing.T) {
	s := StaticProvider{
		Value: Value{
			AccessKeyID:     "AKID",
			SecretAccessKey: "SECRET",
		},
	}

	creds, err := s.Retrieve()
	assert.Nil(t, err, "Expect no error")
	assert.Equal(t, "AKID", creds.AccessKeyID, "Expect access key ID to match")
	assert.Equal(t, "SECRET", creds.SecretAccessKey, "Expect secret access key to match")
}
