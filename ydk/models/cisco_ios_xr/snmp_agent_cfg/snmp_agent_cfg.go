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

// SnmpTos represents Snmp tos
type SnmpTos string

const (
    // SNMP TOS type Precedence
    SnmpTos_precedence SnmpTos = "precedence"

    // SNMP TOS type DSCP
    SnmpTos_dscp SnmpTos = "dscp"
)

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

// SnmpOwnerAccess represents Snmp owner access
type SnmpOwnerAccess string

const (
    // Secure Domain Router Owner permissions
    SnmpOwnerAccess_sdr_owner SnmpOwnerAccess = "sdr-owner"

    // System owner permissions
    SnmpOwnerAccess_system_owner SnmpOwnerAccess = "system-owner"
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

// SnmpMibViewInclusion represents Snmp mib view inclusion
type SnmpMibViewInclusion string

const (
    // MIB View to be included
    SnmpMibViewInclusion_included SnmpMibViewInclusion = "included"

    // MIB View to be excluded
    SnmpMibViewInclusion_excluded SnmpMibViewInclusion = "excluded"
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

// SnmpAccessLevel represents Snmp access level
type SnmpAccessLevel string

const (
    // Read Only Access for a community string
    SnmpAccessLevel_read_only SnmpAccessLevel = "read-only"

    // Read Write Access for a community string
    SnmpAccessLevel_read_write SnmpAccessLevel = "read-write"
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
    // with pattern: b'[a-zA-Z0-9./-]+'.
    TrapSource interface{}

    // Disable authentication traps for packets on a vrf. The type is interface{}.
    VrfAuthenticationTrapDisable interface{}

    // Timeout value in seconds for Inform request (default 15 sec). The type is
    // interface{} with range: 1..42949671. Units are second.
    InformTimeout interface{}

    // Assign an interface for the source IPV6 address of all traps. The type is
    // string with pattern: b'[a-zA-Z0-9./-]+'.
    TrapSourceIpv6 interface{}

    // Largest SNMP packet size. The type is interface{} with range: 484..65500.
    PacketSize interface{}

    // Throttle time for incoming queue (default 0 msec). The type is interface{}
    // with range: 50..1000.
    ThrottleTime interface{}

    // Assign an interface for the source address of all traps. The type is string
    // with pattern: b'[a-zA-Z0-9./-]+'.
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

    snmp.EntityData.Children = make(map[string]types.YChild)
    snmp.EntityData.Children["encrypted-community-maps"] = types.YChild{"EncryptedCommunityMaps", &snmp.EncryptedCommunityMaps}
    snmp.EntityData.Children["views"] = types.YChild{"Views", &snmp.Views}
    snmp.EntityData.Children["logging"] = types.YChild{"Logging", &snmp.Logging}
    snmp.EntityData.Children["administration"] = types.YChild{"Administration", &snmp.Administration}
    snmp.EntityData.Children["agent"] = types.YChild{"Agent", &snmp.Agent}
    snmp.EntityData.Children["trap"] = types.YChild{"Trap", &snmp.Trap}
    snmp.EntityData.Children["drop-packet"] = types.YChild{"DropPacket", &snmp.DropPacket}
    snmp.EntityData.Children["ipv6"] = types.YChild{"Ipv6", &snmp.Ipv6}
    snmp.EntityData.Children["ipv4"] = types.YChild{"Ipv4", &snmp.Ipv4}
    snmp.EntityData.Children["system"] = types.YChild{"System", &snmp.System}
    snmp.EntityData.Children["target"] = types.YChild{"Target", &snmp.Target}
    snmp.EntityData.Children["notification"] = types.YChild{"Notification", &snmp.Notification}
    snmp.EntityData.Children["correlator"] = types.YChild{"Correlator", &snmp.Correlator}
    snmp.EntityData.Children["bulk-stats"] = types.YChild{"BulkStats", &snmp.BulkStats}
    snmp.EntityData.Children["default-community-maps"] = types.YChild{"DefaultCommunityMaps", &snmp.DefaultCommunityMaps}
    snmp.EntityData.Children["overload-control"] = types.YChild{"OverloadControl", &snmp.OverloadControl}
    snmp.EntityData.Children["timeouts"] = types.YChild{"Timeouts", &snmp.Timeouts}
    snmp.EntityData.Children["users"] = types.YChild{"Users", &snmp.Users}
    snmp.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &snmp.Vrfs}
    snmp.EntityData.Children["groups"] = types.YChild{"Groups", &snmp.Groups}
    snmp.EntityData.Children["trap-hosts"] = types.YChild{"TrapHosts", &snmp.TrapHosts}
    snmp.EntityData.Children["contexts"] = types.YChild{"Contexts", &snmp.Contexts}
    snmp.EntityData.Children["context-mappings"] = types.YChild{"ContextMappings", &snmp.ContextMappings}
    snmp.EntityData.Leafs = make(map[string]types.YLeaf)
    snmp.EntityData.Leafs["inform-retries"] = types.YLeaf{"InformRetries", snmp.InformRetries}
    snmp.EntityData.Leafs["trap-port"] = types.YLeaf{"TrapPort", snmp.TrapPort}
    snmp.EntityData.Leafs["oid-poll-stats"] = types.YLeaf{"OidPollStats", snmp.OidPollStats}
    snmp.EntityData.Leafs["trap-source"] = types.YLeaf{"TrapSource", snmp.TrapSource}
    snmp.EntityData.Leafs["vrf-authentication-trap-disable"] = types.YLeaf{"VrfAuthenticationTrapDisable", snmp.VrfAuthenticationTrapDisable}
    snmp.EntityData.Leafs["inform-timeout"] = types.YLeaf{"InformTimeout", snmp.InformTimeout}
    snmp.EntityData.Leafs["trap-source-ipv6"] = types.YLeaf{"TrapSourceIpv6", snmp.TrapSourceIpv6}
    snmp.EntityData.Leafs["packet-size"] = types.YLeaf{"PacketSize", snmp.PacketSize}
    snmp.EntityData.Leafs["throttle-time"] = types.YLeaf{"ThrottleTime", snmp.ThrottleTime}
    snmp.EntityData.Leafs["trap-source-ipv4"] = types.YLeaf{"TrapSourceIpv4", snmp.TrapSourceIpv4}
    snmp.EntityData.Leafs["inform-pending"] = types.YLeaf{"InformPending", snmp.InformPending}
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
    EncryptedCommunityMap []Snmp_EncryptedCommunityMaps_EncryptedCommunityMap
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

    encryptedCommunityMaps.EntityData.Children = make(map[string]types.YChild)
    encryptedCommunityMaps.EntityData.Children["encrypted-community-map"] = types.YChild{"EncryptedCommunityMap", nil}
    for i := range encryptedCommunityMaps.EncryptedCommunityMap {
        encryptedCommunityMaps.EntityData.Children[types.GetSegmentPath(&encryptedCommunityMaps.EncryptedCommunityMap[i])] = types.YChild{"EncryptedCommunityMap", &encryptedCommunityMaps.EncryptedCommunityMap[i]}
    }
    encryptedCommunityMaps.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(encryptedCommunityMaps.EntityData)
}

// Snmp_EncryptedCommunityMaps_EncryptedCommunityMap
// Clear/encrypted SNMP community map
type Snmp_EncryptedCommunityMaps_EncryptedCommunityMap struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SNMP community map. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    encryptedCommunityMap.EntityData.SegmentPath = "encrypted-community-map" + "[community-name='" + fmt.Sprintf("%v", encryptedCommunityMap.CommunityName) + "']"
    encryptedCommunityMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryptedCommunityMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryptedCommunityMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryptedCommunityMap.EntityData.Children = make(map[string]types.YChild)
    encryptedCommunityMap.EntityData.Leafs = make(map[string]types.YLeaf)
    encryptedCommunityMap.EntityData.Leafs["community-name"] = types.YLeaf{"CommunityName", encryptedCommunityMap.CommunityName}
    encryptedCommunityMap.EntityData.Leafs["context"] = types.YLeaf{"Context", encryptedCommunityMap.Context}
    encryptedCommunityMap.EntityData.Leafs["security"] = types.YLeaf{"Security", encryptedCommunityMap.Security}
    encryptedCommunityMap.EntityData.Leafs["target-list"] = types.YLeaf{"TargetList", encryptedCommunityMap.TargetList}
    return &(encryptedCommunityMap.EntityData)
}

// Snmp_Views
// Class to configure a SNMPv2 MIB view
type Snmp_Views struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the view. The type is slice of Snmp_Views_View.
    View []Snmp_Views_View
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

    views.EntityData.Children = make(map[string]types.YChild)
    views.EntityData.Children["view"] = types.YChild{"View", nil}
    for i := range views.View {
        views.EntityData.Children[types.GetSegmentPath(&views.View[i])] = types.YChild{"View", &views.View[i]}
    }
    views.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(views.EntityData)
}

// Snmp_Views_View
// Name of the view
type Snmp_Views_View struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the view. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    view.EntityData.SegmentPath = "view" + "[view-name='" + fmt.Sprintf("%v", view.ViewName) + "']" + "[family='" + fmt.Sprintf("%v", view.Family) + "']"
    view.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    view.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    view.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    view.EntityData.Children = make(map[string]types.YChild)
    view.EntityData.Leafs = make(map[string]types.YLeaf)
    view.EntityData.Leafs["view-name"] = types.YLeaf{"ViewName", view.ViewName}
    view.EntityData.Leafs["family"] = types.YLeaf{"Family", view.Family}
    view.EntityData.Leafs["view-inclusion"] = types.YLeaf{"ViewInclusion", view.ViewInclusion}
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

    logging.EntityData.Children = make(map[string]types.YChild)
    logging.EntityData.Children["threshold"] = types.YChild{"Threshold", &logging.Threshold}
    logging.EntityData.Leafs = make(map[string]types.YLeaf)
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

    threshold.EntityData.Children = make(map[string]types.YChild)
    threshold.EntityData.Leafs = make(map[string]types.YLeaf)
    threshold.EntityData.Leafs["oid-processing"] = types.YLeaf{"OidProcessing", threshold.OidProcessing}
    threshold.EntityData.Leafs["pdu-processing"] = types.YLeaf{"PduProcessing", threshold.PduProcessing}
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

    administration.EntityData.Children = make(map[string]types.YChild)
    administration.EntityData.Children["default-communities"] = types.YChild{"DefaultCommunities", &administration.DefaultCommunities}
    administration.EntityData.Children["encrypted-communities"] = types.YChild{"EncryptedCommunities", &administration.EncryptedCommunities}
    administration.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(administration.EntityData)
}

// Snmp_Administration_DefaultCommunities
// Container class to hold unencrpted communities
type Snmp_Administration_DefaultCommunities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unencrpted SNMP community string and access priviledges. The type is slice
    // of Snmp_Administration_DefaultCommunities_DefaultCommunity.
    DefaultCommunity []Snmp_Administration_DefaultCommunities_DefaultCommunity
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

    defaultCommunities.EntityData.Children = make(map[string]types.YChild)
    defaultCommunities.EntityData.Children["default-community"] = types.YChild{"DefaultCommunity", nil}
    for i := range defaultCommunities.DefaultCommunity {
        defaultCommunities.EntityData.Children[types.GetSegmentPath(&defaultCommunities.DefaultCommunity[i])] = types.YChild{"DefaultCommunity", &defaultCommunities.DefaultCommunity[i]}
    }
    defaultCommunities.EntityData.Leafs = make(map[string]types.YLeaf)
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

func (defaultCommunity *Snmp_Administration_DefaultCommunities_DefaultCommunity) GetEntityData() *types.CommonEntityData {
    defaultCommunity.EntityData.YFilter = defaultCommunity.YFilter
    defaultCommunity.EntityData.YangName = "default-community"
    defaultCommunity.EntityData.BundleName = "cisco_ios_xr"
    defaultCommunity.EntityData.ParentYangName = "default-communities"
    defaultCommunity.EntityData.SegmentPath = "default-community" + "[community-name='" + fmt.Sprintf("%v", defaultCommunity.CommunityName) + "']"
    defaultCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultCommunity.EntityData.Children = make(map[string]types.YChild)
    defaultCommunity.EntityData.Leafs = make(map[string]types.YLeaf)
    defaultCommunity.EntityData.Leafs["community-name"] = types.YLeaf{"CommunityName", defaultCommunity.CommunityName}
    defaultCommunity.EntityData.Leafs["priviledge"] = types.YLeaf{"Priviledge", defaultCommunity.Priviledge}
    defaultCommunity.EntityData.Leafs["view-name"] = types.YLeaf{"ViewName", defaultCommunity.ViewName}
    defaultCommunity.EntityData.Leafs["v4acl-type"] = types.YLeaf{"V4AclType", defaultCommunity.V4AclType}
    defaultCommunity.EntityData.Leafs["v4-access-list"] = types.YLeaf{"V4AccessList", defaultCommunity.V4AccessList}
    defaultCommunity.EntityData.Leafs["v6acl-type"] = types.YLeaf{"V6AclType", defaultCommunity.V6AclType}
    defaultCommunity.EntityData.Leafs["v6-access-list"] = types.YLeaf{"V6AccessList", defaultCommunity.V6AccessList}
    defaultCommunity.EntityData.Leafs["owner"] = types.YLeaf{"Owner", defaultCommunity.Owner}
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
    EncryptedCommunity []Snmp_Administration_EncryptedCommunities_EncryptedCommunity
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

    encryptedCommunities.EntityData.Children = make(map[string]types.YChild)
    encryptedCommunities.EntityData.Children["encrypted-community"] = types.YChild{"EncryptedCommunity", nil}
    for i := range encryptedCommunities.EncryptedCommunity {
        encryptedCommunities.EntityData.Children[types.GetSegmentPath(&encryptedCommunities.EncryptedCommunity[i])] = types.YChild{"EncryptedCommunity", &encryptedCommunities.EncryptedCommunity[i]}
    }
    encryptedCommunities.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(encryptedCommunities.EntityData)
}

// Snmp_Administration_EncryptedCommunities_EncryptedCommunity
// Clear/encrypted SNMP community string and
// access priviledges
type Snmp_Administration_EncryptedCommunities_EncryptedCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SNMP community string. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (encryptedCommunity *Snmp_Administration_EncryptedCommunities_EncryptedCommunity) GetEntityData() *types.CommonEntityData {
    encryptedCommunity.EntityData.YFilter = encryptedCommunity.YFilter
    encryptedCommunity.EntityData.YangName = "encrypted-community"
    encryptedCommunity.EntityData.BundleName = "cisco_ios_xr"
    encryptedCommunity.EntityData.ParentYangName = "encrypted-communities"
    encryptedCommunity.EntityData.SegmentPath = "encrypted-community" + "[community-name='" + fmt.Sprintf("%v", encryptedCommunity.CommunityName) + "']"
    encryptedCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryptedCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryptedCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryptedCommunity.EntityData.Children = make(map[string]types.YChild)
    encryptedCommunity.EntityData.Leafs = make(map[string]types.YLeaf)
    encryptedCommunity.EntityData.Leafs["community-name"] = types.YLeaf{"CommunityName", encryptedCommunity.CommunityName}
    encryptedCommunity.EntityData.Leafs["priviledge"] = types.YLeaf{"Priviledge", encryptedCommunity.Priviledge}
    encryptedCommunity.EntityData.Leafs["view-name"] = types.YLeaf{"ViewName", encryptedCommunity.ViewName}
    encryptedCommunity.EntityData.Leafs["v4acl-type"] = types.YLeaf{"V4AclType", encryptedCommunity.V4AclType}
    encryptedCommunity.EntityData.Leafs["v4-access-list"] = types.YLeaf{"V4AccessList", encryptedCommunity.V4AccessList}
    encryptedCommunity.EntityData.Leafs["v6acl-type"] = types.YLeaf{"V6AclType", encryptedCommunity.V6AclType}
    encryptedCommunity.EntityData.Leafs["v6-access-list"] = types.YLeaf{"V6AccessList", encryptedCommunity.V6AccessList}
    encryptedCommunity.EntityData.Leafs["owner"] = types.YLeaf{"Owner", encryptedCommunity.Owner}
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

    agent.EntityData.Children = make(map[string]types.YChild)
    agent.EntityData.Children["engine-id"] = types.YChild{"EngineId", &agent.EngineId}
    agent.EntityData.Leafs = make(map[string]types.YLeaf)
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

    engineId.EntityData.Children = make(map[string]types.YChild)
    engineId.EntityData.Children["remotes"] = types.YChild{"Remotes", &engineId.Remotes}
    engineId.EntityData.Leafs = make(map[string]types.YLeaf)
    engineId.EntityData.Leafs["local"] = types.YLeaf{"Local", engineId.Local}
    return &(engineId.EntityData)
}

// Snmp_Agent_EngineId_Remotes
// SNMPv3 remote SNMP Entity
type Snmp_Agent_EngineId_Remotes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // engineID of the remote agent. The type is slice of
    // Snmp_Agent_EngineId_Remotes_Remote.
    Remote []Snmp_Agent_EngineId_Remotes_Remote
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

    remotes.EntityData.Children = make(map[string]types.YChild)
    remotes.EntityData.Children["remote"] = types.YChild{"Remote", nil}
    for i := range remotes.Remote {
        remotes.EntityData.Children[types.GetSegmentPath(&remotes.Remote[i])] = types.YChild{"Remote", &remotes.Remote[i]}
    }
    remotes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(remotes.EntityData)
}

// Snmp_Agent_EngineId_Remotes_Remote
// engineID of the remote agent
type Snmp_Agent_EngineId_Remotes_Remote struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of remote SNMP entity. The type is one
    // of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    remote.EntityData.SegmentPath = "remote" + "[remote-address='" + fmt.Sprintf("%v", remote.RemoteAddress) + "']"
    remote.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remote.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remote.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remote.EntityData.Children = make(map[string]types.YChild)
    remote.EntityData.Leafs = make(map[string]types.YLeaf)
    remote.EntityData.Leafs["remote-address"] = types.YLeaf{"RemoteAddress", remote.RemoteAddress}
    remote.EntityData.Leafs["remote-engine-id"] = types.YLeaf{"RemoteEngineId", remote.RemoteEngineId}
    remote.EntityData.Leafs["port"] = types.YLeaf{"Port", remote.Port}
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

    trap.EntityData.Children = make(map[string]types.YChild)
    trap.EntityData.Leafs = make(map[string]types.YLeaf)
    trap.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", trap.Timeout}
    trap.EntityData.Leafs["throttle-time"] = types.YLeaf{"ThrottleTime", trap.ThrottleTime}
    trap.EntityData.Leafs["queue-length"] = types.YLeaf{"QueueLength", trap.QueueLength}
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

    dropPacket.EntityData.Children = make(map[string]types.YChild)
    dropPacket.EntityData.Leafs = make(map[string]types.YLeaf)
    dropPacket.EntityData.Leafs["unknown-user"] = types.YLeaf{"UnknownUser", dropPacket.UnknownUser}
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

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Children["tos"] = types.YChild{"Tos", &ipv6.Tos}
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6.EntityData)
}

// Snmp_Ipv6_Tos
// Type of TOS
type Snmp_Ipv6_Tos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP TOS type DSCP or Precedence. The type is SnmpTos.
    Type_ interface{}

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

    tos.EntityData.Children = make(map[string]types.YChild)
    tos.EntityData.Leafs = make(map[string]types.YLeaf)
    tos.EntityData.Leafs["type"] = types.YLeaf{"Type_", tos.Type_}
    tos.EntityData.Leafs["precedence"] = types.YLeaf{"Precedence", tos.Precedence}
    tos.EntityData.Leafs["dscp"] = types.YLeaf{"Dscp", tos.Dscp}
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

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Children["tos"] = types.YChild{"Tos", &ipv4.Tos}
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4.EntityData)
}

