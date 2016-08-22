package mechturk

import (
	"bytes"
	"errors"
	"strings"
)

type errorT struct {
	code    string
	message string
	keyvals []*KeyValuePair
}

func newError(errs *Errors) error {
	return &errorT{
		code:    errs.Error.Code,
		message: errs.Error.Message,
		keyvals: errs.Error.Data,
	}
}

func (err *errorT) Error() string {
	var buf bytes.Buffer
	if err.code != "" {
		buf.WriteString(err.code)
		buf.WriteString(": ")
	}
	buf.WriteString(err.message)
	for _, kv := range err.keyvals {
		buf.WriteRune(' ')
		buf.WriteString(kv.Key)
		buf.WriteRune('=')
		buf.WriteString(kv.Value)
	}
	return buf.String()
}

func (err *errorT) Code() string {
	return err.code
}

func checkRequest(request *Request) error {
	if request == nil {
		return errors.New("missing <Request> element in response")
	}
	if strings.ToLower(request.IsValid) == "true" {
		return nil
	}
	if request.Errors == nil {
		return errors.New("invalid request")
	}
	return newError(request.Errors)
}

func checkOperationRequest(opreq *OperationRequest) error {
	if opreq == nil {
		return errors.New("missing <OperationRequest> element in response")
	}
	if opreq.Errors == nil {
		return nil
	}
	return newError(opreq.Errors)
}
