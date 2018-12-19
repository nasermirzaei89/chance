package main

import (
	"fmt"
	"time"

	"github.com/nasermirzaei89/chance"
)

func main() {
	ch := chance.New(chance.SetSeed(time.Now().UnixNano()))
	fmt.Println(ch.Bool(chance.SetLikelihood(10)))
	fmt.Println(ch.Character(chance.SetUseLowerCase(), chance.SetNumeric(true)))
}
