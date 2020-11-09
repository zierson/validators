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
