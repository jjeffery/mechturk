package mturk

import (
	"encoding/xml"
	"time"

	"github.com/jjeffery/mechturk/credentials"
	"github.com/jjeffery/mechturk/signer"
)

// Version is the protocol version supported by this package.
const Version = "2012-03-25"

const (
	serviceName   = "AWSMechanicalTurkRequester"
	soapNamespace = "http://requester.mturk.amazonaws.com/doc/" + Version
	soapAction    = "http://soap.amazon.com"
)

type requestPayload struct {
	XMLName        xml.Name
	AWSAccessKeyID string      `xml:"AWSAccessKeyId,omitempty"`
	Timestamp      time.Time   `xml:"Timestamp,omitempty"`
	Signature      string      `xml:"Signature,omitempty"`
	Validate       string      `xml:"Validate,omitempty"`
	Credential     string      `xml:"Credential,omitempty"`
	Request        interface{} `xml:"Request"`
}

func newPayload(operation string, request interface{}, creds credentials.Value) *requestPayload {
	now := time.Now().UTC()
	signer := signer.New(creds.SecretAccessKey)
	p := &requestPayload{
		AWSAccessKeyID: creds.AccessKeyID,
		Timestamp:      now,
		Request:        request,
	}
	p.Signature = signer.Sign("AWSMechanicalTurkRequester", operation, now.Format(time.RFC3339Nano))
	p.XMLName.Space = soapNamespace
	p.XMLName.Local = operation
	return p
}
