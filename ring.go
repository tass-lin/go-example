package main

import (
	"fmt"
	"container/ring"
)

func main() {
	numbers := ring.New(10)
	for i := 0; i < numbers.Len(); i++ {
		numbers.Value = i
		numbers = numbers.Next()
	}

	numbers.Do(func(n interface{}) {
		fmt.Printf("%d ", n.(int))
	})
}

//實際應用上，ring 可以用來管理有限筆數的歷史記錄、輪播等。