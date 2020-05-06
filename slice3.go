package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6}
	slice1 := arr[0:2:4]
	slice2 := arr[0:2]
	fmt.Println(slice1)      // [1 2]
	fmt.Println(len(slice1)) // 2
	fmt.Println(cap(slice1)) // 4
	fmt.Println(slice2)      // [1 2]
	fmt.Println(len(slice2)) // 2
	fmt.Println(cap(slice2)) // 6
}