package validators

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSuspiciousIPv4(t *testing.T) {
	assert.True(t, IsSuspiciousIPv4("54.38.81.231"))
	assert.True(t, IsSuspiciousIPv4("185.220.102.248"))
	assert.False(t, IsSuspiciousIPv4("1.1.1.1"))
	assert.False(t, IsSuspiciousIPv4("130.26.118.215"))
}

func TestIsSuspiciousIPv6(t *testing.T) {
	assert.Panics(t, func() { IsSuspiciousIPv6("::1") })
}
