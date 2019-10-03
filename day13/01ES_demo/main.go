package main

import (
	"fmt"
	"github.com/olivere/elastic/v7"
	"context"
)

type student struct{
	Name string `json:"name"`
	Age int 	`json:"age"`
	Married	bool	`json:"married"`
}

func main(){
	//1.初始化链接得到一个client
	client,err := elastic.NewClient(elastic.SetURL("http://127.0.0.1:9200"))
	if err != nil {
		panic(err)
	}
	fmt.Println("connect to es success")

	p1 := student{Name:"suolong",Age:21,Married:false}

	put1,err := client.Index().
		Index("user").
		Type("go").
		BodyJson(p1).
		Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Printf("indexed user %s to index %s ,type %s\n",put1.Id,put1.Index,put1.Type)



}