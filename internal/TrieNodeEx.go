package internal
import(
	"math"
)
type TrieNodeEx struct {
	Parent *TrieNodeEx    //父节点
	Failure *TrieNodeEx   //失败跳跃节点
	Char int32            //字符
	End bool              //是否结束
	Results []int         //结果
	Values map[int32]*TrieNodeEx   //子节点  
	Merge_values map[int32]*TrieNodeEx  //合并子节点
	Min int32 //子节点最小范围
	Max int32  //子节点最大范围
	Next int  
	Count int  //子节点数量
}

func NewTrieNodeEx() *TrieNodeEx  {
	return &TrieNodeEx{
		Values: make(map[int32]*TrieNodeEx),
		Merge_values: make(map[int32]*TrieNodeEx),
		Results: make([]int, 0),
		Min:  math.MaxInt32,
		Max:  math.MinInt32,
		Next:0,
		Count:0,
	}
}

func (t *TrieNodeEx) GetValue(c int32) (*TrieNodeEx,bool) {
	if t.Min <= c && t.Max >= c {
		if val, ok := t.Values[c]; ok {
			return val,true
		}
	}
	return nil,false
}

func (t *TrieNodeEx) Add(c int32)*TrieNodeEx {
	if val,ok:=t.GetValue(c);ok{
		return val
	}
	// 子节点范围
	if t.Min > c {
		t.Min = c
	}
	if t.Max < c {
		t.Max = c
	}
	node := NewTrieNodeEx()
	node.Parent = t
	node.Char = c
	t.Values[c]=node
	t.Count++
	return node 
}
func (t *TrieNodeEx) SetResults(text int) {
	if !t.End {
		t.End=true;
	}
	for i :=range t.Results  {
        if i == text {
            return
        }
    }
	t.Results=append(t.Results,text)
}

//合并node的节点以及所有失败子节点
func (t *TrieNodeEx)Merge(node *TrieNodeEx) {
	nd:=node
	for	nd.Char != 0 {
		for k,v:= range node.Values{
			if _,ok := t.Values[k]; ok{
				continue
			}
			if _,ok := t.Merge_values[k]; ok{
				continue
			} 
			if t.Min > k {
				t.Min = k
			}
			if t.Max < k {
				t.Max = k
			}
			t.Merge_values[k]=v
			t.Count++
		}
		nd = nd.Failure
	}
}

