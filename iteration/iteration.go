package iteration

func Repeat(s string, num int) (res string) {
	for i := 0; i < num; i++ {
		res += s
	}
	return
}
