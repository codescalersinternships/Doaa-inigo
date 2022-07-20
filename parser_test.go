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
		fmt.Println("got :", got, " want :", want)
		if got != want {
			t.Errorf("got: %q , want: %q", got, want)
		}
	})

	t.Run("Invalid syntax", func(t *testing.T) {
		got, _ := checkLine("; hello i am Doaa")
		want := "comment"
		fmt.Println("got :", got, " want :", want)
		if got != want {
			t.Errorf("got: %q , want: %q", got, want)
		}
	})

	t.Run("Invalid syntax", func(t *testing.T) {
		got, _ := checkLine("name = Doaa = amira ")
		want := " "
		fmt.Println("got :", got, " want :", want)
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
		//fmt.Println("got :", got, " want :", want)
		res1 := reflect.DeepEqual(got, want)
		if !res1 {
			//fmt.Println(err)
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
		fmt.Println("first got", got)
		want := make(map[string]map[string]string)
		want["[owner]"] = make(map[string]string)
		want["[owner]"]["name"] = "John Doe"
		want["[owner]"]["organization"] = "Acme Widgets Inc."
		want["[database]"] = make(map[string]string)

		//fmt.Println("got :", got, " want :", want)
		res1 := reflect.DeepEqual(got, want)
		if !res1 {
			//fmt.Println(err)
			t.Errorf("got: %q , want: %q", got, want)
		}
	})

	// not working
	t.Run("ini text", func(t *testing.T) {
		text := " last modified 1 April = 2001 by John Doe\n" +
			"[owner]\n" + "name=John Doe\n" + "organization=Acme Widgets Inc.\n" +
			"\n" + "[database]\n" + "; use IP address in case network name resolution is not working\n" +
			"server=192.0.2.62\n" + "port=143\n" + "file=payroll.dat\n"

		got, err := Parse(text)
		if err != nil {
			//	t.Error(fmt.Sprintf("Error in parsing: '%v'", err))
			t.Fatalf("Error in parsing: '%v'", err)

		}

		want := make(map[string]map[string]string)
		res1 := reflect.DeepEqual(got, want)
		if !res1 {
			//fmt.Println(err)
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
		fmt.Println("first got", got)
		want := make(map[string]map[string]string)
		want["[owner]"] = make(map[string]string)
		want["[owner]"]["name"] = "John Doe"
		want["[owner]"]["organization"] = "Acme Widgets Inc."
		want["[database]"] = make(map[string]string)
		want["[database]"]["server"] = "192.0.2.62"
		want["[database]"]["port"] = "143"
		want["[database]"]["file"] = ""

		//fmt.Println("got :", got, " want :", want)
		res1 := reflect.DeepEqual(got, want)
		if !res1 {
			//fmt.Println(err)
			t.Errorf("got: %q , want: %q", got, want)
		}
	})

}
func TestGet(t *testing.T) {
	parser := Parser{}
	t.Run("ini text", func(t *testing.T) {
		parser.SetValues("owner", "location", "Cairo")
		parser.SetValues("owner", "Salary", "10000")
		got, _ := parser.Get("database", "location")
		want := ""

		if got != want {
			t.Errorf("got: %q , want: %q", got, want)

		}

	})

	t.Run("ini text", func(t *testing.T) {
		parser.SetValues("owner", "salary", "h")
		parser.SetValues("owner", "", "10000")
		got, _ := parser.Get("owner", "location")
		want := ""

		if got != want {
			t.Errorf("got: %q , want: %q", got, want)

		}

	})

}

/*func TestLoadFromFile(t *testing.T) {
	//not working
	parser := Parser{}
	t.Run("ini file", func(t *testing.T) {

		got := parser.LoadFromFile("doaa.INI")
		want := errors.New("can't open the file with this name")
		if got != want {
			t.Errorf("got: %#v , want: %#v", got, want)
		}
	})

}*/

/*
checking section name
file name
empty file
sections without key, values
more than value (username = doaa =amira )
get sections
save file () invalidname
get value
search for non existing key , section
*/
/*text:="; last modified 1 April 2001 by John Doe \n"+
"[owner]\n"+
"name = John Doe\n"+
"organization = Acme Widgets Inc.\n"+
"\n"+
"[database]\n"+
"; use IP address in case network name resolution is not working\n"+
"server = 192.0.2.62\n"+
"port = 143\n"+
"file = payroll.dat\n"*/
