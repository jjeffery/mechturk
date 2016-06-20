package credentials

// Default credentials for use when credentials are not explicitly
// specified.
var Default = NewChainCredentials(&EnvProvider{}, &SharedCredentialsProvider{})

// A Provider is the interface for any component which will provide credentials.
type Provider interface {
	// Retrieve returns the credentials value, or an error if the
	// credentials cannot be obtained.
	Retrieve() (Value, error)
}

// Value contains the individual AWS credentials fields.
// Note that AWS Mechanical Turk does not support temporary
// credentials, which means that there is no session token.
type Value struct {
	// AWS Access Key ID
	AccessKeyID string

	// AWS Secret Access Key
	SecretAccessKey string

	// Name of the provider that provided these credentials.
	ProviderName string
}

// A Credentials provides synchronous safe retrieval of AWS credentials Value.
//
// This model for credentials is based on the AWS SDK for Go. Note, however,
// that AWS Mechanical Turk does not support temporary credentials, so this
// implementation is significantly simpler.
type Credentials struct {
	provider Provider
}

// NewCredentials returns a pointer to a new Credentials with the provider set.
func NewCredentials(provider Provider) *Credentials {
	if provider == nil {
		provider = &EnvProvider{}
	}
	return &Credentials{
		provider: provider,
	}
}

// Get returns the credentials value, or error if the credentials Value failed
// to be retrieved.
func (c *Credentials) Get() (Value, error) {
	return c.provider.Retrieve()
}
