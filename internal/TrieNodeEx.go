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
	Next int   //子节点起始位置
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

func(t *TrieNodeEx)Rank(has []*TrieNodeEx)int{
	seats:=make([]bool,len(has))//位置
	start:=1

	has[0]=t //索引
	t.Rank2(start,seats,has)
	maxCount:=len(has)-1
	for has[maxCount]==nil{
		maxCount--
	}
	return maxCount
}

func(t *TrieNodeEx)Rank2(start int,seats []bool,has []*TrieNodeEx)int{
	// 没有子节点
	if t.Count==0 {
		return start
	}

	// 子节点
	keys := make([]int32,0,len(t.Merge_values)+len(t.Values))
	for k,_:=range t.Values {
		keys=append(keys,k)
	}
	for k,_:=range t.Merge_values {
		keys=append(keys,k)
	}	

	// 查找下一个起始位置
	for	has[start] != nil{
		start++
	}

	
	s := start
	if start < int (t.Min){
		s= int (t.Min)
	}

	for i:=s;i <= int(t.Max);i++{
		if has[i] == nil{
			// 计算位置
			next:=i-int(t.Min)
			if seats[next] {
				continue
			}

			// 检查位置是否可用
			ok:=true
			for _,item:=range keys{
				if has[next + int(item)] != nil {
					ok =false
					break
				}
			}
			// 可用
			if ok {
				t.Next = next
				seats[next] = true
				t.SetSeats(next,has)
				break
			}
		}
	}
	start +=len(keys)/2
	for _,v:=range t.Merge_values{
		start = v.Rank2(start,seats,has)
	}
	return start
}

func (t *TrieNodeEx)SetSeats(next int,has []*TrieNodeEx){
	for key,value := range t.Merge_values{
		position := next + int (key)
		has[position] = value
	}
	for key,value := range t.Values{
		position := next + int (key) 
		has[position] = value
	}
}
