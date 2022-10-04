package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	number, _ := strconv.Atoi(sc.Text())

	for i := 1; i <= number; i++ {
		sc.Scan()
		name := sc.Text()
		fmt.Println(name)
	}
}
