package validators

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

var (
	mux_susp   sync.Mutex
	cache_susp []string
)

func IsSuspiciousIPv4(ipv4 string) bool {
	mux_susp.Lock()

	if len(cache_susp) == 0 {
		// Backup list file:
		// https://raw.githubusercontent.com/zierson/validators/master/data/suspicious_ip.txt
		req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/stamparm/ipsum/master/suspicious_ip.txt", nil)
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

			cache_susp = append(cache_susp, strings.TrimSpace(ip[0]))
		}
	}

	mux_susp.Unlock()

	for _, v := range cache_susp {
		if v == ipv4 {
			return true
		}
	}

	return false
}

// TODO add IPv6 support
func IsSuspiciousIPv6(ipv6 string) bool {
	panic("validate IPv6 is not supported")
}
