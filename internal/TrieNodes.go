package internal

type TrieNodes struct {
	Items []*TrieNode   //字典树每一层所有节点
}

func NewTrieNodes() *TrieNodes {
	return &TrieNodes{
		Items: make([]*TrieNode, 0),
	}
}
