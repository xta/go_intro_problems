/*

   1. write a go implementation of the md5sum program (in short, given a filename it outputs an md5 hash of its content);

   bonus points if you can feed it by piping into it (cat file.txt | ./md5sum) or feed it from standard input (./md5sum < file.txt);

   more bonus points if it can print an md5sum of itself if given no parameters at all :)

*/

package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func main() {
	filename := "sample.txt"

	log.Println("Opening file: " + filename)

	md5Hash := processFile(filename)

	log.Println("md5hash of", filename, ":", md5Hash)
}

func processFile(filePath string) (hash string) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error: couldnt open file.")
	}
	defer f.Close()

	md5Hash := hashFile(f)

	return md5Hash
}

func hashFile(file *os.File) (hash string) {
	hasher := md5.New()
	io.Copy(hasher, file)
	return hex.EncodeToString(hasher.Sum(nil))
}
