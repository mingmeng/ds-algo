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

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Cache[key]; ok {
		this.moveToHead(node)
		return node.Value
	}
	return -1

}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.Cache[key]; ok {
		node.Value = value
		this.moveToHead(node)
	} else {
		newNode := NewNode(key, value)
		this.addToHead(newNode)
		this.Cache[key] = newNode
		this.Size++
		if this.Size > this.Capacity {
			deletedNode := this.removeTail()
			delete(this.Cache, deletedNode.Key)
			this.Size--
		}
	}
}

func (this *LRUCache) addToHead(node *Node) {
	node.Pre = this.Head
	node.Next = this.Head.Next
	this.Head.Next.Pre = node
	this.Head.Next = node
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *Node) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) removeTail() *Node {
	node := this.Tail.Pre
	this.removeNode(node)
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
