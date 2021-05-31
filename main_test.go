package main

import (
	"io/fs"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/nyanpassu/golang-test/mocks"
)

// TestDummy .
func TestDummy(t *testing.T) {
	os := mocks.OS{}
	os.On("Stat", mock.Anything).Return(func(string) fs.FileInfo {
		return FileInfo{}
	}, func(string) error {
		return nil
	})
	info, err := os.Stat("")
	assert.NotNil(t, info)
	assert.NoError(t, err)
}
