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
    EntityData types.CommonEntityData
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

    // SNMP packet drop config.
    DropPacket Snmp_DropPacket

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

func (snmp *Snmp) GetEntityData() *types.CommonEntityData {
    snmp.EntityData.YFilter = snmp.YFilter
    snmp.EntityData.YangName = "snmp"
    snmp.EntityData.BundleName = "cisco_ios_xr"
    snmp.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-agent-cfg"
    snmp.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-agent-cfg:snmp"
    snmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmp.EntityData.Children = types.NewOrderedMap()
    snmp.EntityData.Children.Append("encrypted-community-maps", types.YChild{"EncryptedCommunityMaps", &snmp.EncryptedCommunityMaps})
    snmp.EntityData.Children.Append("views", types.YChild{"Views", &snmp.Views})
    snmp.EntityData.Children.Append("logging", types.YChild{"Logging", &snmp.Logging})
    snmp.EntityData.Children.Append("administration", types.YChild{"Administration", &snmp.Administration})
    snmp.EntityData.Children.Append("agent", types.YChild{"Agent", &snmp.Agent})
    snmp.EntityData.Children.Append("trap", types.YChild{"Trap", &snmp.Trap})
    snmp.EntityData.Children.Append("drop-packet", types.YChild{"DropPacket", &snmp.DropPacket})
    snmp.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &snmp.Ipv6})
    snmp.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &snmp.Ipv4})
    snmp.EntityData.Children.Append("system", types.YChild{"System", &snmp.System})
    snmp.EntityData.Children.Append("target", types.YChild{"Target", &snmp.Target})
    snmp.EntityData.Children.Append("notification", types.YChild{"Notification", &snmp.Notification})
    snmp.EntityData.Children.Append("correlator", types.YChild{"Correlator", &snmp.Correlator})
    snmp.EntityData.Children.Append("bulk-stats", types.YChild{"BulkStats", &snmp.BulkStats})
    snmp.EntityData.Children.Append("default-community-maps", types.YChild{"DefaultCommunityMaps", &snmp.DefaultCommunityMaps})
    snmp.EntityData.Children.Append("overload-control", types.YChild{"OverloadControl", &snmp.OverloadControl})
    snmp.EntityData.Children.Append("timeouts", types.YChild{"Timeouts", &snmp.Timeouts})
    snmp.EntityData.Children.Append("users", types.YChild{"Users", &snmp.Users})
    snmp.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &snmp.Vrfs})
    snmp.EntityData.Children.Append("groups", types.YChild{"Groups", &snmp.Groups})
    snmp.EntityData.Children.Append("trap-hosts", types.YChild{"TrapHosts", &snmp.TrapHosts})
    snmp.EntityData.Children.Append("contexts", types.YChild{"Contexts", &snmp.Contexts})
    snmp.EntityData.Children.Append("context-mappings", types.YChild{"ContextMappings", &snmp.ContextMappings})
    snmp.EntityData.Leafs = types.NewOrderedMap()
    snmp.EntityData.Leafs.Append("inform-retries", types.YLeaf{"InformRetries", snmp.InformRetries})
    snmp.EntityData.Leafs.Append("trap-port", types.YLeaf{"TrapPort", snmp.TrapPort})
    snmp.EntityData.Leafs.Append("oid-poll-stats", types.YLeaf{"OidPollStats", snmp.OidPollStats})
    snmp.EntityData.Leafs.Append("trap-source", types.YLeaf{"TrapSource", snmp.TrapSource})
    snmp.EntityData.Leafs.Append("vrf-authentication-trap-disable", types.YLeaf{"VrfAuthenticationTrapDisable", snmp.VrfAuthenticationTrapDisable})
    snmp.EntityData.Leafs.Append("inform-timeout", types.YLeaf{"InformTimeout", snmp.InformTimeout})
    snmp.EntityData.Leafs.Append("trap-source-ipv6", types.YLeaf{"TrapSourceIpv6", snmp.TrapSourceIpv6})
    snmp.EntityData.Leafs.Append("packet-size", types.YLeaf{"PacketSize", snmp.PacketSize})
    snmp.EntityData.Leafs.Append("throttle-time", types.YLeaf{"ThrottleTime", snmp.ThrottleTime})
    snmp.EntityData.Leafs.Append("trap-source-ipv4", types.YLeaf{"TrapSourceIpv4", snmp.TrapSourceIpv4})
    snmp.EntityData.Leafs.Append("inform-pending", types.YLeaf{"InformPending", snmp.InformPending})

    snmp.EntityData.YListKeys = []string {}

    return &(snmp.EntityData)
}

// Snmp_EncryptedCommunityMaps
// Container class to hold clear/encrypted
// communitie maps
type Snmp_EncryptedCommunityMaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear/encrypted SNMP community map. The type is slice of
    // Snmp_EncryptedCommunityMaps_EncryptedCommunityMap.
    EncryptedCommunityMap []*Snmp_EncryptedCommunityMaps_EncryptedCommunityMap
}

func (encryptedCommunityMaps *Snmp_EncryptedCommunityMaps) GetEntityData() *types.CommonEntityData {
    encryptedCommunityMaps.EntityData.YFilter = encryptedCommunityMaps.YFilter
    encryptedCommunityMaps.EntityData.YangName = "encrypted-community-maps"
    encryptedCommunityMaps.EntityData.BundleName = "cisco_ios_xr"
    encryptedCommunityMaps.EntityData.ParentYangName = "snmp"
    encryptedCommunityMaps.EntityData.SegmentPath = "encrypted-community-maps"
    encryptedCommunityMaps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryptedCommunityMaps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryptedCommunityMaps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryptedCommunityMaps.EntityData.Children = types.NewOrderedMap()
    encryptedCommunityMaps.EntityData.Children.Append("encrypted-community-map", types.YChild{"EncryptedCommunityMap", nil})
    for i := range encryptedCommunityMaps.EncryptedCommunityMap {
        encryptedCommunityMaps.EntityData.Children.Append(types.GetSegmentPath(encryptedCommunityMaps.EncryptedCommunityMap[i]), types.YChild{"EncryptedCommunityMap", encryptedCommunityMaps.EncryptedCommunityMap[i]})
    }
    encryptedCommunityMaps.EntityData.Leafs = types.NewOrderedMap()

    encryptedCommunityMaps.EntityData.YListKeys = []string {}

    return &(encryptedCommunityMaps.EntityData)
}

// Snmp_EncryptedCommunityMaps_EncryptedCommunityMap
// Clear/encrypted SNMP community map
type Snmp_EncryptedCommunityMaps_EncryptedCommunityMap struct {
    EntityData types.CommonEntityData
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

func (encryptedCommunityMap *Snmp_EncryptedCommunityMaps_EncryptedCommunityMap) GetEntityData() *types.CommonEntityData {
    encryptedCommunityMap.EntityData.YFilter = encryptedCommunityMap.YFilter
    encryptedCommunityMap.EntityData.YangName = "encrypted-community-map"
    encryptedCommunityMap.EntityData.BundleName = "cisco_ios_xr"
    encryptedCommunityMap.EntityData.ParentYangName = "encrypted-community-maps"
    encryptedCommunityMap.EntityData.SegmentPath = "encrypted-community-map" + types.AddKeyToken(encryptedCommunityMap.CommunityName, "community-name")
    encryptedCommunityMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryptedCommunityMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryptedCommunityMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryptedCommunityMap.EntityData.Children = types.NewOrderedMap()
    encryptedCommunityMap.EntityData.Leafs = types.NewOrderedMap()
    encryptedCommunityMap.EntityData.Leafs.Append("community-name", types.YLeaf{"CommunityName", encryptedCommunityMap.CommunityName})
    encryptedCommunityMap.EntityData.Leafs.Append("context", types.YLeaf{"Context", encryptedCommunityMap.Context})
    encryptedCommunityMap.EntityData.Leafs.Append("security", types.YLeaf{"Security", encryptedCommunityMap.Security})
    encryptedCommunityMap.EntityData.Leafs.Append("target-list", types.YLeaf{"TargetList", encryptedCommunityMap.TargetList})

    encryptedCommunityMap.EntityData.YListKeys = []string {"CommunityName"}

    return &(encryptedCommunityMap.EntityData)
}

// Snmp_Views
// Class to configure a SNMPv2 MIB view
type Snmp_Views struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the view. The type is slice of Snmp_Views_View.
    View []*Snmp_Views_View
}

func (views *Snmp_Views) GetEntityData() *types.CommonEntityData {
    views.EntityData.YFilter = views.YFilter
    views.EntityData.YangName = "views"
    views.EntityData.BundleName = "cisco_ios_xr"
    views.EntityData.ParentYangName = "snmp"
    views.EntityData.SegmentPath = "views"
    views.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    views.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    views.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    views.EntityData.Children = types.NewOrderedMap()
    views.EntityData.Children.Append("view", types.YChild{"View", nil})
    for i := range views.View {
        views.EntityData.Children.Append(types.GetSegmentPath(views.View[i]), types.YChild{"View", views.View[i]})
    }
    views.EntityData.Leafs = types.NewOrderedMap()

    views.EntityData.YListKeys = []string {}

    return &(views.EntityData)
}

// Snmp_Views_View
// Name of the view
type Snmp_Views_View struct {
    EntityData types.CommonEntityData
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

func (view *Snmp_Views_View) GetEntityData() *types.CommonEntityData {
    view.EntityData.YFilter = view.YFilter
    view.EntityData.YangName = "view"
    view.EntityData.BundleName = "cisco_ios_xr"
    view.EntityData.ParentYangName = "views"
    view.EntityData.SegmentPath = "view" + types.AddKeyToken(view.ViewName, "view-name") + types.AddKeyToken(view.Family, "family")
    view.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    view.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    view.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    view.EntityData.Children = types.NewOrderedMap()
    view.EntityData.Leafs = types.NewOrderedMap()
    view.EntityData.Leafs.Append("view-name", types.YLeaf{"ViewName", view.ViewName})
    view.EntityData.Leafs.Append("family", types.YLeaf{"Family", view.Family})
    view.EntityData.Leafs.Append("view-inclusion", types.YLeaf{"ViewInclusion", view.ViewInclusion})

    view.EntityData.YListKeys = []string {"ViewName", "Family"}

    return &(view.EntityData)
}

// Snmp_Logging
// SNMP logging
type Snmp_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP logging threshold.
    Threshold Snmp_Logging_Threshold
}

func (logging *Snmp_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "snmp"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Children.Append("threshold", types.YChild{"Threshold", &logging.Threshold})
    logging.EntityData.Leafs = types.NewOrderedMap()

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// Snmp_Logging_Threshold
// SNMP logging threshold
type Snmp_Logging_Threshold struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP logging threshold for OID processing. The type is interface{} with
    // range: 0..20000. The default value is 500.
    OidProcessing interface{}

    // SNMP logging threshold for PDU processing. The type is interface{} with
    // range: 0..20000. The default value is 20000.
    PduProcessing interface{}
}

func (threshold *Snmp_Logging_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "logging"
    threshold.EntityData.SegmentPath = "threshold"
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = types.NewOrderedMap()
    threshold.EntityData.Leafs = types.NewOrderedMap()
    threshold.EntityData.Leafs.Append("oid-processing", types.YLeaf{"OidProcessing", threshold.OidProcessing})
    threshold.EntityData.Leafs.Append("pdu-processing", types.YLeaf{"PduProcessing", threshold.PduProcessing})

    threshold.EntityData.YListKeys = []string {}

    return &(threshold.EntityData)
}

// Snmp_Administration
// Container class for SNMP administration
type Snmp_Administration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Container class to hold unencrpted communities.
    DefaultCommunities Snmp_Administration_DefaultCommunities

    // Container class to hold clear/encrypted communities.
    EncryptedCommunities Snmp_Administration_EncryptedCommunities
}

func (administration *Snmp_Administration) GetEntityData() *types.CommonEntityData {
    administration.EntityData.YFilter = administration.YFilter
    administration.EntityData.YangName = "administration"
    administration.EntityData.BundleName = "cisco_ios_xr"
    administration.EntityData.ParentYangName = "snmp"
    administration.EntityData.SegmentPath = "administration"
    administration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    administration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    administration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    administration.EntityData.Children = types.NewOrderedMap()
    administration.EntityData.Children.Append("default-communities", types.YChild{"DefaultCommunities", &administration.DefaultCommunities})
    administration.EntityData.Children.Append("encrypted-communities", types.YChild{"EncryptedCommunities", &administration.EncryptedCommunities})
    administration.EntityData.Leafs = types.NewOrderedMap()

    administration.EntityData.YListKeys = []string {}

    return &(administration.EntityData)
}

// Snmp_Administration_DefaultCommunities
// Container class to hold unencrpted communities
type Snmp_Administration_DefaultCommunities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unencrpted SNMP community string and access priviledges. The type is slice
    // of Snmp_Administration_DefaultCommunities_DefaultCommunity.
    DefaultCommunity []*Snmp_Administration_DefaultCommunities_DefaultCommunity
}

func (defaultCommunities *Snmp_Administration_DefaultCommunities) GetEntityData() *types.CommonEntityData {
    defaultCommunities.EntityData.YFilter = defaultCommunities.YFilter
    defaultCommunities.EntityData.YangName = "default-communities"
    defaultCommunities.EntityData.BundleName = "cisco_ios_xr"
    defaultCommunities.EntityData.ParentYangName = "administration"
    defaultCommunities.EntityData.SegmentPath = "default-communities"
    defaultCommunities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultCommunities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultCommunities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultCommunities.EntityData.Children = types.NewOrderedMap()
    defaultCommunities.EntityData.Children.Append("default-community", types.YChild{"DefaultCommunity", nil})
    for i := range defaultCommunities.DefaultCommunity {
        defaultCommunities.EntityData.Children.Append(types.GetSegmentPath(defaultCommunities.DefaultCommunity[i]), types.YChild{"DefaultCommunity", defaultCommunities.DefaultCommunity[i]})
    }
    defaultCommunities.EntityData.Leafs = types.NewOrderedMap()

    defaultCommunities.EntityData.YListKeys = []string {}

    return &(defaultCommunities.EntityData)
}

// Snmp_Administration_DefaultCommunities_DefaultCommunity
// Unencrpted SNMP community string and access
// priviledges
type Snmp_Administration_DefaultCommunities_DefaultCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SNMP community string. The type is string with
    // length: 1..128.
    CommunityName interface{}

    // Read/Write Access. The type is SnmpAccessLevel.
    Priviledge interface{}

    // MIB view to which the community has access. The type is string.
    ViewName interface{}

    // Access-list type. The type is Snmpacl.
    V4aclType interface{}

    // Ipv4 Access-list name. The type is string.
    V4AccessList interface{}

    // Access-list type. The type is Snmpacl.
    V6aclType interface{}

    // Ipv6 Access-list name. The type is string.
    V6AccessList interface{}

    // Logical Router or System owner access. The type is SnmpOwnerAccess.
    Owner interface{}
}

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) GetEntityData() *types.CommonEntityData {
    defaultCommunity.EntityData.YFilter = defaultCommunity.YFilter
    defaultCommunity.EntityData.YangName = "default-community"
    defaultCommunity.EntityData.BundleName = "cisco_ios_xr"
    defaultCommunity.EntityData.ParentYangName = "default-communities"
    defaultCommunity.EntityData.SegmentPath = "default-community" + types.AddKeyToken(defaultCommunity.CommunityName, "community-name")
    defaultCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultCommunity.EntityData.Children = types.NewOrderedMap()
    defaultCommunity.EntityData.Leafs = types.NewOrderedMap()
    defaultCommunity.EntityData.Leafs.Append("community-name", types.YLeaf{"CommunityName", defaultCommunity.CommunityName})
    defaultCommunity.EntityData.Leafs.Append("priviledge", types.YLeaf{"Priviledge", defaultCommunity.Priviledge})
    defaultCommunity.EntityData.Leafs.Append("view-name", types.YLeaf{"ViewName", defaultCommunity.ViewName})
    defaultCommunity.EntityData.Leafs.Append("v4acl-type", types.YLeaf{"V4aclType", defaultCommunity.V4aclType})
    defaultCommunity.EntityData.Leafs.Append("v4-access-list", types.YLeaf{"V4AccessList", defaultCommunity.V4AccessList})
    defaultCommunity.EntityData.Leafs.Append("v6acl-type", types.YLeaf{"V6aclType", defaultCommunity.V6aclType})
    defaultCommunity.EntityData.Leafs.Append("v6-access-list", types.YLeaf{"V6AccessList", defaultCommunity.V6AccessList})
    defaultCommunity.EntityData.Leafs.Append("owner", types.YLeaf{"Owner", defaultCommunity.Owner})

    defaultCommunity.EntityData.YListKeys = []string {"CommunityName"}

    return &(defaultCommunity.EntityData)
}

// Snmp_Administration_EncryptedCommunities
// Container class to hold clear/encrypted
// communities
type Snmp_Administration_EncryptedCommunities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear/encrypted SNMP community string and access priviledges. The type is
    // slice of Snmp_Administration_EncryptedCommunities_EncryptedCommunity.
    EncryptedCommunity []*Snmp_Administration_EncryptedCommunities_EncryptedCommunity
}

func (encryptedCommunities *Snmp_Administration_EncryptedCommunities) GetEntityData() *types.CommonEntityData {
    encryptedCommunities.EntityData.YFilter = encryptedCommunities.YFilter
    encryptedCommunities.EntityData.YangName = "encrypted-communities"
    encryptedCommunities.EntityData.BundleName = "cisco_ios_xr"
    encryptedCommunities.EntityData.ParentYangName = "administration"
    encryptedCommunities.EntityData.SegmentPath = "encrypted-communities"
    encryptedCommunities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryptedCommunities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryptedCommunities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryptedCommunities.EntityData.Children = types.NewOrderedMap()
    encryptedCommunities.EntityData.Children.Append("encrypted-community", types.YChild{"EncryptedCommunity", nil})
    for i := range encryptedCommunities.EncryptedCommunity {
        encryptedCommunities.EntityData.Children.Append(types.GetSegmentPath(encryptedCommunities.EncryptedCommunity[i]), types.YChild{"EncryptedCommunity", encryptedCommunities.EncryptedCommunity[i]})
    }
    encryptedCommunities.EntityData.Leafs = types.NewOrderedMap()

    encryptedCommunities.EntityData.YListKeys = []string {}

    return &(encryptedCommunities.EntityData)
}

// Snmp_Administration_EncryptedCommunities_EncryptedCommunity
// Clear/encrypted SNMP community string and
// access priviledges
type Snmp_Administration_EncryptedCommunities_EncryptedCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SNMP community string. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    CommunityName interface{}

    // Read/Write Access. The type is SnmpAccessLevel.
    Priviledge interface{}

    // MIB view to which the community has access. The type is string.
    ViewName interface{}

    // Access-list type. The type is Snmpacl.
    V4aclType interface{}

    // Ipv4 Access-list name. The type is string.
    V4AccessList interface{}

    // Access-list type. The type is Snmpacl.
    V6aclType interface{}

    // Ipv6 Access-list name. The type is string.
    V6AccessList interface{}

    // Logical Router or System owner access. The type is SnmpOwnerAccess.
    Owner interface{}
}

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) GetEntityData() *types.CommonEntityData {
    encryptedCommunity.EntityData.YFilter = encryptedCommunity.YFilter
    encryptedCommunity.EntityData.YangName = "encrypted-community"
    encryptedCommunity.EntityData.BundleName = "cisco_ios_xr"
    encryptedCommunity.EntityData.ParentYangName = "encrypted-communities"
    encryptedCommunity.EntityData.SegmentPath = "encrypted-community" + types.AddKeyToken(encryptedCommunity.CommunityName, "community-name")
    encryptedCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryptedCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryptedCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryptedCommunity.EntityData.Children = types.NewOrderedMap()
    encryptedCommunity.EntityData.Leafs = types.NewOrderedMap()
    encryptedCommunity.EntityData.Leafs.Append("community-name", types.YLeaf{"CommunityName", encryptedCommunity.CommunityName})
    encryptedCommunity.EntityData.Leafs.Append("priviledge", types.YLeaf{"Priviledge", encryptedCommunity.Priviledge})
    encryptedCommunity.EntityData.Leafs.Append("view-name", types.YLeaf{"ViewName", encryptedCommunity.ViewName})
    encryptedCommunity.EntityData.Leafs.Append("v4acl-type", types.YLeaf{"V4aclType", encryptedCommunity.V4aclType})
    encryptedCommunity.EntityData.Leafs.Append("v4-access-list", types.YLeaf{"V4AccessList", encryptedCommunity.V4AccessList})
    encryptedCommunity.EntityData.Leafs.Append("v6acl-type", types.YLeaf{"V6aclType", encryptedCommunity.V6aclType})
    encryptedCommunity.EntityData.Leafs.Append("v6-access-list", types.YLeaf{"V6AccessList", encryptedCommunity.V6AccessList})
    encryptedCommunity.EntityData.Leafs.Append("owner", types.YLeaf{"Owner", encryptedCommunity.Owner})

    encryptedCommunity.EntityData.YListKeys = []string {"CommunityName"}

    return &(encryptedCommunity.EntityData)
}

// Snmp_Agent
// The heirarchy point for SNMP Agent
// configurations
type Snmp_Agent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMPv3 engineID.
    EngineId Snmp_Agent_EngineId
}

func (agent *Snmp_Agent) GetEntityData() *types.CommonEntityData {
    agent.EntityData.YFilter = agent.YFilter
    agent.EntityData.YangName = "agent"
    agent.EntityData.BundleName = "cisco_ios_xr"
    agent.EntityData.ParentYangName = "snmp"
    agent.EntityData.SegmentPath = "agent"
    agent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    agent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    agent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    agent.EntityData.Children = types.NewOrderedMap()
    agent.EntityData.Children.Append("engine-id", types.YChild{"EngineId", &agent.EngineId})
    agent.EntityData.Leafs = types.NewOrderedMap()

    agent.EntityData.YListKeys = []string {}

    return &(agent.EntityData)
}

// Snmp_Agent_EngineId
// SNMPv3 engineID
type Snmp_Agent_EngineId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // engineID of the local agent. The type is string.
    Local interface{}

    // SNMPv3 remote SNMP Entity.
    Remotes Snmp_Agent_EngineId_Remotes
}

func (engineId *Snmp_Agent_EngineId) GetEntityData() *types.CommonEntityData {
    engineId.EntityData.YFilter = engineId.YFilter
    engineId.EntityData.YangName = "engine-id"
    engineId.EntityData.BundleName = "cisco_ios_xr"
    engineId.EntityData.ParentYangName = "agent"
    engineId.EntityData.SegmentPath = "engine-id"
    engineId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    engineId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    engineId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    engineId.EntityData.Children = types.NewOrderedMap()
    engineId.EntityData.Children.Append("remotes", types.YChild{"Remotes", &engineId.Remotes})
    engineId.EntityData.Leafs = types.NewOrderedMap()
    engineId.EntityData.Leafs.Append("local", types.YLeaf{"Local", engineId.Local})

    engineId.EntityData.YListKeys = []string {}

    return &(engineId.EntityData)
}

// Snmp_Agent_EngineId_Remotes
// SNMPv3 remote SNMP Entity
type Snmp_Agent_EngineId_Remotes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // engineID of the remote agent. The type is slice of
    // Snmp_Agent_EngineId_Remotes_Remote.
    Remote []*Snmp_Agent_EngineId_Remotes_Remote
}

func (remotes *Snmp_Agent_EngineId_Remotes) GetEntityData() *types.CommonEntityData {
    remotes.EntityData.YFilter = remotes.YFilter
    remotes.EntityData.YangName = "remotes"
    remotes.EntityData.BundleName = "cisco_ios_xr"
    remotes.EntityData.ParentYangName = "engine-id"
    remotes.EntityData.SegmentPath = "remotes"
    remotes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remotes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remotes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remotes.EntityData.Children = types.NewOrderedMap()
    remotes.EntityData.Children.Append("remote", types.YChild{"Remote", nil})
    for i := range remotes.Remote {
        remotes.EntityData.Children.Append(types.GetSegmentPath(remotes.Remote[i]), types.YChild{"Remote", remotes.Remote[i]})
    }
    remotes.EntityData.Leafs = types.NewOrderedMap()

    remotes.EntityData.YListKeys = []string {}

    return &(remotes.EntityData)
}

// Snmp_Agent_EngineId_Remotes_Remote
// engineID of the remote agent
type Snmp_Agent_EngineId_Remotes_Remote struct {
    EntityData types.CommonEntityData
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

func (remote *Snmp_Agent_EngineId_Remotes_Remote) GetEntityData() *types.CommonEntityData {
    remote.EntityData.YFilter = remote.YFilter
    remote.EntityData.YangName = "remote"
    remote.EntityData.BundleName = "cisco_ios_xr"
    remote.EntityData.ParentYangName = "remotes"
    remote.EntityData.SegmentPath = "remote" + types.AddKeyToken(remote.RemoteAddress, "remote-address")
    remote.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remote.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remote.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remote.EntityData.Children = types.NewOrderedMap()
    remote.EntityData.Leafs = types.NewOrderedMap()
    remote.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", remote.RemoteAddress})
    remote.EntityData.Leafs.Append("remote-engine-id", types.YLeaf{"RemoteEngineId", remote.RemoteEngineId})
    remote.EntityData.Leafs.Append("port", types.YLeaf{"Port", remote.Port})

    remote.EntityData.YListKeys = []string {"RemoteAddress"}

    return &(remote.EntityData)
}

// Snmp_Trap
// Class to hold trap configurations
type Snmp_Trap struct {
    EntityData types.CommonEntityData
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

func (trap *Snmp_Trap) GetEntityData() *types.CommonEntityData {
    trap.EntityData.YFilter = trap.YFilter
    trap.EntityData.YangName = "trap"
    trap.EntityData.BundleName = "cisco_ios_xr"
    trap.EntityData.ParentYangName = "snmp"
    trap.EntityData.SegmentPath = "trap"
    trap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trap.EntityData.Children = types.NewOrderedMap()
    trap.EntityData.Leafs = types.NewOrderedMap()
    trap.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", trap.Timeout})
    trap.EntityData.Leafs.Append("throttle-time", types.YLeaf{"ThrottleTime", trap.ThrottleTime})
    trap.EntityData.Leafs.Append("queue-length", types.YLeaf{"QueueLength", trap.QueueLength})

    trap.EntityData.YListKeys = []string {}

    return &(trap.EntityData)
}

// Snmp_DropPacket
// SNMP packet drop config
type Snmp_DropPacket struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable drop unknown user name. The type is interface{}.
    UnknownUser interface{}
}

func (dropPacket *Snmp_DropPacket) GetEntityData() *types.CommonEntityData {
    dropPacket.EntityData.YFilter = dropPacket.YFilter
    dropPacket.EntityData.YangName = "drop-packet"
    dropPacket.EntityData.BundleName = "cisco_ios_xr"
    dropPacket.EntityData.ParentYangName = "snmp"
    dropPacket.EntityData.SegmentPath = "drop-packet"
    dropPacket.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dropPacket.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dropPacket.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dropPacket.EntityData.Children = types.NewOrderedMap()
    dropPacket.EntityData.Leafs = types.NewOrderedMap()
    dropPacket.EntityData.Leafs.Append("unknown-user", types.YLeaf{"UnknownUser", dropPacket.UnknownUser})

    dropPacket.EntityData.YListKeys = []string {}

    return &(dropPacket.EntityData)
}

// Snmp_Ipv6
// SNMP TOS bit for outgoing packets
type Snmp_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of TOS.
    Tos Snmp_Ipv6_Tos
}

func (ipv6 *Snmp_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "snmp"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("tos", types.YChild{"Tos", &ipv6.Tos})
    ipv6.EntityData.Leafs = types.NewOrderedMap()

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Snmp_Ipv6_Tos
// Type of TOS
type Snmp_Ipv6_Tos struct {
    EntityData types.CommonEntityData
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

func (tos *Snmp_Ipv6_Tos) GetEntityData() *types.CommonEntityData {
    tos.EntityData.YFilter = tos.YFilter
    tos.EntityData.YangName = "tos"
    tos.EntityData.BundleName = "cisco_ios_xr"
    tos.EntityData.ParentYangName = "ipv6"
    tos.EntityData.SegmentPath = "tos"
    tos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tos.EntityData.Children = types.NewOrderedMap()
    tos.EntityData.Leafs = types.NewOrderedMap()
    tos.EntityData.Leafs.Append("type", types.YLeaf{"Type", tos.Type})
    tos.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", tos.Precedence})
    tos.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", tos.Dscp})

    tos.EntityData.YListKeys = []string {}

    return &(tos.EntityData)
}

// Snmp_Ipv4
// SNMP TOS bit for outgoing packets
type Snmp_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of TOS.
    Tos Snmp_Ipv4_Tos
}

func (ipv4 *Snmp_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "snmp"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("tos", types.YChild{"Tos", &ipv4.Tos})
    ipv4.EntityData.Leafs = types.NewOrderedMap()

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Snmp_Ipv4_Tos
// Type of TOS
type Snmp_Ipv4_Tos struct {
    EntityData types.CommonEntityData
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

func (tos *Snmp_Ipv4_Tos) GetEntityData() *types.CommonEntityData {
    tos.EntityData.YFilter = tos.YFilter
    tos.EntityData.YangName = "tos"
    tos.EntityData.BundleName = "cisco_ios_xr"
    tos.EntityData.ParentYangName = "ipv4"
    tos.EntityData.SegmentPath = "tos"
    tos.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tos.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tos.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tos.EntityData.Children = types.NewOrderedMap()
    tos.EntityData.Leafs = types.NewOrderedMap()
    tos.EntityData.Leafs.Append("type", types.YLeaf{"Type", tos.Type})
    tos.EntityData.Leafs.Append("precedence", types.YLeaf{"Precedence", tos.Precedence})
    tos.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", tos.Dscp})

    tos.EntityData.YListKeys = []string {}

    return &(tos.EntityData)
}

// Snmp_System
// container to hold system information
type Snmp_System struct {
    EntityData types.CommonEntityData
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

func (system *Snmp_System) GetEntityData() *types.CommonEntityData {
    system.EntityData.YFilter = system.YFilter
    system.EntityData.YangName = "system"
    system.EntityData.BundleName = "cisco_ios_xr"
    system.EntityData.ParentYangName = "snmp"
    system.EntityData.SegmentPath = "system"
    system.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    system.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    system.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    system.EntityData.Children = types.NewOrderedMap()
    system.EntityData.Leafs = types.NewOrderedMap()
    system.EntityData.Leafs.Append("chassis-id", types.YLeaf{"ChassisId", system.ChassisId})
    system.EntityData.Leafs.Append("location", types.YLeaf{"Location", system.Location})
    system.EntityData.Leafs.Append("contact", types.YLeaf{"Contact", system.Contact})

    system.EntityData.YListKeys = []string {}

    return &(system.EntityData)
}

// Snmp_Target
// SNMP target configurations
type Snmp_Target struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of targets.
    Targets Snmp_Target_Targets
}

func (target *Snmp_Target) GetEntityData() *types.CommonEntityData {
    target.EntityData.YFilter = target.YFilter
    target.EntityData.YangName = "target"
    target.EntityData.BundleName = "cisco_ios_xr"
    target.EntityData.ParentYangName = "snmp"
    target.EntityData.SegmentPath = "target"
    target.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    target.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    target.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    target.EntityData.Children = types.NewOrderedMap()
    target.EntityData.Children.Append("targets", types.YChild{"Targets", &target.Targets})
    target.EntityData.Leafs = types.NewOrderedMap()

    target.EntityData.YListKeys = []string {}

    return &(target.EntityData)
}

// Snmp_Target_Targets
// List of targets
type Snmp_Target_Targets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the target list. The type is slice of Snmp_Target_Targets_Target.
    Target []*Snmp_Target_Targets_Target
}

func (targets *Snmp_Target_Targets) GetEntityData() *types.CommonEntityData {
    targets.EntityData.YFilter = targets.YFilter
    targets.EntityData.YangName = "targets"
    targets.EntityData.BundleName = "cisco_ios_xr"
    targets.EntityData.ParentYangName = "target"
    targets.EntityData.SegmentPath = "targets"
    targets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    targets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    targets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    targets.EntityData.Children = types.NewOrderedMap()
    targets.EntityData.Children.Append("target", types.YChild{"Target", nil})
    for i := range targets.Target {
        targets.EntityData.Children.Append(types.GetSegmentPath(targets.Target[i]), types.YChild{"Target", targets.Target[i]})
    }
    targets.EntityData.Leafs = types.NewOrderedMap()

    targets.EntityData.YListKeys = []string {}

    return &(targets.EntityData)
}

// Snmp_Target_Targets_Target
// Name of the target list
type Snmp_Target_Targets_Target struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the target list. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    TargetListName interface{}

    // List of VRF Name for a target list.
    VrfNames Snmp_Target_Targets_Target_VrfNames

    // SNMP Target address configurations.
    TargetAddresses Snmp_Target_Targets_Target_TargetAddresses
}

func (target *Snmp_Target_Targets_Target) GetEntityData() *types.CommonEntityData {
    target.EntityData.YFilter = target.YFilter
    target.EntityData.YangName = "target"
    target.EntityData.BundleName = "cisco_ios_xr"
    target.EntityData.ParentYangName = "targets"
    target.EntityData.SegmentPath = "target" + types.AddKeyToken(target.TargetListName, "target-list-name")
    target.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    target.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    target.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    target.EntityData.Children = types.NewOrderedMap()
    target.EntityData.Children.Append("vrf-names", types.YChild{"VrfNames", &target.VrfNames})
    target.EntityData.Children.Append("target-addresses", types.YChild{"TargetAddresses", &target.TargetAddresses})
    target.EntityData.Leafs = types.NewOrderedMap()
    target.EntityData.Leafs.Append("target-list-name", types.YLeaf{"TargetListName", target.TargetListName})

    target.EntityData.YListKeys = []string {"TargetListName"}

    return &(target.EntityData)
}

// Snmp_Target_Targets_Target_VrfNames
// List of VRF Name for a target list
type Snmp_Target_Targets_Target_VrfNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name of the target. The type is slice of
    // Snmp_Target_Targets_Target_VrfNames_VrfName.
    VrfName []*Snmp_Target_Targets_Target_VrfNames_VrfName
}

