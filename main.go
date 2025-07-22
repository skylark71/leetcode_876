package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	values := []int{1, 2, 3, 4, 5}
	head := createList(values)

	fmt.Print("Исходный список: ")
	printList(head)

	mid := middleNode(head)

	fmt.Print("Середина и дальше: ")
	printList(mid)
}

func createList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	cur := head
	for _, val := range values[1:] {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}
	return head
}
func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" → ")
		}
		head = head.Next
	}
	fmt.Println()
}

func middleNode(head *ListNode) *ListNode {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}
