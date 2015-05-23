// Uint8.go
package Generic

func compareUint8(a interface{}, b interface{}) int {
	switch {
	case a.(uint8) > b.(uint8):
		return 1
	case a.(uint8) < b.(uint8):
		return -1
	default:
		return 0
	}
}
