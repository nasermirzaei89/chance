# Chance

[![Build Status](https://travis-ci.org/nasermirzaei89/chance.svg?branch=master)](https://travis-ci.org/nasermirzaei89/chance)
[![Go Report Card](https://goreportcard.com/badge/github.com/nasermirzaei89/chance)](https://goreportcard.com/report/github.com/nasermirzaei89/chance)
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

## Basics

### Bool

Returns a random boolean value (`true` or `false`).

```go
// usages
chance.Bool()
```

```go
chance.Bool() // => true
```

### String

Returns a random string value

```go
// usages
chance.String()
chance.String(chance.SetStringLength(12))
```

```go
chance.String() // => "sV8YP"
```

```go
chance.String(chance.SetStringLength(12)) // => "sOG1KmMz]wZ#"
```

Returns a random string with length of 12

```go
chance.String(chance.SetStringPool("aB1")) // => "aaBa1"
```

Returns a random string from custom pool
