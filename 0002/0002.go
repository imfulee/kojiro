package kojiro

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) String() string {
	result := []string{}
	if ln == nil {
		return ""
	}
	if ln.Next == nil {
		return fmt.Sprintf("%d", ln.Val)
	}

	result = append(result, strconv.Itoa(ln.Val))
	current := ln.Next
	for current != nil {
		result = append(result, strconv.Itoa(current.Val))
		current = current.Next
	}

	slices.Reverse(result)
	return strings.Join(result, "")
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

func IsSameListNode(src, dst *ListNode) bool {
	if src == nil && dst == nil {
		return false
	}

	srcP := src
	dstP := dst

	for {
		if srcP == nil && dstP == nil {
			break
		}
		if srcP == nil || dstP == nil {
			return false
		}
		if srcP.Val != dstP.Val {
			return false
		}

		srcP = srcP.Next
		dstP = dstP.Next
	}

	return true
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var root *ListNode = nil
	l1p := l1
	l2p := l2
	lastNode := root
	incrementFromLastStage := 0

	for l1p != nil || l2p != nil {
		var newNode *ListNode = nil
		incrSetThisStage := false
		value := incrementFromLastStage

		if l1p != nil && l2p != nil {
			value += l1p.Val + l2p.Val
			digit := value % 10
			newNode = &ListNode{
				Val: digit,
			}

			l1p = l1p.Next
			l2p = l2p.Next
		} else if l2p == nil {
			value += l1p.Val
			digit := value % 10
			newNode = &ListNode{
				Val: digit,
			}

			l1p = l1p.Next
		} else if l1p == nil {
			value += l2p.Val
			digit := value % 10
			newNode = &ListNode{
				Val: digit,
			}

			l2p = l2p.Next
		}

		if value >= 10 {
			incrementFromLastStage = 1
			incrSetThisStage = true
		}

		if lastNode != nil {
			(*lastNode).Next = newNode
		}
		if root == nil {
			root = newNode
		}
		lastNode = newNode

		if !incrSetThisStage {
			incrementFromLastStage = 0
		}
	}

	if incrementFromLastStage != 0 {
		newNode := &ListNode{
			Val: incrementFromLastStage,
		}
		lastNode.Next = newNode
	}

	return root
}
