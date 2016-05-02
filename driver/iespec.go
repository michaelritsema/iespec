package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*

Turns an iespec file into an ini file for ipfixcat

IE SPEC
input:

protocolIdentifier(44619/4)<unsigned8>

output:

[field "protocolIdentifier"]
type = unsigned8
enterprise = 44619
id = 4

*/

func printWords(words [][]string) {
	for _, line := range words {

		fmt.Printf("[field \"%s\"]\n", line[0])
		fmt.Printf("type = %s\n", line[3])
		fmt.Printf("enterprise = %s\n", line[1])
		fmt.Printf("id = %s\n\n", line[2])
	}
}

func main() {
	// open a file
	filename := flag.String("in", "example.iespec", "IESPEC input file")
	flag.Parse()

	if file, err := os.Open(*filename); err == nil {

		// make sure it gets closed
		defer file.Close()
		var words [][]string

		// create a new scanner and read the file line by line
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			//log.Println(line)
			w := strings.FieldsFunc(line, func(r rune) bool {
				switch r {
				case '(', ')', '/', '<', '>':
					return true
				}
				return false
			})
			//log.Println(w)
			words = append(words, w)
		}
		printWords(words)
		//log.Println(words)
		// check for errors
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}

	} else {
		log.Fatal(err)
	}

}
