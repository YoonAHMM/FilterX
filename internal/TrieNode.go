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

func (this *TrieNode) Add(c int) *TrieNode {
	if val, s := this.Values[c]; s {
		return val
	}
	node := NewTrieNode()
	node.Parent = this
	node.C = c
	this.Values[c] = node
	return node
}

func (this *TrieNode) SetResults(text int) {
	if this.End == false {
		this.End = true
	}
	for i := 0; i < len(this.Results); i++ {
		if this.Results[i] == text {
			return
		}
	}
	this.Results = append(this.Results, text)
}
