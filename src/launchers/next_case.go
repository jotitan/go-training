package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	suffixTestFile = "_test.go"
)

func main() {

	if len(os.Args) == 2 && strings.EqualFold(os.Args[1], "init") {
		initDescribes()
		return
	}

	if len(os.Args) == 2 && strings.EqualFold(os.Args[1], "reset") {
		resetDescribes()
		return
	}

	if len(os.Args) != 3 {
		panic("Impossible, must specify init | reset | status and folder with test")
	}
	if strings.Contains(os.Args[1], "FAIL") {
		fmt.Println("Test fail, no new test to unlock")
		return
	}

	updateDescribes()
}

func resetDescribes() {
	fmt.Println("Remove all XDescribe to Describe")

	// Parcourt tous les repertoires, cherche un fichier repertoire_test.go et modifie
	f, _ := os.Open(".")
	files, _ := f.Readdir(-1)
	for _, file := range files {
		if file.IsDir() {
			filename := filepath.Join(file.Name(), file.Name()+suffixTestFile)
			if dataFileTest, err := ioutil.ReadFile(filename); err == nil {
				regDescribe, _ := regexp.Compile("\t(XDescribe)")
				dataFileTest = regDescribe.ReplaceAll(dataFileTest, []byte("\tDescribe"))
				if f, err := os.OpenFile(filename, os.O_TRUNC|os.O_RDWR, os.ModePerm); err == nil {
					fmt.Println("Reset file", filename)
					f.Write(dataFileTest)
					f.Close()
				}
			}
		}
	}
}

func initDescribes() {
	fmt.Println("Init describes, transform Describe to XDescribe, exept the first")

	// Parcourt tous les repertoires, cherche un fichier repertoire_test.go et modifie
	f, _ := os.Open(".")
	files, _ := f.Readdir(-1)
	for _, file := range files {
		if file.IsDir() {
			filename := filepath.Join(file.Name(), file.Name()+suffixTestFile)
			if dataFileTest, err := ioutil.ReadFile(filename); err == nil {
				regDescribe, _ := regexp.Compile("\t(Describe)")
				// Modify all except first
				describes := regDescribe.FindAllSubmatchIndex(dataFileTest, -1)
				for i := len(describes) - 1; i >= 1; i-- {
					dataFileTest = append(dataFileTest[:describes[i][2]], append([]byte("XDescribe"), dataFileTest[describes[i][3]:]...)...)
				}
				if f, err := os.OpenFile(filename, os.O_TRUNC|os.O_RDWR, os.ModePerm); err == nil {
					fmt.Println("Replace file", filename)
					f.Write(dataFileTest)
					f.Close()
				}
			}
		}
	}
}

func updateDescribes() {
	filename := os.Args[2] + suffixTestFile

	data, _ := ioutil.ReadFile(filename)

	regXDescribe, _ := regexp.Compile("\t(XDescribe)")

	results := regXDescribe.FindAllSubmatchIndex(data, -1)
	if len(results) > 0 {
		data = append(data[:results[0][2]], append([]byte("Describe"), data[results[0][3]:]...)...)

		// Si results2 > 0, on prend le premier et on le passe en describe
		if f, err := os.OpenFile(filename, os.O_TRUNC|os.O_RDWR, os.ModePerm); err == nil {
			f.Write(data)
			f.Close()
		}
	} else {
		fmt.Println("\nAll tests are successful, congrats !!!\n")
	}
}
