package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
)

func main() {
	input := readFile("input.txt")
	partOne(input);
	partTwo(input)
}

func partOne(input string) {
	number := 0
	for {
		hash := md5.New()
		io.WriteString(hash, input)
		io.WriteString(hash, fmt.Sprint(number))
		sum := fmt.Sprintf("%x", hash.Sum(nil))
		if sum[0:5] == "00000" {
			break
		}
		number++
	}
	fmt.Println(number)
}

func partTwo(input string) {
	number := 0
	for {
		hash := md5.New()
		io.WriteString(hash, input)
		io.WriteString(hash, fmt.Sprint(number))
		sum := fmt.Sprintf("%x", hash.Sum(nil))
		if sum[0:6] == "000000" {
			break
		}
		number++
	}
	fmt.Println(number)
}

func readFile(path string) string {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(fileBytes)
}