// 关闭一个通道意味着不能再向这个通道发送值了。该特性可以向通道的接收方传达工作已经完成的信息。
// 但还可以继续读数据

package main

import (
	"fmt"
	"time"
)

func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
			time.Sleep(time.Second)
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    <-done
}