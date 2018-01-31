// This module contains a collection of YANG definitions
// common for all bgp operational data.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
package bgp_common_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package bgp_common_oper"))
}

// TcpFsmState represents TCP state machine states
type TcpFsmState string

const (
    // no connection
    TcpFsmState_closed TcpFsmState = "closed"

    // Waiting for a connection request from any remote TCP
    TcpFsmState_listen TcpFsmState = "listen"

    // Waiting for a matching connection request
    // after having sent a connection request
    TcpFsmState_synsent TcpFsmState = "synsent"

    // Waiting for a confirming connection request acknowledgment
    // after having both received and sent a connection request
    TcpFsmState_synrcvd TcpFsmState = "synrcvd"

    // connection established
    TcpFsmState_established TcpFsmState = "established"

    // Waiting for a connection termination request
    // from the remote TCP, or an acknowledgment of
    // the connection termination request previously sent
    TcpFsmState_finwait1 TcpFsmState = "finwait1"

    // Waiting for a connection termination request from the
    // remote TCP
    TcpFsmState_finwait2 TcpFsmState = "finwait2"

    // Waiting for a connection termination request from
    // the local use
    TcpFsmState_closewait TcpFsmState = "closewait"

    // Waiting for an acknowledgment of the connection termination
    // request previously sent to the remote TCP
    TcpFsmState_lastack TcpFsmState = "lastack"

    // Waiting for a connection termination request acknowledgment
    // from the remote
    TcpFsmState_closing TcpFsmState = "closing"

    // waiting for enough time to pass to be sure the remote TCP
    // received the acknowledgment of its connection termination request
    TcpFsmState_timewait TcpFsmState = "timewait"
)

// AfiSafi represents This identity represents IPv4 address family
type AfiSafi string

const (
    // IPv4 MDT address family
    AfiSafi_ipv4_mdt AfiSafi = "ipv4-mdt"

    // IPv4 Multicast address family
    AfiSafi_ipv4_multicast AfiSafi = "ipv4-multicast"

    // IPv4 Unicast address family
    AfiSafi_ipv4_unicast AfiSafi = "ipv4-unicast"

    // IPV4 MVPN address family
    AfiSafi_ipv4_mvpn AfiSafi = "ipv4-mvpn"

    // IPv4 Flowspec address family
    AfiSafi_ipv4_flowspec AfiSafi = "ipv4-flowspec"

    // IPv6 Multicast address family
    AfiSafi_ipv6_multicast AfiSafi = "ipv6-multicast"

    // IPv6 Unicast address family
    AfiSafi_ipv6_unicast AfiSafi = "ipv6-unicast"

    // IPv6 MVPN address family
    AfiSafi_ipv6_mvpn AfiSafi = "ipv6-mvpn"

    // IPv6 Flowspec address family
    AfiSafi_ipv6_flowspec AfiSafi = "ipv6-flowspec"

    // L2VPN VPLS address family
    AfiSafi_l2vpn_vpls AfiSafi = "l2vpn-vpls"

    // L2VPN EVPN address family
    AfiSafi_l2vpn_e_vpn AfiSafi = "l2vpn-e-vpn"

    // NSAP Unicast address family
    AfiSafi_nsap_unicast AfiSafi = "nsap-unicast"

    // RT Filter Unicast address family
    AfiSafi_rtfilter_unicast AfiSafi = "rtfilter-unicast"

    // VPNv4 Multicast address family
    AfiSafi_vpnv4_multicast AfiSafi = "vpnv4-multicast"

    // VPNv4 Unicast address family
    AfiSafi_vpnv4_unicast AfiSafi = "vpnv4-unicast"

    // VPNv6 Unicast address family
    AfiSafi_vpnv6_unicast AfiSafi = "vpnv6-unicast"

    // VPNv6 Multicast address family
    AfiSafi_vpnv6_multicast AfiSafi = "vpnv6-multicast"

    // VPNv4 Flowspec address family
    AfiSafi_vpnv4_flowspec AfiSafi = "vpnv4-flowspec"

    // VPNv6 Flowspec address family
    AfiSafi_vpnv6_flowspec AfiSafi = "vpnv6-flowspec"
)

