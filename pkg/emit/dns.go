package emit

import (
	"github.com/miekg/dns"
	"log/slog"
)

func EmitDns(value string) error {
	_, valid := dns.IsDomainName(value)
	if !valid {
		return &InvalidValueError{Value: value, Message: "Invalid domain name"}
	}

	slog.Debug("Emit DNS packet", "value", value)

	go func() {
		msg := new(dns.Msg)
		msg.SetQuestion(value, dns.TypeA)

		ans, err := dns.Exchange(msg, "8.8.8.8:53")
		if err == nil {
			slog.Debug("DNS Exchanged", "msg", ans)
		} else {
			slog.Error("DNS Exchange error", err)
		}
	}()

	return nil
}
