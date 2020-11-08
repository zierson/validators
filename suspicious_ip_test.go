package validators

import "testing"
import "github.com/stretchr/testify/assert"

func TestIsSuspiciousIPv4(t *testing.T) {
	// success
	assert.False(t, IsSuspiciousIPv4("1.1.1.1"))
	assert.False(t, IsSuspiciousIPv4("130.26.118.215"))

	// failure
	assert.True(t, IsSuspiciousIPv4("51.77.58.144"))
	assert.True(t, IsSuspiciousIPv4("185.220.102.250"))
}

func TestIsSuspiciousIPv6(t *testing.T) {
	// success
	assert.Panics(t, func() { IsSuspiciousIPv6("::1") })
}
