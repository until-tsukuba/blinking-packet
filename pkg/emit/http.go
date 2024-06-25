package emit

import (
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

func EmitHttp(value string) error {
	_, err := url.ParseRequestURI(value)
	if err != nil {
		return &InvalidValueError{Value: value, Message: "Invalid URL"}
	}

	slog.Debug("Emit HTTP GET request", "value", value)

	go func() {
		resp, err := http.Get(value)
		if err != nil {
			return
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)

		_, _ = io.ReadAll(resp.Body)
	}()

	slog.Debug("Succeed HTTP request", "value", value)

	return nil
}
