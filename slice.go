package main
import "fmt"
func main() {
	
	EvenNum := [9]int{0,2,3,5,9,4,6,8,9}
	for i, value := range EvenNum {
	fmt.Println(value, i)

}
}