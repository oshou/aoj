package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func init() {
	sc.Split(bufio.ScanWords)
}

func readInt() int {
	sc.Scan()
	r, _ := strconv.Atoi(sc.Text())
	return r
}

type block struct {
	h, w int
}

func main() {
	var h, w int
	var blocks []block
	for {
		h, w = readInt(), readInt()
		if h == 0 && w == 0 {
			break
		}
		blocks = append(blocks, block{h: h, w: w})
	}
	for _, block := range blocks {
		for i := 0; i < block.h; i++ {
			for j := 0; j < block.w; j++ {
				if i%2 == 0 {
					if j%2 == 0 {
						fmt.Print("#")
					} else {
						fmt.Print(".")
					}
				} else {
					if j%2 == 0 {
						fmt.Print(".")
					} else {
						fmt.Print("#")
					}
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

}
