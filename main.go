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
		config := minefield.Configuration{
			Height: 10,
			Width: 10,
			Bombs: 20,
		}

		if c.String("size") != "" {
			num, err := strconv.Atoi(c.String("size"))
			if err != nil {
				log.Fatal("size must be a positive integer")
			}
			if num <= 0 {
				log.Fatal("size must be a positive integer")
			}
			config.Height, config.Width = num, num

			// set bombs to be 1/5 of total tiles if unset in -bombs
			config.Bombs = config.Height * config.Width / 5
		}
		if c.String("bombs") != "" {
			num, err := strconv.Atoi(c.String("bombs"))
			if err != nil {
				log.Fatal("bombs must be a positive integer")
			}
			if num <= 0 {
				log.Fatal("bombs must be a positive integer")
			}
			config.Bombs = num
		}

		playGame(config)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func playGame(config minefield.Configuration) {
	game := termloop.NewGame()
	level := termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorDefault,
	})

	// create field
	field := minefield.New(config)

	// add background to level
	level.AddEntity(field.Background)

	// add field to level
	level.AddEntity(field)

	// add selection to level
	level.AddEntity(field.Selection)

	// add tiles to level
	for i := 0; i < len(field.Tiles); i++ {
		for j := 0; j < len(field.Tiles[0]); j++ {
			level.AddEntity(field.Tiles[i][j])
		}
	}

	game.Screen().SetLevel(level)
	game.Start()
}
 