package internal

import "math"

//Ac自动机+字典树
type TrieNode2 struct {
	End      bool  //字符串结束标记
	Results  []int //结果
	Values  map[int32]*TrieNode2 //子节点+失败节点的子节点(非root)+失败的失败的子节点......
	Min  int32 //子节点最小值
	Max  int32 //子节点最大值
}

func NewTrieNode2() *TrieNode2 {
	return &TrieNode2{
		End:      false,
		Values: make(map[int32]*TrieNode2),
		Results:  make([]int, 0),
		Min:  math.MaxInt32,
		Max:  math.MinInt32,
	}
}

func (t *TrieNode2) Add(c int32, node *TrieNode2) {
	if _, ok := t.Values[c]; ok {
		return 
	}
	// 子节点范围
	if t.Min > c {
		t.Min = c
	}
	if t.Max < c {
		t.Max = c
	}
	t.Values[c] = node
}


func (t *TrieNode2) SetResults(text int) {
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

func (t *TrieNode2) GetValue(c int32) (*TrieNode2,bool) {
	if t.Min <= c && t.Max >= c {
		if val, ok := t.Values[c]; ok {
			return val,true
		}
	}
	return nil,false

}