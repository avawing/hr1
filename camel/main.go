package main

import "fmt"

func main() {
	var input string
	fmt.Scanf("%s\n", &input)

	answer := 1

	for _, chr := range input {
		//ascii value! NEAT!
		if chr <= 90 && chr >= 65 {
			answer += 1
		}
	}
	fmt.Println(answer)
}
