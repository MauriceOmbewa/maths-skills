# Math-skills

This program calculates the statistical measures of a dataset provided in a text file (numbers.txt). The supported statistical measures are:

    - Sum
    - Average
    - Median
    - Variance
    - Standard Deviation

## Functions

The program consists of several functions to perform specific tasks in the statistical analysis process. Here's a brief explanation of each function:
### main()

The main function is the entry point of the program. It reads the input file, calculates the statistical measures and outputs the results.

### Average(sum float64, length int) float64

The Average function calculates the average of a dataset by dividing the sum of the dataset by the number of elements.

### Median(data []float64) float64

The Median function calculates the median of a dataset by first sorting the dataset and then finding the middle value. If the dataset has an even number of elements, the median is calculated as the average of the two middle values.

### Variance(data []float64, mean float64) float64

The Variance function calculates the variance of a dataset by finding the average of the squared differences between each element and the mean.

### StdDev(variance float64) float64

The StdDev function calculates the standard deviation of a dataset by taking the square root of the variance.

These functions are called in the main function to perform the required statistical analysis.

## Getting Started
### Dependencies

The program requires the Go programming language (version 1.16 or higher) to be installed on your system.
### Installing

    Clone the repository or download the source code.
    Open a terminal and navigate to the project directory.
    Run go build to compile the source code.

### Executing program

    Prepare a text file containing the dataset, with one number per line.
    Run the compiled program with the text file as an argument:

    ``` go run main.go "numbers.txt" ```

    The program will output the calculated statistical measures.

### Built With

    Go - Programming language

### Author

- Maurice Ombewa @mombewa

### License

This project is licensed under the MIT License - see the LICENSE.md file for details