func (vrfNames *Snmp_Target_Targets_Target_VrfNames) GetEntityData() *types.CommonEntityData {
    vrfNames.EntityData.YFilter = vrfNames.YFilter
    vrfNames.EntityData.YangName = "vrf-names"
    vrfNames.EntityData.BundleName = "cisco_ios_xr"
    vrfNames.EntityData.ParentYangName = "target"
    vrfNames.EntityData.SegmentPath = "vrf-names"
    vrfNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfNames.EntityData.Children = types.NewOrderedMap()
    vrfNames.EntityData.Children.Append("vrf-name", types.YChild{"VrfName", nil})
    for i := range vrfNames.VrfName {
        vrfNames.EntityData.Children.Append(types.GetSegmentPath(vrfNames.VrfName[i]), types.YChild{"VrfName", vrfNames.VrfName[i]})
    }
    vrfNames.EntityData.Leafs = types.NewOrderedMap()

    vrfNames.EntityData.YListKeys = []string {}

    return &(vrfNames.EntityData)
}

// Snmp_Target_Targets_Target_VrfNames_VrfName
// VRF name of the target
type Snmp_Target_Targets_Target_VrfNames_VrfName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}
}

func (vrfName *Snmp_Target_Targets_Target_VrfNames_VrfName) GetEntityData() *types.CommonEntityData {
    vrfName.EntityData.YFilter = vrfName.YFilter
    vrfName.EntityData.YangName = "vrf-name"
    vrfName.EntityData.BundleName = "cisco_ios_xr"
    vrfName.EntityData.ParentYangName = "vrf-names"
    vrfName.EntityData.SegmentPath = "vrf-name" + types.AddKeyToken(vrfName.Name, "name")
    vrfName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfName.EntityData.Children = types.NewOrderedMap()
    vrfName.EntityData.Leafs = types.NewOrderedMap()
    vrfName.EntityData.Leafs.Append("name", types.YLeaf{"Name", vrfName.Name})

    vrfName.EntityData.YListKeys = []string {"Name"}

    return &(vrfName.EntityData)
}

// Snmp_Target_Targets_Target_TargetAddresses
// SNMP Target address configurations
type Snmp_Target_Targets_Target_TargetAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP Address to be configured for the Target. The type is slice of
    // Snmp_Target_Targets_Target_TargetAddresses_TargetAddress.
    TargetAddress []*Snmp_Target_Targets_Target_TargetAddresses_TargetAddress
}

func (targetAddresses *Snmp_Target_Targets_Target_TargetAddresses) GetEntityData() *types.CommonEntityData {
    targetAddresses.EntityData.YFilter = targetAddresses.YFilter
    targetAddresses.EntityData.YangName = "target-addresses"
    targetAddresses.EntityData.BundleName = "cisco_ios_xr"
    targetAddresses.EntityData.ParentYangName = "target"
    targetAddresses.EntityData.SegmentPath = "target-addresses"
    targetAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    targetAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    targetAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    targetAddresses.EntityData.Children = types.NewOrderedMap()
    targetAddresses.EntityData.Children.Append("target-address", types.YChild{"TargetAddress", nil})
    for i := range targetAddresses.TargetAddress {
        targetAddresses.EntityData.Children.Append(types.GetSegmentPath(targetAddresses.TargetAddress[i]), types.YChild{"TargetAddress", targetAddresses.TargetAddress[i]})
    }
    targetAddresses.EntityData.Leafs = types.NewOrderedMap()

    targetAddresses.EntityData.YListKeys = []string {}

    return &(targetAddresses.EntityData)
}

// Snmp_Target_Targets_Target_TargetAddresses_TargetAddress
// IP Address to be configured for the Target
type Snmp_Target_Targets_Target_TargetAddresses_TargetAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4/Ipv6 address. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}
}

func (targetAddress *Snmp_Target_Targets_Target_TargetAddresses_TargetAddress) GetEntityData() *types.CommonEntityData {
    targetAddress.EntityData.YFilter = targetAddress.YFilter
    targetAddress.EntityData.YangName = "target-address"
    targetAddress.EntityData.BundleName = "cisco_ios_xr"
    targetAddress.EntityData.ParentYangName = "target-addresses"
    targetAddress.EntityData.SegmentPath = "target-address" + types.AddKeyToken(targetAddress.IpAddress, "ip-address")
    targetAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    targetAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    targetAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    targetAddress.EntityData.Children = types.NewOrderedMap()
    targetAddress.EntityData.Leafs = types.NewOrderedMap()
    targetAddress.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", targetAddress.IpAddress})

    targetAddress.EntityData.YListKeys = []string {"IpAddress"}

    return &(targetAddress.EntityData)
}

// Snmp_Notification
// Enable SNMP notifications
type Snmp_Notification struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable SNMP RTTMON-MIB IPSLA traps. The type is bool.
    Ipsla interface{}

    // SNMP notification configuration.
    Snmp Snmp_Notification_Snmp

    // Enable CISCO-ENTITY-EXT-MIB notifications.
    CiscoEntityExt Snmp_Notification_CiscoEntityExt

    // CISCO-ENTITY-FRU-CONTROL-MIB notification configuration.
    FruControl Snmp_Notification_FruControl

    // CISCO-NTP-MIB notification configuration.
    Ntp Snmp_Notification_Ntp

    // MPLS-L3VPN-STD-MIB notification configuration.
    MplsL3vpn Snmp_Notification_MplsL3vpn

    // CISCO-IETF-VPLS-GENERIC-MIB notification configuration.
    Vpls Snmp_Notification_Vpls

    // CISCO-IETF-PW-MIB notification configuration.
    L2vpn Snmp_Notification_L2vpn

    // 802.1ag Connectivity Fault Management MIB notification configuration.
    Cfm Snmp_Notification_Cfm

    // CISCO-HSRP-MIB notification configuration.
    Hsrp Snmp_Notification_Hsrp

    // Enable SNMP l2tun traps.
    L2tun Snmp_Notification_L2tun

    // CISCO-SELECTIVE-VRF-DOWNLOAD-MIB notification configuration.
    SelectiveVrfDownload Snmp_Notification_SelectiveVrfDownload

    // BGP4-MIB and CISCO-BGP4-MIB notification configuration.
    Bgp Snmp_Notification_Bgp

    // CISCO-IETF-BFD-MIB notification configuration.
    Bfd Snmp_Notification_Bfd

    // CISCO-FLASH-MIB notification configuration.
    Flash Snmp_Notification_Flash

    // OSPFv3-MIB notification configuration.
    Ospfv3 Snmp_Notification_Ospfv3

    // CISCO-OTN-IF-MIB notification configuration.
    Otn Snmp_Notification_Otn

    // CISCO-FABRIC-HFR-MIB notification configuration.
    FabricCrs Snmp_Notification_FabricCrs

    // CISCO-MPLS-TE-P2MP-STD-MIB notification configuration.
    MplsTeP2mp Snmp_Notification_MplsTeP2mp

    // MPLS-TE-STD-MIB notification configuration.
    MplsTe Snmp_Notification_MplsTe

    // CISCO-IETF-FRR-MIB notification configuration.
    MplsFrr Snmp_Notification_MplsFrr

    // CISCO-ENTITY-SENSOR-MIB notification configuration.
    Sensor Snmp_Notification_Sensor

    // Enable ISIS-MIB notifications.
    Isis Snmp_Notification_Isis

    // OSPF-MIB notification configuration.
    Ospf Snmp_Notification_Ospf

    // CISCO-CONFIG-COPY-MIB notification configuration.
    ConfigCopy Snmp_Notification_ConfigCopy

    // MPLS-LDP-STD-MIB notification configuration.
    MplsLdp Snmp_Notification_MplsLdp

    // CISCO-OPTICAL-MIB notification configuration.
    Optical Snmp_Notification_Optical

    // BRIDGE-MIB notification configuration.
    Bridge Snmp_Notification_Bridge

    // Frequency Synchronization trap configuration.
    FrequencySynchronization Snmp_Notification_FrequencySynchronization

    // VRRP-MIB notification configuration.
    Vrrp Snmp_Notification_Vrrp

    // 802.3 OAM MIB notification configuration.
    Oam Snmp_Notification_Oam

    // CISCO-SYSTEM-MIB notification configuration.
    System Snmp_Notification_System

    // Enable CISCO-IPSEC-FLOW-MONITOR-MIB notifications.
    IpSec Snmp_Notification_IpSec

    // Enable CISCO-IPSEC-FLOW-MONITOR-MIB notifications.
    Isakmp Snmp_Notification_Isakmp

    // CISCO-SYSLOG-MIB notification configuration.
    Syslog Snmp_Notification_Syslog

    // CISCO-ENTITY-REDUNDANCY-MIB notification configuration.
    EntityRedundancy Snmp_Notification_EntityRedundancy

    // Enable ENTITY-MIB notifications.
    Entity Snmp_Notification_Entity

    // Enable RSVP-MIB notifications.
    Rsvp Snmp_Notification_Rsvp

    // CISCO-CONFIG-MAN-MIB notification configurations.
    ConfigMan Snmp_Notification_ConfigMan

    // Subscriber notification commands.
    SubscriberMib Snmp_Notification_SubscriberMib

    // CISCO-OPTICAL-OTS-MIB notification configuration.
    OpticalOts Snmp_Notification_OpticalOts

    // Enable SNMP daps traps.
    AddresspoolMib Snmp_Notification_AddresspoolMib

    // Enable SNMP diameter traps.
    Diametermib Snmp_Notification_Diametermib

    // CISCO-RF-MIB notification configuration.
    Rf Snmp_Notification_Rf

    // ENTITY-STATE-MIB notification configuration.
    EntityState Snmp_Notification_EntityState
}

func (notification *Snmp_Notification) GetEntityData() *types.CommonEntityData {
    notification.EntityData.YFilter = notification.YFilter
    notification.EntityData.YangName = "notification"
    notification.EntityData.BundleName = "cisco_ios_xr"
    notification.EntityData.ParentYangName = "snmp"
    notification.EntityData.SegmentPath = "notification"
    notification.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    notification.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    notification.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    notification.EntityData.Children = types.NewOrderedMap()
    notification.EntityData.Children.Append("snmp", types.YChild{"Snmp", &notification.Snmp})
    notification.EntityData.Children.Append("Cisco-IOS-XR-snmp-entityextmib-cfg:cisco-entity-ext", types.YChild{"CiscoEntityExt", &notification.CiscoEntityExt})
    notification.EntityData.Children.Append("Cisco-IOS-XR-snmp-frucontrolmib-cfg:fru-control", types.YChild{"FruControl", &notification.FruControl})
    notification.EntityData.Children.Append("Cisco-IOS-XR-ip-ntp-cfg:ntp", types.YChild{"Ntp", &notification.Ntp})
    notification.EntityData.Children.Append("Cisco-IOS-XR-mpls-vpn-cfg:mpls-l3vpn", types.YChild{"MplsL3vpn", &notification.MplsL3vpn})
    notification.EntityData.Children.Append("Cisco-IOS-XR-l2vpn-cfg:vpls", types.YChild{"Vpls", &notification.Vpls})
    notification.EntityData.Children.Append("Cisco-IOS-XR-l2vpn-cfg:l2vpn", types.YChild{"L2vpn", &notification.L2vpn})
    notification.EntityData.Children.Append("Cisco-IOS-XR-ethernet-cfm-cfg:cfm", types.YChild{"Cfm", &notification.Cfm})
    notification.EntityData.Children.Append("Cisco-IOS-XR-ipv4-hsrp-cfg:hsrp", types.YChild{"Hsrp", &notification.Hsrp})
    notification.EntityData.Children.Append("Cisco-IOS-XR-tunnel-l2tun-proto-mibs-cfg:l2tun", types.YChild{"L2tun", &notification.L2tun})
    notification.EntityData.Children.Append("Cisco-IOS-XR-infra-rsi-cfg:selective-vrf-download", types.YChild{"SelectiveVrfDownload", &notification.SelectiveVrfDownload})
    notification.EntityData.Children.Append("Cisco-IOS-XR-ipv4-bgp-cfg:bgp", types.YChild{"Bgp", &notification.Bgp})
    notification.EntityData.Children.Append("Cisco-IOS-XR-ip-bfd-cfg:bfd", types.YChild{"Bfd", &notification.Bfd})
    notification.EntityData.Children.Append("Cisco-IOS-XR-flashmib-cfg:flash", types.YChild{"Flash", &notification.Flash})
    notification.EntityData.Children.Append("Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3", types.YChild{"Ospfv3", &notification.Ospfv3})
    notification.EntityData.Children.Append("Cisco-IOS-XR-otnifmib-cfg:otn", types.YChild{"Otn", &notification.Otn})
    notification.EntityData.Children.Append("Cisco-IOS-XR-fabhfr-mib-cfg:fabric-crs", types.YChild{"FabricCrs", &notification.FabricCrs})
    notification.EntityData.Children.Append("Cisco-IOS-XR-mpls-te-cfg:mpls-te-p2mp", types.YChild{"MplsTeP2mp", &notification.MplsTeP2mp})
    notification.EntityData.Children.Append("Cisco-IOS-XR-mpls-te-cfg:mpls-te", types.YChild{"MplsTe", &notification.MplsTe})
    notification.EntityData.Children.Append("Cisco-IOS-XR-mpls-te-cfg:mpls-frr", types.YChild{"MplsFrr", &notification.MplsFrr})
    notification.EntityData.Children.Append("Cisco-IOS-XR-snmp-ciscosensormib-cfg:sensor", types.YChild{"Sensor", &notification.Sensor})
    notification.EntityData.Children.Append("Cisco-IOS-XR-clns-isis-cfg:isis", types.YChild{"Isis", &notification.Isis})
    notification.EntityData.Children.Append("Cisco-IOS-XR-ipv4-ospf-cfg:ospf", types.YChild{"Ospf", &notification.Ospf})
    notification.EntityData.Children.Append("Cisco-IOS-XR-infra-confcopymib-cfg:config-copy", types.YChild{"ConfigCopy", &notification.ConfigCopy})
    notification.EntityData.Children.Append("Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp", types.YChild{"MplsLdp", &notification.MplsLdp})
    notification.EntityData.Children.Append("Cisco-IOS-XR-opticalmib-cfg:optical", types.YChild{"Optical", &notification.Optical})
    notification.EntityData.Children.Append("Cisco-IOS-XR-snmp-bridgemib-cfg:bridge", types.YChild{"Bridge", &notification.Bridge})
    notification.EntityData.Children.Append("Cisco-IOS-XR-freqsync-cfg:frequency-synchronization", types.YChild{"FrequencySynchronization", &notification.FrequencySynchronization})
    notification.EntityData.Children.Append("Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp", types.YChild{"Vrrp", &notification.Vrrp})
    notification.EntityData.Children.Append("Cisco-IOS-XR-ethernet-link-oam-cfg:oam", types.YChild{"Oam", &notification.Oam})
    notification.EntityData.Children.Append("Cisco-IOS-XR-infra-systemmib-cfg:system", types.YChild{"System", &notification.System})
    notification.EntityData.Children.Append("Cisco-IOS-XR-crypto-mibs-ipsecflowmon-cfg:ip-sec", types.YChild{"IpSec", &notification.IpSec})
    notification.EntityData.Children.Append("Cisco-IOS-XR-crypto-mibs-ipsecflowmon-cfg:isakmp", types.YChild{"Isakmp", &notification.Isakmp})
    notification.EntityData.Children.Append("Cisco-IOS-XR-snmp-syslogmib-cfg:syslog", types.YChild{"Syslog", &notification.Syslog})
    notification.EntityData.Children.Append("Cisco-IOS-XR-infra-ceredundancymib-cfg:entity-redundancy", types.YChild{"EntityRedundancy", &notification.EntityRedundancy})
    notification.EntityData.Children.Append("Cisco-IOS-XR-snmp-entitymib-cfg:entity", types.YChild{"Entity", &notification.Entity})
    notification.EntityData.Children.Append("Cisco-IOS-XR-ip-rsvp-cfg:rsvp", types.YChild{"Rsvp", &notification.Rsvp})
    notification.EntityData.Children.Append("Cisco-IOS-XR-config-mibs-cfg:config-man", types.YChild{"ConfigMan", &notification.ConfigMan})
    notification.EntityData.Children.Append("Cisco-IOS-XR-subscriber-session-mon-mibs-cfg:subscriber-mib", types.YChild{"SubscriberMib", &notification.SubscriberMib})
    notification.EntityData.Children.Append("Cisco-IOS-XR-opticalotsmib-cfg:optical-ots", types.YChild{"OpticalOts", &notification.OpticalOts})
    notification.EntityData.Children.Append("Cisco-IOS-XR-ip-daps-mib-cfg:addresspool-mib", types.YChild{"AddresspoolMib", &notification.AddresspoolMib})
    notification.EntityData.Children.Append("Cisco-IOS-XR-aaa-diameter-base-mib-cfg:diametermib", types.YChild{"Diametermib", &notification.Diametermib})
    notification.EntityData.Children.Append("Cisco-IOS-XR-snmp-mib-rfmib-cfg:rf", types.YChild{"Rf", &notification.Rf})
    notification.EntityData.Children.Append("Cisco-IOS-XR-snmp-entstatemib-cfg:entity-state", types.YChild{"EntityState", &notification.EntityState})
    notification.EntityData.Leafs = types.NewOrderedMap()
    notification.EntityData.Leafs.Append("ipsla", types.YLeaf{"Ipsla", notification.Ipsla})

    notification.EntityData.YListKeys = []string {}

    return &(notification.EntityData)
}

// Snmp_Notification_Snmp
// SNMP notification configuration
type Snmp_Notification_Snmp struct {
    EntityData types.CommonEntityData
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

func (snmp *Snmp_Notification_Snmp) GetEntityData() *types.CommonEntityData {
    snmp.EntityData.YFilter = snmp.YFilter
    snmp.EntityData.YangName = "snmp"
    snmp.EntityData.BundleName = "cisco_ios_xr"
    snmp.EntityData.ParentYangName = "notification"
    snmp.EntityData.SegmentPath = "snmp"
    snmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmp.EntityData.Children = types.NewOrderedMap()
    snmp.EntityData.Leafs = types.NewOrderedMap()
    snmp.EntityData.Leafs.Append("authentication", types.YLeaf{"Authentication", snmp.Authentication})
    snmp.EntityData.Leafs.Append("cold-start", types.YLeaf{"ColdStart", snmp.ColdStart})
    snmp.EntityData.Leafs.Append("warm-start", types.YLeaf{"WarmStart", snmp.WarmStart})
    snmp.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", snmp.Enable})
    snmp.EntityData.Leafs.Append("link-down", types.YLeaf{"LinkDown", snmp.LinkDown})
    snmp.EntityData.Leafs.Append("link-up", types.YLeaf{"LinkUp", snmp.LinkUp})

    snmp.EntityData.YListKeys = []string {}

    return &(snmp.EntityData)
}

// Snmp_Notification_CiscoEntityExt
// Enable CISCO-ENTITY-EXT-MIB notifications
type Snmp_Notification_CiscoEntityExt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable CISCO-ENTITY-EXT-MIB notifications. The type is interface{}.
    Enable interface{}
}

func (ciscoEntityExt *Snmp_Notification_CiscoEntityExt) GetEntityData() *types.CommonEntityData {
    ciscoEntityExt.EntityData.YFilter = ciscoEntityExt.YFilter
    ciscoEntityExt.EntityData.YangName = "cisco-entity-ext"
    ciscoEntityExt.EntityData.BundleName = "cisco_ios_xr"
    ciscoEntityExt.EntityData.ParentYangName = "notification"
    ciscoEntityExt.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-entityextmib-cfg:cisco-entity-ext"
    ciscoEntityExt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ciscoEntityExt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ciscoEntityExt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ciscoEntityExt.EntityData.Children = types.NewOrderedMap()
    ciscoEntityExt.EntityData.Leafs = types.NewOrderedMap()
    ciscoEntityExt.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ciscoEntityExt.Enable})

    ciscoEntityExt.EntityData.YListKeys = []string {}

    return &(ciscoEntityExt.EntityData)
}

// Snmp_Notification_FruControl
// CISCO-ENTITY-FRU-CONTROL-MIB notification
// configuration
type Snmp_Notification_FruControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ciscoEntityFRUControlMIB notifications. The type is interface{}.
    Enable interface{}
}

func (fruControl *Snmp_Notification_FruControl) GetEntityData() *types.CommonEntityData {
    fruControl.EntityData.YFilter = fruControl.YFilter
    fruControl.EntityData.YangName = "fru-control"
    fruControl.EntityData.BundleName = "cisco_ios_xr"
    fruControl.EntityData.ParentYangName = "notification"
    fruControl.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-frucontrolmib-cfg:fru-control"
    fruControl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fruControl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fruControl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fruControl.EntityData.Children = types.NewOrderedMap()
    fruControl.EntityData.Leafs = types.NewOrderedMap()
    fruControl.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", fruControl.Enable})

    fruControl.EntityData.YListKeys = []string {}

    return &(fruControl.EntityData)
}

// Snmp_Notification_Ntp
// CISCO-NTP-MIB notification configuration
type Snmp_Notification_Ntp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ciscoNtpMIB notification configuration. The type is interface{}.
    Enable interface{}
}

func (ntp *Snmp_Notification_Ntp) GetEntityData() *types.CommonEntityData {
    ntp.EntityData.YFilter = ntp.YFilter
    ntp.EntityData.YangName = "ntp"
    ntp.EntityData.BundleName = "cisco_ios_xr"
    ntp.EntityData.ParentYangName = "notification"
    ntp.EntityData.SegmentPath = "Cisco-IOS-XR-ip-ntp-cfg:ntp"
    ntp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ntp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ntp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ntp.EntityData.Children = types.NewOrderedMap()
    ntp.EntityData.Leafs = types.NewOrderedMap()
    ntp.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ntp.Enable})

    ntp.EntityData.YListKeys = []string {}

    return &(ntp.EntityData)
}

// Snmp_Notification_MplsL3vpn
// MPLS-L3VPN-STD-MIB notification configuration
type Snmp_Notification_MplsL3vpn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Time interval (secs) for re-issuing max-threshold notification. The type is
    // interface{} with range: 0..4294967295. Units are second.
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

func (mplsL3vpn *Snmp_Notification_MplsL3vpn) GetEntityData() *types.CommonEntityData {
    mplsL3vpn.EntityData.YFilter = mplsL3vpn.YFilter
    mplsL3vpn.EntityData.YangName = "mpls-l3vpn"
    mplsL3vpn.EntityData.BundleName = "cisco_ios_xr"
    mplsL3vpn.EntityData.ParentYangName = "notification"
    mplsL3vpn.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-vpn-cfg:mpls-l3vpn"
    mplsL3vpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsL3vpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsL3vpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsL3vpn.EntityData.Children = types.NewOrderedMap()
    mplsL3vpn.EntityData.Leafs = types.NewOrderedMap()
    mplsL3vpn.EntityData.Leafs.Append("max-threshold-reissue-notification-time", types.YLeaf{"MaxThresholdReissueNotificationTime", mplsL3vpn.MaxThresholdReissueNotificationTime})
    mplsL3vpn.EntityData.Leafs.Append("max-threshold-exceeded", types.YLeaf{"MaxThresholdExceeded", mplsL3vpn.MaxThresholdExceeded})
    mplsL3vpn.EntityData.Leafs.Append("max-threshold-cleared", types.YLeaf{"MaxThresholdCleared", mplsL3vpn.MaxThresholdCleared})
    mplsL3vpn.EntityData.Leafs.Append("mid-threshold-exceeded", types.YLeaf{"MidThresholdExceeded", mplsL3vpn.MidThresholdExceeded})
    mplsL3vpn.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", mplsL3vpn.Enable})
    mplsL3vpn.EntityData.Leafs.Append("vrf-down", types.YLeaf{"VrfDown", mplsL3vpn.VrfDown})
    mplsL3vpn.EntityData.Leafs.Append("vrf-up", types.YLeaf{"VrfUp", mplsL3vpn.VrfUp})

    mplsL3vpn.EntityData.YListKeys = []string {}

    return &(mplsL3vpn.EntityData)
}

// Snmp_Notification_Vpls
// CISCO-IETF-VPLS-GENERIC-MIB notification
// configuration
type Snmp_Notification_Vpls struct {
    EntityData types.CommonEntityData
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

func (vpls *Snmp_Notification_Vpls) GetEntityData() *types.CommonEntityData {
    vpls.EntityData.YFilter = vpls.YFilter
    vpls.EntityData.YangName = "vpls"
    vpls.EntityData.BundleName = "cisco_ios_xr"
    vpls.EntityData.ParentYangName = "notification"
    vpls.EntityData.SegmentPath = "Cisco-IOS-XR-l2vpn-cfg:vpls"
    vpls.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vpls.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vpls.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vpls.EntityData.Children = types.NewOrderedMap()
    vpls.EntityData.Leafs = types.NewOrderedMap()
    vpls.EntityData.Leafs.Append("full-clear", types.YLeaf{"FullClear", vpls.FullClear})
    vpls.EntityData.Leafs.Append("status", types.YLeaf{"Status", vpls.Status})
    vpls.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", vpls.Enable})
    vpls.EntityData.Leafs.Append("full-raise", types.YLeaf{"FullRaise", vpls.FullRaise})

    vpls.EntityData.YListKeys = []string {}

    return &(vpls.EntityData)
}

// Snmp_Notification_L2vpn
// CISCO-IETF-PW-MIB notification configuration
type Snmp_Notification_L2vpn struct {
    EntityData types.CommonEntityData
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

func (l2vpn *Snmp_Notification_L2vpn) GetEntityData() *types.CommonEntityData {
    l2vpn.EntityData.YFilter = l2vpn.YFilter
    l2vpn.EntityData.YangName = "l2vpn"
    l2vpn.EntityData.BundleName = "cisco_ios_xr"
    l2vpn.EntityData.ParentYangName = "notification"
    l2vpn.EntityData.SegmentPath = "Cisco-IOS-XR-l2vpn-cfg:l2vpn"
    l2vpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2vpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2vpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2vpn.EntityData.Children = types.NewOrderedMap()
    l2vpn.EntityData.Leafs = types.NewOrderedMap()
    l2vpn.EntityData.Leafs.Append("cisco", types.YLeaf{"Cisco", l2vpn.Cisco})
    l2vpn.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", l2vpn.Enable})
    l2vpn.EntityData.Leafs.Append("vc-down", types.YLeaf{"VcDown", l2vpn.VcDown})
    l2vpn.EntityData.Leafs.Append("vc-up", types.YLeaf{"VcUp", l2vpn.VcUp})

    l2vpn.EntityData.YListKeys = []string {}

    return &(l2vpn.EntityData)
}

// Snmp_Notification_Cfm
// 802.1ag Connectivity Fault Management MIB
// notification configuration
type Snmp_Notification_Cfm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable 802.1ag Connectivity Fault Management MIB notifications. The type is
    // interface{}.
    Enable interface{}
}

func (cfm *Snmp_Notification_Cfm) GetEntityData() *types.CommonEntityData {
    cfm.EntityData.YFilter = cfm.YFilter
    cfm.EntityData.YangName = "cfm"
    cfm.EntityData.BundleName = "cisco_ios_xr"
    cfm.EntityData.ParentYangName = "notification"
    cfm.EntityData.SegmentPath = "Cisco-IOS-XR-ethernet-cfm-cfg:cfm"
    cfm.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cfm.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cfm.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cfm.EntityData.Children = types.NewOrderedMap()
    cfm.EntityData.Leafs = types.NewOrderedMap()
    cfm.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", cfm.Enable})

    cfm.EntityData.YListKeys = []string {}

    return &(cfm.EntityData)
}

// Snmp_Notification_Hsrp
// CISCO-HSRP-MIB notification configuration
type Snmp_Notification_Hsrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable CISCO-HSRP-MIB notifications. The type is interface{}.
    Enable interface{}
}

func (hsrp *Snmp_Notification_Hsrp) GetEntityData() *types.CommonEntityData {
    hsrp.EntityData.YFilter = hsrp.YFilter
    hsrp.EntityData.YangName = "hsrp"
    hsrp.EntityData.BundleName = "cisco_ios_xr"
    hsrp.EntityData.ParentYangName = "notification"
    hsrp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-hsrp-cfg:hsrp"
    hsrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hsrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hsrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hsrp.EntityData.Children = types.NewOrderedMap()
    hsrp.EntityData.Leafs = types.NewOrderedMap()
    hsrp.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", hsrp.Enable})

    hsrp.EntityData.YListKeys = []string {}

    return &(hsrp.EntityData)
}

// Snmp_Notification_L2tun
// Enable SNMP l2tun traps
type Snmp_Notification_L2tun struct {
    EntityData types.CommonEntityData
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

func (l2tun *Snmp_Notification_L2tun) GetEntityData() *types.CommonEntityData {
    l2tun.EntityData.YFilter = l2tun.YFilter
    l2tun.EntityData.YangName = "l2tun"
    l2tun.EntityData.BundleName = "cisco_ios_xr"
    l2tun.EntityData.ParentYangName = "notification"
    l2tun.EntityData.SegmentPath = "Cisco-IOS-XR-tunnel-l2tun-proto-mibs-cfg:l2tun"
    l2tun.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2tun.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2tun.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2tun.EntityData.Children = types.NewOrderedMap()
    l2tun.EntityData.Leafs = types.NewOrderedMap()
    l2tun.EntityData.Leafs.Append("tunnel-up", types.YLeaf{"TunnelUp", l2tun.TunnelUp})
    l2tun.EntityData.Leafs.Append("tunnel-down", types.YLeaf{"TunnelDown", l2tun.TunnelDown})
    l2tun.EntityData.Leafs.Append("pseudowire-status", types.YLeaf{"PseudowireStatus", l2tun.PseudowireStatus})
    l2tun.EntityData.Leafs.Append("sessions", types.YLeaf{"Sessions", l2tun.Sessions})

    l2tun.EntityData.YListKeys = []string {}

    return &(l2tun.EntityData)
}

// Snmp_Notification_SelectiveVrfDownload
// CISCO-SELECTIVE-VRF-DOWNLOAD-MIB notification
// configuration
type Snmp_Notification_SelectiveVrfDownload struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable csvdEntityRoleChangeNotification notification. The type is
    // interface{}.
    RoleChange interface{}
}

func (selectiveVrfDownload *Snmp_Notification_SelectiveVrfDownload) GetEntityData() *types.CommonEntityData {
    selectiveVrfDownload.EntityData.YFilter = selectiveVrfDownload.YFilter
    selectiveVrfDownload.EntityData.YangName = "selective-vrf-download"
    selectiveVrfDownload.EntityData.BundleName = "cisco_ios_xr"
    selectiveVrfDownload.EntityData.ParentYangName = "notification"
    selectiveVrfDownload.EntityData.SegmentPath = "Cisco-IOS-XR-infra-rsi-cfg:selective-vrf-download"
    selectiveVrfDownload.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    selectiveVrfDownload.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    selectiveVrfDownload.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    selectiveVrfDownload.EntityData.Children = types.NewOrderedMap()
    selectiveVrfDownload.EntityData.Leafs = types.NewOrderedMap()
    selectiveVrfDownload.EntityData.Leafs.Append("role-change", types.YLeaf{"RoleChange", selectiveVrfDownload.RoleChange})

    selectiveVrfDownload.EntityData.YListKeys = []string {}

    return &(selectiveVrfDownload.EntityData)
}

// Snmp_Notification_Bgp
// BGP4-MIB and CISCO-BGP4-MIB notification
// configuration
type Snmp_Notification_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable BGP4-MIB and CISCO-BGP4-MIB IPv4-only notifications:
    // bgpEstablishedNotification, bgpBackwardTransNotification,
    // cbgpFsmStateChange, cbgpBackwardTransition, cbgpPrefixThresholdExceeded,
    // cbgpPrefixThresholdClear.
    Bgp4mib Snmp_Notification_Bgp_Bgp4mib

    // Enable CISCO-BGP4-MIB v2 notifications: cbgpPeer2EstablishedNotification,
    // cbgpPeer2BackwardTransNotification, cbgpPeer2FsmStateChange,
    // cbgpPeer2BackwardTransition, cbgpPeer2PrefixThresholdExceeded,
    // cbgpPeer2PrefixThresholdClear.
    CiscoBgp4mib Snmp_Notification_Bgp_CiscoBgp4mib
}

func (bgp *Snmp_Notification_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "notification"
    bgp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-bgp-cfg:bgp"
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Children.Append("bgp4mib", types.YChild{"Bgp4mib", &bgp.Bgp4mib})
    bgp.EntityData.Children.Append("cisco-bgp4mib", types.YChild{"CiscoBgp4mib", &bgp.CiscoBgp4mib})
    bgp.EntityData.Leafs = types.NewOrderedMap()

    bgp.EntityData.YListKeys = []string {}

    return &(bgp.EntityData)
}

// Snmp_Notification_Bgp_Bgp4mib
// Enable BGP4-MIB and CISCO-BGP4-MIB IPv4-only
// notifications: bgpEstablishedNotification,
// bgpBackwardTransNotification,
// cbgpFsmStateChange, cbgpBackwardTransition,
// cbgpPrefixThresholdExceeded,
// cbgpPrefixThresholdClear.
type Snmp_Notification_Bgp_Bgp4mib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable BGP4-MIB and CISCO-BGP4-MIB IPv4-only notifications. The type is
    // interface{}.
    Enable interface{}

    // Enable BGP4-MIB and CISCO-BGP4-MIB IPv4-only up/down notifications. The
    // type is interface{}.
    UpDown interface{}
}

