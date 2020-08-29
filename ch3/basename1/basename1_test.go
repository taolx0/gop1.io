package basename1

import (
	"fmt"
	"testing"
)

func TestBasename1(t *testing.T) {
	s := basename1("/go/src/main.go")
	fmt.Println(s)
}
