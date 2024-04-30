package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		runners := make([]int, n)
		for j := 0; j < n; j++ {
			var time int
			fmt.Fscan(in, &time)
			runners[j] = time
		}
		places := resolveTask(runners)

		for _, p := range places {
			fmt.Fprint(out, p)
		}
		fmt.Fprintln(out)
	}
}

func resolveTask(runners []int) []int {

	places := make([]int, len(runners))

	sorted := sortArray(runners)
	fmt.Println(sorted)

	return places
}

func sortArray(runners []int) []int {
	sorted := make([]int, len(runners))
	copy(sorted, runners)
	sort.Ints(sorted)
	return sorted
}

func assignPlaces(sorted []int) map[int]int {
	curplace := 1

	places := make(map[int]int)
	buf := 0

	for i := 0; i < len(sorted); i++ {
		if i > 0 {
			if sorted[i]-sorted[i-1] > 1 {
				curplace = curplace + buf
				buf = 0
			} else {
				buf++
			}
		}

		places[sorted[i]] = curplace
	}

	return places
}
