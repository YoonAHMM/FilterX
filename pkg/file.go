package pkg

import (
    "encoding/binary"
	"os"
	"log"
)

func IntToBytes(i int) []byte {
    bs := make([]byte, 4) // int32 占 4 个字节
    binary.BigEndian.PutUint32(bs, uint32(i)) // 转换为大端序的字节数组
    return bs
}

func BytesToInt(bs []byte) int {
    return int(binary.BigEndian.Uint32(bs)) // 从大端序字节数组中恢复整数
}


func ReadBytes(f *os.File, num int) []byte {
    buf := make([]byte, num) // 创建一个长度为 num 的字节切片
    _, err := f.Read(buf)    // 从文件读取 num 字节
    if err != nil {
        log.Fatal(err)       // 处理读取错误
    }
    return buf
}
