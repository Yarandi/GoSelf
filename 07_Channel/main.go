//package main
//
//import (
//	"fmt"
//	"time"
//)

//func main() {
//
//	intChan := make(chan int)
//
//	go CalculateValue(intChan)
//
//	result := <- intChan
//	log.Println(result)
//}
//
//
//func CalculateValue(intChan chan int){
//
//	rand.Seed(time.Now().UnixMicro())
//	rand := rand.Intn(10009)
//
//	intChan <- rand
//}

package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go printNumbers()
	time.Sleep(10 * time.Second)
	fmt.Println("Main function")
}
