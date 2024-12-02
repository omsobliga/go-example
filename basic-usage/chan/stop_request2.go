// 示例：抓取多个 url，其中一个失败，通知所有任务终止
// 子协程通知父协程和其他子协程，任务终止

package main

import (
	"fmt"
	"net/http"
	"sync"
)

func checkClose(close_chan chan struct{}) bool {
	select {
	case <-close_chan:
		fmt.Println("recieved close notification")
		select {
		case close_chan <- struct{}{}:
		default:
		}
		return true
	default:
	}
	return false
}

func closeChan(close_chan chan struct{}) {
	fmt.Println("try notify close")
	select {
	case close_chan <- struct{}{}:
		fmt.Println("notify close success")
	default:
		fmt.Println("notify close failed")
	}
}

func requestUris(num int, uris <-chan string, close_chan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for uri := range uris {
		if checkClose(close_chan) {
			return
		}

		resp, err := http.Get(uri)
		defer resp.Body.Close()

		if err != nil {
			panic(err)
		}
		if resp.StatusCode == 404 { // 请求失败，通知任务终止
			closeChan(close_chan)
		}
		fmt.Println(uri, resp.StatusCode, num)
	}
}

func main() {
	max_num := 10
	uris := make(chan string, max_num)
	close_chan := make(chan struct{}, 1)
	var wg sync.WaitGroup
	for i := 0; i < max_num; i++ {
		wg.Add(1)
		go requestUris(i, uris, close_chan, &wg)
	}
	for i := 0; ; i++ {
		if checkClose(close_chan) {
			break
		}
		uri := "https://www.baidu.com"
		if i%3 == 1 {
			uri += "/err" + fmt.Sprint(i)
		}
		uris <- uri
	}
	wg.Wait()
	fmt.Println("finish")
}
