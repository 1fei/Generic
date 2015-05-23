package Generic

func compareString(a interface{}, b interface{}) int {
	switch {
	case a.(string) > b.(string):
		return 1
	case a.(string) < b.(string):
		return -1
	default:
		return 0
	}
}
