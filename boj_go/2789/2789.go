package main

import "fmt"

func main() {
	var a string
	fmt.Scanln(&a)

	exceptCambridge := func(str string) string {
		var res string
		for i := 0; i < len(str); i++ {
			switch str[i] {
			case 'C':
			case 'A':
			case 'M':
			case 'B':
			case 'R':
			case 'I':
			case 'D':
			case 'G':
			case 'E':
			default:
				res += string(str[i])
			}
		}
		return res
	}
	fmt.Println(exceptCambridge(a))
}
