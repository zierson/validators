package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsDisposableEmailProvider(t *testing.T) {
	assert.True(t, IsDisposableEmailProvider("disposableaddress.com"))
	assert.True(t, IsDisposableEmailProvider("nokiamail.com"))
	assert.False(t, IsDisposableEmailProvider("gmail.com"))
	assert.False(t, IsDisposableEmailProvider("fastmail.com"))
}

func TestIsDisposableEmail(t *testing.T) {
	val, err := IsDisposableEmail("test@nokiamail.com")
	if assert.NoError(t, err) {
		assert.True(t, val)
	}

	val, err = IsDisposableEmail("test@fastmail.com")
	if assert.NoError(t, err) {
		assert.False(t, val)
	}
}
