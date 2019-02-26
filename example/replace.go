package example

import (
	"bytes"
	"fmt"
	"strings"
)

func stringsReplaceUseNegative() {
	text := "aaabbbcccaaabbbccc"
	strings.Replace(text, "aaa", "eee", -1) // want "detect strings.Replace(...) use negative value -1"
	fmt.Println(text)
}

func stringsReplacePositive() {
	text := "aaabbbcccaaabbbccc"
	strings.Replace(text, "aaa", "eee", 2)
	fmt.Println(text)
}

func bytesReplaceUseNegative() {
	b := []byte("aaabbbcccaaabbbccc")
	bytes.Replace(b, []byte("aaa"), []byte("eee"), -1) // want "detect bytes.Replace(...) use negative value -1"
	fmt.Println(b)
}

func bytesReplacePositive() {
	b := []byte("aaabbbcccaaabbbccc")
	bytes.Replace(b, []byte("aaa"), []byte("eee"), 3)
	fmt.Println(b)
}

func Replace(s string, old string, new string, n int) {
	fmt.Println("myreplace")
}

func myReplace() {
	text := "aaabbbcccaaabbbccc"
	Replace(text, "aaa", "eee", -1)
	fmt.Println(text)
}
