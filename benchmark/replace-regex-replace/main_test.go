package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

var t = "12.099.999.9-111-929"

func BenchmarkReplace(b *testing.B) {
	fmt.Println(strings.ReplaceAll(strings.ReplaceAll(t, "-", ""), ".", ""))
}
func BenchmarkReplaceRegexp(b *testing.B) {
	re := regexp.MustCompile(`[\.\-]`)
	fmt.Println(re.ReplaceAllString(t, ""))
}
