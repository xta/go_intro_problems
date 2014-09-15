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
	file := inputFile()
	defer file.Close()

	hash := hashFile(file)
	log.Println(hash)
}

func inputFile() (file *os.File) {
	if stdinPresent() {
		file = os.Stdin
	} else {
		file = openFile("sample.txt")
	}
	return
}

func stdinPresent() bool {
	if stats, _ := os.Stdin.Stat(); stats.Size() > 0 {
		return true
	} else {
		return false
	}
}

func openFile(filePath string) (f *os.File) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error: couldnt open file.")
	}
	return
}

func hashFile(file *os.File) string {
	hasher := md5.New()
	io.Copy(hasher, file)
	return hex.EncodeToString(hasher.Sum(nil))
}
