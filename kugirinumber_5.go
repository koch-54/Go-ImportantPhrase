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
	s := strings.Split(sc.Text(), " ") //空白を区切りにしているいるという意味

	for i := 0; i < 5; i++ {
		fmt.Println(s[i])
	}
}
