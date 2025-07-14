package potato_bugs

type PotatoNode struct {
	Name string
	Next *PotatoNode
}

func hasPotatoCycle(head *PotatoNode) bool {
	if head == nil {
		return false
	}
	slowNode := head
	fastNode := head
	for fastNode != nil && fastNode.Next != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next.Next
		if slowNode == fastNode {
			return true
		}
	}
	return false
}
