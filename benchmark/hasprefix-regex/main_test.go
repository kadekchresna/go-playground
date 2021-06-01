package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

var t = "12.099.999.9-111-929"

func BenchmarkReplace(b *testing.B) {
	fmt.Println(strings.HasPrefix("DIST213", "DIST"))
}
func BenchmarkReplaceRegexp(b *testing.B) {
	re := regexp.MustCompile(`^(DIST)`)
	fmt.Println(re.Match([]byte("DIST213")))
}
