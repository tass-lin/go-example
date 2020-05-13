package main

import (
	"fmt"
	"sort"
)

type keyser interface {
	keys() []string
}

type less func(s1, s2 string) bool

func mapsort(m keyser, f less) []string {
	names := m.keys()
	sort.Slice(names, func(i, j int) bool {
		return f(names[i], names[j])
	})
	return names
}

type data struct {
	grade  int
	height int
}

type mapper map[string]data

func (m mapper) keys() (keys []string) {
	for key := range m {
		keys = append(keys, key)
	}
	return
}

func main() {
	students := mapper{
		"mary":  {grade: 10, height: 167},
		"paul":  {grade: 7, height: 172},
		"peter": {grade: 8, height: 164}}
	names := mapsort(students, func(s1, s2 string) bool {
		return students[s1].grade > students[s2].grade
	})
	fmt.Println(names)
	names = mapsort(students, func(s1, s2 string) bool {
		return students[s1].height > students[s2].height
	})
	fmt.Println(names)
}
