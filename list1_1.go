package main

import (
	"fmt"
	"container/list"
)

type Person struct {
	Name string
	Age  int
}

type PersonQueue struct {
	list *list.List
}

func NewPersonQueue() *PersonQueue {
	return &PersonQueue{list.New()}
}

func (q *PersonQueue) Len() int {
	return q.list.Len()
}

func (q *PersonQueue) Offer(p *Person) {
	q.list.PushBack(p)
}

func (q *PersonQueue) Peek() *Person {
	if q.list.Len() == 0 {
		return nil
	}

	e := q.list.Remove(q.list.Front())
	return e.(*Person)
}

func main() {
	q := NewPersonQueue()

	q.Offer(&Person{"Irene", 12})
	q.Offer(&Person{"Justin", 45})
	q.Offer(&Person{"Monica", 42})

	for p := q.Peek(); p != nil; p = q.Peek() {
		fmt.Printf("姓名：%s\t年齡：%d\n", p.Name, p.Age)
	}
}

//因此，並不是 list.List 不常用，而是你可能很少自行實現資料結構（都拿別人寫好的來用？）；另一種說法「每當想使用 list.List 時，都該思考一下是否優先使用 slice。」的說法也不是完全正確…
//
//若想使用 list.List，應該問的是，你的資料結構在實現上需要雙向鏈結的特性嗎？例如，也許你會需要有個具索引的資料結構，同時底層實現必須是雙向鏈結（像是 Java 的 LinkedList）？那麼就可以考慮透過 list.List 來實現。