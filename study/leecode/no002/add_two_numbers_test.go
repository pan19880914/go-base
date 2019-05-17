package no002

import (
	"fmt"
	"testing"
)

func BenchmarkCacheData(b *testing.B) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, &ListNode{3, nil}}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	b.ResetTimer()
	addTwoNumbers(l1, l2)
	b.StopTimer()
}
func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, &ListNode{4, &ListNode{3, &ListNode{3, nil}}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	actual := addTwoNumbers(l1, l2)
	fmt.Println(actual.Val)
	fmt.Println(actual.Next.Val)
	fmt.Println(actual.Next.Next.Val)
	fmt.Println(actual.Next.Next.Next.Val)
	//assert.Equal(t, expected, actual)
}
