package validators

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func IsSuspiciousIPv4(ipv4 string) bool {
	req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/stamparm/ipsum/master/ipsum.txt", nil)
	if err != nil {
		panic(errors.Wrap(err, "IsSuspiciousIPv4"))
	}

	resp, err := (&http.Client{
		Timeout: 10 * time.Second,
	}).Do(req)
	if err != nil {
		panic(errors.Wrap(err, "IsSuspiciousIPv4"))
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err := resp.Body.Close(); err != nil {
		panic(errors.Wrap(err, "IsSuspiciousIPv4"))
	}

	lines := strings.Split(string(b), "\n")
	for _, l := range lines {
		if strings.HasPrefix(l, "#") {
			continue
		}

		ip := strings.Split(l, "\t")
		if len(ip) != 2 {
			continue
		}

	}

	return false
}

// TODO add IPv6 support
func IsSuspiciousIPv6(address string) bool {
	panic("validate IPv6 is not supported")
}
