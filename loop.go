package main
import "fmt"
func main(){
var a int=14
var b int=0
numbers:=[7]int{1,2,3,4}
for a:=0;a<10;a++{
	fmt.Printf("value of a : %d\n",a)
}
for b<a{
	b++
	fmt.Printf("value of c : %d\n",b)

}
for i,x:=range numbers{
	fmt.Printf("value of x =%d at %d\n",x,i)
}

}