func (bgp4mib *Snmp_Notification_Bgp_Bgp4mib) GetEntityData() *types.CommonEntityData {
    bgp4mib.EntityData.YFilter = bgp4mib.YFilter
    bgp4mib.EntityData.YangName = "bgp4mib"
    bgp4mib.EntityData.BundleName = "cisco_ios_xr"
    bgp4mib.EntityData.ParentYangName = "bgp"
    bgp4mib.EntityData.SegmentPath = "bgp4mib"
    bgp4mib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp4mib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp4mib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp4mib.EntityData.Children = types.NewOrderedMap()
    bgp4mib.EntityData.Leafs = types.NewOrderedMap()
    bgp4mib.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", bgp4mib.Enable})
    bgp4mib.EntityData.Leafs.Append("up-down", types.YLeaf{"UpDown", bgp4mib.UpDown})

    bgp4mib.EntityData.YListKeys = []string {}

    return &(bgp4mib.EntityData)
}

// Snmp_Notification_Bgp_CiscoBgp4mib
// Enable CISCO-BGP4-MIB v2 notifications:
// cbgpPeer2EstablishedNotification,
// cbgpPeer2BackwardTransNotification,
// cbgpPeer2FsmStateChange,
// cbgpPeer2BackwardTransition,
// cbgpPeer2PrefixThresholdExceeded,
// cbgpPeer2PrefixThresholdClear.
type Snmp_Notification_Bgp_CiscoBgp4mib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable CISCO-BGP4-MIB v2 notifications. The type is interface{}.
    Enable interface{}

    // Enable CISCO-BGP4-MIB v2 up/down notifications. The type is interface{}.
    UpDown interface{}
}

func (ciscoBgp4mib *Snmp_Notification_Bgp_CiscoBgp4mib) GetEntityData() *types.CommonEntityData {
    ciscoBgp4mib.EntityData.YFilter = ciscoBgp4mib.YFilter
    ciscoBgp4mib.EntityData.YangName = "cisco-bgp4mib"
    ciscoBgp4mib.EntityData.BundleName = "cisco_ios_xr"
    ciscoBgp4mib.EntityData.ParentYangName = "bgp"
    ciscoBgp4mib.EntityData.SegmentPath = "cisco-bgp4mib"
    ciscoBgp4mib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ciscoBgp4mib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ciscoBgp4mib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ciscoBgp4mib.EntityData.Children = types.NewOrderedMap()
    ciscoBgp4mib.EntityData.Leafs = types.NewOrderedMap()
    ciscoBgp4mib.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ciscoBgp4mib.Enable})
    ciscoBgp4mib.EntityData.Leafs.Append("up-down", types.YLeaf{"UpDown", ciscoBgp4mib.UpDown})

    ciscoBgp4mib.EntityData.YListKeys = []string {}

    return &(ciscoBgp4mib.EntityData)
}

// Snmp_Notification_Bfd
// CISCO-IETF-BFD-MIB notification configuration
type Snmp_Notification_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable CISCO-IETF-BFD-MIB notifications. The type is interface{}.
    Enable interface{}
}

func (bfd *Snmp_Notification_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "notification"
    bfd.EntityData.SegmentPath = "Cisco-IOS-XR-ip-bfd-cfg:bfd"
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", bfd.Enable})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Snmp_Notification_Flash
// CISCO-FLASH-MIB notification configuration
type Snmp_Notification_Flash struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ciscoFlashDeviceInsertedNotif notification. The type is interface{}.
    Insertion interface{}

    // Enable ciscoFlashDeviceRemovedNotif notification. The type is interface{}.
    Removal interface{}
}

func (flash *Snmp_Notification_Flash) GetEntityData() *types.CommonEntityData {
    flash.EntityData.YFilter = flash.YFilter
    flash.EntityData.YangName = "flash"
    flash.EntityData.BundleName = "cisco_ios_xr"
    flash.EntityData.ParentYangName = "notification"
    flash.EntityData.SegmentPath = "Cisco-IOS-XR-flashmib-cfg:flash"
    flash.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    flash.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    flash.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    flash.EntityData.Children = types.NewOrderedMap()
    flash.EntityData.Leafs = types.NewOrderedMap()
    flash.EntityData.Leafs.Append("insertion", types.YLeaf{"Insertion", flash.Insertion})
    flash.EntityData.Leafs.Append("removal", types.YLeaf{"Removal", flash.Removal})

    flash.EntityData.YListKeys = []string {}

    return &(flash.EntityData)
}

// Snmp_Notification_Ospfv3
// OSPFv3-MIB notification configuration
type Snmp_Notification_Ospfv3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP notifications for OSPF errors.
    Error Snmp_Notification_Ospfv3_Error

    // SNMP notifications for OSPF state change.
    StateChange Snmp_Notification_Ospfv3_StateChange
}

func (ospfv3 *Snmp_Notification_Ospfv3) GetEntityData() *types.CommonEntityData {
    ospfv3.EntityData.YFilter = ospfv3.YFilter
    ospfv3.EntityData.YangName = "ospfv3"
    ospfv3.EntityData.BundleName = "cisco_ios_xr"
    ospfv3.EntityData.ParentYangName = "notification"
    ospfv3.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3"
    ospfv3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3.EntityData.Children = types.NewOrderedMap()
    ospfv3.EntityData.Children.Append("error", types.YChild{"Error", &ospfv3.Error})
    ospfv3.EntityData.Children.Append("state-change", types.YChild{"StateChange", &ospfv3.StateChange})
    ospfv3.EntityData.Leafs = types.NewOrderedMap()

    ospfv3.EntityData.YListKeys = []string {}

    return &(ospfv3.EntityData)
}

// Snmp_Notification_Ospfv3_Error
// SNMP notifications for OSPF errors
type Snmp_Notification_Ospfv3_Error struct {
    EntityData types.CommonEntityData
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

func (self *Snmp_Notification_Ospfv3_Error) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "error"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "ospfv3"
    self.EntityData.SegmentPath = "error"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("config-error", types.YLeaf{"ConfigError", self.ConfigError})
    self.EntityData.Leafs.Append("bad-packet", types.YLeaf{"BadPacket", self.BadPacket})
    self.EntityData.Leafs.Append("virtual-bad-packet", types.YLeaf{"VirtualBadPacket", self.VirtualBadPacket})
    self.EntityData.Leafs.Append("virtual-config-error", types.YLeaf{"VirtualConfigError", self.VirtualConfigError})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Snmp_Notification_Ospfv3_StateChange
// SNMP notifications for OSPF state change
type Snmp_Notification_Ospfv3_StateChange struct {
    EntityData types.CommonEntityData
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

func (stateChange *Snmp_Notification_Ospfv3_StateChange) GetEntityData() *types.CommonEntityData {
    stateChange.EntityData.YFilter = stateChange.YFilter
    stateChange.EntityData.YangName = "state-change"
    stateChange.EntityData.BundleName = "cisco_ios_xr"
    stateChange.EntityData.ParentYangName = "ospfv3"
    stateChange.EntityData.SegmentPath = "state-change"
    stateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateChange.EntityData.Children = types.NewOrderedMap()
    stateChange.EntityData.Leafs = types.NewOrderedMap()
    stateChange.EntityData.Leafs.Append("restart-virtual-helper", types.YLeaf{"RestartVirtualHelper", stateChange.RestartVirtualHelper})
    stateChange.EntityData.Leafs.Append("nssa-translator", types.YLeaf{"NssaTranslator", stateChange.NssaTranslator})
    stateChange.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", stateChange.Interface})
    stateChange.EntityData.Leafs.Append("restart", types.YLeaf{"Restart", stateChange.Restart})
    stateChange.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", stateChange.Neighbor})
    stateChange.EntityData.Leafs.Append("virtual-interface", types.YLeaf{"VirtualInterface", stateChange.VirtualInterface})
    stateChange.EntityData.Leafs.Append("restart-helper", types.YLeaf{"RestartHelper", stateChange.RestartHelper})
    stateChange.EntityData.Leafs.Append("virtual-neighbor", types.YLeaf{"VirtualNeighbor", stateChange.VirtualNeighbor})

    stateChange.EntityData.YListKeys = []string {}

    return &(stateChange.EntityData)
}

// Snmp_Notification_Otn
// CISCO-OTN-IF-MIB notification configuration
type Snmp_Notification_Otn struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ciscoOtnIfMIB notifications. The type is interface{}.
    Enable interface{}
}

func (otn *Snmp_Notification_Otn) GetEntityData() *types.CommonEntityData {
    otn.EntityData.YFilter = otn.YFilter
    otn.EntityData.YangName = "otn"
    otn.EntityData.BundleName = "cisco_ios_xr"
    otn.EntityData.ParentYangName = "notification"
    otn.EntityData.SegmentPath = "Cisco-IOS-XR-otnifmib-cfg:otn"
    otn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    otn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    otn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    otn.EntityData.Children = types.NewOrderedMap()
    otn.EntityData.Leafs = types.NewOrderedMap()
    otn.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", otn.Enable})

    otn.EntityData.YListKeys = []string {}

    return &(otn.EntityData)
}

// Snmp_Notification_FabricCrs
// CISCO-FABRIC-HFR-MIB notification configuration
type Snmp_Notification_FabricCrs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable cfhBundleStateNotification notification. The type is interface{}.
    BundleState interface{}

    // Enable cfhPlaneStateNotification notification. The type is interface{}.
    PlaneState interface{}

    // Enable cfhBundleDownedLinkNotification notification. The type is
    // interface{}.
    BundleDownedLink interface{}
}

func (fabricCrs *Snmp_Notification_FabricCrs) GetEntityData() *types.CommonEntityData {
    fabricCrs.EntityData.YFilter = fabricCrs.YFilter
    fabricCrs.EntityData.YangName = "fabric-crs"
    fabricCrs.EntityData.BundleName = "cisco_ios_xr"
    fabricCrs.EntityData.ParentYangName = "notification"
    fabricCrs.EntityData.SegmentPath = "Cisco-IOS-XR-fabhfr-mib-cfg:fabric-crs"
    fabricCrs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fabricCrs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fabricCrs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fabricCrs.EntityData.Children = types.NewOrderedMap()
    fabricCrs.EntityData.Leafs = types.NewOrderedMap()
    fabricCrs.EntityData.Leafs.Append("bundle-state", types.YLeaf{"BundleState", fabricCrs.BundleState})
    fabricCrs.EntityData.Leafs.Append("plane-state", types.YLeaf{"PlaneState", fabricCrs.PlaneState})
    fabricCrs.EntityData.Leafs.Append("bundle-downed-link", types.YLeaf{"BundleDownedLink", fabricCrs.BundleDownedLink})

    fabricCrs.EntityData.YListKeys = []string {}

    return &(fabricCrs.EntityData)
}

// Snmp_Notification_MplsTeP2mp
// CISCO-MPLS-TE-P2MP-STD-MIB notification
// configuration
type Snmp_Notification_MplsTeP2mp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable cmplsTeP2mpTunnelDestUp notification. The type is interface{}.
    Up interface{}

    // Enable cmplsTeP2mpTunnelDestDown notification. The type is interface{}.
    Down interface{}
}

func (mplsTeP2mp *Snmp_Notification_MplsTeP2mp) GetEntityData() *types.CommonEntityData {
    mplsTeP2mp.EntityData.YFilter = mplsTeP2mp.YFilter
    mplsTeP2mp.EntityData.YangName = "mpls-te-p2mp"
    mplsTeP2mp.EntityData.BundleName = "cisco_ios_xr"
    mplsTeP2mp.EntityData.ParentYangName = "notification"
    mplsTeP2mp.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-te-cfg:mpls-te-p2mp"
    mplsTeP2mp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsTeP2mp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsTeP2mp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsTeP2mp.EntityData.Children = types.NewOrderedMap()
    mplsTeP2mp.EntityData.Leafs = types.NewOrderedMap()
    mplsTeP2mp.EntityData.Leafs.Append("up", types.YLeaf{"Up", mplsTeP2mp.Up})
    mplsTeP2mp.EntityData.Leafs.Append("down", types.YLeaf{"Down", mplsTeP2mp.Down})

    mplsTeP2mp.EntityData.YListKeys = []string {}

    return &(mplsTeP2mp.EntityData)
}

// Snmp_Notification_MplsTe
// MPLS-TE-STD-MIB notification configuration
type Snmp_Notification_MplsTe struct {
    EntityData types.CommonEntityData
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

func (mplsTe *Snmp_Notification_MplsTe) GetEntityData() *types.CommonEntityData {
    mplsTe.EntityData.YFilter = mplsTe.YFilter
    mplsTe.EntityData.YangName = "mpls-te"
    mplsTe.EntityData.BundleName = "cisco_ios_xr"
    mplsTe.EntityData.ParentYangName = "notification"
    mplsTe.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-te-cfg:mpls-te"
    mplsTe.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsTe.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsTe.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsTe.EntityData.Children = types.NewOrderedMap()
    mplsTe.EntityData.Children.Append("cisco-extension", types.YChild{"CiscoExtension", &mplsTe.CiscoExtension})
    mplsTe.EntityData.Leafs = types.NewOrderedMap()
    mplsTe.EntityData.Leafs.Append("cisco", types.YLeaf{"Cisco", mplsTe.Cisco})
    mplsTe.EntityData.Leafs.Append("up", types.YLeaf{"Up", mplsTe.Up})
    mplsTe.EntityData.Leafs.Append("reoptimize", types.YLeaf{"Reoptimize", mplsTe.Reoptimize})
    mplsTe.EntityData.Leafs.Append("reroute", types.YLeaf{"Reroute", mplsTe.Reroute})
    mplsTe.EntityData.Leafs.Append("down", types.YLeaf{"Down", mplsTe.Down})

    mplsTe.EntityData.YListKeys = []string {}

    return &(mplsTe.EntityData)
}

// Snmp_Notification_MplsTe_CiscoExtension
// CISCO-MPLS-TE-STD-EXT-MIB notification
// configuration
type Snmp_Notification_MplsTe_CiscoExtension struct {
    EntityData types.CommonEntityData
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

func (ciscoExtension *Snmp_Notification_MplsTe_CiscoExtension) GetEntityData() *types.CommonEntityData {
    ciscoExtension.EntityData.YFilter = ciscoExtension.YFilter
    ciscoExtension.EntityData.YangName = "cisco-extension"
    ciscoExtension.EntityData.BundleName = "cisco_ios_xr"
    ciscoExtension.EntityData.ParentYangName = "mpls-te"
    ciscoExtension.EntityData.SegmentPath = "cisco-extension"
    ciscoExtension.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ciscoExtension.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ciscoExtension.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ciscoExtension.EntityData.Children = types.NewOrderedMap()
    ciscoExtension.EntityData.Leafs = types.NewOrderedMap()
    ciscoExtension.EntityData.Leafs.Append("preempt", types.YLeaf{"Preempt", ciscoExtension.Preempt})
    ciscoExtension.EntityData.Leafs.Append("insufficient-bandwidth", types.YLeaf{"InsufficientBandwidth", ciscoExtension.InsufficientBandwidth})
    ciscoExtension.EntityData.Leafs.Append("re-route-pending-clear", types.YLeaf{"ReRoutePendingClear", ciscoExtension.ReRoutePendingClear})
    ciscoExtension.EntityData.Leafs.Append("bringup-fail", types.YLeaf{"BringupFail", ciscoExtension.BringupFail})
    ciscoExtension.EntityData.Leafs.Append("re-route-pending", types.YLeaf{"ReRoutePending", ciscoExtension.ReRoutePending})

    ciscoExtension.EntityData.YListKeys = []string {}

    return &(ciscoExtension.EntityData)
}

// Snmp_Notification_MplsFrr
// CISCO-IETF-FRR-MIB notification configuration
type Snmp_Notification_MplsFrr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable cmplsFrrUnProtected notification. The type is interface{}.
    Unprotected interface{}

    // Enable cmplsFrrMIB notifications. The type is interface{}.
    Enable interface{}

    // Enable cmplsFrrProtected notification. The type is interface{}.
    Protected interface{}
}

func (mplsFrr *Snmp_Notification_MplsFrr) GetEntityData() *types.CommonEntityData {
    mplsFrr.EntityData.YFilter = mplsFrr.YFilter
    mplsFrr.EntityData.YangName = "mpls-frr"
    mplsFrr.EntityData.BundleName = "cisco_ios_xr"
    mplsFrr.EntityData.ParentYangName = "notification"
    mplsFrr.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-te-cfg:mpls-frr"
    mplsFrr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsFrr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsFrr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsFrr.EntityData.Children = types.NewOrderedMap()
    mplsFrr.EntityData.Leafs = types.NewOrderedMap()
    mplsFrr.EntityData.Leafs.Append("unprotected", types.YLeaf{"Unprotected", mplsFrr.Unprotected})
    mplsFrr.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", mplsFrr.Enable})
    mplsFrr.EntityData.Leafs.Append("protected", types.YLeaf{"Protected", mplsFrr.Protected})

    mplsFrr.EntityData.YListKeys = []string {}

    return &(mplsFrr.EntityData)
}

// Snmp_Notification_Sensor
// CISCO-ENTITY-SENSOR-MIB notification
// configuration
type Snmp_Notification_Sensor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable entitySensorMIB notifications. The type is interface{}.
    Enable interface{}
}

func (sensor *Snmp_Notification_Sensor) GetEntityData() *types.CommonEntityData {
    sensor.EntityData.YFilter = sensor.YFilter
    sensor.EntityData.YangName = "sensor"
    sensor.EntityData.BundleName = "cisco_ios_xr"
    sensor.EntityData.ParentYangName = "notification"
    sensor.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-ciscosensormib-cfg:sensor"
    sensor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sensor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sensor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sensor.EntityData.Children = types.NewOrderedMap()
    sensor.EntityData.Leafs = types.NewOrderedMap()
    sensor.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", sensor.Enable})

    sensor.EntityData.YListKeys = []string {}

    return &(sensor.EntityData)
}

// Snmp_Notification_Isis
// Enable ISIS-MIB notifications
type Snmp_Notification_Isis struct {
    EntityData types.CommonEntityData
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

func (isis *Snmp_Notification_Isis) GetEntityData() *types.CommonEntityData {
    isis.EntityData.YFilter = isis.YFilter
    isis.EntityData.YangName = "isis"
    isis.EntityData.BundleName = "cisco_ios_xr"
    isis.EntityData.ParentYangName = "notification"
    isis.EntityData.SegmentPath = "Cisco-IOS-XR-clns-isis-cfg:isis"
    isis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isis.EntityData.Children = types.NewOrderedMap()
    isis.EntityData.Leafs = types.NewOrderedMap()
    isis.EntityData.Leafs.Append("database-overflow", types.YLeaf{"DatabaseOverflow", isis.DatabaseOverflow})
    isis.EntityData.Leafs.Append("manual-address-drops", types.YLeaf{"ManualAddressDrops", isis.ManualAddressDrops})
    isis.EntityData.Leafs.Append("corrupted-lsp-detected", types.YLeaf{"CorruptedLspDetected", isis.CorruptedLspDetected})
    isis.EntityData.Leafs.Append("attempt-to-exceed-max-sequence", types.YLeaf{"AttemptToExceedMaxSequence", isis.AttemptToExceedMaxSequence})
    isis.EntityData.Leafs.Append("id-length-mismatch", types.YLeaf{"IdLengthMismatch", isis.IdLengthMismatch})
    isis.EntityData.Leafs.Append("max-area-address-mismatch", types.YLeaf{"MaxAreaAddressMismatch", isis.MaxAreaAddressMismatch})
    isis.EntityData.Leafs.Append("own-lsp-purge", types.YLeaf{"OwnLspPurge", isis.OwnLspPurge})
    isis.EntityData.Leafs.Append("sequence-number-skip", types.YLeaf{"SequenceNumberSkip", isis.SequenceNumberSkip})
    isis.EntityData.Leafs.Append("authentication-type-failure", types.YLeaf{"AuthenticationTypeFailure", isis.AuthenticationTypeFailure})
    isis.EntityData.Leafs.Append("authentication-failure", types.YLeaf{"AuthenticationFailure", isis.AuthenticationFailure})
    isis.EntityData.Leafs.Append("version-skew", types.YLeaf{"VersionSkew", isis.VersionSkew})
    isis.EntityData.Leafs.Append("area-mismatch", types.YLeaf{"AreaMismatch", isis.AreaMismatch})
    isis.EntityData.Leafs.Append("rejected-adjacency", types.YLeaf{"RejectedAdjacency", isis.RejectedAdjacency})
    isis.EntityData.Leafs.Append("lsp-too-large-to-propagate", types.YLeaf{"LspTooLargeToPropagate", isis.LspTooLargeToPropagate})
    isis.EntityData.Leafs.Append("originated-lsp-buffer-size-mismatch", types.YLeaf{"OriginatedLspBufferSizeMismatch", isis.OriginatedLspBufferSizeMismatch})
    isis.EntityData.Leafs.Append("protocols-supported-mismatch", types.YLeaf{"ProtocolsSupportedMismatch", isis.ProtocolsSupportedMismatch})
    isis.EntityData.Leafs.Append("adjacency-change", types.YLeaf{"AdjacencyChange", isis.AdjacencyChange})
    isis.EntityData.Leafs.Append("lsp-error-detected", types.YLeaf{"LspErrorDetected", isis.LspErrorDetected})
    isis.EntityData.Leafs.Append("all", types.YLeaf{"All", isis.All})

    isis.EntityData.YListKeys = []string {}

    return &(isis.EntityData)
}

// Snmp_Notification_Ospf
// OSPF-MIB notification configuration
type Snmp_Notification_Ospf struct {
    EntityData types.CommonEntityData
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

func (ospf *Snmp_Notification_Ospf) GetEntityData() *types.CommonEntityData {
    ospf.EntityData.YFilter = ospf.YFilter
    ospf.EntityData.YangName = "ospf"
    ospf.EntityData.BundleName = "cisco_ios_xr"
    ospf.EntityData.ParentYangName = "notification"
    ospf.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-ospf-cfg:ospf"
    ospf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospf.EntityData.Children = types.NewOrderedMap()
    ospf.EntityData.Children.Append("lsa", types.YChild{"Lsa", &ospf.Lsa})
    ospf.EntityData.Children.Append("state-change", types.YChild{"StateChange", &ospf.StateChange})
    ospf.EntityData.Children.Append("retransmit", types.YChild{"Retransmit", &ospf.Retransmit})
    ospf.EntityData.Children.Append("error", types.YChild{"Error", &ospf.Error})
    ospf.EntityData.Leafs = types.NewOrderedMap()

    ospf.EntityData.YListKeys = []string {}

    return &(ospf.EntityData)
}

// Snmp_Notification_Ospf_Lsa
// SNMP notifications related to LSAs
type Snmp_Notification_Ospf_Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ospfMaxAgeLsa notification. The type is interface{}.
    MaxAgeLsa interface{}

    // Enable ospfOriginateLsa notification. The type is interface{}.
    OriginateLsa interface{}
}

func (lsa *Snmp_Notification_Ospf_Lsa) GetEntityData() *types.CommonEntityData {
    lsa.EntityData.YFilter = lsa.YFilter
    lsa.EntityData.YangName = "lsa"
    lsa.EntityData.BundleName = "cisco_ios_xr"
    lsa.EntityData.ParentYangName = "ospf"
    lsa.EntityData.SegmentPath = "lsa"
    lsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsa.EntityData.Children = types.NewOrderedMap()
    lsa.EntityData.Leafs = types.NewOrderedMap()
    lsa.EntityData.Leafs.Append("max-age-lsa", types.YLeaf{"MaxAgeLsa", lsa.MaxAgeLsa})
    lsa.EntityData.Leafs.Append("originate-lsa", types.YLeaf{"OriginateLsa", lsa.OriginateLsa})

    lsa.EntityData.YListKeys = []string {}

    return &(lsa.EntityData)
}

// Snmp_Notification_Ospf_StateChange
// SNMP notifications for OSPF state change
type Snmp_Notification_Ospf_StateChange struct {
    EntityData types.CommonEntityData
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

func (stateChange *Snmp_Notification_Ospf_StateChange) GetEntityData() *types.CommonEntityData {
    stateChange.EntityData.YFilter = stateChange.YFilter
    stateChange.EntityData.YangName = "state-change"
    stateChange.EntityData.BundleName = "cisco_ios_xr"
    stateChange.EntityData.ParentYangName = "ospf"
    stateChange.EntityData.SegmentPath = "state-change"
    stateChange.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stateChange.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stateChange.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stateChange.EntityData.Children = types.NewOrderedMap()
    stateChange.EntityData.Leafs = types.NewOrderedMap()
    stateChange.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", stateChange.Interface})
    stateChange.EntityData.Leafs.Append("virtual-interface", types.YLeaf{"VirtualInterface", stateChange.VirtualInterface})
    stateChange.EntityData.Leafs.Append("virtual-neighbor", types.YLeaf{"VirtualNeighbor", stateChange.VirtualNeighbor})
    stateChange.EntityData.Leafs.Append("neighbor", types.YLeaf{"Neighbor", stateChange.Neighbor})

    stateChange.EntityData.YListKeys = []string {}

    return &(stateChange.EntityData)
}

// Snmp_Notification_Ospf_Retransmit
// SNMP notifications for packet retransmissions
type Snmp_Notification_Ospf_Retransmit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ospfVirtIfTxRetransmit notification. The type is interface{}.
    VirtualPacket interface{}

    // Enable ospfTxRetransmit notification. The type is interface{}.
    Packet interface{}
}

func (retransmit *Snmp_Notification_Ospf_Retransmit) GetEntityData() *types.CommonEntityData {
    retransmit.EntityData.YFilter = retransmit.YFilter
    retransmit.EntityData.YangName = "retransmit"
    retransmit.EntityData.BundleName = "cisco_ios_xr"
    retransmit.EntityData.ParentYangName = "ospf"
    retransmit.EntityData.SegmentPath = "retransmit"
    retransmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    retransmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    retransmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    retransmit.EntityData.Children = types.NewOrderedMap()
    retransmit.EntityData.Leafs = types.NewOrderedMap()
    retransmit.EntityData.Leafs.Append("virtual-packet", types.YLeaf{"VirtualPacket", retransmit.VirtualPacket})
    retransmit.EntityData.Leafs.Append("packet", types.YLeaf{"Packet", retransmit.Packet})

    retransmit.EntityData.YListKeys = []string {}

    return &(retransmit.EntityData)
}

// Snmp_Notification_Ospf_Error
// SNMP notifications for OSPF errors
type Snmp_Notification_Ospf_Error struct {
    EntityData types.CommonEntityData
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

func (self *Snmp_Notification_Ospf_Error) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "error"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "ospf"
    self.EntityData.SegmentPath = "error"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("config-error", types.YLeaf{"ConfigError", self.ConfigError})
    self.EntityData.Leafs.Append("authentication-failure", types.YLeaf{"AuthenticationFailure", self.AuthenticationFailure})
    self.EntityData.Leafs.Append("virtual-config-error", types.YLeaf{"VirtualConfigError", self.VirtualConfigError})
    self.EntityData.Leafs.Append("virtual-authentication-failure", types.YLeaf{"VirtualAuthenticationFailure", self.VirtualAuthenticationFailure})
    self.EntityData.Leafs.Append("bad-packet", types.YLeaf{"BadPacket", self.BadPacket})
    self.EntityData.Leafs.Append("virtual-bad-packet", types.YLeaf{"VirtualBadPacket", self.VirtualBadPacket})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// Snmp_Notification_ConfigCopy
// CISCO-CONFIG-COPY-MIB notification configuration
type Snmp_Notification_ConfigCopy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ccCopyCompletion notification. The type is interface{}.
    Completion interface{}
}

func (configCopy *Snmp_Notification_ConfigCopy) GetEntityData() *types.CommonEntityData {
    configCopy.EntityData.YFilter = configCopy.YFilter
    configCopy.EntityData.YangName = "config-copy"
    configCopy.EntityData.BundleName = "cisco_ios_xr"
    configCopy.EntityData.ParentYangName = "notification"
    configCopy.EntityData.SegmentPath = "Cisco-IOS-XR-infra-confcopymib-cfg:config-copy"
    configCopy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configCopy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configCopy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configCopy.EntityData.Children = types.NewOrderedMap()
    configCopy.EntityData.Leafs = types.NewOrderedMap()
    configCopy.EntityData.Leafs.Append("completion", types.YLeaf{"Completion", configCopy.Completion})

    configCopy.EntityData.YListKeys = []string {}

    return &(configCopy.EntityData)
}

// Snmp_Notification_MplsLdp
// MPLS-LDP-STD-MIB notification configuration
type Snmp_Notification_MplsLdp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable mplsLdpSessionUp notification. The type is interface{}.
    SessionUp interface{}

    // Enable mplsLdpInitSessionThresholdExceeded notification. The type is
    // interface{}.
    InitSessionThresholdExceeded interface{}

    // Enable mplsLdpSessionDown notification. The type is interface{}.
    SessionDown interface{}
}

func (mplsLdp *Snmp_Notification_MplsLdp) GetEntityData() *types.CommonEntityData {
    mplsLdp.EntityData.YFilter = mplsLdp.YFilter
    mplsLdp.EntityData.YangName = "mpls-ldp"
    mplsLdp.EntityData.BundleName = "cisco_ios_xr"
    mplsLdp.EntityData.ParentYangName = "notification"
    mplsLdp.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp"
    mplsLdp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsLdp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsLdp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsLdp.EntityData.Children = types.NewOrderedMap()
    mplsLdp.EntityData.Leafs = types.NewOrderedMap()
    mplsLdp.EntityData.Leafs.Append("session-up", types.YLeaf{"SessionUp", mplsLdp.SessionUp})
    mplsLdp.EntityData.Leafs.Append("init-session-threshold-exceeded", types.YLeaf{"InitSessionThresholdExceeded", mplsLdp.InitSessionThresholdExceeded})
    mplsLdp.EntityData.Leafs.Append("session-down", types.YLeaf{"SessionDown", mplsLdp.SessionDown})

    mplsLdp.EntityData.YListKeys = []string {}

    return &(mplsLdp.EntityData)
}

// Snmp_Notification_Optical
// CISCO-OPTICAL-MIB notification configuration
type Snmp_Notification_Optical struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Opticalmib notifications. The type is interface{}.
    Enable interface{}
}

func (optical *Snmp_Notification_Optical) GetEntityData() *types.CommonEntityData {
    optical.EntityData.YFilter = optical.YFilter
    optical.EntityData.YangName = "optical"
    optical.EntityData.BundleName = "cisco_ios_xr"
    optical.EntityData.ParentYangName = "notification"
    optical.EntityData.SegmentPath = "Cisco-IOS-XR-opticalmib-cfg:optical"
    optical.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optical.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optical.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optical.EntityData.Children = types.NewOrderedMap()
    optical.EntityData.Leafs = types.NewOrderedMap()
    optical.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", optical.Enable})

    optical.EntityData.YListKeys = []string {}

    return &(optical.EntityData)
}

// Snmp_Notification_Bridge
// BRIDGE-MIB notification configuration
type Snmp_Notification_Bridge struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable dot1dBridge notifications. The type is interface{}.
    Enable interface{}
}

func (bridge *Snmp_Notification_Bridge) GetEntityData() *types.CommonEntityData {
    bridge.EntityData.YFilter = bridge.YFilter
    bridge.EntityData.YangName = "bridge"
    bridge.EntityData.BundleName = "cisco_ios_xr"
    bridge.EntityData.ParentYangName = "notification"
    bridge.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-bridgemib-cfg:bridge"
    bridge.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bridge.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bridge.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bridge.EntityData.Children = types.NewOrderedMap()
    bridge.EntityData.Leafs = types.NewOrderedMap()
    bridge.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", bridge.Enable})

    bridge.EntityData.YListKeys = []string {}

    return &(bridge.EntityData)
}

// Snmp_Notification_FrequencySynchronization
// Frequency Synchronization trap configuration
type Snmp_Notification_FrequencySynchronization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable Frequency Synchronization traps. The type is interface{}.
    Enable interface{}
}

func (frequencySynchronization *Snmp_Notification_FrequencySynchronization) GetEntityData() *types.CommonEntityData {
    frequencySynchronization.EntityData.YFilter = frequencySynchronization.YFilter
    frequencySynchronization.EntityData.YangName = "frequency-synchronization"
    frequencySynchronization.EntityData.BundleName = "cisco_ios_xr"
    frequencySynchronization.EntityData.ParentYangName = "notification"
    frequencySynchronization.EntityData.SegmentPath = "Cisco-IOS-XR-freqsync-cfg:frequency-synchronization"
    frequencySynchronization.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    frequencySynchronization.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    frequencySynchronization.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    frequencySynchronization.EntityData.Children = types.NewOrderedMap()
    frequencySynchronization.EntityData.Leafs = types.NewOrderedMap()
    frequencySynchronization.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", frequencySynchronization.Enable})

    frequencySynchronization.EntityData.YListKeys = []string {}

    return &(frequencySynchronization.EntityData)
}

// Snmp_Notification_Vrrp
// VRRP-MIB notification configuration
type Snmp_Notification_Vrrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable VRRP-MIB notifications. The type is interface{}.
    Enable interface{}
}

func (vrrp *Snmp_Notification_Vrrp) GetEntityData() *types.CommonEntityData {
    vrrp.EntityData.YFilter = vrrp.YFilter
    vrrp.EntityData.YangName = "vrrp"
    vrrp.EntityData.BundleName = "cisco_ios_xr"
    vrrp.EntityData.ParentYangName = "notification"
    vrrp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp"
    vrrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrrp.EntityData.Children = types.NewOrderedMap()
    vrrp.EntityData.Leafs = types.NewOrderedMap()
    vrrp.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", vrrp.Enable})

    vrrp.EntityData.YListKeys = []string {}

    return &(vrrp.EntityData)
}

// Snmp_Notification_Oam
// 802.3 OAM MIB notification configuration
type Snmp_Notification_Oam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable 802.3 OAM MIB notifications. The type is interface{}.
    Enable interface{}
}

