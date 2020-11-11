package main
import "fmt"
func main() {
	
 age :=18
switch age {
case 16 : fmt.Println("Get ready for collge")
case 18 : fmt.Println("Don't run behind the girls")
case 20 : fmt.Println("Get yourself a job!")
case 24 : fmt.Println("Get ready to marry")
default : fmt.Println("How old are you?")
}
}