package signer

import "testing"

func TestSign(t *testing.T) {
	tests := []struct {
		key       string
		service   string
		operation string
		timestamp string
		signature string
	}{
		// TODO(jpj): It would be nice to find sample test date in the AWS
		// documentation.
		{
			key:       "----SECRET-ACCESS-KEY----",
			service:   "MechanicalTurk",
			operation: "CreateHIT",
			timestamp: "2099-12-31T23:59:59Z",
			signature: "JV1Si7RZLNwczfT/jzfS8fr8dfEtkxrFmxKRoSoxaVk=",
		},
	}
	for _, tt := range tests {
		signer := New(tt.key)
		signature := signer.Sign(tt.service, tt.operation, tt.timestamp)
		if signature != tt.signature {
			t.Errorf("signature: expected=%s, actual=%s", tt.signature, signature)
		}
	}
}
