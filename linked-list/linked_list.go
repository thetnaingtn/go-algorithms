package linkedlist

type Node struct {
	data     string
	nextNode *Node
}

type LinkedList struct {
	firstNode *Node
}

func NewNode(d string) *Node {
	return &Node{data: d}
}

func NewLinkedList(n *Node) LinkedList {
	return LinkedList{firstNode: n}
}

func (l LinkedList) Read(index int) string {
	var (
		currentIndex = 0
		currentNode  = l.firstNode
	)

	for currentIndex < index {
		currentNode = currentNode.nextNode
		currentIndex += 1
	}

	if currentNode == nil {
		return ""
	}

	return currentNode.data
}

func (l LinkedList) Search(term string) int {

	var (
		currentIndex = 0
		currentNode  = l.firstNode
	)

	for currentNode.data != term {
		if currentNode.nextNode == nil {
			return -1
		}
		currentIndex += 1
		currentNode = currentNode.nextNode
	}

	return currentIndex
}

func (l LinkedList) Delete(index int) {
	var precedingNode *Node
	currentNode := l.firstNode
	currentIndex := 0
	if index == 0 {
		l.firstNode = currentNode.nextNode
		return
	}

	for currentIndex < index {
		if currentIndex == index-1 {
			precedingNode = currentNode
		}

		currentIndex++
		currentNode = currentNode.nextNode
	}

	precedingNode.nextNode = currentNode.nextNode
}
