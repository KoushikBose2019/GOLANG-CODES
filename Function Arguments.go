package main
import "fmt"
func main(){
/* local variable declaration*/
var a int=100
var b int=700

fmt.Printf("before swap the value of a is %d\n",a)
fmt.Printf("before swap the value of b is %d\n",b)

swap(a,b)

fmt.Printf("after swap the value of a is %d\n",a)
fmt.Printf("after swap the value of b is %d\n",b)
}
func swap(x,y int)int{
	var temp  int
	temp=x
	x=y
	y=temp
	return temp;
}











}