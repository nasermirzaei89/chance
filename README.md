# Chance

[![Build Status](https://travis-ci.org/nasermirzaei89/chance.svg?branch=master)](https://travis-ci.org/nasermirzaei89/chance)
[![Go Report Card](https://goreportcard.com/badge/github.com/nasermirzaei89/chance)](https://goreportcard.com/report/github.com/nasermirzaei89/chance)
[![GoDoc](https://godoc.org/github.com/nasermirzaei89/chance?status.svg)](https://godoc.org/github.com/nasermirzaei89/chance)
[![GitHub license](https://img.shields.io/github/license/nasermirzaei89/chance.svg)](https://github.com/nasermirzaei89/chance/blob/master/LICENSE)

## Instal

```sh
go get -u github.com/nasermirzaei89/chance
```

## Usage

### Sample

```go
package main

import (
    "fmt"
    "time"

    "github.com/nasermirzaei89/chance"
)

func main() {
    ch := chance.New(chance.SetSeed(time.Now().UnixNano()))
    fmt.Println(ch.Bool())
    fmt.Println(ch.String())
}
```

## Documentation
[https://godoc.org/github.com/nasermirzaei89/chance](https://godoc.org/github.com/nasermirzaei89/chance)
