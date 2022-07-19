package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

func isValidSectionName(result string) bool {
	regularExpression := regexp.MustCompile("[[a-z|A-Z]+]")
	if regularExpression.MatchString(result) {
		return true

	}

	return false

}

type Parser struct {
	Data map[string]map[string]string
	//Comments [10]string
}

func (p *Parser) SetValues(section, key, value string) {
	if _, ok := p.Data[section]; !ok {
		p.SetSections(section)

	}

	p.Data[section][key] = value
}

func (p *Parser) GetSectionNames() []string {

	keys := make([]string, 0, len(p.Data))
	for k := range p.Data {
		keys = append(keys, k)

	}

	fmt.Println("sections are :", keys)
	return keys
}

func (p *Parser) GetValue(section string, key string) (string, error) {
	if _, ok := p.Data[section]; !ok {
		return "", errors.New("No section with this name")
	}
	if _, ok := p.Data[section][key]; !ok {
		return "", errors.New("No key with this name")
	}

	return p.Data[section][key], nil
}

func (p *Parser) SetSections(section string) {
	_, ok := p.Data[section]
	if !ok {
		p.Data[section] = make(map[string]string)
	}
	fmt.Println("setsections")
}
func (p *Parser) SaveToFile(name string, dictionary map[string]map[string]string) (err error) {

	file, ferr := os.Create(name)
	if ferr != nil {
		return errors.New("can't open file with this name")
	}
	for k := range dictionary {
		_, err := file.WriteString(k + "\n")
		for key, value := range dictionary[k] {
			file.WriteString(key + " = " + value + "\n")

		}
		file.WriteString("\n")
		if err != nil {

			return errors.New("can't open file with this name")

		}

	}
	return nil
}
func (p *Parser) LoadFromFile(name string) error {
	f, ferr := os.Open(name)
	fmt.Println("openedfile")
	if ferr != nil {
		return errors.New("can't open the file with this name")
	}
	//code to read file
	content,err  :=io.ReadAll(f)
	if err ....

	return p.LoadFromString(string(content))
}

func (p *Parser) LoadFromString(content string) error {

	err = p.Parse(content)
	// fmt.Println(p.Data)
	// p.SaveToFile("k.txt", p.Data)

}

func (p *Parser) Parse(content string) error {
	scanner := bufio.NewScanner(strings.NewReader(content))

	var key string
	var value string
	var section string
	SectionFlag := false
	d := make(map[string]map[string]string)
	//d = p.Data

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		fmt.Println((items[0]))
		if items[0] == ";" {

			continue

		} else if isValidSectionName(items[0]) {
			p.SetSections(items[0])
			section = items[0]
			SectionFlag = true

		} else if SectionFlag == true {
			split_equal := strings.Split(line, "=")

			if len(split_equal) == 2 {

				key = split_equal[0]
				value = split_equal[1]

				p.SetValues(section, key, value)

			}

		} else if len(items) == 1 {
			if items[0] == " " {

			}

		}

	}
	dictionary = d
	p.SaveToFile("name.txt", p.Data)
	return dictionary, err
}

/*func check_name(name string) bool {
	regularExpression := regexp.MustCompile("[a-z|A-Z]+")
	if regularExpression.MatchString(name) {
		return true

	}
	return false
}

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
}*/

func main() {
	Data := make(map[string]map[string]string)
	var info Parser
	var name = "text.INI"
	info.LoadFromFile(name)
	fmt.Println("printed from main", Data)
	info.SaveToFile("name.txt", info.Data)
	/*d := make(map[string]map[string]string)
	d["database"] = make(map[string]string)
	d["database"]["username"] = "abc"
	d["database"]["password"] = "dmdm"
	d["owner"] = make(map[string]string)*/

}
