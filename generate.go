package mturk

// This documents the command used to generate the WSDL types.
// See comment in wsdltypes.go regarding hand-editing of generated output.

//go:generate gowsdl -p mturk -o wsdltypes.go AWSMechanicalTurkRequester-2012-03-25.wsdl
