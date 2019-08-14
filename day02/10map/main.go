package main


import "fmt"

func main(){
	var m1 map[string]int
	fmt.Println(m1 == nil)
	m1 = make(map[string]int,10)
	m1["北京"] = 18
	m1["南京"] = 11
	fmt.Println(m1)
	fmt.Println(m1["北京"])
	fmt.Println(m1["西安"])
	value,ok := m1["西安"]
	if !ok {
		fmt.Println("无此key")

	}else{
		fmt.Println(value)
	}

	for k := range m1{
		fmt.Println(k)
	}

	for _,v := range m1{
		fmt.Println(v)
	}
	delete(m1,"北京")
	fmt.Println(m1)
	delete(m1,"西安")
}