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
	regularExpression := regexp.MustCompile("[[a-z|A-Z]+]")
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
	Data map[string]map[string]string

	Comments [10]string
}

func Search(dictionary map[string]map[string]string, section string, key string) string {
	return dictionary[section][key]
}

func LoadFromFile(name string, info Inforamation) {
	file, ferr := os.Open(name)
	if ferr != nil {
		panic(ferr)
	}

	scanner := bufio.NewScanner(file)
	LoadFromString(scanner, info)
}

func LoadFromString(scanner *bufio.Scanner, info Inforamation) {

	var section string
	cont := make(map[string]string)

	m := make(map[string]map[string]string)
	length := 0
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		//fmt.Println("number of items:", len(items), items[0])

		if items[0] == ";" {

			for i := 0; i < len(items); i++ {
				info.Comments[length] = line
			}
			length++

		} else if Checking(items[0]) {
			section = items[0]

		} else if len(items) == 1 && items[0] == " " {
			section = " "

		} else {

			split_equal := strings.Split(line, "=")

			if len(split_equal) == 2 {
				cont[split_equal[0]] = split_equal[1]
				m[section] = cont
				//	info.Data = m

				/*fmt.Println(cont)
				fmt.Println("before_delete", m)

				delete(cont, split_equal[0])
				fmt.Println("after delete", cont)
				fmt.Println("sub_map", m)*/
			}
			info.Data = m
			for k := range cont {
				delete(cont, k)
			}

		}

	}
	//info.Data = m
	fmt.Println(info.Data)

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
	var info Inforamation
	var name = "text.INI"
	LoadFromFile(name, info)

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
	//fmt.Println("checking", Checking("[hhk]"))

}

/*checking on section should begin with [ and end with ']'
section name has only letters
section must have key
the key must have value
*/
