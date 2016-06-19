// Package signer provides a signer type that signs
// Mechanical Turk requests.
package signer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// Signer signs AWS Mechanical Turk requests.
// The byte array representation is the AWS secret access key.
type Signer []byte

// New returns a new signer for the secret access key.
func New(key string) Signer {
	return Signer(key)
}

// Sign returns a signature based on the service, operation and
// timestamp.
func (s Signer) Sign(service, operation, timestamp string) string {
	hash := hmac.New(sha256.New, []byte(s))
	hash.Write([]byte(service))
	hash.Write([]byte(operation))
	hash.Write([]byte(timestamp))
	data := hash.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(data)
	return signature
}
