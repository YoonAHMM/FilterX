package logic

import (
	"math"
	"os"

	"github.com/FilterX/internal"
	. "github.com/FilterX/pkg"
)

type BaseSearchEx struct {
	KeyWords 	[]string  //模式串
	Guides 		[][]int	  //模式串索引
	Key 		[]int     //当前位置字符	
	Next  		[]int	  //查询子节点的指针位置
	Check  		[]int	  //当前位置是否命中
	Dict 		[]int	  //字典
}

//检测是否有保存的初始化二进制文件
func(b *BaseSearchEx)checkFileIsExist(filepath string) bool{
	if _,err :=os.Stat(filepath);os.IsNotExist(err){
		return false
	}
	return true
}

// 保存到二进制文件
func(b *BaseSearchEx)Save(filepath string) {
	//如果文件不存在则创建该文件,以只写模式打开文件
	f,_:=os.OpenFile(filepath,os.O_CREATE|os.O_WRONLY,0666)
	defer f.Close()

	b.Save2(f)
}

func(b *BaseSearchEx)Save2(f *os.File) {
	
	//KeyWords
	f.Write(IntToBytes(len(b.KeyWords)))
	for _,key:=range b.KeyWords {
		f.Write(IntToBytes(len(key)))
		f.Write([]byte (key))
	}

	//Guides
	f.Write(IntToBytes(len(b.Guides)))
	for _,guide:=range b.Guides {
		f.Write(IntToBytes(len(guide)))
		for _,g:=range guide {
			f.Write(IntToBytes(g))
		}
	}

	//Key
	f.Write(IntToBytes(len(b.Key)))
	for _,key:=range b.Key {
		f.Write(IntToBytes (key))
	}

	//Next
	f.Write(IntToBytes(len(b.Next)))
	for _,next:=range b.Next {
		f.Write(IntToBytes(next))
	}

	//Check
	f.Write(IntToBytes(len(b.Check)))
	for _,check:= range b.Check{
		f.Write(IntToBytes(check))
	}

	//Dict
	f.Write(IntToBytes(len(b.Dict)))
	for _,item:= range b.Dict{
		f.Write(IntToBytes(item))
	}

}

// 从二进制文件读取
func(b *BaseSearchEx)Load(filepath string) {
	//如果文件不存在则创建该文件,以只写模式打开文件
	f,_:=os.OpenFile(filepath,os.O_RDONLY,0666)
	defer f.Close()

	b.Load2(f)
}

func (b *BaseSearchEx)Load2(f *os.File) {
	//KeyWords
	length:=BytesToInt(ReadBytes(f,4))
	b.KeyWords = make([]string,length)
	for i:=0;i<length;i++ {
		len:=BytesToInt(ReadBytes(f,4))
		b.KeyWords[i] = string(ReadBytes(f,len))
	}

	//Guides
	length=BytesToInt(ReadBytes(f,4))
	b.Guides = make([][]int,length)
	for i:=0;i<length;i++ {
		len:=BytesToInt(ReadBytes(f,4))
		b.Guides[i] = make([]int,length)
		for j:=0;j<len;j++ {
			b.Guides[i][j] = BytesToInt(ReadBytes(f,4))
		}
	}

	//Key
	length=BytesToInt(ReadBytes(f,4))
	b.Key = make([]int,length)
	for i:=0;i<length;i++ {
		b.Key[i] = BytesToInt(ReadBytes(f,4))
	}

	//Next
	length=BytesToInt(ReadBytes(f,4))
	b.Next = make([]int,length)
	for i:=0;i<length;i++ {
		b.Next[i] = BytesToInt(ReadBytes(f,4))
	}

	//Check
	length=BytesToInt(ReadBytes(f,4))
	b.Check = make([]int,length)
	for i:=0;i<length;i++ {
		b.Check[i] = BytesToInt(ReadBytes(f,4))
	}

	//Dict
	length=BytesToInt(ReadBytes(f,4))
	b.Dict = make([]int,length)
	for i:=0;i<length;i++ {
		b.Dict[i] = BytesToInt(ReadBytes(f,4))
	}

}





func (b *BaseSearchEx)CreateDict(keywords []string) int {
	dictionary:= make(map[int32]int, 0)

	//根据字符出现次数加权
	for	_,keyword := range keywords{
		for _,item:=range keyword{
			if v,ok:= dictionary[item];ok{
				if v > 0 {
					dictionary[item]=dictionary[item]+1
				}
			}else{
				dictionary[item]=1
			}

		}
	}

	//根据加权排序所有字符
	list:=SortMap(dictionary)

	//index:864201357(0左边是偶数顺序，0右边是奇数顺序)
	index:=make([]int,0,len(list))	

	
	for i:=0;i<len(list);i=i+2 {
		index= append(index, i)
	}

	length:= len(index)

	
	for i := 0 ; i < length/2 ; i++ {
		index[i], index[length -i - 1] = index[length - i -1 ], index[i]
	}

	
	for i:=1;i<len(list);i=i+2 {
		index = append(index, i)
	}
	

	//根据index重排list到list2
	list2:=make([]int32,0)
	for i:=0;i<len(list);i++ {
		list2 = append(list2,list[index[i]])
	}

	b.Dict = make([]int,math.MaxInt32)   
	
	//以list2顺序添加到字典
	for i,v:=range list2{
		b.Dict[v] = i + 1 //为0的位置说明字符不存在

	}
	return len(dictionary)  
}

func(b *BaseSearchEx)tryLinks(node *internal.TrieNodeEx) {
	node.Merge(node.Failure)
	for _,item:=range node.Merge_values{
		b.tryLinks(item)
	}
}

func(b *BaseSearchEx)build(root *internal.TrieNodeEx,length int) {
	has:= make([]*internal.TrieNodeEx,0x00FFFFFF)

	length = root.Rank(has) + length + 1;
	b.Key=make([]int,length)
	b.Next=make([]int,length)
	b.Check=make([]int,length)
	var guides [][]int

	// 根节点
	guides=append(guides,[]int{0})

	for i := 0; i < length; i++  {
		item := has[i];
		if item==nil{
			continue
		}
		b.Key[i] = int (item.Char) ;
		b.Next[i] = item.Next;
		if  item.End==true  {
			b.Check[i] = len(guides) 
			guides=append(guides,item.Results)
		}
	}
	b.Guides=guides
	return
}