package control

import "net/netip"

type ConnectionInfo struct {
	Src      netip.AddrPort
	Dst      string
	Sniffed  string
	Dscp     uint8
	L4proto  string
	Dialer   string
	Outbound string
	Mac      string
	Pname    string
}
