package Generic

func compareInt64(a interface{}, b interface{}) int {
	switch {
	case a.(int64) > b.(int64):
		return 1
	case a.(int64) < b.(int64):
		return -1
	default:
		return 0
	}
}
