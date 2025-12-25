package kojiro

import (
	"testing"
)

func TestNewListNode(t *testing.T) {
	result := NewListNode(1234)
	expect := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	if !IsSameListNode(result, expect) {
		t.Errorf("New list node does not produce the correct ListNode")
	}
}

func TestListNodeString(t *testing.T) {
	result := &ListNode{
		Val: 4,
		Next: &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val:  1,
					Next: nil,
				},
			},
		},
	}

	if result.String() != "1594" {
		t.Errorf("List node string function invalid")
	}
}

func TestAddTwoNumbersSameLength(t *testing.T) {
	result := addTwoNumbers(&ListNode{Val: 1}, &ListNode{Val: 2})
	if !IsSameListNode(result, &ListNode{Val: 3}) {
		t.Errorf("Add two numbers of same length incorrect")
	}
}

func TestAddTwoNumbersDifferentLength(t *testing.T) {
	result := addTwoNumbers(&ListNode{Val: 1, Next: &ListNode{Val: 5}}, &ListNode{Val: 2})
	expect := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 5,
		},
	}
	if !IsSameListNode(result, expect) {
		t.Errorf("Add two numbers of different length incorrect")
	}
}

func TestAddTwoNumbersIncrementLength(t *testing.T) {
	x := &ListNode{Val: 9}
	y := &ListNode{Val: 7}
	expect := &ListNode{
		Val: 6,
		Next: &ListNode{
			Val: 1,
		},
	}
	result := addTwoNumbers(x, y)
	if !IsSameListNode(result, expect) {
		t.Errorf("Add two numbers to increment length failed")
	}
}

func TestAddTwoNumbersCase1(t *testing.T) {
	x := NewListNode(342)
	y := NewListNode(465)
	expect := NewListNode(807)
	result := addTwoNumbers(x, y)
	if !IsSameListNode(result, expect) {
		t.Errorf("342 + 465 should = 807 but equals %s", result)
	}
}

func TestAddTwoNumbersCase2(t *testing.T) {
	x := NewListNode(9999999)
	y := NewListNode(9999)
	expect := NewListNode(10009998)
	result := addTwoNumbers(x, y)
	if !IsSameListNode(result, expect) {
		t.Errorf("9999999 + 9999 should = 10009998 but equals %s", result)
	}
}
