package main

import "fmt"

type LinkNode struct {
	data     int
	nextNode *LinkNode
}

func (l *LinkNode) add(data int) {
	node := l

	for node.nextNode != nil {
		node = node.nextNode
	}

	node.nextNode = &LinkNode{
		data: data,
	}
}

func (l *LinkNode) Print() {
	node := l

	for node != nil {
		fmt.Println(node.data)
		node = node.nextNode
	}
}

func InitLink() *LinkNode {
	node := new(LinkNode)

	node.add(1)
	node.add(2)
	node.add(3)

	return node
}

func main() {
	link := InitLink()
	link.Print()
}
