/*package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Oppgave 3")
	file, err := os.Open("Oblig2/src/oppgave3/testfile.txt")
	if err != nil {

	}
	buf := make([]byte, 1024)
	fmt.Println(file.Read(buf))
}*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
	"os"
)

// ReadInts reads whitespace-separated ints from r. If there's an error, it
// returns the ints successfully read so far as well as the error value.
func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func main() {
	tf := "1\n2\n3\n4\n5\n6"
	fmt.Println("Oppgave 3")
	file, err := os.Open("Oblig2/src/oppgave3/testfile.txt")
	if err != nil {

	}
	ints, err := ReadInts(strings.NewReader(file))
	fmt.Println(ints, err)
	//was looking at https://stackoverflow.com/questions/9862443/golang-is-there-a-better-way-read-a-file-of-integers-into-an-array
}
