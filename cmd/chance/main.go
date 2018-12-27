package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nasermirzaei89/chance"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "chance"
	app.Usage = "Random generator"

	var seed int64
	app.Flags = []cli.Flag{
		cli.Int64Flag{
			Name:        "seed",
			Usage:       "seed for random generator",
			Value:       1,
			Destination: &seed,
		},
	}

	app.Commands = []cli.Command{
		cmdBool(seed),
		cmdString(seed),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func cmdBool(seed int64) cli.Command {
	return cli.Command{
		Name:  "bool",
		Usage: "Generates a random bool value",
		Action: func(c *cli.Context) error {
			fmt.Print(chance.New(chance.SetSeed(seed)).Bool())
			return nil
		},
	}
}

func cmdString(seed int64) cli.Command {
	var length int
	var pool string
	return cli.Command{
		Name:  "string",
		Usage: "Generates a random string value",
		Flags: []cli.Flag{
			cli.IntFlag{
				Name:        "length",
				Usage:       "Sets length of random string",
				Value:       5,
				Destination: &length,
			},
			cli.StringFlag{
				Name:        "pool",
				Usage:       "Sets pool for random string",
				Value:       "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()[]",
				Destination: &pool,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Print(chance.New(chance.SetSeed(seed)).String(chance.SetStringLength(length), chance.SetStringPool(pool)))
			return nil
		},
	}
}
