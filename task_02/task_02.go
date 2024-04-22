package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var inputString string
	fmt.Fscan(in, &inputString)

	m := len(inputString)

	output := make([]byte, m)

	for i := range inputString {
		output[i] = inputString[i]
	}

	var n int

	fmt.Fscan(in, &n)

	for i := 0; i < n; i++ {
		var l, r int
		var new string

		fmt.Fscan(in, &l, &r, &new)

		for j := l - 1; j < r; j++ {
			output[j] = new[j-l+1]
		}

	}

	fmt.Fprint(out, string(output))

}
