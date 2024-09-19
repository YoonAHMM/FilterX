package logic

import (
	"unicode/utf8"

	"github.com/FilterX/format"
	"github.com/FilterX/internal"
)

//算法：AC自动机+Trie tree

type Search struct {
	first    map[int32]*internal.TrieNode2   //所有root下第一层节点
	Keywords []string      //模式串
}

func NewSearch()*Search{
	return &Search{
		first: make(map[int32]*internal.TrieNode2),
	}
}

func(s *Search)SetKeywords(keywords []string){
	s.Keywords = keywords

//Trie tree 构建

	//root节点
	root:=internal.NewTrieNode()
	//每一层节点的深度级别
	nodeLevel:=make(map[int]*internal.TrieNodes)

	for index,str:=range keywords{
		node := root
		//字符位置
		r:=[]rune(str)
		for i,ch:=range r{
			node = node.Add(ch)

			//新节点
			if node.Level == 0{
				node.Level = i + 1
				if trieNodes, ok := nodeLevel[node.Level]; ok {
					trieNodes.Items = append(trieNodes.Items, node)
				} else {
					trieNodes = internal.NewTrieNodes()
					nodeLevel[node.Level] = trieNodes
					trieNodes.Items = append(trieNodes.Items, node)
				}
			}
		}
		node.SetResults(index)
	}

//ac自动机 构建	
	allNode := make([]*internal.TrieNode, 1)  //所有节点
	allNode[0] = root
	
	//根据层高把节点写入数组
	for i := 1; i <=len(nodeLevel); i++ {
		nodes:= nodeLevel[i].Items
		for j := 0; j < len(nodes); j++ {
			allNode = append(allNode, nodes[j])
		}
	}

	//根据所有节点构建ac自动机
	//设置失败节点，从上向下一层层设(普普通通的ac自动机hh)
	for i := 1; i < len(allNode); i++ {
		node  := allNode[i]
		node.Id = i
		c:= node.C

		fail := node.Parent.Failure

		//找到失败节点,从失败节点的子节点开始找，找不到就继续找失败节点的失败节点，直到找到或者失败节点为nil
		for {
			if fail != nil {
				if _, ok := fail.Values[c]; ok {
					break
				} else {
					fail = fail.Failure
				}
			} else {
				break
			}
		}

		//如果失败节点为nil，就设置为root节点
		if fail == nil {
			node.Failure = root
		} else {
			//将失败节点的数据也存到本节点(那这个节点就会有当前模式串的结果和模式串字串的所有结果)
			node.Failure = fail.Values[c]
			for j:=range  node.Failure.Results {
				node.SetResults(j)
			}
		}	
	}
	root.Failure=root

//Trie2 (ac+Trie) 构建
	allNode2 := make([]*internal.TrieNode2, len(allNode))

	for i := 0; i < len(allNode); i++ {
		allNode2[i] = internal.NewTrieNode2()
	}

	for i := 0; i < len(allNode); i++ {
		oldNode := allNode[i]
		newNode := allNode2[i]

		//子节点正常加
		for k, v := range oldNode.Values {
			newNode.Add(k, allNode2[v.Id])
		}
		//结果正常加
		for j := 0; j < len(oldNode.Results); j++ {
			newNode.SetResults(oldNode.Results[j])
		}

		oldNode = oldNode.Failure
		//递归走到root
		for oldNode != root {
			//把失败节点的路(以及失败节点的失败节点。。。)(本节点没有的)
			for c, v := range oldNode.Values {
				newNode.Add(c, allNode2[v.Id])
			}
			//失败节点结果加
			for j:=range oldNode.Failure.Results {
				newNode.SetResults(j)
			}
			oldNode = oldNode.Failure
		}
	}

	//父节点为root放入first内	
	for key, val := range allNode2[0].Values {
		s.first[key] = val
	}
	
}

