package main

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strings"
)

const (
	Section  string = "section"
	Empty    string = "empty"
	Keyvalue string = "keyvalue"
	Comment  string = "comment"
)

var (
	ErrSectionNotFound = errors.New("section not found")
	ErrKeyNotFound     = errors.New("key not found")
	ErrInvalid         = errors.New("Invalid input")
)

// check if the line is section have (only one '[' and one ']' and any sequence of letters) or not
// return true if it is section or false if it isn't
func isValidSectionName(result string) bool {
	regularExpression := regexp.MustCompile("[[a-z|A-Z]+]")
	if regularExpression.MatchString(result) {
		return true

	}

	return false

}

// parser is struct have map to save the data of the ini file
// has all the functions of the parser
type Parser struct {
	data map[string]map[string]string
}

// set the key and value into map linked with the section
func (p *Parser) Set(section, key, value string) {
	if _, ok := p.data[section]; !ok {
		p.SetSections(section)

	}

	p.data[section][key] = value
}

// create new key [section] and linked it with new map[string]string
func (p *Parser) SetSections(section string) {
	if len(p.data) == 0 {
		p.data = map[string]map[string]string{}
	}
	_, ok := p.data[section]
	if !ok {

		p.data[section] = make(map[string]string)
	}

}

// get the map and return it

func (p *Parser) GetSections() map[string]map[string]string {

	return p.data
}

// get all the sectionsNames [keys]of the map
// return the kays
func (p *Parser) GetSectionNames() []string {

	keys := make([]string, 0, len(p.data))
	for k := range p.data {
		keys = append(keys, k)

	}

	return keys
}

//get the value of existing section and key , otherwise return error
func (p *Parser) Get(section string, key string) (string, error) {
	if _, ok := p.data[section]; !ok {
		return "", ErrSectionNotFound
	}
	if _, ok := p.data[section][key]; !ok {
		return "", ErrKeyNotFound
	}

	return p.data[section][key], nil
}

// search if section is exist or not
func (p *Parser) SearchSection(section string) (bool, error) {
	_, ok := p.data[section]
	if !ok {

		return false, ErrSectionNotFound
	}
	return true, nil
}

// save all the valid data which stored in map
// return error if the file can't be opened
func (p *Parser) SaveToFile(name string) (err error) {

	file, err := os.Create(name)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err = file.WriteString(p.String())

	if err != nil {

		return err

	}

	return nil
}

// try to open INIFile and get it's content
// return error if the file can't be opened
func (p *Parser) LoadFromFile(name string) error {
	f, err := os.ReadFile(name)
	if err != nil {
		return err
	}

	content := string(f)

	return p.LoadFromString((content))
}

//Inilize the map then call the parser and save to the file at the end
// return error from the parser
func (p *Parser) LoadFromString(content string) (err error) {
	if len(p.data) == 0 {
		p.data = map[string]map[string]string{}
	}
	p.data, err = parse(content)
	//p.SaveToFile("SavedFile.ini")

	return err

}

// check each line to determine if it is section , comment , KeyValue , empty or InvalidSyntax
// return the type of the line and error if it is Invalid input
func checkLine(line string) (string, error) {
	splits := strings.Split(line, " ")
	equal := strings.Split(line, " = ")
	equal2 := strings.Split(line, "=")

	if isValidSectionName(line) && len(line) == 1 {
		return Section, nil

	} else if splits[0] == ";" {
		return Comment, nil

	} else if line == "\n" {
		return Empty, nil

	} else if len(equal) == 2 || len(equal2) == 2 {
		return Keyvalue, nil
	}
	return " ", ErrInvalid
}

// extract sections , keys and values from the content then save it in a map
// return map which has valid syntax of sections , keys and values
// return error if the content contains Ivalid syntax
func parse(content string) (map[string]map[string]string, error) {
	scanner := bufio.NewScanner(strings.NewReader(content))
	ini := make(map[string]map[string]string)
	var key string
	var value string
	var section string
	SectionFlag := false

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		if items[0] == ";" {

			continue

		} else if isValidSectionName(line) {
			section = items[0]
			_, ok := ini[section]
			if !ok {

				ini[section] = make(map[string]string)
			}
			SectionFlag = true

		} else if SectionFlag == true {
			split_equal := strings.Split(line, "=")
			split_equal2 := strings.Split(line, " = ")

			if len(split_equal) == 2 {

				key = split_equal[0]
				value = split_equal[1]

				ini[section][key] = value

			} else if len(split_equal2) == 2 {

				key = split_equal2[0]
				value = split_equal2[1]

				ini[section][key] = value
			} else if len(split_equal) > 2 || len(split_equal2) > 2 {
				return map[string]map[string]string{}, errors.New("more than one value")
			} else if len(line) != 0 {
				return map[string]map[string]string{}, ErrInvalid
			}

		} else if len(line) == 0 {
			continue

		} else {

			return map[string]map[string]string{}, ErrInvalid
		}

	}
	return ini, nil
}
func (p *Parser) String() string {

	initext := ""
	for k := range p.data {
		initext += (k + "\n")
		for key, value := range p.data[k] {
			initext += (key + " = " + value + "\n")

		}
		initext += ("\n")
	}
	return initext
}
