package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 工作池本质上是生产者消费者模型
// 可以有效控制goroutine数量，防止暴涨
// 需求：
//    计算一个数字的各个位数之和，例如数字123，结果为1+2+3=6
//    随机生成数字进行计算

type Job struct {
	Id      int
	RandNum int // 需要计算的随机数
}

type Result struct {
	job *Job // 这里必须传对象实例
	sum int  // 求和
}

func main() {
	// 随机种子
	rand.Seed(time.Now().UnixNano())
	// 创建通道
	jobChan := make(chan *Job, 128)
	resultChan := make(chan *Result, 128)

	// 创建工作池,也可以理解为消费者
	CreatePool(64, jobChan, resultChan)

	// 开启打印结构协程
	go func(resultChan chan *Result) {
		for r := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", r.job.Id, r.job.RandNum, r.sum)
		}
	}(resultChan)

	// 循环创建job,输入到通道，也可以理解为生产者
	var id int
	for {
		id++
		rNum := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: rNum,
		}
		jobChan <- job
	}
}

// 创建工作池, 消费者
func CreatePool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去跑程序
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobChan {
				rNum := job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for rNum != 0 {
					// 一直取最大的一位
					tmp := rNum % 10
					sum += tmp
					rNum /= 10
				}
				// 将结构塞入通道
				r := &Result{
					job: job,
					sum: sum,
				}
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}