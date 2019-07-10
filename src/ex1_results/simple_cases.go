package ex1_results

import (
	"errors"
	"regexp"
	"strconv"
)

//HelloWorld Traditionnal first method
func HelloWorld() string {
	return "Hello World"
}

func AddInts(value1, value2 int) int {
	return value1 + value2
}

func AddIntsDifferents(value1 int64, value2 int32) int {
	return int(value1) + int(value2)
}

//Substring standard method, don't use library
func Substring(value string, left, right int) string {
	if left > len(value) || right < left || right > len(value) {
		return ""
	}
	return value[left:right]
}

// SubstringWithErrors return substringed value but can generate errors
func SubstringWithErrors(value string, left, right int) (string, error) {
	switch {
	case left > len(value):
		return "", errors.New("Left bound must me lower than value length")
	case left > right:
		return "", errors.New("Left bound must me lower than right bound")
	case right > len(value):
		return "", errors.New("Right bound must me lower than value length")
	default:
		return value[left:right], nil
	}
}

func ExtractNumbersFromString(value string) []int {

	reg, err := regexp.Compile("([0-9]+)")

	if err != nil {
		return []int{}
	}

	founded := reg.FindAllStringSubmatch(value, -1)
	results := make([]int, 0, len(founded))
	for _, found := range founded {
		if value, err := strconv.ParseInt(found[1], 10, 32); err == nil {
			results = append(results, int(value))
		}
	}
	return results
}

// CountTypes count number of numbers, string and others
func CountTypes(values ...interface{}) (nbNumber, nbString, nbUnknown int) {
	nbNumber = 0
	nbString = 0
	nbUnknown = 0

	for _, value := range values {
		switch value.(type) {
		case string:
			nbString++
		case int32:
			nbNumber++
		case int64:
			nbNumber++
		case int:
			nbNumber++
		case float32:
			nbNumber++
		case float64:
			nbNumber++
		default:
			nbUnknown++
		}
	}
	return
}

//CreateSet create a set from a list of values
func CreateSet(values ...interface{}) map[interface{}]struct{} {
	set := make(map[interface{}]struct{}, len(values))
	for _, value := range values {
		set[value] = struct{}{}
	}
	return set
}

//GetEndList return the end of the list
func GetEndList(list []string, from int) ([]string, error) {
	if from > len(list) {
		return []string{}, errors.New("From must be lower than list length")
	}

	if from < 0 {
		return []string{}, errors.New("From must be greater or equal to zero")
	}
	return list[from:], nil
}
