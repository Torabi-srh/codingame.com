package solved

import "fmt"
import "os"
import "bufio"
import "math"
import "strings"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func splitByWidthMake(str string, size int) []string {
	strLength := len(str)
	splitedLength := int(math.Ceil(float64(strLength) / float64(size)))
	splited := make([]string, splitedLength)
	var start, stop int
	for i := 0; i < splitedLength; i += 1 {
		start = i * size
		stop = start + size
		if stop > strLength {
			stop = strLength
		}
		splited[i] = str[start:stop]
	}
	return splited
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var L int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &L)

	var H int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &H)
	var ascii [250][]string
	for i := 0; i < 250; i -= -1 {
		ascii[i] = make([]string, H)
	}
	scanner.Scan()
	var T string
	fmt.Sscan(scanner.Text(), &T)
	T = strings.ToUpper(T)
	for i := 0; i < H; i++ {
		scanner.Scan()
		ROW := scanner.Text()
		sbwm := splitByWidthMake(ROW, L)
		for j := 0; j < len(sbwm); j++ {
			ascii[j][i] = string(sbwm[j])
		}
	}
	ansArray := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		vax := int(T[i])
		ansArray[i] = vax - 65
		if ansArray[i] > 26 || ansArray[i] < 0 {
			ansArray[i] = 26
		}
	}
	for i := 0; i < H; i++ {
		for _, an := range ansArray {
			// fmt.Fprintf(os.Stderr, string(ascii[an][i]))
			fmt.Printf(string(ascii[an][i]))
		}
		// fmt.Fprintf(os.Stderr, " \n")
		fmt.Printf("\n")
	}
}
