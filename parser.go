package main

import (
	"bufio"
	"errors"
	"fmt"
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
}

func (p *Parser) SetValues(section, key, value string) {
	if _, ok := p.Data[section]; !ok {
		p.SetSections(section)

	}

	p.Data[section][key] = value
}
func (p *Parser) GetSections() map[string]map[string]string {

	return p.Data
}

func (p *Parser) GetSectionNames() []string {

	keys := make([]string, 0, len(p.Data))
	for k := range p.Data {
		keys = append(keys, k)

	}

	return keys
}

func (p *Parser) Get(section string, key string) (string, error) {
	if _, ok := p.Data[section]; !ok {
		return "", errors.New("No section with this name")
	}
	if _, ok := p.Data[section][key]; !ok {
		return "", errors.New("No key with this name")
	}

	return p.Data[section][key], nil
}
func (p *Parser) SearchSection(section string) (err error) {
	_, ok := p.Data[section]
	if !ok {

		return errors.New("No section with this name")
	}
	return nil
}
func (p *Parser) SetSections(section string) {
	if len(p.Data) == 0 {
		p.Data = map[string]map[string]string{}
	}
	_, ok := p.Data[section]
	if !ok {

		p.Data[section] = make(map[string]string)
	}

}

func (p *Parser) SaveToFile(name string, dictionary map[string]map[string]string) (err error) {

	file, ferr := os.Create(name)
	defer file.Close()
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
	f, ferr := os.ReadFile(name)
	fmt.Println("openedfile")
	if ferr != nil {
		return errors.New("can't open the file with this name")
	}

	content := string(f)

	return p.LoadFromString(string(content))
}

func (p *Parser) LoadFromString(content string) (err error) {

	err = p.Parse(content)
	fmt.Println("error", err)
	return err

}
func checkLine(line string) (string, error) {
	splits := strings.Split(line, " ")
	equal := strings.Split(line, "=")
	if isValidSectionName(splits[0]) {
		return "section", nil

	} else if splits[0] == ";" {
		return "comment", nil

	} else if line == "\n" {
		return "empty", nil

	} else if len(equal) == 2 {
		return "KeyValue", nil
	}
	return "", errors.New("Invalid input")
}

func (p *Parser) Parse(content string) (err error) {
	scanner := bufio.NewScanner(strings.NewReader(content))

	var key string
	var value string
	var section string
	SectionFlag := false

	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")
		fmt.Println((items[0]))
		if items[0] == ";" {

			continue

		} else if isValidSectionName(items[0]) {
			section = items[0]

			fmt.Println("section:", section)
			p.SetSections(section)
			SectionFlag = true

		} else if SectionFlag == true {
			split_equal := strings.Split(line, "=")

			if len(split_equal) == 2 {

				key = split_equal[0]
				value = split_equal[1]
				p.SetValues(section, key, value)

			} else if len(split_equal) > 2 {
				return errors.New("more than one value")
			}

		} else if len(line) == 0 {
			continue

		} else {

			return errors.New("Invalid input Name")
		}

	}
	p.SaveToFile("name.txt", p.Data)
	return err
}
