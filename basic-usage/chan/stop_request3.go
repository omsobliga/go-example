// 示例：抓取多个 url，其中一个失败，通知所有任务终止
// 子协程通知父协程和其他子协程，任务终止

package main

import (
	"fmt"
	"net/http"
	"sync"
)

func CheckStop(stopChan chan struct{}) bool {
	select {
	case <-stopChan:
		fmt.Println("recieved close notification")
		return true
	default:
	}
	return false
}

func ToStop(toStop chan struct{}) {
	fmt.Println("try notify close")
	select {
	case toStop <- struct{}{}:
		fmt.Println("notify close success")
	default:
		fmt.Println("notify close failed")
	}
}

func requestUris(num int, uris <-chan string, toStop chan struct{}, stopChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for uri := range uris {
		if CheckStop(stopChan) {
			return
		}

		resp, err := http.Get(uri)
		defer resp.Body.Close()

		if err != nil {
			panic(err)
		}
		if resp.StatusCode == 404 { // 请求失败，通知任务终止
			ToStop(toStop)
		}
		fmt.Println(uri, resp.StatusCode, num)
	}
}

func main() {
	max_num := 10
	uris := make(chan string, max_num)
	stopChan := make(chan struct{})  // 判断任务是否终止
	toStop := make(chan struct{}, 1)  // 终止操作的协调器
	go func() {  // 实现 stopChan 只被 close 一次
		<-toStop
		close(stopChan)
	}()
	var wg sync.WaitGroup
	for i := 0; i < max_num; i++ {
		wg.Add(1)
		go requestUris(i, uris, toStop, stopChan, &wg)
	}
	for i := 0; ; i++ {
		if CheckStop(stopChan) {
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
