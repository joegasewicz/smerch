package smerch

type Locker struct {
	LockPath string
}

// NewLocker
// e.g smerch.lock
func NewLocker(lock string) *Locker {
	return &Locker{
		LockPath: lock,
	}
}
