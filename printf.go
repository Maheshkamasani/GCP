package main
import "fmt"
func main() {
	var name string = "Maheswar Reddy Kamasani"
	const pi float64 = 3.1412
	win :=true
	x := 7
fmt.Println(len(name))
fmt.Println(name + " is a Trainee")
fmt.Printf("%.3f \n", pi)
fmt.Printf("%T \n", name)
fmt.Printf("%t \n", win)
fmt.Printf("%d \n", x)
fmt.Printf("%b \n", 35)
fmt.Printf("%c \n", 19)
fmt.Printf("%x \n", 25)
fmt.Printf("%e \n", pi)
}