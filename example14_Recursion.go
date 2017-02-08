package main
import "fmt"

func main(){
	
	var num int
	num = 5
	fmt.Printf("Factorial of %d is %d",num,fact(num))
}

func fact (num int) int {
	if(num <= 1){
		return 1
	}
	return num*fact(num-1)
}