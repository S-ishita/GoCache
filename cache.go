package main

import "fmt"

type Cache struct {
	Queue Queue
	Hash  Hash
}

func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("add: %s\n", n.Val)
	tmp := c.Queue.Head.Right

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n

	c.Queue.Length++
	if c.Queue.Length > SIZE {
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("remove: %s\n", n.Val)
	left := n.Left
	right := n.Right

	right.Left = left
	left.Right = right

	c.Queue.Length -= 1
	delete(c.Hash, n.Val)

	return n
}

func (c *Cache) Display() {
	c.Queue.Display()
}

// if element exists remove it and then adds to front else creates new and adds to the front
func (c *Cache) Check(str string) {
	node := &Node{}

	if val, ok := c.Hash[str]; ok {
		node = c.Remove(val)
	} else {
		node = &Node{Val: str}
	}
	c.Add(node)
	c.Hash[str] = node

}