// Snmp_Ipv4_Tos
// Type of TOS
type Snmp_Ipv4_Tos struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP TOS type DSCP or Precedence. The type is SnmpTos.
    Type_ interface{}

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

    tos.EntityData.Children = make(map[string]types.YChild)
    tos.EntityData.Leafs = make(map[string]types.YLeaf)
    tos.EntityData.Leafs["type"] = types.YLeaf{"Type_", tos.Type_}
    tos.EntityData.Leafs["precedence"] = types.YLeaf{"Precedence", tos.Precedence}
    tos.EntityData.Leafs["dscp"] = types.YLeaf{"Dscp", tos.Dscp}
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

    system.EntityData.Children = make(map[string]types.YChild)
    system.EntityData.Leafs = make(map[string]types.YLeaf)
    system.EntityData.Leafs["chassis-id"] = types.YLeaf{"ChassisId", system.ChassisId}
    system.EntityData.Leafs["location"] = types.YLeaf{"Location", system.Location}
    system.EntityData.Leafs["contact"] = types.YLeaf{"Contact", system.Contact}
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

    target.EntityData.Children = make(map[string]types.YChild)
    target.EntityData.Children["targets"] = types.YChild{"Targets", &target.Targets}
    target.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(target.EntityData)
}

// Snmp_Target_Targets
// List of targets
type Snmp_Target_Targets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the target list. The type is slice of Snmp_Target_Targets_Target.
    Target []Snmp_Target_Targets_Target_
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

    targets.EntityData.Children = make(map[string]types.YChild)
    targets.EntityData.Children["target"] = types.YChild{"Target", nil}
    for i := range targets.Target {
        targets.EntityData.Children[types.GetSegmentPath(&targets.Target[i])] = types.YChild{"Target", &targets.Target[i]}
    }
    targets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(targets.EntityData)
}

// Snmp_Target_Targets_Target_
// Name of the target list
type Snmp_Target_Targets_Target_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the target list. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TargetListName interface{}

    // List of VRF Name for a target list.
    VrfNames Snmp_Target_Targets_Target__VrfNames

    // SNMP Target address configurations.
    TargetAddresses Snmp_Target_Targets_Target__TargetAddresses
}

func (target_ *Snmp_Target_Targets_Target_) GetEntityData() *types.CommonEntityData {
    target_.EntityData.YFilter = target_.YFilter
    target_.EntityData.YangName = "target"
    target_.EntityData.BundleName = "cisco_ios_xr"
    target_.EntityData.ParentYangName = "targets"
    target_.EntityData.SegmentPath = "target" + "[target-list-name='" + fmt.Sprintf("%v", target_.TargetListName) + "']"
    target_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    target_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    target_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    target_.EntityData.Children = make(map[string]types.YChild)
    target_.EntityData.Children["vrf-names"] = types.YChild{"VrfNames", &target_.VrfNames}
    target_.EntityData.Children["target-addresses"] = types.YChild{"TargetAddresses", &target_.TargetAddresses}
    target_.EntityData.Leafs = make(map[string]types.YLeaf)
    target_.EntityData.Leafs["target-list-name"] = types.YLeaf{"TargetListName", target_.TargetListName}
    return &(target_.EntityData)
}

// Snmp_Target_Targets_Target__VrfNames
// List of VRF Name for a target list
type Snmp_Target_Targets_Target__VrfNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name of the target. The type is slice of
    // Snmp_Target_Targets_Target__VrfNames_VrfName.
    VrfName []Snmp_Target_Targets_Target__VrfNames_VrfName
}

func (vrfNames *Snmp_Target_Targets_Target__VrfNames) GetEntityData() *types.CommonEntityData {
    vrfNames.EntityData.YFilter = vrfNames.YFilter
    vrfNames.EntityData.YangName = "vrf-names"
    vrfNames.EntityData.BundleName = "cisco_ios_xr"
    vrfNames.EntityData.ParentYangName = "target"
    vrfNames.EntityData.SegmentPath = "vrf-names"
    vrfNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfNames.EntityData.Children = make(map[string]types.YChild)
    vrfNames.EntityData.Children["vrf-name"] = types.YChild{"VrfName", nil}
    for i := range vrfNames.VrfName {
        vrfNames.EntityData.Children[types.GetSegmentPath(&vrfNames.VrfName[i])] = types.YChild{"VrfName", &vrfNames.VrfName[i]}
    }
    vrfNames.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfNames.EntityData)
}

// Snmp_Target_Targets_Target__VrfNames_VrfName
// VRF name of the target
type Snmp_Target_Targets_Target__VrfNames_VrfName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Name interface{}
}

func (vrfName *Snmp_Target_Targets_Target__VrfNames_VrfName) GetEntityData() *types.CommonEntityData {
    vrfName.EntityData.YFilter = vrfName.YFilter
    vrfName.EntityData.YangName = "vrf-name"
    vrfName.EntityData.BundleName = "cisco_ios_xr"
    vrfName.EntityData.ParentYangName = "vrf-names"
    vrfName.EntityData.SegmentPath = "vrf-name" + "[name='" + fmt.Sprintf("%v", vrfName.Name) + "']"
    vrfName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfName.EntityData.Children = make(map[string]types.YChild)
    vrfName.EntityData.Leafs = make(map[string]types.YLeaf)
    vrfName.EntityData.Leafs["name"] = types.YLeaf{"Name", vrfName.Name}
    return &(vrfName.EntityData)
}

// Snmp_Target_Targets_Target__TargetAddresses
// SNMP Target address configurations
type Snmp_Target_Targets_Target__TargetAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP Address to be configured for the Target. The type is slice of
    // Snmp_Target_Targets_Target__TargetAddresses_TargetAddress.
    TargetAddress []Snmp_Target_Targets_Target__TargetAddresses_TargetAddress
}

func (targetAddresses *Snmp_Target_Targets_Target__TargetAddresses) GetEntityData() *types.CommonEntityData {
    targetAddresses.EntityData.YFilter = targetAddresses.YFilter
    targetAddresses.EntityData.YangName = "target-addresses"
    targetAddresses.EntityData.BundleName = "cisco_ios_xr"
    targetAddresses.EntityData.ParentYangName = "target"
    targetAddresses.EntityData.SegmentPath = "target-addresses"
    targetAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    targetAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    targetAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    targetAddresses.EntityData.Children = make(map[string]types.YChild)
    targetAddresses.EntityData.Children["target-address"] = types.YChild{"TargetAddress", nil}
    for i := range targetAddresses.TargetAddress {
        targetAddresses.EntityData.Children[types.GetSegmentPath(&targetAddresses.TargetAddress[i])] = types.YChild{"TargetAddress", &targetAddresses.TargetAddress[i]}
    }
    targetAddresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(targetAddresses.EntityData)
}

// Snmp_Target_Targets_Target__TargetAddresses_TargetAddress
// IP Address to be configured for the Target
type Snmp_Target_Targets_Target__TargetAddresses_TargetAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4/Ipv6 address. The type is one of the
    // following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    IpAddress interface{}
}

