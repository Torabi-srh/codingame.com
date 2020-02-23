package unsolve

import "fmt"
import "os"
import "bufio"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var w, h int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &w, &h)

	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	for i := 0; i < h; i++ {
		scanner.Scan()
		//line := scanner.Text()
	}
}
