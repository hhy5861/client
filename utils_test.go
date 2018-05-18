package client
//
//import "testing"
//
//var (
//	ut *Utils
//)
//
//func init() {
//	ut = NewUtil()
//}
//
//func Test_GenerateRangeNum(t *testing.T) {
//	min, max := 1, 3
//
//	num := ut.GenerateRangeNum(min, max)
//	if min > num || max < num {
//		t.Log("failure")
//	} else {
//		t.Log("ok")
//	}
//
//}
//
//func Benchmark_GenerateRangeNum(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		min, max := 1, 3
//		ut.GenerateRangeNum(min, max)
//	}
//
//	b.Log("ok")
//}
