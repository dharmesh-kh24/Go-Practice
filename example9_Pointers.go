package main
import "fmt"

func main(){
	var a int	
	a = 5
	var ip *int
	ip = &a
	fmt.Printf("%d %d",a,*ip)
}