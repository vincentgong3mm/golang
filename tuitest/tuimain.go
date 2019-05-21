package main

import tl "github.com/JoelOtter/termloop"

// Player is
type Player struct {
	*tl.Entity
	id int
}

// Tick is
func (player *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey { // Is it a keyboard event?

		if player.id == 1 {
			x, y := player.Position()
			switch event.Key { // If so, switch on the pressed key.
			case tl.KeyArrowRight:
				player.SetPosition(x+1, y)
			case tl.KeyArrowLeft:
				player.SetPosition(x-1, y)
				/*
					case tl.KeyArrowUp:
						player.SetPosition(x, y-1)
					case tl.KeyArrowDown:
						player.SetPosition(x, y+1)
				*/
			}
		} else {
			x, y := player.Position()
			switch event.Key { // If so, switch on the pressed key.
			case tl.KeyArrowUp:
				player.SetPosition(x+1, y)
			case tl.KeyArrowDown:
				player.SetPosition(x-1, y)
			}
		}
	}
}

func main() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: 'v',
	})
	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))

	player := Player{tl.NewEntity(1, 1, 1, 1), 1}
	// Set the character at position (0, 0) on the entity.
	//player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '@'})
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '0'})
	player.SetPosition(10, 10)
	level.AddEntity(&player)

	player2 := Player{tl.NewEntity(1, 1, 1, 1), 2}
	player2.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: 'X'})
	level.AddEntity(&player2)

	game.Screen().SetLevel(level)
	game.Start()
}
