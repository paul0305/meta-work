package task2

import (
	"fmt"
	"math"
	"sync"
	"sync/atomic"
	"time"
)

/**
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，
然后在主函数中调用该函数并输出修改后的值。
*/

func fun1(a *int) int {
	var b int
	b = *a + 10
	return b
}

/*
*
实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
*/
func angel(par *[]int) *[]int {
	for i := range *par {
		(*par)[i] *= 2
	}
	return par
}

/**
编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数
*/

func printOddAndEven() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i += 2 {
			fmt.Println(i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 2; i <= 10; i += 2 {
			fmt.Println(i)
		}
	}()
	wg.Wait()
	fmt.Println("Done")
}

/*
*
设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间
*/
func scheduleExe(task []func()) {
	var wg sync.WaitGroup
	wg.Add(len(task))
	for _, v := range task {
		go func(v func()) {
			defer wg.Done()
			start := time.Now()
			v()
			end := time.Since(start)
			fmt.Println("执行时间:", end)
		}(v)
	}
	wg.Wait()
}

/**
定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。
然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，
创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
*/

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	//面积
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	//周长
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	//面积
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	//周长
	return 2 * math.Pi * c.radius
}

/**
使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，
再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。
为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
*/

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	EmployeeID int
}

func (e *Employee) PrintInfo() {
	fmt.Printf("员工Name=%s,age=%d,employeeId=%d\n", e.Name, e.Age, e.EmployeeID)
}

/**
编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
*/

func chan1ToCha2() {
	var chan1 = make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			chan1 <- i
		}
		close(chan1)
	}()
	go func() {
		for {
			select {
			case num, ok := <-chan1:
				if !ok {
					fmt.Println("通道已关闭")
					return
				}
				fmt.Println("接收到数据:", num)
			case <-time.After(time.Second * 5):
				fmt.Println("超时")
				return
			default:
				fmt.Println("没有数据")
			}

		}
	}()
	time.Sleep(time.Second * 10)
}

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印
func buffChan() {
	var buffChan = make(chan int, 10)
	go func() {
		for i := 0; i < 100; i++ {
			buffChan <- i
		}
		close(buffChan)
	}()
	go func() {
		for {
			select {
			case i, ok := <-buffChan:
				if !ok {
					fmt.Println("通道已关闭")
					return
				}
				fmt.Println(i)
			default:
				fmt.Println("没有数据")
			}
		}
	}()
}

/**
编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
*/

func getSumResult() int {
	var sumMutex sync.Mutex
	var sum int
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				sumMutex.Lock()
				sum++
				sumMutex.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("sum:", sum)
	return sum
}

// 使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func atomUse() {
	var sum int32
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 1000; j++ {
				atomic.AddInt32(&sum, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("sum:", sum)
}
