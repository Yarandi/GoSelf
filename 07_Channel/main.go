package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {

	intChan := make(chan int)

	go CalculateValue(intChan)

	result := <- intChan
	log.Println(result)
}


func CalculateValue(intChan chan int){

	rand.Seed(time.Now().UnixMicro())
	rand := rand.Intn(10009)

	intChan <- rand
}

