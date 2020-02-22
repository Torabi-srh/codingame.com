package solved

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "math"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var result int = 1000000
	// n: the number of temperatures to analyse
	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	scanner.Scan()
	inputs := strings.Split(scanner.Text(), " ")
	if n == 0 {
		result = n
	}
	for i := 0; i < n; i++ {
		// t: a temperature expressed as an integer ranging from -273 to 5526
		t, _ := strconv.ParseInt(inputs[i], 10, 32)
		if math.Abs(float64(result)) >= math.Abs(float64(t)) {
			if math.Abs(float64(result)) == math.Abs(float64(t)) {
				if int(result) < 0 && int(t) > 0 {
					result = int(t)
				}
			} else {
				result = int(t)
			}
		}
	}
	fmt.Println(result) // Write answer to stdout
}
