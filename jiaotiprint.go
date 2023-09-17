package main

import "fmt"

func main() {
	chanNum := 3                                // 定义通道数量 == 协程数量
	chanQueue := make([]chan struct{}, chanNum) // 创建通道队列的切片

	var count int                   // 计数器变量
	exitChan := make(chan struct{}) // 创建推出信号通道

	// 初始化通道队列！
	for i := 0; i < chanNum; i++ {
		chanQueue[i] = make(chan struct{})
	}

	// 必须使用协程对第一个通道发送唤醒信号
	// 因为通道是无缓冲的，直接发送的话主协程会一直阻塞在这导致死锁！
	go func() {
		chanQueue[0] <- struct{}{}
	}()

	// 启动对应数量的协程
	for i := 0; i < chanNum; i++ {
		var nextChan chan struct{} // 获取下一个协程的通道，因为当前协程执行完后要往里面发送信号
		if i == chanNum-1 {
			nextChan = chanQueue[0]
		} else {
			nextChan = chanQueue[i+1]
		}
		curChan := chanQueue[i] // 获取当前协程的通道

		// 开启协程
		go func(i int) {
			// 必须使用无限循环for{}来阻塞等待前一个协程做发送信号操作
			// 否则每个协程只会接收和发送一次，那么当最后一个协程接收完执行发送操作时，没有协程再接收；由于通道无缓冲，此时会长时间阻塞导致deadlock
			for {
				<-curChan // 读取上一个协程往当前通道发送的信号，表示可以输出；若没有则一直阻塞
				if count >= 100 {
					exitChan <- struct{}{}
					return
				}
				count++
				fmt.Printf("当前输出协程：%d，数字：%d\n", i, count)
				nextChan <- struct{}{} // 发送通道信号给下一个协程，相当于将其唤醒
			}
		}(i)
	}

	<-exitChan // 主协程阻塞，等待退出信号
	fmt.Print("Print Done")
}
