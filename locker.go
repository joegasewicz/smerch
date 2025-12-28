package smerch

import (
	"errors"
	"os"
)

type Locker struct {
	LockPath      string // e.g smerch.lock
	PyprojectData *map[string]any
}

// NewLocker
func NewLocker(lock string, pyprojectData *map[string]any) *Locker {
	return &Locker{
		LockPath:      lock,
		PyprojectData: pyprojectData,
	}
}

// IsLocked Check if the lock file exists.
func (l *Locker) IsLocked() bool {
	if _, err := os.Stat(l.LockPath); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
