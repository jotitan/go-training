package ex1

func HelloWorld() string {
	return ""
}

func AddInts(value1, value2 int) int {
	return 0
}

func AddIntsDifferents(value1 int64, value2 int32) int {
	return 0
}

//Substring standard method, don't use library
func Substring(value string, left, right int) string {
	return ""
}

// SubstringWithErrors return substringed value but can generate errors
func SubstringWithErrors(value string, left, right int) (string, error) {
	return "", nil
}

func ExtractNumbersFromString(value string) []int {
	return []int{}
}

// CountTypes count number of numbers, string and others
func CountTypes(values ...interface{}) (nbNumber, nbString, nbUnknown int) {
	return 0, 0, 0
}

//CreateSet create a set from a list of values
func CreateSet(values ...interface{}) map[interface{}]struct{} {
	return make(map[interface{}]struct{})
}

//GetEndList return the end of the list
func GetEndList(list []string, from int) ([]string, error) {
	return []string{}, nil
}
