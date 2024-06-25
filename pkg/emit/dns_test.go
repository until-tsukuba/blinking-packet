package emit

import (
	"errors"
	"testing"
)

func TestExchange(t *testing.T) {
	err := EmitDns("www.tsukuba.ac.jp")
	if err != nil {
		t.Errorf("An error occured while exchange dns")
	}
}

func TestInvalidDomainName(t *testing.T) {
	err := EmitDns(".invalid")
	var invalidValueError *InvalidValueError
	ok := errors.As(err, &invalidValueError)

	if !ok {
		t.Errorf("Failed validation of invalid domain name: %v", err)
	}
}
