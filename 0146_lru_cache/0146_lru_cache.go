package leetcode

import "runtime/debug"

type Node struct {
	Key  int
	Val  int
	Next *Node
	Prev *Node
}

type LRUCache struct {
	capacity int
	m        map[int]*Node
	head     *Node
	tail     *Node
	size     int
}

// i dont know why should do this, but this can make runtime faster
func init() { debug.SetGCPercent(-1) }

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		m:        make(map[int]*Node),
		head:     nil,
		tail:     nil,
		size:     0,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.m[key]; ok {
		this.removeNode(node)
		this.pushFront(node.Key, node.Val)
		return node.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.m[key]; ok {
		this.removeNode(node)
		this.pushFront(key, value)
		return
	}

	if this.size >= this.capacity {
		this.removeNode(this.tail)
	}
	this.pushFront(key, value)
}

func (this *LRUCache) removeNode(node *Node) {
	if this.size == 0 {
		return
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	if this.head == node {
		this.head = node.Next
	}
	if this.tail == node {
		this.tail = node.Prev
	}
	node.Next = nil
	node.Prev = nil
	delete(this.m, node.Key)
	this.size--
}

func (this *LRUCache) pushFront(key int, value int) {
	newNode := &Node{
		Key: key,
		Val: value,
	}
	if this.size == 0 {
		this.head = newNode
		this.tail = newNode
	} else {
		newNode.Next = this.head
		this.head.Prev = newNode
		this.head = newNode
	}
	this.m[key] = newNode
	this.size++
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
