package main
import "fmt"

func main(){
	var numbers[10] int
	for i := 1;i <= 10;i++ {
		numbers[i-1] = i;
	}
	
	for i := range numbers {
		fmt.Printf("%d ",numbers[i])
	}
	
	var arr[] int

	arr = append(arr,11)
	arr = append(arr,12)
	arr = append(arr,13)

	fmt.Println()
	for i := range arr {
		fmt.Printf("%d ",arr[i])
	}
}