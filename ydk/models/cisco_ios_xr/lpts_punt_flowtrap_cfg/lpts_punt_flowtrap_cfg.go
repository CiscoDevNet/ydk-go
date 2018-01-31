// This module contains a collection of YANG definitions
// for Cisco IOS-XR lpts-punt-flowtrap package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-lpts-lib-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package lpts_punt_flowtrap_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package lpts_punt_flowtrap_cfg"))
}

// LptsPuntFlowtrapProtoId represents Lpts punt flowtrap proto id
type LptsPuntFlowtrapProtoId string

const (
    // ARP
    LptsPuntFlowtrapProtoId_arp LptsPuntFlowtrapProtoId = "arp"

    // Internet Control Message Protocol
    LptsPuntFlowtrapProtoId_icmp LptsPuntFlowtrapProtoId = "icmp"

    // Dynamic Host Configuration Protocol
    LptsPuntFlowtrapProtoId_dhcp LptsPuntFlowtrapProtoId = "dhcp"

    // PPP over Ethernet
    LptsPuntFlowtrapProtoId_pppoe LptsPuntFlowtrapProtoId = "pppoe"

    // Point to point Protocol
    LptsPuntFlowtrapProtoId_ppp LptsPuntFlowtrapProtoId = "ppp"

    // Internet Gateway Message Protocol
    LptsPuntFlowtrapProtoId_igmp LptsPuntFlowtrapProtoId = "igmp"

    // IPv4
    LptsPuntFlowtrapProtoId_ipv4 LptsPuntFlowtrapProtoId = "ipv4"

    // Layer2 Tunneling Protocol
    LptsPuntFlowtrapProtoId_l2tp LptsPuntFlowtrapProtoId = "l2tp"

    // Unclassified Source
    LptsPuntFlowtrapProtoId_unclassified LptsPuntFlowtrapProtoId = "unclassified"

    // OSPF
    LptsPuntFlowtrapProtoId_ospf LptsPuntFlowtrapProtoId = "ospf"

    // BGP
    LptsPuntFlowtrapProtoId_bgp LptsPuntFlowtrapProtoId = "bgp"

    // All protocols
    LptsPuntFlowtrapProtoId_default_ LptsPuntFlowtrapProtoId = "default"
)

