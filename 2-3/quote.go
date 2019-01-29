package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()
	scanner := bufio.NewScanner(f) // bufio.Scanner

	fmt.Print("What is the quote? ")
	scanner.Scan() // read one line from os.Stdin
	quote := scanner.Text()

	fmt.Print("Who said it? ")
	scanner.Scan()
	who := scanner.Text()

	// fmt.Printf("%v says, \"%v\"\n", who, quote)
	fmt.Println(who, "says,", "\""+quote+"\"")
}
