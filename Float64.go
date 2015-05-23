// Float64.go
package Generic

func compareFloat64(a interface{}, b interface{}) int {
	switch {
	case a.(float64) > b.(float64):
		return 1
	case a.(float64) < b.(float64):
		return -1
	default:
		return 0
	}
}