func (oam *Snmp_Notification_Oam) GetEntityData() *types.CommonEntityData {
    oam.EntityData.YFilter = oam.YFilter
    oam.EntityData.YangName = "oam"
    oam.EntityData.BundleName = "cisco_ios_xr"
    oam.EntityData.ParentYangName = "notification"
    oam.EntityData.SegmentPath = "Cisco-IOS-XR-ethernet-link-oam-cfg:oam"
    oam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    oam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    oam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    oam.EntityData.Children = types.NewOrderedMap()
    oam.EntityData.Leafs = types.NewOrderedMap()
    oam.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", oam.Enable})

    oam.EntityData.YListKeys = []string {}

    return &(oam.EntityData)
}

// Snmp_Notification_System
// CISCO-SYSTEM-MIB notification configuration
type Snmp_Notification_System struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ciscoSystemMIB notifications. The type is interface{}.
    Enable interface{}
}

func (system *Snmp_Notification_System) GetEntityData() *types.CommonEntityData {
    system.EntityData.YFilter = system.YFilter
    system.EntityData.YangName = "system"
    system.EntityData.BundleName = "cisco_ios_xr"
    system.EntityData.ParentYangName = "notification"
    system.EntityData.SegmentPath = "Cisco-IOS-XR-infra-systemmib-cfg:system"
    system.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    system.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    system.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    system.EntityData.Children = types.NewOrderedMap()
    system.EntityData.Leafs = types.NewOrderedMap()
    system.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", system.Enable})

    system.EntityData.YListKeys = []string {}

    return &(system.EntityData)
}

// Snmp_Notification_IpSec
// Enable CISCO-IPSEC-FLOW-MONITOR-MIB
// notifications
type Snmp_Notification_IpSec struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable cipSecTunnelStop notification. The type is interface{}.
    TunnelStop interface{}

    // Enable cipSecTunnelStart notification. The type is interface{}.
    TunnelStart interface{}
}

func (ipSec *Snmp_Notification_IpSec) GetEntityData() *types.CommonEntityData {
    ipSec.EntityData.YFilter = ipSec.YFilter
    ipSec.EntityData.YangName = "ip-sec"
    ipSec.EntityData.BundleName = "cisco_ios_xr"
    ipSec.EntityData.ParentYangName = "notification"
    ipSec.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-mibs-ipsecflowmon-cfg:ip-sec"
    ipSec.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipSec.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipSec.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipSec.EntityData.Children = types.NewOrderedMap()
    ipSec.EntityData.Leafs = types.NewOrderedMap()
    ipSec.EntityData.Leafs.Append("tunnel-stop", types.YLeaf{"TunnelStop", ipSec.TunnelStop})
    ipSec.EntityData.Leafs.Append("tunnel-start", types.YLeaf{"TunnelStart", ipSec.TunnelStart})

    ipSec.EntityData.YListKeys = []string {}

    return &(ipSec.EntityData)
}

// Snmp_Notification_Isakmp
// Enable CISCO-IPSEC-FLOW-MONITOR-MIB
// notifications
type Snmp_Notification_Isakmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable cikeTunnelStop notification. The type is interface{}.
    TunnelStop interface{}

    // Enable cikeTunnelStart notification. The type is interface{}.
    TunnelStart interface{}
}

func (isakmp *Snmp_Notification_Isakmp) GetEntityData() *types.CommonEntityData {
    isakmp.EntityData.YFilter = isakmp.YFilter
    isakmp.EntityData.YangName = "isakmp"
    isakmp.EntityData.BundleName = "cisco_ios_xr"
    isakmp.EntityData.ParentYangName = "notification"
    isakmp.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-mibs-ipsecflowmon-cfg:isakmp"
    isakmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    isakmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    isakmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    isakmp.EntityData.Children = types.NewOrderedMap()
    isakmp.EntityData.Leafs = types.NewOrderedMap()
    isakmp.EntityData.Leafs.Append("tunnel-stop", types.YLeaf{"TunnelStop", isakmp.TunnelStop})
    isakmp.EntityData.Leafs.Append("tunnel-start", types.YLeaf{"TunnelStart", isakmp.TunnelStart})

    isakmp.EntityData.YListKeys = []string {}

    return &(isakmp.EntityData)
}

// Snmp_Notification_Syslog
// CISCO-SYSLOG-MIB notification configuration
type Snmp_Notification_Syslog struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ciscoSyslogMIB notifications. The type is interface{}.
    Enable interface{}
}

func (syslog *Snmp_Notification_Syslog) GetEntityData() *types.CommonEntityData {
    syslog.EntityData.YFilter = syslog.YFilter
    syslog.EntityData.YangName = "syslog"
    syslog.EntityData.BundleName = "cisco_ios_xr"
    syslog.EntityData.ParentYangName = "notification"
    syslog.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-syslogmib-cfg:syslog"
    syslog.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    syslog.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    syslog.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    syslog.EntityData.Children = types.NewOrderedMap()
    syslog.EntityData.Leafs = types.NewOrderedMap()
    syslog.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", syslog.Enable})

    syslog.EntityData.YListKeys = []string {}

    return &(syslog.EntityData)
}

// Snmp_Notification_EntityRedundancy
// CISCO-ENTITY-REDUNDANCY-MIB notification
// configuration
type Snmp_Notification_EntityRedundancy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable switchover notification. The type is interface{}.
    Switchover interface{}

    // Enable CISCO-ENTITY-REDUNDANCY-MIB notification. The type is interface{}.
    Enable interface{}

    // Enable status change notification. The type is interface{}.
    Status interface{}
}

func (entityRedundancy *Snmp_Notification_EntityRedundancy) GetEntityData() *types.CommonEntityData {
    entityRedundancy.EntityData.YFilter = entityRedundancy.YFilter
    entityRedundancy.EntityData.YangName = "entity-redundancy"
    entityRedundancy.EntityData.BundleName = "cisco_ios_xr"
    entityRedundancy.EntityData.ParentYangName = "notification"
    entityRedundancy.EntityData.SegmentPath = "Cisco-IOS-XR-infra-ceredundancymib-cfg:entity-redundancy"
    entityRedundancy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entityRedundancy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entityRedundancy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entityRedundancy.EntityData.Children = types.NewOrderedMap()
    entityRedundancy.EntityData.Leafs = types.NewOrderedMap()
    entityRedundancy.EntityData.Leafs.Append("switchover", types.YLeaf{"Switchover", entityRedundancy.Switchover})
    entityRedundancy.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", entityRedundancy.Enable})
    entityRedundancy.EntityData.Leafs.Append("status", types.YLeaf{"Status", entityRedundancy.Status})

    entityRedundancy.EntityData.YListKeys = []string {}

    return &(entityRedundancy.EntityData)
}

// Snmp_Notification_Entity
// Enable ENTITY-MIB notifications
type Snmp_Notification_Entity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable entityMIB notifications. The type is interface{}.
    Enable interface{}
}

func (entity *Snmp_Notification_Entity) GetEntityData() *types.CommonEntityData {
    entity.EntityData.YFilter = entity.YFilter
    entity.EntityData.YangName = "entity"
    entity.EntityData.BundleName = "cisco_ios_xr"
    entity.EntityData.ParentYangName = "notification"
    entity.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-entitymib-cfg:entity"
    entity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entity.EntityData.Children = types.NewOrderedMap()
    entity.EntityData.Leafs = types.NewOrderedMap()
    entity.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", entity.Enable})

    entity.EntityData.YListKeys = []string {}

    return &(entity.EntityData)
}

// Snmp_Notification_Rsvp
// Enable RSVP-MIB notifications
type Snmp_Notification_Rsvp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable lostFlow notification. The type is interface{}.
    LostFlow interface{}

    // Enable newFlow notification. The type is interface{}.
    NewFlow interface{}

    // Enable all RSVP notifications. The type is interface{}.
    Enable interface{}
}

func (rsvp *Snmp_Notification_Rsvp) GetEntityData() *types.CommonEntityData {
    rsvp.EntityData.YFilter = rsvp.YFilter
    rsvp.EntityData.YangName = "rsvp"
    rsvp.EntityData.BundleName = "cisco_ios_xr"
    rsvp.EntityData.ParentYangName = "notification"
    rsvp.EntityData.SegmentPath = "Cisco-IOS-XR-ip-rsvp-cfg:rsvp"
    rsvp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rsvp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rsvp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rsvp.EntityData.Children = types.NewOrderedMap()
    rsvp.EntityData.Leafs = types.NewOrderedMap()
    rsvp.EntityData.Leafs.Append("lost-flow", types.YLeaf{"LostFlow", rsvp.LostFlow})
    rsvp.EntityData.Leafs.Append("new-flow", types.YLeaf{"NewFlow", rsvp.NewFlow})
    rsvp.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", rsvp.Enable})

    rsvp.EntityData.YListKeys = []string {}

    return &(rsvp.EntityData)
}

// Snmp_Notification_ConfigMan
// CISCO-CONFIG-MAN-MIB notification configurations
type Snmp_Notification_ConfigMan struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ciscoConfigManMIB notifications. The type is interface{}.
    Enable interface{}
}

func (configMan *Snmp_Notification_ConfigMan) GetEntityData() *types.CommonEntityData {
    configMan.EntityData.YFilter = configMan.YFilter
    configMan.EntityData.YangName = "config-man"
    configMan.EntityData.BundleName = "cisco_ios_xr"
    configMan.EntityData.ParentYangName = "notification"
    configMan.EntityData.SegmentPath = "Cisco-IOS-XR-config-mibs-cfg:config-man"
    configMan.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    configMan.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    configMan.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    configMan.EntityData.Children = types.NewOrderedMap()
    configMan.EntityData.Leafs = types.NewOrderedMap()
    configMan.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", configMan.Enable})

    configMan.EntityData.YListKeys = []string {}

    return &(configMan.EntityData)
}

// Snmp_Notification_SubscriberMib
// Subscriber notification commands
type Snmp_Notification_SubscriberMib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Session aggregation.
    SessionAggregate Snmp_Notification_SubscriberMib_SessionAggregate
}

func (subscriberMib *Snmp_Notification_SubscriberMib) GetEntityData() *types.CommonEntityData {
    subscriberMib.EntityData.YFilter = subscriberMib.YFilter
    subscriberMib.EntityData.YangName = "subscriber-mib"
    subscriberMib.EntityData.BundleName = "cisco_ios_xr"
    subscriberMib.EntityData.ParentYangName = "notification"
    subscriberMib.EntityData.SegmentPath = "Cisco-IOS-XR-subscriber-session-mon-mibs-cfg:subscriber-mib"
    subscriberMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriberMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriberMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriberMib.EntityData.Children = types.NewOrderedMap()
    subscriberMib.EntityData.Children.Append("session-aggregate", types.YChild{"SessionAggregate", &subscriberMib.SessionAggregate})
    subscriberMib.EntityData.Leafs = types.NewOrderedMap()

    subscriberMib.EntityData.YListKeys = []string {}

    return &(subscriberMib.EntityData)
}

// Snmp_Notification_SubscriberMib_SessionAggregate
// Session aggregation
type Snmp_Notification_SubscriberMib_SessionAggregate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber notification at node level. The type is interface{}.
    Node interface{}

    // Subscriber notification at access interface level. The type is interface{}.
    AccessInterface interface{}
}

func (sessionAggregate *Snmp_Notification_SubscriberMib_SessionAggregate) GetEntityData() *types.CommonEntityData {
    sessionAggregate.EntityData.YFilter = sessionAggregate.YFilter
    sessionAggregate.EntityData.YangName = "session-aggregate"
    sessionAggregate.EntityData.BundleName = "cisco_ios_xr"
    sessionAggregate.EntityData.ParentYangName = "subscriber-mib"
    sessionAggregate.EntityData.SegmentPath = "session-aggregate"
    sessionAggregate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessionAggregate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessionAggregate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessionAggregate.EntityData.Children = types.NewOrderedMap()
    sessionAggregate.EntityData.Leafs = types.NewOrderedMap()
    sessionAggregate.EntityData.Leafs.Append("node", types.YLeaf{"Node", sessionAggregate.Node})
    sessionAggregate.EntityData.Leafs.Append("access-interface", types.YLeaf{"AccessInterface", sessionAggregate.AccessInterface})

    sessionAggregate.EntityData.YListKeys = []string {}

    return &(sessionAggregate.EntityData)
}

// Snmp_Notification_OpticalOts
// CISCO-OPTICAL-OTS-MIB notification configuration
type Snmp_Notification_OpticalOts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable OpticalOtsmib notifications. The type is interface{}.
    Enable interface{}
}

func (opticalOts *Snmp_Notification_OpticalOts) GetEntityData() *types.CommonEntityData {
    opticalOts.EntityData.YFilter = opticalOts.YFilter
    opticalOts.EntityData.YangName = "optical-ots"
    opticalOts.EntityData.BundleName = "cisco_ios_xr"
    opticalOts.EntityData.ParentYangName = "notification"
    opticalOts.EntityData.SegmentPath = "Cisco-IOS-XR-opticalotsmib-cfg:optical-ots"
    opticalOts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    opticalOts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    opticalOts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    opticalOts.EntityData.Children = types.NewOrderedMap()
    opticalOts.EntityData.Leafs = types.NewOrderedMap()
    opticalOts.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", opticalOts.Enable})

    opticalOts.EntityData.YListKeys = []string {}

    return &(opticalOts.EntityData)
}

// Snmp_Notification_AddresspoolMib
// Enable SNMP daps traps
type Snmp_Notification_AddresspoolMib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable SNMP Address Pool High Threshold trap. The type is bool.
    High interface{}

    // Enable SNMP Address Pool Low Threshold trap. The type is bool.
    Low interface{}
}

func (addresspoolMib *Snmp_Notification_AddresspoolMib) GetEntityData() *types.CommonEntityData {
    addresspoolMib.EntityData.YFilter = addresspoolMib.YFilter
    addresspoolMib.EntityData.YangName = "addresspool-mib"
    addresspoolMib.EntityData.BundleName = "cisco_ios_xr"
    addresspoolMib.EntityData.ParentYangName = "notification"
    addresspoolMib.EntityData.SegmentPath = "Cisco-IOS-XR-ip-daps-mib-cfg:addresspool-mib"
    addresspoolMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addresspoolMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addresspoolMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addresspoolMib.EntityData.Children = types.NewOrderedMap()
    addresspoolMib.EntityData.Leafs = types.NewOrderedMap()
    addresspoolMib.EntityData.Leafs.Append("high", types.YLeaf{"High", addresspoolMib.High})
    addresspoolMib.EntityData.Leafs.Append("low", types.YLeaf{"Low", addresspoolMib.Low})

    addresspoolMib.EntityData.YListKeys = []string {}

    return &(addresspoolMib.EntityData)
}

// Snmp_Notification_Diametermib
// Enable SNMP diameter traps
type Snmp_Notification_Diametermib struct {
    EntityData types.CommonEntityData
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

func (diametermib *Snmp_Notification_Diametermib) GetEntityData() *types.CommonEntityData {
    diametermib.EntityData.YFilter = diametermib.YFilter
    diametermib.EntityData.YangName = "diametermib"
    diametermib.EntityData.BundleName = "cisco_ios_xr"
    diametermib.EntityData.ParentYangName = "notification"
    diametermib.EntityData.SegmentPath = "Cisco-IOS-XR-aaa-diameter-base-mib-cfg:diametermib"
    diametermib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    diametermib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    diametermib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    diametermib.EntityData.Children = types.NewOrderedMap()
    diametermib.EntityData.Leafs = types.NewOrderedMap()
    diametermib.EntityData.Leafs.Append("protocolerror", types.YLeaf{"Protocolerror", diametermib.Protocolerror})
    diametermib.EntityData.Leafs.Append("permanentfail", types.YLeaf{"Permanentfail", diametermib.Permanentfail})
    diametermib.EntityData.Leafs.Append("peerdown", types.YLeaf{"Peerdown", diametermib.Peerdown})
    diametermib.EntityData.Leafs.Append("peerup", types.YLeaf{"Peerup", diametermib.Peerup})
    diametermib.EntityData.Leafs.Append("transientfail", types.YLeaf{"Transientfail", diametermib.Transientfail})

    diametermib.EntityData.YListKeys = []string {}

    return &(diametermib.EntityData)
}

// Snmp_Notification_Rf
// CISCO-RF-MIB notification configuration
type Snmp_Notification_Rf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ciscoRFMIB notifications. The type is interface{}.
    Enable interface{}
}

func (rf *Snmp_Notification_Rf) GetEntityData() *types.CommonEntityData {
    rf.EntityData.YFilter = rf.YFilter
    rf.EntityData.YangName = "rf"
    rf.EntityData.BundleName = "cisco_ios_xr"
    rf.EntityData.ParentYangName = "notification"
    rf.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-mib-rfmib-cfg:rf"
    rf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rf.EntityData.Children = types.NewOrderedMap()
    rf.EntityData.Leafs = types.NewOrderedMap()
    rf.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", rf.Enable})

    rf.EntityData.YListKeys = []string {}

    return &(rf.EntityData)
}

// Snmp_Notification_EntityState
// ENTITY-STATE-MIB notification configuration
type Snmp_Notification_EntityState struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ceStateExtStandbySwitchover notification. The type is interface{}.
    Switchover interface{}

    // Enable entStateOperEnable and entStateOperDisable notifications. The type
    // is interface{}.
    OperStatus interface{}
}

func (entityState *Snmp_Notification_EntityState) GetEntityData() *types.CommonEntityData {
    entityState.EntityData.YFilter = entityState.YFilter
    entityState.EntityData.YangName = "entity-state"
    entityState.EntityData.BundleName = "cisco_ios_xr"
    entityState.EntityData.ParentYangName = "notification"
    entityState.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-entstatemib-cfg:entity-state"
    entityState.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entityState.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entityState.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entityState.EntityData.Children = types.NewOrderedMap()
    entityState.EntityData.Leafs = types.NewOrderedMap()
    entityState.EntityData.Leafs.Append("switchover", types.YLeaf{"Switchover", entityState.Switchover})
    entityState.EntityData.Leafs.Append("oper-status", types.YLeaf{"OperStatus", entityState.OperStatus})

    entityState.EntityData.YListKeys = []string {}

    return &(entityState.EntityData)
}

// Snmp_Correlator
// Configure properties of the trap correlator
type Snmp_Correlator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure size of the correlator buffer. The type is interface{} with
    // range: 1024..52428800.
    BufferSize interface{}

    // Table of configured rules.
    Rules Snmp_Correlator_Rules

    // Table of configured rulesets.
    RuleSets Snmp_Correlator_RuleSets
}

func (correlator *Snmp_Correlator) GetEntityData() *types.CommonEntityData {
    correlator.EntityData.YFilter = correlator.YFilter
    correlator.EntityData.YangName = "correlator"
    correlator.EntityData.BundleName = "cisco_ios_xr"
    correlator.EntityData.ParentYangName = "snmp"
    correlator.EntityData.SegmentPath = "correlator"
    correlator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    correlator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    correlator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    correlator.EntityData.Children = types.NewOrderedMap()
    correlator.EntityData.Children.Append("rules", types.YChild{"Rules", &correlator.Rules})
    correlator.EntityData.Children.Append("rule-sets", types.YChild{"RuleSets", &correlator.RuleSets})
    correlator.EntityData.Leafs = types.NewOrderedMap()
    correlator.EntityData.Leafs.Append("buffer-size", types.YLeaf{"BufferSize", correlator.BufferSize})

    correlator.EntityData.YListKeys = []string {}

    return &(correlator.EntityData)
}

// Snmp_Correlator_Rules
// Table of configured rules
type Snmp_Correlator_Rules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rule name. The type is slice of Snmp_Correlator_Rules_Rule.
    Rule []*Snmp_Correlator_Rules_Rule
}

func (rules *Snmp_Correlator_Rules) GetEntityData() *types.CommonEntityData {
    rules.EntityData.YFilter = rules.YFilter
    rules.EntityData.YangName = "rules"
    rules.EntityData.BundleName = "cisco_ios_xr"
    rules.EntityData.ParentYangName = "correlator"
    rules.EntityData.SegmentPath = "rules"
    rules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rules.EntityData.Children = types.NewOrderedMap()
    rules.EntityData.Children.Append("rule", types.YChild{"Rule", nil})
    for i := range rules.Rule {
        rules.EntityData.Children.Append(types.GetSegmentPath(rules.Rule[i]), types.YChild{"Rule", rules.Rule[i]})
    }
    rules.EntityData.Leafs = types.NewOrderedMap()

    rules.EntityData.YListKeys = []string {}

    return &(rules.EntityData)
}

// Snmp_Correlator_Rules_Rule
// Rule name
type Snmp_Correlator_Rules_Rule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Rule name. The type is string with length: 1..32.
    Name interface{}

    // Timeout (time to wait for active correlation) in milliseconds. The type is
    // interface{} with range: 1..600000. Units are millisecond.
    Timeout interface{}

    // Table of configured rootcause (only one entry allowed).
    RootCauses Snmp_Correlator_Rules_Rule_RootCauses

    // Table of configured non-rootcause.
    NonRootCauses Snmp_Correlator_Rules_Rule_NonRootCauses

    // Applied to the Rule or Ruleset.
    AppliedTo Snmp_Correlator_Rules_Rule_AppliedTo
}

func (rule *Snmp_Correlator_Rules_Rule) GetEntityData() *types.CommonEntityData {
    rule.EntityData.YFilter = rule.YFilter
    rule.EntityData.YangName = "rule"
    rule.EntityData.BundleName = "cisco_ios_xr"
    rule.EntityData.ParentYangName = "rules"
    rule.EntityData.SegmentPath = "rule" + types.AddKeyToken(rule.Name, "name")
    rule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rule.EntityData.Children = types.NewOrderedMap()
    rule.EntityData.Children.Append("root-causes", types.YChild{"RootCauses", &rule.RootCauses})
    rule.EntityData.Children.Append("non-root-causes", types.YChild{"NonRootCauses", &rule.NonRootCauses})
    rule.EntityData.Children.Append("applied-to", types.YChild{"AppliedTo", &rule.AppliedTo})
    rule.EntityData.Leafs = types.NewOrderedMap()
    rule.EntityData.Leafs.Append("name", types.YLeaf{"Name", rule.Name})
    rule.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", rule.Timeout})

    rule.EntityData.YListKeys = []string {"Name"}

    return &(rule.EntityData)
}

// Snmp_Correlator_Rules_Rule_RootCauses
// Table of configured rootcause (only one entry
// allowed)
type Snmp_Correlator_Rules_Rule_RootCauses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The rootcause - maximum of one can be configured per rule. The type is
    // slice of Snmp_Correlator_Rules_Rule_RootCauses_RootCause.
    RootCause []*Snmp_Correlator_Rules_Rule_RootCauses_RootCause
}

func (rootCauses *Snmp_Correlator_Rules_Rule_RootCauses) GetEntityData() *types.CommonEntityData {
    rootCauses.EntityData.YFilter = rootCauses.YFilter
    rootCauses.EntityData.YangName = "root-causes"
    rootCauses.EntityData.BundleName = "cisco_ios_xr"
    rootCauses.EntityData.ParentYangName = "rule"
    rootCauses.EntityData.SegmentPath = "root-causes"
    rootCauses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rootCauses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rootCauses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rootCauses.EntityData.Children = types.NewOrderedMap()
    rootCauses.EntityData.Children.Append("root-cause", types.YChild{"RootCause", nil})
    for i := range rootCauses.RootCause {
        rootCauses.EntityData.Children.Append(types.GetSegmentPath(rootCauses.RootCause[i]), types.YChild{"RootCause", rootCauses.RootCause[i]})
    }
    rootCauses.EntityData.Leafs = types.NewOrderedMap()

    rootCauses.EntityData.YListKeys = []string {}

    return &(rootCauses.EntityData)
}

// Snmp_Correlator_Rules_Rule_RootCauses_RootCause
// The rootcause - maximum of one can be
// configured per rule
type Snmp_Correlator_Rules_Rule_RootCauses_RootCause struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OID of rootcause trap (dotted decimal). The type
    // is string.
    Oid interface{}

    // Create rootcause. The type is interface{}.
    Created interface{}

    // Varbinds to match.
    VarBinds Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds
}

func (rootCause *Snmp_Correlator_Rules_Rule_RootCauses_RootCause) GetEntityData() *types.CommonEntityData {
    rootCause.EntityData.YFilter = rootCause.YFilter
    rootCause.EntityData.YangName = "root-cause"
    rootCause.EntityData.BundleName = "cisco_ios_xr"
    rootCause.EntityData.ParentYangName = "root-causes"
    rootCause.EntityData.SegmentPath = "root-cause" + types.AddKeyToken(rootCause.Oid, "oid")
    rootCause.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rootCause.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rootCause.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rootCause.EntityData.Children = types.NewOrderedMap()
    rootCause.EntityData.Children.Append("var-binds", types.YChild{"VarBinds", &rootCause.VarBinds})
    rootCause.EntityData.Leafs = types.NewOrderedMap()
    rootCause.EntityData.Leafs.Append("oid", types.YLeaf{"Oid", rootCause.Oid})
    rootCause.EntityData.Leafs.Append("created", types.YLeaf{"Created", rootCause.Created})

    rootCause.EntityData.YListKeys = []string {"Oid"}

    return &(rootCause.EntityData)
}

// Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds
// Varbinds to match
type Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Varbind match conditions. The type is slice of
    // Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds_VarBind.
    VarBind []*Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds_VarBind
}

func (varBinds *Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds) GetEntityData() *types.CommonEntityData {
    varBinds.EntityData.YFilter = varBinds.YFilter
    varBinds.EntityData.YangName = "var-binds"
    varBinds.EntityData.BundleName = "cisco_ios_xr"
    varBinds.EntityData.ParentYangName = "root-cause"
    varBinds.EntityData.SegmentPath = "var-binds"
    varBinds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    varBinds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    varBinds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    varBinds.EntityData.Children = types.NewOrderedMap()
    varBinds.EntityData.Children.Append("var-bind", types.YChild{"VarBind", nil})
    for i := range varBinds.VarBind {
        varBinds.EntityData.Children.Append(types.GetSegmentPath(varBinds.VarBind[i]), types.YChild{"VarBind", varBinds.VarBind[i]})
    }
    varBinds.EntityData.Leafs = types.NewOrderedMap()

    varBinds.EntityData.YListKeys = []string {}

    return &(varBinds.EntityData)
}

// Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds_VarBind
// Varbind match conditions
type Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds_VarBind struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OID of varbind (dotted decimal). The type is
    // string.
    Oid interface{}

    // VarBind match conditions.
    Match Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds_VarBind_Match
}

func (varBind *Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds_VarBind) GetEntityData() *types.CommonEntityData {
    varBind.EntityData.YFilter = varBind.YFilter
    varBind.EntityData.YangName = "var-bind"
    varBind.EntityData.BundleName = "cisco_ios_xr"
    varBind.EntityData.ParentYangName = "var-binds"
    varBind.EntityData.SegmentPath = "var-bind" + types.AddKeyToken(varBind.Oid, "oid")
    varBind.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    varBind.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    varBind.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    varBind.EntityData.Children = types.NewOrderedMap()
    varBind.EntityData.Children.Append("match", types.YChild{"Match", &varBind.Match})
    varBind.EntityData.Leafs = types.NewOrderedMap()
    varBind.EntityData.Leafs.Append("oid", types.YLeaf{"Oid", varBind.Oid})

    varBind.EntityData.YListKeys = []string {"Oid"}

    return &(varBind.EntityData)
}

// Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds_VarBind_Match
// VarBind match conditions
type Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds_VarBind_Match struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Regular Expression to match value. The type is string.
    Value interface{}

    // Regular Expression to match index. The type is string.
    Index interface{}
}

func (match *Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds_VarBind_Match) GetEntityData() *types.CommonEntityData {
    match.EntityData.YFilter = match.YFilter
    match.EntityData.YangName = "match"
    match.EntityData.BundleName = "cisco_ios_xr"
    match.EntityData.ParentYangName = "var-bind"
    match.EntityData.SegmentPath = "match"
    match.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    match.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    match.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    match.EntityData.Children = types.NewOrderedMap()
    match.EntityData.Leafs = types.NewOrderedMap()
    match.EntityData.Leafs.Append("value", types.YLeaf{"Value", match.Value})
    match.EntityData.Leafs.Append("index", types.YLeaf{"Index", match.Index})

    match.EntityData.YListKeys = []string {}

    return &(match.EntityData)
}

// Snmp_Correlator_Rules_Rule_NonRootCauses
// Table of configured non-rootcause
type Snmp_Correlator_Rules_Rule_NonRootCauses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A non-rootcause. The type is slice of
    // Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause.
    NonRootCause []*Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause
}

func (nonRootCauses *Snmp_Correlator_Rules_Rule_NonRootCauses) GetEntityData() *types.CommonEntityData {
    nonRootCauses.EntityData.YFilter = nonRootCauses.YFilter
    nonRootCauses.EntityData.YangName = "non-root-causes"
    nonRootCauses.EntityData.BundleName = "cisco_ios_xr"
    nonRootCauses.EntityData.ParentYangName = "rule"
    nonRootCauses.EntityData.SegmentPath = "non-root-causes"
    nonRootCauses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonRootCauses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonRootCauses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonRootCauses.EntityData.Children = types.NewOrderedMap()
    nonRootCauses.EntityData.Children.Append("non-root-cause", types.YChild{"NonRootCause", nil})
    for i := range nonRootCauses.NonRootCause {
        nonRootCauses.EntityData.Children.Append(types.GetSegmentPath(nonRootCauses.NonRootCause[i]), types.YChild{"NonRootCause", nonRootCauses.NonRootCause[i]})
    }
    nonRootCauses.EntityData.Leafs = types.NewOrderedMap()

    nonRootCauses.EntityData.YListKeys = []string {}

    return &(nonRootCauses.EntityData)
}

// Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause
// A non-rootcause
type Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OID of nonrootcause trap (dotted decimal). The
    // type is string.
    Oid interface{}

    // Create nonrootcause. The type is interface{}.
    Created interface{}

    // Varbinds to match.
    VarBinds Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds
}

func (nonRootCause *Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause) GetEntityData() *types.CommonEntityData {
    nonRootCause.EntityData.YFilter = nonRootCause.YFilter
    nonRootCause.EntityData.YangName = "non-root-cause"
    nonRootCause.EntityData.BundleName = "cisco_ios_xr"
    nonRootCause.EntityData.ParentYangName = "non-root-causes"
    nonRootCause.EntityData.SegmentPath = "non-root-cause" + types.AddKeyToken(nonRootCause.Oid, "oid")
    nonRootCause.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonRootCause.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonRootCause.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonRootCause.EntityData.Children = types.NewOrderedMap()
    nonRootCause.EntityData.Children.Append("var-binds", types.YChild{"VarBinds", &nonRootCause.VarBinds})
    nonRootCause.EntityData.Leafs = types.NewOrderedMap()
    nonRootCause.EntityData.Leafs.Append("oid", types.YLeaf{"Oid", nonRootCause.Oid})
    nonRootCause.EntityData.Leafs.Append("created", types.YLeaf{"Created", nonRootCause.Created})

    nonRootCause.EntityData.YListKeys = []string {"Oid"}

    return &(nonRootCause.EntityData)
}

// Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds
// Varbinds to match
type Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Varbind match conditions. The type is slice of
    // Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds_VarBind.
    VarBind []*Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds_VarBind
}

func (varBinds *Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds) GetEntityData() *types.CommonEntityData {
    varBinds.EntityData.YFilter = varBinds.YFilter
    varBinds.EntityData.YangName = "var-binds"
    varBinds.EntityData.BundleName = "cisco_ios_xr"
    varBinds.EntityData.ParentYangName = "non-root-cause"
    varBinds.EntityData.SegmentPath = "var-binds"
    varBinds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    varBinds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    varBinds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    varBinds.EntityData.Children = types.NewOrderedMap()
    varBinds.EntityData.Children.Append("var-bind", types.YChild{"VarBind", nil})
    for i := range varBinds.VarBind {
        varBinds.EntityData.Children.Append(types.GetSegmentPath(varBinds.VarBind[i]), types.YChild{"VarBind", varBinds.VarBind[i]})
    }
    varBinds.EntityData.Leafs = types.NewOrderedMap()

    varBinds.EntityData.YListKeys = []string {}

    return &(varBinds.EntityData)
}

// Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds_VarBind
// Varbind match conditions
type Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds_VarBind struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. OID of varbind (dotted decimal). The type is
    // string.
    Oid interface{}

    // VarBind match conditions.
    Match Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds_VarBind_Match
}

func (varBind *Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds_VarBind) GetEntityData() *types.CommonEntityData {
    varBind.EntityData.YFilter = varBind.YFilter
    varBind.EntityData.YangName = "var-bind"
    varBind.EntityData.BundleName = "cisco_ios_xr"
    varBind.EntityData.ParentYangName = "var-binds"
    varBind.EntityData.SegmentPath = "var-bind" + types.AddKeyToken(varBind.Oid, "oid")
    varBind.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    varBind.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    varBind.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    varBind.EntityData.Children = types.NewOrderedMap()
    varBind.EntityData.Children.Append("match", types.YChild{"Match", &varBind.Match})
    varBind.EntityData.Leafs = types.NewOrderedMap()
    varBind.EntityData.Leafs.Append("oid", types.YLeaf{"Oid", varBind.Oid})

    varBind.EntityData.YListKeys = []string {"Oid"}

    return &(varBind.EntityData)
}

// Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds_VarBind_Match
// VarBind match conditions
type Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds_VarBind_Match struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Regular Expression to match value. The type is string.
    Value interface{}

    // Regular Expression to match index. The type is string.
    Index interface{}
}

func (match *Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds_VarBind_Match) GetEntityData() *types.CommonEntityData {
    match.EntityData.YFilter = match.YFilter
    match.EntityData.YangName = "match"
    match.EntityData.BundleName = "cisco_ios_xr"
    match.EntityData.ParentYangName = "var-bind"
    match.EntityData.SegmentPath = "match"
    match.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    match.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    match.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    match.EntityData.Children = types.NewOrderedMap()
    match.EntityData.Leafs = types.NewOrderedMap()
    match.EntityData.Leafs.Append("value", types.YLeaf{"Value", match.Value})
    match.EntityData.Leafs.Append("index", types.YLeaf{"Index", match.Index})

    match.EntityData.YListKeys = []string {}

    return &(match.EntityData)
}

