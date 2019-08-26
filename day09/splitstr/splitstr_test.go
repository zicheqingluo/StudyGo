package splitstr

import (
	"reflect"
	"testing"
)

func TestSplit(t *testing.T) {
	ret := Split("babcbef","b")
	want := []string{"","a","c","ef"}
	if !reflect.DeepEqual(ret,want) {
		//测试用例失败
		t.Errorf("want:%v  but got:%v \n",want,ret)
	}
}

func Test2Split(t *testing.T) {
	ret := Split("a:b:c",":")
	want := []string{"a","b","c"}
	if !reflect.DeepEqual(ret,want) {
		//测试用例失败
		t.Errorf("want:%v  but got:%v \n",want,ret)
	}
}