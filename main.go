package main

import (
	"fmt"
	"math"
	"sync"
)

func fibonacci(n int, fb chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	a, b,c := 0, 1,0
	for i := 1; i <= n; i++ {
		a=b
		b=c
		c=a+b
	}
	fb <- c 
}

func sqrt1(n float64, sqrt chan<- float64, wg *sync.WaitGroup) {
	defer wg.Done() 
	sqrt <- math.Sqrt(n) 
}

func fanIn(n int ,fb <-chan int, n1 float64, sqrt<-chan float64,out chan interface{}, wg *sync.WaitGroup){

    go func() {
        defer close(out)
        defer wg.Done()
        for {
            select {
            case f, ok := <-fb:
                if !ok {
                    return
                }
                out <- fmt.Sprintf("%v th Fibanocci number is %v",n,f)
            case s, ok := <-sqrt:
                if !ok {
                    return
                }
                out <- fmt.Sprintf("square root of %v is %v",n1,s)
				break
            }
        }
    }()
}


func main() {
	n:=0
	var n1 float64
	fmt.Print("Enter the number: ")
	fmt.Scanln(&n)
	fmt.Print("Input the number: ")
	fmt.Scanln(&n1)
	fb := make(chan int, 1)
	sqrt := make(chan float64, 1)
	out:=make(chan interface{},2)
	var wg sync.WaitGroup
	wg.Add(2)

	go fibonacci(n, fb, &wg)
	go sqrt1(n1, sqrt, &wg)
	fanIn(n,fb,n1,sqrt,out, &wg)
	for i:=0;i<2;i++{
		fmt.Println(<-out)
	}
}
