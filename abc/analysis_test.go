package abc

import (
	"testing"
)


func TestABC(t *testing.T) {

	instance := New()
	t.Log(instance.Calculate())

}
