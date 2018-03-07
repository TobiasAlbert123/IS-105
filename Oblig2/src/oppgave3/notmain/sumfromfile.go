package notmain

import (
	"fmt"
	"strconv"
	"strings"
	"io/ioutil"
	"os"
)

func addup(tall, tall2 int) int{
	return tall + tall2
}

//var testfile = "Oblig2/src/oppgave3/testfile.txt"

/*func main() {
	filen, err := ioutil.ReadFile(testfile)
	file, err := os.Open(testfile)
	if err != nil {

	}
	fileScanner := bufio.NewScanner(file)
	fileString := string(filen)
	splitString := []string(strings.Split(fileString, "\n"))

	for fileScanner.Scan() {
		fmt.Println(filen)
	}

}*/

func WriteSum(number int) {
	file, err := os.Create("../testfile.txt")
	if err != nil {
		fmt.Println("Error creating file")
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%d", number))
	if err != nil {
		fmt.Println("Error writing")
	} else {
		fmt.Printf("Wrote sum %d to testfile.txt\n", number)
	}
}

func SumFromFile() int{
	fmt.Println("Oppgave 3")
	file, err := ioutil.ReadFile("../testfile.txt")
	if err != nil {

	}
	fileString := string(file)
	splitString := []string(strings.Split(fileString, "\n"))
	var ints [2]int
	for i := 0; i < len(splitString); i++ {
		ints[i], err = strconv.Atoi(splitString[i])
		if err != nil {
			fmt.Println("int conversion error")
		}
	}
	total := addup(ints[0], ints[1])
	fmt.Printf("Total is %d\n", total)
	return total
	//ints, err := ReadInts(strings.NewReader(file))
	//fmt.Println(ints, err)
	//was looking at https://stackoverflow.com/questions/9862443/golang-is-there-a-better-way-read-a-file-of-integers-into-an-array
}
