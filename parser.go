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
		return true

	}

	return false

}

type Inforamation struct {
	Data     map[string]map[string]string
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
	var key string
	var value string
	d := make(map[string]map[string]string)

	var section string

	length := 0
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")

		if items[0] == ";" {

			for i := 0; i < len(items); i++ {
				info.Comments[length] = line
			}
			length++

		} else if Checking(items[0]) {
			section = items[0]
			d[section] = make(map[string]string)

		} else if len(items) == 1 && items[0] == " " {
			section = " "

		} else {

			split_equal := strings.Split(line, "=")

			if len(split_equal) == 2 {

				key = split_equal[0]
				value = split_equal[1]

				d[section][key] = value

			}

		}

	}
	info.Data = d
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

	var info Inforamation
	var name = "text.INI"
	LoadFromFile(name, info)
	//hey
}
