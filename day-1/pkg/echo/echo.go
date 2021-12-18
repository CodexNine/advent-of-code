package echo

import (
	"fmt"
)

func Shout(input string) (output string) {
	output = fmt.Sprintf("%v", input)
	return output
}