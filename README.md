# golang generic 


func TestMin(t *testing.T) {

	t.Log("test begin")
	{
		min := Min("int")
		t.Log("min(9, 10) =", min(9, 10))
		min = Min("string")
		t.Log("min('9', '10') =", min("9", "10"))
	}
	t.Log("test check type")
	{
		var min func(a, b int) int
		MakeFunc(Min("int"), &min)
		var ret int = min(9, 8)
		t.Log("test3 func ", ret)
	}
	{
		var min func(a, b string) string
		MakeFunc(Min("string"), &min)
		var ret string = min("10", "8")
		t.Log("ret = ", ret)
	}
	
}
