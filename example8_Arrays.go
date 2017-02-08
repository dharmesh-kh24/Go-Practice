package main

import "fmt"

func main(){
	var numbers[10] int
	for i := 0;i < 10;i++ {
		numbers[i] = i+1
	}
	
	for i := 0;i < len(numbers);i++ {
		fmt.Printf("%d ",numbers[i])
	}
}