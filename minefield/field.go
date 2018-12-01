package minefield

import (
	"math/rand"
	"time"

	"github.com/JoelOtter/termloop"
)

type Configuration struct {
	Width  int
	Height int
	Bombs  int
}

// Field holds the tiles, the selected tile, and the background
type Field struct {
	Tiles        [][]*Tile
	SelectedTile *Tile
	Background   *termloop.Rectangle
	Selection    *termloop.Rectangle
	tileWidth    int
	tileHeight   int
}

// New initializes a background entity, tile entities, and the selected tile, then returns the collected fields in a Field struct pointer
func New(config Configuration) *Field {
	tileSize := 2
	tileWidth := tileSize * 2
	tileHeight := tileSize

	field := &Field{}

	field.tileHeight = tileSize
	field.tileWidth = tileSize * 2

	field.Background = termloop.NewRectangle(0, 0, (config.Width)*tileWidth*2+tileWidth, (config.Height)*tileHeight*2+tileHeight, termloop.ColorBlue)
	field.Selection = termloop.NewRectangle(0, 0, tileWidth*3, tileHeight*3, termloop.ColorCyan)

	field.Tiles = make([][]*Tile, config.Width)
	for column := range field.Tiles {
		field.Tiles[column] = make([]*Tile, config.Height)
	}

	// create tiles
	for tileX := 0; tileX < config.Width; tileX++ {
		for tileY := 0; tileY < config.Height; tileY++ {
			// add tile to field
			field.Tiles[tileX][tileY] = &Tile{
				Entity:   termloop.NewEntity(2*tileWidth*(tileX)+tileWidth, 2*tileHeight*(tileY)+tileHeight, tileWidth, tileHeight),
				Position: Position{X: tileX, Y: tileY},
			}
		}
	}

	field.setBombs(config.Bombs)

	// set selected tile as tile at (0,0)
	field.SelectedTile = field.Tiles[0][0]

	return field
}

// Tick checks for arrow key events, updates SelectedTile, and sets the color of the tiles in the field
func (field *Field) Tick(event termloop.Event) {
	if event.Type == termloop.EventKey {
		switch event.Key {
		case termloop.KeyArrowRight:
			if field.SelectedTile.X < len(field.Tiles)-1 {
				field.SelectedTile = field.Tiles[field.SelectedTile.X+1][field.SelectedTile.Y]
			}
		case termloop.KeyArrowLeft:
			if field.SelectedTile.X > 0 {
				field.SelectedTile = field.Tiles[field.SelectedTile.X-1][field.SelectedTile.Y]
			}
		case termloop.KeyArrowUp:
			if field.SelectedTile.Y > 0 {
				field.SelectedTile = field.Tiles[field.SelectedTile.X][field.SelectedTile.Y-1]
			}
		case termloop.KeyArrowDown:
			if field.SelectedTile.Y < len(field.Tiles[0])-1 {
				field.SelectedTile = field.Tiles[field.SelectedTile.X][field.SelectedTile.Y+1]
			}
		case termloop.KeySpace:
			field.SelectedTile.Flagged = !field.SelectedTile.Flagged
		case termloop.KeyEnter:
			if field.SelectedTile.Reveal() {
				field.revealAll()
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

func (field *Field) setBombs(bombs int) {
	numTiles := len(field.Tiles) * len(field.Tiles[0])
	if bombs > numTiles {
		// all are bombs, game is lost
	}

	hash := make(map[int]int)
	rand.Seed(time.Now().Unix())

	// generate bombs using Fisher-Yates shuffle
	for i := 0; i < bombs; i++ {
		j := i + rand.Intn(numTiles-i)
		value, ok := hash[j]
		if ok {
			q, r := divmod(value, len(field.Tiles))
			field.Tiles[q][r].Bomb = true
			delete(hash, j)
		} else {
			q, r := divmod(j, len(field.Tiles))
			field.Tiles[q][r].Bomb = true
		}
		if j > i {
			value, ok := hash[i]
			if ok {
				hash[j] = value
				delete(hash, i)
			} else {
				hash[j] = i
			}
		}
	}

	field.calcBombCounts()
}

func (field *Field) calcBombCounts() {
	for x := 0; x < len(field.Tiles); x++ {
		for y := 0; y < len(field.Tiles[0]); y++ {
			for tileX := x - 1; tileX < x+2; tileX++ {
				for tileY := y - 1; tileY < y+2; tileY++ {
					if tileX >= 0 && tileY >= 0 && tileX < len(field.Tiles) && tileY < len(field.Tiles[0]) && !(tileX == x && tileY == y) && field.Tiles[tileX][tileY].Bomb {
						field.Tiles[x][y].BombCount++
					}
				}
			}
		}
	}
}

func (field *Field) revealAll() {
	for x := 0; x < len(field.Tiles); x++ {
		for y := 0; y < len(field.Tiles[0]); y++ {
			if field.Tiles[x][y].Bomb {
				field.Tiles[x][y].Reveal()
			}
		}
	}
}

func divmod(numerator, denominator int) (quotient, remainder int) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}
