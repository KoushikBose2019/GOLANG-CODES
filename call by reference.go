package main
import "fmt"
/* call by reference*/
func swap(x* int,y* int){
	var temp int
	temp=*x
	*x=*y
	*y=temp
}
func main(){
var a int=100
var b int=200
fmt.Printf("before swap value of a: %d\n",a)
fmt.Printf("before swap value odf b:%d\n",b)
swap(&a,&b)
fmt.Printf("after swap value of a: %d\n",a)
fmt.Printf("after swap value of b:%d\n",b)
}