// This module defines common types for use in models requiring
// data definitions related to packet matches.
package packet_match_types

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package packet_match_types"))
}

type TCPACK struct {
}

func (id TCPACK) String() string {
	return "openconfig-packet-match-types:TCP_ACK"
}

type IPPROTOCOL struct {
}

func (id IPPROTOCOL) String() string {
	return "openconfig-packet-match-types:IP_PROTOCOL"
}

type IPUDP struct {
}

func (id IPUDP) String() string {
	return "openconfig-packet-match-types:IP_UDP"
}

type ETHERTYPEARP struct {
}

func (id ETHERTYPEARP) String() string {
	return "openconfig-packet-match-types:ETHERTYPE_ARP"
}

type TCPSYN struct {
}

func (id TCPSYN) String() string {
	return "openconfig-packet-match-types:TCP_SYN"
}

type ETHERTYPEVLAN struct {
}

func (id ETHERTYPEVLAN) String() string {
	return "openconfig-packet-match-types:ETHERTYPE_VLAN"
}

type TCPECE struct {
}

func (id TCPECE) String() string {
	return "openconfig-packet-match-types:TCP_ECE"
}

type IPICMP struct {
}

func (id IPICMP) String() string {
	return "openconfig-packet-match-types:IP_ICMP"
}

type TCPFIN struct {
}

func (id TCPFIN) String() string {
	return "openconfig-packet-match-types:TCP_FIN"
}

type ETHERTYPEROCE struct {
}

func (id ETHERTYPEROCE) String() string {
	return "openconfig-packet-match-types:ETHERTYPE_ROCE"
}

type IPPIM struct {
}

func (id IPPIM) String() string {
	return "openconfig-packet-match-types:IP_PIM"
}

type TCPFLAGS struct {
}

func (id TCPFLAGS) String() string {
	return "openconfig-packet-match-types:TCP_FLAGS"
}

type TCPRST struct {
}

func (id TCPRST) String() string {
	return "openconfig-packet-match-types:TCP_RST"
}

type ETHERTYPE struct {
}

func (id ETHERTYPE) String() string {
	return "openconfig-packet-match-types:ETHERTYPE"
}

type ETHERTYPEIPV4 struct {
}

func (id ETHERTYPEIPV4) String() string {
	return "openconfig-packet-match-types:ETHERTYPE_IPV4"
}

type IPTCP struct {
}

func (id IPTCP) String() string {
	return "openconfig-packet-match-types:IP_TCP"
}

type ETHERTYPEIPV6 struct {
}

func (id ETHERTYPEIPV6) String() string {
	return "openconfig-packet-match-types:ETHERTYPE_IPV6"
}

type TCPURG struct {
}

func (id TCPURG) String() string {
	return "openconfig-packet-match-types:TCP_URG"
}

type IPRSVP struct {
}

func (id IPRSVP) String() string {
	return "openconfig-packet-match-types:IP_RSVP"
}

type ETHERTYPEMPLS struct {
}

func (id ETHERTYPEMPLS) String() string {
	return "openconfig-packet-match-types:ETHERTYPE_MPLS"
}

type TCPPSH struct {
}

func (id TCPPSH) String() string {
	return "openconfig-packet-match-types:TCP_PSH"
}

type IPAUTH struct {
}

func (id IPAUTH) String() string {
	return "openconfig-packet-match-types:IP_AUTH"
}

type IPGRE struct {
}

func (id IPGRE) String() string {
	return "openconfig-packet-match-types:IP_GRE"
}

type IPIGMP struct {
}

func (id IPIGMP) String() string {
	return "openconfig-packet-match-types:IP_IGMP"
}

type ETHERTYPELLDP struct {
}

func (id ETHERTYPELLDP) String() string {
	return "openconfig-packet-match-types:ETHERTYPE_LLDP"
}

type TCPCWR struct {
}

func (id TCPCWR) String() string {
	return "openconfig-packet-match-types:TCP_CWR"
}

type IPL2TP struct {
}

func (id IPL2TP) String() string {
	return "openconfig-packet-match-types:IP_L2TP"
}

// PortNumRange represents indicate a wildcard.
type PortNumRange string

const (
    // Indicates any valid port number (e.g., wildcard)
    PortNumRange_ANY PortNumRange = "ANY"
)

