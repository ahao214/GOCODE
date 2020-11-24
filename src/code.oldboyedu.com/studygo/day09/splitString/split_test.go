package splitString

import(
	"testing"
	"reflect"
)



// func TestSplitGroup(t *testing.T){
// 	type testCase struct{
// 		str string
// 		seq string
// 		want[]string
// 	}

// 	testGroup:=[]testCase{
// 		testCase{"babcbef","b",[]string{"","a","c","ef"}},
// 		testCase{"a:b:c",":",[]string{"a","b","c"}},
// 		testCase{"abcef","bc",[]string{"a","ef"}},
// 	}

// 	for _,tc:=range testGroup{
// 		got:=Split(tc.str,tc.seq)
// 		if !reflect.DeepEqual(got,tc.want){
// 			t.Fatalf("want:%v but got:%v\n",tc.want,got)
// 		}
// 	}
// }


//子测试
func TestSplitMap(t *testing.T){
	type testCase struct{
		str string
		seq string
		want[]string
	}


	testGroup:=map[string]testCase{
		"case 1":testCase{"babcbef","b",[]string{"","a","c","ef"}},
		"case 2":testCase{"a:b:c",":",[]string{"a","b","c"}},
		"case 3":testCase{"abcef","bc",[]string{"a","ef"}},
	}


	for name,tc:=range testGroup{
		t.Run(name,func(t *testing.T){
			got:=Split(tc.str,tc.seq)
			if !reflect.DeepEqual(got,tc.want){
				t.Fatalf("want:%#v but got:%#v\n",tc.want,got)
			}
		})
	}
}

//基准测试
func BenchmarkSplit(b *testing.B){
	for i:=0;i<b.N;i++{
		Split("a:b:c:d:e:f:g",":")
	}
}

//性能测试
func benchmarkFib(b *testing.B,n int){
	for i:=0;i<b.N;i++{
		Fib(n)
	}
}

func BenchmarkFib1(b *testing.B){
	benchmarkFib(b,1)
}

func BenchmarkFib2(b *testing.B){
	benchmarkFib(b,2)
}

func BenchmarkFib3(b *testing.B){
	benchmarkFib(b,3)
}
