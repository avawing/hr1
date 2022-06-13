package main

import "fmt"

func main() {
	var delta int
	var input string

	fmt.Scanf("%d\n", &delta)
	fmt.Scanf("%s\n", &input)

	var ret []rune
	for _, chr := range input {
		ret = append(ret, rotateAny(chr, delta))
	}

	fmt.Println(string(ret))
	return
}

func cipher(r rune, base, delta int) rune {
	tmp := int(r) - base
	tmp = (tmp + delta) % 26
	return rune(tmp + base)
}

func rotateAny(r rune, delta int) rune {
	if r >= 'A' && r <= 'Z' {
		return cipher(r, 'A', delta)
	}
	if r >= 'a' && r <= 'z' {
		return cipher(r, 'a', delta)
	}
	return r
}
