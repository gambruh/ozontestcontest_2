package main

import (
	"bufio"
	"fmt"
	"os"
)

type messenger struct {
	lastMessageID     int
	lastGlobalMessage int
	customers         map[int]int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q int
	fmt.Fscan(in, &n, &q)

	m := NewMessenger()

	for i := 0; i < q; i++ {
		var t, id, messageID int
		fmt.Fscan(in, &t, &id)

		if t == 1 {
			m.sendMessage(id)
		}

		if t == 2 {
			messageID = m.getLastMessageNum(id)
			fmt.Fprintln(out, messageID)
		}

	}

}

func NewMessenger() messenger {
	return messenger{
		lastMessageID:     0,
		lastGlobalMessage: 0,
		customers:         make(map[int]int),
	}
}

func (m *messenger) sendMessage(id int) {
	if id == 0 {
		m.sendGlobalMessage()
		return
	}

	m.lastMessageID++
	m.customers[id] = m.lastMessageID
}

func (m *messenger) sendGlobalMessage() {
	m.lastMessageID++
	m.lastGlobalMessage = m.lastMessageID
	for k := range m.customers {
		m.customers[k] = m.lastMessageID
	}
}

func (m messenger) getLastMessageNum(id int) int {
	var ans int
	if num, ok := m.customers[id]; ok {
		ans = num
	} else {
		ans = m.lastGlobalMessage
	}
	return ans
}
