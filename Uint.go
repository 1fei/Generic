// Uint.go
package Generic

func compareUint(a interface{}, b interface{}) int {
	switch {
	case a.(uint) > b.(uint):
		return 1
	case a.(uint) < b.(uint):
		return -1
	default:
		return 0
	}
}
