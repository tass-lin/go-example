package main

import (
	"fmt"
	"sort"
)

type data struct {
	grade  int
	height int
}

func extractkeys(m map[string]data) (keys []string) {
	for key := range m {
		keys = append(keys, key)
	}
	return
}

func mapsort(m map[string]data) []string {
	names := extractkeys(m)
	sort.Slice(names, func(i, j int) bool {
		return m[names[i]].height > m[names[j]].height
	})
	return names
}

func main() {
	students := map[string]data{
		"mary":  {grade: 10, height: 167},
		"paul":  {grade: 7, height: 172},
		"peter": {grade: 8, height: 164}}
	fmt.Println(mapsort(students))
}
