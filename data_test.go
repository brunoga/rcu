package rcu

import (
	"testing"
)

func TestData_Simple(t *testing.T) {
	value0 := int(0)
	data := NewData(&value0)
	if *data.GetValue() != 0 {
		t.Error("*data.GetValue() != 0")
	}

	value1 := int(1)
	data.SetValue(&value1)
	if *data.GetValue() != 1 {
		t.Error("*data.GetValue() != 1")
	}
}
