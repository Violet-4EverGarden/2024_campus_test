package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	//var s = []int{1, 2, 3, 4, 5, 6}
	//fmt.Println(len(s), cap(s))
	//s1 := s[2:4]
	//fmt.Println(len(s1), cap(s1))
	//s1 = append(s1, 11) // 会发生什么变化？
	//fmt.Println(s, s1, len(s1), cap(s1))
	//s1 = append(s1, []int{12, 13}...) // 会发生什么变化？
	//fmt.Println(s, s1, len(s1), cap(s1))
	chanNum := 2
	chanQueue := make([]chan struct{}, chanNum)
	//count1 := 1 // count1控制数字打印，14次
	//count2 := 1 // count2控制字符打印，13次
	count := 1

	for i := 0; i < chanNum; i++ {
		chanQueue[i] = make(chan struct{}) // 无buffer的channel
	}

	exitChan := make(chan struct{}) // 退出通道，无buffer

	go func() {
		chanQueue[0] <- struct{}{} // 发送开始信号，启动第一个协程
	}()

	for i := 0; i < chanNum; i++ {
		curChan := chanQueue[i]
		var nextChan chan struct{}
		if i == chanNum-1 {
			nextChan = chanQueue[0]
		} else {
			nextChan = chanQueue[i+1]
		}

		go func() {
			// 必须使用无限循环for{}来阻塞等待前一个协程做发送信号操作
			// 否则每个协程只会接收和发送一次，那么当最后一个协程接收完执行发送操作时，没有协程再接收；由于通道无缓冲，此时会长时间阻塞导致deadlock
			for {
				<-curChan
				if count > 27 {
					exitChan <- struct{}{}
					return
				}
				if count%2 == 0 { // count为偶数打印字符
					fmt.Printf("%c%c", 'A'+count-2, 'A'+count-1)
				} else { // count为奇数打印数字
					fmt.Printf("%s%s", string(IntToBytes(count)), string(IntToBytes(count+1)))
				}
				count++
				nextChan <- struct{}{}
			}
		}()
	}

	<-exitChan
	fmt.Print("\n")
}

func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

//5 3
//192.168.0.1 1
//192.168.0.2 2
//192.168.0.3 3
//192.168.0.4 4
//192.168.0.5 5
//1 2
//2 3
//4 5
//3
//192.168.0.1 192.168.0.2
//192.168.0.2 192.168.0.3
//192.168.0.3 192.168.0.4
