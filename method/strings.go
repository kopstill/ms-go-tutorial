package method

import "strings"

type upperstring string

func (us upperstring) toUpper() string {
	return strings.ToUpper(string(us))
}
