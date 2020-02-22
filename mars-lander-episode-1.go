package solved

import "fmt"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	// surfaceN: the number of points used to draw the surface of Mars.
	var surfaceN int
	fmt.Scan(&surfaceN)
	var surface [][2]int = make([][2]int, surfaceN)
	var flatS, flatE int
	for i := 0; i < surfaceN; i++ {
		// landX: X coordinate of a surface point. (0 to 6999)
		// landY: Y coordinate of a surface point. By linking all the points together in a sequential fashion, you form the surface of Mars.
		var landX, landY int
		fmt.Scan(&landX, &landY)
		if i > 0 {
			if surface[i-1][1] == landY {
				flatE = landX
				flatS = surface[i-1][1]
			}
		}
		surface[i][0] = landX
		surface[i][1] = landY
	}
	_ = flatE
	_ = flatS
	var engine, tilt, lastVS int
	var X, Y, hSpeed, vSpeed, fuel, rotate, power int
	vSpeed = 0
	for {
		// hSpeed: the horizontal speed (in m/s), can be negative.
		// vSpeed: the vertical speed (in m/s), can be negative.
		// fuel: the quantity of remaining fuel in liters.
		// rotate: the rotation angle in degrees (-90 to 90).
		// power: the thrust power (0 to 4).
		lastVS = vSpeed
		fmt.Scan(&X, &Y, &hSpeed, &vSpeed, &fuel, &rotate, &power)

		if vSpeed >= 40 {
			if lastVS < vSpeed {
				engine--
			}
		} else if vSpeed <= -40 {
			if lastVS > vSpeed {
				engine++
			}
		} else {
			engine--
		}
		if engine >= 4 {
			engine = 4
		}
		if engine <= 0 {
			engine = 0
		}
		fmt.Printf("%d %d\n", tilt, engine)
	}
}
