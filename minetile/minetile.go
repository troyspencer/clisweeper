package minetile

import "github.com/JoelOtter/termloop"

// Position is the (X, Y) coordinate of the tile
type Position struct {
	X, Y int
}

// Tile is an entity that will be drawn with space on all sides in a grid
type Tile struct {
	*termloop.Entity
	Position
	Color termloop.Attr
}

// SetColor sets the color of the tile, to be drawn on the next Tick call
func (tile *Tile) SetColor(color termloop.Attr) {
	tile.Color = color
}

func (tile *Tile) drawColor() {
	// draw tile
	width, height := tile.Size()
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			tile.SetCell(x, y, &termloop.Cell{Bg: tile.Color})
		}
	}
}

// Tick draws the color of the tile
func (tile *Tile) Tick(event termloop.Event) {
	tile.drawColor()
}
