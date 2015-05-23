package Generic

func compareInt16(a interface{}, b interface{}) int {
	switch {
	case a.(int16) > b.(int16):
		return 1
	case a.(int16) < b.(int16):
		return -1
	default:
		return 0
	}
}
