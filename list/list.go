package list

import (
	"github.com/brunoga/rcu"
)

type Node struct {
	previous rcu.Data[Node]
	next     rcu.Data[Node]
	value    interface{}
}

type List struct {
	head rcu.Data[Node]
	tail rcu.Data[Node]
}

func (l *List) Insert(node *Node, value interface{}) *Node {
	newNode := &Node{
		value: value,
	}

	if node == nil {
		newNode.previous = rcu.NewData((*Node)(nil))
		newNode.next = l.head
		l.head.SetValue(newNode)
	} else {
		newNode.previous = rcu.NewData(node)
		newNode.next = node.next
		node.next.SetValue(newNode)
	}

	nextNode := newNode.next.GetValue()
	if nextNode == nil {
		l.tail.SetValue(newNode)
	} else {
		nextNode.previous.SetValue(newNode)
	}

	return newNode
}

func (l *List) Remove(node *Node) {
	if node == nil {
		return
	}

	previousNode := node.previous.GetValue()
	nextNode := node.next.GetValue()

	if previousNode == nil {
		l.head.SetValue(nextNode)
	} else {
		previousNode.next.SetValue(nextNode)
	}

	if nextNode == nil {
		l.tail.SetValue(previousNode)
	} else {
		nextNode.previous.SetValue(previousNode)
	}
}

func (l *List) Head() *Node {
	return l.head.GetValue()
}

func (l *List) Tail() *Node {
	return l.tail.GetValue()
}

func (n *Node) Next() *Node {
	return n.next.GetValue()
}

func (n *Node) Previous() *Node {
	return n.previous.GetValue()
}
