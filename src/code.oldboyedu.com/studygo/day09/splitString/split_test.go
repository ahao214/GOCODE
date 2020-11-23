package splitString

import(
	"testing"
	"reflect"
)


func TestSplit(t *testing.T){
	got:=Split("babcef","b")
	want:=[]string{"", "a", "cef"}
	if !reflect.DeepEqual(got,want){
		//测试用例失败了
		t.Errorf("want:%v but got:%v\n",want,got)
	}
}



func TestSplit2(t *testing.T){
	got:=Split("a:b:c",":")
	want:=[]string{"a", "b", "c"}
	if !reflect.DeepEqual(got,want){
		//测试用例失败了
		t.Errorf("want:%v but got:%v\n",want,got)
	}
}