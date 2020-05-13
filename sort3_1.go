package main

import (
	"fmt"
	"sort"
)

func extractkeys(m map[string]int) (keys []string) {
	for key := range m {
		keys = append(keys, key)
	}
	return
}

func mapsort(m map[string]int) []string {
	names := extractkeys(m)
	sort.Slice(names, func(i, j int) bool {
		return m[names[i]] > m[names[j]]
	})
	return names
}

func main() {
	grades := map[string]int{"mary": 10, "paul": 7, "peter": 8}
	fmt.Println(mapsort(grades))
}
