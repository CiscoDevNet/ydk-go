// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-dhcpd package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ipv4-dhcpd: DHCP IPV4 configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_dhcpd_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_dhcpd_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-dhcpd-cfg ipv4-dhcpd}", reflect.TypeOf(Ipv4Dhcpd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd", reflect.TypeOf(Ipv4Dhcpd{}))
}

// Matchaction represents Matchaction
type Matchaction string

const (
    // Allow DHCP Discover
    Matchaction_allow Matchaction = "allow"

    // Drop DHCP Discover
    Matchaction_drop Matchaction = "drop"
)

// Dhcpv4AuthUsername represents Dhcpv4 auth username
type Dhcpv4AuthUsername string

const (
    // Authentication Username formating mac
    Dhcpv4AuthUsername_auth_username_mac Dhcpv4AuthUsername = "auth-username-mac"

    // Authentication Username formating giaddr
    Dhcpv4AuthUsername_auth_username_giaddr Dhcpv4AuthUsername = "auth-username-giaddr"
)

// Policy represents Policy
type Policy string

const (
    // Ignore the broadcast policy
    Policy_ignore Policy = "ignore"

    // Check for broadcast flag
    Policy_check Policy = "check"

    // Always Unicast the reply
    Policy_unicastalways Policy = "unicastalways"
)

// Matchoption represents Matchoption
type Matchoption string

const (
    // Match circuit id of option 82 Relay-agent
    // specific class
    Matchoption_circuitid Matchoption = "circuitid"

    // Match remote id of option 82 Relay-agent
    // specific class
    Matchoption_remoteid Matchoption = "remoteid"

    // Match option 60 vendor class id
    Matchoption_Y_60 Matchoption = "60"

    // Match option 77 user class
    Matchoption_Y_77 Matchoption = "77"

    // Match option 124 vendor-identifying vendor
    // class
    Matchoption_Y_124 Matchoption = "124"

    // Match option 125 vendor-indentifying
    // vendor-specific info
    Matchoption_Y_125 Matchoption = "125"
)

// BaseAction represents Base action
type BaseAction string

const (
    // Allow vendor specific DHCP Discover
    BaseAction_allow BaseAction = "allow"

    // Drop vendor specific DHCP Discover
    BaseAction_drop BaseAction = "drop"
)

// Dhcpv4LimitLease1 represents Dhcpv4 limit lease1
type Dhcpv4LimitLease1 string

const (
    // Interface
    Dhcpv4LimitLease1_interface_ Dhcpv4LimitLease1 = "interface"

    // Circuit ID
    Dhcpv4LimitLease1_circuit_id Dhcpv4LimitLease1 = "circuit-id"

    // Remote ID
    Dhcpv4LimitLease1_remote_id Dhcpv4LimitLease1 = "remote-id"

    // Circuit ID Remote ID
    Dhcpv4LimitLease1_circuit_id_remote_id Dhcpv4LimitLease1 = "circuit-id-remote-id"
)

// Ipv4dhcpdLayer represents Ipv4dhcpd layer
type Ipv4dhcpdLayer string

const (
    // Layer2
    Ipv4dhcpdLayer_layer2 Ipv4dhcpdLayer = "layer2"

    // Layer3
    Ipv4dhcpdLayer_layer3 Ipv4dhcpdLayer = "layer3"
)

// Ipv4dhcpdGiaddrPolicy represents Ipv4dhcpd giaddr policy
type Ipv4dhcpdGiaddrPolicy string

const (
    // Giaddr Policy Keep
    Ipv4dhcpdGiaddrPolicy_giaddr_policy_keep Ipv4dhcpdGiaddrPolicy = "giaddr-policy-keep"
)

// Ipv4dhcpdMode represents Ipv4dhcpd mode
type Ipv4dhcpdMode string

const (
    // Base
    Ipv4dhcpdMode_base Ipv4dhcpdMode = "base"

    // Relay
    Ipv4dhcpdMode_relay Ipv4dhcpdMode = "relay"

    // Snoop
    Ipv4dhcpdMode_snoop Ipv4dhcpdMode = "snoop"

    // Server
    Ipv4dhcpdMode_server Ipv4dhcpdMode = "server"

    // Proxy
    Ipv4dhcpdMode_proxy Ipv4dhcpdMode = "proxy"

    // Base2
    Ipv4dhcpdMode_base2 Ipv4dhcpdMode = "base2"
)

// Ipv4dhcpdFmtSpecifier represents Ipv4dhcpd fmt specifier
type Ipv4dhcpdFmtSpecifier string

const (
    // Physical chassis
    Ipv4dhcpdFmtSpecifier_physical_chassis Ipv4dhcpdFmtSpecifier = "physical-chassis"

    // Physical slot
    Ipv4dhcpdFmtSpecifier_physical_slot Ipv4dhcpdFmtSpecifier = "physical-slot"

    // Physical sub-slot
    Ipv4dhcpdFmtSpecifier_physical_sub_slot Ipv4dhcpdFmtSpecifier = "physical-sub-slot"

    // Physical port
    Ipv4dhcpdFmtSpecifier_physical_port Ipv4dhcpdFmtSpecifier = "physical-port"

    // Physical sub-port
    Ipv4dhcpdFmtSpecifier_physical_sub_port Ipv4dhcpdFmtSpecifier = "physical-sub-port"

    // Inner VLAN ID
    Ipv4dhcpdFmtSpecifier_inner_vlan_id Ipv4dhcpdFmtSpecifier = "inner-vlan-id"

    // Outer VLAN ID
    Ipv4dhcpdFmtSpecifier_outer_vlan_id Ipv4dhcpdFmtSpecifier = "outer-vlan-id"

    // L2 Interface
    Ipv4dhcpdFmtSpecifier_l2_interface Ipv4dhcpdFmtSpecifier = "l2-interface"

    // L3 Interface
    Ipv4dhcpdFmtSpecifier_l3_interface Ipv4dhcpdFmtSpecifier = "l3-interface"

    // Hostname
    Ipv4dhcpdFmtSpecifier_host_name Ipv4dhcpdFmtSpecifier = "host-name"
)

// Ipv4dhcpdRelayInfoOptionPolicy represents Ipv4dhcpd relay info option policy
type Ipv4dhcpdRelayInfoOptionPolicy string

const (
    // Replace
    Ipv4dhcpdRelayInfoOptionPolicy_replace Ipv4dhcpdRelayInfoOptionPolicy = "replace"

    // Keep
    Ipv4dhcpdRelayInfoOptionPolicy_keep Ipv4dhcpdRelayInfoOptionPolicy = "keep"

    // Drop
    Ipv4dhcpdRelayInfoOptionPolicy_drop Ipv4dhcpdRelayInfoOptionPolicy = "drop"

    // Encapsulate
    Ipv4dhcpdRelayInfoOptionPolicy_encapsulate Ipv4dhcpdRelayInfoOptionPolicy = "encapsulate"
)

// MacMismatchAction represents Mac mismatch action
type MacMismatchAction string

const (
    // Forward
    MacMismatchAction_forward MacMismatchAction = "forward"

    // Drop
    MacMismatchAction_drop MacMismatchAction = "drop"
)

// Ipv4dhcpdBroadcastFlagPolicy represents Ipv4dhcpd broadcast flag policy
type Ipv4dhcpdBroadcastFlagPolicy string

const (
    // Ignore
    Ipv4dhcpdBroadcastFlagPolicy_ignore Ipv4dhcpdBroadcastFlagPolicy = "ignore"

    // check
    Ipv4dhcpdBroadcastFlagPolicy_check Ipv4dhcpdBroadcastFlagPolicy = "check"

    // Unicast always
    Ipv4dhcpdBroadcastFlagPolicy_unicast_always Ipv4dhcpdBroadcastFlagPolicy = "unicast-always"
)

// Ipv4dhcpdFmt represents Ipv4dhcpd fmt
type Ipv4dhcpdFmt string

const (
    // Not a Format String
    Ipv4dhcpdFmt_no_format Ipv4dhcpdFmt = "no-format"

    // Hex Format String
    Ipv4dhcpdFmt_hex Ipv4dhcpdFmt = "hex"

    // Ascii Format String
    Ipv4dhcpdFmt_ascii Ipv4dhcpdFmt = "ascii"

    // Extended Format String
    Ipv4dhcpdFmt_extended Ipv4dhcpdFmt = "extended"
)

// Ipv4dhcpdRelayInfoOptionvpnMode represents Ipv4dhcpd relay info optionvpn mode
type Ipv4dhcpdRelayInfoOptionvpnMode string

const (
    // RFC
    Ipv4dhcpdRelayInfoOptionvpnMode_rfc Ipv4dhcpdRelayInfoOptionvpnMode = "rfc"

    // Cisco
    Ipv4dhcpdRelayInfoOptionvpnMode_cisco Ipv4dhcpdRelayInfoOptionvpnMode = "cisco"
)

// ProxyAction represents Proxy action
type ProxyAction string

const (
    // Allow vendor specific DHCP Discover
    ProxyAction_allow ProxyAction = "allow"

    // Drop vendor specific DHCP Discover
    ProxyAction_drop ProxyAction = "drop"

    // Relay vendor-id specific DHCP packets unaltered
    ProxyAction_relay ProxyAction = "relay"
)

// LeaseLimitValue represents Lease limit value
type LeaseLimitValue string

const (
    // Insert the limit lease type interface
    LeaseLimitValue_per_interface LeaseLimitValue = "per-interface"

    // Insert the limit lease type circuit-id
    LeaseLimitValue_per_circuit_id LeaseLimitValue = "per-circuit-id"

    // Insert the limit lease type remote-id
    LeaseLimitValue_per_remote_id LeaseLimitValue = "per-remote-id"
)

// Dhcpv4MatchOption represents Dhcpv4 match option
type Dhcpv4MatchOption string

const (
    // Vendor class ID
    Dhcpv4MatchOption_Y_60__FWD_SLASH__60 Dhcpv4MatchOption = "60/60"

    // 77 User class
    Dhcpv4MatchOption_Y_77__FWD_SLASH__77 Dhcpv4MatchOption = "77/77"

    // Vendor identifying class
    Dhcpv4MatchOption_Y_124__FWD_SLASH__124 Dhcpv4MatchOption = "124/124"

    // Vendor specific information
    Dhcpv4MatchOption_Y_125__FWD_SLASH__125 Dhcpv4MatchOption = "125/125"
)

// Ipv4dhcpdRelayInfoOptionAuthenticate represents Ipv4dhcpd relay info option authenticate
type Ipv4dhcpdRelayInfoOptionAuthenticate string

const (
    // Received
    Ipv4dhcpdRelayInfoOptionAuthenticate_received Ipv4dhcpdRelayInfoOptionAuthenticate = "received"

    // Inserted
    Ipv4dhcpdRelayInfoOptionAuthenticate_inserted Ipv4dhcpdRelayInfoOptionAuthenticate = "inserted"
)

// Ipv4dhcpdRelayGiaddrPolicy represents Ipv4dhcpd relay giaddr policy
type Ipv4dhcpdRelayGiaddrPolicy string

const (
    // Replace
    Ipv4dhcpdRelayGiaddrPolicy_replace Ipv4dhcpdRelayGiaddrPolicy = "replace"

    // Drop
    Ipv4dhcpdRelayGiaddrPolicy_drop Ipv4dhcpdRelayGiaddrPolicy = "drop"
)

// Ipv4Dhcpd
// DHCP IPV4 configuration
// This type is a presence type.
type Ipv4Dhcpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // DHCP IPV4 configuration. The type is interface{}. This attribute is
    // mandatory.
    Enable interface{}

    // Configure outer cos values for dhcp packets. The type is interface{} with
    // range: 0..7.
    OuterCos interface{}

    // For BNG session, allow client id change for a client MAC. The type is
    // interface{}.
    AllowClientIdChange interface{}

    // Configure inner cos values for dhcp packets. The type is interface{} with
    // range: 0..7.
    InnerCos interface{}

    // VRF Table.
    Vrfs Ipv4Dhcpd_Vrfs

    // DHCP IPV4 Profile Table.
    Profiles Ipv4Dhcpd_Profiles

    // Enable DHCP binding database storage to file system.
    Database Ipv4Dhcpd_Database

    // DHCP IPV4 Interface Table.
    Interfaces Ipv4Dhcpd_Interfaces

    // Allow Duplicate MAC Address.
    DuplicateMacAllowed Ipv4Dhcpd_DuplicateMacAllowed

    // Rate limit ingress packets.
    RateLimit Ipv4Dhcpd_RateLimit
}

func (ipv4Dhcpd *Ipv4Dhcpd) GetEntityData() *types.CommonEntityData {
    ipv4Dhcpd.EntityData.YFilter = ipv4Dhcpd.YFilter
    ipv4Dhcpd.EntityData.YangName = "ipv4-dhcpd"
    ipv4Dhcpd.EntityData.BundleName = "cisco_ios_xr"
    ipv4Dhcpd.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-dhcpd-cfg"
    ipv4Dhcpd.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd"
    ipv4Dhcpd.EntityData.AbsolutePath = ipv4Dhcpd.EntityData.SegmentPath
    ipv4Dhcpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Dhcpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Dhcpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Dhcpd.EntityData.Children = types.NewOrderedMap()
    ipv4Dhcpd.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &ipv4Dhcpd.Vrfs})
    ipv4Dhcpd.EntityData.Children.Append("profiles", types.YChild{"Profiles", &ipv4Dhcpd.Profiles})
    ipv4Dhcpd.EntityData.Children.Append("database", types.YChild{"Database", &ipv4Dhcpd.Database})
    ipv4Dhcpd.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &ipv4Dhcpd.Interfaces})
    ipv4Dhcpd.EntityData.Children.Append("duplicate-mac-allowed", types.YChild{"DuplicateMacAllowed", &ipv4Dhcpd.DuplicateMacAllowed})
    ipv4Dhcpd.EntityData.Children.Append("rate-limit", types.YChild{"RateLimit", &ipv4Dhcpd.RateLimit})
    ipv4Dhcpd.EntityData.Leafs = types.NewOrderedMap()
    ipv4Dhcpd.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ipv4Dhcpd.Enable})
    ipv4Dhcpd.EntityData.Leafs.Append("outer-cos", types.YLeaf{"OuterCos", ipv4Dhcpd.OuterCos})
    ipv4Dhcpd.EntityData.Leafs.Append("allow-client-id-change", types.YLeaf{"AllowClientIdChange", ipv4Dhcpd.AllowClientIdChange})
    ipv4Dhcpd.EntityData.Leafs.Append("inner-cos", types.YLeaf{"InnerCos", ipv4Dhcpd.InnerCos})

    ipv4Dhcpd.EntityData.YListKeys = []string {}

    return &(ipv4Dhcpd.EntityData)
}

// Ipv4Dhcpd_Vrfs
// VRF Table
type Ipv4Dhcpd_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF table. The type is slice of Ipv4Dhcpd_Vrfs_Vrf.
    Vrf []*Ipv4Dhcpd_Vrfs_Vrf
}

func (vrfs *Ipv4Dhcpd_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "ipv4-dhcpd"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/" + vrfs.EntityData.SegmentPath
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

// Ipv4Dhcpd_Vrfs_Vrf
// VRF table
type Ipv4Dhcpd_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Profile name and mode.
    Profile Ipv4Dhcpd_Vrfs_Vrf_Profile
}

func (vrf *Ipv4Dhcpd_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("profile", types.YChild{"Profile", &vrf.Profile})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Ipv4Dhcpd_Vrfs_Vrf_Profile
// Profile name and mode
// This type is a presence type.
type Ipv4Dhcpd_Vrfs_Vrf_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Profile name. The type is string. This attribute is mandatory.
    VrfProfileName interface{}

    // Dhcp mode. The type is Ipv4dhcpdMode. This attribute is mandatory.
    Mode interface{}
}

func (profile *Ipv4Dhcpd_Vrfs_Vrf_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "vrf"
    profile.EntityData.SegmentPath = "profile"
    profile.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/vrfs/vrf/" + profile.EntityData.SegmentPath
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = types.NewOrderedMap()
    profile.EntityData.Leafs = types.NewOrderedMap()
    profile.EntityData.Leafs.Append("vrf-profile-name", types.YLeaf{"VrfProfileName", profile.VrfProfileName})
    profile.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", profile.Mode})

    profile.EntityData.YListKeys = []string {}

    return &(profile.EntityData)
}

// Ipv4Dhcpd_Profiles
// DHCP IPV4 Profile Table
type Ipv4Dhcpd_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP IPV4 Profile. The type is slice of Ipv4Dhcpd_Profiles_Profile.
    Profile []*Ipv4Dhcpd_Profiles_Profile
}

