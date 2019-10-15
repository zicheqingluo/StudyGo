package main

import (
	"fmt"
  	"github.com/aliyun/alibaba-cloud-sdk-go/services/cms"
  
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
  
)

func main() {
	client, err := cms.NewClientWithAccessKey("cn-beijing", "LTAI6GkrOCY5TAnV", "RB43X4Es2yF2XHMcNpY4qcOgzbsIUj")

	request := cms.CreateDescribeAlertHistoryListRequest()
	request.Scheme = "https"

  request.StartTime = "1567267200000"
  request.EndTime = "1569859200000"
  request.PageSize = requests.NewInteger(100)
  

	response, err := client.DescribeAlertHistoryList(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %T\n", response)
}
