package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {
	result := ""
	for i, char := range strings.Split("Rogerio", "") {
		if !(i%4 == 3) {
			result += char
		} else {
			result += "_"
		}
	}
	fmt.Println(result) //Rog_rio

	var text = "{\"data\":\"key=abcde, age=58, key=aaaaa, age=49\"}" 
	p(strings.Split(text[9:len(text)-2], ", ")[1]) // 'age=58'

	p("Contains:  ", strings.Contains("test", "es"))          //Contains:   true
	p("ContainsAny:  ", strings.ContainsAny("test", "tz"))    //ContainsAny:   true
	p("Count:     ", strings.Count("test", "t"))              //Count:      2
	p("HasPrefix: ", strings.HasPrefix("test", "te"))         //HasPrefix:  true
	p("HasSuffix: ", strings.HasSuffix("test", "st"))         //HasSuffix:  true
	p("Index:     ", strings.Index("test", "es"))             //Index:      1
	p("IndexAny:  ", strings.IndexAny("Japan", "abc"))        //IndexAny:  1
	p("LastIndex: ", strings.LastIndex("Japan", "abc"))       //LastIndex:  -1
	p("LastIndexAny: ", strings.LastIndexAny("Japan", "abc")) //LastIndexAny:3
	p("Join:      ", strings.Join([]string{"a", "b"}, "-"))   //Join:       a-b
	p("Repeat:    ", strings.Repeat("a", 5))                  //Repeat:     aaaaa
	p("Replace:   ", strings.Replace("foo", "o", "0", -1))    //Replace:    f00
	p("Replace:   ", strings.Replace("foo", "o", "0", 1))     //Replace:    f0o
	p("Split:     ", strings.Split("a-b-c-d-e", "-"))         //Split:      [a b c d e]
	p("ToLower:   ", strings.ToLower("TEST"))                 //ToLower:    test
	p("ToUpper:   ", strings.ToUpper("test"))                 //ToUpper:    TEST
	p("Fields:   ", strings.Fields(" a\t b\n"))               //Fields:    ["a" "b"]
	p("Split:   ", strings.Split("a,b", ","))                 //Split:    ["a" "b"]
	p("SplitAfter:   ", strings.SplitAfter("a,b", ","))       //SplitAfter:    ["a," "b"]

	bytes := [4]byte{'0', '1', '2', '3'}
	p("Bytes array to string", string(bytes[:]))
	p("Array to slice", bytes[:])
	p("Rogerio"[0:2])

	pi := 2
	p("Insert into string", "Rorio"[:pi]+"ge"+"Rorio"[pi:])          // Rogerio
	p("Replace in string", strings.Replace("Rugeriu", "u", "o", -1)) // Rogerio
	pr, ncr := 2, 2
	p("Remove from string", "Rodegerio"[:pr]+"Rodegerio"[pr+ncr:]) //Rogerio

	insert := func(s string, t string, p int) string {
		return s[:pi] + t + s[pi:]
	}
	p("Insert into string (function)", insert("Rorio", "ge", 2)) //Rogerio
	remove := func(s string, p int, nc int) string {
		return s[:pr] + s[pr+nc:]
	}
	p("Remove from string (function)", remove("Rodegerio", 2, 2)) //Rogerio

	p(strings.Cut("Part1,Part2", ",")) // Part1 Part2 true
}
