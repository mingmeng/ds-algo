package main

import "math"

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

func maxSubArray(nums []int) int {
	var max int

	max = nums[0]

	for i := 1; i < len(nums); i++ {
		nums[i] = max(nums[i], nums[i-1]+nums[i])

		if nums[i] > max {
			max = nums[i]
		}

	}

	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sortArray(nums []int) []int {
	var sort func(nums []int, left, right int) []int
	sort = func(nums []int, left, right int) []int {
		if left > right {
			return nil
		}
		pivot := nums[left]
		i, j := left, right
		for i < j {
			for i < j && nums[j] >= pivot {
				j--
			}
			for i < j && nums[i] <= pivot {
				i++
			}

			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[i], nums[left] = nums[left], nums[i]
		sort(nums, left, i-1)
		sort(nums, i+1, right)
		math.MaxInt64
		return nums
	}
	return sort(nums, 0, len(nums)-1)
}

func coinChange(coins []int, amount int) int {
	memo := make([]int, (amount)+1)
	for i := 0; i < len(memo); i++ {
		memo[i] = -100
	}

	var dp func(coins []int, amount int) int
	dp = func(coins []int, amount int) int {
		if amount == 0 {
			return 0
		} else if amount < 0 {
			return -1
		} else if memo[amount] != -100 {
			return memo[amount]
		}

		result := math.MaxInt
		for _, coin := range coins {

			subAmount := dp(coins, amount-coin)
			if subAmount == -1 {
				continue
			}

			result = min(result, subAmount+1)
		}

		if result == math.MaxInt {
			memo[amount] = -1
		} else {
			memo[amount] = result
		}
		return memo[amount]
	}

	return dp(coins, amount)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
