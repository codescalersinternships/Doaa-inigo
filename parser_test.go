package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCheckLine(t *testing.T) {
	t.Run("Invalid syntax", func(t *testing.T) {
		got, _ := checkLine("[[owner")
		want := " "
		if got != want {
			t.Errorf("got: %q , want: %q", got, want)
		}
	})
	t.Run("Invalid syntax", func(t *testing.T) {
		got, _ := checkLine("[section] ; comment")
		want := " "
		if got != want {
			t.Errorf("got: %q , want: %q", got, want)
		}
	})

	t.Run("Invalid syntax", func(t *testing.T) {
		got, _ := checkLine("; hello i am Doaa")
		want := "comment"
		if got != want {
			t.Errorf("got: %q , want: %q", got, want)
		}
	})

	t.Run("Invalid syntax", func(t *testing.T) {
		got, _ := checkLine("name = Doaa = amira ")
		want := " "
		if got != want {
			t.Errorf("got: %q , want: %q", got, want)
		}
	})

}
func TestParse(t *testing.T) {
	t.Run("ini text", func(t *testing.T) {

		text := "; last modified 1 April 2001 by John Doe\n" +
			"[owner]\n" + "name=John Doe\n" + "organization=Acme Widgets Inc.\n" +
			"\n" + "[database]\n" + "; use IP address in case network name resolution is not working\n" +
			"server=192.0.2.62\n" + "port=143\n" + "file=payroll.dat\n"

		got, err := Parse(text)
		if err != nil {
			t.Error(fmt.Sprintf("Error in parsing: '%v'", err))
		}

		want := make(map[string]map[string]string)
		want["[owner]"] = make(map[string]string)
		want["[owner]"]["name"] = "John Doe"
		want["[owner]"]["organization"] = "Acme Widgets Inc."
		want["[database]"] = make(map[string]string)
		want["[database]"]["server"] = "192.0.2.62"
		want["[database]"]["port"] = "143"
		want["[database]"]["file"] = "payroll.dat"
		res1 := reflect.DeepEqual(got, want)
		if !res1 {
			t.Errorf("got: %q , want: %q", got, want)

		}

	})
	t.Run("ini text", func(t *testing.T) {

		text :=
			"[owner]\n" + "name=John Doe\n" + "organization=Acme Widgets Inc.\n" +
				"\n" + "[database]\n"

		got, err := Parse(text)
		if err != nil {
			t.Error(fmt.Sprintf("Error in parsing: '%v'", err))
		}
		want := make(map[string]map[string]string)
		want["[owner]"] = make(map[string]string)
		want["[owner]"]["name"] = "John Doe"
		want["[owner]"]["organization"] = "Acme Widgets Inc."
		want["[database]"] = make(map[string]string)

		res1 := reflect.DeepEqual(got, want)
		if !res1 {
			t.Errorf("got: %q , want: %q", got, want)
		}
	})

	t.Run("ini text", func(t *testing.T) {
		text := " last modified 1 April = 2001 by John Doe\n" +
			"[owner]\n" + "name=John Doe\n" + "organization=Acme Widgets Inc.\n" +
			"\n" + "[database]\n" + "; use IP address in case network name resolution is not working\n" +
			"server=192.0.2.62\n" + "port=143\n" + "file=payroll.dat\n"

		got, _ := Parse(text)

		want := make(map[string]map[string]string)
		res1 := reflect.DeepEqual(got, want)
		if !res1 {
			t.Errorf("got: %#v , want: %#v", got, want)

		}
	})

	t.Run("ini text", func(t *testing.T) {

		text := "; last modified 1 April 2001 by John Doe\n" +
			"[owner]\n" + "name=John Doe\n" + "organization=Acme Widgets Inc.\n" +
			"\n" + "[database]\n" + "; use IP address in case network name resolution is not working\n" +
			"server=192.0.2.62\n" + "port=143\n" + "file="
		got, err := Parse(text)
		if err != nil {
			t.Error(fmt.Sprintf("Error in parsing: '%v'", err))
		}
		want := make(map[string]map[string]string)
		want["[owner]"] = make(map[string]string)
		want["[owner]"]["name"] = "John Doe"
		want["[owner]"]["organization"] = "Acme Widgets Inc."
		want["[database]"] = make(map[string]string)
		want["[database]"]["server"] = "192.0.2.62"
		want["[database]"]["port"] = "143"
		want["[database]"]["file"] = ""

		res1 := reflect.DeepEqual(got, want)
		if !res1 {
			t.Errorf("got: %q , want: %q", got, want)
		}
	})

}

func TestGet(t *testing.T) {
	parser := Parser{}

	t.Run("CheckingValues", func(t *testing.T) {
		parser.SetValues("owner", "salary", "h")
		parser.SetValues("owner", "", "10000")
		got, _ := parser.Get("owner", "location")
		want := ""

		if got != want {
			t.Errorf("got: %q , want: %q", got, want)

		}

	})

}

func TestLoadFromFile(t *testing.T) {
	parser := Parser{}
	t.Run("File Test", func(t *testing.T) {
		//checking on existing file
		got := parser.LoadFromFile("text.INI")

		if got != nil {
			t.Errorf("got: %#v ", got)
		}
	})

}

func TestGetSectionNames(t *testing.T) {
	parser := Parser{}
	t.Run("SectionsName", func(t *testing.T) {

		parser.SetValues("owner", "location", "Cairo")
		parser.SetValues("database", "Salary", "10000")
		got := parser.GetSectionNames()
		want := []string{"owner", "database"}
		res := reflect.DeepEqual(got, want)
		if !res {
			t.Errorf("got: %q , want: %q", got, want)

		}
	})

}
func TestGetSections(t *testing.T) {
	parser := Parser{}
	t.Run("CheckingMap", func(t *testing.T) {

		parser.SetValues("[owner]", "location", "Cairo")
		parser.SetValues("[database]", "Salary", "10000")
		parser.SetValues("[database]", "port", "143")

		got := parser.GetSections()
		want := make(map[string]map[string]string)
		want["[owner]"] = make(map[string]string)
		want["[owner]"]["location"] = "Cairo"
		want["[database]"] = make(map[string]string)
		want["[database]"]["Salary"] = "10000"
		want["[database]"]["port"] = "143"

		res1 := reflect.DeepEqual(got, want)
		if !res1 {
			t.Errorf("got: %q , want: %q", got, want)
		}
	})

}
func TestSearchSection(t *testing.T) {
	parser := Parser{}
	t.Run("SearchForSectionsNames", func(t *testing.T) {
		parser.SetValues("[owner]", "location", "Cairo")
		parser.SetValues("[database]", "Salary", "10000")
		got := parser.SearchSection("[owner]")
		if got != nil {
			t.Errorf("Inavalid")

		}

	})

}
