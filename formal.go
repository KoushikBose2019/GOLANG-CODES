package main
import "fmt"

var a int=24;
func main(){
var a int=10
var b int=40
var c int=0
fmt.Printf("value of a in main()=%d\n ",a);
c=sum(a,b);
fmt.Printf("value of b in main()=%d\n",c);
}
func sum(a,b int)int{
	fmt.Printf("value of a in sum()=%d\n",a);
	fmt.Printf("value of b in sum()=%d\n",b);
	return a+b;
}










	
}