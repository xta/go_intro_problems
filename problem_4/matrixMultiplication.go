package main

import (
	"log"
)

func main() {
	// input matrix
	matrix := matrix()

	// check if it's square and within 10x10
	err := validate(matrix)
	if err {
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
	return [3][3]int{[3]int{1, 2, 3}, [3]int{2, 2, 2}, [3]int{-1, 0, 1}}
}

func validate(matrix [3][3]int) (err bool) {
	height := len(matrix)
	firstWidth := len(matrix[0])
	if height == firstWidth && height <= 10 && firstWidth <= 10 {
		err = false
	} else {
		err = true
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
