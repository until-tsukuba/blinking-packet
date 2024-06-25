package emit

import (
	"errors"
	"testing"
)

func TestEchoRequestWithIpv4(t *testing.T) {
	err := EmitPing("8.8.8.8")
	if err != nil {
		t.Errorf("An error occured while send echo request")
	}
}

func TestEchoRequestWithIpv6(t *testing.T) {
	err := EmitPing("2001:4860:4860::8888")
	if err != nil {
		t.Errorf("An error occured while send echo request")
	}
}

func TestEchoRequestWithDomainName(t *testing.T) {
	err := EmitPing("www.tsukuba.ac.jp")
	if err != nil {
		t.Errorf("An error occured while send echo request")
	}
}

func TestInvalidHost(t *testing.T) {
	err := EmitPing(".invalid")
	var invalidValueError *InvalidValueError
	ok := errors.As(err, &invalidValueError)

	if !ok {
		t.Errorf("Failed validation of invalid host: %v", err)
	}
}
