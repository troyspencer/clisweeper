package minefield

import (
	"strconv"

	"github.com/JoelOtter/termloop"
)

// Position is the (X, Y) coordinate of the tile
type Position struct {
	X, Y int
}

// Tile is an entity that will be drawn with space on all sides in a grid
type Tile struct {
	*termloop.Entity
	Position
	Color     termloop.Attr
	Flagged   bool
	Bomb      bool
	revealed  bool
	BombCount int
}

// SetColor sets the color of the tile, to be drawn on the next Tick call
func (tile *Tile) SetColor(color termloop.Attr) {
	tile.Color = color
}

func (tile *Tile) drawColor() {
	// draw tile
	width, height := tile.Size()
	if tile.revealed {
		if tile.Bomb {
			tile.SetColor(termloop.ColorBlack)
		} else {
			tile.SetColor(termloop.ColorYellow)
		}
	} else if tile.Flagged {
		tile.SetColor(termloop.ColorRed)
	} else {
		tile.SetColor(termloop.ColorWhite)
	}
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if x == 0 && y == 0 && tile.Revealed() {
				number := strconv.Itoa(tile.BombCount)
				tile.SetCell(x, y, &termloop.Cell{Bg: tile.Color, Fg: termloop.ColorBlack, Ch: rune(number[0])})
			} else {
				tile.SetCell(x, y, &termloop.Cell{Bg: tile.Color})
			}

		}
	}
}

// Tick draws the color of the tile
func (tile *Tile) Tick(event termloop.Event) {
	tile.drawColor()
}

func (tile *Tile) Revealed() bool {
	return tile.revealed
}

func (tile *Tile) Reveal() bool {
	tile.revealed = true
	return tile.Bomb
}
