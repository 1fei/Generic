// Uint32.go
package Generic

func compareUint32(a interface{}, b interface{}) int {
	switch {
	case a.(uint32) > b.(uint32):
		return 1
	case a.(uint32) < b.(uint32):
		return -1
	default:
		return 0
	}
}
