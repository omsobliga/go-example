// 示例：抓取多个 url，其中一个失败，通知所有任务终止
// 子协程通知父协程，任务终止

package main

import (
	"fmt"
	"net/http"
	"sync"
)

func checkClose(close <-chan struct{}) bool {
	select {
	case <-close:
		fmt.Println("recieved close notification")
		return true
	default:
	}
	return false
}

func closeChan(close chan<- struct{}) {
	fmt.Println("try notify close")
	select {
	case close <- struct{}{}:
		fmt.Println("notify close success")
	default:
		fmt.Println("notify close failed")
	}
}

func requestUri(num int, uri string, close chan<- struct{}) {
	if num%3 == 1 {
		uri += "/err"
	}
	resp, err := http.Get(uri)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 404 { // 请求失败，通知任务终止
		closeChan(close)
	}
	defer resp.Body.Close()
	fmt.Println(num, uri, resp.StatusCode)
}

func main() {
	max_num := 10
	pool := make(chan struct{}, max_num)
	close := make(chan struct{}, 1)
	var wg sync.WaitGroup
	for i := 0; ; i++ {
		if checkClose(close) {
			break
		}
		wg.Add(1)
		pool <- struct{}{}
		go func(i int) {
			uri := "https://www.baidu.com"
			requestUri(i, uri, close)
			wg.Done()
			<-pool
		}(i)
	}
	wg.Wait()
	fmt.Println("finish")
}
