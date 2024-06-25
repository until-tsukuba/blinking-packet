package emit

import "fmt"

type InvalidTypeError struct {
	PacketType string
}

func (e *InvalidTypeError) Error() string {
	return fmt.Sprintf("Invalid packet type: %s", e.PacketType)
}

type InvalidValueError struct {
	Value   string
	Message string
}

func (e *InvalidValueError) Error() string {
	return fmt.Sprintf("Invalid value: %s / %s", e.Value, e.Message)
}
