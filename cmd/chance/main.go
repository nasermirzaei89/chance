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
