package Generic

func compareInt32(a interface{}, b interface{}) int {
	switch {
	case a.(int32) > b.(int32):
		return 1
	case a.(int32) < b.(int32):
		return -1
	default:
		return 0
	}
}
