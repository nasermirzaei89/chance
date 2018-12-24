package main

import (
	"fmt"
	"github.com/nasermirzaei89/chance"
)

func main() {
	ch := chance.New()
	fmt.Println(ch.Bool())
	fmt.Println(ch.String())
	fmt.Println(ch.Int(chance.SetIntMax(-5), chance.SetIntMin(5)))
}
