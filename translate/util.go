package translate

import (
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	"github.com/robertkrimen/otto"
)

func encodeURI(s string) (string, error) {
	eUri := `eUri = encodeURI(sourceText);`
	vm := otto.New()
	err := vm.Set("sourceText", s)
	if err != nil {
		return "err", errors.New("Error setting js variable")
	}
	_, err = vm.Run(eUri)
	if err != nil {
		return "err", errors.New("Error executing jscript")
	}
	val, err := vm.Get("eUri")
	if err != nil {
		return "err", errors.New("Error getting variable value from js")
	}
	v, err := val.ToString()
	if err != nil {
		return "err", errors.New("Error converting js var to string")
	}
	return v, nil
}

func httpGet(url string) ([]byte, error) {
	r, err := http.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "httpGet. url:%s", url)
	}

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "reading response body. url:%s", url)
	}

	return body, nil
}