package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	// 호출한 함수가 종료 직전에 실행된다.
	// 주로 cleanup 에 사용된다.
	// control flow mechanism, 지연 실행
	defer f.Close()
	scanner := bufio.NewScanner(f) // bufio.Scanner

	fmt.Print("What is your name? ")
	scanner.Scan() // read one line from os.Stdin
	name := scanner.Text()
	fmt.Printf("Hello, %s, nice to meet you! \n", name)
}
