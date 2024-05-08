// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	 sum := 0
// 	// average := 0
// 	// median := 0
// 	// variance := 0
// 	// stnddev := 0

// 	inputfilename := os.Args[1]
// 	inputfile, err := os.Open(inputfilename)
// 	if err != nil {
// 		fmt.Println("Error opening file ", err)
// 		return
// 	}

// 	scanner := bufio.NewScanner(inputfile)
// 	//var new string
// 	//var line2 []int
// 	var splitted []string

// 	for scanner.Scan(){
// 		new := scanner.Text()
// 		// fmt.Println(len(new))
// 		splitted = strings.Split(new, "\n")
// 		// fmt.Println(splitted)

// 		for i := 0; i < len(splitted); i++{
// 			atc, err := strconv.Atoi(string(splitted[i]))
// 			if err != nil{
// 				fmt.Println(err)
// 			}
// 			sum += atc
// 		}
// 	}
// 	count1 := 0
// 	for i, char := range splitted{
// 		fmt.Println(char[i])
// 		count1++
// 	}
// 	// fmt.Println(count1)

// 	fmt.Println(sum)

// 	// Average(sum, )

// }

// func Average(sum int, length int) int{
// 	ave := sum / length
// 	return ave
// }

// ________________________________________________________________

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	 sum := 0
// 	// average := 0
// 	// median := 0
// 	// variance := 0
// 	// stnddev := 0

// 	inputfilename := os.Args[1]
// 	inputfile, err := os.Open(inputfilename)
// 	if err != nil {
// 		fmt.Println("Error opening file ", err)
// 		return
// 	}

// 	scanner := bufio.NewScanner(inputfile)
// 	//var new string
// 	//var line2 []int

// 	for scanner.Scan(){
// 		new := scanner.Text()
// 		splitted := strings.Split(new, "\n")

// 		sum = Sum(splitted)

// 	}
// 	fmt.Println(sum)

// }

// func Sum(splitted []string) int{
// 	sum := 0
// 	for i := 0; i < len(splitted); i++{
// 		atc, err := strconv.Atoi(string(splitted[i]))
// 		if err != nil{
// 			fmt.Println(err)
// 		}
// 		sum += atc
// 	}
// 	return sum
// }

// ________________________________________________________________

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0
	// average := 0
	// median := 0
	// variance := 0
	// stnddev := 0

	inputfilename := os.Args[1]
	inputfile, err := os.Open(inputfilename)
	if err != nil {
		fmt.Println("Error opening file ", err)
		return
	}

	scanner := bufio.NewScanner(inputfile)

	var data []int

	for scanner.Scan() {
		new := scanner.Text()

		atc, err := strconv.Atoi(new)
		if err != nil {
			fmt.Println(err)
		}
		data = append(data, atc)

	}
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	length := len(data)
	fmt.Println(sum)

	// fmt.Println(len(data))
	// fmt.Println(data[3])

	average := Average(sum, length)
	median := Median(data)
	fmt.Println(average)
	fmt.Println(median)
}

func Average(sum int, length int) int {
	ave := sum / length
	return ave
}

func Median(data []int) int {
	for i := 0; i < len(data); i++ {
        for j := i + 1; j < len(data); j++ {
            if data[i] > data[j] {
                data[i], data[j] = data[j], data[i]
            }
        }
    }
	var median int
	mid := len(data) / 2

	if len(data)%2 == 0 {
		median = (data[mid] + data[mid-1]) / 2
	} else {
		median = data[mid]
	}
	return median
}
