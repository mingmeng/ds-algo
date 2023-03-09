package main

type LRUCache struct {
	Cache          map[int]*Node
	Capacity, Size int
	Head, Tail     *Node
}

func Constructor(capacity int) LRUCache {
	container := LRUCache{
		Cache:    map[int]*Node{},
		Capacity: capacity,
		Head:     NewNode(0, 0),
		Tail:     NewNode(0, 0),
	}

	container.Head.Next = container.Tail
	container.Tail.Pre = container.Head

	return container
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.Cache[key]; ok {
		c.moveToHead(node)
		return node.Value
	}
	return -1

}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.Cache[key]; ok {
		node.Value = value
		c.moveToHead(node)
	} else {
		newNode := NewNode(key, value)
		c.addToHead(newNode)
		c.Cache[key] = newNode
		c.Size++
		if c.Size > c.Capacity {
			deletedNode := c.removeTail()
			delete(c.Cache, deletedNode.Key)
			c.Size--
		}
	}
}

func (c *LRUCache) addToHead(node *Node) {
	node.Pre = c.Head
	node.Next = c.Head.Next
	c.Head.Next.Pre = node
	c.Head.Next = node
}

func (c *LRUCache) moveToHead(node *Node) {
	c.removeNode(node)
	c.addToHead(node)
}

func (c *LRUCache) removeNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (c *LRUCache) removeTail() *Node {
	node := c.Tail.Pre
	c.removeNode(node)
	return node
}

type Node struct {
	Key, Value int
	Pre, Next  *Node
}

func NewNode(k, v int) *Node {
	return &Node{Key: k, Value: v}
}

func main() {
	LRU := Constructor(1)
	LRU.Put(2, 1)
	LRU.Get(2)
}
