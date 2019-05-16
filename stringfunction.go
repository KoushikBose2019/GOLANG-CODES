package main
 
import (
  "fmt"
  "strings"
)
 
func main() {
  fmt.Println(strings.Compare("A", "B"))  // A < B
  fmt.Println(strings.Compare("B", "A")) 
  fmt.Println(strings.Compare("india","Germany"))
  fmt.Println(strings.Compare("germany","us"))
  fmt.Println(strings.Compare(""," "))
  fmt.Println(strings.Compare(""," "))*/
  fmt.Println(strings.Contains("india","inda"))
  fmt.Println(strings.Contains("Australia", "Australian")) 
    fmt.Println(strings.Contains("Japan", "JAPAN")) // Case sensitive
    fmt.Println(strings.Contains("Japan", "JAP")) // Case sensitive
    fmt.Println(strings.Contains("Become inspired to travel to Australia.", "Australia"))
    fmt.Println(strings.Contains("", ""))
    fmt.Println(strings.Contains("  ", " ")) // space also consider as string
	fmt.Println(strings.Contains("12554", "1"))*/
	 
    fmt.Println(strings.ContainsAny("india","a"))
    fmt.Println(strings.ContainsAny("Australia", "r & a"))
   fmt.Println(strings.ContainsAny("JAPAN", "j"))
   fmt.Println(strings.ContainsAny("JAPAN", "J"))
   fmt.Println(strings.ContainsAny("JAPAN", "JAPAN"))  
   fmt.Println(strings.ContainsAny("JAPAN", "japan"))
	fmt.Println(strings.ContainsAny("Shell-12541", "1"))
	 //  Contains vs ContainsAny
   fmt.Println(strings.ContainsAny("Shell-12541", "1-2")) // true
  fmt.Println(strings.Contains("Shell-12541", "1-2"))   // false

/* count func*/
fmt.Println(strings.Count("india","i"))
//Equalfold func compare two strings are equal or not//
fmt.Println(ings.Equalfold("india","INDIA"))






}