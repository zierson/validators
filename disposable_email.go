package validators

import (
	"github.com/asaskevich/govalidator"
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
		// Backup list file:
		// https://raw.githubusercontent.com/zierson/validators/master/data/disposable_email_list.txt
		req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/martenson/disposable-email-domains/master/disposable_email_list.txt", nil)
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
				panic(errors.WithMessage(errors.New("invalid domain name"), "IsDisposableEmail"))
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

func IsDisposableEmail(email string) (bool, error) {
	if !govalidator.IsEmail(email) {
		return false, errors.WithMessage(errors.New("invalid email format"), "IsDisposableEmail")
	}

	e := strings.Split(email, "@")
	if len(e) != 2 {
		return false, errors.WithMessage(errors.New("invalid email format"), "IsDisposableEmail")
	}

	if IsDisposableEmailProvider(e[1]) {
		return true, nil
	}

	return false, nil
}
