package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt64)
}

func readString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	s := readString()
	cnt := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "1" {
			cnt++
		}
	}
	fmt.Println(cnt)
}
