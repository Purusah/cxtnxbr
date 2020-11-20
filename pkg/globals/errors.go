package globals

import "fmt"

type KeyNotFound struct {
	key string
}

func (e KeyNotFound) Error() string {
	return fmt.Sprint("key not found ", e.key)
}

type BrokenKey struct {
	key string
}

func (e BrokenKey) Error() string {
	return fmt.Sprint("key broken ", e.key)
}

type StorageError struct{}

func (e StorageError) Error() string {
	return "storage problem"
}
