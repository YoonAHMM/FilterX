package internal

type TrieNode struct {
	Id    int   //id
	Depth    int   //深度
	End      bool  //是否结束
	C     int   //字符
	Results  []int //结果
	Values map[int]*TrieNode  //子节点
	Failure  *TrieNode    //失败节点
	Parent   *TrieNode    //父节点
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		End:      false,
		Values: make(map[int]*TrieNode),
		Results:  make([]int, 0),
	}
}

func (t *TrieNode) Add(c int) *TrieNode {
	if val, s := t.Values[c]; s {
		return val
	}
	node := NewTrieNode()
	node.Parent = t
	node.C = c
	t.Values[c] = node
	return node
}

func (t *TrieNode) SetResults(text int) {
	if t.End == false {
		t.End = true
	}
	for i := 0; i < len(t.Results); i++ {
		if t.Results[i] == text {
			return
		}
	}
	t.Results = append(t.Results, text)
}
