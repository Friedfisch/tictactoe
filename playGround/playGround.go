package playground

import (
	"errors"
	"fmt"
)

type Stone struct {
	X, Y int
}

type PlayGround struct {
	fields  [][]byte
	events  [][]Stone
	size    int
	players byte
}

func NewPlayGround(size int, players byte) PlayGround {
	if size < 1 {
		panic("Size must be a positive int")
	}
	var r = make([][]byte, size)
	for i := range r {
		r[i] = make([]byte, size)
	}

	var s = make([][]Stone, players)
	for i := range s {
		s[i] = make([]Stone, 0, size*size/2+1)
	}

	return PlayGround{r, s, size, players}
}

func (pg PlayGround) Size() int {
	return pg.size
}

func (pg PlayGround) Players() byte {
	return pg.players
}

func (pg PlayGround) Board() [][]byte {
	return pg.fields
}

func (pg PlayGround) Log(player byte) []Stone {
	if player < 1 || player > pg.players {
		panic("player out of range")
	}
	return pg.events[player-1]
}

func (pg PlayGround) Set(x, y int, player byte, overwrite bool) error {
	if x < 0 || x >= pg.size {
		panic("x out of range")
	}
	if y < 0 || y >= pg.size {
		panic("y out of range")
	}

	stone := Stone{x, y}
	if !overwrite {
		if pg.fields[x][y] > 0 {
			return errors.New("field in use")
		}
		// TODO overwrite check for event
		for i := range pg.events {
			for j := range pg.events[i] {
				if pg.events[i][j].X == stone.X && pg.events[i][j].Y == stone.Y {
					return errors.New("field in use")
				}
			}
		}
	}
	if player >= 1 && player <= pg.players {
		pg.events[player-1] = append(pg.events[player-1], stone)
	}
	pg.fields[x][y] = player
	return nil
}

func (pg PlayGround) Reset() {
	for i := range pg.events {
		pg.events[i] = []Stone{}
	}
	// TODO there must be some better way
	for i := 0; i < pg.size; i++ {
		for j := 0; j < pg.size; j++ {
			pg.Set(i, j, 0, true)
		}
	}
}

func (pg PlayGround) HasWonEvents(player byte) (result bool, i int, hrm string) {
	s := pg.size
	events := pg.Log(player)
	var rows = make([]int, s)
	var cols = make([]int, s)
	var ltr = 0 // Top-Left to Bottom-Right
	var rtl = 0 // Bottom-Left to Top-Right
	for _, event := range events {
		i++
		rows[event.X]++
		if rows[event.X] == s {
			return true, i, fmt.Sprintf("Hit row %d", event.X)
		}
		cols[event.Y]++
		if cols[event.Y] == s {
			return true, i, fmt.Sprintf("Hit col %d", event.Y)
		}

		// TODO Keep them in seperate cases until tests are done
		if event.X*2 == s && event.Y*2 == s {
			ltr++
			rtl++
		} else if event.X*2 >= s && event.Y*2 >= s {
			ltr++
		} else if event.X*2 >= s && event.Y*2 <= s {
			rtl++
		} else if event.X*2 <= s && event.Y*2 >= s {
			rtl++
		} else if event.X*2 <= s && event.Y*2 <= s {
			ltr++
		}
		// else: X or Y is on middle line but not in center - cant be a hit

		if ltr == s {
			return true, i, "Hit Top-Left to Bottom-Right"
		}
		if rtl == s {
			return true, i, "Hit Bottom-Left to Top-Right"
		}
	}

	return false, i, "Miss"
}

func (pg PlayGround) HasWonBoard(player byte) (result bool, i int, hrm string) {
	var s = pg.size
	var f = pg.fields
	var fr = 0
	var fc = 0
	var fd = 0
	var fg = false
	var y = 0
	for x := 0; x < s && y < s; x++ {
		i++
		fg = false
		if f[x][y] == player {
			fr++
			if fr == s {
				return true, i, fmt.Sprintf("Hit row %d", x)
			}
			if x == y {
				fg = true
			}
		}
		if f[y][x] == player {
			fc++
			if fc == s {
				return true, i, fmt.Sprintf("Hit col %d", y)
			}
			if s-1-y == x {
				fg = true
			}
		}
		if fg {
			fd++
			if fd == s {
				return true, i, "Hit diag"
			}
		}

		if x == s-1 {
			x = -1
			y++
			fr = 0
		}
	}
	return false, i, "Missed"
}
