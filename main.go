package main

import "fmt"

func main() {
	input := "Here's my spammy page: https://hehefouls.netHAHAHA see you."
	output := spam(input)

	fmt.Println(output)
}

func spam(msg string) string {

	buffer := []byte(msg)

	for i := 0; i < len(buffer); i++ {
		if i+8 <= len(buffer) && string(buffer[i:i+8]) == "https://" {
			x := i + 8
			for x < len(buffer) && buffer[x] != ' ' {
				buffer[x] = '*'
				x++
			}
			i = x
		}
	}
	return string(buffer)
}
