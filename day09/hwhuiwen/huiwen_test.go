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
	type test []struct {
		input string
		want string
	}
	tests := test{
		{input:"abcba",want:"是回文"},
		{input:"上海",want:"不是回文"},
	}
	for _,tc :=range tests {
		got := BackToText(tc.input)
		if got  != tc.want{
			
 		t.Errorf("not huiwen,want:%v,but %v",tc.want,got)
		}
	}
}