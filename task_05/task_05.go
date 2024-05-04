package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Folder struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	// Wrap bufio.Reader into Scanner

	for i := 0; i < n; i++ {
		data := readJSON(in)
		fmt.Println(data)
	}
}

func readJSON(in *bufio.Reader) Folder {
	var buf bytes.Buffer
	var lines int
	scanner := bufio.NewScanner(in)
	fmt.Fscan(in, &lines)

	for j := 0; j <= lines; j++ {
		scanner.Scan()
		buf.WriteString(scanner.Text())
	}

	decoder := json.NewDecoder(&buf)
	var data Folder
	if err := decoder.Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return Folder{}
	}

	return data
}

func resolveTask(data Folder, corrupted bool, counter int) int {

	for i, file := range data.Files {
		parts := strings.Split(file, ".")
		if parts[len(parts)-1] == "hack" {
			res += len(data.Files)
		}
	}
	return res
}
