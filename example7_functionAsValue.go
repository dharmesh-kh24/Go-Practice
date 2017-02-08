package main
import(
	"fmt"
	"math"
)

func main(){
	getSquareRoot := func(x float64) float64{
				return math.Sqrt(x);
				}
	fmt.Print(getSquareRoot(9));
}