// Snmp_Correlator_Rules_Rule_AppliedTo
// Applied to the Rule or Ruleset
type Snmp_Correlator_Rules_Rule_AppliedTo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Apply to all of the device. The type is interface{}.
    All interface{}

    // Table of configured hosts to apply rules to.
    Hosts Snmp_Correlator_Rules_Rule_AppliedTo_Hosts
}

func (appliedTo *Snmp_Correlator_Rules_Rule_AppliedTo) GetEntityData() *types.CommonEntityData {
    appliedTo.EntityData.YFilter = appliedTo.YFilter
    appliedTo.EntityData.YangName = "applied-to"
    appliedTo.EntityData.BundleName = "cisco_ios_xr"
    appliedTo.EntityData.ParentYangName = "rule"
    appliedTo.EntityData.SegmentPath = "applied-to"
    appliedTo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appliedTo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appliedTo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appliedTo.EntityData.Children = types.NewOrderedMap()
    appliedTo.EntityData.Children.Append("hosts", types.YChild{"Hosts", &appliedTo.Hosts})
    appliedTo.EntityData.Leafs = types.NewOrderedMap()
    appliedTo.EntityData.Leafs.Append("all", types.YLeaf{"All", appliedTo.All})

    appliedTo.EntityData.YListKeys = []string {}

    return &(appliedTo.EntityData)
}

// Snmp_Correlator_Rules_Rule_AppliedTo_Hosts
// Table of configured hosts to apply rules to
type Snmp_Correlator_Rules_Rule_AppliedTo_Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A destination host. The type is slice of
    // Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host.
    Host []*Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host
}

func (hosts *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts) GetEntityData() *types.CommonEntityData {
    hosts.EntityData.YFilter = hosts.YFilter
    hosts.EntityData.YangName = "hosts"
    hosts.EntityData.BundleName = "cisco_ios_xr"
    hosts.EntityData.ParentYangName = "applied-to"
    hosts.EntityData.SegmentPath = "hosts"
    hosts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hosts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hosts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hosts.EntityData.Children = types.NewOrderedMap()
    hosts.EntityData.Children.Append("host", types.YChild{"Host", nil})
    for i := range hosts.Host {
        hosts.EntityData.Children.Append(types.GetSegmentPath(hosts.Host[i]), types.YChild{"Host", hosts.Host[i]})
    }
    hosts.EntityData.Leafs = types.NewOrderedMap()

    hosts.EntityData.YListKeys = []string {}

    return &(hosts.EntityData)
}

// Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host
// A destination host
type Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host struct {
    EntityData types.CommonEntityData
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

func (host *Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host) GetEntityData() *types.CommonEntityData {
    host.EntityData.YFilter = host.YFilter
    host.EntityData.YangName = "host"
    host.EntityData.BundleName = "cisco_ios_xr"
    host.EntityData.ParentYangName = "hosts"
    host.EntityData.SegmentPath = "host" + types.AddKeyToken(host.IpAddress, "ip-address") + types.AddKeyToken(host.Port, "port")
    host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    host.EntityData.Children = types.NewOrderedMap()
    host.EntityData.Leafs = types.NewOrderedMap()
    host.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", host.IpAddress})
    host.EntityData.Leafs.Append("port", types.YLeaf{"Port", host.Port})

    host.EntityData.YListKeys = []string {"IpAddress", "Port"}

    return &(host.EntityData)
}

// Snmp_Correlator_RuleSets
// Table of configured rulesets
type Snmp_Correlator_RuleSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ruleset name. The type is slice of Snmp_Correlator_RuleSets_RuleSet.
    RuleSet []*Snmp_Correlator_RuleSets_RuleSet
}

func (ruleSets *Snmp_Correlator_RuleSets) GetEntityData() *types.CommonEntityData {
    ruleSets.EntityData.YFilter = ruleSets.YFilter
    ruleSets.EntityData.YangName = "rule-sets"
    ruleSets.EntityData.BundleName = "cisco_ios_xr"
    ruleSets.EntityData.ParentYangName = "correlator"
    ruleSets.EntityData.SegmentPath = "rule-sets"
    ruleSets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSets.EntityData.Children = types.NewOrderedMap()
    ruleSets.EntityData.Children.Append("rule-set", types.YChild{"RuleSet", nil})
    for i := range ruleSets.RuleSet {
        ruleSets.EntityData.Children.Append(types.GetSegmentPath(ruleSets.RuleSet[i]), types.YChild{"RuleSet", ruleSets.RuleSet[i]})
    }
    ruleSets.EntityData.Leafs = types.NewOrderedMap()

    ruleSets.EntityData.YListKeys = []string {}

    return &(ruleSets.EntityData)
}

// Snmp_Correlator_RuleSets_RuleSet
// Ruleset name
type Snmp_Correlator_RuleSets_RuleSet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Ruleset name. The type is string with length:
    // 1..32.
    Name interface{}

    // Table of configured rulenames.
    Rulenames Snmp_Correlator_RuleSets_RuleSet_Rulenames

    // Applied to the Rule or Ruleset.
    AppliedTo Snmp_Correlator_RuleSets_RuleSet_AppliedTo
}

func (ruleSet *Snmp_Correlator_RuleSets_RuleSet) GetEntityData() *types.CommonEntityData {
    ruleSet.EntityData.YFilter = ruleSet.YFilter
    ruleSet.EntityData.YangName = "rule-set"
    ruleSet.EntityData.BundleName = "cisco_ios_xr"
    ruleSet.EntityData.ParentYangName = "rule-sets"
    ruleSet.EntityData.SegmentPath = "rule-set" + types.AddKeyToken(ruleSet.Name, "name")
    ruleSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSet.EntityData.Children = types.NewOrderedMap()
    ruleSet.EntityData.Children.Append("rulenames", types.YChild{"Rulenames", &ruleSet.Rulenames})
    ruleSet.EntityData.Children.Append("applied-to", types.YChild{"AppliedTo", &ruleSet.AppliedTo})
    ruleSet.EntityData.Leafs = types.NewOrderedMap()
    ruleSet.EntityData.Leafs.Append("name", types.YLeaf{"Name", ruleSet.Name})

    ruleSet.EntityData.YListKeys = []string {"Name"}

    return &(ruleSet.EntityData)
}

// Snmp_Correlator_RuleSets_RuleSet_Rulenames
// Table of configured rulenames
type Snmp_Correlator_RuleSets_RuleSet_Rulenames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A rulename. The type is slice of
    // Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename.
    Rulename []*Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename
}

func (rulenames *Snmp_Correlator_RuleSets_RuleSet_Rulenames) GetEntityData() *types.CommonEntityData {
    rulenames.EntityData.YFilter = rulenames.YFilter
    rulenames.EntityData.YangName = "rulenames"
    rulenames.EntityData.BundleName = "cisco_ios_xr"
    rulenames.EntityData.ParentYangName = "rule-set"
    rulenames.EntityData.SegmentPath = "rulenames"
    rulenames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rulenames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rulenames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rulenames.EntityData.Children = types.NewOrderedMap()
    rulenames.EntityData.Children.Append("rulename", types.YChild{"Rulename", nil})
    for i := range rulenames.Rulename {
        rulenames.EntityData.Children.Append(types.GetSegmentPath(rulenames.Rulename[i]), types.YChild{"Rulename", rulenames.Rulename[i]})
    }
    rulenames.EntityData.Leafs = types.NewOrderedMap()

    rulenames.EntityData.YListKeys = []string {}

    return &(rulenames.EntityData)
}

// Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename
// A rulename
type Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Rule name. The type is string with length: 1..32.
    Rulename interface{}
}

func (rulename *Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename) GetEntityData() *types.CommonEntityData {
    rulename.EntityData.YFilter = rulename.YFilter
    rulename.EntityData.YangName = "rulename"
    rulename.EntityData.BundleName = "cisco_ios_xr"
    rulename.EntityData.ParentYangName = "rulenames"
    rulename.EntityData.SegmentPath = "rulename" + types.AddKeyToken(rulename.Rulename, "rulename")
    rulename.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rulename.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rulename.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rulename.EntityData.Children = types.NewOrderedMap()
    rulename.EntityData.Leafs = types.NewOrderedMap()
    rulename.EntityData.Leafs.Append("rulename", types.YLeaf{"Rulename", rulename.Rulename})

    rulename.EntityData.YListKeys = []string {"Rulename"}

    return &(rulename.EntityData)
}

// Snmp_Correlator_RuleSets_RuleSet_AppliedTo
// Applied to the Rule or Ruleset
type Snmp_Correlator_RuleSets_RuleSet_AppliedTo struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Apply to all of the device. The type is interface{}.
    All interface{}

    // Table of configured hosts to apply rules to.
    Hosts Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts
}

func (appliedTo *Snmp_Correlator_RuleSets_RuleSet_AppliedTo) GetEntityData() *types.CommonEntityData {
    appliedTo.EntityData.YFilter = appliedTo.YFilter
    appliedTo.EntityData.YangName = "applied-to"
    appliedTo.EntityData.BundleName = "cisco_ios_xr"
    appliedTo.EntityData.ParentYangName = "rule-set"
    appliedTo.EntityData.SegmentPath = "applied-to"
    appliedTo.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    appliedTo.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    appliedTo.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    appliedTo.EntityData.Children = types.NewOrderedMap()
    appliedTo.EntityData.Children.Append("hosts", types.YChild{"Hosts", &appliedTo.Hosts})
    appliedTo.EntityData.Leafs = types.NewOrderedMap()
    appliedTo.EntityData.Leafs.Append("all", types.YLeaf{"All", appliedTo.All})

    appliedTo.EntityData.YListKeys = []string {}

    return &(appliedTo.EntityData)
}

// Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts
// Table of configured hosts to apply rules to
type Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A destination host. The type is slice of
    // Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host.
    Host []*Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host
}

func (hosts *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts) GetEntityData() *types.CommonEntityData {
    hosts.EntityData.YFilter = hosts.YFilter
    hosts.EntityData.YangName = "hosts"
    hosts.EntityData.BundleName = "cisco_ios_xr"
    hosts.EntityData.ParentYangName = "applied-to"
    hosts.EntityData.SegmentPath = "hosts"
    hosts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hosts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hosts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hosts.EntityData.Children = types.NewOrderedMap()
    hosts.EntityData.Children.Append("host", types.YChild{"Host", nil})
    for i := range hosts.Host {
        hosts.EntityData.Children.Append(types.GetSegmentPath(hosts.Host[i]), types.YChild{"Host", hosts.Host[i]})
    }
    hosts.EntityData.Leafs = types.NewOrderedMap()

    hosts.EntityData.YListKeys = []string {}

    return &(hosts.EntityData)
}

// Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host
// A destination host
type Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host struct {
    EntityData types.CommonEntityData
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

func (host *Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host) GetEntityData() *types.CommonEntityData {
    host.EntityData.YFilter = host.YFilter
    host.EntityData.YangName = "host"
    host.EntityData.BundleName = "cisco_ios_xr"
    host.EntityData.ParentYangName = "hosts"
    host.EntityData.SegmentPath = "host" + types.AddKeyToken(host.IpAddress, "ip-address") + types.AddKeyToken(host.Port, "port")
    host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    host.EntityData.Children = types.NewOrderedMap()
    host.EntityData.Leafs = types.NewOrderedMap()
    host.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", host.IpAddress})
    host.EntityData.Leafs.Append("port", types.YLeaf{"Port", host.Port})

    host.EntityData.YListKeys = []string {"IpAddress", "Port"}

    return &(host.EntityData)
}

// Snmp_BulkStats
// SNMP bulk stats configuration commands
type Snmp_BulkStats struct {
    EntityData types.CommonEntityData
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

func (bulkStats *Snmp_BulkStats) GetEntityData() *types.CommonEntityData {
    bulkStats.EntityData.YFilter = bulkStats.YFilter
    bulkStats.EntityData.YangName = "bulk-stats"
    bulkStats.EntityData.BundleName = "cisco_ios_xr"
    bulkStats.EntityData.ParentYangName = "snmp"
    bulkStats.EntityData.SegmentPath = "bulk-stats"
    bulkStats.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bulkStats.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bulkStats.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bulkStats.EntityData.Children = types.NewOrderedMap()
    bulkStats.EntityData.Children.Append("schemas", types.YChild{"Schemas", &bulkStats.Schemas})
    bulkStats.EntityData.Children.Append("objects", types.YChild{"Objects", &bulkStats.Objects})
    bulkStats.EntityData.Children.Append("transfers", types.YChild{"Transfers", &bulkStats.Transfers})
    bulkStats.EntityData.Leafs = types.NewOrderedMap()
    bulkStats.EntityData.Leafs.Append("memory", types.YLeaf{"Memory", bulkStats.Memory})

    bulkStats.EntityData.YListKeys = []string {}

    return &(bulkStats.EntityData)
}

// Snmp_BulkStats_Schemas
// Configure schema definition
type Snmp_BulkStats_Schemas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the Schema. The type is slice of Snmp_BulkStats_Schemas_Schema.
    Schema []*Snmp_BulkStats_Schemas_Schema
}

func (schemas *Snmp_BulkStats_Schemas) GetEntityData() *types.CommonEntityData {
    schemas.EntityData.YFilter = schemas.YFilter
    schemas.EntityData.YangName = "schemas"
    schemas.EntityData.BundleName = "cisco_ios_xr"
    schemas.EntityData.ParentYangName = "bulk-stats"
    schemas.EntityData.SegmentPath = "schemas"
    schemas.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    schemas.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    schemas.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    schemas.EntityData.Children = types.NewOrderedMap()
    schemas.EntityData.Children.Append("schema", types.YChild{"Schema", nil})
    for i := range schemas.Schema {
        schemas.EntityData.Children.Append(types.GetSegmentPath(schemas.Schema[i]), types.YChild{"Schema", schemas.Schema[i]})
    }
    schemas.EntityData.Leafs = types.NewOrderedMap()

    schemas.EntityData.YListKeys = []string {}

    return &(schemas.EntityData)
}

// Snmp_BulkStats_Schemas_Schema
// The name of the Schema
type Snmp_BulkStats_Schemas_Schema struct {
    EntityData types.CommonEntityData
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

func (schema *Snmp_BulkStats_Schemas_Schema) GetEntityData() *types.CommonEntityData {
    schema.EntityData.YFilter = schema.YFilter
    schema.EntityData.YangName = "schema"
    schema.EntityData.BundleName = "cisco_ios_xr"
    schema.EntityData.ParentYangName = "schemas"
    schema.EntityData.SegmentPath = "schema" + types.AddKeyToken(schema.SchemaName, "schema-name")
    schema.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    schema.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    schema.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    schema.EntityData.Children = types.NewOrderedMap()
    schema.EntityData.Children.Append("instance", types.YChild{"Instance", &schema.Instance})
    schema.EntityData.Leafs = types.NewOrderedMap()
    schema.EntityData.Leafs.Append("schema-name", types.YLeaf{"SchemaName", schema.SchemaName})
    schema.EntityData.Leafs.Append("type", types.YLeaf{"Type", schema.Type})
    schema.EntityData.Leafs.Append("schema-object-list", types.YLeaf{"SchemaObjectList", schema.SchemaObjectList})
    schema.EntityData.Leafs.Append("poll-interval", types.YLeaf{"PollInterval", schema.PollInterval})

    schema.EntityData.YListKeys = []string {"SchemaName"}

    return &(schema.EntityData)
}

// Snmp_BulkStats_Schemas_Schema_Instance
// Object instance information
// This type is a presence type.
type Snmp_BulkStats_Schemas_Schema_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

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
    // 0..4294967295. This attribute is mandatory.
    Max interface{}

    // Include all the subinterface. The type is bool. This attribute is
    // mandatory.
    SubInterface interface{}
}

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "schema"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = types.NewOrderedMap()
    instance.EntityData.Leafs = types.NewOrderedMap()
    instance.EntityData.Leafs.Append("type", types.YLeaf{"Type", instance.Type})
    instance.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", instance.Instance})
    instance.EntityData.Leafs.Append("start", types.YLeaf{"Start", instance.Start})
    instance.EntityData.Leafs.Append("end", types.YLeaf{"End", instance.End})
    instance.EntityData.Leafs.Append("max", types.YLeaf{"Max", instance.Max})
    instance.EntityData.Leafs.Append("sub-interface", types.YLeaf{"SubInterface", instance.SubInterface})

    instance.EntityData.YListKeys = []string {}

    return &(instance.EntityData)
}

// Snmp_BulkStats_Objects
// Configure an Object List 
type Snmp_BulkStats_Objects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the object List. The type is slice of
    // Snmp_BulkStats_Objects_Object.
    Object []*Snmp_BulkStats_Objects_Object
}

func (objects *Snmp_BulkStats_Objects) GetEntityData() *types.CommonEntityData {
    objects.EntityData.YFilter = objects.YFilter
    objects.EntityData.YangName = "objects"
    objects.EntityData.BundleName = "cisco_ios_xr"
    objects.EntityData.ParentYangName = "bulk-stats"
    objects.EntityData.SegmentPath = "objects"
    objects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objects.EntityData.Children = types.NewOrderedMap()
    objects.EntityData.Children.Append("object", types.YChild{"Object", nil})
    for i := range objects.Object {
        objects.EntityData.Children.Append(types.GetSegmentPath(objects.Object[i]), types.YChild{"Object", objects.Object[i]})
    }
    objects.EntityData.Leafs = types.NewOrderedMap()

    objects.EntityData.YListKeys = []string {}

    return &(objects.EntityData)
}

// Snmp_BulkStats_Objects_Object
// Name of the object List
type Snmp_BulkStats_Objects_Object struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the object List. The type is string with
    // length: 1..32.
    ObjectListName interface{}

    // Configure object list name. The type is interface{}.
    Type interface{}

    // Configure an object List.
    Objects Snmp_BulkStats_Objects_Object_Objects
}

func (object *Snmp_BulkStats_Objects_Object) GetEntityData() *types.CommonEntityData {
    object.EntityData.YFilter = object.YFilter
    object.EntityData.YangName = "object"
    object.EntityData.BundleName = "cisco_ios_xr"
    object.EntityData.ParentYangName = "objects"
    object.EntityData.SegmentPath = "object" + types.AddKeyToken(object.ObjectListName, "object-list-name")
    object.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    object.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    object.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    object.EntityData.Children = types.NewOrderedMap()
    object.EntityData.Children.Append("objects", types.YChild{"Objects", &object.Objects})
    object.EntityData.Leafs = types.NewOrderedMap()
    object.EntityData.Leafs.Append("object-list-name", types.YLeaf{"ObjectListName", object.ObjectListName})
    object.EntityData.Leafs.Append("type", types.YLeaf{"Type", object.Type})

    object.EntityData.YListKeys = []string {"ObjectListName"}

    return &(object.EntityData)
}

// Snmp_BulkStats_Objects_Object_Objects
// Configure an object List
type Snmp_BulkStats_Objects_Object_Objects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object name or OID. The type is slice of
    // Snmp_BulkStats_Objects_Object_Objects_Object.
    Object []*Snmp_BulkStats_Objects_Object_Objects_Object
}

func (objects *Snmp_BulkStats_Objects_Object_Objects) GetEntityData() *types.CommonEntityData {
    objects.EntityData.YFilter = objects.YFilter
    objects.EntityData.YangName = "objects"
    objects.EntityData.BundleName = "cisco_ios_xr"
    objects.EntityData.ParentYangName = "object"
    objects.EntityData.SegmentPath = "objects"
    objects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objects.EntityData.Children = types.NewOrderedMap()
    objects.EntityData.Children.Append("object", types.YChild{"Object", nil})
    for i := range objects.Object {
        objects.EntityData.Children.Append(types.GetSegmentPath(objects.Object[i]), types.YChild{"Object", objects.Object[i]})
    }
    objects.EntityData.Leafs = types.NewOrderedMap()

    objects.EntityData.YListKeys = []string {}

    return &(objects.EntityData)
}

// Snmp_BulkStats_Objects_Object_Objects_Object
// Object name or OID
type Snmp_BulkStats_Objects_Object_Objects_Object struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Object name or OID . The type is string.
    Oid interface{}
}

func (object *Snmp_BulkStats_Objects_Object_Objects_Object) GetEntityData() *types.CommonEntityData {
    object.EntityData.YFilter = object.YFilter
    object.EntityData.YangName = "object"
    object.EntityData.BundleName = "cisco_ios_xr"
    object.EntityData.ParentYangName = "objects"
    object.EntityData.SegmentPath = "object" + types.AddKeyToken(object.Oid, "oid")
    object.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    object.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    object.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    object.EntityData.Children = types.NewOrderedMap()
    object.EntityData.Leafs = types.NewOrderedMap()
    object.EntityData.Leafs.Append("oid", types.YLeaf{"Oid", object.Oid})

    object.EntityData.YListKeys = []string {"Oid"}

    return &(object.EntityData)
}

// Snmp_BulkStats_Transfers
// Periodicity for the transfer of bulk data in
// minutes
type Snmp_BulkStats_Transfers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of bulk transfer. The type is slice of
    // Snmp_BulkStats_Transfers_Transfer.
    Transfer []*Snmp_BulkStats_Transfers_Transfer
}

func (transfers *Snmp_BulkStats_Transfers) GetEntityData() *types.CommonEntityData {
    transfers.EntityData.YFilter = transfers.YFilter
    transfers.EntityData.YangName = "transfers"
    transfers.EntityData.BundleName = "cisco_ios_xr"
    transfers.EntityData.ParentYangName = "bulk-stats"
    transfers.EntityData.SegmentPath = "transfers"
    transfers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transfers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transfers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transfers.EntityData.Children = types.NewOrderedMap()
    transfers.EntityData.Children.Append("transfer", types.YChild{"Transfer", nil})
    for i := range transfers.Transfer {
        transfers.EntityData.Children.Append(types.GetSegmentPath(transfers.Transfer[i]), types.YChild{"Transfer", transfers.Transfer[i]})
    }
    transfers.EntityData.Leafs = types.NewOrderedMap()

    transfers.EntityData.YListKeys = []string {}

    return &(transfers.EntityData)
}

// Snmp_BulkStats_Transfers_Transfer
// Name of bulk transfer
type Snmp_BulkStats_Transfers_Transfer struct {
    EntityData types.CommonEntityData
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
    // interface{} with range: 0..4294967295. Units are minute.
    Interval interface{}

    // Schema that contains objects to be collected.
    TransferSchemas Snmp_BulkStats_Transfers_Transfer_TransferSchemas
}

func (transfer *Snmp_BulkStats_Transfers_Transfer) GetEntityData() *types.CommonEntityData {
    transfer.EntityData.YFilter = transfer.YFilter
    transfer.EntityData.YangName = "transfer"
    transfer.EntityData.BundleName = "cisco_ios_xr"
    transfer.EntityData.ParentYangName = "transfers"
    transfer.EntityData.SegmentPath = "transfer" + types.AddKeyToken(transfer.TransferName, "transfer-name")
    transfer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transfer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transfer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transfer.EntityData.Children = types.NewOrderedMap()
    transfer.EntityData.Children.Append("transfer-schemas", types.YChild{"TransferSchemas", &transfer.TransferSchemas})
    transfer.EntityData.Leafs = types.NewOrderedMap()
    transfer.EntityData.Leafs.Append("transfer-name", types.YLeaf{"TransferName", transfer.TransferName})
    transfer.EntityData.Leafs.Append("secondary", types.YLeaf{"Secondary", transfer.Secondary})
    transfer.EntityData.Leafs.Append("type", types.YLeaf{"Type", transfer.Type})
    transfer.EntityData.Leafs.Append("buffer-size", types.YLeaf{"BufferSize", transfer.BufferSize})
    transfer.EntityData.Leafs.Append("retain", types.YLeaf{"Retain", transfer.Retain})
    transfer.EntityData.Leafs.Append("format", types.YLeaf{"Format", transfer.Format})
    transfer.EntityData.Leafs.Append("retry", types.YLeaf{"Retry", transfer.Retry})
    transfer.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", transfer.Enable})
    transfer.EntityData.Leafs.Append("primary", types.YLeaf{"Primary", transfer.Primary})
    transfer.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", transfer.Interval})

    transfer.EntityData.YListKeys = []string {"TransferName"}

    return &(transfer.EntityData)
}

// Snmp_BulkStats_Transfers_Transfer_TransferSchemas
// Schema that contains objects to be collected
type Snmp_BulkStats_Transfers_Transfer_TransferSchemas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Schema that contains objects to be collected. The type is slice of
    // Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema.
    TransferSchema []*Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema
}

func (transferSchemas *Snmp_BulkStats_Transfers_Transfer_TransferSchemas) GetEntityData() *types.CommonEntityData {
    transferSchemas.EntityData.YFilter = transferSchemas.YFilter
    transferSchemas.EntityData.YangName = "transfer-schemas"
    transferSchemas.EntityData.BundleName = "cisco_ios_xr"
    transferSchemas.EntityData.ParentYangName = "transfer"
    transferSchemas.EntityData.SegmentPath = "transfer-schemas"
    transferSchemas.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transferSchemas.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transferSchemas.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transferSchemas.EntityData.Children = types.NewOrderedMap()
    transferSchemas.EntityData.Children.Append("transfer-schema", types.YChild{"TransferSchema", nil})
    for i := range transferSchemas.TransferSchema {
        transferSchemas.EntityData.Children.Append(types.GetSegmentPath(transferSchemas.TransferSchema[i]), types.YChild{"TransferSchema", transferSchemas.TransferSchema[i]})
    }
    transferSchemas.EntityData.Leafs = types.NewOrderedMap()

    transferSchemas.EntityData.YListKeys = []string {}

    return &(transferSchemas.EntityData)
}

// Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema
// Schema that contains objects to be collected
type Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Schema that contains objects to be collected. The
    // type is string with length: 1..32.
    SchemaName interface{}
}

func (transferSchema *Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema) GetEntityData() *types.CommonEntityData {
    transferSchema.EntityData.YFilter = transferSchema.YFilter
    transferSchema.EntityData.YangName = "transfer-schema"
    transferSchema.EntityData.BundleName = "cisco_ios_xr"
    transferSchema.EntityData.ParentYangName = "transfer-schemas"
    transferSchema.EntityData.SegmentPath = "transfer-schema" + types.AddKeyToken(transferSchema.SchemaName, "schema-name")
    transferSchema.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transferSchema.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transferSchema.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transferSchema.EntityData.Children = types.NewOrderedMap()
    transferSchema.EntityData.Leafs = types.NewOrderedMap()
    transferSchema.EntityData.Leafs.Append("schema-name", types.YLeaf{"SchemaName", transferSchema.SchemaName})

    transferSchema.EntityData.YListKeys = []string {"SchemaName"}

    return &(transferSchema.EntityData)
}

// Snmp_DefaultCommunityMaps
// Container class to hold unencrpted community map
type Snmp_DefaultCommunityMaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unencrpted SNMP community map name . The type is slice of
    // Snmp_DefaultCommunityMaps_DefaultCommunityMap.
    DefaultCommunityMap []*Snmp_DefaultCommunityMaps_DefaultCommunityMap
}

func (defaultCommunityMaps *Snmp_DefaultCommunityMaps) GetEntityData() *types.CommonEntityData {
    defaultCommunityMaps.EntityData.YFilter = defaultCommunityMaps.YFilter
    defaultCommunityMaps.EntityData.YangName = "default-community-maps"
    defaultCommunityMaps.EntityData.BundleName = "cisco_ios_xr"
    defaultCommunityMaps.EntityData.ParentYangName = "snmp"
    defaultCommunityMaps.EntityData.SegmentPath = "default-community-maps"
    defaultCommunityMaps.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultCommunityMaps.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultCommunityMaps.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultCommunityMaps.EntityData.Children = types.NewOrderedMap()
    defaultCommunityMaps.EntityData.Children.Append("default-community-map", types.YChild{"DefaultCommunityMap", nil})
    for i := range defaultCommunityMaps.DefaultCommunityMap {
        defaultCommunityMaps.EntityData.Children.Append(types.GetSegmentPath(defaultCommunityMaps.DefaultCommunityMap[i]), types.YChild{"DefaultCommunityMap", defaultCommunityMaps.DefaultCommunityMap[i]})
    }
    defaultCommunityMaps.EntityData.Leafs = types.NewOrderedMap()

    defaultCommunityMaps.EntityData.YListKeys = []string {}

    return &(defaultCommunityMaps.EntityData)
}

// Snmp_DefaultCommunityMaps_DefaultCommunityMap
// Unencrpted SNMP community map name 
type Snmp_DefaultCommunityMaps_DefaultCommunityMap struct {
    EntityData types.CommonEntityData
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

func (defaultCommunityMap *Snmp_DefaultCommunityMaps_DefaultCommunityMap) GetEntityData() *types.CommonEntityData {
    defaultCommunityMap.EntityData.YFilter = defaultCommunityMap.YFilter
    defaultCommunityMap.EntityData.YangName = "default-community-map"
    defaultCommunityMap.EntityData.BundleName = "cisco_ios_xr"
    defaultCommunityMap.EntityData.ParentYangName = "default-community-maps"
    defaultCommunityMap.EntityData.SegmentPath = "default-community-map" + types.AddKeyToken(defaultCommunityMap.CommunityName, "community-name")
    defaultCommunityMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultCommunityMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultCommunityMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultCommunityMap.EntityData.Children = types.NewOrderedMap()
    defaultCommunityMap.EntityData.Leafs = types.NewOrderedMap()
    defaultCommunityMap.EntityData.Leafs.Append("community-name", types.YLeaf{"CommunityName", defaultCommunityMap.CommunityName})
    defaultCommunityMap.EntityData.Leafs.Append("context", types.YLeaf{"Context", defaultCommunityMap.Context})
    defaultCommunityMap.EntityData.Leafs.Append("security", types.YLeaf{"Security", defaultCommunityMap.Security})
    defaultCommunityMap.EntityData.Leafs.Append("target-list", types.YLeaf{"TargetList", defaultCommunityMap.TargetList})

    defaultCommunityMap.EntityData.YListKeys = []string {"CommunityName"}

    return &(defaultCommunityMap.EntityData)
}

// Snmp_OverloadControl
// Set overload control params for handling
// incoming messages
// This type is a presence type.
type Snmp_OverloadControl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Drop time in seconds for incoming queue (default 1 sec). The type is
    // interface{} with range: 0..300. This attribute is mandatory. Units are
    // second.
    DropTime interface{}

    // Throttle time in milliseconds for incoming queue (default 500 msec). The
    // type is interface{} with range: 0..1000. This attribute is mandatory. Units
    // are millisecond.
    ThrottleRate interface{}
}

func (overloadControl *Snmp_OverloadControl) GetEntityData() *types.CommonEntityData {
    overloadControl.EntityData.YFilter = overloadControl.YFilter
    overloadControl.EntityData.YangName = "overload-control"
    overloadControl.EntityData.BundleName = "cisco_ios_xr"
    overloadControl.EntityData.ParentYangName = "snmp"
    overloadControl.EntityData.SegmentPath = "overload-control"
    overloadControl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    overloadControl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    overloadControl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    overloadControl.EntityData.Children = types.NewOrderedMap()
    overloadControl.EntityData.Leafs = types.NewOrderedMap()
    overloadControl.EntityData.Leafs.Append("drop-time", types.YLeaf{"DropTime", overloadControl.DropTime})
    overloadControl.EntityData.Leafs.Append("throttle-rate", types.YLeaf{"ThrottleRate", overloadControl.ThrottleRate})

    overloadControl.EntityData.YListKeys = []string {}

    return &(overloadControl.EntityData)
}

// Snmp_Timeouts
// SNMP timeouts
type Snmp_Timeouts struct {
    EntityData types.CommonEntityData
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

func (timeouts *Snmp_Timeouts) GetEntityData() *types.CommonEntityData {
    timeouts.EntityData.YFilter = timeouts.YFilter
    timeouts.EntityData.YangName = "timeouts"
    timeouts.EntityData.BundleName = "cisco_ios_xr"
    timeouts.EntityData.ParentYangName = "snmp"
    timeouts.EntityData.SegmentPath = "timeouts"
    timeouts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeouts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeouts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeouts.EntityData.Children = types.NewOrderedMap()
    timeouts.EntityData.Leafs = types.NewOrderedMap()
    timeouts.EntityData.Leafs.Append("duplicates", types.YLeaf{"Duplicates", timeouts.Duplicates})
    timeouts.EntityData.Leafs.Append("in-qdrop", types.YLeaf{"InQdrop", timeouts.InQdrop})
    timeouts.EntityData.Leafs.Append("subagent", types.YLeaf{"Subagent", timeouts.Subagent})
    timeouts.EntityData.Leafs.Append("pdu-stats", types.YLeaf{"PduStats", timeouts.PduStats})

    timeouts.EntityData.YListKeys = []string {}

    return &(timeouts.EntityData)
}

// Snmp_Users
// Define a user who can access the SNMP engine
type Snmp_Users struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the user. The type is slice of Snmp_Users_User.
    User []*Snmp_Users_User
}

func (users *Snmp_Users) GetEntityData() *types.CommonEntityData {
    users.EntityData.YFilter = users.YFilter
    users.EntityData.YangName = "users"
    users.EntityData.BundleName = "cisco_ios_xr"
    users.EntityData.ParentYangName = "snmp"
    users.EntityData.SegmentPath = "users"
    users.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    users.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    users.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    users.EntityData.Children = types.NewOrderedMap()
    users.EntityData.Children.Append("user", types.YChild{"User", nil})
    for i := range users.User {
        users.EntityData.Children.Append(types.GetSegmentPath(users.User[i]), types.YChild{"User", users.User[i]})
    }
    users.EntityData.Leafs = types.NewOrderedMap()

    users.EntityData.YListKeys = []string {}

    return &(users.EntityData)
}

