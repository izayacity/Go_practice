package cache

/*
	Nodes in the doubly linked list of the LRU Cache
*/
type LRUNode struct {
	key int
	value int
	prev *LRUNode
	next *LRUNode
}
