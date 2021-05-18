# go-datagovgr

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://raw.githubusercontent.com/chanioxaris/go-datagovgr/master/LICENSE)
[![GoDoc](https://godoc.org/github.com/chanioxaris/go-opap?status.svg)](https://pkg.go.dev/github.com/chanioxaris/go-datagovgr)
[![codecov](https://codecov.io/gh/chanioxaris/go-datagovgr/branch/master/graph/badge.svg?token=CZKBOS132N)](https://codecov.io/gh/chanioxaris/go-datagovgr)
[![goReportCard](https://goreportcard.com/badge/github.com/chanioxaris/go-opap)](https://goreportcard.com/report/github.com/chanioxaris/go-datagovgr)

An open source Golang SDK client to access the available open data from [data.gov.gr](https://www.data.gov.gr/), 
provided from the Greek government.

## Getting started

### API Token

First you need to obtain an API Token to access and use the web services of the data.gov.gr developer network. 
You can get one [here](https://www.data.gov.gr/token/).

### Install
`$ go get github.com/chanioxaris/go-datagovgr`

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