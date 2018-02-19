package dstructs

import (
	"fmt"
	"strings"
)

type node struct {
	prev *node
	val int
	key int
	next *node
}

func (n *node) String() string {
	str := "(Prev : "
	if n.prev != nil {
		str += fmt.Sprintf("%d;", n.prev.key)
	} else {
		str += "NIL; "
	}
	str += fmt.Sprintf(" K, V : %d, %d; Next : ", n.key, n.val)
	if n.next != nil {
		str += fmt.Sprintf("%d)", n.next.key)
	} else {
		str += "NIL)"
	}
	return str
}

func (n *node) GetPrev() *node {
	return n.prev
}

func (n *node) GetNext() *node {
	return n.next
}

type dList struct {
	head *node
	tail *node
}

func NewDList() *dList {
	return new(dList)
}

func (l *dList) String() string {
	n := l.head
	soln := make([]int, 0)
	for n != nil {
		soln = append(soln, n.key)
		n = n.next
	}
	return strings.Join(strings.Split(fmt.Sprint(soln), " "), ", ")
}

func (root *dList) GetHead() (*node) {
	return root.head
}

func (root *dList) GetTail() (*node) {
	return root.tail
}

func (root *dList) Prepend(key, value int) {
	newNode := new(node)
	newNode.val = value
	newNode.key = key
	if root.head == nil {
		root.head = newNode
		root.tail = newNode
	} else {
		newNode.next = root.head
		root.head.prev = newNode
		root.head = newNode
	}
}

func (root *dList) EvictTail() {
	if root.tail != nil {
		if root.tail.prev == nil {
			root.head = nil
			root.tail = nil
		} else {
			root.tail.prev.next = nil
			root.tail = root.tail.prev
		}
	}
}

func (root *dList) MoveToHead(n *node) {
	if n.prev != nil {
		if n.next == nil {
			n.prev.next = nil
			root.tail = n.prev
		} else {
			n.prev.next = n.next
			n.next.prev = n.prev
		}
		n.prev = nil
		root.head.prev = n
		n.next = root.head
		root.head = n
	}
}