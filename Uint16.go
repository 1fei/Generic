// Uint16.go
package Generic

func compareUint16(a interface{}, b interface{}) int {
	switch {
	case a.(uint16) > b.(uint16):
		return 1
	case a.(uint16) < b.(uint16):
		return -1
	default:
		return 0
	}
}
