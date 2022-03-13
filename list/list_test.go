package list

import (
	"testing"
)

func countNodes(l *List) int {
	i := 0
	for node := l.head.GetValue(); node != nil; node = node.next.GetValue() {
		i++
	}
	return i
}

func check(t *testing.T, node, previous, next *Node, l *List, numNodes int) {
	if node.previous.GetValue() != previous {
		t.Errorf("node.previous != previous\n")
	}
	if node.next.GetValue() != next {
		t.Errorf("node.next != next\n")
	}
	if previous != nil && previous.next.GetValue() != node {
		t.Errorf("previous.next != node\n")
	}
	if next != nil && next.previous.GetValue() != node {
		t.Errorf("next.previous != node\n")
	}
	if countNodes(l) != numNodes {
		t.Errorf("countNodes(&l) != numNodes\n")
	}
}

func TestList(t *testing.T) {
	l := List{}

	node0 := l.Insert(nil, 0)
	check(t, node0, nil, nil, &l, 1)

	node1 := l.Insert(node0, 1)
	check(t, node1, node0, nil, &l, 2)

	node2 := l.Insert(node0, 2)
	check(t, node2, node0, node1, &l, 3)

	node3 := l.Insert(nil, 3)
	check(t, node3, nil, node0, &l, 4)
}
