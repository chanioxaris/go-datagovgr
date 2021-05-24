# go-datagovgr

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://raw.githubusercontent.com/chanioxaris/go-datagovgr/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/chanioxaris/go-opap?status.svg)](https://pkg.go.dev/github.com/chanioxaris/go-datagovgr)
[![codecov](https://codecov.io/gh/chanioxaris/go-datagovgr/branch/master/graph/badge.svg?token=CZKBOS132N)](https://codecov.io/gh/chanioxaris/go-datagovgr)
[![goReportCard](https://goreportcard.com/badge/github.com/chanioxaris/go-opap)](https://goreportcard.com/report/github.com/chanioxaris/go-datagovgr)

A Golang SDK client to access the available open data from [data.gov.gr](https://www.data.gov.gr/), 
provided by the Greek government.

## Getting started

### API Token

First you need to obtain an API Token to access and use the web services of the [data.gov.gr](https://www.data.gov.gr/) developer network. 
You can get one [here](https://www.data.gov.gr/token/).

### Install
`$ go get github.com/chanioxaris/go-datagovgr`

## Endpoints

The client supports the following endpoints derived by topic:
- [x] [Business and Economy](https://pkg.go.dev/github.com/chanioxaris/go-datagovgr/api#BusinessEconomy)
- [x] [Crime and Justice](https://pkg.go.dev/github.com/chanioxaris/go-datagovgr/api#CrimeJustice)
- [x] [Education](https://pkg.go.dev/github.com/chanioxaris/go-datagovgr/api#Education)
- [x] [Environment](https://pkg.go.dev/github.com/chanioxaris/go-datagovgr/api#Environment)
- [x] [Health](https://pkg.go.dev/github.com/chanioxaris/go-datagovgr/api#Health)
- [x] [Society](https://pkg.go.dev/github.com/chanioxaris/go-datagovgr/api#Society)
- [x] [Technology](https://pkg.go.dev/github.com/chanioxaris/go-datagovgr/api#Technology)
- [x] [Telecommunications](https://pkg.go.dev/github.com/chanioxaris/go-datagovgr/api#Telcos)


## Usage
```golang
package main

import (
	"github.com/chanioxaris/go-datagovgr/datagovgr"
)

const token = "<YOUR_API_TOKEN>"

func main() {
	c, err := datagovgr.NewClient(token)
	if err != nil {
		// Handle error
	}
	
	// Retrieve and access data
	...
}
```

## License

go-datagovgr is [MIT licensed](LICENSE).