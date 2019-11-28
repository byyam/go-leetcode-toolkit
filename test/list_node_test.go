package test

import (
	toolkit "github.com/byyam/go-leetcode-toolkit"
	"testing"
)

func Test_ListNode_String(t *testing.T) {
	l1 := &toolkit.ListNode{Val: -1}
	l1.AppendList([]int{5, 1, 3})
	l1.Next.String()
}
