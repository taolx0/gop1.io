package basename1

import (
	"fmt"
	"testing"
)

func TestBasename(t *testing.T) {
	s := basename("/go/src/main.go")
	fmt.Println(s)
}