func (targetAddress *Snmp_Target_Targets_Target__TargetAddresses_TargetAddress) GetEntityData() *types.CommonEntityData {
    targetAddress.EntityData.YFilter = targetAddress.YFilter
    targetAddress.EntityData.YangName = "target-address"
    targetAddress.EntityData.BundleName = "cisco_ios_xr"
    targetAddress.EntityData.ParentYangName = "target-addresses"
    targetAddress.EntityData.SegmentPath = "target-address" + "[ip-address='" + fmt.Sprintf("%v", targetAddress.IpAddress) + "']"
    targetAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    targetAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    targetAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    targetAddress.EntityData.Children = make(map[string]types.YChild)
    targetAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    targetAddress.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", targetAddress.IpAddress}
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
    Snmp Snmp_Notification_Snmp_

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
    FrequencySynchronization Snmp_Notification_FrequencySynchronization

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

    // Enable CISCO-ENTITY-EXT-MIB notifications.
    CiscoEntityExt Snmp_Notification_CiscoEntityExt

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

func (notification *Snmp_Notification) GetEntityData() *types.CommonEntityData {
    notification.EntityData.YFilter = notification.YFilter
    notification.EntityData.YangName = "notification"
    notification.EntityData.BundleName = "cisco_ios_xr"
    notification.EntityData.ParentYangName = "snmp"
    notification.EntityData.SegmentPath = "notification"
    notification.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    notification.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    notification.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    notification.EntityData.Children = make(map[string]types.YChild)
    notification.EntityData.Children["snmp"] = types.YChild{"Snmp", &notification.Snmp}
    notification.EntityData.Children["Cisco-IOS-XR-aaa-diameter-base-mib-cfg:diametermib"] = types.YChild{"Diametermib", &notification.Diametermib}
    notification.EntityData.Children["Cisco-IOS-XR-l2vpn-cfg:vpls"] = types.YChild{"Vpls", &notification.Vpls}
    notification.EntityData.Children["Cisco-IOS-XR-l2vpn-cfg:l2vpn"] = types.YChild{"L2Vpn", &notification.L2Vpn}
    notification.EntityData.Children["Cisco-IOS-XR-clns-isis-cfg:isis"] = types.YChild{"Isis", &notification.Isis}
    notification.EntityData.Children["Cisco-IOS-XR-config-mibs-cfg:config-man"] = types.YChild{"ConfigMan", &notification.ConfigMan}
    notification.EntityData.Children["Cisco-IOS-XR-ethernet-cfm-cfg:cfm"] = types.YChild{"Cfm", &notification.Cfm}
    notification.EntityData.Children["Cisco-IOS-XR-ethernet-link-oam-cfg:oam"] = types.YChild{"Oam", &notification.Oam}
    notification.EntityData.Children["Cisco-IOS-XR-fabhfr-mib-cfg:fabric-crs"] = types.YChild{"FabricCrs", &notification.FabricCrs}
    notification.EntityData.Children["Cisco-IOS-XR-flashmib-cfg:flash"] = types.YChild{"Flash", &notification.Flash}
    notification.EntityData.Children["Cisco-IOS-XR-freqsync-cfg:frequency-synchronization"] = types.YChild{"FrequencySynchronization", &notification.FrequencySynchronization}
    notification.EntityData.Children["Cisco-IOS-XR-infra-ceredundancymib-cfg:entity-redundancy"] = types.YChild{"EntityRedundancy", &notification.EntityRedundancy}
    notification.EntityData.Children["Cisco-IOS-XR-infra-confcopymib-cfg:config-copy"] = types.YChild{"ConfigCopy", &notification.ConfigCopy}
    notification.EntityData.Children["Cisco-IOS-XR-infra-rsi-cfg:selective-vrf-download"] = types.YChild{"SelectiveVrfDownload", &notification.SelectiveVrfDownload}
    notification.EntityData.Children["Cisco-IOS-XR-infra-systemmib-cfg:system"] = types.YChild{"System", &notification.System}
    notification.EntityData.Children["Cisco-IOS-XR-ip-bfd-cfg:bfd"] = types.YChild{"Bfd", &notification.Bfd}
    notification.EntityData.Children["Cisco-IOS-XR-ip-daps-mib-cfg:addresspool-mib"] = types.YChild{"AddresspoolMib", &notification.AddresspoolMib}
    notification.EntityData.Children["Cisco-IOS-XR-ip-ntp-cfg:ntp"] = types.YChild{"Ntp", &notification.Ntp}
    notification.EntityData.Children["Cisco-IOS-XR-ip-rsvp-cfg:rsvp"] = types.YChild{"Rsvp", &notification.Rsvp}
    notification.EntityData.Children["Cisco-IOS-XR-ipv4-bgp-cfg:bgp"] = types.YChild{"Bgp", &notification.Bgp}
    notification.EntityData.Children["Cisco-IOS-XR-ipv4-hsrp-cfg:hsrp"] = types.YChild{"Hsrp", &notification.Hsrp}
    notification.EntityData.Children["Cisco-IOS-XR-ipv4-ospf-cfg:ospf"] = types.YChild{"Ospf", &notification.Ospf}
    notification.EntityData.Children["Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp"] = types.YChild{"Vrrp", &notification.Vrrp}
    notification.EntityData.Children["Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3"] = types.YChild{"Ospfv3", &notification.Ospfv3}
    notification.EntityData.Children["Cisco-IOS-XR-mpls-ldp-cfg:mpls-ldp"] = types.YChild{"MplsLdp", &notification.MplsLdp}
    notification.EntityData.Children["Cisco-IOS-XR-mpls-te-cfg:mpls-te-p2mp"] = types.YChild{"MplsTeP2Mp", &notification.MplsTeP2Mp}
    notification.EntityData.Children["Cisco-IOS-XR-mpls-te-cfg:mpls-te"] = types.YChild{"MplsTe", &notification.MplsTe}
    notification.EntityData.Children["Cisco-IOS-XR-mpls-te-cfg:mpls-frr"] = types.YChild{"MplsFrr", &notification.MplsFrr}
    notification.EntityData.Children["Cisco-IOS-XR-mpls-vpn-cfg:mpls-l3vpn"] = types.YChild{"MplsL3Vpn", &notification.MplsL3Vpn}
    notification.EntityData.Children["Cisco-IOS-XR-opticalmib-cfg:optical"] = types.YChild{"Optical", &notification.Optical}
    notification.EntityData.Children["Cisco-IOS-XR-opticalotsmib-cfg:optical-ots"] = types.YChild{"OpticalOts", &notification.OpticalOts}
    notification.EntityData.Children["Cisco-IOS-XR-otnifmib-cfg:otn"] = types.YChild{"Otn", &notification.Otn}
    notification.EntityData.Children["Cisco-IOS-XR-snmp-bridgemib-cfg:bridge"] = types.YChild{"Bridge", &notification.Bridge}
    notification.EntityData.Children["Cisco-IOS-XR-snmp-ciscosensormib-cfg:sensor"] = types.YChild{"Sensor", &notification.Sensor}
    notification.EntityData.Children["Cisco-IOS-XR-snmp-entityextmib-cfg:cisco-entity-ext"] = types.YChild{"CiscoEntityExt", &notification.CiscoEntityExt}
    notification.EntityData.Children["Cisco-IOS-XR-snmp-entitymib-cfg:entity"] = types.YChild{"Entity", &notification.Entity}
    notification.EntityData.Children["Cisco-IOS-XR-snmp-entstatemib-cfg:entity-state"] = types.YChild{"EntityState", &notification.EntityState}
    notification.EntityData.Children["Cisco-IOS-XR-snmp-frucontrolmib-cfg:fru-control"] = types.YChild{"FruControl", &notification.FruControl}
    notification.EntityData.Children["Cisco-IOS-XR-snmp-mib-rfmib-cfg:rf"] = types.YChild{"Rf", &notification.Rf}
    notification.EntityData.Children["Cisco-IOS-XR-snmp-syslogmib-cfg:syslog"] = types.YChild{"Syslog", &notification.Syslog}
    notification.EntityData.Children["Cisco-IOS-XR-subscriber-session-mon-mibs-cfg:subscriber-mib"] = types.YChild{"SubscriberMib", &notification.SubscriberMib}
    notification.EntityData.Children["Cisco-IOS-XR-tunnel-l2tun-proto-mibs-cfg:l2tun"] = types.YChild{"L2Tun", &notification.L2Tun}
    notification.EntityData.Leafs = make(map[string]types.YLeaf)
    notification.EntityData.Leafs["ipsla"] = types.YLeaf{"Ipsla", notification.Ipsla}
    return &(notification.EntityData)
}

// Snmp_Notification_Snmp_
// SNMP notification configuration
type Snmp_Notification_Snmp_ struct {
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

func (snmp_ *Snmp_Notification_Snmp_) GetEntityData() *types.CommonEntityData {
    snmp_.EntityData.YFilter = snmp_.YFilter
    snmp_.EntityData.YangName = "snmp"
    snmp_.EntityData.BundleName = "cisco_ios_xr"
    snmp_.EntityData.ParentYangName = "notification"
    snmp_.EntityData.SegmentPath = "snmp"
    snmp_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmp_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmp_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmp_.EntityData.Children = make(map[string]types.YChild)
    snmp_.EntityData.Leafs = make(map[string]types.YLeaf)
    snmp_.EntityData.Leafs["authentication"] = types.YLeaf{"Authentication", snmp_.Authentication}
    snmp_.EntityData.Leafs["cold-start"] = types.YLeaf{"ColdStart", snmp_.ColdStart}
    snmp_.EntityData.Leafs["warm-start"] = types.YLeaf{"WarmStart", snmp_.WarmStart}
    snmp_.EntityData.Leafs["enable"] = types.YLeaf{"Enable", snmp_.Enable}
    snmp_.EntityData.Leafs["link-down"] = types.YLeaf{"LinkDown", snmp_.LinkDown}
    snmp_.EntityData.Leafs["link-up"] = types.YLeaf{"LinkUp", snmp_.LinkUp}
    return &(snmp_.EntityData)
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

    diametermib.EntityData.Children = make(map[string]types.YChild)
    diametermib.EntityData.Leafs = make(map[string]types.YLeaf)
    diametermib.EntityData.Leafs["protocolerror"] = types.YLeaf{"Protocolerror", diametermib.Protocolerror}
    diametermib.EntityData.Leafs["permanentfail"] = types.YLeaf{"Permanentfail", diametermib.Permanentfail}
    diametermib.EntityData.Leafs["peerdown"] = types.YLeaf{"Peerdown", diametermib.Peerdown}
    diametermib.EntityData.Leafs["peerup"] = types.YLeaf{"Peerup", diametermib.Peerup}
    diametermib.EntityData.Leafs["transientfail"] = types.YLeaf{"Transientfail", diametermib.Transientfail}
    return &(diametermib.EntityData)
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

    vpls.EntityData.Children = make(map[string]types.YChild)
    vpls.EntityData.Leafs = make(map[string]types.YLeaf)
    vpls.EntityData.Leafs["full-clear"] = types.YLeaf{"FullClear", vpls.FullClear}
    vpls.EntityData.Leafs["status"] = types.YLeaf{"Status", vpls.Status}
    vpls.EntityData.Leafs["enable"] = types.YLeaf{"Enable", vpls.Enable}
    vpls.EntityData.Leafs["full-raise"] = types.YLeaf{"FullRaise", vpls.FullRaise}
    return &(vpls.EntityData)
}

// Snmp_Notification_L2Vpn
// CISCO-IETF-PW-MIB notification configuration
type Snmp_Notification_L2Vpn struct {
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

func (l2Vpn *Snmp_Notification_L2Vpn) GetEntityData() *types.CommonEntityData {
    l2Vpn.EntityData.YFilter = l2Vpn.YFilter
    l2Vpn.EntityData.YangName = "l2vpn"
    l2Vpn.EntityData.BundleName = "cisco_ios_xr"
    l2Vpn.EntityData.ParentYangName = "notification"
    l2Vpn.EntityData.SegmentPath = "Cisco-IOS-XR-l2vpn-cfg:l2vpn"
    l2Vpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2Vpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2Vpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2Vpn.EntityData.Children = make(map[string]types.YChild)
    l2Vpn.EntityData.Leafs = make(map[string]types.YLeaf)
    l2Vpn.EntityData.Leafs["cisco"] = types.YLeaf{"Cisco", l2Vpn.Cisco}
    l2Vpn.EntityData.Leafs["enable"] = types.YLeaf{"Enable", l2Vpn.Enable}
    l2Vpn.EntityData.Leafs["vc-down"] = types.YLeaf{"VcDown", l2Vpn.VcDown}
    l2Vpn.EntityData.Leafs["vc-up"] = types.YLeaf{"VcUp", l2Vpn.VcUp}
    return &(l2Vpn.EntityData)
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

    isis.EntityData.Children = make(map[string]types.YChild)
    isis.EntityData.Leafs = make(map[string]types.YLeaf)
    isis.EntityData.Leafs["database-overflow"] = types.YLeaf{"DatabaseOverflow", isis.DatabaseOverflow}
    isis.EntityData.Leafs["manual-address-drops"] = types.YLeaf{"ManualAddressDrops", isis.ManualAddressDrops}
    isis.EntityData.Leafs["corrupted-lsp-detected"] = types.YLeaf{"CorruptedLspDetected", isis.CorruptedLspDetected}
    isis.EntityData.Leafs["attempt-to-exceed-max-sequence"] = types.YLeaf{"AttemptToExceedMaxSequence", isis.AttemptToExceedMaxSequence}
    isis.EntityData.Leafs["id-length-mismatch"] = types.YLeaf{"IdLengthMismatch", isis.IdLengthMismatch}
    isis.EntityData.Leafs["max-area-address-mismatch"] = types.YLeaf{"MaxAreaAddressMismatch", isis.MaxAreaAddressMismatch}
    isis.EntityData.Leafs["own-lsp-purge"] = types.YLeaf{"OwnLspPurge", isis.OwnLspPurge}
    isis.EntityData.Leafs["sequence-number-skip"] = types.YLeaf{"SequenceNumberSkip", isis.SequenceNumberSkip}
    isis.EntityData.Leafs["authentication-type-failure"] = types.YLeaf{"AuthenticationTypeFailure", isis.AuthenticationTypeFailure}
    isis.EntityData.Leafs["authentication-failure"] = types.YLeaf{"AuthenticationFailure", isis.AuthenticationFailure}
    isis.EntityData.Leafs["version-skew"] = types.YLeaf{"VersionSkew", isis.VersionSkew}
    isis.EntityData.Leafs["area-mismatch"] = types.YLeaf{"AreaMismatch", isis.AreaMismatch}
    isis.EntityData.Leafs["rejected-adjacency"] = types.YLeaf{"RejectedAdjacency", isis.RejectedAdjacency}
    isis.EntityData.Leafs["lsp-too-large-to-propagate"] = types.YLeaf{"LspTooLargeToPropagate", isis.LspTooLargeToPropagate}
    isis.EntityData.Leafs["originated-lsp-buffer-size-mismatch"] = types.YLeaf{"OriginatedLspBufferSizeMismatch", isis.OriginatedLspBufferSizeMismatch}
    isis.EntityData.Leafs["protocols-supported-mismatch"] = types.YLeaf{"ProtocolsSupportedMismatch", isis.ProtocolsSupportedMismatch}
    isis.EntityData.Leafs["adjacency-change"] = types.YLeaf{"AdjacencyChange", isis.AdjacencyChange}
    isis.EntityData.Leafs["lsp-error-detected"] = types.YLeaf{"LspErrorDetected", isis.LspErrorDetected}
    isis.EntityData.Leafs["all"] = types.YLeaf{"All", isis.All}
    return &(isis.EntityData)
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

    configMan.EntityData.Children = make(map[string]types.YChild)
    configMan.EntityData.Leafs = make(map[string]types.YLeaf)
    configMan.EntityData.Leafs["enable"] = types.YLeaf{"Enable", configMan.Enable}
    return &(configMan.EntityData)
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

    cfm.EntityData.Children = make(map[string]types.YChild)
    cfm.EntityData.Leafs = make(map[string]types.YLeaf)
    cfm.EntityData.Leafs["enable"] = types.YLeaf{"Enable", cfm.Enable}
    return &(cfm.EntityData)
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

    oam.EntityData.Children = make(map[string]types.YChild)
    oam.EntityData.Leafs = make(map[string]types.YLeaf)
    oam.EntityData.Leafs["enable"] = types.YLeaf{"Enable", oam.Enable}
    return &(oam.EntityData)
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

    fabricCrs.EntityData.Children = make(map[string]types.YChild)
    fabricCrs.EntityData.Leafs = make(map[string]types.YLeaf)
    fabricCrs.EntityData.Leafs["bundle-state"] = types.YLeaf{"BundleState", fabricCrs.BundleState}
    fabricCrs.EntityData.Leafs["plane-state"] = types.YLeaf{"PlaneState", fabricCrs.PlaneState}
    fabricCrs.EntityData.Leafs["bundle-downed-link"] = types.YLeaf{"BundleDownedLink", fabricCrs.BundleDownedLink}
    return &(fabricCrs.EntityData)
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

    flash.EntityData.Children = make(map[string]types.YChild)
    flash.EntityData.Leafs = make(map[string]types.YLeaf)
    flash.EntityData.Leafs["insertion"] = types.YLeaf{"Insertion", flash.Insertion}
    flash.EntityData.Leafs["removal"] = types.YLeaf{"Removal", flash.Removal}
    return &(flash.EntityData)
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

    frequencySynchronization.EntityData.Children = make(map[string]types.YChild)
    frequencySynchronization.EntityData.Leafs = make(map[string]types.YLeaf)
    frequencySynchronization.EntityData.Leafs["enable"] = types.YLeaf{"Enable", frequencySynchronization.Enable}
    return &(frequencySynchronization.EntityData)
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

    entityRedundancy.EntityData.Children = make(map[string]types.YChild)
    entityRedundancy.EntityData.Leafs = make(map[string]types.YLeaf)
    entityRedundancy.EntityData.Leafs["switchover"] = types.YLeaf{"Switchover", entityRedundancy.Switchover}
    entityRedundancy.EntityData.Leafs["enable"] = types.YLeaf{"Enable", entityRedundancy.Enable}
    entityRedundancy.EntityData.Leafs["status"] = types.YLeaf{"Status", entityRedundancy.Status}
    return &(entityRedundancy.EntityData)
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

    configCopy.EntityData.Children = make(map[string]types.YChild)
    configCopy.EntityData.Leafs = make(map[string]types.YLeaf)
    configCopy.EntityData.Leafs["completion"] = types.YLeaf{"Completion", configCopy.Completion}
    return &(configCopy.EntityData)
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

    selectiveVrfDownload.EntityData.Children = make(map[string]types.YChild)
    selectiveVrfDownload.EntityData.Leafs = make(map[string]types.YLeaf)
    selectiveVrfDownload.EntityData.Leafs["role-change"] = types.YLeaf{"RoleChange", selectiveVrfDownload.RoleChange}
    return &(selectiveVrfDownload.EntityData)
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

    system.EntityData.Children = make(map[string]types.YChild)
    system.EntityData.Leafs = make(map[string]types.YLeaf)
    system.EntityData.Leafs["enable"] = types.YLeaf{"Enable", system.Enable}
    return &(system.EntityData)
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

    bfd.EntityData.Children = make(map[string]types.YChild)
    bfd.EntityData.Leafs = make(map[string]types.YLeaf)
    bfd.EntityData.Leafs["enable"] = types.YLeaf{"Enable", bfd.Enable}
    return &(bfd.EntityData)
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

    addresspoolMib.EntityData.Children = make(map[string]types.YChild)
    addresspoolMib.EntityData.Leafs = make(map[string]types.YLeaf)
    addresspoolMib.EntityData.Leafs["high"] = types.YLeaf{"High", addresspoolMib.High}
    addresspoolMib.EntityData.Leafs["low"] = types.YLeaf{"Low", addresspoolMib.Low}
    return &(addresspoolMib.EntityData)
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

    ntp.EntityData.Children = make(map[string]types.YChild)
    ntp.EntityData.Leafs = make(map[string]types.YLeaf)
    ntp.EntityData.Leafs["enable"] = types.YLeaf{"Enable", ntp.Enable}
    return &(ntp.EntityData)
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

    rsvp.EntityData.Children = make(map[string]types.YChild)
    rsvp.EntityData.Leafs = make(map[string]types.YLeaf)
    rsvp.EntityData.Leafs["lost-flow"] = types.YLeaf{"LostFlow", rsvp.LostFlow}
    rsvp.EntityData.Leafs["new-flow"] = types.YLeaf{"NewFlow", rsvp.NewFlow}
    rsvp.EntityData.Leafs["enable"] = types.YLeaf{"Enable", rsvp.Enable}
    return &(rsvp.EntityData)
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
    Bgp4Mib Snmp_Notification_Bgp_Bgp4Mib

    // Enable CISCO-BGP4-MIB v2 notifications: cbgpPeer2EstablishedNotification,
    // cbgpPeer2BackwardTransNotification, cbgpPeer2FsmStateChange,
    // cbgpPeer2BackwardTransition, cbgpPeer2PrefixThresholdExceeded,
    // cbgpPeer2PrefixThresholdClear.
    CiscoBgp4Mib Snmp_Notification_Bgp_CiscoBgp4Mib
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

    bgp.EntityData.Children = make(map[string]types.YChild)
    bgp.EntityData.Children["bgp4mib"] = types.YChild{"Bgp4Mib", &bgp.Bgp4Mib}
    bgp.EntityData.Children["cisco-bgp4mib"] = types.YChild{"CiscoBgp4Mib", &bgp.CiscoBgp4Mib}
    bgp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(bgp.EntityData)
}

// Snmp_Notification_Bgp_Bgp4Mib
// Enable BGP4-MIB and CISCO-BGP4-MIB IPv4-only
// notifications: bgpEstablishedNotification,
// bgpBackwardTransNotification,
// cbgpFsmStateChange, cbgpBackwardTransition,
// cbgpPrefixThresholdExceeded,
// cbgpPrefixThresholdClear.
type Snmp_Notification_Bgp_Bgp4Mib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable BGP4-MIB and CISCO-BGP4-MIB IPv4-only notifications. The type is
    // interface{}.
    Enable interface{}

    // Enable BGP4-MIB and CISCO-BGP4-MIB IPv4-only up/down notifications. The
    // type is interface{}.
    UpDown interface{}
}

func (bgp4Mib *Snmp_Notification_Bgp_Bgp4Mib) GetEntityData() *types.CommonEntityData {
    bgp4Mib.EntityData.YFilter = bgp4Mib.YFilter
    bgp4Mib.EntityData.YangName = "bgp4mib"
    bgp4Mib.EntityData.BundleName = "cisco_ios_xr"
    bgp4Mib.EntityData.ParentYangName = "bgp"
    bgp4Mib.EntityData.SegmentPath = "bgp4mib"
    bgp4Mib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp4Mib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp4Mib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp4Mib.EntityData.Children = make(map[string]types.YChild)
    bgp4Mib.EntityData.Leafs = make(map[string]types.YLeaf)
    bgp4Mib.EntityData.Leafs["enable"] = types.YLeaf{"Enable", bgp4Mib.Enable}
    bgp4Mib.EntityData.Leafs["up-down"] = types.YLeaf{"UpDown", bgp4Mib.UpDown}
    return &(bgp4Mib.EntityData)
}

// Snmp_Notification_Bgp_CiscoBgp4Mib
// Enable CISCO-BGP4-MIB v2 notifications:
// cbgpPeer2EstablishedNotification,
// cbgpPeer2BackwardTransNotification,
// cbgpPeer2FsmStateChange,
// cbgpPeer2BackwardTransition,
// cbgpPeer2PrefixThresholdExceeded,
// cbgpPeer2PrefixThresholdClear.
type Snmp_Notification_Bgp_CiscoBgp4Mib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable CISCO-BGP4-MIB v2 notifications. The type is interface{}.
    Enable interface{}

    // Enable CISCO-BGP4-MIB v2 up/down notifications. The type is interface{}.
    UpDown interface{}
}

func (ciscoBgp4Mib *Snmp_Notification_Bgp_CiscoBgp4Mib) GetEntityData() *types.CommonEntityData {
    ciscoBgp4Mib.EntityData.YFilter = ciscoBgp4Mib.YFilter
    ciscoBgp4Mib.EntityData.YangName = "cisco-bgp4mib"
    ciscoBgp4Mib.EntityData.BundleName = "cisco_ios_xr"
    ciscoBgp4Mib.EntityData.ParentYangName = "bgp"
    ciscoBgp4Mib.EntityData.SegmentPath = "cisco-bgp4mib"
    ciscoBgp4Mib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ciscoBgp4Mib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ciscoBgp4Mib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ciscoBgp4Mib.EntityData.Children = make(map[string]types.YChild)
    ciscoBgp4Mib.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoBgp4Mib.EntityData.Leafs["enable"] = types.YLeaf{"Enable", ciscoBgp4Mib.Enable}
    ciscoBgp4Mib.EntityData.Leafs["up-down"] = types.YLeaf{"UpDown", ciscoBgp4Mib.UpDown}
    return &(ciscoBgp4Mib.EntityData)
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

    hsrp.EntityData.Children = make(map[string]types.YChild)
    hsrp.EntityData.Leafs = make(map[string]types.YLeaf)
    hsrp.EntityData.Leafs["enable"] = types.YLeaf{"Enable", hsrp.Enable}
    return &(hsrp.EntityData)
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

    ospf.EntityData.Children = make(map[string]types.YChild)
    ospf.EntityData.Children["lsa"] = types.YChild{"Lsa", &ospf.Lsa}
    ospf.EntityData.Children["state-change"] = types.YChild{"StateChange", &ospf.StateChange}
    ospf.EntityData.Children["retransmit"] = types.YChild{"Retransmit", &ospf.Retransmit}
    ospf.EntityData.Children["error"] = types.YChild{"Error", &ospf.Error}
    ospf.EntityData.Leafs = make(map[string]types.YLeaf)
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

    lsa.EntityData.Children = make(map[string]types.YChild)
    lsa.EntityData.Leafs = make(map[string]types.YLeaf)
    lsa.EntityData.Leafs["max-age-lsa"] = types.YLeaf{"MaxAgeLsa", lsa.MaxAgeLsa}
    lsa.EntityData.Leafs["originate-lsa"] = types.YLeaf{"OriginateLsa", lsa.OriginateLsa}
    return &(lsa.EntityData)
}

// Snmp_Notification_Ospf_StateChange
// SNMP notifications for OSPF state change
type Snmp_Notification_Ospf_StateChange struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ospfIfStateChange notification. The type is interface{}.
    Interface_ interface{}

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

    stateChange.EntityData.Children = make(map[string]types.YChild)
    stateChange.EntityData.Leafs = make(map[string]types.YLeaf)
    stateChange.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", stateChange.Interface_}
    stateChange.EntityData.Leafs["virtual-interface"] = types.YLeaf{"VirtualInterface", stateChange.VirtualInterface}
    stateChange.EntityData.Leafs["virtual-neighbor"] = types.YLeaf{"VirtualNeighbor", stateChange.VirtualNeighbor}
    stateChange.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", stateChange.Neighbor}
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

    retransmit.EntityData.Children = make(map[string]types.YChild)
    retransmit.EntityData.Leafs = make(map[string]types.YLeaf)
    retransmit.EntityData.Leafs["virtual-packet"] = types.YLeaf{"VirtualPacket", retransmit.VirtualPacket}
    retransmit.EntityData.Leafs["packet"] = types.YLeaf{"Packet", retransmit.Packet}
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

func (error *Snmp_Notification_Ospf_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "ospf"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["config-error"] = types.YLeaf{"ConfigError", error.ConfigError}
    error.EntityData.Leafs["authentication-failure"] = types.YLeaf{"AuthenticationFailure", error.AuthenticationFailure}
    error.EntityData.Leafs["virtual-config-error"] = types.YLeaf{"VirtualConfigError", error.VirtualConfigError}
    error.EntityData.Leafs["virtual-authentication-failure"] = types.YLeaf{"VirtualAuthenticationFailure", error.VirtualAuthenticationFailure}
    error.EntityData.Leafs["bad-packet"] = types.YLeaf{"BadPacket", error.BadPacket}
    error.EntityData.Leafs["virtual-bad-packet"] = types.YLeaf{"VirtualBadPacket", error.VirtualBadPacket}
    return &(error.EntityData)
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

    vrrp.EntityData.Children = make(map[string]types.YChild)
    vrrp.EntityData.Leafs = make(map[string]types.YLeaf)
    vrrp.EntityData.Leafs["enable"] = types.YLeaf{"Enable", vrrp.Enable}
    return &(vrrp.EntityData)
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

    ospfv3.EntityData.Children = make(map[string]types.YChild)
    ospfv3.EntityData.Children["error"] = types.YChild{"Error", &ospfv3.Error}
    ospfv3.EntityData.Children["state-change"] = types.YChild{"StateChange", &ospfv3.StateChange}
    ospfv3.EntityData.Leafs = make(map[string]types.YLeaf)
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

func (error *Snmp_Notification_Ospfv3_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "ospfv3"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["config-error"] = types.YLeaf{"ConfigError", error.ConfigError}
    error.EntityData.Leafs["bad-packet"] = types.YLeaf{"BadPacket", error.BadPacket}
    error.EntityData.Leafs["virtual-bad-packet"] = types.YLeaf{"VirtualBadPacket", error.VirtualBadPacket}
    error.EntityData.Leafs["virtual-config-error"] = types.YLeaf{"VirtualConfigError", error.VirtualConfigError}
    return &(error.EntityData)
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
    Interface_ interface{}

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

    stateChange.EntityData.Children = make(map[string]types.YChild)
    stateChange.EntityData.Leafs = make(map[string]types.YLeaf)
    stateChange.EntityData.Leafs["restart-virtual-helper"] = types.YLeaf{"RestartVirtualHelper", stateChange.RestartVirtualHelper}
    stateChange.EntityData.Leafs["nssa-translator"] = types.YLeaf{"NssaTranslator", stateChange.NssaTranslator}
    stateChange.EntityData.Leafs["interface"] = types.YLeaf{"Interface_", stateChange.Interface_}
    stateChange.EntityData.Leafs["restart"] = types.YLeaf{"Restart", stateChange.Restart}
    stateChange.EntityData.Leafs["neighbor"] = types.YLeaf{"Neighbor", stateChange.Neighbor}
    stateChange.EntityData.Leafs["virtual-interface"] = types.YLeaf{"VirtualInterface", stateChange.VirtualInterface}
    stateChange.EntityData.Leafs["restart-helper"] = types.YLeaf{"RestartHelper", stateChange.RestartHelper}
    stateChange.EntityData.Leafs["virtual-neighbor"] = types.YLeaf{"VirtualNeighbor", stateChange.VirtualNeighbor}
    return &(stateChange.EntityData)
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

    mplsLdp.EntityData.Children = make(map[string]types.YChild)
    mplsLdp.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsLdp.EntityData.Leafs["session-up"] = types.YLeaf{"SessionUp", mplsLdp.SessionUp}
    mplsLdp.EntityData.Leafs["init-session-threshold-exceeded"] = types.YLeaf{"InitSessionThresholdExceeded", mplsLdp.InitSessionThresholdExceeded}
    mplsLdp.EntityData.Leafs["session-down"] = types.YLeaf{"SessionDown", mplsLdp.SessionDown}
    return &(mplsLdp.EntityData)
}

// Snmp_Notification_MplsTeP2Mp
// CISCO-MPLS-TE-P2MP-STD-MIB notification
// configuration
type Snmp_Notification_MplsTeP2Mp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable cmplsTeP2mpTunnelDestUp notification. The type is interface{}.
    Up interface{}

    // Enable cmplsTeP2mpTunnelDestDown notification. The type is interface{}.
    Down interface{}
}

func (mplsTeP2Mp *Snmp_Notification_MplsTeP2Mp) GetEntityData() *types.CommonEntityData {
    mplsTeP2Mp.EntityData.YFilter = mplsTeP2Mp.YFilter
    mplsTeP2Mp.EntityData.YangName = "mpls-te-p2mp"
    mplsTeP2Mp.EntityData.BundleName = "cisco_ios_xr"
    mplsTeP2Mp.EntityData.ParentYangName = "notification"
    mplsTeP2Mp.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-te-cfg:mpls-te-p2mp"
    mplsTeP2Mp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsTeP2Mp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsTeP2Mp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsTeP2Mp.EntityData.Children = make(map[string]types.YChild)
    mplsTeP2Mp.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsTeP2Mp.EntityData.Leafs["up"] = types.YLeaf{"Up", mplsTeP2Mp.Up}
    mplsTeP2Mp.EntityData.Leafs["down"] = types.YLeaf{"Down", mplsTeP2Mp.Down}
    return &(mplsTeP2Mp.EntityData)
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

    mplsTe.EntityData.Children = make(map[string]types.YChild)
    mplsTe.EntityData.Children["cisco-extension"] = types.YChild{"CiscoExtension", &mplsTe.CiscoExtension}
    mplsTe.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsTe.EntityData.Leafs["cisco"] = types.YLeaf{"Cisco", mplsTe.Cisco}
    mplsTe.EntityData.Leafs["up"] = types.YLeaf{"Up", mplsTe.Up}
    mplsTe.EntityData.Leafs["reoptimize"] = types.YLeaf{"Reoptimize", mplsTe.Reoptimize}
    mplsTe.EntityData.Leafs["reroute"] = types.YLeaf{"Reroute", mplsTe.Reroute}
    mplsTe.EntityData.Leafs["down"] = types.YLeaf{"Down", mplsTe.Down}
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

    ciscoExtension.EntityData.Children = make(map[string]types.YChild)
    ciscoExtension.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoExtension.EntityData.Leafs["preempt"] = types.YLeaf{"Preempt", ciscoExtension.Preempt}
    ciscoExtension.EntityData.Leafs["insufficient-bandwidth"] = types.YLeaf{"InsufficientBandwidth", ciscoExtension.InsufficientBandwidth}
    ciscoExtension.EntityData.Leafs["re-route-pending-clear"] = types.YLeaf{"ReRoutePendingClear", ciscoExtension.ReRoutePendingClear}
    ciscoExtension.EntityData.Leafs["bringup-fail"] = types.YLeaf{"BringupFail", ciscoExtension.BringupFail}
    ciscoExtension.EntityData.Leafs["re-route-pending"] = types.YLeaf{"ReRoutePending", ciscoExtension.ReRoutePending}
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

    mplsFrr.EntityData.Children = make(map[string]types.YChild)
    mplsFrr.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsFrr.EntityData.Leafs["unprotected"] = types.YLeaf{"Unprotected", mplsFrr.Unprotected}
    mplsFrr.EntityData.Leafs["enable"] = types.YLeaf{"Enable", mplsFrr.Enable}
    mplsFrr.EntityData.Leafs["protected"] = types.YLeaf{"Protected", mplsFrr.Protected}
    return &(mplsFrr.EntityData)
}

// Snmp_Notification_MplsL3Vpn
// MPLS-L3VPN-STD-MIB notification configuration
type Snmp_Notification_MplsL3Vpn struct {
    EntityData types.CommonEntityData
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

func (mplsL3Vpn *Snmp_Notification_MplsL3Vpn) GetEntityData() *types.CommonEntityData {
    mplsL3Vpn.EntityData.YFilter = mplsL3Vpn.YFilter
    mplsL3Vpn.EntityData.YangName = "mpls-l3vpn"
    mplsL3Vpn.EntityData.BundleName = "cisco_ios_xr"
    mplsL3Vpn.EntityData.ParentYangName = "notification"
    mplsL3Vpn.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-vpn-cfg:mpls-l3vpn"
    mplsL3Vpn.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsL3Vpn.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsL3Vpn.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsL3Vpn.EntityData.Children = make(map[string]types.YChild)
    mplsL3Vpn.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsL3Vpn.EntityData.Leafs["max-threshold-reissue-notification-time"] = types.YLeaf{"MaxThresholdReissueNotificationTime", mplsL3Vpn.MaxThresholdReissueNotificationTime}
    mplsL3Vpn.EntityData.Leafs["max-threshold-exceeded"] = types.YLeaf{"MaxThresholdExceeded", mplsL3Vpn.MaxThresholdExceeded}
    mplsL3Vpn.EntityData.Leafs["max-threshold-cleared"] = types.YLeaf{"MaxThresholdCleared", mplsL3Vpn.MaxThresholdCleared}
    mplsL3Vpn.EntityData.Leafs["mid-threshold-exceeded"] = types.YLeaf{"MidThresholdExceeded", mplsL3Vpn.MidThresholdExceeded}
    mplsL3Vpn.EntityData.Leafs["enable"] = types.YLeaf{"Enable", mplsL3Vpn.Enable}
    mplsL3Vpn.EntityData.Leafs["vrf-down"] = types.YLeaf{"VrfDown", mplsL3Vpn.VrfDown}
    mplsL3Vpn.EntityData.Leafs["vrf-up"] = types.YLeaf{"VrfUp", mplsL3Vpn.VrfUp}
    return &(mplsL3Vpn.EntityData)
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

    optical.EntityData.Children = make(map[string]types.YChild)
    optical.EntityData.Leafs = make(map[string]types.YLeaf)
    optical.EntityData.Leafs["enable"] = types.YLeaf{"Enable", optical.Enable}
    return &(optical.EntityData)
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

    opticalOts.EntityData.Children = make(map[string]types.YChild)
    opticalOts.EntityData.Leafs = make(map[string]types.YLeaf)
    opticalOts.EntityData.Leafs["enable"] = types.YLeaf{"Enable", opticalOts.Enable}
    return &(opticalOts.EntityData)
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

    otn.EntityData.Children = make(map[string]types.YChild)
    otn.EntityData.Leafs = make(map[string]types.YLeaf)
    otn.EntityData.Leafs["enable"] = types.YLeaf{"Enable", otn.Enable}
    return &(otn.EntityData)
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

    bridge.EntityData.Children = make(map[string]types.YChild)
    bridge.EntityData.Leafs = make(map[string]types.YLeaf)
    bridge.EntityData.Leafs["enable"] = types.YLeaf{"Enable", bridge.Enable}
    return &(bridge.EntityData)
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

    sensor.EntityData.Children = make(map[string]types.YChild)
    sensor.EntityData.Leafs = make(map[string]types.YLeaf)
    sensor.EntityData.Leafs["enable"] = types.YLeaf{"Enable", sensor.Enable}
    return &(sensor.EntityData)
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

    ciscoEntityExt.EntityData.Children = make(map[string]types.YChild)
    ciscoEntityExt.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoEntityExt.EntityData.Leafs["enable"] = types.YLeaf{"Enable", ciscoEntityExt.Enable}
    return &(ciscoEntityExt.EntityData)
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

    entity.EntityData.Children = make(map[string]types.YChild)
    entity.EntityData.Leafs = make(map[string]types.YLeaf)
    entity.EntityData.Leafs["enable"] = types.YLeaf{"Enable", entity.Enable}
    return &(entity.EntityData)
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

    entityState.EntityData.Children = make(map[string]types.YChild)
    entityState.EntityData.Leafs = make(map[string]types.YLeaf)
    entityState.EntityData.Leafs["switchover"] = types.YLeaf{"Switchover", entityState.Switchover}
    entityState.EntityData.Leafs["oper-status"] = types.YLeaf{"OperStatus", entityState.OperStatus}
    return &(entityState.EntityData)
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

    fruControl.EntityData.Children = make(map[string]types.YChild)
    fruControl.EntityData.Leafs = make(map[string]types.YLeaf)
    fruControl.EntityData.Leafs["enable"] = types.YLeaf{"Enable", fruControl.Enable}
    return &(fruControl.EntityData)
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

    rf.EntityData.Children = make(map[string]types.YChild)
    rf.EntityData.Leafs = make(map[string]types.YLeaf)
    rf.EntityData.Leafs["enable"] = types.YLeaf{"Enable", rf.Enable}
    return &(rf.EntityData)
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

    syslog.EntityData.Children = make(map[string]types.YChild)
    syslog.EntityData.Leafs = make(map[string]types.YLeaf)
    syslog.EntityData.Leafs["enable"] = types.YLeaf{"Enable", syslog.Enable}
    return &(syslog.EntityData)
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

    subscriberMib.EntityData.Children = make(map[string]types.YChild)
    subscriberMib.EntityData.Children["session-aggregate"] = types.YChild{"SessionAggregate", &subscriberMib.SessionAggregate}
    subscriberMib.EntityData.Leafs = make(map[string]types.YLeaf)
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

    sessionAggregate.EntityData.Children = make(map[string]types.YChild)
    sessionAggregate.EntityData.Leafs = make(map[string]types.YLeaf)
    sessionAggregate.EntityData.Leafs["node"] = types.YLeaf{"Node", sessionAggregate.Node}
    sessionAggregate.EntityData.Leafs["access-interface"] = types.YLeaf{"AccessInterface", sessionAggregate.AccessInterface}
    return &(sessionAggregate.EntityData)
}

// Snmp_Notification_L2Tun
// Enable SNMP l2tun traps
type Snmp_Notification_L2Tun struct {
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

func (l2Tun *Snmp_Notification_L2Tun) GetEntityData() *types.CommonEntityData {
    l2Tun.EntityData.YFilter = l2Tun.YFilter
    l2Tun.EntityData.YangName = "l2tun"
    l2Tun.EntityData.BundleName = "cisco_ios_xr"
    l2Tun.EntityData.ParentYangName = "notification"
    l2Tun.EntityData.SegmentPath = "Cisco-IOS-XR-tunnel-l2tun-proto-mibs-cfg:l2tun"
    l2Tun.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2Tun.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2Tun.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2Tun.EntityData.Children = make(map[string]types.YChild)
    l2Tun.EntityData.Leafs = make(map[string]types.YLeaf)
    l2Tun.EntityData.Leafs["tunnel-up"] = types.YLeaf{"TunnelUp", l2Tun.TunnelUp}
    l2Tun.EntityData.Leafs["tunnel-down"] = types.YLeaf{"TunnelDown", l2Tun.TunnelDown}
    l2Tun.EntityData.Leafs["pseudowire-status"] = types.YLeaf{"PseudowireStatus", l2Tun.PseudowireStatus}
    l2Tun.EntityData.Leafs["sessions"] = types.YLeaf{"Sessions", l2Tun.Sessions}
    return &(l2Tun.EntityData)
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

    correlator.EntityData.Children = make(map[string]types.YChild)
    correlator.EntityData.Children["rules"] = types.YChild{"Rules", &correlator.Rules}
    correlator.EntityData.Children["rule-sets"] = types.YChild{"RuleSets", &correlator.RuleSets}
    correlator.EntityData.Leafs = make(map[string]types.YLeaf)
    correlator.EntityData.Leafs["buffer-size"] = types.YLeaf{"BufferSize", correlator.BufferSize}
    return &(correlator.EntityData)
}

// Snmp_Correlator_Rules
// Table of configured rules
type Snmp_Correlator_Rules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rule name. The type is slice of Snmp_Correlator_Rules_Rule.
    Rule []Snmp_Correlator_Rules_Rule
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

    rules.EntityData.Children = make(map[string]types.YChild)
    rules.EntityData.Children["rule"] = types.YChild{"Rule", nil}
    for i := range rules.Rule {
        rules.EntityData.Children[types.GetSegmentPath(&rules.Rule[i])] = types.YChild{"Rule", &rules.Rule[i]}
    }
    rules.EntityData.Leafs = make(map[string]types.YLeaf)
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
    rule.EntityData.SegmentPath = "rule" + "[name='" + fmt.Sprintf("%v", rule.Name) + "']"
    rule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rule.EntityData.Children = make(map[string]types.YChild)
    rule.EntityData.Children["root-causes"] = types.YChild{"RootCauses", &rule.RootCauses}
    rule.EntityData.Children["non-root-causes"] = types.YChild{"NonRootCauses", &rule.NonRootCauses}
    rule.EntityData.Children["applied-to"] = types.YChild{"AppliedTo", &rule.AppliedTo}
    rule.EntityData.Leafs = make(map[string]types.YLeaf)
    rule.EntityData.Leafs["name"] = types.YLeaf{"Name", rule.Name}
    rule.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", rule.Timeout}
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
    RootCause []Snmp_Correlator_Rules_Rule_RootCauses_RootCause
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

    rootCauses.EntityData.Children = make(map[string]types.YChild)
    rootCauses.EntityData.Children["root-cause"] = types.YChild{"RootCause", nil}
    for i := range rootCauses.RootCause {
        rootCauses.EntityData.Children[types.GetSegmentPath(&rootCauses.RootCause[i])] = types.YChild{"RootCause", &rootCauses.RootCause[i]}
    }
    rootCauses.EntityData.Leafs = make(map[string]types.YLeaf)
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
    rootCause.EntityData.SegmentPath = "root-cause" + "[oid='" + fmt.Sprintf("%v", rootCause.Oid) + "']"
    rootCause.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rootCause.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rootCause.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rootCause.EntityData.Children = make(map[string]types.YChild)
    rootCause.EntityData.Children["var-binds"] = types.YChild{"VarBinds", &rootCause.VarBinds}
    rootCause.EntityData.Leafs = make(map[string]types.YLeaf)
    rootCause.EntityData.Leafs["oid"] = types.YLeaf{"Oid", rootCause.Oid}
    rootCause.EntityData.Leafs["created"] = types.YLeaf{"Created", rootCause.Created}
    return &(rootCause.EntityData)
}

// Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds
// Varbinds to match
type Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Varbind match conditions. The type is slice of
    // Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds_VarBind.
    VarBind []Snmp_Correlator_Rules_Rule_RootCauses_RootCause_VarBinds_VarBind
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

    varBinds.EntityData.Children = make(map[string]types.YChild)
    varBinds.EntityData.Children["var-bind"] = types.YChild{"VarBind", nil}
    for i := range varBinds.VarBind {
        varBinds.EntityData.Children[types.GetSegmentPath(&varBinds.VarBind[i])] = types.YChild{"VarBind", &varBinds.VarBind[i]}
    }
    varBinds.EntityData.Leafs = make(map[string]types.YLeaf)
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
    varBind.EntityData.SegmentPath = "var-bind" + "[oid='" + fmt.Sprintf("%v", varBind.Oid) + "']"
    varBind.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    varBind.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    varBind.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    varBind.EntityData.Children = make(map[string]types.YChild)
    varBind.EntityData.Children["match"] = types.YChild{"Match", &varBind.Match}
    varBind.EntityData.Leafs = make(map[string]types.YLeaf)
    varBind.EntityData.Leafs["oid"] = types.YLeaf{"Oid", varBind.Oid}
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

    match.EntityData.Children = make(map[string]types.YChild)
    match.EntityData.Leafs = make(map[string]types.YLeaf)
    match.EntityData.Leafs["value"] = types.YLeaf{"Value", match.Value}
    match.EntityData.Leafs["index"] = types.YLeaf{"Index", match.Index}
    return &(match.EntityData)
}

// Snmp_Correlator_Rules_Rule_NonRootCauses
// Table of configured non-rootcause
type Snmp_Correlator_Rules_Rule_NonRootCauses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A non-rootcause. The type is slice of
    // Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause.
    NonRootCause []Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause
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

    nonRootCauses.EntityData.Children = make(map[string]types.YChild)
    nonRootCauses.EntityData.Children["non-root-cause"] = types.YChild{"NonRootCause", nil}
    for i := range nonRootCauses.NonRootCause {
        nonRootCauses.EntityData.Children[types.GetSegmentPath(&nonRootCauses.NonRootCause[i])] = types.YChild{"NonRootCause", &nonRootCauses.NonRootCause[i]}
    }
    nonRootCauses.EntityData.Leafs = make(map[string]types.YLeaf)
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
    nonRootCause.EntityData.SegmentPath = "non-root-cause" + "[oid='" + fmt.Sprintf("%v", nonRootCause.Oid) + "']"
    nonRootCause.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nonRootCause.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nonRootCause.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nonRootCause.EntityData.Children = make(map[string]types.YChild)
    nonRootCause.EntityData.Children["var-binds"] = types.YChild{"VarBinds", &nonRootCause.VarBinds}
    nonRootCause.EntityData.Leafs = make(map[string]types.YLeaf)
    nonRootCause.EntityData.Leafs["oid"] = types.YLeaf{"Oid", nonRootCause.Oid}
    nonRootCause.EntityData.Leafs["created"] = types.YLeaf{"Created", nonRootCause.Created}
    return &(nonRootCause.EntityData)
}

// Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds
// Varbinds to match
type Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Varbind match conditions. The type is slice of
    // Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds_VarBind.
    VarBind []Snmp_Correlator_Rules_Rule_NonRootCauses_NonRootCause_VarBinds_VarBind
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

    varBinds.EntityData.Children = make(map[string]types.YChild)
    varBinds.EntityData.Children["var-bind"] = types.YChild{"VarBind", nil}
    for i := range varBinds.VarBind {
        varBinds.EntityData.Children[types.GetSegmentPath(&varBinds.VarBind[i])] = types.YChild{"VarBind", &varBinds.VarBind[i]}
    }
    varBinds.EntityData.Leafs = make(map[string]types.YLeaf)
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
    varBind.EntityData.SegmentPath = "var-bind" + "[oid='" + fmt.Sprintf("%v", varBind.Oid) + "']"
    varBind.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    varBind.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    varBind.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    varBind.EntityData.Children = make(map[string]types.YChild)
    varBind.EntityData.Children["match"] = types.YChild{"Match", &varBind.Match}
    varBind.EntityData.Leafs = make(map[string]types.YLeaf)
    varBind.EntityData.Leafs["oid"] = types.YLeaf{"Oid", varBind.Oid}
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

    match.EntityData.Children = make(map[string]types.YChild)
    match.EntityData.Leafs = make(map[string]types.YLeaf)
    match.EntityData.Leafs["value"] = types.YLeaf{"Value", match.Value}
    match.EntityData.Leafs["index"] = types.YLeaf{"Index", match.Index}
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

    appliedTo.EntityData.Children = make(map[string]types.YChild)
    appliedTo.EntityData.Children["hosts"] = types.YChild{"Hosts", &appliedTo.Hosts}
    appliedTo.EntityData.Leafs = make(map[string]types.YLeaf)
    appliedTo.EntityData.Leafs["all"] = types.YLeaf{"All", appliedTo.All}
    return &(appliedTo.EntityData)
}

// Snmp_Correlator_Rules_Rule_AppliedTo_Hosts
// Table of configured hosts to apply rules to
type Snmp_Correlator_Rules_Rule_AppliedTo_Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A destination host. The type is slice of
    // Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host.
    Host []Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host
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

    hosts.EntityData.Children = make(map[string]types.YChild)
    hosts.EntityData.Children["host"] = types.YChild{"Host", nil}
    for i := range hosts.Host {
        hosts.EntityData.Children[types.GetSegmentPath(&hosts.Host[i])] = types.YChild{"Host", &hosts.Host[i]}
    }
    hosts.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hosts.EntityData)
}

// Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host
// A destination host
type Snmp_Correlator_Rules_Rule_AppliedTo_Hosts_Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    host.EntityData.SegmentPath = "host" + "[ip-address='" + fmt.Sprintf("%v", host.IpAddress) + "']" + "[port='" + fmt.Sprintf("%v", host.Port) + "']"
    host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    host.EntityData.Children = make(map[string]types.YChild)
    host.EntityData.Leafs = make(map[string]types.YLeaf)
    host.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", host.IpAddress}
    host.EntityData.Leafs["port"] = types.YLeaf{"Port", host.Port}
    return &(host.EntityData)
}

// Snmp_Correlator_RuleSets
// Table of configured rulesets
type Snmp_Correlator_RuleSets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Ruleset name. The type is slice of Snmp_Correlator_RuleSets_RuleSet.
    RuleSet []Snmp_Correlator_RuleSets_RuleSet
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

    ruleSets.EntityData.Children = make(map[string]types.YChild)
    ruleSets.EntityData.Children["rule-set"] = types.YChild{"RuleSet", nil}
    for i := range ruleSets.RuleSet {
        ruleSets.EntityData.Children[types.GetSegmentPath(&ruleSets.RuleSet[i])] = types.YChild{"RuleSet", &ruleSets.RuleSet[i]}
    }
    ruleSets.EntityData.Leafs = make(map[string]types.YLeaf)
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
    ruleSet.EntityData.SegmentPath = "rule-set" + "[name='" + fmt.Sprintf("%v", ruleSet.Name) + "']"
    ruleSet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ruleSet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ruleSet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ruleSet.EntityData.Children = make(map[string]types.YChild)
    ruleSet.EntityData.Children["rulenames"] = types.YChild{"Rulenames", &ruleSet.Rulenames}
    ruleSet.EntityData.Children["applied-to"] = types.YChild{"AppliedTo", &ruleSet.AppliedTo}
    ruleSet.EntityData.Leafs = make(map[string]types.YLeaf)
    ruleSet.EntityData.Leafs["name"] = types.YLeaf{"Name", ruleSet.Name}
    return &(ruleSet.EntityData)
}

