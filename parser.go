package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func Checking(result string) bool {
	regularExpression := regexp.MustCompile("'['[a-z|A-Z]+']")
	if regularExpression.MatchString(result) {
		//fmt.Println("result is :", result, " pass")
		return true

	}
	/*if !regularExpression.MatchString(result) {

		//fmt.Println("result is :", result, " failed")
		return false
	}*/
	return false

}

type Inforamation struct {
	Data     map[string]map[string]string
	Comments []string
}

func Search(dictionary map[string]map[string]string, section string, key string) string {
	return dictionary[section][key]
}

func LoadFromFile(name string) {
	file, ferr := os.Open(name)
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)
	LoadFromString(scanner)
}

func LoadFromString(scanner *bufio.Scanner) {

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")

		fmt.Println(items[0])
	}

}
func check_name(name string) bool {
	regularExpression := regexp.MustCompile("[a-z|A-Z]+")
	if regularExpression.MatchString(name) {
		return true

	}
	return false
}

//"1K345"
func check_port(port string) bool {
	for _, c := range port {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}
func check_FileName(name string) bool {

	regularExpression := regexp.MustCompile("[.].[txt|dat]")
	if regularExpression.MatchString(name) {
		return true

	}
	return false
}

func check_server(server string) bool {

	regularExpression := regexp.MustCompile("[0-9]+[.][0-9]+[.][0-9]+[.][0-9]+")
	if regularExpression.MatchString(server) {
		return true

	}
	return false
}

func check_org(org string) bool {

	regularExpression := regexp.MustCompile("[a-z|A-Z]+[.]")
	if regularExpression.MatchString(org) {
		return true

	}
	return false
}

func main() {

	/*type SectionOptions = map[string]string
	   INIFile = map[string]SectionOptions
	   INI ={
		"section1" : {
			"key1": "val1"
		},
	   }*/
	// var info []Information
	var name = "text.INI"
	LoadFromFile(name)

	/*Data := map[string]map[string]string{
		"section1": {
			"name": "doaa",
			"age":  "21",
		},
		"section2": {
			"key1":  "val1",
			"key2 ": "val2"},
	}
	//sections := [2]string{"user","database"}
	//dictionary1 := map[string]string{"test1": "this is just a test1", "key2": "value2"}
	//dictionary2 := map[string]string{"test2": "this is just a test2"}
	fmt.Println("return value 1 : ", Search(Data, "section1", "name"))
	fmt.Println("return value 2 : ", Search(Data, "section2", "key1"))

	Checking("jj")*/
	fmt.Println("checking", check_org("Acme Widgets Inc."))

}

/*checking on section should begin with [ and end with ']'
section name has only letters
section must have key
the key must have value
*/
