package main
import "fmt"

//切片
func main(){
	var s1 []int //定义一个存放int类型元素的切片
	var s2 []string // 定义一个存放string类型的元素
	//初始化
	s1 = []int{1,2,3}
	s2 = []string{"沙河","西安","丽江"}
	fmt.Println(s1,s2)

	a := [5]int{55,56,57,58,59}
	b := a[1:4]
	fmt.Println(b)
	fmt.Printf("type of a:%T \n",a)
	fmt.Printf("type of b:%T \n",b)

	c := [...]string{"北京", "上海", "广州", "深圳", "成都", "重庆"}
	fmt.Printf("c: %v  type:%T len:%d cap:%d \n",c,c,len(c),cap(c))
	d := c[1:3]
	fmt.Printf("d: %v type:%T len:%d cap:%d \n",d,d,len(d),cap(d))
	e := d[1:5]
	fmt.Printf("e : %v type:%T len:%d cap:%d \n",e,e,len(e),cap(e))
	
}