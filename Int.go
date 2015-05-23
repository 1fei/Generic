// Int
package Generic

func compareInt(a interface{}, b interface{}) int {
	switch {
	case a.(int) > b.(int):
		return 1
	case a.(int) < b.(int):
		return -1
	default:
		return 0
	}
}
