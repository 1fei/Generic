// Uint64.go
package Generic

func compareUint64(a interface{}, b interface{}) int {
	switch {
	case a.(uint64) > b.(uint64):
		return 1
	case a.(uint64) < b.(uint64):
		return -1
	default:
		return 0
	}
}
