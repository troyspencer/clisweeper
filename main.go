package main

import (
	"log"
	"os"
	"strconv"

	"github.com/troyspencer/clisweeper/minefield"
	"github.com/JoelOtter/termloop"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "clisweeper"
	app.Usage = "play minesweeper with CLI"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "bombs, b",
			Usage: "Creates a minefield with b bombs, defaults to 20",
		},
		cli.StringFlag{
			Name:  "size, s",
			Usage: "Creates a minefield of dimensions size x size, defaults to 10",
		},
	}
	app.Action = func(c *cli.Context) error {
		height := 10
		width := 10
		bombs := 20

		if c.String("size") != "" {
			num, err := strconv.Atoi(c.String("size"))
			if err != nil {
				log.Fatal("size must be a positive integer")
			}
			if num <= 0 {
				log.Fatal("size must be a positive integer")
			}
			height, width = num, num


			// set bombs to be 1/5 of total tiles if unset in -bombs
			bombs= width * width / 5
		}
		if c.String("bombs") != "" {
			num, err := strconv.Atoi(c.String("bombs"))
			if err != nil {
				log.Fatal("bombs must be a positive integer")
			}
			if num <= 0 {
				log.Fatal("bombs must be a positive integer")
			}
			bombs = num
		}

		game := termloop.NewGame()
		game.Screen().SetLevel(minefield.NewLevel(height, width, bombs, 1))
		game.Start()
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}