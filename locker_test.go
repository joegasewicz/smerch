package smerch_test

import (
	"testing"

	"github.com/joegasewicz/smerch"
	"github.com/stretchr/testify/assert"
)

func TestLocker_IsLocked(t *testing.T) {
	pyprojectData := map[string]any{}
	locker := smerch.NewLocker("examples/smerch.lock", &pyprojectData)
	result := locker.IsLocked()
	assert.Equal(t, result, true)
}
