package solved

import "fmt"
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
	var PX0, PY0, PX1, PY1 int
	var minX, minY, maxX, maxY, midX, midY float64 = 0, 0, 0, 0, 0, 0
	fmt.Scan(&X0, &Y0)
	var PD1, PD2 string
	PD1 = ""
	PD2 = ""
	X0++
	Y0++
	_, _ = PX1, PY1
	v := true
	for {
		// bombDir: the direction of the bombs from batman's current location (U, UR, R, DR, D, DL, L or UL)
		var bombDir string
		fmt.Scan(&bombDir)
		if v {
			v = false
			if string(bombDir) == "D" {
				minY = float64(Y0)
				maxY = float64(H)
			} else if string(bombDir) == "U" {
				minY = 0.0
				maxY = float64(Y0)
			} else if string(bombDir) == "R" {
				minX = float64(X0)
				maxX = float64(W)
			} else if string(bombDir) == "L" {
				minX = 0.0
				maxX = float64(X0)
			} else if string(bombDir) == "DR" {
				minY = float64(Y0)
				maxY = float64(H)
				minX = float64(X0)
				maxX = float64(W)
			} else if string(bombDir) == "DL" {
				minY = float64(Y0)
				maxY = float64(H)
				minX = 0.0
				maxX = float64(X0)
			} else if string(bombDir) == "UR" {
				minY = 0.0
				maxY = float64(Y0)
				minX = float64(X0)
				maxX = float64(W)
			} else if string(bombDir) == "UL" {
				minY = 0.0
				maxY = float64(Y0)
				minX = 0.0
				maxX = float64(X0)
			}
		}
		for _, dir := range bombDir {
			if string(dir) == "D" {
				if PD1 == "U" {
					minY = float64(PY0)
					maxY = float64(PY1)
				}
				PD1 = string(dir)
				midY = math.Floor(((minY + maxY) / 2))
				minY = midY + 1
			} else if string(dir) == "U" {
				if PD1 == "D" {
					minY = float64(PY1)
					maxY = float64(PY0)
				}
				PD1 = string(dir)
				midY = math.Floor(((minY + maxY) / 2))
				maxY = midY - 1
			} else if string(dir) == "R" {
				if PD2 == "L" {
					minX = float64(PX0)
					maxX = float64(PX1)
				}
				PD2 = string(dir)
				midX = math.Floor(((minX + maxX) / 2))
				minX = midX + 1
			} else if string(dir) == "L" {
				if PD2 == "R" {
					minX = float64(PX1)
					maxX = float64(PX0)
				}
				PD2 = string(dir)
				midX = math.Floor(((minX + maxX) / 2))
				maxX = midX - 1
			}
		}

		X0 = int(midX)
		Y0 = int(midY)
		PY1 = PY0
		PX1 = PX0
		PX0 = X0
		PY0 = Y0

		// the location of the next window Batman should jump to.
		fmt.Printf("%d %d\n", X0, Y0)
	}
}
