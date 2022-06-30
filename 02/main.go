package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"io"
	"io/ioutil"
)

const target = 2020

type policy struct {
	low, high int
	value     string
}

type data struct {
	policy
	password string
}

func main() {
	dataBytes, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	xe := readData(bytes.NewReader(dataBytes))

	invalidData := make([]data, 0)
	validData := make([]data, 0)

	for _, data := range xe {
		_, ok := valid(data)
		if ok {
			validData = append(validData, data)
			continue
		}

		invalidData = append(invalidData, data)

	}
	fmt.Printf("\n\n\n %v invalid passwords \n\n\n", len(invalidData))
	fmt.Printf("\n\n\n %v valid passwords \n\n\n", len(validData))
	//fmt.Println(invalidData)

}

func valid(input data) (string, bool) {
	instances := strings.Count(input.password, input.value)

	if (instances < input.low) || (instances > input.high) {
		return input.password, false
	}

	return "", true
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

func readData(reader io.Reader) []data {
	scanner := bufio.NewScanner(reader)
	result := make([]data, 0)

	for scanner.Scan() {
		line := scanner.Text()
		xs := strings.SplitAfter(line, " ")
		e := processData(xs)

		result = append(result, e)
	}
	return result
}

func processData(input []string) data {
	xs := make([]string, 0)
	for _, s := range input {
		trimmed := strings.TrimSpace(s)
		xs = append(xs, trimmed)
	}
	lowAndHigh := strings.Split(xs[0], "-")
	l, err := strconv.Atoi(lowAndHigh[0])
	if err != nil {
		panic(err)
	}
	h, err := strconv.Atoi(lowAndHigh[1])
	if err != nil {
		panic(err)
	}

	value := strings.Trim(xs[1], ":")
	return data{
		policy: policy{
			low:   l,
			high:  h,
			value: value,
		},
		password: xs[2],
	}
}
