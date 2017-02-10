package utils

import (
	"fmt"
	"strings"
	"testing"
)

func TestRead2String(t *testing.T) {
	reader := strings.NewReader("abcdefghij")
	s, _ := Read2String(reader)
	fmt.Println(s)
}
