package main

import(
	"fmt"
	"sync"
)

func main(){
	var wg sync.WaitGroup
	wg.Add(6)
	for i:=1;i<=6;i++{
		go func(id int){
			defer wg.Done()
			sum(id)
		}(i)
	}
	wg.Wait()
}

func sum(id int){
	var r int
	for i:=0;i<10000000;i++{
		r += i
	}
	fmt.Println(id,r)	
}