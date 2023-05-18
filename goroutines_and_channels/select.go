// TODO(wwfyde) 这种代码逻辑, 还不太熟悉

package main

func fibInSelect(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			println(<-c)
		}
		quit <- 0
	}()
	fibInSelect(c, quit)
}
