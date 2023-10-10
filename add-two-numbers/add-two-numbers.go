package add_two_numbers

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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	toAdd := 0
	finalL := ListNode{}

	currentNode := &finalL
	prevNode := currentNode
	for l1 != nil || l2 != nil || toAdd != 0 {
		l1V, l2V := 0, 0
		if l1 != nil {
			l1V = l1.Val

			l1 = l1.Next
		} else {
			l1 = nil
		}

		if l2 != nil {
			l2V = l2.Val

			l2 = l2.Next
		} else {
			l2 = nil
		}

		println(l1V, l2V, toAdd)
		currentV := l1V + l2V + toAdd
		if currentV >= 10 {
			currentV -= 10
			toAdd = 1
		} else {
			toAdd = 0
		}

		currentNode.Val = currentV

		newNode := &ListNode{}
		currentNode.Next = newNode

		prevNode = currentNode
		currentNode = newNode

	}

	prevNode.Next = nil

	return &finalL
}

func Run() {
	l1 := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}

	l2 := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}

	println(1111)
	l := addTwoNumbers(l1, l2)

	println(2222)
	println(l.Val)
	for l.Next != nil {
		l = l.Next
		println(l.Val)
	}

}
