package no002

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ll := &ListNode{0, nil}
	cur := ll
	c, s, t1, t2 := 0, 0, 0, 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			t1 = l1.Val
			l1 = l1.Next
		} else {
			t1 = 0
		}
		if l2 != nil {
			t2 = l2.Val
			l2 = l2.Next
		} else {
			t2 = 0
		}
		s = t1 + t2 + c
		c = s / 10
		cur.Next = &ListNode{Val: s % 10}
		cur = cur.Next
		if c >= 1 {
			cur.Next = &ListNode{Val: 1}
		}
	}
	return ll.Next
}

//
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	list := &ListNode{}
//	var link *ListNode
//	val, n := 0, 0
//	for l1 != nil || l2 != nil {
//		val, n = add(l1, l2, n)
//		l1,l2 = next(l1),next(l2)
//		node := &ListNode{val,link}
//		if link == nil {
//			list.Next,link = node,node
//			continue
//		}
//		link.Next = node
//		link = node
//	}
//	return list.Next
//}
//func next(l *ListNode) *ListNode {
//	if l != nil {
//		return l.Next
//	}
//	return nil
//}
//func add(n1, n2 *ListNode, i int) (v, n int) {
//	if n1 != nil {
//		v += n1.Val
//	}
//	if n2 != nil {
//		v += n2.Val
//	}
//	v += i
//	if v > 9 {
//		v %= 10
//		n = 1
//	}
//	return
//}
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
//	list := &ListNode{}
//	var indexNode *ListNode
//	temp1 := l1.Val
//	temp2 := l2.Val
//	high := 0
//	for i := 0 ;l1 != nil || l2 != nil;i++{
//		temp1,temp2 = 0,0
//		if l1 != nil {
//			temp1 ,l1= l1.Val,l1.Next
//		}
//		if l2 != nil {
//			temp2,l2 = l2.Val,l2.Next
//		}
//		val := temp1 + temp2 + high
//		if val >= 10 {
//			high,val = 1,val%10
//		}else{
//			high = 0
//		}
//		node := &ListNode{val,indexNode}
//		if i == 0 {
//			list.Next,indexNode = node,node
//			continue
//		}
//		indexNode.Next,indexNode = node,node
//	}
//	return list.Next
//}

//func reverseList(list *ListNode) (newlist *ListNode) {
//	head := &ListNode{}
//	head.Next = list
//	cursor := list
//	link := cursor.Next
//	for link != nil {
//		cursor.Next = link.Next
//		link.Next = head.Next
//		head.Next = link
//		link = cursor.Next
//	}
//	return head.Next
//}
