package playground_test

import (
	"fmt"
	"strconv"
	"testing"

	playground "github.com/Friedfisch/tictactoe/playGround"
)

var errorCount int

func setRow(t *testing.T, pg playground.PlayGround, y int, player byte) {
	t.Helper()
	for i := 0; i < pg.Size(); i++ {
		pg.Set(i, y, player)
	}
}

func setCol(t *testing.T, pg playground.PlayGround, x int, player byte) {
	t.Helper()
	for i := 0; i < pg.Size(); i++ {
		pg.Set(x, i, player)
	}
}

func setDiagLtR(t *testing.T, pg playground.PlayGround, player byte) {
	t.Helper()
	for i := 0; i < pg.Size(); i++ {
		pg.Set(i, i, player)
	}
}

func setDiagRtL(t *testing.T, pg playground.PlayGround, player byte) {
	t.Helper()
	for i := 0; i < pg.Size(); i++ {
		pg.Set(pg.Size()-1-i, i, player)
	}
}

func draw(t *testing.T, pg playground.PlayGround) string {
	t.Helper()
	var fields [][]byte = make([][]byte, pg.Size())
	for p := range fields {
		fields[p] = make([]byte, pg.Size())
	}
	for p := byte(1); p < pg.Players(); p++ {
		for _, k := range pg.Moves(p) {
			fields[k.X][k.Y] = byte(p + 1)
		}
	}
	result := ""
	for x := range fields {
		var line string
		for _, p := range fields[x] {
			c := " "
			switch p {
			case 0:
			case 1:
				c = "X"
			case 2:
				c = "O"
			default:
				c = fmt.Sprintf("%d", p)
			}
			line = line + fmt.Sprintf("\t[%s]", c)
		}
		result += fmt.Sprintf("%s\n", line)
	}
	return result
}

func assertEvents(t *testing.T, pg playground.PlayGround, name string, player byte, expected bool) bool {
	t.Helper()
	v, i, msg := pg.HasWon(player)
	fmt.Printf("P1: %d; P2: %d ; Player: %d; Won: %t; Iter: %d; Test: %s; Msg: %s\n", pg.Moves(1), pg.Moves(2), player, v, i, name, msg)
	if v != expected {
		errorCount++
		draw(t, pg)
		fmt.Printf("******************** Assertion \"%t\" failed in \"%s %d\", got \"%t\"\n", expected, name, pg.Size(), v)
	}
	pg.Reset()
	return v
}

func TestSet(t *testing.T) {
	// TODO
}

func TestReset(t *testing.T) {
	f := playground.NewPlayGround(5, 5)
	setDiagLtR(t, f, 1)
	setDiagRtL(t, f, 5)
	r := f.Moves(1)
	if len(r) != 4 {
		t.Error("There should be 4 moves for player 1")
		t.Fail()
	}
	r = f.Moves(5)
	if len(r) != 5 {
		t.Error("There should be 5 moves for player 5")
		t.Fail()
	}
	f.Reset()
	r = f.Moves(1)
	if len(r) != 0 {
		t.Error("There should be 0 moves")
		t.Fail()
	}
}

func TestHasWon(t *testing.T) {

	/*	f := playground.NewPlayGround(3, 2)
		setDiagRtL(t, f, 1)
		assertEvents(t, f, "DUT", 1, true)
		return*/

	for size := 1; size <= 17; size++ {
		f := playground.NewPlayGround(size, 2)

		// Overwrite for events is not implemented
		assertEvents(t, f, "E0", 1, false)

		f.Set(0, 0, 1)
		assertEvents(t, f, "E1", 1, size == 1)

		f.Set(size-1, size-1, 1)
		assertEvents(t, f, "E2", 1, size == 1)

		for i := 0; i < size; i++ {
			setRow(t, f, i, 1)
			assertEvents(t, f, "RT"+strconv.Itoa(i), 1, true)

			setRow(t, f, i, 1)
			f.Set(0, i, 2)
			assertEvents(t, f, "RF"+strconv.Itoa(i), 1, false)

			setCol(t, f, i, 1)
			assertEvents(t, f, "CT"+strconv.Itoa(i), 1, true)

			setCol(t, f, i, 1)
			f.Set(i, 0, 2)
			assertEvents(t, f, "CF"+strconv.Itoa(i), 1, false)
		}

		setDiagRtL(t, f, 1)
		assertEvents(t, f, "DUT", 1, true)

		setDiagRtL(t, f, 1)
		f.Set(0, size-1, 2)
		assertEvents(t, f, "DUF", 1, false)

		setDiagLtR(t, f, 1)
		assertEvents(t, f, "DDT", 1, true)

		setDiagRtL(t, f, 1)
		f.Set(size-1, 0, 2)
		assertEvents(t, f, "DDF", 1, false)

		if size == 4 {
			/*
				xxoo
				oxox
				xoxo
				ooxx
			*/
			f.Set(0, 0, 1)
			f.Set(1, 0, 1)
			f.Set(2, 0, 2)
			f.Set(3, 0, 2)

			f.Set(0, 1, 2)
			f.Set(1, 1, 1)
			f.Set(2, 1, 2)
			f.Set(3, 1, 1)

			f.Set(0, 2, 1)
			f.Set(1, 2, 2)
			f.Set(2, 2, 1)
			f.Set(3, 2, 2)

			f.Set(0, 3, 2)
			f.Set(1, 3, 2)
			f.Set(2, 3, 1)
			f.Set(3, 3, 1)

			assertEvents(t, f, "MIX", 1, true)
		}
	}
	fmt.Printf("ERRORCOUNT: %d\n", errorCount)
}
