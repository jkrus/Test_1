package main

var justString string

func someFunc() {
	v := createHugeString(1 << 62)
	bytes := make([]byte, 100)
	_ = copy(bytes, v[0:100])
	justString = string(bytes)

}

func createHugeString(i int64) string {
	var s string
	return s
}

func main() {
	someFunc()
}
