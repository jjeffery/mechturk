package mechturk

import (
	"github.com/aws/aws-sdk-go/aws/awsutil"
)

// Prettify returns a string representation of the data
// structure v.
//
// The current implementation uses the `awsutil.Prettify`
// function from the AWS SDK for Go (github.com/aws/aws-sdk-go).
func Prettify(v interface{}) string {
	return awsutil.Prettify(v)
}
