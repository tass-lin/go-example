package main

import (
	"fmt"
	"container/ring"
)

type Person struct {
	Number int
}

func main() {
	persons := ring.New(41)
	// 給每個人編號
	for i := 1; i <= persons.Len(); i++ {
		persons.Value = &Person{i}
		persons = persons.Next()
	}

	persons = persons.Prev()
	// 最後只留下兩人
	for persons.Len() > 2 {
		for i := 1; i <= 2; i++ {
			persons = persons.Next()
		}

		// 報數 3 Out
		persons.Unlink(1)

		persons.Do(func(p interface{}) {
			person := p.(*Person)
			fmt.Printf("%d ", person.Number)
		})
		fmt.Println("\n")
	}

	fmt.Print("安全位置：")
	persons.Do(func(p interface{}) {
		person := p.(*Person)
		fmt.Printf("%d ", person.Number)
	})
}

//據說著名猶太歷史/數學家約瑟夫（Josephus）有過以下的故事：在羅馬人佔領喬塔帕特後，40個猶太士兵與約瑟夫躲到一個洞中，眼見脫逃無望，一群人決定集體自殺，然而私下約瑟夫與某個傢伙並不贊成，於是約瑟夫建議自殺方式，41個人排成圓圈，由第1個人 開始報數，每報數到3的人就必須自殺，然後由下一個重新報數，直到所有人都自殺身亡為止。約瑟夫與不想自殺的那個人分別排在第16個與第31個位置，於是逃過了這場死亡遊戲。