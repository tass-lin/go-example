package main

import "fmt"

type Consumer = func(int)

func forEach(elems []int, consumer Consumer) {
	for _, elem := range elems {
		consumer(elem)
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, elem := range numbers {
		sum += elem
	}
	fmt.Println(sum) // 15
}

//然而意義完全不同。在使用 forEach 函式的範例中，sum 變數被匿名函式包覆並傳入 forEach 之中，在 forEach 執行迴圈的過程中，每次呼叫傳入的函式（被 consumer 參考），就會改變 sum 的值，因此，最後得到的是加總後的值 15。
//
//實際上，使用 forEach 函式的範例中，建立了一個閉包，閉包本質上就是一個匿名函式，sum 變數被閉包包覆，讓 sum 變數可以存活於閉包的範疇中，其實，更之前從 funcA 傳回函式的範例中，也建立了閉包，funcA 的 x 區域變數被閉包包覆，因此，你執行傳回的函式時，即使 funcA 已執行完畢，x 變數依然是存活著在傳回的閉包範疇中，所以，你指定的引數總是會與 x 的值進行相加。