// Snmp_Correlator_RuleSets_RuleSet_Rulenames
// Table of configured rulenames
type Snmp_Correlator_RuleSets_RuleSet_Rulenames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A rulename. The type is slice of
    // Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename.
    Rulename []Snmp_Correlator_RuleSets_RuleSet_Rulenames_Rulename
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

    rulenames.EntityData.Children = make(map[string]types.YChild)
    rulenames.EntityData.Children["rulename"] = types.YChild{"Rulename", nil}
    for i := range rulenames.Rulename {
        rulenames.EntityData.Children[types.GetSegmentPath(&rulenames.Rulename[i])] = types.YChild{"Rulename", &rulenames.Rulename[i]}
    }
    rulenames.EntityData.Leafs = make(map[string]types.YLeaf)
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
    rulename.EntityData.SegmentPath = "rulename" + "[rulename='" + fmt.Sprintf("%v", rulename.Rulename) + "']"
    rulename.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rulename.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rulename.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rulename.EntityData.Children = make(map[string]types.YChild)
    rulename.EntityData.Leafs = make(map[string]types.YLeaf)
    rulename.EntityData.Leafs["rulename"] = types.YLeaf{"Rulename", rulename.Rulename}
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

    appliedTo.EntityData.Children = make(map[string]types.YChild)
    appliedTo.EntityData.Children["hosts"] = types.YChild{"Hosts", &appliedTo.Hosts}
    appliedTo.EntityData.Leafs = make(map[string]types.YLeaf)
    appliedTo.EntityData.Leafs["all"] = types.YLeaf{"All", appliedTo.All}
    return &(appliedTo.EntityData)
}

// Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts
// Table of configured hosts to apply rules to
type Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A destination host. The type is slice of
    // Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host.
    Host []Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host
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

    hosts.EntityData.Children = make(map[string]types.YChild)
    hosts.EntityData.Children["host"] = types.YChild{"Host", nil}
    for i := range hosts.Host {
        hosts.EntityData.Children[types.GetSegmentPath(&hosts.Host[i])] = types.YChild{"Host", &hosts.Host[i]}
    }
    hosts.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hosts.EntityData)
}

// Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host
// A destination host
type Snmp_Correlator_RuleSets_RuleSet_AppliedTo_Hosts_Host struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address. The type is one of the following
    // types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    host.EntityData.SegmentPath = "host" + "[ip-address='" + fmt.Sprintf("%v", host.IpAddress) + "']" + "[port='" + fmt.Sprintf("%v", host.Port) + "']"
    host.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    host.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    host.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    host.EntityData.Children = make(map[string]types.YChild)
    host.EntityData.Leafs = make(map[string]types.YLeaf)
    host.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", host.IpAddress}
    host.EntityData.Leafs["port"] = types.YLeaf{"Port", host.Port}
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

    bulkStats.EntityData.Children = make(map[string]types.YChild)
    bulkStats.EntityData.Children["schemas"] = types.YChild{"Schemas", &bulkStats.Schemas}
    bulkStats.EntityData.Children["objects"] = types.YChild{"Objects", &bulkStats.Objects}
    bulkStats.EntityData.Children["transfers"] = types.YChild{"Transfers", &bulkStats.Transfers}
    bulkStats.EntityData.Leafs = make(map[string]types.YLeaf)
    bulkStats.EntityData.Leafs["memory"] = types.YLeaf{"Memory", bulkStats.Memory}
    return &(bulkStats.EntityData)
}

// Snmp_BulkStats_Schemas
// Configure schema definition
type Snmp_BulkStats_Schemas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the Schema. The type is slice of Snmp_BulkStats_Schemas_Schema.
    Schema []Snmp_BulkStats_Schemas_Schema
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

    schemas.EntityData.Children = make(map[string]types.YChild)
    schemas.EntityData.Children["schema"] = types.YChild{"Schema", nil}
    for i := range schemas.Schema {
        schemas.EntityData.Children[types.GetSegmentPath(&schemas.Schema[i])] = types.YChild{"Schema", &schemas.Schema[i]}
    }
    schemas.EntityData.Leafs = make(map[string]types.YLeaf)
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
    Type_ interface{}

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
    schema.EntityData.SegmentPath = "schema" + "[schema-name='" + fmt.Sprintf("%v", schema.SchemaName) + "']"
    schema.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    schema.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    schema.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    schema.EntityData.Children = make(map[string]types.YChild)
    schema.EntityData.Children["instance"] = types.YChild{"Instance", &schema.Instance}
    schema.EntityData.Leafs = make(map[string]types.YLeaf)
    schema.EntityData.Leafs["schema-name"] = types.YLeaf{"SchemaName", schema.SchemaName}
    schema.EntityData.Leafs["type"] = types.YLeaf{"Type_", schema.Type_}
    schema.EntityData.Leafs["schema-object-list"] = types.YLeaf{"SchemaObjectList", schema.SchemaObjectList}
    schema.EntityData.Leafs["poll-interval"] = types.YLeaf{"PollInterval", schema.PollInterval}
    return &(schema.EntityData)
}

// Snmp_BulkStats_Schemas_Schema_Instance
// Object instance information
// This type is a presence type.
type Snmp_BulkStats_Schemas_Schema_Instance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of the instance. The type is SnmpBulkstatSchema. This attribute is
    // mandatory.
    Type_ interface{}

    // Instance of the schema. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

func (instance *Snmp_BulkStats_Schemas_Schema_Instance) GetEntityData() *types.CommonEntityData {
    instance.EntityData.YFilter = instance.YFilter
    instance.EntityData.YangName = "instance"
    instance.EntityData.BundleName = "cisco_ios_xr"
    instance.EntityData.ParentYangName = "schema"
    instance.EntityData.SegmentPath = "instance"
    instance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    instance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    instance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    instance.EntityData.Children = make(map[string]types.YChild)
    instance.EntityData.Leafs = make(map[string]types.YLeaf)
    instance.EntityData.Leafs["type"] = types.YLeaf{"Type_", instance.Type_}
    instance.EntityData.Leafs["instance"] = types.YLeaf{"Instance", instance.Instance}
    instance.EntityData.Leafs["start"] = types.YLeaf{"Start", instance.Start}
    instance.EntityData.Leafs["end"] = types.YLeaf{"End", instance.End}
    instance.EntityData.Leafs["max"] = types.YLeaf{"Max", instance.Max}
    instance.EntityData.Leafs["sub-interface"] = types.YLeaf{"SubInterface", instance.SubInterface}
    return &(instance.EntityData)
}

// Snmp_BulkStats_Objects
// Configure an Object List 
type Snmp_BulkStats_Objects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the object List. The type is slice of
    // Snmp_BulkStats_Objects_Object.
    Object []Snmp_BulkStats_Objects_Object
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

    objects.EntityData.Children = make(map[string]types.YChild)
    objects.EntityData.Children["object"] = types.YChild{"Object", nil}
    for i := range objects.Object {
        objects.EntityData.Children[types.GetSegmentPath(&objects.Object[i])] = types.YChild{"Object", &objects.Object[i]}
    }
    objects.EntityData.Leafs = make(map[string]types.YLeaf)
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
    Type_ interface{}

    // Configure an object List.
    Objects Snmp_BulkStats_Objects_Object_Objects_
}

func (object *Snmp_BulkStats_Objects_Object) GetEntityData() *types.CommonEntityData {
    object.EntityData.YFilter = object.YFilter
    object.EntityData.YangName = "object"
    object.EntityData.BundleName = "cisco_ios_xr"
    object.EntityData.ParentYangName = "objects"
    object.EntityData.SegmentPath = "object" + "[object-list-name='" + fmt.Sprintf("%v", object.ObjectListName) + "']"
    object.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    object.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    object.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    object.EntityData.Children = make(map[string]types.YChild)
    object.EntityData.Children["objects"] = types.YChild{"Objects", &object.Objects}
    object.EntityData.Leafs = make(map[string]types.YLeaf)
    object.EntityData.Leafs["object-list-name"] = types.YLeaf{"ObjectListName", object.ObjectListName}
    object.EntityData.Leafs["type"] = types.YLeaf{"Type_", object.Type_}
    return &(object.EntityData)
}

// Snmp_BulkStats_Objects_Object_Objects_
// Configure an object List
type Snmp_BulkStats_Objects_Object_Objects_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object name or OID. The type is slice of
    // Snmp_BulkStats_Objects_Object_Objects__Object.
    Object []Snmp_BulkStats_Objects_Object_Objects__Object_
}

func (objects_ *Snmp_BulkStats_Objects_Object_Objects_) GetEntityData() *types.CommonEntityData {
    objects_.EntityData.YFilter = objects_.YFilter
    objects_.EntityData.YangName = "objects"
    objects_.EntityData.BundleName = "cisco_ios_xr"
    objects_.EntityData.ParentYangName = "object"
    objects_.EntityData.SegmentPath = "objects"
    objects_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    objects_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    objects_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    objects_.EntityData.Children = make(map[string]types.YChild)
    objects_.EntityData.Children["object"] = types.YChild{"Object", nil}
    for i := range objects_.Object {
        objects_.EntityData.Children[types.GetSegmentPath(&objects_.Object[i])] = types.YChild{"Object", &objects_.Object[i]}
    }
    objects_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(objects_.EntityData)
}

// Snmp_BulkStats_Objects_Object_Objects__Object_
// Object name or OID
type Snmp_BulkStats_Objects_Object_Objects__Object_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Object name or OID . The type is string.
    Oid interface{}
}

func (object_ *Snmp_BulkStats_Objects_Object_Objects__Object_) GetEntityData() *types.CommonEntityData {
    object_.EntityData.YFilter = object_.YFilter
    object_.EntityData.YangName = "object"
    object_.EntityData.BundleName = "cisco_ios_xr"
    object_.EntityData.ParentYangName = "objects"
    object_.EntityData.SegmentPath = "object" + "[oid='" + fmt.Sprintf("%v", object_.Oid) + "']"
    object_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    object_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    object_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    object_.EntityData.Children = make(map[string]types.YChild)
    object_.EntityData.Leafs = make(map[string]types.YLeaf)
    object_.EntityData.Leafs["oid"] = types.YLeaf{"Oid", object_.Oid}
    return &(object_.EntityData)
}

// Snmp_BulkStats_Transfers
// Periodicity for the transfer of bulk data in
// minutes
type Snmp_BulkStats_Transfers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of bulk transfer. The type is slice of
    // Snmp_BulkStats_Transfers_Transfer.
    Transfer []Snmp_BulkStats_Transfers_Transfer
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

    transfers.EntityData.Children = make(map[string]types.YChild)
    transfers.EntityData.Children["transfer"] = types.YChild{"Transfer", nil}
    for i := range transfers.Transfer {
        transfers.EntityData.Children[types.GetSegmentPath(&transfers.Transfer[i])] = types.YChild{"Transfer", &transfers.Transfer[i]}
    }
    transfers.EntityData.Leafs = make(map[string]types.YLeaf)
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
    Type_ interface{}

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

func (transfer *Snmp_BulkStats_Transfers_Transfer) GetEntityData() *types.CommonEntityData {
    transfer.EntityData.YFilter = transfer.YFilter
    transfer.EntityData.YangName = "transfer"
    transfer.EntityData.BundleName = "cisco_ios_xr"
    transfer.EntityData.ParentYangName = "transfers"
    transfer.EntityData.SegmentPath = "transfer" + "[transfer-name='" + fmt.Sprintf("%v", transfer.TransferName) + "']"
    transfer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transfer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transfer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transfer.EntityData.Children = make(map[string]types.YChild)
    transfer.EntityData.Children["transfer-schemas"] = types.YChild{"TransferSchemas", &transfer.TransferSchemas}
    transfer.EntityData.Leafs = make(map[string]types.YLeaf)
    transfer.EntityData.Leafs["transfer-name"] = types.YLeaf{"TransferName", transfer.TransferName}
    transfer.EntityData.Leafs["secondary"] = types.YLeaf{"Secondary", transfer.Secondary}
    transfer.EntityData.Leafs["type"] = types.YLeaf{"Type_", transfer.Type_}
    transfer.EntityData.Leafs["buffer-size"] = types.YLeaf{"BufferSize", transfer.BufferSize}
    transfer.EntityData.Leafs["retain"] = types.YLeaf{"Retain", transfer.Retain}
    transfer.EntityData.Leafs["format"] = types.YLeaf{"Format", transfer.Format}
    transfer.EntityData.Leafs["retry"] = types.YLeaf{"Retry", transfer.Retry}
    transfer.EntityData.Leafs["enable"] = types.YLeaf{"Enable", transfer.Enable}
    transfer.EntityData.Leafs["primary"] = types.YLeaf{"Primary", transfer.Primary}
    transfer.EntityData.Leafs["interval"] = types.YLeaf{"Interval", transfer.Interval}
    return &(transfer.EntityData)
}

// Snmp_BulkStats_Transfers_Transfer_TransferSchemas
// Schema that contains objects to be collected
type Snmp_BulkStats_Transfers_Transfer_TransferSchemas struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Schema that contains objects to be collected. The type is slice of
    // Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema.
    TransferSchema []Snmp_BulkStats_Transfers_Transfer_TransferSchemas_TransferSchema
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

    transferSchemas.EntityData.Children = make(map[string]types.YChild)
    transferSchemas.EntityData.Children["transfer-schema"] = types.YChild{"TransferSchema", nil}
    for i := range transferSchemas.TransferSchema {
        transferSchemas.EntityData.Children[types.GetSegmentPath(&transferSchemas.TransferSchema[i])] = types.YChild{"TransferSchema", &transferSchemas.TransferSchema[i]}
    }
    transferSchemas.EntityData.Leafs = make(map[string]types.YLeaf)
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
    transferSchema.EntityData.SegmentPath = "transfer-schema" + "[schema-name='" + fmt.Sprintf("%v", transferSchema.SchemaName) + "']"
    transferSchema.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    transferSchema.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    transferSchema.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    transferSchema.EntityData.Children = make(map[string]types.YChild)
    transferSchema.EntityData.Leafs = make(map[string]types.YLeaf)
    transferSchema.EntityData.Leafs["schema-name"] = types.YLeaf{"SchemaName", transferSchema.SchemaName}
    return &(transferSchema.EntityData)
}

// Snmp_DefaultCommunityMaps
// Container class to hold unencrpted community map
type Snmp_DefaultCommunityMaps struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Unencrpted SNMP community map name . The type is slice of
    // Snmp_DefaultCommunityMaps_DefaultCommunityMap.
    DefaultCommunityMap []Snmp_DefaultCommunityMaps_DefaultCommunityMap
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

    defaultCommunityMaps.EntityData.Children = make(map[string]types.YChild)
    defaultCommunityMaps.EntityData.Children["default-community-map"] = types.YChild{"DefaultCommunityMap", nil}
    for i := range defaultCommunityMaps.DefaultCommunityMap {
        defaultCommunityMaps.EntityData.Children[types.GetSegmentPath(&defaultCommunityMaps.DefaultCommunityMap[i])] = types.YChild{"DefaultCommunityMap", &defaultCommunityMaps.DefaultCommunityMap[i]}
    }
    defaultCommunityMaps.EntityData.Leafs = make(map[string]types.YLeaf)
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
    defaultCommunityMap.EntityData.SegmentPath = "default-community-map" + "[community-name='" + fmt.Sprintf("%v", defaultCommunityMap.CommunityName) + "']"
    defaultCommunityMap.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultCommunityMap.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultCommunityMap.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultCommunityMap.EntityData.Children = make(map[string]types.YChild)
    defaultCommunityMap.EntityData.Leafs = make(map[string]types.YLeaf)
    defaultCommunityMap.EntityData.Leafs["community-name"] = types.YLeaf{"CommunityName", defaultCommunityMap.CommunityName}
    defaultCommunityMap.EntityData.Leafs["context"] = types.YLeaf{"Context", defaultCommunityMap.Context}
    defaultCommunityMap.EntityData.Leafs["security"] = types.YLeaf{"Security", defaultCommunityMap.Security}
    defaultCommunityMap.EntityData.Leafs["target-list"] = types.YLeaf{"TargetList", defaultCommunityMap.TargetList}
    return &(defaultCommunityMap.EntityData)
}

// Snmp_OverloadControl
// Set overload control params for handling
// incoming messages
// This type is a presence type.
type Snmp_OverloadControl struct {
    EntityData types.CommonEntityData
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

func (overloadControl *Snmp_OverloadControl) GetEntityData() *types.CommonEntityData {
    overloadControl.EntityData.YFilter = overloadControl.YFilter
    overloadControl.EntityData.YangName = "overload-control"
    overloadControl.EntityData.BundleName = "cisco_ios_xr"
    overloadControl.EntityData.ParentYangName = "snmp"
    overloadControl.EntityData.SegmentPath = "overload-control"
    overloadControl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    overloadControl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    overloadControl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    overloadControl.EntityData.Children = make(map[string]types.YChild)
    overloadControl.EntityData.Leafs = make(map[string]types.YLeaf)
    overloadControl.EntityData.Leafs["drop-time"] = types.YLeaf{"DropTime", overloadControl.DropTime}
    overloadControl.EntityData.Leafs["throttle-rate"] = types.YLeaf{"ThrottleRate", overloadControl.ThrottleRate}
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

    timeouts.EntityData.Children = make(map[string]types.YChild)
    timeouts.EntityData.Leafs = make(map[string]types.YLeaf)
    timeouts.EntityData.Leafs["duplicates"] = types.YLeaf{"Duplicates", timeouts.Duplicates}
    timeouts.EntityData.Leafs["in-qdrop"] = types.YLeaf{"InQdrop", timeouts.InQdrop}
    timeouts.EntityData.Leafs["subagent"] = types.YLeaf{"Subagent", timeouts.Subagent}
    timeouts.EntityData.Leafs["pdu-stats"] = types.YLeaf{"PduStats", timeouts.PduStats}
    return &(timeouts.EntityData)
}

// Snmp_Users
// Define a user who can access the SNMP engine
type Snmp_Users struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the user. The type is slice of Snmp_Users_User.
    User []Snmp_Users_User
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

    users.EntityData.Children = make(map[string]types.YChild)
    users.EntityData.Children["user"] = types.YChild{"User", nil}
    for i := range users.User {
        users.EntityData.Children[types.GetSegmentPath(&users.User[i])] = types.YChild{"User", &users.User[i]}
    }
    users.EntityData.Leafs = make(map[string]types.YLeaf)
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
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}

    // Flag to indicate that the privacy password is configured for version 3. The
    // type is interface{}.
    PrivacyPasswordConfigured interface{}

    // The algorithm used des56 or aes128 or aes192or aes256 or 3des. The type is
    // SnmpPrivAlgorithm.
    PrivAlgorithm interface{}

    // The privacy password. The type is string with pattern: b'(!.+)|([^!].+)'.
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    RemoteAddress interface{}

    // UDP port number. The type is interface{} with range: 1..65535.
    Port interface{}
}

func (user *Snmp_Users_User) GetEntityData() *types.CommonEntityData {
    user.EntityData.YFilter = user.YFilter
    user.EntityData.YangName = "user"
    user.EntityData.BundleName = "cisco_ios_xr"
    user.EntityData.ParentYangName = "users"
    user.EntityData.SegmentPath = "user" + "[user-name='" + fmt.Sprintf("%v", user.UserName) + "']"
    user.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    user.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    user.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    user.EntityData.Children = make(map[string]types.YChild)
    user.EntityData.Leafs = make(map[string]types.YLeaf)
    user.EntityData.Leafs["user-name"] = types.YLeaf{"UserName", user.UserName}
    user.EntityData.Leafs["group-name"] = types.YLeaf{"GroupName", user.GroupName}
    user.EntityData.Leafs["version"] = types.YLeaf{"Version", user.Version}
    user.EntityData.Leafs["authentication-password-configured"] = types.YLeaf{"AuthenticationPasswordConfigured", user.AuthenticationPasswordConfigured}
    user.EntityData.Leafs["algorithm"] = types.YLeaf{"Algorithm", user.Algorithm}
    user.EntityData.Leafs["authentication-password"] = types.YLeaf{"AuthenticationPassword", user.AuthenticationPassword}
    user.EntityData.Leafs["privacy-password-configured"] = types.YLeaf{"PrivacyPasswordConfigured", user.PrivacyPasswordConfigured}
    user.EntityData.Leafs["priv-algorithm"] = types.YLeaf{"PrivAlgorithm", user.PrivAlgorithm}
    user.EntityData.Leafs["privacy-password"] = types.YLeaf{"PrivacyPassword", user.PrivacyPassword}
    user.EntityData.Leafs["v4acl-type"] = types.YLeaf{"V4AclType", user.V4AclType}
    user.EntityData.Leafs["v4-access-list"] = types.YLeaf{"V4AccessList", user.V4AccessList}
    user.EntityData.Leafs["v6acl-type"] = types.YLeaf{"V6AclType", user.V6AclType}
    user.EntityData.Leafs["v6-access-list"] = types.YLeaf{"V6AccessList", user.V6AccessList}
    user.EntityData.Leafs["owner"] = types.YLeaf{"Owner", user.Owner}
    user.EntityData.Leafs["remote-address"] = types.YLeaf{"RemoteAddress", user.RemoteAddress}
    user.EntityData.Leafs["port"] = types.YLeaf{"Port", user.Port}
    return &(user.EntityData)
}

