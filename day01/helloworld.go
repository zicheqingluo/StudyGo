package main
import "fmt"

var name string
var (
    a string
    b int 
    c bool
    d float32
)
func foo() (int,string) {
    return 10,"yxk4"
}
func main(){
    var name1 string
    name1 = "yxk"
    var name2 string = "yxk1"
    var name3,age = "yxk3",20
    n := 10
                fmt.Printf("hello world %s\n",name1)
    fmt.Printf("%s 今年 %d 岁 \n",name3,age)

        fmt.Printf("%s 今年 %d 个硬币 \n",name2,n)

    x,_ := foo()
    _,y := foo()
    fmt.Printf("%d %s \n",x,y)
}