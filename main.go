package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		if err := stdin.Err(); err != nil {
			return
		}
		line := stdin.Text()
		gs1 := GS1(line)
		if gs1.CheckDigitOK() {
			fmt.Println(gs1.ToJAN())
		} else {
			fmt.Println("check digit error")
		}
	}
}
