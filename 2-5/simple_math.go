package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()
	scanner := bufio.NewScanner(f)

	print("What is the first number? ")
	scanner.Scan()
	ops1 := scanner.Text()
	op1, _ := strconv.Atoi(ops1)

	print("What is the second number? ")
	scanner.Scan()
	ops2 := scanner.Text()
	op2, _ := strconv.Atoi(ops2)

	fmt.Printf("%v + %v = %v\n", op1, op2, op1+op2)
	fmt.Printf("%v - %v = %v\n", op1, op2, op1-op2)
	fmt.Printf("%v * %v = %v\n", op1, op2, op1*op2)
	fmt.Printf("%v / %v = %v\n", op1, op2, op1/op2)
}
