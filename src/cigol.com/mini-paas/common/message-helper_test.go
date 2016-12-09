package common

import (
	"fmt"
	"testing"
)

func TestMessageHelper(t *testing.T) {
	ret := WrapMessage("Result", true, "Message", "good")
	fmt.Println(ret)
}
