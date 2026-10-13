package main

import "fmt"

func hello() {
	fmt.Println("hi")
}

//func main() {
//	go hello()// 这里的go不是调用函数并且等待而是启动函数之后立即继续往下执行
//
//	channel := make(chan int) //goroutine之间传递数据的一个通道 这个通道是int类型的
//								//是无缓冲的channel 发送和接收必须准备好
//	channel <- 5	//发送5
//	number := <-channel		//接收数据
//		//如果两个都同时在一个goroutine中执行 程序回卡住 因为发送时没有其他的goroutine接收
//
//	channel := make(chan int, 6)	//不是六条通道是buffer的容量是6 这个chan最多存6个int
//
//	channel <- 5
//	number := <-channel
//		//可以在同一个goroutine中运行 因为5会被先放进buffer 然后再取出来
//
//	//无缓冲channel 发送和接收必须同步
//	//有缓冲channel buffer没满的时候可以先发送 buffer不为空的时候可以接收
//		//buffer满了还发送 就block
//		//buffer空了还接收 也block
//
//}

// 只发送channel
func sendNum(ch chan<- int) {
	ch <- 5
}

// 只接收channel
func receiveNum(ch <-chan int) {
	number := <-ch
	fmt.Println(number)
}

func main() {
	numbers := make(chan int)

	go sendNum(numbers)
	receiveNum(numbers)
}
