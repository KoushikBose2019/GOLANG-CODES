package main
import "fmt"
func main(){
var a=[4][2]int{{0,0},{1,2},{2,4},{4,6}}
var i,j int
for i=0;i<4;i++{
	for j=0;j<2;j++{
		fmt.Printf("a[%d][%d]=%d\n",i,j,a[i][j])
	}
}

}