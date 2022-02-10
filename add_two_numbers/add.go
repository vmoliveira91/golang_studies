package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getListValue(list *ListNode) int {
	count := 0
	var value int

	for list != nil {
		value += list.Val * int(math.Pow(10, float64(count)))

		list = list.Next
		count++
	}

	return value
}

func addTwoNumber(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	var result *ListNode

	for l1 != nil || l2 != nil {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		sum = sum % 10

		if result == nil {
			result = &ListNode{Val: sum}
		} else {
			aux := result
			for aux.Next != nil {
				aux = aux.Next
			}
			aux.Next = &ListNode{Val: sum}
		}

	}

	if carry != 0 {
		aux := result
		for aux.Next != nil {
			aux = aux.Next
		}
		aux.Next = &ListNode{Val: carry}
	}

	return result
}

func main() {
	l1 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}
	l2 := &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}

	fmt.Println(getListValue(addTwoNumber(l1, l2)))
}
