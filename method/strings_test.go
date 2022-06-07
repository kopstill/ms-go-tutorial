package method

import (
	"fmt"
	"testing"
)

func TestToUpper(t *testing.T) {
	us := upperstring("Learning Go!")
	fmt.Println(us)
	fmt.Println(us.toUpper())
}
