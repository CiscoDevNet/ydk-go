// This module contains a collection of YANG definitions
// for Cisco IOS-XR snmp-agent package configuration.
// 
// This module contains definitions
// for the following management objects:
//   snmp: The heirarchy point for all the SNMP configurations
//   mib: mib
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package snmp_agent_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmp_agent_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-agent-cfg snmp}", reflect.TypeOf(Snmp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-agent-cfg:snmp", reflect.TypeOf(Snmp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-snmp-agent-cfg mib}", reflect.TypeOf(Mib{}))
    ydk.RegisterEntity("Cisco-IOS-XR-snmp-agent-cfg:mib", reflect.TypeOf(Mib{}))
}

// SnmpHashAlgorithm represents Snmp hash algorithm
type SnmpHashAlgorithm string

const (
    // No authentication required
    SnmpHashAlgorithm_none SnmpHashAlgorithm = "none"

    // Standard Message Digest algorithm
    SnmpHashAlgorithm_md5 SnmpHashAlgorithm = "md5"

    // SHA algorithm
    SnmpHashAlgorithm_sha SnmpHashAlgorithm = "sha"
)

// SnmpAccessLevel represents Snmp access level
type SnmpAccessLevel string

const (
    // Read Only Access for a community string
    SnmpAccessLevel_read_only SnmpAccessLevel = "read-only"

    // Read Write Access for a community string
    SnmpAccessLevel_read_write SnmpAccessLevel = "read-write"
)

// SnmpBulkstatSchema represents Snmp bulkstat schema
type SnmpBulkstatSchema string

const (
    // Exact Interface
    SnmpBulkstatSchema_exact_interface SnmpBulkstatSchema = "exact-interface"

    // Exact OID
    SnmpBulkstatSchema_exact_oid SnmpBulkstatSchema = "exact-oid"

    // Wild Interface
    SnmpBulkstatSchema_wild_interface SnmpBulkstatSchema = "wild-interface"

    // Wild OID
    SnmpBulkstatSchema_wild_oid SnmpBulkstatSchema = "wild-oid"

    // Range of OID
    SnmpBulkstatSchema_range_oid SnmpBulkstatSchema = "range-oid"

    // Repeated the instance
    SnmpBulkstatSchema_repeat_oid SnmpBulkstatSchema = "repeat-oid"
)

// GroupSnmpVersion represents Group snmp version
type GroupSnmpVersion string

const (
    // SNMP version 1
    GroupSnmpVersion_v1 GroupSnmpVersion = "v1"

    // SNMP version 2
    GroupSnmpVersion_v2c GroupSnmpVersion = "v2c"

    // SNMP version 3
    GroupSnmpVersion_v3 GroupSnmpVersion = "v3"
)

// SnmpOwnerAccess represents Snmp owner access
type SnmpOwnerAccess string

const (
    // Secure Domain Router Owner permissions
    SnmpOwnerAccess_sdr_owner SnmpOwnerAccess = "sdr-owner"

    // System owner permissions
    SnmpOwnerAccess_system_owner SnmpOwnerAccess = "system-owner"
)

// SnmpBulkstatFileFormat represents Snmp bulkstat file format
type SnmpBulkstatFileFormat string

const (
    // Tranfer file in schema Ascii format
    SnmpBulkstatFileFormat_schema_ascii SnmpBulkstatFileFormat = "schema-ascii"

    // Tranfer file in Bulk Ascii format
    SnmpBulkstatFileFormat_bulk_ascii SnmpBulkstatFileFormat = "bulk-ascii"

    // Tranfer file in Bulk binary format
    SnmpBulkstatFileFormat_bulk_binary SnmpBulkstatFileFormat = "bulk-binary"
)

// SnmpSecurityModel represents Snmp security model
type SnmpSecurityModel string

const (
    // No Authentication required
    SnmpSecurityModel_no_authentication SnmpSecurityModel = "no-authentication"

    // Authentication password alone required for
    // access
    SnmpSecurityModel_authentication SnmpSecurityModel = "authentication"

    // Authentication and privacy password required
    // for access
    SnmpSecurityModel_privacy SnmpSecurityModel = "privacy"
)

// SnmpTos represents Snmp tos
type SnmpTos string

const (
    // SNMP TOS type Precedence
    SnmpTos_precedence SnmpTos = "precedence"

    // SNMP TOS type DSCP
    SnmpTos_dscp SnmpTos = "dscp"
)

// Snmpacl represents Snmpacl
type Snmpacl string

const (
    // Ipv4 Access-list
    Snmpacl_ipv4 Snmpacl = "ipv4"

    // Ipv6 Access-list
    Snmpacl_ipv6 Snmpacl = "ipv6"
)

// SnmpDscpValue represents Snmp dscp value
type SnmpDscpValue string

const (
    // Applicable to DSCP: bits 000000
    SnmpDscpValue_default_ SnmpDscpValue = "default"

    // Applicable to DSCP: bits 001010
    SnmpDscpValue_af11 SnmpDscpValue = "af11"

    // Applicable to DSCP: bits 001100
    SnmpDscpValue_af12 SnmpDscpValue = "af12"

    // Applicable to DSCP: bits 001110
    SnmpDscpValue_af13 SnmpDscpValue = "af13"

    // Applicable to DSCP: bits 010010
    SnmpDscpValue_af21 SnmpDscpValue = "af21"

    // Applicable to DSCP: bits 010100
    SnmpDscpValue_af22 SnmpDscpValue = "af22"

    // Applicable to DSCP: bits 010110
    SnmpDscpValue_af23 SnmpDscpValue = "af23"

    // Applicable to DSCP: bits 011010
    SnmpDscpValue_af31 SnmpDscpValue = "af31"

    // Applicable to DSCP: bits 011100
    SnmpDscpValue_af32 SnmpDscpValue = "af32"

    // Applicable to DSCP: bits 011110
    SnmpDscpValue_af33 SnmpDscpValue = "af33"

    // Applicable to DSCP: bits 100010
    SnmpDscpValue_af41 SnmpDscpValue = "af41"

    // Applicable to DSCP: bits 100100
    SnmpDscpValue_af42 SnmpDscpValue = "af42"

    // Applicable to DSCP: bits 100110
    SnmpDscpValue_af43 SnmpDscpValue = "af43"

    // Applicable to DSCP: bits 101110
    SnmpDscpValue_ef SnmpDscpValue = "ef"

    // Applicable to DSCP: bits 001000
    SnmpDscpValue_cs1 SnmpDscpValue = "cs1"

    // Applicable to DSCP: bits 010000
    SnmpDscpValue_cs2 SnmpDscpValue = "cs2"

    // Applicable to DSCP: bits 011000
    SnmpDscpValue_cs3 SnmpDscpValue = "cs3"

    // Applicable to DSCP: bits 100000
    SnmpDscpValue_cs4 SnmpDscpValue = "cs4"

    // Applicable to DSCP: bits 101000
    SnmpDscpValue_cs5 SnmpDscpValue = "cs5"

    // Applicable to DSCP: bits 110000
    SnmpDscpValue_cs6 SnmpDscpValue = "cs6"

    // Applicable to DSCP: bits 111000
    SnmpDscpValue_cs7 SnmpDscpValue = "cs7"
)

// UserSnmpVersion represents User snmp version
type UserSnmpVersion string

const (
    // SNMP version 1
    UserSnmpVersion_v1 UserSnmpVersion = "v1"

    // SNMP version 2
    UserSnmpVersion_v2c UserSnmpVersion = "v2c"

    // SNMP version 3
    UserSnmpVersion_v3 UserSnmpVersion = "v3"
)

// SnmpPrecedenceValue1 represents Snmp precedence value1
type SnmpPrecedenceValue1 string

const (
    // Applicable to Precedence: value 0
    SnmpPrecedenceValue1_routine SnmpPrecedenceValue1 = "routine"

    // Applicable to Precedence: value 1
    SnmpPrecedenceValue1_priority SnmpPrecedenceValue1 = "priority"

    // Applicable to Precedence: value 2
    SnmpPrecedenceValue1_immediate SnmpPrecedenceValue1 = "immediate"

    // Applicable to Precedence: value 3
    SnmpPrecedenceValue1_flash SnmpPrecedenceValue1 = "flash"

    // Applicable to Precedence: value 4
    SnmpPrecedenceValue1_flash_override SnmpPrecedenceValue1 = "flash-override"

    // Applicable to Precedence: value 5
    SnmpPrecedenceValue1_critical SnmpPrecedenceValue1 = "critical"

    // Applicable to Precedence: value 6
    SnmpPrecedenceValue1_internet SnmpPrecedenceValue1 = "internet"

    // Applicable to Precedence: value 7
    SnmpPrecedenceValue1_network SnmpPrecedenceValue1 = "network"
)

// SnmpPrivAlgorithm represents Snmp priv algorithm
type SnmpPrivAlgorithm string

const (
    // No Privacy
    SnmpPrivAlgorithm_none SnmpPrivAlgorithm = "none"

    // Des algorithm
    SnmpPrivAlgorithm_des SnmpPrivAlgorithm = "des"

    // 3des algorithm
    SnmpPrivAlgorithm_Y_3des SnmpPrivAlgorithm = "3des"

    // aes128 algorithm
    SnmpPrivAlgorithm_aes128 SnmpPrivAlgorithm = "aes128"

    // aes192 algorithm
    SnmpPrivAlgorithm_aes192 SnmpPrivAlgorithm = "aes192"

    // aes256 algorithm
    SnmpPrivAlgorithm_aes256 SnmpPrivAlgorithm = "aes256"
)

// SnmpMibViewInclusion represents Snmp mib view inclusion
type SnmpMibViewInclusion string

const (
    // MIB View to be included
    SnmpMibViewInclusion_included SnmpMibViewInclusion = "included"

    // MIB View to be excluded
    SnmpMibViewInclusion_excluded SnmpMibViewInclusion = "excluded"
)

// SnmpContext represents Snmp context
type SnmpContext string

const (
    // VRF feature
    SnmpContext_vrf SnmpContext = "vrf"

    // BRIDGE feature
    SnmpContext_bridge SnmpContext = "bridge"

    // OSPF feature
    SnmpContext_ospf SnmpContext = "ospf"

    // OSPFv3 feature
    SnmpContext_ospfv3 SnmpContext = "ospfv3"
)

// Snmp
// The heirarchy point for all the SNMP
// configurations
type Snmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of times to retry an Inform request (default 3). The type is
    // interface{} with range: 0..100.
    InformRetries interface{}

    // Change the source port of all traps. The type is interface{} with range:
    // 1024..65535.
    TrapPort interface{}

    // Enable Poll OID statistics. The type is interface{}.
    OidPollStats interface{}

    // Assign an interface for the source address of all traps. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    TrapSource interface{}

    // Disable authentication traps for packets on a vrf. The type is interface{}.
    VrfAuthenticationTrapDisable interface{}

    // Timeout value in seconds for Inform request (default 15 sec). The type is
    // interface{} with range: 1..42949671. Units are second.
    InformTimeout interface{}

    // Assign an interface for the source IPV6 address of all traps. The type is
    // string with pattern: [a-zA-Z0-9./-]+.
    TrapSourceIpv6 interface{}

    // Largest SNMP packet size. The type is interface{} with range: 484..65500.
    PacketSize interface{}

    // Throttle time for incoming queue (default 0 msec). The type is interface{}
    // with range: 50..1000.
    ThrottleTime interface{}

    // Assign an interface for the source address of all traps. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    TrapSourceIpv4 interface{}

    // Max nmber of informs to hold in queue, (default 25). The type is
    // interface{} with range: 0..4294967295.
    InformPending interface{}

    // Container class to hold clear/encrypted communitie maps.
    EncryptedCommunityMaps Snmp_EncryptedCommunityMaps

    // Class to configure a SNMPv2 MIB view.
    Views Snmp_Views

    // SNMP logging.
    Logging Snmp_Logging

    // Container class for SNMP administration.
    Administration Snmp_Administration

    // The heirarchy point for SNMP Agent configurations.
    Agent Snmp_Agent

    // Class to hold trap configurations.
    Trap Snmp_Trap

    // SNMP TOS bit for outgoing packets.
    Ipv6 Snmp_Ipv6

    // SNMP TOS bit for outgoing packets.
    Ipv4 Snmp_Ipv4

    // container to hold system information.
    System Snmp_System

    // SNMP target configurations.
    Target Snmp_Target

    // Enable SNMP notifications.
    Notification Snmp_Notification

    // Configure properties of the trap correlator.
    Correlator Snmp_Correlator

    // SNMP bulk stats configuration commands.
    BulkStats Snmp_BulkStats

    // Container class to hold unencrpted community map.
    DefaultCommunityMaps Snmp_DefaultCommunityMaps

    // Set overload control params for handling incoming messages.
    OverloadControl Snmp_OverloadControl

    // SNMP timeouts.
    Timeouts Snmp_Timeouts

    // Define a user who can access the SNMP engine.
    Users Snmp_Users

    // SNMP VRF configuration commands.
    Vrfs Snmp_Vrfs

    // Define a User Security Model group.
    Groups Snmp_Groups

    // Specify hosts to receive SNMP notifications.
    TrapHosts Snmp_TrapHosts

    // List of Context Names.
    Contexts Snmp_Contexts

    // List of context names.
    ContextMappings Snmp_ContextMappings
}

func (snmp *Snmp) GetFilter() yfilter.YFilter { return snmp.YFilter }

func (snmp *Snmp) SetFilter(yf yfilter.YFilter) { snmp.YFilter = yf }

func (snmp *Snmp) GetGoName(yname string) string {
    if yname == "inform-retries" { return "InformRetries" }
    if yname == "trap-port" { return "TrapPort" }
    if yname == "oid-poll-stats" { return "OidPollStats" }
    if yname == "trap-source" { return "TrapSource" }
    if yname == "vrf-authentication-trap-disable" { return "VrfAuthenticationTrapDisable" }
    if yname == "inform-timeout" { return "InformTimeout" }
    if yname == "trap-source-ipv6" { return "TrapSourceIpv6" }
    if yname == "packet-size" { return "PacketSize" }
    if yname == "throttle-time" { return "ThrottleTime" }
    if yname == "trap-source-ipv4" { return "TrapSourceIpv4" }
    if yname == "inform-pending" { return "InformPending" }
    if yname == "encrypted-community-maps" { return "EncryptedCommunityMaps" }
    if yname == "views" { return "Views" }
    if yname == "logging" { return "Logging" }
    if yname == "administration" { return "Administration" }
    if yname == "agent" { return "Agent" }
    if yname == "trap" { return "Trap" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "system" { return "System" }
    if yname == "target" { return "Target" }
    if yname == "notification" { return "Notification" }
    if yname == "correlator" { return "Correlator" }
    if yname == "bulk-stats" { return "BulkStats" }
    if yname == "default-community-maps" { return "DefaultCommunityMaps" }
    if yname == "overload-control" { return "OverloadControl" }
    if yname == "timeouts" { return "Timeouts" }
    if yname == "users" { return "Users" }
    if yname == "vrfs" { return "Vrfs" }
    if yname == "groups" { return "Groups" }
    if yname == "trap-hosts" { return "TrapHosts" }
    if yname == "contexts" { return "Contexts" }
    if yname == "context-mappings" { return "ContextMappings" }
    return ""
}

func (snmp *Snmp) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-agent-cfg:snmp"
}

func (snmp *Snmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "encrypted-community-maps" {
        return &snmp.EncryptedCommunityMaps
    }
    if childYangName == "views" {
        return &snmp.Views
    }
    if childYangName == "logging" {
        return &snmp.Logging
    }
    if childYangName == "administration" {
        return &snmp.Administration
    }
    if childYangName == "agent" {
        return &snmp.Agent
    }
    if childYangName == "trap" {
        return &snmp.Trap
    }
    if childYangName == "ipv6" {
        return &snmp.Ipv6
    }
    if childYangName == "ipv4" {
        return &snmp.Ipv4
    }
    if childYangName == "system" {
        return &snmp.System
    }
    if childYangName == "target" {
        return &snmp.Target
    }
    if childYangName == "notification" {
        return &snmp.Notification
    }
    if childYangName == "correlator" {
        return &snmp.Correlator
    }
    if childYangName == "bulk-stats" {
        return &snmp.BulkStats
    }
    if childYangName == "default-community-maps" {
        return &snmp.DefaultCommunityMaps
    }
    if childYangName == "overload-control" {
        return &snmp.OverloadControl
    }
    if childYangName == "timeouts" {
        return &snmp.Timeouts
    }
    if childYangName == "users" {
        return &snmp.Users
    }
    if childYangName == "vrfs" {
        return &snmp.Vrfs
    }
    if childYangName == "groups" {
        return &snmp.Groups
    }
    if childYangName == "trap-hosts" {
        return &snmp.TrapHosts
    }
    if childYangName == "contexts" {
        return &snmp.Contexts
    }
    if childYangName == "context-mappings" {
        return &snmp.ContextMappings
    }
    return nil
}

func (snmp *Snmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["encrypted-community-maps"] = &snmp.EncryptedCommunityMaps
    children["views"] = &snmp.Views
    children["logging"] = &snmp.Logging
    children["administration"] = &snmp.Administration
    children["agent"] = &snmp.Agent
    children["trap"] = &snmp.Trap
    children["ipv6"] = &snmp.Ipv6
    children["ipv4"] = &snmp.Ipv4
    children["system"] = &snmp.System
    children["target"] = &snmp.Target
    children["notification"] = &snmp.Notification
    children["correlator"] = &snmp.Correlator
    children["bulk-stats"] = &snmp.BulkStats
    children["default-community-maps"] = &snmp.DefaultCommunityMaps
    children["overload-control"] = &snmp.OverloadControl
    children["timeouts"] = &snmp.Timeouts
    children["users"] = &snmp.Users
    children["vrfs"] = &snmp.Vrfs
    children["groups"] = &snmp.Groups
    children["trap-hosts"] = &snmp.TrapHosts
    children["contexts"] = &snmp.Contexts
    children["context-mappings"] = &snmp.ContextMappings
    return children
}

func (snmp *Snmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["inform-retries"] = snmp.InformRetries
    leafs["trap-port"] = snmp.TrapPort
    leafs["oid-poll-stats"] = snmp.OidPollStats
    leafs["trap-source"] = snmp.TrapSource
    leafs["vrf-authentication-trap-disable"] = snmp.VrfAuthenticationTrapDisable
    leafs["inform-timeout"] = snmp.InformTimeout
    leafs["trap-source-ipv6"] = snmp.TrapSourceIpv6
    leafs["packet-size"] = snmp.PacketSize
    leafs["throttle-time"] = snmp.ThrottleTime
    leafs["trap-source-ipv4"] = snmp.TrapSourceIpv4
    leafs["inform-pending"] = snmp.InformPending
    return leafs
}

func (snmp *Snmp) GetBundleName() string { return "cisco_ios_xr" }

func (snmp *Snmp) GetYangName() string { return "snmp" }

func (snmp *Snmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (snmp *Snmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (snmp *Snmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (snmp *Snmp) SetParent(parent types.Entity) { snmp.parent = parent }

func (snmp *Snmp) GetParent() types.Entity { return snmp.parent }

func (snmp *Snmp) GetParentYangName() string { return "Cisco-IOS-XR-snmp-agent-cfg" }

// Snmp_EncryptedCommunityMaps
// Container class to hold clear/encrypted
// communitie maps
type Snmp_EncryptedCommunityMaps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear/encrypted SNMP community map. The type is slice of
    // Snmp_EncryptedCommunityMaps_EncryptedCommunityMap.
    EncryptedCommunityMap []Snmp_EncryptedCommunityMaps_EncryptedCommunityMap
}

func (encryptedCommunityMaps *Snmp_EncryptedCommunityMaps) GetFilter() yfilter.YFilter { return encryptedCommunityMaps.YFilter }

func (encryptedCommunityMaps *Snmp_EncryptedCommunityMaps) SetFilter(yf yfilter.YFilter) { encryptedCommunityMaps.YFilter = yf }

func (encryptedCommunityMaps *Snmp_EncryptedCommunityMaps) GetGoName(yname string) string {
    if yname == "encrypted-community-map" { return "EncryptedCommunityMap" }
    return ""
}

func (encryptedCommunityMaps *Snmp_EncryptedCommunityMaps) GetSegmentPath() string {
    return "encrypted-community-maps"
}

func (encryptedCommunityMaps *Snmp_EncryptedCommunityMaps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "encrypted-community-map" {
        for _, c := range encryptedCommunityMaps.EncryptedCommunityMap {
            if encryptedCommunityMaps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_EncryptedCommunityMaps_EncryptedCommunityMap{}
        encryptedCommunityMaps.EncryptedCommunityMap = append(encryptedCommunityMaps.EncryptedCommunityMap, child)
        return &encryptedCommunityMaps.EncryptedCommunityMap[len(encryptedCommunityMaps.EncryptedCommunityMap)-1]
    }
    return nil
}

func (encryptedCommunityMaps *Snmp_EncryptedCommunityMaps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range encryptedCommunityMaps.EncryptedCommunityMap {
        children[encryptedCommunityMaps.EncryptedCommunityMap[i].GetSegmentPath()] = &encryptedCommunityMaps.EncryptedCommunityMap[i]
    }
    return children
}

func (encryptedCommunityMaps *Snmp_EncryptedCommunityMaps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (encryptedCommunityMaps *Snmp_EncryptedCommunityMaps) GetBundleName() string { return "cisco_ios_xr" }

func (encryptedCommunityMaps *Snmp_EncryptedCommunityMaps) GetYangName() string { return "encrypted-community-maps" }

func (encryptedCommunityMaps *Snmp_EncryptedCommunityMaps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (encryptedCommunityMaps *Snmp_EncryptedCommunityMaps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (encryptedCommunityMaps *Snmp_EncryptedCommunityMaps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (encryptedCommunityMaps *Snmp_EncryptedCommunityMaps) SetParent(parent types.Entity) { encryptedCommunityMaps.parent = parent }

func (encryptedCommunityMaps *Snmp_EncryptedCommunityMaps) GetParent() types.Entity { return encryptedCommunityMaps.parent }

func (encryptedCommunityMaps *Snmp_EncryptedCommunityMaps) GetParentYangName() string { return "snmp" }

// Snmp_EncryptedCommunityMaps_EncryptedCommunityMap
// Clear/encrypted SNMP community map
type Snmp_EncryptedCommunityMaps_EncryptedCommunityMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SNMP community map. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    CommunityName interface{}

    // SNMP Context Name . The type is string.
    Context interface{}

    // SNMP Security Name . The type is string.
    Security interface{}

    // target list name . The type is string.
    TargetList interface{}
}

func (encryptedCommunityMap *Snmp_EncryptedCommunityMaps_EncryptedCommunityMap) GetFilter() yfilter.YFilter { return encryptedCommunityMap.YFilter }

func (encryptedCommunityMap *Snmp_EncryptedCommunityMaps_EncryptedCommunityMap) SetFilter(yf yfilter.YFilter) { encryptedCommunityMap.YFilter = yf }

func (encryptedCommunityMap *Snmp_EncryptedCommunityMaps_EncryptedCommunityMap) GetGoName(yname string) string {
    if yname == "community-name" { return "CommunityName" }
    if yname == "context" { return "Context" }
    if yname == "security" { return "Security" }
    if yname == "target-list" { return "TargetList" }
    return ""
}

func (encryptedCommunityMap *Snmp_EncryptedCommunityMaps_EncryptedCommunityMap) GetSegmentPath() string {
    return "encrypted-community-map" + "[community-name='" + fmt.Sprintf("%v", encryptedCommunityMap.CommunityName) + "']"
}

func (encryptedCommunityMap *Snmp_EncryptedCommunityMaps_EncryptedCommunityMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (encryptedCommunityMap *Snmp_EncryptedCommunityMaps_EncryptedCommunityMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (encryptedCommunityMap *Snmp_EncryptedCommunityMaps_EncryptedCommunityMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-name"] = encryptedCommunityMap.CommunityName
    leafs["context"] = encryptedCommunityMap.Context
    leafs["security"] = encryptedCommunityMap.Security
    leafs["target-list"] = encryptedCommunityMap.TargetList
    return leafs
}

func (encryptedCommunityMap *Snmp_EncryptedCommunityMaps_EncryptedCommunityMap) GetBundleName() string { return "cisco_ios_xr" }

func (encryptedCommunityMap *Snmp_EncryptedCommunityMaps_EncryptedCommunityMap) GetYangName() string { return "encrypted-community-map" }

func (encryptedCommunityMap *Snmp_EncryptedCommunityMaps_EncryptedCommunityMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (encryptedCommunityMap *Snmp_EncryptedCommunityMaps_EncryptedCommunityMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (encryptedCommunityMap *Snmp_EncryptedCommunityMaps_EncryptedCommunityMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (encryptedCommunityMap *Snmp_EncryptedCommunityMaps_EncryptedCommunityMap) SetParent(parent types.Entity) { encryptedCommunityMap.parent = parent }

func (encryptedCommunityMap *Snmp_EncryptedCommunityMaps_EncryptedCommunityMap) GetParent() types.Entity { return encryptedCommunityMap.parent }

func (encryptedCommunityMap *Snmp_EncryptedCommunityMaps_EncryptedCommunityMap) GetParentYangName() string { return "encrypted-community-maps" }

// Snmp_Views
// Class to configure a SNMPv2 MIB view
type Snmp_Views struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the view. The type is slice of Snmp_Views_View.
    View []Snmp_Views_View
}

func (views *Snmp_Views) GetFilter() yfilter.YFilter { return views.YFilter }

func (views *Snmp_Views) SetFilter(yf yfilter.YFilter) { views.YFilter = yf }

func (views *Snmp_Views) GetGoName(yname string) string {
    if yname == "view" { return "View" }
    return ""
}

func (views *Snmp_Views) GetSegmentPath() string {
    return "views"
}

func (views *Snmp_Views) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "view" {
        for _, c := range views.View {
            if views.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Views_View{}
        views.View = append(views.View, child)
        return &views.View[len(views.View)-1]
    }
    return nil
}

func (views *Snmp_Views) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range views.View {
        children[views.View[i].GetSegmentPath()] = &views.View[i]
    }
    return children
}

func (views *Snmp_Views) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (views *Snmp_Views) GetBundleName() string { return "cisco_ios_xr" }

func (views *Snmp_Views) GetYangName() string { return "views" }

func (views *Snmp_Views) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (views *Snmp_Views) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (views *Snmp_Views) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (views *Snmp_Views) SetParent(parent types.Entity) { views.parent = parent }

func (views *Snmp_Views) GetParent() types.Entity { return views.parent }

func (views *Snmp_Views) GetParentYangName() string { return "snmp" }

// Snmp_Views_View
// Name of the view
type Snmp_Views_View struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the view. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ViewName interface{}

    // This attribute is a key. MIB view family name. The type is string.
    Family interface{}

    // MIB view to be included or excluded. The type is SnmpMibViewInclusion. This
    // attribute is mandatory.
    ViewInclusion interface{}
}

func (view *Snmp_Views_View) GetFilter() yfilter.YFilter { return view.YFilter }

func (view *Snmp_Views_View) SetFilter(yf yfilter.YFilter) { view.YFilter = yf }

func (view *Snmp_Views_View) GetGoName(yname string) string {
    if yname == "view-name" { return "ViewName" }
    if yname == "family" { return "Family" }
    if yname == "view-inclusion" { return "ViewInclusion" }
    return ""
}

func (view *Snmp_Views_View) GetSegmentPath() string {
    return "view" + "[view-name='" + fmt.Sprintf("%v", view.ViewName) + "']" + "[family='" + fmt.Sprintf("%v", view.Family) + "']"
}

func (view *Snmp_Views_View) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (view *Snmp_Views_View) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (view *Snmp_Views_View) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["view-name"] = view.ViewName
    leafs["family"] = view.Family
    leafs["view-inclusion"] = view.ViewInclusion
    return leafs
}

func (view *Snmp_Views_View) GetBundleName() string { return "cisco_ios_xr" }

func (view *Snmp_Views_View) GetYangName() string { return "view" }

func (view *Snmp_Views_View) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (view *Snmp_Views_View) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (view *Snmp_Views_View) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (view *Snmp_Views_View) SetParent(parent types.Entity) { view.parent = parent }

func (view *Snmp_Views_View) GetParent() types.Entity { return view.parent }

func (view *Snmp_Views_View) GetParentYangName() string { return "views" }

// Snmp_Logging
// SNMP logging
type Snmp_Logging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP logging threshold.
    Threshold Snmp_Logging_Threshold
}

func (logging *Snmp_Logging) GetFilter() yfilter.YFilter { return logging.YFilter }

func (logging *Snmp_Logging) SetFilter(yf yfilter.YFilter) { logging.YFilter = yf }

func (logging *Snmp_Logging) GetGoName(yname string) string {
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (logging *Snmp_Logging) GetSegmentPath() string {
    return "logging"
}

func (logging *Snmp_Logging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold" {
        return &logging.Threshold
    }
    return nil
}

func (logging *Snmp_Logging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["threshold"] = &logging.Threshold
    return children
}

func (logging *Snmp_Logging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (logging *Snmp_Logging) GetBundleName() string { return "cisco_ios_xr" }

func (logging *Snmp_Logging) GetYangName() string { return "logging" }

func (logging *Snmp_Logging) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logging *Snmp_Logging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logging *Snmp_Logging) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logging *Snmp_Logging) SetParent(parent types.Entity) { logging.parent = parent }

func (logging *Snmp_Logging) GetParent() types.Entity { return logging.parent }

func (logging *Snmp_Logging) GetParentYangName() string { return "snmp" }

// Snmp_Logging_Threshold
// SNMP logging threshold
type Snmp_Logging_Threshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP logging threshold for OID processing. The type is interface{} with
    // range: 0..20000. The default value is 500.
    OidProcessing interface{}

    // SNMP logging threshold for PDU processing. The type is interface{} with
    // range: 0..20000. The default value is 20000.
    PduProcessing interface{}
}

func (threshold *Snmp_Logging_Threshold) GetFilter() yfilter.YFilter { return threshold.YFilter }

func (threshold *Snmp_Logging_Threshold) SetFilter(yf yfilter.YFilter) { threshold.YFilter = yf }

func (threshold *Snmp_Logging_Threshold) GetGoName(yname string) string {
    if yname == "oid-processing" { return "OidProcessing" }
    if yname == "pdu-processing" { return "PduProcessing" }
    return ""
}

func (threshold *Snmp_Logging_Threshold) GetSegmentPath() string {
    return "threshold"
}

func (threshold *Snmp_Logging_Threshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (threshold *Snmp_Logging_Threshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (threshold *Snmp_Logging_Threshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oid-processing"] = threshold.OidProcessing
    leafs["pdu-processing"] = threshold.PduProcessing
    return leafs
}

func (threshold *Snmp_Logging_Threshold) GetBundleName() string { return "cisco_ios_xr" }

func (threshold *Snmp_Logging_Threshold) GetYangName() string { return "threshold" }

func (threshold *Snmp_Logging_Threshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (threshold *Snmp_Logging_Threshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (threshold *Snmp_Logging_Threshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (threshold *Snmp_Logging_Threshold) SetParent(parent types.Entity) { threshold.parent = parent }

func (threshold *Snmp_Logging_Threshold) GetParent() types.Entity { return threshold.parent }

func (threshold *Snmp_Logging_Threshold) GetParentYangName() string { return "logging" }

// Snmp_Administration
// Container class for SNMP administration
type Snmp_Administration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Container class to hold unencrpted communities.
    DefaultCommunities Snmp_Administration_DefaultCommunities

    // Container class to hold clear/encrypted communities.
    EncryptedCommunities Snmp_Administration_EncryptedCommunities
}

func (administration *Snmp_Administration) GetFilter() yfilter.YFilter { return administration.YFilter }

func (administration *Snmp_Administration) SetFilter(yf yfilter.YFilter) { administration.YFilter = yf }

func (administration *Snmp_Administration) GetGoName(yname string) string {
    if yname == "default-communities" { return "DefaultCommunities" }
    if yname == "encrypted-communities" { return "EncryptedCommunities" }
    return ""
}

func (administration *Snmp_Administration) GetSegmentPath() string {
    return "administration"
}

func (administration *Snmp_Administration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-communities" {
        return &administration.DefaultCommunities
    }
    if childYangName == "encrypted-communities" {
        return &administration.EncryptedCommunities
    }
    return nil
}

func (administration *Snmp_Administration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-communities"] = &administration.DefaultCommunities
    children["encrypted-communities"] = &administration.EncryptedCommunities
    return children
}

func (administration *Snmp_Administration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (administration *Snmp_Administration) GetBundleName() string { return "cisco_ios_xr" }

func (administration *Snmp_Administration) GetYangName() string { return "administration" }

func (administration *Snmp_Administration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (administration *Snmp_Administration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (administration *Snmp_Administration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (administration *Snmp_Administration) SetParent(parent types.Entity) { administration.parent = parent }

func (administration *Snmp_Administration) GetParent() types.Entity { return administration.parent }

func (administration *Snmp_Administration) GetParentYangName() string { return "snmp" }

// Snmp_Administration_DefaultCommunities
// Container class to hold unencrpted communities
type Snmp_Administration_DefaultCommunities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unencrpted SNMP community string and access priviledges. The type is slice
    // of Snmp_Administration_DefaultCommunities_DefaultCommunity.
    DefaultCommunity []Snmp_Administration_DefaultCommunities_DefaultCommunity
}

func (defaultCommunities *Snmp_Administration_DefaultCommunities) GetFilter() yfilter.YFilter { return defaultCommunities.YFilter }

func (defaultCommunities *Snmp_Administration_DefaultCommunities) SetFilter(yf yfilter.YFilter) { defaultCommunities.YFilter = yf }

func (defaultCommunities *Snmp_Administration_DefaultCommunities) GetGoName(yname string) string {
    if yname == "default-community" { return "DefaultCommunity" }
    return ""
}

func (defaultCommunities *Snmp_Administration_DefaultCommunities) GetSegmentPath() string {
    return "default-communities"
}

func (defaultCommunities *Snmp_Administration_DefaultCommunities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-community" {
        for _, c := range defaultCommunities.DefaultCommunity {
            if defaultCommunities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Administration_DefaultCommunities_DefaultCommunity{}
        defaultCommunities.DefaultCommunity = append(defaultCommunities.DefaultCommunity, child)
        return &defaultCommunities.DefaultCommunity[len(defaultCommunities.DefaultCommunity)-1]
    }
    return nil
}

func (defaultCommunities *Snmp_Administration_DefaultCommunities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range defaultCommunities.DefaultCommunity {
        children[defaultCommunities.DefaultCommunity[i].GetSegmentPath()] = &defaultCommunities.DefaultCommunity[i]
    }
    return children
}

func (defaultCommunities *Snmp_Administration_DefaultCommunities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (defaultCommunities *Snmp_Administration_DefaultCommunities) GetBundleName() string { return "cisco_ios_xr" }

func (defaultCommunities *Snmp_Administration_DefaultCommunities) GetYangName() string { return "default-communities" }

func (defaultCommunities *Snmp_Administration_DefaultCommunities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultCommunities *Snmp_Administration_DefaultCommunities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultCommunities *Snmp_Administration_DefaultCommunities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultCommunities *Snmp_Administration_DefaultCommunities) SetParent(parent types.Entity) { defaultCommunities.parent = parent }

func (defaultCommunities *Snmp_Administration_DefaultCommunities) GetParent() types.Entity { return defaultCommunities.parent }

func (defaultCommunities *Snmp_Administration_DefaultCommunities) GetParentYangName() string { return "administration" }

// Snmp_Administration_DefaultCommunities_DefaultCommunity
// Unencrpted SNMP community string and access
// priviledges
type Snmp_Administration_DefaultCommunities_DefaultCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SNMP community string. The type is string with
    // length: 1..128.
    CommunityName interface{}

    // Read/Write Access. The type is SnmpAccessLevel.
    Priviledge interface{}

    // MIB view to which the community has access. The type is string.
    ViewName interface{}

    // Access-list type. The type is Snmpacl.
    V4AclType interface{}

    // Ipv4 Access-list name. The type is string.
    V4AccessList interface{}

    // Access-list type. The type is Snmpacl.
    V6AclType interface{}

    // Ipv6 Access-list name. The type is string.
    V6AccessList interface{}

    // Logical Router or System owner access. The type is SnmpOwnerAccess.
    Owner interface{}
}

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) GetFilter() yfilter.YFilter { return defaultCommunity.YFilter }

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) SetFilter(yf yfilter.YFilter) { defaultCommunity.YFilter = yf }

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) GetGoName(yname string) string {
    if yname == "community-name" { return "CommunityName" }
    if yname == "priviledge" { return "Priviledge" }
    if yname == "view-name" { return "ViewName" }
    if yname == "v4acl-type" { return "V4AclType" }
    if yname == "v4-access-list" { return "V4AccessList" }
    if yname == "v6acl-type" { return "V6AclType" }
    if yname == "v6-access-list" { return "V6AccessList" }
    if yname == "owner" { return "Owner" }
    return ""
}

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) GetSegmentPath() string {
    return "default-community" + "[community-name='" + fmt.Sprintf("%v", defaultCommunity.CommunityName) + "']"
}

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-name"] = defaultCommunity.CommunityName
    leafs["priviledge"] = defaultCommunity.Priviledge
    leafs["view-name"] = defaultCommunity.ViewName
    leafs["v4acl-type"] = defaultCommunity.V4AclType
    leafs["v4-access-list"] = defaultCommunity.V4AccessList
    leafs["v6acl-type"] = defaultCommunity.V6AclType
    leafs["v6-access-list"] = defaultCommunity.V6AccessList
    leafs["owner"] = defaultCommunity.Owner
    return leafs
}

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) GetYangName() string { return "default-community" }

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) SetParent(parent types.Entity) { defaultCommunity.parent = parent }

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) GetParent() types.Entity { return defaultCommunity.parent }

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) GetParentYangName() string { return "default-communities" }

// Snmp_Administration_EncryptedCommunities
// Container class to hold clear/encrypted
// communities
type Snmp_Administration_EncryptedCommunities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear/encrypted SNMP community string and access priviledges. The type is
    // slice of Snmp_Administration_EncryptedCommunities_EncryptedCommunity.
    EncryptedCommunity []Snmp_Administration_EncryptedCommunities_EncryptedCommunity
}

func (encryptedCommunities *Snmp_Administration_EncryptedCommunities) GetFilter() yfilter.YFilter { return encryptedCommunities.YFilter }

func (encryptedCommunities *Snmp_Administration_EncryptedCommunities) SetFilter(yf yfilter.YFilter) { encryptedCommunities.YFilter = yf }

func (encryptedCommunities *Snmp_Administration_EncryptedCommunities) GetGoName(yname string) string {
    if yname == "encrypted-community" { return "EncryptedCommunity" }
    return ""
}

func (encryptedCommunities *Snmp_Administration_EncryptedCommunities) GetSegmentPath() string {
    return "encrypted-communities"
}

func (encryptedCommunities *Snmp_Administration_EncryptedCommunities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "encrypted-community" {
        for _, c := range encryptedCommunities.EncryptedCommunity {
            if encryptedCommunities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Administration_EncryptedCommunities_EncryptedCommunity{}
        encryptedCommunities.EncryptedCommunity = append(encryptedCommunities.EncryptedCommunity, child)
        return &encryptedCommunities.EncryptedCommunity[len(encryptedCommunities.EncryptedCommunity)-1]
    }
    return nil
}

func (encryptedCommunities *Snmp_Administration_EncryptedCommunities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range encryptedCommunities.EncryptedCommunity {
        children[encryptedCommunities.EncryptedCommunity[i].GetSegmentPath()] = &encryptedCommunities.EncryptedCommunity[i]
    }
    return children
}

func (encryptedCommunities *Snmp_Administration_EncryptedCommunities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (encryptedCommunities *Snmp_Administration_EncryptedCommunities) GetBundleName() string { return "cisco_ios_xr" }

func (encryptedCommunities *Snmp_Administration_EncryptedCommunities) GetYangName() string { return "encrypted-communities" }

func (encryptedCommunities *Snmp_Administration_EncryptedCommunities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (encryptedCommunities *Snmp_Administration_EncryptedCommunities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (encryptedCommunities *Snmp_Administration_EncryptedCommunities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (encryptedCommunities *Snmp_Administration_EncryptedCommunities) SetParent(parent types.Entity) { encryptedCommunities.parent = parent }

func (encryptedCommunities *Snmp_Administration_EncryptedCommunities) GetParent() types.Entity { return encryptedCommunities.parent }

func (encryptedCommunities *Snmp_Administration_EncryptedCommunities) GetParentYangName() string { return "administration" }

// Snmp_Administration_EncryptedCommunities_EncryptedCommunity
// Clear/encrypted SNMP community string and
// access priviledges
type Snmp_Administration_EncryptedCommunities_EncryptedCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SNMP community string. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    CommunityName interface{}

    // Read/Write Access. The type is SnmpAccessLevel.
    Priviledge interface{}

    // MIB view to which the community has access. The type is string.
    ViewName interface{}

    // Access-list type. The type is Snmpacl.
    V4AclType interface{}

    // Ipv4 Access-list name. The type is string.
    V4AccessList interface{}

    // Access-list type. The type is Snmpacl.
    V6AclType interface{}

    // Ipv6 Access-list name. The type is string.
    V6AccessList interface{}

    // Logical Router or System owner access. The type is SnmpOwnerAccess.
    Owner interface{}
}

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) GetFilter() yfilter.YFilter { return encryptedCommunity.YFilter }

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) SetFilter(yf yfilter.YFilter) { encryptedCommunity.YFilter = yf }

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) GetGoName(yname string) string {
    if yname == "community-name" { return "CommunityName" }
    if yname == "priviledge" { return "Priviledge" }
    if yname == "view-name" { return "ViewName" }
    if yname == "v4acl-type" { return "V4AclType" }
    if yname == "v4-access-list" { return "V4AccessList" }
    if yname == "v6acl-type" { return "V6AclType" }
    if yname == "v6-access-list" { return "V6AccessList" }
    if yname == "owner" { return "Owner" }
    return ""
}

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) GetSegmentPath() string {
    return "encrypted-community" + "[community-name='" + fmt.Sprintf("%v", encryptedCommunity.CommunityName) + "']"
}

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-name"] = encryptedCommunity.CommunityName
    leafs["priviledge"] = encryptedCommunity.Priviledge
    leafs["view-name"] = encryptedCommunity.ViewName
    leafs["v4acl-type"] = encryptedCommunity.V4AclType
    leafs["v4-access-list"] = encryptedCommunity.V4AccessList
    leafs["v6acl-type"] = encryptedCommunity.V6AclType
    leafs["v6-access-list"] = encryptedCommunity.V6AccessList
    leafs["owner"] = encryptedCommunity.Owner
    return leafs
}

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) GetYangName() string { return "encrypted-community" }

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) SetParent(parent types.Entity) { encryptedCommunity.parent = parent }

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) GetParent() types.Entity { return encryptedCommunity.parent }

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) GetParentYangName() string { return "encrypted-communities" }

// Snmp_Agent
// The heirarchy point for SNMP Agent
// configurations
type Snmp_Agent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMPv3 engineID.
    EngineId Snmp_Agent_EngineId
}

func (agent *Snmp_Agent) GetFilter() yfilter.YFilter { return agent.YFilter }

func (agent *Snmp_Agent) SetFilter(yf yfilter.YFilter) { agent.YFilter = yf }

func (agent *Snmp_Agent) GetGoName(yname string) string {
    if yname == "engine-id" { return "EngineId" }
    return ""
}

func (agent *Snmp_Agent) GetSegmentPath() string {
    return "agent"
}

func (agent *Snmp_Agent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "engine-id" {
        return &agent.EngineId
    }
    return nil
}

func (agent *Snmp_Agent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["engine-id"] = &agent.EngineId
    return children
}

func (agent *Snmp_Agent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (agent *Snmp_Agent) GetBundleName() string { return "cisco_ios_xr" }

func (agent *Snmp_Agent) GetYangName() string { return "agent" }

func (agent *Snmp_Agent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (agent *Snmp_Agent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (agent *Snmp_Agent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (agent *Snmp_Agent) SetParent(parent types.Entity) { agent.parent = parent }

func (agent *Snmp_Agent) GetParent() types.Entity { return agent.parent }

func (agent *Snmp_Agent) GetParentYangName() string { return "snmp" }

// Snmp_Agent_EngineId
// SNMPv3 engineID
type Snmp_Agent_EngineId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // engineID of the local agent. The type is string.
    Local interface{}

    // SNMPv3 remote SNMP Entity.
    Remotes Snmp_Agent_EngineId_Remotes
}

func (engineId *Snmp_Agent_EngineId) GetFilter() yfilter.YFilter { return engineId.YFilter }

func (engineId *Snmp_Agent_EngineId) SetFilter(yf yfilter.YFilter) { engineId.YFilter = yf }

func (engineId *Snmp_Agent_EngineId) GetGoName(yname string) string {
    if yname == "local" { return "Local" }
    if yname == "remotes" { return "Remotes" }
    return ""
}

func (engineId *Snmp_Agent_EngineId) GetSegmentPath() string {
    return "engine-id"
}

func (engineId *Snmp_Agent_EngineId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remotes" {
        return &engineId.Remotes
    }
    return nil
}

func (engineId *Snmp_Agent_EngineId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["remotes"] = &engineId.Remotes
    return children
}

func (engineId *Snmp_Agent_EngineId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local"] = engineId.Local
    return leafs
}

func (engineId *Snmp_Agent_EngineId) GetBundleName() string { return "cisco_ios_xr" }

func (engineId *Snmp_Agent_EngineId) GetYangName() string { return "engine-id" }

func (engineId *Snmp_Agent_EngineId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (engineId *Snmp_Agent_EngineId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (engineId *Snmp_Agent_EngineId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (engineId *Snmp_Agent_EngineId) SetParent(parent types.Entity) { engineId.parent = parent }

func (engineId *Snmp_Agent_EngineId) GetParent() types.Entity { return engineId.parent }

func (engineId *Snmp_Agent_EngineId) GetParentYangName() string { return "agent" }

// Snmp_Agent_EngineId_Remotes
// SNMPv3 remote SNMP Entity
type Snmp_Agent_EngineId_Remotes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // engineID of the remote agent. The type is slice of
    // Snmp_Agent_EngineId_Remotes_Remote.
    Remote []Snmp_Agent_EngineId_Remotes_Remote
}

func (remotes *Snmp_Agent_EngineId_Remotes) GetFilter() yfilter.YFilter { return remotes.YFilter }

func (remotes *Snmp_Agent_EngineId_Remotes) SetFilter(yf yfilter.YFilter) { remotes.YFilter = yf }

func (remotes *Snmp_Agent_EngineId_Remotes) GetGoName(yname string) string {
    if yname == "remote" { return "Remote" }
    return ""
}

func (remotes *Snmp_Agent_EngineId_Remotes) GetSegmentPath() string {
    return "remotes"
}

func (remotes *Snmp_Agent_EngineId_Remotes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote" {
        for _, c := range remotes.Remote {
            if remotes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Agent_EngineId_Remotes_Remote{}
        remotes.Remote = append(remotes.Remote, child)
        return &remotes.Remote[len(remotes.Remote)-1]
    }
    return nil
}

func (remotes *Snmp_Agent_EngineId_Remotes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range remotes.Remote {
        children[remotes.Remote[i].GetSegmentPath()] = &remotes.Remote[i]
    }
    return children
}

func (remotes *Snmp_Agent_EngineId_Remotes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (remotes *Snmp_Agent_EngineId_Remotes) GetBundleName() string { return "cisco_ios_xr" }

func (remotes *Snmp_Agent_EngineId_Remotes) GetYangName() string { return "remotes" }

func (remotes *Snmp_Agent_EngineId_Remotes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remotes *Snmp_Agent_EngineId_Remotes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remotes *Snmp_Agent_EngineId_Remotes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remotes *Snmp_Agent_EngineId_Remotes) SetParent(parent types.Entity) { remotes.parent = parent }

func (remotes *Snmp_Agent_EngineId_Remotes) GetParent() types.Entity { return remotes.parent }

func (remotes *Snmp_Agent_EngineId_Remotes) GetParentYangName() string { return "engine-id" }

// Snmp_Agent_EngineId_Remotes_Remote
// engineID of the remote agent
type Snmp_Agent_EngineId_Remotes_Remote struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of remote SNMP entity. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RemoteAddress interface{}

    // engine ID octet string. The type is string.
    RemoteEngineId interface{}

    // UDP port number. The type is interface{} with range: 1..65535.
    Port interface{}
}

func (remote *Snmp_Agent_EngineId_Remotes_Remote) GetFilter() yfilter.YFilter { return remote.YFilter }

func (remote *Snmp_Agent_EngineId_Remotes_Remote) SetFilter(yf yfilter.YFilter) { remote.YFilter = yf }

func (remote *Snmp_Agent_EngineId_Remotes_Remote) GetGoName(yname string) string {
    if yname == "remote-address" { return "RemoteAddress" }
    if yname == "remote-engine-id" { return "RemoteEngineId" }
    if yname == "port" { return "Port" }
    return ""
}

func (remote *Snmp_Agent_EngineId_Remotes_Remote) GetSegmentPath() string {
    return "remote" + "[remote-address='" + fmt.Sprintf("%v", remote.RemoteAddress) + "']"
}

func (remote *Snmp_Agent_EngineId_Remotes_Remote) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remote *Snmp_Agent_EngineId_Remotes_Remote) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remote *Snmp_Agent_EngineId_Remotes_Remote) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["remote-address"] = remote.RemoteAddress
    leafs["remote-engine-id"] = remote.RemoteEngineId
    leafs["port"] = remote.Port
    return leafs
}

func (remote *Snmp_Agent_EngineId_Remotes_Remote) GetBundleName() string { return "cisco_ios_xr" }

func (remote *Snmp_Agent_EngineId_Remotes_Remote) GetYangName() string { return "remote" }

func (remote *Snmp_Agent_EngineId_Remotes_Remote) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remote *Snmp_Agent_EngineId_Remotes_Remote) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remote *Snmp_Agent_EngineId_Remotes_Remote) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remote *Snmp_Agent_EngineId_Remotes_Remote) SetParent(parent types.Entity) { remote.parent = parent }

func (remote *Snmp_Agent_EngineId_Remotes_Remote) GetParent() types.Entity { return remote.parent }

func (remote *Snmp_Agent_EngineId_Remotes_Remote) GetParentYangName() string { return "remotes" }

// Snmp_Trap
// Class to hold trap configurations
type Snmp_Trap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timeout for TRAP message retransmissions. The type is interface{} with
    // range: 1..1000.
    Timeout interface{}

    // Set throttle time for handling traps. The type is interface{} with range:
    // 10..500. Units are millisecond.
    ThrottleTime interface{}

    // Message queue length for each TRAP host. The type is interface{} with
    // range: 1..5000.
    QueueLength interface{}
}

func (trap *Snmp_Trap) GetFilter() yfilter.YFilter { return trap.YFilter }

func (trap *Snmp_Trap) SetFilter(yf yfilter.YFilter) { trap.YFilter = yf }

func (trap *Snmp_Trap) GetGoName(yname string) string {
    if yname == "timeout" { return "Timeout" }
    if yname == "throttle-time" { return "ThrottleTime" }
    if yname == "queue-length" { return "QueueLength" }
    return ""
}

func (trap *Snmp_Trap) GetSegmentPath() string {
    return "trap"
}

func (trap *Snmp_Trap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trap *Snmp_Trap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trap *Snmp_Trap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timeout"] = trap.Timeout
    leafs["throttle-time"] = trap.ThrottleTime
    leafs["queue-length"] = trap.QueueLength
    return leafs
}

func (trap *Snmp_Trap) GetBundleName() string { return "cisco_ios_xr" }

func (trap *Snmp_Trap) GetYangName() string { return "trap" }

func (trap *Snmp_Trap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trap *Snmp_Trap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trap *Snmp_Trap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trap *Snmp_Trap) SetParent(parent types.Entity) { trap.parent = parent }

func (trap *Snmp_Trap) GetParent() types.Entity { return trap.parent }

func (trap *Snmp_Trap) GetParentYangName() string { return "snmp" }

// Snmp_Ipv6
// SNMP TOS bit for outgoing packets
type Snmp_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type of TOS.
    Tos Snmp_Ipv6_Tos
}

func (ipv6 *Snmp_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Snmp_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Snmp_Ipv6) GetGoName(yname string) string {
    if yname == "tos" { return "Tos" }
    return ""
}

func (ipv6 *Snmp_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Snmp_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tos" {
        return &ipv6.Tos
    }
    return nil
}

func (ipv6 *Snmp_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tos"] = &ipv6.Tos
    return children
}

func (ipv6 *Snmp_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *Snmp_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Snmp_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Snmp_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Snmp_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Snmp_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Snmp_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Snmp_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Snmp_Ipv6) GetParentYangName() string { return "snmp" }

// Snmp_Ipv6_Tos
// Type of TOS
type Snmp_Ipv6_Tos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP TOS type DSCP or Precedence. The type is SnmpTos.
    Type interface{}

    // SNMP Precedence value. The type is one of the following types: enumeration
    // SnmpPrecedenceValue1, or int with range: 0..7.
    Precedence interface{}

    // SNMP DSCP value. The type is one of the following types: enumeration
    // SnmpDscpValue, or int with range: 0..63.
    Dscp interface{}
}

func (tos *Snmp_Ipv6_Tos) GetFilter() yfilter.YFilter { return tos.YFilter }

func (tos *Snmp_Ipv6_Tos) SetFilter(yf yfilter.YFilter) { tos.YFilter = yf }

func (tos *Snmp_Ipv6_Tos) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "precedence" { return "Precedence" }
    if yname == "dscp" { return "Dscp" }
    return ""
}

func (tos *Snmp_Ipv6_Tos) GetSegmentPath() string {
    return "tos"
}

func (tos *Snmp_Ipv6_Tos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tos *Snmp_Ipv6_Tos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tos *Snmp_Ipv6_Tos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = tos.Type
    leafs["precedence"] = tos.Precedence
    leafs["dscp"] = tos.Dscp
    return leafs
}

func (tos *Snmp_Ipv6_Tos) GetBundleName() string { return "cisco_ios_xr" }

func (tos *Snmp_Ipv6_Tos) GetYangName() string { return "tos" }

func (tos *Snmp_Ipv6_Tos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tos *Snmp_Ipv6_Tos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tos *Snmp_Ipv6_Tos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tos *Snmp_Ipv6_Tos) SetParent(parent types.Entity) { tos.parent = parent }

func (tos *Snmp_Ipv6_Tos) GetParent() types.Entity { return tos.parent }

func (tos *Snmp_Ipv6_Tos) GetParentYangName() string { return "ipv6" }

// Snmp_Ipv4
// SNMP TOS bit for outgoing packets
type Snmp_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type of TOS.
    Tos Snmp_Ipv4_Tos
}

func (ipv4 *Snmp_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Snmp_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Snmp_Ipv4) GetGoName(yname string) string {
    if yname == "tos" { return "Tos" }
    return ""
}

func (ipv4 *Snmp_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Snmp_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tos" {
        return &ipv4.Tos
    }
    return nil
}

func (ipv4 *Snmp_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tos"] = &ipv4.Tos
    return children
}

func (ipv4 *Snmp_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *Snmp_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Snmp_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Snmp_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Snmp_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Snmp_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Snmp_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Snmp_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Snmp_Ipv4) GetParentYangName() string { return "snmp" }

// Snmp_Ipv4_Tos
// Type of TOS
type Snmp_Ipv4_Tos struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP TOS type DSCP or Precedence. The type is SnmpTos.
    Type interface{}

    // SNMP Precedence value. The type is one of the following types: enumeration
    // SnmpPrecedenceValue1, or int with range: 0..7.
    Precedence interface{}

    // SNMP DSCP value. The type is one of the following types: enumeration
    // SnmpDscpValue, or int with range: 0..63.
    Dscp interface{}
}

func (tos *Snmp_Ipv4_Tos) GetFilter() yfilter.YFilter { return tos.YFilter }

func (tos *Snmp_Ipv4_Tos) SetFilter(yf yfilter.YFilter) { tos.YFilter = yf }

func (tos *Snmp_Ipv4_Tos) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "precedence" { return "Precedence" }
    if yname == "dscp" { return "Dscp" }
    return ""
}

func (tos *Snmp_Ipv4_Tos) GetSegmentPath() string {
    return "tos"
}

func (tos *Snmp_Ipv4_Tos) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (tos *Snmp_Ipv4_Tos) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (tos *Snmp_Ipv4_Tos) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = tos.Type
    leafs["precedence"] = tos.Precedence
    leafs["dscp"] = tos.Dscp
    return leafs
}

func (tos *Snmp_Ipv4_Tos) GetBundleName() string { return "cisco_ios_xr" }

func (tos *Snmp_Ipv4_Tos) GetYangName() string { return "tos" }

func (tos *Snmp_Ipv4_Tos) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tos *Snmp_Ipv4_Tos) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tos *Snmp_Ipv4_Tos) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tos *Snmp_Ipv4_Tos) SetParent(parent types.Entity) { tos.parent = parent }

func (tos *Snmp_Ipv4_Tos) GetParent() types.Entity { return tos.parent }

func (tos *Snmp_Ipv4_Tos) GetParentYangName() string { return "ipv4" }

// Snmp_System
// container to hold system information
type Snmp_System struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // String to uniquely identify this chassis. The type is string with length:
    // 1..255.
    ChassisId interface{}

    // The physical location of this node. The type is string with length: 1..255.
    Location interface{}

    // identification of the contact person for this managed node. The type is
    // string with length: 1..255.
    Contact interface{}
}

func (system *Snmp_System) GetFilter() yfilter.YFilter { return system.YFilter }

func (system *Snmp_System) SetFilter(yf yfilter.YFilter) { system.YFilter = yf }

func (system *Snmp_System) GetGoName(yname string) string {
    if yname == "chassis-id" { return "ChassisId" }
    if yname == "location" { return "Location" }
    if yname == "contact" { return "Contact" }
    return ""
}

func (system *Snmp_System) GetSegmentPath() string {
    return "system"
}

func (system *Snmp_System) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (system *Snmp_System) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (system *Snmp_System) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["chassis-id"] = system.ChassisId
    leafs["location"] = system.Location
    leafs["contact"] = system.Contact
    return leafs
}

func (system *Snmp_System) GetBundleName() string { return "cisco_ios_xr" }

func (system *Snmp_System) GetYangName() string { return "system" }

func (system *Snmp_System) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (system *Snmp_System) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (system *Snmp_System) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (system *Snmp_System) SetParent(parent types.Entity) { system.parent = parent }

func (system *Snmp_System) GetParent() types.Entity { return system.parent }

func (system *Snmp_System) GetParentYangName() string { return "snmp" }

// Snmp_Target
// SNMP target configurations
type Snmp_Target struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of targets.
    Targets Snmp_Target_Targets
}

func (target *Snmp_Target) GetFilter() yfilter.YFilter { return target.YFilter }

func (target *Snmp_Target) SetFilter(yf yfilter.YFilter) { target.YFilter = yf }

func (target *Snmp_Target) GetGoName(yname string) string {
    if yname == "targets" { return "Targets" }
    return ""
}

func (target *Snmp_Target) GetSegmentPath() string {
    return "target"
}

func (target *Snmp_Target) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "targets" {
        return &target.Targets
    }
    return nil
}

func (target *Snmp_Target) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["targets"] = &target.Targets
    return children
}

func (target *Snmp_Target) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (target *Snmp_Target) GetBundleName() string { return "cisco_ios_xr" }

func (target *Snmp_Target) GetYangName() string { return "target" }

func (target *Snmp_Target) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (target *Snmp_Target) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (target *Snmp_Target) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (target *Snmp_Target) SetParent(parent types.Entity) { target.parent = parent }

func (target *Snmp_Target) GetParent() types.Entity { return target.parent }

func (target *Snmp_Target) GetParentYangName() string { return "snmp" }

// Snmp_Target_Targets
// List of targets
type Snmp_Target_Targets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the target list. The type is slice of Snmp_Target_Targets_Target.
    Target []Snmp_Target_Targets_Target
}

func (targets *Snmp_Target_Targets) GetFilter() yfilter.YFilter { return targets.YFilter }

func (targets *Snmp_Target_Targets) SetFilter(yf yfilter.YFilter) { targets.YFilter = yf }

func (targets *Snmp_Target_Targets) GetGoName(yname string) string {
    if yname == "target" { return "Target" }
    return ""
}

func (targets *Snmp_Target_Targets) GetSegmentPath() string {
    return "targets"
}

func (targets *Snmp_Target_Targets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "target" {
        for _, c := range targets.Target {
            if targets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Target_Targets_Target{}
        targets.Target = append(targets.Target, child)
        return &targets.Target[len(targets.Target)-1]
    }
    return nil
}

func (targets *Snmp_Target_Targets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range targets.Target {
        children[targets.Target[i].GetSegmentPath()] = &targets.Target[i]
    }
    return children
}

func (targets *Snmp_Target_Targets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (targets *Snmp_Target_Targets) GetBundleName() string { return "cisco_ios_xr" }

func (targets *Snmp_Target_Targets) GetYangName() string { return "targets" }

func (targets *Snmp_Target_Targets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (targets *Snmp_Target_Targets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (targets *Snmp_Target_Targets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (targets *Snmp_Target_Targets) SetParent(parent types.Entity) { targets.parent = parent }

func (targets *Snmp_Target_Targets) GetParent() types.Entity { return targets.parent }

func (targets *Snmp_Target_Targets) GetParentYangName() string { return "target" }

// Snmp_Target_Targets_Target
// Name of the target list
type Snmp_Target_Targets_Target struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the target list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    TargetListName interface{}

    // List of VRF Name for a target list.
    VrfNames Snmp_Target_Targets_Target_VrfNames

    // SNMP Target address configurations.
    TargetAddresses Snmp_Target_Targets_Target_TargetAddresses
}

func (target *Snmp_Target_Targets_Target) GetFilter() yfilter.YFilter { return target.YFilter }

func (target *Snmp_Target_Targets_Target) SetFilter(yf yfilter.YFilter) { target.YFilter = yf }

func (target *Snmp_Target_Targets_Target) GetGoName(yname string) string {
    if yname == "target-list-name" { return "TargetListName" }
    if yname == "vrf-names" { return "VrfNames" }
    if yname == "target-addresses" { return "TargetAddresses" }
    return ""
}

func (target *Snmp_Target_Targets_Target) GetSegmentPath() string {
    return "target" + "[target-list-name='" + fmt.Sprintf("%v", target.TargetListName) + "']"
}

func (target *Snmp_Target_Targets_Target) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf-names" {
        return &target.VrfNames
    }
    if childYangName == "target-addresses" {
        return &target.TargetAddresses
    }
    return nil
}

func (target *Snmp_Target_Targets_Target) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrf-names"] = &target.VrfNames
    children["target-addresses"] = &target.TargetAddresses
    return children
}

func (target *Snmp_Target_Targets_Target) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["target-list-name"] = target.TargetListName
    return leafs
}

func (target *Snmp_Target_Targets_Target) GetBundleName() string { return "cisco_ios_xr" }

func (target *Snmp_Target_Targets_Target) GetYangName() string { return "target" }

func (target *Snmp_Target_Targets_Target) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (target *Snmp_Target_Targets_Target) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (target *Snmp_Target_Targets_Target) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (target *Snmp_Target_Targets_Target) SetParent(parent types.Entity) { target.parent = parent }

func (target *Snmp_Target_Targets_Target) GetParent() types.Entity { return target.parent }

func (target *Snmp_Target_Targets_Target) GetParentYangName() string { return "targets" }

// Snmp_Target_Targets_Target_VrfNames
// List of VRF Name for a target list
type Snmp_Target_Targets_Target_VrfNames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name of the target. The type is slice of
    // Snmp_Target_Targets_Target_VrfNames_VrfName.
    VrfName []Snmp_Target_Targets_Target_VrfNames_VrfName
}

func (vrfNames *Snmp_Target_Targets_Target_VrfNames) GetFilter() yfilter.YFilter { return vrfNames.YFilter }

func (vrfNames *Snmp_Target_Targets_Target_VrfNames) SetFilter(yf yfilter.YFilter) { vrfNames.YFilter = yf }

func (vrfNames *Snmp_Target_Targets_Target_VrfNames) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    return ""
}

func (vrfNames *Snmp_Target_Targets_Target_VrfNames) GetSegmentPath() string {
    return "vrf-names"
}

func (vrfNames *Snmp_Target_Targets_Target_VrfNames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf-name" {
        for _, c := range vrfNames.VrfName {
            if vrfNames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Target_Targets_Target_VrfNames_VrfName{}
        vrfNames.VrfName = append(vrfNames.VrfName, child)
        return &vrfNames.VrfName[len(vrfNames.VrfName)-1]
    }
    return nil
}

func (vrfNames *Snmp_Target_Targets_Target_VrfNames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfNames.VrfName {
        children[vrfNames.VrfName[i].GetSegmentPath()] = &vrfNames.VrfName[i]
    }
    return children
}

func (vrfNames *Snmp_Target_Targets_Target_VrfNames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfNames *Snmp_Target_Targets_Target_VrfNames) GetBundleName() string { return "cisco_ios_xr" }

func (vrfNames *Snmp_Target_Targets_Target_VrfNames) GetYangName() string { return "vrf-names" }

func (vrfNames *Snmp_Target_Targets_Target_VrfNames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfNames *Snmp_Target_Targets_Target_VrfNames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfNames *Snmp_Target_Targets_Target_VrfNames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfNames *Snmp_Target_Targets_Target_VrfNames) SetParent(parent types.Entity) { vrfNames.parent = parent }

func (vrfNames *Snmp_Target_Targets_Target_VrfNames) GetParent() types.Entity { return vrfNames.parent }

func (vrfNames *Snmp_Target_Targets_Target_VrfNames) GetParentYangName() string { return "target" }

// Snmp_Target_Targets_Target_VrfNames_VrfName
// VRF name of the target
type Snmp_Target_Targets_Target_VrfNames_VrfName struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}
}

func (vrfName *Snmp_Target_Targets_Target_VrfNames_VrfName) GetFilter() yfilter.YFilter { return vrfName.YFilter }

func (vrfName *Snmp_Target_Targets_Target_VrfNames_VrfName) SetFilter(yf yfilter.YFilter) { vrfName.YFilter = yf }

func (vrfName *Snmp_Target_Targets_Target_VrfNames_VrfName) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    return ""
}

func (vrfName *Snmp_Target_Targets_Target_VrfNames_VrfName) GetSegmentPath() string {
    return "vrf-name" + "[name='" + fmt.Sprintf("%v", vrfName.Name) + "']"
}

func (vrfName *Snmp_Target_Targets_Target_VrfNames_VrfName) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrfName *Snmp_Target_Targets_Target_VrfNames_VrfName) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrfName *Snmp_Target_Targets_Target_VrfNames_VrfName) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = vrfName.Name
    return leafs
}

func (vrfName *Snmp_Target_Targets_Target_VrfNames_VrfName) GetBundleName() string { return "cisco_ios_xr" }

func (vrfName *Snmp_Target_Targets_Target_VrfNames_VrfName) GetYangName() string { return "vrf-name" }

func (vrfName *Snmp_Target_Targets_Target_VrfNames_VrfName) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfName *Snmp_Target_Targets_Target_VrfNames_VrfName) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfName *Snmp_Target_Targets_Target_VrfNames_VrfName) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfName *Snmp_Target_Targets_Target_VrfNames_VrfName) SetParent(parent types.Entity) { vrfName.parent = parent }

func (vrfName *Snmp_Target_Targets_Target_VrfNames_VrfName) GetParent() types.Entity { return vrfName.parent }

func (vrfName *Snmp_Target_Targets_Target_VrfNames_VrfName) GetParentYangName() string { return "vrf-names" }

// Snmp_Target_Targets_Target_TargetAddresses
// SNMP Target address configurations
type Snmp_Target_Targets_Target_TargetAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP Address to be configured for the Target. The type is slice of
    // Snmp_Target_Targets_Target_TargetAddresses_TargetAddress.
    TargetAddress []Snmp_Target_Targets_Target_TargetAddresses_TargetAddress
}

func (targetAddresses *Snmp_Target_Targets_Target_TargetAddresses) GetFilter() yfilter.YFilter { return targetAddresses.YFilter }

func (targetAddresses *Snmp_Target_Targets_Target_TargetAddresses) SetFilter(yf yfilter.YFilter) { targetAddresses.YFilter = yf }

func (targetAddresses *Snmp_Target_Targets_Target_TargetAddresses) GetGoName(yname string) string {
    if yname == "target-address" { return "TargetAddress" }
    return ""
}

func (targetAddresses *Snmp_Target_Targets_Target_TargetAddresses) GetSegmentPath() string {
    return "target-addresses"
}

func (targetAddresses *Snmp_Target_Targets_Target_TargetAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "target-address" {
        for _, c := range targetAddresses.TargetAddress {
            if targetAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Target_Targets_Target_TargetAddresses_TargetAddress{}
        targetAddresses.TargetAddress = append(targetAddresses.TargetAddress, child)
        return &targetAddresses.TargetAddress[len(targetAddresses.TargetAddress)-1]
    }
    return nil
}

func (targetAddresses *Snmp_Target_Targets_Target_TargetAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range targetAddresses.TargetAddress {
        children[targetAddresses.TargetAddress[i].GetSegmentPath()] = &targetAddresses.TargetAddress[i]
    }
    return children
}

func (targetAddresses *Snmp_Target_Targets_Target_TargetAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (targetAddresses *Snmp_Target_Targets_Target_TargetAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (targetAddresses *Snmp_Target_Targets_Target_TargetAddresses) GetYangName() string { return "target-addresses" }

func (targetAddresses *Snmp_Target_Targets_Target_TargetAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (targetAddresses *Snmp_Target_Targets_Target_TargetAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (targetAddresses *Snmp_Target_Targets_Target_TargetAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (targetAddresses *Snmp_Target_Targets_Target_TargetAddresses) SetParent(parent types.Entity) { targetAddresses.parent = parent }

func (targetAddresses *Snmp_Target_Targets_Target_TargetAddresses) GetParent() types.Entity { return targetAddresses.parent }

func (targetAddresses *Snmp_Target_Targets_Target_TargetAddresses) GetParentYangName() string { return "target" }

// Snmp_Target_Targets_Target_TargetAddresses_TargetAddress
// IP Address to be configured for the Target
type Snmp_Target_Targets_Target_TargetAddresses_TargetAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4/Ipv6 address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}
}

func (targetAddress *Snmp_Target_Targets_Target_TargetAddresses_TargetAddress) GetFilter() yfilter.YFilter { return targetAddress.YFilter }

func (targetAddress *Snmp_Target_Targets_Target_TargetAddresses_TargetAddress) SetFilter(yf yfilter.YFilter) { targetAddress.YFilter = yf }

func (targetAddress *Snmp_Target_Targets_Target_TargetAddresses_TargetAddress) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    return ""
}

func (targetAddress *Snmp_Target_Targets_Target_TargetAddresses_TargetAddress) GetSegmentPath() string {
    return "target-address" + "[ip-address='" + fmt.Sprintf("%v", targetAddress.IpAddress) + "']"
}

func (targetAddress *Snmp_Target_Targets_Target_TargetAddresses_TargetAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (targetAddress *Snmp_Target_Targets_Target_TargetAddresses_TargetAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (targetAddress *Snmp_Target_Targets_Target_TargetAddresses_TargetAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = targetAddress.IpAddress
    return leafs
}

func (targetAddress *Snmp_Target_Targets_Target_TargetAddresses_TargetAddress) GetBundleName() string { return "cisco_ios_xr" }

func (targetAddress *Snmp_Target_Targets_Target_TargetAddresses_TargetAddress) GetYangName() string { return "target-address" }

func (targetAddress *Snmp_Target_Targets_Target_TargetAddresses_TargetAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (targetAddress *Snmp_Target_Targets_Target_TargetAddresses_TargetAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (targetAddress *Snmp_Target_Targets_Target_TargetAddresses_TargetAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (targetAddress *Snmp_Target_Targets_Target_TargetAddresses_TargetAddress) SetParent(parent types.Entity) { targetAddress.parent = parent }

func (targetAddress *Snmp_Target_Targets_Target_TargetAddresses_TargetAddress) GetParent() types.Entity { return targetAddress.parent }

func (targetAddress *Snmp_Target_Targets_Target_TargetAddresses_TargetAddress) GetParentYangName() string { return "target-addresses" }

// Snmp_Notification
// Enable SNMP notifications
type Snmp_Notification struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP notification configuration.
    Snmp Snmp_Notification_Snmp

    // Enable SNMP diameter traps.
    Diametermib Snmp_Notification_Diametermib

    // CISCO-IETF-VPLS-GENERIC-MIB notification configuration.
    Vpls Snmp_Notification_Vpls

    // CISCO-IETF-PW-MIB notification configuration.
    L2Vpn Snmp_Notification_L2Vpn

    // Enable ISIS-MIB notifications.
    Isis Snmp_Notification_Isis

    // CISCO-CONFIG-MAN-MIB notification configurations.
    ConfigMan Snmp_Notification_ConfigMan

    // 802.1ag Connectivity Fault Management MIB notification configuration.
    Cfm Snmp_Notification_Cfm

    // 802.3 OAM MIB notification configuration.
    Oam Snmp_Notification_Oam

    // CISCO-FABRIC-HFR-MIB notification configuration.
    FabricCrs Snmp_Notification_FabricCrs

    // CISCO-FLASH-MIB notification configuration.
    Flash Snmp_Notification_Flash

    // Frequency Synchronization trap configuration.
    CiscoIosXrFreqsyncCfgFrequencySynchronization Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization

    // CISCO-ENTITY-REDUNDANCY-MIB notification configuration.
    EntityRedundancy Snmp_Notification_EntityRedundancy

    // CISCO-CONFIG-COPY-MIB notification configuration.
    ConfigCopy Snmp_Notification_ConfigCopy

    // CISCO-SELECTIVE-VRF-DOWNLOAD-MIB notification configuration.
    SelectiveVrfDownload Snmp_Notification_SelectiveVrfDownload

    // CISCO-SYSTEM-MIB notification configuration.
    System Snmp_Notification_System

    // CISCO-IETF-BFD-MIB notification configuration.
    Bfd Snmp_Notification_Bfd

    // Enable SNMP daps traps.
    AddresspoolMib Snmp_Notification_AddresspoolMib

    // CISCO-NTP-MIB notification configuration.
    Ntp Snmp_Notification_Ntp

    // Enable RSVP-MIB notifications.
    Rsvp Snmp_Notification_Rsvp

    // BGP4-MIB and CISCO-BGP4-MIB notification configuration.
    Bgp Snmp_Notification_Bgp

    // CISCO-HSRP-MIB notification configuration.
    Hsrp Snmp_Notification_Hsrp

    // OSPF-MIB notification configuration.
    Ospf Snmp_Notification_Ospf

    // VRRP-MIB notification configuration.
    Vrrp Snmp_Notification_Vrrp

    // OSPFv3-MIB notification configuration.
    Ospfv3 Snmp_Notification_Ospfv3

    // MPLS-LDP-STD-MIB notification configuration.
    MplsLdp Snmp_Notification_MplsLdp

    // CISCO-MPLS-TE-P2MP-STD-MIB notification configuration.
    MplsTeP2Mp Snmp_Notification_MplsTeP2Mp

    // MPLS-TE-STD-MIB notification configuration.
    MplsTe Snmp_Notification_MplsTe

    // CISCO-IETF-FRR-MIB notification configuration.
    MplsFrr Snmp_Notification_MplsFrr

    // MPLS-L3VPN-STD-MIB notification configuration.
    MplsL3Vpn Snmp_Notification_MplsL3Vpn

    // Frequency Synchronization trap configuration.
    CiscoIosXrNcs4KFreqsyncCfgFrequencySynchronization Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization

    // CISCO-OPTICAL-MIB notification configuration.
    Optical Snmp_Notification_Optical

    // CISCO-OPTICAL-OTS-MIB notification configuration.
    OpticalOts Snmp_Notification_OpticalOts

    // CISCO-OTN-IF-MIB notification configuration.
    Otn Snmp_Notification_Otn

    // BRIDGE-MIB notification configuration.
    Bridge Snmp_Notification_Bridge

    // CISCO-ENTITY-SENSOR-MIB notification configuration.
    Sensor Snmp_Notification_Sensor

    // Enable ENTITY-MIB notifications.
    Entity Snmp_Notification_Entity

    // ENTITY-STATE-MIB notification configuration.
    EntityState Snmp_Notification_EntityState

    // CISCO-ENTITY-FRU-CONTROL-MIB notification configuration.
    FruControl Snmp_Notification_FruControl

    // CISCO-RF-MIB notification configuration.
    Rf Snmp_Notification_Rf

    // CISCO-SYSLOG-MIB notification configuration.
    Syslog Snmp_Notification_Syslog

    // Subscriber notification commands.
    SubscriberMib Snmp_Notification_SubscriberMib

    // Enable SNMP l2tun traps.
    L2Tun Snmp_Notification_L2Tun
}

func (notification *Snmp_Notification) GetFilter() yfilter.YFilter { return notification.YFilter }

func (notification *Snmp_Notification) SetFilter(yf yfilter.YFilter) { notification.YFilter = yf }

func (notification *Snmp_Notification) GetGoName(yname string) string {
    if yname == "snmp" { return "Snmp" }
    if yname == "Cisco-IOS-XR-aaa-diameter-base-mib-cfg:diametermib" { return "Diametermib" }
    if yname == "Cisco-IOS-XR-l2vpn-cfg:vpls" { return "Vpls" }
    if yname == "Cisco-IOS-XR-l2vpn-cfg:l2vpn" { return "L2Vpn" }
    if yname == "Cisco-IOS-XR-clns-isis-cfg:isis" { return "Isis" }
    if yname == "Cisco-IOS-XR-config-mibs-cfg:config-man" { return "ConfigMan" }
    if yname == "Cisco-IOS-XR-ethernet-cfm-cfg:cfm" { return "Cfm" }
    if yname == "Cisco-IOS-XR-ethernet-link-oam-cfg:oam" { return "Oam" }
    if yname == "Cisco-IOS-XR-fabhfr-mib-cfg:fabric-crs" { return "FabricCrs" }
    if yname == "Cisco-IOS-XR-flashmib-cfg:flash" { return "Flash" }
    if yname == "Cisco-IOS-XR-freqsync-cfg:frequency-synchronization" { return "CiscoIosXrFreqsyncCfgFrequencySynchronization" }
    if yname == "Cisco-IOS-XR-infra-ceredundancymib-cfg:entity-redundancy" { return "EntityRedundancy" }
    if yname == "Cisco-IOS-XR-infra-confcopymib-cfg:config-copy" { return "ConfigCopy" }
    if yname == "Cisco-IOS-XR-infra-rsi-cfg:selective-vrf-download" { return "SelectiveVrfDownload" }
    if yname == "Cisco-IOS-XR-infra-systemmib-cfg:system" { return "System" }
    if yname == "Cisco-IOS-XR-ip-bfd-cfg:bfd" { return "Bfd" }
    if yname == "Cisco-IOS-XR-ip-daps-mib-cfg:addresspool-mib" { return "AddresspoolMib" }
    if yname == "Cisco-IOS-XR-ip-ntp-cfg:ntp" { return "Ntp" }
    if yname == "Cisco-IOS-XR-ip-rsvp-cfg:rsvp" { return "Rsvp" }
    if yname == "Cisco-IOS-XR-ipv4-bgp-cfg:bgp" { return "Bgp" }
    if yname == "Cisco-IOS-XR-ipv4-hsrp-cfg:hsrp" { return "Hsrp" }
    if yname == "Cisco-IOS-XR-ipv4-ospf-cfg:ospf" { return "Ospf" }
    if yname == "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp" { return "Vrrp" }
    if yname == "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3" { return "Ospfv3" }
    if yname == "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp" { return "MplsLdp" }
    if yname == "Cisco-IOS-XR-mpls-te-cfg:mpls-te-p2mp" { return "MplsTeP2Mp" }
    if yname == "Cisco-IOS-XR-mpls-te-cfg:mpls-te" { return "MplsTe" }
    if yname == "Cisco-IOS-XR-mpls-te-cfg:mpls-frr" { return "MplsFrr" }
    if yname == "Cisco-IOS-XR-mpls-vpn-cfg:mpls-l3vpn" { return "MplsL3Vpn" }
    if yname == "Cisco-IOS-XR-ncs4k-freqsync-cfg:frequency-synchronization" { return "CiscoIosXrNcs4KFreqsyncCfgFrequencySynchronization" }
    if yname == "Cisco-IOS-XR-opticalmib-cfg:optical" { return "Optical" }
    if yname == "Cisco-IOS-XR-opticalotsmib-cfg:optical-ots" { return "OpticalOts" }
    if yname == "Cisco-IOS-XR-otnifmib-cfg:otn" { return "Otn" }
    if yname == "Cisco-IOS-XR-snmp-bridgemib-cfg:bridge" { return "Bridge" }
    if yname == "Cisco-IOS-XR-snmp-ciscosensormib-cfg:sensor" { return "Sensor" }
    if yname == "Cisco-IOS-XR-snmp-entitymib-cfg:entity" { return "Entity" }
    if yname == "Cisco-IOS-XR-snmp-entstatemib-cfg:entity-state" { return "EntityState" }
    if yname == "Cisco-IOS-XR-snmp-frucontrolmib-cfg:fru-control" { return "FruControl" }
    if yname == "Cisco-IOS-XR-snmp-mib-rfmib-cfg:rf" { return "Rf" }
    if yname == "Cisco-IOS-XR-snmp-syslogmib-cfg:syslog" { return "Syslog" }
    if yname == "Cisco-IOS-XR-subscriber-session-mon-mibs-cfg:subscriber-mib" { return "SubscriberMib" }
    if yname == "Cisco-IOS-XR-tunnel-l2tun-proto-mibs-cfg:l2tun" { return "L2Tun" }
    return ""
}

func (notification *Snmp_Notification) GetSegmentPath() string {
    return "notification"
}

func (notification *Snmp_Notification) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "snmp" {
        return &notification.Snmp
    }
    if childYangName == "Cisco-IOS-XR-aaa-diameter-base-mib-cfg:diametermib" {
        return &notification.Diametermib
    }
    if childYangName == "Cisco-IOS-XR-l2vpn-cfg:vpls" {
        return &notification.Vpls
    }
    if childYangName == "Cisco-IOS-XR-l2vpn-cfg:l2vpn" {
        return &notification.L2Vpn
    }
    if childYangName == "Cisco-IOS-XR-clns-isis-cfg:isis" {
        return &notification.Isis
    }
    if childYangName == "Cisco-IOS-XR-config-mibs-cfg:config-man" {
        return &notification.ConfigMan
    }
    if childYangName == "Cisco-IOS-XR-ethernet-cfm-cfg:cfm" {
        return &notification.Cfm
    }
    if childYangName == "Cisco-IOS-XR-ethernet-link-oam-cfg:oam" {
        return &notification.Oam
    }
    if childYangName == "Cisco-IOS-XR-fabhfr-mib-cfg:fabric-crs" {
        return &notification.FabricCrs
    }
    if childYangName == "Cisco-IOS-XR-flashmib-cfg:flash" {
        return &notification.Flash
    }
    if childYangName == "Cisco-IOS-XR-freqsync-cfg:frequency-synchronization" {
        return &notification.CiscoIosXrFreqsyncCfgFrequencySynchronization
    }
    if childYangName == "Cisco-IOS-XR-infra-ceredundancymib-cfg:entity-redundancy" {
        return &notification.EntityRedundancy
    }
    if childYangName == "Cisco-IOS-XR-infra-confcopymib-cfg:config-copy" {
        return &notification.ConfigCopy
    }
    if childYangName == "Cisco-IOS-XR-infra-rsi-cfg:selective-vrf-download" {
        return &notification.SelectiveVrfDownload
    }
    if childYangName == "Cisco-IOS-XR-infra-systemmib-cfg:system" {
        return &notification.System
    }
    if childYangName == "Cisco-IOS-XR-ip-bfd-cfg:bfd" {
        return &notification.Bfd
    }
    if childYangName == "Cisco-IOS-XR-ip-daps-mib-cfg:addresspool-mib" {
        return &notification.AddresspoolMib
    }
    if childYangName == "Cisco-IOS-XR-ip-ntp-cfg:ntp" {
        return &notification.Ntp
    }
    if childYangName == "Cisco-IOS-XR-ip-rsvp-cfg:rsvp" {
        return &notification.Rsvp
    }
    if childYangName == "Cisco-IOS-XR-ipv4-bgp-cfg:bgp" {
        return &notification.Bgp
    }
    if childYangName == "Cisco-IOS-XR-ipv4-hsrp-cfg:hsrp" {
        return &notification.Hsrp
    }
    if childYangName == "Cisco-IOS-XR-ipv4-ospf-cfg:ospf" {
        return &notification.Ospf
    }
    if childYangName == "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp" {
        return &notification.Vrrp
    }
    if childYangName == "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3" {
        return &notification.Ospfv3
    }
    if childYangName == "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp" {
        return &notification.MplsLdp
    }
    if childYangName == "Cisco-IOS-XR-mpls-te-cfg:mpls-te-p2mp" {
        return &notification.MplsTeP2Mp
    }
    if childYangName == "Cisco-IOS-XR-mpls-te-cfg:mpls-te" {
        return &notification.MplsTe
    }
    if childYangName == "Cisco-IOS-XR-mpls-te-cfg:mpls-frr" {
        return &notification.MplsFrr
    }
    if childYangName == "Cisco-IOS-XR-mpls-vpn-cfg:mpls-l3vpn" {
        return &notification.MplsL3Vpn
    }
    if childYangName == "Cisco-IOS-XR-ncs4k-freqsync-cfg:frequency-synchronization" {
        return &notification.CiscoIosXrNcs4KFreqsyncCfgFrequencySynchronization
    }
    if childYangName == "Cisco-IOS-XR-opticalmib-cfg:optical" {
        return &notification.Optical
    }
    if childYangName == "Cisco-IOS-XR-opticalotsmib-cfg:optical-ots" {
        return &notification.OpticalOts
    }
    if childYangName == "Cisco-IOS-XR-otnifmib-cfg:otn" {
        return &notification.Otn
    }
    if childYangName == "Cisco-IOS-XR-snmp-bridgemib-cfg:bridge" {
        return &notification.Bridge
    }
    if childYangName == "Cisco-IOS-XR-snmp-ciscosensormib-cfg:sensor" {
        return &notification.Sensor
    }
    if childYangName == "Cisco-IOS-XR-snmp-entitymib-cfg:entity" {
        return &notification.Entity
    }
    if childYangName == "Cisco-IOS-XR-snmp-entstatemib-cfg:entity-state" {
        return &notification.EntityState
    }
    if childYangName == "Cisco-IOS-XR-snmp-frucontrolmib-cfg:fru-control" {
        return &notification.FruControl
    }
    if childYangName == "Cisco-IOS-XR-snmp-mib-rfmib-cfg:rf" {
        return &notification.Rf
    }
    if childYangName == "Cisco-IOS-XR-snmp-syslogmib-cfg:syslog" {
        return &notification.Syslog
    }
    if childYangName == "Cisco-IOS-XR-subscriber-session-mon-mibs-cfg:subscriber-mib" {
        return &notification.SubscriberMib
    }
    if childYangName == "Cisco-IOS-XR-tunnel-l2tun-proto-mibs-cfg:l2tun" {
        return &notification.L2Tun
    }
    return nil
}

func (notification *Snmp_Notification) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["snmp"] = &notification.Snmp
    children["Cisco-IOS-XR-aaa-diameter-base-mib-cfg:diametermib"] = &notification.Diametermib
    children["Cisco-IOS-XR-l2vpn-cfg:vpls"] = &notification.Vpls
    children["Cisco-IOS-XR-l2vpn-cfg:l2vpn"] = &notification.L2Vpn
    children["Cisco-IOS-XR-clns-isis-cfg:isis"] = &notification.Isis
    children["Cisco-IOS-XR-config-mibs-cfg:config-man"] = &notification.ConfigMan
    children["Cisco-IOS-XR-ethernet-cfm-cfg:cfm"] = &notification.Cfm
    children["Cisco-IOS-XR-ethernet-link-oam-cfg:oam"] = &notification.Oam
    children["Cisco-IOS-XR-fabhfr-mib-cfg:fabric-crs"] = &notification.FabricCrs
    children["Cisco-IOS-XR-flashmib-cfg:flash"] = &notification.Flash
    children["Cisco-IOS-XR-freqsync-cfg:frequency-synchronization"] = &notification.CiscoIosXrFreqsyncCfgFrequencySynchronization
    children["Cisco-IOS-XR-infra-ceredundancymib-cfg:entity-redundancy"] = &notification.EntityRedundancy
    children["Cisco-IOS-XR-infra-confcopymib-cfg:config-copy"] = &notification.ConfigCopy
    children["Cisco-IOS-XR-infra-rsi-cfg:selective-vrf-download"] = &notification.SelectiveVrfDownload
    children["Cisco-IOS-XR-infra-systemmib-cfg:system"] = &notification.System
    children["Cisco-IOS-XR-ip-bfd-cfg:bfd"] = &notification.Bfd
    children["Cisco-IOS-XR-ip-daps-mib-cfg:addresspool-mib"] = &notification.AddresspoolMib
    children["Cisco-IOS-XR-ip-ntp-cfg:ntp"] = &notification.Ntp
    children["Cisco-IOS-XR-ip-rsvp-cfg:rsvp"] = &notification.Rsvp
    children["Cisco-IOS-XR-ipv4-bgp-cfg:bgp"] = &notification.Bgp
    children["Cisco-IOS-XR-ipv4-hsrp-cfg:hsrp"] = &notification.Hsrp
    children["Cisco-IOS-XR-ipv4-ospf-cfg:ospf"] = &notification.Ospf
    children["Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp"] = &notification.Vrrp
    children["Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3"] = &notification.Ospfv3
    children["Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp"] = &notification.MplsLdp
    children["Cisco-IOS-XR-mpls-te-cfg:mpls-te-p2mp"] = &notification.MplsTeP2Mp
    children["Cisco-IOS-XR-mpls-te-cfg:mpls-te"] = &notification.MplsTe
    children["Cisco-IOS-XR-mpls-te-cfg:mpls-frr"] = &notification.MplsFrr
    children["Cisco-IOS-XR-mpls-vpn-cfg:mpls-l3vpn"] = &notification.MplsL3Vpn
    children["Cisco-IOS-XR-ncs4k-freqsync-cfg:frequency-synchronization"] = &notification.CiscoIosXrNcs4KFreqsyncCfgFrequencySynchronization
    children["Cisco-IOS-XR-opticalmib-cfg:optical"] = &notification.Optical
    children["Cisco-IOS-XR-opticalotsmib-cfg:optical-ots"] = &notification.OpticalOts
    children["Cisco-IOS-XR-otnifmib-cfg:otn"] = &notification.Otn
    children["Cisco-IOS-XR-snmp-bridgemib-cfg:bridge"] = &notification.Bridge
    children["Cisco-IOS-XR-snmp-ciscosensormib-cfg:sensor"] = &notification.Sensor
    children["Cisco-IOS-XR-snmp-entitymib-cfg:entity"] = &notification.Entity
    children["Cisco-IOS-XR-snmp-entstatemib-cfg:entity-state"] = &notification.EntityState
    children["Cisco-IOS-XR-snmp-frucontrolmib-cfg:fru-control"] = &notification.FruControl
    children["Cisco-IOS-XR-snmp-mib-rfmib-cfg:rf"] = &notification.Rf
    children["Cisco-IOS-XR-snmp-syslogmib-cfg:syslog"] = &notification.Syslog
    children["Cisco-IOS-XR-subscriber-session-mon-mibs-cfg:subscriber-mib"] = &notification.SubscriberMib
    children["Cisco-IOS-XR-tunnel-l2tun-proto-mibs-cfg:l2tun"] = &notification.L2Tun
    return children
}

func (notification *Snmp_Notification) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (notification *Snmp_Notification) GetBundleName() string { return "cisco_ios_xr" }

func (notification *Snmp_Notification) GetYangName() string { return "notification" }

func (notification *Snmp_Notification) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (notification *Snmp_Notification) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (notification *Snmp_Notification) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (notification *Snmp_Notification) SetParent(parent types.Entity) { notification.parent = parent }

func (notification *Snmp_Notification) GetParent() types.Entity { return notification.parent }

func (notification *Snmp_Notification) GetParentYangName() string { return "snmp" }

// Snmp_Notification_Snmp
// SNMP notification configuration
type Snmp_Notification_Snmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable authentication notification. The type is interface{}.
    Authentication interface{}

    // Enable cold start notification. The type is interface{}.
    ColdStart interface{}

    // Enable warm start notification. The type is interface{}.
    WarmStart interface{}

    // Enable SNMP notifications. The type is interface{}.
    Enable interface{}

    // Enable link down notification. The type is interface{}.
    LinkDown interface{}

    // Enable link up notification. The type is interface{}.
    LinkUp interface{}
}

func (snmp *Snmp_Notification_Snmp) GetFilter() yfilter.YFilter { return snmp.YFilter }

func (snmp *Snmp_Notification_Snmp) SetFilter(yf yfilter.YFilter) { snmp.YFilter = yf }

func (snmp *Snmp_Notification_Snmp) GetGoName(yname string) string {
    if yname == "authentication" { return "Authentication" }
    if yname == "cold-start" { return "ColdStart" }
    if yname == "warm-start" { return "WarmStart" }
    if yname == "enable" { return "Enable" }
    if yname == "link-down" { return "LinkDown" }
    if yname == "link-up" { return "LinkUp" }
    return ""
}

func (snmp *Snmp_Notification_Snmp) GetSegmentPath() string {
    return "snmp"
}

func (snmp *Snmp_Notification_Snmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (snmp *Snmp_Notification_Snmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (snmp *Snmp_Notification_Snmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["authentication"] = snmp.Authentication
    leafs["cold-start"] = snmp.ColdStart
    leafs["warm-start"] = snmp.WarmStart
    leafs["enable"] = snmp.Enable
    leafs["link-down"] = snmp.LinkDown
    leafs["link-up"] = snmp.LinkUp
    return leafs
}

func (snmp *Snmp_Notification_Snmp) GetBundleName() string { return "cisco_ios_xr" }

func (snmp *Snmp_Notification_Snmp) GetYangName() string { return "snmp" }

func (snmp *Snmp_Notification_Snmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (snmp *Snmp_Notification_Snmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (snmp *Snmp_Notification_Snmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (snmp *Snmp_Notification_Snmp) SetParent(parent types.Entity) { snmp.parent = parent }

func (snmp *Snmp_Notification_Snmp) GetParent() types.Entity { return snmp.parent }

func (snmp *Snmp_Notification_Snmp) GetParentYangName() string { return "notification" }

// Snmp_Notification_Diametermib
// Enable SNMP diameter traps
type Snmp_Notification_Diametermib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable SNMP diameter protocol error notification. The type is bool.
    Protocolerror interface{}

    // Enable SNMP diameter permanent failure notification. The type is bool.
    Permanentfail interface{}

    // Enable SNMP diameter peer connection down notification. The type is bool.
    Peerdown interface{}

    // Enable SNMP diameter peer connection up notification. The type is bool.
    Peerup interface{}

    // Enable SNMP diameter transient failure notification. The type is bool.
    Transientfail interface{}
}

func (diametermib *Snmp_Notification_Diametermib) GetFilter() yfilter.YFilter { return diametermib.YFilter }

func (diametermib *Snmp_Notification_Diametermib) SetFilter(yf yfilter.YFilter) { diametermib.YFilter = yf }

func (diametermib *Snmp_Notification_Diametermib) GetGoName(yname string) string {
    if yname == "protocolerror" { return "Protocolerror" }
    if yname == "permanentfail" { return "Permanentfail" }
    if yname == "peerdown" { return "Peerdown" }
    if yname == "peerup" { return "Peerup" }
    if yname == "transientfail" { return "Transientfail" }
    return ""
}

func (diametermib *Snmp_Notification_Diametermib) GetSegmentPath() string {
    return "Cisco-IOS-XR-aaa-diameter-base-mib-cfg:diametermib"
}

func (diametermib *Snmp_Notification_Diametermib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (diametermib *Snmp_Notification_Diametermib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (diametermib *Snmp_Notification_Diametermib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["protocolerror"] = diametermib.Protocolerror
    leafs["permanentfail"] = diametermib.Permanentfail
    leafs["peerdown"] = diametermib.Peerdown
    leafs["peerup"] = diametermib.Peerup
    leafs["transientfail"] = diametermib.Transientfail
    return leafs
}

func (diametermib *Snmp_Notification_Diametermib) GetBundleName() string { return "cisco_ios_xr" }

func (diametermib *Snmp_Notification_Diametermib) GetYangName() string { return "diametermib" }

func (diametermib *Snmp_Notification_Diametermib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (diametermib *Snmp_Notification_Diametermib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (diametermib *Snmp_Notification_Diametermib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (diametermib *Snmp_Notification_Diametermib) SetParent(parent types.Entity) { diametermib.parent = parent }

func (diametermib *Snmp_Notification_Diametermib) GetParent() types.Entity { return diametermib.parent }

func (diametermib *Snmp_Notification_Diametermib) GetParentYangName() string { return "notification" }

// Snmp_Notification_Vpls
// CISCO-IETF-VPLS-GENERIC-MIB notification
// configuration
type Snmp_Notification_Vpls struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable cvplsFwdFullAlarmCleared notification. The type is interface{}.
    FullClear interface{}

    // Enable cvplsStatusChanged notification. The type is interface{}.
    Status interface{}

    // Enable CISCO-IETF-VPLS-GENERIC-MIB notifications. The type is interface{}.
    Enable interface{}

    // Enable cvplsFwdFullAlarmRaised notification. The type is interface{}.
    FullRaise interface{}
}

func (vpls *Snmp_Notification_Vpls) GetFilter() yfilter.YFilter { return vpls.YFilter }

func (vpls *Snmp_Notification_Vpls) SetFilter(yf yfilter.YFilter) { vpls.YFilter = yf }

func (vpls *Snmp_Notification_Vpls) GetGoName(yname string) string {
    if yname == "full-clear" { return "FullClear" }
    if yname == "status" { return "Status" }
    if yname == "enable" { return "Enable" }
    if yname == "full-raise" { return "FullRaise" }
    return ""
}

func (vpls *Snmp_Notification_Vpls) GetSegmentPath() string {
    return "Cisco-IOS-XR-l2vpn-cfg:vpls"
}

func (vpls *Snmp_Notification_Vpls) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vpls *Snmp_Notification_Vpls) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vpls *Snmp_Notification_Vpls) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["full-clear"] = vpls.FullClear
    leafs["status"] = vpls.Status
    leafs["enable"] = vpls.Enable
    leafs["full-raise"] = vpls.FullRaise
    return leafs
}

func (vpls *Snmp_Notification_Vpls) GetBundleName() string { return "cisco_ios_xr" }

func (vpls *Snmp_Notification_Vpls) GetYangName() string { return "vpls" }

func (vpls *Snmp_Notification_Vpls) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vpls *Snmp_Notification_Vpls) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vpls *Snmp_Notification_Vpls) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vpls *Snmp_Notification_Vpls) SetParent(parent types.Entity) { vpls.parent = parent }

func (vpls *Snmp_Notification_Vpls) GetParent() types.Entity { return vpls.parent }

func (vpls *Snmp_Notification_Vpls) GetParentYangName() string { return "notification" }

// Snmp_Notification_L2Vpn
// CISCO-IETF-PW-MIB notification configuration
type Snmp_Notification_L2Vpn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Cisco format including extra varbinds. The type is interface{}.
    Cisco interface{}

    // Enable CISCO-IETF-PW-MIB notifications. The type is interface{}.
    Enable interface{}

    // Enable cpwVcDown notification. The type is interface{}.
    VcDown interface{}

    // Enable cpwVcUp notification. The type is interface{}.
    VcUp interface{}
}

func (l2Vpn *Snmp_Notification_L2Vpn) GetFilter() yfilter.YFilter { return l2Vpn.YFilter }

func (l2Vpn *Snmp_Notification_L2Vpn) SetFilter(yf yfilter.YFilter) { l2Vpn.YFilter = yf }

func (l2Vpn *Snmp_Notification_L2Vpn) GetGoName(yname string) string {
    if yname == "cisco" { return "Cisco" }
    if yname == "enable" { return "Enable" }
    if yname == "vc-down" { return "VcDown" }
    if yname == "vc-up" { return "VcUp" }
    return ""
}

func (l2Vpn *Snmp_Notification_L2Vpn) GetSegmentPath() string {
    return "Cisco-IOS-XR-l2vpn-cfg:l2vpn"
}

func (l2Vpn *Snmp_Notification_L2Vpn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (l2Vpn *Snmp_Notification_L2Vpn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (l2Vpn *Snmp_Notification_L2Vpn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cisco"] = l2Vpn.Cisco
    leafs["enable"] = l2Vpn.Enable
    leafs["vc-down"] = l2Vpn.VcDown
    leafs["vc-up"] = l2Vpn.VcUp
    return leafs
}

func (l2Vpn *Snmp_Notification_L2Vpn) GetBundleName() string { return "cisco_ios_xr" }

func (l2Vpn *Snmp_Notification_L2Vpn) GetYangName() string { return "l2vpn" }

func (l2Vpn *Snmp_Notification_L2Vpn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (l2Vpn *Snmp_Notification_L2Vpn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (l2Vpn *Snmp_Notification_L2Vpn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (l2Vpn *Snmp_Notification_L2Vpn) SetParent(parent types.Entity) { l2Vpn.parent = parent }

func (l2Vpn *Snmp_Notification_L2Vpn) GetParent() types.Entity { return l2Vpn.parent }

func (l2Vpn *Snmp_Notification_L2Vpn) GetParentYangName() string { return "notification" }

// Snmp_Notification_Isis
// Enable ISIS-MIB notifications
type Snmp_Notification_Isis struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable or disable. The type is IsisMibDatabaseOverFlowBoolean. The default
    // value is false.
    DatabaseOverflow interface{}

    // Enable or disable. The type is IsisMibManualAddressDropsBoolean. The
    // default value is false.
    ManualAddressDrops interface{}

    // Enable or disable. The type is IsisMibCorruptedLspDetectedBoolean. The
    // default value is false.
    CorruptedLspDetected interface{}

    // Enable or disable. The type is IsisMibAttemptToExceedMaxSequenceBoolean.
    // The default value is false.
    AttemptToExceedMaxSequence interface{}

    // Enable or disable. The type is IsisMibIdLengthMismatchBoolean. The default
    // value is false.
    IdLengthMismatch interface{}

    // Enable or disable. The type is IsisMibMaxAreaAddressMismatchBoolean. The
    // default value is false.
    MaxAreaAddressMismatch interface{}

    // Enable or disable. The type is IsisMibOwnLspPurgeBoolean. The default value
    // is false.
    OwnLspPurge interface{}

    // Enable or disable. The type is IsisMibSequenceNumberSkipBoolean. The
    // default value is false.
    SequenceNumberSkip interface{}

    // Enable or disable. The type is IsisMibAuthenticationTypeFailureBoolean. The
    // default value is false.
    AuthenticationTypeFailure interface{}

    // Enable or disable. The type is IsisMibAuthenticationFailureBoolean. The
    // default value is false.
    AuthenticationFailure interface{}

    // Enable or disable. The type is IsisMibVersionSkewBoolean. The default value
    // is false.
    VersionSkew interface{}

    // Enable or disable. The type is IsisMibAreaMismatchBoolean. The default
    // value is false.
    AreaMismatch interface{}

    // Enable or disable. The type is IsisMibRejectedAdjacencyBoolean. The default
    // value is false.
    RejectedAdjacency interface{}

    // Enable or disable. The type is IsisMibLspTooLargeToPropagateBoolean. The
    // default value is false.
    LspTooLargeToPropagate interface{}

    // Enable or disable. The type is
    // IsisMibOriginatedLspBufferSizeMismatchBoolean. The default value is false.
    OriginatedLspBufferSizeMismatch interface{}

    // Enable or disable. The type is IsisMibProtocolsSupportedMismatchBoolean.
    // The default value is false.
    ProtocolsSupportedMismatch interface{}

    // Enable or disable. The type is IsisMibAdjacencyChangeBoolean. The default
    // value is false.
    AdjacencyChange interface{}

    // Enable or disable. The type is IsisMibLspErrorDetectedBoolean. The default
    // value is false.
    LspErrorDetected interface{}

    // Enable all isisMIB notifications. The type is IsisMibAllBoolean. The
    // default value is false.
    All interface{}
}

func (isis *Snmp_Notification_Isis) GetFilter() yfilter.YFilter { return isis.YFilter }

func (isis *Snmp_Notification_Isis) SetFilter(yf yfilter.YFilter) { isis.YFilter = yf }

func (isis *Snmp_Notification_Isis) GetGoName(yname string) string {
    if yname == "database-overflow" { return "DatabaseOverflow" }
    if yname == "manual-address-drops" { return "ManualAddressDrops" }
    if yname == "corrupted-lsp-detected" { return "CorruptedLspDetected" }
    if yname == "attempt-to-exceed-max-sequence" { return "AttemptToExceedMaxSequence" }
    if yname == "id-length-mismatch" { return "IdLengthMismatch" }
    if yname == "max-area-address-mismatch" { return "MaxAreaAddressMismatch" }
    if yname == "own-lsp-purge" { return "OwnLspPurge" }
    if yname == "sequence-number-skip" { return "SequenceNumberSkip" }
    if yname == "authentication-type-failure" { return "AuthenticationTypeFailure" }
    if yname == "authentication-failure" { return "AuthenticationFailure" }
    if yname == "version-skew" { return "VersionSkew" }
    if yname == "area-mismatch" { return "AreaMismatch" }
    if yname == "rejected-adjacency" { return "RejectedAdjacency" }
    if yname == "lsp-too-large-to-propagate" { return "LspTooLargeToPropagate" }
    if yname == "originated-lsp-buffer-size-mismatch" { return "OriginatedLspBufferSizeMismatch" }
    if yname == "protocols-supported-mismatch" { return "ProtocolsSupportedMismatch" }
    if yname == "adjacency-change" { return "AdjacencyChange" }
    if yname == "lsp-error-detected" { return "LspErrorDetected" }
    if yname == "all" { return "All" }
    return ""
}

func (isis *Snmp_Notification_Isis) GetSegmentPath() string {
    return "Cisco-IOS-XR-clns-isis-cfg:isis"
}

func (isis *Snmp_Notification_Isis) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (isis *Snmp_Notification_Isis) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (isis *Snmp_Notification_Isis) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["database-overflow"] = isis.DatabaseOverflow
    leafs["manual-address-drops"] = isis.ManualAddressDrops
    leafs["corrupted-lsp-detected"] = isis.CorruptedLspDetected
    leafs["attempt-to-exceed-max-sequence"] = isis.AttemptToExceedMaxSequence
    leafs["id-length-mismatch"] = isis.IdLengthMismatch
    leafs["max-area-address-mismatch"] = isis.MaxAreaAddressMismatch
    leafs["own-lsp-purge"] = isis.OwnLspPurge
    leafs["sequence-number-skip"] = isis.SequenceNumberSkip
    leafs["authentication-type-failure"] = isis.AuthenticationTypeFailure
    leafs["authentication-failure"] = isis.AuthenticationFailure
    leafs["version-skew"] = isis.VersionSkew
    leafs["area-mismatch"] = isis.AreaMismatch
    leafs["rejected-adjacency"] = isis.RejectedAdjacency
    leafs["lsp-too-large-to-propagate"] = isis.LspTooLargeToPropagate
    leafs["originated-lsp-buffer-size-mismatch"] = isis.OriginatedLspBufferSizeMismatch
    leafs["protocols-supported-mismatch"] = isis.ProtocolsSupportedMismatch
    leafs["adjacency-change"] = isis.AdjacencyChange
    leafs["lsp-error-detected"] = isis.LspErrorDetected
    leafs["all"] = isis.All
    return leafs
}

func (isis *Snmp_Notification_Isis) GetBundleName() string { return "cisco_ios_xr" }

func (isis *Snmp_Notification_Isis) GetYangName() string { return "isis" }

func (isis *Snmp_Notification_Isis) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (isis *Snmp_Notification_Isis) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (isis *Snmp_Notification_Isis) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (isis *Snmp_Notification_Isis) SetParent(parent types.Entity) { isis.parent = parent }

func (isis *Snmp_Notification_Isis) GetParent() types.Entity { return isis.parent }

func (isis *Snmp_Notification_Isis) GetParentYangName() string { return "notification" }

// Snmp_Notification_ConfigMan
// CISCO-CONFIG-MAN-MIB notification configurations
type Snmp_Notification_ConfigMan struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ciscoConfigManMIB notifications. The type is interface{}.
    Enable interface{}
}

func (configMan *Snmp_Notification_ConfigMan) GetFilter() yfilter.YFilter { return configMan.YFilter }

func (configMan *Snmp_Notification_ConfigMan) SetFilter(yf yfilter.YFilter) { configMan.YFilter = yf }

func (configMan *Snmp_Notification_ConfigMan) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (configMan *Snmp_Notification_ConfigMan) GetSegmentPath() string {
    return "Cisco-IOS-XR-config-mibs-cfg:config-man"
}

func (configMan *Snmp_Notification_ConfigMan) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configMan *Snmp_Notification_ConfigMan) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configMan *Snmp_Notification_ConfigMan) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = configMan.Enable
    return leafs
}

func (configMan *Snmp_Notification_ConfigMan) GetBundleName() string { return "cisco_ios_xr" }

func (configMan *Snmp_Notification_ConfigMan) GetYangName() string { return "config-man" }

func (configMan *Snmp_Notification_ConfigMan) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configMan *Snmp_Notification_ConfigMan) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configMan *Snmp_Notification_ConfigMan) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configMan *Snmp_Notification_ConfigMan) SetParent(parent types.Entity) { configMan.parent = parent }

func (configMan *Snmp_Notification_ConfigMan) GetParent() types.Entity { return configMan.parent }

func (configMan *Snmp_Notification_ConfigMan) GetParentYangName() string { return "notification" }

// Snmp_Notification_Cfm
// 802.1ag Connectivity Fault Management MIB
// notification configuration
type Snmp_Notification_Cfm struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable 802.1ag Connectivity Fault Management MIB notifications. The type is
    // interface{}.
    Enable interface{}
}

func (cfm *Snmp_Notification_Cfm) GetFilter() yfilter.YFilter { return cfm.YFilter }

func (cfm *Snmp_Notification_Cfm) SetFilter(yf yfilter.YFilter) { cfm.YFilter = yf }

func (cfm *Snmp_Notification_Cfm) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (cfm *Snmp_Notification_Cfm) GetSegmentPath() string {
    return "Cisco-IOS-XR-ethernet-cfm-cfg:cfm"
}

func (cfm *Snmp_Notification_Cfm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cfm *Snmp_Notification_Cfm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cfm *Snmp_Notification_Cfm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = cfm.Enable
    return leafs
}

func (cfm *Snmp_Notification_Cfm) GetBundleName() string { return "cisco_ios_xr" }

func (cfm *Snmp_Notification_Cfm) GetYangName() string { return "cfm" }

func (cfm *Snmp_Notification_Cfm) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cfm *Snmp_Notification_Cfm) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cfm *Snmp_Notification_Cfm) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cfm *Snmp_Notification_Cfm) SetParent(parent types.Entity) { cfm.parent = parent }

func (cfm *Snmp_Notification_Cfm) GetParent() types.Entity { return cfm.parent }

func (cfm *Snmp_Notification_Cfm) GetParentYangName() string { return "notification" }

// Snmp_Notification_Oam
// 802.3 OAM MIB notification configuration
type Snmp_Notification_Oam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable 802.3 OAM MIB notifications. The type is interface{}.
    Enable interface{}
}

func (oam *Snmp_Notification_Oam) GetFilter() yfilter.YFilter { return oam.YFilter }

func (oam *Snmp_Notification_Oam) SetFilter(yf yfilter.YFilter) { oam.YFilter = yf }

func (oam *Snmp_Notification_Oam) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (oam *Snmp_Notification_Oam) GetSegmentPath() string {
    return "Cisco-IOS-XR-ethernet-link-oam-cfg:oam"
}

func (oam *Snmp_Notification_Oam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (oam *Snmp_Notification_Oam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (oam *Snmp_Notification_Oam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = oam.Enable
    return leafs
}

func (oam *Snmp_Notification_Oam) GetBundleName() string { return "cisco_ios_xr" }

func (oam *Snmp_Notification_Oam) GetYangName() string { return "oam" }

func (oam *Snmp_Notification_Oam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (oam *Snmp_Notification_Oam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (oam *Snmp_Notification_Oam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (oam *Snmp_Notification_Oam) SetParent(parent types.Entity) { oam.parent = parent }

func (oam *Snmp_Notification_Oam) GetParent() types.Entity { return oam.parent }

func (oam *Snmp_Notification_Oam) GetParentYangName() string { return "notification" }

// Snmp_Notification_FabricCrs
// CISCO-FABRIC-HFR-MIB notification configuration
type Snmp_Notification_FabricCrs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable cfhBundleStateNotification notification. The type is interface{}.
    BundleState interface{}

    // Enable cfhPlaneStateNotification notification. The type is interface{}.
    PlaneState interface{}

    // Enable cfhBundleDownedLinkNotification notification. The type is
    // interface{}.
    BundleDownedLink interface{}
}

func (fabricCrs *Snmp_Notification_FabricCrs) GetFilter() yfilter.YFilter { return fabricCrs.YFilter }

func (fabricCrs *Snmp_Notification_FabricCrs) SetFilter(yf yfilter.YFilter) { fabricCrs.YFilter = yf }

func (fabricCrs *Snmp_Notification_FabricCrs) GetGoName(yname string) string {
    if yname == "bundle-state" { return "BundleState" }
    if yname == "plane-state" { return "PlaneState" }
    if yname == "bundle-downed-link" { return "BundleDownedLink" }
    return ""
}

func (fabricCrs *Snmp_Notification_FabricCrs) GetSegmentPath() string {
    return "Cisco-IOS-XR-fabhfr-mib-cfg:fabric-crs"
}

func (fabricCrs *Snmp_Notification_FabricCrs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fabricCrs *Snmp_Notification_FabricCrs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fabricCrs *Snmp_Notification_FabricCrs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bundle-state"] = fabricCrs.BundleState
    leafs["plane-state"] = fabricCrs.PlaneState
    leafs["bundle-downed-link"] = fabricCrs.BundleDownedLink
    return leafs
}

func (fabricCrs *Snmp_Notification_FabricCrs) GetBundleName() string { return "cisco_ios_xr" }

func (fabricCrs *Snmp_Notification_FabricCrs) GetYangName() string { return "fabric-crs" }

func (fabricCrs *Snmp_Notification_FabricCrs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fabricCrs *Snmp_Notification_FabricCrs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fabricCrs *Snmp_Notification_FabricCrs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fabricCrs *Snmp_Notification_FabricCrs) SetParent(parent types.Entity) { fabricCrs.parent = parent }

func (fabricCrs *Snmp_Notification_FabricCrs) GetParent() types.Entity { return fabricCrs.parent }

func (fabricCrs *Snmp_Notification_FabricCrs) GetParentYangName() string { return "notification" }

// Snmp_Notification_Flash
// CISCO-FLASH-MIB notification configuration
type Snmp_Notification_Flash struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ciscoFlashDeviceInsertedNotif notification. The type is interface{}.
    Insertion interface{}

    // Enable ciscoFlashDeviceRemovedNotif notification. The type is interface{}.
    Removal interface{}
}

func (flash *Snmp_Notification_Flash) GetFilter() yfilter.YFilter { return flash.YFilter }

func (flash *Snmp_Notification_Flash) SetFilter(yf yfilter.YFilter) { flash.YFilter = yf }

func (flash *Snmp_Notification_Flash) GetGoName(yname string) string {
    if yname == "insertion" { return "Insertion" }
    if yname == "removal" { return "Removal" }
    return ""
}

func (flash *Snmp_Notification_Flash) GetSegmentPath() string {
    return "Cisco-IOS-XR-flashmib-cfg:flash"
}

func (flash *Snmp_Notification_Flash) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (flash *Snmp_Notification_Flash) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (flash *Snmp_Notification_Flash) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["insertion"] = flash.Insertion
    leafs["removal"] = flash.Removal
    return leafs
}

func (flash *Snmp_Notification_Flash) GetBundleName() string { return "cisco_ios_xr" }

func (flash *Snmp_Notification_Flash) GetYangName() string { return "flash" }

func (flash *Snmp_Notification_Flash) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flash *Snmp_Notification_Flash) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flash *Snmp_Notification_Flash) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flash *Snmp_Notification_Flash) SetParent(parent types.Entity) { flash.parent = parent }

func (flash *Snmp_Notification_Flash) GetParent() types.Entity { return flash.parent }

func (flash *Snmp_Notification_Flash) GetParentYangName() string { return "notification" }

// Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization
// Frequency Synchronization trap configuration
type Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Frequency Synchronization traps. The type is interface{}.
    Enable interface{}
}

func (ciscoIOSXRFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization) GetFilter() yfilter.YFilter { return ciscoIOSXRFreqsyncCfgFrequencySynchronization.YFilter }

func (ciscoIOSXRFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization) SetFilter(yf yfilter.YFilter) { ciscoIOSXRFreqsyncCfgFrequencySynchronization.YFilter = yf }

func (ciscoIOSXRFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (ciscoIOSXRFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization) GetSegmentPath() string {
    return "Cisco-IOS-XR-freqsync-cfg:frequency-synchronization"
}

func (ciscoIOSXRFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoIOSXRFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoIOSXRFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = ciscoIOSXRFreqsyncCfgFrequencySynchronization.Enable
    return leafs
}

func (ciscoIOSXRFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization) GetBundleName() string { return "cisco_ios_xr" }

func (ciscoIOSXRFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization) GetYangName() string { return "frequency-synchronization" }

func (ciscoIOSXRFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ciscoIOSXRFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ciscoIOSXRFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ciscoIOSXRFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization) SetParent(parent types.Entity) { ciscoIOSXRFreqsyncCfgFrequencySynchronization.parent = parent }

func (ciscoIOSXRFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization) GetParent() types.Entity { return ciscoIOSXRFreqsyncCfgFrequencySynchronization.parent }

func (ciscoIOSXRFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRFreqsyncCfgFrequencySynchronization) GetParentYangName() string { return "notification" }

// Snmp_Notification_EntityRedundancy
// CISCO-ENTITY-REDUNDANCY-MIB notification
// configuration
type Snmp_Notification_EntityRedundancy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable switchover notification. The type is interface{}.
    Switchover interface{}

    // Enable CISCO-ENTITY-REDUNDANCY-MIB notification. The type is interface{}.
    Enable interface{}

    // Enable status change notification. The type is interface{}.
    Status interface{}
}

func (entityRedundancy *Snmp_Notification_EntityRedundancy) GetFilter() yfilter.YFilter { return entityRedundancy.YFilter }

func (entityRedundancy *Snmp_Notification_EntityRedundancy) SetFilter(yf yfilter.YFilter) { entityRedundancy.YFilter = yf }

func (entityRedundancy *Snmp_Notification_EntityRedundancy) GetGoName(yname string) string {
    if yname == "switchover" { return "Switchover" }
    if yname == "enable" { return "Enable" }
    if yname == "status" { return "Status" }
    return ""
}

func (entityRedundancy *Snmp_Notification_EntityRedundancy) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-ceredundancymib-cfg:entity-redundancy"
}

func (entityRedundancy *Snmp_Notification_EntityRedundancy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entityRedundancy *Snmp_Notification_EntityRedundancy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entityRedundancy *Snmp_Notification_EntityRedundancy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["switchover"] = entityRedundancy.Switchover
    leafs["enable"] = entityRedundancy.Enable
    leafs["status"] = entityRedundancy.Status
    return leafs
}

func (entityRedundancy *Snmp_Notification_EntityRedundancy) GetBundleName() string { return "cisco_ios_xr" }

func (entityRedundancy *Snmp_Notification_EntityRedundancy) GetYangName() string { return "entity-redundancy" }

func (entityRedundancy *Snmp_Notification_EntityRedundancy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entityRedundancy *Snmp_Notification_EntityRedundancy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entityRedundancy *Snmp_Notification_EntityRedundancy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entityRedundancy *Snmp_Notification_EntityRedundancy) SetParent(parent types.Entity) { entityRedundancy.parent = parent }

func (entityRedundancy *Snmp_Notification_EntityRedundancy) GetParent() types.Entity { return entityRedundancy.parent }

func (entityRedundancy *Snmp_Notification_EntityRedundancy) GetParentYangName() string { return "notification" }

// Snmp_Notification_ConfigCopy
// CISCO-CONFIG-COPY-MIB notification configuration
type Snmp_Notification_ConfigCopy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ccCopyCompletion notification. The type is interface{}.
    Completion interface{}
}

func (configCopy *Snmp_Notification_ConfigCopy) GetFilter() yfilter.YFilter { return configCopy.YFilter }

func (configCopy *Snmp_Notification_ConfigCopy) SetFilter(yf yfilter.YFilter) { configCopy.YFilter = yf }

func (configCopy *Snmp_Notification_ConfigCopy) GetGoName(yname string) string {
    if yname == "completion" { return "Completion" }
    return ""
}

func (configCopy *Snmp_Notification_ConfigCopy) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-confcopymib-cfg:config-copy"
}

func (configCopy *Snmp_Notification_ConfigCopy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (configCopy *Snmp_Notification_ConfigCopy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (configCopy *Snmp_Notification_ConfigCopy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["completion"] = configCopy.Completion
    return leafs
}

func (configCopy *Snmp_Notification_ConfigCopy) GetBundleName() string { return "cisco_ios_xr" }

func (configCopy *Snmp_Notification_ConfigCopy) GetYangName() string { return "config-copy" }

func (configCopy *Snmp_Notification_ConfigCopy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (configCopy *Snmp_Notification_ConfigCopy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (configCopy *Snmp_Notification_ConfigCopy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (configCopy *Snmp_Notification_ConfigCopy) SetParent(parent types.Entity) { configCopy.parent = parent }

func (configCopy *Snmp_Notification_ConfigCopy) GetParent() types.Entity { return configCopy.parent }

func (configCopy *Snmp_Notification_ConfigCopy) GetParentYangName() string { return "notification" }

// Snmp_Notification_SelectiveVrfDownload
// CISCO-SELECTIVE-VRF-DOWNLOAD-MIB notification
// configuration
type Snmp_Notification_SelectiveVrfDownload struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable csvdEntityRoleChangeNotification notification. The type is
    // interface{}.
    RoleChange interface{}
}

func (selectiveVrfDownload *Snmp_Notification_SelectiveVrfDownload) GetFilter() yfilter.YFilter { return selectiveVrfDownload.YFilter }

func (selectiveVrfDownload *Snmp_Notification_SelectiveVrfDownload) SetFilter(yf yfilter.YFilter) { selectiveVrfDownload.YFilter = yf }

func (selectiveVrfDownload *Snmp_Notification_SelectiveVrfDownload) GetGoName(yname string) string {
    if yname == "role-change" { return "RoleChange" }
    return ""
}

func (selectiveVrfDownload *Snmp_Notification_SelectiveVrfDownload) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-rsi-cfg:selective-vrf-download"
}

func (selectiveVrfDownload *Snmp_Notification_SelectiveVrfDownload) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (selectiveVrfDownload *Snmp_Notification_SelectiveVrfDownload) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (selectiveVrfDownload *Snmp_Notification_SelectiveVrfDownload) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["role-change"] = selectiveVrfDownload.RoleChange
    return leafs
}

func (selectiveVrfDownload *Snmp_Notification_SelectiveVrfDownload) GetBundleName() string { return "cisco_ios_xr" }

func (selectiveVrfDownload *Snmp_Notification_SelectiveVrfDownload) GetYangName() string { return "selective-vrf-download" }

func (selectiveVrfDownload *Snmp_Notification_SelectiveVrfDownload) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (selectiveVrfDownload *Snmp_Notification_SelectiveVrfDownload) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (selectiveVrfDownload *Snmp_Notification_SelectiveVrfDownload) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (selectiveVrfDownload *Snmp_Notification_SelectiveVrfDownload) SetParent(parent types.Entity) { selectiveVrfDownload.parent = parent }

func (selectiveVrfDownload *Snmp_Notification_SelectiveVrfDownload) GetParent() types.Entity { return selectiveVrfDownload.parent }

func (selectiveVrfDownload *Snmp_Notification_SelectiveVrfDownload) GetParentYangName() string { return "notification" }

// Snmp_Notification_System
// CISCO-SYSTEM-MIB notification configuration
type Snmp_Notification_System struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ciscoSystemMIB notifications. The type is interface{}.
    Enable interface{}
}

func (system *Snmp_Notification_System) GetFilter() yfilter.YFilter { return system.YFilter }

func (system *Snmp_Notification_System) SetFilter(yf yfilter.YFilter) { system.YFilter = yf }

func (system *Snmp_Notification_System) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (system *Snmp_Notification_System) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-systemmib-cfg:system"
}

func (system *Snmp_Notification_System) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (system *Snmp_Notification_System) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (system *Snmp_Notification_System) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = system.Enable
    return leafs
}

func (system *Snmp_Notification_System) GetBundleName() string { return "cisco_ios_xr" }

func (system *Snmp_Notification_System) GetYangName() string { return "system" }

func (system *Snmp_Notification_System) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (system *Snmp_Notification_System) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (system *Snmp_Notification_System) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (system *Snmp_Notification_System) SetParent(parent types.Entity) { system.parent = parent }

func (system *Snmp_Notification_System) GetParent() types.Entity { return system.parent }

func (system *Snmp_Notification_System) GetParentYangName() string { return "notification" }

// Snmp_Notification_Bfd
// CISCO-IETF-BFD-MIB notification configuration
type Snmp_Notification_Bfd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable CISCO-IETF-BFD-MIB notifications. The type is interface{}.
    Enable interface{}
}

func (bfd *Snmp_Notification_Bfd) GetFilter() yfilter.YFilter { return bfd.YFilter }

func (bfd *Snmp_Notification_Bfd) SetFilter(yf yfilter.YFilter) { bfd.YFilter = yf }

func (bfd *Snmp_Notification_Bfd) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (bfd *Snmp_Notification_Bfd) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-bfd-cfg:bfd"
}

func (bfd *Snmp_Notification_Bfd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfd *Snmp_Notification_Bfd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfd *Snmp_Notification_Bfd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = bfd.Enable
    return leafs
}

func (bfd *Snmp_Notification_Bfd) GetBundleName() string { return "cisco_ios_xr" }

func (bfd *Snmp_Notification_Bfd) GetYangName() string { return "bfd" }

func (bfd *Snmp_Notification_Bfd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfd *Snmp_Notification_Bfd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfd *Snmp_Notification_Bfd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfd *Snmp_Notification_Bfd) SetParent(parent types.Entity) { bfd.parent = parent }

func (bfd *Snmp_Notification_Bfd) GetParent() types.Entity { return bfd.parent }

func (bfd *Snmp_Notification_Bfd) GetParentYangName() string { return "notification" }

// Snmp_Notification_AddresspoolMib
// Enable SNMP daps traps
type Snmp_Notification_AddresspoolMib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable SNMP Address Pool High Threshold trap. The type is bool.
    High interface{}

    // Enable SNMP Address Pool Low Threshold trap. The type is bool.
    Low interface{}
}

func (addresspoolMib *Snmp_Notification_AddresspoolMib) GetFilter() yfilter.YFilter { return addresspoolMib.YFilter }

func (addresspoolMib *Snmp_Notification_AddresspoolMib) SetFilter(yf yfilter.YFilter) { addresspoolMib.YFilter = yf }

func (addresspoolMib *Snmp_Notification_AddresspoolMib) GetGoName(yname string) string {
    if yname == "high" { return "High" }
    if yname == "low" { return "Low" }
    return ""
}

func (addresspoolMib *Snmp_Notification_AddresspoolMib) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-daps-mib-cfg:addresspool-mib"
}

func (addresspoolMib *Snmp_Notification_AddresspoolMib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (addresspoolMib *Snmp_Notification_AddresspoolMib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (addresspoolMib *Snmp_Notification_AddresspoolMib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["high"] = addresspoolMib.High
    leafs["low"] = addresspoolMib.Low
    return leafs
}

func (addresspoolMib *Snmp_Notification_AddresspoolMib) GetBundleName() string { return "cisco_ios_xr" }

func (addresspoolMib *Snmp_Notification_AddresspoolMib) GetYangName() string { return "addresspool-mib" }

func (addresspoolMib *Snmp_Notification_AddresspoolMib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (addresspoolMib *Snmp_Notification_AddresspoolMib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (addresspoolMib *Snmp_Notification_AddresspoolMib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (addresspoolMib *Snmp_Notification_AddresspoolMib) SetParent(parent types.Entity) { addresspoolMib.parent = parent }

func (addresspoolMib *Snmp_Notification_AddresspoolMib) GetParent() types.Entity { return addresspoolMib.parent }

func (addresspoolMib *Snmp_Notification_AddresspoolMib) GetParentYangName() string { return "notification" }

// Snmp_Notification_Ntp
// CISCO-NTP-MIB notification configuration
type Snmp_Notification_Ntp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ciscoNtpMIB notification configuration. The type is interface{}.
    Enable interface{}
}

func (ntp *Snmp_Notification_Ntp) GetFilter() yfilter.YFilter { return ntp.YFilter }

func (ntp *Snmp_Notification_Ntp) SetFilter(yf yfilter.YFilter) { ntp.YFilter = yf }

func (ntp *Snmp_Notification_Ntp) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (ntp *Snmp_Notification_Ntp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-ntp-cfg:ntp"
}

func (ntp *Snmp_Notification_Ntp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ntp *Snmp_Notification_Ntp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ntp *Snmp_Notification_Ntp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = ntp.Enable
    return leafs
}

func (ntp *Snmp_Notification_Ntp) GetBundleName() string { return "cisco_ios_xr" }

func (ntp *Snmp_Notification_Ntp) GetYangName() string { return "ntp" }

func (ntp *Snmp_Notification_Ntp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ntp *Snmp_Notification_Ntp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ntp *Snmp_Notification_Ntp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ntp *Snmp_Notification_Ntp) SetParent(parent types.Entity) { ntp.parent = parent }

func (ntp *Snmp_Notification_Ntp) GetParent() types.Entity { return ntp.parent }

func (ntp *Snmp_Notification_Ntp) GetParentYangName() string { return "notification" }

// Snmp_Notification_Rsvp
// Enable RSVP-MIB notifications
type Snmp_Notification_Rsvp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable lostFlow notification. The type is interface{}.
    LostFlow interface{}

    // Enable newFlow notification. The type is interface{}.
    NewFlow interface{}

    // Enable all RSVP notifications. The type is interface{}.
    Enable interface{}
}

func (rsvp *Snmp_Notification_Rsvp) GetFilter() yfilter.YFilter { return rsvp.YFilter }

func (rsvp *Snmp_Notification_Rsvp) SetFilter(yf yfilter.YFilter) { rsvp.YFilter = yf }

func (rsvp *Snmp_Notification_Rsvp) GetGoName(yname string) string {
    if yname == "lost-flow" { return "LostFlow" }
    if yname == "new-flow" { return "NewFlow" }
    if yname == "enable" { return "Enable" }
    return ""
}

func (rsvp *Snmp_Notification_Rsvp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-rsvp-cfg:rsvp"
}

func (rsvp *Snmp_Notification_Rsvp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rsvp *Snmp_Notification_Rsvp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rsvp *Snmp_Notification_Rsvp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lost-flow"] = rsvp.LostFlow
    leafs["new-flow"] = rsvp.NewFlow
    leafs["enable"] = rsvp.Enable
    return leafs
}

func (rsvp *Snmp_Notification_Rsvp) GetBundleName() string { return "cisco_ios_xr" }

func (rsvp *Snmp_Notification_Rsvp) GetYangName() string { return "rsvp" }

func (rsvp *Snmp_Notification_Rsvp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rsvp *Snmp_Notification_Rsvp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rsvp *Snmp_Notification_Rsvp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rsvp *Snmp_Notification_Rsvp) SetParent(parent types.Entity) { rsvp.parent = parent }

func (rsvp *Snmp_Notification_Rsvp) GetParent() types.Entity { return rsvp.parent }

func (rsvp *Snmp_Notification_Rsvp) GetParentYangName() string { return "notification" }

// Snmp_Notification_Bgp
// BGP4-MIB and CISCO-BGP4-MIB notification
// configuration
type Snmp_Notification_Bgp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable BGP4-MIB and CISCO-BGP4-MIB IPv4-only notifications:
    // bgpEstablishedNotification, bgpBackwardTransNotification,
    // cbgpFsmStateChange, cbgpBackwardTransition, cbgpPrefixThresholdExceeded,
    // cbgpPrefixThresholdClear.
    Bgp4Mib Snmp_Notification_Bgp_Bgp4Mib

    // Enable CISCO-BGP4-MIB v2 notifications: cbgpPeer2EstablishedNotification,
    // cbgpPeer2BackwardTransNotification, cbgpPeer2FsmStateChange,
    // cbgpPeer2BackwardTransition, cbgpPeer2PrefixThresholdExceeded,
    // cbgpPeer2PrefixThresholdClear.
    CiscoBgp4Mib Snmp_Notification_Bgp_CiscoBgp4Mib
}

func (bgp *Snmp_Notification_Bgp) GetFilter() yfilter.YFilter { return bgp.YFilter }

func (bgp *Snmp_Notification_Bgp) SetFilter(yf yfilter.YFilter) { bgp.YFilter = yf }

func (bgp *Snmp_Notification_Bgp) GetGoName(yname string) string {
    if yname == "bgp4mib" { return "Bgp4Mib" }
    if yname == "cisco-bgp4mib" { return "CiscoBgp4Mib" }
    return ""
}

func (bgp *Snmp_Notification_Bgp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-bgp-cfg:bgp"
}

func (bgp *Snmp_Notification_Bgp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bgp4mib" {
        return &bgp.Bgp4Mib
    }
    if childYangName == "cisco-bgp4mib" {
        return &bgp.CiscoBgp4Mib
    }
    return nil
}

func (bgp *Snmp_Notification_Bgp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bgp4mib"] = &bgp.Bgp4Mib
    children["cisco-bgp4mib"] = &bgp.CiscoBgp4Mib
    return children
}

func (bgp *Snmp_Notification_Bgp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bgp *Snmp_Notification_Bgp) GetBundleName() string { return "cisco_ios_xr" }

func (bgp *Snmp_Notification_Bgp) GetYangName() string { return "bgp" }

func (bgp *Snmp_Notification_Bgp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp *Snmp_Notification_Bgp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp *Snmp_Notification_Bgp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp *Snmp_Notification_Bgp) SetParent(parent types.Entity) { bgp.parent = parent }

func (bgp *Snmp_Notification_Bgp) GetParent() types.Entity { return bgp.parent }

func (bgp *Snmp_Notification_Bgp) GetParentYangName() string { return "notification" }

// Snmp_Notification_Bgp_Bgp4Mib
// Enable BGP4-MIB and CISCO-BGP4-MIB IPv4-only
// notifications: bgpEstablishedNotification,
// bgpBackwardTransNotification,
// cbgpFsmStateChange, cbgpBackwardTransition,
// cbgpPrefixThresholdExceeded,
// cbgpPrefixThresholdClear.
type Snmp_Notification_Bgp_Bgp4Mib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable BGP4-MIB and CISCO-BGP4-MIB IPv4-only notifications. The type is
    // interface{}.
    Enable interface{}

    // Enable BGP4-MIB and CISCO-BGP4-MIB IPv4-only up/down notifications. The
    // type is interface{}.
    UpDown interface{}
}

func (bgp4Mib *Snmp_Notification_Bgp_Bgp4Mib) GetFilter() yfilter.YFilter { return bgp4Mib.YFilter }

func (bgp4Mib *Snmp_Notification_Bgp_Bgp4Mib) SetFilter(yf yfilter.YFilter) { bgp4Mib.YFilter = yf }

func (bgp4Mib *Snmp_Notification_Bgp_Bgp4Mib) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "up-down" { return "UpDown" }
    return ""
}

func (bgp4Mib *Snmp_Notification_Bgp_Bgp4Mib) GetSegmentPath() string {
    return "bgp4mib"
}

func (bgp4Mib *Snmp_Notification_Bgp_Bgp4Mib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bgp4Mib *Snmp_Notification_Bgp_Bgp4Mib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bgp4Mib *Snmp_Notification_Bgp_Bgp4Mib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = bgp4Mib.Enable
    leafs["up-down"] = bgp4Mib.UpDown
    return leafs
}

func (bgp4Mib *Snmp_Notification_Bgp_Bgp4Mib) GetBundleName() string { return "cisco_ios_xr" }

func (bgp4Mib *Snmp_Notification_Bgp_Bgp4Mib) GetYangName() string { return "bgp4mib" }

func (bgp4Mib *Snmp_Notification_Bgp_Bgp4Mib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bgp4Mib *Snmp_Notification_Bgp_Bgp4Mib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bgp4Mib *Snmp_Notification_Bgp_Bgp4Mib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bgp4Mib *Snmp_Notification_Bgp_Bgp4Mib) SetParent(parent types.Entity) { bgp4Mib.parent = parent }

func (bgp4Mib *Snmp_Notification_Bgp_Bgp4Mib) GetParent() types.Entity { return bgp4Mib.parent }

func (bgp4Mib *Snmp_Notification_Bgp_Bgp4Mib) GetParentYangName() string { return "bgp" }

// Snmp_Notification_Bgp_CiscoBgp4Mib
// Enable CISCO-BGP4-MIB v2 notifications:
// cbgpPeer2EstablishedNotification,
// cbgpPeer2BackwardTransNotification,
// cbgpPeer2FsmStateChange,
// cbgpPeer2BackwardTransition,
// cbgpPeer2PrefixThresholdExceeded,
// cbgpPeer2PrefixThresholdClear.
type Snmp_Notification_Bgp_CiscoBgp4Mib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable CISCO-BGP4-MIB v2 notifications. The type is interface{}.
    Enable interface{}

    // Enable CISCO-BGP4-MIB v2 up/down notifications. The type is interface{}.
    UpDown interface{}
}

func (ciscoBgp4Mib *Snmp_Notification_Bgp_CiscoBgp4Mib) GetFilter() yfilter.YFilter { return ciscoBgp4Mib.YFilter }

func (ciscoBgp4Mib *Snmp_Notification_Bgp_CiscoBgp4Mib) SetFilter(yf yfilter.YFilter) { ciscoBgp4Mib.YFilter = yf }

func (ciscoBgp4Mib *Snmp_Notification_Bgp_CiscoBgp4Mib) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "up-down" { return "UpDown" }
    return ""
}

func (ciscoBgp4Mib *Snmp_Notification_Bgp_CiscoBgp4Mib) GetSegmentPath() string {
    return "cisco-bgp4mib"
}

func (ciscoBgp4Mib *Snmp_Notification_Bgp_CiscoBgp4Mib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoBgp4Mib *Snmp_Notification_Bgp_CiscoBgp4Mib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoBgp4Mib *Snmp_Notification_Bgp_CiscoBgp4Mib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = ciscoBgp4Mib.Enable
    leafs["up-down"] = ciscoBgp4Mib.UpDown
    return leafs
}

func (ciscoBgp4Mib *Snmp_Notification_Bgp_CiscoBgp4Mib) GetBundleName() string { return "cisco_ios_xr" }

func (ciscoBgp4Mib *Snmp_Notification_Bgp_CiscoBgp4Mib) GetYangName() string { return "cisco-bgp4mib" }

func (ciscoBgp4Mib *Snmp_Notification_Bgp_CiscoBgp4Mib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ciscoBgp4Mib *Snmp_Notification_Bgp_CiscoBgp4Mib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ciscoBgp4Mib *Snmp_Notification_Bgp_CiscoBgp4Mib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ciscoBgp4Mib *Snmp_Notification_Bgp_CiscoBgp4Mib) SetParent(parent types.Entity) { ciscoBgp4Mib.parent = parent }

func (ciscoBgp4Mib *Snmp_Notification_Bgp_CiscoBgp4Mib) GetParent() types.Entity { return ciscoBgp4Mib.parent }

func (ciscoBgp4Mib *Snmp_Notification_Bgp_CiscoBgp4Mib) GetParentYangName() string { return "bgp" }

// Snmp_Notification_Hsrp
// CISCO-HSRP-MIB notification configuration
type Snmp_Notification_Hsrp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable CISCO-HSRP-MIB notifications. The type is interface{}.
    Enable interface{}
}

func (hsrp *Snmp_Notification_Hsrp) GetFilter() yfilter.YFilter { return hsrp.YFilter }

func (hsrp *Snmp_Notification_Hsrp) SetFilter(yf yfilter.YFilter) { hsrp.YFilter = yf }

func (hsrp *Snmp_Notification_Hsrp) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (hsrp *Snmp_Notification_Hsrp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-hsrp-cfg:hsrp"
}

func (hsrp *Snmp_Notification_Hsrp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (hsrp *Snmp_Notification_Hsrp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (hsrp *Snmp_Notification_Hsrp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = hsrp.Enable
    return leafs
}

func (hsrp *Snmp_Notification_Hsrp) GetBundleName() string { return "cisco_ios_xr" }

func (hsrp *Snmp_Notification_Hsrp) GetYangName() string { return "hsrp" }

func (hsrp *Snmp_Notification_Hsrp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hsrp *Snmp_Notification_Hsrp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hsrp *Snmp_Notification_Hsrp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hsrp *Snmp_Notification_Hsrp) SetParent(parent types.Entity) { hsrp.parent = parent }

func (hsrp *Snmp_Notification_Hsrp) GetParent() types.Entity { return hsrp.parent }

func (hsrp *Snmp_Notification_Hsrp) GetParentYangName() string { return "notification" }

// Snmp_Notification_Ospf
// OSPF-MIB notification configuration
type Snmp_Notification_Ospf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP notifications related to LSAs.
    Lsa Snmp_Notification_Ospf_Lsa

    // SNMP notifications for OSPF state change.
    StateChange Snmp_Notification_Ospf_StateChange

    // SNMP notifications for packet retransmissions.
    Retransmit Snmp_Notification_Ospf_Retransmit

    // SNMP notifications for OSPF errors.
    Error Snmp_Notification_Ospf_Error
}

func (ospf *Snmp_Notification_Ospf) GetFilter() yfilter.YFilter { return ospf.YFilter }

func (ospf *Snmp_Notification_Ospf) SetFilter(yf yfilter.YFilter) { ospf.YFilter = yf }

func (ospf *Snmp_Notification_Ospf) GetGoName(yname string) string {
    if yname == "lsa" { return "Lsa" }
    if yname == "state-change" { return "StateChange" }
    if yname == "retransmit" { return "Retransmit" }
    if yname == "error" { return "Error" }
    return ""
}

func (ospf *Snmp_Notification_Ospf) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-ospf-cfg:ospf"
}

func (ospf *Snmp_Notification_Ospf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "lsa" {
        return &ospf.Lsa
    }
    if childYangName == "state-change" {
        return &ospf.StateChange
    }
    if childYangName == "retransmit" {
        return &ospf.Retransmit
    }
    if childYangName == "error" {
        return &ospf.Error
    }
    return nil
}

func (ospf *Snmp_Notification_Ospf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["lsa"] = &ospf.Lsa
    children["state-change"] = &ospf.StateChange
    children["retransmit"] = &ospf.Retransmit
    children["error"] = &ospf.Error
    return children
}

func (ospf *Snmp_Notification_Ospf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospf *Snmp_Notification_Ospf) GetBundleName() string { return "cisco_ios_xr" }

func (ospf *Snmp_Notification_Ospf) GetYangName() string { return "ospf" }

func (ospf *Snmp_Notification_Ospf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospf *Snmp_Notification_Ospf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospf *Snmp_Notification_Ospf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospf *Snmp_Notification_Ospf) SetParent(parent types.Entity) { ospf.parent = parent }

func (ospf *Snmp_Notification_Ospf) GetParent() types.Entity { return ospf.parent }

func (ospf *Snmp_Notification_Ospf) GetParentYangName() string { return "notification" }

// Snmp_Notification_Ospf_Lsa
// SNMP notifications related to LSAs
type Snmp_Notification_Ospf_Lsa struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ospfMaxAgeLsa notification. The type is interface{}.
    MaxAgeLsa interface{}

    // Enable ospfOriginateLsa notification. The type is interface{}.
    OriginateLsa interface{}
}

func (lsa *Snmp_Notification_Ospf_Lsa) GetFilter() yfilter.YFilter { return lsa.YFilter }

func (lsa *Snmp_Notification_Ospf_Lsa) SetFilter(yf yfilter.YFilter) { lsa.YFilter = yf }

func (lsa *Snmp_Notification_Ospf_Lsa) GetGoName(yname string) string {
    if yname == "max-age-lsa" { return "MaxAgeLsa" }
    if yname == "originate-lsa" { return "OriginateLsa" }
    return ""
}

func (lsa *Snmp_Notification_Ospf_Lsa) GetSegmentPath() string {
    return "lsa"
}

func (lsa *Snmp_Notification_Ospf_Lsa) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lsa *Snmp_Notification_Ospf_Lsa) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lsa *Snmp_Notification_Ospf_Lsa) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-age-lsa"] = lsa.MaxAgeLsa
    leafs["originate-lsa"] = lsa.OriginateLsa
    return leafs
}

func (lsa *Snmp_Notification_Ospf_Lsa) GetBundleName() string { return "cisco_ios_xr" }

func (lsa *Snmp_Notification_Ospf_Lsa) GetYangName() string { return "lsa" }

func (lsa *Snmp_Notification_Ospf_Lsa) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lsa *Snmp_Notification_Ospf_Lsa) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lsa *Snmp_Notification_Ospf_Lsa) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lsa *Snmp_Notification_Ospf_Lsa) SetParent(parent types.Entity) { lsa.parent = parent }

func (lsa *Snmp_Notification_Ospf_Lsa) GetParent() types.Entity { return lsa.parent }

func (lsa *Snmp_Notification_Ospf_Lsa) GetParentYangName() string { return "ospf" }

// Snmp_Notification_Ospf_StateChange
// SNMP notifications for OSPF state change
type Snmp_Notification_Ospf_StateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ospfIfStateChange notification. The type is interface{}.
    Interface interface{}

    // Enable ospfVirtIfStateChange notification. The type is interface{}.
    VirtualInterface interface{}

    // Enable ospfVirtNbrStateChange notification. The type is interface{}.
    VirtualNeighbor interface{}

    // Enable ospfNbrStateChange notification. The type is interface{}.
    Neighbor interface{}
}

func (stateChange *Snmp_Notification_Ospf_StateChange) GetFilter() yfilter.YFilter { return stateChange.YFilter }

func (stateChange *Snmp_Notification_Ospf_StateChange) SetFilter(yf yfilter.YFilter) { stateChange.YFilter = yf }

func (stateChange *Snmp_Notification_Ospf_StateChange) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    if yname == "virtual-interface" { return "VirtualInterface" }
    if yname == "virtual-neighbor" { return "VirtualNeighbor" }
    if yname == "neighbor" { return "Neighbor" }
    return ""
}

func (stateChange *Snmp_Notification_Ospf_StateChange) GetSegmentPath() string {
    return "state-change"
}

func (stateChange *Snmp_Notification_Ospf_StateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stateChange *Snmp_Notification_Ospf_StateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stateChange *Snmp_Notification_Ospf_StateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface"] = stateChange.Interface
    leafs["virtual-interface"] = stateChange.VirtualInterface
    leafs["virtual-neighbor"] = stateChange.VirtualNeighbor
    leafs["neighbor"] = stateChange.Neighbor
    return leafs
}

func (stateChange *Snmp_Notification_Ospf_StateChange) GetBundleName() string { return "cisco_ios_xr" }

func (stateChange *Snmp_Notification_Ospf_StateChange) GetYangName() string { return "state-change" }

func (stateChange *Snmp_Notification_Ospf_StateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stateChange *Snmp_Notification_Ospf_StateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stateChange *Snmp_Notification_Ospf_StateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stateChange *Snmp_Notification_Ospf_StateChange) SetParent(parent types.Entity) { stateChange.parent = parent }

func (stateChange *Snmp_Notification_Ospf_StateChange) GetParent() types.Entity { return stateChange.parent }

func (stateChange *Snmp_Notification_Ospf_StateChange) GetParentYangName() string { return "ospf" }

// Snmp_Notification_Ospf_Retransmit
// SNMP notifications for packet retransmissions
type Snmp_Notification_Ospf_Retransmit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ospfVirtIfTxRetransmit notification. The type is interface{}.
    VirtualPacket interface{}

    // Enable ospfTxRetransmit notification. The type is interface{}.
    Packet interface{}
}

func (retransmit *Snmp_Notification_Ospf_Retransmit) GetFilter() yfilter.YFilter { return retransmit.YFilter }

func (retransmit *Snmp_Notification_Ospf_Retransmit) SetFilter(yf yfilter.YFilter) { retransmit.YFilter = yf }

func (retransmit *Snmp_Notification_Ospf_Retransmit) GetGoName(yname string) string {
    if yname == "virtual-packet" { return "VirtualPacket" }
    if yname == "packet" { return "Packet" }
    return ""
}

func (retransmit *Snmp_Notification_Ospf_Retransmit) GetSegmentPath() string {
    return "retransmit"
}

func (retransmit *Snmp_Notification_Ospf_Retransmit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (retransmit *Snmp_Notification_Ospf_Retransmit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (retransmit *Snmp_Notification_Ospf_Retransmit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["virtual-packet"] = retransmit.VirtualPacket
    leafs["packet"] = retransmit.Packet
    return leafs
}

func (retransmit *Snmp_Notification_Ospf_Retransmit) GetBundleName() string { return "cisco_ios_xr" }

func (retransmit *Snmp_Notification_Ospf_Retransmit) GetYangName() string { return "retransmit" }

func (retransmit *Snmp_Notification_Ospf_Retransmit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (retransmit *Snmp_Notification_Ospf_Retransmit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (retransmit *Snmp_Notification_Ospf_Retransmit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (retransmit *Snmp_Notification_Ospf_Retransmit) SetParent(parent types.Entity) { retransmit.parent = parent }

func (retransmit *Snmp_Notification_Ospf_Retransmit) GetParent() types.Entity { return retransmit.parent }

func (retransmit *Snmp_Notification_Ospf_Retransmit) GetParentYangName() string { return "ospf" }

// Snmp_Notification_Ospf_Error
// SNMP notifications for OSPF errors
type Snmp_Notification_Ospf_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ospfIfConfigError notification. The type is interface{}.
    ConfigError interface{}

    // Enable ospfIfAuthFailure notification. The type is interface{}.
    AuthenticationFailure interface{}

    // Enable ospfVirtIfConfigError notification. The type is interface{}.
    VirtualConfigError interface{}

    // Enable ospfVirtIfAuthFailure notification. The type is interface{}.
    VirtualAuthenticationFailure interface{}

    // Enable ospfIfRxBadPacket notification. The type is interface{}.
    BadPacket interface{}

    // Enable ospfVirtIfRxBadPacket notification. The type is interface{}.
    VirtualBadPacket interface{}
}

func (error *Snmp_Notification_Ospf_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *Snmp_Notification_Ospf_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *Snmp_Notification_Ospf_Error) GetGoName(yname string) string {
    if yname == "config-error" { return "ConfigError" }
    if yname == "authentication-failure" { return "AuthenticationFailure" }
    if yname == "virtual-config-error" { return "VirtualConfigError" }
    if yname == "virtual-authentication-failure" { return "VirtualAuthenticationFailure" }
    if yname == "bad-packet" { return "BadPacket" }
    if yname == "virtual-bad-packet" { return "VirtualBadPacket" }
    return ""
}

func (error *Snmp_Notification_Ospf_Error) GetSegmentPath() string {
    return "error"
}

func (error *Snmp_Notification_Ospf_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (error *Snmp_Notification_Ospf_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (error *Snmp_Notification_Ospf_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["config-error"] = error.ConfigError
    leafs["authentication-failure"] = error.AuthenticationFailure
    leafs["virtual-config-error"] = error.VirtualConfigError
    leafs["virtual-authentication-failure"] = error.VirtualAuthenticationFailure
    leafs["bad-packet"] = error.BadPacket
    leafs["virtual-bad-packet"] = error.VirtualBadPacket
    return leafs
}

func (error *Snmp_Notification_Ospf_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *Snmp_Notification_Ospf_Error) GetYangName() string { return "error" }

func (error *Snmp_Notification_Ospf_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *Snmp_Notification_Ospf_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *Snmp_Notification_Ospf_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *Snmp_Notification_Ospf_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *Snmp_Notification_Ospf_Error) GetParent() types.Entity { return error.parent }

func (error *Snmp_Notification_Ospf_Error) GetParentYangName() string { return "ospf" }

// Snmp_Notification_Vrrp
// VRRP-MIB notification configuration
type Snmp_Notification_Vrrp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable VRRP-MIB notifications. The type is interface{}.
    Enable interface{}
}

func (vrrp *Snmp_Notification_Vrrp) GetFilter() yfilter.YFilter { return vrrp.YFilter }

func (vrrp *Snmp_Notification_Vrrp) SetFilter(yf yfilter.YFilter) { vrrp.YFilter = yf }

func (vrrp *Snmp_Notification_Vrrp) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (vrrp *Snmp_Notification_Vrrp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp"
}

func (vrrp *Snmp_Notification_Vrrp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrrp *Snmp_Notification_Vrrp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrrp *Snmp_Notification_Vrrp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = vrrp.Enable
    return leafs
}

func (vrrp *Snmp_Notification_Vrrp) GetBundleName() string { return "cisco_ios_xr" }

func (vrrp *Snmp_Notification_Vrrp) GetYangName() string { return "vrrp" }

func (vrrp *Snmp_Notification_Vrrp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrrp *Snmp_Notification_Vrrp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrrp *Snmp_Notification_Vrrp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrrp *Snmp_Notification_Vrrp) SetParent(parent types.Entity) { vrrp.parent = parent }

func (vrrp *Snmp_Notification_Vrrp) GetParent() types.Entity { return vrrp.parent }

func (vrrp *Snmp_Notification_Vrrp) GetParentYangName() string { return "notification" }

// Snmp_Notification_Ospfv3
// OSPFv3-MIB notification configuration
type Snmp_Notification_Ospfv3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SNMP notifications for OSPF errors.
    Error Snmp_Notification_Ospfv3_Error

    // SNMP notifications for OSPF state change.
    StateChange Snmp_Notification_Ospfv3_StateChange
}

func (ospfv3 *Snmp_Notification_Ospfv3) GetFilter() yfilter.YFilter { return ospfv3.YFilter }

func (ospfv3 *Snmp_Notification_Ospfv3) SetFilter(yf yfilter.YFilter) { ospfv3.YFilter = yf }

func (ospfv3 *Snmp_Notification_Ospfv3) GetGoName(yname string) string {
    if yname == "error" { return "Error" }
    if yname == "state-change" { return "StateChange" }
    return ""
}

func (ospfv3 *Snmp_Notification_Ospfv3) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3"
}

func (ospfv3 *Snmp_Notification_Ospfv3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "error" {
        return &ospfv3.Error
    }
    if childYangName == "state-change" {
        return &ospfv3.StateChange
    }
    return nil
}

func (ospfv3 *Snmp_Notification_Ospfv3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["error"] = &ospfv3.Error
    children["state-change"] = &ospfv3.StateChange
    return children
}

func (ospfv3 *Snmp_Notification_Ospfv3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ospfv3 *Snmp_Notification_Ospfv3) GetBundleName() string { return "cisco_ios_xr" }

func (ospfv3 *Snmp_Notification_Ospfv3) GetYangName() string { return "ospfv3" }

func (ospfv3 *Snmp_Notification_Ospfv3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ospfv3 *Snmp_Notification_Ospfv3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ospfv3 *Snmp_Notification_Ospfv3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ospfv3 *Snmp_Notification_Ospfv3) SetParent(parent types.Entity) { ospfv3.parent = parent }

func (ospfv3 *Snmp_Notification_Ospfv3) GetParent() types.Entity { return ospfv3.parent }

func (ospfv3 *Snmp_Notification_Ospfv3) GetParentYangName() string { return "notification" }

// Snmp_Notification_Ospfv3_Error
// SNMP notifications for OSPF errors
type Snmp_Notification_Ospfv3_Error struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ospfv3IfConfigError notification. The type is interface{}.
    ConfigError interface{}

    // Enable ospfv3IfRxBadPacket notification. The type is interface{}.
    BadPacket interface{}

    // Enable ospfv3VirtIfRxBadPacket notification. The type is interface{}.
    VirtualBadPacket interface{}

    // Enable ospfv3VirtIfConfigError notification. The type is interface{}.
    VirtualConfigError interface{}
}

func (error *Snmp_Notification_Ospfv3_Error) GetFilter() yfilter.YFilter { return error.YFilter }

func (error *Snmp_Notification_Ospfv3_Error) SetFilter(yf yfilter.YFilter) { error.YFilter = yf }

func (error *Snmp_Notification_Ospfv3_Error) GetGoName(yname string) string {
    if yname == "config-error" { return "ConfigError" }
    if yname == "bad-packet" { return "BadPacket" }
    if yname == "virtual-bad-packet" { return "VirtualBadPacket" }
    if yname == "virtual-config-error" { return "VirtualConfigError" }
    return ""
}

func (error *Snmp_Notification_Ospfv3_Error) GetSegmentPath() string {
    return "error"
}

func (error *Snmp_Notification_Ospfv3_Error) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (error *Snmp_Notification_Ospfv3_Error) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (error *Snmp_Notification_Ospfv3_Error) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["config-error"] = error.ConfigError
    leafs["bad-packet"] = error.BadPacket
    leafs["virtual-bad-packet"] = error.VirtualBadPacket
    leafs["virtual-config-error"] = error.VirtualConfigError
    return leafs
}

func (error *Snmp_Notification_Ospfv3_Error) GetBundleName() string { return "cisco_ios_xr" }

func (error *Snmp_Notification_Ospfv3_Error) GetYangName() string { return "error" }

func (error *Snmp_Notification_Ospfv3_Error) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (error *Snmp_Notification_Ospfv3_Error) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (error *Snmp_Notification_Ospfv3_Error) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (error *Snmp_Notification_Ospfv3_Error) SetParent(parent types.Entity) { error.parent = parent }

func (error *Snmp_Notification_Ospfv3_Error) GetParent() types.Entity { return error.parent }

func (error *Snmp_Notification_Ospfv3_Error) GetParentYangName() string { return "ospfv3" }

// Snmp_Notification_Ospfv3_StateChange
// SNMP notifications for OSPF state change
type Snmp_Notification_Ospfv3_StateChange struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ospfv3VirtNbrRestartHelperStatusChange notification. The type is
    // interface{}.
    RestartVirtualHelper interface{}

    // Enable ospfv3NssaTranslatorStatusChange notification. The type is
    // interface{}.
    NssaTranslator interface{}

    // Enable ospfv3IfStateChange notification. The type is interface{}.
    Interface interface{}

    // Enable ospfv3RestartStatusChange notification. The type is interface{}.
    Restart interface{}

    // Enable ospfv3NbrStateChange notification. The type is interface{}.
    Neighbor interface{}

    // Enable ospfv3VirtIfStateChange notification. The type is interface{}.
    VirtualInterface interface{}

    // Enable ospfv3NbrRestartHelperStatusChange notification. The type is
    // interface{}.
    RestartHelper interface{}

    // Enable ospfv3VirtNbrStateChange notification. The type is interface{}.
    VirtualNeighbor interface{}
}

func (stateChange *Snmp_Notification_Ospfv3_StateChange) GetFilter() yfilter.YFilter { return stateChange.YFilter }

func (stateChange *Snmp_Notification_Ospfv3_StateChange) SetFilter(yf yfilter.YFilter) { stateChange.YFilter = yf }

func (stateChange *Snmp_Notification_Ospfv3_StateChange) GetGoName(yname string) string {
    if yname == "restart-virtual-helper" { return "RestartVirtualHelper" }
    if yname == "nssa-translator" { return "NssaTranslator" }
    if yname == "interface" { return "Interface" }
    if yname == "restart" { return "Restart" }
    if yname == "neighbor" { return "Neighbor" }
    if yname == "virtual-interface" { return "VirtualInterface" }
    if yname == "restart-helper" { return "RestartHelper" }
    if yname == "virtual-neighbor" { return "VirtualNeighbor" }
    return ""
}

func (stateChange *Snmp_Notification_Ospfv3_StateChange) GetSegmentPath() string {
    return "state-change"
}

func (stateChange *Snmp_Notification_Ospfv3_StateChange) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (stateChange *Snmp_Notification_Ospfv3_StateChange) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (stateChange *Snmp_Notification_Ospfv3_StateChange) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["restart-virtual-helper"] = stateChange.RestartVirtualHelper
    leafs["nssa-translator"] = stateChange.NssaTranslator
    leafs["interface"] = stateChange.Interface
    leafs["restart"] = stateChange.Restart
    leafs["neighbor"] = stateChange.Neighbor
    leafs["virtual-interface"] = stateChange.VirtualInterface
    leafs["restart-helper"] = stateChange.RestartHelper
    leafs["virtual-neighbor"] = stateChange.VirtualNeighbor
    return leafs
}

func (stateChange *Snmp_Notification_Ospfv3_StateChange) GetBundleName() string { return "cisco_ios_xr" }

func (stateChange *Snmp_Notification_Ospfv3_StateChange) GetYangName() string { return "state-change" }

func (stateChange *Snmp_Notification_Ospfv3_StateChange) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (stateChange *Snmp_Notification_Ospfv3_StateChange) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (stateChange *Snmp_Notification_Ospfv3_StateChange) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (stateChange *Snmp_Notification_Ospfv3_StateChange) SetParent(parent types.Entity) { stateChange.parent = parent }

func (stateChange *Snmp_Notification_Ospfv3_StateChange) GetParent() types.Entity { return stateChange.parent }

func (stateChange *Snmp_Notification_Ospfv3_StateChange) GetParentYangName() string { return "ospfv3" }

// Snmp_Notification_MplsLdp
// MPLS-LDP-STD-MIB notification configuration
type Snmp_Notification_MplsLdp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable mplsLdpSessionUp notification. The type is interface{}.
    SessionUp interface{}

    // Enable mplsLdpInitSessionThresholdExceeded notification. The type is
    // interface{}.
    InitSessionThresholdExceeded interface{}

    // Enable mplsLdpSessionDown notification. The type is interface{}.
    SessionDown interface{}
}

func (mplsLdp *Snmp_Notification_MplsLdp) GetFilter() yfilter.YFilter { return mplsLdp.YFilter }

func (mplsLdp *Snmp_Notification_MplsLdp) SetFilter(yf yfilter.YFilter) { mplsLdp.YFilter = yf }

func (mplsLdp *Snmp_Notification_MplsLdp) GetGoName(yname string) string {
    if yname == "session-up" { return "SessionUp" }
    if yname == "init-session-threshold-exceeded" { return "InitSessionThresholdExceeded" }
    if yname == "session-down" { return "SessionDown" }
    return ""
}

func (mplsLdp *Snmp_Notification_MplsLdp) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp"
}

func (mplsLdp *Snmp_Notification_MplsLdp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsLdp *Snmp_Notification_MplsLdp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsLdp *Snmp_Notification_MplsLdp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["session-up"] = mplsLdp.SessionUp
    leafs["init-session-threshold-exceeded"] = mplsLdp.InitSessionThresholdExceeded
    leafs["session-down"] = mplsLdp.SessionDown
    return leafs
}

func (mplsLdp *Snmp_Notification_MplsLdp) GetBundleName() string { return "cisco_ios_xr" }

func (mplsLdp *Snmp_Notification_MplsLdp) GetYangName() string { return "mpls-ldp" }

func (mplsLdp *Snmp_Notification_MplsLdp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsLdp *Snmp_Notification_MplsLdp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsLdp *Snmp_Notification_MplsLdp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsLdp *Snmp_Notification_MplsLdp) SetParent(parent types.Entity) { mplsLdp.parent = parent }

func (mplsLdp *Snmp_Notification_MplsLdp) GetParent() types.Entity { return mplsLdp.parent }

func (mplsLdp *Snmp_Notification_MplsLdp) GetParentYangName() string { return "notification" }

// Snmp_Notification_MplsTeP2Mp
// CISCO-MPLS-TE-P2MP-STD-MIB notification
// configuration
type Snmp_Notification_MplsTeP2Mp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable cmplsTeP2mpTunnelDestUp notification. The type is interface{}.
    Up interface{}

    // Enable cmplsTeP2mpTunnelDestDown notification. The type is interface{}.
    Down interface{}
}

func (mplsTeP2Mp *Snmp_Notification_MplsTeP2Mp) GetFilter() yfilter.YFilter { return mplsTeP2Mp.YFilter }

func (mplsTeP2Mp *Snmp_Notification_MplsTeP2Mp) SetFilter(yf yfilter.YFilter) { mplsTeP2Mp.YFilter = yf }

func (mplsTeP2Mp *Snmp_Notification_MplsTeP2Mp) GetGoName(yname string) string {
    if yname == "up" { return "Up" }
    if yname == "down" { return "Down" }
    return ""
}

func (mplsTeP2Mp *Snmp_Notification_MplsTeP2Mp) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-te-cfg:mpls-te-p2mp"
}

func (mplsTeP2Mp *Snmp_Notification_MplsTeP2Mp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsTeP2Mp *Snmp_Notification_MplsTeP2Mp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsTeP2Mp *Snmp_Notification_MplsTeP2Mp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["up"] = mplsTeP2Mp.Up
    leafs["down"] = mplsTeP2Mp.Down
    return leafs
}

func (mplsTeP2Mp *Snmp_Notification_MplsTeP2Mp) GetBundleName() string { return "cisco_ios_xr" }

func (mplsTeP2Mp *Snmp_Notification_MplsTeP2Mp) GetYangName() string { return "mpls-te-p2mp" }

func (mplsTeP2Mp *Snmp_Notification_MplsTeP2Mp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsTeP2Mp *Snmp_Notification_MplsTeP2Mp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsTeP2Mp *Snmp_Notification_MplsTeP2Mp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsTeP2Mp *Snmp_Notification_MplsTeP2Mp) SetParent(parent types.Entity) { mplsTeP2Mp.parent = parent }

func (mplsTeP2Mp *Snmp_Notification_MplsTeP2Mp) GetParent() types.Entity { return mplsTeP2Mp.parent }

func (mplsTeP2Mp *Snmp_Notification_MplsTeP2Mp) GetParentYangName() string { return "notification" }

// Snmp_Notification_MplsTe
// MPLS-TE-STD-MIB notification configuration
type Snmp_Notification_MplsTe struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable MPLS TE tunnel Cisco format (default IETF) notification. The type is
    // interface{}.
    Cisco interface{}

    // Enable mplsTunnelUp notification. The type is interface{}.
    Up interface{}

    // Enable mplsTunnelReoptimized notification. The type is interface{}.
    Reoptimize interface{}

    // Enable mplsTunnelRerouted notification. The type is interface{}.
    Reroute interface{}

    // Enable mplsTunnelDown notification. The type is interface{}.
    Down interface{}

    // CISCO-MPLS-TE-STD-EXT-MIB notification configuration.
    CiscoExtension Snmp_Notification_MplsTe_CiscoExtension
}

func (mplsTe *Snmp_Notification_MplsTe) GetFilter() yfilter.YFilter { return mplsTe.YFilter }

func (mplsTe *Snmp_Notification_MplsTe) SetFilter(yf yfilter.YFilter) { mplsTe.YFilter = yf }

func (mplsTe *Snmp_Notification_MplsTe) GetGoName(yname string) string {
    if yname == "cisco" { return "Cisco" }
    if yname == "up" { return "Up" }
    if yname == "reoptimize" { return "Reoptimize" }
    if yname == "reroute" { return "Reroute" }
    if yname == "down" { return "Down" }
    if yname == "cisco-extension" { return "CiscoExtension" }
    return ""
}

func (mplsTe *Snmp_Notification_MplsTe) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-te-cfg:mpls-te"
}

func (mplsTe *Snmp_Notification_MplsTe) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cisco-extension" {
        return &mplsTe.CiscoExtension
    }
    return nil
}

func (mplsTe *Snmp_Notification_MplsTe) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cisco-extension"] = &mplsTe.CiscoExtension
    return children
}

func (mplsTe *Snmp_Notification_MplsTe) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cisco"] = mplsTe.Cisco
    leafs["up"] = mplsTe.Up
    leafs["reoptimize"] = mplsTe.Reoptimize
    leafs["reroute"] = mplsTe.Reroute
    leafs["down"] = mplsTe.Down
    return leafs
}

func (mplsTe *Snmp_Notification_MplsTe) GetBundleName() string { return "cisco_ios_xr" }

func (mplsTe *Snmp_Notification_MplsTe) GetYangName() string { return "mpls-te" }

func (mplsTe *Snmp_Notification_MplsTe) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsTe *Snmp_Notification_MplsTe) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsTe *Snmp_Notification_MplsTe) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsTe *Snmp_Notification_MplsTe) SetParent(parent types.Entity) { mplsTe.parent = parent }

func (mplsTe *Snmp_Notification_MplsTe) GetParent() types.Entity { return mplsTe.parent }

func (mplsTe *Snmp_Notification_MplsTe) GetParentYangName() string { return "notification" }

// Snmp_Notification_MplsTe_CiscoExtension
// CISCO-MPLS-TE-STD-EXT-MIB notification
// configuration
type Snmp_Notification_MplsTe_CiscoExtension struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable cmplsTunnelPreempt notification. The type is interface{}.
    Preempt interface{}

    // Enable cmplsTunnelInsuffBW notification. The type is interface{}.
    InsufficientBandwidth interface{}

    // Enable cmplsTunnelReRoutePendingClear notification. The type is
    // interface{}.
    ReRoutePendingClear interface{}

    // Enable cmplsTunnelBringupFail notification. The type is interface{}.
    BringupFail interface{}

    // Enable cmplsTunnelReRoutePending notification. The type is interface{}.
    ReRoutePending interface{}
}

func (ciscoExtension *Snmp_Notification_MplsTe_CiscoExtension) GetFilter() yfilter.YFilter { return ciscoExtension.YFilter }

func (ciscoExtension *Snmp_Notification_MplsTe_CiscoExtension) SetFilter(yf yfilter.YFilter) { ciscoExtension.YFilter = yf }

func (ciscoExtension *Snmp_Notification_MplsTe_CiscoExtension) GetGoName(yname string) string {
    if yname == "preempt" { return "Preempt" }
    if yname == "insufficient-bandwidth" { return "InsufficientBandwidth" }
    if yname == "re-route-pending-clear" { return "ReRoutePendingClear" }
    if yname == "bringup-fail" { return "BringupFail" }
    if yname == "re-route-pending" { return "ReRoutePending" }
    return ""
}

func (ciscoExtension *Snmp_Notification_MplsTe_CiscoExtension) GetSegmentPath() string {
    return "cisco-extension"
}

func (ciscoExtension *Snmp_Notification_MplsTe_CiscoExtension) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoExtension *Snmp_Notification_MplsTe_CiscoExtension) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoExtension *Snmp_Notification_MplsTe_CiscoExtension) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["preempt"] = ciscoExtension.Preempt
    leafs["insufficient-bandwidth"] = ciscoExtension.InsufficientBandwidth
    leafs["re-route-pending-clear"] = ciscoExtension.ReRoutePendingClear
    leafs["bringup-fail"] = ciscoExtension.BringupFail
    leafs["re-route-pending"] = ciscoExtension.ReRoutePending
    return leafs
}

func (ciscoExtension *Snmp_Notification_MplsTe_CiscoExtension) GetBundleName() string { return "cisco_ios_xr" }

func (ciscoExtension *Snmp_Notification_MplsTe_CiscoExtension) GetYangName() string { return "cisco-extension" }

func (ciscoExtension *Snmp_Notification_MplsTe_CiscoExtension) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ciscoExtension *Snmp_Notification_MplsTe_CiscoExtension) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ciscoExtension *Snmp_Notification_MplsTe_CiscoExtension) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ciscoExtension *Snmp_Notification_MplsTe_CiscoExtension) SetParent(parent types.Entity) { ciscoExtension.parent = parent }

func (ciscoExtension *Snmp_Notification_MplsTe_CiscoExtension) GetParent() types.Entity { return ciscoExtension.parent }

func (ciscoExtension *Snmp_Notification_MplsTe_CiscoExtension) GetParentYangName() string { return "mpls-te" }

// Snmp_Notification_MplsFrr
// CISCO-IETF-FRR-MIB notification configuration
type Snmp_Notification_MplsFrr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable cmplsFrrUnProtected notification. The type is interface{}.
    Unprotected interface{}

    // Enable cmplsFrrMIB notifications. The type is interface{}.
    Enable interface{}

    // Enable cmplsFrrProtected notification. The type is interface{}.
    Protected interface{}
}

func (mplsFrr *Snmp_Notification_MplsFrr) GetFilter() yfilter.YFilter { return mplsFrr.YFilter }

func (mplsFrr *Snmp_Notification_MplsFrr) SetFilter(yf yfilter.YFilter) { mplsFrr.YFilter = yf }

func (mplsFrr *Snmp_Notification_MplsFrr) GetGoName(yname string) string {
    if yname == "unprotected" { return "Unprotected" }
    if yname == "enable" { return "Enable" }
    if yname == "protected" { return "Protected" }
    return ""
}

func (mplsFrr *Snmp_Notification_MplsFrr) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-te-cfg:mpls-frr"
}

func (mplsFrr *Snmp_Notification_MplsFrr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsFrr *Snmp_Notification_MplsFrr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsFrr *Snmp_Notification_MplsFrr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["unprotected"] = mplsFrr.Unprotected
    leafs["enable"] = mplsFrr.Enable
    leafs["protected"] = mplsFrr.Protected
    return leafs
}

func (mplsFrr *Snmp_Notification_MplsFrr) GetBundleName() string { return "cisco_ios_xr" }

func (mplsFrr *Snmp_Notification_MplsFrr) GetYangName() string { return "mpls-frr" }

func (mplsFrr *Snmp_Notification_MplsFrr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsFrr *Snmp_Notification_MplsFrr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsFrr *Snmp_Notification_MplsFrr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsFrr *Snmp_Notification_MplsFrr) SetParent(parent types.Entity) { mplsFrr.parent = parent }

func (mplsFrr *Snmp_Notification_MplsFrr) GetParent() types.Entity { return mplsFrr.parent }

func (mplsFrr *Snmp_Notification_MplsFrr) GetParentYangName() string { return "notification" }

// Snmp_Notification_MplsL3Vpn
// MPLS-L3VPN-STD-MIB notification configuration
type Snmp_Notification_MplsL3Vpn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Time interval (secs) for re-issuing max-threshold notification. The type is
    // interface{} with range: -2147483648..2147483647. Units are second.
    MaxThresholdReissueNotificationTime interface{}

    // Enable mplsL3VpnVrfNumVrfRouteMaxThreshExceeded notification. The type is
    // interface{}.
    MaxThresholdExceeded interface{}

    // Enable mplsL3VpnNumVrfRouteMaxThreshCleared notification. The type is
    // interface{}.
    MaxThresholdCleared interface{}

    // Enable mplsL3VpnVrfRouteMidThreshExceeded notification. The type is
    // interface{}.
    MidThresholdExceeded interface{}

    // Enable mplsL3VpnMIB notifications. The type is interface{}.
    Enable interface{}

    // Enable mplsL3VpnVrfDown notification. The type is interface{}.
    VrfDown interface{}

    // Enable mplsL3VpnVrfUp notification. The type is interface{}.
    VrfUp interface{}
}

func (mplsL3Vpn *Snmp_Notification_MplsL3Vpn) GetFilter() yfilter.YFilter { return mplsL3Vpn.YFilter }

func (mplsL3Vpn *Snmp_Notification_MplsL3Vpn) SetFilter(yf yfilter.YFilter) { mplsL3Vpn.YFilter = yf }

func (mplsL3Vpn *Snmp_Notification_MplsL3Vpn) GetGoName(yname string) string {
    if yname == "max-threshold-reissue-notification-time" { return "MaxThresholdReissueNotificationTime" }
    if yname == "max-threshold-exceeded" { return "MaxThresholdExceeded" }
    if yname == "max-threshold-cleared" { return "MaxThresholdCleared" }
    if yname == "mid-threshold-exceeded" { return "MidThresholdExceeded" }
    if yname == "enable" { return "Enable" }
    if yname == "vrf-down" { return "VrfDown" }
    if yname == "vrf-up" { return "VrfUp" }
    return ""
}

func (mplsL3Vpn *Snmp_Notification_MplsL3Vpn) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-vpn-cfg:mpls-l3vpn"
}

func (mplsL3Vpn *Snmp_Notification_MplsL3Vpn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsL3Vpn *Snmp_Notification_MplsL3Vpn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsL3Vpn *Snmp_Notification_MplsL3Vpn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["max-threshold-reissue-notification-time"] = mplsL3Vpn.MaxThresholdReissueNotificationTime
    leafs["max-threshold-exceeded"] = mplsL3Vpn.MaxThresholdExceeded
    leafs["max-threshold-cleared"] = mplsL3Vpn.MaxThresholdCleared
    leafs["mid-threshold-exceeded"] = mplsL3Vpn.MidThresholdExceeded
    leafs["enable"] = mplsL3Vpn.Enable
    leafs["vrf-down"] = mplsL3Vpn.VrfDown
    leafs["vrf-up"] = mplsL3Vpn.VrfUp
    return leafs
}

func (mplsL3Vpn *Snmp_Notification_MplsL3Vpn) GetBundleName() string { return "cisco_ios_xr" }

func (mplsL3Vpn *Snmp_Notification_MplsL3Vpn) GetYangName() string { return "mpls-l3vpn" }

func (mplsL3Vpn *Snmp_Notification_MplsL3Vpn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsL3Vpn *Snmp_Notification_MplsL3Vpn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsL3Vpn *Snmp_Notification_MplsL3Vpn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsL3Vpn *Snmp_Notification_MplsL3Vpn) SetParent(parent types.Entity) { mplsL3Vpn.parent = parent }

func (mplsL3Vpn *Snmp_Notification_MplsL3Vpn) GetParent() types.Entity { return mplsL3Vpn.parent }

func (mplsL3Vpn *Snmp_Notification_MplsL3Vpn) GetParentYangName() string { return "notification" }

// Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization
// Frequency Synchronization trap configuration
type Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Frequency Synchronization traps. The type is interface{}.
    Enable interface{}
}

func (ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization) GetFilter() yfilter.YFilter { return ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization.YFilter }

func (ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization) SetFilter(yf yfilter.YFilter) { ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization.YFilter = yf }

func (ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization) GetSegmentPath() string {
    return "Cisco-IOS-XR-ncs4k-freqsync-cfg:frequency-synchronization"
}

func (ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization.Enable
    return leafs
}

func (ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization) GetBundleName() string { return "cisco_ios_xr" }

func (ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization) GetYangName() string { return "frequency-synchronization" }

func (ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization) SetParent(parent types.Entity) { ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization.parent = parent }

func (ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization) GetParent() types.Entity { return ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization.parent }

func (ciscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization *Snmp_Notification_CiscoIOSXRNcs4KFreqsyncCfgFrequencySynchronization) GetParentYangName() string { return "notification" }

// Snmp_Notification_Optical
// CISCO-OPTICAL-MIB notification configuration
type Snmp_Notification_Optical struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable Opticalmib notifications. The type is interface{}.
    Enable interface{}
}

func (optical *Snmp_Notification_Optical) GetFilter() yfilter.YFilter { return optical.YFilter }

func (optical *Snmp_Notification_Optical) SetFilter(yf yfilter.YFilter) { optical.YFilter = yf }

func (optical *Snmp_Notification_Optical) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (optical *Snmp_Notification_Optical) GetSegmentPath() string {
    return "Cisco-IOS-XR-opticalmib-cfg:optical"
}

func (optical *Snmp_Notification_Optical) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (optical *Snmp_Notification_Optical) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (optical *Snmp_Notification_Optical) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = optical.Enable
    return leafs
}

func (optical *Snmp_Notification_Optical) GetBundleName() string { return "cisco_ios_xr" }

func (optical *Snmp_Notification_Optical) GetYangName() string { return "optical" }

func (optical *Snmp_Notification_Optical) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (optical *Snmp_Notification_Optical) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (optical *Snmp_Notification_Optical) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (optical *Snmp_Notification_Optical) SetParent(parent types.Entity) { optical.parent = parent }

func (optical *Snmp_Notification_Optical) GetParent() types.Entity { return optical.parent }

func (optical *Snmp_Notification_Optical) GetParentYangName() string { return "notification" }

// Snmp_Notification_OpticalOts
// CISCO-OPTICAL-OTS-MIB notification configuration
type Snmp_Notification_OpticalOts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable OpticalOtsmib notifications. The type is interface{}.
    Enable interface{}
}

func (opticalOts *Snmp_Notification_OpticalOts) GetFilter() yfilter.YFilter { return opticalOts.YFilter }

func (opticalOts *Snmp_Notification_OpticalOts) SetFilter(yf yfilter.YFilter) { opticalOts.YFilter = yf }

func (opticalOts *Snmp_Notification_OpticalOts) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (opticalOts *Snmp_Notification_OpticalOts) GetSegmentPath() string {
    return "Cisco-IOS-XR-opticalotsmib-cfg:optical-ots"
}

func (opticalOts *Snmp_Notification_OpticalOts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (opticalOts *Snmp_Notification_OpticalOts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (opticalOts *Snmp_Notification_OpticalOts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = opticalOts.Enable
    return leafs
}

func (opticalOts *Snmp_Notification_OpticalOts) GetBundleName() string { return "cisco_ios_xr" }

func (opticalOts *Snmp_Notification_OpticalOts) GetYangName() string { return "optical-ots" }

func (opticalOts *Snmp_Notification_OpticalOts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (opticalOts *Snmp_Notification_OpticalOts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (opticalOts *Snmp_Notification_OpticalOts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (opticalOts *Snmp_Notification_OpticalOts) SetParent(parent types.Entity) { opticalOts.parent = parent }

func (opticalOts *Snmp_Notification_OpticalOts) GetParent() types.Entity { return opticalOts.parent }

func (opticalOts *Snmp_Notification_OpticalOts) GetParentYangName() string { return "notification" }

// Snmp_Notification_Otn
// CISCO-OTN-IF-MIB notification configuration
type Snmp_Notification_Otn struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ciscoOtnIfMIB notifications. The type is interface{}.
    Enable interface{}
}

func (otn *Snmp_Notification_Otn) GetFilter() yfilter.YFilter { return otn.YFilter }

func (otn *Snmp_Notification_Otn) SetFilter(yf yfilter.YFilter) { otn.YFilter = yf }

func (otn *Snmp_Notification_Otn) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (otn *Snmp_Notification_Otn) GetSegmentPath() string {
    return "Cisco-IOS-XR-otnifmib-cfg:otn"
}

func (otn *Snmp_Notification_Otn) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (otn *Snmp_Notification_Otn) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (otn *Snmp_Notification_Otn) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = otn.Enable
    return leafs
}

func (otn *Snmp_Notification_Otn) GetBundleName() string { return "cisco_ios_xr" }

func (otn *Snmp_Notification_Otn) GetYangName() string { return "otn" }

func (otn *Snmp_Notification_Otn) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (otn *Snmp_Notification_Otn) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (otn *Snmp_Notification_Otn) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (otn *Snmp_Notification_Otn) SetParent(parent types.Entity) { otn.parent = parent }

func (otn *Snmp_Notification_Otn) GetParent() types.Entity { return otn.parent }

func (otn *Snmp_Notification_Otn) GetParentYangName() string { return "notification" }

// Snmp_Notification_Bridge
// BRIDGE-MIB notification configuration
type Snmp_Notification_Bridge struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable dot1dBridge notifications. The type is interface{}.
    Enable interface{}
}

func (bridge *Snmp_Notification_Bridge) GetFilter() yfilter.YFilter { return bridge.YFilter }

func (bridge *Snmp_Notification_Bridge) SetFilter(yf yfilter.YFilter) { bridge.YFilter = yf }

func (bridge *Snmp_Notification_Bridge) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (bridge *Snmp_Notification_Bridge) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-bridgemib-cfg:bridge"
}

func (bridge *Snmp_Notification_Bridge) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bridge *Snmp_Notification_Bridge) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bridge *Snmp_Notification_Bridge) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = bridge.Enable
    return leafs
}

func (bridge *Snmp_Notification_Bridge) GetBundleName() string { return "cisco_ios_xr" }

func (bridge *Snmp_Notification_Bridge) GetYangName() string { return "bridge" }

func (bridge *Snmp_Notification_Bridge) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bridge *Snmp_Notification_Bridge) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bridge *Snmp_Notification_Bridge) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bridge *Snmp_Notification_Bridge) SetParent(parent types.Entity) { bridge.parent = parent }

func (bridge *Snmp_Notification_Bridge) GetParent() types.Entity { return bridge.parent }

func (bridge *Snmp_Notification_Bridge) GetParentYangName() string { return "notification" }

// Snmp_Notification_Sensor
// CISCO-ENTITY-SENSOR-MIB notification
// configuration
type Snmp_Notification_Sensor struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable entitySensorMIB notifications. The type is interface{}.
    Enable interface{}
}

func (sensor *Snmp_Notification_Sensor) GetFilter() yfilter.YFilter { return sensor.YFilter }

func (sensor *Snmp_Notification_Sensor) SetFilter(yf yfilter.YFilter) { sensor.YFilter = yf }

func (sensor *Snmp_Notification_Sensor) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (sensor *Snmp_Notification_Sensor) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-ciscosensormib-cfg:sensor"
}

func (sensor *Snmp_Notification_Sensor) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sensor *Snmp_Notification_Sensor) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sensor *Snmp_Notification_Sensor) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = sensor.Enable
    return leafs
}

func (sensor *Snmp_Notification_Sensor) GetBundleName() string { return "cisco_ios_xr" }

func (sensor *Snmp_Notification_Sensor) GetYangName() string { return "sensor" }

func (sensor *Snmp_Notification_Sensor) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sensor *Snmp_Notification_Sensor) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sensor *Snmp_Notification_Sensor) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sensor *Snmp_Notification_Sensor) SetParent(parent types.Entity) { sensor.parent = parent }

func (sensor *Snmp_Notification_Sensor) GetParent() types.Entity { return sensor.parent }

func (sensor *Snmp_Notification_Sensor) GetParentYangName() string { return "notification" }

// Snmp_Notification_Entity
// Enable ENTITY-MIB notifications
type Snmp_Notification_Entity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable entityMIB notifications. The type is interface{}.
    Enable interface{}
}

func (entity *Snmp_Notification_Entity) GetFilter() yfilter.YFilter { return entity.YFilter }

func (entity *Snmp_Notification_Entity) SetFilter(yf yfilter.YFilter) { entity.YFilter = yf }

func (entity *Snmp_Notification_Entity) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (entity *Snmp_Notification_Entity) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-entitymib-cfg:entity"
}

func (entity *Snmp_Notification_Entity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entity *Snmp_Notification_Entity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entity *Snmp_Notification_Entity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = entity.Enable
    return leafs
}

func (entity *Snmp_Notification_Entity) GetBundleName() string { return "cisco_ios_xr" }

func (entity *Snmp_Notification_Entity) GetYangName() string { return "entity" }

func (entity *Snmp_Notification_Entity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entity *Snmp_Notification_Entity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entity *Snmp_Notification_Entity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entity *Snmp_Notification_Entity) SetParent(parent types.Entity) { entity.parent = parent }

func (entity *Snmp_Notification_Entity) GetParent() types.Entity { return entity.parent }

func (entity *Snmp_Notification_Entity) GetParentYangName() string { return "notification" }

// Snmp_Notification_EntityState
// ENTITY-STATE-MIB notification configuration
type Snmp_Notification_EntityState struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ceStateExtStandbySwitchover notification. The type is interface{}.
    Switchover interface{}

    // Enable entStateOperEnable and entStateOperDisable notifications. The type
    // is interface{}.
    OperStatus interface{}
}

func (entityState *Snmp_Notification_EntityState) GetFilter() yfilter.YFilter { return entityState.YFilter }

func (entityState *Snmp_Notification_EntityState) SetFilter(yf yfilter.YFilter) { entityState.YFilter = yf }

func (entityState *Snmp_Notification_EntityState) GetGoName(yname string) string {
    if yname == "switchover" { return "Switchover" }
    if yname == "oper-status" { return "OperStatus" }
    return ""
}

func (entityState *Snmp_Notification_EntityState) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-entstatemib-cfg:entity-state"
}

func (entityState *Snmp_Notification_EntityState) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entityState *Snmp_Notification_EntityState) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entityState *Snmp_Notification_EntityState) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["switchover"] = entityState.Switchover
    leafs["oper-status"] = entityState.OperStatus
    return leafs
}

func (entityState *Snmp_Notification_EntityState) GetBundleName() string { return "cisco_ios_xr" }

func (entityState *Snmp_Notification_EntityState) GetYangName() string { return "entity-state" }

func (entityState *Snmp_Notification_EntityState) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entityState *Snmp_Notification_EntityState) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entityState *Snmp_Notification_EntityState) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entityState *Snmp_Notification_EntityState) SetParent(parent types.Entity) { entityState.parent = parent }

func (entityState *Snmp_Notification_EntityState) GetParent() types.Entity { return entityState.parent }

func (entityState *Snmp_Notification_EntityState) GetParentYangName() string { return "notification" }

// Snmp_Notification_FruControl
// CISCO-ENTITY-FRU-CONTROL-MIB notification
// configuration
type Snmp_Notification_FruControl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ciscoEntityFRUControlMIB notifications. The type is interface{}.
    Enable interface{}
}

func (fruControl *Snmp_Notification_FruControl) GetFilter() yfilter.YFilter { return fruControl.YFilter }

func (fruControl *Snmp_Notification_FruControl) SetFilter(yf yfilter.YFilter) { fruControl.YFilter = yf }

func (fruControl *Snmp_Notification_FruControl) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (fruControl *Snmp_Notification_FruControl) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-frucontrolmib-cfg:fru-control"
}

func (fruControl *Snmp_Notification_FruControl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (fruControl *Snmp_Notification_FruControl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (fruControl *Snmp_Notification_FruControl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = fruControl.Enable
    return leafs
}

func (fruControl *Snmp_Notification_FruControl) GetBundleName() string { return "cisco_ios_xr" }

func (fruControl *Snmp_Notification_FruControl) GetYangName() string { return "fru-control" }

func (fruControl *Snmp_Notification_FruControl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (fruControl *Snmp_Notification_FruControl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (fruControl *Snmp_Notification_FruControl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (fruControl *Snmp_Notification_FruControl) SetParent(parent types.Entity) { fruControl.parent = parent }

func (fruControl *Snmp_Notification_FruControl) GetParent() types.Entity { return fruControl.parent }

func (fruControl *Snmp_Notification_FruControl) GetParentYangName() string { return "notification" }

// Snmp_Notification_Rf
// CISCO-RF-MIB notification configuration
type Snmp_Notification_Rf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ciscoRFMIB notifications. The type is interface{}.
    Enable interface{}
}

func (rf *Snmp_Notification_Rf) GetFilter() yfilter.YFilter { return rf.YFilter }

func (rf *Snmp_Notification_Rf) SetFilter(yf yfilter.YFilter) { rf.YFilter = yf }

func (rf *Snmp_Notification_Rf) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (rf *Snmp_Notification_Rf) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-mib-rfmib-cfg:rf"
}

func (rf *Snmp_Notification_Rf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rf *Snmp_Notification_Rf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rf *Snmp_Notification_Rf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = rf.Enable
    return leafs
}

func (rf *Snmp_Notification_Rf) GetBundleName() string { return "cisco_ios_xr" }

func (rf *Snmp_Notification_Rf) GetYangName() string { return "rf" }

func (rf *Snmp_Notification_Rf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rf *Snmp_Notification_Rf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rf *Snmp_Notification_Rf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rf *Snmp_Notification_Rf) SetParent(parent types.Entity) { rf.parent = parent }

func (rf *Snmp_Notification_Rf) GetParent() types.Entity { return rf.parent }

func (rf *Snmp_Notification_Rf) GetParentYangName() string { return "notification" }

// Snmp_Notification_Syslog
// CISCO-SYSLOG-MIB notification configuration
type Snmp_Notification_Syslog struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable ciscoSyslogMIB notifications. The type is interface{}.
    Enable interface{}
}

func (syslog *Snmp_Notification_Syslog) GetFilter() yfilter.YFilter { return syslog.YFilter }

func (syslog *Snmp_Notification_Syslog) SetFilter(yf yfilter.YFilter) { syslog.YFilter = yf }

func (syslog *Snmp_Notification_Syslog) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    return ""
}

func (syslog *Snmp_Notification_Syslog) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-syslogmib-cfg:syslog"
}

func (syslog *Snmp_Notification_Syslog) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (syslog *Snmp_Notification_Syslog) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (syslog *Snmp_Notification_Syslog) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = syslog.Enable
    return leafs
}

func (syslog *Snmp_Notification_Syslog) GetBundleName() string { return "cisco_ios_xr" }

func (syslog *Snmp_Notification_Syslog) GetYangName() string { return "syslog" }

func (syslog *Snmp_Notification_Syslog) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (syslog *Snmp_Notification_Syslog) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (syslog *Snmp_Notification_Syslog) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (syslog *Snmp_Notification_Syslog) SetParent(parent types.Entity) { syslog.parent = parent }

func (syslog *Snmp_Notification_Syslog) GetParent() types.Entity { return syslog.parent }

func (syslog *Snmp_Notification_Syslog) GetParentYangName() string { return "notification" }

// Snmp_Notification_SubscriberMib
// Subscriber notification commands
type Snmp_Notification_SubscriberMib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Session aggregation.
    SessionAggregate Snmp_Notification_SubscriberMib_SessionAggregate
}

func (subscriberMib *Snmp_Notification_SubscriberMib) GetFilter() yfilter.YFilter { return subscriberMib.YFilter }

func (subscriberMib *Snmp_Notification_SubscriberMib) SetFilter(yf yfilter.YFilter) { subscriberMib.YFilter = yf }

func (subscriberMib *Snmp_Notification_SubscriberMib) GetGoName(yname string) string {
    if yname == "session-aggregate" { return "SessionAggregate" }
    return ""
}

func (subscriberMib *Snmp_Notification_SubscriberMib) GetSegmentPath() string {
    return "Cisco-IOS-XR-subscriber-session-mon-mibs-cfg:subscriber-mib"
}

func (subscriberMib *Snmp_Notification_SubscriberMib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "session-aggregate" {
        return &subscriberMib.SessionAggregate
    }
    return nil
}

func (subscriberMib *Snmp_Notification_SubscriberMib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["session-aggregate"] = &subscriberMib.SessionAggregate
    return children
}

func (subscriberMib *Snmp_Notification_SubscriberMib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscriberMib *Snmp_Notification_SubscriberMib) GetBundleName() string { return "cisco_ios_xr" }

func (subscriberMib *Snmp_Notification_SubscriberMib) GetYangName() string { return "subscriber-mib" }

func (subscriberMib *Snmp_Notification_SubscriberMib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriberMib *Snmp_Notification_SubscriberMib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriberMib *Snmp_Notification_SubscriberMib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriberMib *Snmp_Notification_SubscriberMib) SetParent(parent types.Entity) { subscriberMib.parent = parent }

func (subscriberMib *Snmp_Notification_SubscriberMib) GetParent() types.Entity { return subscriberMib.parent }

func (subscriberMib *Snmp_Notification_SubscriberMib) GetParentYangName() string { return "notification" }

// Snmp_Notification_SubscriberMib_SessionAggregate
// Session aggregation
type Snmp_Notification_SubscriberMib_SessionAggregate struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber notification at node level. The type is interface{}.
    Node interface{}

    // Subscriber notification at access interface level. The type is interface{}.
    AccessInterface interface{}
}

func (sessionAggregate *Snmp_Notification_SubscriberMib_SessionAggregate) GetFilter() yfilter.YFilter { return sessionAggregate.YFilter }

func (sessionAggregate *Snmp_Notification_SubscriberMib_SessionAggregate) SetFilter(yf yfilter.YFilter) { sessionAggregate.YFilter = yf }

func (sessionAggregate *Snmp_Notification_SubscriberMib_SessionAggregate) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    if yname == "access-interface" { return "AccessInterface" }
    return ""
}

func (sessionAggregate *Snmp_Notification_SubscriberMib_SessionAggregate) GetSegmentPath() string {
    return "session-aggregate"
}

func (sessionAggregate *Snmp_Notification_SubscriberMib_SessionAggregate) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sessionAggregate *Snmp_Notification_SubscriberMib_SessionAggregate) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sessionAggregate *Snmp_Notification_SubscriberMib_SessionAggregate) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node"] = sessionAggregate.Node
    leafs["access-interface"] = sessionAggregate.AccessInterface
    return leafs
}

func (sessionAggregate *Snmp_Notification_SubscriberMib_SessionAggregate) GetBundleName() string { return "cisco_ios_xr" }

func (sessionAggregate *Snmp_Notification_SubscriberMib_SessionAggregate) GetYangName() string { return "session-aggregate" }

func (sessionAggregate *Snmp_Notification_SubscriberMib_SessionAggregate) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessionAggregate *Snmp_Notification_SubscriberMib_SessionAggregate) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessionAggregate *Snmp_Notification_SubscriberMib_SessionAggregate) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessionAggregate *Snmp_Notification_SubscriberMib_SessionAggregate) SetParent(parent types.Entity) { sessionAggregate.parent = parent }

func (sessionAggregate *Snmp_Notification_SubscriberMib_SessionAggregate) GetParent() types.Entity { return sessionAggregate.parent }

func (sessionAggregate *Snmp_Notification_SubscriberMib_SessionAggregate) GetParentYangName() string { return "subscriber-mib" }

// Snmp_Notification_L2Tun
// Enable SNMP l2tun traps
type Snmp_Notification_L2Tun struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable L2TUN tunnel UP traps. The type is bool.
    TunnelUp interface{}

    // Enable L2TUN tunnel DOWN traps. The type is bool.
    TunnelDown interface{}

    // Enable traps for L2TPv3 PW status. The type is bool.
    PseudowireStatus interface{}

    // Enable L2TUN sessions traps. The type is bool.
    Sessions interface{}
}

func (l2Tun *Snmp_Notification_L2Tun) GetFilter() yfilter.YFilter { return l2Tun.YFilter }

func (l2Tun *Snmp_Notification_L2Tun) SetFilter(yf yfilter.YFilter) { l2Tun.YFilter = yf }

func (l2Tun *Snmp_Notification_L2Tun) GetGoName(yname string) string {
    if yname == "tunnel-up" { return "TunnelUp" }
    if yname == "tunnel-down" { return "TunnelDown" }
    if yname == "pseudowire-status" { return "PseudowireStatus" }
    if yname == "sessions" { return "Sessions" }
    return ""
}

func (l2Tun *Snmp_Notification_L2Tun) GetSegmentPath() string {
    return "Cisco-IOS-XR-tunnel-l2tun-proto-mibs-cfg:l2tun"
}

func (l2Tun *Snmp_Notification_L2Tun) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (l2Tun *Snmp_Notification_L2Tun) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (l2Tun *Snmp_Notification_L2Tun) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["tunnel-up"] = l2Tun.TunnelUp
    leafs["tunnel-down"] = l2Tun.TunnelDown
    leafs["pseudowire-status"] = l2Tun.PseudowireStatus
    leafs["sessions"] = l2Tun.Sessions
    return leafs
}

func (l2Tun *Snmp_Notification_L2Tun) GetBundleName() string { return "cisco_ios_xr" }

func (l2Tun *Snmp_Notification_L2Tun) GetYangName() string { return "l2tun" }

func (l2Tun *Snmp_Notification_L2Tun) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (l2Tun *Snmp_Notification_L2Tun) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (l2Tun *Snmp_Notification_L2Tun) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (l2Tun *Snmp_Notification_L2Tun) SetParent(parent types.Entity) { l2Tun.parent = parent }

func (l2Tun *Snmp_Notification_L2Tun) GetParent() types.Entity { return l2Tun.parent }

func (l2Tun *Snmp_Notification_L2Tun) GetParentYangName() string { return "notification" }

// Snmp_Correlator
// Configure properties of the trap correlator
type Snmp_Correlator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure size of the correlator buffer. The type is interface{} with
    // range: 1024..52428800.
    BufferSize interface{}

    // Table of configured rules.
    Rules Snmp_Correlator_Rules

    // Table of configured rulesets.
    RuleSets Snmp_Correlator_RuleSets
}

func (correlator *Snmp_Correlator) GetFilter() yfilter.YFilter { return correlator.YFilter }

func (correlator *Snmp_Correlator) SetFilter(yf yfilter.YFilter) { correlator.YFilter = yf }

func (correlator *Snmp_Correlator) GetGoName(yname string) string {
    if yname == "buffer-size" { return "BufferSize" }
    if yname == "rules" { return "Rules" }
    if yname == "rule-sets" { return "RuleSets" }
    return ""
}

func (correlator *Snmp_Correlator) GetSegmentPath() string {
    return "correlator"
}

func (correlator *Snmp_Correlator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rules" {
        return &correlator.Rules
    }
    if childYangName == "rule-sets" {
        return &correlator.RuleSets
    }
    return nil
}

func (correlator *Snmp_Correlator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rules"] = &correlator.Rules
    children["rule-sets"] = &correlator.RuleSets
    return children
}

func (correlator *Snmp_Correlator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["buffer-size"] = correlator.BufferSize
    return leafs
}

func (correlator *Snmp_Correlator) GetBundleName() string { return "cisco_ios_xr" }

func (correlator *Snmp_Correlator) GetYangName() string { return "correlator" }

func (correlator *Snmp_Correlator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (correlator *Snmp_Correlator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (correlator *Snmp_Correlator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (correlator *Snmp_Correlator) SetParent(parent types.Entity) { correlator.parent = parent }

func (correlator *Snmp_Correlator) GetParent() types.Entity { return correlator.parent }

func (correlator *Snmp_Correlator) GetParentYangName() string { return "snmp" }

// Snmp_Correlator_Rules
// Table of configured rules
type Snmp_Correlator_Rules struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rule name. The type is slice of Snmp_Correlator_Rules_Rule.
    Rule []Snmp_Correlator_Rules_Rule
}

func (rules *Snmp_Correlator_Rules) GetFilter() yfilter.YFilter { return rules.YFilter }

func (rules *Snmp_Correlator_Rules) SetFilter(yf yfilter.YFilter) { rules.YFilter = yf }

func (rules *Snmp_Correlator_Rules) GetGoName(yname string) string {
    if yname == "rule" { return "Rule" }
    return ""
}

func (rules *Snmp_Correlator_Rules) GetSegmentPath() string {
    return "rules"
}

func (rules *Snmp_Correlator_Rules) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule" {
        for _, c := range rules.Rule {
            if rules.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_Rules_Rule{}
        rules.Rule = append(rules.Rule, child)
        return &rules.Rule[len(rules.Rule)-1]
    }
    return nil
}

func (rules *Snmp_Correlator_Rules) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rules.Rule {
        children[rules.Rule[i].GetSegmentPath()] = &rules.Rule[i]
    }
    return children
}

func (rules *Snmp_Correlator_Rules) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rules *Snmp_Correlator_Rules) GetBundleName() string { return "cisco_ios_xr" }

func (rules *Snmp_Correlator_Rules) GetYangName() string { return "rules" }

func (rules *Snmp_Correlator_Rules) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rules *Snmp_Correlator_Rules) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rules *Snmp_Correlator_Rules) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rules *Snmp_Correlator_Rules) SetParent(parent types.Entity) { rules.parent = parent }

func (rules *Snmp_Correlator_Rules) GetParent() types.Entity { return rules.parent }

func (rules *Snmp_Correlator_Rules) GetParentYangName() string { return "correlator" }

// Snmp_Correlator_Rules_Rule
// Rule name
type Snmp_Correlator_Rules_Rule struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rule name. The type is string with length: 1..32.
    Name interface{}

    // The Non-Stateful Rule Type.
    NonStateful Snmp_Correlator_Rules_Rule_NonStateful

    // Applied to the Rule or Ruleset.
    AppliedTo Snmp_Correlator_Rules_Rule_AppliedTo
}

func (rule *Snmp_Correlator_Rules_Rule) GetFilter() yfilter.YFilter { return rule.YFilter }

func (rule *Snmp_Correlator_Rules_Rule) SetFilter(yf yfilter.YFilter) { rule.YFilter = yf }

func (rule *Snmp_Correlator_Rules_Rule) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "non-stateful" { return "NonStateful" }
    if yname == "applied-to" { return "AppliedTo" }
    return ""
}

func (rule *Snmp_Correlator_Rules_Rule) GetSegmentPath() string {
    return "rule" + "[name='" + fmt.Sprintf("%v", rule.Name) + "']"
}

func (rule *Snmp_Correlator_Rules_Rule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "non-stateful" {
        return &rule.NonStateful
    }
    if childYangName == "applied-to" {
        return &rule.AppliedTo
    }
    return nil
}

func (rule *Snmp_Correlator_Rules_Rule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["non-stateful"] = &rule.NonStateful
    children["applied-to"] = &rule.AppliedTo
    return children
}

func (rule *Snmp_Correlator_Rules_Rule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = rule.Name
    return leafs
}

func (rule *Snmp_Correlator_Rules_Rule) GetBundleName() string { return "cisco_ios_xr" }

func (rule *Snmp_Correlator_Rules_Rule) GetYangName() string { return "rule" }

func (rule *Snmp_Correlator_Rules_Rule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rule *Snmp_Correlator_Rules_Rule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rule *Snmp_Correlator_Rules_Rule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rule *Snmp_Correlator_Rules_Rule) SetParent(parent types.Entity) { rule.parent = parent }

func (rule *Snmp_Correlator_Rules_Rule) GetParent() types.Entity { return rule.parent }

func (rule *Snmp_Correlator_Rules_Rule) GetParentYangName() string { return "rules" }

// Snmp_Correlator_Rules_Rule_NonStateful
// The Non-Stateful Rule Type
// This type is a presence type.
type Snmp_Correlator_Rules_Rule_NonStateful struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Timeout (time to wait for active correlation) in milliseconds. The type is
    // interface{} with range: 1..600000. Units are millisecond.
    Timeout interface{}

    // Table of configured rootcause (only one entry allowed).
    RootCauses Snmp_Correlator_Rules_Rule_NonStateful_RootCauses

    // Table of configured non-rootcause.
    NonRootCauses Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses
}

func (nonStateful *Snmp_Correlator_Rules_Rule_NonStateful) GetFilter() yfilter.YFilter { return nonStateful.YFilter }

func (nonStateful *Snmp_Correlator_Rules_Rule_NonStateful) SetFilter(yf yfilter.YFilter) { nonStateful.YFilter = yf }

func (nonStateful *Snmp_Correlator_Rules_Rule_NonStateful) GetGoName(yname string) string {
    if yname == "timeout" { return "Timeout" }
    if yname == "root-causes" { return "RootCauses" }
    if yname == "non-root-causes" { return "NonRootCauses" }
    return ""
}

func (nonStateful *Snmp_Correlator_Rules_Rule_NonStateful) GetSegmentPath() string {
    return "non-stateful"
}

func (nonStateful *Snmp_Correlator_Rules_Rule_NonStateful) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "root-causes" {
        return &nonStateful.RootCauses
    }
    if childYangName == "non-root-causes" {
        return &nonStateful.NonRootCauses
    }
    return nil
}

func (nonStateful *Snmp_Correlator_Rules_Rule_NonStateful) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["root-causes"] = &nonStateful.RootCauses
    children["non-root-causes"] = &nonStateful.NonRootCauses
    return children
}

func (nonStateful *Snmp_Correlator_Rules_Rule_NonStateful) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timeout"] = nonStateful.Timeout
    return leafs
}

func (nonStateful *Snmp_Correlator_Rules_Rule_NonStateful) GetBundleName() string { return "cisco_ios_xr" }

func (nonStateful *Snmp_Correlator_Rules_Rule_NonStateful) GetYangName() string { return "non-stateful" }

func (nonStateful *Snmp_Correlator_Rules_Rule_NonStateful) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nonStateful *Snmp_Correlator_Rules_Rule_NonStateful) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nonStateful *Snmp_Correlator_Rules_Rule_NonStateful) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nonStateful *Snmp_Correlator_Rules_Rule_NonStateful) SetParent(parent types.Entity) { nonStateful.parent = parent }

func (nonStateful *Snmp_Correlator_Rules_Rule_NonStateful) GetParent() types.Entity { return nonStateful.parent }

func (nonStateful *Snmp_Correlator_Rules_Rule_NonStateful) GetParentYangName() string { return "rule" }

// Snmp_Correlator_Rules_Rule_NonStateful_RootCauses
// Table of configured rootcause (only one
// entry allowed)
type Snmp_Correlator_Rules_Rule_NonStateful_RootCauses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The rootcause - maximum of one can be configured per rule. The type is
    // slice of Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause.
    RootCause []Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause
}

func (rootCauses *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses) GetFilter() yfilter.YFilter { return rootCauses.YFilter }

func (rootCauses *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses) SetFilter(yf yfilter.YFilter) { rootCauses.YFilter = yf }

func (rootCauses *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses) GetGoName(yname string) string {
    if yname == "root-cause" { return "RootCause" }
    return ""
}

func (rootCauses *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses) GetSegmentPath() string {
    return "root-causes"
}

func (rootCauses *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "root-cause" {
        for _, c := range rootCauses.RootCause {
            if rootCauses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause{}
        rootCauses.RootCause = append(rootCauses.RootCause, child)
        return &rootCauses.RootCause[len(rootCauses.RootCause)-1]
    }
    return nil
}

func (rootCauses *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rootCauses.RootCause {
        children[rootCauses.RootCause[i].GetSegmentPath()] = &rootCauses.RootCause[i]
    }
    return children
}

func (rootCauses *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rootCauses *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses) GetBundleName() string { return "cisco_ios_xr" }

func (rootCauses *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses) GetYangName() string { return "root-causes" }

func (rootCauses *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rootCauses *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rootCauses *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rootCauses *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses) SetParent(parent types.Entity) { rootCauses.parent = parent }

func (rootCauses *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses) GetParent() types.Entity { return rootCauses.parent }

func (rootCauses *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses) GetParentYangName() string { return "non-stateful" }

// Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause
// The rootcause - maximum of one can be
// configured per rule
type Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. OID of rootcause trap (dotted decimal). The type
    // is string.
    Oid interface{}

    // Create rootcause. The type is interface{}.
    Created interface{}

    // Varbinds to match.
    VarBinds Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds
}

func (rootCause *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause) GetFilter() yfilter.YFilter { return rootCause.YFilter }

func (rootCause *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause) SetFilter(yf yfilter.YFilter) { rootCause.YFilter = yf }

func (rootCause *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause) GetGoName(yname string) string {
    if yname == "oid" { return "Oid" }
    if yname == "created" { return "Created" }
    if yname == "var-binds" { return "VarBinds" }
    return ""
}

func (rootCause *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause) GetSegmentPath() string {
    return "root-cause" + "[oid='" + fmt.Sprintf("%v", rootCause.Oid) + "']"
}

func (rootCause *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "var-binds" {
        return &rootCause.VarBinds
    }
    return nil
}

func (rootCause *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["var-binds"] = &rootCause.VarBinds
    return children
}

func (rootCause *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oid"] = rootCause.Oid
    leafs["created"] = rootCause.Created
    return leafs
}

func (rootCause *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause) GetBundleName() string { return "cisco_ios_xr" }

func (rootCause *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause) GetYangName() string { return "root-cause" }

func (rootCause *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rootCause *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rootCause *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rootCause *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause) SetParent(parent types.Entity) { rootCause.parent = parent }

func (rootCause *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause) GetParent() types.Entity { return rootCause.parent }

func (rootCause *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause) GetParentYangName() string { return "root-causes" }

// Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds
// Varbinds to match
type Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Varbind match conditions. The type is slice of
    // Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind.
    VarBind []Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind
}

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds) GetFilter() yfilter.YFilter { return varBinds.YFilter }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds) SetFilter(yf yfilter.YFilter) { varBinds.YFilter = yf }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds) GetGoName(yname string) string {
    if yname == "var-bind" { return "VarBind" }
    return ""
}

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds) GetSegmentPath() string {
    return "var-binds"
}

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "var-bind" {
        for _, c := range varBinds.VarBind {
            if varBinds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind{}
        varBinds.VarBind = append(varBinds.VarBind, child)
        return &varBinds.VarBind[len(varBinds.VarBind)-1]
    }
    return nil
}

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range varBinds.VarBind {
        children[varBinds.VarBind[i].GetSegmentPath()] = &varBinds.VarBind[i]
    }
    return children
}

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds) GetBundleName() string { return "cisco_ios_xr" }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds) GetYangName() string { return "var-binds" }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds) SetParent(parent types.Entity) { varBinds.parent = parent }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds) GetParent() types.Entity { return varBinds.parent }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds) GetParentYangName() string { return "root-cause" }

// Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind
// Varbind match conditions
type Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. OID of varbind (dotted decimal). The type is
    // string.
    Oid interface{}

    // VarBind match conditions.
    Match Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match
}

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind) GetFilter() yfilter.YFilter { return varBind.YFilter }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind) SetFilter(yf yfilter.YFilter) { varBind.YFilter = yf }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind) GetGoName(yname string) string {
    if yname == "oid" { return "Oid" }
    if yname == "match" { return "Match" }
    return ""
}

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind) GetSegmentPath() string {
    return "var-bind" + "[oid='" + fmt.Sprintf("%v", varBind.Oid) + "']"
}

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "match" {
        return &varBind.Match
    }
    return nil
}

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["match"] = &varBind.Match
    return children
}

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oid"] = varBind.Oid
    return leafs
}

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind) GetBundleName() string { return "cisco_ios_xr" }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind) GetYangName() string { return "var-bind" }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind) SetParent(parent types.Entity) { varBind.parent = parent }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind) GetParent() types.Entity { return varBind.parent }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind) GetParentYangName() string { return "var-binds" }

// Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match
// VarBind match conditions
type Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Regular Expression to match value. The type is string.
    Value interface{}

    // Regular Expression to match index. The type is string.
    Index interface{}
}

func (match *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match) GetFilter() yfilter.YFilter { return match.YFilter }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match) SetFilter(yf yfilter.YFilter) { match.YFilter = yf }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "index" { return "Index" }
    return ""
}

func (match *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match) GetSegmentPath() string {
    return "match"
}

func (match *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (match *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (match *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = match.Value
    leafs["index"] = match.Index
    return leafs
}

func (match *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match) GetBundleName() string { return "cisco_ios_xr" }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match) GetYangName() string { return "match" }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match) SetParent(parent types.Entity) { match.parent = parent }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match) GetParent() types.Entity { return match.parent }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_RootCauses_RootCause_VarBinds_VarBind_Match) GetParentYangName() string { return "var-bind" }

// Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses
// Table of configured non-rootcause
type Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A non-rootcause. The type is slice of
    // Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause.
    NonRootCause []Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause
}

func (nonRootCauses *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetFilter() yfilter.YFilter { return nonRootCauses.YFilter }

func (nonRootCauses *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses) SetFilter(yf yfilter.YFilter) { nonRootCauses.YFilter = yf }

func (nonRootCauses *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetGoName(yname string) string {
    if yname == "non-root-cause" { return "NonRootCause" }
    return ""
}

func (nonRootCauses *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetSegmentPath() string {
    return "non-root-causes"
}

func (nonRootCauses *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "non-root-cause" {
        for _, c := range nonRootCauses.NonRootCause {
            if nonRootCauses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause{}
        nonRootCauses.NonRootCause = append(nonRootCauses.NonRootCause, child)
        return &nonRootCauses.NonRootCause[len(nonRootCauses.NonRootCause)-1]
    }
    return nil
}

func (nonRootCauses *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nonRootCauses.NonRootCause {
        children[nonRootCauses.NonRootCause[i].GetSegmentPath()] = &nonRootCauses.NonRootCause[i]
    }
    return children
}

func (nonRootCauses *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nonRootCauses *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetBundleName() string { return "cisco_ios_xr" }

func (nonRootCauses *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetYangName() string { return "non-root-causes" }

func (nonRootCauses *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nonRootCauses *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nonRootCauses *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nonRootCauses *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses) SetParent(parent types.Entity) { nonRootCauses.parent = parent }

func (nonRootCauses *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetParent() types.Entity { return nonRootCauses.parent }

func (nonRootCauses *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses) GetParentYangName() string { return "non-stateful" }

// Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause
// A non-rootcause
type Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. OID of nonrootcause trap (dotted decimal). The
    // type is string.
    Oid interface{}

    // Create nonrootcause. The type is interface{}.
    Created interface{}

    // Varbinds to match.
    VarBinds Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds
}

func (nonRootCause *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetFilter() yfilter.YFilter { return nonRootCause.YFilter }

func (nonRootCause *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) SetFilter(yf yfilter.YFilter) { nonRootCause.YFilter = yf }

func (nonRootCause *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetGoName(yname string) string {
    if yname == "oid" { return "Oid" }
    if yname == "created" { return "Created" }
    if yname == "var-binds" { return "VarBinds" }
    return ""
}

func (nonRootCause *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetSegmentPath() string {
    return "non-root-cause" + "[oid='" + fmt.Sprintf("%v", nonRootCause.Oid) + "']"
}

func (nonRootCause *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "var-binds" {
        return &nonRootCause.VarBinds
    }
    return nil
}

func (nonRootCause *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["var-binds"] = &nonRootCause.VarBinds
    return children
}

func (nonRootCause *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oid"] = nonRootCause.Oid
    leafs["created"] = nonRootCause.Created
    return leafs
}

func (nonRootCause *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetBundleName() string { return "cisco_ios_xr" }

func (nonRootCause *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetYangName() string { return "non-root-cause" }

func (nonRootCause *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nonRootCause *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nonRootCause *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nonRootCause *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) SetParent(parent types.Entity) { nonRootCause.parent = parent }

func (nonRootCause *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetParent() types.Entity { return nonRootCause.parent }

func (nonRootCause *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause) GetParentYangName() string { return "non-root-causes" }

// Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds
// Varbinds to match
type Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Varbind match conditions. The type is slice of
    // Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind.
    VarBind []Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind
}

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds) GetFilter() yfilter.YFilter { return varBinds.YFilter }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds) SetFilter(yf yfilter.YFilter) { varBinds.YFilter = yf }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds) GetGoName(yname string) string {
    if yname == "var-bind" { return "VarBind" }
    return ""
}

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds) GetSegmentPath() string {
    return "var-binds"
}

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "var-bind" {
        for _, c := range varBinds.VarBind {
            if varBinds.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind{}
        varBinds.VarBind = append(varBinds.VarBind, child)
        return &varBinds.VarBind[len(varBinds.VarBind)-1]
    }
    return nil
}

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range varBinds.VarBind {
        children[varBinds.VarBind[i].GetSegmentPath()] = &varBinds.VarBind[i]
    }
    return children
}

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds) GetBundleName() string { return "cisco_ios_xr" }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds) GetYangName() string { return "var-binds" }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds) SetParent(parent types.Entity) { varBinds.parent = parent }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds) GetParent() types.Entity { return varBinds.parent }

func (varBinds *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds) GetParentYangName() string { return "non-root-cause" }

// Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind
// Varbind match conditions
type Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. OID of varbind (dotted decimal). The type is
    // string.
    Oid interface{}

    // VarBind match conditions.
    Match Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match
}

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind) GetFilter() yfilter.YFilter { return varBind.YFilter }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind) SetFilter(yf yfilter.YFilter) { varBind.YFilter = yf }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind) GetGoName(yname string) string {
    if yname == "oid" { return "Oid" }
    if yname == "match" { return "Match" }
    return ""
}

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind) GetSegmentPath() string {
    return "var-bind" + "[oid='" + fmt.Sprintf("%v", varBind.Oid) + "']"
}

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "match" {
        return &varBind.Match
    }
    return nil
}

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["match"] = &varBind.Match
    return children
}

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oid"] = varBind.Oid
    return leafs
}

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind) GetBundleName() string { return "cisco_ios_xr" }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind) GetYangName() string { return "var-bind" }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind) SetParent(parent types.Entity) { varBind.parent = parent }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind) GetParent() types.Entity { return varBind.parent }

func (varBind *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind) GetParentYangName() string { return "var-binds" }

// Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match
// VarBind match conditions
type Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Regular Expression to match value. The type is string.
    Value interface{}

    // Regular Expression to match index. The type is string.
    Index interface{}
}

func (match *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match) GetFilter() yfilter.YFilter { return match.YFilter }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match) SetFilter(yf yfilter.YFilter) { match.YFilter = yf }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match) GetGoName(yname string) string {
    if yname == "value" { return "Value" }
    if yname == "index" { return "Index" }
    return ""
}

func (match *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match) GetSegmentPath() string {
    return "match"
}

func (match *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (match *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (match *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["value"] = match.Value
    leafs["index"] = match.Index
    return leafs
}

func (match *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match) GetBundleName() string { return "cisco_ios_xr" }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match) GetYangName() string { return "match" }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match) SetParent(parent types.Entity) { match.parent = parent }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match) GetParent() types.Entity { return match.parent }

func (match *Snmp_Correlator_Rules_Rule_NonStateful_NonRootCauses_NonRootCause_VarBinds_VarBind_Match) GetParentYangName() string { return "var-bind" }

// Snmp_Correlator_Rules_Rule_AppliedTo
// Applied to the Rule or Ruleset
type Snmp_Correlator_Rules_Rule_AppliedTo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Apply to all of the device. The type is interface{}.
    All interface{}

    // Table of configured hosts to apply rules to.
    Hosts Snmp_Correlator_Rules_Rule_AppliedTo_Hosts
}

func (appliedTo *Snmp_Correlator_Rules_Rule_AppliedTo) GetFilter() yfilter.YFilter { return appliedTo.YFilter }

func (appliedTo *Snmp_Correlator_Rules_Rule_AppliedTo) SetFilter(yf yfilter.YFilter) { appliedTo.YFilter = yf }

func (appliedTo *Snmp_Correlator_Rules_Rule_AppliedTo) GetGoName(yname string) string {
    if yname == "all" { return "All" }
    if yname == "hosts" { return "Hosts" }
    return ""
}

func (appliedTo *Snmp_Correlator_Rules_Rule_AppliedTo) GetSegmentPath() string {
    return "applied-to"
}

func (appliedTo *Snmp_Correlator_Rules_Rule_AppliedTo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hosts" {
        return &appliedTo.Hosts
    }
    return nil
}

func (appliedTo *Snmp_Correlator_Rules_Rule_AppliedTo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hosts"] = &appliedTo.Hosts
    return children
}

func (appliedTo *Snmp_Correlator_Rules_Rule_AppliedTo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["all"] = appliedTo.All
    return leafs
}

func (appliedTo *Snmp_Correlator_Rules_Rule_AppliedTo) GetBundleName() string { return "cisco_ios_xr" }

func (appliedTo *Snmp_Correlator_Rules_Rule_AppliedTo) GetYangName() string { return "applied-to" }

func (appliedTo *Snmp_Correlator_Rules_Rule_AppliedTo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appliedTo *Snmp_Correlator_Rules_Rule_AppliedTo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appliedTo *Snmp_Correlator_Rules_Rule_AppliedTo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appliedTo *Snmp_Correlator_Rules_Rule_AppliedTo) SetParent(parent types.Entity) { appliedTo.parent = parent }

func (appliedTo *Snmp_Correlator_Rules_Rule_AppliedTo) GetParent() types.Entity { return appliedTo.parent }

func (appliedTo *Snmp_Correlator_Rules_Rule_AppliedTo) GetParentYangName() string { return "rule" }

// Snmp_Correlator_Rules_Rule_AppliedTo_Hosts
// Table of configured hosts to apply rules to
type Snmp_Correlator_Rules_Rule_AppliedTo_Hosts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A destination host. The type is slice of
    // Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host.
    Host []Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host
}

func (hosts *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts) GetFilter() yfilter.YFilter { return hosts.YFilter }

func (hosts *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts) SetFilter(yf yfilter.YFilter) { hosts.YFilter = yf }

func (hosts *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts) GetGoName(yname string) string {
    if yname == "host" { return "Host" }
    return ""
}

func (hosts *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts) GetSegmentPath() string {
    return "hosts"
}

func (hosts *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host" {
        for _, c := range hosts.Host {
            if hosts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host{}
        hosts.Host = append(hosts.Host, child)
        return &hosts.Host[len(hosts.Host)-1]
    }
    return nil
}

func (hosts *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hosts.Host {
        children[hosts.Host[i].GetSegmentPath()] = &hosts.Host[i]
    }
    return children
}

func (hosts *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hosts *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts) GetBundleName() string { return "cisco_ios_xr" }

func (hosts *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts) GetYangName() string { return "hosts" }

func (hosts *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hosts *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hosts *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hosts *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts) SetParent(parent types.Entity) { hosts.parent = parent }

func (hosts *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts) GetParent() types.Entity { return hosts.parent }

func (hosts *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts) GetParentYangName() string { return "applied-to" }

// Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host
// A destination host
type Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // This attribute is a key. Port (specify 162 for default). The type is
    // interface{} with range: 1..65535.
    Port interface{}
}

func (host *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host) GetFilter() yfilter.YFilter { return host.YFilter }

func (host *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host) SetFilter(yf yfilter.YFilter) { host.YFilter = yf }

func (host *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "port" { return "Port" }
    return ""
}

func (host *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host) GetSegmentPath() string {
    return "host" + "[ip-address='" + fmt.Sprintf("%v", host.IpAddress) + "']" + "[port='" + fmt.Sprintf("%v", host.Port) + "']"
}

func (host *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (host *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (host *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = host.IpAddress
    leafs["port"] = host.Port
    return leafs
}

func (host *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host) GetBundleName() string { return "cisco_ios_xr" }

func (host *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host) GetYangName() string { return "host" }

func (host *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (host *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (host *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (host *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host) SetParent(parent types.Entity) { host.parent = parent }

func (host *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host) GetParent() types.Entity { return host.parent }

func (host *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host) GetParentYangName() string { return "hosts" }

// Snmp_Correlator_RuleSets
// Table of configured rulesets
type Snmp_Correlator_RuleSets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Ruleset name. The type is slice of Snmp_Correlator_RuleSets_RuleSet.
    RuleSet []Snmp_Correlator_RuleSets_RuleSet
}

func (ruleSets *Snmp_Correlator_RuleSets) GetFilter() yfilter.YFilter { return ruleSets.YFilter }

func (ruleSets *Snmp_Correlator_RuleSets) SetFilter(yf yfilter.YFilter) { ruleSets.YFilter = yf }

func (ruleSets *Snmp_Correlator_RuleSets) GetGoName(yname string) string {
    if yname == "rule-set" { return "RuleSet" }
    return ""
}

func (ruleSets *Snmp_Correlator_RuleSets) GetSegmentPath() string {
    return "rule-sets"
}

func (ruleSets *Snmp_Correlator_RuleSets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rule-set" {
        for _, c := range ruleSets.RuleSet {
            if ruleSets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_RuleSets_RuleSet{}
        ruleSets.RuleSet = append(ruleSets.RuleSet, child)
        return &ruleSets.RuleSet[len(ruleSets.RuleSet)-1]
    }
    return nil
}

func (ruleSets *Snmp_Correlator_RuleSets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ruleSets.RuleSet {
        children[ruleSets.RuleSet[i].GetSegmentPath()] = &ruleSets.RuleSet[i]
    }
    return children
}

func (ruleSets *Snmp_Correlator_RuleSets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ruleSets *Snmp_Correlator_RuleSets) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSets *Snmp_Correlator_RuleSets) GetYangName() string { return "rule-sets" }

func (ruleSets *Snmp_Correlator_RuleSets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSets *Snmp_Correlator_RuleSets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSets *Snmp_Correlator_RuleSets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSets *Snmp_Correlator_RuleSets) SetParent(parent types.Entity) { ruleSets.parent = parent }

func (ruleSets *Snmp_Correlator_RuleSets) GetParent() types.Entity { return ruleSets.parent }

func (ruleSets *Snmp_Correlator_RuleSets) GetParentYangName() string { return "correlator" }

// Snmp_Correlator_RuleSets_RuleSet
// Ruleset name
type Snmp_Correlator_RuleSets_RuleSet struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Ruleset name. The type is string with length:
    // 1..32.
    Name interface{}

    // Table of configured rulenames.
    Rulenames Snmp_Correlator_RuleSets_RuleSet_Rulenames

    // Applied to the Rule or Ruleset.
    AppliedTo Snmp_Correlator_RuleSets_RuleSet_AppliedTo
}

func (ruleSet *Snmp_Correlator_RuleSets_RuleSet) GetFilter() yfilter.YFilter { return ruleSet.YFilter }

func (ruleSet *Snmp_Correlator_RuleSets_RuleSet) SetFilter(yf yfilter.YFilter) { ruleSet.YFilter = yf }

func (ruleSet *Snmp_Correlator_RuleSets_RuleSet) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "rulenames" { return "Rulenames" }
    if yname == "applied-to" { return "AppliedTo" }
    return ""
}

func (ruleSet *Snmp_Correlator_RuleSets_RuleSet) GetSegmentPath() string {
    return "rule-set" + "[name='" + fmt.Sprintf("%v", ruleSet.Name) + "']"
}

func (ruleSet *Snmp_Correlator_RuleSets_RuleSet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rulenames" {
        return &ruleSet.Rulenames
    }
    if childYangName == "applied-to" {
        return &ruleSet.AppliedTo
    }
    return nil
}

func (ruleSet *Snmp_Correlator_RuleSets_RuleSet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rulenames"] = &ruleSet.Rulenames
    children["applied-to"] = &ruleSet.AppliedTo
    return children
}

func (ruleSet *Snmp_Correlator_RuleSets_RuleSet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = ruleSet.Name
    return leafs
}

func (ruleSet *Snmp_Correlator_RuleSets_RuleSet) GetBundleName() string { return "cisco_ios_xr" }

func (ruleSet *Snmp_Correlator_RuleSets_RuleSet) GetYangName() string { return "rule-set" }

func (ruleSet *Snmp_Correlator_RuleSets_RuleSet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ruleSet *Snmp_Correlator_RuleSets_RuleSet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ruleSet *Snmp_Correlator_RuleSets_RuleSet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ruleSet *Snmp_Correlator_RuleSets_RuleSet) SetParent(parent types.Entity) { ruleSet.parent = parent }

func (ruleSet *Snmp_Correlator_RuleSets_RuleSet) GetParent() types.Entity { return ruleSet.parent }

func (ruleSet *Snmp_Correlator_RuleSets_RuleSet) GetParentYangName() string { return "rule-sets" }

// Snmp_Correlator_RuleSets_RuleSet_Rulenames
// Table of configured rulenames
type Snmp_Correlator_RuleSets_RuleSet_Rulenames struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A rulename. The type is slice of
    // Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename.
    Rulename []Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename
}

func (rulenames *Snmp_Correlator_RuleSets_RuleSet_Rulenames) GetFilter() yfilter.YFilter { return rulenames.YFilter }

func (rulenames *Snmp_Correlator_RuleSets_RuleSet_Rulenames) SetFilter(yf yfilter.YFilter) { rulenames.YFilter = yf }

func (rulenames *Snmp_Correlator_RuleSets_RuleSet_Rulenames) GetGoName(yname string) string {
    if yname == "rulename" { return "Rulename" }
    return ""
}

func (rulenames *Snmp_Correlator_RuleSets_RuleSet_Rulenames) GetSegmentPath() string {
    return "rulenames"
}

func (rulenames *Snmp_Correlator_RuleSets_RuleSet_Rulenames) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rulename" {
        for _, c := range rulenames.Rulename {
            if rulenames.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename{}
        rulenames.Rulename = append(rulenames.Rulename, child)
        return &rulenames.Rulename[len(rulenames.Rulename)-1]
    }
    return nil
}

func (rulenames *Snmp_Correlator_RuleSets_RuleSet_Rulenames) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range rulenames.Rulename {
        children[rulenames.Rulename[i].GetSegmentPath()] = &rulenames.Rulename[i]
    }
    return children
}

func (rulenames *Snmp_Correlator_RuleSets_RuleSet_Rulenames) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rulenames *Snmp_Correlator_RuleSets_RuleSet_Rulenames) GetBundleName() string { return "cisco_ios_xr" }

func (rulenames *Snmp_Correlator_RuleSets_RuleSet_Rulenames) GetYangName() string { return "rulenames" }

func (rulenames *Snmp_Correlator_RuleSets_RuleSet_Rulenames) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rulenames *Snmp_Correlator_RuleSets_RuleSet_Rulenames) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rulenames *Snmp_Correlator_RuleSets_RuleSet_Rulenames) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rulenames *Snmp_Correlator_RuleSets_RuleSet_Rulenames) SetParent(parent types.Entity) { rulenames.parent = parent }

func (rulenames *Snmp_Correlator_RuleSets_RuleSet_Rulenames) GetParent() types.Entity { return rulenames.parent }

func (rulenames *Snmp_Correlator_RuleSets_RuleSet_Rulenames) GetParentYangName() string { return "rule-set" }

// Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename
// A rulename
type Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Rule name. The type is string with length: 1..32.
    Rulename interface{}
}

func (rulename *Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetFilter() yfilter.YFilter { return rulename.YFilter }

func (rulename *Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename) SetFilter(yf yfilter.YFilter) { rulename.YFilter = yf }

func (rulename *Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetGoName(yname string) string {
    if yname == "rulename" { return "Rulename" }
    return ""
}

func (rulename *Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetSegmentPath() string {
    return "rulename" + "[rulename='" + fmt.Sprintf("%v", rulename.Rulename) + "']"
}

func (rulename *Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rulename *Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rulename *Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["rulename"] = rulename.Rulename
    return leafs
}

func (rulename *Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetBundleName() string { return "cisco_ios_xr" }

func (rulename *Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetYangName() string { return "rulename" }

func (rulename *Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rulename *Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rulename *Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rulename *Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename) SetParent(parent types.Entity) { rulename.parent = parent }

func (rulename *Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetParent() types.Entity { return rulename.parent }

func (rulename *Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetParentYangName() string { return "rulenames" }

// Snmp_Correlator_RuleSets_RuleSet_AppliedTo
// Applied to the Rule or Ruleset
type Snmp_Correlator_RuleSets_RuleSet_AppliedTo struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Apply to all of the device. The type is interface{}.
    All interface{}

    // Table of configured hosts to apply rules to.
    Hosts Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts
}

func (appliedTo *Snmp_Correlator_RuleSets_RuleSet_AppliedTo) GetFilter() yfilter.YFilter { return appliedTo.YFilter }

func (appliedTo *Snmp_Correlator_RuleSets_RuleSet_AppliedTo) SetFilter(yf yfilter.YFilter) { appliedTo.YFilter = yf }

func (appliedTo *Snmp_Correlator_RuleSets_RuleSet_AppliedTo) GetGoName(yname string) string {
    if yname == "all" { return "All" }
    if yname == "hosts" { return "Hosts" }
    return ""
}

func (appliedTo *Snmp_Correlator_RuleSets_RuleSet_AppliedTo) GetSegmentPath() string {
    return "applied-to"
}

func (appliedTo *Snmp_Correlator_RuleSets_RuleSet_AppliedTo) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "hosts" {
        return &appliedTo.Hosts
    }
    return nil
}

func (appliedTo *Snmp_Correlator_RuleSets_RuleSet_AppliedTo) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["hosts"] = &appliedTo.Hosts
    return children
}

func (appliedTo *Snmp_Correlator_RuleSets_RuleSet_AppliedTo) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["all"] = appliedTo.All
    return leafs
}

func (appliedTo *Snmp_Correlator_RuleSets_RuleSet_AppliedTo) GetBundleName() string { return "cisco_ios_xr" }

func (appliedTo *Snmp_Correlator_RuleSets_RuleSet_AppliedTo) GetYangName() string { return "applied-to" }

func (appliedTo *Snmp_Correlator_RuleSets_RuleSet_AppliedTo) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (appliedTo *Snmp_Correlator_RuleSets_RuleSet_AppliedTo) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (appliedTo *Snmp_Correlator_RuleSets_RuleSet_AppliedTo) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (appliedTo *Snmp_Correlator_RuleSets_RuleSet_AppliedTo) SetParent(parent types.Entity) { appliedTo.parent = parent }

func (appliedTo *Snmp_Correlator_RuleSets_RuleSet_AppliedTo) GetParent() types.Entity { return appliedTo.parent }

func (appliedTo *Snmp_Correlator_RuleSets_RuleSet_AppliedTo) GetParentYangName() string { return "rule-set" }

// Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts
// Table of configured hosts to apply rules to
type Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A destination host. The type is slice of
    // Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host.
    Host []Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host
}

func (hosts *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts) GetFilter() yfilter.YFilter { return hosts.YFilter }

func (hosts *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts) SetFilter(yf yfilter.YFilter) { hosts.YFilter = yf }

func (hosts *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts) GetGoName(yname string) string {
    if yname == "host" { return "Host" }
    return ""
}

func (hosts *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts) GetSegmentPath() string {
    return "hosts"
}

func (hosts *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "host" {
        for _, c := range hosts.Host {
            if hosts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host{}
        hosts.Host = append(hosts.Host, child)
        return &hosts.Host[len(hosts.Host)-1]
    }
    return nil
}

func (hosts *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range hosts.Host {
        children[hosts.Host[i].GetSegmentPath()] = &hosts.Host[i]
    }
    return children
}

func (hosts *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hosts *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts) GetBundleName() string { return "cisco_ios_xr" }

func (hosts *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts) GetYangName() string { return "hosts" }

func (hosts *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hosts *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hosts *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hosts *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts) SetParent(parent types.Entity) { hosts.parent = parent }

func (hosts *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts) GetParent() types.Entity { return hosts.parent }

func (hosts *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts) GetParentYangName() string { return "applied-to" }

// Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host
// A destination host
type Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // This attribute is a key. Port (specify 162 for default). The type is
    // interface{} with range: 1..65535.
    Port interface{}
}

func (host *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host) GetFilter() yfilter.YFilter { return host.YFilter }

func (host *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host) SetFilter(yf yfilter.YFilter) { host.YFilter = yf }

func (host *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "port" { return "Port" }
    return ""
}

func (host *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host) GetSegmentPath() string {
    return "host" + "[ip-address='" + fmt.Sprintf("%v", host.IpAddress) + "']" + "[port='" + fmt.Sprintf("%v", host.Port) + "']"
}

func (host *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (host *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (host *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = host.IpAddress
    leafs["port"] = host.Port
    return leafs
}

func (host *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host) GetBundleName() string { return "cisco_ios_xr" }

func (host *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host) GetYangName() string { return "host" }

func (host *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (host *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (host *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (host *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host) SetParent(parent types.Entity) { host.parent = parent }

func (host *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host) GetParent() types.Entity { return host.parent }

func (host *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host) GetParentYangName() string { return "hosts" }

// Snmp_BulkStats
// SNMP bulk stats configuration commands
type Snmp_BulkStats struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // per process memory limit in kilo bytes. The type is interface{} with range:
    // 100..200000. Units are kilobyte.
    Memory interface{}

    // Configure schema definition.
    Schemas Snmp_BulkStats_Schemas

    // Configure an Object List .
    Objects Snmp_BulkStats_Objects

    // Periodicity for the transfer of bulk data in minutes.
    Transfers Snmp_BulkStats_Transfers
}

func (bulkStats *Snmp_BulkStats) GetFilter() yfilter.YFilter { return bulkStats.YFilter }

func (bulkStats *Snmp_BulkStats) SetFilter(yf yfilter.YFilter) { bulkStats.YFilter = yf }

func (bulkStats *Snmp_BulkStats) GetGoName(yname string) string {
    if yname == "memory" { return "Memory" }
    if yname == "schemas" { return "Schemas" }
    if yname == "objects" { return "Objects" }
    if yname == "transfers" { return "Transfers" }
    return ""
}

func (bulkStats *Snmp_BulkStats) GetSegmentPath() string {
    return "bulk-stats"
}

func (bulkStats *Snmp_BulkStats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "schemas" {
        return &bulkStats.Schemas
    }
    if childYangName == "objects" {
        return &bulkStats.Objects
    }
    if childYangName == "transfers" {
        return &bulkStats.Transfers
    }
    return nil
}

func (bulkStats *Snmp_BulkStats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["schemas"] = &bulkStats.Schemas
    children["objects"] = &bulkStats.Objects
    children["transfers"] = &bulkStats.Transfers
    return children
}

func (bulkStats *Snmp_BulkStats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["memory"] = bulkStats.Memory
    return leafs
}

func (bulkStats *Snmp_BulkStats) GetBundleName() string { return "cisco_ios_xr" }

func (bulkStats *Snmp_BulkStats) GetYangName() string { return "bulk-stats" }

func (bulkStats *Snmp_BulkStats) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bulkStats *Snmp_BulkStats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bulkStats *Snmp_BulkStats) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bulkStats *Snmp_BulkStats) SetParent(parent types.Entity) { bulkStats.parent = parent }

func (bulkStats *Snmp_BulkStats) GetParent() types.Entity { return bulkStats.parent }

func (bulkStats *Snmp_BulkStats) GetParentYangName() string { return "snmp" }

// Snmp_BulkStats_Schemas
// Configure schema definition
type Snmp_BulkStats_Schemas struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the Schema. The type is slice of Snmp_BulkStats_Schemas_Schema.
    Schema []Snmp_BulkStats_Schemas_Schema
}

func (schemas *Snmp_BulkStats_Schemas) GetFilter() yfilter.YFilter { return schemas.YFilter }

func (schemas *Snmp_BulkStats_Schemas) SetFilter(yf yfilter.YFilter) { schemas.YFilter = yf }

func (schemas *Snmp_BulkStats_Schemas) GetGoName(yname string) string {
    if yname == "schema" { return "Schema" }
    return ""
}

func (schemas *Snmp_BulkStats_Schemas) GetSegmentPath() string {
    return "schemas"
}

func (schemas *Snmp_BulkStats_Schemas) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "schema" {
        for _, c := range schemas.Schema {
            if schemas.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_BulkStats_Schemas_Schema{}
        schemas.Schema = append(schemas.Schema, child)
        return &schemas.Schema[len(schemas.Schema)-1]
    }
    return nil
}

func (schemas *Snmp_BulkStats_Schemas) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range schemas.Schema {
        children[schemas.Schema[i].GetSegmentPath()] = &schemas.Schema[i]
    }
    return children
}

func (schemas *Snmp_BulkStats_Schemas) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (schemas *Snmp_BulkStats_Schemas) GetBundleName() string { return "cisco_ios_xr" }

func (schemas *Snmp_BulkStats_Schemas) GetYangName() string { return "schemas" }

func (schemas *Snmp_BulkStats_Schemas) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (schemas *Snmp_BulkStats_Schemas) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (schemas *Snmp_BulkStats_Schemas) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (schemas *Snmp_BulkStats_Schemas) SetParent(parent types.Entity) { schemas.parent = parent }

func (schemas *Snmp_BulkStats_Schemas) GetParent() types.Entity { return schemas.parent }

func (schemas *Snmp_BulkStats_Schemas) GetParentYangName() string { return "bulk-stats" }

// Snmp_BulkStats_Schemas_Schema
// The name of the Schema
type Snmp_BulkStats_Schemas_Schema struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the schema. The type is string with
    // length: 1..32.
    SchemaName interface{}

    // Configure schema name. The type is interface{}.
    Type interface{}

    // Name of an object List. The type is string with length: 1..32.
    SchemaObjectList interface{}

    // Periodicity for polling of objects in this schema in minutes. The type is
    // interface{} with range: 1..20000. Units are minute.
    PollInterval interface{}

    // Object instance information.
    Instance Snmp_BulkStats_Schemas_Schema_Instance
}

func (schema *Snmp_BulkStats_Schemas_Schema) GetFilter() yfilter.YFilter { return schema.YFilter }

func (schema *Snmp_BulkStats_Schemas_Schema) SetFilter(yf yfilter.YFilter) { schema.YFilter = yf }

func (schema *Snmp_BulkStats_Schemas_Schema) GetGoName(yname string) string {
    if yname == "schema-name" { return "SchemaName" }
    if yname == "type" { return "Type" }
    if yname == "schema-object-list" { return "SchemaObjectList" }
    if yname == "poll-interval" { return "PollInterval" }
    if yname == "instance" { return "Instance" }
    return ""
}

func (schema *Snmp_BulkStats_Schemas_Schema) GetSegmentPath() string {
    return "schema" + "[schema-name='" + fmt.Sprintf("%v", schema.SchemaName) + "']"
}

func (schema *Snmp_BulkStats_Schemas_Schema) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "instance" {
        return &schema.Instance
    }
    return nil
}

func (schema *Snmp_BulkStats_Schemas_Schema) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["instance"] = &schema.Instance
    return children
}

func (schema *Snmp_BulkStats_Schemas_Schema) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["schema-name"] = schema.SchemaName
    leafs["type"] = schema.Type
    leafs["schema-object-list"] = schema.SchemaObjectList
    leafs["poll-interval"] = schema.PollInterval
    return leafs
}

func (schema *Snmp_BulkStats_Schemas_Schema) GetBundleName() string { return "cisco_ios_xr" }

func (schema *Snmp_BulkStats_Schemas_Schema) GetYangName() string { return "schema" }

func (schema *Snmp_BulkStats_Schemas_Schema) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (schema *Snmp_BulkStats_Schemas_Schema) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (schema *Snmp_BulkStats_Schemas_Schema) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (schema *Snmp_BulkStats_Schemas_Schema) SetParent(parent types.Entity) { schema.parent = parent }

func (schema *Snmp_BulkStats_Schemas_Schema) GetParent() types.Entity { return schema.parent }

func (schema *Snmp_BulkStats_Schemas_Schema) GetParentYangName() string { return "schemas" }

// Snmp_BulkStats_Schemas_Schema_Instance
// Object instance information
// This type is a presence type.
type Snmp_BulkStats_Schemas_Schema_Instance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Type of the instance. The type is SnmpBulkstatSchema. This attribute is
    // mandatory.
    Type interface{}

    // Instance of the schema. The type is string with pattern: [a-zA-Z0-9./-]+.
    Instance interface{}

    // Start Instance OID for repetition. The type is string. This attribute is
    // mandatory.
    Start interface{}

    // End Instance OID for repetition. The type is string. This attribute is
    // mandatory.
    End interface{}

    // Max value of Instance repetition. The type is interface{} with range:
    // -2147483648..2147483647. This attribute is mandatory.
    Max interface{}

    // Include all the subinterface. The type is bool. This attribute is
    // mandatory.
    SubInterface interface{}
}

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) GetFilter() yfilter.YFilter { return instance.YFilter }

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) SetFilter(yf yfilter.YFilter) { instance.YFilter = yf }

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "instance" { return "Instance" }
    if yname == "start" { return "Start" }
    if yname == "end" { return "End" }
    if yname == "max" { return "Max" }
    if yname == "sub-interface" { return "SubInterface" }
    return ""
}

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) GetSegmentPath() string {
    return "instance"
}

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = instance.Type
    leafs["instance"] = instance.Instance
    leafs["start"] = instance.Start
    leafs["end"] = instance.End
    leafs["max"] = instance.Max
    leafs["sub-interface"] = instance.SubInterface
    return leafs
}

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) GetBundleName() string { return "cisco_ios_xr" }

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) GetYangName() string { return "instance" }

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) SetParent(parent types.Entity) { instance.parent = parent }

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) GetParent() types.Entity { return instance.parent }

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) GetParentYangName() string { return "schema" }

// Snmp_BulkStats_Objects
// Configure an Object List 
type Snmp_BulkStats_Objects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the object List. The type is slice of
    // Snmp_BulkStats_Objects_Object.
    Object []Snmp_BulkStats_Objects_Object
}

func (objects *Snmp_BulkStats_Objects) GetFilter() yfilter.YFilter { return objects.YFilter }

func (objects *Snmp_BulkStats_Objects) SetFilter(yf yfilter.YFilter) { objects.YFilter = yf }

func (objects *Snmp_BulkStats_Objects) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (objects *Snmp_BulkStats_Objects) GetSegmentPath() string {
    return "objects"
}

func (objects *Snmp_BulkStats_Objects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "object" {
        for _, c := range objects.Object {
            if objects.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_BulkStats_Objects_Object{}
        objects.Object = append(objects.Object, child)
        return &objects.Object[len(objects.Object)-1]
    }
    return nil
}

func (objects *Snmp_BulkStats_Objects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range objects.Object {
        children[objects.Object[i].GetSegmentPath()] = &objects.Object[i]
    }
    return children
}

func (objects *Snmp_BulkStats_Objects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (objects *Snmp_BulkStats_Objects) GetBundleName() string { return "cisco_ios_xr" }

func (objects *Snmp_BulkStats_Objects) GetYangName() string { return "objects" }

func (objects *Snmp_BulkStats_Objects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (objects *Snmp_BulkStats_Objects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (objects *Snmp_BulkStats_Objects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (objects *Snmp_BulkStats_Objects) SetParent(parent types.Entity) { objects.parent = parent }

func (objects *Snmp_BulkStats_Objects) GetParent() types.Entity { return objects.parent }

func (objects *Snmp_BulkStats_Objects) GetParentYangName() string { return "bulk-stats" }

// Snmp_BulkStats_Objects_Object
// Name of the object List
type Snmp_BulkStats_Objects_Object struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the object List. The type is string with
    // length: 1..32.
    ObjectListName interface{}

    // Configure object list name. The type is interface{}.
    Type interface{}

    // Configure an object List.
    Objects Snmp_BulkStats_Objects_Object_Objects
}

func (object *Snmp_BulkStats_Objects_Object) GetFilter() yfilter.YFilter { return object.YFilter }

func (object *Snmp_BulkStats_Objects_Object) SetFilter(yf yfilter.YFilter) { object.YFilter = yf }

func (object *Snmp_BulkStats_Objects_Object) GetGoName(yname string) string {
    if yname == "object-list-name" { return "ObjectListName" }
    if yname == "type" { return "Type" }
    if yname == "objects" { return "Objects" }
    return ""
}

func (object *Snmp_BulkStats_Objects_Object) GetSegmentPath() string {
    return "object" + "[object-list-name='" + fmt.Sprintf("%v", object.ObjectListName) + "']"
}

func (object *Snmp_BulkStats_Objects_Object) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "objects" {
        return &object.Objects
    }
    return nil
}

func (object *Snmp_BulkStats_Objects_Object) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["objects"] = &object.Objects
    return children
}

func (object *Snmp_BulkStats_Objects_Object) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-list-name"] = object.ObjectListName
    leafs["type"] = object.Type
    return leafs
}

func (object *Snmp_BulkStats_Objects_Object) GetBundleName() string { return "cisco_ios_xr" }

func (object *Snmp_BulkStats_Objects_Object) GetYangName() string { return "object" }

func (object *Snmp_BulkStats_Objects_Object) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (object *Snmp_BulkStats_Objects_Object) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (object *Snmp_BulkStats_Objects_Object) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (object *Snmp_BulkStats_Objects_Object) SetParent(parent types.Entity) { object.parent = parent }

func (object *Snmp_BulkStats_Objects_Object) GetParent() types.Entity { return object.parent }

func (object *Snmp_BulkStats_Objects_Object) GetParentYangName() string { return "objects" }

// Snmp_BulkStats_Objects_Object_Objects
// Configure an object List
type Snmp_BulkStats_Objects_Object_Objects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object name or OID. The type is slice of
    // Snmp_BulkStats_Objects_Object_Objects_Object.
    Object []Snmp_BulkStats_Objects_Object_Objects_Object
}

func (objects *Snmp_BulkStats_Objects_Object_Objects) GetFilter() yfilter.YFilter { return objects.YFilter }

func (objects *Snmp_BulkStats_Objects_Object_Objects) SetFilter(yf yfilter.YFilter) { objects.YFilter = yf }

func (objects *Snmp_BulkStats_Objects_Object_Objects) GetGoName(yname string) string {
    if yname == "object" { return "Object" }
    return ""
}

func (objects *Snmp_BulkStats_Objects_Object_Objects) GetSegmentPath() string {
    return "objects"
}

func (objects *Snmp_BulkStats_Objects_Object_Objects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "object" {
        for _, c := range objects.Object {
            if objects.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_BulkStats_Objects_Object_Objects_Object{}
        objects.Object = append(objects.Object, child)
        return &objects.Object[len(objects.Object)-1]
    }
    return nil
}

func (objects *Snmp_BulkStats_Objects_Object_Objects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range objects.Object {
        children[objects.Object[i].GetSegmentPath()] = &objects.Object[i]
    }
    return children
}

func (objects *Snmp_BulkStats_Objects_Object_Objects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (objects *Snmp_BulkStats_Objects_Object_Objects) GetBundleName() string { return "cisco_ios_xr" }

func (objects *Snmp_BulkStats_Objects_Object_Objects) GetYangName() string { return "objects" }

func (objects *Snmp_BulkStats_Objects_Object_Objects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (objects *Snmp_BulkStats_Objects_Object_Objects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (objects *Snmp_BulkStats_Objects_Object_Objects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (objects *Snmp_BulkStats_Objects_Object_Objects) SetParent(parent types.Entity) { objects.parent = parent }

func (objects *Snmp_BulkStats_Objects_Object_Objects) GetParent() types.Entity { return objects.parent }

func (objects *Snmp_BulkStats_Objects_Object_Objects) GetParentYangName() string { return "object" }

// Snmp_BulkStats_Objects_Object_Objects_Object
// Object name or OID
type Snmp_BulkStats_Objects_Object_Objects_Object struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object name or OID . The type is string.
    Oid interface{}
}

func (object *Snmp_BulkStats_Objects_Object_Objects_Object) GetFilter() yfilter.YFilter { return object.YFilter }

func (object *Snmp_BulkStats_Objects_Object_Objects_Object) SetFilter(yf yfilter.YFilter) { object.YFilter = yf }

func (object *Snmp_BulkStats_Objects_Object_Objects_Object) GetGoName(yname string) string {
    if yname == "oid" { return "Oid" }
    return ""
}

func (object *Snmp_BulkStats_Objects_Object_Objects_Object) GetSegmentPath() string {
    return "object" + "[oid='" + fmt.Sprintf("%v", object.Oid) + "']"
}

func (object *Snmp_BulkStats_Objects_Object_Objects_Object) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (object *Snmp_BulkStats_Objects_Object_Objects_Object) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (object *Snmp_BulkStats_Objects_Object_Objects_Object) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["oid"] = object.Oid
    return leafs
}

func (object *Snmp_BulkStats_Objects_Object_Objects_Object) GetBundleName() string { return "cisco_ios_xr" }

func (object *Snmp_BulkStats_Objects_Object_Objects_Object) GetYangName() string { return "object" }

func (object *Snmp_BulkStats_Objects_Object_Objects_Object) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (object *Snmp_BulkStats_Objects_Object_Objects_Object) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (object *Snmp_BulkStats_Objects_Object_Objects_Object) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (object *Snmp_BulkStats_Objects_Object_Objects_Object) SetParent(parent types.Entity) { object.parent = parent }

func (object *Snmp_BulkStats_Objects_Object_Objects_Object) GetParent() types.Entity { return object.parent }

func (object *Snmp_BulkStats_Objects_Object_Objects_Object) GetParentYangName() string { return "objects" }

// Snmp_BulkStats_Transfers
// Periodicity for the transfer of bulk data in
// minutes
type Snmp_BulkStats_Transfers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of bulk transfer. The type is slice of
    // Snmp_BulkStats_Transfers_Transfer.
    Transfer []Snmp_BulkStats_Transfers_Transfer
}

func (transfers *Snmp_BulkStats_Transfers) GetFilter() yfilter.YFilter { return transfers.YFilter }

func (transfers *Snmp_BulkStats_Transfers) SetFilter(yf yfilter.YFilter) { transfers.YFilter = yf }

func (transfers *Snmp_BulkStats_Transfers) GetGoName(yname string) string {
    if yname == "transfer" { return "Transfer" }
    return ""
}

func (transfers *Snmp_BulkStats_Transfers) GetSegmentPath() string {
    return "transfers"
}

func (transfers *Snmp_BulkStats_Transfers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transfer" {
        for _, c := range transfers.Transfer {
            if transfers.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_BulkStats_Transfers_Transfer{}
        transfers.Transfer = append(transfers.Transfer, child)
        return &transfers.Transfer[len(transfers.Transfer)-1]
    }
    return nil
}

func (transfers *Snmp_BulkStats_Transfers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range transfers.Transfer {
        children[transfers.Transfer[i].GetSegmentPath()] = &transfers.Transfer[i]
    }
    return children
}

func (transfers *Snmp_BulkStats_Transfers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (transfers *Snmp_BulkStats_Transfers) GetBundleName() string { return "cisco_ios_xr" }

func (transfers *Snmp_BulkStats_Transfers) GetYangName() string { return "transfers" }

func (transfers *Snmp_BulkStats_Transfers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transfers *Snmp_BulkStats_Transfers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transfers *Snmp_BulkStats_Transfers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transfers *Snmp_BulkStats_Transfers) SetParent(parent types.Entity) { transfers.parent = parent }

func (transfers *Snmp_BulkStats_Transfers) GetParent() types.Entity { return transfers.parent }

func (transfers *Snmp_BulkStats_Transfers) GetParentYangName() string { return "bulk-stats" }

// Snmp_BulkStats_Transfers_Transfer
// Name of bulk transfer
type Snmp_BulkStats_Transfers_Transfer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of bulk transfer. The type is string with
    // length: 1..32.
    TransferName interface{}

    // FTP or rcp or TFTP can be used for file transfer. The type is string.
    Secondary interface{}

    // Configure transfer list name. The type is interface{}.
    Type interface{}

    // Bulkstat data file maximum size in bytes. The type is interface{} with
    // range: 1024..2147483647. Units are byte.
    BufferSize interface{}

    // Retention period in minutes. The type is interface{} with range: 0..20000.
    // Units are minute.
    Retain interface{}

    // Format of the bulk data file. The type is SnmpBulkstatFileFormat.
    Format interface{}

    // Number of transmission retries. The type is interface{} with range: 0..100.
    Retry interface{}

    // Start Data Collection for this Configuration. The type is interface{}.
    Enable interface{}

    // FTP or rcp or TFTP can be used for file transfer. The type is string.
    Primary interface{}

    // Periodicity for the transfer of bulk data in minutes. The type is
    // interface{} with range: -2147483648..2147483647. Units are minute.
    Interval interface{}

    // Schema that contains objects to be collected.
    TransferSchemas Snmp_BulkStats_Transfers_Transfer_TransferSchemas
}

func (transfer *Snmp_BulkStats_Transfers_Transfer) GetFilter() yfilter.YFilter { return transfer.YFilter }

func (transfer *Snmp_BulkStats_Transfers_Transfer) SetFilter(yf yfilter.YFilter) { transfer.YFilter = yf }

func (transfer *Snmp_BulkStats_Transfers_Transfer) GetGoName(yname string) string {
    if yname == "transfer-name" { return "TransferName" }
    if yname == "secondary" { return "Secondary" }
    if yname == "type" { return "Type" }
    if yname == "buffer-size" { return "BufferSize" }
    if yname == "retain" { return "Retain" }
    if yname == "format" { return "Format" }
    if yname == "retry" { return "Retry" }
    if yname == "enable" { return "Enable" }
    if yname == "primary" { return "Primary" }
    if yname == "interval" { return "Interval" }
    if yname == "transfer-schemas" { return "TransferSchemas" }
    return ""
}

func (transfer *Snmp_BulkStats_Transfers_Transfer) GetSegmentPath() string {
    return "transfer" + "[transfer-name='" + fmt.Sprintf("%v", transfer.TransferName) + "']"
}

func (transfer *Snmp_BulkStats_Transfers_Transfer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transfer-schemas" {
        return &transfer.TransferSchemas
    }
    return nil
}

func (transfer *Snmp_BulkStats_Transfers_Transfer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["transfer-schemas"] = &transfer.TransferSchemas
    return children
}

func (transfer *Snmp_BulkStats_Transfers_Transfer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["transfer-name"] = transfer.TransferName
    leafs["secondary"] = transfer.Secondary
    leafs["type"] = transfer.Type
    leafs["buffer-size"] = transfer.BufferSize
    leafs["retain"] = transfer.Retain
    leafs["format"] = transfer.Format
    leafs["retry"] = transfer.Retry
    leafs["enable"] = transfer.Enable
    leafs["primary"] = transfer.Primary
    leafs["interval"] = transfer.Interval
    return leafs
}

func (transfer *Snmp_BulkStats_Transfers_Transfer) GetBundleName() string { return "cisco_ios_xr" }

func (transfer *Snmp_BulkStats_Transfers_Transfer) GetYangName() string { return "transfer" }

func (transfer *Snmp_BulkStats_Transfers_Transfer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transfer *Snmp_BulkStats_Transfers_Transfer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transfer *Snmp_BulkStats_Transfers_Transfer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transfer *Snmp_BulkStats_Transfers_Transfer) SetParent(parent types.Entity) { transfer.parent = parent }

func (transfer *Snmp_BulkStats_Transfers_Transfer) GetParent() types.Entity { return transfer.parent }

func (transfer *Snmp_BulkStats_Transfers_Transfer) GetParentYangName() string { return "transfers" }

// Snmp_BulkStats_Transfers_Transfer_TransferSchemas
// Schema that contains objects to be collected
type Snmp_BulkStats_Transfers_Transfer_TransferSchemas struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Schema that contains objects to be collected. The type is slice of
    // Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema.
    TransferSchema []Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema
}

func (transferSchemas *Snmp_BulkStats_Transfers_Transfer_TransferSchemas) GetFilter() yfilter.YFilter { return transferSchemas.YFilter }

func (transferSchemas *Snmp_BulkStats_Transfers_Transfer_TransferSchemas) SetFilter(yf yfilter.YFilter) { transferSchemas.YFilter = yf }

func (transferSchemas *Snmp_BulkStats_Transfers_Transfer_TransferSchemas) GetGoName(yname string) string {
    if yname == "transfer-schema" { return "TransferSchema" }
    return ""
}

func (transferSchemas *Snmp_BulkStats_Transfers_Transfer_TransferSchemas) GetSegmentPath() string {
    return "transfer-schemas"
}

func (transferSchemas *Snmp_BulkStats_Transfers_Transfer_TransferSchemas) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "transfer-schema" {
        for _, c := range transferSchemas.TransferSchema {
            if transferSchemas.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema{}
        transferSchemas.TransferSchema = append(transferSchemas.TransferSchema, child)
        return &transferSchemas.TransferSchema[len(transferSchemas.TransferSchema)-1]
    }
    return nil
}

func (transferSchemas *Snmp_BulkStats_Transfers_Transfer_TransferSchemas) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range transferSchemas.TransferSchema {
        children[transferSchemas.TransferSchema[i].GetSegmentPath()] = &transferSchemas.TransferSchema[i]
    }
    return children
}

func (transferSchemas *Snmp_BulkStats_Transfers_Transfer_TransferSchemas) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (transferSchemas *Snmp_BulkStats_Transfers_Transfer_TransferSchemas) GetBundleName() string { return "cisco_ios_xr" }

func (transferSchemas *Snmp_BulkStats_Transfers_Transfer_TransferSchemas) GetYangName() string { return "transfer-schemas" }

func (transferSchemas *Snmp_BulkStats_Transfers_Transfer_TransferSchemas) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transferSchemas *Snmp_BulkStats_Transfers_Transfer_TransferSchemas) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transferSchemas *Snmp_BulkStats_Transfers_Transfer_TransferSchemas) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transferSchemas *Snmp_BulkStats_Transfers_Transfer_TransferSchemas) SetParent(parent types.Entity) { transferSchemas.parent = parent }

func (transferSchemas *Snmp_BulkStats_Transfers_Transfer_TransferSchemas) GetParent() types.Entity { return transferSchemas.parent }

func (transferSchemas *Snmp_BulkStats_Transfers_Transfer_TransferSchemas) GetParentYangName() string { return "transfer" }

// Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema
// Schema that contains objects to be collected
type Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Schema that contains objects to be collected. The
    // type is string with length: 1..32.
    SchemaName interface{}
}

func (transferSchema *Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema) GetFilter() yfilter.YFilter { return transferSchema.YFilter }

func (transferSchema *Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema) SetFilter(yf yfilter.YFilter) { transferSchema.YFilter = yf }

func (transferSchema *Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema) GetGoName(yname string) string {
    if yname == "schema-name" { return "SchemaName" }
    return ""
}

func (transferSchema *Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema) GetSegmentPath() string {
    return "transfer-schema" + "[schema-name='" + fmt.Sprintf("%v", transferSchema.SchemaName) + "']"
}

func (transferSchema *Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (transferSchema *Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (transferSchema *Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["schema-name"] = transferSchema.SchemaName
    return leafs
}

func (transferSchema *Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema) GetBundleName() string { return "cisco_ios_xr" }

func (transferSchema *Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema) GetYangName() string { return "transfer-schema" }

func (transferSchema *Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (transferSchema *Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (transferSchema *Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (transferSchema *Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema) SetParent(parent types.Entity) { transferSchema.parent = parent }

func (transferSchema *Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema) GetParent() types.Entity { return transferSchema.parent }

func (transferSchema *Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema) GetParentYangName() string { return "transfer-schemas" }

// Snmp_DefaultCommunityMaps
// Container class to hold unencrpted community map
type Snmp_DefaultCommunityMaps struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unencrpted SNMP community map name . The type is slice of
    // Snmp_DefaultCommunityMaps_DefaultCommunityMap.
    DefaultCommunityMap []Snmp_DefaultCommunityMaps_DefaultCommunityMap
}

func (defaultCommunityMaps *Snmp_DefaultCommunityMaps) GetFilter() yfilter.YFilter { return defaultCommunityMaps.YFilter }

func (defaultCommunityMaps *Snmp_DefaultCommunityMaps) SetFilter(yf yfilter.YFilter) { defaultCommunityMaps.YFilter = yf }

func (defaultCommunityMaps *Snmp_DefaultCommunityMaps) GetGoName(yname string) string {
    if yname == "default-community-map" { return "DefaultCommunityMap" }
    return ""
}

func (defaultCommunityMaps *Snmp_DefaultCommunityMaps) GetSegmentPath() string {
    return "default-community-maps"
}

func (defaultCommunityMaps *Snmp_DefaultCommunityMaps) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-community-map" {
        for _, c := range defaultCommunityMaps.DefaultCommunityMap {
            if defaultCommunityMaps.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_DefaultCommunityMaps_DefaultCommunityMap{}
        defaultCommunityMaps.DefaultCommunityMap = append(defaultCommunityMaps.DefaultCommunityMap, child)
        return &defaultCommunityMaps.DefaultCommunityMap[len(defaultCommunityMaps.DefaultCommunityMap)-1]
    }
    return nil
}

func (defaultCommunityMaps *Snmp_DefaultCommunityMaps) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range defaultCommunityMaps.DefaultCommunityMap {
        children[defaultCommunityMaps.DefaultCommunityMap[i].GetSegmentPath()] = &defaultCommunityMaps.DefaultCommunityMap[i]
    }
    return children
}

func (defaultCommunityMaps *Snmp_DefaultCommunityMaps) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (defaultCommunityMaps *Snmp_DefaultCommunityMaps) GetBundleName() string { return "cisco_ios_xr" }

func (defaultCommunityMaps *Snmp_DefaultCommunityMaps) GetYangName() string { return "default-community-maps" }

func (defaultCommunityMaps *Snmp_DefaultCommunityMaps) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultCommunityMaps *Snmp_DefaultCommunityMaps) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultCommunityMaps *Snmp_DefaultCommunityMaps) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultCommunityMaps *Snmp_DefaultCommunityMaps) SetParent(parent types.Entity) { defaultCommunityMaps.parent = parent }

func (defaultCommunityMaps *Snmp_DefaultCommunityMaps) GetParent() types.Entity { return defaultCommunityMaps.parent }

func (defaultCommunityMaps *Snmp_DefaultCommunityMaps) GetParentYangName() string { return "snmp" }

// Snmp_DefaultCommunityMaps_DefaultCommunityMap
// Unencrpted SNMP community map name 
type Snmp_DefaultCommunityMaps_DefaultCommunityMap struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SNMP community map. The type is string with
    // length: 1..128.
    CommunityName interface{}

    // SNMP Context Name . The type is string.
    Context interface{}

    // SNMP Security Name . The type is string.
    Security interface{}

    // target list name . The type is string.
    TargetList interface{}
}

func (defaultCommunityMap *Snmp_DefaultCommunityMaps_DefaultCommunityMap) GetFilter() yfilter.YFilter { return defaultCommunityMap.YFilter }

func (defaultCommunityMap *Snmp_DefaultCommunityMaps_DefaultCommunityMap) SetFilter(yf yfilter.YFilter) { defaultCommunityMap.YFilter = yf }

func (defaultCommunityMap *Snmp_DefaultCommunityMaps_DefaultCommunityMap) GetGoName(yname string) string {
    if yname == "community-name" { return "CommunityName" }
    if yname == "context" { return "Context" }
    if yname == "security" { return "Security" }
    if yname == "target-list" { return "TargetList" }
    return ""
}

func (defaultCommunityMap *Snmp_DefaultCommunityMaps_DefaultCommunityMap) GetSegmentPath() string {
    return "default-community-map" + "[community-name='" + fmt.Sprintf("%v", defaultCommunityMap.CommunityName) + "']"
}

func (defaultCommunityMap *Snmp_DefaultCommunityMaps_DefaultCommunityMap) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defaultCommunityMap *Snmp_DefaultCommunityMaps_DefaultCommunityMap) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defaultCommunityMap *Snmp_DefaultCommunityMaps_DefaultCommunityMap) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-name"] = defaultCommunityMap.CommunityName
    leafs["context"] = defaultCommunityMap.Context
    leafs["security"] = defaultCommunityMap.Security
    leafs["target-list"] = defaultCommunityMap.TargetList
    return leafs
}

func (defaultCommunityMap *Snmp_DefaultCommunityMaps_DefaultCommunityMap) GetBundleName() string { return "cisco_ios_xr" }

func (defaultCommunityMap *Snmp_DefaultCommunityMaps_DefaultCommunityMap) GetYangName() string { return "default-community-map" }

func (defaultCommunityMap *Snmp_DefaultCommunityMaps_DefaultCommunityMap) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultCommunityMap *Snmp_DefaultCommunityMaps_DefaultCommunityMap) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultCommunityMap *Snmp_DefaultCommunityMaps_DefaultCommunityMap) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultCommunityMap *Snmp_DefaultCommunityMaps_DefaultCommunityMap) SetParent(parent types.Entity) { defaultCommunityMap.parent = parent }

func (defaultCommunityMap *Snmp_DefaultCommunityMaps_DefaultCommunityMap) GetParent() types.Entity { return defaultCommunityMap.parent }

func (defaultCommunityMap *Snmp_DefaultCommunityMaps_DefaultCommunityMap) GetParentYangName() string { return "default-community-maps" }

// Snmp_OverloadControl
// Set overload control params for handling
// incoming messages
// This type is a presence type.
type Snmp_OverloadControl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Drop time in seconds for incoming queue (default 1 sec). The type is
    // interface{} with range: 0..300. This attribute is mandatory. Units are
    // second.
    DropTime interface{}

    // Throttle time in milliseconds for incoming queue (default 500 msec). The
    // type is interface{} with range: 0..1000. This attribute is mandatory. Units
    // are millisecond.
    ThrottleRate interface{}
}

func (overloadControl *Snmp_OverloadControl) GetFilter() yfilter.YFilter { return overloadControl.YFilter }

func (overloadControl *Snmp_OverloadControl) SetFilter(yf yfilter.YFilter) { overloadControl.YFilter = yf }

func (overloadControl *Snmp_OverloadControl) GetGoName(yname string) string {
    if yname == "drop-time" { return "DropTime" }
    if yname == "throttle-rate" { return "ThrottleRate" }
    return ""
}

func (overloadControl *Snmp_OverloadControl) GetSegmentPath() string {
    return "overload-control"
}

func (overloadControl *Snmp_OverloadControl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (overloadControl *Snmp_OverloadControl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (overloadControl *Snmp_OverloadControl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["drop-time"] = overloadControl.DropTime
    leafs["throttle-rate"] = overloadControl.ThrottleRate
    return leafs
}

func (overloadControl *Snmp_OverloadControl) GetBundleName() string { return "cisco_ios_xr" }

func (overloadControl *Snmp_OverloadControl) GetYangName() string { return "overload-control" }

func (overloadControl *Snmp_OverloadControl) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (overloadControl *Snmp_OverloadControl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (overloadControl *Snmp_OverloadControl) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (overloadControl *Snmp_OverloadControl) SetParent(parent types.Entity) { overloadControl.parent = parent }

func (overloadControl *Snmp_OverloadControl) GetParent() types.Entity { return overloadControl.parent }

func (overloadControl *Snmp_OverloadControl) GetParentYangName() string { return "snmp" }

// Snmp_Timeouts
// SNMP timeouts
type Snmp_Timeouts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Duplicate request feature timeout. The type is interface{} with range:
    // 0..20. Units are second. The default value is 1.
    Duplicates interface{}

    // incoming queue drop feature timeout. The type is interface{} with range:
    // 0..20. Units are second. The default value is 10.
    InQdrop interface{}

    // Sub-Agent Request timeout. The type is interface{} with range: 1..20. Units
    // are second. The default value is 10.
    Subagent interface{}

    // SNMP pdu statistics timeout. The type is interface{} with range: 1..10.
    // Units are second. The default value is 2.
    PduStats interface{}
}

func (timeouts *Snmp_Timeouts) GetFilter() yfilter.YFilter { return timeouts.YFilter }

func (timeouts *Snmp_Timeouts) SetFilter(yf yfilter.YFilter) { timeouts.YFilter = yf }

func (timeouts *Snmp_Timeouts) GetGoName(yname string) string {
    if yname == "duplicates" { return "Duplicates" }
    if yname == "in-qdrop" { return "InQdrop" }
    if yname == "subagent" { return "Subagent" }
    if yname == "pdu-stats" { return "PduStats" }
    return ""
}

func (timeouts *Snmp_Timeouts) GetSegmentPath() string {
    return "timeouts"
}

func (timeouts *Snmp_Timeouts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timeouts *Snmp_Timeouts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timeouts *Snmp_Timeouts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["duplicates"] = timeouts.Duplicates
    leafs["in-qdrop"] = timeouts.InQdrop
    leafs["subagent"] = timeouts.Subagent
    leafs["pdu-stats"] = timeouts.PduStats
    return leafs
}

func (timeouts *Snmp_Timeouts) GetBundleName() string { return "cisco_ios_xr" }

func (timeouts *Snmp_Timeouts) GetYangName() string { return "timeouts" }

func (timeouts *Snmp_Timeouts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timeouts *Snmp_Timeouts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timeouts *Snmp_Timeouts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timeouts *Snmp_Timeouts) SetParent(parent types.Entity) { timeouts.parent = parent }

func (timeouts *Snmp_Timeouts) GetParent() types.Entity { return timeouts.parent }

func (timeouts *Snmp_Timeouts) GetParentYangName() string { return "snmp" }

// Snmp_Users
// Define a user who can access the SNMP engine
type Snmp_Users struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the user. The type is slice of Snmp_Users_User.
    User []Snmp_Users_User
}

func (users *Snmp_Users) GetFilter() yfilter.YFilter { return users.YFilter }

func (users *Snmp_Users) SetFilter(yf yfilter.YFilter) { users.YFilter = yf }

func (users *Snmp_Users) GetGoName(yname string) string {
    if yname == "user" { return "User" }
    return ""
}

func (users *Snmp_Users) GetSegmentPath() string {
    return "users"
}

func (users *Snmp_Users) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "user" {
        for _, c := range users.User {
            if users.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Users_User{}
        users.User = append(users.User, child)
        return &users.User[len(users.User)-1]
    }
    return nil
}

func (users *Snmp_Users) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range users.User {
        children[users.User[i].GetSegmentPath()] = &users.User[i]
    }
    return children
}

func (users *Snmp_Users) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (users *Snmp_Users) GetBundleName() string { return "cisco_ios_xr" }

func (users *Snmp_Users) GetYangName() string { return "users" }

func (users *Snmp_Users) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (users *Snmp_Users) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (users *Snmp_Users) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (users *Snmp_Users) SetParent(parent types.Entity) { users.parent = parent }

func (users *Snmp_Users) GetParent() types.Entity { return users.parent }

func (users *Snmp_Users) GetParentYangName() string { return "snmp" }

// Snmp_Users_User
// Name of the user
type Snmp_Users_User struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the user. The type is string.
    UserName interface{}

    // Group to which the user belongs. The type is string. This attribute is
    // mandatory.
    GroupName interface{}

    // SNMP version to be used. v1,v2c or v3. The type is UserSnmpVersion. This
    // attribute is mandatory.
    Version interface{}

    // Flag to indicate that authentication password is configred for version 3.
    // The type is interface{}.
    AuthenticationPasswordConfigured interface{}

    // The algorithm used md5 or sha. The type is SnmpHashAlgorithm.
    Algorithm interface{}

    // The authentication password. The type is string with pattern:
    // (!.+)|([^!].+).
    AuthenticationPassword interface{}

    // Flag to indicate that the privacy password is configured for version 3. The
    // type is interface{}.
    PrivacyPasswordConfigured interface{}

    // The algorithm used des56 or aes128 or aes192or aes256 or 3des. The type is
    // SnmpPrivAlgorithm.
    PrivAlgorithm interface{}

    // The privacy password. The type is string with pattern: (!.+)|([^!].+).
    PrivacyPassword interface{}

    // Access-list type. The type is Snmpacl.
    V4AclType interface{}

    // Ipv4 Access-list name. The type is string.
    V4AccessList interface{}

    // Access-list type. The type is Snmpacl.
    V6AclType interface{}

    // Ipv6 Access-list name. The type is string.
    V6AccessList interface{}

    // The system access either SDROwner or SystemOwner. The type is
    // SnmpOwnerAccess.
    Owner interface{}

    // IP address of remote SNMP entity. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    RemoteAddress interface{}

    // UDP port number. The type is interface{} with range: 1..65535.
    Port interface{}
}

func (user *Snmp_Users_User) GetFilter() yfilter.YFilter { return user.YFilter }

func (user *Snmp_Users_User) SetFilter(yf yfilter.YFilter) { user.YFilter = yf }

func (user *Snmp_Users_User) GetGoName(yname string) string {
    if yname == "user-name" { return "UserName" }
    if yname == "group-name" { return "GroupName" }
    if yname == "version" { return "Version" }
    if yname == "authentication-password-configured" { return "AuthenticationPasswordConfigured" }
    if yname == "algorithm" { return "Algorithm" }
    if yname == "authentication-password" { return "AuthenticationPassword" }
    if yname == "privacy-password-configured" { return "PrivacyPasswordConfigured" }
    if yname == "priv-algorithm" { return "PrivAlgorithm" }
    if yname == "privacy-password" { return "PrivacyPassword" }
    if yname == "v4acl-type" { return "V4AclType" }
    if yname == "v4-access-list" { return "V4AccessList" }
    if yname == "v6acl-type" { return "V6AclType" }
    if yname == "v6-access-list" { return "V6AccessList" }
    if yname == "owner" { return "Owner" }
    if yname == "remote-address" { return "RemoteAddress" }
    if yname == "port" { return "Port" }
    return ""
}

func (user *Snmp_Users_User) GetSegmentPath() string {
    return "user" + "[user-name='" + fmt.Sprintf("%v", user.UserName) + "']"
}

func (user *Snmp_Users_User) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (user *Snmp_Users_User) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (user *Snmp_Users_User) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["user-name"] = user.UserName
    leafs["group-name"] = user.GroupName
    leafs["version"] = user.Version
    leafs["authentication-password-configured"] = user.AuthenticationPasswordConfigured
    leafs["algorithm"] = user.Algorithm
    leafs["authentication-password"] = user.AuthenticationPassword
    leafs["privacy-password-configured"] = user.PrivacyPasswordConfigured
    leafs["priv-algorithm"] = user.PrivAlgorithm
    leafs["privacy-password"] = user.PrivacyPassword
    leafs["v4acl-type"] = user.V4AclType
    leafs["v4-access-list"] = user.V4AccessList
    leafs["v6acl-type"] = user.V6AclType
    leafs["v6-access-list"] = user.V6AccessList
    leafs["owner"] = user.Owner
    leafs["remote-address"] = user.RemoteAddress
    leafs["port"] = user.Port
    return leafs
}

func (user *Snmp_Users_User) GetBundleName() string { return "cisco_ios_xr" }

func (user *Snmp_Users_User) GetYangName() string { return "user" }

func (user *Snmp_Users_User) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (user *Snmp_Users_User) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (user *Snmp_Users_User) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (user *Snmp_Users_User) SetParent(parent types.Entity) { user.parent = parent }

func (user *Snmp_Users_User) GetParent() types.Entity { return user.parent }

func (user *Snmp_Users_User) GetParentYangName() string { return "users" }

// Snmp_Vrfs
// SNMP VRF configuration commands
type Snmp_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name. The type is slice of Snmp_Vrfs_Vrf.
    Vrf []Snmp_Vrfs_Vrf
}

func (vrfs *Snmp_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Snmp_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Snmp_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Snmp_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Snmp_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Snmp_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Snmp_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Snmp_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Snmp_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Snmp_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Snmp_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Snmp_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Snmp_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Snmp_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Snmp_Vrfs) GetParentYangName() string { return "snmp" }

// Snmp_Vrfs_Vrf
// VRF name
type Snmp_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // Specify hosts to receive SNMP notifications.
    TrapHosts Snmp_Vrfs_Vrf_TrapHosts

    // List of Context Names.
    Contexts Snmp_Vrfs_Vrf_Contexts

    // List of context names.
    ContextMappings Snmp_Vrfs_Vrf_ContextMappings
}

func (vrf *Snmp_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Snmp_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Snmp_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "trap-hosts" { return "TrapHosts" }
    if yname == "contexts" { return "Contexts" }
    if yname == "context-mappings" { return "ContextMappings" }
    return ""
}

func (vrf *Snmp_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[name='" + fmt.Sprintf("%v", vrf.Name) + "']"
}

func (vrf *Snmp_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trap-hosts" {
        return &vrf.TrapHosts
    }
    if childYangName == "contexts" {
        return &vrf.Contexts
    }
    if childYangName == "context-mappings" {
        return &vrf.ContextMappings
    }
    return nil
}

func (vrf *Snmp_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["trap-hosts"] = &vrf.TrapHosts
    children["contexts"] = &vrf.Contexts
    children["context-mappings"] = &vrf.ContextMappings
    return children
}

func (vrf *Snmp_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = vrf.Name
    return leafs
}

func (vrf *Snmp_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Snmp_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Snmp_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Snmp_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Snmp_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Snmp_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Snmp_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Snmp_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Snmp_Vrfs_Vrf_TrapHosts
// Specify hosts to receive SNMP notifications
type Snmp_Vrfs_Vrf_TrapHosts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify hosts to receive SNMP notifications. The type is slice of
    // Snmp_Vrfs_Vrf_TrapHosts_TrapHost.
    TrapHost []Snmp_Vrfs_Vrf_TrapHosts_TrapHost
}

func (trapHosts *Snmp_Vrfs_Vrf_TrapHosts) GetFilter() yfilter.YFilter { return trapHosts.YFilter }

func (trapHosts *Snmp_Vrfs_Vrf_TrapHosts) SetFilter(yf yfilter.YFilter) { trapHosts.YFilter = yf }

func (trapHosts *Snmp_Vrfs_Vrf_TrapHosts) GetGoName(yname string) string {
    if yname == "trap-host" { return "TrapHost" }
    return ""
}

func (trapHosts *Snmp_Vrfs_Vrf_TrapHosts) GetSegmentPath() string {
    return "trap-hosts"
}

func (trapHosts *Snmp_Vrfs_Vrf_TrapHosts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trap-host" {
        for _, c := range trapHosts.TrapHost {
            if trapHosts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Vrfs_Vrf_TrapHosts_TrapHost{}
        trapHosts.TrapHost = append(trapHosts.TrapHost, child)
        return &trapHosts.TrapHost[len(trapHosts.TrapHost)-1]
    }
    return nil
}

func (trapHosts *Snmp_Vrfs_Vrf_TrapHosts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trapHosts.TrapHost {
        children[trapHosts.TrapHost[i].GetSegmentPath()] = &trapHosts.TrapHost[i]
    }
    return children
}

func (trapHosts *Snmp_Vrfs_Vrf_TrapHosts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trapHosts *Snmp_Vrfs_Vrf_TrapHosts) GetBundleName() string { return "cisco_ios_xr" }

func (trapHosts *Snmp_Vrfs_Vrf_TrapHosts) GetYangName() string { return "trap-hosts" }

func (trapHosts *Snmp_Vrfs_Vrf_TrapHosts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapHosts *Snmp_Vrfs_Vrf_TrapHosts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapHosts *Snmp_Vrfs_Vrf_TrapHosts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapHosts *Snmp_Vrfs_Vrf_TrapHosts) SetParent(parent types.Entity) { trapHosts.parent = parent }

func (trapHosts *Snmp_Vrfs_Vrf_TrapHosts) GetParent() types.Entity { return trapHosts.parent }

func (trapHosts *Snmp_Vrfs_Vrf_TrapHosts) GetParentYangName() string { return "vrf" }

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost
// Specify hosts to receive SNMP notifications
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of SNMP notification host. The type is
    // one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Container class for defining Clear/encrypt communities for a trap host.
    EncryptedUserCommunities Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities

    // Container class for defining notification type for a Inform host.
    InformHost Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost

    // Container class for defining communities for a trap host.
    DefaultUserCommunities Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities
}

func (trapHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost) GetFilter() yfilter.YFilter { return trapHost.YFilter }

func (trapHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost) SetFilter(yf yfilter.YFilter) { trapHost.YFilter = yf }

func (trapHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "encrypted-user-communities" { return "EncryptedUserCommunities" }
    if yname == "inform-host" { return "InformHost" }
    if yname == "default-user-communities" { return "DefaultUserCommunities" }
    return ""
}

func (trapHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost) GetSegmentPath() string {
    return "trap-host" + "[ip-address='" + fmt.Sprintf("%v", trapHost.IpAddress) + "']"
}

func (trapHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "encrypted-user-communities" {
        return &trapHost.EncryptedUserCommunities
    }
    if childYangName == "inform-host" {
        return &trapHost.InformHost
    }
    if childYangName == "default-user-communities" {
        return &trapHost.DefaultUserCommunities
    }
    return nil
}

func (trapHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["encrypted-user-communities"] = &trapHost.EncryptedUserCommunities
    children["inform-host"] = &trapHost.InformHost
    children["default-user-communities"] = &trapHost.DefaultUserCommunities
    return children
}

func (trapHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = trapHost.IpAddress
    return leafs
}

func (trapHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost) GetBundleName() string { return "cisco_ios_xr" }

func (trapHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost) GetYangName() string { return "trap-host" }

func (trapHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost) SetParent(parent types.Entity) { trapHost.parent = parent }

func (trapHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost) GetParent() types.Entity { return trapHost.parent }

func (trapHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost) GetParentYangName() string { return "trap-hosts" }

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities
// Container class for defining Clear/encrypt
// communities for a trap host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear/Encrypt Community name associated with a trap host. The type is slice
    // of
    // Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity.
    EncryptedUserCommunity []Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity
}

func (encryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities) GetFilter() yfilter.YFilter { return encryptedUserCommunities.YFilter }

func (encryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities) SetFilter(yf yfilter.YFilter) { encryptedUserCommunities.YFilter = yf }

func (encryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities) GetGoName(yname string) string {
    if yname == "encrypted-user-community" { return "EncryptedUserCommunity" }
    return ""
}

func (encryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities) GetSegmentPath() string {
    return "encrypted-user-communities"
}

func (encryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "encrypted-user-community" {
        for _, c := range encryptedUserCommunities.EncryptedUserCommunity {
            if encryptedUserCommunities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity{}
        encryptedUserCommunities.EncryptedUserCommunity = append(encryptedUserCommunities.EncryptedUserCommunity, child)
        return &encryptedUserCommunities.EncryptedUserCommunity[len(encryptedUserCommunities.EncryptedUserCommunity)-1]
    }
    return nil
}

func (encryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range encryptedUserCommunities.EncryptedUserCommunity {
        children[encryptedUserCommunities.EncryptedUserCommunity[i].GetSegmentPath()] = &encryptedUserCommunities.EncryptedUserCommunity[i]
    }
    return children
}

func (encryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (encryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities) GetBundleName() string { return "cisco_ios_xr" }

func (encryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities) GetYangName() string { return "encrypted-user-communities" }

func (encryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (encryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (encryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (encryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities) SetParent(parent types.Entity) { encryptedUserCommunities.parent = parent }

func (encryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities) GetParent() types.Entity { return encryptedUserCommunities.parent }

func (encryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities) GetParentYangName() string { return "trap-host" }

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity
// Clear/Encrypt Community name associated with
// a trap host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SNMPv1/v2c community string or SNMPv3 user. The
    // type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    CommunityName interface{}

    // UDP port number. The type is interface{} with range: 1..65535.
    Port interface{}

    // SNMP Version to be used v1/v2c/v3. The type is string. This attribute is
    // mandatory.
    Version interface{}

    // Security level to be used noauth/auth/priv. The type is SnmpSecurityModel.
    SecurityLevel interface{}

    // Number to signify the feature traps that needs to be setBasicTrapTypes is
    // used for all traps except copy-completeSet this value to an integer
    // corresponding to the trapBGP 8192, CONFIG 4096,SYSLOG 131072,SNMP_TRAP
    // 1COPY_COMPLETE_TRAP 64To provide a combination of trap Add the respective
    // numbersValue must be set to 0 for all traps. The type is interface{} with
    // range: -2147483648..2147483647. This attribute is mandatory.
    BasicTrapTypes interface{}

    // Number to signify the feature traps that needs to be setUse this for
    // providing copy-complete trapValue must be set to 0 if not used. The type is
    // interface{} with range: -2147483648..2147483647. This attribute is
    // mandatory.
    AdvancedTrapTypes1 interface{}

    // Number to signify the feature traps that needs to be setvalue should always
    // to set as 0. The type is interface{} with range: -2147483648..2147483647.
    // This attribute is mandatory.
    AdvancedTrapTypes2 interface{}
}

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetFilter() yfilter.YFilter { return encryptedUserCommunity.YFilter }

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) SetFilter(yf yfilter.YFilter) { encryptedUserCommunity.YFilter = yf }

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetGoName(yname string) string {
    if yname == "community-name" { return "CommunityName" }
    if yname == "port" { return "Port" }
    if yname == "version" { return "Version" }
    if yname == "security-level" { return "SecurityLevel" }
    if yname == "basic-trap-types" { return "BasicTrapTypes" }
    if yname == "advanced-trap-types1" { return "AdvancedTrapTypes1" }
    if yname == "advanced-trap-types2" { return "AdvancedTrapTypes2" }
    return ""
}

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetSegmentPath() string {
    return "encrypted-user-community" + "[community-name='" + fmt.Sprintf("%v", encryptedUserCommunity.CommunityName) + "']"
}

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-name"] = encryptedUserCommunity.CommunityName
    leafs["port"] = encryptedUserCommunity.Port
    leafs["version"] = encryptedUserCommunity.Version
    leafs["security-level"] = encryptedUserCommunity.SecurityLevel
    leafs["basic-trap-types"] = encryptedUserCommunity.BasicTrapTypes
    leafs["advanced-trap-types1"] = encryptedUserCommunity.AdvancedTrapTypes1
    leafs["advanced-trap-types2"] = encryptedUserCommunity.AdvancedTrapTypes2
    return leafs
}

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetYangName() string { return "encrypted-user-community" }

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) SetParent(parent types.Entity) { encryptedUserCommunity.parent = parent }

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetParent() types.Entity { return encryptedUserCommunity.parent }

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetParentYangName() string { return "encrypted-user-communities" }

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost
// Container class for defining notification type
// for a Inform host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Container class for defining communities for a inform host.
    InformUserCommunities Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities

    // Container class for defining Clear/encrypt communities for a inform host.
    InformEncryptedUserCommunities Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities
}

func (informHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost) GetFilter() yfilter.YFilter { return informHost.YFilter }

func (informHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost) SetFilter(yf yfilter.YFilter) { informHost.YFilter = yf }

func (informHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost) GetGoName(yname string) string {
    if yname == "inform-user-communities" { return "InformUserCommunities" }
    if yname == "inform-encrypted-user-communities" { return "InformEncryptedUserCommunities" }
    return ""
}

func (informHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost) GetSegmentPath() string {
    return "inform-host"
}

func (informHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inform-user-communities" {
        return &informHost.InformUserCommunities
    }
    if childYangName == "inform-encrypted-user-communities" {
        return &informHost.InformEncryptedUserCommunities
    }
    return nil
}

func (informHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["inform-user-communities"] = &informHost.InformUserCommunities
    children["inform-encrypted-user-communities"] = &informHost.InformEncryptedUserCommunities
    return children
}

func (informHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (informHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost) GetBundleName() string { return "cisco_ios_xr" }

func (informHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost) GetYangName() string { return "inform-host" }

func (informHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (informHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (informHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (informHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost) SetParent(parent types.Entity) { informHost.parent = parent }

func (informHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost) GetParent() types.Entity { return informHost.parent }

func (informHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost) GetParentYangName() string { return "trap-host" }

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities
// Container class for defining communities for
// a inform host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unencrpted Community name associated with a inform host. The type is slice
    // of
    // Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity.
    InformUserCommunity []Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity
}

func (informUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetFilter() yfilter.YFilter { return informUserCommunities.YFilter }

func (informUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities) SetFilter(yf yfilter.YFilter) { informUserCommunities.YFilter = yf }

func (informUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetGoName(yname string) string {
    if yname == "inform-user-community" { return "InformUserCommunity" }
    return ""
}

func (informUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetSegmentPath() string {
    return "inform-user-communities"
}

func (informUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inform-user-community" {
        for _, c := range informUserCommunities.InformUserCommunity {
            if informUserCommunities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity{}
        informUserCommunities.InformUserCommunity = append(informUserCommunities.InformUserCommunity, child)
        return &informUserCommunities.InformUserCommunity[len(informUserCommunities.InformUserCommunity)-1]
    }
    return nil
}

func (informUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range informUserCommunities.InformUserCommunity {
        children[informUserCommunities.InformUserCommunity[i].GetSegmentPath()] = &informUserCommunities.InformUserCommunity[i]
    }
    return children
}

func (informUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (informUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetBundleName() string { return "cisco_ios_xr" }

func (informUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetYangName() string { return "inform-user-communities" }

func (informUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (informUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (informUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (informUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities) SetParent(parent types.Entity) { informUserCommunities.parent = parent }

func (informUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetParent() types.Entity { return informUserCommunities.parent }

func (informUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetParentYangName() string { return "inform-host" }

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity
// Unencrpted Community name associated with a
// inform host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SNMPv2c community string or SNMPv3 user. The type
    // is string with length: 1..128.
    CommunityName interface{}

    // UDP port number. The type is interface{} with range: 1..65535.
    Port interface{}

    // SNMP Version to be used v2c/v3. The type is string. This attribute is
    // mandatory.
    Version interface{}

    // Security level to be used noauth/auth/priv. The type is SnmpSecurityModel.
    SecurityLevel interface{}

    // Number to signify the feature traps that needs to be setBasicTrapTypes is
    // used for all traps except copy-completeSet this value to an integer
    // corresponding to the trapBGP 8192, CONFIG 4096,SYSLOG 131072 ,SNMP_TRAP
    // 1COPY_COMPLETE_TRAP 64To provide a combination of trap Add the respective
    // numbersValue must be set to 0 for all traps. The type is interface{} with
    // range: -2147483648..2147483647. This attribute is mandatory.
    BasicTrapTypes interface{}

    // Number to signify the feature traps that needs to be setUse this for
    // providing copy-complete trapValue must be set to 0 if not used. The type is
    // interface{} with range: -2147483648..2147483647. This attribute is
    // mandatory.
    AdvancedTrapTypes1 interface{}

    // Number to signify the feature traps that needs to be setvalue should always
    // to set as 0. The type is interface{} with range: -2147483648..2147483647.
    // This attribute is mandatory.
    AdvancedTrapTypes2 interface{}
}

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetFilter() yfilter.YFilter { return informUserCommunity.YFilter }

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) SetFilter(yf yfilter.YFilter) { informUserCommunity.YFilter = yf }

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetGoName(yname string) string {
    if yname == "community-name" { return "CommunityName" }
    if yname == "port" { return "Port" }
    if yname == "version" { return "Version" }
    if yname == "security-level" { return "SecurityLevel" }
    if yname == "basic-trap-types" { return "BasicTrapTypes" }
    if yname == "advanced-trap-types1" { return "AdvancedTrapTypes1" }
    if yname == "advanced-trap-types2" { return "AdvancedTrapTypes2" }
    return ""
}

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetSegmentPath() string {
    return "inform-user-community" + "[community-name='" + fmt.Sprintf("%v", informUserCommunity.CommunityName) + "']"
}

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-name"] = informUserCommunity.CommunityName
    leafs["port"] = informUserCommunity.Port
    leafs["version"] = informUserCommunity.Version
    leafs["security-level"] = informUserCommunity.SecurityLevel
    leafs["basic-trap-types"] = informUserCommunity.BasicTrapTypes
    leafs["advanced-trap-types1"] = informUserCommunity.AdvancedTrapTypes1
    leafs["advanced-trap-types2"] = informUserCommunity.AdvancedTrapTypes2
    return leafs
}

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetYangName() string { return "inform-user-community" }

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) SetParent(parent types.Entity) { informUserCommunity.parent = parent }

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetParent() types.Entity { return informUserCommunity.parent }

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetParentYangName() string { return "inform-user-communities" }

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities
// Container class for defining Clear/encrypt
// communities for a inform host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear/Encrypt Community name associated with a inform host. The type is
    // slice of
    // Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity.
    InformEncryptedUserCommunity []Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity
}

func (informEncryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetFilter() yfilter.YFilter { return informEncryptedUserCommunities.YFilter }

func (informEncryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) SetFilter(yf yfilter.YFilter) { informEncryptedUserCommunities.YFilter = yf }

func (informEncryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetGoName(yname string) string {
    if yname == "inform-encrypted-user-community" { return "InformEncryptedUserCommunity" }
    return ""
}

func (informEncryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetSegmentPath() string {
    return "inform-encrypted-user-communities"
}

func (informEncryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inform-encrypted-user-community" {
        for _, c := range informEncryptedUserCommunities.InformEncryptedUserCommunity {
            if informEncryptedUserCommunities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity{}
        informEncryptedUserCommunities.InformEncryptedUserCommunity = append(informEncryptedUserCommunities.InformEncryptedUserCommunity, child)
        return &informEncryptedUserCommunities.InformEncryptedUserCommunity[len(informEncryptedUserCommunities.InformEncryptedUserCommunity)-1]
    }
    return nil
}

func (informEncryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range informEncryptedUserCommunities.InformEncryptedUserCommunity {
        children[informEncryptedUserCommunities.InformEncryptedUserCommunity[i].GetSegmentPath()] = &informEncryptedUserCommunities.InformEncryptedUserCommunity[i]
    }
    return children
}

func (informEncryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (informEncryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetBundleName() string { return "cisco_ios_xr" }

func (informEncryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetYangName() string { return "inform-encrypted-user-communities" }

func (informEncryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (informEncryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (informEncryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (informEncryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) SetParent(parent types.Entity) { informEncryptedUserCommunities.parent = parent }

func (informEncryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetParent() types.Entity { return informEncryptedUserCommunities.parent }

func (informEncryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetParentYangName() string { return "inform-host" }

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity
// Clear/Encrypt Community name associated with
// a inform host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SNMPv2c community string or SNMPv3 user. The type
    // is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    CommunityName interface{}

    // UDP port number. The type is interface{} with range: 1..65535.
    Port interface{}

    // SNMP Version to be used v2c/v3. The type is string. This attribute is
    // mandatory.
    Version interface{}

    // Security level to be used noauth/auth/priv. The type is SnmpSecurityModel.
    SecurityLevel interface{}

    // Number to signify the feature traps that needs to be setBasicTrapTypes is
    // used for all traps except copy-completeSet this value to an integer
    // corresponding to the trapBGP 8192, CONFIG 4096,SYSLOG 131072 ,SNMP_TRAP
    // 1COPY_COMPLETE_TRAP 64To provide a combination of trap Add the respective
    // numbersValue must be set to 0 for all traps. The type is interface{} with
    // range: -2147483648..2147483647. This attribute is mandatory.
    BasicTrapTypes interface{}

    // Number to signify the feature traps that needs to be setUse this for
    // providing copy-complete trapValue must be set to 0 if not used. The type is
    // interface{} with range: -2147483648..2147483647. This attribute is
    // mandatory.
    AdvancedTrapTypes1 interface{}

    // Number to signify the feature traps that needs to be setvalue should always
    // to set as 0. The type is interface{} with range: -2147483648..2147483647.
    // This attribute is mandatory.
    AdvancedTrapTypes2 interface{}
}

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetFilter() yfilter.YFilter { return informEncryptedUserCommunity.YFilter }

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) SetFilter(yf yfilter.YFilter) { informEncryptedUserCommunity.YFilter = yf }

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetGoName(yname string) string {
    if yname == "community-name" { return "CommunityName" }
    if yname == "port" { return "Port" }
    if yname == "version" { return "Version" }
    if yname == "security-level" { return "SecurityLevel" }
    if yname == "basic-trap-types" { return "BasicTrapTypes" }
    if yname == "advanced-trap-types1" { return "AdvancedTrapTypes1" }
    if yname == "advanced-trap-types2" { return "AdvancedTrapTypes2" }
    return ""
}

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetSegmentPath() string {
    return "inform-encrypted-user-community" + "[community-name='" + fmt.Sprintf("%v", informEncryptedUserCommunity.CommunityName) + "']"
}

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-name"] = informEncryptedUserCommunity.CommunityName
    leafs["port"] = informEncryptedUserCommunity.Port
    leafs["version"] = informEncryptedUserCommunity.Version
    leafs["security-level"] = informEncryptedUserCommunity.SecurityLevel
    leafs["basic-trap-types"] = informEncryptedUserCommunity.BasicTrapTypes
    leafs["advanced-trap-types1"] = informEncryptedUserCommunity.AdvancedTrapTypes1
    leafs["advanced-trap-types2"] = informEncryptedUserCommunity.AdvancedTrapTypes2
    return leafs
}

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetYangName() string { return "inform-encrypted-user-community" }

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) SetParent(parent types.Entity) { informEncryptedUserCommunity.parent = parent }

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetParent() types.Entity { return informEncryptedUserCommunity.parent }

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetParentYangName() string { return "inform-encrypted-user-communities" }

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities
// Container class for defining communities for a
// trap host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unencrpted Community name associated with a trap host. The type is slice of
    // Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity.
    DefaultUserCommunity []Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity
}

func (defaultUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities) GetFilter() yfilter.YFilter { return defaultUserCommunities.YFilter }

func (defaultUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities) SetFilter(yf yfilter.YFilter) { defaultUserCommunities.YFilter = yf }

func (defaultUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities) GetGoName(yname string) string {
    if yname == "default-user-community" { return "DefaultUserCommunity" }
    return ""
}

func (defaultUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities) GetSegmentPath() string {
    return "default-user-communities"
}

func (defaultUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-user-community" {
        for _, c := range defaultUserCommunities.DefaultUserCommunity {
            if defaultUserCommunities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity{}
        defaultUserCommunities.DefaultUserCommunity = append(defaultUserCommunities.DefaultUserCommunity, child)
        return &defaultUserCommunities.DefaultUserCommunity[len(defaultUserCommunities.DefaultUserCommunity)-1]
    }
    return nil
}

func (defaultUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range defaultUserCommunities.DefaultUserCommunity {
        children[defaultUserCommunities.DefaultUserCommunity[i].GetSegmentPath()] = &defaultUserCommunities.DefaultUserCommunity[i]
    }
    return children
}

func (defaultUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (defaultUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities) GetBundleName() string { return "cisco_ios_xr" }

func (defaultUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities) GetYangName() string { return "default-user-communities" }

func (defaultUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities) SetParent(parent types.Entity) { defaultUserCommunities.parent = parent }

func (defaultUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities) GetParent() types.Entity { return defaultUserCommunities.parent }

func (defaultUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities) GetParentYangName() string { return "trap-host" }

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity
// Unencrpted Community name associated with a
// trap host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SNMPv1/v2c community string or SNMPv3 user. The
    // type is string with length: 1..128.
    CommunityName interface{}

    // UDP port number. The type is interface{} with range: 1..65535.
    Port interface{}

    // SNMP Version to be used v1/v2c/v3. The type is string. This attribute is
    // mandatory.
    Version interface{}

    // Security level to be used noauth/auth/priv. The type is SnmpSecurityModel.
    SecurityLevel interface{}

    // Number to signify the feature traps that needs to be setBasicTrapTypes is
    // used for all traps except copy-completeSet this value to an integer
    // corresponding to the trapBGP 8192, CONFIG 4096,SYSLOG 131072,SNMP_TRAP
    // 1COPY_COMPLETE_TRAP 64To provide a combination of trap Add the respective
    // numbersValue must be set to 0 for all traps. The type is interface{} with
    // range: -2147483648..2147483647. This attribute is mandatory.
    BasicTrapTypes interface{}

    // Number to signify the feature traps that needs to be setUse this for
    // providing copy-complete trapValue must be set to 0 if not used. The type is
    // interface{} with range: -2147483648..2147483647. This attribute is
    // mandatory.
    AdvancedTrapTypes1 interface{}

    // Number to signify the feature traps that needs to be setvalue should always
    // to set as 0. The type is interface{} with range: -2147483648..2147483647.
    // This attribute is mandatory.
    AdvancedTrapTypes2 interface{}
}

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetFilter() yfilter.YFilter { return defaultUserCommunity.YFilter }

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) SetFilter(yf yfilter.YFilter) { defaultUserCommunity.YFilter = yf }

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetGoName(yname string) string {
    if yname == "community-name" { return "CommunityName" }
    if yname == "port" { return "Port" }
    if yname == "version" { return "Version" }
    if yname == "security-level" { return "SecurityLevel" }
    if yname == "basic-trap-types" { return "BasicTrapTypes" }
    if yname == "advanced-trap-types1" { return "AdvancedTrapTypes1" }
    if yname == "advanced-trap-types2" { return "AdvancedTrapTypes2" }
    return ""
}

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetSegmentPath() string {
    return "default-user-community" + "[community-name='" + fmt.Sprintf("%v", defaultUserCommunity.CommunityName) + "']"
}

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-name"] = defaultUserCommunity.CommunityName
    leafs["port"] = defaultUserCommunity.Port
    leafs["version"] = defaultUserCommunity.Version
    leafs["security-level"] = defaultUserCommunity.SecurityLevel
    leafs["basic-trap-types"] = defaultUserCommunity.BasicTrapTypes
    leafs["advanced-trap-types1"] = defaultUserCommunity.AdvancedTrapTypes1
    leafs["advanced-trap-types2"] = defaultUserCommunity.AdvancedTrapTypes2
    return leafs
}

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetYangName() string { return "default-user-community" }

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) SetParent(parent types.Entity) { defaultUserCommunity.parent = parent }

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetParent() types.Entity { return defaultUserCommunity.parent }

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetParentYangName() string { return "default-user-communities" }

// Snmp_Vrfs_Vrf_Contexts
// List of Context Names
type Snmp_Vrfs_Vrf_Contexts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Context Name. The type is slice of Snmp_Vrfs_Vrf_Contexts_Context.
    Context []Snmp_Vrfs_Vrf_Contexts_Context
}

func (contexts *Snmp_Vrfs_Vrf_Contexts) GetFilter() yfilter.YFilter { return contexts.YFilter }

func (contexts *Snmp_Vrfs_Vrf_Contexts) SetFilter(yf yfilter.YFilter) { contexts.YFilter = yf }

func (contexts *Snmp_Vrfs_Vrf_Contexts) GetGoName(yname string) string {
    if yname == "context" { return "Context" }
    return ""
}

func (contexts *Snmp_Vrfs_Vrf_Contexts) GetSegmentPath() string {
    return "contexts"
}

func (contexts *Snmp_Vrfs_Vrf_Contexts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "context" {
        for _, c := range contexts.Context {
            if contexts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Vrfs_Vrf_Contexts_Context{}
        contexts.Context = append(contexts.Context, child)
        return &contexts.Context[len(contexts.Context)-1]
    }
    return nil
}

func (contexts *Snmp_Vrfs_Vrf_Contexts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range contexts.Context {
        children[contexts.Context[i].GetSegmentPath()] = &contexts.Context[i]
    }
    return children
}

func (contexts *Snmp_Vrfs_Vrf_Contexts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (contexts *Snmp_Vrfs_Vrf_Contexts) GetBundleName() string { return "cisco_ios_xr" }

func (contexts *Snmp_Vrfs_Vrf_Contexts) GetYangName() string { return "contexts" }

func (contexts *Snmp_Vrfs_Vrf_Contexts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contexts *Snmp_Vrfs_Vrf_Contexts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contexts *Snmp_Vrfs_Vrf_Contexts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contexts *Snmp_Vrfs_Vrf_Contexts) SetParent(parent types.Entity) { contexts.parent = parent }

func (contexts *Snmp_Vrfs_Vrf_Contexts) GetParent() types.Entity { return contexts.parent }

func (contexts *Snmp_Vrfs_Vrf_Contexts) GetParentYangName() string { return "vrf" }

// Snmp_Vrfs_Vrf_Contexts_Context
// Context Name
type Snmp_Vrfs_Vrf_Contexts_Context struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Context Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ContextName interface{}
}

func (context *Snmp_Vrfs_Vrf_Contexts_Context) GetFilter() yfilter.YFilter { return context.YFilter }

func (context *Snmp_Vrfs_Vrf_Contexts_Context) SetFilter(yf yfilter.YFilter) { context.YFilter = yf }

func (context *Snmp_Vrfs_Vrf_Contexts_Context) GetGoName(yname string) string {
    if yname == "context-name" { return "ContextName" }
    return ""
}

func (context *Snmp_Vrfs_Vrf_Contexts_Context) GetSegmentPath() string {
    return "context" + "[context-name='" + fmt.Sprintf("%v", context.ContextName) + "']"
}

func (context *Snmp_Vrfs_Vrf_Contexts_Context) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (context *Snmp_Vrfs_Vrf_Contexts_Context) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (context *Snmp_Vrfs_Vrf_Contexts_Context) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["context-name"] = context.ContextName
    return leafs
}

func (context *Snmp_Vrfs_Vrf_Contexts_Context) GetBundleName() string { return "cisco_ios_xr" }

func (context *Snmp_Vrfs_Vrf_Contexts_Context) GetYangName() string { return "context" }

func (context *Snmp_Vrfs_Vrf_Contexts_Context) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (context *Snmp_Vrfs_Vrf_Contexts_Context) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (context *Snmp_Vrfs_Vrf_Contexts_Context) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (context *Snmp_Vrfs_Vrf_Contexts_Context) SetParent(parent types.Entity) { context.parent = parent }

func (context *Snmp_Vrfs_Vrf_Contexts_Context) GetParent() types.Entity { return context.parent }

func (context *Snmp_Vrfs_Vrf_Contexts_Context) GetParentYangName() string { return "contexts" }

// Snmp_Vrfs_Vrf_ContextMappings
// List of context names
type Snmp_Vrfs_Vrf_ContextMappings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Context mapping name. The type is slice of
    // Snmp_Vrfs_Vrf_ContextMappings_ContextMapping.
    ContextMapping []Snmp_Vrfs_Vrf_ContextMappings_ContextMapping
}

func (contextMappings *Snmp_Vrfs_Vrf_ContextMappings) GetFilter() yfilter.YFilter { return contextMappings.YFilter }

func (contextMappings *Snmp_Vrfs_Vrf_ContextMappings) SetFilter(yf yfilter.YFilter) { contextMappings.YFilter = yf }

func (contextMappings *Snmp_Vrfs_Vrf_ContextMappings) GetGoName(yname string) string {
    if yname == "context-mapping" { return "ContextMapping" }
    return ""
}

func (contextMappings *Snmp_Vrfs_Vrf_ContextMappings) GetSegmentPath() string {
    return "context-mappings"
}

func (contextMappings *Snmp_Vrfs_Vrf_ContextMappings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "context-mapping" {
        for _, c := range contextMappings.ContextMapping {
            if contextMappings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Vrfs_Vrf_ContextMappings_ContextMapping{}
        contextMappings.ContextMapping = append(contextMappings.ContextMapping, child)
        return &contextMappings.ContextMapping[len(contextMappings.ContextMapping)-1]
    }
    return nil
}

func (contextMappings *Snmp_Vrfs_Vrf_ContextMappings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range contextMappings.ContextMapping {
        children[contextMappings.ContextMapping[i].GetSegmentPath()] = &contextMappings.ContextMapping[i]
    }
    return children
}

func (contextMappings *Snmp_Vrfs_Vrf_ContextMappings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (contextMappings *Snmp_Vrfs_Vrf_ContextMappings) GetBundleName() string { return "cisco_ios_xr" }

func (contextMappings *Snmp_Vrfs_Vrf_ContextMappings) GetYangName() string { return "context-mappings" }

func (contextMappings *Snmp_Vrfs_Vrf_ContextMappings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contextMappings *Snmp_Vrfs_Vrf_ContextMappings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contextMappings *Snmp_Vrfs_Vrf_ContextMappings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contextMappings *Snmp_Vrfs_Vrf_ContextMappings) SetParent(parent types.Entity) { contextMappings.parent = parent }

func (contextMappings *Snmp_Vrfs_Vrf_ContextMappings) GetParent() types.Entity { return contextMappings.parent }

func (contextMappings *Snmp_Vrfs_Vrf_ContextMappings) GetParentYangName() string { return "vrf" }

// Snmp_Vrfs_Vrf_ContextMappings_ContextMapping
// Context mapping name
type Snmp_Vrfs_Vrf_ContextMappings_ContextMapping struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Context mapping name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ContextMappingName interface{}

    // SNMP context feature type. The type is SnmpContext.
    Context interface{}

    // OSPF protocol instance. The type is string.
    InstanceName interface{}

    // VRF name associated with the context. The type is string.
    VrfName interface{}

    // Topology name associated with the context. The type is string.
    TopologyName interface{}
}

func (contextMapping *Snmp_Vrfs_Vrf_ContextMappings_ContextMapping) GetFilter() yfilter.YFilter { return contextMapping.YFilter }

func (contextMapping *Snmp_Vrfs_Vrf_ContextMappings_ContextMapping) SetFilter(yf yfilter.YFilter) { contextMapping.YFilter = yf }

func (contextMapping *Snmp_Vrfs_Vrf_ContextMappings_ContextMapping) GetGoName(yname string) string {
    if yname == "context-mapping-name" { return "ContextMappingName" }
    if yname == "context" { return "Context" }
    if yname == "instance-name" { return "InstanceName" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "topology-name" { return "TopologyName" }
    return ""
}

func (contextMapping *Snmp_Vrfs_Vrf_ContextMappings_ContextMapping) GetSegmentPath() string {
    return "context-mapping" + "[context-mapping-name='" + fmt.Sprintf("%v", contextMapping.ContextMappingName) + "']"
}

func (contextMapping *Snmp_Vrfs_Vrf_ContextMappings_ContextMapping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (contextMapping *Snmp_Vrfs_Vrf_ContextMappings_ContextMapping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (contextMapping *Snmp_Vrfs_Vrf_ContextMappings_ContextMapping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["context-mapping-name"] = contextMapping.ContextMappingName
    leafs["context"] = contextMapping.Context
    leafs["instance-name"] = contextMapping.InstanceName
    leafs["vrf-name"] = contextMapping.VrfName
    leafs["topology-name"] = contextMapping.TopologyName
    return leafs
}

func (contextMapping *Snmp_Vrfs_Vrf_ContextMappings_ContextMapping) GetBundleName() string { return "cisco_ios_xr" }

func (contextMapping *Snmp_Vrfs_Vrf_ContextMappings_ContextMapping) GetYangName() string { return "context-mapping" }

func (contextMapping *Snmp_Vrfs_Vrf_ContextMappings_ContextMapping) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contextMapping *Snmp_Vrfs_Vrf_ContextMappings_ContextMapping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contextMapping *Snmp_Vrfs_Vrf_ContextMappings_ContextMapping) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contextMapping *Snmp_Vrfs_Vrf_ContextMappings_ContextMapping) SetParent(parent types.Entity) { contextMapping.parent = parent }

func (contextMapping *Snmp_Vrfs_Vrf_ContextMappings_ContextMapping) GetParent() types.Entity { return contextMapping.parent }

func (contextMapping *Snmp_Vrfs_Vrf_ContextMappings_ContextMapping) GetParentYangName() string { return "context-mappings" }

// Snmp_Groups
// Define a User Security Model group
type Snmp_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the group. The type is slice of Snmp_Groups_Group.
    Group []Snmp_Groups_Group
}

func (groups *Snmp_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *Snmp_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *Snmp_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *Snmp_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *Snmp_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *Snmp_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *Snmp_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *Snmp_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *Snmp_Groups) GetYangName() string { return "groups" }

func (groups *Snmp_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *Snmp_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *Snmp_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *Snmp_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *Snmp_Groups) GetParent() types.Entity { return groups.parent }

func (groups *Snmp_Groups) GetParentYangName() string { return "snmp" }

// Snmp_Groups_Group
// Name of the group
type Snmp_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the group. The type is string with length:
    // 1..128.
    Name interface{}

    // snmp version. The type is GroupSnmpVersion. This attribute is mandatory.
    SnmpVersion interface{}

    // security model like auth/noAuth/Priv applicable for v3. The type is
    // SnmpSecurityModel.
    SecurityModel interface{}

    // notify view name. The type is string.
    NotifyView interface{}

    // read view name. The type is string.
    ReadView interface{}

    // write view name. The type is string.
    WriteView interface{}

    // Access-list type. The type is Snmpacl.
    V4AclType interface{}

    // Ipv4 Access-list name. The type is string.
    V4AccessList interface{}

    // Access-list type. The type is Snmpacl.
    V6AclType interface{}

    // Ipv6 Access-list name. The type is string.
    V6AccessList interface{}

    // Context name. The type is string.
    ContextName interface{}
}

func (group *Snmp_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *Snmp_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *Snmp_Groups_Group) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "snmp-version" { return "SnmpVersion" }
    if yname == "security-model" { return "SecurityModel" }
    if yname == "notify-view" { return "NotifyView" }
    if yname == "read-view" { return "ReadView" }
    if yname == "write-view" { return "WriteView" }
    if yname == "v4acl-type" { return "V4AclType" }
    if yname == "v4-access-list" { return "V4AccessList" }
    if yname == "v6acl-type" { return "V6AclType" }
    if yname == "v6-access-list" { return "V6AccessList" }
    if yname == "context-name" { return "ContextName" }
    return ""
}

func (group *Snmp_Groups_Group) GetSegmentPath() string {
    return "group" + "[name='" + fmt.Sprintf("%v", group.Name) + "']"
}

func (group *Snmp_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (group *Snmp_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (group *Snmp_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = group.Name
    leafs["snmp-version"] = group.SnmpVersion
    leafs["security-model"] = group.SecurityModel
    leafs["notify-view"] = group.NotifyView
    leafs["read-view"] = group.ReadView
    leafs["write-view"] = group.WriteView
    leafs["v4acl-type"] = group.V4AclType
    leafs["v4-access-list"] = group.V4AccessList
    leafs["v6acl-type"] = group.V6AclType
    leafs["v6-access-list"] = group.V6AccessList
    leafs["context-name"] = group.ContextName
    return leafs
}

func (group *Snmp_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *Snmp_Groups_Group) GetYangName() string { return "group" }

func (group *Snmp_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *Snmp_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *Snmp_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *Snmp_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *Snmp_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *Snmp_Groups_Group) GetParentYangName() string { return "groups" }

// Snmp_TrapHosts
// Specify hosts to receive SNMP notifications
type Snmp_TrapHosts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify hosts to receive SNMP notifications. The type is slice of
    // Snmp_TrapHosts_TrapHost.
    TrapHost []Snmp_TrapHosts_TrapHost
}

func (trapHosts *Snmp_TrapHosts) GetFilter() yfilter.YFilter { return trapHosts.YFilter }

func (trapHosts *Snmp_TrapHosts) SetFilter(yf yfilter.YFilter) { trapHosts.YFilter = yf }

func (trapHosts *Snmp_TrapHosts) GetGoName(yname string) string {
    if yname == "trap-host" { return "TrapHost" }
    return ""
}

func (trapHosts *Snmp_TrapHosts) GetSegmentPath() string {
    return "trap-hosts"
}

func (trapHosts *Snmp_TrapHosts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "trap-host" {
        for _, c := range trapHosts.TrapHost {
            if trapHosts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_TrapHosts_TrapHost{}
        trapHosts.TrapHost = append(trapHosts.TrapHost, child)
        return &trapHosts.TrapHost[len(trapHosts.TrapHost)-1]
    }
    return nil
}

func (trapHosts *Snmp_TrapHosts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trapHosts.TrapHost {
        children[trapHosts.TrapHost[i].GetSegmentPath()] = &trapHosts.TrapHost[i]
    }
    return children
}

func (trapHosts *Snmp_TrapHosts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trapHosts *Snmp_TrapHosts) GetBundleName() string { return "cisco_ios_xr" }

func (trapHosts *Snmp_TrapHosts) GetYangName() string { return "trap-hosts" }

func (trapHosts *Snmp_TrapHosts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapHosts *Snmp_TrapHosts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapHosts *Snmp_TrapHosts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapHosts *Snmp_TrapHosts) SetParent(parent types.Entity) { trapHosts.parent = parent }

func (trapHosts *Snmp_TrapHosts) GetParent() types.Entity { return trapHosts.parent }

func (trapHosts *Snmp_TrapHosts) GetParentYangName() string { return "snmp" }

// Snmp_TrapHosts_TrapHost
// Specify hosts to receive SNMP notifications
type Snmp_TrapHosts_TrapHost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of SNMP notification host. The type is
    // one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Container class for defining Clear/encrypt communities for a trap host.
    EncryptedUserCommunities Snmp_TrapHosts_TrapHost_EncryptedUserCommunities

    // Container class for defining notification type for a Inform host.
    InformHost Snmp_TrapHosts_TrapHost_InformHost

    // Container class for defining communities for a trap host.
    DefaultUserCommunities Snmp_TrapHosts_TrapHost_DefaultUserCommunities
}

func (trapHost *Snmp_TrapHosts_TrapHost) GetFilter() yfilter.YFilter { return trapHost.YFilter }

func (trapHost *Snmp_TrapHosts_TrapHost) SetFilter(yf yfilter.YFilter) { trapHost.YFilter = yf }

func (trapHost *Snmp_TrapHosts_TrapHost) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "encrypted-user-communities" { return "EncryptedUserCommunities" }
    if yname == "inform-host" { return "InformHost" }
    if yname == "default-user-communities" { return "DefaultUserCommunities" }
    return ""
}

func (trapHost *Snmp_TrapHosts_TrapHost) GetSegmentPath() string {
    return "trap-host" + "[ip-address='" + fmt.Sprintf("%v", trapHost.IpAddress) + "']"
}

func (trapHost *Snmp_TrapHosts_TrapHost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "encrypted-user-communities" {
        return &trapHost.EncryptedUserCommunities
    }
    if childYangName == "inform-host" {
        return &trapHost.InformHost
    }
    if childYangName == "default-user-communities" {
        return &trapHost.DefaultUserCommunities
    }
    return nil
}

func (trapHost *Snmp_TrapHosts_TrapHost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["encrypted-user-communities"] = &trapHost.EncryptedUserCommunities
    children["inform-host"] = &trapHost.InformHost
    children["default-user-communities"] = &trapHost.DefaultUserCommunities
    return children
}

func (trapHost *Snmp_TrapHosts_TrapHost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = trapHost.IpAddress
    return leafs
}

func (trapHost *Snmp_TrapHosts_TrapHost) GetBundleName() string { return "cisco_ios_xr" }

func (trapHost *Snmp_TrapHosts_TrapHost) GetYangName() string { return "trap-host" }

func (trapHost *Snmp_TrapHosts_TrapHost) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trapHost *Snmp_TrapHosts_TrapHost) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trapHost *Snmp_TrapHosts_TrapHost) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trapHost *Snmp_TrapHosts_TrapHost) SetParent(parent types.Entity) { trapHost.parent = parent }

func (trapHost *Snmp_TrapHosts_TrapHost) GetParent() types.Entity { return trapHost.parent }

func (trapHost *Snmp_TrapHosts_TrapHost) GetParentYangName() string { return "trap-hosts" }

// Snmp_TrapHosts_TrapHost_EncryptedUserCommunities
// Container class for defining Clear/encrypt
// communities for a trap host
type Snmp_TrapHosts_TrapHost_EncryptedUserCommunities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear/Encrypt Community name associated with a trap host. The type is slice
    // of Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity.
    EncryptedUserCommunity []Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity
}

func (encryptedUserCommunities *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities) GetFilter() yfilter.YFilter { return encryptedUserCommunities.YFilter }

func (encryptedUserCommunities *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities) SetFilter(yf yfilter.YFilter) { encryptedUserCommunities.YFilter = yf }

func (encryptedUserCommunities *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities) GetGoName(yname string) string {
    if yname == "encrypted-user-community" { return "EncryptedUserCommunity" }
    return ""
}

func (encryptedUserCommunities *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities) GetSegmentPath() string {
    return "encrypted-user-communities"
}

func (encryptedUserCommunities *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "encrypted-user-community" {
        for _, c := range encryptedUserCommunities.EncryptedUserCommunity {
            if encryptedUserCommunities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity{}
        encryptedUserCommunities.EncryptedUserCommunity = append(encryptedUserCommunities.EncryptedUserCommunity, child)
        return &encryptedUserCommunities.EncryptedUserCommunity[len(encryptedUserCommunities.EncryptedUserCommunity)-1]
    }
    return nil
}

func (encryptedUserCommunities *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range encryptedUserCommunities.EncryptedUserCommunity {
        children[encryptedUserCommunities.EncryptedUserCommunity[i].GetSegmentPath()] = &encryptedUserCommunities.EncryptedUserCommunity[i]
    }
    return children
}

func (encryptedUserCommunities *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (encryptedUserCommunities *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities) GetBundleName() string { return "cisco_ios_xr" }

func (encryptedUserCommunities *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities) GetYangName() string { return "encrypted-user-communities" }

func (encryptedUserCommunities *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (encryptedUserCommunities *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (encryptedUserCommunities *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (encryptedUserCommunities *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities) SetParent(parent types.Entity) { encryptedUserCommunities.parent = parent }

func (encryptedUserCommunities *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities) GetParent() types.Entity { return encryptedUserCommunities.parent }

func (encryptedUserCommunities *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities) GetParentYangName() string { return "trap-host" }

// Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity
// Clear/Encrypt Community name associated with
// a trap host
type Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SNMPv1/v2c community string or SNMPv3 user. The
    // type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    CommunityName interface{}

    // UDP port number. The type is interface{} with range: 1..65535.
    Port interface{}

    // SNMP Version to be used v1/v2c/v3. The type is string. This attribute is
    // mandatory.
    Version interface{}

    // Security level to be used noauth/auth/priv. The type is SnmpSecurityModel.
    SecurityLevel interface{}

    // Number to signify the feature traps that needs to be setBasicTrapTypes is
    // used for all traps except copy-completeSet this value to an integer
    // corresponding to the trapBGP 8192, CONFIG 4096,SYSLOG 131072,SNMP_TRAP
    // 1COPY_COMPLETE_TRAP 64To provide a combination of trap Add the respective
    // numbersValue must be set to 0 for all traps. The type is interface{} with
    // range: -2147483648..2147483647. This attribute is mandatory.
    BasicTrapTypes interface{}

    // Number to signify the feature traps that needs to be setUse this for
    // providing copy-complete trapValue must be set to 0 if not used. The type is
    // interface{} with range: -2147483648..2147483647. This attribute is
    // mandatory.
    AdvancedTrapTypes1 interface{}

    // Number to signify the feature traps that needs to be setvalue should always
    // to set as 0. The type is interface{} with range: -2147483648..2147483647.
    // This attribute is mandatory.
    AdvancedTrapTypes2 interface{}
}

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetFilter() yfilter.YFilter { return encryptedUserCommunity.YFilter }

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) SetFilter(yf yfilter.YFilter) { encryptedUserCommunity.YFilter = yf }

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetGoName(yname string) string {
    if yname == "community-name" { return "CommunityName" }
    if yname == "port" { return "Port" }
    if yname == "version" { return "Version" }
    if yname == "security-level" { return "SecurityLevel" }
    if yname == "basic-trap-types" { return "BasicTrapTypes" }
    if yname == "advanced-trap-types1" { return "AdvancedTrapTypes1" }
    if yname == "advanced-trap-types2" { return "AdvancedTrapTypes2" }
    return ""
}

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetSegmentPath() string {
    return "encrypted-user-community" + "[community-name='" + fmt.Sprintf("%v", encryptedUserCommunity.CommunityName) + "']"
}

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-name"] = encryptedUserCommunity.CommunityName
    leafs["port"] = encryptedUserCommunity.Port
    leafs["version"] = encryptedUserCommunity.Version
    leafs["security-level"] = encryptedUserCommunity.SecurityLevel
    leafs["basic-trap-types"] = encryptedUserCommunity.BasicTrapTypes
    leafs["advanced-trap-types1"] = encryptedUserCommunity.AdvancedTrapTypes1
    leafs["advanced-trap-types2"] = encryptedUserCommunity.AdvancedTrapTypes2
    return leafs
}

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetYangName() string { return "encrypted-user-community" }

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) SetParent(parent types.Entity) { encryptedUserCommunity.parent = parent }

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetParent() types.Entity { return encryptedUserCommunity.parent }

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetParentYangName() string { return "encrypted-user-communities" }

// Snmp_TrapHosts_TrapHost_InformHost
// Container class for defining notification type
// for a Inform host
type Snmp_TrapHosts_TrapHost_InformHost struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Container class for defining communities for a inform host.
    InformUserCommunities Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities

    // Container class for defining Clear/encrypt communities for a inform host.
    InformEncryptedUserCommunities Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities
}

func (informHost *Snmp_TrapHosts_TrapHost_InformHost) GetFilter() yfilter.YFilter { return informHost.YFilter }

func (informHost *Snmp_TrapHosts_TrapHost_InformHost) SetFilter(yf yfilter.YFilter) { informHost.YFilter = yf }

func (informHost *Snmp_TrapHosts_TrapHost_InformHost) GetGoName(yname string) string {
    if yname == "inform-user-communities" { return "InformUserCommunities" }
    if yname == "inform-encrypted-user-communities" { return "InformEncryptedUserCommunities" }
    return ""
}

func (informHost *Snmp_TrapHosts_TrapHost_InformHost) GetSegmentPath() string {
    return "inform-host"
}

func (informHost *Snmp_TrapHosts_TrapHost_InformHost) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inform-user-communities" {
        return &informHost.InformUserCommunities
    }
    if childYangName == "inform-encrypted-user-communities" {
        return &informHost.InformEncryptedUserCommunities
    }
    return nil
}

func (informHost *Snmp_TrapHosts_TrapHost_InformHost) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["inform-user-communities"] = &informHost.InformUserCommunities
    children["inform-encrypted-user-communities"] = &informHost.InformEncryptedUserCommunities
    return children
}

func (informHost *Snmp_TrapHosts_TrapHost_InformHost) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (informHost *Snmp_TrapHosts_TrapHost_InformHost) GetBundleName() string { return "cisco_ios_xr" }

func (informHost *Snmp_TrapHosts_TrapHost_InformHost) GetYangName() string { return "inform-host" }

func (informHost *Snmp_TrapHosts_TrapHost_InformHost) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (informHost *Snmp_TrapHosts_TrapHost_InformHost) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (informHost *Snmp_TrapHosts_TrapHost_InformHost) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (informHost *Snmp_TrapHosts_TrapHost_InformHost) SetParent(parent types.Entity) { informHost.parent = parent }

func (informHost *Snmp_TrapHosts_TrapHost_InformHost) GetParent() types.Entity { return informHost.parent }

func (informHost *Snmp_TrapHosts_TrapHost_InformHost) GetParentYangName() string { return "trap-host" }

// Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities
// Container class for defining communities for
// a inform host
type Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unencrpted Community name associated with a inform host. The type is slice
    // of
    // Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity.
    InformUserCommunity []Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity
}

func (informUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetFilter() yfilter.YFilter { return informUserCommunities.YFilter }

func (informUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities) SetFilter(yf yfilter.YFilter) { informUserCommunities.YFilter = yf }

func (informUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetGoName(yname string) string {
    if yname == "inform-user-community" { return "InformUserCommunity" }
    return ""
}

func (informUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetSegmentPath() string {
    return "inform-user-communities"
}

func (informUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inform-user-community" {
        for _, c := range informUserCommunities.InformUserCommunity {
            if informUserCommunities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity{}
        informUserCommunities.InformUserCommunity = append(informUserCommunities.InformUserCommunity, child)
        return &informUserCommunities.InformUserCommunity[len(informUserCommunities.InformUserCommunity)-1]
    }
    return nil
}

func (informUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range informUserCommunities.InformUserCommunity {
        children[informUserCommunities.InformUserCommunity[i].GetSegmentPath()] = &informUserCommunities.InformUserCommunity[i]
    }
    return children
}

func (informUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (informUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetBundleName() string { return "cisco_ios_xr" }

func (informUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetYangName() string { return "inform-user-communities" }

func (informUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (informUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (informUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (informUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities) SetParent(parent types.Entity) { informUserCommunities.parent = parent }

func (informUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetParent() types.Entity { return informUserCommunities.parent }

func (informUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetParentYangName() string { return "inform-host" }

// Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity
// Unencrpted Community name associated with a
// inform host
type Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SNMPv2c community string or SNMPv3 user. The type
    // is string with length: 1..128.
    CommunityName interface{}

    // UDP port number. The type is interface{} with range: 1..65535.
    Port interface{}

    // SNMP Version to be used v2c/v3. The type is string. This attribute is
    // mandatory.
    Version interface{}

    // Security level to be used noauth/auth/priv. The type is SnmpSecurityModel.
    SecurityLevel interface{}

    // Number to signify the feature traps that needs to be setBasicTrapTypes is
    // used for all traps except copy-completeSet this value to an integer
    // corresponding to the trapBGP 8192, CONFIG 4096,SYSLOG 131072 ,SNMP_TRAP
    // 1COPY_COMPLETE_TRAP 64To provide a combination of trap Add the respective
    // numbersValue must be set to 0 for all traps. The type is interface{} with
    // range: -2147483648..2147483647. This attribute is mandatory.
    BasicTrapTypes interface{}

    // Number to signify the feature traps that needs to be setUse this for
    // providing copy-complete trapValue must be set to 0 if not used. The type is
    // interface{} with range: -2147483648..2147483647. This attribute is
    // mandatory.
    AdvancedTrapTypes1 interface{}

    // Number to signify the feature traps that needs to be setvalue should always
    // to set as 0. The type is interface{} with range: -2147483648..2147483647.
    // This attribute is mandatory.
    AdvancedTrapTypes2 interface{}
}

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetFilter() yfilter.YFilter { return informUserCommunity.YFilter }

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) SetFilter(yf yfilter.YFilter) { informUserCommunity.YFilter = yf }

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetGoName(yname string) string {
    if yname == "community-name" { return "CommunityName" }
    if yname == "port" { return "Port" }
    if yname == "version" { return "Version" }
    if yname == "security-level" { return "SecurityLevel" }
    if yname == "basic-trap-types" { return "BasicTrapTypes" }
    if yname == "advanced-trap-types1" { return "AdvancedTrapTypes1" }
    if yname == "advanced-trap-types2" { return "AdvancedTrapTypes2" }
    return ""
}

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetSegmentPath() string {
    return "inform-user-community" + "[community-name='" + fmt.Sprintf("%v", informUserCommunity.CommunityName) + "']"
}

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-name"] = informUserCommunity.CommunityName
    leafs["port"] = informUserCommunity.Port
    leafs["version"] = informUserCommunity.Version
    leafs["security-level"] = informUserCommunity.SecurityLevel
    leafs["basic-trap-types"] = informUserCommunity.BasicTrapTypes
    leafs["advanced-trap-types1"] = informUserCommunity.AdvancedTrapTypes1
    leafs["advanced-trap-types2"] = informUserCommunity.AdvancedTrapTypes2
    return leafs
}

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetYangName() string { return "inform-user-community" }

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) SetParent(parent types.Entity) { informUserCommunity.parent = parent }

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetParent() types.Entity { return informUserCommunity.parent }

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetParentYangName() string { return "inform-user-communities" }

// Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities
// Container class for defining Clear/encrypt
// communities for a inform host
type Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Clear/Encrypt Community name associated with a inform host. The type is
    // slice of
    // Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity.
    InformEncryptedUserCommunity []Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity
}

func (informEncryptedUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetFilter() yfilter.YFilter { return informEncryptedUserCommunities.YFilter }

func (informEncryptedUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) SetFilter(yf yfilter.YFilter) { informEncryptedUserCommunities.YFilter = yf }

func (informEncryptedUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetGoName(yname string) string {
    if yname == "inform-encrypted-user-community" { return "InformEncryptedUserCommunity" }
    return ""
}

func (informEncryptedUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetSegmentPath() string {
    return "inform-encrypted-user-communities"
}

func (informEncryptedUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "inform-encrypted-user-community" {
        for _, c := range informEncryptedUserCommunities.InformEncryptedUserCommunity {
            if informEncryptedUserCommunities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity{}
        informEncryptedUserCommunities.InformEncryptedUserCommunity = append(informEncryptedUserCommunities.InformEncryptedUserCommunity, child)
        return &informEncryptedUserCommunities.InformEncryptedUserCommunity[len(informEncryptedUserCommunities.InformEncryptedUserCommunity)-1]
    }
    return nil
}

func (informEncryptedUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range informEncryptedUserCommunities.InformEncryptedUserCommunity {
        children[informEncryptedUserCommunities.InformEncryptedUserCommunity[i].GetSegmentPath()] = &informEncryptedUserCommunities.InformEncryptedUserCommunity[i]
    }
    return children
}

func (informEncryptedUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (informEncryptedUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetBundleName() string { return "cisco_ios_xr" }

func (informEncryptedUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetYangName() string { return "inform-encrypted-user-communities" }

func (informEncryptedUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (informEncryptedUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (informEncryptedUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (informEncryptedUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) SetParent(parent types.Entity) { informEncryptedUserCommunities.parent = parent }

func (informEncryptedUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetParent() types.Entity { return informEncryptedUserCommunities.parent }

func (informEncryptedUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetParentYangName() string { return "inform-host" }

// Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity
// Clear/Encrypt Community name associated with
// a inform host
type Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SNMPv2c community string or SNMPv3 user. The type
    // is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    CommunityName interface{}

    // UDP port number. The type is interface{} with range: 1..65535.
    Port interface{}

    // SNMP Version to be used v2c/v3. The type is string. This attribute is
    // mandatory.
    Version interface{}

    // Security level to be used noauth/auth/priv. The type is SnmpSecurityModel.
    SecurityLevel interface{}

    // Number to signify the feature traps that needs to be setBasicTrapTypes is
    // used for all traps except copy-completeSet this value to an integer
    // corresponding to the trapBGP 8192, CONFIG 4096,SYSLOG 131072 ,SNMP_TRAP
    // 1COPY_COMPLETE_TRAP 64To provide a combination of trap Add the respective
    // numbersValue must be set to 0 for all traps. The type is interface{} with
    // range: -2147483648..2147483647. This attribute is mandatory.
    BasicTrapTypes interface{}

    // Number to signify the feature traps that needs to be setUse this for
    // providing copy-complete trapValue must be set to 0 if not used. The type is
    // interface{} with range: -2147483648..2147483647. This attribute is
    // mandatory.
    AdvancedTrapTypes1 interface{}

    // Number to signify the feature traps that needs to be setvalue should always
    // to set as 0. The type is interface{} with range: -2147483648..2147483647.
    // This attribute is mandatory.
    AdvancedTrapTypes2 interface{}
}

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetFilter() yfilter.YFilter { return informEncryptedUserCommunity.YFilter }

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) SetFilter(yf yfilter.YFilter) { informEncryptedUserCommunity.YFilter = yf }

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetGoName(yname string) string {
    if yname == "community-name" { return "CommunityName" }
    if yname == "port" { return "Port" }
    if yname == "version" { return "Version" }
    if yname == "security-level" { return "SecurityLevel" }
    if yname == "basic-trap-types" { return "BasicTrapTypes" }
    if yname == "advanced-trap-types1" { return "AdvancedTrapTypes1" }
    if yname == "advanced-trap-types2" { return "AdvancedTrapTypes2" }
    return ""
}

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetSegmentPath() string {
    return "inform-encrypted-user-community" + "[community-name='" + fmt.Sprintf("%v", informEncryptedUserCommunity.CommunityName) + "']"
}

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-name"] = informEncryptedUserCommunity.CommunityName
    leafs["port"] = informEncryptedUserCommunity.Port
    leafs["version"] = informEncryptedUserCommunity.Version
    leafs["security-level"] = informEncryptedUserCommunity.SecurityLevel
    leafs["basic-trap-types"] = informEncryptedUserCommunity.BasicTrapTypes
    leafs["advanced-trap-types1"] = informEncryptedUserCommunity.AdvancedTrapTypes1
    leafs["advanced-trap-types2"] = informEncryptedUserCommunity.AdvancedTrapTypes2
    return leafs
}

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetYangName() string { return "inform-encrypted-user-community" }

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) SetParent(parent types.Entity) { informEncryptedUserCommunity.parent = parent }

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetParent() types.Entity { return informEncryptedUserCommunity.parent }

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetParentYangName() string { return "inform-encrypted-user-communities" }

// Snmp_TrapHosts_TrapHost_DefaultUserCommunities
// Container class for defining communities for a
// trap host
type Snmp_TrapHosts_TrapHost_DefaultUserCommunities struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Unencrpted Community name associated with a trap host. The type is slice of
    // Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity.
    DefaultUserCommunity []Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity
}

func (defaultUserCommunities *Snmp_TrapHosts_TrapHost_DefaultUserCommunities) GetFilter() yfilter.YFilter { return defaultUserCommunities.YFilter }

func (defaultUserCommunities *Snmp_TrapHosts_TrapHost_DefaultUserCommunities) SetFilter(yf yfilter.YFilter) { defaultUserCommunities.YFilter = yf }

func (defaultUserCommunities *Snmp_TrapHosts_TrapHost_DefaultUserCommunities) GetGoName(yname string) string {
    if yname == "default-user-community" { return "DefaultUserCommunity" }
    return ""
}

func (defaultUserCommunities *Snmp_TrapHosts_TrapHost_DefaultUserCommunities) GetSegmentPath() string {
    return "default-user-communities"
}

func (defaultUserCommunities *Snmp_TrapHosts_TrapHost_DefaultUserCommunities) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-user-community" {
        for _, c := range defaultUserCommunities.DefaultUserCommunity {
            if defaultUserCommunities.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity{}
        defaultUserCommunities.DefaultUserCommunity = append(defaultUserCommunities.DefaultUserCommunity, child)
        return &defaultUserCommunities.DefaultUserCommunity[len(defaultUserCommunities.DefaultUserCommunity)-1]
    }
    return nil
}

func (defaultUserCommunities *Snmp_TrapHosts_TrapHost_DefaultUserCommunities) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range defaultUserCommunities.DefaultUserCommunity {
        children[defaultUserCommunities.DefaultUserCommunity[i].GetSegmentPath()] = &defaultUserCommunities.DefaultUserCommunity[i]
    }
    return children
}

func (defaultUserCommunities *Snmp_TrapHosts_TrapHost_DefaultUserCommunities) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (defaultUserCommunities *Snmp_TrapHosts_TrapHost_DefaultUserCommunities) GetBundleName() string { return "cisco_ios_xr" }

func (defaultUserCommunities *Snmp_TrapHosts_TrapHost_DefaultUserCommunities) GetYangName() string { return "default-user-communities" }

func (defaultUserCommunities *Snmp_TrapHosts_TrapHost_DefaultUserCommunities) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultUserCommunities *Snmp_TrapHosts_TrapHost_DefaultUserCommunities) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultUserCommunities *Snmp_TrapHosts_TrapHost_DefaultUserCommunities) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultUserCommunities *Snmp_TrapHosts_TrapHost_DefaultUserCommunities) SetParent(parent types.Entity) { defaultUserCommunities.parent = parent }

func (defaultUserCommunities *Snmp_TrapHosts_TrapHost_DefaultUserCommunities) GetParent() types.Entity { return defaultUserCommunities.parent }

func (defaultUserCommunities *Snmp_TrapHosts_TrapHost_DefaultUserCommunities) GetParentYangName() string { return "trap-host" }

// Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity
// Unencrpted Community name associated with a
// trap host
type Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. SNMPv1/v2c community string or SNMPv3 user. The
    // type is string with length: 1..128.
    CommunityName interface{}

    // UDP port number. The type is interface{} with range: 1..65535.
    Port interface{}

    // SNMP Version to be used v1/v2c/v3. The type is string. This attribute is
    // mandatory.
    Version interface{}

    // Security level to be used noauth/auth/priv. The type is SnmpSecurityModel.
    SecurityLevel interface{}

    // Number to signify the feature traps that needs to be setBasicTrapTypes is
    // used for all traps except copy-completeSet this value to an integer
    // corresponding to the trapBGP 8192, CONFIG 4096,SYSLOG 131072,SNMP_TRAP
    // 1COPY_COMPLETE_TRAP 64To provide a combination of trap Add the respective
    // numbersValue must be set to 0 for all traps. The type is interface{} with
    // range: -2147483648..2147483647. This attribute is mandatory.
    BasicTrapTypes interface{}

    // Number to signify the feature traps that needs to be setUse this for
    // providing copy-complete trapValue must be set to 0 if not used. The type is
    // interface{} with range: -2147483648..2147483647. This attribute is
    // mandatory.
    AdvancedTrapTypes1 interface{}

    // Number to signify the feature traps that needs to be setvalue should always
    // to set as 0. The type is interface{} with range: -2147483648..2147483647.
    // This attribute is mandatory.
    AdvancedTrapTypes2 interface{}
}

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetFilter() yfilter.YFilter { return defaultUserCommunity.YFilter }

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) SetFilter(yf yfilter.YFilter) { defaultUserCommunity.YFilter = yf }

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetGoName(yname string) string {
    if yname == "community-name" { return "CommunityName" }
    if yname == "port" { return "Port" }
    if yname == "version" { return "Version" }
    if yname == "security-level" { return "SecurityLevel" }
    if yname == "basic-trap-types" { return "BasicTrapTypes" }
    if yname == "advanced-trap-types1" { return "AdvancedTrapTypes1" }
    if yname == "advanced-trap-types2" { return "AdvancedTrapTypes2" }
    return ""
}

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetSegmentPath() string {
    return "default-user-community" + "[community-name='" + fmt.Sprintf("%v", defaultUserCommunity.CommunityName) + "']"
}

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["community-name"] = defaultUserCommunity.CommunityName
    leafs["port"] = defaultUserCommunity.Port
    leafs["version"] = defaultUserCommunity.Version
    leafs["security-level"] = defaultUserCommunity.SecurityLevel
    leafs["basic-trap-types"] = defaultUserCommunity.BasicTrapTypes
    leafs["advanced-trap-types1"] = defaultUserCommunity.AdvancedTrapTypes1
    leafs["advanced-trap-types2"] = defaultUserCommunity.AdvancedTrapTypes2
    return leafs
}

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetBundleName() string { return "cisco_ios_xr" }

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetYangName() string { return "default-user-community" }

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) SetParent(parent types.Entity) { defaultUserCommunity.parent = parent }

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetParent() types.Entity { return defaultUserCommunity.parent }

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetParentYangName() string { return "default-user-communities" }

// Snmp_Contexts
// List of Context Names
type Snmp_Contexts struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Context Name. The type is slice of Snmp_Contexts_Context.
    Context []Snmp_Contexts_Context
}

func (contexts *Snmp_Contexts) GetFilter() yfilter.YFilter { return contexts.YFilter }

func (contexts *Snmp_Contexts) SetFilter(yf yfilter.YFilter) { contexts.YFilter = yf }

func (contexts *Snmp_Contexts) GetGoName(yname string) string {
    if yname == "context" { return "Context" }
    return ""
}

func (contexts *Snmp_Contexts) GetSegmentPath() string {
    return "contexts"
}

func (contexts *Snmp_Contexts) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "context" {
        for _, c := range contexts.Context {
            if contexts.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_Contexts_Context{}
        contexts.Context = append(contexts.Context, child)
        return &contexts.Context[len(contexts.Context)-1]
    }
    return nil
}

func (contexts *Snmp_Contexts) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range contexts.Context {
        children[contexts.Context[i].GetSegmentPath()] = &contexts.Context[i]
    }
    return children
}

func (contexts *Snmp_Contexts) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (contexts *Snmp_Contexts) GetBundleName() string { return "cisco_ios_xr" }

func (contexts *Snmp_Contexts) GetYangName() string { return "contexts" }

func (contexts *Snmp_Contexts) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contexts *Snmp_Contexts) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contexts *Snmp_Contexts) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contexts *Snmp_Contexts) SetParent(parent types.Entity) { contexts.parent = parent }

func (contexts *Snmp_Contexts) GetParent() types.Entity { return contexts.parent }

func (contexts *Snmp_Contexts) GetParentYangName() string { return "snmp" }

// Snmp_Contexts_Context
// Context Name
type Snmp_Contexts_Context struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Context Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ContextName interface{}
}

func (context *Snmp_Contexts_Context) GetFilter() yfilter.YFilter { return context.YFilter }

func (context *Snmp_Contexts_Context) SetFilter(yf yfilter.YFilter) { context.YFilter = yf }

func (context *Snmp_Contexts_Context) GetGoName(yname string) string {
    if yname == "context-name" { return "ContextName" }
    return ""
}

func (context *Snmp_Contexts_Context) GetSegmentPath() string {
    return "context" + "[context-name='" + fmt.Sprintf("%v", context.ContextName) + "']"
}

func (context *Snmp_Contexts_Context) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (context *Snmp_Contexts_Context) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (context *Snmp_Contexts_Context) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["context-name"] = context.ContextName
    return leafs
}

func (context *Snmp_Contexts_Context) GetBundleName() string { return "cisco_ios_xr" }

func (context *Snmp_Contexts_Context) GetYangName() string { return "context" }

func (context *Snmp_Contexts_Context) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (context *Snmp_Contexts_Context) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (context *Snmp_Contexts_Context) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (context *Snmp_Contexts_Context) SetParent(parent types.Entity) { context.parent = parent }

func (context *Snmp_Contexts_Context) GetParent() types.Entity { return context.parent }

func (context *Snmp_Contexts_Context) GetParentYangName() string { return "contexts" }

// Snmp_ContextMappings
// List of context names
type Snmp_ContextMappings struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Context mapping name. The type is slice of
    // Snmp_ContextMappings_ContextMapping.
    ContextMapping []Snmp_ContextMappings_ContextMapping
}

func (contextMappings *Snmp_ContextMappings) GetFilter() yfilter.YFilter { return contextMappings.YFilter }

func (contextMappings *Snmp_ContextMappings) SetFilter(yf yfilter.YFilter) { contextMappings.YFilter = yf }

func (contextMappings *Snmp_ContextMappings) GetGoName(yname string) string {
    if yname == "context-mapping" { return "ContextMapping" }
    return ""
}

func (contextMappings *Snmp_ContextMappings) GetSegmentPath() string {
    return "context-mappings"
}

func (contextMappings *Snmp_ContextMappings) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "context-mapping" {
        for _, c := range contextMappings.ContextMapping {
            if contextMappings.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Snmp_ContextMappings_ContextMapping{}
        contextMappings.ContextMapping = append(contextMappings.ContextMapping, child)
        return &contextMappings.ContextMapping[len(contextMappings.ContextMapping)-1]
    }
    return nil
}

func (contextMappings *Snmp_ContextMappings) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range contextMappings.ContextMapping {
        children[contextMappings.ContextMapping[i].GetSegmentPath()] = &contextMappings.ContextMapping[i]
    }
    return children
}

func (contextMappings *Snmp_ContextMappings) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (contextMappings *Snmp_ContextMappings) GetBundleName() string { return "cisco_ios_xr" }

func (contextMappings *Snmp_ContextMappings) GetYangName() string { return "context-mappings" }

func (contextMappings *Snmp_ContextMappings) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contextMappings *Snmp_ContextMappings) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contextMappings *Snmp_ContextMappings) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contextMappings *Snmp_ContextMappings) SetParent(parent types.Entity) { contextMappings.parent = parent }

func (contextMappings *Snmp_ContextMappings) GetParent() types.Entity { return contextMappings.parent }

func (contextMappings *Snmp_ContextMappings) GetParentYangName() string { return "snmp" }

// Snmp_ContextMappings_ContextMapping
// Context mapping name
type Snmp_ContextMappings_ContextMapping struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Context mapping name. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ContextMappingName interface{}

    // SNMP context feature type. The type is SnmpContext.
    Context interface{}

    // OSPF protocol instance. The type is string.
    InstanceName interface{}

    // VRF name associated with the context. The type is string.
    VrfName interface{}

    // Topology name associated with the context. The type is string.
    TopologyName interface{}
}

func (contextMapping *Snmp_ContextMappings_ContextMapping) GetFilter() yfilter.YFilter { return contextMapping.YFilter }

func (contextMapping *Snmp_ContextMappings_ContextMapping) SetFilter(yf yfilter.YFilter) { contextMapping.YFilter = yf }

func (contextMapping *Snmp_ContextMappings_ContextMapping) GetGoName(yname string) string {
    if yname == "context-mapping-name" { return "ContextMappingName" }
    if yname == "context" { return "Context" }
    if yname == "instance-name" { return "InstanceName" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "topology-name" { return "TopologyName" }
    return ""
}

func (contextMapping *Snmp_ContextMappings_ContextMapping) GetSegmentPath() string {
    return "context-mapping" + "[context-mapping-name='" + fmt.Sprintf("%v", contextMapping.ContextMappingName) + "']"
}

func (contextMapping *Snmp_ContextMappings_ContextMapping) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (contextMapping *Snmp_ContextMappings_ContextMapping) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (contextMapping *Snmp_ContextMappings_ContextMapping) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["context-mapping-name"] = contextMapping.ContextMappingName
    leafs["context"] = contextMapping.Context
    leafs["instance-name"] = contextMapping.InstanceName
    leafs["vrf-name"] = contextMapping.VrfName
    leafs["topology-name"] = contextMapping.TopologyName
    return leafs
}

func (contextMapping *Snmp_ContextMappings_ContextMapping) GetBundleName() string { return "cisco_ios_xr" }

func (contextMapping *Snmp_ContextMappings_ContextMapping) GetYangName() string { return "context-mapping" }

func (contextMapping *Snmp_ContextMappings_ContextMapping) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (contextMapping *Snmp_ContextMappings_ContextMapping) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (contextMapping *Snmp_ContextMappings_ContextMapping) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (contextMapping *Snmp_ContextMappings_ContextMapping) SetParent(parent types.Entity) { contextMapping.parent = parent }

func (contextMapping *Snmp_ContextMappings_ContextMapping) GetParent() types.Entity { return contextMapping.parent }

func (contextMapping *Snmp_ContextMappings_ContextMapping) GetParentYangName() string { return "context-mappings" }

// Mib
// mib
type Mib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Get cached Sesnsor MIB statistics. The type is interface{}.
    SensorMibCache interface{}

    // MPLS TE MIB configuration.
    MplsTeMib Mib_MplsTeMib

    // MPLS P2MP MIB configuration.
    MplsP2MpMib Mib_MplsP2MpMib

    // MPLS TE EXT STD MIB configuration.
    MplsTeExtStdMib Mib_MplsTeExtStdMib

    // MPLS TE EXT MIB configuration.
    MplsTeExtMib Mib_MplsTeExtMib

    // FRR MIB configuration.
    MplsFrrMib Mib_MplsFrrMib

    // CBQoSMIB configuration.
    CbQosmib Mib_CbQosmib

    // Entity MIB.
    EntityMib Mib_EntityMib

    // Interface MIB configuration.
    InterfaceMib Mib_InterfaceMib

    // Subscriber threshold commands.
    Subscriber Mib_Subscriber
}

func (mib *Mib) GetFilter() yfilter.YFilter { return mib.YFilter }

func (mib *Mib) SetFilter(yf yfilter.YFilter) { mib.YFilter = yf }

func (mib *Mib) GetGoName(yname string) string {
    if yname == "sensor-mib-cache" { return "SensorMibCache" }
    if yname == "Cisco-IOS-XR-mpls-te-cfg:mpls-te-mib" { return "MplsTeMib" }
    if yname == "Cisco-IOS-XR-mpls-te-cfg:mpls-p2mp-mib" { return "MplsP2MpMib" }
    if yname == "Cisco-IOS-XR-mpls-te-cfg:mpls-te-ext-std-mib" { return "MplsTeExtStdMib" }
    if yname == "Cisco-IOS-XR-mpls-te-cfg:mpls-te-ext-mib" { return "MplsTeExtMib" }
    if yname == "Cisco-IOS-XR-mpls-te-cfg:mpls-frr-mib" { return "MplsFrrMib" }
    if yname == "Cisco-IOS-XR-qos-mibs-cfg:cb-qosmib" { return "CbQosmib" }
    if yname == "Cisco-IOS-XR-snmp-entitymib-cfg:entity-mib" { return "EntityMib" }
    if yname == "Cisco-IOS-XR-snmp-ifmib-cfg:interface-mib" { return "InterfaceMib" }
    if yname == "Cisco-IOS-XR-subscriber-session-mon-mibs-cfg:subscriber" { return "Subscriber" }
    return ""
}

func (mib *Mib) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-agent-cfg:mib"
}

func (mib *Mib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "Cisco-IOS-XR-mpls-te-cfg:mpls-te-mib" {
        return &mib.MplsTeMib
    }
    if childYangName == "Cisco-IOS-XR-mpls-te-cfg:mpls-p2mp-mib" {
        return &mib.MplsP2MpMib
    }
    if childYangName == "Cisco-IOS-XR-mpls-te-cfg:mpls-te-ext-std-mib" {
        return &mib.MplsTeExtStdMib
    }
    if childYangName == "Cisco-IOS-XR-mpls-te-cfg:mpls-te-ext-mib" {
        return &mib.MplsTeExtMib
    }
    if childYangName == "Cisco-IOS-XR-mpls-te-cfg:mpls-frr-mib" {
        return &mib.MplsFrrMib
    }
    if childYangName == "Cisco-IOS-XR-qos-mibs-cfg:cb-qosmib" {
        return &mib.CbQosmib
    }
    if childYangName == "Cisco-IOS-XR-snmp-entitymib-cfg:entity-mib" {
        return &mib.EntityMib
    }
    if childYangName == "Cisco-IOS-XR-snmp-ifmib-cfg:interface-mib" {
        return &mib.InterfaceMib
    }
    if childYangName == "Cisco-IOS-XR-subscriber-session-mon-mibs-cfg:subscriber" {
        return &mib.Subscriber
    }
    return nil
}

func (mib *Mib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["Cisco-IOS-XR-mpls-te-cfg:mpls-te-mib"] = &mib.MplsTeMib
    children["Cisco-IOS-XR-mpls-te-cfg:mpls-p2mp-mib"] = &mib.MplsP2MpMib
    children["Cisco-IOS-XR-mpls-te-cfg:mpls-te-ext-std-mib"] = &mib.MplsTeExtStdMib
    children["Cisco-IOS-XR-mpls-te-cfg:mpls-te-ext-mib"] = &mib.MplsTeExtMib
    children["Cisco-IOS-XR-mpls-te-cfg:mpls-frr-mib"] = &mib.MplsFrrMib
    children["Cisco-IOS-XR-qos-mibs-cfg:cb-qosmib"] = &mib.CbQosmib
    children["Cisco-IOS-XR-snmp-entitymib-cfg:entity-mib"] = &mib.EntityMib
    children["Cisco-IOS-XR-snmp-ifmib-cfg:interface-mib"] = &mib.InterfaceMib
    children["Cisco-IOS-XR-subscriber-session-mon-mibs-cfg:subscriber"] = &mib.Subscriber
    return children
}

func (mib *Mib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["sensor-mib-cache"] = mib.SensorMibCache
    return leafs
}

func (mib *Mib) GetBundleName() string { return "cisco_ios_xr" }

func (mib *Mib) GetYangName() string { return "mib" }

func (mib *Mib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mib *Mib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mib *Mib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mib *Mib) SetParent(parent types.Entity) { mib.parent = parent }

func (mib *Mib) GetParent() types.Entity { return mib.parent }

func (mib *Mib) GetParentYangName() string { return "Cisco-IOS-XR-snmp-agent-cfg" }

// Mib_MplsTeMib
// MPLS TE MIB configuration
type Mib_MplsTeMib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the cache garbage collector time for the mib. The type is
    // interface{} with range: 0..3600. Units are second. The default value is
    // 1800.
    CacheGarbageCollectTimer interface{}

    // Configure the cache time for the mib. The type is interface{} with range:
    // 0..600. Units are second. The default value is 60.
    CacheTimer interface{}
}

func (mplsTeMib *Mib_MplsTeMib) GetFilter() yfilter.YFilter { return mplsTeMib.YFilter }

func (mplsTeMib *Mib_MplsTeMib) SetFilter(yf yfilter.YFilter) { mplsTeMib.YFilter = yf }

func (mplsTeMib *Mib_MplsTeMib) GetGoName(yname string) string {
    if yname == "cache-garbage-collect-timer" { return "CacheGarbageCollectTimer" }
    if yname == "cache-timer" { return "CacheTimer" }
    return ""
}

func (mplsTeMib *Mib_MplsTeMib) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-te-cfg:mpls-te-mib"
}

func (mplsTeMib *Mib_MplsTeMib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsTeMib *Mib_MplsTeMib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsTeMib *Mib_MplsTeMib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cache-garbage-collect-timer"] = mplsTeMib.CacheGarbageCollectTimer
    leafs["cache-timer"] = mplsTeMib.CacheTimer
    return leafs
}

func (mplsTeMib *Mib_MplsTeMib) GetBundleName() string { return "cisco_ios_xr" }

func (mplsTeMib *Mib_MplsTeMib) GetYangName() string { return "mpls-te-mib" }

func (mplsTeMib *Mib_MplsTeMib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsTeMib *Mib_MplsTeMib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsTeMib *Mib_MplsTeMib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsTeMib *Mib_MplsTeMib) SetParent(parent types.Entity) { mplsTeMib.parent = parent }

func (mplsTeMib *Mib_MplsTeMib) GetParent() types.Entity { return mplsTeMib.parent }

func (mplsTeMib *Mib_MplsTeMib) GetParentYangName() string { return "mib" }

// Mib_MplsP2MpMib
// MPLS P2MP MIB configuration
type Mib_MplsP2MpMib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the cache time for the mib. The type is interface{} with range:
    // 0..600. Units are second. The default value is 60.
    CacheTimer interface{}
}

func (mplsP2MpMib *Mib_MplsP2MpMib) GetFilter() yfilter.YFilter { return mplsP2MpMib.YFilter }

func (mplsP2MpMib *Mib_MplsP2MpMib) SetFilter(yf yfilter.YFilter) { mplsP2MpMib.YFilter = yf }

func (mplsP2MpMib *Mib_MplsP2MpMib) GetGoName(yname string) string {
    if yname == "cache-timer" { return "CacheTimer" }
    return ""
}

func (mplsP2MpMib *Mib_MplsP2MpMib) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-te-cfg:mpls-p2mp-mib"
}

func (mplsP2MpMib *Mib_MplsP2MpMib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsP2MpMib *Mib_MplsP2MpMib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsP2MpMib *Mib_MplsP2MpMib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cache-timer"] = mplsP2MpMib.CacheTimer
    return leafs
}

func (mplsP2MpMib *Mib_MplsP2MpMib) GetBundleName() string { return "cisco_ios_xr" }

func (mplsP2MpMib *Mib_MplsP2MpMib) GetYangName() string { return "mpls-p2mp-mib" }

func (mplsP2MpMib *Mib_MplsP2MpMib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsP2MpMib *Mib_MplsP2MpMib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsP2MpMib *Mib_MplsP2MpMib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsP2MpMib *Mib_MplsP2MpMib) SetParent(parent types.Entity) { mplsP2MpMib.parent = parent }

func (mplsP2MpMib *Mib_MplsP2MpMib) GetParent() types.Entity { return mplsP2MpMib.parent }

func (mplsP2MpMib *Mib_MplsP2MpMib) GetParentYangName() string { return "mib" }

// Mib_MplsTeExtStdMib
// MPLS TE EXT STD MIB configuration
type Mib_MplsTeExtStdMib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the cache time for the mib. The type is interface{} with range:
    // 0..600. Units are second. The default value is 60.
    CacheTimer interface{}
}

func (mplsTeExtStdMib *Mib_MplsTeExtStdMib) GetFilter() yfilter.YFilter { return mplsTeExtStdMib.YFilter }

func (mplsTeExtStdMib *Mib_MplsTeExtStdMib) SetFilter(yf yfilter.YFilter) { mplsTeExtStdMib.YFilter = yf }

func (mplsTeExtStdMib *Mib_MplsTeExtStdMib) GetGoName(yname string) string {
    if yname == "cache-timer" { return "CacheTimer" }
    return ""
}

func (mplsTeExtStdMib *Mib_MplsTeExtStdMib) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-te-cfg:mpls-te-ext-std-mib"
}

func (mplsTeExtStdMib *Mib_MplsTeExtStdMib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsTeExtStdMib *Mib_MplsTeExtStdMib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsTeExtStdMib *Mib_MplsTeExtStdMib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cache-timer"] = mplsTeExtStdMib.CacheTimer
    return leafs
}

func (mplsTeExtStdMib *Mib_MplsTeExtStdMib) GetBundleName() string { return "cisco_ios_xr" }

func (mplsTeExtStdMib *Mib_MplsTeExtStdMib) GetYangName() string { return "mpls-te-ext-std-mib" }

func (mplsTeExtStdMib *Mib_MplsTeExtStdMib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsTeExtStdMib *Mib_MplsTeExtStdMib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsTeExtStdMib *Mib_MplsTeExtStdMib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsTeExtStdMib *Mib_MplsTeExtStdMib) SetParent(parent types.Entity) { mplsTeExtStdMib.parent = parent }

func (mplsTeExtStdMib *Mib_MplsTeExtStdMib) GetParent() types.Entity { return mplsTeExtStdMib.parent }

func (mplsTeExtStdMib *Mib_MplsTeExtStdMib) GetParentYangName() string { return "mib" }

// Mib_MplsTeExtMib
// MPLS TE EXT MIB configuration
type Mib_MplsTeExtMib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the cache time for the mib. The type is interface{} with range:
    // 0..600. Units are second. The default value is 60.
    CacheTimer interface{}
}

func (mplsTeExtMib *Mib_MplsTeExtMib) GetFilter() yfilter.YFilter { return mplsTeExtMib.YFilter }

func (mplsTeExtMib *Mib_MplsTeExtMib) SetFilter(yf yfilter.YFilter) { mplsTeExtMib.YFilter = yf }

func (mplsTeExtMib *Mib_MplsTeExtMib) GetGoName(yname string) string {
    if yname == "cache-timer" { return "CacheTimer" }
    return ""
}

func (mplsTeExtMib *Mib_MplsTeExtMib) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-te-cfg:mpls-te-ext-mib"
}

func (mplsTeExtMib *Mib_MplsTeExtMib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsTeExtMib *Mib_MplsTeExtMib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsTeExtMib *Mib_MplsTeExtMib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cache-timer"] = mplsTeExtMib.CacheTimer
    return leafs
}

func (mplsTeExtMib *Mib_MplsTeExtMib) GetBundleName() string { return "cisco_ios_xr" }

func (mplsTeExtMib *Mib_MplsTeExtMib) GetYangName() string { return "mpls-te-ext-mib" }

func (mplsTeExtMib *Mib_MplsTeExtMib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsTeExtMib *Mib_MplsTeExtMib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsTeExtMib *Mib_MplsTeExtMib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsTeExtMib *Mib_MplsTeExtMib) SetParent(parent types.Entity) { mplsTeExtMib.parent = parent }

func (mplsTeExtMib *Mib_MplsTeExtMib) GetParent() types.Entity { return mplsTeExtMib.parent }

func (mplsTeExtMib *Mib_MplsTeExtMib) GetParentYangName() string { return "mib" }

// Mib_MplsFrrMib
// FRR MIB configuration
type Mib_MplsFrrMib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure the cache time for the mib. The type is interface{} with range:
    // 0..600. Units are second. The default value is 60.
    CacheTimer interface{}
}

func (mplsFrrMib *Mib_MplsFrrMib) GetFilter() yfilter.YFilter { return mplsFrrMib.YFilter }

func (mplsFrrMib *Mib_MplsFrrMib) SetFilter(yf yfilter.YFilter) { mplsFrrMib.YFilter = yf }

func (mplsFrrMib *Mib_MplsFrrMib) GetGoName(yname string) string {
    if yname == "cache-timer" { return "CacheTimer" }
    return ""
}

func (mplsFrrMib *Mib_MplsFrrMib) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-te-cfg:mpls-frr-mib"
}

func (mplsFrrMib *Mib_MplsFrrMib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (mplsFrrMib *Mib_MplsFrrMib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (mplsFrrMib *Mib_MplsFrrMib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cache-timer"] = mplsFrrMib.CacheTimer
    return leafs
}

func (mplsFrrMib *Mib_MplsFrrMib) GetBundleName() string { return "cisco_ios_xr" }

func (mplsFrrMib *Mib_MplsFrrMib) GetYangName() string { return "mpls-frr-mib" }

func (mplsFrrMib *Mib_MplsFrrMib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsFrrMib *Mib_MplsFrrMib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsFrrMib *Mib_MplsFrrMib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsFrrMib *Mib_MplsFrrMib) SetParent(parent types.Entity) { mplsFrrMib.parent = parent }

func (mplsFrrMib *Mib_MplsFrrMib) GetParent() types.Entity { return mplsFrrMib.parent }

func (mplsFrrMib *Mib_MplsFrrMib) GetParentYangName() string { return "mib" }

// Mib_CbQosmib
// CBQoSMIB configuration
type Mib_CbQosmib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable bundle member interface statistics retrieval. The type is
    // interface{}.
    MemberInterfaceStats interface{}

    // Persist CBQoSMIB config, service-policy and object indices. The type is
    // interface{}.
    Persist interface{}

    // CBQoSMIB statistics data caching.
    Cache Mib_CbQosmib_Cache
}

func (cbQosmib *Mib_CbQosmib) GetFilter() yfilter.YFilter { return cbQosmib.YFilter }

func (cbQosmib *Mib_CbQosmib) SetFilter(yf yfilter.YFilter) { cbQosmib.YFilter = yf }

func (cbQosmib *Mib_CbQosmib) GetGoName(yname string) string {
    if yname == "member-interface-stats" { return "MemberInterfaceStats" }
    if yname == "persist" { return "Persist" }
    if yname == "cache" { return "Cache" }
    return ""
}

func (cbQosmib *Mib_CbQosmib) GetSegmentPath() string {
    return "Cisco-IOS-XR-qos-mibs-cfg:cb-qosmib"
}

func (cbQosmib *Mib_CbQosmib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cache" {
        return &cbQosmib.Cache
    }
    return nil
}

func (cbQosmib *Mib_CbQosmib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cache"] = &cbQosmib.Cache
    return children
}

func (cbQosmib *Mib_CbQosmib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["member-interface-stats"] = cbQosmib.MemberInterfaceStats
    leafs["persist"] = cbQosmib.Persist
    return leafs
}

func (cbQosmib *Mib_CbQosmib) GetBundleName() string { return "cisco_ios_xr" }

func (cbQosmib *Mib_CbQosmib) GetYangName() string { return "cb-qosmib" }

func (cbQosmib *Mib_CbQosmib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cbQosmib *Mib_CbQosmib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cbQosmib *Mib_CbQosmib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cbQosmib *Mib_CbQosmib) SetParent(parent types.Entity) { cbQosmib.parent = parent }

func (cbQosmib *Mib_CbQosmib) GetParent() types.Entity { return cbQosmib.parent }

func (cbQosmib *Mib_CbQosmib) GetParentYangName() string { return "mib" }

// Mib_CbQosmib_Cache
// CBQoSMIB statistics data caching
type Mib_CbQosmib_Cache struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable CBQoSMIB statistics data caching. The type is interface{}.
    Enable interface{}

    // Cache refresh time in seconds. The type is interface{} with range: 5..60.
    // Units are second.
    RefreshTime interface{}

    // Maximum number of service policies to cache the statistics for. The type is
    // interface{} with range: 1..5000.
    ServicePolicyCount interface{}
}

func (cache *Mib_CbQosmib_Cache) GetFilter() yfilter.YFilter { return cache.YFilter }

func (cache *Mib_CbQosmib_Cache) SetFilter(yf yfilter.YFilter) { cache.YFilter = yf }

func (cache *Mib_CbQosmib_Cache) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "refresh-time" { return "RefreshTime" }
    if yname == "service-policy-count" { return "ServicePolicyCount" }
    return ""
}

func (cache *Mib_CbQosmib_Cache) GetSegmentPath() string {
    return "cache"
}

func (cache *Mib_CbQosmib_Cache) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cache *Mib_CbQosmib_Cache) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cache *Mib_CbQosmib_Cache) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = cache.Enable
    leafs["refresh-time"] = cache.RefreshTime
    leafs["service-policy-count"] = cache.ServicePolicyCount
    return leafs
}

func (cache *Mib_CbQosmib_Cache) GetBundleName() string { return "cisco_ios_xr" }

func (cache *Mib_CbQosmib_Cache) GetYangName() string { return "cache" }

func (cache *Mib_CbQosmib_Cache) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (cache *Mib_CbQosmib_Cache) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (cache *Mib_CbQosmib_Cache) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (cache *Mib_CbQosmib_Cache) SetParent(parent types.Entity) { cache.parent = parent }

func (cache *Mib_CbQosmib_Cache) GetParent() types.Entity { return cache.parent }

func (cache *Mib_CbQosmib_Cache) GetParentYangName() string { return "cb-qosmib" }

// Mib_EntityMib
// Entity MIB
type Mib_EntityMib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable entPhysicalIndex persistence. The type is interface{}.
    EntityIndexPersistence interface{}
}

func (entityMib *Mib_EntityMib) GetFilter() yfilter.YFilter { return entityMib.YFilter }

func (entityMib *Mib_EntityMib) SetFilter(yf yfilter.YFilter) { entityMib.YFilter = yf }

func (entityMib *Mib_EntityMib) GetGoName(yname string) string {
    if yname == "entity-index-persistence" { return "EntityIndexPersistence" }
    return ""
}

func (entityMib *Mib_EntityMib) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-entitymib-cfg:entity-mib"
}

func (entityMib *Mib_EntityMib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (entityMib *Mib_EntityMib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (entityMib *Mib_EntityMib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entity-index-persistence"] = entityMib.EntityIndexPersistence
    return leafs
}

func (entityMib *Mib_EntityMib) GetBundleName() string { return "cisco_ios_xr" }

func (entityMib *Mib_EntityMib) GetYangName() string { return "entity-mib" }

func (entityMib *Mib_EntityMib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (entityMib *Mib_EntityMib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (entityMib *Mib_EntityMib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (entityMib *Mib_EntityMib) SetParent(parent types.Entity) { entityMib.parent = parent }

func (entityMib *Mib_EntityMib) GetParent() types.Entity { return entityMib.parent }

func (entityMib *Mib_EntityMib) GetParentYangName() string { return "mib" }

// Mib_InterfaceMib
// Interface MIB configuration
type Mib_InterfaceMib struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Get cached interface statistics. The type is interface{} with range: 0..60.
    // The default value is 15.
    InternalCache interface{}

    // Enable support for ifAlias values longer than 64 characters. The type is
    // interface{}.
    InterfaceAliasLong interface{}

    // Enable IP subscriber interfaces in IFMIB. The type is interface{}.
    IpSubscriber interface{}

    // Enable ifindex persistence. The type is interface{}.
    InterfaceIndexPersistence interface{}

    // Enable cached interface statistics. The type is interface{}.
    StatisticsCache interface{}

    // Enter the SNMP interface configuration commands.
    Interfaces Mib_InterfaceMib_Interfaces

    // MIB notification configuration.
    Notification Mib_InterfaceMib_Notification

    // Add configuration for an interface subset.
    Subsets Mib_InterfaceMib_Subsets
}

func (interfaceMib *Mib_InterfaceMib) GetFilter() yfilter.YFilter { return interfaceMib.YFilter }

func (interfaceMib *Mib_InterfaceMib) SetFilter(yf yfilter.YFilter) { interfaceMib.YFilter = yf }

func (interfaceMib *Mib_InterfaceMib) GetGoName(yname string) string {
    if yname == "internal-cache" { return "InternalCache" }
    if yname == "interface-alias-long" { return "InterfaceAliasLong" }
    if yname == "ip-subscriber" { return "IpSubscriber" }
    if yname == "interface-index-persistence" { return "InterfaceIndexPersistence" }
    if yname == "statistics-cache" { return "StatisticsCache" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "notification" { return "Notification" }
    if yname == "subsets" { return "Subsets" }
    return ""
}

func (interfaceMib *Mib_InterfaceMib) GetSegmentPath() string {
    return "Cisco-IOS-XR-snmp-ifmib-cfg:interface-mib"
}

func (interfaceMib *Mib_InterfaceMib) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &interfaceMib.Interfaces
    }
    if childYangName == "notification" {
        return &interfaceMib.Notification
    }
    if childYangName == "subsets" {
        return &interfaceMib.Subsets
    }
    return nil
}

func (interfaceMib *Mib_InterfaceMib) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &interfaceMib.Interfaces
    children["notification"] = &interfaceMib.Notification
    children["subsets"] = &interfaceMib.Subsets
    return children
}

func (interfaceMib *Mib_InterfaceMib) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["internal-cache"] = interfaceMib.InternalCache
    leafs["interface-alias-long"] = interfaceMib.InterfaceAliasLong
    leafs["ip-subscriber"] = interfaceMib.IpSubscriber
    leafs["interface-index-persistence"] = interfaceMib.InterfaceIndexPersistence
    leafs["statistics-cache"] = interfaceMib.StatisticsCache
    return leafs
}

func (interfaceMib *Mib_InterfaceMib) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceMib *Mib_InterfaceMib) GetYangName() string { return "interface-mib" }

func (interfaceMib *Mib_InterfaceMib) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceMib *Mib_InterfaceMib) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceMib *Mib_InterfaceMib) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceMib *Mib_InterfaceMib) SetParent(parent types.Entity) { interfaceMib.parent = parent }

func (interfaceMib *Mib_InterfaceMib) GetParent() types.Entity { return interfaceMib.parent }

func (interfaceMib *Mib_InterfaceMib) GetParentYangName() string { return "mib" }

// Mib_InterfaceMib_Interfaces
// Enter the SNMP interface configuration commands
type Mib_InterfaceMib_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface to configure. The type is slice of
    // Mib_InterfaceMib_Interfaces_Interface.
    Interface []Mib_InterfaceMib_Interfaces_Interface
}

func (interfaces *Mib_InterfaceMib_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Mib_InterfaceMib_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Mib_InterfaceMib_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Mib_InterfaceMib_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Mib_InterfaceMib_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mib_InterfaceMib_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Mib_InterfaceMib_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Mib_InterfaceMib_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Mib_InterfaceMib_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Mib_InterfaceMib_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Mib_InterfaceMib_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Mib_InterfaceMib_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Mib_InterfaceMib_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Mib_InterfaceMib_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Mib_InterfaceMib_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Mib_InterfaceMib_Interfaces) GetParentYangName() string { return "interface-mib" }

// Mib_InterfaceMib_Interfaces_Interface
// Interface to configure
type Mib_InterfaceMib_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Enable or disable LinkUpDown notification. The type is bool.
    LinkUpDown interface{}

    // Enable or disable index persistence. The type is bool.
    IndexPersistence interface{}
}

func (self *Mib_InterfaceMib_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Mib_InterfaceMib_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Mib_InterfaceMib_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "link-up-down" { return "LinkUpDown" }
    if yname == "index-persistence" { return "IndexPersistence" }
    return ""
}

func (self *Mib_InterfaceMib_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Mib_InterfaceMib_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Mib_InterfaceMib_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Mib_InterfaceMib_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["link-up-down"] = self.LinkUpDown
    leafs["index-persistence"] = self.IndexPersistence
    return leafs
}

func (self *Mib_InterfaceMib_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Mib_InterfaceMib_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Mib_InterfaceMib_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Mib_InterfaceMib_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Mib_InterfaceMib_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Mib_InterfaceMib_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Mib_InterfaceMib_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Mib_InterfaceMib_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Mib_InterfaceMib_Notification
// MIB notification configuration
type Mib_InterfaceMib_Notification struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set the varbind of linkupdown trap to the RFC specified varbinds (default
    // cisco). The type is interface{}.
    LinkIetf interface{}
}

func (notification *Mib_InterfaceMib_Notification) GetFilter() yfilter.YFilter { return notification.YFilter }

func (notification *Mib_InterfaceMib_Notification) SetFilter(yf yfilter.YFilter) { notification.YFilter = yf }

func (notification *Mib_InterfaceMib_Notification) GetGoName(yname string) string {
    if yname == "link-ietf" { return "LinkIetf" }
    return ""
}

func (notification *Mib_InterfaceMib_Notification) GetSegmentPath() string {
    return "notification"
}

func (notification *Mib_InterfaceMib_Notification) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (notification *Mib_InterfaceMib_Notification) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (notification *Mib_InterfaceMib_Notification) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["link-ietf"] = notification.LinkIetf
    return leafs
}

func (notification *Mib_InterfaceMib_Notification) GetBundleName() string { return "cisco_ios_xr" }

func (notification *Mib_InterfaceMib_Notification) GetYangName() string { return "notification" }

func (notification *Mib_InterfaceMib_Notification) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (notification *Mib_InterfaceMib_Notification) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (notification *Mib_InterfaceMib_Notification) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (notification *Mib_InterfaceMib_Notification) SetParent(parent types.Entity) { notification.parent = parent }

func (notification *Mib_InterfaceMib_Notification) GetParent() types.Entity { return notification.parent }

func (notification *Mib_InterfaceMib_Notification) GetParentYangName() string { return "interface-mib" }

// Mib_InterfaceMib_Subsets
// Add configuration for an interface subset
type Mib_InterfaceMib_Subsets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subset priorityID to group ifNames based on Regular Expression. The type is
    // slice of Mib_InterfaceMib_Subsets_Subset.
    Subset []Mib_InterfaceMib_Subsets_Subset
}

func (subsets *Mib_InterfaceMib_Subsets) GetFilter() yfilter.YFilter { return subsets.YFilter }

func (subsets *Mib_InterfaceMib_Subsets) SetFilter(yf yfilter.YFilter) { subsets.YFilter = yf }

func (subsets *Mib_InterfaceMib_Subsets) GetGoName(yname string) string {
    if yname == "subset" { return "Subset" }
    return ""
}

func (subsets *Mib_InterfaceMib_Subsets) GetSegmentPath() string {
    return "subsets"
}

func (subsets *Mib_InterfaceMib_Subsets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "subset" {
        for _, c := range subsets.Subset {
            if subsets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mib_InterfaceMib_Subsets_Subset{}
        subsets.Subset = append(subsets.Subset, child)
        return &subsets.Subset[len(subsets.Subset)-1]
    }
    return nil
}

func (subsets *Mib_InterfaceMib_Subsets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range subsets.Subset {
        children[subsets.Subset[i].GetSegmentPath()] = &subsets.Subset[i]
    }
    return children
}

func (subsets *Mib_InterfaceMib_Subsets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subsets *Mib_InterfaceMib_Subsets) GetBundleName() string { return "cisco_ios_xr" }

func (subsets *Mib_InterfaceMib_Subsets) GetYangName() string { return "subsets" }

func (subsets *Mib_InterfaceMib_Subsets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subsets *Mib_InterfaceMib_Subsets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subsets *Mib_InterfaceMib_Subsets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subsets *Mib_InterfaceMib_Subsets) SetParent(parent types.Entity) { subsets.parent = parent }

func (subsets *Mib_InterfaceMib_Subsets) GetParent() types.Entity { return subsets.parent }

func (subsets *Mib_InterfaceMib_Subsets) GetParentYangName() string { return "interface-mib" }

// Mib_InterfaceMib_Subsets_Subset
// Subset priorityID to group ifNames based on
// Regular Expression
type Mib_InterfaceMib_Subsets_Subset struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The interface subset PriorityID. The type is
    // interface{} with range: 1..255.
    SubsetId interface{}

    // SNMP linkUp and linkDown notifications.
    LinkUpDown Mib_InterfaceMib_Subsets_Subset_LinkUpDown
}

func (subset *Mib_InterfaceMib_Subsets_Subset) GetFilter() yfilter.YFilter { return subset.YFilter }

func (subset *Mib_InterfaceMib_Subsets_Subset) SetFilter(yf yfilter.YFilter) { subset.YFilter = yf }

func (subset *Mib_InterfaceMib_Subsets_Subset) GetGoName(yname string) string {
    if yname == "subset-id" { return "SubsetId" }
    if yname == "link-up-down" { return "LinkUpDown" }
    return ""
}

func (subset *Mib_InterfaceMib_Subsets_Subset) GetSegmentPath() string {
    return "subset" + "[subset-id='" + fmt.Sprintf("%v", subset.SubsetId) + "']"
}

func (subset *Mib_InterfaceMib_Subsets_Subset) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-up-down" {
        return &subset.LinkUpDown
    }
    return nil
}

func (subset *Mib_InterfaceMib_Subsets_Subset) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["link-up-down"] = &subset.LinkUpDown
    return children
}

func (subset *Mib_InterfaceMib_Subsets_Subset) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subset-id"] = subset.SubsetId
    return leafs
}

func (subset *Mib_InterfaceMib_Subsets_Subset) GetBundleName() string { return "cisco_ios_xr" }

func (subset *Mib_InterfaceMib_Subsets_Subset) GetYangName() string { return "subset" }

func (subset *Mib_InterfaceMib_Subsets_Subset) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subset *Mib_InterfaceMib_Subsets_Subset) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subset *Mib_InterfaceMib_Subsets_Subset) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subset *Mib_InterfaceMib_Subsets_Subset) SetParent(parent types.Entity) { subset.parent = parent }

func (subset *Mib_InterfaceMib_Subsets_Subset) GetParent() types.Entity { return subset.parent }

func (subset *Mib_InterfaceMib_Subsets_Subset) GetParentYangName() string { return "subsets" }

// Mib_InterfaceMib_Subsets_Subset_LinkUpDown
// SNMP linkUp and linkDown notifications
type Mib_InterfaceMib_Subsets_Subset_LinkUpDown struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable or disable linkupdown notification. The type is bool.
    Enable interface{}

    // Regular expression to match ifName. The type is string.
    RegularExpression interface{}
}

func (linkUpDown *Mib_InterfaceMib_Subsets_Subset_LinkUpDown) GetFilter() yfilter.YFilter { return linkUpDown.YFilter }

func (linkUpDown *Mib_InterfaceMib_Subsets_Subset_LinkUpDown) SetFilter(yf yfilter.YFilter) { linkUpDown.YFilter = yf }

func (linkUpDown *Mib_InterfaceMib_Subsets_Subset_LinkUpDown) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "regular-expression" { return "RegularExpression" }
    return ""
}

func (linkUpDown *Mib_InterfaceMib_Subsets_Subset_LinkUpDown) GetSegmentPath() string {
    return "link-up-down"
}

func (linkUpDown *Mib_InterfaceMib_Subsets_Subset_LinkUpDown) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkUpDown *Mib_InterfaceMib_Subsets_Subset_LinkUpDown) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkUpDown *Mib_InterfaceMib_Subsets_Subset_LinkUpDown) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = linkUpDown.Enable
    leafs["regular-expression"] = linkUpDown.RegularExpression
    return leafs
}

func (linkUpDown *Mib_InterfaceMib_Subsets_Subset_LinkUpDown) GetBundleName() string { return "cisco_ios_xr" }

func (linkUpDown *Mib_InterfaceMib_Subsets_Subset_LinkUpDown) GetYangName() string { return "link-up-down" }

func (linkUpDown *Mib_InterfaceMib_Subsets_Subset_LinkUpDown) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkUpDown *Mib_InterfaceMib_Subsets_Subset_LinkUpDown) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkUpDown *Mib_InterfaceMib_Subsets_Subset_LinkUpDown) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkUpDown *Mib_InterfaceMib_Subsets_Subset_LinkUpDown) SetParent(parent types.Entity) { linkUpDown.parent = parent }

func (linkUpDown *Mib_InterfaceMib_Subsets_Subset_LinkUpDown) GetParent() types.Entity { return linkUpDown.parent }

func (linkUpDown *Mib_InterfaceMib_Subsets_Subset_LinkUpDown) GetParentYangName() string { return "subset" }

// Mib_Subscriber
// Subscriber threshold commands
type Mib_Subscriber struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subscriber threshold commands.
    Threshold Mib_Subscriber_Threshold
}

func (subscriber *Mib_Subscriber) GetFilter() yfilter.YFilter { return subscriber.YFilter }

func (subscriber *Mib_Subscriber) SetFilter(yf yfilter.YFilter) { subscriber.YFilter = yf }

func (subscriber *Mib_Subscriber) GetGoName(yname string) string {
    if yname == "threshold" { return "Threshold" }
    return ""
}

func (subscriber *Mib_Subscriber) GetSegmentPath() string {
    return "Cisco-IOS-XR-subscriber-session-mon-mibs-cfg:subscriber"
}

func (subscriber *Mib_Subscriber) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "threshold" {
        return &subscriber.Threshold
    }
    return nil
}

func (subscriber *Mib_Subscriber) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["threshold"] = &subscriber.Threshold
    return children
}

func (subscriber *Mib_Subscriber) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subscriber *Mib_Subscriber) GetBundleName() string { return "cisco_ios_xr" }

func (subscriber *Mib_Subscriber) GetYangName() string { return "subscriber" }

func (subscriber *Mib_Subscriber) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subscriber *Mib_Subscriber) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subscriber *Mib_Subscriber) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subscriber *Mib_Subscriber) SetParent(parent types.Entity) { subscriber.parent = parent }

func (subscriber *Mib_Subscriber) GetParent() types.Entity { return subscriber.parent }

func (subscriber *Mib_Subscriber) GetParentYangName() string { return "mib" }

// Mib_Subscriber_Threshold
// Subscriber threshold commands
type Mib_Subscriber_Threshold struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Delta loss keyword.
    Delta Mib_Subscriber_Threshold_Delta

    // Access interface for regular expression.
    AccessInterfaceSub Mib_Subscriber_Threshold_AccessInterfaceSub

    // Falling threshold.
    Falling Mib_Subscriber_Threshold_Falling

    // Rising threshold.
    Rising Mib_Subscriber_Threshold_Rising
}

func (threshold *Mib_Subscriber_Threshold) GetFilter() yfilter.YFilter { return threshold.YFilter }

func (threshold *Mib_Subscriber_Threshold) SetFilter(yf yfilter.YFilter) { threshold.YFilter = yf }

func (threshold *Mib_Subscriber_Threshold) GetGoName(yname string) string {
    if yname == "delta" { return "Delta" }
    if yname == "access-interface-sub" { return "AccessInterfaceSub" }
    if yname == "falling" { return "Falling" }
    if yname == "rising" { return "Rising" }
    return ""
}

func (threshold *Mib_Subscriber_Threshold) GetSegmentPath() string {
    return "threshold"
}

func (threshold *Mib_Subscriber_Threshold) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "delta" {
        return &threshold.Delta
    }
    if childYangName == "access-interface-sub" {
        return &threshold.AccessInterfaceSub
    }
    if childYangName == "falling" {
        return &threshold.Falling
    }
    if childYangName == "rising" {
        return &threshold.Rising
    }
    return nil
}

func (threshold *Mib_Subscriber_Threshold) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["delta"] = &threshold.Delta
    children["access-interface-sub"] = &threshold.AccessInterfaceSub
    children["falling"] = &threshold.Falling
    children["rising"] = &threshold.Rising
    return children
}

func (threshold *Mib_Subscriber_Threshold) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (threshold *Mib_Subscriber_Threshold) GetBundleName() string { return "cisco_ios_xr" }

func (threshold *Mib_Subscriber_Threshold) GetYangName() string { return "threshold" }

func (threshold *Mib_Subscriber_Threshold) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (threshold *Mib_Subscriber_Threshold) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (threshold *Mib_Subscriber_Threshold) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (threshold *Mib_Subscriber_Threshold) SetParent(parent types.Entity) { threshold.parent = parent }

func (threshold *Mib_Subscriber_Threshold) GetParent() types.Entity { return threshold.parent }

func (threshold *Mib_Subscriber_Threshold) GetParentYangName() string { return "subscriber" }

// Mib_Subscriber_Threshold_Delta
// Delta loss keyword
type Mib_Subscriber_Threshold_Delta struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Evaluation keyword.
    Evaluation Mib_Subscriber_Threshold_Delta_Evaluation

    // Delta loss percent.
    Percent Mib_Subscriber_Threshold_Delta_Percent
}

func (delta *Mib_Subscriber_Threshold_Delta) GetFilter() yfilter.YFilter { return delta.YFilter }

func (delta *Mib_Subscriber_Threshold_Delta) SetFilter(yf yfilter.YFilter) { delta.YFilter = yf }

func (delta *Mib_Subscriber_Threshold_Delta) GetGoName(yname string) string {
    if yname == "evaluation" { return "Evaluation" }
    if yname == "percent" { return "Percent" }
    return ""
}

func (delta *Mib_Subscriber_Threshold_Delta) GetSegmentPath() string {
    return "delta"
}

func (delta *Mib_Subscriber_Threshold_Delta) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "evaluation" {
        return &delta.Evaluation
    }
    if childYangName == "percent" {
        return &delta.Percent
    }
    return nil
}

func (delta *Mib_Subscriber_Threshold_Delta) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["evaluation"] = &delta.Evaluation
    children["percent"] = &delta.Percent
    return children
}

func (delta *Mib_Subscriber_Threshold_Delta) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (delta *Mib_Subscriber_Threshold_Delta) GetBundleName() string { return "cisco_ios_xr" }

func (delta *Mib_Subscriber_Threshold_Delta) GetYangName() string { return "delta" }

func (delta *Mib_Subscriber_Threshold_Delta) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delta *Mib_Subscriber_Threshold_Delta) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delta *Mib_Subscriber_Threshold_Delta) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delta *Mib_Subscriber_Threshold_Delta) SetParent(parent types.Entity) { delta.parent = parent }

func (delta *Mib_Subscriber_Threshold_Delta) GetParent() types.Entity { return delta.parent }

func (delta *Mib_Subscriber_Threshold_Delta) GetParentYangName() string { return "threshold" }

// Mib_Subscriber_Threshold_Delta_Evaluation
// Evaluation keyword
type Mib_Subscriber_Threshold_Delta_Evaluation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of AccessInterface.
    AccessInterfaces Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces

    // Table of Node.
    Nodes Mib_Subscriber_Threshold_Delta_Evaluation_Nodes
}

func (evaluation *Mib_Subscriber_Threshold_Delta_Evaluation) GetFilter() yfilter.YFilter { return evaluation.YFilter }

func (evaluation *Mib_Subscriber_Threshold_Delta_Evaluation) SetFilter(yf yfilter.YFilter) { evaluation.YFilter = yf }

func (evaluation *Mib_Subscriber_Threshold_Delta_Evaluation) GetGoName(yname string) string {
    if yname == "access-interfaces" { return "AccessInterfaces" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (evaluation *Mib_Subscriber_Threshold_Delta_Evaluation) GetSegmentPath() string {
    return "evaluation"
}

func (evaluation *Mib_Subscriber_Threshold_Delta_Evaluation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-interfaces" {
        return &evaluation.AccessInterfaces
    }
    if childYangName == "nodes" {
        return &evaluation.Nodes
    }
    return nil
}

func (evaluation *Mib_Subscriber_Threshold_Delta_Evaluation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-interfaces"] = &evaluation.AccessInterfaces
    children["nodes"] = &evaluation.Nodes
    return children
}

func (evaluation *Mib_Subscriber_Threshold_Delta_Evaluation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (evaluation *Mib_Subscriber_Threshold_Delta_Evaluation) GetBundleName() string { return "cisco_ios_xr" }

func (evaluation *Mib_Subscriber_Threshold_Delta_Evaluation) GetYangName() string { return "evaluation" }

func (evaluation *Mib_Subscriber_Threshold_Delta_Evaluation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (evaluation *Mib_Subscriber_Threshold_Delta_Evaluation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (evaluation *Mib_Subscriber_Threshold_Delta_Evaluation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (evaluation *Mib_Subscriber_Threshold_Delta_Evaluation) SetParent(parent types.Entity) { evaluation.parent = parent }

func (evaluation *Mib_Subscriber_Threshold_Delta_Evaluation) GetParent() types.Entity { return evaluation.parent }

func (evaluation *Mib_Subscriber_Threshold_Delta_Evaluation) GetParentYangName() string { return "delta" }

// Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces
// Table of AccessInterface
type Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access interface. The type is slice of
    // Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface.
    AccessInterface []Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface
}

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces) GetFilter() yfilter.YFilter { return accessInterfaces.YFilter }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces) SetFilter(yf yfilter.YFilter) { accessInterfaces.YFilter = yf }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces) GetGoName(yname string) string {
    if yname == "access-interface" { return "AccessInterface" }
    return ""
}

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces) GetSegmentPath() string {
    return "access-interfaces"
}

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-interface" {
        for _, c := range accessInterfaces.AccessInterface {
            if accessInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface{}
        accessInterfaces.AccessInterface = append(accessInterfaces.AccessInterface, child)
        return &accessInterfaces.AccessInterface[len(accessInterfaces.AccessInterface)-1]
    }
    return nil
}

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessInterfaces.AccessInterface {
        children[accessInterfaces.AccessInterface[i].GetSegmentPath()] = &accessInterfaces.AccessInterface[i]
    }
    return children
}

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces) GetYangName() string { return "access-interfaces" }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces) SetParent(parent types.Entity) { accessInterfaces.parent = parent }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces) GetParent() types.Entity { return accessInterfaces.parent }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces) GetParentYangName() string { return "evaluation" }

// Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface
// Access interface
type Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Session count. The type is interface{} with range: 1..4294967294.
    SessionCount interface{}

    // Interval value in multiples of 10. The type is interface{} with range:
    // 30..3600.
    Interval interface{}
}

func (accessInterface *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface) GetFilter() yfilter.YFilter { return accessInterface.YFilter }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface) SetFilter(yf yfilter.YFilter) { accessInterface.YFilter = yf }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "session-count" { return "SessionCount" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (accessInterface *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface) GetSegmentPath() string {
    return "access-interface" + "[interface-name='" + fmt.Sprintf("%v", accessInterface.InterfaceName) + "']"
}

func (accessInterface *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accessInterface *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accessInterface *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = accessInterface.InterfaceName
    leafs["session-count"] = accessInterface.SessionCount
    leafs["interval"] = accessInterface.Interval
    return leafs
}

func (accessInterface *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface) GetBundleName() string { return "cisco_ios_xr" }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface) GetYangName() string { return "access-interface" }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface) SetParent(parent types.Entity) { accessInterface.parent = parent }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface) GetParent() types.Entity { return accessInterface.parent }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface) GetParentYangName() string { return "access-interfaces" }

// Mib_Subscriber_Threshold_Delta_Evaluation_Nodes
// Table of Node
type Mib_Subscriber_Threshold_Delta_Evaluation_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rising node level. The type is slice of
    // Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node.
    Node []Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node
}

func (nodes *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes) GetYangName() string { return "nodes" }

func (nodes *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes) GetParentYangName() string { return "evaluation" }

// Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node
// Rising node level
type Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Session count. The type is interface{} with range: 1..4294967294.
    SessionCount interface{}

    // interval value in multiples of 10. The type is interface{} with range:
    // 30..3600.
    Interval interface{}
}

func (node *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "session-count" { return "SessionCount" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (node *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    leafs["session-count"] = node.SessionCount
    leafs["interval"] = node.Interval
    return leafs
}

func (node *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node) GetYangName() string { return "node" }

func (node *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node) GetParentYangName() string { return "nodes" }

// Mib_Subscriber_Threshold_Delta_Percent
// Delta loss percent
type Mib_Subscriber_Threshold_Delta_Percent struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of AccessInterface.
    AccessInterfaces Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces

    // Table of Node.
    Nodes Mib_Subscriber_Threshold_Delta_Percent_Nodes
}

func (percent *Mib_Subscriber_Threshold_Delta_Percent) GetFilter() yfilter.YFilter { return percent.YFilter }

func (percent *Mib_Subscriber_Threshold_Delta_Percent) SetFilter(yf yfilter.YFilter) { percent.YFilter = yf }

func (percent *Mib_Subscriber_Threshold_Delta_Percent) GetGoName(yname string) string {
    if yname == "access-interfaces" { return "AccessInterfaces" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (percent *Mib_Subscriber_Threshold_Delta_Percent) GetSegmentPath() string {
    return "percent"
}

func (percent *Mib_Subscriber_Threshold_Delta_Percent) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-interfaces" {
        return &percent.AccessInterfaces
    }
    if childYangName == "nodes" {
        return &percent.Nodes
    }
    return nil
}

func (percent *Mib_Subscriber_Threshold_Delta_Percent) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-interfaces"] = &percent.AccessInterfaces
    children["nodes"] = &percent.Nodes
    return children
}

func (percent *Mib_Subscriber_Threshold_Delta_Percent) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (percent *Mib_Subscriber_Threshold_Delta_Percent) GetBundleName() string { return "cisco_ios_xr" }

func (percent *Mib_Subscriber_Threshold_Delta_Percent) GetYangName() string { return "percent" }

func (percent *Mib_Subscriber_Threshold_Delta_Percent) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (percent *Mib_Subscriber_Threshold_Delta_Percent) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (percent *Mib_Subscriber_Threshold_Delta_Percent) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (percent *Mib_Subscriber_Threshold_Delta_Percent) SetParent(parent types.Entity) { percent.parent = parent }

func (percent *Mib_Subscriber_Threshold_Delta_Percent) GetParent() types.Entity { return percent.parent }

func (percent *Mib_Subscriber_Threshold_Delta_Percent) GetParentYangName() string { return "delta" }

// Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces
// Table of AccessInterface
type Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access interface. The type is slice of
    // Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface.
    AccessInterface []Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface
}

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces) GetFilter() yfilter.YFilter { return accessInterfaces.YFilter }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces) SetFilter(yf yfilter.YFilter) { accessInterfaces.YFilter = yf }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces) GetGoName(yname string) string {
    if yname == "access-interface" { return "AccessInterface" }
    return ""
}

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces) GetSegmentPath() string {
    return "access-interfaces"
}

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-interface" {
        for _, c := range accessInterfaces.AccessInterface {
            if accessInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface{}
        accessInterfaces.AccessInterface = append(accessInterfaces.AccessInterface, child)
        return &accessInterfaces.AccessInterface[len(accessInterfaces.AccessInterface)-1]
    }
    return nil
}

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessInterfaces.AccessInterface {
        children[accessInterfaces.AccessInterface[i].GetSegmentPath()] = &accessInterfaces.AccessInterface[i]
    }
    return children
}

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces) GetYangName() string { return "access-interfaces" }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces) SetParent(parent types.Entity) { accessInterfaces.parent = parent }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces) GetParent() types.Entity { return accessInterfaces.parent }

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces) GetParentYangName() string { return "percent" }

// Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface
// Access interface
type Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Session count. The type is interface{} with range: 1..4294967294.
    SessionCount interface{}

    // Interval value in multiples of 10. The type is interface{} with range:
    // 30..3600.
    Interval interface{}
}

func (accessInterface *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface) GetFilter() yfilter.YFilter { return accessInterface.YFilter }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface) SetFilter(yf yfilter.YFilter) { accessInterface.YFilter = yf }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "session-count" { return "SessionCount" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (accessInterface *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface) GetSegmentPath() string {
    return "access-interface" + "[interface-name='" + fmt.Sprintf("%v", accessInterface.InterfaceName) + "']"
}

func (accessInterface *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accessInterface *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accessInterface *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = accessInterface.InterfaceName
    leafs["session-count"] = accessInterface.SessionCount
    leafs["interval"] = accessInterface.Interval
    return leafs
}

func (accessInterface *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface) GetBundleName() string { return "cisco_ios_xr" }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface) GetYangName() string { return "access-interface" }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface) SetParent(parent types.Entity) { accessInterface.parent = parent }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface) GetParent() types.Entity { return accessInterface.parent }

func (accessInterface *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface) GetParentYangName() string { return "access-interfaces" }

// Mib_Subscriber_Threshold_Delta_Percent_Nodes
// Table of Node
type Mib_Subscriber_Threshold_Delta_Percent_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rising node level. The type is slice of
    // Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node.
    Node []Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node
}

func (nodes *Mib_Subscriber_Threshold_Delta_Percent_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Mib_Subscriber_Threshold_Delta_Percent_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Mib_Subscriber_Threshold_Delta_Percent_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Mib_Subscriber_Threshold_Delta_Percent_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Mib_Subscriber_Threshold_Delta_Percent_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Mib_Subscriber_Threshold_Delta_Percent_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Mib_Subscriber_Threshold_Delta_Percent_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Mib_Subscriber_Threshold_Delta_Percent_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Mib_Subscriber_Threshold_Delta_Percent_Nodes) GetYangName() string { return "nodes" }

func (nodes *Mib_Subscriber_Threshold_Delta_Percent_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Mib_Subscriber_Threshold_Delta_Percent_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Mib_Subscriber_Threshold_Delta_Percent_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Mib_Subscriber_Threshold_Delta_Percent_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Mib_Subscriber_Threshold_Delta_Percent_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Mib_Subscriber_Threshold_Delta_Percent_Nodes) GetParentYangName() string { return "percent" }

// Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node
// Rising node level
type Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Session count. The type is interface{} with range: 1..4294967294.
    SessionCount interface{}

    // interval value in multiples of 10. The type is interface{} with range:
    // 30..3600.
    Interval interface{}
}

func (node *Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "session-count" { return "SessionCount" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (node *Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    leafs["session-count"] = node.SessionCount
    leafs["interval"] = node.Interval
    return leafs
}

func (node *Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node) GetYangName() string { return "node" }

func (node *Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node) GetParentYangName() string { return "nodes" }

// Mib_Subscriber_Threshold_AccessInterfaceSub
// Access interface for regular expression
type Mib_Subscriber_Threshold_AccessInterfaceSub struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of Subset.
    Subsets Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets
}

func (accessInterfaceSub *Mib_Subscriber_Threshold_AccessInterfaceSub) GetFilter() yfilter.YFilter { return accessInterfaceSub.YFilter }

func (accessInterfaceSub *Mib_Subscriber_Threshold_AccessInterfaceSub) SetFilter(yf yfilter.YFilter) { accessInterfaceSub.YFilter = yf }

func (accessInterfaceSub *Mib_Subscriber_Threshold_AccessInterfaceSub) GetGoName(yname string) string {
    if yname == "subsets" { return "Subsets" }
    return ""
}

func (accessInterfaceSub *Mib_Subscriber_Threshold_AccessInterfaceSub) GetSegmentPath() string {
    return "access-interface-sub"
}

func (accessInterfaceSub *Mib_Subscriber_Threshold_AccessInterfaceSub) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "subsets" {
        return &accessInterfaceSub.Subsets
    }
    return nil
}

func (accessInterfaceSub *Mib_Subscriber_Threshold_AccessInterfaceSub) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["subsets"] = &accessInterfaceSub.Subsets
    return children
}

func (accessInterfaceSub *Mib_Subscriber_Threshold_AccessInterfaceSub) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessInterfaceSub *Mib_Subscriber_Threshold_AccessInterfaceSub) GetBundleName() string { return "cisco_ios_xr" }

func (accessInterfaceSub *Mib_Subscriber_Threshold_AccessInterfaceSub) GetYangName() string { return "access-interface-sub" }

func (accessInterfaceSub *Mib_Subscriber_Threshold_AccessInterfaceSub) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessInterfaceSub *Mib_Subscriber_Threshold_AccessInterfaceSub) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessInterfaceSub *Mib_Subscriber_Threshold_AccessInterfaceSub) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessInterfaceSub *Mib_Subscriber_Threshold_AccessInterfaceSub) SetParent(parent types.Entity) { accessInterfaceSub.parent = parent }

func (accessInterfaceSub *Mib_Subscriber_Threshold_AccessInterfaceSub) GetParent() types.Entity { return accessInterfaceSub.parent }

func (accessInterfaceSub *Mib_Subscriber_Threshold_AccessInterfaceSub) GetParentYangName() string { return "threshold" }

// Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets
// Table of Subset
type Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Subset command. The type is slice of
    // Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset.
    Subset []Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset
}

func (subsets *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets) GetFilter() yfilter.YFilter { return subsets.YFilter }

func (subsets *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets) SetFilter(yf yfilter.YFilter) { subsets.YFilter = yf }

func (subsets *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets) GetGoName(yname string) string {
    if yname == "subset" { return "Subset" }
    return ""
}

func (subsets *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets) GetSegmentPath() string {
    return "subsets"
}

func (subsets *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "subset" {
        for _, c := range subsets.Subset {
            if subsets.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset{}
        subsets.Subset = append(subsets.Subset, child)
        return &subsets.Subset[len(subsets.Subset)-1]
    }
    return nil
}

func (subsets *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range subsets.Subset {
        children[subsets.Subset[i].GetSegmentPath()] = &subsets.Subset[i]
    }
    return children
}

func (subsets *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (subsets *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets) GetBundleName() string { return "cisco_ios_xr" }

func (subsets *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets) GetYangName() string { return "subsets" }

func (subsets *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subsets *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subsets *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subsets *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets) SetParent(parent types.Entity) { subsets.parent = parent }

func (subsets *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets) GetParent() types.Entity { return subsets.parent }

func (subsets *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets) GetParentYangName() string { return "access-interface-sub" }

// Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset
// Subset command
type Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Subset number. The type is interface{} with range:
    // 1..255.
    SubsetId interface{}

    // Regular expression.
    RegularExpression Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression
}

func (subset *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset) GetFilter() yfilter.YFilter { return subset.YFilter }

func (subset *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset) SetFilter(yf yfilter.YFilter) { subset.YFilter = yf }

func (subset *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset) GetGoName(yname string) string {
    if yname == "subset-id" { return "SubsetId" }
    if yname == "regular-expression" { return "RegularExpression" }
    return ""
}

func (subset *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset) GetSegmentPath() string {
    return "subset" + "[subset-id='" + fmt.Sprintf("%v", subset.SubsetId) + "']"
}

func (subset *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "regular-expression" {
        return &subset.RegularExpression
    }
    return nil
}

func (subset *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["regular-expression"] = &subset.RegularExpression
    return children
}

func (subset *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subset-id"] = subset.SubsetId
    return leafs
}

func (subset *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset) GetBundleName() string { return "cisco_ios_xr" }

func (subset *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset) GetYangName() string { return "subset" }

func (subset *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (subset *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (subset *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (subset *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset) SetParent(parent types.Entity) { subset.parent = parent }

func (subset *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset) GetParent() types.Entity { return subset.parent }

func (subset *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset) GetParentYangName() string { return "subsets" }

// Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression
// Regular expression
type Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Notification keyword.
    Notification Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification
}

func (regularExpression *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression) GetFilter() yfilter.YFilter { return regularExpression.YFilter }

func (regularExpression *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression) SetFilter(yf yfilter.YFilter) { regularExpression.YFilter = yf }

func (regularExpression *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression) GetGoName(yname string) string {
    if yname == "notification" { return "Notification" }
    return ""
}

func (regularExpression *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression) GetSegmentPath() string {
    return "regular-expression"
}

func (regularExpression *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "notification" {
        return &regularExpression.Notification
    }
    return nil
}

func (regularExpression *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["notification"] = &regularExpression.Notification
    return children
}

func (regularExpression *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (regularExpression *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression) GetBundleName() string { return "cisco_ios_xr" }

func (regularExpression *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression) GetYangName() string { return "regular-expression" }

func (regularExpression *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (regularExpression *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (regularExpression *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (regularExpression *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression) SetParent(parent types.Entity) { regularExpression.parent = parent }

func (regularExpression *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression) GetParent() types.Entity { return regularExpression.parent }

func (regularExpression *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression) GetParentYangName() string { return "subset" }

// Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification
// Notification keyword
type Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rising-falling threshold.
    RisingFalling Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling
}

func (notification *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification) GetFilter() yfilter.YFilter { return notification.YFilter }

func (notification *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification) SetFilter(yf yfilter.YFilter) { notification.YFilter = yf }

func (notification *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification) GetGoName(yname string) string {
    if yname == "rising-falling" { return "RisingFalling" }
    return ""
}

func (notification *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification) GetSegmentPath() string {
    return "notification"
}

func (notification *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rising-falling" {
        return &notification.RisingFalling
    }
    return nil
}

func (notification *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rising-falling"] = &notification.RisingFalling
    return children
}

func (notification *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (notification *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification) GetBundleName() string { return "cisco_ios_xr" }

func (notification *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification) GetYangName() string { return "notification" }

func (notification *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (notification *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (notification *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (notification *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification) SetParent(parent types.Entity) { notification.parent = parent }

func (notification *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification) GetParent() types.Entity { return notification.parent }

func (notification *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification) GetParentYangName() string { return "regular-expression" }

// Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling
// Rising-falling threshold
type Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Disable the notifications on access interfaces. The type is string.
    Disable interface{}
}

func (risingFalling *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling) GetFilter() yfilter.YFilter { return risingFalling.YFilter }

func (risingFalling *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling) SetFilter(yf yfilter.YFilter) { risingFalling.YFilter = yf }

func (risingFalling *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling) GetGoName(yname string) string {
    if yname == "disable" { return "Disable" }
    return ""
}

func (risingFalling *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling) GetSegmentPath() string {
    return "rising-falling"
}

func (risingFalling *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (risingFalling *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (risingFalling *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disable"] = risingFalling.Disable
    return leafs
}

func (risingFalling *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling) GetBundleName() string { return "cisco_ios_xr" }

func (risingFalling *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling) GetYangName() string { return "rising-falling" }

func (risingFalling *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (risingFalling *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (risingFalling *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (risingFalling *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling) SetParent(parent types.Entity) { risingFalling.parent = parent }

func (risingFalling *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling) GetParent() types.Entity { return risingFalling.parent }

func (risingFalling *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling) GetParentYangName() string { return "notification" }

// Mib_Subscriber_Threshold_Falling
// Falling threshold
type Mib_Subscriber_Threshold_Falling struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of AccessInterface.
    AccessInterfaces Mib_Subscriber_Threshold_Falling_AccessInterfaces

    // Table of Node.
    Nodes Mib_Subscriber_Threshold_Falling_Nodes
}

func (falling *Mib_Subscriber_Threshold_Falling) GetFilter() yfilter.YFilter { return falling.YFilter }

func (falling *Mib_Subscriber_Threshold_Falling) SetFilter(yf yfilter.YFilter) { falling.YFilter = yf }

func (falling *Mib_Subscriber_Threshold_Falling) GetGoName(yname string) string {
    if yname == "access-interfaces" { return "AccessInterfaces" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (falling *Mib_Subscriber_Threshold_Falling) GetSegmentPath() string {
    return "falling"
}

func (falling *Mib_Subscriber_Threshold_Falling) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-interfaces" {
        return &falling.AccessInterfaces
    }
    if childYangName == "nodes" {
        return &falling.Nodes
    }
    return nil
}

func (falling *Mib_Subscriber_Threshold_Falling) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-interfaces"] = &falling.AccessInterfaces
    children["nodes"] = &falling.Nodes
    return children
}

func (falling *Mib_Subscriber_Threshold_Falling) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (falling *Mib_Subscriber_Threshold_Falling) GetBundleName() string { return "cisco_ios_xr" }

func (falling *Mib_Subscriber_Threshold_Falling) GetYangName() string { return "falling" }

func (falling *Mib_Subscriber_Threshold_Falling) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (falling *Mib_Subscriber_Threshold_Falling) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (falling *Mib_Subscriber_Threshold_Falling) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (falling *Mib_Subscriber_Threshold_Falling) SetParent(parent types.Entity) { falling.parent = parent }

func (falling *Mib_Subscriber_Threshold_Falling) GetParent() types.Entity { return falling.parent }

func (falling *Mib_Subscriber_Threshold_Falling) GetParentYangName() string { return "threshold" }

// Mib_Subscriber_Threshold_Falling_AccessInterfaces
// Table of AccessInterface
type Mib_Subscriber_Threshold_Falling_AccessInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access interface. The type is slice of
    // Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface.
    AccessInterface []Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface
}

func (accessInterfaces *Mib_Subscriber_Threshold_Falling_AccessInterfaces) GetFilter() yfilter.YFilter { return accessInterfaces.YFilter }

func (accessInterfaces *Mib_Subscriber_Threshold_Falling_AccessInterfaces) SetFilter(yf yfilter.YFilter) { accessInterfaces.YFilter = yf }

func (accessInterfaces *Mib_Subscriber_Threshold_Falling_AccessInterfaces) GetGoName(yname string) string {
    if yname == "access-interface" { return "AccessInterface" }
    return ""
}

func (accessInterfaces *Mib_Subscriber_Threshold_Falling_AccessInterfaces) GetSegmentPath() string {
    return "access-interfaces"
}

func (accessInterfaces *Mib_Subscriber_Threshold_Falling_AccessInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-interface" {
        for _, c := range accessInterfaces.AccessInterface {
            if accessInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface{}
        accessInterfaces.AccessInterface = append(accessInterfaces.AccessInterface, child)
        return &accessInterfaces.AccessInterface[len(accessInterfaces.AccessInterface)-1]
    }
    return nil
}

func (accessInterfaces *Mib_Subscriber_Threshold_Falling_AccessInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessInterfaces.AccessInterface {
        children[accessInterfaces.AccessInterface[i].GetSegmentPath()] = &accessInterfaces.AccessInterface[i]
    }
    return children
}

func (accessInterfaces *Mib_Subscriber_Threshold_Falling_AccessInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessInterfaces *Mib_Subscriber_Threshold_Falling_AccessInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (accessInterfaces *Mib_Subscriber_Threshold_Falling_AccessInterfaces) GetYangName() string { return "access-interfaces" }

func (accessInterfaces *Mib_Subscriber_Threshold_Falling_AccessInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessInterfaces *Mib_Subscriber_Threshold_Falling_AccessInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessInterfaces *Mib_Subscriber_Threshold_Falling_AccessInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessInterfaces *Mib_Subscriber_Threshold_Falling_AccessInterfaces) SetParent(parent types.Entity) { accessInterfaces.parent = parent }

func (accessInterfaces *Mib_Subscriber_Threshold_Falling_AccessInterfaces) GetParent() types.Entity { return accessInterfaces.parent }

func (accessInterfaces *Mib_Subscriber_Threshold_Falling_AccessInterfaces) GetParentYangName() string { return "falling" }

// Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface
// Access interface
type Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Session count. The type is interface{} with range: 1..4294967294.
    SessionCount interface{}

    // Interval value in multiples of 10. The type is interface{} with range:
    // 30..3600.
    Interval interface{}
}

func (accessInterface *Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface) GetFilter() yfilter.YFilter { return accessInterface.YFilter }

func (accessInterface *Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface) SetFilter(yf yfilter.YFilter) { accessInterface.YFilter = yf }

func (accessInterface *Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "session-count" { return "SessionCount" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (accessInterface *Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface) GetSegmentPath() string {
    return "access-interface" + "[interface-name='" + fmt.Sprintf("%v", accessInterface.InterfaceName) + "']"
}

func (accessInterface *Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accessInterface *Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accessInterface *Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = accessInterface.InterfaceName
    leafs["session-count"] = accessInterface.SessionCount
    leafs["interval"] = accessInterface.Interval
    return leafs
}

func (accessInterface *Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface) GetBundleName() string { return "cisco_ios_xr" }

func (accessInterface *Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface) GetYangName() string { return "access-interface" }

func (accessInterface *Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessInterface *Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessInterface *Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessInterface *Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface) SetParent(parent types.Entity) { accessInterface.parent = parent }

func (accessInterface *Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface) GetParent() types.Entity { return accessInterface.parent }

func (accessInterface *Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface) GetParentYangName() string { return "access-interfaces" }

// Mib_Subscriber_Threshold_Falling_Nodes
// Table of Node
type Mib_Subscriber_Threshold_Falling_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rising node level. The type is slice of
    // Mib_Subscriber_Threshold_Falling_Nodes_Node.
    Node []Mib_Subscriber_Threshold_Falling_Nodes_Node
}

func (nodes *Mib_Subscriber_Threshold_Falling_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Mib_Subscriber_Threshold_Falling_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Mib_Subscriber_Threshold_Falling_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Mib_Subscriber_Threshold_Falling_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Mib_Subscriber_Threshold_Falling_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mib_Subscriber_Threshold_Falling_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Mib_Subscriber_Threshold_Falling_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Mib_Subscriber_Threshold_Falling_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Mib_Subscriber_Threshold_Falling_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Mib_Subscriber_Threshold_Falling_Nodes) GetYangName() string { return "nodes" }

func (nodes *Mib_Subscriber_Threshold_Falling_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Mib_Subscriber_Threshold_Falling_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Mib_Subscriber_Threshold_Falling_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Mib_Subscriber_Threshold_Falling_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Mib_Subscriber_Threshold_Falling_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Mib_Subscriber_Threshold_Falling_Nodes) GetParentYangName() string { return "falling" }

// Mib_Subscriber_Threshold_Falling_Nodes_Node
// Rising node level
type Mib_Subscriber_Threshold_Falling_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Session count. The type is interface{} with range: 1..4294967294.
    SessionCount interface{}

    // interval value in multiples of 10. The type is interface{} with range:
    // 30..3600.
    Interval interface{}
}

func (node *Mib_Subscriber_Threshold_Falling_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Mib_Subscriber_Threshold_Falling_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Mib_Subscriber_Threshold_Falling_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "session-count" { return "SessionCount" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (node *Mib_Subscriber_Threshold_Falling_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Mib_Subscriber_Threshold_Falling_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *Mib_Subscriber_Threshold_Falling_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *Mib_Subscriber_Threshold_Falling_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    leafs["session-count"] = node.SessionCount
    leafs["interval"] = node.Interval
    return leafs
}

func (node *Mib_Subscriber_Threshold_Falling_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Mib_Subscriber_Threshold_Falling_Nodes_Node) GetYangName() string { return "node" }

func (node *Mib_Subscriber_Threshold_Falling_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Mib_Subscriber_Threshold_Falling_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Mib_Subscriber_Threshold_Falling_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Mib_Subscriber_Threshold_Falling_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Mib_Subscriber_Threshold_Falling_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Mib_Subscriber_Threshold_Falling_Nodes_Node) GetParentYangName() string { return "nodes" }

// Mib_Subscriber_Threshold_Rising
// Rising threshold
type Mib_Subscriber_Threshold_Rising struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of AccessInterface.
    AccessInterfaces Mib_Subscriber_Threshold_Rising_AccessInterfaces

    // Table of Node.
    Nodes Mib_Subscriber_Threshold_Rising_Nodes
}

func (rising *Mib_Subscriber_Threshold_Rising) GetFilter() yfilter.YFilter { return rising.YFilter }

func (rising *Mib_Subscriber_Threshold_Rising) SetFilter(yf yfilter.YFilter) { rising.YFilter = yf }

func (rising *Mib_Subscriber_Threshold_Rising) GetGoName(yname string) string {
    if yname == "access-interfaces" { return "AccessInterfaces" }
    if yname == "nodes" { return "Nodes" }
    return ""
}

func (rising *Mib_Subscriber_Threshold_Rising) GetSegmentPath() string {
    return "rising"
}

func (rising *Mib_Subscriber_Threshold_Rising) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-interfaces" {
        return &rising.AccessInterfaces
    }
    if childYangName == "nodes" {
        return &rising.Nodes
    }
    return nil
}

func (rising *Mib_Subscriber_Threshold_Rising) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["access-interfaces"] = &rising.AccessInterfaces
    children["nodes"] = &rising.Nodes
    return children
}

func (rising *Mib_Subscriber_Threshold_Rising) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rising *Mib_Subscriber_Threshold_Rising) GetBundleName() string { return "cisco_ios_xr" }

func (rising *Mib_Subscriber_Threshold_Rising) GetYangName() string { return "rising" }

func (rising *Mib_Subscriber_Threshold_Rising) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rising *Mib_Subscriber_Threshold_Rising) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rising *Mib_Subscriber_Threshold_Rising) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rising *Mib_Subscriber_Threshold_Rising) SetParent(parent types.Entity) { rising.parent = parent }

func (rising *Mib_Subscriber_Threshold_Rising) GetParent() types.Entity { return rising.parent }

func (rising *Mib_Subscriber_Threshold_Rising) GetParentYangName() string { return "threshold" }

// Mib_Subscriber_Threshold_Rising_AccessInterfaces
// Table of AccessInterface
type Mib_Subscriber_Threshold_Rising_AccessInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Access interface. The type is slice of
    // Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface.
    AccessInterface []Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface
}

func (accessInterfaces *Mib_Subscriber_Threshold_Rising_AccessInterfaces) GetFilter() yfilter.YFilter { return accessInterfaces.YFilter }

func (accessInterfaces *Mib_Subscriber_Threshold_Rising_AccessInterfaces) SetFilter(yf yfilter.YFilter) { accessInterfaces.YFilter = yf }

func (accessInterfaces *Mib_Subscriber_Threshold_Rising_AccessInterfaces) GetGoName(yname string) string {
    if yname == "access-interface" { return "AccessInterface" }
    return ""
}

func (accessInterfaces *Mib_Subscriber_Threshold_Rising_AccessInterfaces) GetSegmentPath() string {
    return "access-interfaces"
}

func (accessInterfaces *Mib_Subscriber_Threshold_Rising_AccessInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "access-interface" {
        for _, c := range accessInterfaces.AccessInterface {
            if accessInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface{}
        accessInterfaces.AccessInterface = append(accessInterfaces.AccessInterface, child)
        return &accessInterfaces.AccessInterface[len(accessInterfaces.AccessInterface)-1]
    }
    return nil
}

func (accessInterfaces *Mib_Subscriber_Threshold_Rising_AccessInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range accessInterfaces.AccessInterface {
        children[accessInterfaces.AccessInterface[i].GetSegmentPath()] = &accessInterfaces.AccessInterface[i]
    }
    return children
}

func (accessInterfaces *Mib_Subscriber_Threshold_Rising_AccessInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (accessInterfaces *Mib_Subscriber_Threshold_Rising_AccessInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (accessInterfaces *Mib_Subscriber_Threshold_Rising_AccessInterfaces) GetYangName() string { return "access-interfaces" }

func (accessInterfaces *Mib_Subscriber_Threshold_Rising_AccessInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessInterfaces *Mib_Subscriber_Threshold_Rising_AccessInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessInterfaces *Mib_Subscriber_Threshold_Rising_AccessInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessInterfaces *Mib_Subscriber_Threshold_Rising_AccessInterfaces) SetParent(parent types.Entity) { accessInterfaces.parent = parent }

func (accessInterfaces *Mib_Subscriber_Threshold_Rising_AccessInterfaces) GetParent() types.Entity { return accessInterfaces.parent }

func (accessInterfaces *Mib_Subscriber_Threshold_Rising_AccessInterfaces) GetParentYangName() string { return "rising" }

// Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface
// Access interface
type Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Session count. The type is interface{} with range: 1..4294967294.
    SessionCount interface{}

    // Interval value in multiples of 10. The type is interface{} with range:
    // 30..3600.
    Interval interface{}
}

func (accessInterface *Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface) GetFilter() yfilter.YFilter { return accessInterface.YFilter }

func (accessInterface *Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface) SetFilter(yf yfilter.YFilter) { accessInterface.YFilter = yf }

func (accessInterface *Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "session-count" { return "SessionCount" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (accessInterface *Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface) GetSegmentPath() string {
    return "access-interface" + "[interface-name='" + fmt.Sprintf("%v", accessInterface.InterfaceName) + "']"
}

func (accessInterface *Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (accessInterface *Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (accessInterface *Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = accessInterface.InterfaceName
    leafs["session-count"] = accessInterface.SessionCount
    leafs["interval"] = accessInterface.Interval
    return leafs
}

func (accessInterface *Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface) GetBundleName() string { return "cisco_ios_xr" }

func (accessInterface *Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface) GetYangName() string { return "access-interface" }

func (accessInterface *Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (accessInterface *Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (accessInterface *Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (accessInterface *Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface) SetParent(parent types.Entity) { accessInterface.parent = parent }

func (accessInterface *Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface) GetParent() types.Entity { return accessInterface.parent }

func (accessInterface *Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface) GetParentYangName() string { return "access-interfaces" }

// Mib_Subscriber_Threshold_Rising_Nodes
// Table of Node
type Mib_Subscriber_Threshold_Rising_Nodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rising node level. The type is slice of
    // Mib_Subscriber_Threshold_Rising_Nodes_Node.
    Node []Mib_Subscriber_Threshold_Rising_Nodes_Node
}

func (nodes *Mib_Subscriber_Threshold_Rising_Nodes) GetFilter() yfilter.YFilter { return nodes.YFilter }

func (nodes *Mib_Subscriber_Threshold_Rising_Nodes) SetFilter(yf yfilter.YFilter) { nodes.YFilter = yf }

func (nodes *Mib_Subscriber_Threshold_Rising_Nodes) GetGoName(yname string) string {
    if yname == "node" { return "Node" }
    return ""
}

func (nodes *Mib_Subscriber_Threshold_Rising_Nodes) GetSegmentPath() string {
    return "nodes"
}

func (nodes *Mib_Subscriber_Threshold_Rising_Nodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "node" {
        for _, c := range nodes.Node {
            if nodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mib_Subscriber_Threshold_Rising_Nodes_Node{}
        nodes.Node = append(nodes.Node, child)
        return &nodes.Node[len(nodes.Node)-1]
    }
    return nil
}

func (nodes *Mib_Subscriber_Threshold_Rising_Nodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range nodes.Node {
        children[nodes.Node[i].GetSegmentPath()] = &nodes.Node[i]
    }
    return children
}

func (nodes *Mib_Subscriber_Threshold_Rising_Nodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (nodes *Mib_Subscriber_Threshold_Rising_Nodes) GetBundleName() string { return "cisco_ios_xr" }

func (nodes *Mib_Subscriber_Threshold_Rising_Nodes) GetYangName() string { return "nodes" }

func (nodes *Mib_Subscriber_Threshold_Rising_Nodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (nodes *Mib_Subscriber_Threshold_Rising_Nodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (nodes *Mib_Subscriber_Threshold_Rising_Nodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (nodes *Mib_Subscriber_Threshold_Rising_Nodes) SetParent(parent types.Entity) { nodes.parent = parent }

func (nodes *Mib_Subscriber_Threshold_Rising_Nodes) GetParent() types.Entity { return nodes.parent }

func (nodes *Mib_Subscriber_Threshold_Rising_Nodes) GetParentYangName() string { return "rising" }

// Mib_Subscriber_Threshold_Rising_Nodes_Node
// Rising node level
type Mib_Subscriber_Threshold_Rising_Nodes_Node struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    NodeName interface{}

    // Session count. The type is interface{} with range: 1..4294967294.
    SessionCount interface{}

    // interval value in multiples of 10. The type is interface{} with range:
    // 30..3600.
    Interval interface{}
}

func (node *Mib_Subscriber_Threshold_Rising_Nodes_Node) GetFilter() yfilter.YFilter { return node.YFilter }

func (node *Mib_Subscriber_Threshold_Rising_Nodes_Node) SetFilter(yf yfilter.YFilter) { node.YFilter = yf }

func (node *Mib_Subscriber_Threshold_Rising_Nodes_Node) GetGoName(yname string) string {
    if yname == "node-name" { return "NodeName" }
    if yname == "session-count" { return "SessionCount" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (node *Mib_Subscriber_Threshold_Rising_Nodes_Node) GetSegmentPath() string {
    return "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
}

func (node *Mib_Subscriber_Threshold_Rising_Nodes_Node) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (node *Mib_Subscriber_Threshold_Rising_Nodes_Node) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (node *Mib_Subscriber_Threshold_Rising_Nodes_Node) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["node-name"] = node.NodeName
    leafs["session-count"] = node.SessionCount
    leafs["interval"] = node.Interval
    return leafs
}

func (node *Mib_Subscriber_Threshold_Rising_Nodes_Node) GetBundleName() string { return "cisco_ios_xr" }

func (node *Mib_Subscriber_Threshold_Rising_Nodes_Node) GetYangName() string { return "node" }

func (node *Mib_Subscriber_Threshold_Rising_Nodes_Node) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (node *Mib_Subscriber_Threshold_Rising_Nodes_Node) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (node *Mib_Subscriber_Threshold_Rising_Nodes_Node) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (node *Mib_Subscriber_Threshold_Rising_Nodes_Node) SetParent(parent types.Entity) { node.parent = parent }

func (node *Mib_Subscriber_Threshold_Rising_Nodes_Node) GetParent() types.Entity { return node.parent }

func (node *Mib_Subscriber_Threshold_Rising_Nodes_Node) GetParentYangName() string { return "nodes" }

