package lrucache

type LRUCache struct {
	head     *Node
	tail     *Node
	hashmap  map[int]*Node
	capacity int
}

type Node struct {
	Prev, Next *Node
	Key, Value int
}

func Constructor(capacity int) LRUCache {
	head, tail := &Node{}, &Node{}

	head.Next = tail
	tail.Prev = head
	return LRUCache{
		head:     head,
		tail:     tail,
		hashmap:  make(map[int]*Node),
		capacity: capacity,
	}
}

func (cache *LRUCache) Get(key int) int {
	if n, ok := cache.hashmap[key]; ok {
		cache.remove(n)
		cache.insert(n)
		return n.Value
	}

	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	if _, ok := cache.hashmap[key]; ok {
		cache.remove(cache.hashmap[key])
	}

	if len(cache.hashmap) == cache.capacity {
		cache.remove(cache.tail.Prev)
	}

	cache.insert(&Node{Key: key, Value: value})
}

func (cache *LRUCache) remove(node *Node) {
	delete(cache.hashmap, node.Key)

	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (cache *LRUCache) insert(node *Node) {
	cache.hashmap[node.Key] = node

	next := cache.head.Next
	cache.head.Next = node
	node.Prev = cache.head
	next.Prev = node
	node.Next = next
}
