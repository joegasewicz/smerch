package smerch_test

import (
	"testing"

	"github.com/joegasewicz/smerch"
	"github.com/stretchr/testify/assert"
)

func TestLocker_IsLocked(t *testing.T) {
	locker := smerch.NewLocker("examples/smerch.lock")
	result := locker.IsLocked()
	assert.Equal(t, result, true)
}
