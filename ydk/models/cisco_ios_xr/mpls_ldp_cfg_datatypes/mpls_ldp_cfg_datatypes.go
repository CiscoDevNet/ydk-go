// This module contains a collection of YANG definitions
// for Cisco IOS-XR mpls-ldp-cfg-datat package configuration.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package mpls_ldp_cfg_datatypes

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_ldp_cfg_datatypes"))
}

// MplsLdpNbrPassword represents Mpls ldp nbr password
type MplsLdpNbrPassword string

const (
    // Disable the global default password for this
    // neighbor
    MplsLdpNbrPassword_disable MplsLdpNbrPassword = "disable"

    // Specify a password for this neighbor
    MplsLdpNbrPassword_specified MplsLdpNbrPassword = "specified"
)

// MplsLdpDownstreamOnDemand represents Mpls ldp downstream on demand
type MplsLdpDownstreamOnDemand string

const (
    // Downstream on Demand peers permitted by ACL
    MplsLdpDownstreamOnDemand_peer_acl MplsLdpDownstreamOnDemand = "peer-acl"
)

// MplsLdpRouterId represents Mpls ldp router id
type MplsLdpRouterId string

const (
    // Use given IP address as LDP Router ID
    MplsLdpRouterId_address MplsLdpRouterId = "address"
)

// MplsLdpafName represents Mpls ldpaf name
type MplsLdpafName string

const (
    // IPv4
    MplsLdpafName_ipv4 MplsLdpafName = "ipv4"

    // IPv6
    MplsLdpafName_ipv6 MplsLdpafName = "ipv6"
)

// MplsLdpSessionProtection represents Mpls ldp session protection
type MplsLdpSessionProtection string

const (
    // Protect all peer sessions
    MplsLdpSessionProtection_all MplsLdpSessionProtection = "all"

    // Protect peer session(s) permitted by peer ACL
    MplsLdpSessionProtection_for_ MplsLdpSessionProtection = "for"

    // Protect all peer sessions and holdup protection
    // for given duration
    MplsLdpSessionProtection_all_with_duration MplsLdpSessionProtection = "all-with-duration"

    // Protect peer session(s) permitted by peer ACL
    // and holdup protection for given duration
    MplsLdpSessionProtection_for_with_duration MplsLdpSessionProtection = "for-with-duration"

    // Protect all peer sessions and holdup protection
    // forever
    MplsLdpSessionProtection_all_with_forever MplsLdpSessionProtection = "all-with-forever"

    // Protect peer session(s) permitted by peer ACL
    // and holdup protection forever
    MplsLdpSessionProtection_for_with_forever MplsLdpSessionProtection = "for-with-forever"
)

