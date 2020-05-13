package main

import (
	"fmt"
	"sort"
)

type mapsorter struct {
	names  []string
	record map[string]int
}

func (m *mapsorter) Len() int {
	return len(m.names)
}

func (m *mapsorter) Swap(i, j int) {
	m.names[i], m.names[j] = m.names[j], m.names[i]
}

func (m *mapsorter) Less(i, j int) bool {
	return m.record[m.names[i]] > m.record[m.names[j]]
}

func extractkeys(m map[string]int) (keys []string) {
	for key := range m {
		keys = append(keys, key)
	}
	return
}

func main() {
	grades := map[string]int{"mary": 10, "paul": 7, "peter": 8}
	ms := &mapsorter{names: extractkeys(grades), record: grades}
	sort.Sort(ms)
	fmt.Println(ms.names)
}

//https://www.facebook.com/groups/269001993248363/?ref=bookmarks