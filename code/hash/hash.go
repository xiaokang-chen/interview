package hash

// HashNode hash表节点
type HashNode struct {
	Key   string
	Value string
	Next  *HashNode
}

// HashList redis 哈希表
type HashList struct {
	Node []*HashNode
}
