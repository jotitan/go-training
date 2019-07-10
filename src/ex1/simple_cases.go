package ex1

// HelloWorld classic function
// return a string that says 'Hello World' (without the quotes)
func HelloWorld() string {
	return ""
}

// AddInts
// return the sum of two integers (value1 and value2)
func AddInts(value1, value2 int) int {
	return 0
}

// AddIntsDifferents
// return sum of two different types of integers: value1(64bits) and value2(32bits)
func AddIntsDifferents(value1 int64, value2 int32) int {
	return 0
}

// Substring standard method
// return a substring of value starting on the index 'left' until 'right'
// don't use any libraries
func Substring(value string, left, right int) string {
	return ""
}

// SubstringWithErrors
// same as above, but should throw errors when 'left' and 'right' are not appropriate
// Q: "But what do you consider 'appropriate'?"
// A: I don't know, what do _you_ consider appropriate?
func SubstringWithErrors(value string, left, right int) (string, error) {
	return "", nil
}

// ExtractNumbersFromString
// return a slice of integers containing all the numbers found in
// the string 'value'
func ExtractNumbersFromString(value string) []int {
	return []int{}
}

// CountTypes
// count the amount of numbers (nbNumber), strings (nbString) and other types (nbUnknown)
// passed as arguments (hint: read about empty interfaces)
func CountTypes(values ...interface{}) (nbNumber, nbString, nbUnknown int) {
	return 0, 0, 0
}

// CreateSet
// create a set from a list of values
// the set must not contain repeated values
// the 'values' (the variables passed as arguments) are the keys of the map
// the values the map holds are empty structs (hint: read about empty structs)
func CreateSet(values ...interface{}) map[interface{}]struct{} {
	return make(map[interface{}]struct{})
}

// GetEndList
// return a string slice containing the elements of 'list' beginning at index 'from'
func GetEndList(list []string, from int) ([]string, error) {
	return []string{}, nil
}
