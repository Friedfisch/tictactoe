package playground

import (
	"errors"
	"fmt"
)

type Move struct {
	X, Y int
}

type PlayGround struct {
	moves   [][]Move
	size    int
	players byte
}

func NewPlayGround(size int, players byte) PlayGround {
	if size < 1 {
		panic("Size must be a positive int")
	}

	var s = make([][]Move, players)
	for i := range s {
		s[i] = make([]Move, 0, size*size/2)
	}

	return PlayGround{s, size, players}
}

func (pg PlayGround) Size() int {
	return pg.size
}

func (pg PlayGround) Players() byte {
	return pg.players
}

func (pg PlayGround) Moves(player byte) []Move {
	if player < 1 || player > pg.players {
		panic("player out of range")
	}
	return pg.moves[player-1]
}

func (pg PlayGround) Set(x, y int, player byte) error {
	if x < 0 || x >= pg.size {
		return errors.New("x out of range")
	}
	if y < 0 || y >= pg.size {
		return errors.New("y out of range")
	}
	if player > pg.players {
		return errors.New("player out of range")
	}

	stone := Move{x, y}

	// We always overwrite, that means we need to remove it everywhere
out:
	for i := range pg.moves {
		for j := range pg.moves[i] {
			if pg.moves[i][j].X == stone.X && pg.moves[i][j].Y == stone.Y {
				pg.moves[i] = append(pg.moves[i][:j], pg.moves[i][j+1:]...)
				break out
			}
		}
	}

	if player > 0 {
		pg.moves[player-1] = append(pg.moves[player-1], stone)
	}
	return nil
}

func (pg PlayGround) Reset() {
	for i := range pg.moves {
		pg.moves[i] = []Move{}
	}
}

func (pg PlayGround) HasWon(player byte) (result bool, i int, hrm string) {
	s := pg.size
	// Counters, we search for rows or columns or diagonals where "hits" match size.
	var rows = make([]int, s)
	var cols = make([]int, s)
	var ltr = 0
	var rtl = 0
	for _, event := range pg.Moves(player) {
		i++
		rows[event.X]++
		if rows[event.X] == s {
			return true, i, fmt.Sprintf("Hit row %d", event.X)
		}
		cols[event.Y]++
		if cols[event.Y] == s {
			return true, i, fmt.Sprintf("Hit col %d", event.Y)
		}

		// We are looking for half of the board without using float that works with un/even sizes
		var x2 = max(event.X*2, s)
		var y2 = max(event.Y*2, s)

		if (x2 >= s && y2 >= s) || (x2 <= s && y2 <= s) {
			ltr++
			if ltr == s {
				return true, i, "Hit Top-Left to Bottom-Right"
			}
		}

		if (x2 >= s && y2 <= s) || (x2 <= s && y2 >= s) {
			rtl++
			if rtl == s {
				return true, i, "Hit Bottom-Left to Top-Right"
			}
		}
	}

	return false, i, "Miss"
}
