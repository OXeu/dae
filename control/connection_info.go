package control

type ConnectionInfo struct {
    Src      string
    Dst      string
    Sniffed  string
    Dscp     uint8
    L4proto  string
    Dialer   string
    Outbound string
    Mac      string
    Pname    string
}
