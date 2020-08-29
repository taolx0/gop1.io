package basename2

import (
	"fmt"
	"testing"
)

func TestBasename(t *testing.T) {
	s := basename("/go/src/data.go")
	fmt.Println(s)
}
