package main

import (
	"fmt"
	"time"
)
var FI_RESULT map[Integer]Integer = make(map[Integer]Integer,10)

type Integer int

func main() {
	now := time.Now()
	fmt.Println(fn(45))
	since := time.Since(now)
	fmt.Println(since.Seconds())
}

func fn(i Integer) Integer{
	if i==1 || i==2{
		return 1
	}
	tempResult := FI_RESULT[i]
	if tempResult != 0{
		return tempResult
	}
	result := fn(i-1) + fn(i-2)
	FI_RESULT[i] = result
	return result
}
