package basename1

import (
	"fmt"
	"testing"
)

func TestBasename(t *testing.T) {
	s := basename("/go/src/eval.go")
	fmt.Println(s)
}