// Snmp_Users_User
// Name of the user
type Snmp_Users_User struct {
    EntityData types.CommonEntityData
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
    V4aclType interface{}

    // Ipv4 Access-list name. The type is string.
    V4AccessList interface{}

    // Access-list type. The type is Snmpacl.
    V6aclType interface{}

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

func (user *Snmp_Users_User) GetEntityData() *types.CommonEntityData {
    user.EntityData.YFilter = user.YFilter
    user.EntityData.YangName = "user"
    user.EntityData.BundleName = "cisco_ios_xr"
    user.EntityData.ParentYangName = "users"
    user.EntityData.SegmentPath = "user" + types.AddKeyToken(user.UserName, "user-name")
    user.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    user.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    user.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    user.EntityData.Children = types.NewOrderedMap()
    user.EntityData.Leafs = types.NewOrderedMap()
    user.EntityData.Leafs.Append("user-name", types.YLeaf{"UserName", user.UserName})
    user.EntityData.Leafs.Append("group-name", types.YLeaf{"GroupName", user.GroupName})
    user.EntityData.Leafs.Append("version", types.YLeaf{"Version", user.Version})
    user.EntityData.Leafs.Append("authentication-password-configured", types.YLeaf{"AuthenticationPasswordConfigured", user.AuthenticationPasswordConfigured})
    user.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", user.Algorithm})
    user.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", user.AuthenticationPassword})
    user.EntityData.Leafs.Append("privacy-password-configured", types.YLeaf{"PrivacyPasswordConfigured", user.PrivacyPasswordConfigured})
    user.EntityData.Leafs.Append("priv-algorithm", types.YLeaf{"PrivAlgorithm", user.PrivAlgorithm})
    user.EntityData.Leafs.Append("privacy-password", types.YLeaf{"PrivacyPassword", user.PrivacyPassword})
    user.EntityData.Leafs.Append("v4acl-type", types.YLeaf{"V4aclType", user.V4aclType})
    user.EntityData.Leafs.Append("v4-access-list", types.YLeaf{"V4AccessList", user.V4AccessList})
    user.EntityData.Leafs.Append("v6acl-type", types.YLeaf{"V6aclType", user.V6aclType})
    user.EntityData.Leafs.Append("v6-access-list", types.YLeaf{"V6AccessList", user.V6AccessList})
    user.EntityData.Leafs.Append("owner", types.YLeaf{"Owner", user.Owner})
    user.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", user.RemoteAddress})
    user.EntityData.Leafs.Append("port", types.YLeaf{"Port", user.Port})

    user.EntityData.YListKeys = []string {"UserName"}

    return &(user.EntityData)
}

// Snmp_Vrfs
// SNMP VRF configuration commands
type Snmp_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is slice of Snmp_Vrfs_Vrf.
    Vrf []*Snmp_Vrfs_Vrf
}

func (vrfs *Snmp_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "snmp"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// Snmp_Vrfs_Vrf
// VRF name
type Snmp_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
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

func (vrf *Snmp_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.Name, "name")
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("trap-hosts", types.YChild{"TrapHosts", &vrf.TrapHosts})
    vrf.EntityData.Children.Append("contexts", types.YChild{"Contexts", &vrf.Contexts})
    vrf.EntityData.Children.Append("context-mappings", types.YChild{"ContextMappings", &vrf.ContextMappings})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("name", types.YLeaf{"Name", vrf.Name})

    vrf.EntityData.YListKeys = []string {"Name"}

    return &(vrf.EntityData)
}

// Snmp_Vrfs_Vrf_TrapHosts
// Specify hosts to receive SNMP notifications
type Snmp_Vrfs_Vrf_TrapHosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify hosts to receive SNMP notifications. The type is slice of
    // Snmp_Vrfs_Vrf_TrapHosts_TrapHost.
    TrapHost []*Snmp_Vrfs_Vrf_TrapHosts_TrapHost
}

func (trapHosts *Snmp_Vrfs_Vrf_TrapHosts) GetEntityData() *types.CommonEntityData {
    trapHosts.EntityData.YFilter = trapHosts.YFilter
    trapHosts.EntityData.YangName = "trap-hosts"
    trapHosts.EntityData.BundleName = "cisco_ios_xr"
    trapHosts.EntityData.ParentYangName = "vrf"
    trapHosts.EntityData.SegmentPath = "trap-hosts"
    trapHosts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapHosts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapHosts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapHosts.EntityData.Children = types.NewOrderedMap()
    trapHosts.EntityData.Children.Append("trap-host", types.YChild{"TrapHost", nil})
    for i := range trapHosts.TrapHost {
        trapHosts.EntityData.Children.Append(types.GetSegmentPath(trapHosts.TrapHost[i]), types.YChild{"TrapHost", trapHosts.TrapHost[i]})
    }
    trapHosts.EntityData.Leafs = types.NewOrderedMap()

    trapHosts.EntityData.YListKeys = []string {}

    return &(trapHosts.EntityData)
}

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost
// Specify hosts to receive SNMP notifications
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost struct {
    EntityData types.CommonEntityData
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

func (trapHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost) GetEntityData() *types.CommonEntityData {
    trapHost.EntityData.YFilter = trapHost.YFilter
    trapHost.EntityData.YangName = "trap-host"
    trapHost.EntityData.BundleName = "cisco_ios_xr"
    trapHost.EntityData.ParentYangName = "trap-hosts"
    trapHost.EntityData.SegmentPath = "trap-host" + types.AddKeyToken(trapHost.IpAddress, "ip-address")
    trapHost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapHost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapHost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapHost.EntityData.Children = types.NewOrderedMap()
    trapHost.EntityData.Children.Append("encrypted-user-communities", types.YChild{"EncryptedUserCommunities", &trapHost.EncryptedUserCommunities})
    trapHost.EntityData.Children.Append("inform-host", types.YChild{"InformHost", &trapHost.InformHost})
    trapHost.EntityData.Children.Append("default-user-communities", types.YChild{"DefaultUserCommunities", &trapHost.DefaultUserCommunities})
    trapHost.EntityData.Leafs = types.NewOrderedMap()
    trapHost.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", trapHost.IpAddress})

    trapHost.EntityData.YListKeys = []string {"IpAddress"}

    return &(trapHost.EntityData)
}

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities
// Container class for defining Clear/encrypt
// communities for a trap host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear/Encrypt Community name associated with a trap host. The type is slice
    // of
    // Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity.
    EncryptedUserCommunity []*Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity
}

func (encryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities) GetEntityData() *types.CommonEntityData {
    encryptedUserCommunities.EntityData.YFilter = encryptedUserCommunities.YFilter
    encryptedUserCommunities.EntityData.YangName = "encrypted-user-communities"
    encryptedUserCommunities.EntityData.BundleName = "cisco_ios_xr"
    encryptedUserCommunities.EntityData.ParentYangName = "trap-host"
    encryptedUserCommunities.EntityData.SegmentPath = "encrypted-user-communities"
    encryptedUserCommunities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryptedUserCommunities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryptedUserCommunities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryptedUserCommunities.EntityData.Children = types.NewOrderedMap()
    encryptedUserCommunities.EntityData.Children.Append("encrypted-user-community", types.YChild{"EncryptedUserCommunity", nil})
    for i := range encryptedUserCommunities.EncryptedUserCommunity {
        encryptedUserCommunities.EntityData.Children.Append(types.GetSegmentPath(encryptedUserCommunities.EncryptedUserCommunity[i]), types.YChild{"EncryptedUserCommunity", encryptedUserCommunities.EncryptedUserCommunity[i]})
    }
    encryptedUserCommunities.EntityData.Leafs = types.NewOrderedMap()

    encryptedUserCommunities.EntityData.YListKeys = []string {}

    return &(encryptedUserCommunities.EntityData)
}

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity
// Clear/Encrypt Community name associated with
// a trap host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity struct {
    EntityData types.CommonEntityData
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
    // range: 0..4294967295. This attribute is mandatory.
    BasicTrapTypes interface{}

    // Number to signify the feature traps that needs to be setUse this for
    // providing copy-complete trapValue must be set to 0 if not used. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    AdvancedTrapTypes1 interface{}

    // Number to signify the feature traps that needs to be setvalue should always
    // to set as 0. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    AdvancedTrapTypes2 interface{}
}

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetEntityData() *types.CommonEntityData {
    encryptedUserCommunity.EntityData.YFilter = encryptedUserCommunity.YFilter
    encryptedUserCommunity.EntityData.YangName = "encrypted-user-community"
    encryptedUserCommunity.EntityData.BundleName = "cisco_ios_xr"
    encryptedUserCommunity.EntityData.ParentYangName = "encrypted-user-communities"
    encryptedUserCommunity.EntityData.SegmentPath = "encrypted-user-community" + types.AddKeyToken(encryptedUserCommunity.CommunityName, "community-name")
    encryptedUserCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryptedUserCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryptedUserCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryptedUserCommunity.EntityData.Children = types.NewOrderedMap()
    encryptedUserCommunity.EntityData.Leafs = types.NewOrderedMap()
    encryptedUserCommunity.EntityData.Leafs.Append("community-name", types.YLeaf{"CommunityName", encryptedUserCommunity.CommunityName})
    encryptedUserCommunity.EntityData.Leafs.Append("port", types.YLeaf{"Port", encryptedUserCommunity.Port})
    encryptedUserCommunity.EntityData.Leafs.Append("version", types.YLeaf{"Version", encryptedUserCommunity.Version})
    encryptedUserCommunity.EntityData.Leafs.Append("security-level", types.YLeaf{"SecurityLevel", encryptedUserCommunity.SecurityLevel})
    encryptedUserCommunity.EntityData.Leafs.Append("basic-trap-types", types.YLeaf{"BasicTrapTypes", encryptedUserCommunity.BasicTrapTypes})
    encryptedUserCommunity.EntityData.Leafs.Append("advanced-trap-types1", types.YLeaf{"AdvancedTrapTypes1", encryptedUserCommunity.AdvancedTrapTypes1})
    encryptedUserCommunity.EntityData.Leafs.Append("advanced-trap-types2", types.YLeaf{"AdvancedTrapTypes2", encryptedUserCommunity.AdvancedTrapTypes2})

    encryptedUserCommunity.EntityData.YListKeys = []string {"CommunityName"}

    return &(encryptedUserCommunity.EntityData)
}

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost
// Container class for defining notification type
// for a Inform host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Container class for defining communities for a inform host.
    InformUserCommunities Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities

    // Container class for defining Clear/encrypt communities for a inform host.
    InformEncryptedUserCommunities Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities
}

func (informHost *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost) GetEntityData() *types.CommonEntityData {
    informHost.EntityData.YFilter = informHost.YFilter
    informHost.EntityData.YangName = "inform-host"
    informHost.EntityData.BundleName = "cisco_ios_xr"
    informHost.EntityData.ParentYangName = "trap-host"
    informHost.EntityData.SegmentPath = "inform-host"
    informHost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    informHost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    informHost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    informHost.EntityData.Children = types.NewOrderedMap()
    informHost.EntityData.Children.Append("inform-user-communities", types.YChild{"InformUserCommunities", &informHost.InformUserCommunities})
    informHost.EntityData.Children.Append("inform-encrypted-user-communities", types.YChild{"InformEncryptedUserCommunities", &informHost.InformEncryptedUserCommunities})
    informHost.EntityData.Leafs = types.NewOrderedMap()

    informHost.EntityData.YListKeys = []string {}

    return &(informHost.EntityData)
}

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities
// Container class for defining communities for
// a inform host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unencrpted Community name associated with a inform host. The type is slice
    // of
    // Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity.
    InformUserCommunity []*Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity
}

func (informUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetEntityData() *types.CommonEntityData {
    informUserCommunities.EntityData.YFilter = informUserCommunities.YFilter
    informUserCommunities.EntityData.YangName = "inform-user-communities"
    informUserCommunities.EntityData.BundleName = "cisco_ios_xr"
    informUserCommunities.EntityData.ParentYangName = "inform-host"
    informUserCommunities.EntityData.SegmentPath = "inform-user-communities"
    informUserCommunities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    informUserCommunities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    informUserCommunities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    informUserCommunities.EntityData.Children = types.NewOrderedMap()
    informUserCommunities.EntityData.Children.Append("inform-user-community", types.YChild{"InformUserCommunity", nil})
    for i := range informUserCommunities.InformUserCommunity {
        informUserCommunities.EntityData.Children.Append(types.GetSegmentPath(informUserCommunities.InformUserCommunity[i]), types.YChild{"InformUserCommunity", informUserCommunities.InformUserCommunity[i]})
    }
    informUserCommunities.EntityData.Leafs = types.NewOrderedMap()

    informUserCommunities.EntityData.YListKeys = []string {}

    return &(informUserCommunities.EntityData)
}

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity
// Unencrpted Community name associated with a
// inform host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity struct {
    EntityData types.CommonEntityData
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
    // range: 0..4294967295. This attribute is mandatory.
    BasicTrapTypes interface{}

    // Number to signify the feature traps that needs to be setUse this for
    // providing copy-complete trapValue must be set to 0 if not used. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    AdvancedTrapTypes1 interface{}

    // Number to signify the feature traps that needs to be setvalue should always
    // to set as 0. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    AdvancedTrapTypes2 interface{}
}

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetEntityData() *types.CommonEntityData {
    informUserCommunity.EntityData.YFilter = informUserCommunity.YFilter
    informUserCommunity.EntityData.YangName = "inform-user-community"
    informUserCommunity.EntityData.BundleName = "cisco_ios_xr"
    informUserCommunity.EntityData.ParentYangName = "inform-user-communities"
    informUserCommunity.EntityData.SegmentPath = "inform-user-community" + types.AddKeyToken(informUserCommunity.CommunityName, "community-name")
    informUserCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    informUserCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    informUserCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    informUserCommunity.EntityData.Children = types.NewOrderedMap()
    informUserCommunity.EntityData.Leafs = types.NewOrderedMap()
    informUserCommunity.EntityData.Leafs.Append("community-name", types.YLeaf{"CommunityName", informUserCommunity.CommunityName})
    informUserCommunity.EntityData.Leafs.Append("port", types.YLeaf{"Port", informUserCommunity.Port})
    informUserCommunity.EntityData.Leafs.Append("version", types.YLeaf{"Version", informUserCommunity.Version})
    informUserCommunity.EntityData.Leafs.Append("security-level", types.YLeaf{"SecurityLevel", informUserCommunity.SecurityLevel})
    informUserCommunity.EntityData.Leafs.Append("basic-trap-types", types.YLeaf{"BasicTrapTypes", informUserCommunity.BasicTrapTypes})
    informUserCommunity.EntityData.Leafs.Append("advanced-trap-types1", types.YLeaf{"AdvancedTrapTypes1", informUserCommunity.AdvancedTrapTypes1})
    informUserCommunity.EntityData.Leafs.Append("advanced-trap-types2", types.YLeaf{"AdvancedTrapTypes2", informUserCommunity.AdvancedTrapTypes2})

    informUserCommunity.EntityData.YListKeys = []string {"CommunityName"}

    return &(informUserCommunity.EntityData)
}

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities
// Container class for defining Clear/encrypt
// communities for a inform host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear/Encrypt Community name associated with a inform host. The type is
    // slice of
    // Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity.
    InformEncryptedUserCommunity []*Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity
}

func (informEncryptedUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetEntityData() *types.CommonEntityData {
    informEncryptedUserCommunities.EntityData.YFilter = informEncryptedUserCommunities.YFilter
    informEncryptedUserCommunities.EntityData.YangName = "inform-encrypted-user-communities"
    informEncryptedUserCommunities.EntityData.BundleName = "cisco_ios_xr"
    informEncryptedUserCommunities.EntityData.ParentYangName = "inform-host"
    informEncryptedUserCommunities.EntityData.SegmentPath = "inform-encrypted-user-communities"
    informEncryptedUserCommunities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    informEncryptedUserCommunities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    informEncryptedUserCommunities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    informEncryptedUserCommunities.EntityData.Children = types.NewOrderedMap()
    informEncryptedUserCommunities.EntityData.Children.Append("inform-encrypted-user-community", types.YChild{"InformEncryptedUserCommunity", nil})
    for i := range informEncryptedUserCommunities.InformEncryptedUserCommunity {
        informEncryptedUserCommunities.EntityData.Children.Append(types.GetSegmentPath(informEncryptedUserCommunities.InformEncryptedUserCommunity[i]), types.YChild{"InformEncryptedUserCommunity", informEncryptedUserCommunities.InformEncryptedUserCommunity[i]})
    }
    informEncryptedUserCommunities.EntityData.Leafs = types.NewOrderedMap()

    informEncryptedUserCommunities.EntityData.YListKeys = []string {}

    return &(informEncryptedUserCommunities.EntityData)
}

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity
// Clear/Encrypt Community name associated with
// a inform host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity struct {
    EntityData types.CommonEntityData
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
    // range: 0..4294967295. This attribute is mandatory.
    BasicTrapTypes interface{}

    // Number to signify the feature traps that needs to be setUse this for
    // providing copy-complete trapValue must be set to 0 if not used. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    AdvancedTrapTypes1 interface{}

    // Number to signify the feature traps that needs to be setvalue should always
    // to set as 0. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    AdvancedTrapTypes2 interface{}
}

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetEntityData() *types.CommonEntityData {
    informEncryptedUserCommunity.EntityData.YFilter = informEncryptedUserCommunity.YFilter
    informEncryptedUserCommunity.EntityData.YangName = "inform-encrypted-user-community"
    informEncryptedUserCommunity.EntityData.BundleName = "cisco_ios_xr"
    informEncryptedUserCommunity.EntityData.ParentYangName = "inform-encrypted-user-communities"
    informEncryptedUserCommunity.EntityData.SegmentPath = "inform-encrypted-user-community" + types.AddKeyToken(informEncryptedUserCommunity.CommunityName, "community-name")
    informEncryptedUserCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    informEncryptedUserCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    informEncryptedUserCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    informEncryptedUserCommunity.EntityData.Children = types.NewOrderedMap()
    informEncryptedUserCommunity.EntityData.Leafs = types.NewOrderedMap()
    informEncryptedUserCommunity.EntityData.Leafs.Append("community-name", types.YLeaf{"CommunityName", informEncryptedUserCommunity.CommunityName})
    informEncryptedUserCommunity.EntityData.Leafs.Append("port", types.YLeaf{"Port", informEncryptedUserCommunity.Port})
    informEncryptedUserCommunity.EntityData.Leafs.Append("version", types.YLeaf{"Version", informEncryptedUserCommunity.Version})
    informEncryptedUserCommunity.EntityData.Leafs.Append("security-level", types.YLeaf{"SecurityLevel", informEncryptedUserCommunity.SecurityLevel})
    informEncryptedUserCommunity.EntityData.Leafs.Append("basic-trap-types", types.YLeaf{"BasicTrapTypes", informEncryptedUserCommunity.BasicTrapTypes})
    informEncryptedUserCommunity.EntityData.Leafs.Append("advanced-trap-types1", types.YLeaf{"AdvancedTrapTypes1", informEncryptedUserCommunity.AdvancedTrapTypes1})
    informEncryptedUserCommunity.EntityData.Leafs.Append("advanced-trap-types2", types.YLeaf{"AdvancedTrapTypes2", informEncryptedUserCommunity.AdvancedTrapTypes2})

    informEncryptedUserCommunity.EntityData.YListKeys = []string {"CommunityName"}

    return &(informEncryptedUserCommunity.EntityData)
}

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities
// Container class for defining communities for a
// trap host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unencrpted Community name associated with a trap host. The type is slice of
    // Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity.
    DefaultUserCommunity []*Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity
}

func (defaultUserCommunities *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities) GetEntityData() *types.CommonEntityData {
    defaultUserCommunities.EntityData.YFilter = defaultUserCommunities.YFilter
    defaultUserCommunities.EntityData.YangName = "default-user-communities"
    defaultUserCommunities.EntityData.BundleName = "cisco_ios_xr"
    defaultUserCommunities.EntityData.ParentYangName = "trap-host"
    defaultUserCommunities.EntityData.SegmentPath = "default-user-communities"
    defaultUserCommunities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultUserCommunities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultUserCommunities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultUserCommunities.EntityData.Children = types.NewOrderedMap()
    defaultUserCommunities.EntityData.Children.Append("default-user-community", types.YChild{"DefaultUserCommunity", nil})
    for i := range defaultUserCommunities.DefaultUserCommunity {
        defaultUserCommunities.EntityData.Children.Append(types.GetSegmentPath(defaultUserCommunities.DefaultUserCommunity[i]), types.YChild{"DefaultUserCommunity", defaultUserCommunities.DefaultUserCommunity[i]})
    }
    defaultUserCommunities.EntityData.Leafs = types.NewOrderedMap()

    defaultUserCommunities.EntityData.YListKeys = []string {}

    return &(defaultUserCommunities.EntityData)
}

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity
// Unencrpted Community name associated with a
// trap host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity struct {
    EntityData types.CommonEntityData
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
    // range: 0..4294967295. This attribute is mandatory.
    BasicTrapTypes interface{}

    // Number to signify the feature traps that needs to be setUse this for
    // providing copy-complete trapValue must be set to 0 if not used. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    AdvancedTrapTypes1 interface{}

    // Number to signify the feature traps that needs to be setvalue should always
    // to set as 0. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    AdvancedTrapTypes2 interface{}
}

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetEntityData() *types.CommonEntityData {
    defaultUserCommunity.EntityData.YFilter = defaultUserCommunity.YFilter
    defaultUserCommunity.EntityData.YangName = "default-user-community"
    defaultUserCommunity.EntityData.BundleName = "cisco_ios_xr"
    defaultUserCommunity.EntityData.ParentYangName = "default-user-communities"
    defaultUserCommunity.EntityData.SegmentPath = "default-user-community" + types.AddKeyToken(defaultUserCommunity.CommunityName, "community-name")
    defaultUserCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultUserCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultUserCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultUserCommunity.EntityData.Children = types.NewOrderedMap()
    defaultUserCommunity.EntityData.Leafs = types.NewOrderedMap()
    defaultUserCommunity.EntityData.Leafs.Append("community-name", types.YLeaf{"CommunityName", defaultUserCommunity.CommunityName})
    defaultUserCommunity.EntityData.Leafs.Append("port", types.YLeaf{"Port", defaultUserCommunity.Port})
    defaultUserCommunity.EntityData.Leafs.Append("version", types.YLeaf{"Version", defaultUserCommunity.Version})
    defaultUserCommunity.EntityData.Leafs.Append("security-level", types.YLeaf{"SecurityLevel", defaultUserCommunity.SecurityLevel})
    defaultUserCommunity.EntityData.Leafs.Append("basic-trap-types", types.YLeaf{"BasicTrapTypes", defaultUserCommunity.BasicTrapTypes})
    defaultUserCommunity.EntityData.Leafs.Append("advanced-trap-types1", types.YLeaf{"AdvancedTrapTypes1", defaultUserCommunity.AdvancedTrapTypes1})
    defaultUserCommunity.EntityData.Leafs.Append("advanced-trap-types2", types.YLeaf{"AdvancedTrapTypes2", defaultUserCommunity.AdvancedTrapTypes2})

    defaultUserCommunity.EntityData.YListKeys = []string {"CommunityName"}

    return &(defaultUserCommunity.EntityData)
}

// Snmp_Vrfs_Vrf_Contexts
// List of Context Names
type Snmp_Vrfs_Vrf_Contexts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Context Name. The type is slice of Snmp_Vrfs_Vrf_Contexts_Context.
    Context []*Snmp_Vrfs_Vrf_Contexts_Context
}

func (contexts *Snmp_Vrfs_Vrf_Contexts) GetEntityData() *types.CommonEntityData {
    contexts.EntityData.YFilter = contexts.YFilter
    contexts.EntityData.YangName = "contexts"
    contexts.EntityData.BundleName = "cisco_ios_xr"
    contexts.EntityData.ParentYangName = "vrf"
    contexts.EntityData.SegmentPath = "contexts"
    contexts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contexts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contexts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contexts.EntityData.Children = types.NewOrderedMap()
    contexts.EntityData.Children.Append("context", types.YChild{"Context", nil})
    for i := range contexts.Context {
        contexts.EntityData.Children.Append(types.GetSegmentPath(contexts.Context[i]), types.YChild{"Context", contexts.Context[i]})
    }
    contexts.EntityData.Leafs = types.NewOrderedMap()

    contexts.EntityData.YListKeys = []string {}

    return &(contexts.EntityData)
}

// Snmp_Vrfs_Vrf_Contexts_Context
// Context Name
type Snmp_Vrfs_Vrf_Contexts_Context struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Context Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ContextName interface{}
}

func (context *Snmp_Vrfs_Vrf_Contexts_Context) GetEntityData() *types.CommonEntityData {
    context.EntityData.YFilter = context.YFilter
    context.EntityData.YangName = "context"
    context.EntityData.BundleName = "cisco_ios_xr"
    context.EntityData.ParentYangName = "contexts"
    context.EntityData.SegmentPath = "context" + types.AddKeyToken(context.ContextName, "context-name")
    context.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    context.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    context.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    context.EntityData.Children = types.NewOrderedMap()
    context.EntityData.Leafs = types.NewOrderedMap()
    context.EntityData.Leafs.Append("context-name", types.YLeaf{"ContextName", context.ContextName})

    context.EntityData.YListKeys = []string {"ContextName"}

    return &(context.EntityData)
}

// Snmp_Vrfs_Vrf_ContextMappings
// List of context names
type Snmp_Vrfs_Vrf_ContextMappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Context mapping name. The type is slice of
    // Snmp_Vrfs_Vrf_ContextMappings_ContextMapping.
    ContextMapping []*Snmp_Vrfs_Vrf_ContextMappings_ContextMapping
}

func (contextMappings *Snmp_Vrfs_Vrf_ContextMappings) GetEntityData() *types.CommonEntityData {
    contextMappings.EntityData.YFilter = contextMappings.YFilter
    contextMappings.EntityData.YangName = "context-mappings"
    contextMappings.EntityData.BundleName = "cisco_ios_xr"
    contextMappings.EntityData.ParentYangName = "vrf"
    contextMappings.EntityData.SegmentPath = "context-mappings"
    contextMappings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextMappings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextMappings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextMappings.EntityData.Children = types.NewOrderedMap()
    contextMappings.EntityData.Children.Append("context-mapping", types.YChild{"ContextMapping", nil})
    for i := range contextMappings.ContextMapping {
        contextMappings.EntityData.Children.Append(types.GetSegmentPath(contextMappings.ContextMapping[i]), types.YChild{"ContextMapping", contextMappings.ContextMapping[i]})
    }
    contextMappings.EntityData.Leafs = types.NewOrderedMap()

    contextMappings.EntityData.YListKeys = []string {}

    return &(contextMappings.EntityData)
}

// Snmp_Vrfs_Vrf_ContextMappings_ContextMapping
// Context mapping name
type Snmp_Vrfs_Vrf_ContextMappings_ContextMapping struct {
    EntityData types.CommonEntityData
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

func (contextMapping *Snmp_Vrfs_Vrf_ContextMappings_ContextMapping) GetEntityData() *types.CommonEntityData {
    contextMapping.EntityData.YFilter = contextMapping.YFilter
    contextMapping.EntityData.YangName = "context-mapping"
    contextMapping.EntityData.BundleName = "cisco_ios_xr"
    contextMapping.EntityData.ParentYangName = "context-mappings"
    contextMapping.EntityData.SegmentPath = "context-mapping" + types.AddKeyToken(contextMapping.ContextMappingName, "context-mapping-name")
    contextMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextMapping.EntityData.Children = types.NewOrderedMap()
    contextMapping.EntityData.Leafs = types.NewOrderedMap()
    contextMapping.EntityData.Leafs.Append("context-mapping-name", types.YLeaf{"ContextMappingName", contextMapping.ContextMappingName})
    contextMapping.EntityData.Leafs.Append("context", types.YLeaf{"Context", contextMapping.Context})
    contextMapping.EntityData.Leafs.Append("instance-name", types.YLeaf{"InstanceName", contextMapping.InstanceName})
    contextMapping.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", contextMapping.VrfName})
    contextMapping.EntityData.Leafs.Append("topology-name", types.YLeaf{"TopologyName", contextMapping.TopologyName})

    contextMapping.EntityData.YListKeys = []string {"ContextMappingName"}

    return &(contextMapping.EntityData)
}

// Snmp_Groups
// Define a User Security Model group
type Snmp_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the group. The type is slice of Snmp_Groups_Group.
    Group []*Snmp_Groups_Group
}

func (groups *Snmp_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "snmp"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = types.NewOrderedMap()
    groups.EntityData.Children.Append("group", types.YChild{"Group", nil})
    for i := range groups.Group {
        groups.EntityData.Children.Append(types.GetSegmentPath(groups.Group[i]), types.YChild{"Group", groups.Group[i]})
    }
    groups.EntityData.Leafs = types.NewOrderedMap()

    groups.EntityData.YListKeys = []string {}

    return &(groups.EntityData)
}

// Snmp_Groups_Group
// Name of the group
type Snmp_Groups_Group struct {
    EntityData types.CommonEntityData
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
    V4aclType interface{}

    // Ipv4 Access-list name. The type is string.
    V4AccessList interface{}

    // Access-list type. The type is Snmpacl.
    V6aclType interface{}

    // Ipv6 Access-list name. The type is string.
    V6AccessList interface{}

    // Context name. The type is string.
    ContextName interface{}
}

func (group *Snmp_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + types.AddKeyToken(group.Name, "name")
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = types.NewOrderedMap()
    group.EntityData.Leafs = types.NewOrderedMap()
    group.EntityData.Leafs.Append("name", types.YLeaf{"Name", group.Name})
    group.EntityData.Leafs.Append("snmp-version", types.YLeaf{"SnmpVersion", group.SnmpVersion})
    group.EntityData.Leafs.Append("security-model", types.YLeaf{"SecurityModel", group.SecurityModel})
    group.EntityData.Leafs.Append("notify-view", types.YLeaf{"NotifyView", group.NotifyView})
    group.EntityData.Leafs.Append("read-view", types.YLeaf{"ReadView", group.ReadView})
    group.EntityData.Leafs.Append("write-view", types.YLeaf{"WriteView", group.WriteView})
    group.EntityData.Leafs.Append("v4acl-type", types.YLeaf{"V4aclType", group.V4aclType})
    group.EntityData.Leafs.Append("v4-access-list", types.YLeaf{"V4AccessList", group.V4AccessList})
    group.EntityData.Leafs.Append("v6acl-type", types.YLeaf{"V6aclType", group.V6aclType})
    group.EntityData.Leafs.Append("v6-access-list", types.YLeaf{"V6AccessList", group.V6AccessList})
    group.EntityData.Leafs.Append("context-name", types.YLeaf{"ContextName", group.ContextName})

    group.EntityData.YListKeys = []string {"Name"}

    return &(group.EntityData)
}

// Snmp_TrapHosts
// Specify hosts to receive SNMP notifications
type Snmp_TrapHosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify hosts to receive SNMP notifications. The type is slice of
    // Snmp_TrapHosts_TrapHost.
    TrapHost []*Snmp_TrapHosts_TrapHost
}

func (trapHosts *Snmp_TrapHosts) GetEntityData() *types.CommonEntityData {
    trapHosts.EntityData.YFilter = trapHosts.YFilter
    trapHosts.EntityData.YangName = "trap-hosts"
    trapHosts.EntityData.BundleName = "cisco_ios_xr"
    trapHosts.EntityData.ParentYangName = "snmp"
    trapHosts.EntityData.SegmentPath = "trap-hosts"
    trapHosts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapHosts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapHosts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapHosts.EntityData.Children = types.NewOrderedMap()
    trapHosts.EntityData.Children.Append("trap-host", types.YChild{"TrapHost", nil})
    for i := range trapHosts.TrapHost {
        trapHosts.EntityData.Children.Append(types.GetSegmentPath(trapHosts.TrapHost[i]), types.YChild{"TrapHost", trapHosts.TrapHost[i]})
    }
    trapHosts.EntityData.Leafs = types.NewOrderedMap()

    trapHosts.EntityData.YListKeys = []string {}

    return &(trapHosts.EntityData)
}

// Snmp_TrapHosts_TrapHost
// Specify hosts to receive SNMP notifications
type Snmp_TrapHosts_TrapHost struct {
    EntityData types.CommonEntityData
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

func (trapHost *Snmp_TrapHosts_TrapHost) GetEntityData() *types.CommonEntityData {
    trapHost.EntityData.YFilter = trapHost.YFilter
    trapHost.EntityData.YangName = "trap-host"
    trapHost.EntityData.BundleName = "cisco_ios_xr"
    trapHost.EntityData.ParentYangName = "trap-hosts"
    trapHost.EntityData.SegmentPath = "trap-host" + types.AddKeyToken(trapHost.IpAddress, "ip-address")
    trapHost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapHost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapHost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapHost.EntityData.Children = types.NewOrderedMap()
    trapHost.EntityData.Children.Append("encrypted-user-communities", types.YChild{"EncryptedUserCommunities", &trapHost.EncryptedUserCommunities})
    trapHost.EntityData.Children.Append("inform-host", types.YChild{"InformHost", &trapHost.InformHost})
    trapHost.EntityData.Children.Append("default-user-communities", types.YChild{"DefaultUserCommunities", &trapHost.DefaultUserCommunities})
    trapHost.EntityData.Leafs = types.NewOrderedMap()
    trapHost.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", trapHost.IpAddress})

    trapHost.EntityData.YListKeys = []string {"IpAddress"}

    return &(trapHost.EntityData)
}

// Snmp_TrapHosts_TrapHost_EncryptedUserCommunities
// Container class for defining Clear/encrypt
// communities for a trap host
type Snmp_TrapHosts_TrapHost_EncryptedUserCommunities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear/Encrypt Community name associated with a trap host. The type is slice
    // of Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity.
    EncryptedUserCommunity []*Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity
}

