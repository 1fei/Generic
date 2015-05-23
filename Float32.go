// Float32.go
package Generic

func compareFloat32(a interface{}, b interface{}) int {
	switch {
	case a.(float32) > b.(float32):
		return 1
	case a.(float32) < b.(float32):
		return -1
	default:
		return 0
	}
}
