package kojiro

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	result := []int{}
	if ln == nil {
		return ""
	}
	if ln.Next == nil {
		return fmt.Sprintf("%d", ln.Val)
	}

	for current := ln.Next; current != nil; {

	}

	return result.String()
}

func NewListNode(number int) *ListNode {
	numChars := strings.Split(fmt.Sprintf("%d", number), "")

	var last *ListNode = nil

	for _, numChar := range numChars {
		num, err := strconv.Atoi(numChar)
		if err != nil {
			log.Fatalln(fmt.Errorf("Convert to int error %w", err))
		}
		newNode := &ListNode{
			Val: num,
		}
		if last != nil {
			newNode.Next = last
		}

		last = newNode
	}

	return last
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var root *ListNode = nil
	l1p := l1
	l2p := l2
	lastNode := root

	for l1p != nil || l2p != nil {
		var newNode *ListNode = nil

		if l1p != nil && l2p != nil {
			newNode = &ListNode{
				Val: l1p.Val + l2p.Val,
			}

			l1p = l1p.Next
			l2p = l2p.Next
		} else if l2p == nil {
			newNode = &ListNode{
				Val: l1p.Val,
			}

			l1p = l1p.Next
		} else if l1p == nil {
			newNode = &ListNode{
				Val: l2p.Val,
			}

			l2p = l2p.Next
		}

		if lastNode != nil {
			(*lastNode).Next = newNode
		}
		if root == nil {
			root = newNode
		}
		lastNode = newNode
	}

	return root
}
