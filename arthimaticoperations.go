package main
import "fmt"
func main() {
	x,y := 10,20
fmt.Println("x + y =", x+y)
fmt.Println("x - y =", x-y)
fmt.Println("x * y =", x*y)
fmt.Println("x / y =", x/y)
fmt.Println("x mod y =", x%y)

isbool := true
hate := false

fmt.Println(isbool && hate)
fmt.Println(isbool || hate)
fmt.Println(isbool)

}