package emit

func EmitPacket(packetType string, value string) error {
	switch packetType {
	case "ping":
		return EmitPing(value)
	case "http":
		return EmitHttp(value)
	case "dns":
		return EmitDns(value)
	default:
		return &InvalidTypeError{PacketType: packetType}
	}
}
