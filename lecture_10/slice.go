package main
import "fmt"

func main(){
	c := [...]string{"abdullah","asgor","jony","ismile"}
	fmt.Println(c)
	fmt.Println(len(c))
	
	//slice
	d :=c[0:4]
	fmt.Println(d)
	//x := make([]string,10)
	fmt.Println("10 represents the capacity of the underlying array which the slice")
	x := make([]float64, 5, 5)
	fmt.Println(x)
	fmt.Println("")
	var fruits []string
	fruits = append(fruits, "Apple","Banana")
	fmt.Println(fruits)
	
	//Slice Functions
	fmt.Println("")
	fmt.Println("Slice Functions Append & Copy")
	slice1 := []int{1,2,3}
	slice2 := append(slice1, 4, 5)
	fmt.Println(slice1, slice2)
	
	// Slice copy	
	slicecopy1 := []int{1,2,3,}
	slicecopy2 := make([]int, 2)
	copy(slicecopy2, slicecopy1)
	fmt.Println(slicecopy1, slicecopy2)
	
	
}