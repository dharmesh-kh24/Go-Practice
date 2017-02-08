package main
import "fmt"

func main(){
	var a int
	var b int	
	
	a,b = 3,5
	swap(&a,&b)
	fmt.Printf("%d %d",a,b)
}

func swap (a,b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}