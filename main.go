package main

import "fmt"

func main() {
	input := "Here's my spammy page: http://hehefouls.netHAHAHA see you."
	output := spam(input)

	fmt.Println(output)
}
func spam(msg string) string {

	buffer := []byte(msg)

	for i := 0; i < len(buffer); i++ {
		if i+7 <= len(buffer) && string(buffer[i:i+7]) == "http://" {
			x := i + 7
			for x < len(buffer) && buffer[x] != ' ' {
				buffer[x] = '*'
				x++
			}
			i = x
		}
	}
	return string(buffer)
}
