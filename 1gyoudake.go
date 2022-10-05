// N n_1 n_2 ... n_N　のうち最初の一つ以外の入力値を縦に表示する。
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
	s := strings.Split(sc.Text(), " ")
	n := len(s)

	for i := 1; i < n; i++ {
		fmt.Println(s[i])
	}
}
