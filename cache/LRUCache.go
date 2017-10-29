package cache

/*
	A Golang implementation of the LRU cache, composed of a map and a doubly linked list
*/
type LRUCache struct {
	count int
	capacity int
	head *LRUNode
	tail *LRUNode
	dict map[int]*LRUNode
}


/*
	Return an instance of LRUCache, capacity as the parameter
*/
func GetLRUCache (c int) LRUCache {
	return LRUCache{0, c, nil, nil, make(map[int]*LRUNode)}
}


/*
	Add the node as the new head
*/
func (lru *LRUCache) addToHead (node *LRUNode) {
	if node == nil {
		return
	}
	if lru.head == nil {
		lru.head = node
		lru.tail = node
	} else {
		node.next = lru.head
		lru.head.prev = node
		node.prev = nil
		lru.head = node
	}
}


/*
	Node is passed through a reference got from the LRUCache; ensured existence.
*/
func (lru *LRUCache) deleteNode (node *LRUNode) {
	if node == nil {
		return
	}
	if node == lru.head {
		lru.head = node.next
		if node.next != nil {
			node.next.prev = nil
		}
	} else if node == lru.tail {
		lru.tail = node.prev
		if node.prev != nil {
			node.prev.next = nil
		}
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
}


/*
	Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
*/
func (lru *LRUCache) Get(key int) int {
	if lru.dict[key] != nil {
		node := lru.dict[key]
		result := node.value
		lru.deleteNode(node)
		lru.addToHead(node)
		return result
	}
	return -1
}


/*
	Set or insert the value if the key is not already present.
	When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.
*/
func (lru *LRUCache) Put(key int, value int)  {
	if lru.dict[key] != nil {
		node := lru.dict[key]
		node.value = value
		lru.deleteNode(node)
		lru.addToHead(node)
	} else {
		node := &LRUNode{key, value, nil, nil}
		lru.dict[key] = node

		if lru.count < lru.capacity {
			lru.addToHead(node)
			lru.count++
		} else {
			delete(lru.dict, lru.tail.key)
			lru.deleteNode(lru.tail)
			lru.addToHead(node)
		}
	}
}
