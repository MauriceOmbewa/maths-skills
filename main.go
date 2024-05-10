package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	maths"maths_skills/maths"
)

func main() {
	var sum float64

	if len(os.Args) != 2 {
		fmt.Println("Error: Incorrect number of arguments")
		return
	}

	inputfilename := os.Args[1]
	inputfile, err := os.Open(inputfilename)
	if err != nil {
		fmt.Println("Error opening file ", err)
		return
	}
	defer inputfile.Close()

	fileinfo, err := inputfile.Stat()
	if err != nil{
		fmt.Println("Error: ",err)
		return
	}
	if fileinfo.Size() == 0{
		fmt.Println("Error: File is empty")
		return
	}

	scanner := bufio.NewScanner(inputfile)


	var data []float64

	for scanner.Scan() {
		new := scanner.Text()

		atc, err := strconv.ParseFloat(new, 64)
		if err != nil {
			fmt.Println("Error: ",err)
			return
		}
		data = append(data, float64(atc))

	}
	if err := scanner.Err() ; err != nil{
		fmt.Println("Error: ",err)
		return
	} 
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	length := float64(len(data))
	fmt.Println(sum)


	average := maths.Average(float64(sum), length)
	median := maths.Median(data)
	variance := maths.Variance(data, float64(average))
	stddev := maths.StdDev(variance)
	fmt.Println("Average:",int(average))
	fmt.Println("Median:",int(median))
	fmt.Println("Variance:",int(variance))
	fmt.Println("Standard Deviation:",int(stddev))
}

