package playground_test

import (
	"fmt"
	"strconv"
	"testing"

	playground "github.com/Friedfisch/tictactoe/playGround"
)

func assertEvents(t *testing.T, pg playground.PlayGround, name string, player byte, expected bool) bool {
	t.Helper()
	v, i, msg := pg.HasWonEvents(player)
	fmt.Printf("P1: %d; P2: %d ; Player: %d; Won: %t; Iter: %d; Test: %s; Msg: %s\n", pg.Log(1), pg.Log(2), player, v, i, name, msg)
	if v != expected {
		errorCount++
		t.Errorf("******************** Assertion \"%t\" failed in \"%s\", got \"%t\"\n", expected, name, v)
		fmt.Printf("******************** Assertion \"%t\" failed in \"%s\", got \"%t\"\n", expected, name, v)
	}
	pg.Reset()
	return v
}

func TestHasWonEvents(t *testing.T) {
	for size := 1; size < 42; size++ {
		f := playground.NewPlayGround(size, 2)

		// Overwrite for events is not implemented
		assertEvents(t, f, "E0", 1, false)

		f.Set(0, 0, 1, false)
		assertEvents(t, f, "E1", 1, size == 1)

		f.Set(size-1, size-1, 1, false)
		assertEvents(t, f, "E2", 1, size == 1)

		for i := 0; i < size; i++ {
			setRow(t, f, i, 1, false)
			assertEvents(t, f, "RT"+strconv.Itoa(i), 1, true)

			f.Set(0, i, 2, false)
			setRow(t, f, i, 1, false)
			assertEvents(t, f, "RF"+strconv.Itoa(i), 1, false)

			setCol(t, f, i, 1, false)
			assertEvents(t, f, "CT"+strconv.Itoa(i), 1, true)

			f.Set(i, 0, 2, false)
			setCol(t, f, i, 1, false)
			assertEvents(t, f, "CF"+strconv.Itoa(i), 1, false)
		}

		setDiagUp(t, f, 1, false)
		assertEvents(t, f, "DUT", 1, true)

		f.Set(0, size-1, 2, false)
		setDiagUp(t, f, 1, false)
		assertEvents(t, f, "DUF", 1, false)

		setDiagDown(t, f, 1, false)
		assertEvents(t, f, "DDT", 1, true)

		f.Set(size-1, 0, 2, false)
		setDiagUp(t, f, 1, false)
		assertEvents(t, f, "DDF", 1, false)

		if size == 4 {
			/*
				xxoo
				oxox
				xoxo
				ooxx
			*/
			f.Set(0, 0, 1, false)
			f.Set(1, 0, 1, false)
			f.Set(2, 0, 2, false)
			f.Set(3, 0, 2, false)

			f.Set(0, 1, 2, false)
			f.Set(1, 1, 1, false)
			f.Set(2, 1, 2, false)
			f.Set(3, 1, 1, false)

			f.Set(0, 2, 1, false)
			f.Set(1, 2, 2, false)
			f.Set(2, 2, 1, false)
			f.Set(3, 2, 2, false)

			f.Set(0, 3, 2, false)
			f.Set(1, 3, 2, false)
			f.Set(2, 3, 1, false)
			f.Set(3, 3, 1, false)

			assertEvents(t, f, "MIX", 1, true)
		}
	}
	fmt.Printf("ERRORCOUNT: %d\n", errorCount)
}

var errorCount int

func setRow(t *testing.T, pg playground.PlayGround, y int, player byte, overwrite bool) {
	t.Helper()
	for i := 0; i < pg.Size(); i++ {
		pg.Set(i, y, player, overwrite)
	}
}

func setCol(t *testing.T, pg playground.PlayGround, x int, player byte, overwrite bool) {
	t.Helper()
	for i := 0; i < pg.Size(); i++ {
		pg.Set(x, i, player, overwrite)
	}
}

func setDiagDown(t *testing.T, pg playground.PlayGround, player byte, overwrite bool) {
	t.Helper()
	for i := 0; i < pg.Size(); i++ {
		pg.Set(i, i, player, overwrite)
	}
}

func setDiagUp(t *testing.T, pg playground.PlayGround, player byte, overwrite bool) {
	t.Helper()
	for i := 0; i < pg.Size(); i++ {
		pg.Set(pg.Size()-1-i, i, player, overwrite)
	}
}

func assertBoard(t *testing.T, pg playground.PlayGround, name string, player byte, expected bool) bool {
	t.Helper()
	v, i, msg := pg.HasWonBoard(player)
	fmt.Printf("PG: %d; Player: %d; Won: %t; Iter: %d; Test: %s; Msg: %s\n", pg.Board(), player, v, i, name, msg)
	if v != expected {
		errorCount++
		fmt.Printf("******************** Assertion \"%t\" failed in \"%s\", got \"%t\"\n", expected, name, v)
	}
	pg.Reset()
	return v
}

func TestHasWonBoard(t *testing.T) {
	return
	/*size := 3
	f := playground.NewPlayGround(size, 3)
	assertBoard(t, f, "E0", 1, false)

	f.Set(0, 0, 1, true)
	assertBoard(t, f, "E1", 1, size == 1)

	f.Set(size-1, size-1, 1, true)
	assertBoard(t, f, "E2", 1, size == 1)

	for i := 0; i < size; i++ {
		setRow(t, f, i, 1)
		assertBoard(t, f, "RT"+strconv.Itoa(i), 1, true)

		setRow(t, f, i, 1)
		f.Set(0, i, 2, true)
		assertBoard(t, f, "RF"+strconv.Itoa(i), 1, false)

		setCol(t, f, i, 1)
		assertBoard(t, f, "CT"+strconv.Itoa(i), 1, true)

		setCol(t, f, i, 1)
		f.Set(i, 0, 2, true)
		assertBoard(t, f, "CF"+strconv.Itoa(i), 1, false)

		setCol(t, f, 0, 1)
		setRow(t, f, 0, 2)
		setCol(t, f, 1, 1)
		setRow(t, f, 1, 2)
		setCol(t, f, 2, 1)
		setRow(t, f, 2, 3)
		assertBoard(t, f, "MIX"+strconv.Itoa(i), 1, false)
	}

	setDiagUp(t, f, 1)
	assertBoard(t, f, "DUT", 1, true)

	setDiagUp(t, f, 1)
	f.Set(0, size-1, 0, true)
	assertBoard(t, f, "DUF", 1, false)

	setDiagDown(t, f, 1)
	assertBoard(t, f, "DDT", 1, true)

	setDiagUp(t, f, 1)
	f.Set(size-1, 0, 0, true)
	assertBoard(t, f, "DDF", 1, false)

	t.Logf("ERRORCOUNT: %d\n", errorCount)*/
}
