# Doaa-inigo

INI Parser
Go package provides read and write for INI files.

Features
Load from files and strings.
Set values and add keys and sections.
Get parsed data as a map.
Get sections names as a slice.
Get parsed data as a string.
Export parsed data to INI file.


How To Use
parser := Parser{}
you can parse from a file:

parser.LoadFromFile("file.ini")
or from a string:

iniText := `; last modified 1 April 2001 by John Doe
name = Test
[owner]
name = John Doe
organization = Acme Widgets Inc.

[database]
; use IP address in case network name resolution is not working
server = 192.0.2.62     
port = 143
file = "payroll.dat"`
parser.LoadFromString(iniText)

use parser.SetValues and sections("Section","Key","Value") to set keys and values to existing section
use parser.GetSections() get all the keys and values of the map
use parser.Get("section","key") get value of the section and key 
use parser.GetSectionsName() get all the keys of the map
use parser.SetSections("section") create new section and link it with new map[string]string
use parser.SaveToFile("file.ini") save all the keys and values of the map in file
use parser.SearchSection(section) search if the section is exist or no




