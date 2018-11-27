package main

import (
	"log"
	"os"
	"strconv"

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
		config := Configuration{}
		config.tilesH = 10
		config.tilesW = 10
		config.bombs = 20

		if c.String("size") != "" {
			num, err := strconv.Atoi(c.String("size"))
			if err != nil {
				log.Fatal()
			}
			config.tilesH, config.tilesW = num, num
		}
		if c.String("bombs") != "" {
			num, err := strconv.Atoi(c.String("bombs"))
			if err != nil {
				log.Fatal()
			}
			config.bombs = num
		}

		playGame(config)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func playGame(config Configuration) {
	game := termloop.NewGame()
	level := termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorDefault,
	})

	tileSize := 1
	tileWidth := tileSize * 2
	tileHeight := tileSize

	// Create enough space for tiles surrounded by space
	level.AddEntity(termloop.NewRectangle(0, 0, (config.tilesW)*tileWidth*2+tileWidth, (config.tilesH)*tileHeight*2+tileHeight, termloop.ColorBlue))

	// create tiles
	for tileX := 0; tileX < config.tilesW; tileX++ {
		for tileY := 0; tileY < config.tilesH; tileY++ {
			// create tile
			tile := Tile{
				Entity: termloop.NewEntity(2*tileWidth*(tileX)+tileWidth, 2*tileHeight*(tileY)+tileHeight, tileWidth, tileHeight),
			}
			// draw tile
			for x := 0; x < tileWidth; x++ {
				for y := 0; y < tileHeight; y++ {
					tile.SetCell(x, y, &termloop.Cell{Bg: termloop.ColorWhite})
				}
			}
			// add tile to level
			level.AddEntity(&tile)
		}
	}

	game.Screen().SetLevel(level)
	game.Start()
}

type Configuration struct {
	tilesW int
	tilesH int
	bombs  int
}

// Tile is an entity that will be drawn with space on all sides in a grid
type Tile struct {
	*termloop.Entity
}
