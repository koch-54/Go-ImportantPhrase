package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	sc.Text()

	sc.Scan()
	s := strings.Split(sc.Text(), " ")

	for _, v := range s {
		fmt.Println(v)
	}
}
