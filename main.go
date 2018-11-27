package main

import "github.com/JoelOtter/termloop"

func main() {
	game := termloop.NewGame()
	level := termloop.NewBaseLevel(termloop.Cell{
		Bg: termloop.ColorDefault,
	})

	tilesW := 10
	tilesH := 10
	tileSize := 1
	tileWidth := tileSize * 2
	tileHeight := tileSize

	// Create enough space for tiles surrounded by space
	level.AddEntity(termloop.NewRectangle(0, 0, (tilesW)*tileWidth*2+tileWidth, (tilesH)*tileHeight*2+tileHeight, termloop.ColorBlue))

	// create tiles
	for tileX := 0; tileX < tilesW; tileX++ {
		for tileY := 0; tileY < tilesH; tileY++ {
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

// Tile is an entity that will be drawn with space on all sides in a grid
type Tile struct {
	*termloop.Entity
}
