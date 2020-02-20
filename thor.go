package main

import (
	"fmt"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 * ---
 * Hint: You can use the debug stream to print initialTX and initialTY, if Thor seems not follow your orders.
 **/

func main() {
	// lightX: the X position of the light of power
	// lightY: the Y position of the light of power
	// initialTx: Thor's starting X position
	// initialTy: Thor's starting Y position
	var lightX, lightY, initialTx, initialTy int
	fmt.Scan(&lightX, &lightY, &initialTx, &initialTy)
	const (
		N  string = "N"
		NE string = "NE"
		E  string = "E"
		SE string = "SE"
		S  string = "S"
		SW string = "SW"
		W  string = "W"
		NW string = "NW"
	)
	var turn string

	remainX := lightX - initialTx
	remainY := lightY - initialTy
	for {
		var remainingTurns int
		fmt.Scan(&remainingTurns)

		if remainX > 0 && remainY > 0 {
			turn = SE
			remainX--
			remainY--
		} else if remainX > 0 && remainY < 0 {
			turn = NE
			remainX--
			remainY++
		} else if remainX < 0 && remainY < 0 {
			turn = NW
			remainX++
			remainY++
		} else if remainX < 0 && remainY > 0 {
			turn = SW
			remainX++
			remainY--
		} else if remainX == 0 && remainY > 0 {
			turn = S
			remainY--
		} else if remainX > 0 && remainY == 0 {
			turn = E
			remainX--
		} else if remainX == 0 && remainY < 0 {
			turn = N
			remainY++
		} else if remainX < 0 && remainY == 0 {
			turn = W
			remainX++
		}

		fmt.Println(turn)
	}
}
