package exercism

import "fmt"

func PascalsT(num int) string {
	if num < 1 {
		return ""
	}
	out := make([][]int, num)
	out[0] = []int{1}
	for i := 0; i < num; i++ {
		out[i] = make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				out[i][j] = 1
			} else {
				out[i][j] = out[i-1][j] + out[i-1][j-1]
			}
		}
	}
	res := ""
	for _, a := range out {
		for _, b := range a {
			res += fmt.Sprint(b) + " "
		}
		res += "\n"
	}
	return res
}
