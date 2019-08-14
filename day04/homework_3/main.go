package main

import (
	"fmt"
	"os"
)

var (
	allStudent map[int64]*student  //学生信息存储
)
//定义学生的结构体
type student struct {
	id	int64
	name	string

}
//构造函数
func newStudent(id int64,name string) *student {
	return &student{
		id: id,
		name:name,
	}
}
//查看所有学生信息
func (s student) showAllStudent() {
	for k,v := range allStudent{
		fmt.Printf("学号：%d	姓名：%s \n",k,v.name)
	}
}
//添加学生
func (s student) addStudent() {
	var (
		id	int64
		name string
	)
	fmt.Println("请输入学号：")
	fmt.Scanln(&id)
	fmt.Println("请输入学生姓名：")
	fmt.Scanln(&name)
	
	allStudent[id] = newStudent(id,name)

}
//删除学生
func (s student) delStudent() {
	var delId int64
	fmt.Println("请输入学号：")
	fmt.Scanln(&delId)
	delete(allStudent,delId)


}


func main(){
	allStudent = make(map[int64]*student,10) //初始化map
	for {
		fmt.Println("欢迎光临学生管理系统！")
		fmt.Println(`
		1.查看所有学生信息
		2.增加学生信息
		3.删除学生信息
		4.退出
		`)
		fmt.Println("请输入操作序号：")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你的选项是【%d】\n",choice)
		var s student
		switch choice {
		case 1:
			
			s.showAllStudent()
		case 2:
			s.addStudent()
		case 3:
			s.delStudent()
		case 4:
			os.Exit(1)
		default:
			fmt.Println("您输入的操作有误,请重新输入！")
		}
	}
	
}