package rcu

import (
	"sync/atomic"
	"unsafe"
)

// Data represents a thread-safe lock-free RCU value. The zero value is valid
// and represents a Data[T] with a nil value pointer.
type Data[T any] struct {
	value *T
}

// NewData returns a new Data[T] associated with the given value.
func NewData[T any](value T) *Data[T] {
	return &Data[T]{value: &value}
}

// SetValue sets the value associated with Data[T] to the given one.
func (d *Data[T]) SetValue(value T) {
	d.SetValuePtr(&value)
}

// SetValuePtr sets the value pointer associated with Data[T] to the given one.
func (d *Data[T]) SetValuePtr(valuePtr *T) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&d.value)),
		unsafe.Pointer(valuePtr))
}

// GetValue returns the current value associated with Data[T]. It the internal
// value pointer is nil, returns the zero value of T.
func (d *Data[T]) GetValue() T {
	valuePtr := d.GetValuePtr()
	if valuePtr == nil {
		var value T
		return value
	}

	return *valuePtr
}

// GetValuePtr returns the current value pointer associated with Data[T].
func (d *Data[T]) GetValuePtr() *T {
	return (*T)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(
		&d.value))))
}
