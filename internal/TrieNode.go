package internal

type TrieNode struct {
	Id    int   //id
	Level    int   //深度
	End      bool  //是否结束
	C     int32   //字符
	Results  []int //结果
	Values map[int32]*TrieNode  //子节点
	Failure  *TrieNode    //失败节点
	Parent   *TrieNode    //父节点
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		End:      false,
		Values: make(map[int32]*TrieNode),
		Results:  make([]int, 0),
	}
}

func (t *TrieNode) Add(c int32) *TrieNode {
	if val, ok := t.Values[c]; ok {
		return val
	}
	node := NewTrieNode()
	node.Parent = t
	node.C = c
	t.Values[c] = node
	return node
}

func (t *TrieNode) SetResults(text int) {
	if !t.End {
		t.End = true
	}
	for i :=range t.Results  {
        if i == text {
            return
        }
    }
	t.Results = append(t.Results, text)
}
