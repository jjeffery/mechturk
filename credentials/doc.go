// Package credentials provides credential retrieval and management.
//
// Almost all of the code in this directory has been adapted from the
// AWS SDK for Go. https://github.com/aws/aws-sdk-go.
//
// A key simplification is that Mechanical Turk does not support temporary
// credentials (session tokens), and so credentials do not need to support
// expiry.
package credentials
