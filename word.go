package word

import (
	"fmt"
	"strings"
)

func Wordcount(s string) map[string]int {
	fmt.Println("Ning")
	str := strings.Split(s, " ")
	data := make(map[string]int)
	for _, word := range str {
		data[word] += 1
	}
	return data
}
