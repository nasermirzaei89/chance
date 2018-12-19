# Chance

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
    fmt.Println(ch.Bool(chance.SetLikelihood(10)))
    fmt.Println(ch.Character(chance.SetAlpha(true), chance.SetUseLowerCase(), chance.SetNumeric(true)))
}
```

## Basics

### Bool

```go
// usages
chance.bool()
chance.bool({ likelihood: 30 })
```

Return a random boolean value (`true` or `false`).

```go
chance.Bool() // => true
```

The default likelihood of success (returning `true`) is 50%. Can optionally specify the likelihood in percent:

```go
chance.Bool(chance.SetLikelihood(30)) // => false
```

In this case only a 30% likelihood of `true`, and a 70% likelihood of `false`.

### Character

```go
// usages
chance.Character()
chance.character(chance.SetPool("abcde"))
chance.character(chance.SetAlpha(true))
chance.character(chance.SetUseLowerCase())
chance.character(chance.SetSymbols(true))
```

Return a random character.

```go
chance.Character() // => "v"
```

By default it will return a string with random character from the following pool.

```go
"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
```

Optionally specify a pool and the character will be generated with characters only from that pool.

```go
chance.character(chance.SetPool("abcde")) // => "c"
```

Optionally specify alpha for only an alphanumeric character.

```go
chance.character(chance.SetAlpha(true)) // => "N"
```

Default includes both upper and lower case. It's possible to specify one or the other.

```go
chance.character(chance.SetUseLowerCase()) // => "j"
```

Note, wanted to call this key just `case` but unfortunately that's a reserved word in JavaScript for use in a switch statement

Optionally return only symbols

```go
chance.character(chance.SetSymbols(true)) // => '%'
```
