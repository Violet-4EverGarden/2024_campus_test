package main

import (
	"fmt"
)

//func IntToBytes(n int) []byte {
//	data := int64(n)
//	bytebuf := bytes.NewBuffer([]byte{})
//	binary.Write(bytebuf, binary.BigEndian, data)
//	return bytebuf.Bytes()
//}

func main() {
	num := 1234567
	bytes := IntToBytes(num)
	fmt.Println(bytes) // 输出：[73 150 2 210 31 15 66]
}