func (encryptedUserCommunities *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities) GetEntityData() *types.CommonEntityData {
    encryptedUserCommunities.EntityData.YFilter = encryptedUserCommunities.YFilter
    encryptedUserCommunities.EntityData.YangName = "encrypted-user-communities"
    encryptedUserCommunities.EntityData.BundleName = "cisco_ios_xr"
    encryptedUserCommunities.EntityData.ParentYangName = "trap-host"
    encryptedUserCommunities.EntityData.SegmentPath = "encrypted-user-communities"
    encryptedUserCommunities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryptedUserCommunities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryptedUserCommunities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryptedUserCommunities.EntityData.Children = types.NewOrderedMap()
    encryptedUserCommunities.EntityData.Children.Append("encrypted-user-community", types.YChild{"EncryptedUserCommunity", nil})
    for i := range encryptedUserCommunities.EncryptedUserCommunity {
        encryptedUserCommunities.EntityData.Children.Append(types.GetSegmentPath(encryptedUserCommunities.EncryptedUserCommunity[i]), types.YChild{"EncryptedUserCommunity", encryptedUserCommunities.EncryptedUserCommunity[i]})
    }
    encryptedUserCommunities.EntityData.Leafs = types.NewOrderedMap()

    encryptedUserCommunities.EntityData.YListKeys = []string {}

    return &(encryptedUserCommunities.EntityData)
}

// Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity
// Clear/Encrypt Community name associated with
// a trap host
type Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity struct {
    EntityData types.CommonEntityData
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
    // range: 0..4294967295. This attribute is mandatory.
    BasicTrapTypes interface{}

    // Number to signify the feature traps that needs to be setUse this for
    // providing copy-complete trapValue must be set to 0 if not used. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    AdvancedTrapTypes1 interface{}

    // Number to signify the feature traps that needs to be setvalue should always
    // to set as 0. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    AdvancedTrapTypes2 interface{}
}

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetEntityData() *types.CommonEntityData {
    encryptedUserCommunity.EntityData.YFilter = encryptedUserCommunity.YFilter
    encryptedUserCommunity.EntityData.YangName = "encrypted-user-community"
    encryptedUserCommunity.EntityData.BundleName = "cisco_ios_xr"
    encryptedUserCommunity.EntityData.ParentYangName = "encrypted-user-communities"
    encryptedUserCommunity.EntityData.SegmentPath = "encrypted-user-community" + types.AddKeyToken(encryptedUserCommunity.CommunityName, "community-name")
    encryptedUserCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryptedUserCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryptedUserCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryptedUserCommunity.EntityData.Children = types.NewOrderedMap()
    encryptedUserCommunity.EntityData.Leafs = types.NewOrderedMap()
    encryptedUserCommunity.EntityData.Leafs.Append("community-name", types.YLeaf{"CommunityName", encryptedUserCommunity.CommunityName})
    encryptedUserCommunity.EntityData.Leafs.Append("port", types.YLeaf{"Port", encryptedUserCommunity.Port})
    encryptedUserCommunity.EntityData.Leafs.Append("version", types.YLeaf{"Version", encryptedUserCommunity.Version})
    encryptedUserCommunity.EntityData.Leafs.Append("security-level", types.YLeaf{"SecurityLevel", encryptedUserCommunity.SecurityLevel})
    encryptedUserCommunity.EntityData.Leafs.Append("basic-trap-types", types.YLeaf{"BasicTrapTypes", encryptedUserCommunity.BasicTrapTypes})
    encryptedUserCommunity.EntityData.Leafs.Append("advanced-trap-types1", types.YLeaf{"AdvancedTrapTypes1", encryptedUserCommunity.AdvancedTrapTypes1})
    encryptedUserCommunity.EntityData.Leafs.Append("advanced-trap-types2", types.YLeaf{"AdvancedTrapTypes2", encryptedUserCommunity.AdvancedTrapTypes2})

    encryptedUserCommunity.EntityData.YListKeys = []string {"CommunityName"}

    return &(encryptedUserCommunity.EntityData)
}

// Snmp_TrapHosts_TrapHost_InformHost
// Container class for defining notification type
// for a Inform host
type Snmp_TrapHosts_TrapHost_InformHost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Container class for defining communities for a inform host.
    InformUserCommunities Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities

    // Container class for defining Clear/encrypt communities for a inform host.
    InformEncryptedUserCommunities Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities
}

func (informHost *Snmp_TrapHosts_TrapHost_InformHost) GetEntityData() *types.CommonEntityData {
    informHost.EntityData.YFilter = informHost.YFilter
    informHost.EntityData.YangName = "inform-host"
    informHost.EntityData.BundleName = "cisco_ios_xr"
    informHost.EntityData.ParentYangName = "trap-host"
    informHost.EntityData.SegmentPath = "inform-host"
    informHost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    informHost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    informHost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    informHost.EntityData.Children = types.NewOrderedMap()
    informHost.EntityData.Children.Append("inform-user-communities", types.YChild{"InformUserCommunities", &informHost.InformUserCommunities})
    informHost.EntityData.Children.Append("inform-encrypted-user-communities", types.YChild{"InformEncryptedUserCommunities", &informHost.InformEncryptedUserCommunities})
    informHost.EntityData.Leafs = types.NewOrderedMap()

    informHost.EntityData.YListKeys = []string {}

    return &(informHost.EntityData)
}

// Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities
// Container class for defining communities for
// a inform host
type Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unencrpted Community name associated with a inform host. The type is slice
    // of
    // Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity.
    InformUserCommunity []*Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity
}

func (informUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities) GetEntityData() *types.CommonEntityData {
    informUserCommunities.EntityData.YFilter = informUserCommunities.YFilter
    informUserCommunities.EntityData.YangName = "inform-user-communities"
    informUserCommunities.EntityData.BundleName = "cisco_ios_xr"
    informUserCommunities.EntityData.ParentYangName = "inform-host"
    informUserCommunities.EntityData.SegmentPath = "inform-user-communities"
    informUserCommunities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    informUserCommunities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    informUserCommunities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    informUserCommunities.EntityData.Children = types.NewOrderedMap()
    informUserCommunities.EntityData.Children.Append("inform-user-community", types.YChild{"InformUserCommunity", nil})
    for i := range informUserCommunities.InformUserCommunity {
        informUserCommunities.EntityData.Children.Append(types.GetSegmentPath(informUserCommunities.InformUserCommunity[i]), types.YChild{"InformUserCommunity", informUserCommunities.InformUserCommunity[i]})
    }
    informUserCommunities.EntityData.Leafs = types.NewOrderedMap()

    informUserCommunities.EntityData.YListKeys = []string {}

    return &(informUserCommunities.EntityData)
}

// Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity
// Unencrpted Community name associated with a
// inform host
type Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity struct {
    EntityData types.CommonEntityData
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
    // range: 0..4294967295. This attribute is mandatory.
    BasicTrapTypes interface{}

    // Number to signify the feature traps that needs to be setUse this for
    // providing copy-complete trapValue must be set to 0 if not used. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    AdvancedTrapTypes1 interface{}

    // Number to signify the feature traps that needs to be setvalue should always
    // to set as 0. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    AdvancedTrapTypes2 interface{}
}

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetEntityData() *types.CommonEntityData {
    informUserCommunity.EntityData.YFilter = informUserCommunity.YFilter
    informUserCommunity.EntityData.YangName = "inform-user-community"
    informUserCommunity.EntityData.BundleName = "cisco_ios_xr"
    informUserCommunity.EntityData.ParentYangName = "inform-user-communities"
    informUserCommunity.EntityData.SegmentPath = "inform-user-community" + types.AddKeyToken(informUserCommunity.CommunityName, "community-name")
    informUserCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    informUserCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    informUserCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    informUserCommunity.EntityData.Children = types.NewOrderedMap()
    informUserCommunity.EntityData.Leafs = types.NewOrderedMap()
    informUserCommunity.EntityData.Leafs.Append("community-name", types.YLeaf{"CommunityName", informUserCommunity.CommunityName})
    informUserCommunity.EntityData.Leafs.Append("port", types.YLeaf{"Port", informUserCommunity.Port})
    informUserCommunity.EntityData.Leafs.Append("version", types.YLeaf{"Version", informUserCommunity.Version})
    informUserCommunity.EntityData.Leafs.Append("security-level", types.YLeaf{"SecurityLevel", informUserCommunity.SecurityLevel})
    informUserCommunity.EntityData.Leafs.Append("basic-trap-types", types.YLeaf{"BasicTrapTypes", informUserCommunity.BasicTrapTypes})
    informUserCommunity.EntityData.Leafs.Append("advanced-trap-types1", types.YLeaf{"AdvancedTrapTypes1", informUserCommunity.AdvancedTrapTypes1})
    informUserCommunity.EntityData.Leafs.Append("advanced-trap-types2", types.YLeaf{"AdvancedTrapTypes2", informUserCommunity.AdvancedTrapTypes2})

    informUserCommunity.EntityData.YListKeys = []string {"CommunityName"}

    return &(informUserCommunity.EntityData)
}

// Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities
// Container class for defining Clear/encrypt
// communities for a inform host
type Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Clear/Encrypt Community name associated with a inform host. The type is
    // slice of
    // Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity.
    InformEncryptedUserCommunity []*Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity
}

func (informEncryptedUserCommunities *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities) GetEntityData() *types.CommonEntityData {
    informEncryptedUserCommunities.EntityData.YFilter = informEncryptedUserCommunities.YFilter
    informEncryptedUserCommunities.EntityData.YangName = "inform-encrypted-user-communities"
    informEncryptedUserCommunities.EntityData.BundleName = "cisco_ios_xr"
    informEncryptedUserCommunities.EntityData.ParentYangName = "inform-host"
    informEncryptedUserCommunities.EntityData.SegmentPath = "inform-encrypted-user-communities"
    informEncryptedUserCommunities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    informEncryptedUserCommunities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    informEncryptedUserCommunities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    informEncryptedUserCommunities.EntityData.Children = types.NewOrderedMap()
    informEncryptedUserCommunities.EntityData.Children.Append("inform-encrypted-user-community", types.YChild{"InformEncryptedUserCommunity", nil})
    for i := range informEncryptedUserCommunities.InformEncryptedUserCommunity {
        informEncryptedUserCommunities.EntityData.Children.Append(types.GetSegmentPath(informEncryptedUserCommunities.InformEncryptedUserCommunity[i]), types.YChild{"InformEncryptedUserCommunity", informEncryptedUserCommunities.InformEncryptedUserCommunity[i]})
    }
    informEncryptedUserCommunities.EntityData.Leafs = types.NewOrderedMap()

    informEncryptedUserCommunities.EntityData.YListKeys = []string {}

    return &(informEncryptedUserCommunities.EntityData)
}

// Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity
// Clear/Encrypt Community name associated with
// a inform host
type Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity struct {
    EntityData types.CommonEntityData
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
    // range: 0..4294967295. This attribute is mandatory.
    BasicTrapTypes interface{}

    // Number to signify the feature traps that needs to be setUse this for
    // providing copy-complete trapValue must be set to 0 if not used. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    AdvancedTrapTypes1 interface{}

    // Number to signify the feature traps that needs to be setvalue should always
    // to set as 0. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    AdvancedTrapTypes2 interface{}
}

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetEntityData() *types.CommonEntityData {
    informEncryptedUserCommunity.EntityData.YFilter = informEncryptedUserCommunity.YFilter
    informEncryptedUserCommunity.EntityData.YangName = "inform-encrypted-user-community"
    informEncryptedUserCommunity.EntityData.BundleName = "cisco_ios_xr"
    informEncryptedUserCommunity.EntityData.ParentYangName = "inform-encrypted-user-communities"
    informEncryptedUserCommunity.EntityData.SegmentPath = "inform-encrypted-user-community" + types.AddKeyToken(informEncryptedUserCommunity.CommunityName, "community-name")
    informEncryptedUserCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    informEncryptedUserCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    informEncryptedUserCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    informEncryptedUserCommunity.EntityData.Children = types.NewOrderedMap()
    informEncryptedUserCommunity.EntityData.Leafs = types.NewOrderedMap()
    informEncryptedUserCommunity.EntityData.Leafs.Append("community-name", types.YLeaf{"CommunityName", informEncryptedUserCommunity.CommunityName})
    informEncryptedUserCommunity.EntityData.Leafs.Append("port", types.YLeaf{"Port", informEncryptedUserCommunity.Port})
    informEncryptedUserCommunity.EntityData.Leafs.Append("version", types.YLeaf{"Version", informEncryptedUserCommunity.Version})
    informEncryptedUserCommunity.EntityData.Leafs.Append("security-level", types.YLeaf{"SecurityLevel", informEncryptedUserCommunity.SecurityLevel})
    informEncryptedUserCommunity.EntityData.Leafs.Append("basic-trap-types", types.YLeaf{"BasicTrapTypes", informEncryptedUserCommunity.BasicTrapTypes})
    informEncryptedUserCommunity.EntityData.Leafs.Append("advanced-trap-types1", types.YLeaf{"AdvancedTrapTypes1", informEncryptedUserCommunity.AdvancedTrapTypes1})
    informEncryptedUserCommunity.EntityData.Leafs.Append("advanced-trap-types2", types.YLeaf{"AdvancedTrapTypes2", informEncryptedUserCommunity.AdvancedTrapTypes2})

    informEncryptedUserCommunity.EntityData.YListKeys = []string {"CommunityName"}

    return &(informEncryptedUserCommunity.EntityData)
}

// Snmp_TrapHosts_TrapHost_DefaultUserCommunities
// Container class for defining communities for a
// trap host
type Snmp_TrapHosts_TrapHost_DefaultUserCommunities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unencrpted Community name associated with a trap host. The type is slice of
    // Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity.
    DefaultUserCommunity []*Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity
}

func (defaultUserCommunities *Snmp_TrapHosts_TrapHost_DefaultUserCommunities) GetEntityData() *types.CommonEntityData {
    defaultUserCommunities.EntityData.YFilter = defaultUserCommunities.YFilter
    defaultUserCommunities.EntityData.YangName = "default-user-communities"
    defaultUserCommunities.EntityData.BundleName = "cisco_ios_xr"
    defaultUserCommunities.EntityData.ParentYangName = "trap-host"
    defaultUserCommunities.EntityData.SegmentPath = "default-user-communities"
    defaultUserCommunities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultUserCommunities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultUserCommunities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultUserCommunities.EntityData.Children = types.NewOrderedMap()
    defaultUserCommunities.EntityData.Children.Append("default-user-community", types.YChild{"DefaultUserCommunity", nil})
    for i := range defaultUserCommunities.DefaultUserCommunity {
        defaultUserCommunities.EntityData.Children.Append(types.GetSegmentPath(defaultUserCommunities.DefaultUserCommunity[i]), types.YChild{"DefaultUserCommunity", defaultUserCommunities.DefaultUserCommunity[i]})
    }
    defaultUserCommunities.EntityData.Leafs = types.NewOrderedMap()

    defaultUserCommunities.EntityData.YListKeys = []string {}

    return &(defaultUserCommunities.EntityData)
}

// Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity
// Unencrpted Community name associated with a
// trap host
type Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity struct {
    EntityData types.CommonEntityData
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
    // range: 0..4294967295. This attribute is mandatory.
    BasicTrapTypes interface{}

    // Number to signify the feature traps that needs to be setUse this for
    // providing copy-complete trapValue must be set to 0 if not used. The type is
    // interface{} with range: 0..4294967295. This attribute is mandatory.
    AdvancedTrapTypes1 interface{}

    // Number to signify the feature traps that needs to be setvalue should always
    // to set as 0. The type is interface{} with range: 0..4294967295. This
    // attribute is mandatory.
    AdvancedTrapTypes2 interface{}
}

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetEntityData() *types.CommonEntityData {
    defaultUserCommunity.EntityData.YFilter = defaultUserCommunity.YFilter
    defaultUserCommunity.EntityData.YangName = "default-user-community"
    defaultUserCommunity.EntityData.BundleName = "cisco_ios_xr"
    defaultUserCommunity.EntityData.ParentYangName = "default-user-communities"
    defaultUserCommunity.EntityData.SegmentPath = "default-user-community" + types.AddKeyToken(defaultUserCommunity.CommunityName, "community-name")
    defaultUserCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultUserCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultUserCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultUserCommunity.EntityData.Children = types.NewOrderedMap()
    defaultUserCommunity.EntityData.Leafs = types.NewOrderedMap()
    defaultUserCommunity.EntityData.Leafs.Append("community-name", types.YLeaf{"CommunityName", defaultUserCommunity.CommunityName})
    defaultUserCommunity.EntityData.Leafs.Append("port", types.YLeaf{"Port", defaultUserCommunity.Port})
    defaultUserCommunity.EntityData.Leafs.Append("version", types.YLeaf{"Version", defaultUserCommunity.Version})
    defaultUserCommunity.EntityData.Leafs.Append("security-level", types.YLeaf{"SecurityLevel", defaultUserCommunity.SecurityLevel})
    defaultUserCommunity.EntityData.Leafs.Append("basic-trap-types", types.YLeaf{"BasicTrapTypes", defaultUserCommunity.BasicTrapTypes})
    defaultUserCommunity.EntityData.Leafs.Append("advanced-trap-types1", types.YLeaf{"AdvancedTrapTypes1", defaultUserCommunity.AdvancedTrapTypes1})
    defaultUserCommunity.EntityData.Leafs.Append("advanced-trap-types2", types.YLeaf{"AdvancedTrapTypes2", defaultUserCommunity.AdvancedTrapTypes2})

    defaultUserCommunity.EntityData.YListKeys = []string {"CommunityName"}

    return &(defaultUserCommunity.EntityData)
}

// Snmp_Contexts
// List of Context Names
type Snmp_Contexts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Context Name. The type is slice of Snmp_Contexts_Context.
    Context []*Snmp_Contexts_Context
}

func (contexts *Snmp_Contexts) GetEntityData() *types.CommonEntityData {
    contexts.EntityData.YFilter = contexts.YFilter
    contexts.EntityData.YangName = "contexts"
    contexts.EntityData.BundleName = "cisco_ios_xr"
    contexts.EntityData.ParentYangName = "snmp"
    contexts.EntityData.SegmentPath = "contexts"
    contexts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contexts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contexts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contexts.EntityData.Children = types.NewOrderedMap()
    contexts.EntityData.Children.Append("context", types.YChild{"Context", nil})
    for i := range contexts.Context {
        contexts.EntityData.Children.Append(types.GetSegmentPath(contexts.Context[i]), types.YChild{"Context", contexts.Context[i]})
    }
    contexts.EntityData.Leafs = types.NewOrderedMap()

    contexts.EntityData.YListKeys = []string {}

    return &(contexts.EntityData)
}

// Snmp_Contexts_Context
// Context Name
type Snmp_Contexts_Context struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Context Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ContextName interface{}
}

func (context *Snmp_Contexts_Context) GetEntityData() *types.CommonEntityData {
    context.EntityData.YFilter = context.YFilter
    context.EntityData.YangName = "context"
    context.EntityData.BundleName = "cisco_ios_xr"
    context.EntityData.ParentYangName = "contexts"
    context.EntityData.SegmentPath = "context" + types.AddKeyToken(context.ContextName, "context-name")
    context.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    context.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    context.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    context.EntityData.Children = types.NewOrderedMap()
    context.EntityData.Leafs = types.NewOrderedMap()
    context.EntityData.Leafs.Append("context-name", types.YLeaf{"ContextName", context.ContextName})

    context.EntityData.YListKeys = []string {"ContextName"}

    return &(context.EntityData)
}

// Snmp_ContextMappings
// List of context names
type Snmp_ContextMappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Context mapping name. The type is slice of
    // Snmp_ContextMappings_ContextMapping.
    ContextMapping []*Snmp_ContextMappings_ContextMapping
}

func (contextMappings *Snmp_ContextMappings) GetEntityData() *types.CommonEntityData {
    contextMappings.EntityData.YFilter = contextMappings.YFilter
    contextMappings.EntityData.YangName = "context-mappings"
    contextMappings.EntityData.BundleName = "cisco_ios_xr"
    contextMappings.EntityData.ParentYangName = "snmp"
    contextMappings.EntityData.SegmentPath = "context-mappings"
    contextMappings.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextMappings.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextMappings.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextMappings.EntityData.Children = types.NewOrderedMap()
    contextMappings.EntityData.Children.Append("context-mapping", types.YChild{"ContextMapping", nil})
    for i := range contextMappings.ContextMapping {
        contextMappings.EntityData.Children.Append(types.GetSegmentPath(contextMappings.ContextMapping[i]), types.YChild{"ContextMapping", contextMappings.ContextMapping[i]})
    }
    contextMappings.EntityData.Leafs = types.NewOrderedMap()

    contextMappings.EntityData.YListKeys = []string {}

    return &(contextMappings.EntityData)
}

// Snmp_ContextMappings_ContextMapping
// Context mapping name
type Snmp_ContextMappings_ContextMapping struct {
    EntityData types.CommonEntityData
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

func (contextMapping *Snmp_ContextMappings_ContextMapping) GetEntityData() *types.CommonEntityData {
    contextMapping.EntityData.YFilter = contextMapping.YFilter
    contextMapping.EntityData.YangName = "context-mapping"
    contextMapping.EntityData.BundleName = "cisco_ios_xr"
    contextMapping.EntityData.ParentYangName = "context-mappings"
    contextMapping.EntityData.SegmentPath = "context-mapping" + types.AddKeyToken(contextMapping.ContextMappingName, "context-mapping-name")
    contextMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextMapping.EntityData.Children = types.NewOrderedMap()
    contextMapping.EntityData.Leafs = types.NewOrderedMap()
    contextMapping.EntityData.Leafs.Append("context-mapping-name", types.YLeaf{"ContextMappingName", contextMapping.ContextMappingName})
    contextMapping.EntityData.Leafs.Append("context", types.YLeaf{"Context", contextMapping.Context})
    contextMapping.EntityData.Leafs.Append("instance-name", types.YLeaf{"InstanceName", contextMapping.InstanceName})
    contextMapping.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", contextMapping.VrfName})
    contextMapping.EntityData.Leafs.Append("topology-name", types.YLeaf{"TopologyName", contextMapping.TopologyName})

    contextMapping.EntityData.YListKeys = []string {"ContextMappingName"}

    return &(contextMapping.EntityData)
}

// Mib
// mib
type Mib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Get cached Sesnsor MIB statistics. The type is interface{}.
    SensorMibCache interface{}

    // Interface MIB configuration.
    InterfaceMib Mib_InterfaceMib

    // MPLS TE MIB configuration.
    MplsTeMib Mib_MplsTeMib

    // MPLS P2MP MIB configuration.
    MplsP2mpMib Mib_MplsP2mpMib

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

    // Subscriber threshold commands.
    Subscriber Mib_Subscriber
}

func (mib *Mib) GetEntityData() *types.CommonEntityData {
    mib.EntityData.YFilter = mib.YFilter
    mib.EntityData.YangName = "mib"
    mib.EntityData.BundleName = "cisco_ios_xr"
    mib.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-agent-cfg"
    mib.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-agent-cfg:mib"
    mib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mib.EntityData.Children = types.NewOrderedMap()
    mib.EntityData.Children.Append("Cisco-IOS-XR-snmp-ifmib-cfg:interface-mib", types.YChild{"InterfaceMib", &mib.InterfaceMib})
    mib.EntityData.Children.Append("Cisco-IOS-XR-mpls-te-cfg:mpls-te-mib", types.YChild{"MplsTeMib", &mib.MplsTeMib})
    mib.EntityData.Children.Append("Cisco-IOS-XR-mpls-te-cfg:mpls-p2mp-mib", types.YChild{"MplsP2mpMib", &mib.MplsP2mpMib})
    mib.EntityData.Children.Append("Cisco-IOS-XR-mpls-te-cfg:mpls-te-ext-std-mib", types.YChild{"MplsTeExtStdMib", &mib.MplsTeExtStdMib})
    mib.EntityData.Children.Append("Cisco-IOS-XR-mpls-te-cfg:mpls-te-ext-mib", types.YChild{"MplsTeExtMib", &mib.MplsTeExtMib})
    mib.EntityData.Children.Append("Cisco-IOS-XR-mpls-te-cfg:mpls-frr-mib", types.YChild{"MplsFrrMib", &mib.MplsFrrMib})
    mib.EntityData.Children.Append("Cisco-IOS-XR-qos-mibs-cfg:cb-qosmib", types.YChild{"CbQosmib", &mib.CbQosmib})
    mib.EntityData.Children.Append("Cisco-IOS-XR-snmp-entitymib-cfg:entity-mib", types.YChild{"EntityMib", &mib.EntityMib})
    mib.EntityData.Children.Append("Cisco-IOS-XR-subscriber-session-mon-mibs-cfg:subscriber", types.YChild{"Subscriber", &mib.Subscriber})
    mib.EntityData.Leafs = types.NewOrderedMap()
    mib.EntityData.Leafs.Append("sensor-mib-cache", types.YLeaf{"SensorMibCache", mib.SensorMibCache})

    mib.EntityData.YListKeys = []string {}

    return &(mib.EntityData)
}

// Mib_InterfaceMib
// Interface MIB configuration
type Mib_InterfaceMib struct {
    EntityData types.CommonEntityData
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

func (interfaceMib *Mib_InterfaceMib) GetEntityData() *types.CommonEntityData {
    interfaceMib.EntityData.YFilter = interfaceMib.YFilter
    interfaceMib.EntityData.YangName = "interface-mib"
    interfaceMib.EntityData.BundleName = "cisco_ios_xr"
    interfaceMib.EntityData.ParentYangName = "mib"
    interfaceMib.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-ifmib-cfg:interface-mib"
    interfaceMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceMib.EntityData.Children = types.NewOrderedMap()
    interfaceMib.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &interfaceMib.Interfaces})
    interfaceMib.EntityData.Children.Append("notification", types.YChild{"Notification", &interfaceMib.Notification})
    interfaceMib.EntityData.Children.Append("subsets", types.YChild{"Subsets", &interfaceMib.Subsets})
    interfaceMib.EntityData.Leafs = types.NewOrderedMap()
    interfaceMib.EntityData.Leafs.Append("internal-cache", types.YLeaf{"InternalCache", interfaceMib.InternalCache})
    interfaceMib.EntityData.Leafs.Append("interface-alias-long", types.YLeaf{"InterfaceAliasLong", interfaceMib.InterfaceAliasLong})
    interfaceMib.EntityData.Leafs.Append("ip-subscriber", types.YLeaf{"IpSubscriber", interfaceMib.IpSubscriber})
    interfaceMib.EntityData.Leafs.Append("interface-index-persistence", types.YLeaf{"InterfaceIndexPersistence", interfaceMib.InterfaceIndexPersistence})
    interfaceMib.EntityData.Leafs.Append("statistics-cache", types.YLeaf{"StatisticsCache", interfaceMib.StatisticsCache})

    interfaceMib.EntityData.YListKeys = []string {}

    return &(interfaceMib.EntityData)
}

// Mib_InterfaceMib_Interfaces
// Enter the SNMP interface configuration commands
type Mib_InterfaceMib_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface to configure. The type is slice of
    // Mib_InterfaceMib_Interfaces_Interface.
    Interface []*Mib_InterfaceMib_Interfaces_Interface
}

func (interfaces *Mib_InterfaceMib_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "interface-mib"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Mib_InterfaceMib_Interfaces_Interface
// Interface to configure
type Mib_InterfaceMib_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Enable or disable LinkUpDown notification. The type is bool.
    LinkUpDown interface{}

    // Enable or disable index persistence. The type is bool.
    IndexPersistence interface{}
}

func (self *Mib_InterfaceMib_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("link-up-down", types.YLeaf{"LinkUpDown", self.LinkUpDown})
    self.EntityData.Leafs.Append("index-persistence", types.YLeaf{"IndexPersistence", self.IndexPersistence})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Mib_InterfaceMib_Notification
// MIB notification configuration
type Mib_InterfaceMib_Notification struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set the varbind of linkupdown trap to the RFC specified varbinds (default
    // cisco). The type is interface{}.
    LinkIetf interface{}
}

func (notification *Mib_InterfaceMib_Notification) GetEntityData() *types.CommonEntityData {
    notification.EntityData.YFilter = notification.YFilter
    notification.EntityData.YangName = "notification"
    notification.EntityData.BundleName = "cisco_ios_xr"
    notification.EntityData.ParentYangName = "interface-mib"
    notification.EntityData.SegmentPath = "notification"
    notification.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    notification.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    notification.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    notification.EntityData.Children = types.NewOrderedMap()
    notification.EntityData.Leafs = types.NewOrderedMap()
    notification.EntityData.Leafs.Append("link-ietf", types.YLeaf{"LinkIetf", notification.LinkIetf})

    notification.EntityData.YListKeys = []string {}

    return &(notification.EntityData)
}

// Mib_InterfaceMib_Subsets
// Add configuration for an interface subset
type Mib_InterfaceMib_Subsets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subset priorityID to group ifNames based on Regular Expression. The type is
    // slice of Mib_InterfaceMib_Subsets_Subset.
    Subset []*Mib_InterfaceMib_Subsets_Subset
}

func (subsets *Mib_InterfaceMib_Subsets) GetEntityData() *types.CommonEntityData {
    subsets.EntityData.YFilter = subsets.YFilter
    subsets.EntityData.YangName = "subsets"
    subsets.EntityData.BundleName = "cisco_ios_xr"
    subsets.EntityData.ParentYangName = "interface-mib"
    subsets.EntityData.SegmentPath = "subsets"
    subsets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subsets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subsets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subsets.EntityData.Children = types.NewOrderedMap()
    subsets.EntityData.Children.Append("subset", types.YChild{"Subset", nil})
    for i := range subsets.Subset {
        subsets.EntityData.Children.Append(types.GetSegmentPath(subsets.Subset[i]), types.YChild{"Subset", subsets.Subset[i]})
    }
    subsets.EntityData.Leafs = types.NewOrderedMap()

    subsets.EntityData.YListKeys = []string {}

    return &(subsets.EntityData)
}

// Mib_InterfaceMib_Subsets_Subset
// Subset priorityID to group ifNames based on
// Regular Expression
type Mib_InterfaceMib_Subsets_Subset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The interface subset PriorityID. The type is
    // interface{} with range: 1..255.
    SubsetId interface{}

    // SNMP linkUp and linkDown notifications.
    LinkUpDown Mib_InterfaceMib_Subsets_Subset_LinkUpDown
}

func (subset *Mib_InterfaceMib_Subsets_Subset) GetEntityData() *types.CommonEntityData {
    subset.EntityData.YFilter = subset.YFilter
    subset.EntityData.YangName = "subset"
    subset.EntityData.BundleName = "cisco_ios_xr"
    subset.EntityData.ParentYangName = "subsets"
    subset.EntityData.SegmentPath = "subset" + types.AddKeyToken(subset.SubsetId, "subset-id")
    subset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subset.EntityData.Children = types.NewOrderedMap()
    subset.EntityData.Children.Append("link-up-down", types.YChild{"LinkUpDown", &subset.LinkUpDown})
    subset.EntityData.Leafs = types.NewOrderedMap()
    subset.EntityData.Leafs.Append("subset-id", types.YLeaf{"SubsetId", subset.SubsetId})

    subset.EntityData.YListKeys = []string {"SubsetId"}

    return &(subset.EntityData)
}

// Mib_InterfaceMib_Subsets_Subset_LinkUpDown
// SNMP linkUp and linkDown notifications
type Mib_InterfaceMib_Subsets_Subset_LinkUpDown struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable or disable linkupdown notification. The type is bool.
    Enable interface{}

    // Regular expression to match ifName. The type is string.
    RegularExpression interface{}
}

func (linkUpDown *Mib_InterfaceMib_Subsets_Subset_LinkUpDown) GetEntityData() *types.CommonEntityData {
    linkUpDown.EntityData.YFilter = linkUpDown.YFilter
    linkUpDown.EntityData.YangName = "link-up-down"
    linkUpDown.EntityData.BundleName = "cisco_ios_xr"
    linkUpDown.EntityData.ParentYangName = "subset"
    linkUpDown.EntityData.SegmentPath = "link-up-down"
    linkUpDown.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkUpDown.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkUpDown.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkUpDown.EntityData.Children = types.NewOrderedMap()
    linkUpDown.EntityData.Leafs = types.NewOrderedMap()
    linkUpDown.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", linkUpDown.Enable})
    linkUpDown.EntityData.Leafs.Append("regular-expression", types.YLeaf{"RegularExpression", linkUpDown.RegularExpression})

    linkUpDown.EntityData.YListKeys = []string {}

    return &(linkUpDown.EntityData)
}

// Mib_MplsTeMib
// MPLS TE MIB configuration
type Mib_MplsTeMib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the cache garbage collector time for the mib. The type is
    // interface{} with range: 0..3600. Units are second. The default value is
    // 1800.
    CacheGarbageCollectTimer interface{}

    // Configure the cache time for the mib. The type is interface{} with range:
    // 0..600. Units are second. The default value is 60.
    CacheTimer interface{}
}

func (mplsTeMib *Mib_MplsTeMib) GetEntityData() *types.CommonEntityData {
    mplsTeMib.EntityData.YFilter = mplsTeMib.YFilter
    mplsTeMib.EntityData.YangName = "mpls-te-mib"
    mplsTeMib.EntityData.BundleName = "cisco_ios_xr"
    mplsTeMib.EntityData.ParentYangName = "mib"
    mplsTeMib.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-te-cfg:mpls-te-mib"
    mplsTeMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsTeMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsTeMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsTeMib.EntityData.Children = types.NewOrderedMap()
    mplsTeMib.EntityData.Leafs = types.NewOrderedMap()
    mplsTeMib.EntityData.Leafs.Append("cache-garbage-collect-timer", types.YLeaf{"CacheGarbageCollectTimer", mplsTeMib.CacheGarbageCollectTimer})
    mplsTeMib.EntityData.Leafs.Append("cache-timer", types.YLeaf{"CacheTimer", mplsTeMib.CacheTimer})

    mplsTeMib.EntityData.YListKeys = []string {}

    return &(mplsTeMib.EntityData)
}

