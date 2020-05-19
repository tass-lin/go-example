package main

import "fmt"

func fibonacci1(num int) int{
	if num == 0{
		return 0
	}
	if num == 1{
		return 1
	}
	return fibonacci1(num-1) + fibonacci1(num-2)
}

func fibonacci() func() int {
	a, b := -1, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	fmt.Println("----")
	fmt.Println(fibonacci1(10))
}