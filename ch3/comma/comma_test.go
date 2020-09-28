package comma

import (
	"fmt"
	"testing"
)

func Test_comma(t *testing.T) {
	s := comma("12348797345355")
	fmt.Printf("%s\n", s)
}