//findfirst 找到的是第一个匹配串
func (s *Search) GetWordsFindFirst(text string) *format.WordsSearchResult {
	var lnode *internal.TrieNode2   //前一个trieNode2值

	i:=0
	//遍历所有字符，如果当前节点的字符表里没有则到first里找
	for _, t := range text {
		var node *internal.TrieNode2  //当前trieNode2值
		
		if lnode == nil {
			node = s.first[t]
		} else {
			var ok bool
			if node, ok = lnode.GetValue(t);!ok {
				node = s.first[t]
			}
		}

		if node != nil {
			if node.End {
				k := s.Keywords[node.Results[0]]
				length :=  utf8.RuneCountInString(k)
				return format.NewWordsSearchResult(k, i+1-length+1, i+1, node.Results[0])
			}
		}
		lnode = node
		i++
	}
	return nil
}

func (s *Search) GetStringFindFirst(text string) string{
	var lnode *internal.TrieNode2   //前一个trieNode2值

	
	//遍历所有字符，如果当前节点的字符表里没有则到first里找
	for _, t := range text {
		var node *internal.TrieNode2  //当前trieNode2值
		
		if lnode == nil {
			node = s.first[t]
		} else {
			var ok bool
			if node, ok = lnode.GetValue(t);!ok {
				node = s.first[t]
			}
		}

		if node != nil && node.End{
			return s.Keywords[node.Results[0]]
		}
		lnode = node
	}
	return ""
}


func (s *Search)  GetWordsFindAll(text string) []*format.WordsSearchResult {
	list := make([]*format.WordsSearchResult, 0)

	var lnode *internal.TrieNode2   //前一个trieNode2值

	i:=0
	//遍历所有字符，如果当前节点的字符表里没有则到first里找
	for _, t := range text {
		var node *internal.TrieNode2  //当前trieNode2值
		
		if lnode == nil {
			node = s.first[t]
		} else {
			var ok bool
			if node, ok = lnode.GetValue(t);!ok {
				node = s.first[t]
			}
		}

			if node != nil {
				if node.End {
					for _,index:=range node.Results {
						k := s.Keywords[index]
						length := utf8.RuneCountInString(k)
						r := format.NewWordsSearchResult(k, i+1-length+1, i+1, index)
						list = append(list, r)
					}
			}
			lnode = node
			i++
		}
	}
	return list
}


func (s *Search)  GetStringFindAll(text string) []string {
	list := make([]string, 0)

	var lnode *internal.TrieNode2   //前一个trieNode2值

	i:=0
	//遍历所有字符，如果当前节点的字符表里没有则到first里找
	for _, t := range text {
		var node *internal.TrieNode2  //当前trieNode2值
		
		if lnode == nil {
			node = s.first[t]
		} else {
			var ok bool
			if node, ok = lnode.GetValue(t);!ok  {
				node = s.first[t]
			}
		}

			if node != nil && node.End{
				for _,index:=range node.Results {
					k := s.Keywords[index]
					list = append(list, k)
				}
			}
			lnode = node
			i++
	}
	return list
}





//获取text是否包含违禁词
func (s *Search) ContainsAny(text string) bool {
	var lnode *internal.TrieNode2   //前一个trieNode2值


	//遍历所有字符，如果当前节点的字符表里没有则到first里找
	for _, t := range text {
		var node *internal.TrieNode2  //当前trieNode2值
		
		if lnode == nil {
			node = s.first[t]
		} else {
			var ok bool
			if node, ok = lnode.GetValue(t);!ok  {
				node = s.first[t]
			}
		}

		if node != nil {
			if node.End {
				return true
			}
		}
		lnode = node
	}
	return false
}


func (s *Search) Replace(text string,replaceChar int32) string {
	result := []rune(text)

	var lnode *internal.TrieNode2   //前一个trieNode2值

	i:=0
	//遍历所有字符，如果当前节点的字符表里没有则到first里找
	for _, t := range text {
		var node *internal.TrieNode2  //当前trieNode2值
		
		if lnode == nil {
			node = s.first[t]
		} else {
			var ok bool
			if node, ok = lnode.GetValue(t);!ok  {
				node = s.first[t]
			}
		}

		if node != nil && node.End{
			length := len([]rune(s.Keywords[node.Results[0]]))
			start := i + 1 - length
			for j := start; j <= i; j++ {
				result[j] = replaceChar
			}
		}
		lnode = node
		i++
	}
	return string(result)

}

