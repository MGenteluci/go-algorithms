package main

type ListNode struct {
	Val int
	Next *ListNode
}

var l, r *ListNode

func removeNth(head *ListNode, n int) {
	if r.Next == nil {
		l.Next = l.Next.Next
		return
	}
	l = l.Next
	r = r.Next
	removeNth(head.Next, n)
}

// 19. Remove Nth Node From End of List
// Time complexity: 
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: 0, Next: head}
	l = dummy
	r = dummy
	for i := 0; i < n; i++ {
		r = r.Next
	}

	removeNth(dummy, n)
    return dummy.Next
}

func main() {}