// Snmp_Vrfs
// SNMP VRF configuration commands
type Snmp_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is slice of Snmp_Vrfs_Vrf.
    Vrf []Snmp_Vrfs_Vrf
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

    vrfs.EntityData.Children = make(map[string]types.YChild)
    vrfs.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children[types.GetSegmentPath(&vrfs.Vrf[i])] = types.YChild{"Vrf", &vrfs.Vrf[i]}
    }
    vrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfs.EntityData)
}

// Snmp_Vrfs_Vrf
// VRF name
type Snmp_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    vrf.EntityData.SegmentPath = "vrf" + "[name='" + fmt.Sprintf("%v", vrf.Name) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["trap-hosts"] = types.YChild{"TrapHosts", &vrf.TrapHosts}
    vrf.EntityData.Children["contexts"] = types.YChild{"Contexts", &vrf.Contexts}
    vrf.EntityData.Children["context-mappings"] = types.YChild{"ContextMappings", &vrf.ContextMappings}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["name"] = types.YLeaf{"Name", vrf.Name}
    return &(vrf.EntityData)
}

// Snmp_Vrfs_Vrf_TrapHosts
// Specify hosts to receive SNMP notifications
type Snmp_Vrfs_Vrf_TrapHosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify hosts to receive SNMP notifications. The type is slice of
    // Snmp_Vrfs_Vrf_TrapHosts_TrapHost.
    TrapHost []Snmp_Vrfs_Vrf_TrapHosts_TrapHost
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

    trapHosts.EntityData.Children = make(map[string]types.YChild)
    trapHosts.EntityData.Children["trap-host"] = types.YChild{"TrapHost", nil}
    for i := range trapHosts.TrapHost {
        trapHosts.EntityData.Children[types.GetSegmentPath(&trapHosts.TrapHost[i])] = types.YChild{"TrapHost", &trapHosts.TrapHost[i]}
    }
    trapHosts.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trapHosts.EntityData)
}

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost
// Specify hosts to receive SNMP notifications
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of SNMP notification host. The type is
    // one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    trapHost.EntityData.SegmentPath = "trap-host" + "[ip-address='" + fmt.Sprintf("%v", trapHost.IpAddress) + "']"
    trapHost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapHost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapHost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapHost.EntityData.Children = make(map[string]types.YChild)
    trapHost.EntityData.Children["encrypted-user-communities"] = types.YChild{"EncryptedUserCommunities", &trapHost.EncryptedUserCommunities}
    trapHost.EntityData.Children["inform-host"] = types.YChild{"InformHost", &trapHost.InformHost}
    trapHost.EntityData.Children["default-user-communities"] = types.YChild{"DefaultUserCommunities", &trapHost.DefaultUserCommunities}
    trapHost.EntityData.Leafs = make(map[string]types.YLeaf)
    trapHost.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", trapHost.IpAddress}
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
    EncryptedUserCommunity []Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity
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

    encryptedUserCommunities.EntityData.Children = make(map[string]types.YChild)
    encryptedUserCommunities.EntityData.Children["encrypted-user-community"] = types.YChild{"EncryptedUserCommunity", nil}
    for i := range encryptedUserCommunities.EncryptedUserCommunity {
        encryptedUserCommunities.EntityData.Children[types.GetSegmentPath(&encryptedUserCommunities.EncryptedUserCommunity[i])] = types.YChild{"EncryptedUserCommunity", &encryptedUserCommunities.EncryptedUserCommunity[i]}
    }
    encryptedUserCommunities.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(encryptedUserCommunities.EntityData)
}

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity
// Clear/Encrypt Community name associated with
// a trap host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SNMPv1/v2c community string or SNMPv3 user. The
    // type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (encryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetEntityData() *types.CommonEntityData {
    encryptedUserCommunity.EntityData.YFilter = encryptedUserCommunity.YFilter
    encryptedUserCommunity.EntityData.YangName = "encrypted-user-community"
    encryptedUserCommunity.EntityData.BundleName = "cisco_ios_xr"
    encryptedUserCommunity.EntityData.ParentYangName = "encrypted-user-communities"
    encryptedUserCommunity.EntityData.SegmentPath = "encrypted-user-community" + "[community-name='" + fmt.Sprintf("%v", encryptedUserCommunity.CommunityName) + "']"
    encryptedUserCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryptedUserCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryptedUserCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryptedUserCommunity.EntityData.Children = make(map[string]types.YChild)
    encryptedUserCommunity.EntityData.Leafs = make(map[string]types.YLeaf)
    encryptedUserCommunity.EntityData.Leafs["community-name"] = types.YLeaf{"CommunityName", encryptedUserCommunity.CommunityName}
    encryptedUserCommunity.EntityData.Leafs["port"] = types.YLeaf{"Port", encryptedUserCommunity.Port}
    encryptedUserCommunity.EntityData.Leafs["version"] = types.YLeaf{"Version", encryptedUserCommunity.Version}
    encryptedUserCommunity.EntityData.Leafs["security-level"] = types.YLeaf{"SecurityLevel", encryptedUserCommunity.SecurityLevel}
    encryptedUserCommunity.EntityData.Leafs["basic-trap-types"] = types.YLeaf{"BasicTrapTypes", encryptedUserCommunity.BasicTrapTypes}
    encryptedUserCommunity.EntityData.Leafs["advanced-trap-types1"] = types.YLeaf{"AdvancedTrapTypes1", encryptedUserCommunity.AdvancedTrapTypes1}
    encryptedUserCommunity.EntityData.Leafs["advanced-trap-types2"] = types.YLeaf{"AdvancedTrapTypes2", encryptedUserCommunity.AdvancedTrapTypes2}
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

    informHost.EntityData.Children = make(map[string]types.YChild)
    informHost.EntityData.Children["inform-user-communities"] = types.YChild{"InformUserCommunities", &informHost.InformUserCommunities}
    informHost.EntityData.Children["inform-encrypted-user-communities"] = types.YChild{"InformEncryptedUserCommunities", &informHost.InformEncryptedUserCommunities}
    informHost.EntityData.Leafs = make(map[string]types.YLeaf)
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
    InformUserCommunity []Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity
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

    informUserCommunities.EntityData.Children = make(map[string]types.YChild)
    informUserCommunities.EntityData.Children["inform-user-community"] = types.YChild{"InformUserCommunity", nil}
    for i := range informUserCommunities.InformUserCommunity {
        informUserCommunities.EntityData.Children[types.GetSegmentPath(&informUserCommunities.InformUserCommunity[i])] = types.YChild{"InformUserCommunity", &informUserCommunities.InformUserCommunity[i]}
    }
    informUserCommunities.EntityData.Leafs = make(map[string]types.YLeaf)
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

func (informUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetEntityData() *types.CommonEntityData {
    informUserCommunity.EntityData.YFilter = informUserCommunity.YFilter
    informUserCommunity.EntityData.YangName = "inform-user-community"
    informUserCommunity.EntityData.BundleName = "cisco_ios_xr"
    informUserCommunity.EntityData.ParentYangName = "inform-user-communities"
    informUserCommunity.EntityData.SegmentPath = "inform-user-community" + "[community-name='" + fmt.Sprintf("%v", informUserCommunity.CommunityName) + "']"
    informUserCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    informUserCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    informUserCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    informUserCommunity.EntityData.Children = make(map[string]types.YChild)
    informUserCommunity.EntityData.Leafs = make(map[string]types.YLeaf)
    informUserCommunity.EntityData.Leafs["community-name"] = types.YLeaf{"CommunityName", informUserCommunity.CommunityName}
    informUserCommunity.EntityData.Leafs["port"] = types.YLeaf{"Port", informUserCommunity.Port}
    informUserCommunity.EntityData.Leafs["version"] = types.YLeaf{"Version", informUserCommunity.Version}
    informUserCommunity.EntityData.Leafs["security-level"] = types.YLeaf{"SecurityLevel", informUserCommunity.SecurityLevel}
    informUserCommunity.EntityData.Leafs["basic-trap-types"] = types.YLeaf{"BasicTrapTypes", informUserCommunity.BasicTrapTypes}
    informUserCommunity.EntityData.Leafs["advanced-trap-types1"] = types.YLeaf{"AdvancedTrapTypes1", informUserCommunity.AdvancedTrapTypes1}
    informUserCommunity.EntityData.Leafs["advanced-trap-types2"] = types.YLeaf{"AdvancedTrapTypes2", informUserCommunity.AdvancedTrapTypes2}
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
    InformEncryptedUserCommunity []Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity
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

    informEncryptedUserCommunities.EntityData.Children = make(map[string]types.YChild)
    informEncryptedUserCommunities.EntityData.Children["inform-encrypted-user-community"] = types.YChild{"InformEncryptedUserCommunity", nil}
    for i := range informEncryptedUserCommunities.InformEncryptedUserCommunity {
        informEncryptedUserCommunities.EntityData.Children[types.GetSegmentPath(&informEncryptedUserCommunities.InformEncryptedUserCommunity[i])] = types.YChild{"InformEncryptedUserCommunity", &informEncryptedUserCommunities.InformEncryptedUserCommunity[i]}
    }
    informEncryptedUserCommunities.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(informEncryptedUserCommunities.EntityData)
}

// Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity
// Clear/Encrypt Community name associated with
// a inform host
type Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SNMPv2c community string or SNMPv3 user. The type
    // is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (informEncryptedUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetEntityData() *types.CommonEntityData {
    informEncryptedUserCommunity.EntityData.YFilter = informEncryptedUserCommunity.YFilter
    informEncryptedUserCommunity.EntityData.YangName = "inform-encrypted-user-community"
    informEncryptedUserCommunity.EntityData.BundleName = "cisco_ios_xr"
    informEncryptedUserCommunity.EntityData.ParentYangName = "inform-encrypted-user-communities"
    informEncryptedUserCommunity.EntityData.SegmentPath = "inform-encrypted-user-community" + "[community-name='" + fmt.Sprintf("%v", informEncryptedUserCommunity.CommunityName) + "']"
    informEncryptedUserCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    informEncryptedUserCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    informEncryptedUserCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    informEncryptedUserCommunity.EntityData.Children = make(map[string]types.YChild)
    informEncryptedUserCommunity.EntityData.Leafs = make(map[string]types.YLeaf)
    informEncryptedUserCommunity.EntityData.Leafs["community-name"] = types.YLeaf{"CommunityName", informEncryptedUserCommunity.CommunityName}
    informEncryptedUserCommunity.EntityData.Leafs["port"] = types.YLeaf{"Port", informEncryptedUserCommunity.Port}
    informEncryptedUserCommunity.EntityData.Leafs["version"] = types.YLeaf{"Version", informEncryptedUserCommunity.Version}
    informEncryptedUserCommunity.EntityData.Leafs["security-level"] = types.YLeaf{"SecurityLevel", informEncryptedUserCommunity.SecurityLevel}
    informEncryptedUserCommunity.EntityData.Leafs["basic-trap-types"] = types.YLeaf{"BasicTrapTypes", informEncryptedUserCommunity.BasicTrapTypes}
    informEncryptedUserCommunity.EntityData.Leafs["advanced-trap-types1"] = types.YLeaf{"AdvancedTrapTypes1", informEncryptedUserCommunity.AdvancedTrapTypes1}
    informEncryptedUserCommunity.EntityData.Leafs["advanced-trap-types2"] = types.YLeaf{"AdvancedTrapTypes2", informEncryptedUserCommunity.AdvancedTrapTypes2}
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
    DefaultUserCommunity []Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity
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

    defaultUserCommunities.EntityData.Children = make(map[string]types.YChild)
    defaultUserCommunities.EntityData.Children["default-user-community"] = types.YChild{"DefaultUserCommunity", nil}
    for i := range defaultUserCommunities.DefaultUserCommunity {
        defaultUserCommunities.EntityData.Children[types.GetSegmentPath(&defaultUserCommunities.DefaultUserCommunity[i])] = types.YChild{"DefaultUserCommunity", &defaultUserCommunities.DefaultUserCommunity[i]}
    }
    defaultUserCommunities.EntityData.Leafs = make(map[string]types.YLeaf)
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

func (defaultUserCommunity *Snmp_Vrfs_Vrf_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetEntityData() *types.CommonEntityData {
    defaultUserCommunity.EntityData.YFilter = defaultUserCommunity.YFilter
    defaultUserCommunity.EntityData.YangName = "default-user-community"
    defaultUserCommunity.EntityData.BundleName = "cisco_ios_xr"
    defaultUserCommunity.EntityData.ParentYangName = "default-user-communities"
    defaultUserCommunity.EntityData.SegmentPath = "default-user-community" + "[community-name='" + fmt.Sprintf("%v", defaultUserCommunity.CommunityName) + "']"
    defaultUserCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultUserCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultUserCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultUserCommunity.EntityData.Children = make(map[string]types.YChild)
    defaultUserCommunity.EntityData.Leafs = make(map[string]types.YLeaf)
    defaultUserCommunity.EntityData.Leafs["community-name"] = types.YLeaf{"CommunityName", defaultUserCommunity.CommunityName}
    defaultUserCommunity.EntityData.Leafs["port"] = types.YLeaf{"Port", defaultUserCommunity.Port}
    defaultUserCommunity.EntityData.Leafs["version"] = types.YLeaf{"Version", defaultUserCommunity.Version}
    defaultUserCommunity.EntityData.Leafs["security-level"] = types.YLeaf{"SecurityLevel", defaultUserCommunity.SecurityLevel}
    defaultUserCommunity.EntityData.Leafs["basic-trap-types"] = types.YLeaf{"BasicTrapTypes", defaultUserCommunity.BasicTrapTypes}
    defaultUserCommunity.EntityData.Leafs["advanced-trap-types1"] = types.YLeaf{"AdvancedTrapTypes1", defaultUserCommunity.AdvancedTrapTypes1}
    defaultUserCommunity.EntityData.Leafs["advanced-trap-types2"] = types.YLeaf{"AdvancedTrapTypes2", defaultUserCommunity.AdvancedTrapTypes2}
    return &(defaultUserCommunity.EntityData)
}

// Snmp_Vrfs_Vrf_Contexts
// List of Context Names
type Snmp_Vrfs_Vrf_Contexts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Context Name. The type is slice of Snmp_Vrfs_Vrf_Contexts_Context.
    Context []Snmp_Vrfs_Vrf_Contexts_Context
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

    contexts.EntityData.Children = make(map[string]types.YChild)
    contexts.EntityData.Children["context"] = types.YChild{"Context", nil}
    for i := range contexts.Context {
        contexts.EntityData.Children[types.GetSegmentPath(&contexts.Context[i])] = types.YChild{"Context", &contexts.Context[i]}
    }
    contexts.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(contexts.EntityData)
}

// Snmp_Vrfs_Vrf_Contexts_Context
// Context Name
type Snmp_Vrfs_Vrf_Contexts_Context struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Context Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ContextName interface{}
}

func (context *Snmp_Vrfs_Vrf_Contexts_Context) GetEntityData() *types.CommonEntityData {
    context.EntityData.YFilter = context.YFilter
    context.EntityData.YangName = "context"
    context.EntityData.BundleName = "cisco_ios_xr"
    context.EntityData.ParentYangName = "contexts"
    context.EntityData.SegmentPath = "context" + "[context-name='" + fmt.Sprintf("%v", context.ContextName) + "']"
    context.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    context.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    context.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    context.EntityData.Children = make(map[string]types.YChild)
    context.EntityData.Leafs = make(map[string]types.YLeaf)
    context.EntityData.Leafs["context-name"] = types.YLeaf{"ContextName", context.ContextName}
    return &(context.EntityData)
}

// Snmp_Vrfs_Vrf_ContextMappings
// List of context names
type Snmp_Vrfs_Vrf_ContextMappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Context mapping name. The type is slice of
    // Snmp_Vrfs_Vrf_ContextMappings_ContextMapping.
    ContextMapping []Snmp_Vrfs_Vrf_ContextMappings_ContextMapping
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

    contextMappings.EntityData.Children = make(map[string]types.YChild)
    contextMappings.EntityData.Children["context-mapping"] = types.YChild{"ContextMapping", nil}
    for i := range contextMappings.ContextMapping {
        contextMappings.EntityData.Children[types.GetSegmentPath(&contextMappings.ContextMapping[i])] = types.YChild{"ContextMapping", &contextMappings.ContextMapping[i]}
    }
    contextMappings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(contextMappings.EntityData)
}

// Snmp_Vrfs_Vrf_ContextMappings_ContextMapping
// Context mapping name
type Snmp_Vrfs_Vrf_ContextMappings_ContextMapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Context mapping name. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    contextMapping.EntityData.SegmentPath = "context-mapping" + "[context-mapping-name='" + fmt.Sprintf("%v", contextMapping.ContextMappingName) + "']"
    contextMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextMapping.EntityData.Children = make(map[string]types.YChild)
    contextMapping.EntityData.Leafs = make(map[string]types.YLeaf)
    contextMapping.EntityData.Leafs["context-mapping-name"] = types.YLeaf{"ContextMappingName", contextMapping.ContextMappingName}
    contextMapping.EntityData.Leafs["context"] = types.YLeaf{"Context", contextMapping.Context}
    contextMapping.EntityData.Leafs["instance-name"] = types.YLeaf{"InstanceName", contextMapping.InstanceName}
    contextMapping.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", contextMapping.VrfName}
    contextMapping.EntityData.Leafs["topology-name"] = types.YLeaf{"TopologyName", contextMapping.TopologyName}
    return &(contextMapping.EntityData)
}

// Snmp_Groups
// Define a User Security Model group
type Snmp_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the group. The type is slice of Snmp_Groups_Group.
    Group []Snmp_Groups_Group
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

    groups.EntityData.Children = make(map[string]types.YChild)
    groups.EntityData.Children["group"] = types.YChild{"Group", nil}
    for i := range groups.Group {
        groups.EntityData.Children[types.GetSegmentPath(&groups.Group[i])] = types.YChild{"Group", &groups.Group[i]}
    }
    groups.EntityData.Leafs = make(map[string]types.YLeaf)
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

func (group *Snmp_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + "[name='" + fmt.Sprintf("%v", group.Name) + "']"
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = make(map[string]types.YChild)
    group.EntityData.Leafs = make(map[string]types.YLeaf)
    group.EntityData.Leafs["name"] = types.YLeaf{"Name", group.Name}
    group.EntityData.Leafs["snmp-version"] = types.YLeaf{"SnmpVersion", group.SnmpVersion}
    group.EntityData.Leafs["security-model"] = types.YLeaf{"SecurityModel", group.SecurityModel}
    group.EntityData.Leafs["notify-view"] = types.YLeaf{"NotifyView", group.NotifyView}
    group.EntityData.Leafs["read-view"] = types.YLeaf{"ReadView", group.ReadView}
    group.EntityData.Leafs["write-view"] = types.YLeaf{"WriteView", group.WriteView}
    group.EntityData.Leafs["v4acl-type"] = types.YLeaf{"V4AclType", group.V4AclType}
    group.EntityData.Leafs["v4-access-list"] = types.YLeaf{"V4AccessList", group.V4AccessList}
    group.EntityData.Leafs["v6acl-type"] = types.YLeaf{"V6AclType", group.V6AclType}
    group.EntityData.Leafs["v6-access-list"] = types.YLeaf{"V6AccessList", group.V6AccessList}
    group.EntityData.Leafs["context-name"] = types.YLeaf{"ContextName", group.ContextName}
    return &(group.EntityData)
}

