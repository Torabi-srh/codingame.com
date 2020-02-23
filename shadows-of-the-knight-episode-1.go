package main

import "fmt"
import "os"
import "math"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// W: width of the building.
	// H: height of the building.
	var W, H int
	fmt.Scan(&W, &H)

	// N: maximum number of turns before game over.
	var N int
	fmt.Scan(&N)

	var X0, Y0 int
	var minX, minY, maxX, maxY, midX, midY float64
	fmt.Scan(&X0, &Y0)
	for {
		// bombDir: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
		var bombDir string
		fmt.Scan(&bombDir)
		for _, dir := range bombDir {
			if string(dir) == "D" {
				midY = math.Floor(((minY + maxY) / 2))
				minY = midY + 1
			} else if string(dir) == "U" {
				midY = math.Floor(((minY + maxY) / 2))
				maxY = midY - 1
			} else if string(dir) == "R" {
				midX = math.Floor(((minX + maxX) / 2))
				minX = midX + 1
			} else if string(dir) == "L" {
				midX = math.Floor(((minX + maxX) / 2))
				maxX = midX - 1
			}
		}
		X0 = int(midX)
		Y0 = int(midY)
		fmt.Fprintln(os.Stderr, bombDir, X0, Y0)

		// the location of the next window Batman should jump to.
		fmt.Printf("%d %d\n", X0, Y0)
	}
}
