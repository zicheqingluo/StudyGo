package hwhuiwen

//import "fmt"

//BackToText 回文函数
func BackToText(ss string) string {

	//ss := "a山西运煤车煤运西山a"
	bd := ",'"
	var flag bool
	
	r := make([]rune,0,len(ss))
	for _,c := range ss {
		flag = true
		for _,b := range bd {
			
			if c == b{
				flag = false
			}
		}
		if flag {
			r = append(r,c)

		}
		
		
	}
	//fmt.Println("[]rune:",r)
	for i :=0;i<len(r)/2;i++ {
		if r[i] != r[len(r)-1-i] {
			//fmt.Println("不是回文")
			return "不是回文"
		}
	}
	//fmt.Println("是回文")
	return "是回文"


}