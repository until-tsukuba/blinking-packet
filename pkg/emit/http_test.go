package emit

import (
	"errors"
	"testing"
)

func TestHttpRequest(t *testing.T) {
	err := EmitHttp("http://www.tsukuba.ac.jp/")
	if err != nil {
		t.Errorf("An error occured while http request")
	}
}

func TestInvalidUrl(t *testing.T) {
	err := EmitHttp(".invalid")
	var invalidValueError *InvalidValueError
	ok := errors.As(err, &invalidValueError)

	if !ok {
		t.Errorf("Failed validation of invalid url: %v", err)
	}
}
