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

	print("Enter a noun ")
	scanner.Scan()
	noun := scanner.Text()

	print("Enter a verb ")
	scanner.Scan()
	verb := scanner.Text()

	print("Enter a adjective ")
	scanner.Scan()
	adjective := scanner.Text()

	print("Enter aadverb ")
	scanner.Scan()
	adverb := scanner.Text()

	fmt.Printf("Do you %v your %v %v %v? That's hilarious!\n",
		noun, verb, adjective, adverb)
}
