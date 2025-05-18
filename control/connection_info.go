package control

type ConnectionInfo struct {
	Src      string
	Dst      string
	Sniffed  string
	Dscp     int
	L4proto  string
	Dialer   string
	Outbound string
	Mac      string
	Pname    string
}
