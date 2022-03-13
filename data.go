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
func NewData[T any](value *T) Data[T] {
	return Data[T]{value: value}
}

// SetValue sets the value pointer associated with Data[T] to the given one.
func (d *Data[T]) SetValue(value *T) {
	atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&d.value)),
		unsafe.Pointer(value))
}

// SetIfNilValue sets the value pointer associated with Data[T] to the given
// value iff the current value pointer is nil. Returns true if the value was
// set and false otherwise.
func (d *Data[T]) SetIfNilValue(value *T) bool {
	return atomic.CompareAndSwapPointer((*unsafe.Pointer)(unsafe.Pointer(&d.value)),
		unsafe.Pointer(nil), unsafe.Pointer(value))
}

// GetValue returns the current value pointer associated with Data[T].
func (d *Data[T]) GetValue() *T {
	return (*T)(atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(
		&d.value))))
}
