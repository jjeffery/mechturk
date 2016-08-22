# AWS Mechanical Turk SDK for Go
[![GoDoc](https://godoc.org/github.com/jjeffery/mturk?status.svg)](https://godoc.org/github.com/jjeffery/mechturk)
[![Build Status (Linux)](https://travis-ci.org/jjeffery/mechturk.svg?branch=master)](https://travis-ci.org/jjeffery/mechturk)
[![Build status (Windows)](https://ci.appveyor.com/api/projects/status/vcmay27do6wtkrsk?svg=true)](https://ci.appveyor.com/project/jjeffery/mechturk)
[![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/spkg/nullable/master/LICENSE.md)

## SDK

Package mechturk provides a Go SDK for the AWS Mechanical Turk Requester API.

Here is a simple example that invokes the [GetAccountBalance](http://docs.aws.amazon.com/AWSMechTurk/latest/AWSMturkAPI/ApiReference_GetAccountBalanceOperation.html) operation:

```go
package main

import (
  "log"
  "github.com/jjeffery/mechturk"
)

func main() {
	mt := mechturk.New()
	result, err := mt.GetAccountBalance(&mechturk.GetAccountBalanceRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(mechturk.Prettify(result))
}
```

This program will produce output similar to:

```
{
  Request: {
    IsValid: "True"
  },
  AvailableBalance: {
    Amount: 10000,
    CurrencyCode: "USD",
    FormattedPrice: "$10,000.00"
  }
}
```

## CLI

The subdirectory `cmd/mechturk` contains a command line client. Not all Mechanical Turk
commands are supported yet -- commands are added as they are needed. Pull requests
are welcome.

Example usage:

```
$ mechturk --sandbox get-account-balance
{
  Request: {
    IsValid: "True"
  },
  AvailableBalance: {
    Amount: 10000,
    CurrencyCode: "USD",
    FormattedPrice: "$10,000.00"
  }
}
```

## API Stability

This project is still under development, and there is currently no guarantee
of API stability. If you are planning to make use of this code in any sort
of production environment, please consider
[vendoring](https://golang.org/cmd/go/#hdr-Vendor_Directories) the version
you are using with your project.
