package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func main() {
	// input matrix
	matrix := matrix()

	// check if it's square and within 10x10
	ok := validate(matrix)
	if !ok {
		log.Fatal("Error. Please ensure your matrix is square and 10x10 or smaller.")
	}

	// transpose the matrix
	transposed := transposeMatrix(matrix)

	// multiply
	product := multiplyMatrices(matrix, transposed)
	log.Println("Product:", product)
}

// TODO: accept various size 2d array (matrix). Besides 3x3
func matrix() [3][3]int {
	return parseFile(inputFile())
}

func inputFile() (f *os.File) {
	filename := "input.csv"
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error: couldnt open file.")
	}
	return
}

func parseFile(file *os.File) (matrix [3][3]int) {
	defer file.Close()
	matrix = [...][3]int{[3]int{}, [3]int{}, [3]int{}}

	reader := csv.NewReader(file)
	csvLines, err := reader.ReadAll()
	if err != nil {
		log.Println("Error reading csv lines: ", err)
	}

	for i := range csvLines {
		yCoord, err := strconv.Atoi(csvLines[i][0])
		xCoord, err := strconv.Atoi(csvLines[i][1])
		value, err := strconv.Atoi(csvLines[i][2])
		if err != nil {
			log.Println("Error converting string to int from csv: ", err)
		}
		matrix[yCoord-1][xCoord-1] = value
	}
	return
}

func validate(matrix [3][3]int) (err bool) {
	height := len(matrix)
	firstWidth := len(matrix[0])
	if height == firstWidth && height <= 10 && firstWidth <= 10 {
		err = true
	} else {
		err = false
	}
	return
}

func transposeMatrix(original [3][3]int) (transposed [3][3]int) {
	transposed = [...][3]int{[3]int{}, [3]int{}, [3]int{}}
	height := len(original)
	firstWidth := len(original[0])
	width := firstWidth
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			val := original[h][w]
			transposed[w][h] = val
		}
	}
	return
}

func multiplyMatrices(original [3][3]int, transposed [3][3]int) (product [3][3]int) {
	product = [...][3]int{[3]int{}, [3]int{}, [3]int{}}
	height := len(original)
	firstWidth := len(original[0])
	width := firstWidth
	for h := 0; h < height; h++ {
		for w := 0; w < width; w++ {
			product[h][w] = 0
			for i := 0; i < width; i++ {
				product[h][w] += original[h][i] * transposed[i][w]
			}
		}
	}
	return
}