// Snmp_TrapHosts
// Specify hosts to receive SNMP notifications
type Snmp_TrapHosts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify hosts to receive SNMP notifications. The type is slice of
    // Snmp_TrapHosts_TrapHost.
    TrapHost []Snmp_TrapHosts_TrapHost
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

    trapHosts.EntityData.Children = make(map[string]types.YChild)
    trapHosts.EntityData.Children["trap-host"] = types.YChild{"TrapHost", nil}
    for i := range trapHosts.TrapHost {
        trapHosts.EntityData.Children[types.GetSegmentPath(&trapHosts.TrapHost[i])] = types.YChild{"TrapHost", &trapHosts.TrapHost[i]}
    }
    trapHosts.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trapHosts.EntityData)
}

// Snmp_TrapHosts_TrapHost
// Specify hosts to receive SNMP notifications
type Snmp_TrapHosts_TrapHost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of SNMP notification host. The type is
    // one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    trapHost.EntityData.SegmentPath = "trap-host" + "[ip-address='" + fmt.Sprintf("%v", trapHost.IpAddress) + "']"
    trapHost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapHost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapHost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapHost.EntityData.Children = make(map[string]types.YChild)
    trapHost.EntityData.Children["encrypted-user-communities"] = types.YChild{"EncryptedUserCommunities", &trapHost.EncryptedUserCommunities}
    trapHost.EntityData.Children["inform-host"] = types.YChild{"InformHost", &trapHost.InformHost}
    trapHost.EntityData.Children["default-user-communities"] = types.YChild{"DefaultUserCommunities", &trapHost.DefaultUserCommunities}
    trapHost.EntityData.Leafs = make(map[string]types.YLeaf)
    trapHost.EntityData.Leafs["ip-address"] = types.YLeaf{"IpAddress", trapHost.IpAddress}
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
    EncryptedUserCommunity []Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity
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

    encryptedUserCommunities.EntityData.Children = make(map[string]types.YChild)
    encryptedUserCommunities.EntityData.Children["encrypted-user-community"] = types.YChild{"EncryptedUserCommunity", nil}
    for i := range encryptedUserCommunities.EncryptedUserCommunity {
        encryptedUserCommunities.EntityData.Children[types.GetSegmentPath(&encryptedUserCommunities.EncryptedUserCommunity[i])] = types.YChild{"EncryptedUserCommunity", &encryptedUserCommunities.EncryptedUserCommunity[i]}
    }
    encryptedUserCommunities.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(encryptedUserCommunities.EntityData)
}

// Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity
// Clear/Encrypt Community name associated with
// a trap host
type Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SNMPv1/v2c community string or SNMPv3 user. The
    // type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (encryptedUserCommunity *Snmp_TrapHosts_TrapHost_EncryptedUserCommunities_EncryptedUserCommunity) GetEntityData() *types.CommonEntityData {
    encryptedUserCommunity.EntityData.YFilter = encryptedUserCommunity.YFilter
    encryptedUserCommunity.EntityData.YangName = "encrypted-user-community"
    encryptedUserCommunity.EntityData.BundleName = "cisco_ios_xr"
    encryptedUserCommunity.EntityData.ParentYangName = "encrypted-user-communities"
    encryptedUserCommunity.EntityData.SegmentPath = "encrypted-user-community" + "[community-name='" + fmt.Sprintf("%v", encryptedUserCommunity.CommunityName) + "']"
    encryptedUserCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryptedUserCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryptedUserCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryptedUserCommunity.EntityData.Children = make(map[string]types.YChild)
    encryptedUserCommunity.EntityData.Leafs = make(map[string]types.YLeaf)
    encryptedUserCommunity.EntityData.Leafs["community-name"] = types.YLeaf{"CommunityName", encryptedUserCommunity.CommunityName}
    encryptedUserCommunity.EntityData.Leafs["port"] = types.YLeaf{"Port", encryptedUserCommunity.Port}
    encryptedUserCommunity.EntityData.Leafs["version"] = types.YLeaf{"Version", encryptedUserCommunity.Version}
    encryptedUserCommunity.EntityData.Leafs["security-level"] = types.YLeaf{"SecurityLevel", encryptedUserCommunity.SecurityLevel}
    encryptedUserCommunity.EntityData.Leafs["basic-trap-types"] = types.YLeaf{"BasicTrapTypes", encryptedUserCommunity.BasicTrapTypes}
    encryptedUserCommunity.EntityData.Leafs["advanced-trap-types1"] = types.YLeaf{"AdvancedTrapTypes1", encryptedUserCommunity.AdvancedTrapTypes1}
    encryptedUserCommunity.EntityData.Leafs["advanced-trap-types2"] = types.YLeaf{"AdvancedTrapTypes2", encryptedUserCommunity.AdvancedTrapTypes2}
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

    informHost.EntityData.Children = make(map[string]types.YChild)
    informHost.EntityData.Children["inform-user-communities"] = types.YChild{"InformUserCommunities", &informHost.InformUserCommunities}
    informHost.EntityData.Children["inform-encrypted-user-communities"] = types.YChild{"InformEncryptedUserCommunities", &informHost.InformEncryptedUserCommunities}
    informHost.EntityData.Leafs = make(map[string]types.YLeaf)
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
    InformUserCommunity []Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity
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

    informUserCommunities.EntityData.Children = make(map[string]types.YChild)
    informUserCommunities.EntityData.Children["inform-user-community"] = types.YChild{"InformUserCommunity", nil}
    for i := range informUserCommunities.InformUserCommunity {
        informUserCommunities.EntityData.Children[types.GetSegmentPath(&informUserCommunities.InformUserCommunity[i])] = types.YChild{"InformUserCommunity", &informUserCommunities.InformUserCommunity[i]}
    }
    informUserCommunities.EntityData.Leafs = make(map[string]types.YLeaf)
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

func (informUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformUserCommunities_InformUserCommunity) GetEntityData() *types.CommonEntityData {
    informUserCommunity.EntityData.YFilter = informUserCommunity.YFilter
    informUserCommunity.EntityData.YangName = "inform-user-community"
    informUserCommunity.EntityData.BundleName = "cisco_ios_xr"
    informUserCommunity.EntityData.ParentYangName = "inform-user-communities"
    informUserCommunity.EntityData.SegmentPath = "inform-user-community" + "[community-name='" + fmt.Sprintf("%v", informUserCommunity.CommunityName) + "']"
    informUserCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    informUserCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    informUserCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    informUserCommunity.EntityData.Children = make(map[string]types.YChild)
    informUserCommunity.EntityData.Leafs = make(map[string]types.YLeaf)
    informUserCommunity.EntityData.Leafs["community-name"] = types.YLeaf{"CommunityName", informUserCommunity.CommunityName}
    informUserCommunity.EntityData.Leafs["port"] = types.YLeaf{"Port", informUserCommunity.Port}
    informUserCommunity.EntityData.Leafs["version"] = types.YLeaf{"Version", informUserCommunity.Version}
    informUserCommunity.EntityData.Leafs["security-level"] = types.YLeaf{"SecurityLevel", informUserCommunity.SecurityLevel}
    informUserCommunity.EntityData.Leafs["basic-trap-types"] = types.YLeaf{"BasicTrapTypes", informUserCommunity.BasicTrapTypes}
    informUserCommunity.EntityData.Leafs["advanced-trap-types1"] = types.YLeaf{"AdvancedTrapTypes1", informUserCommunity.AdvancedTrapTypes1}
    informUserCommunity.EntityData.Leafs["advanced-trap-types2"] = types.YLeaf{"AdvancedTrapTypes2", informUserCommunity.AdvancedTrapTypes2}
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
    InformEncryptedUserCommunity []Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity
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

    informEncryptedUserCommunities.EntityData.Children = make(map[string]types.YChild)
    informEncryptedUserCommunities.EntityData.Children["inform-encrypted-user-community"] = types.YChild{"InformEncryptedUserCommunity", nil}
    for i := range informEncryptedUserCommunities.InformEncryptedUserCommunity {
        informEncryptedUserCommunities.EntityData.Children[types.GetSegmentPath(&informEncryptedUserCommunities.InformEncryptedUserCommunity[i])] = types.YChild{"InformEncryptedUserCommunity", &informEncryptedUserCommunities.InformEncryptedUserCommunity[i]}
    }
    informEncryptedUserCommunities.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(informEncryptedUserCommunities.EntityData)
}

// Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity
// Clear/Encrypt Community name associated with
// a inform host
type Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. SNMPv2c community string or SNMPv3 user. The type
    // is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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

func (informEncryptedUserCommunity *Snmp_TrapHosts_TrapHost_InformHost_InformEncryptedUserCommunities_InformEncryptedUserCommunity) GetEntityData() *types.CommonEntityData {
    informEncryptedUserCommunity.EntityData.YFilter = informEncryptedUserCommunity.YFilter
    informEncryptedUserCommunity.EntityData.YangName = "inform-encrypted-user-community"
    informEncryptedUserCommunity.EntityData.BundleName = "cisco_ios_xr"
    informEncryptedUserCommunity.EntityData.ParentYangName = "inform-encrypted-user-communities"
    informEncryptedUserCommunity.EntityData.SegmentPath = "inform-encrypted-user-community" + "[community-name='" + fmt.Sprintf("%v", informEncryptedUserCommunity.CommunityName) + "']"
    informEncryptedUserCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    informEncryptedUserCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    informEncryptedUserCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    informEncryptedUserCommunity.EntityData.Children = make(map[string]types.YChild)
    informEncryptedUserCommunity.EntityData.Leafs = make(map[string]types.YLeaf)
    informEncryptedUserCommunity.EntityData.Leafs["community-name"] = types.YLeaf{"CommunityName", informEncryptedUserCommunity.CommunityName}
    informEncryptedUserCommunity.EntityData.Leafs["port"] = types.YLeaf{"Port", informEncryptedUserCommunity.Port}
    informEncryptedUserCommunity.EntityData.Leafs["version"] = types.YLeaf{"Version", informEncryptedUserCommunity.Version}
    informEncryptedUserCommunity.EntityData.Leafs["security-level"] = types.YLeaf{"SecurityLevel", informEncryptedUserCommunity.SecurityLevel}
    informEncryptedUserCommunity.EntityData.Leafs["basic-trap-types"] = types.YLeaf{"BasicTrapTypes", informEncryptedUserCommunity.BasicTrapTypes}
    informEncryptedUserCommunity.EntityData.Leafs["advanced-trap-types1"] = types.YLeaf{"AdvancedTrapTypes1", informEncryptedUserCommunity.AdvancedTrapTypes1}
    informEncryptedUserCommunity.EntityData.Leafs["advanced-trap-types2"] = types.YLeaf{"AdvancedTrapTypes2", informEncryptedUserCommunity.AdvancedTrapTypes2}
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
    DefaultUserCommunity []Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity
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

    defaultUserCommunities.EntityData.Children = make(map[string]types.YChild)
    defaultUserCommunities.EntityData.Children["default-user-community"] = types.YChild{"DefaultUserCommunity", nil}
    for i := range defaultUserCommunities.DefaultUserCommunity {
        defaultUserCommunities.EntityData.Children[types.GetSegmentPath(&defaultUserCommunities.DefaultUserCommunity[i])] = types.YChild{"DefaultUserCommunity", &defaultUserCommunities.DefaultUserCommunity[i]}
    }
    defaultUserCommunities.EntityData.Leafs = make(map[string]types.YLeaf)
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

func (defaultUserCommunity *Snmp_TrapHosts_TrapHost_DefaultUserCommunities_DefaultUserCommunity) GetEntityData() *types.CommonEntityData {
    defaultUserCommunity.EntityData.YFilter = defaultUserCommunity.YFilter
    defaultUserCommunity.EntityData.YangName = "default-user-community"
    defaultUserCommunity.EntityData.BundleName = "cisco_ios_xr"
    defaultUserCommunity.EntityData.ParentYangName = "default-user-communities"
    defaultUserCommunity.EntityData.SegmentPath = "default-user-community" + "[community-name='" + fmt.Sprintf("%v", defaultUserCommunity.CommunityName) + "']"
    defaultUserCommunity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultUserCommunity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultUserCommunity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultUserCommunity.EntityData.Children = make(map[string]types.YChild)
    defaultUserCommunity.EntityData.Leafs = make(map[string]types.YLeaf)
    defaultUserCommunity.EntityData.Leafs["community-name"] = types.YLeaf{"CommunityName", defaultUserCommunity.CommunityName}
    defaultUserCommunity.EntityData.Leafs["port"] = types.YLeaf{"Port", defaultUserCommunity.Port}
    defaultUserCommunity.EntityData.Leafs["version"] = types.YLeaf{"Version", defaultUserCommunity.Version}
    defaultUserCommunity.EntityData.Leafs["security-level"] = types.YLeaf{"SecurityLevel", defaultUserCommunity.SecurityLevel}
    defaultUserCommunity.EntityData.Leafs["basic-trap-types"] = types.YLeaf{"BasicTrapTypes", defaultUserCommunity.BasicTrapTypes}
    defaultUserCommunity.EntityData.Leafs["advanced-trap-types1"] = types.YLeaf{"AdvancedTrapTypes1", defaultUserCommunity.AdvancedTrapTypes1}
    defaultUserCommunity.EntityData.Leafs["advanced-trap-types2"] = types.YLeaf{"AdvancedTrapTypes2", defaultUserCommunity.AdvancedTrapTypes2}
    return &(defaultUserCommunity.EntityData)
}

// Snmp_Contexts
// List of Context Names
type Snmp_Contexts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Context Name. The type is slice of Snmp_Contexts_Context.
    Context []Snmp_Contexts_Context
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

    contexts.EntityData.Children = make(map[string]types.YChild)
    contexts.EntityData.Children["context"] = types.YChild{"Context", nil}
    for i := range contexts.Context {
        contexts.EntityData.Children[types.GetSegmentPath(&contexts.Context[i])] = types.YChild{"Context", &contexts.Context[i]}
    }
    contexts.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(contexts.EntityData)
}

// Snmp_Contexts_Context
// Context Name
type Snmp_Contexts_Context struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Context Name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ContextName interface{}
}

func (context *Snmp_Contexts_Context) GetEntityData() *types.CommonEntityData {
    context.EntityData.YFilter = context.YFilter
    context.EntityData.YangName = "context"
    context.EntityData.BundleName = "cisco_ios_xr"
    context.EntityData.ParentYangName = "contexts"
    context.EntityData.SegmentPath = "context" + "[context-name='" + fmt.Sprintf("%v", context.ContextName) + "']"
    context.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    context.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    context.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    context.EntityData.Children = make(map[string]types.YChild)
    context.EntityData.Leafs = make(map[string]types.YLeaf)
    context.EntityData.Leafs["context-name"] = types.YLeaf{"ContextName", context.ContextName}
    return &(context.EntityData)
}

// Snmp_ContextMappings
// List of context names
type Snmp_ContextMappings struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Context mapping name. The type is slice of
    // Snmp_ContextMappings_ContextMapping.
    ContextMapping []Snmp_ContextMappings_ContextMapping
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

    contextMappings.EntityData.Children = make(map[string]types.YChild)
    contextMappings.EntityData.Children["context-mapping"] = types.YChild{"ContextMapping", nil}
    for i := range contextMappings.ContextMapping {
        contextMappings.EntityData.Children[types.GetSegmentPath(&contextMappings.ContextMapping[i])] = types.YChild{"ContextMapping", &contextMappings.ContextMapping[i]}
    }
    contextMappings.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(contextMappings.EntityData)
}

// Snmp_ContextMappings_ContextMapping
// Context mapping name
type Snmp_ContextMappings_ContextMapping struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Context mapping name. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    contextMapping.EntityData.SegmentPath = "context-mapping" + "[context-mapping-name='" + fmt.Sprintf("%v", contextMapping.ContextMappingName) + "']"
    contextMapping.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    contextMapping.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    contextMapping.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    contextMapping.EntityData.Children = make(map[string]types.YChild)
    contextMapping.EntityData.Leafs = make(map[string]types.YLeaf)
    contextMapping.EntityData.Leafs["context-mapping-name"] = types.YLeaf{"ContextMappingName", contextMapping.ContextMappingName}
    contextMapping.EntityData.Leafs["context"] = types.YLeaf{"Context", contextMapping.Context}
    contextMapping.EntityData.Leafs["instance-name"] = types.YLeaf{"InstanceName", contextMapping.InstanceName}
    contextMapping.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", contextMapping.VrfName}
    contextMapping.EntityData.Leafs["topology-name"] = types.YLeaf{"TopologyName", contextMapping.TopologyName}
    return &(contextMapping.EntityData)
}

// Mib
// mib
type Mib struct {
    EntityData types.CommonEntityData
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

func (mib *Mib) GetEntityData() *types.CommonEntityData {
    mib.EntityData.YFilter = mib.YFilter
    mib.EntityData.YangName = "mib"
    mib.EntityData.BundleName = "cisco_ios_xr"
    mib.EntityData.ParentYangName = "Cisco-IOS-XR-snmp-agent-cfg"
    mib.EntityData.SegmentPath = "Cisco-IOS-XR-snmp-agent-cfg:mib"
    mib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mib.EntityData.Children = make(map[string]types.YChild)
    mib.EntityData.Children["Cisco-IOS-XR-mpls-te-cfg:mpls-te-mib"] = types.YChild{"MplsTeMib", &mib.MplsTeMib}
    mib.EntityData.Children["Cisco-IOS-XR-mpls-te-cfg:mpls-p2mp-mib"] = types.YChild{"MplsP2MpMib", &mib.MplsP2MpMib}
    mib.EntityData.Children["Cisco-IOS-XR-mpls-te-cfg:mpls-te-ext-std-mib"] = types.YChild{"MplsTeExtStdMib", &mib.MplsTeExtStdMib}
    mib.EntityData.Children["Cisco-IOS-XR-mpls-te-cfg:mpls-te-ext-mib"] = types.YChild{"MplsTeExtMib", &mib.MplsTeExtMib}
    mib.EntityData.Children["Cisco-IOS-XR-mpls-te-cfg:mpls-frr-mib"] = types.YChild{"MplsFrrMib", &mib.MplsFrrMib}
    mib.EntityData.Children["Cisco-IOS-XR-qos-mibs-cfg:cb-qosmib"] = types.YChild{"CbQosmib", &mib.CbQosmib}
    mib.EntityData.Children["Cisco-IOS-XR-snmp-entitymib-cfg:entity-mib"] = types.YChild{"EntityMib", &mib.EntityMib}
    mib.EntityData.Children["Cisco-IOS-XR-snmp-ifmib-cfg:interface-mib"] = types.YChild{"InterfaceMib", &mib.InterfaceMib}
    mib.EntityData.Children["Cisco-IOS-XR-subscriber-session-mon-mibs-cfg:subscriber"] = types.YChild{"Subscriber", &mib.Subscriber}
    mib.EntityData.Leafs = make(map[string]types.YLeaf)
    mib.EntityData.Leafs["sensor-mib-cache"] = types.YLeaf{"SensorMibCache", mib.SensorMibCache}
    return &(mib.EntityData)
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

    mplsTeMib.EntityData.Children = make(map[string]types.YChild)
    mplsTeMib.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsTeMib.EntityData.Leafs["cache-garbage-collect-timer"] = types.YLeaf{"CacheGarbageCollectTimer", mplsTeMib.CacheGarbageCollectTimer}
    mplsTeMib.EntityData.Leafs["cache-timer"] = types.YLeaf{"CacheTimer", mplsTeMib.CacheTimer}
    return &(mplsTeMib.EntityData)
}

// Mib_MplsP2MpMib
// MPLS P2MP MIB configuration
type Mib_MplsP2MpMib struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure the cache time for the mib. The type is interface{} with range:
    // 0..600. Units are second. The default value is 60.
    CacheTimer interface{}
}

