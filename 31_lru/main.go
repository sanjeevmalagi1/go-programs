package main

import "fmt"

type Node struct {
	value string
	left  *Node
	right *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
}

func (l *DoubleLinkedList) insertAtBeginning(val string) *Node {

	newNode := &Node{value: val}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		newNode.left = l.head
		l.head.right = newNode
		l.head = newNode
	}

	return l.head
}

func (l *DoubleLinkedList) insertAtEnd(val string) {
	newNode := &Node{value: val}

	if l.tail == nil {
		l.tail = newNode
		l.head = newNode
	} else {
		newNode.right = l.tail
		l.tail.left = newNode
		l.tail = newNode
	}

	l.tail = newNode
}

func (dll *DoubleLinkedList) remove(node *Node) {
	if node == dll.tail {
		dll.tail = dll.tail.right
		dll.tail.left = nil
		return
	}

	left, right := node.left, node.right
	left.right = right
	right.left = left
}

func (l *DoubleLinkedList) display() {
	current := l.head
	for current != nil {
		fmt.Printf("%s -> ", current.value)
		current = current.left
	}
	fmt.Println("nil")
}

type HashMap struct {
	data map[string]*Node
}

func (h *HashMap) get(key string) *Node {
	return h.data[key]
}

func (h *HashMap) put(key string, value *Node) {
	if h.data == nil { // Ensure the map is initialized
		h.data = make(map[string]*Node)
	}
	h.data[key] = value
}

type LRUCache struct {
	hashmap HashMap
	ll      DoubleLinkedList
	maxSize int
}

func (c *LRUCache) get(key string) string {
	node := c.hashmap.get(key)

	if node == nil {
		panic("IDK")
	}

	c.ll.insertAtBeginning(node.value)
	c.ll.remove(node)

	return node.value
}

func (c *LRUCache) put(key, value string) {
	insertedNode := c.ll.insertAtBeginning(value)
	c.hashmap.put(key, insertedNode)

	if len(c.hashmap.data) > c.maxSize {
		c.ll.remove(c.ll.tail)
	}
}

func main() {
	cache := LRUCache{hashmap: HashMap{data: make(map[string]*Node)}, maxSize: 5}

	cache.put("1", "123")
	cache.put("2", "222")
	cache.put("3", "333")
	cache.put("4", "444")
	cache.put("5", "555")

	cache.ll.display()

	cache.get("2")
	cache.get("2")
	cache.ll.display()

}
