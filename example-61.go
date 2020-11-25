package main


func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	q := head
	for q.Next == nil {
		return head
	}
	i := 1
	for q.Next != nil {
		q = q.Next
		i++
	}
	k = k % i
	var h, prev *ListNode
	for k > 0 {
		h = head
		if head.Next == nil {
			return head
		}
		for head.Next != nil {
			prev = head
			head = head.Next
		}
		prev.Next = nil
		head.Next = h
		k--
	}
	return head
}
