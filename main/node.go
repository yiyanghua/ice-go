package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	val  int
	next *ListNode
}

func (l *ListNode) sum() int {
	sum := 0
	n := l
	for {
		if sum == 0 {
			sum = n.val
		} else {
			len := len(strconv.Itoa(sum))
			sum = n.val*len*10 + sum
		}
		if n.next == nil {
			break
		}
		n = n.next
	}
	return sum
}
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := l1.sum() + l2.sum()
	a := strings.Split(strconv.Itoa(sum), "")
	r := &ListNode{}
	var pre *ListNode
	for _, v := range a {
		c := &ListNode{}
		va, _ := strconv.Atoi(v)
		c.val = va

		if pre == nil {
			r = c
		} else {
			pre.next = c
		}
		pre = c
	}
	return r
}

func main() {
	p := &ListNode{
		val: 2,
	}
	c := &ListNode{
		val: 3,
	}
	p.next = c

	sum := addTwoNumbers(p, p)

	fmt.Println(sum.sum())
}
