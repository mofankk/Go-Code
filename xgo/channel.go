package main

// channel 阻塞解除测试
func main() {
	println("start main")
	ch := make(chan bool)
	go func() {
		println("come into goroutine")
		ch <- true
	}()

	//ch <- false   // fatal error: all goroutines are asleep - deadlock!
	println("do something else")
	<-ch
	close(ch)

	println("end main")
}
