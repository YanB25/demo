package add

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
}

func TestCatSay(t *testing.T) {
	cat := Cat{}
	s := cat.Say()
	if s != "cat" {
		e := fmt.Sprintf("Err. cat.Say() should return \"cat\" but not \"%s\"", s)
		t.Error(e)
	}
}
