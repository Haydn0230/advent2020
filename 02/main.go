package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"

	"io"
	"io/ioutil"
)

const target = 2020

func main () {
	dataBytes, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}

	result := readInts(bytes.NewReader(dataBytes))

	var answer int
	for _, firstValue := range result {
		for _, secondValue := range result {
			sum := firstValue + secondValue
			if sum == target {
				answer = firstValue * secondValue
			}
		}
	}

	fmt.Println(answer)

}

func readInts(reader io.Reader) []int {
	scanner := bufio.NewScanner(reader)
	result := make([]int, 0)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		result = append(result, i)
	}
	return result
}