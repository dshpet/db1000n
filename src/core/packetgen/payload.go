package packetgen

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"

	"github.com/Arriven/db1000n/src/utils"
)

func BuildPayload(c LayerConfig) (gopacket.Layer, error) {
	switch c.Type {
	case "":
		return nil, nil
	case "raw":
		var packetConfig struct {
			Payload string
		}

		if err := utils.Decode(c.Data, &packetConfig); err != nil {
			return nil, err
		}

		return gopacket.Payload([]byte(packetConfig.Payload)), nil
	case "icmpv4":
		var packetConfig ICMPV4PacketConfig
		if err := utils.Decode(c.Data, &packetConfig); err != nil {
			return nil, err
		}

		return buildICMPV4Packet(packetConfig), nil
	case "dns":
		var packetConfig DNSPacketConfig
		if err := utils.Decode(c.Data, &packetConfig); err != nil {
			return nil, err
		}

		return buildDNSPacket(packetConfig), nil
	default:
		return nil, fmt.Errorf("unsupported layer type %s", c.Type)
	}
}

type ICMPV4PacketConfig struct {
	TypeCode uint16 `mapstructure:"code"`
	ID       uint16
	Seq      uint16
}

func buildICMPV4Packet(c ICMPV4PacketConfig) *layers.ICMPv4 {
	return &layers.ICMPv4{
		TypeCode: layers.ICMPv4TypeCode(c.TypeCode),
		Id:       c.ID,
		Seq:      c.Seq,
	}
}

type DNSPacketConfig struct {
	ID      uint16
	Qr      bool
	OpCode  uint8 `mapstructure:"code"`
	QDCount uint16
}

func buildDNSPacket(c DNSPacketConfig) *layers.DNS {
	return &layers.DNS{
		ID:      c.ID,
		QR:      c.Qr,
		OpCode:  layers.DNSOpCode(c.OpCode),
		QDCount: c.QDCount,
	}
}
