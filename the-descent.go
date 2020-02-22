package solved

import (
	"fmt"
)

func main() {
	for {
		var mountainM, ix int = 0, 0
		for i := 0; i < 8; i++ {
			var mountainH int
			fmt.Scan(&mountainH)
			if mountainH > mountainM {
				mountainM = mountainH
				ix = i
			}
		}
		fmt.Println(ix)
	}
}
