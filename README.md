# AWS Mechanical Turk SDK for Go
[![GoDoc](https://godoc.org/github.com/jjeffery/mturk?status.svg)](https://godoc.org/github.com/jjeffery/mechturk)
[![Build Status (Linux)](https://travis-ci.org/jjeffery/mechturk.svg?branch=master)](https://travis-ci.org/jjeffery/mechturk)
[![Build status (Windows)](https://ci.appveyor.com/api/projects/status/vcmay27do6wtkrsk?svg=true)](https://ci.appveyor.com/project/jjeffery/mechturk)
[![license](http://img.shields.io/badge/license-MIT-green.svg?style=flat)](https://raw.githubusercontent.com/spkg/nullable/master/LICENSE.md)

## SDK

Package mechturk provides a Go SDK for the AWS Mechanical Turk Requester API.

Here is about the simplest example possible:

```go
package main

import (
  "log"
  "github.com/jjeffery/mechturk"
)

func main() {
	mt := mechturk.New()
	resp, err := mt.GetAccountBalance(&mechturk.GetAccountBalanceRequest{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(mechturk.Prettify(resp))
}
```

This program will produce output similar to:

```
{
  Request: {
    IsValid: "True"
  },
  AvailableBalance: {
    Amount: 0,
    CurrencyCode: "USD",
    FormattedPrice: "$0.00"
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
