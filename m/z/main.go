package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

//写代码实现两个 goroutine，其中一个


//产生随机数并写入到 go channel 中，另外一个从 channel 中读取数字并打印到标准输出。
//最终输出五个随机数。

func main() {

	test := make(chan int, 5)

	wg := new(sync.WaitGroup)
	wg.Add(1)

	rand.Intn(10)

	go func() {
		for i := 0; i < 5; i++ {
			ran := rand.Int()
			test<-ran
		}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(<-test)
		}
		wg.Done()
	}()
	wg.Wait()

	var people [][2]int
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || a[0] == b[0] && a[1] < b[1]
	})
}