func (mplsP2MpMib *Mib_MplsP2MpMib) GetEntityData() *types.CommonEntityData {
    mplsP2MpMib.EntityData.YFilter = mplsP2MpMib.YFilter
    mplsP2MpMib.EntityData.YangName = "mpls-p2mp-mib"
    mplsP2MpMib.EntityData.BundleName = "cisco_ios_xr"
    mplsP2MpMib.EntityData.ParentYangName = "mib"
    mplsP2MpMib.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-te-cfg:mpls-p2mp-mib"
    mplsP2MpMib.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsP2MpMib.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsP2MpMib.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsP2MpMib.EntityData.Children = make(map[string]types.YChild)
    mplsP2MpMib.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsP2MpMib.EntityData.Leafs["cache-timer"] = types.YLeaf{"CacheTimer", mplsP2MpMib.CacheTimer}
    return &(mplsP2MpMib.EntityData)
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

    mplsTeExtStdMib.EntityData.Children = make(map[string]types.YChild)
    mplsTeExtStdMib.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsTeExtStdMib.EntityData.Leafs["cache-timer"] = types.YLeaf{"CacheTimer", mplsTeExtStdMib.CacheTimer}
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

    mplsTeExtMib.EntityData.Children = make(map[string]types.YChild)
    mplsTeExtMib.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsTeExtMib.EntityData.Leafs["cache-timer"] = types.YLeaf{"CacheTimer", mplsTeExtMib.CacheTimer}
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

    mplsFrrMib.EntityData.Children = make(map[string]types.YChild)
    mplsFrrMib.EntityData.Leafs = make(map[string]types.YLeaf)
    mplsFrrMib.EntityData.Leafs["cache-timer"] = types.YLeaf{"CacheTimer", mplsFrrMib.CacheTimer}
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

    cbQosmib.EntityData.Children = make(map[string]types.YChild)
    cbQosmib.EntityData.Children["cache"] = types.YChild{"Cache", &cbQosmib.Cache}
    cbQosmib.EntityData.Leafs = make(map[string]types.YLeaf)
    cbQosmib.EntityData.Leafs["member-interface-stats"] = types.YLeaf{"MemberInterfaceStats", cbQosmib.MemberInterfaceStats}
    cbQosmib.EntityData.Leafs["persist"] = types.YLeaf{"Persist", cbQosmib.Persist}
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

    cache.EntityData.Children = make(map[string]types.YChild)
    cache.EntityData.Leafs = make(map[string]types.YLeaf)
    cache.EntityData.Leafs["enable"] = types.YLeaf{"Enable", cache.Enable}
    cache.EntityData.Leafs["refresh-time"] = types.YLeaf{"RefreshTime", cache.RefreshTime}
    cache.EntityData.Leafs["service-policy-count"] = types.YLeaf{"ServicePolicyCount", cache.ServicePolicyCount}
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

    entityMib.EntityData.Children = make(map[string]types.YChild)
    entityMib.EntityData.Leafs = make(map[string]types.YLeaf)
    entityMib.EntityData.Leafs["entity-index-persistence"] = types.YLeaf{"EntityIndexPersistence", entityMib.EntityIndexPersistence}
    return &(entityMib.EntityData)
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

    interfaceMib.EntityData.Children = make(map[string]types.YChild)
    interfaceMib.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &interfaceMib.Interfaces}
    interfaceMib.EntityData.Children["notification"] = types.YChild{"Notification", &interfaceMib.Notification}
    interfaceMib.EntityData.Children["subsets"] = types.YChild{"Subsets", &interfaceMib.Subsets}
    interfaceMib.EntityData.Leafs = make(map[string]types.YLeaf)
    interfaceMib.EntityData.Leafs["internal-cache"] = types.YLeaf{"InternalCache", interfaceMib.InternalCache}
    interfaceMib.EntityData.Leafs["interface-alias-long"] = types.YLeaf{"InterfaceAliasLong", interfaceMib.InterfaceAliasLong}
    interfaceMib.EntityData.Leafs["ip-subscriber"] = types.YLeaf{"IpSubscriber", interfaceMib.IpSubscriber}
    interfaceMib.EntityData.Leafs["interface-index-persistence"] = types.YLeaf{"InterfaceIndexPersistence", interfaceMib.InterfaceIndexPersistence}
    interfaceMib.EntityData.Leafs["statistics-cache"] = types.YLeaf{"StatisticsCache", interfaceMib.StatisticsCache}
    return &(interfaceMib.EntityData)
}

// Mib_InterfaceMib_Interfaces
// Enter the SNMP interface configuration commands
type Mib_InterfaceMib_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface to configure. The type is slice of
    // Mib_InterfaceMib_Interfaces_Interface_.
    Interface_ []Mib_InterfaceMib_Interfaces_Interface
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

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Mib_InterfaceMib_Interfaces_Interface
// Interface to configure
type Mib_InterfaceMib_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the interface. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
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
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["link-up-down"] = types.YLeaf{"LinkUpDown", self.LinkUpDown}
    self.EntityData.Leafs["index-persistence"] = types.YLeaf{"IndexPersistence", self.IndexPersistence}
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

    notification.EntityData.Children = make(map[string]types.YChild)
    notification.EntityData.Leafs = make(map[string]types.YLeaf)
    notification.EntityData.Leafs["link-ietf"] = types.YLeaf{"LinkIetf", notification.LinkIetf}
    return &(notification.EntityData)
}

// Mib_InterfaceMib_Subsets
// Add configuration for an interface subset
type Mib_InterfaceMib_Subsets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subset priorityID to group ifNames based on Regular Expression. The type is
    // slice of Mib_InterfaceMib_Subsets_Subset.
    Subset []Mib_InterfaceMib_Subsets_Subset
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

    subsets.EntityData.Children = make(map[string]types.YChild)
    subsets.EntityData.Children["subset"] = types.YChild{"Subset", nil}
    for i := range subsets.Subset {
        subsets.EntityData.Children[types.GetSegmentPath(&subsets.Subset[i])] = types.YChild{"Subset", &subsets.Subset[i]}
    }
    subsets.EntityData.Leafs = make(map[string]types.YLeaf)
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
    subset.EntityData.SegmentPath = "subset" + "[subset-id='" + fmt.Sprintf("%v", subset.SubsetId) + "']"
    subset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subset.EntityData.Children = make(map[string]types.YChild)
    subset.EntityData.Children["link-up-down"] = types.YChild{"LinkUpDown", &subset.LinkUpDown}
    subset.EntityData.Leafs = make(map[string]types.YLeaf)
    subset.EntityData.Leafs["subset-id"] = types.YLeaf{"SubsetId", subset.SubsetId}
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

    linkUpDown.EntityData.Children = make(map[string]types.YChild)
    linkUpDown.EntityData.Leafs = make(map[string]types.YLeaf)
    linkUpDown.EntityData.Leafs["enable"] = types.YLeaf{"Enable", linkUpDown.Enable}
    linkUpDown.EntityData.Leafs["regular-expression"] = types.YLeaf{"RegularExpression", linkUpDown.RegularExpression}
    return &(linkUpDown.EntityData)
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

    subscriber.EntityData.Children = make(map[string]types.YChild)
    subscriber.EntityData.Children["threshold"] = types.YChild{"Threshold", &subscriber.Threshold}
    subscriber.EntityData.Leafs = make(map[string]types.YLeaf)
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

    threshold.EntityData.Children = make(map[string]types.YChild)
    threshold.EntityData.Children["delta"] = types.YChild{"Delta", &threshold.Delta}
    threshold.EntityData.Children["access-interface-sub"] = types.YChild{"AccessInterfaceSub", &threshold.AccessInterfaceSub}
    threshold.EntityData.Children["falling"] = types.YChild{"Falling", &threshold.Falling}
    threshold.EntityData.Children["rising"] = types.YChild{"Rising", &threshold.Rising}
    threshold.EntityData.Leafs = make(map[string]types.YLeaf)
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

    delta.EntityData.Children = make(map[string]types.YChild)
    delta.EntityData.Children["evaluation"] = types.YChild{"Evaluation", &delta.Evaluation}
    delta.EntityData.Children["percent"] = types.YChild{"Percent", &delta.Percent}
    delta.EntityData.Leafs = make(map[string]types.YLeaf)
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

    evaluation.EntityData.Children = make(map[string]types.YChild)
    evaluation.EntityData.Children["access-interfaces"] = types.YChild{"AccessInterfaces", &evaluation.AccessInterfaces}
    evaluation.EntityData.Children["nodes"] = types.YChild{"Nodes", &evaluation.Nodes}
    evaluation.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(evaluation.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces
// Table of AccessInterface
type Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access interface. The type is slice of
    // Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface.
    AccessInterface []Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface
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

    accessInterfaces.EntityData.Children = make(map[string]types.YChild)
    accessInterfaces.EntityData.Children["access-interface"] = types.YChild{"AccessInterface", nil}
    for i := range accessInterfaces.AccessInterface {
        accessInterfaces.EntityData.Children[types.GetSegmentPath(&accessInterfaces.AccessInterface[i])] = types.YChild{"AccessInterface", &accessInterfaces.AccessInterface[i]}
    }
    accessInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accessInterfaces.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface
// Access interface
type Mib_Subscriber_Threshold_Delta_Evaluation_AccessInterfaces_AccessInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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
    accessInterface.EntityData.SegmentPath = "access-interface" + "[interface-name='" + fmt.Sprintf("%v", accessInterface.InterfaceName) + "']"
    accessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterface.EntityData.Children = make(map[string]types.YChild)
    accessInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    accessInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", accessInterface.InterfaceName}
    accessInterface.EntityData.Leafs["session-count"] = types.YLeaf{"SessionCount", accessInterface.SessionCount}
    accessInterface.EntityData.Leafs["interval"] = types.YLeaf{"Interval", accessInterface.Interval}
    return &(accessInterface.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Evaluation_Nodes
// Table of Node
type Mib_Subscriber_Threshold_Delta_Evaluation_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rising node level. The type is slice of
    // Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node.
    Node []Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node
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

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node
// Rising node level
type Mib_Subscriber_Threshold_Delta_Evaluation_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. location. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    node.EntityData.Leafs["session-count"] = types.YLeaf{"SessionCount", node.SessionCount}
    node.EntityData.Leafs["interval"] = types.YLeaf{"Interval", node.Interval}
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

    percent.EntityData.Children = make(map[string]types.YChild)
    percent.EntityData.Children["access-interfaces"] = types.YChild{"AccessInterfaces", &percent.AccessInterfaces}
    percent.EntityData.Children["nodes"] = types.YChild{"Nodes", &percent.Nodes}
    percent.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(percent.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces
// Table of AccessInterface
type Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access interface. The type is slice of
    // Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface.
    AccessInterface []Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface
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

    accessInterfaces.EntityData.Children = make(map[string]types.YChild)
    accessInterfaces.EntityData.Children["access-interface"] = types.YChild{"AccessInterface", nil}
    for i := range accessInterfaces.AccessInterface {
        accessInterfaces.EntityData.Children[types.GetSegmentPath(&accessInterfaces.AccessInterface[i])] = types.YChild{"AccessInterface", &accessInterfaces.AccessInterface[i]}
    }
    accessInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accessInterfaces.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface
// Access interface
type Mib_Subscriber_Threshold_Delta_Percent_AccessInterfaces_AccessInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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
    accessInterface.EntityData.SegmentPath = "access-interface" + "[interface-name='" + fmt.Sprintf("%v", accessInterface.InterfaceName) + "']"
    accessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterface.EntityData.Children = make(map[string]types.YChild)
    accessInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    accessInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", accessInterface.InterfaceName}
    accessInterface.EntityData.Leafs["session-count"] = types.YLeaf{"SessionCount", accessInterface.SessionCount}
    accessInterface.EntityData.Leafs["interval"] = types.YLeaf{"Interval", accessInterface.Interval}
    return &(accessInterface.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Percent_Nodes
// Table of Node
type Mib_Subscriber_Threshold_Delta_Percent_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rising node level. The type is slice of
    // Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node.
    Node []Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node
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

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node
// Rising node level
type Mib_Subscriber_Threshold_Delta_Percent_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. location. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    node.EntityData.Leafs["session-count"] = types.YLeaf{"SessionCount", node.SessionCount}
    node.EntityData.Leafs["interval"] = types.YLeaf{"Interval", node.Interval}
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

    accessInterfaceSub.EntityData.Children = make(map[string]types.YChild)
    accessInterfaceSub.EntityData.Children["subsets"] = types.YChild{"Subsets", &accessInterfaceSub.Subsets}
    accessInterfaceSub.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accessInterfaceSub.EntityData)
}

// Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets
// Table of Subset
type Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Subset command. The type is slice of
    // Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset.
    Subset []Mib_Subscriber_Threshold_AccessInterfaceSub_Subsets_Subset
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

    subsets.EntityData.Children = make(map[string]types.YChild)
    subsets.EntityData.Children["subset"] = types.YChild{"Subset", nil}
    for i := range subsets.Subset {
        subsets.EntityData.Children[types.GetSegmentPath(&subsets.Subset[i])] = types.YChild{"Subset", &subsets.Subset[i]}
    }
    subsets.EntityData.Leafs = make(map[string]types.YLeaf)
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
    subset.EntityData.SegmentPath = "subset" + "[subset-id='" + fmt.Sprintf("%v", subset.SubsetId) + "']"
    subset.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    subset.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    subset.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    subset.EntityData.Children = make(map[string]types.YChild)
    subset.EntityData.Children["regular-expression"] = types.YChild{"RegularExpression", &subset.RegularExpression}
    subset.EntityData.Leafs = make(map[string]types.YLeaf)
    subset.EntityData.Leafs["subset-id"] = types.YLeaf{"SubsetId", subset.SubsetId}
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

    regularExpression.EntityData.Children = make(map[string]types.YChild)
    regularExpression.EntityData.Children["notification"] = types.YChild{"Notification", &regularExpression.Notification}
    regularExpression.EntityData.Leafs = make(map[string]types.YLeaf)
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

    notification.EntityData.Children = make(map[string]types.YChild)
    notification.EntityData.Children["rising-falling"] = types.YChild{"RisingFalling", &notification.RisingFalling}
    notification.EntityData.Leafs = make(map[string]types.YLeaf)
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

    risingFalling.EntityData.Children = make(map[string]types.YChild)
    risingFalling.EntityData.Leafs = make(map[string]types.YLeaf)
    risingFalling.EntityData.Leafs["disable"] = types.YLeaf{"Disable", risingFalling.Disable}
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

    falling.EntityData.Children = make(map[string]types.YChild)
    falling.EntityData.Children["access-interfaces"] = types.YChild{"AccessInterfaces", &falling.AccessInterfaces}
    falling.EntityData.Children["nodes"] = types.YChild{"Nodes", &falling.Nodes}
    falling.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(falling.EntityData)
}

// Mib_Subscriber_Threshold_Falling_AccessInterfaces
// Table of AccessInterface
type Mib_Subscriber_Threshold_Falling_AccessInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access interface. The type is slice of
    // Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface.
    AccessInterface []Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface
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

    accessInterfaces.EntityData.Children = make(map[string]types.YChild)
    accessInterfaces.EntityData.Children["access-interface"] = types.YChild{"AccessInterface", nil}
    for i := range accessInterfaces.AccessInterface {
        accessInterfaces.EntityData.Children[types.GetSegmentPath(&accessInterfaces.AccessInterface[i])] = types.YChild{"AccessInterface", &accessInterfaces.AccessInterface[i]}
    }
    accessInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accessInterfaces.EntityData)
}

// Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface
// Access interface
type Mib_Subscriber_Threshold_Falling_AccessInterfaces_AccessInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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
    accessInterface.EntityData.SegmentPath = "access-interface" + "[interface-name='" + fmt.Sprintf("%v", accessInterface.InterfaceName) + "']"
    accessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterface.EntityData.Children = make(map[string]types.YChild)
    accessInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    accessInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", accessInterface.InterfaceName}
    accessInterface.EntityData.Leafs["session-count"] = types.YLeaf{"SessionCount", accessInterface.SessionCount}
    accessInterface.EntityData.Leafs["interval"] = types.YLeaf{"Interval", accessInterface.Interval}
    return &(accessInterface.EntityData)
}

// Mib_Subscriber_Threshold_Falling_Nodes
// Table of Node
type Mib_Subscriber_Threshold_Falling_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rising node level. The type is slice of
    // Mib_Subscriber_Threshold_Falling_Nodes_Node.
    Node []Mib_Subscriber_Threshold_Falling_Nodes_Node
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

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Mib_Subscriber_Threshold_Falling_Nodes_Node
// Rising node level
type Mib_Subscriber_Threshold_Falling_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. location. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    node.EntityData.Leafs["session-count"] = types.YLeaf{"SessionCount", node.SessionCount}
    node.EntityData.Leafs["interval"] = types.YLeaf{"Interval", node.Interval}
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

    rising.EntityData.Children = make(map[string]types.YChild)
    rising.EntityData.Children["access-interfaces"] = types.YChild{"AccessInterfaces", &rising.AccessInterfaces}
    rising.EntityData.Children["nodes"] = types.YChild{"Nodes", &rising.Nodes}
    rising.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(rising.EntityData)
}

// Mib_Subscriber_Threshold_Rising_AccessInterfaces
// Table of AccessInterface
type Mib_Subscriber_Threshold_Rising_AccessInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Access interface. The type is slice of
    // Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface.
    AccessInterface []Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface
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

    accessInterfaces.EntityData.Children = make(map[string]types.YChild)
    accessInterfaces.EntityData.Children["access-interface"] = types.YChild{"AccessInterface", nil}
    for i := range accessInterfaces.AccessInterface {
        accessInterfaces.EntityData.Children[types.GetSegmentPath(&accessInterfaces.AccessInterface[i])] = types.YChild{"AccessInterface", &accessInterfaces.AccessInterface[i]}
    }
    accessInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(accessInterfaces.EntityData)
}

// Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface
// Access interface
type Mib_Subscriber_Threshold_Rising_AccessInterfaces_AccessInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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
    accessInterface.EntityData.SegmentPath = "access-interface" + "[interface-name='" + fmt.Sprintf("%v", accessInterface.InterfaceName) + "']"
    accessInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    accessInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    accessInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    accessInterface.EntityData.Children = make(map[string]types.YChild)
    accessInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    accessInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", accessInterface.InterfaceName}
    accessInterface.EntityData.Leafs["session-count"] = types.YLeaf{"SessionCount", accessInterface.SessionCount}
    accessInterface.EntityData.Leafs["interval"] = types.YLeaf{"Interval", accessInterface.Interval}
    return &(accessInterface.EntityData)
}

// Mib_Subscriber_Threshold_Rising_Nodes
// Table of Node
type Mib_Subscriber_Threshold_Rising_Nodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rising node level. The type is slice of
    // Mib_Subscriber_Threshold_Rising_Nodes_Node.
    Node []Mib_Subscriber_Threshold_Rising_Nodes_Node
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

    nodes.EntityData.Children = make(map[string]types.YChild)
    nodes.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range nodes.Node {
        nodes.EntityData.Children[types.GetSegmentPath(&nodes.Node[i])] = types.YChild{"Node", &nodes.Node[i]}
    }
    nodes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(nodes.EntityData)
}

// Mib_Subscriber_Threshold_Rising_Nodes_Node
// Rising node level
type Mib_Subscriber_Threshold_Rising_Nodes_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. location. The type is string with pattern:
    // b'([a-zA-Z0-9_]*\\d+/){1,2}([a-zA-Z0-9_]*\\d+)'.
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
    node.EntityData.SegmentPath = "node" + "[node-name='" + fmt.Sprintf("%v", node.NodeName) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-name"] = types.YLeaf{"NodeName", node.NodeName}
    node.EntityData.Leafs["session-count"] = types.YLeaf{"SessionCount", node.SessionCount}
    node.EntityData.Leafs["interval"] = types.YLeaf{"Interval", node.Interval}
    return &(node.EntityData)
}

