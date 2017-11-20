package faker

import (
	"testing"
)

// get random 10 male's name
func TestMale(t *testing.T) {
	names := make(map[string]bool)
	for i := 0; i < 10; i++ {
		name := male()
		if _, ok := names[name]; ok {
			t.Errorf("the males() function return the duplicate value or previous generated value ")
		} else {
			names[name] = true
		}
	}

}
