package emit

import "github.com/go-ping/ping"

func EmitPing(value string) error {
	pinger, err := ping.NewPinger(value)
	if err != nil {
		return &InvalidValueError{Value: value, Message: "Invalid host"}
	}

	go func() {
		err = pinger.Run()
	}()

	return nil
}
