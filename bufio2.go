package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	var filename string
	fmt.Print("檔案名稱：")
	fmt.Scanf("%s", &filename);

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bufio.NewReader(f).WriteTo(os.Stdout)
}

//NewWriter 預設的緩衝區為 4096 位元組，由於這邊使用標準輸出，在緩衝區未滿前，資料不會寫出，可以使用 Flush 來出清緩衝區中的資料。
//
//事實上，對於需要逐行讀取的需求，使用 bufio.Scanner 會比較方便，可以使用 NewScanner 來建立實例，建立之後有以下的方法可以使用：