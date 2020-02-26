package main

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

	scanner.Scan()
	magicPhrase := scanner.Text()
	var oph string = ""
	fmt.Fprintln(os.Stderr, magicPhrase)
	for _, mp := range magicPhrase {
		var mpi int = int(mp) - 64
		var fst bool = false
		spi := 0
		if mpi > 13 {
			fst = true
			spi = mpi
			mpi = 27
		}
		fmt.Fprintln(os.Stderr, spi, mpi, string(mp), fst)
		for i := spi; i < mpi; i++ {
			if fst {
				oph += "-"
			} else {
				oph += "+"
			}
		}
		oph += ".>"
	}
	fmt.Println(oph) // Write action to stdout
}
