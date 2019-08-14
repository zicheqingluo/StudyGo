package main


import "fmt"


func main() {
	s1 := make([]int,5,10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d \n",s1,len(s1),cap(s1))
	var nems []int
	for i:=0;i<10;i++{
		nems = append(nems,i)
		fmt.Printf("%v len(nems)%d cap(nems)%d ptf:%p \n",nems,len(nems),cap(nems),nems)

	}
	nems = append(nems[:3],nems[4:]...)
	fmt.Println(nems)
}