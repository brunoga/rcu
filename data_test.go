package rcu

import (
	"testing"
)

func TestData_Simple(t *testing.T) {
	data := NewData(0)
	if data.GetValue() != 0 {
		t.Error("data.GetValue() != 0")
	}

	data.SetValue(1)
	if data.GetValue() != 1 {
		t.Error("data.GetValue() != 1")
	}

	value := 2
	data.SetValuePtr(&value)
	if data.GetValue() != 2 {
		t.Error("data.GetValue() != 2")
	}
	valuePtr := data.GetValuePtr()
	if valuePtr != &value {
		t.Error("valuePtr != &value")
	}

	data = &Data[int]{(*int)(nil)}
	if data.GetValue() != 0 {
		t.Error("data.GetValue() != 0")
	}
}
