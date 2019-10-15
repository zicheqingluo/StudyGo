package main

import (
	"encoding/json"
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
  request.PageSize = requests.NewInteger(1)
  

	response, err := client.DescribeAlertHistoryList(request)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println(response)
	
	
	type ContactSmses struct{
		ContactSms []interface{}  `json:"ContactSms"`
	}
	type ContactMails struct {
		ContactMail []interface{}`json:"ContactMail"`
	}	

	type Contacts   struct {
		Contact []interface{} `json:"Contact"`
	}

	type ContactALIIMs struct {
		ContactALIIM []interface{} `json:"ContactALIIM"`
	}

	type ContactGroups struct {
		ContactGroup []interface{} `json:"ContactGroup"`
	}

	

	type AlarmHistory []struct {
		Value        string `json:"Value"`
		LastTime     int    `json:"LastTime"`
		Webhooks     string `json:"Webhooks"`
		ContactSmses ContactSmses `json:"ContactSmses"` 
		RuleName        string `json:"RuleName"`
		GroupID         string `json:"GroupId"`
		AlertName       string `json:"AlertName"`
		EvaluationCount int    `json:"EvaluationCount"`
		Status          int    `json:"Status"`
		AlertState      string `json:"AlertState"`
		MetricName      string `json:"MetricName"`
		ContactMails ContactMails `json:"ContactMails"`
		AlertTime  int64  `json:"AlertTime"`
		Dimensions string `json:"Dimensions"`
		RuleID     string `json:"RuleId"`
		Contacts Contacts `json:"Contacts"`
		Namespace     string `json:"Namespace"`
		ContactALIIMs ContactALIIMs  `json:"ContactALIIMs"`
		ContactGroups ContactGroups `json:"ContactGroups"`
		Expression   string `json:"Expression"`
		Level        string `json:"Level"`
		InstanceName string `json:"InstanceName"`
	}

    type AlarmHistoryList struct {
		AlarmHistory	AlarmHistory `json:"AlarmHistory"`

	} 

	type resData struct {
		AlarmHistoryList	AlarmHistoryList	`json:"AlarmHistoryList"`
		RequestID string `json:"RequestId"`
		Success   bool   `json:"Success"`
		Code      string `json:"Code"`
		Total     int    `json:"Total"`
	}



	var Data resData
	res := fmt.Sprintf(`%v`,response)
	fmt.Println(res)
	res1 := `{"AlarmHistoryList":{"AlarmHistory":[{"Value":"100","LastTime":1037048011,"Webhooks":"","ContactSmses":{"ContactSms":[]},"RuleName":"diskusage_utilization","GroupId":"14103","AlertName":"applyTemplate371eca68-5217-4655-95e7-084649c2a852","EvaluationCount":3,"Status":2,"AlertState":"ALARM","MetricName":"diskusage_utilization#60","ContactMails":{"ContactMail":[]},"AlertTime":1569824888000,"Dimensions":"{\"device\":\"/dev/vdb1\",\"instanceId\":\"i-25u9dm3xi\"}","RuleId":"","Contacts":{"Contact":[]},"Namespace":"acs_ecs","ContactALIIMs":{"ContactALIIM":[]},"ContactGroups":{"ContactGroup":[]},"Expression":"$Average>=90","Level":"P2","InstanceName":""}]},"RequestId":"42827853-C3F7-411F-88BA-BABFF6770467","Success":true,"Code":"200","Total":2039}`
	err = json.Unmarshal([]byte(res1),&Data)
	if err != nil {
		fmt.Println("json unmarshal err:",err)
	}
	fmt.Printf("%#v",Data)


}
