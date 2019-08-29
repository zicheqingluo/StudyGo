package hwhuiwen

import (
	"testing"
)

// func Test1BackToText(t *testing.T) {
// 	got := BackToText("abcba")
// 	want := "是回文"
// 	if got != want {
// 		t.Errorf("not huiwen,want:%v,but %v",want,got)
// 	}
// }
// func Test2BackToText(t *testing.T) {
// 	got := BackToText("上海")
// 	want := "是回文"
// 	if got != want {
// 		t.Errorf("not huiwen,want:%v,but %v",want,got)
// 	}
// }

func TestBackToText(t *testing.T) {
	type test struct {
		input string
		want string
	}
	tests := map[string]test{
		"test1":{input:"abc,'ba",want:"是回文"},
		"test2":{input:"油灯少灯油",want:"是回文"},
	}
	for name,tc :=range tests {
		t.Run(name,func(t *testing.T){
			got := BackToText(tc.input)
			if got  != tc.want{
				
			 t.Errorf("not huiwen,want:%v,but %v",tc.want,got)
		}
		})
		}
	}


func BenchmarkBackToText(b *testing.B){
	for i:=0;i<b.N;i++{
		BackToText("沙河有沙又有河")
	}
}