// Mib_MplsP2mpMib
// MPLS P2MP MIB configuration
type Mib_MplsP2mpMib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the cache time for the mib. The type is interface{} with range:
    // 0..600. Units are second. The default value is 60.
    CacheTimer interface{}
}

func (mplsP2mpMib *Mib_MplsP2mpMib) GetEntityData() *types.CommonEntityData {
    mplsP2mpMib.EntityData.YFilter = mplsP2mpMib.YFilter
    mplsP2mpMib.EntityData.YangName = "mpls-p2mp-mib"
    mplsP2mpMib.EntityData.BundleName = "cisco_ios_xr"
    mplsP2mpMib.EntityData.ParentYangName = "mib"
    mplsP2mpMib.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-te-cfg:mpls-p2mp-mib"
    mplsP2mpMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsP2mpMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsP2mpMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsP2mpMib.EntityData.Children = types.NewOrderedMap()
    mplsP2mpMib.EntityData.Leafs = types.NewOrderedMap()
    mplsP2mpMib.EntityData.Leafs.Append("cache-timer", types.YLeaf{"CacheTimer", mplsP2mpMib.CacheTimer})

    mplsP2mpMib.EntityData.YListKeys = []string {}

    return &(mplsP2mpMib.EntityData)
}

// Mib_MplsTeExtStdMib
// MPLS TE EXT STD MIB configuration
type Mib_MplsTeExtStdMib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the cache time for the mib. The type is interface{} with range:
    // 0..600. Units are second. The default value is 60.
    CacheTimer interface{}
}

func (mplsTeExtStdMib *Mib_MplsTeExtStdMib) GetEntityData() *types.CommonEntityData {
    mplsTeExtStdMib.EntityData.YFilter = mplsTeExtStdMib.YFilter
    mplsTeExtStdMib.EntityData.YangName = "mpls-te-ext-std-mib"
    mplsTeExtStdMib.EntityData.BundleName = "cisco_ios_xr"
    mplsTeExtStdMib.EntityData.ParentYangName = "mib"
    mplsTeExtStdMib.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-te-cfg:mpls-te-ext-std-mib"
    mplsTeExtStdMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsTeExtStdMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsTeExtStdMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsTeExtStdMib.EntityData.Children = types.NewOrderedMap()
    mplsTeExtStdMib.EntityData.Leafs = types.NewOrderedMap()
    mplsTeExtStdMib.EntityData.Leafs.Append("cache-timer", types.YLeaf{"CacheTimer", mplsTeExtStdMib.CacheTimer})

    mplsTeExtStdMib.EntityData.YListKeys = []string {}

    return &(mplsTeExtStdMib.EntityData)
}

// Mib_MplsTeExtMib
// MPLS TE EXT MIB configuration
type Mib_MplsTeExtMib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the cache time for the mib. The type is interface{} with range:
    // 0..600. Units are second. The default value is 60.
    CacheTimer interface{}
}

func (mplsTeExtMib *Mib_MplsTeExtMib) GetEntityData() *types.CommonEntityData {
    mplsTeExtMib.EntityData.YFilter = mplsTeExtMib.YFilter
    mplsTeExtMib.EntityData.YangName = "mpls-te-ext-mib"
    mplsTeExtMib.EntityData.BundleName = "cisco_ios_xr"
    mplsTeExtMib.EntityData.ParentYangName = "mib"
    mplsTeExtMib.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-te-cfg:mpls-te-ext-mib"
    mplsTeExtMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsTeExtMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsTeExtMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsTeExtMib.EntityData.Children = types.NewOrderedMap()
    mplsTeExtMib.EntityData.Leafs = types.NewOrderedMap()
    mplsTeExtMib.EntityData.Leafs.Append("cache-timer", types.YLeaf{"CacheTimer", mplsTeExtMib.CacheTimer})

    mplsTeExtMib.EntityData.YListKeys = []string {}

    return &(mplsTeExtMib.EntityData)
}

// Mib_MplsFrrMib
// FRR MIB configuration
type Mib_MplsFrrMib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the cache time for the mib. The type is interface{} with range:
    // 0..600. Units are second. The default value is 60.
    CacheTimer interface{}
}

func (mplsFrrMib *Mib_MplsFrrMib) GetEntityData() *types.CommonEntityData {
    mplsFrrMib.EntityData.YFilter = mplsFrrMib.YFilter
    mplsFrrMib.EntityData.YangName = "mpls-frr-mib"
    mplsFrrMib.EntityData.BundleName = "cisco_ios_xr"
    mplsFrrMib.EntityData.ParentYangName = "mib"
    mplsFrrMib.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-te-cfg:mpls-frr-mib"
    mplsFrrMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsFrrMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsFrrMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsFrrMib.EntityData.Children = types.NewOrderedMap()
    mplsFrrMib.EntityData.Leafs = types.NewOrderedMap()
    mplsFrrMib.EntityData.Leafs.Append("cache-timer", types.YLeaf{"CacheTimer", mplsFrrMib.CacheTimer})

    mplsFrrMib.EntityData.YListKeys = []string {}

    return &(mplsFrrMib.EntityData)
}

// Mib_CbQosmib
// CBQoSMIB configuration
type Mib_CbQosmib struct {
    EntityData types.CommonEntityData
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

func (cbQosmib *Mib_CbQosmib) GetEntityData() *types.CommonEntityData {
    cbQosmib.EntityData.YFilter = cbQosmib.YFilter
    cbQosmib.EntityData.YangName = "cb-qosmib"
    cbQosmib.EntityData.BundleName = "cisco_ios_xr"
    cbQosmib.EntityData.ParentYangName = "mib"
    cbQosmib.EntityData.SegmentPath = "Cisco-IOS-XR-qos-mibs-cfg:cb-qosmib"
    cbQosmib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cbQosmib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cbQosmib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cbQosmib.EntityData.Children = types.NewOrderedMap()
    cbQosmib.EntityData.Children.Append("cache", types.YChild{"Cache", &cbQosmib.Cache})
    cbQosmib.EntityData.Leafs = types.NewOrderedMap()
    cbQosmib.EntityData.Leafs.Append("member-interface-stats", types.YLeaf{"MemberInterfaceStats", cbQosmib.MemberInterfaceStats})
    cbQosmib.EntityData.Leafs.Append("persist", types.YLeaf{"Persist", cbQosmib.Persist})

    cbQosmib.EntityData.YListKeys = []string {}

    return &(cbQosmib.EntityData)
}

// Mib_CbQosmib_Cache
// CBQoSMIB statistics data caching
type Mib_CbQosmib_Cache struct {
    EntityData types.CommonEntityData
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

func (cache *Mib_CbQosmib_Cache) GetEntityData() *types.CommonEntityData {
    cache.EntityData.YFilter = cache.YFilter
    cache.EntityData.YangName = "cache"
    cache.EntityData.BundleName = "cisco_ios_xr"
    cache.EntityData.ParentYangName = "cb-qosmib"
    cache.EntityData.SegmentPath = "cache"
    cache.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cache.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cache.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cache.EntityData.Children = types.NewOrderedMap()
    cache.EntityData.Leafs = types.NewOrderedMap()
    cache.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", cache.Enable})
    cache.EntityData.Leafs.Append("refresh-time", types.YLeaf{"RefreshTime", cache.RefreshTime})
    cache.EntityData.Leafs.Append("service-policy-count", types.YLeaf{"ServicePolicyCount", cache.ServicePolicyCount})

    cache.EntityData.YListKeys = []string {}

    return &(cache.EntityData)
}

// Mib_EntityMib
// Entity MIB
type Mib_EntityMib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable entPhysicalIndex persistence. The type is interface{}.
    EntityIndexPersistence interface{}
}

func (entityMib *Mib_EntityMib) GetEntityData() *types.CommonEntityData {
    entityMib.EntityData.YFilter = entityMib.YFilter
    entityMib.EntityData.YangName = "entity-mib"
    entityMib.EntityData.BundleName = "cisco_ios_xr"
    entityMib.EntityData.ParentYangName = "mib"
    entityMib.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-entitymib-cfg:entity-mib"
    entityMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    entityMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    entityMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    entityMib.EntityData.Children = types.NewOrderedMap()
    entityMib.EntityData.Leafs = types.NewOrderedMap()
    entityMib.EntityData.Leafs.Append("entity-index-persistence", types.YLeaf{"EntityIndexPersistence", entityMib.EntityIndexPersistence})

    entityMib.EntityData.YListKeys = []string {}

    return &(entityMib.EntityData)
}

// Mib_Subscriber
// Subscriber threshold commands
type Mib_Subscriber struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subscriber threshold commands.
    Threshold Mib_Subscriber_Threshold
}

func (subscriber *Mib_Subscriber) GetEntityData() *types.CommonEntityData {
    subscriber.EntityData.YFilter = subscriber.YFilter
    subscriber.EntityData.YangName = "subscriber"
    subscriber.EntityData.BundleName = "cisco_ios_xr"
    subscriber.EntityData.ParentYangName = "mib"
    subscriber.EntityData.SegmentPath = "Cisco-IOS-XR-subscriber-session-mon-mibs-cfg:subscriber"
    subscriber.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subscriber.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subscriber.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subscriber.EntityData.Children = types.NewOrderedMap()
    subscriber.EntityData.Children.Append("threshold", types.YChild{"Threshold", &subscriber.Threshold})
    subscriber.EntityData.Leafs = types.NewOrderedMap()

    subscriber.EntityData.YListKeys = []string {}

    return &(subscriber.EntityData)
}

// Mib_Subscriber_Threshold
// Subscriber threshold commands
type Mib_Subscriber_Threshold struct {
    EntityData types.CommonEntityData
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

func (threshold *Mib_Subscriber_Threshold) GetEntityData() *types.CommonEntityData {
    threshold.EntityData.YFilter = threshold.YFilter
    threshold.EntityData.YangName = "threshold"
    threshold.EntityData.BundleName = "cisco_ios_xr"
    threshold.EntityData.ParentYangName = "subscriber"
    threshold.EntityData.SegmentPath = "threshold"
    threshold.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    threshold.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    threshold.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    threshold.EntityData.Children = types.NewOrderedMap()
    threshold.EntityData.Children.Append("delta", types.YChild{"Delta", &threshold.Delta})
    threshold.EntityData.Children.Append("access-interface-sub", types.YChild{"AccessInterfaceSub", &threshold.AccessInterfaceSub})
    threshold.EntityData.Children.Append("falling", types.YChild{"Falling", &threshold.Falling})
    threshold.EntityData.Children.Append("rising", types.YChild{"Rising", &threshold.Rising})
    threshold.EntityData.Leafs = types.NewOrderedMap()

    threshold.EntityData.YListKeys = []string {}

    return &(threshold.EntityData)
}

// Mib_Subscriber_Threshold_Delta
// Delta loss keyword
type Mib_Subscriber_Threshold_Delta struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Evaluation keyword.
    Evaluation Mib_Subscriber_Threshold_Delta_Evaluation

    // Delta loss percent.
    Percent Mib_Subscriber_Threshold_Delta_Percent
}

func (delta *Mib_Subscriber_Threshold_Delta) GetEntityData() *types.CommonEntityData {
    delta.EntityData.YFilter = delta.YFilter
    delta.EntityData.YangName = "delta"
    delta.EntityData.BundleName = "cisco_ios_xr"
    delta.EntityData.ParentYangName = "threshold"
    delta.EntityData.SegmentPath = "delta"
    delta.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delta.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delta.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delta.EntityData.Children = types.NewOrderedMap()
    delta.EntityData.Children.Append("evaluation", types.YChild{"Evaluation", &delta.Evaluation})
    delta.EntityData.Children.Append("percent", types.YChild{"Percent", &delta.Percent})
    delta.EntityData.Leafs = types.NewOrderedMap()

    delta.EntityData.YListKeys = []string {}

    return &(delta.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Evaluation
// Evaluation keyword
type Mib_Subscriber_Threshold_Delta_Evaluation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of AccessInterface.
    AccessInterfaces Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces

    // Table of Node.
    Nodes Mib_Subscriber_Threshold_Delta_Evaluation_Nodes
}

func (evaluation *Mib_Subscriber_Threshold_Delta_Evaluation) GetEntityData() *types.CommonEntityData {
    evaluation.EntityData.YFilter = evaluation.YFilter
    evaluation.EntityData.YangName = "evaluation"
    evaluation.EntityData.BundleName = "cisco_ios_xr"
    evaluation.EntityData.ParentYangName = "delta"
    evaluation.EntityData.SegmentPath = "evaluation"
    evaluation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    evaluation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    evaluation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    evaluation.EntityData.Children = types.NewOrderedMap()
    evaluation.EntityData.Children.Append("access-interfaces", types.YChild{"AccessInterfaces", &evaluation.AccessInterfaces})
    evaluation.EntityData.Children.Append("nodes", types.YChild{"Nodes", &evaluation.Nodes})
    evaluation.EntityData.Leafs = types.NewOrderedMap()

    evaluation.EntityData.YListKeys = []string {}

    return &(evaluation.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces
// Table of AccessInterface
type Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access interface. The type is slice of
    // Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface.
    AccessInterface []*Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface
}

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces) GetEntityData() *types.CommonEntityData {
    accessInterfaces.EntityData.YFilter = accessInterfaces.YFilter
    accessInterfaces.EntityData.YangName = "access-interfaces"
    accessInterfaces.EntityData.BundleName = "cisco_ios_xr"
    accessInterfaces.EntityData.ParentYangName = "evaluation"
    accessInterfaces.EntityData.SegmentPath = "access-interfaces"
    accessInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterfaces.EntityData.Children = types.NewOrderedMap()
    accessInterfaces.EntityData.Children.Append("access-interface", types.YChild{"AccessInterface", nil})
    for i := range accessInterfaces.AccessInterface {
        accessInterfaces.EntityData.Children.Append(types.GetSegmentPath(accessInterfaces.AccessInterface[i]), types.YChild{"AccessInterface", accessInterfaces.AccessInterface[i]})
    }
    accessInterfaces.EntityData.Leafs = types.NewOrderedMap()

    accessInterfaces.EntityData.YListKeys = []string {}

    return &(accessInterfaces.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface
// Access interface
type Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface struct {
    EntityData types.CommonEntityData
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

func (accessInterface *Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface) GetEntityData() *types.CommonEntityData {
    accessInterface.EntityData.YFilter = accessInterface.YFilter
    accessInterface.EntityData.YangName = "access-interface"
    accessInterface.EntityData.BundleName = "cisco_ios_xr"
    accessInterface.EntityData.ParentYangName = "access-interfaces"
    accessInterface.EntityData.SegmentPath = "access-interface" + types.AddKeyToken(accessInterface.InterfaceName, "interface-name")
    accessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterface.EntityData.Children = types.NewOrderedMap()
    accessInterface.EntityData.Leafs = types.NewOrderedMap()
    accessInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", accessInterface.InterfaceName})
    accessInterface.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", accessInterface.SessionCount})
    accessInterface.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", accessInterface.Interval})

    accessInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(accessInterface.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Evaluation_Nodes
// Table of Node
type Mib_Subscriber_Threshold_Delta_Evaluation_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rising node level. The type is slice of
    // Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node.
    Node []*Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node
}

func (nodes *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "evaluation"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node
// Rising node level
type Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node struct {
    EntityData types.CommonEntityData
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

func (node *Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})
    node.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", node.SessionCount})
    node.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", node.Interval})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Percent
// Delta loss percent
type Mib_Subscriber_Threshold_Delta_Percent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of AccessInterface.
    AccessInterfaces Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces

    // Table of Node.
    Nodes Mib_Subscriber_Threshold_Delta_Percent_Nodes
}

func (percent *Mib_Subscriber_Threshold_Delta_Percent) GetEntityData() *types.CommonEntityData {
    percent.EntityData.YFilter = percent.YFilter
    percent.EntityData.YangName = "percent"
    percent.EntityData.BundleName = "cisco_ios_xr"
    percent.EntityData.ParentYangName = "delta"
    percent.EntityData.SegmentPath = "percent"
    percent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    percent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    percent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    percent.EntityData.Children = types.NewOrderedMap()
    percent.EntityData.Children.Append("access-interfaces", types.YChild{"AccessInterfaces", &percent.AccessInterfaces})
    percent.EntityData.Children.Append("nodes", types.YChild{"Nodes", &percent.Nodes})
    percent.EntityData.Leafs = types.NewOrderedMap()

    percent.EntityData.YListKeys = []string {}

    return &(percent.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces
// Table of AccessInterface
type Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access interface. The type is slice of
    // Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface.
    AccessInterface []*Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface
}

func (accessInterfaces *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces) GetEntityData() *types.CommonEntityData {
    accessInterfaces.EntityData.YFilter = accessInterfaces.YFilter
    accessInterfaces.EntityData.YangName = "access-interfaces"
    accessInterfaces.EntityData.BundleName = "cisco_ios_xr"
    accessInterfaces.EntityData.ParentYangName = "percent"
    accessInterfaces.EntityData.SegmentPath = "access-interfaces"
    accessInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterfaces.EntityData.Children = types.NewOrderedMap()
    accessInterfaces.EntityData.Children.Append("access-interface", types.YChild{"AccessInterface", nil})
    for i := range accessInterfaces.AccessInterface {
        accessInterfaces.EntityData.Children.Append(types.GetSegmentPath(accessInterfaces.AccessInterface[i]), types.YChild{"AccessInterface", accessInterfaces.AccessInterface[i]})
    }
    accessInterfaces.EntityData.Leafs = types.NewOrderedMap()

    accessInterfaces.EntityData.YListKeys = []string {}

    return &(accessInterfaces.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface
// Access interface
type Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface struct {
    EntityData types.CommonEntityData
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

func (accessInterface *Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface) GetEntityData() *types.CommonEntityData {
    accessInterface.EntityData.YFilter = accessInterface.YFilter
    accessInterface.EntityData.YangName = "access-interface"
    accessInterface.EntityData.BundleName = "cisco_ios_xr"
    accessInterface.EntityData.ParentYangName = "access-interfaces"
    accessInterface.EntityData.SegmentPath = "access-interface" + types.AddKeyToken(accessInterface.InterfaceName, "interface-name")
    accessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterface.EntityData.Children = types.NewOrderedMap()
    accessInterface.EntityData.Leafs = types.NewOrderedMap()
    accessInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", accessInterface.InterfaceName})
    accessInterface.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", accessInterface.SessionCount})
    accessInterface.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", accessInterface.Interval})

    accessInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(accessInterface.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Percent_Nodes
// Table of Node
type Mib_Subscriber_Threshold_Delta_Percent_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rising node level. The type is slice of
    // Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node.
    Node []*Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node
}

func (nodes *Mib_Subscriber_Threshold_Delta_Percent_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "percent"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node
// Rising node level
type Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node struct {
    EntityData types.CommonEntityData
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

func (node *Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})
    node.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", node.SessionCount})
    node.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", node.Interval})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Mib_Subscriber_Threshold_AccessInterfaceSub
// Access interface for regular expression
type Mib_Subscriber_Threshold_AccessInterfaceSub struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Subset.
    Subsets Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets
}

func (accessInterfaceSub *Mib_Subscriber_Threshold_AccessInterfaceSub) GetEntityData() *types.CommonEntityData {
    accessInterfaceSub.EntityData.YFilter = accessInterfaceSub.YFilter
    accessInterfaceSub.EntityData.YangName = "access-interface-sub"
    accessInterfaceSub.EntityData.BundleName = "cisco_ios_xr"
    accessInterfaceSub.EntityData.ParentYangName = "threshold"
    accessInterfaceSub.EntityData.SegmentPath = "access-interface-sub"
    accessInterfaceSub.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterfaceSub.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterfaceSub.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterfaceSub.EntityData.Children = types.NewOrderedMap()
    accessInterfaceSub.EntityData.Children.Append("subsets", types.YChild{"Subsets", &accessInterfaceSub.Subsets})
    accessInterfaceSub.EntityData.Leafs = types.NewOrderedMap()

    accessInterfaceSub.EntityData.YListKeys = []string {}

    return &(accessInterfaceSub.EntityData)
}

// Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets
// Table of Subset
type Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subset command. The type is slice of
    // Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset.
    Subset []*Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset
}

func (subsets *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets) GetEntityData() *types.CommonEntityData {
    subsets.EntityData.YFilter = subsets.YFilter
    subsets.EntityData.YangName = "subsets"
    subsets.EntityData.BundleName = "cisco_ios_xr"
    subsets.EntityData.ParentYangName = "access-interface-sub"
    subsets.EntityData.SegmentPath = "subsets"
    subsets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subsets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subsets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subsets.EntityData.Children = types.NewOrderedMap()
    subsets.EntityData.Children.Append("subset", types.YChild{"Subset", nil})
    for i := range subsets.Subset {
        subsets.EntityData.Children.Append(types.GetSegmentPath(subsets.Subset[i]), types.YChild{"Subset", subsets.Subset[i]})
    }
    subsets.EntityData.Leafs = types.NewOrderedMap()

    subsets.EntityData.YListKeys = []string {}

    return &(subsets.EntityData)
}

// Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset
// Subset command
type Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Subset number. The type is interface{} with range:
    // 1..255.
    SubsetId interface{}

    // Regular expression.
    RegularExpression Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression
}

func (subset *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset) GetEntityData() *types.CommonEntityData {
    subset.EntityData.YFilter = subset.YFilter
    subset.EntityData.YangName = "subset"
    subset.EntityData.BundleName = "cisco_ios_xr"
    subset.EntityData.ParentYangName = "subsets"
    subset.EntityData.SegmentPath = "subset" + types.AddKeyToken(subset.SubsetId, "subset-id")
    subset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subset.EntityData.Children = types.NewOrderedMap()
    subset.EntityData.Children.Append("regular-expression", types.YChild{"RegularExpression", &subset.RegularExpression})
    subset.EntityData.Leafs = types.NewOrderedMap()
    subset.EntityData.Leafs.Append("subset-id", types.YLeaf{"SubsetId", subset.SubsetId})

    subset.EntityData.YListKeys = []string {"SubsetId"}

    return &(subset.EntityData)
}

// Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression
// Regular expression
type Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Notification keyword.
    Notification Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification
}

func (regularExpression *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression) GetEntityData() *types.CommonEntityData {
    regularExpression.EntityData.YFilter = regularExpression.YFilter
    regularExpression.EntityData.YangName = "regular-expression"
    regularExpression.EntityData.BundleName = "cisco_ios_xr"
    regularExpression.EntityData.ParentYangName = "subset"
    regularExpression.EntityData.SegmentPath = "regular-expression"
    regularExpression.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    regularExpression.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    regularExpression.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    regularExpression.EntityData.Children = types.NewOrderedMap()
    regularExpression.EntityData.Children.Append("notification", types.YChild{"Notification", &regularExpression.Notification})
    regularExpression.EntityData.Leafs = types.NewOrderedMap()

    regularExpression.EntityData.YListKeys = []string {}

    return &(regularExpression.EntityData)
}

// Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification
// Notification keyword
type Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rising-falling threshold.
    RisingFalling Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling
}

func (notification *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification) GetEntityData() *types.CommonEntityData {
    notification.EntityData.YFilter = notification.YFilter
    notification.EntityData.YangName = "notification"
    notification.EntityData.BundleName = "cisco_ios_xr"
    notification.EntityData.ParentYangName = "regular-expression"
    notification.EntityData.SegmentPath = "notification"
    notification.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    notification.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    notification.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    notification.EntityData.Children = types.NewOrderedMap()
    notification.EntityData.Children.Append("rising-falling", types.YChild{"RisingFalling", &notification.RisingFalling})
    notification.EntityData.Leafs = types.NewOrderedMap()

    notification.EntityData.YListKeys = []string {}

    return &(notification.EntityData)
}

// Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling
// Rising-falling threshold
type Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable the notifications on access interfaces. The type is string.
    Disable interface{}
}

func (risingFalling *Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset_RegularExpression_Notification_RisingFalling) GetEntityData() *types.CommonEntityData {
    risingFalling.EntityData.YFilter = risingFalling.YFilter
    risingFalling.EntityData.YangName = "rising-falling"
    risingFalling.EntityData.BundleName = "cisco_ios_xr"
    risingFalling.EntityData.ParentYangName = "notification"
    risingFalling.EntityData.SegmentPath = "rising-falling"
    risingFalling.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    risingFalling.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    risingFalling.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    risingFalling.EntityData.Children = types.NewOrderedMap()
    risingFalling.EntityData.Leafs = types.NewOrderedMap()
    risingFalling.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", risingFalling.Disable})

    risingFalling.EntityData.YListKeys = []string {}

    return &(risingFalling.EntityData)
}

// Mib_Subscriber_Threshold_Falling
// Falling threshold
type Mib_Subscriber_Threshold_Falling struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of AccessInterface.
    AccessInterfaces Mib_Subscriber_Threshold_Falling_AccessInterfaces

    // Table of Node.
    Nodes Mib_Subscriber_Threshold_Falling_Nodes
}

func (falling *Mib_Subscriber_Threshold_Falling) GetEntityData() *types.CommonEntityData {
    falling.EntityData.YFilter = falling.YFilter
    falling.EntityData.YangName = "falling"
    falling.EntityData.BundleName = "cisco_ios_xr"
    falling.EntityData.ParentYangName = "threshold"
    falling.EntityData.SegmentPath = "falling"
    falling.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    falling.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    falling.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    falling.EntityData.Children = types.NewOrderedMap()
    falling.EntityData.Children.Append("access-interfaces", types.YChild{"AccessInterfaces", &falling.AccessInterfaces})
    falling.EntityData.Children.Append("nodes", types.YChild{"Nodes", &falling.Nodes})
    falling.EntityData.Leafs = types.NewOrderedMap()

    falling.EntityData.YListKeys = []string {}

    return &(falling.EntityData)
}

// Mib_Subscriber_Threshold_Falling_AccessInterfaces
// Table of AccessInterface
type Mib_Subscriber_Threshold_Falling_AccessInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access interface. The type is slice of
    // Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface.
    AccessInterface []*Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface
}

func (accessInterfaces *Mib_Subscriber_Threshold_Falling_AccessInterfaces) GetEntityData() *types.CommonEntityData {
    accessInterfaces.EntityData.YFilter = accessInterfaces.YFilter
    accessInterfaces.EntityData.YangName = "access-interfaces"
    accessInterfaces.EntityData.BundleName = "cisco_ios_xr"
    accessInterfaces.EntityData.ParentYangName = "falling"
    accessInterfaces.EntityData.SegmentPath = "access-interfaces"
    accessInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterfaces.EntityData.Children = types.NewOrderedMap()
    accessInterfaces.EntityData.Children.Append("access-interface", types.YChild{"AccessInterface", nil})
    for i := range accessInterfaces.AccessInterface {
        accessInterfaces.EntityData.Children.Append(types.GetSegmentPath(accessInterfaces.AccessInterface[i]), types.YChild{"AccessInterface", accessInterfaces.AccessInterface[i]})
    }
    accessInterfaces.EntityData.Leafs = types.NewOrderedMap()

    accessInterfaces.EntityData.YListKeys = []string {}

    return &(accessInterfaces.EntityData)
}

// Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface
// Access interface
type Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface struct {
    EntityData types.CommonEntityData
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

func (accessInterface *Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface) GetEntityData() *types.CommonEntityData {
    accessInterface.EntityData.YFilter = accessInterface.YFilter
    accessInterface.EntityData.YangName = "access-interface"
    accessInterface.EntityData.BundleName = "cisco_ios_xr"
    accessInterface.EntityData.ParentYangName = "access-interfaces"
    accessInterface.EntityData.SegmentPath = "access-interface" + types.AddKeyToken(accessInterface.InterfaceName, "interface-name")
    accessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterface.EntityData.Children = types.NewOrderedMap()
    accessInterface.EntityData.Leafs = types.NewOrderedMap()
    accessInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", accessInterface.InterfaceName})
    accessInterface.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", accessInterface.SessionCount})
    accessInterface.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", accessInterface.Interval})

    accessInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(accessInterface.EntityData)
}

// Mib_Subscriber_Threshold_Falling_Nodes
// Table of Node
type Mib_Subscriber_Threshold_Falling_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rising node level. The type is slice of
    // Mib_Subscriber_Threshold_Falling_Nodes_Node.
    Node []*Mib_Subscriber_Threshold_Falling_Nodes_Node
}

func (nodes *Mib_Subscriber_Threshold_Falling_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "falling"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// Mib_Subscriber_Threshold_Falling_Nodes_Node
// Rising node level
type Mib_Subscriber_Threshold_Falling_Nodes_Node struct {
    EntityData types.CommonEntityData
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

func (node *Mib_Subscriber_Threshold_Falling_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})
    node.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", node.SessionCount})
    node.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", node.Interval})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

// Mib_Subscriber_Threshold_Rising
// Rising threshold
type Mib_Subscriber_Threshold_Rising struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of AccessInterface.
    AccessInterfaces Mib_Subscriber_Threshold_Rising_AccessInterfaces

    // Table of Node.
    Nodes Mib_Subscriber_Threshold_Rising_Nodes
}

func (rising *Mib_Subscriber_Threshold_Rising) GetEntityData() *types.CommonEntityData {
    rising.EntityData.YFilter = rising.YFilter
    rising.EntityData.YangName = "rising"
    rising.EntityData.BundleName = "cisco_ios_xr"
    rising.EntityData.ParentYangName = "threshold"
    rising.EntityData.SegmentPath = "rising"
    rising.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rising.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rising.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rising.EntityData.Children = types.NewOrderedMap()
    rising.EntityData.Children.Append("access-interfaces", types.YChild{"AccessInterfaces", &rising.AccessInterfaces})
    rising.EntityData.Children.Append("nodes", types.YChild{"Nodes", &rising.Nodes})
    rising.EntityData.Leafs = types.NewOrderedMap()

    rising.EntityData.YListKeys = []string {}

    return &(rising.EntityData)
}

// Mib_Subscriber_Threshold_Rising_AccessInterfaces
// Table of AccessInterface
type Mib_Subscriber_Threshold_Rising_AccessInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access interface. The type is slice of
    // Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface.
    AccessInterface []*Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface
}

func (accessInterfaces *Mib_Subscriber_Threshold_Rising_AccessInterfaces) GetEntityData() *types.CommonEntityData {
    accessInterfaces.EntityData.YFilter = accessInterfaces.YFilter
    accessInterfaces.EntityData.YangName = "access-interfaces"
    accessInterfaces.EntityData.BundleName = "cisco_ios_xr"
    accessInterfaces.EntityData.ParentYangName = "rising"
    accessInterfaces.EntityData.SegmentPath = "access-interfaces"
    accessInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterfaces.EntityData.Children = types.NewOrderedMap()
    accessInterfaces.EntityData.Children.Append("access-interface", types.YChild{"AccessInterface", nil})
    for i := range accessInterfaces.AccessInterface {
        accessInterfaces.EntityData.Children.Append(types.GetSegmentPath(accessInterfaces.AccessInterface[i]), types.YChild{"AccessInterface", accessInterfaces.AccessInterface[i]})
    }
    accessInterfaces.EntityData.Leafs = types.NewOrderedMap()

    accessInterfaces.EntityData.YListKeys = []string {}

    return &(accessInterfaces.EntityData)
}

// Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface
// Access interface
type Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface struct {
    EntityData types.CommonEntityData
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

func (accessInterface *Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface) GetEntityData() *types.CommonEntityData {
    accessInterface.EntityData.YFilter = accessInterface.YFilter
    accessInterface.EntityData.YangName = "access-interface"
    accessInterface.EntityData.BundleName = "cisco_ios_xr"
    accessInterface.EntityData.ParentYangName = "access-interfaces"
    accessInterface.EntityData.SegmentPath = "access-interface" + types.AddKeyToken(accessInterface.InterfaceName, "interface-name")
    accessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterface.EntityData.Children = types.NewOrderedMap()
    accessInterface.EntityData.Leafs = types.NewOrderedMap()
    accessInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", accessInterface.InterfaceName})
    accessInterface.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", accessInterface.SessionCount})
    accessInterface.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", accessInterface.Interval})

    accessInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(accessInterface.EntityData)
}

// Mib_Subscriber_Threshold_Rising_Nodes
// Table of Node
type Mib_Subscriber_Threshold_Rising_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rising node level. The type is slice of
    // Mib_Subscriber_Threshold_Rising_Nodes_Node.
    Node []*Mib_Subscriber_Threshold_Rising_Nodes_Node
}

func (nodes *Mib_Subscriber_Threshold_Rising_Nodes) GetEntityData() *types.CommonEntityData {
    nodes.EntityData.YFilter = nodes.YFilter
    nodes.EntityData.YangName = "nodes"
    nodes.EntityData.BundleName = "cisco_ios_xr"
    nodes.EntityData.ParentYangName = "rising"
    nodes.EntityData.SegmentPath = "nodes"
    nodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nodes.EntityData.Children = types.NewOrderedMap()
    nodes.EntityData.Children.Append("node", types.YChild{"Node", nil})
    for i := range nodes.Node {
        nodes.EntityData.Children.Append(types.GetSegmentPath(nodes.Node[i]), types.YChild{"Node", nodes.Node[i]})
    }
    nodes.EntityData.Leafs = types.NewOrderedMap()

    nodes.EntityData.YListKeys = []string {}

    return &(nodes.EntityData)
}

// Mib_Subscriber_Threshold_Rising_Nodes_Node
// Rising node level
type Mib_Subscriber_Threshold_Rising_Nodes_Node struct {
    EntityData types.CommonEntityData
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

func (node *Mib_Subscriber_Threshold_Rising_Nodes_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "nodes"
    node.EntityData.SegmentPath = "node" + types.AddKeyToken(node.NodeName, "node-name")
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = types.NewOrderedMap()
    node.EntityData.Leafs = types.NewOrderedMap()
    node.EntityData.Leafs.Append("node-name", types.YLeaf{"NodeName", node.NodeName})
    node.EntityData.Leafs.Append("session-count", types.YLeaf{"SessionCount", node.SessionCount})
    node.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", node.Interval})

    node.EntityData.YListKeys = []string {"NodeName"}

    return &(node.EntityData)
}

