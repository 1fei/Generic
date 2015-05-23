// Generic project Generic.go
package Generic

import "fmt"
import "reflect"

//golang没有泛型被广为吐槽，在golang中要实现泛型都是以下方式实现的：
//1、把用到的类型转为interface{}
//2、把算法中用到的函数封装成接口
//在运行中通常碰到两类问题：
//1、类型转为interface{}丢失了类型信息
//2、大量的基础类型需要重载接口，代码显的很丑。
//本解决方案中，实现了以下工作：
//1、通过反射把各常见类型中需要的通用函数同一化。
//2、通过一次性的类型检查，校验类型的完整度
type Number struct {
	Compare func(interface{}, interface{}) int
}

type Compare struct {
	Compare func(interface{}, interface{}) int
}

//var MapNumber = map[reflect.Kind]*Number{
//	reflect.Int: &Number{
//		Compare: compareInt},
//	reflect.Int8: &Number{
//		Compare: compareInt8},
//	reflect.Int16: &Number{
//		Compare: compareInt16},
//	reflect.Int32: &Number{
//		Compare: compareInt32},
//	reflect.Int64: &Number{
//		Compare: compareInt64},
//	reflect.Uint: &Number{
//		Compare: compareUint},
//	reflect.Uint8: &Number{
//		Compare: compareUint},
//	reflect.Uint16: &Number{
//		Compare: compareUint},
//	reflect.Uint32: &Number{
//		Compare: compareUint},
//	reflect.Uint64: &Number{
//		Compare: compareUint},
//}
var MapCompare = map[string]*Compare{
	"int": &Compare{
		Compare: compareInt},
	"int8": &Compare{
		Compare: compareInt8},
	"int16": &Compare{
		Compare: compareInt16},
	"int32": &Compare{
		Compare: compareInt32},
	"int64": &Compare{
		Compare: compareInt64},
	"uint": &Compare{
		Compare: compareUint},
	"uint8": &Compare{
		Compare: compareUint8},
	"uint16": &Compare{
		Compare: compareUint16},
	"uint32": &Compare{
		Compare: compareUint32},
	"uint64": &Compare{
		Compare: compareUint64},
	"string": &Compare{
		Compare: compareString},
}

func GetCompare(tp string) *Compare {
	compare := MapCompare[tp]
	if compare == nil {
		panic(fmt.Sprintf("不支持的类型：%v", tp))
	}
	return compare
}

func GetComparefunc(tp string) func(interface{}, interface{}) int {
	return GetCompare(tp).Compare
}

func MakeFunc(fIn interface{}, fOut interface{}) {
	inFunc := reflect.ValueOf(fIn)
	inType := reflect.TypeOf(fIn)
	outFunc := reflect.ValueOf(fOut).Elem()
	outType := outFunc.Type()
	if inType.NumIn() != outType.NumIn() && inType.NumOut() != outType.NumOut() {
		panic(fmt.Errorf("参数不匹配：in[%d:%d] out[%d:%d]",
			inType.NumIn(), inType.NumOut(), outType.NumIn(), outType.NumOut()))
	}

	for i := 0; i < outType.NumIn(); i++ {
		ot := outType.In(i)
		it := inType.In(i)
		if ot.Kind() != it.Kind() && it.Kind() != reflect.Interface {
			panic(fmt.Errorf("in:ot != it ot=%s, it=%s", ot.Kind(), it.Kind()))
		}
	}

	for i := 0; i < outType.NumOut(); i++ {
		ot := outType.Out(i)
		it := inType.Out(i)
		if ot.Kind() != it.Kind() {
			if it.Kind() != reflect.Interface {
				panic(fmt.Errorf("out:ot != it ot=%s, it=%s", ot.Kind(), it.Kind()))
			} else {

			}
		}
	}

	f := func(in []reflect.Value) []reflect.Value {
		ret := inFunc.Call(in) //返回值
		for i := 0; i < outType.NumOut(); i++ {
			ot := outType.Out(i)
			it := inType.Out(i)
			if ot.Kind() != it.Kind() && it.Kind() == reflect.Interface {
				if ret[i].CanInterface() {
					e := ret[i].Elem()
					ret[0] = e
				}
			}
		}
		return ret
	}
	v := reflect.MakeFunc(outFunc.Type(), f)
	outFunc.Set(v)
}

func Min(tp string) func(interface{}, interface{}) interface{} {
	compare := GetCompare(tp).Compare
	return func(a, b interface{}) interface{} {
		if compare(a, b) > 0 {
			return b
		}
		return a
	}
}
