package main

import (
	"container/heap"
	"fmt"
	"sort"
	"strings"
)

type mq []int

func (q mq) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
}

func (q mq) Less(i, j int) bool {
	return q[i]<q[j]
}

func (q mq) Len() int {
	return len(q)
}

func (q *mq) Push(x interface{}) {
	qv := *q
	item := x.(int)
	*q = append(qv, item)
	//fmt.Println(*q)
}

func (q *mq) Pop() interface{} {
	qv := *q
	leng := len(qv)
	item := qv[leng-1]
	*q = qv[:leng-1]
	return item
}

func main() {
	raw :=[]int{45,1,21,30,3,2,17,56,23,26,91,73,66}
	q := make(mq,0)
	heap.Init(&q)

	for i:=0; i<len(raw); i++ {
		heap.Push(&q, raw[i])
	}
	heap.Push(&q, 54)
	for i:=0; i<len(raw); i++ {
		r := heap.Pop(&q)
		fmt.Println(r)
	}

	sh := show{1}
	try(&sh)

	trysort()

	trystring()
}

type show struct {
	v int
}

type function interface {
	up(n int)
	down(n int)
}

func (s *show) up(n int) {
	s.v = n
	print("up: ", s)
}

func (s show) down(n int) {
	s.v = n
	print("down: ", &s)
}

func try(f function) {
	f.up(13)
	fmt.Println(f)
	f.down(2)
	fmt.Println(f)
}

func trysort() {
	a := []int{3, 5, 4, -1, 9, 11, -14}
	sort.Ints(a)
	fmt.Println(a)

	ss := []string{"surface", "ipad", "mac pro", "mac air", "think pad", "idea pad"}
	sort.Strings(ss)
	for _, item := range ss {
		fmt.Println(item)
	}

	sort.Sort(sort.Reverse(sort.StringSlice(ss)))
	fmt.Printf("After reverse: %v\n", ss)
	for _, item := range ss {
		fmt.Println(item)
	}

}

func trystring() {
	fmt.Println(strings.Count("haha0980hahauahah", "haha"))
}