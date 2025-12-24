package kojiro

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbersSameLength(t *testing.T) {
	result := addTwoNumbers(&ListNode{Val: 1}, &ListNode{Val: 2})
	fmt.Println(result)
}
