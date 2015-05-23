// Int8.go
package Generic

func compareInt8(a interface{}, b interface{}) int {
	switch {
	case a.(int8) > b.(int8):
		return 1
	case a.(int8) < b.(int8):
		return -1
	default:
		return 0
	}
}
