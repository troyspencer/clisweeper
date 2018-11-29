package minefield

import (
	"github.com/JoelOtter/termloop"
	"github.com/troyspencer/clisweeper/minetile"
)

// Field holds the tiles, the selected tile, and the background
type Field struct {
	Tiles        [][]*minetile.Tile
	SelectedTile minetile.Position
	Background   *termloop.Rectangle
	Selection    *termloop.Rectangle
	tileWidth    int
	tileHeight   int
}

// Tick checks for arrow key events, updates SelectedTile, and sets the color of the tiles in the field
func (field *Field) Tick(event termloop.Event) {
	if event.Type == termloop.EventKey {
		switch event.Key {
		case termloop.KeyArrowRight:
			field.SelectedTile.X++
			if field.SelectedTile.X >= len(field.Tiles[0]) {
				field.SelectedTile.X = len(field.Tiles[0]) - 1
			}
		case termloop.KeyArrowLeft:
			field.SelectedTile.X--
			if field.SelectedTile.X < 0 {
				field.SelectedTile.X = 0
			}
		case termloop.KeyArrowUp:
			field.SelectedTile.Y--
			if field.SelectedTile.Y < 0 {
				field.SelectedTile.Y = 0
			}
		case termloop.KeyArrowDown:
			field.SelectedTile.Y++
			if field.SelectedTile.Y >= len(field.Tiles[0]) {
				field.SelectedTile.Y = len(field.Tiles[0]) - 1
			}
		case termloop.KeySpace:
			field.Tiles[field.SelectedTile.X][field.SelectedTile.Y].Flagged = !field.Tiles[field.SelectedTile.X][field.SelectedTile.Y].Flagged
		}
	}
	for tileX := 0; tileX < len(field.Tiles); tileX++ {
		for tileY := 0; tileY < len(field.Tiles[0]); tileY++ {
			if field.Tiles[tileX][tileY].Flagged {
				field.Tiles[tileX][tileY].SetColor(termloop.ColorGreen)
			} else {
				field.Tiles[tileX][tileY].SetColor(termloop.ColorWhite)
			}
		}
	}

	field.Selection.SetPosition(field.SelectedTile.X*field.tileWidth*2, field.SelectedTile.Y*field.tileHeight*2)

}

// Draw is left empty as the field itself does not need to draw,
// but does need to respond to key events,
// so it implements termloop.Drawable to achieve this
func (field *Field) Draw(screen *termloop.Screen) {

}

// New initializes a background entity, tile entities, and the selected tile, then returns the collected fields in a Field struct pointer
func New(width int, height int) *Field {
	tileSize := 1
	tileWidth := tileSize * 2
	tileHeight := tileSize

	field := &Field{}

	field.tileHeight = tileSize
	field.tileWidth = tileSize * 2

	field.Background = termloop.NewRectangle(0, 0, (width)*tileWidth*2+tileWidth, (height)*tileHeight*2+tileHeight, termloop.ColorBlue)
	field.SelectedTile = minetile.Position{X: 0, Y: 0}

	field.Selection = termloop.NewRectangle(0, 0, tileWidth*3, tileHeight*3, termloop.ColorRed)

	field.Tiles = make([][]*minetile.Tile, width)
	for column := range field.Tiles {
		field.Tiles[column] = make([]*minetile.Tile, height)
	}

	// create tiles
	for tileX := 0; tileX < width; tileX++ {
		for tileY := 0; tileY < height; tileY++ {
			// create tile
			tile := minetile.Tile{
				Entity:   termloop.NewEntity(2*tileWidth*(tileX)+tileWidth, 2*tileHeight*(tileY)+tileHeight, tileWidth, tileHeight),
				Position: minetile.Position{X: tileX, Y: tileY},
			}

			tile.SetColor(termloop.ColorWhite)
			// add tile to field
			field.Tiles[tileX][tileY] = &tile
		}
	}

	return field
}
