package Generic

import (
	"reflect"
	"testing"
)

/*
golang没有泛型被广为吐槽，在golang中要实现泛型都是以下方式实现的：
1、把用到的类型转为interface{}
2、把算法中用到的函数封装成接口
在运行中通常碰到两类问题：
1、类型转为interface{}丢失了类型信息
2、大量的基础类型需要重载接口，代码显的很丑。
本解决方案中，实现了以下工作：
1、构建一个通用包，封装了基础类型泛型过程中需要做的工作
2、通过反射把各常见类型中需要的通用函数同一化。
3、通过一次性的类型检查，校验类型的完整度

用户使用时只需引用包就可以了
使用前，通过类型读取相关的函数（函数集）比较，数值运算等
算法中读取到的函数便可。
*/
func Sort(ay []interface{}) {
	nLen := len(ay)
	if nLen == 0 {
		return
	}
	compare := GetCompare(reflect.TypeOf(ay[0]).Name()).Compare //用的接口多用这个函数
	for i := 0; i < nLen-1; i++ {
		for j := i + 1; j < nLen; j++ {
			if compare(ay[i], ay[j]) > 0 {
				k := ay[i]
				ay[i] = ay[j]
				ay[j] = k
			}
		}
	}
}

func TestSort(t *testing.T) {
	t.Log("test begin")
	{
		ay := []interface{}{"10", "1", "9", "8"}
		Sort(ay)
		t.Log("sort string ay=", ay)
	}
	{
		ay := []interface{}{int(10), int(1), int(9), int(8)}
		Sort(ay)
		t.Log("sort int ay=", ay)
	}
	{
		ay := []interface{}{uint(10), uint(1), uint(9), uint(8)}
		Sort(ay)
		t.Log("sort uint ay=", ay)
	}

}

func TestMin(t *testing.T) {
	t.Log("test begin")
	{
		min := Min("int")
		t.Log("min(9, 10) =", min(9, 10))
	}

	t.Log("test check type")
	{
		var min func(a, b int) int
		MakeFunc(Min("int"), &min)
		var ret int = min(9, 8)
		t.Log("test3 func ", ret)
	}
	{
		var min func(a, b string) string
		MakeFunc(Min("string"), &min)
		var ret string = min("10", "8")
		t.Log("ret = ", ret)
	}

}
