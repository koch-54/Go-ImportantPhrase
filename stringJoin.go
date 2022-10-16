package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	s := make([]string, 0) //空の配列をわたす

	for i := 0; i < 3; i++ {
		sc.Scan()
		s = append(s, sc.Text())
	}

	result := strings.Join(s, "|") //配列を"|"で区切った文字列を渡す

	fmt.Println(result)
}