func (profiles *Ipv4Dhcpd_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "ipv4-dhcpd"
    profiles.EntityData.SegmentPath = "profiles"
    profiles.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/" + profiles.EntityData.SegmentPath
    profiles.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profiles.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profiles.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profiles.EntityData.Children = types.NewOrderedMap()
    profiles.EntityData.Children.Append("profile", types.YChild{"Profile", nil})
    for i := range profiles.Profile {
        profiles.EntityData.Children.Append(types.GetSegmentPath(profiles.Profile[i]), types.YChild{"Profile", profiles.Profile[i]})
    }
    profiles.EntityData.Leafs = types.NewOrderedMap()

    profiles.EntityData.YListKeys = []string {}

    return &(profiles.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile
// DHCP IPV4 Profile
type Ipv4Dhcpd_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Profile Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // DHCP IPV4 Profile modes.
    Modes Ipv4Dhcpd_Profiles_Profile_Modes
}

func (profile *Ipv4Dhcpd_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + types.AddKeyToken(profile.ProfileName, "profile-name")
    profile.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/" + profile.EntityData.SegmentPath
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = types.NewOrderedMap()
    profile.EntityData.Children.Append("modes", types.YChild{"Modes", &profile.Modes})
    profile.EntityData.Leafs = types.NewOrderedMap()
    profile.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", profile.ProfileName})

    profile.EntityData.YListKeys = []string {"ProfileName"}

    return &(profile.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes
// DHCP IPV4 Profile modes
type Ipv4Dhcpd_Profiles_Profile_Modes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP IPV4 Profile mode. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode.
    Mode []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode
}

func (modes *Ipv4Dhcpd_Profiles_Profile_Modes) GetEntityData() *types.CommonEntityData {
    modes.EntityData.YFilter = modes.YFilter
    modes.EntityData.YangName = "modes"
    modes.EntityData.BundleName = "cisco_ios_xr"
    modes.EntityData.ParentYangName = "profile"
    modes.EntityData.SegmentPath = "modes"
    modes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/" + modes.EntityData.SegmentPath
    modes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    modes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    modes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    modes.EntityData.Children = types.NewOrderedMap()
    modes.EntityData.Children.Append("mode", types.YChild{"Mode", nil})
    for i := range modes.Mode {
        modes.EntityData.Children.Append(types.GetSegmentPath(modes.Mode[i]), types.YChild{"Mode", modes.Mode[i]})
    }
    modes.EntityData.Leafs = types.NewOrderedMap()

    modes.EntityData.YListKeys = []string {}

    return &(modes.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode
// DHCP IPV4 Profile mode
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. DHCP IPV4 Profile mode. The type is Ipv4dhcpdMode.
    Mode interface{}

    // Enable the DHCP IPV4 Profile mode. The type is interface{}.
    Enable interface{}

    // DHCP Snoop profile.
    Snoop Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop

    // DHCP Base Profile.
    Base Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base

    // DHCP Server profile.
    Server Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server

    // DHCP Relay profile.
    Relay Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay

    // DHCP proxy profile.
    Proxy Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy
}

func (mode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode) GetEntityData() *types.CommonEntityData {
    mode.EntityData.YFilter = mode.YFilter
    mode.EntityData.YangName = "mode"
    mode.EntityData.BundleName = "cisco_ios_xr"
    mode.EntityData.ParentYangName = "modes"
    mode.EntityData.SegmentPath = "mode" + types.AddKeyToken(mode.Mode, "mode")
    mode.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/" + mode.EntityData.SegmentPath
    mode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mode.EntityData.Children = types.NewOrderedMap()
    mode.EntityData.Children.Append("snoop", types.YChild{"Snoop", &mode.Snoop})
    mode.EntityData.Children.Append("base", types.YChild{"Base", &mode.Base})
    mode.EntityData.Children.Append("server", types.YChild{"Server", &mode.Server})
    mode.EntityData.Children.Append("relay", types.YChild{"Relay", &mode.Relay})
    mode.EntityData.Children.Append("proxy", types.YChild{"Proxy", &mode.Proxy})
    mode.EntityData.Leafs = types.NewOrderedMap()
    mode.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", mode.Mode})
    mode.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", mode.Enable})

    mode.EntityData.YListKeys = []string {"Mode"}

    return &(mode.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop
// DHCP Snoop profile
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trusted sources. The type is interface{}.
    Trusted interface{}

    // DHCP Snoop profile.
    RelayInformationOption Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption
}

func (snoop *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop) GetEntityData() *types.CommonEntityData {
    snoop.EntityData.YFilter = snoop.YFilter
    snoop.EntityData.YangName = "snoop"
    snoop.EntityData.BundleName = "cisco_ios_xr"
    snoop.EntityData.ParentYangName = "mode"
    snoop.EntityData.SegmentPath = "snoop"
    snoop.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/" + snoop.EntityData.SegmentPath
    snoop.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snoop.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snoop.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snoop.EntityData.Children = types.NewOrderedMap()
    snoop.EntityData.Children.Append("relay-information-option", types.YChild{"RelayInformationOption", &snoop.RelayInformationOption})
    snoop.EntityData.Leafs = types.NewOrderedMap()
    snoop.EntityData.Leafs.Append("trusted", types.YLeaf{"Trusted", snoop.Trusted})

    snoop.EntityData.YListKeys = []string {}

    return &(snoop.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption
// DHCP Snoop profile
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Insert Relay Agent Information circuit ID and remote ID suboptions in
    // client request. The type is interface{}.
    Insert interface{}

    // Forward untrusted packets. The type is interface{}.
    AllowUntrusted interface{}

    // Relay information option policy. The type is
    // Ipv4dhcpdRelayInfoOptionPolicy.
    Policy interface{}

    // Enter remote-id value.
    RemoteId Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId
}

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption) GetEntityData() *types.CommonEntityData {
    relayInformationOption.EntityData.YFilter = relayInformationOption.YFilter
    relayInformationOption.EntityData.YangName = "relay-information-option"
    relayInformationOption.EntityData.BundleName = "cisco_ios_xr"
    relayInformationOption.EntityData.ParentYangName = "snoop"
    relayInformationOption.EntityData.SegmentPath = "relay-information-option"
    relayInformationOption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/snoop/" + relayInformationOption.EntityData.SegmentPath
    relayInformationOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayInformationOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayInformationOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayInformationOption.EntityData.Children = types.NewOrderedMap()
    relayInformationOption.EntityData.Children.Append("remote-id", types.YChild{"RemoteId", &relayInformationOption.RemoteId})
    relayInformationOption.EntityData.Leafs = types.NewOrderedMap()
    relayInformationOption.EntityData.Leafs.Append("insert", types.YLeaf{"Insert", relayInformationOption.Insert})
    relayInformationOption.EntityData.Leafs.Append("allow-untrusted", types.YLeaf{"AllowUntrusted", relayInformationOption.AllowUntrusted})
    relayInformationOption.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", relayInformationOption.Policy})

    relayInformationOption.EntityData.YListKeys = []string {}

    return &(relayInformationOption.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId
// Enter remote-id value
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Format type, 1. Hex 2. ASCII. The type is interface{} with range: 1..2.
    FormatType interface{}

    // Enter remote-id value. The type is string.
    RemoteIdValue interface{}
}

func (remoteId *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId) GetEntityData() *types.CommonEntityData {
    remoteId.EntityData.YFilter = remoteId.YFilter
    remoteId.EntityData.YangName = "remote-id"
    remoteId.EntityData.BundleName = "cisco_ios_xr"
    remoteId.EntityData.ParentYangName = "relay-information-option"
    remoteId.EntityData.SegmentPath = "remote-id"
    remoteId.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/snoop/relay-information-option/" + remoteId.EntityData.SegmentPath
    remoteId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteId.EntityData.Children = types.NewOrderedMap()
    remoteId.EntityData.Leafs = types.NewOrderedMap()
    remoteId.EntityData.Leafs.Append("format-type", types.YLeaf{"FormatType", remoteId.FormatType})
    remoteId.EntityData.Leafs.Append("remote-id-value", types.YLeaf{"RemoteIdValue", remoteId.RemoteIdValue})

    remoteId.EntityData.YListKeys = []string {}

    return &(remoteId.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base
// DHCP Base Profile
// This type is a presence type.
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Enable the DHCP IPv4 Base Profile. The type is interface{}. This attribute
    // is mandatory.
    Enable interface{}

    // Enable the default profile.
    DefaultProfile Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile

    // Insert match keyword.
    Match Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match

    // Insert Relay Agent Information circuit ID and remote ID suboptions in
    // client request.
    BaseRelayOpt Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt

    // Insert match keyword.
    BaseMatch Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch
}

func (base *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base) GetEntityData() *types.CommonEntityData {
    base.EntityData.YFilter = base.YFilter
    base.EntityData.YangName = "base"
    base.EntityData.BundleName = "cisco_ios_xr"
    base.EntityData.ParentYangName = "mode"
    base.EntityData.SegmentPath = "base"
    base.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/" + base.EntityData.SegmentPath
    base.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    base.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    base.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    base.EntityData.Children = types.NewOrderedMap()
    base.EntityData.Children.Append("default-profile", types.YChild{"DefaultProfile", &base.DefaultProfile})
    base.EntityData.Children.Append("match", types.YChild{"Match", &base.Match})
    base.EntityData.Children.Append("base-relay-opt", types.YChild{"BaseRelayOpt", &base.BaseRelayOpt})
    base.EntityData.Children.Append("base-match", types.YChild{"BaseMatch", &base.BaseMatch})
    base.EntityData.Leafs = types.NewOrderedMap()
    base.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", base.Enable})

    base.EntityData.YListKeys = []string {}

    return &(base.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile
// Enable the default profile
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile name. The type is string with length: 1..64.
    ProfileName interface{}

    // none. The type is interface{} with range: 0..4294967295.
    ProfileMode interface{}
}

func (defaultProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile) GetEntityData() *types.CommonEntityData {
    defaultProfile.EntityData.YFilter = defaultProfile.YFilter
    defaultProfile.EntityData.YangName = "default-profile"
    defaultProfile.EntityData.BundleName = "cisco_ios_xr"
    defaultProfile.EntityData.ParentYangName = "base"
    defaultProfile.EntityData.SegmentPath = "default-profile"
    defaultProfile.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/base/" + defaultProfile.EntityData.SegmentPath
    defaultProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultProfile.EntityData.Children = types.NewOrderedMap()
    defaultProfile.EntityData.Leafs = types.NewOrderedMap()
    defaultProfile.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", defaultProfile.ProfileName})
    defaultProfile.EntityData.Leafs.Append("profile-mode", types.YLeaf{"ProfileMode", defaultProfile.ProfileMode})

    defaultProfile.EntityData.YListKeys = []string {}

    return &(defaultProfile.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match
// Insert match keyword
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Option.
    OptionFilters Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters

    // Table of Option.
    DefOptions Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match) GetEntityData() *types.CommonEntityData {
    match.EntityData.YFilter = match.YFilter
    match.EntityData.YangName = "match"
    match.EntityData.BundleName = "cisco_ios_xr"
    match.EntityData.ParentYangName = "base"
    match.EntityData.SegmentPath = "match"
    match.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/base/" + match.EntityData.SegmentPath
    match.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    match.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    match.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    match.EntityData.Children = types.NewOrderedMap()
    match.EntityData.Children.Append("option-filters", types.YChild{"OptionFilters", &match.OptionFilters})
    match.EntityData.Children.Append("def-options", types.YChild{"DefOptions", &match.DefOptions})
    match.EntityData.Leafs = types.NewOrderedMap()

    match.EntityData.YListKeys = []string {}

    return &(match.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters
// Table of Option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify match option. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter.
    OptionFilter []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter
}

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters) GetEntityData() *types.CommonEntityData {
    optionFilters.EntityData.YFilter = optionFilters.YFilter
    optionFilters.EntityData.YangName = "option-filters"
    optionFilters.EntityData.BundleName = "cisco_ios_xr"
    optionFilters.EntityData.ParentYangName = "match"
    optionFilters.EntityData.SegmentPath = "option-filters"
    optionFilters.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/base/match/" + optionFilters.EntityData.SegmentPath
    optionFilters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optionFilters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optionFilters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optionFilters.EntityData.Children = types.NewOrderedMap()
    optionFilters.EntityData.Children.Append("option-filter", types.YChild{"OptionFilter", nil})
    for i := range optionFilters.OptionFilter {
        optionFilters.EntityData.Children.Append(types.GetSegmentPath(optionFilters.OptionFilter[i]), types.YChild{"OptionFilter", optionFilters.OptionFilter[i]})
    }
    optionFilters.EntityData.Leafs = types.NewOrderedMap()

    optionFilters.EntityData.YListKeys = []string {}

    return &(optionFilters.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter
// Specify match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Match option 60. The type is interface{} with
    // range: 0..4294967295.
    Matchoption interface{}

    // This attribute is a key. Enter hex pattern string. The type is string with
    // length: 1..64.
    Pattern interface{}

    // This attribute is a key. Set constant integer. The type is interface{} with
    // range: 0..4294967295.
    Format interface{}

    // Vendor action. The type is BaseAction.
    OptionAction interface{}
}

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter) GetEntityData() *types.CommonEntityData {
    optionFilter.EntityData.YFilter = optionFilter.YFilter
    optionFilter.EntityData.YangName = "option-filter"
    optionFilter.EntityData.BundleName = "cisco_ios_xr"
    optionFilter.EntityData.ParentYangName = "option-filters"
    optionFilter.EntityData.SegmentPath = "option-filter" + types.AddKeyToken(optionFilter.Matchoption, "matchoption") + types.AddKeyToken(optionFilter.Pattern, "pattern") + types.AddKeyToken(optionFilter.Format, "format")
    optionFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/base/match/option-filters/" + optionFilter.EntityData.SegmentPath
    optionFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optionFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optionFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optionFilter.EntityData.Children = types.NewOrderedMap()
    optionFilter.EntityData.Leafs = types.NewOrderedMap()
    optionFilter.EntityData.Leafs.Append("matchoption", types.YLeaf{"Matchoption", optionFilter.Matchoption})
    optionFilter.EntityData.Leafs.Append("pattern", types.YLeaf{"Pattern", optionFilter.Pattern})
    optionFilter.EntityData.Leafs.Append("format", types.YLeaf{"Format", optionFilter.Format})
    optionFilter.EntityData.Leafs.Append("option-action", types.YLeaf{"OptionAction", optionFilter.OptionAction})

    optionFilter.EntityData.YListKeys = []string {"Matchoption", "Pattern", "Format"}

    return &(optionFilter.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions
// Table of Option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify match option. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption.
    DefOption []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption
}

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions) GetEntityData() *types.CommonEntityData {
    defOptions.EntityData.YFilter = defOptions.YFilter
    defOptions.EntityData.YangName = "def-options"
    defOptions.EntityData.BundleName = "cisco_ios_xr"
    defOptions.EntityData.ParentYangName = "match"
    defOptions.EntityData.SegmentPath = "def-options"
    defOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/base/match/" + defOptions.EntityData.SegmentPath
    defOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defOptions.EntityData.Children = types.NewOrderedMap()
    defOptions.EntityData.Children.Append("def-option", types.YChild{"DefOption", nil})
    for i := range defOptions.DefOption {
        defOptions.EntityData.Children.Append(types.GetSegmentPath(defOptions.DefOption[i]), types.YChild{"DefOption", defOptions.DefOption[i]})
    }
    defOptions.EntityData.Leafs = types.NewOrderedMap()

    defOptions.EntityData.YListKeys = []string {}

    return &(defOptions.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption
// Specify match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Match option 60. The type is interface{} with
    // range: 0..4294967295.
    DefMatchoption interface{}

    // Vendor action. The type is BaseAction. This attribute is mandatory.
    DefMatchaction interface{}
}

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption) GetEntityData() *types.CommonEntityData {
    defOption.EntityData.YFilter = defOption.YFilter
    defOption.EntityData.YangName = "def-option"
    defOption.EntityData.BundleName = "cisco_ios_xr"
    defOption.EntityData.ParentYangName = "def-options"
    defOption.EntityData.SegmentPath = "def-option" + types.AddKeyToken(defOption.DefMatchoption, "def-matchoption")
    defOption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/base/match/def-options/" + defOption.EntityData.SegmentPath
    defOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defOption.EntityData.Children = types.NewOrderedMap()
    defOption.EntityData.Leafs = types.NewOrderedMap()
    defOption.EntityData.Leafs.Append("def-matchoption", types.YLeaf{"DefMatchoption", defOption.DefMatchoption})
    defOption.EntityData.Leafs.Append("def-matchaction", types.YLeaf{"DefMatchaction", defOption.DefMatchaction})

    defOption.EntityData.YListKeys = []string {"DefMatchoption"}

    return &(defOption.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt
// Insert Relay Agent Information circuit ID
// and remote ID suboptions in client request
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enter remote-id value. The type is string with length: 1..256.
    RemoteId interface{}

    // Specify Relay Agent Information Option authenticate. The type is
    // interface{} with range: 0..4294967295.
    Authenticate interface{}
}

func (baseRelayOpt *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt) GetEntityData() *types.CommonEntityData {
    baseRelayOpt.EntityData.YFilter = baseRelayOpt.YFilter
    baseRelayOpt.EntityData.YangName = "base-relay-opt"
    baseRelayOpt.EntityData.BundleName = "cisco_ios_xr"
    baseRelayOpt.EntityData.ParentYangName = "base"
    baseRelayOpt.EntityData.SegmentPath = "base-relay-opt"
    baseRelayOpt.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/base/" + baseRelayOpt.EntityData.SegmentPath
    baseRelayOpt.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    baseRelayOpt.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    baseRelayOpt.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    baseRelayOpt.EntityData.Children = types.NewOrderedMap()
    baseRelayOpt.EntityData.Leafs = types.NewOrderedMap()
    baseRelayOpt.EntityData.Leafs.Append("remote-id", types.YLeaf{"RemoteId", baseRelayOpt.RemoteId})
    baseRelayOpt.EntityData.Leafs.Append("authenticate", types.YLeaf{"Authenticate", baseRelayOpt.Authenticate})

    baseRelayOpt.EntityData.YListKeys = []string {}

    return &(baseRelayOpt.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch
// Insert match keyword
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify match option.
    Options Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options
}

func (baseMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch) GetEntityData() *types.CommonEntityData {
    baseMatch.EntityData.YFilter = baseMatch.YFilter
    baseMatch.EntityData.YangName = "base-match"
    baseMatch.EntityData.BundleName = "cisco_ios_xr"
    baseMatch.EntityData.ParentYangName = "base"
    baseMatch.EntityData.SegmentPath = "base-match"
    baseMatch.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/base/" + baseMatch.EntityData.SegmentPath
    baseMatch.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    baseMatch.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    baseMatch.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    baseMatch.EntityData.Children = types.NewOrderedMap()
    baseMatch.EntityData.Children.Append("options", types.YChild{"Options", &baseMatch.Options})
    baseMatch.EntityData.Leafs = types.NewOrderedMap()

    baseMatch.EntityData.YListKeys = []string {}

    return &(baseMatch.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options
// Specify match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // none. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option.
    Option []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option
}

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options) GetEntityData() *types.CommonEntityData {
    options.EntityData.YFilter = options.YFilter
    options.EntityData.YangName = "options"
    options.EntityData.BundleName = "cisco_ios_xr"
    options.EntityData.ParentYangName = "base-match"
    options.EntityData.SegmentPath = "options"
    options.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/base/base-match/" + options.EntityData.SegmentPath
    options.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    options.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    options.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    options.EntityData.Children = types.NewOrderedMap()
    options.EntityData.Children.Append("option", types.YChild{"Option", nil})
    for i := range options.Option {
        options.EntityData.Children.Append(types.GetSegmentPath(options.Option[i]), types.YChild{"Option", options.Option[i]})
    }
    options.EntityData.Leafs = types.NewOrderedMap()

    options.EntityData.YListKeys = []string {}

    return &(options.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option
// none
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. none. The type is interface{} with range:
    // 0..4294967295.
    Opt60 interface{}

    // This attribute is a key. Enter hex pattern string. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Opt60HexStr interface{}

    // This attribute is a key. Set constant integer. The type is interface{} with
    // range: 0..4294967295.
    Format interface{}

    // Enter a profile.
    OptionProfile Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option) GetEntityData() *types.CommonEntityData {
    option.EntityData.YFilter = option.YFilter
    option.EntityData.YangName = "option"
    option.EntityData.BundleName = "cisco_ios_xr"
    option.EntityData.ParentYangName = "options"
    option.EntityData.SegmentPath = "option" + types.AddKeyToken(option.Opt60, "opt60") + types.AddKeyToken(option.Opt60HexStr, "opt60-hex-str") + types.AddKeyToken(option.Format, "format")
    option.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/base/base-match/options/" + option.EntityData.SegmentPath
    option.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    option.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    option.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    option.EntityData.Children = types.NewOrderedMap()
    option.EntityData.Children.Append("option-profile", types.YChild{"OptionProfile", &option.OptionProfile})
    option.EntityData.Leafs = types.NewOrderedMap()
    option.EntityData.Leafs.Append("opt60", types.YLeaf{"Opt60", option.Opt60})
    option.EntityData.Leafs.Append("opt60-hex-str", types.YLeaf{"Opt60HexStr", option.Opt60HexStr})
    option.EntityData.Leafs.Append("format", types.YLeaf{"Format", option.Format})

    option.EntityData.YListKeys = []string {"Opt60", "Opt60HexStr", "Format"}

    return &(option.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile
// Enter a profile
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Profile name. The type is string with length: 1..64.
    ProfileName interface{}

    // none. The type is interface{} with range: 0..4294967295.
    ProfileMode interface{}
}

func (optionProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile) GetEntityData() *types.CommonEntityData {
    optionProfile.EntityData.YFilter = optionProfile.YFilter
    optionProfile.EntityData.YangName = "option-profile"
    optionProfile.EntityData.BundleName = "cisco_ios_xr"
    optionProfile.EntityData.ParentYangName = "option"
    optionProfile.EntityData.SegmentPath = "option-profile"
    optionProfile.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/base/base-match/options/option/" + optionProfile.EntityData.SegmentPath
    optionProfile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optionProfile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optionProfile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optionProfile.EntityData.Children = types.NewOrderedMap()
    optionProfile.EntityData.Leafs = types.NewOrderedMap()
    optionProfile.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", optionProfile.ProfileName})
    optionProfile.EntityData.Leafs.Append("profile-mode", types.YLeaf{"ProfileMode", optionProfile.ProfileMode})

    optionProfile.EntityData.YListKeys = []string {}

    return &(optionProfile.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server
// DHCP Server profile
// This type is a presence type.
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Allow dhcp subscriber move. The type is interface{}.
    ServerAllowMove interface{}

    // DHCP IPV4 profile mode enable. The type is interface{}. This attribute is
    // mandatory.
    Enable interface{}

    // Configure Subnet Mask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SubnetMask interface{}

    // Specify the Pool name. The type is string with length: 1..64.
    Pool interface{}

    // Domain name. The type is string with length: 1..256.
    DomainName interface{}

    // Enable Secure Arp. The type is interface{}.
    SecureArp interface{}

    // Skip ARP installation for standalone sessions. The type is interface{}.
    ArpInstalSkipStdalone interface{}

    // Boot Filename. The type is string with length: 1..128.
    BootFilename interface{}

    // Configure the tftp-server IP to be used by the client. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextServer interface{}

    // Validate server ID check.
    ServerIdCheck Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck

    // Specify limit lease.
    LeaseLimit Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit

    // Validate Requested IP Address.
    RequestedIpAddress Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress

    // Enable aaa dhcp option force-insert.
    AaaServer Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer

    // default routers.
    DefaultRouters Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters

    // Delete binding on receiving discover.
    DeleteBindingOnDiscover Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DeleteBindingOnDiscover

    // NetBIOS name servers.
    NetBiosNameServers Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers

    // Insert match keyword.
    Match Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match

    // None.
    BroadcastFlag Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag

    // Change sessions configuration.
    Session Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session

    // Table of Class.
    Classes Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes

    // Specify Relay Agent Information Option configuration.
    Relay Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay

    // lease.
    Lease Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease

    // NetBIOS node type.
    NetbiosNodeType Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType

    // DNS servers.
    DnsServers Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers

    // Enable to provide the list of options need to send to aaa.
    DhcpToAaa Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DhcpToAaa

    // Table of OptionCode.
    OptionCodes Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes
}

func (server *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server) GetEntityData() *types.CommonEntityData {
    server.EntityData.YFilter = server.YFilter
    server.EntityData.YangName = "server"
    server.EntityData.BundleName = "cisco_ios_xr"
    server.EntityData.ParentYangName = "mode"
    server.EntityData.SegmentPath = "server"
    server.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/" + server.EntityData.SegmentPath
    server.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    server.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    server.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    server.EntityData.Children = types.NewOrderedMap()
    server.EntityData.Children.Append("server-id-check", types.YChild{"ServerIdCheck", &server.ServerIdCheck})
    server.EntityData.Children.Append("lease-limit", types.YChild{"LeaseLimit", &server.LeaseLimit})
    server.EntityData.Children.Append("requested-ip-address", types.YChild{"RequestedIpAddress", &server.RequestedIpAddress})
    server.EntityData.Children.Append("aaa-server", types.YChild{"AaaServer", &server.AaaServer})
    server.EntityData.Children.Append("default-routers", types.YChild{"DefaultRouters", &server.DefaultRouters})
    server.EntityData.Children.Append("delete-binding-on-discover", types.YChild{"DeleteBindingOnDiscover", &server.DeleteBindingOnDiscover})
    server.EntityData.Children.Append("net-bios-name-servers", types.YChild{"NetBiosNameServers", &server.NetBiosNameServers})
    server.EntityData.Children.Append("match", types.YChild{"Match", &server.Match})
    server.EntityData.Children.Append("broadcast-flag", types.YChild{"BroadcastFlag", &server.BroadcastFlag})
    server.EntityData.Children.Append("session", types.YChild{"Session", &server.Session})
    server.EntityData.Children.Append("classes", types.YChild{"Classes", &server.Classes})
    server.EntityData.Children.Append("relay", types.YChild{"Relay", &server.Relay})
    server.EntityData.Children.Append("lease", types.YChild{"Lease", &server.Lease})
    server.EntityData.Children.Append("netbios-node-type", types.YChild{"NetbiosNodeType", &server.NetbiosNodeType})
    server.EntityData.Children.Append("dns-servers", types.YChild{"DnsServers", &server.DnsServers})
    server.EntityData.Children.Append("dhcp-to-aaa", types.YChild{"DhcpToAaa", &server.DhcpToAaa})
    server.EntityData.Children.Append("option-codes", types.YChild{"OptionCodes", &server.OptionCodes})
    server.EntityData.Leafs = types.NewOrderedMap()
    server.EntityData.Leafs.Append("server-allow-move", types.YLeaf{"ServerAllowMove", server.ServerAllowMove})
    server.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", server.Enable})
    server.EntityData.Leafs.Append("subnet-mask", types.YLeaf{"SubnetMask", server.SubnetMask})
    server.EntityData.Leafs.Append("pool", types.YLeaf{"Pool", server.Pool})
    server.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", server.DomainName})
    server.EntityData.Leafs.Append("secure-arp", types.YLeaf{"SecureArp", server.SecureArp})
    server.EntityData.Leafs.Append("arp-instal-skip-stdalone", types.YLeaf{"ArpInstalSkipStdalone", server.ArpInstalSkipStdalone})
    server.EntityData.Leafs.Append("boot-filename", types.YLeaf{"BootFilename", server.BootFilename})
    server.EntityData.Leafs.Append("next-server", types.YLeaf{"NextServer", server.NextServer})

    server.EntityData.YListKeys = []string {}

    return &(server.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck
// Validate server ID check
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // specify server-id-check disable. The type is interface{}.
    Check interface{}
}

func (serverIdCheck *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck) GetEntityData() *types.CommonEntityData {
    serverIdCheck.EntityData.YFilter = serverIdCheck.YFilter
    serverIdCheck.EntityData.YangName = "server-id-check"
    serverIdCheck.EntityData.BundleName = "cisco_ios_xr"
    serverIdCheck.EntityData.ParentYangName = "server"
    serverIdCheck.EntityData.SegmentPath = "server-id-check"
    serverIdCheck.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + serverIdCheck.EntityData.SegmentPath
    serverIdCheck.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverIdCheck.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverIdCheck.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverIdCheck.EntityData.Children = types.NewOrderedMap()
    serverIdCheck.EntityData.Leafs = types.NewOrderedMap()
    serverIdCheck.EntityData.Leafs.Append("check", types.YLeaf{"Check", serverIdCheck.Check})

    serverIdCheck.EntityData.YListKeys = []string {}

    return &(serverIdCheck.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit
// Specify limit lease
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure Lease limit value. The type is LeaseLimitValue.
    LeaseLimitValue interface{}

    // Value of limit lease count in Decimal. The type is interface{} with range:
    // 1..240000.
    Range interface{}
}

func (leaseLimit *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit) GetEntityData() *types.CommonEntityData {
    leaseLimit.EntityData.YFilter = leaseLimit.YFilter
    leaseLimit.EntityData.YangName = "lease-limit"
    leaseLimit.EntityData.BundleName = "cisco_ios_xr"
    leaseLimit.EntityData.ParentYangName = "server"
    leaseLimit.EntityData.SegmentPath = "lease-limit"
    leaseLimit.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + leaseLimit.EntityData.SegmentPath
    leaseLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseLimit.EntityData.Children = types.NewOrderedMap()
    leaseLimit.EntityData.Leafs = types.NewOrderedMap()
    leaseLimit.EntityData.Leafs.Append("lease-limit-value", types.YLeaf{"LeaseLimitValue", leaseLimit.LeaseLimitValue})
    leaseLimit.EntityData.Leafs.Append("range", types.YLeaf{"Range", leaseLimit.Range})

    leaseLimit.EntityData.YListKeys = []string {}

    return &(leaseLimit.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress
// Validate Requested IP Address
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // specify requested-ip-address-check disable. The type is interface{}.
    Check interface{}
}

func (requestedIpAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress) GetEntityData() *types.CommonEntityData {
    requestedIpAddress.EntityData.YFilter = requestedIpAddress.YFilter
    requestedIpAddress.EntityData.YangName = "requested-ip-address"
    requestedIpAddress.EntityData.BundleName = "cisco_ios_xr"
    requestedIpAddress.EntityData.ParentYangName = "server"
    requestedIpAddress.EntityData.SegmentPath = "requested-ip-address"
    requestedIpAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + requestedIpAddress.EntityData.SegmentPath
    requestedIpAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    requestedIpAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    requestedIpAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    requestedIpAddress.EntityData.Children = types.NewOrderedMap()
    requestedIpAddress.EntityData.Leafs = types.NewOrderedMap()
    requestedIpAddress.EntityData.Leafs.Append("check", types.YLeaf{"Check", requestedIpAddress.Check})

    requestedIpAddress.EntityData.YListKeys = []string {}

    return &(requestedIpAddress.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer
// Enable aaa dhcp option force-insert
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable aaa dhcp option force-insert.
    DhcpOption Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption
}

func (aaaServer *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer) GetEntityData() *types.CommonEntityData {
    aaaServer.EntityData.YFilter = aaaServer.YFilter
    aaaServer.EntityData.YangName = "aaa-server"
    aaaServer.EntityData.BundleName = "cisco_ios_xr"
    aaaServer.EntityData.ParentYangName = "server"
    aaaServer.EntityData.SegmentPath = "aaa-server"
    aaaServer.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + aaaServer.EntityData.SegmentPath
    aaaServer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    aaaServer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    aaaServer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    aaaServer.EntityData.Children = types.NewOrderedMap()
    aaaServer.EntityData.Children.Append("dhcp-option", types.YChild{"DhcpOption", &aaaServer.DhcpOption})
    aaaServer.EntityData.Leafs = types.NewOrderedMap()

    aaaServer.EntityData.YListKeys = []string {}

    return &(aaaServer.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption
// Enable aaa dhcp option force-insert
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable aaa dhcp option force-insert. The type is interface{}.
    ForceInsert interface{}
}

func (dhcpOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption) GetEntityData() *types.CommonEntityData {
    dhcpOption.EntityData.YFilter = dhcpOption.YFilter
    dhcpOption.EntityData.YangName = "dhcp-option"
    dhcpOption.EntityData.BundleName = "cisco_ios_xr"
    dhcpOption.EntityData.ParentYangName = "aaa-server"
    dhcpOption.EntityData.SegmentPath = "dhcp-option"
    dhcpOption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/aaa-server/" + dhcpOption.EntityData.SegmentPath
    dhcpOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpOption.EntityData.Children = types.NewOrderedMap()
    dhcpOption.EntityData.Leafs = types.NewOrderedMap()
    dhcpOption.EntityData.Leafs.Append("force-insert", types.YLeaf{"ForceInsert", dhcpOption.ForceInsert})

    dhcpOption.EntityData.YListKeys = []string {}

    return &(dhcpOption.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters
// default routers
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router's IP address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DefaultRouter []interface{}
}

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters) GetEntityData() *types.CommonEntityData {
    defaultRouters.EntityData.YFilter = defaultRouters.YFilter
    defaultRouters.EntityData.YangName = "default-routers"
    defaultRouters.EntityData.BundleName = "cisco_ios_xr"
    defaultRouters.EntityData.ParentYangName = "server"
    defaultRouters.EntityData.SegmentPath = "default-routers"
    defaultRouters.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + defaultRouters.EntityData.SegmentPath
    defaultRouters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultRouters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultRouters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultRouters.EntityData.Children = types.NewOrderedMap()
    defaultRouters.EntityData.Leafs = types.NewOrderedMap()
    defaultRouters.EntityData.Leafs.Append("default-router", types.YLeaf{"DefaultRouter", defaultRouters.DefaultRouter})

    defaultRouters.EntityData.YListKeys = []string {}

    return &(defaultRouters.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DeleteBindingOnDiscover
// Delete binding on receiving discover
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DeleteBindingOnDiscover struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable delete binding on discover. The type is interface{}.
    Disable interface{}
}

func (deleteBindingOnDiscover *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DeleteBindingOnDiscover) GetEntityData() *types.CommonEntityData {
    deleteBindingOnDiscover.EntityData.YFilter = deleteBindingOnDiscover.YFilter
    deleteBindingOnDiscover.EntityData.YangName = "delete-binding-on-discover"
    deleteBindingOnDiscover.EntityData.BundleName = "cisco_ios_xr"
    deleteBindingOnDiscover.EntityData.ParentYangName = "server"
    deleteBindingOnDiscover.EntityData.SegmentPath = "delete-binding-on-discover"
    deleteBindingOnDiscover.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + deleteBindingOnDiscover.EntityData.SegmentPath
    deleteBindingOnDiscover.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    deleteBindingOnDiscover.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    deleteBindingOnDiscover.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    deleteBindingOnDiscover.EntityData.Children = types.NewOrderedMap()
    deleteBindingOnDiscover.EntityData.Leafs = types.NewOrderedMap()
    deleteBindingOnDiscover.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", deleteBindingOnDiscover.Disable})

    deleteBindingOnDiscover.EntityData.YListKeys = []string {}

    return &(deleteBindingOnDiscover.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers
// NetBIOS name servers
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NetBIOSNameServer's IP address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NetBiosNameServer []interface{}
}

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers) GetEntityData() *types.CommonEntityData {
    netBiosNameServers.EntityData.YFilter = netBiosNameServers.YFilter
    netBiosNameServers.EntityData.YangName = "net-bios-name-servers"
    netBiosNameServers.EntityData.BundleName = "cisco_ios_xr"
    netBiosNameServers.EntityData.ParentYangName = "server"
    netBiosNameServers.EntityData.SegmentPath = "net-bios-name-servers"
    netBiosNameServers.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + netBiosNameServers.EntityData.SegmentPath
    netBiosNameServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netBiosNameServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netBiosNameServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netBiosNameServers.EntityData.Children = types.NewOrderedMap()
    netBiosNameServers.EntityData.Leafs = types.NewOrderedMap()
    netBiosNameServers.EntityData.Leafs.Append("net-bios-name-server", types.YLeaf{"NetBiosNameServer", netBiosNameServers.NetBiosNameServer})

    netBiosNameServers.EntityData.YListKeys = []string {}

    return &(netBiosNameServers.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match
// Insert match keyword
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of OptionDefault.
    OptionDefaults Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults

    // Table of Option.
    Options Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match) GetEntityData() *types.CommonEntityData {
    match.EntityData.YFilter = match.YFilter
    match.EntityData.YangName = "match"
    match.EntityData.BundleName = "cisco_ios_xr"
    match.EntityData.ParentYangName = "server"
    match.EntityData.SegmentPath = "match"
    match.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + match.EntityData.SegmentPath
    match.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    match.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    match.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    match.EntityData.Children = types.NewOrderedMap()
    match.EntityData.Children.Append("option-defaults", types.YChild{"OptionDefaults", &match.OptionDefaults})
    match.EntityData.Children.Append("options", types.YChild{"Options", &match.Options})
    match.EntityData.Leafs = types.NewOrderedMap()

    match.EntityData.YListKeys = []string {}

    return &(match.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults
// Table of OptionDefault
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify match option. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault.
    OptionDefault []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault
}

func (optionDefaults *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults) GetEntityData() *types.CommonEntityData {
    optionDefaults.EntityData.YFilter = optionDefaults.YFilter
    optionDefaults.EntityData.YangName = "option-defaults"
    optionDefaults.EntityData.BundleName = "cisco_ios_xr"
    optionDefaults.EntityData.ParentYangName = "match"
    optionDefaults.EntityData.SegmentPath = "option-defaults"
    optionDefaults.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/match/" + optionDefaults.EntityData.SegmentPath
    optionDefaults.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optionDefaults.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optionDefaults.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optionDefaults.EntityData.Children = types.NewOrderedMap()
    optionDefaults.EntityData.Children.Append("option-default", types.YChild{"OptionDefault", nil})
    for i := range optionDefaults.OptionDefault {
        optionDefaults.EntityData.Children.Append(types.GetSegmentPath(optionDefaults.OptionDefault[i]), types.YChild{"OptionDefault", optionDefaults.OptionDefault[i]})
    }
    optionDefaults.EntityData.Leafs = types.NewOrderedMap()

    optionDefaults.EntityData.YListKeys = []string {}

    return &(optionDefaults.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault
// Specify match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Match option 60. The type is Matchoption.
    Matchoption interface{}

    // Vendor action. The type is Matchaction. This attribute is mandatory.
    Matchaction interface{}
}

func (optionDefault *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault) GetEntityData() *types.CommonEntityData {
    optionDefault.EntityData.YFilter = optionDefault.YFilter
    optionDefault.EntityData.YangName = "option-default"
    optionDefault.EntityData.BundleName = "cisco_ios_xr"
    optionDefault.EntityData.ParentYangName = "option-defaults"
    optionDefault.EntityData.SegmentPath = "option-default" + types.AddKeyToken(optionDefault.Matchoption, "matchoption")
    optionDefault.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/match/option-defaults/" + optionDefault.EntityData.SegmentPath
    optionDefault.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optionDefault.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optionDefault.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optionDefault.EntityData.Children = types.NewOrderedMap()
    optionDefault.EntityData.Leafs = types.NewOrderedMap()
    optionDefault.EntityData.Leafs.Append("matchoption", types.YLeaf{"Matchoption", optionDefault.Matchoption})
    optionDefault.EntityData.Leafs.Append("matchaction", types.YLeaf{"Matchaction", optionDefault.Matchaction})

    optionDefault.EntityData.YListKeys = []string {"Matchoption"}

    return &(optionDefault.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options
// Table of Option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify match option. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option.
    Option []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option
}

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options) GetEntityData() *types.CommonEntityData {
    options.EntityData.YFilter = options.YFilter
    options.EntityData.YangName = "options"
    options.EntityData.BundleName = "cisco_ios_xr"
    options.EntityData.ParentYangName = "match"
    options.EntityData.SegmentPath = "options"
    options.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/match/" + options.EntityData.SegmentPath
    options.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    options.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    options.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    options.EntityData.Children = types.NewOrderedMap()
    options.EntityData.Children.Append("option", types.YChild{"Option", nil})
    for i := range options.Option {
        options.EntityData.Children.Append(types.GetSegmentPath(options.Option[i]), types.YChild{"Option", options.Option[i]})
    }
    options.EntityData.Leafs = types.NewOrderedMap()

    options.EntityData.YListKeys = []string {}

    return &(options.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option
// Specify match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Match option 60. The type is Matchoption.
    Matchoption interface{}

    // This attribute is a key. Enter hex pattern string. The type is string with
    // length: 1..64.
    Pattern interface{}

    // This attribute is a key. Set constant integer. The type is interface{} with
    // range: 0..4294967295.
    Format interface{}

    // Vendor action. The type is Matchaction. This attribute is mandatory.
    Matchaction interface{}
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option) GetEntityData() *types.CommonEntityData {
    option.EntityData.YFilter = option.YFilter
    option.EntityData.YangName = "option"
    option.EntityData.BundleName = "cisco_ios_xr"
    option.EntityData.ParentYangName = "options"
    option.EntityData.SegmentPath = "option" + types.AddKeyToken(option.Matchoption, "matchoption") + types.AddKeyToken(option.Pattern, "pattern") + types.AddKeyToken(option.Format, "format")
    option.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/match/options/" + option.EntityData.SegmentPath
    option.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    option.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    option.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    option.EntityData.Children = types.NewOrderedMap()
    option.EntityData.Leafs = types.NewOrderedMap()
    option.EntityData.Leafs.Append("matchoption", types.YLeaf{"Matchoption", option.Matchoption})
    option.EntityData.Leafs.Append("pattern", types.YLeaf{"Pattern", option.Pattern})
    option.EntityData.Leafs.Append("format", types.YLeaf{"Format", option.Format})
    option.EntityData.Leafs.Append("matchaction", types.YLeaf{"Matchaction", option.Matchaction})

    option.EntityData.YListKeys = []string {"Matchoption", "Pattern", "Format"}

    return &(option.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag
// None
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify broadcast flag policy. The type is Policy.
    Policy interface{}
}

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag) GetEntityData() *types.CommonEntityData {
    broadcastFlag.EntityData.YFilter = broadcastFlag.YFilter
    broadcastFlag.EntityData.YangName = "broadcast-flag"
    broadcastFlag.EntityData.BundleName = "cisco_ios_xr"
    broadcastFlag.EntityData.ParentYangName = "server"
    broadcastFlag.EntityData.SegmentPath = "broadcast-flag"
    broadcastFlag.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + broadcastFlag.EntityData.SegmentPath
    broadcastFlag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    broadcastFlag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    broadcastFlag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    broadcastFlag.EntityData.Children = types.NewOrderedMap()
    broadcastFlag.EntityData.Leafs = types.NewOrderedMap()
    broadcastFlag.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", broadcastFlag.Policy})

    broadcastFlag.EntityData.YListKeys = []string {}

    return &(broadcastFlag.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session
// Change sessions configuration
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Throttle DHCP sessions based on MAC address.
    ThrottleType Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType
}

func (session *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "server"
    session.EntityData.SegmentPath = "session"
    session.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + session.EntityData.SegmentPath
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Children.Append("throttle-type", types.YChild{"ThrottleType", &session.ThrottleType})
    session.EntityData.Leafs = types.NewOrderedMap()

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType
// Throttle DHCP sessions based on MAC
// address
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Throttle DHCP sessions from any one MAC address.
    MacThrottle Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle
}

func (throttleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType) GetEntityData() *types.CommonEntityData {
    throttleType.EntityData.YFilter = throttleType.YFilter
    throttleType.EntityData.YangName = "throttle-type"
    throttleType.EntityData.BundleName = "cisco_ios_xr"
    throttleType.EntityData.ParentYangName = "session"
    throttleType.EntityData.SegmentPath = "throttle-type"
    throttleType.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/session/" + throttleType.EntityData.SegmentPath
    throttleType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttleType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttleType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttleType.EntityData.Children = types.NewOrderedMap()
    throttleType.EntityData.Children.Append("mac-throttle", types.YChild{"MacThrottle", &throttleType.MacThrottle})
    throttleType.EntityData.Leafs = types.NewOrderedMap()

    throttleType.EntityData.YListKeys = []string {}

    return &(throttleType.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle
// Throttle DHCP sessions from any one MAC
// address
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of discovers at which to throttle. The type is interface{} with
    // range: 1..65535.
    NumDiscover interface{}

    // Throttle request period (in secs). The type is interface{} with range:
    // 1..100. Units are second.
    NumRequest interface{}

    // Throttle blocking period (in secs). The type is interface{} with range:
    // 1..100. Units are second.
    NumBlock interface{}
}

func (macThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle) GetEntityData() *types.CommonEntityData {
    macThrottle.EntityData.YFilter = macThrottle.YFilter
    macThrottle.EntityData.YangName = "mac-throttle"
    macThrottle.EntityData.BundleName = "cisco_ios_xr"
    macThrottle.EntityData.ParentYangName = "throttle-type"
    macThrottle.EntityData.SegmentPath = "mac-throttle"
    macThrottle.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/session/throttle-type/" + macThrottle.EntityData.SegmentPath
    macThrottle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    macThrottle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    macThrottle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    macThrottle.EntityData.Children = types.NewOrderedMap()
    macThrottle.EntityData.Leafs = types.NewOrderedMap()
    macThrottle.EntityData.Leafs.Append("num-discover", types.YLeaf{"NumDiscover", macThrottle.NumDiscover})
    macThrottle.EntityData.Leafs.Append("num-request", types.YLeaf{"NumRequest", macThrottle.NumRequest})
    macThrottle.EntityData.Leafs.Append("num-block", types.YLeaf{"NumBlock", macThrottle.NumBlock})

    macThrottle.EntityData.YListKeys = []string {}

    return &(macThrottle.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes
// Table of Class
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Create or enter server profile class. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class.
    Class []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class
}

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes) GetEntityData() *types.CommonEntityData {
    classes.EntityData.YFilter = classes.YFilter
    classes.EntityData.YangName = "classes"
    classes.EntityData.BundleName = "cisco_ios_xr"
    classes.EntityData.ParentYangName = "server"
    classes.EntityData.SegmentPath = "classes"
    classes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + classes.EntityData.SegmentPath
    classes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classes.EntityData.Children = types.NewOrderedMap()
    classes.EntityData.Children.Append("class", types.YChild{"Class", nil})
    for i := range classes.Class {
        classes.EntityData.Children.Append(types.GetSegmentPath(classes.Class[i]), types.YChild{"Class", classes.Class[i]})
    }
    classes.EntityData.Leafs = types.NewOrderedMap()

    classes.EntityData.YListKeys = []string {}

    return &(classes.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class
// Create or enter server profile class
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. class name. The type is string with length:
    // 1..128.
    ClassName interface{}

    // Configure Subnet Mask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SubnetMask interface{}

    // Specify the pool. The type is string.
    Pool interface{}

    // Enable Create or enter server profile class. Deletion of this object also
    // causes deletion of all associated objects under Class. The type is
    // interface{}.
    Enable interface{}

    // Domain name. The type is string with length: 1..256.
    DomainName interface{}

    // Boot Filename. The type is string with length: 1..128.
    BootFilename interface{}

    // Configure the tftp-server IP to be used by the client. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NextServer interface{}

    // default routers.
    DefaultRouters Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters

    // NetBIOS name servers.
    NetBiosNameServers Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers

    // Insert match keyword.
    ClassMatch Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch

    // lease.
    Lease Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease

    // NetBIOS node type.
    NetbiosNodeType Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType

    // DNS servers.
    DnsServers Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers

    // Table of OptionCode.
    OptionCodes Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes
}

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + types.AddKeyToken(class.ClassName, "class-name")
    class.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/classes/" + class.EntityData.SegmentPath
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = types.NewOrderedMap()
    class.EntityData.Children.Append("default-routers", types.YChild{"DefaultRouters", &class.DefaultRouters})
    class.EntityData.Children.Append("net-bios-name-servers", types.YChild{"NetBiosNameServers", &class.NetBiosNameServers})
    class.EntityData.Children.Append("class-match", types.YChild{"ClassMatch", &class.ClassMatch})
    class.EntityData.Children.Append("lease", types.YChild{"Lease", &class.Lease})
    class.EntityData.Children.Append("netbios-node-type", types.YChild{"NetbiosNodeType", &class.NetbiosNodeType})
    class.EntityData.Children.Append("dns-servers", types.YChild{"DnsServers", &class.DnsServers})
    class.EntityData.Children.Append("option-codes", types.YChild{"OptionCodes", &class.OptionCodes})
    class.EntityData.Leafs = types.NewOrderedMap()
    class.EntityData.Leafs.Append("class-name", types.YLeaf{"ClassName", class.ClassName})
    class.EntityData.Leafs.Append("subnet-mask", types.YLeaf{"SubnetMask", class.SubnetMask})
    class.EntityData.Leafs.Append("pool", types.YLeaf{"Pool", class.Pool})
    class.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", class.Enable})
    class.EntityData.Leafs.Append("domain-name", types.YLeaf{"DomainName", class.DomainName})
    class.EntityData.Leafs.Append("boot-filename", types.YLeaf{"BootFilename", class.BootFilename})
    class.EntityData.Leafs.Append("next-server", types.YLeaf{"NextServer", class.NextServer})

    class.EntityData.YListKeys = []string {"ClassName"}

    return &(class.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters
// default routers
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router's IP address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DefaultRouter []interface{}
}

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters) GetEntityData() *types.CommonEntityData {
    defaultRouters.EntityData.YFilter = defaultRouters.YFilter
    defaultRouters.EntityData.YangName = "default-routers"
    defaultRouters.EntityData.BundleName = "cisco_ios_xr"
    defaultRouters.EntityData.ParentYangName = "class"
    defaultRouters.EntityData.SegmentPath = "default-routers"
    defaultRouters.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/classes/class/" + defaultRouters.EntityData.SegmentPath
    defaultRouters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultRouters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultRouters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultRouters.EntityData.Children = types.NewOrderedMap()
    defaultRouters.EntityData.Leafs = types.NewOrderedMap()
    defaultRouters.EntityData.Leafs.Append("default-router", types.YLeaf{"DefaultRouter", defaultRouters.DefaultRouter})

    defaultRouters.EntityData.YListKeys = []string {}

    return &(defaultRouters.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers
// NetBIOS name servers
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NetBIOSNameServer's IP address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NetBiosNameServer []interface{}
}

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers) GetEntityData() *types.CommonEntityData {
    netBiosNameServers.EntityData.YFilter = netBiosNameServers.YFilter
    netBiosNameServers.EntityData.YangName = "net-bios-name-servers"
    netBiosNameServers.EntityData.BundleName = "cisco_ios_xr"
    netBiosNameServers.EntityData.ParentYangName = "class"
    netBiosNameServers.EntityData.SegmentPath = "net-bios-name-servers"
    netBiosNameServers.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/classes/class/" + netBiosNameServers.EntityData.SegmentPath
    netBiosNameServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netBiosNameServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netBiosNameServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netBiosNameServers.EntityData.Children = types.NewOrderedMap()
    netBiosNameServers.EntityData.Leafs = types.NewOrderedMap()
    netBiosNameServers.EntityData.Leafs.Append("net-bios-name-server", types.YLeaf{"NetBiosNameServer", netBiosNameServers.NetBiosNameServer})

    netBiosNameServers.EntityData.YListKeys = []string {}

    return &(netBiosNameServers.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch
// Insert match keyword
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify match l2-interface. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    L2Interface interface{}

    // Specify match VRF. The type is string with length: 1..32.
    Vrf interface{}

    // Table of Class-Option.
    ClassOptions Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions
}

func (classMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch) GetEntityData() *types.CommonEntityData {
    classMatch.EntityData.YFilter = classMatch.YFilter
    classMatch.EntityData.YangName = "class-match"
    classMatch.EntityData.BundleName = "cisco_ios_xr"
    classMatch.EntityData.ParentYangName = "class"
    classMatch.EntityData.SegmentPath = "class-match"
    classMatch.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/classes/class/" + classMatch.EntityData.SegmentPath
    classMatch.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classMatch.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classMatch.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classMatch.EntityData.Children = types.NewOrderedMap()
    classMatch.EntityData.Children.Append("class-options", types.YChild{"ClassOptions", &classMatch.ClassOptions})
    classMatch.EntityData.Leafs = types.NewOrderedMap()
    classMatch.EntityData.Leafs.Append("l2-interface", types.YLeaf{"L2Interface", classMatch.L2Interface})
    classMatch.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", classMatch.Vrf})

    classMatch.EntityData.YListKeys = []string {}

    return &(classMatch.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions
// Table of Class-Option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify match option. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption.
    ClassOption []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption
}

func (classOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions) GetEntityData() *types.CommonEntityData {
    classOptions.EntityData.YFilter = classOptions.YFilter
    classOptions.EntityData.YangName = "class-options"
    classOptions.EntityData.BundleName = "cisco_ios_xr"
    classOptions.EntityData.ParentYangName = "class-match"
    classOptions.EntityData.SegmentPath = "class-options"
    classOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/classes/class/class-match/" + classOptions.EntityData.SegmentPath
    classOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classOptions.EntityData.Children = types.NewOrderedMap()
    classOptions.EntityData.Children.Append("class-option", types.YChild{"ClassOption", nil})
    for i := range classOptions.ClassOption {
        classOptions.EntityData.Children.Append(types.GetSegmentPath(classOptions.ClassOption[i]), types.YChild{"ClassOption", classOptions.ClassOption[i]})
    }
    classOptions.EntityData.Leafs = types.NewOrderedMap()

    classOptions.EntityData.YListKeys = []string {}

    return &(classOptions.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption
// Specify match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Match options. The type is Matchoption.
    Matchoption interface{}

    // Enter hex pattern string. The type is string with length: 1..64.
    Pattern interface{}

    // Enter bit mask pattern string. The type is string with length: 1..64.
    BitMask interface{}
}

func (classOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption) GetEntityData() *types.CommonEntityData {
    classOption.EntityData.YFilter = classOption.YFilter
    classOption.EntityData.YangName = "class-option"
    classOption.EntityData.BundleName = "cisco_ios_xr"
    classOption.EntityData.ParentYangName = "class-options"
    classOption.EntityData.SegmentPath = "class-option" + types.AddKeyToken(classOption.Matchoption, "matchoption")
    classOption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/classes/class/class-match/class-options/" + classOption.EntityData.SegmentPath
    classOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classOption.EntityData.Children = types.NewOrderedMap()
    classOption.EntityData.Leafs = types.NewOrderedMap()
    classOption.EntityData.Leafs.Append("matchoption", types.YLeaf{"Matchoption", classOption.Matchoption})
    classOption.EntityData.Leafs.Append("pattern", types.YLeaf{"Pattern", classOption.Pattern})
    classOption.EntityData.Leafs.Append("bit-mask", types.YLeaf{"BitMask", classOption.BitMask})

    classOption.EntityData.YListKeys = []string {"Matchoption"}

    return &(classOption.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease
// lease
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set string. The type is string.
    Infinite interface{}

    // Days. The type is interface{} with range: 0..365. Units are day.
    Days interface{}

    // Hours. The type is interface{} with range: 0..23. Units are hour.
    Hours interface{}

    // Minutes. The type is interface{} with range: 0..59. Units are minute.
    Minutes interface{}
}

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease) GetEntityData() *types.CommonEntityData {
    lease.EntityData.YFilter = lease.YFilter
    lease.EntityData.YangName = "lease"
    lease.EntityData.BundleName = "cisco_ios_xr"
    lease.EntityData.ParentYangName = "class"
    lease.EntityData.SegmentPath = "lease"
    lease.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/classes/class/" + lease.EntityData.SegmentPath
    lease.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lease.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lease.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lease.EntityData.Children = types.NewOrderedMap()
    lease.EntityData.Leafs = types.NewOrderedMap()
    lease.EntityData.Leafs.Append("infinite", types.YLeaf{"Infinite", lease.Infinite})
    lease.EntityData.Leafs.Append("days", types.YLeaf{"Days", lease.Days})
    lease.EntityData.Leafs.Append("hours", types.YLeaf{"Hours", lease.Hours})
    lease.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", lease.Minutes})

    lease.EntityData.YListKeys = []string {}

    return &(lease.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType
// NetBIOS node type
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set string. The type is string.
    BroadcastNode interface{}

    // Set string. The type is string.
    HybridNode interface{}

    // Set string. The type is string.
    MixedNode interface{}

    // Set string. The type is string.
    PeerToPeerNode interface{}

    // Hexadecimal number. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Hexadecimal interface{}
}

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType) GetEntityData() *types.CommonEntityData {
    netbiosNodeType.EntityData.YFilter = netbiosNodeType.YFilter
    netbiosNodeType.EntityData.YangName = "netbios-node-type"
    netbiosNodeType.EntityData.BundleName = "cisco_ios_xr"
    netbiosNodeType.EntityData.ParentYangName = "class"
    netbiosNodeType.EntityData.SegmentPath = "netbios-node-type"
    netbiosNodeType.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/classes/class/" + netbiosNodeType.EntityData.SegmentPath
    netbiosNodeType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netbiosNodeType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netbiosNodeType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netbiosNodeType.EntityData.Children = types.NewOrderedMap()
    netbiosNodeType.EntityData.Leafs = types.NewOrderedMap()
    netbiosNodeType.EntityData.Leafs.Append("broadcast-node", types.YLeaf{"BroadcastNode", netbiosNodeType.BroadcastNode})
    netbiosNodeType.EntityData.Leafs.Append("hybrid-node", types.YLeaf{"HybridNode", netbiosNodeType.HybridNode})
    netbiosNodeType.EntityData.Leafs.Append("mixed-node", types.YLeaf{"MixedNode", netbiosNodeType.MixedNode})
    netbiosNodeType.EntityData.Leafs.Append("peer-to-peer-node", types.YLeaf{"PeerToPeerNode", netbiosNodeType.PeerToPeerNode})
    netbiosNodeType.EntityData.Leafs.Append("hexadecimal", types.YLeaf{"Hexadecimal", netbiosNodeType.Hexadecimal})

    netbiosNodeType.EntityData.YListKeys = []string {}

    return &(netbiosNodeType.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers
// DNS servers
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DNS Server's IP address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DnsServer []interface{}
}

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers) GetEntityData() *types.CommonEntityData {
    dnsServers.EntityData.YFilter = dnsServers.YFilter
    dnsServers.EntityData.YangName = "dns-servers"
    dnsServers.EntityData.BundleName = "cisco_ios_xr"
    dnsServers.EntityData.ParentYangName = "class"
    dnsServers.EntityData.SegmentPath = "dns-servers"
    dnsServers.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/classes/class/" + dnsServers.EntityData.SegmentPath
    dnsServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dnsServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dnsServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dnsServers.EntityData.Children = types.NewOrderedMap()
    dnsServers.EntityData.Leafs = types.NewOrderedMap()
    dnsServers.EntityData.Leafs.Append("dns-server", types.YLeaf{"DnsServer", dnsServers.DnsServer})

    dnsServers.EntityData.YListKeys = []string {}

    return &(dnsServers.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes
// Table of OptionCode
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP option code. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode.
    OptionCode []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode
}

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes) GetEntityData() *types.CommonEntityData {
    optionCodes.EntityData.YFilter = optionCodes.YFilter
    optionCodes.EntityData.YangName = "option-codes"
    optionCodes.EntityData.BundleName = "cisco_ios_xr"
    optionCodes.EntityData.ParentYangName = "class"
    optionCodes.EntityData.SegmentPath = "option-codes"
    optionCodes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/classes/class/" + optionCodes.EntityData.SegmentPath
    optionCodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optionCodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optionCodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optionCodes.EntityData.Children = types.NewOrderedMap()
    optionCodes.EntityData.Children.Append("option-code", types.YChild{"OptionCode", nil})
    for i := range optionCodes.OptionCode {
        optionCodes.EntityData.Children.Append(types.GetSegmentPath(optionCodes.OptionCode[i]), types.YChild{"OptionCode", optionCodes.OptionCode[i]})
    }
    optionCodes.EntityData.Leafs = types.NewOrderedMap()

    optionCodes.EntityData.YListKeys = []string {}

    return &(optionCodes.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode
// DHCP option code
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. DHCP option code. The type is interface{} with
    // range: 0..255.
    OptionCode interface{}

    // ASCII string. The type is string.
    AsciiString interface{}

    // Hexadecimal string. The type is string.
    HexString interface{}

    // Set constant integer. The type is interface{} with range: 0..4294967295.
    ForceInsert interface{}

    // Server's IP address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress []interface{}
}

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode) GetEntityData() *types.CommonEntityData {
    optionCode.EntityData.YFilter = optionCode.YFilter
    optionCode.EntityData.YangName = "option-code"
    optionCode.EntityData.BundleName = "cisco_ios_xr"
    optionCode.EntityData.ParentYangName = "option-codes"
    optionCode.EntityData.SegmentPath = "option-code" + types.AddKeyToken(optionCode.OptionCode, "option-code")
    optionCode.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/classes/class/option-codes/" + optionCode.EntityData.SegmentPath
    optionCode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optionCode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optionCode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optionCode.EntityData.Children = types.NewOrderedMap()
    optionCode.EntityData.Leafs = types.NewOrderedMap()
    optionCode.EntityData.Leafs.Append("option-code", types.YLeaf{"OptionCode", optionCode.OptionCode})
    optionCode.EntityData.Leafs.Append("ascii-string", types.YLeaf{"AsciiString", optionCode.AsciiString})
    optionCode.EntityData.Leafs.Append("hex-string", types.YLeaf{"HexString", optionCode.HexString})
    optionCode.EntityData.Leafs.Append("force-insert", types.YLeaf{"ForceInsert", optionCode.ForceInsert})
    optionCode.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", optionCode.IpAddress})

    optionCode.EntityData.YListKeys = []string {"OptionCode"}

    return &(optionCode.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay
// Specify Relay Agent Information Option
// configuration
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify Relay Agent Information Option authenticate. The type is
    // interface{} with range: 0..4294967295.
    Authenticate interface{}
}

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay) GetEntityData() *types.CommonEntityData {
    relay.EntityData.YFilter = relay.YFilter
    relay.EntityData.YangName = "relay"
    relay.EntityData.BundleName = "cisco_ios_xr"
    relay.EntityData.ParentYangName = "server"
    relay.EntityData.SegmentPath = "relay"
    relay.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + relay.EntityData.SegmentPath
    relay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relay.EntityData.Children = types.NewOrderedMap()
    relay.EntityData.Leafs = types.NewOrderedMap()
    relay.EntityData.Leafs.Append("authenticate", types.YLeaf{"Authenticate", relay.Authenticate})

    relay.EntityData.YListKeys = []string {}

    return &(relay.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease
// lease
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set string. The type is string.
    Infinite interface{}

    // Days. The type is interface{} with range: 0..365. Units are day.
    Days interface{}

    // Hours. The type is interface{} with range: 0..23. Units are hour.
    Hours interface{}

    // Minutes. The type is interface{} with range: 0..59. Units are minute.
    Minutes interface{}
}

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease) GetEntityData() *types.CommonEntityData {
    lease.EntityData.YFilter = lease.YFilter
    lease.EntityData.YangName = "lease"
    lease.EntityData.BundleName = "cisco_ios_xr"
    lease.EntityData.ParentYangName = "server"
    lease.EntityData.SegmentPath = "lease"
    lease.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + lease.EntityData.SegmentPath
    lease.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lease.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lease.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lease.EntityData.Children = types.NewOrderedMap()
    lease.EntityData.Leafs = types.NewOrderedMap()
    lease.EntityData.Leafs.Append("infinite", types.YLeaf{"Infinite", lease.Infinite})
    lease.EntityData.Leafs.Append("days", types.YLeaf{"Days", lease.Days})
    lease.EntityData.Leafs.Append("hours", types.YLeaf{"Hours", lease.Hours})
    lease.EntityData.Leafs.Append("minutes", types.YLeaf{"Minutes", lease.Minutes})

    lease.EntityData.YListKeys = []string {}

    return &(lease.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType
// NetBIOS node type
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set string. The type is string.
    BroadcastNode interface{}

    // Set string. The type is string.
    HybridNode interface{}

    // Set string. The type is string.
    MixedNode interface{}

    // Set string. The type is string.
    PeerToPeerNode interface{}

    // Hexadecimal number. The type is string with pattern: [0-9a-fA-F]{1,8}.
    Hexadecimal interface{}
}

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType) GetEntityData() *types.CommonEntityData {
    netbiosNodeType.EntityData.YFilter = netbiosNodeType.YFilter
    netbiosNodeType.EntityData.YangName = "netbios-node-type"
    netbiosNodeType.EntityData.BundleName = "cisco_ios_xr"
    netbiosNodeType.EntityData.ParentYangName = "server"
    netbiosNodeType.EntityData.SegmentPath = "netbios-node-type"
    netbiosNodeType.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + netbiosNodeType.EntityData.SegmentPath
    netbiosNodeType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netbiosNodeType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netbiosNodeType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netbiosNodeType.EntityData.Children = types.NewOrderedMap()
    netbiosNodeType.EntityData.Leafs = types.NewOrderedMap()
    netbiosNodeType.EntityData.Leafs.Append("broadcast-node", types.YLeaf{"BroadcastNode", netbiosNodeType.BroadcastNode})
    netbiosNodeType.EntityData.Leafs.Append("hybrid-node", types.YLeaf{"HybridNode", netbiosNodeType.HybridNode})
    netbiosNodeType.EntityData.Leafs.Append("mixed-node", types.YLeaf{"MixedNode", netbiosNodeType.MixedNode})
    netbiosNodeType.EntityData.Leafs.Append("peer-to-peer-node", types.YLeaf{"PeerToPeerNode", netbiosNodeType.PeerToPeerNode})
    netbiosNodeType.EntityData.Leafs.Append("hexadecimal", types.YLeaf{"Hexadecimal", netbiosNodeType.Hexadecimal})

    netbiosNodeType.EntityData.YListKeys = []string {}

    return &(netbiosNodeType.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers
// DNS servers
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DNS Server's IP address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DnsServer []interface{}
}

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers) GetEntityData() *types.CommonEntityData {
    dnsServers.EntityData.YFilter = dnsServers.YFilter
    dnsServers.EntityData.YangName = "dns-servers"
    dnsServers.EntityData.BundleName = "cisco_ios_xr"
    dnsServers.EntityData.ParentYangName = "server"
    dnsServers.EntityData.SegmentPath = "dns-servers"
    dnsServers.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + dnsServers.EntityData.SegmentPath
    dnsServers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dnsServers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dnsServers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dnsServers.EntityData.Children = types.NewOrderedMap()
    dnsServers.EntityData.Leafs = types.NewOrderedMap()
    dnsServers.EntityData.Leafs.Append("dns-server", types.YLeaf{"DnsServer", dnsServers.DnsServer})

    dnsServers.EntityData.YListKeys = []string {}

    return &(dnsServers.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DhcpToAaa
// Enable to provide the list of options need
// to send to aaa
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DhcpToAaa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // option type.
    Option Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DhcpToAaa_Option
}

func (dhcpToAaa *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DhcpToAaa) GetEntityData() *types.CommonEntityData {
    dhcpToAaa.EntityData.YFilter = dhcpToAaa.YFilter
    dhcpToAaa.EntityData.YangName = "dhcp-to-aaa"
    dhcpToAaa.EntityData.BundleName = "cisco_ios_xr"
    dhcpToAaa.EntityData.ParentYangName = "server"
    dhcpToAaa.EntityData.SegmentPath = "dhcp-to-aaa"
    dhcpToAaa.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + dhcpToAaa.EntityData.SegmentPath
    dhcpToAaa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpToAaa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpToAaa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpToAaa.EntityData.Children = types.NewOrderedMap()
    dhcpToAaa.EntityData.Children.Append("option", types.YChild{"Option", &dhcpToAaa.Option})
    dhcpToAaa.EntityData.Leafs = types.NewOrderedMap()

    dhcpToAaa.EntityData.YListKeys = []string {}

    return &(dhcpToAaa.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DhcpToAaa_Option
// option type
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DhcpToAaa_Option struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of options.
    List Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DhcpToAaa_Option_List
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DhcpToAaa_Option) GetEntityData() *types.CommonEntityData {
    option.EntityData.YFilter = option.YFilter
    option.EntityData.YangName = "option"
    option.EntityData.BundleName = "cisco_ios_xr"
    option.EntityData.ParentYangName = "dhcp-to-aaa"
    option.EntityData.SegmentPath = "option"
    option.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/dhcp-to-aaa/" + option.EntityData.SegmentPath
    option.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    option.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    option.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    option.EntityData.Children = types.NewOrderedMap()
    option.EntityData.Children.Append("list", types.YChild{"List", &option.List})
    option.EntityData.Leafs = types.NewOrderedMap()

    option.EntityData.YListKeys = []string {}

    return &(option.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DhcpToAaa_Option_List
// List of options
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DhcpToAaa_Option_List struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set constant integer. The type is interface{} with range: 0..4294967295.
    OptionAll interface{}

    // Option number. The type is slice of interface{} with range: 0..4294967295.
    OptionNumber []interface{}
}

func (list *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DhcpToAaa_Option_List) GetEntityData() *types.CommonEntityData {
    list.EntityData.YFilter = list.YFilter
    list.EntityData.YangName = "list"
    list.EntityData.BundleName = "cisco_ios_xr"
    list.EntityData.ParentYangName = "option"
    list.EntityData.SegmentPath = "list"
    list.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/dhcp-to-aaa/option/" + list.EntityData.SegmentPath
    list.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    list.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    list.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    list.EntityData.Children = types.NewOrderedMap()
    list.EntityData.Leafs = types.NewOrderedMap()
    list.EntityData.Leafs.Append("option-all", types.YLeaf{"OptionAll", list.OptionAll})
    list.EntityData.Leafs.Append("option-number", types.YLeaf{"OptionNumber", list.OptionNumber})

    list.EntityData.YListKeys = []string {}

    return &(list.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes
// Table of OptionCode
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP option code. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode.
    OptionCode []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode
}

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes) GetEntityData() *types.CommonEntityData {
    optionCodes.EntityData.YFilter = optionCodes.YFilter
    optionCodes.EntityData.YangName = "option-codes"
    optionCodes.EntityData.BundleName = "cisco_ios_xr"
    optionCodes.EntityData.ParentYangName = "server"
    optionCodes.EntityData.SegmentPath = "option-codes"
    optionCodes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/" + optionCodes.EntityData.SegmentPath
    optionCodes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optionCodes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optionCodes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optionCodes.EntityData.Children = types.NewOrderedMap()
    optionCodes.EntityData.Children.Append("option-code", types.YChild{"OptionCode", nil})
    for i := range optionCodes.OptionCode {
        optionCodes.EntityData.Children.Append(types.GetSegmentPath(optionCodes.OptionCode[i]), types.YChild{"OptionCode", optionCodes.OptionCode[i]})
    }
    optionCodes.EntityData.Leafs = types.NewOrderedMap()

    optionCodes.EntityData.YListKeys = []string {}

    return &(optionCodes.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode
// DHCP option code
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. DHCP option code. The type is interface{} with
    // range: 0..255.
    OptionCode interface{}

    // ASCII string. The type is string.
    AsciiString interface{}

    // Hexadecimal string. The type is string.
    HexString interface{}

    // Set constant integer. The type is interface{} with range: 0..4294967295.
    ForceInsert interface{}

    // Server's IP address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress []interface{}
}

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode) GetEntityData() *types.CommonEntityData {
    optionCode.EntityData.YFilter = optionCode.YFilter
    optionCode.EntityData.YangName = "option-code"
    optionCode.EntityData.BundleName = "cisco_ios_xr"
    optionCode.EntityData.ParentYangName = "option-codes"
    optionCode.EntityData.SegmentPath = "option-code" + types.AddKeyToken(optionCode.OptionCode, "option-code")
    optionCode.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/server/option-codes/" + optionCode.EntityData.SegmentPath
    optionCode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optionCode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optionCode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optionCode.EntityData.Children = types.NewOrderedMap()
    optionCode.EntityData.Leafs = types.NewOrderedMap()
    optionCode.EntityData.Leafs.Append("option-code", types.YLeaf{"OptionCode", optionCode.OptionCode})
    optionCode.EntityData.Leafs.Append("ascii-string", types.YLeaf{"AsciiString", optionCode.AsciiString})
    optionCode.EntityData.Leafs.Append("hex-string", types.YLeaf{"HexString", optionCode.HexString})
    optionCode.EntityData.Leafs.Append("force-insert", types.YLeaf{"ForceInsert", optionCode.ForceInsert})
    optionCode.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", optionCode.IpAddress})

    optionCode.EntityData.YListKeys = []string {"OptionCode"}

    return &(optionCode.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay
// DHCP Relay profile
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Action to take if L2 header source Mac and dhcp header mac address don't
    // match. The type is MacMismatchAction.
    MacMismatchAction interface{}

    // GIADDR policy.
    GiAddrPolicy Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy

    // VRF Helper Addresses.
    Vrfs Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs

    // Relay agent information option.
    RelayInformationOption Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption

    // Broadcast Flag policy.
    BroadcastPolicy Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy
}

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay) GetEntityData() *types.CommonEntityData {
    relay.EntityData.YFilter = relay.YFilter
    relay.EntityData.YangName = "relay"
    relay.EntityData.BundleName = "cisco_ios_xr"
    relay.EntityData.ParentYangName = "mode"
    relay.EntityData.SegmentPath = "relay"
    relay.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/" + relay.EntityData.SegmentPath
    relay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relay.EntityData.Children = types.NewOrderedMap()
    relay.EntityData.Children.Append("gi-addr-policy", types.YChild{"GiAddrPolicy", &relay.GiAddrPolicy})
    relay.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &relay.Vrfs})
    relay.EntityData.Children.Append("relay-information-option", types.YChild{"RelayInformationOption", &relay.RelayInformationOption})
    relay.EntityData.Children.Append("broadcast-policy", types.YChild{"BroadcastPolicy", &relay.BroadcastPolicy})
    relay.EntityData.Leafs = types.NewOrderedMap()
    relay.EntityData.Leafs.Append("mac-mismatch-action", types.YLeaf{"MacMismatchAction", relay.MacMismatchAction})

    relay.EntityData.YListKeys = []string {}

    return &(relay.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy
// GIADDR policy
// This type is a presence type.
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // GIADDR policy. The type is Ipv4dhcpdRelayGiaddrPolicy. This attribute is
    // mandatory.
    Policy interface{}
}

func (giAddrPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy) GetEntityData() *types.CommonEntityData {
    giAddrPolicy.EntityData.YFilter = giAddrPolicy.YFilter
    giAddrPolicy.EntityData.YangName = "gi-addr-policy"
    giAddrPolicy.EntityData.BundleName = "cisco_ios_xr"
    giAddrPolicy.EntityData.ParentYangName = "relay"
    giAddrPolicy.EntityData.SegmentPath = "gi-addr-policy"
    giAddrPolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/relay/" + giAddrPolicy.EntityData.SegmentPath
    giAddrPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    giAddrPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    giAddrPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    giAddrPolicy.EntityData.Children = types.NewOrderedMap()
    giAddrPolicy.EntityData.Leafs = types.NewOrderedMap()
    giAddrPolicy.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", giAddrPolicy.Policy})

    giAddrPolicy.EntityData.YListKeys = []string {}

    return &(giAddrPolicy.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs
// VRF Helper Addresses
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF Name. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf.
    Vrf []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "relay"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/relay/" + vrfs.EntityData.SegmentPath
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

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf
// VRF Name
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Helper Addresses.
    HelperAddresses Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/relay/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("helper-addresses", types.YChild{"HelperAddresses", &vrf.HelperAddresses})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses
// Helper Addresses
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Helper Address. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress.
    HelperAddress []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses) GetEntityData() *types.CommonEntityData {
    helperAddresses.EntityData.YFilter = helperAddresses.YFilter
    helperAddresses.EntityData.YangName = "helper-addresses"
    helperAddresses.EntityData.BundleName = "cisco_ios_xr"
    helperAddresses.EntityData.ParentYangName = "vrf"
    helperAddresses.EntityData.SegmentPath = "helper-addresses"
    helperAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/relay/vrfs/vrf/" + helperAddresses.EntityData.SegmentPath
    helperAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddresses.EntityData.Children = types.NewOrderedMap()
    helperAddresses.EntityData.Children.Append("helper-address", types.YChild{"HelperAddress", nil})
    for i := range helperAddresses.HelperAddress {
        helperAddresses.EntityData.Children.Append(types.GetSegmentPath(helperAddresses.HelperAddress[i]), types.YChild{"HelperAddress", helperAddresses.HelperAddress[i]})
    }
    helperAddresses.EntityData.Leafs = types.NewOrderedMap()

    helperAddresses.EntityData.YListKeys = []string {}

    return &(helperAddresses.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress
// Helper Address
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPV4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Enable helper - deprecated. The type is interface{}.
    Enable interface{}

    // GatewayAddress. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    GatewayAddress interface{}
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress) GetEntityData() *types.CommonEntityData {
    helperAddress.EntityData.YFilter = helperAddress.YFilter
    helperAddress.EntityData.YangName = "helper-address"
    helperAddress.EntityData.BundleName = "cisco_ios_xr"
    helperAddress.EntityData.ParentYangName = "helper-addresses"
    helperAddress.EntityData.SegmentPath = "helper-address" + types.AddKeyToken(helperAddress.IpAddress, "ip-address")
    helperAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/relay/vrfs/vrf/helper-addresses/" + helperAddress.EntityData.SegmentPath
    helperAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddress.EntityData.Children = types.NewOrderedMap()
    helperAddress.EntityData.Leafs = types.NewOrderedMap()
    helperAddress.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", helperAddress.IpAddress})
    helperAddress.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", helperAddress.Enable})
    helperAddress.EntityData.Leafs.Append("gateway-address", types.YLeaf{"GatewayAddress", helperAddress.GatewayAddress})

    helperAddress.EntityData.YListKeys = []string {"IpAddress"}

    return &(helperAddress.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption
// Relay agent information option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VPN Mode. The type is Ipv4dhcpdRelayInfoOptionvpnMode.
    VpnMode interface{}

    // Subscriber ID. The type is string.
    SubscriberId interface{}

    // Insert Relay Agent Information circuit ID and remote ID suboptions in
    // client requests. The type is interface{}.
    Insert interface{}

    // Check Relay Agent Information Option in server reply. The type is
    // interface{}.
    Check interface{}

    // Insert VPN options. The type is interface{}.
    Vpn interface{}

    // Forward untrusted packets. The type is interface{}.
    AllowUntrusted interface{}

    // Relay information option policy. The type is
    // Ipv4dhcpdRelayInfoOptionPolicy.
    Policy interface{}
}

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption) GetEntityData() *types.CommonEntityData {
    relayInformationOption.EntityData.YFilter = relayInformationOption.YFilter
    relayInformationOption.EntityData.YangName = "relay-information-option"
    relayInformationOption.EntityData.BundleName = "cisco_ios_xr"
    relayInformationOption.EntityData.ParentYangName = "relay"
    relayInformationOption.EntityData.SegmentPath = "relay-information-option"
    relayInformationOption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/relay/" + relayInformationOption.EntityData.SegmentPath
    relayInformationOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayInformationOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayInformationOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayInformationOption.EntityData.Children = types.NewOrderedMap()
    relayInformationOption.EntityData.Leafs = types.NewOrderedMap()
    relayInformationOption.EntityData.Leafs.Append("vpn-mode", types.YLeaf{"VpnMode", relayInformationOption.VpnMode})
    relayInformationOption.EntityData.Leafs.Append("subscriber-id", types.YLeaf{"SubscriberId", relayInformationOption.SubscriberId})
    relayInformationOption.EntityData.Leafs.Append("insert", types.YLeaf{"Insert", relayInformationOption.Insert})
    relayInformationOption.EntityData.Leafs.Append("check", types.YLeaf{"Check", relayInformationOption.Check})
    relayInformationOption.EntityData.Leafs.Append("vpn", types.YLeaf{"Vpn", relayInformationOption.Vpn})
    relayInformationOption.EntityData.Leafs.Append("allow-untrusted", types.YLeaf{"AllowUntrusted", relayInformationOption.AllowUntrusted})
    relayInformationOption.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", relayInformationOption.Policy})

    relayInformationOption.EntityData.YListKeys = []string {}

    return &(relayInformationOption.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy
// Broadcast Flag policy
// This type is a presence type.
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Broadcast flag policy. The type is Ipv4dhcpdBroadcastFlagPolicy. This
    // attribute is mandatory.
    Policy interface{}
}

func (broadcastPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy) GetEntityData() *types.CommonEntityData {
    broadcastPolicy.EntityData.YFilter = broadcastPolicy.YFilter
    broadcastPolicy.EntityData.YangName = "broadcast-policy"
    broadcastPolicy.EntityData.BundleName = "cisco_ios_xr"
    broadcastPolicy.EntityData.ParentYangName = "relay"
    broadcastPolicy.EntityData.SegmentPath = "broadcast-policy"
    broadcastPolicy.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/relay/" + broadcastPolicy.EntityData.SegmentPath
    broadcastPolicy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    broadcastPolicy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    broadcastPolicy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    broadcastPolicy.EntityData.Children = types.NewOrderedMap()
    broadcastPolicy.EntityData.Leafs = types.NewOrderedMap()
    broadcastPolicy.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", broadcastPolicy.Policy})

    broadcastPolicy.EntityData.YListKeys = []string {}

    return &(broadcastPolicy.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy
// DHCP proxy profile
// This type is a presence type.
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Allow dhcp subscriber move. The type is interface{}.
    ProxyAllowMove interface{}

    // DHCP IPV4 profile proxy secure-arp enable. The type is interface{}.
    SecureArp interface{}

    // For BNG session, delay the authentication. The type is interface{}.
    DelayedAuthenProxy interface{}

    // DHCP IPV4 profile mode enable. The type is interface{}. This attribute is
    // mandatory.
    Enable interface{}

    // Specify gateway address policy.
    Giaddr Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr

    // DHCP class table.
    Classes Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes

    // Authentication Username formating.
    AuthUsername Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername

    // Relay agent information option.
    RelayInformation Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation

    // Enable to provide the list of options need to send to aaa.
    DhcpToAaa Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_DhcpToAaa

    // List of VRFs.
    Vrfs Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs

    // Change sessions configuration.
    Sessions Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions

    // Proxy limit lease.
    LimitLease Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease

    // DHCPv4 lease proxy.
    LeaseProxy Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy

    // Specify broadcast flag.
    BroadcastFlag Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag

    // Insert match keyword.
    Match Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match
}

func (proxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy) GetEntityData() *types.CommonEntityData {
    proxy.EntityData.YFilter = proxy.YFilter
    proxy.EntityData.YangName = "proxy"
    proxy.EntityData.BundleName = "cisco_ios_xr"
    proxy.EntityData.ParentYangName = "mode"
    proxy.EntityData.SegmentPath = "proxy"
    proxy.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/" + proxy.EntityData.SegmentPath
    proxy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proxy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proxy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proxy.EntityData.Children = types.NewOrderedMap()
    proxy.EntityData.Children.Append("giaddr", types.YChild{"Giaddr", &proxy.Giaddr})
    proxy.EntityData.Children.Append("classes", types.YChild{"Classes", &proxy.Classes})
    proxy.EntityData.Children.Append("auth-username", types.YChild{"AuthUsername", &proxy.AuthUsername})
    proxy.EntityData.Children.Append("relay-information", types.YChild{"RelayInformation", &proxy.RelayInformation})
    proxy.EntityData.Children.Append("dhcp-to-aaa", types.YChild{"DhcpToAaa", &proxy.DhcpToAaa})
    proxy.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &proxy.Vrfs})
    proxy.EntityData.Children.Append("sessions", types.YChild{"Sessions", &proxy.Sessions})
    proxy.EntityData.Children.Append("limit-lease", types.YChild{"LimitLease", &proxy.LimitLease})
    proxy.EntityData.Children.Append("lease-proxy", types.YChild{"LeaseProxy", &proxy.LeaseProxy})
    proxy.EntityData.Children.Append("broadcast-flag", types.YChild{"BroadcastFlag", &proxy.BroadcastFlag})
    proxy.EntityData.Children.Append("match", types.YChild{"Match", &proxy.Match})
    proxy.EntityData.Leafs = types.NewOrderedMap()
    proxy.EntityData.Leafs.Append("proxy-allow-move", types.YLeaf{"ProxyAllowMove", proxy.ProxyAllowMove})
    proxy.EntityData.Leafs.Append("secure-arp", types.YLeaf{"SecureArp", proxy.SecureArp})
    proxy.EntityData.Leafs.Append("delayed-authen-proxy", types.YLeaf{"DelayedAuthenProxy", proxy.DelayedAuthenProxy})
    proxy.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", proxy.Enable})

    proxy.EntityData.YListKeys = []string {}

    return &(proxy.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr
// Specify gateway address policy
// This type is a presence type.
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Gateway address policy. The type is Ipv4dhcpdGiaddrPolicy. This attribute
    // is mandatory.
    Policy interface{}
}

func (giaddr *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr) GetEntityData() *types.CommonEntityData {
    giaddr.EntityData.YFilter = giaddr.YFilter
    giaddr.EntityData.YangName = "giaddr"
    giaddr.EntityData.BundleName = "cisco_ios_xr"
    giaddr.EntityData.ParentYangName = "proxy"
    giaddr.EntityData.SegmentPath = "giaddr"
    giaddr.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/" + giaddr.EntityData.SegmentPath
    giaddr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    giaddr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    giaddr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    giaddr.EntityData.Children = types.NewOrderedMap()
    giaddr.EntityData.Leafs = types.NewOrderedMap()
    giaddr.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", giaddr.Policy})

    giaddr.EntityData.YListKeys = []string {}

    return &(giaddr.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes
// DHCP class table
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP class. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class.
    Class []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class
}

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes) GetEntityData() *types.CommonEntityData {
    classes.EntityData.YFilter = classes.YFilter
    classes.EntityData.YangName = "classes"
    classes.EntityData.BundleName = "cisco_ios_xr"
    classes.EntityData.ParentYangName = "proxy"
    classes.EntityData.SegmentPath = "classes"
    classes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/" + classes.EntityData.SegmentPath
    classes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classes.EntityData.Children = types.NewOrderedMap()
    classes.EntityData.Children.Append("class", types.YChild{"Class", nil})
    for i := range classes.Class {
        classes.EntityData.Children.Append(types.GetSegmentPath(classes.Class[i]), types.YChild{"Class", classes.Class[i]})
    }
    classes.EntityData.Leafs = types.NewOrderedMap()

    classes.EntityData.YListKeys = []string {}

    return &(classes.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class
// DHCP class
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Class name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClassName interface{}

    // Enable the DHCP IPV4 proxy class. The type is interface{}. This attribute
    // is mandatory.
    Enable interface{}

    // Match option.
    Match Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match

    // List of VRFs.
    Vrfs Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs
}

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + types.AddKeyToken(class.ClassName, "class-name")
    class.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/classes/" + class.EntityData.SegmentPath
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = types.NewOrderedMap()
    class.EntityData.Children.Append("match", types.YChild{"Match", &class.Match})
    class.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &class.Vrfs})
    class.EntityData.Leafs = types.NewOrderedMap()
    class.EntityData.Leafs.Append("class-name", types.YLeaf{"ClassName", class.ClassName})
    class.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", class.Enable})

    class.EntityData.YListKeys = []string {"ClassName"}

    return &(class.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match
// Match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Match VRF name. The type is string.
    Vrf interface{}

    // Match option.
    Option Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match) GetEntityData() *types.CommonEntityData {
    match.EntityData.YFilter = match.YFilter
    match.EntityData.YangName = "match"
    match.EntityData.BundleName = "cisco_ios_xr"
    match.EntityData.ParentYangName = "class"
    match.EntityData.SegmentPath = "match"
    match.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/classes/class/" + match.EntityData.SegmentPath
    match.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    match.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    match.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    match.EntityData.Children = types.NewOrderedMap()
    match.EntityData.Children.Append("option", types.YChild{"Option", &match.Option})
    match.EntityData.Leafs = types.NewOrderedMap()
    match.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", match.Vrf})

    match.EntityData.YListKeys = []string {}

    return &(match.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option
// Match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Match option. The type is Dhcpv4MatchOption.
    OptionType interface{}

    // Hex pattern string. The type is string.
    Pattern interface{}

    // Bit mask pattern. The type is string.
    BitMask interface{}
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option) GetEntityData() *types.CommonEntityData {
    option.EntityData.YFilter = option.YFilter
    option.EntityData.YangName = "option"
    option.EntityData.BundleName = "cisco_ios_xr"
    option.EntityData.ParentYangName = "match"
    option.EntityData.SegmentPath = "option"
    option.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/classes/class/match/" + option.EntityData.SegmentPath
    option.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    option.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    option.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    option.EntityData.Children = types.NewOrderedMap()
    option.EntityData.Leafs = types.NewOrderedMap()
    option.EntityData.Leafs.Append("option-type", types.YLeaf{"OptionType", option.OptionType})
    option.EntityData.Leafs.Append("pattern", types.YLeaf{"Pattern", option.Pattern})
    option.EntityData.Leafs.Append("bit-mask", types.YLeaf{"BitMask", option.BitMask})

    option.EntityData.YListKeys = []string {}

    return &(option.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs
// List of VRFs
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf.
    Vrf []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "class"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/classes/class/" + vrfs.EntityData.SegmentPath
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

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf
// VRF name
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Helper addresses.
    HelperAddresses Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/classes/class/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("helper-addresses", types.YChild{"HelperAddresses", &vrf.HelperAddresses})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses
// Helper addresses
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Helper address. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress.
    HelperAddress []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses) GetEntityData() *types.CommonEntityData {
    helperAddresses.EntityData.YFilter = helperAddresses.YFilter
    helperAddresses.EntityData.YangName = "helper-addresses"
    helperAddresses.EntityData.BundleName = "cisco_ios_xr"
    helperAddresses.EntityData.ParentYangName = "vrf"
    helperAddresses.EntityData.SegmentPath = "helper-addresses"
    helperAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/classes/class/vrfs/vrf/" + helperAddresses.EntityData.SegmentPath
    helperAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddresses.EntityData.Children = types.NewOrderedMap()
    helperAddresses.EntityData.Children.Append("helper-address", types.YChild{"HelperAddress", nil})
    for i := range helperAddresses.HelperAddress {
        helperAddresses.EntityData.Children.Append(types.GetSegmentPath(helperAddresses.HelperAddress[i]), types.YChild{"HelperAddress", helperAddresses.HelperAddress[i]})
    }
    helperAddresses.EntityData.Leafs = types.NewOrderedMap()

    helperAddresses.EntityData.YListKeys = []string {}

    return &(helperAddresses.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress
// Helper address
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ServerAddress interface{}

    // Gateway address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    GatewayAddress interface{}
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress) GetEntityData() *types.CommonEntityData {
    helperAddress.EntityData.YFilter = helperAddress.YFilter
    helperAddress.EntityData.YangName = "helper-address"
    helperAddress.EntityData.BundleName = "cisco_ios_xr"
    helperAddress.EntityData.ParentYangName = "helper-addresses"
    helperAddress.EntityData.SegmentPath = "helper-address" + types.AddKeyToken(helperAddress.ServerAddress, "server-address")
    helperAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/classes/class/vrfs/vrf/helper-addresses/" + helperAddress.EntityData.SegmentPath
    helperAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddress.EntityData.Children = types.NewOrderedMap()
    helperAddress.EntityData.Leafs = types.NewOrderedMap()
    helperAddress.EntityData.Leafs.Append("server-address", types.YLeaf{"ServerAddress", helperAddress.ServerAddress})
    helperAddress.EntityData.Leafs.Append("gateway-address", types.YLeaf{"GatewayAddress", helperAddress.GatewayAddress})

    helperAddress.EntityData.YListKeys = []string {"ServerAddress"}

    return &(helperAddress.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername
// Authentication Username formating
// This type is a presence type.
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Username Formatting first argument . The type is Dhcpv4AuthUsername. This
    // attribute is mandatory.
    Arg1 interface{}

    // Username Formatting second argument . The type is Dhcpv4AuthUsername.
    Arg2 interface{}
}

func (authUsername *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername) GetEntityData() *types.CommonEntityData {
    authUsername.EntityData.YFilter = authUsername.YFilter
    authUsername.EntityData.YangName = "auth-username"
    authUsername.EntityData.BundleName = "cisco_ios_xr"
    authUsername.EntityData.ParentYangName = "proxy"
    authUsername.EntityData.SegmentPath = "auth-username"
    authUsername.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/" + authUsername.EntityData.SegmentPath
    authUsername.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authUsername.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authUsername.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authUsername.EntityData.Children = types.NewOrderedMap()
    authUsername.EntityData.Leafs = types.NewOrderedMap()
    authUsername.EntityData.Leafs.Append("arg1", types.YLeaf{"Arg1", authUsername.Arg1})
    authUsername.EntityData.Leafs.Append("arg2", types.YLeaf{"Arg2", authUsername.Arg2})

    authUsername.EntityData.YListKeys = []string {}

    return &(authUsername.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation
// Relay agent information option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Insert relay rgent information circuit ID and remote ID suboptions in
    // client requests. The type is interface{}.
    Option interface{}

    // Insert VPN options. The type is interface{}.
    Vpn interface{}

    // Forward untrusted packets. The type is interface{}.
    AllowUntrusted interface{}

    // Insert Circuit-id sub-option. The type is interface{}.
    CircuitId interface{}

    // Relay information option policy. The type is
    // Ipv4dhcpdRelayInfoOptionPolicy.
    Policy interface{}

    // VPN Mode. The type is Ipv4dhcpdRelayInfoOptionvpnMode.
    VpnMode interface{}

    // Insert Remote-id sub-option. The type is interface{}.
    RemoteIdXr interface{}

    // Suppress Remote ID. The type is interface{}.
    RemoteIdSuppress interface{}

    // Check relay agent information option in server reply. The type is
    // interface{}.
    Check interface{}

    // Remote ID. The type is string.
    RemoteId interface{}

    // Relay information option authenticate. The type is
    // Ipv4dhcpdRelayInfoOptionAuthenticate.
    Authenticate interface{}
}

func (relayInformation *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation) GetEntityData() *types.CommonEntityData {
    relayInformation.EntityData.YFilter = relayInformation.YFilter
    relayInformation.EntityData.YangName = "relay-information"
    relayInformation.EntityData.BundleName = "cisco_ios_xr"
    relayInformation.EntityData.ParentYangName = "proxy"
    relayInformation.EntityData.SegmentPath = "relay-information"
    relayInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/" + relayInformation.EntityData.SegmentPath
    relayInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayInformation.EntityData.Children = types.NewOrderedMap()
    relayInformation.EntityData.Leafs = types.NewOrderedMap()
    relayInformation.EntityData.Leafs.Append("option", types.YLeaf{"Option", relayInformation.Option})
    relayInformation.EntityData.Leafs.Append("vpn", types.YLeaf{"Vpn", relayInformation.Vpn})
    relayInformation.EntityData.Leafs.Append("allow-untrusted", types.YLeaf{"AllowUntrusted", relayInformation.AllowUntrusted})
    relayInformation.EntityData.Leafs.Append("circuit-id", types.YLeaf{"CircuitId", relayInformation.CircuitId})
    relayInformation.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", relayInformation.Policy})
    relayInformation.EntityData.Leafs.Append("vpn-mode", types.YLeaf{"VpnMode", relayInformation.VpnMode})
    relayInformation.EntityData.Leafs.Append("remote-id-xr", types.YLeaf{"RemoteIdXr", relayInformation.RemoteIdXr})
    relayInformation.EntityData.Leafs.Append("remote-id-suppress", types.YLeaf{"RemoteIdSuppress", relayInformation.RemoteIdSuppress})
    relayInformation.EntityData.Leafs.Append("check", types.YLeaf{"Check", relayInformation.Check})
    relayInformation.EntityData.Leafs.Append("remote-id", types.YLeaf{"RemoteId", relayInformation.RemoteId})
    relayInformation.EntityData.Leafs.Append("authenticate", types.YLeaf{"Authenticate", relayInformation.Authenticate})

    relayInformation.EntityData.YListKeys = []string {}

    return &(relayInformation.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_DhcpToAaa
// Enable to provide the list of options need
// to send to aaa
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_DhcpToAaa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // option type.
    Option Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_DhcpToAaa_Option
}

func (dhcpToAaa *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_DhcpToAaa) GetEntityData() *types.CommonEntityData {
    dhcpToAaa.EntityData.YFilter = dhcpToAaa.YFilter
    dhcpToAaa.EntityData.YangName = "dhcp-to-aaa"
    dhcpToAaa.EntityData.BundleName = "cisco_ios_xr"
    dhcpToAaa.EntityData.ParentYangName = "proxy"
    dhcpToAaa.EntityData.SegmentPath = "dhcp-to-aaa"
    dhcpToAaa.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/" + dhcpToAaa.EntityData.SegmentPath
    dhcpToAaa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpToAaa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpToAaa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpToAaa.EntityData.Children = types.NewOrderedMap()
    dhcpToAaa.EntityData.Children.Append("option", types.YChild{"Option", &dhcpToAaa.Option})
    dhcpToAaa.EntityData.Leafs = types.NewOrderedMap()

    dhcpToAaa.EntityData.YListKeys = []string {}

    return &(dhcpToAaa.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_DhcpToAaa_Option
// option type
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_DhcpToAaa_Option struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of options.
    List Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_DhcpToAaa_Option_List
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_DhcpToAaa_Option) GetEntityData() *types.CommonEntityData {
    option.EntityData.YFilter = option.YFilter
    option.EntityData.YangName = "option"
    option.EntityData.BundleName = "cisco_ios_xr"
    option.EntityData.ParentYangName = "dhcp-to-aaa"
    option.EntityData.SegmentPath = "option"
    option.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/dhcp-to-aaa/" + option.EntityData.SegmentPath
    option.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    option.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    option.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    option.EntityData.Children = types.NewOrderedMap()
    option.EntityData.Children.Append("list", types.YChild{"List", &option.List})
    option.EntityData.Leafs = types.NewOrderedMap()

    option.EntityData.YListKeys = []string {}

    return &(option.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_DhcpToAaa_Option_List
// List of options
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_DhcpToAaa_Option_List struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // option all. The type is interface{} with range: 0..4294967295.
    OptionAll interface{}

    // List of options. The type is slice of interface{} with range: 1..255.
    Option []interface{}
}

func (list *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_DhcpToAaa_Option_List) GetEntityData() *types.CommonEntityData {
    list.EntityData.YFilter = list.YFilter
    list.EntityData.YangName = "list"
    list.EntityData.BundleName = "cisco_ios_xr"
    list.EntityData.ParentYangName = "option"
    list.EntityData.SegmentPath = "list"
    list.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/dhcp-to-aaa/option/" + list.EntityData.SegmentPath
    list.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    list.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    list.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    list.EntityData.Children = types.NewOrderedMap()
    list.EntityData.Leafs = types.NewOrderedMap()
    list.EntityData.Leafs.Append("option-all", types.YLeaf{"OptionAll", list.OptionAll})
    list.EntityData.Leafs.Append("option", types.YLeaf{"Option", list.Option})

    list.EntityData.YListKeys = []string {}

    return &(list.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs
// List of VRFs
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf.
    Vrf []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "proxy"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/" + vrfs.EntityData.SegmentPath
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

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf
// VRF name
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Helper addresses.
    HelperAddresses Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("helper-addresses", types.YChild{"HelperAddresses", &vrf.HelperAddresses})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses
// Helper addresses
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Helper address. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress.
    HelperAddress []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses) GetEntityData() *types.CommonEntityData {
    helperAddresses.EntityData.YFilter = helperAddresses.YFilter
    helperAddresses.EntityData.YangName = "helper-addresses"
    helperAddresses.EntityData.BundleName = "cisco_ios_xr"
    helperAddresses.EntityData.ParentYangName = "vrf"
    helperAddresses.EntityData.SegmentPath = "helper-addresses"
    helperAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/vrfs/vrf/" + helperAddresses.EntityData.SegmentPath
    helperAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddresses.EntityData.Children = types.NewOrderedMap()
    helperAddresses.EntityData.Children.Append("helper-address", types.YChild{"HelperAddress", nil})
    for i := range helperAddresses.HelperAddress {
        helperAddresses.EntityData.Children.Append(types.GetSegmentPath(helperAddresses.HelperAddress[i]), types.YChild{"HelperAddress", helperAddresses.HelperAddress[i]})
    }
    helperAddresses.EntityData.Leafs = types.NewOrderedMap()

    helperAddresses.EntityData.YListKeys = []string {}

    return &(helperAddresses.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress
// Helper address
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ServerAddress interface{}

    // Gateway address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    GatewayAddress interface{}
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetEntityData() *types.CommonEntityData {
    helperAddress.EntityData.YFilter = helperAddress.YFilter
    helperAddress.EntityData.YangName = "helper-address"
    helperAddress.EntityData.BundleName = "cisco_ios_xr"
    helperAddress.EntityData.ParentYangName = "helper-addresses"
    helperAddress.EntityData.SegmentPath = "helper-address" + types.AddKeyToken(helperAddress.ServerAddress, "server-address")
    helperAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/vrfs/vrf/helper-addresses/" + helperAddress.EntityData.SegmentPath
    helperAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    helperAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    helperAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    helperAddress.EntityData.Children = types.NewOrderedMap()
    helperAddress.EntityData.Leafs = types.NewOrderedMap()
    helperAddress.EntityData.Leafs.Append("server-address", types.YLeaf{"ServerAddress", helperAddress.ServerAddress})
    helperAddress.EntityData.Leafs.Append("gateway-address", types.YLeaf{"GatewayAddress", helperAddress.GatewayAddress})

    helperAddress.EntityData.YListKeys = []string {"ServerAddress"}

    return &(helperAddress.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions
// Change sessions configuration
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Throttle DHCP sessions based on MAC address.
    ProxyThrottleType Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType
}

func (sessions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions) GetEntityData() *types.CommonEntityData {
    sessions.EntityData.YFilter = sessions.YFilter
    sessions.EntityData.YangName = "sessions"
    sessions.EntityData.BundleName = "cisco_ios_xr"
    sessions.EntityData.ParentYangName = "proxy"
    sessions.EntityData.SegmentPath = "sessions"
    sessions.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/" + sessions.EntityData.SegmentPath
    sessions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sessions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sessions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sessions.EntityData.Children = types.NewOrderedMap()
    sessions.EntityData.Children.Append("proxy-throttle-type", types.YChild{"ProxyThrottleType", &sessions.ProxyThrottleType})
    sessions.EntityData.Leafs = types.NewOrderedMap()

    sessions.EntityData.YListKeys = []string {}

    return &(sessions.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType
// Throttle DHCP sessions based on MAC
// address
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Throttle DHCP sessions from any one MAC address.
    ProxyMacThrottle Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle
}

func (proxyThrottleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType) GetEntityData() *types.CommonEntityData {
    proxyThrottleType.EntityData.YFilter = proxyThrottleType.YFilter
    proxyThrottleType.EntityData.YangName = "proxy-throttle-type"
    proxyThrottleType.EntityData.BundleName = "cisco_ios_xr"
    proxyThrottleType.EntityData.ParentYangName = "sessions"
    proxyThrottleType.EntityData.SegmentPath = "proxy-throttle-type"
    proxyThrottleType.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/sessions/" + proxyThrottleType.EntityData.SegmentPath
    proxyThrottleType.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proxyThrottleType.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proxyThrottleType.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proxyThrottleType.EntityData.Children = types.NewOrderedMap()
    proxyThrottleType.EntityData.Children.Append("proxy-mac-throttle", types.YChild{"ProxyMacThrottle", &proxyThrottleType.ProxyMacThrottle})
    proxyThrottleType.EntityData.Leafs = types.NewOrderedMap()

    proxyThrottleType.EntityData.YListKeys = []string {}

    return &(proxyThrottleType.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle
// Throttle DHCP sessions from any one MAC
// address
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of discovers at which to throttle. The type is interface{} with
    // range: 1..65535.
    NumDiscover interface{}

    // Throttle request period (in secs). The type is interface{} with range:
    // 1..100. Units are second.
    NumRequest interface{}

    // Throttle blocking period (in secs). The type is interface{} with range:
    // 1..100. Units are second.
    NumBlock interface{}
}

func (proxyMacThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle) GetEntityData() *types.CommonEntityData {
    proxyMacThrottle.EntityData.YFilter = proxyMacThrottle.YFilter
    proxyMacThrottle.EntityData.YangName = "proxy-mac-throttle"
    proxyMacThrottle.EntityData.BundleName = "cisco_ios_xr"
    proxyMacThrottle.EntityData.ParentYangName = "proxy-throttle-type"
    proxyMacThrottle.EntityData.SegmentPath = "proxy-mac-throttle"
    proxyMacThrottle.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/sessions/proxy-throttle-type/" + proxyMacThrottle.EntityData.SegmentPath
    proxyMacThrottle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proxyMacThrottle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proxyMacThrottle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proxyMacThrottle.EntityData.Children = types.NewOrderedMap()
    proxyMacThrottle.EntityData.Leafs = types.NewOrderedMap()
    proxyMacThrottle.EntityData.Leafs.Append("num-discover", types.YLeaf{"NumDiscover", proxyMacThrottle.NumDiscover})
    proxyMacThrottle.EntityData.Leafs.Append("num-request", types.YLeaf{"NumRequest", proxyMacThrottle.NumRequest})
    proxyMacThrottle.EntityData.Leafs.Append("num-block", types.YLeaf{"NumBlock", proxyMacThrottle.NumBlock})

    proxyMacThrottle.EntityData.YListKeys = []string {}

    return &(proxyMacThrottle.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease
// Proxy limit lease
// This type is a presence type.
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Lease limit type. The type is Dhcpv4LimitLease1. This attribute is
    // mandatory.
    LimitType interface{}

    // Limit lease count. The type is interface{} with range: 1..240000. This
    // attribute is mandatory.
    LimitLeaseCount interface{}
}

func (limitLease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease) GetEntityData() *types.CommonEntityData {
    limitLease.EntityData.YFilter = limitLease.YFilter
    limitLease.EntityData.YangName = "limit-lease"
    limitLease.EntityData.BundleName = "cisco_ios_xr"
    limitLease.EntityData.ParentYangName = "proxy"
    limitLease.EntityData.SegmentPath = "limit-lease"
    limitLease.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/" + limitLease.EntityData.SegmentPath
    limitLease.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    limitLease.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    limitLease.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    limitLease.EntityData.Children = types.NewOrderedMap()
    limitLease.EntityData.Leafs = types.NewOrderedMap()
    limitLease.EntityData.Leafs.Append("limit-type", types.YLeaf{"LimitType", limitLease.LimitType})
    limitLease.EntityData.Leafs.Append("limit-lease-count", types.YLeaf{"LimitLeaseCount", limitLease.LimitLeaseCount})

    limitLease.EntityData.YListKeys = []string {}

    return &(limitLease.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy
// DHCPv4 lease proxy
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify client lease proxy time. The type is interface{} with range:
    // 300..4294967295.
    ClientLeaseTime interface{}

    // Set DHCP server sent options in lease proxy generating ACK. The type is
    // interface{}.
    SetServerOptions interface{}
}

func (leaseProxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy) GetEntityData() *types.CommonEntityData {
    leaseProxy.EntityData.YFilter = leaseProxy.YFilter
    leaseProxy.EntityData.YangName = "lease-proxy"
    leaseProxy.EntityData.BundleName = "cisco_ios_xr"
    leaseProxy.EntityData.ParentYangName = "proxy"
    leaseProxy.EntityData.SegmentPath = "lease-proxy"
    leaseProxy.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/" + leaseProxy.EntityData.SegmentPath
    leaseProxy.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    leaseProxy.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    leaseProxy.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    leaseProxy.EntityData.Children = types.NewOrderedMap()
    leaseProxy.EntityData.Leafs = types.NewOrderedMap()
    leaseProxy.EntityData.Leafs.Append("client-lease-time", types.YLeaf{"ClientLeaseTime", leaseProxy.ClientLeaseTime})
    leaseProxy.EntityData.Leafs.Append("set-server-options", types.YLeaf{"SetServerOptions", leaseProxy.SetServerOptions})

    leaseProxy.EntityData.YListKeys = []string {}

    return &(leaseProxy.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag
// Specify broadcast flag
// This type is a presence type.
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Broadcast flag policy. The type is Ipv4dhcpdBroadcastFlagPolicy. This
    // attribute is mandatory.
    Policy interface{}
}

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag) GetEntityData() *types.CommonEntityData {
    broadcastFlag.EntityData.YFilter = broadcastFlag.YFilter
    broadcastFlag.EntityData.YangName = "broadcast-flag"
    broadcastFlag.EntityData.BundleName = "cisco_ios_xr"
    broadcastFlag.EntityData.ParentYangName = "proxy"
    broadcastFlag.EntityData.SegmentPath = "broadcast-flag"
    broadcastFlag.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/" + broadcastFlag.EntityData.SegmentPath
    broadcastFlag.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    broadcastFlag.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    broadcastFlag.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    broadcastFlag.EntityData.Children = types.NewOrderedMap()
    broadcastFlag.EntityData.Leafs = types.NewOrderedMap()
    broadcastFlag.EntityData.Leafs.Append("policy", types.YLeaf{"Policy", broadcastFlag.Policy})

    broadcastFlag.EntityData.YListKeys = []string {}

    return &(broadcastFlag.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match
// Insert match keyword
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of Option.
    DefOptions Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions

    // Table of Option.
    OptionFilters Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match) GetEntityData() *types.CommonEntityData {
    match.EntityData.YFilter = match.YFilter
    match.EntityData.YangName = "match"
    match.EntityData.BundleName = "cisco_ios_xr"
    match.EntityData.ParentYangName = "proxy"
    match.EntityData.SegmentPath = "match"
    match.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/" + match.EntityData.SegmentPath
    match.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    match.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    match.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    match.EntityData.Children = types.NewOrderedMap()
    match.EntityData.Children.Append("def-options", types.YChild{"DefOptions", &match.DefOptions})
    match.EntityData.Children.Append("option-filters", types.YChild{"OptionFilters", &match.OptionFilters})
    match.EntityData.Leafs = types.NewOrderedMap()

    match.EntityData.YListKeys = []string {}

    return &(match.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions
// Table of Option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify match option. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption.
    DefOption []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption
}

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions) GetEntityData() *types.CommonEntityData {
    defOptions.EntityData.YFilter = defOptions.YFilter
    defOptions.EntityData.YangName = "def-options"
    defOptions.EntityData.BundleName = "cisco_ios_xr"
    defOptions.EntityData.ParentYangName = "match"
    defOptions.EntityData.SegmentPath = "def-options"
    defOptions.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/match/" + defOptions.EntityData.SegmentPath
    defOptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defOptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defOptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defOptions.EntityData.Children = types.NewOrderedMap()
    defOptions.EntityData.Children.Append("def-option", types.YChild{"DefOption", nil})
    for i := range defOptions.DefOption {
        defOptions.EntityData.Children.Append(types.GetSegmentPath(defOptions.DefOption[i]), types.YChild{"DefOption", defOptions.DefOption[i]})
    }
    defOptions.EntityData.Leafs = types.NewOrderedMap()

    defOptions.EntityData.YListKeys = []string {}

    return &(defOptions.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption
// Specify match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Match option 60. The type is interface{} with
    // range: 0..4294967295.
    DefMatchoption interface{}

    // Vendor action. The type is ProxyAction. This attribute is mandatory.
    DefMatchaction interface{}
}

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption) GetEntityData() *types.CommonEntityData {
    defOption.EntityData.YFilter = defOption.YFilter
    defOption.EntityData.YangName = "def-option"
    defOption.EntityData.BundleName = "cisco_ios_xr"
    defOption.EntityData.ParentYangName = "def-options"
    defOption.EntityData.SegmentPath = "def-option" + types.AddKeyToken(defOption.DefMatchoption, "def-matchoption")
    defOption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/match/def-options/" + defOption.EntityData.SegmentPath
    defOption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defOption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defOption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defOption.EntityData.Children = types.NewOrderedMap()
    defOption.EntityData.Leafs = types.NewOrderedMap()
    defOption.EntityData.Leafs.Append("def-matchoption", types.YLeaf{"DefMatchoption", defOption.DefMatchoption})
    defOption.EntityData.Leafs.Append("def-matchaction", types.YLeaf{"DefMatchaction", defOption.DefMatchaction})

    defOption.EntityData.YListKeys = []string {"DefMatchoption"}

    return &(defOption.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters
// Table of Option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify match option. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter.
    OptionFilter []*Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter
}

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters) GetEntityData() *types.CommonEntityData {
    optionFilters.EntityData.YFilter = optionFilters.YFilter
    optionFilters.EntityData.YangName = "option-filters"
    optionFilters.EntityData.BundleName = "cisco_ios_xr"
    optionFilters.EntityData.ParentYangName = "match"
    optionFilters.EntityData.SegmentPath = "option-filters"
    optionFilters.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/match/" + optionFilters.EntityData.SegmentPath
    optionFilters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optionFilters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optionFilters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optionFilters.EntityData.Children = types.NewOrderedMap()
    optionFilters.EntityData.Children.Append("option-filter", types.YChild{"OptionFilter", nil})
    for i := range optionFilters.OptionFilter {
        optionFilters.EntityData.Children.Append(types.GetSegmentPath(optionFilters.OptionFilter[i]), types.YChild{"OptionFilter", optionFilters.OptionFilter[i]})
    }
    optionFilters.EntityData.Leafs = types.NewOrderedMap()

    optionFilters.EntityData.YListKeys = []string {}

    return &(optionFilters.EntityData)
}

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter
// Specify match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Match option 60. The type is interface{} with
    // range: 0..4294967295.
    Matchoption interface{}

    // This attribute is a key. Enter hex pattern string. The type is string with
    // length: 1..64.
    Pattern interface{}

    // This attribute is a key. Set constant integer. The type is interface{} with
    // range: 0..4294967295.
    Format interface{}

    // Vendor action. The type is ProxyAction. This attribute is mandatory.
    Matchaction interface{}
}

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter) GetEntityData() *types.CommonEntityData {
    optionFilter.EntityData.YFilter = optionFilter.YFilter
    optionFilter.EntityData.YangName = "option-filter"
    optionFilter.EntityData.BundleName = "cisco_ios_xr"
    optionFilter.EntityData.ParentYangName = "option-filters"
    optionFilter.EntityData.SegmentPath = "option-filter" + types.AddKeyToken(optionFilter.Matchoption, "matchoption") + types.AddKeyToken(optionFilter.Pattern, "pattern") + types.AddKeyToken(optionFilter.Format, "format")
    optionFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/profiles/profile/modes/mode/proxy/match/option-filters/" + optionFilter.EntityData.SegmentPath
    optionFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    optionFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    optionFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    optionFilter.EntityData.Children = types.NewOrderedMap()
    optionFilter.EntityData.Leafs = types.NewOrderedMap()
    optionFilter.EntityData.Leafs.Append("matchoption", types.YLeaf{"Matchoption", optionFilter.Matchoption})
    optionFilter.EntityData.Leafs.Append("pattern", types.YLeaf{"Pattern", optionFilter.Pattern})
    optionFilter.EntityData.Leafs.Append("format", types.YLeaf{"Format", optionFilter.Format})
    optionFilter.EntityData.Leafs.Append("matchaction", types.YLeaf{"Matchaction", optionFilter.Matchaction})

    optionFilter.EntityData.YListKeys = []string {"Matchoption", "Pattern", "Format"}

    return &(optionFilter.EntityData)
}

// Ipv4Dhcpd_Database
// Enable DHCP binding database storage to file
// system
type Ipv4Dhcpd_Database struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable DHCP proxy binding database storage to file system. The type is
    // interface{}.
    Proxy interface{}

    // Enable DHCP server binding database storage to file system. The type is
    // interface{}.
    Server interface{}

    // Enable DHCP snoop binding database storage to file system. The type is
    // interface{}.
    Snoop interface{}

    // Full file write interval (default 10 minutes). The type is interface{} with
    // range: 1..1440. The default value is 10.
    FullWriteInterval interface{}

    // Incremental file write interval (default 1 minutes). The type is
    // interface{} with range: 1..1440. The default value is 1.
    IncrementalWriteInterval interface{}
}

func (database *Ipv4Dhcpd_Database) GetEntityData() *types.CommonEntityData {
    database.EntityData.YFilter = database.YFilter
    database.EntityData.YangName = "database"
    database.EntityData.BundleName = "cisco_ios_xr"
    database.EntityData.ParentYangName = "ipv4-dhcpd"
    database.EntityData.SegmentPath = "database"
    database.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/" + database.EntityData.SegmentPath
    database.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    database.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    database.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    database.EntityData.Children = types.NewOrderedMap()
    database.EntityData.Leafs = types.NewOrderedMap()
    database.EntityData.Leafs.Append("proxy", types.YLeaf{"Proxy", database.Proxy})
    database.EntityData.Leafs.Append("server", types.YLeaf{"Server", database.Server})
    database.EntityData.Leafs.Append("snoop", types.YLeaf{"Snoop", database.Snoop})
    database.EntityData.Leafs.Append("full-write-interval", types.YLeaf{"FullWriteInterval", database.FullWriteInterval})
    database.EntityData.Leafs.Append("incremental-write-interval", types.YLeaf{"IncrementalWriteInterval", database.IncrementalWriteInterval})

    database.EntityData.YListKeys = []string {}

    return &(database.EntityData)
}

// Ipv4Dhcpd_Interfaces
// DHCP IPV4 Interface Table
type Ipv4Dhcpd_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP IPV4 Interface. The type is slice of Ipv4Dhcpd_Interfaces_Interface.
    Interface []*Ipv4Dhcpd_Interfaces_Interface
}

func (interfaces *Ipv4Dhcpd_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "ipv4-dhcpd"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/" + interfaces.EntityData.SegmentPath
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

// Ipv4Dhcpd_Interfaces_Interface
// DHCP IPV4 Interface
type Ipv4Dhcpd_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // DHCP IPv4 proxy information.
    ProxyInterface Ipv4Dhcpd_Interfaces_Interface_ProxyInterface

    // DHCP IPv4 Base profile information.
    BaseInterface Ipv4Dhcpd_Interfaces_Interface_BaseInterface

    // DHCP IPv4 relay information.
    RelayInterface Ipv4Dhcpd_Interfaces_Interface_RelayInterface

    // Static Table Entries containing MAC address to IP address bindings.
    StaticMode Ipv4Dhcpd_Interfaces_Interface_StaticMode

    // Profile name and mode.
    Profile Ipv4Dhcpd_Interfaces_Interface_Profile

    // DHCP IPv4 Server information.
    ServerInterface Ipv4Dhcpd_Interfaces_Interface_ServerInterface

    // DHCP IPv4 snoop information.
    SnoopInterface Ipv4Dhcpd_Interfaces_Interface_SnoopInterface
}

func (self *Ipv4Dhcpd_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("proxy-interface", types.YChild{"ProxyInterface", &self.ProxyInterface})
    self.EntityData.Children.Append("base-interface", types.YChild{"BaseInterface", &self.BaseInterface})
    self.EntityData.Children.Append("relay-interface", types.YChild{"RelayInterface", &self.RelayInterface})
    self.EntityData.Children.Append("static-mode", types.YChild{"StaticMode", &self.StaticMode})
    self.EntityData.Children.Append("profile", types.YChild{"Profile", &self.Profile})
    self.EntityData.Children.Append("server-interface", types.YChild{"ServerInterface", &self.ServerInterface})
    self.EntityData.Children.Append("snoop-interface", types.YChild{"SnoopInterface", &self.SnoopInterface})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Ipv4Dhcpd_Interfaces_Interface_ProxyInterface
// DHCP IPv4 proxy information
type Ipv4Dhcpd_Interfaces_Interface_ProxyInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface profile name. The type is string.
    Profile interface{}

    // Circuit ID value.
    DhcpCircuitId Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId
}

func (proxyInterface *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface) GetEntityData() *types.CommonEntityData {
    proxyInterface.EntityData.YFilter = proxyInterface.YFilter
    proxyInterface.EntityData.YangName = "proxy-interface"
    proxyInterface.EntityData.BundleName = "cisco_ios_xr"
    proxyInterface.EntityData.ParentYangName = "interface"
    proxyInterface.EntityData.SegmentPath = "proxy-interface"
    proxyInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/interfaces/interface/" + proxyInterface.EntityData.SegmentPath
    proxyInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proxyInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proxyInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proxyInterface.EntityData.Children = types.NewOrderedMap()
    proxyInterface.EntityData.Children.Append("dhcp-circuit-id", types.YChild{"DhcpCircuitId", &proxyInterface.DhcpCircuitId})
    proxyInterface.EntityData.Leafs = types.NewOrderedMap()
    proxyInterface.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", proxyInterface.Profile})

    proxyInterface.EntityData.YListKeys = []string {}

    return &(proxyInterface.EntityData)
}

// Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId
// Circuit ID value
// This type is a presence type.
type Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // DHCP IPv4 circuit ID value. The type is string. This attribute is
    // mandatory.
    CircuitId interface{}

    // Format String. The type is Ipv4dhcpdFmt. This attribute is mandatory.
    Format interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument1 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument2 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument3 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument4 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument5 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument6 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument7 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument8 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument9 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument10 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument11 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument12 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument13 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument14 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument15 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument16 interface{}
}

func (dhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId) GetEntityData() *types.CommonEntityData {
    dhcpCircuitId.EntityData.YFilter = dhcpCircuitId.YFilter
    dhcpCircuitId.EntityData.YangName = "dhcp-circuit-id"
    dhcpCircuitId.EntityData.BundleName = "cisco_ios_xr"
    dhcpCircuitId.EntityData.ParentYangName = "proxy-interface"
    dhcpCircuitId.EntityData.SegmentPath = "dhcp-circuit-id"
    dhcpCircuitId.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/interfaces/interface/proxy-interface/" + dhcpCircuitId.EntityData.SegmentPath
    dhcpCircuitId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dhcpCircuitId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dhcpCircuitId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dhcpCircuitId.EntityData.Children = types.NewOrderedMap()
    dhcpCircuitId.EntityData.Leafs = types.NewOrderedMap()
    dhcpCircuitId.EntityData.Leafs.Append("circuit-id", types.YLeaf{"CircuitId", dhcpCircuitId.CircuitId})
    dhcpCircuitId.EntityData.Leafs.Append("format", types.YLeaf{"Format", dhcpCircuitId.Format})
    dhcpCircuitId.EntityData.Leafs.Append("argument1", types.YLeaf{"Argument1", dhcpCircuitId.Argument1})
    dhcpCircuitId.EntityData.Leafs.Append("argument2", types.YLeaf{"Argument2", dhcpCircuitId.Argument2})
    dhcpCircuitId.EntityData.Leafs.Append("argument3", types.YLeaf{"Argument3", dhcpCircuitId.Argument3})
    dhcpCircuitId.EntityData.Leafs.Append("argument4", types.YLeaf{"Argument4", dhcpCircuitId.Argument4})
    dhcpCircuitId.EntityData.Leafs.Append("argument5", types.YLeaf{"Argument5", dhcpCircuitId.Argument5})
    dhcpCircuitId.EntityData.Leafs.Append("argument6", types.YLeaf{"Argument6", dhcpCircuitId.Argument6})
    dhcpCircuitId.EntityData.Leafs.Append("argument7", types.YLeaf{"Argument7", dhcpCircuitId.Argument7})
    dhcpCircuitId.EntityData.Leafs.Append("argument8", types.YLeaf{"Argument8", dhcpCircuitId.Argument8})
    dhcpCircuitId.EntityData.Leafs.Append("argument9", types.YLeaf{"Argument9", dhcpCircuitId.Argument9})
    dhcpCircuitId.EntityData.Leafs.Append("argument10", types.YLeaf{"Argument10", dhcpCircuitId.Argument10})
    dhcpCircuitId.EntityData.Leafs.Append("argument11", types.YLeaf{"Argument11", dhcpCircuitId.Argument11})
    dhcpCircuitId.EntityData.Leafs.Append("argument12", types.YLeaf{"Argument12", dhcpCircuitId.Argument12})
    dhcpCircuitId.EntityData.Leafs.Append("argument13", types.YLeaf{"Argument13", dhcpCircuitId.Argument13})
    dhcpCircuitId.EntityData.Leafs.Append("argument14", types.YLeaf{"Argument14", dhcpCircuitId.Argument14})
    dhcpCircuitId.EntityData.Leafs.Append("argument15", types.YLeaf{"Argument15", dhcpCircuitId.Argument15})
    dhcpCircuitId.EntityData.Leafs.Append("argument16", types.YLeaf{"Argument16", dhcpCircuitId.Argument16})

    dhcpCircuitId.EntityData.YListKeys = []string {}

    return &(dhcpCircuitId.EntityData)
}

// Ipv4Dhcpd_Interfaces_Interface_BaseInterface
// DHCP IPv4 Base profile information
type Ipv4Dhcpd_Interfaces_Interface_BaseInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface profile name. The type is string.
    Profile interface{}

    // Circuit ID value.
    BaseDhcpCircuitId Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId
}

func (baseInterface *Ipv4Dhcpd_Interfaces_Interface_BaseInterface) GetEntityData() *types.CommonEntityData {
    baseInterface.EntityData.YFilter = baseInterface.YFilter
    baseInterface.EntityData.YangName = "base-interface"
    baseInterface.EntityData.BundleName = "cisco_ios_xr"
    baseInterface.EntityData.ParentYangName = "interface"
    baseInterface.EntityData.SegmentPath = "base-interface"
    baseInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/interfaces/interface/" + baseInterface.EntityData.SegmentPath
    baseInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    baseInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    baseInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    baseInterface.EntityData.Children = types.NewOrderedMap()
    baseInterface.EntityData.Children.Append("base-dhcp-circuit-id", types.YChild{"BaseDhcpCircuitId", &baseInterface.BaseDhcpCircuitId})
    baseInterface.EntityData.Leafs = types.NewOrderedMap()
    baseInterface.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", baseInterface.Profile})

    baseInterface.EntityData.YListKeys = []string {}

    return &(baseInterface.EntityData)
}

// Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId
// Circuit ID value
// This type is a presence type.
type Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // DHCP IPv4 circuit ID value. The type is string. This attribute is
    // mandatory.
    CircuitId interface{}

    // Format String. The type is Ipv4dhcpdFmt. This attribute is mandatory.
    Format interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument1 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument2 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument3 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument4 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument5 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument6 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument7 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument8 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument9 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument10 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument11 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument12 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument13 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument14 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument15 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument16 interface{}
}

func (baseDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId) GetEntityData() *types.CommonEntityData {
    baseDhcpCircuitId.EntityData.YFilter = baseDhcpCircuitId.YFilter
    baseDhcpCircuitId.EntityData.YangName = "base-dhcp-circuit-id"
    baseDhcpCircuitId.EntityData.BundleName = "cisco_ios_xr"
    baseDhcpCircuitId.EntityData.ParentYangName = "base-interface"
    baseDhcpCircuitId.EntityData.SegmentPath = "base-dhcp-circuit-id"
    baseDhcpCircuitId.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/interfaces/interface/base-interface/" + baseDhcpCircuitId.EntityData.SegmentPath
    baseDhcpCircuitId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    baseDhcpCircuitId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    baseDhcpCircuitId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    baseDhcpCircuitId.EntityData.Children = types.NewOrderedMap()
    baseDhcpCircuitId.EntityData.Leafs = types.NewOrderedMap()
    baseDhcpCircuitId.EntityData.Leafs.Append("circuit-id", types.YLeaf{"CircuitId", baseDhcpCircuitId.CircuitId})
    baseDhcpCircuitId.EntityData.Leafs.Append("format", types.YLeaf{"Format", baseDhcpCircuitId.Format})
    baseDhcpCircuitId.EntityData.Leafs.Append("argument1", types.YLeaf{"Argument1", baseDhcpCircuitId.Argument1})
    baseDhcpCircuitId.EntityData.Leafs.Append("argument2", types.YLeaf{"Argument2", baseDhcpCircuitId.Argument2})
    baseDhcpCircuitId.EntityData.Leafs.Append("argument3", types.YLeaf{"Argument3", baseDhcpCircuitId.Argument3})
    baseDhcpCircuitId.EntityData.Leafs.Append("argument4", types.YLeaf{"Argument4", baseDhcpCircuitId.Argument4})
    baseDhcpCircuitId.EntityData.Leafs.Append("argument5", types.YLeaf{"Argument5", baseDhcpCircuitId.Argument5})
    baseDhcpCircuitId.EntityData.Leafs.Append("argument6", types.YLeaf{"Argument6", baseDhcpCircuitId.Argument6})
    baseDhcpCircuitId.EntityData.Leafs.Append("argument7", types.YLeaf{"Argument7", baseDhcpCircuitId.Argument7})
    baseDhcpCircuitId.EntityData.Leafs.Append("argument8", types.YLeaf{"Argument8", baseDhcpCircuitId.Argument8})
    baseDhcpCircuitId.EntityData.Leafs.Append("argument9", types.YLeaf{"Argument9", baseDhcpCircuitId.Argument9})
    baseDhcpCircuitId.EntityData.Leafs.Append("argument10", types.YLeaf{"Argument10", baseDhcpCircuitId.Argument10})
    baseDhcpCircuitId.EntityData.Leafs.Append("argument11", types.YLeaf{"Argument11", baseDhcpCircuitId.Argument11})
    baseDhcpCircuitId.EntityData.Leafs.Append("argument12", types.YLeaf{"Argument12", baseDhcpCircuitId.Argument12})
    baseDhcpCircuitId.EntityData.Leafs.Append("argument13", types.YLeaf{"Argument13", baseDhcpCircuitId.Argument13})
    baseDhcpCircuitId.EntityData.Leafs.Append("argument14", types.YLeaf{"Argument14", baseDhcpCircuitId.Argument14})
    baseDhcpCircuitId.EntityData.Leafs.Append("argument15", types.YLeaf{"Argument15", baseDhcpCircuitId.Argument15})
    baseDhcpCircuitId.EntityData.Leafs.Append("argument16", types.YLeaf{"Argument16", baseDhcpCircuitId.Argument16})

    baseDhcpCircuitId.EntityData.YListKeys = []string {}

    return &(baseDhcpCircuitId.EntityData)
}

// Ipv4Dhcpd_Interfaces_Interface_RelayInterface
// DHCP IPv4 relay information
type Ipv4Dhcpd_Interfaces_Interface_RelayInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Circuit ID value.
    RelayDhcpCircuitId Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId
}

func (relayInterface *Ipv4Dhcpd_Interfaces_Interface_RelayInterface) GetEntityData() *types.CommonEntityData {
    relayInterface.EntityData.YFilter = relayInterface.YFilter
    relayInterface.EntityData.YangName = "relay-interface"
    relayInterface.EntityData.BundleName = "cisco_ios_xr"
    relayInterface.EntityData.ParentYangName = "interface"
    relayInterface.EntityData.SegmentPath = "relay-interface"
    relayInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/interfaces/interface/" + relayInterface.EntityData.SegmentPath
    relayInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayInterface.EntityData.Children = types.NewOrderedMap()
    relayInterface.EntityData.Children.Append("relay-dhcp-circuit-id", types.YChild{"RelayDhcpCircuitId", &relayInterface.RelayDhcpCircuitId})
    relayInterface.EntityData.Leafs = types.NewOrderedMap()

    relayInterface.EntityData.YListKeys = []string {}

    return &(relayInterface.EntityData)
}

// Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId
// Circuit ID value
// This type is a presence type.
type Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // DHCP IPv4 circuit ID value. The type is string. This attribute is
    // mandatory.
    CircuitId interface{}

    // Format String. The type is Ipv4dhcpdFmt. This attribute is mandatory.
    Format interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument1 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument2 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument3 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument4 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument5 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument6 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument7 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument8 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument9 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument10 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument11 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument12 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument13 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument14 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument15 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument16 interface{}
}

func (relayDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId) GetEntityData() *types.CommonEntityData {
    relayDhcpCircuitId.EntityData.YFilter = relayDhcpCircuitId.YFilter
    relayDhcpCircuitId.EntityData.YangName = "relay-dhcp-circuit-id"
    relayDhcpCircuitId.EntityData.BundleName = "cisco_ios_xr"
    relayDhcpCircuitId.EntityData.ParentYangName = "relay-interface"
    relayDhcpCircuitId.EntityData.SegmentPath = "relay-dhcp-circuit-id"
    relayDhcpCircuitId.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/interfaces/interface/relay-interface/" + relayDhcpCircuitId.EntityData.SegmentPath
    relayDhcpCircuitId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    relayDhcpCircuitId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    relayDhcpCircuitId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    relayDhcpCircuitId.EntityData.Children = types.NewOrderedMap()
    relayDhcpCircuitId.EntityData.Leafs = types.NewOrderedMap()
    relayDhcpCircuitId.EntityData.Leafs.Append("circuit-id", types.YLeaf{"CircuitId", relayDhcpCircuitId.CircuitId})
    relayDhcpCircuitId.EntityData.Leafs.Append("format", types.YLeaf{"Format", relayDhcpCircuitId.Format})
    relayDhcpCircuitId.EntityData.Leafs.Append("argument1", types.YLeaf{"Argument1", relayDhcpCircuitId.Argument1})
    relayDhcpCircuitId.EntityData.Leafs.Append("argument2", types.YLeaf{"Argument2", relayDhcpCircuitId.Argument2})
    relayDhcpCircuitId.EntityData.Leafs.Append("argument3", types.YLeaf{"Argument3", relayDhcpCircuitId.Argument3})
    relayDhcpCircuitId.EntityData.Leafs.Append("argument4", types.YLeaf{"Argument4", relayDhcpCircuitId.Argument4})
    relayDhcpCircuitId.EntityData.Leafs.Append("argument5", types.YLeaf{"Argument5", relayDhcpCircuitId.Argument5})
    relayDhcpCircuitId.EntityData.Leafs.Append("argument6", types.YLeaf{"Argument6", relayDhcpCircuitId.Argument6})
    relayDhcpCircuitId.EntityData.Leafs.Append("argument7", types.YLeaf{"Argument7", relayDhcpCircuitId.Argument7})
    relayDhcpCircuitId.EntityData.Leafs.Append("argument8", types.YLeaf{"Argument8", relayDhcpCircuitId.Argument8})
    relayDhcpCircuitId.EntityData.Leafs.Append("argument9", types.YLeaf{"Argument9", relayDhcpCircuitId.Argument9})
    relayDhcpCircuitId.EntityData.Leafs.Append("argument10", types.YLeaf{"Argument10", relayDhcpCircuitId.Argument10})
    relayDhcpCircuitId.EntityData.Leafs.Append("argument11", types.YLeaf{"Argument11", relayDhcpCircuitId.Argument11})
    relayDhcpCircuitId.EntityData.Leafs.Append("argument12", types.YLeaf{"Argument12", relayDhcpCircuitId.Argument12})
    relayDhcpCircuitId.EntityData.Leafs.Append("argument13", types.YLeaf{"Argument13", relayDhcpCircuitId.Argument13})
    relayDhcpCircuitId.EntityData.Leafs.Append("argument14", types.YLeaf{"Argument14", relayDhcpCircuitId.Argument14})
    relayDhcpCircuitId.EntityData.Leafs.Append("argument15", types.YLeaf{"Argument15", relayDhcpCircuitId.Argument15})
    relayDhcpCircuitId.EntityData.Leafs.Append("argument16", types.YLeaf{"Argument16", relayDhcpCircuitId.Argument16})

    relayDhcpCircuitId.EntityData.YListKeys = []string {}

    return &(relayDhcpCircuitId.EntityData)
}

// Ipv4Dhcpd_Interfaces_Interface_StaticMode
// Static Table Entries containing MAC address to
// IP address bindings
type Ipv4Dhcpd_Interfaces_Interface_StaticMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Static Table Entries containing MAC address to IP address bindings.
    Statics Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics
}

func (staticMode *Ipv4Dhcpd_Interfaces_Interface_StaticMode) GetEntityData() *types.CommonEntityData {
    staticMode.EntityData.YFilter = staticMode.YFilter
    staticMode.EntityData.YangName = "static-mode"
    staticMode.EntityData.BundleName = "cisco_ios_xr"
    staticMode.EntityData.ParentYangName = "interface"
    staticMode.EntityData.SegmentPath = "static-mode"
    staticMode.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/interfaces/interface/" + staticMode.EntityData.SegmentPath
    staticMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticMode.EntityData.Children = types.NewOrderedMap()
    staticMode.EntityData.Children.Append("statics", types.YChild{"Statics", &staticMode.Statics})
    staticMode.EntityData.Leafs = types.NewOrderedMap()

    staticMode.EntityData.YListKeys = []string {}

    return &(staticMode.EntityData)
}

// Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics
// Static Table Entries containing MAC address
// to IP address bindings
type Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // DHCP static binding of Mac address to IP address. The type is slice of
    // Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static.
    Static []*Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static
}

func (statics *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics) GetEntityData() *types.CommonEntityData {
    statics.EntityData.YFilter = statics.YFilter
    statics.EntityData.YangName = "statics"
    statics.EntityData.BundleName = "cisco_ios_xr"
    statics.EntityData.ParentYangName = "static-mode"
    statics.EntityData.SegmentPath = "statics"
    statics.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/interfaces/interface/static-mode/" + statics.EntityData.SegmentPath
    statics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statics.EntityData.Children = types.NewOrderedMap()
    statics.EntityData.Children.Append("static", types.YChild{"Static", nil})
    for i := range statics.Static {
        statics.EntityData.Children.Append(types.GetSegmentPath(statics.Static[i]), types.YChild{"Static", statics.Static[i]})
    }
    statics.EntityData.Leafs = types.NewOrderedMap()

    statics.EntityData.YListKeys = []string {}

    return &(statics.EntityData)
}

// Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static
// DHCP static binding of Mac address to IP
// address
type Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. MACAddress. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // This attribute is a key. Client Id. The type is interface{} with range:
    // 0..4294967295.
    ClientId interface{}

    // This attribute is a key. DHCP IPV4 Static layer. The type is
    // Ipv4dhcpdLayer.
    Layer interface{}

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    StaticAddress interface{}
}

func (static *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static) GetEntityData() *types.CommonEntityData {
    static.EntityData.YFilter = static.YFilter
    static.EntityData.YangName = "static"
    static.EntityData.BundleName = "cisco_ios_xr"
    static.EntityData.ParentYangName = "statics"
    static.EntityData.SegmentPath = "static" + types.AddKeyToken(static.MacAddress, "mac-address") + types.AddKeyToken(static.ClientId, "client-id") + types.AddKeyToken(static.Layer, "layer")
    static.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/interfaces/interface/static-mode/statics/" + static.EntityData.SegmentPath
    static.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    static.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    static.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    static.EntityData.Children = types.NewOrderedMap()
    static.EntityData.Leafs = types.NewOrderedMap()
    static.EntityData.Leafs.Append("mac-address", types.YLeaf{"MacAddress", static.MacAddress})
    static.EntityData.Leafs.Append("client-id", types.YLeaf{"ClientId", static.ClientId})
    static.EntityData.Leafs.Append("layer", types.YLeaf{"Layer", static.Layer})
    static.EntityData.Leafs.Append("static-address", types.YLeaf{"StaticAddress", static.StaticAddress})

    static.EntityData.YListKeys = []string {"MacAddress", "ClientId", "Layer"}

    return &(static.EntityData)
}

// Ipv4Dhcpd_Interfaces_Interface_Profile
// Profile name and mode
// This type is a presence type.
type Ipv4Dhcpd_Interfaces_Interface_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Profile name. The type is string. This attribute is mandatory.
    ProfileName interface{}

    // DHCP mode. The type is Ipv4dhcpdMode. This attribute is mandatory.
    Mode interface{}
}

func (profile *Ipv4Dhcpd_Interfaces_Interface_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "interface"
    profile.EntityData.SegmentPath = "profile"
    profile.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/interfaces/interface/" + profile.EntityData.SegmentPath
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = types.NewOrderedMap()
    profile.EntityData.Leafs = types.NewOrderedMap()
    profile.EntityData.Leafs.Append("profile-name", types.YLeaf{"ProfileName", profile.ProfileName})
    profile.EntityData.Leafs.Append("mode", types.YLeaf{"Mode", profile.Mode})

    profile.EntityData.YListKeys = []string {}

    return &(profile.EntityData)
}

// Ipv4Dhcpd_Interfaces_Interface_ServerInterface
// DHCP IPv4 Server information
type Ipv4Dhcpd_Interfaces_Interface_ServerInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface profile name. The type is string.
    Profile interface{}

    // Circuit ID value.
    ServerDhcpCircuitId Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId
}

func (serverInterface *Ipv4Dhcpd_Interfaces_Interface_ServerInterface) GetEntityData() *types.CommonEntityData {
    serverInterface.EntityData.YFilter = serverInterface.YFilter
    serverInterface.EntityData.YangName = "server-interface"
    serverInterface.EntityData.BundleName = "cisco_ios_xr"
    serverInterface.EntityData.ParentYangName = "interface"
    serverInterface.EntityData.SegmentPath = "server-interface"
    serverInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/interfaces/interface/" + serverInterface.EntityData.SegmentPath
    serverInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverInterface.EntityData.Children = types.NewOrderedMap()
    serverInterface.EntityData.Children.Append("server-dhcp-circuit-id", types.YChild{"ServerDhcpCircuitId", &serverInterface.ServerDhcpCircuitId})
    serverInterface.EntityData.Leafs = types.NewOrderedMap()
    serverInterface.EntityData.Leafs.Append("profile", types.YLeaf{"Profile", serverInterface.Profile})

    serverInterface.EntityData.YListKeys = []string {}

    return &(serverInterface.EntityData)
}

// Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId
// Circuit ID value
// This type is a presence type.
type Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // DHCP IPv4 circuit ID value. The type is string. This attribute is
    // mandatory.
    CircuitId interface{}

    // Format String. The type is Ipv4dhcpdFmt. This attribute is mandatory.
    Format interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument1 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument2 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument3 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument4 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument5 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument6 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument7 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument8 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument9 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument10 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument11 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument12 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument13 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument14 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument15 interface{}

    // Argument. The type is Ipv4dhcpdFmtSpecifier.
    Argument16 interface{}
}

func (serverDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId) GetEntityData() *types.CommonEntityData {
    serverDhcpCircuitId.EntityData.YFilter = serverDhcpCircuitId.YFilter
    serverDhcpCircuitId.EntityData.YangName = "server-dhcp-circuit-id"
    serverDhcpCircuitId.EntityData.BundleName = "cisco_ios_xr"
    serverDhcpCircuitId.EntityData.ParentYangName = "server-interface"
    serverDhcpCircuitId.EntityData.SegmentPath = "server-dhcp-circuit-id"
    serverDhcpCircuitId.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/interfaces/interface/server-interface/" + serverDhcpCircuitId.EntityData.SegmentPath
    serverDhcpCircuitId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    serverDhcpCircuitId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    serverDhcpCircuitId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    serverDhcpCircuitId.EntityData.Children = types.NewOrderedMap()
    serverDhcpCircuitId.EntityData.Leafs = types.NewOrderedMap()
    serverDhcpCircuitId.EntityData.Leafs.Append("circuit-id", types.YLeaf{"CircuitId", serverDhcpCircuitId.CircuitId})
    serverDhcpCircuitId.EntityData.Leafs.Append("format", types.YLeaf{"Format", serverDhcpCircuitId.Format})
    serverDhcpCircuitId.EntityData.Leafs.Append("argument1", types.YLeaf{"Argument1", serverDhcpCircuitId.Argument1})
    serverDhcpCircuitId.EntityData.Leafs.Append("argument2", types.YLeaf{"Argument2", serverDhcpCircuitId.Argument2})
    serverDhcpCircuitId.EntityData.Leafs.Append("argument3", types.YLeaf{"Argument3", serverDhcpCircuitId.Argument3})
    serverDhcpCircuitId.EntityData.Leafs.Append("argument4", types.YLeaf{"Argument4", serverDhcpCircuitId.Argument4})
    serverDhcpCircuitId.EntityData.Leafs.Append("argument5", types.YLeaf{"Argument5", serverDhcpCircuitId.Argument5})
    serverDhcpCircuitId.EntityData.Leafs.Append("argument6", types.YLeaf{"Argument6", serverDhcpCircuitId.Argument6})
    serverDhcpCircuitId.EntityData.Leafs.Append("argument7", types.YLeaf{"Argument7", serverDhcpCircuitId.Argument7})
    serverDhcpCircuitId.EntityData.Leafs.Append("argument8", types.YLeaf{"Argument8", serverDhcpCircuitId.Argument8})
    serverDhcpCircuitId.EntityData.Leafs.Append("argument9", types.YLeaf{"Argument9", serverDhcpCircuitId.Argument9})
    serverDhcpCircuitId.EntityData.Leafs.Append("argument10", types.YLeaf{"Argument10", serverDhcpCircuitId.Argument10})
    serverDhcpCircuitId.EntityData.Leafs.Append("argument11", types.YLeaf{"Argument11", serverDhcpCircuitId.Argument11})
    serverDhcpCircuitId.EntityData.Leafs.Append("argument12", types.YLeaf{"Argument12", serverDhcpCircuitId.Argument12})
    serverDhcpCircuitId.EntityData.Leafs.Append("argument13", types.YLeaf{"Argument13", serverDhcpCircuitId.Argument13})
    serverDhcpCircuitId.EntityData.Leafs.Append("argument14", types.YLeaf{"Argument14", serverDhcpCircuitId.Argument14})
    serverDhcpCircuitId.EntityData.Leafs.Append("argument15", types.YLeaf{"Argument15", serverDhcpCircuitId.Argument15})
    serverDhcpCircuitId.EntityData.Leafs.Append("argument16", types.YLeaf{"Argument16", serverDhcpCircuitId.Argument16})

    serverDhcpCircuitId.EntityData.YListKeys = []string {}

    return &(serverDhcpCircuitId.EntityData)
}

// Ipv4Dhcpd_Interfaces_Interface_SnoopInterface
// DHCP IPv4 snoop information
type Ipv4Dhcpd_Interfaces_Interface_SnoopInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure circuit ID for snoop 1. Hex 2. ASCII.
    SnoopCircuitId Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId
}

func (snoopInterface *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface) GetEntityData() *types.CommonEntityData {
    snoopInterface.EntityData.YFilter = snoopInterface.YFilter
    snoopInterface.EntityData.YangName = "snoop-interface"
    snoopInterface.EntityData.BundleName = "cisco_ios_xr"
    snoopInterface.EntityData.ParentYangName = "interface"
    snoopInterface.EntityData.SegmentPath = "snoop-interface"
    snoopInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/interfaces/interface/" + snoopInterface.EntityData.SegmentPath
    snoopInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snoopInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snoopInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snoopInterface.EntityData.Children = types.NewOrderedMap()
    snoopInterface.EntityData.Children.Append("snoop-circuit-id", types.YChild{"SnoopCircuitId", &snoopInterface.SnoopCircuitId})
    snoopInterface.EntityData.Leafs = types.NewOrderedMap()

    snoopInterface.EntityData.YListKeys = []string {}

    return &(snoopInterface.EntityData)
}

// Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId
// Configure circuit ID for snoop 1. Hex 2.
// ASCII
type Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Format type, 1. Hex 2. ASCII. The type is interface{} with range: 1..2.
    FormatType interface{}

    // Enter circuit-id value. The type is string.
    CircuitIdValue interface{}
}

func (snoopCircuitId *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId) GetEntityData() *types.CommonEntityData {
    snoopCircuitId.EntityData.YFilter = snoopCircuitId.YFilter
    snoopCircuitId.EntityData.YangName = "snoop-circuit-id"
    snoopCircuitId.EntityData.BundleName = "cisco_ios_xr"
    snoopCircuitId.EntityData.ParentYangName = "snoop-interface"
    snoopCircuitId.EntityData.SegmentPath = "snoop-circuit-id"
    snoopCircuitId.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/interfaces/interface/snoop-interface/" + snoopCircuitId.EntityData.SegmentPath
    snoopCircuitId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snoopCircuitId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snoopCircuitId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snoopCircuitId.EntityData.Children = types.NewOrderedMap()
    snoopCircuitId.EntityData.Leafs = types.NewOrderedMap()
    snoopCircuitId.EntityData.Leafs.Append("format-type", types.YLeaf{"FormatType", snoopCircuitId.FormatType})
    snoopCircuitId.EntityData.Leafs.Append("circuit-id-value", types.YLeaf{"CircuitIdValue", snoopCircuitId.CircuitIdValue})

    snoopCircuitId.EntityData.YListKeys = []string {}

    return &(snoopCircuitId.EntityData)
}

// Ipv4Dhcpd_DuplicateMacAllowed
// Allow Duplicate MAC Address
// This type is a presence type.
type Ipv4Dhcpd_DuplicateMacAllowed struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Duplicate mac is allowed. The type is interface{}. This attribute is
    // mandatory.
    DuplicateMac interface{}

    // Exclude vlan. The type is interface{}.
    ExcludeVlan interface{}

    // Include giaddr. The type is interface{}.
    IncludeGiaddr interface{}
}

func (duplicateMacAllowed *Ipv4Dhcpd_DuplicateMacAllowed) GetEntityData() *types.CommonEntityData {
    duplicateMacAllowed.EntityData.YFilter = duplicateMacAllowed.YFilter
    duplicateMacAllowed.EntityData.YangName = "duplicate-mac-allowed"
    duplicateMacAllowed.EntityData.BundleName = "cisco_ios_xr"
    duplicateMacAllowed.EntityData.ParentYangName = "ipv4-dhcpd"
    duplicateMacAllowed.EntityData.SegmentPath = "duplicate-mac-allowed"
    duplicateMacAllowed.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/" + duplicateMacAllowed.EntityData.SegmentPath
    duplicateMacAllowed.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    duplicateMacAllowed.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    duplicateMacAllowed.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    duplicateMacAllowed.EntityData.Children = types.NewOrderedMap()
    duplicateMacAllowed.EntityData.Leafs = types.NewOrderedMap()
    duplicateMacAllowed.EntityData.Leafs.Append("duplicate-mac", types.YLeaf{"DuplicateMac", duplicateMacAllowed.DuplicateMac})
    duplicateMacAllowed.EntityData.Leafs.Append("exclude-vlan", types.YLeaf{"ExcludeVlan", duplicateMacAllowed.ExcludeVlan})
    duplicateMacAllowed.EntityData.Leafs.Append("include-giaddr", types.YLeaf{"IncludeGiaddr", duplicateMacAllowed.IncludeGiaddr})

    duplicateMacAllowed.EntityData.YListKeys = []string {}

    return &(duplicateMacAllowed.EntityData)
}

// Ipv4Dhcpd_RateLimit
// Rate limit ingress packets
type Ipv4Dhcpd_RateLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Rate limiter period in msec (default: 200 msec). The type is interface{}
    // with range: 1..1000. The default value is 200.
    NumPeriod interface{}

    // Max DISCOVER packets per rate-limiter period (default 100). The type is
    // interface{} with range: 0..1000. The default value is 100.
    NumDiscover interface{}
}

func (rateLimit *Ipv4Dhcpd_RateLimit) GetEntityData() *types.CommonEntityData {
    rateLimit.EntityData.YFilter = rateLimit.YFilter
    rateLimit.EntityData.YangName = "rate-limit"
    rateLimit.EntityData.BundleName = "cisco_ios_xr"
    rateLimit.EntityData.ParentYangName = "ipv4-dhcpd"
    rateLimit.EntityData.SegmentPath = "rate-limit"
    rateLimit.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd/" + rateLimit.EntityData.SegmentPath
    rateLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rateLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rateLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rateLimit.EntityData.Children = types.NewOrderedMap()
    rateLimit.EntityData.Leafs = types.NewOrderedMap()
    rateLimit.EntityData.Leafs.Append("num-period", types.YLeaf{"NumPeriod", rateLimit.NumPeriod})
    rateLimit.EntityData.Leafs.Append("num-discover", types.YLeaf{"NumDiscover", rateLimit.NumDiscover})

    rateLimit.EntityData.YListKeys = []string {}

    return &(rateLimit.EntityData)
}

