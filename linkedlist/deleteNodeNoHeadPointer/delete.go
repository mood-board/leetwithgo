package deleteNodeNoHeadPointer

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	// node.Val = node.Next.Val
	// node.Next = node.Next.Next

	nextNode := node.Next
	nextNodeVal := node.Next.Val
	followingNode := nextNode.Next

	node.Val = nextNodeVal
	node.Next = followingNode
}
