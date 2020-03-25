package lru

import (
	"sync"
	"time"
)

// List 定义链表
type List struct {
	head   *Node
	tail   *Node
	length int
	sync.Mutex
}

// Node 链表节点
type Node struct {
	prev     *Node
	next     *Node
	expire   time.Duration
	expireAt time.Time
	val      string
}

// AddFirst 链表添加元素
func (list *List) AddFirst(node *Node) {
	list.Lock()
	defer list.Unlock()
	list.length++
	if list.head == nil {
		list.head = node
		list.tail = node
		return
	}
	list.head.prev = node
	node.next = list.head
	list.head = node
}

// RemoveLast 删除最后一个节点
func (list *List) RemoveLast() {
	if list.tail != nil {
		list.Remove(list.tail)
	}
}

// Remove 链表节点删除
func (list *List) Remove(node *Node) {
	list.Lock()
	defer list.Unlock()

	// emtpy list
	if list.head == nil {
		return
	}

	list.length--

	// 表头
	if list.head == node {
		// 唯一元素
		if list.head == list.tail {
			list.head = nil
			list.tail = nil
			return
		}
		list.head = node.next
		node.next.prev = nil
		return
	}

	// 表尾
	if list.tail == node {
		list.tail.prev.next = nil
		list.tail = list.tail.prev
		return
	}

	// 中间元素
	node.prev.next = node.next
	node.next.prev = node.prev
	return
}

// Clear 清空链表
func (list *List) Clear() {
	list.head = nil
	list.tail = nil
	list.length = 0
}
