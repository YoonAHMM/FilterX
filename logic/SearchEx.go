package logic

import "github.com/FilterX/internal"

type SearchEx struct {
	BaseSearchEx
}

func NewWordsSearchEx() *SearchEx{
	return &SearchEx{
	}
}

func(s *SearchEx)SetKeyWords(keywords []string) {
	s.KeyWords = keywords

	//创建字典
	length:=s.CreateDict(keywords)
	root:=internal.NewTrieNodeEx()
	for i,keyword:= range keywords {
		nd:=root
		p:= []rune(keyword)
		for _,c := range p{
			nd = nd.Add(int32(s.Dict[c]))
		}
		nd.SetResults(i)
	}

	
	nodes:=make([]*internal.TrieNodeEx,0)

	//设置二层节点的失败节点，记录三层节点
	for _,value := range root.Values{
		value.Failure=root
		for _,trans := range value.Values{
			nodes=append(nodes,trans)
		}
	}

	//设置失败节点
	for len(nodes) > 0 {
		newNodes:=make([]*internal.TrieNodeEx,0)
		for _,nd :=range nodes {
			fail := nd.Parent.Failure
			c := nd.Char
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
			if fail==nil{
				nd.Failure = root
			}else{
				nd.Failure = fail.Values[c];
				for j:=range  nd.Failure.Results {
					nd.SetResults(j)
				}
			}

			for _,child:= range nd.Values{
				newNodes=append(newNodes,child)
			}
		}
		nodes = newNodes
	}
	root.Failure = root

	for _,item :=range root.Values{
		s.tryLinks(item)
	}
	s.build(root, length)
}

func(s *SearchEx)FindFirst(text string) string{
	p:=0

	for _,c:=range text{
		t:=s.Dict[c]
		if t == 0{
			p = 0
			continue
		}
		next:=s.Next[p]+t
		find:=s.Key[next] == t
		if !find && p != 0{
			p = 0
			next = s.Next[0]+t
			find = s.Key[next] == t
		}

		if find {
			index:=s.Check[next]
			if index > 0{
				return s.KeyWords[s.Guides[index][0]]
			}
			p = next
		}
	}
	return ""
}