package main
import "fmt"

func main(){
	var a,b int;
	a = 5
	b = 10
	var res int
	res = max(a,b)
	fmt.Print(res)
}

func max(a,b int) int{
	if a > b {
		return a
	}
	return b
}