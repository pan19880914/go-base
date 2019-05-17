package no002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	return nil
}

func reverseList(list *ListNode) (newlist *ListNode) {
	head := &ListNode{}
	head.Next = list
	cursor := list
	link := cursor.Next
	for link != nil {
		cursor.Next = link.Next
		link.Next = head.Next
		head.Next = link
		link = cursor.Next
	}
	return head.Next
}
