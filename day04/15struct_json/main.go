package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	Name string `json:"name"`
	Age int `json:"age"`
}



func main() {
	p1 := person {
		Name: "山治",
		Age: 100,
	}
	scoreMap := make(map[string]string)

	//序列化
	b,err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed,err:%v",err)
		return
	}
	fmt.Printf("%v \n",string(b))

	//反序列化
	str := `{"name":"理想","age":"18"}`
	
	json.Unmarshal([]byte(str),&scoreMap)
	fmt.Printf("%#v \n",scoreMap)

}