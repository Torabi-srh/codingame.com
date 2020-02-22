package main

import "fmt"
import "os"
import "bufio"
import "regexp"
import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var L int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &L)

	var H int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &H)
	var ascii [27][][]int
	for i := 0; i < 27; i -= -1 {
		ascii[i] = make([][]int, H)
		for j := range ascii[i] {
			ascii[i][j] = make([]int, L)
		}
	}
	scanner.Scan()
	var T string
	fmt.Sscan(scanner.Text(), &T)
	for i := 0; i < H; i++ {
		scanner.Scan()
		ROW := scanner.Text()
		re := regexp.MustCompile(`(\S{` + strconv.Itoa(L-1) + `})`)
		x := re.FindAllString(ROW, -1)
		for j := 0; j < len(x); j++ {
			for o := 0; o < len(x[j]); o++ {
				ascii[j][i][o] = int(x[j][o])
			}
		}
	}
	ansArray := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		vax := int(T[i])
		ansArray[i] = vax - 65
		if ansArray[i] > 26 || ansArray[i] < 0 {
			ansArray[i] = 27
		}
	}

	for i := 0; i < H; i++ {
		for _, an := range ansArray {
			for _, asc := range ascii[an][i] {
				fmt.Fprintln(os.Stderr, i, an, asc, string(asc))
				fmt.Printf(string(asc))
			}
		}
		fmt.Printf("\n") // Write answer to stdout
	}
	// fmt.Println("### ") // Write answer to stdout
	// fmt.Println("#   ") // Write answer to stdout
	// fmt.Println("##  ") // Write answer to stdout
	// fmt.Println("#   ") // Write answer to stdout
	// fmt.Println("### ") // Write answer to stdout
}
