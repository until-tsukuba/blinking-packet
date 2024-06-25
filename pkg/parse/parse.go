package parse

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"net"
	"strings"
)

func ProtocolToColor(protocol string) (int, int, int) {
	protocol = strings.ToUpper(protocol)
	switch protocol {
	case "ETHERNET":
		// brown
		return 3, 2, 0
	case "ARP":
		// pink
		return 3, 1, 3
	case "IP":
		// cyan
		return 1, 3, 3
	case "ICMP":
		// white
		return 2, 2, 2
	case "TCP":
		// green
		return 1, 3, 1
	case "UDP":
		// red
		return 3, 1, 1
	case "DNS":
		// orange
		return 4, 2, 0
	case "HTTP":
		// blue
		return 1, 1, 3
	default:
		// yellow
		return 3, 3, 1
	}
}

type ParsedPacket struct {
	LayerList []Layer
	Reversed  bool
}

type Layer struct {
	ProtocolName string
}

func Parse(packet gopacket.Packet, interfaceInfo net.Interface) ParsedPacket {
	var layerList = make([]Layer, 0)
	var reversed = true

	if ethLayer := packet.Layer(layers.LayerTypeEthernet); ethLayer != nil {
		layerList = append(layerList, Layer{ProtocolName: "ETHERNET"})
		source := ethLayer.(*layers.Ethernet).SrcMAC
		interfaceMac := interfaceInfo.HardwareAddr
		if source.String() == interfaceMac.String() {
			reversed = false
		}
	}

	if arpLayer := packet.Layer(layers.LayerTypeARP); arpLayer != nil {
		layerList = append(layerList, Layer{ProtocolName: "ARP"})
	}

	if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
		layerList = append(layerList, Layer{ProtocolName: "IP"})
	}

	if icmpLayer := packet.Layer(layers.LayerTypeICMPv4); icmpLayer != nil {
		layerList = append(layerList, Layer{ProtocolName: "ICMP"})
	}

	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		layerList = append(layerList, Layer{ProtocolName: "TCP"})
		dstPort := tcpLayer.(*layers.TCP).DstPort
		if dstPort == 80 {
			layerList = append(layerList, Layer{ProtocolName: "HTTP"})
		}
		if dstPort == 443 {
			layerList = append(layerList, Layer{ProtocolName: "HTTPS"})
		}
	}

	if udpLayer := packet.Layer(layers.LayerTypeUDP); udpLayer != nil {
		layerList = append(layerList, Layer{ProtocolName: "UDP"})
	}

	if dnsLayer := packet.Layer(layers.LayerTypeDNS); dnsLayer != nil {
		layerList = append(layerList, Layer{ProtocolName: "DNS"})
	}

	return ParsedPacket{LayerList: layerList, Reversed: reversed}
}
