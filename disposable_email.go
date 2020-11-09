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
	mux_disp   sync.Mutex
	cache_disp []string
)

func IsDisposableEmailProvider(domain string) bool {
	mux_disp.Lock()

	if len(cache_disp) == 0 {
		req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/martenson/disposable-email-domains/master/disposable_email_blocklist.conf", nil)
		if err != nil {
			panic(errors.Wrap(err, "IsDisposableEmail"))
		}

		resp, err := (&http.Client{
			Timeout: 10 * time.Second,
		}).Do(req)
		if err != nil {
			panic(errors.Wrap(err, "IsDisposableEmail"))
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err := resp.Body.Close(); err != nil {
			panic(errors.Wrap(err, "IsDisposableEmail"))
		}

		lines := strings.Split(string(b), "\n")
		for _, l := range lines {
			if len(strings.Split(l, ".")) == 0 {
				panic(errors.Wrap(errors.New("invalid domain name"), "IsDisposableEmail"))
			}

			cache_disp = append(cache_disp, strings.TrimSpace(l))
		}
	}

	mux_disp.Unlock()

	for _, v := range cache_disp {
		if v == domain {
			return true
		}
	}

	return false
}
