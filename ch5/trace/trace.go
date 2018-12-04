package trace

import (
	"log"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")()
	// ...这里是一些处理...
	time.Sleep(10 * time.Second) // 通过休眠仿真慢操作
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
