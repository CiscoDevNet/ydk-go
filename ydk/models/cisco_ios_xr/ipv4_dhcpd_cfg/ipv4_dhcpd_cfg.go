// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-dhcpd package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ipv4-dhcpd: DHCP IPV4 configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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

// ProxyAction represents Proxy action
type ProxyAction string

const (
    // Allow vendor specific DHCP Discover
    ProxyAction_allow ProxyAction = "allow"

    // Drop vendor specific DHCP Discover
    ProxyAction_drop ProxyAction = "drop"
)

// Ipv4dhcpdLayer represents Ipv4dhcpd layer
type Ipv4dhcpdLayer string

const (
    // Layer2
    Ipv4dhcpdLayer_layer2 Ipv4dhcpdLayer = "layer2"

    // Layer3
    Ipv4dhcpdLayer_layer3 Ipv4dhcpdLayer = "layer3"
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
)

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

// Ipv4dhcpdRelayInfoOptionAuthenticate represents Ipv4dhcpd relay info option authenticate
type Ipv4dhcpdRelayInfoOptionAuthenticate string

const (
    // Received
    Ipv4dhcpdRelayInfoOptionAuthenticate_received Ipv4dhcpdRelayInfoOptionAuthenticate = "received"

    // Inserted
    Ipv4dhcpdRelayInfoOptionAuthenticate_inserted Ipv4dhcpdRelayInfoOptionAuthenticate = "inserted"
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

// Ipv4dhcpdRelayInfoOptionvpnMode represents Ipv4dhcpd relay info optionvpn mode
type Ipv4dhcpdRelayInfoOptionvpnMode string

const (
    // RFC
    Ipv4dhcpdRelayInfoOptionvpnMode_rfc Ipv4dhcpdRelayInfoOptionvpnMode = "rfc"

    // Cisco
    Ipv4dhcpdRelayInfoOptionvpnMode_cisco Ipv4dhcpdRelayInfoOptionvpnMode = "cisco"
)

// Ipv4dhcpdGiaddrPolicy represents Ipv4dhcpd giaddr policy
type Ipv4dhcpdGiaddrPolicy string

const (
    // Giaddr Policy Keep
    Ipv4dhcpdGiaddrPolicy_giaddr_policy_keep Ipv4dhcpdGiaddrPolicy = "giaddr-policy-keep"
)

// Ipv4dhcpdFmt represents Ipv4dhcpd fmt
type Ipv4dhcpdFmt string

const (
    // Not a Format String
    Ipv4dhcpdFmt_no_format Ipv4dhcpdFmt = "no-format"

    // Format String
    Ipv4dhcpdFmt_format Ipv4dhcpdFmt = "format"
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

// MacMismatchAction represents Mac mismatch action
type MacMismatchAction string

const (
    // Forward
    MacMismatchAction_forward MacMismatchAction = "forward"

    // Drop
    MacMismatchAction_drop MacMismatchAction = "drop"
)

// BaseAction represents Base action
type BaseAction string

const (
    // Allow vendor specific DHCP Discover
    BaseAction_allow BaseAction = "allow"

    // Drop vendor specific DHCP Discover
    BaseAction_drop BaseAction = "drop"
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

// Ipv4Dhcpd
// DHCP IPV4 configuration
type Ipv4Dhcpd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP IPV4 configuration. The type is interface{}.
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

func (ipv4Dhcpd *Ipv4Dhcpd) GetFilter() yfilter.YFilter { return ipv4Dhcpd.YFilter }

func (ipv4Dhcpd *Ipv4Dhcpd) SetFilter(yf yfilter.YFilter) { ipv4Dhcpd.YFilter = yf }

func (ipv4Dhcpd *Ipv4Dhcpd) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "outer-cos" { return "OuterCos" }
    if yname == "allow-client-id-change" { return "AllowClientIdChange" }
    if yname == "inner-cos" { return "InnerCos" }
    if yname == "vrfs" { return "Vrfs" }
    if yname == "profiles" { return "Profiles" }
    if yname == "database" { return "Database" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "duplicate-mac-allowed" { return "DuplicateMacAllowed" }
    if yname == "rate-limit" { return "RateLimit" }
    return ""
}

func (ipv4Dhcpd *Ipv4Dhcpd) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-dhcpd-cfg:ipv4-dhcpd"
}

func (ipv4Dhcpd *Ipv4Dhcpd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &ipv4Dhcpd.Vrfs
    }
    if childYangName == "profiles" {
        return &ipv4Dhcpd.Profiles
    }
    if childYangName == "database" {
        return &ipv4Dhcpd.Database
    }
    if childYangName == "interfaces" {
        return &ipv4Dhcpd.Interfaces
    }
    if childYangName == "duplicate-mac-allowed" {
        return &ipv4Dhcpd.DuplicateMacAllowed
    }
    if childYangName == "rate-limit" {
        return &ipv4Dhcpd.RateLimit
    }
    return nil
}

func (ipv4Dhcpd *Ipv4Dhcpd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &ipv4Dhcpd.Vrfs
    children["profiles"] = &ipv4Dhcpd.Profiles
    children["database"] = &ipv4Dhcpd.Database
    children["interfaces"] = &ipv4Dhcpd.Interfaces
    children["duplicate-mac-allowed"] = &ipv4Dhcpd.DuplicateMacAllowed
    children["rate-limit"] = &ipv4Dhcpd.RateLimit
    return children
}

func (ipv4Dhcpd *Ipv4Dhcpd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = ipv4Dhcpd.Enable
    leafs["outer-cos"] = ipv4Dhcpd.OuterCos
    leafs["allow-client-id-change"] = ipv4Dhcpd.AllowClientIdChange
    leafs["inner-cos"] = ipv4Dhcpd.InnerCos
    return leafs
}

func (ipv4Dhcpd *Ipv4Dhcpd) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Dhcpd *Ipv4Dhcpd) GetYangName() string { return "ipv4-dhcpd" }

func (ipv4Dhcpd *Ipv4Dhcpd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Dhcpd *Ipv4Dhcpd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Dhcpd *Ipv4Dhcpd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Dhcpd *Ipv4Dhcpd) SetParent(parent types.Entity) { ipv4Dhcpd.parent = parent }

func (ipv4Dhcpd *Ipv4Dhcpd) GetParent() types.Entity { return ipv4Dhcpd.parent }

func (ipv4Dhcpd *Ipv4Dhcpd) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-dhcpd-cfg" }

// Ipv4Dhcpd_Vrfs
// VRF Table
type Ipv4Dhcpd_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF table. The type is slice of Ipv4Dhcpd_Vrfs_Vrf.
    Vrf []Ipv4Dhcpd_Vrfs_Vrf
}

func (vrfs *Ipv4Dhcpd_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Ipv4Dhcpd_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Ipv4Dhcpd_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Ipv4Dhcpd_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Ipv4Dhcpd_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Ipv4Dhcpd_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Ipv4Dhcpd_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Ipv4Dhcpd_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Ipv4Dhcpd_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Ipv4Dhcpd_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Ipv4Dhcpd_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Ipv4Dhcpd_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Ipv4Dhcpd_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Ipv4Dhcpd_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Ipv4Dhcpd_Vrfs) GetParentYangName() string { return "ipv4-dhcpd" }

// Ipv4Dhcpd_Vrfs_Vrf
// VRF table
type Ipv4Dhcpd_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Profile name and mode.
    Profile Ipv4Dhcpd_Vrfs_Vrf_Profile
}

func (vrf *Ipv4Dhcpd_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Ipv4Dhcpd_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Ipv4Dhcpd_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "profile" { return "Profile" }
    return ""
}

func (vrf *Ipv4Dhcpd_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Ipv4Dhcpd_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        return &vrf.Profile
    }
    return nil
}

func (vrf *Ipv4Dhcpd_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["profile"] = &vrf.Profile
    return children
}

func (vrf *Ipv4Dhcpd_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Ipv4Dhcpd_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Ipv4Dhcpd_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Ipv4Dhcpd_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Ipv4Dhcpd_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Ipv4Dhcpd_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Ipv4Dhcpd_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Ipv4Dhcpd_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Ipv4Dhcpd_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Ipv4Dhcpd_Vrfs_Vrf_Profile
// Profile name and mode
// This type is a presence type.
type Ipv4Dhcpd_Vrfs_Vrf_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Profile name. The type is string. This attribute is mandatory.
    VrfProfileName interface{}

    // Dhcp mode. The type is Ipv4dhcpdMode. This attribute is mandatory.
    Mode interface{}
}

func (profile *Ipv4Dhcpd_Vrfs_Vrf_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Ipv4Dhcpd_Vrfs_Vrf_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Ipv4Dhcpd_Vrfs_Vrf_Profile) GetGoName(yname string) string {
    if yname == "vrf-profile-name" { return "VrfProfileName" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (profile *Ipv4Dhcpd_Vrfs_Vrf_Profile) GetSegmentPath() string {
    return "profile"
}

func (profile *Ipv4Dhcpd_Vrfs_Vrf_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (profile *Ipv4Dhcpd_Vrfs_Vrf_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (profile *Ipv4Dhcpd_Vrfs_Vrf_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-profile-name"] = profile.VrfProfileName
    leafs["mode"] = profile.Mode
    return leafs
}

func (profile *Ipv4Dhcpd_Vrfs_Vrf_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Ipv4Dhcpd_Vrfs_Vrf_Profile) GetYangName() string { return "profile" }

func (profile *Ipv4Dhcpd_Vrfs_Vrf_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Ipv4Dhcpd_Vrfs_Vrf_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Ipv4Dhcpd_Vrfs_Vrf_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Ipv4Dhcpd_Vrfs_Vrf_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Ipv4Dhcpd_Vrfs_Vrf_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Ipv4Dhcpd_Vrfs_Vrf_Profile) GetParentYangName() string { return "vrf" }

// Ipv4Dhcpd_Profiles
// DHCP IPV4 Profile Table
type Ipv4Dhcpd_Profiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP IPV4 Profile. The type is slice of Ipv4Dhcpd_Profiles_Profile.
    Profile []Ipv4Dhcpd_Profiles_Profile
}

func (profiles *Ipv4Dhcpd_Profiles) GetFilter() yfilter.YFilter { return profiles.YFilter }

func (profiles *Ipv4Dhcpd_Profiles) SetFilter(yf yfilter.YFilter) { profiles.YFilter = yf }

func (profiles *Ipv4Dhcpd_Profiles) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (profiles *Ipv4Dhcpd_Profiles) GetSegmentPath() string {
    return "profiles"
}

func (profiles *Ipv4Dhcpd_Profiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        for _, c := range profiles.Profile {
            if profiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile{}
        profiles.Profile = append(profiles.Profile, child)
        return &profiles.Profile[len(profiles.Profile)-1]
    }
    return nil
}

func (profiles *Ipv4Dhcpd_Profiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range profiles.Profile {
        children[profiles.Profile[i].GetSegmentPath()] = &profiles.Profile[i]
    }
    return children
}

func (profiles *Ipv4Dhcpd_Profiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (profiles *Ipv4Dhcpd_Profiles) GetBundleName() string { return "cisco_ios_xr" }

func (profiles *Ipv4Dhcpd_Profiles) GetYangName() string { return "profiles" }

func (profiles *Ipv4Dhcpd_Profiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profiles *Ipv4Dhcpd_Profiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profiles *Ipv4Dhcpd_Profiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profiles *Ipv4Dhcpd_Profiles) SetParent(parent types.Entity) { profiles.parent = parent }

func (profiles *Ipv4Dhcpd_Profiles) GetParent() types.Entity { return profiles.parent }

func (profiles *Ipv4Dhcpd_Profiles) GetParentYangName() string { return "ipv4-dhcpd" }

// Ipv4Dhcpd_Profiles_Profile
// DHCP IPV4 Profile
type Ipv4Dhcpd_Profiles_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Profile Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // DHCP IPV4 Profile modes.
    Modes Ipv4Dhcpd_Profiles_Profile_Modes
}

func (profile *Ipv4Dhcpd_Profiles_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Ipv4Dhcpd_Profiles_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Ipv4Dhcpd_Profiles_Profile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "modes" { return "Modes" }
    return ""
}

func (profile *Ipv4Dhcpd_Profiles_Profile) GetSegmentPath() string {
    return "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
}

func (profile *Ipv4Dhcpd_Profiles_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "modes" {
        return &profile.Modes
    }
    return nil
}

func (profile *Ipv4Dhcpd_Profiles_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["modes"] = &profile.Modes
    return children
}

func (profile *Ipv4Dhcpd_Profiles_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = profile.ProfileName
    return leafs
}

func (profile *Ipv4Dhcpd_Profiles_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Ipv4Dhcpd_Profiles_Profile) GetYangName() string { return "profile" }

func (profile *Ipv4Dhcpd_Profiles_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Ipv4Dhcpd_Profiles_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Ipv4Dhcpd_Profiles_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Ipv4Dhcpd_Profiles_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Ipv4Dhcpd_Profiles_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Ipv4Dhcpd_Profiles_Profile) GetParentYangName() string { return "profiles" }

// Ipv4Dhcpd_Profiles_Profile_Modes
// DHCP IPV4 Profile modes
type Ipv4Dhcpd_Profiles_Profile_Modes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP IPV4 Profile mode. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode.
    Mode []Ipv4Dhcpd_Profiles_Profile_Modes_Mode
}

func (modes *Ipv4Dhcpd_Profiles_Profile_Modes) GetFilter() yfilter.YFilter { return modes.YFilter }

func (modes *Ipv4Dhcpd_Profiles_Profile_Modes) SetFilter(yf yfilter.YFilter) { modes.YFilter = yf }

func (modes *Ipv4Dhcpd_Profiles_Profile_Modes) GetGoName(yname string) string {
    if yname == "mode" { return "Mode" }
    return ""
}

func (modes *Ipv4Dhcpd_Profiles_Profile_Modes) GetSegmentPath() string {
    return "modes"
}

func (modes *Ipv4Dhcpd_Profiles_Profile_Modes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mode" {
        for _, c := range modes.Mode {
            if modes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode{}
        modes.Mode = append(modes.Mode, child)
        return &modes.Mode[len(modes.Mode)-1]
    }
    return nil
}

func (modes *Ipv4Dhcpd_Profiles_Profile_Modes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range modes.Mode {
        children[modes.Mode[i].GetSegmentPath()] = &modes.Mode[i]
    }
    return children
}

func (modes *Ipv4Dhcpd_Profiles_Profile_Modes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (modes *Ipv4Dhcpd_Profiles_Profile_Modes) GetBundleName() string { return "cisco_ios_xr" }

func (modes *Ipv4Dhcpd_Profiles_Profile_Modes) GetYangName() string { return "modes" }

func (modes *Ipv4Dhcpd_Profiles_Profile_Modes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (modes *Ipv4Dhcpd_Profiles_Profile_Modes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (modes *Ipv4Dhcpd_Profiles_Profile_Modes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (modes *Ipv4Dhcpd_Profiles_Profile_Modes) SetParent(parent types.Entity) { modes.parent = parent }

func (modes *Ipv4Dhcpd_Profiles_Profile_Modes) GetParent() types.Entity { return modes.parent }

func (modes *Ipv4Dhcpd_Profiles_Profile_Modes) GetParentYangName() string { return "profile" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode
// DHCP IPV4 Profile mode
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (mode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode) GetFilter() yfilter.YFilter { return mode.YFilter }

func (mode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode) SetFilter(yf yfilter.YFilter) { mode.YFilter = yf }

func (mode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode) GetGoName(yname string) string {
    if yname == "mode" { return "Mode" }
    if yname == "enable" { return "Enable" }
    if yname == "snoop" { return "Snoop" }
    if yname == "base" { return "Base" }
    if yname == "server" { return "Server" }
    if yname == "relay" { return "Relay" }
    if yname == "proxy" { return "Proxy" }
    return ""
}

func (mode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode) GetSegmentPath() string {
    return "mode" + "[mode='" + fmt.Sprintf("%v", mode.Mode) + "']"
}

func (mode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "snoop" {
        return &mode.Snoop
    }
    if childYangName == "base" {
        return &mode.Base
    }
    if childYangName == "server" {
        return &mode.Server
    }
    if childYangName == "relay" {
        return &mode.Relay
    }
    if childYangName == "proxy" {
        return &mode.Proxy
    }
    return nil
}

func (mode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["snoop"] = &mode.Snoop
    children["base"] = &mode.Base
    children["server"] = &mode.Server
    children["relay"] = &mode.Relay
    children["proxy"] = &mode.Proxy
    return children
}

func (mode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mode"] = mode.Mode
    leafs["enable"] = mode.Enable
    return leafs
}

func (mode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode) GetBundleName() string { return "cisco_ios_xr" }

func (mode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode) GetYangName() string { return "mode" }

func (mode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode) SetParent(parent types.Entity) { mode.parent = parent }

func (mode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode) GetParent() types.Entity { return mode.parent }

func (mode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode) GetParentYangName() string { return "modes" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop
// DHCP Snoop profile
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Trusted sources. The type is interface{}.
    Trusted interface{}

    // DHCP Snoop profile.
    RelayInformationOption Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption
}

func (snoop *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop) GetFilter() yfilter.YFilter { return snoop.YFilter }

func (snoop *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop) SetFilter(yf yfilter.YFilter) { snoop.YFilter = yf }

func (snoop *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop) GetGoName(yname string) string {
    if yname == "trusted" { return "Trusted" }
    if yname == "relay-information-option" { return "RelayInformationOption" }
    return ""
}

func (snoop *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop) GetSegmentPath() string {
    return "snoop"
}

func (snoop *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "relay-information-option" {
        return &snoop.RelayInformationOption
    }
    return nil
}

func (snoop *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["relay-information-option"] = &snoop.RelayInformationOption
    return children
}

func (snoop *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["trusted"] = snoop.Trusted
    return leafs
}

func (snoop *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop) GetBundleName() string { return "cisco_ios_xr" }

func (snoop *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop) GetYangName() string { return "snoop" }

func (snoop *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (snoop *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (snoop *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (snoop *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop) SetParent(parent types.Entity) { snoop.parent = parent }

func (snoop *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop) GetParent() types.Entity { return snoop.parent }

func (snoop *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop) GetParentYangName() string { return "mode" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption
// DHCP Snoop profile
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption struct {
    parent types.Entity
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

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption) GetFilter() yfilter.YFilter { return relayInformationOption.YFilter }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption) SetFilter(yf yfilter.YFilter) { relayInformationOption.YFilter = yf }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption) GetGoName(yname string) string {
    if yname == "insert" { return "Insert" }
    if yname == "allow-untrusted" { return "AllowUntrusted" }
    if yname == "policy" { return "Policy" }
    if yname == "remote-id" { return "RemoteId" }
    return ""
}

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption) GetSegmentPath() string {
    return "relay-information-option"
}

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-id" {
        return &relayInformationOption.RemoteId
    }
    return nil
}

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["remote-id"] = &relayInformationOption.RemoteId
    return children
}

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["insert"] = relayInformationOption.Insert
    leafs["allow-untrusted"] = relayInformationOption.AllowUntrusted
    leafs["policy"] = relayInformationOption.Policy
    return leafs
}

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption) GetBundleName() string { return "cisco_ios_xr" }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption) GetYangName() string { return "relay-information-option" }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption) SetParent(parent types.Entity) { relayInformationOption.parent = parent }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption) GetParent() types.Entity { return relayInformationOption.parent }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption) GetParentYangName() string { return "snoop" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId
// Enter remote-id value
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Format type, 1. Hex 2. ASCII. The type is interface{} with range: 1..2.
    FormatType interface{}

    // Enter remote-id value. The type is string.
    RemoteIdValue interface{}
}

func (remoteId *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId) GetFilter() yfilter.YFilter { return remoteId.YFilter }

func (remoteId *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId) SetFilter(yf yfilter.YFilter) { remoteId.YFilter = yf }

func (remoteId *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId) GetGoName(yname string) string {
    if yname == "format-type" { return "FormatType" }
    if yname == "remote-id-value" { return "RemoteIdValue" }
    return ""
}

func (remoteId *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId) GetSegmentPath() string {
    return "remote-id"
}

func (remoteId *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteId *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteId *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["format-type"] = remoteId.FormatType
    leafs["remote-id-value"] = remoteId.RemoteIdValue
    return leafs
}

func (remoteId *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId) GetBundleName() string { return "cisco_ios_xr" }

func (remoteId *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId) GetYangName() string { return "remote-id" }

func (remoteId *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteId *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteId *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteId *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId) SetParent(parent types.Entity) { remoteId.parent = parent }

func (remoteId *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId) GetParent() types.Entity { return remoteId.parent }

func (remoteId *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Snoop_RelayInformationOption_RemoteId) GetParentYangName() string { return "relay-information-option" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base
// DHCP Base Profile
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable the DHCP IPv4 Base Profile. The type is interface{}.
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

func (base *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base) GetFilter() yfilter.YFilter { return base.YFilter }

func (base *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base) SetFilter(yf yfilter.YFilter) { base.YFilter = yf }

func (base *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "default-profile" { return "DefaultProfile" }
    if yname == "match" { return "Match" }
    if yname == "base-relay-opt" { return "BaseRelayOpt" }
    if yname == "base-match" { return "BaseMatch" }
    return ""
}

func (base *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base) GetSegmentPath() string {
    return "base"
}

func (base *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-profile" {
        return &base.DefaultProfile
    }
    if childYangName == "match" {
        return &base.Match
    }
    if childYangName == "base-relay-opt" {
        return &base.BaseRelayOpt
    }
    if childYangName == "base-match" {
        return &base.BaseMatch
    }
    return nil
}

func (base *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-profile"] = &base.DefaultProfile
    children["match"] = &base.Match
    children["base-relay-opt"] = &base.BaseRelayOpt
    children["base-match"] = &base.BaseMatch
    return children
}

func (base *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = base.Enable
    return leafs
}

func (base *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base) GetBundleName() string { return "cisco_ios_xr" }

func (base *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base) GetYangName() string { return "base" }

func (base *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (base *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (base *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (base *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base) SetParent(parent types.Entity) { base.parent = parent }

func (base *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base) GetParent() types.Entity { return base.parent }

func (base *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base) GetParentYangName() string { return "mode" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile
// Enable the default profile
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Profile name. The type is string with length: 1..64.
    ProfileName interface{}

    // none. The type is interface{} with range: -2147483648..2147483647.
    ProfileMode interface{}
}

func (defaultProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile) GetFilter() yfilter.YFilter { return defaultProfile.YFilter }

func (defaultProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile) SetFilter(yf yfilter.YFilter) { defaultProfile.YFilter = yf }

func (defaultProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "profile-mode" { return "ProfileMode" }
    return ""
}

func (defaultProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile) GetSegmentPath() string {
    return "default-profile"
}

func (defaultProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defaultProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defaultProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = defaultProfile.ProfileName
    leafs["profile-mode"] = defaultProfile.ProfileMode
    return leafs
}

func (defaultProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile) GetBundleName() string { return "cisco_ios_xr" }

func (defaultProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile) GetYangName() string { return "default-profile" }

func (defaultProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile) SetParent(parent types.Entity) { defaultProfile.parent = parent }

func (defaultProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile) GetParent() types.Entity { return defaultProfile.parent }

func (defaultProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_DefaultProfile) GetParentYangName() string { return "base" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match
// Insert match keyword
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of Option.
    OptionFilters Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters

    // Table of Option.
    DefOptions Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match) GetFilter() yfilter.YFilter { return match.YFilter }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match) SetFilter(yf yfilter.YFilter) { match.YFilter = yf }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match) GetGoName(yname string) string {
    if yname == "option-filters" { return "OptionFilters" }
    if yname == "def-options" { return "DefOptions" }
    return ""
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match) GetSegmentPath() string {
    return "match"
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "option-filters" {
        return &match.OptionFilters
    }
    if childYangName == "def-options" {
        return &match.DefOptions
    }
    return nil
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["option-filters"] = &match.OptionFilters
    children["def-options"] = &match.DefOptions
    return children
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match) GetBundleName() string { return "cisco_ios_xr" }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match) GetYangName() string { return "match" }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match) SetParent(parent types.Entity) { match.parent = parent }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match) GetParent() types.Entity { return match.parent }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match) GetParentYangName() string { return "base" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters
// Table of Option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify match option. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter.
    OptionFilter []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter
}

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters) GetFilter() yfilter.YFilter { return optionFilters.YFilter }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters) SetFilter(yf yfilter.YFilter) { optionFilters.YFilter = yf }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters) GetGoName(yname string) string {
    if yname == "option-filter" { return "OptionFilter" }
    return ""
}

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters) GetSegmentPath() string {
    return "option-filters"
}

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "option-filter" {
        for _, c := range optionFilters.OptionFilter {
            if optionFilters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter{}
        optionFilters.OptionFilter = append(optionFilters.OptionFilter, child)
        return &optionFilters.OptionFilter[len(optionFilters.OptionFilter)-1]
    }
    return nil
}

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range optionFilters.OptionFilter {
        children[optionFilters.OptionFilter[i].GetSegmentPath()] = &optionFilters.OptionFilter[i]
    }
    return children
}

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters) GetBundleName() string { return "cisco_ios_xr" }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters) GetYangName() string { return "option-filters" }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters) SetParent(parent types.Entity) { optionFilters.parent = parent }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters) GetParent() types.Entity { return optionFilters.parent }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters) GetParentYangName() string { return "match" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter
// Specify match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Match option 60. The type is interface{} with
    // range: -2147483648..2147483647.
    Matchoption interface{}

    // This attribute is a key. Enter hex pattern string. The type is string with
    // length: 1..64.
    Pattern interface{}

    // This attribute is a key. Set constant integer. The type is interface{} with
    // range: -2147483648..2147483647.
    Format interface{}

    // Vendor action. The type is BaseAction.
    OptionAction interface{}
}

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter) GetFilter() yfilter.YFilter { return optionFilter.YFilter }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter) SetFilter(yf yfilter.YFilter) { optionFilter.YFilter = yf }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter) GetGoName(yname string) string {
    if yname == "matchoption" { return "Matchoption" }
    if yname == "pattern" { return "Pattern" }
    if yname == "format" { return "Format" }
    if yname == "option-action" { return "OptionAction" }
    return ""
}

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter) GetSegmentPath() string {
    return "option-filter" + "[matchoption='" + fmt.Sprintf("%v", optionFilter.Matchoption) + "']" + "[pattern='" + fmt.Sprintf("%v", optionFilter.Pattern) + "']" + "[format='" + fmt.Sprintf("%v", optionFilter.Format) + "']"
}

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["matchoption"] = optionFilter.Matchoption
    leafs["pattern"] = optionFilter.Pattern
    leafs["format"] = optionFilter.Format
    leafs["option-action"] = optionFilter.OptionAction
    return leafs
}

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter) GetBundleName() string { return "cisco_ios_xr" }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter) GetYangName() string { return "option-filter" }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter) SetParent(parent types.Entity) { optionFilter.parent = parent }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter) GetParent() types.Entity { return optionFilter.parent }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_OptionFilters_OptionFilter) GetParentYangName() string { return "option-filters" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions
// Table of Option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify match option. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption.
    DefOption []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption
}

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions) GetFilter() yfilter.YFilter { return defOptions.YFilter }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions) SetFilter(yf yfilter.YFilter) { defOptions.YFilter = yf }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions) GetGoName(yname string) string {
    if yname == "def-option" { return "DefOption" }
    return ""
}

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions) GetSegmentPath() string {
    return "def-options"
}

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "def-option" {
        for _, c := range defOptions.DefOption {
            if defOptions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption{}
        defOptions.DefOption = append(defOptions.DefOption, child)
        return &defOptions.DefOption[len(defOptions.DefOption)-1]
    }
    return nil
}

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range defOptions.DefOption {
        children[defOptions.DefOption[i].GetSegmentPath()] = &defOptions.DefOption[i]
    }
    return children
}

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions) GetBundleName() string { return "cisco_ios_xr" }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions) GetYangName() string { return "def-options" }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions) SetParent(parent types.Entity) { defOptions.parent = parent }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions) GetParent() types.Entity { return defOptions.parent }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions) GetParentYangName() string { return "match" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption
// Specify match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Match option 60. The type is interface{} with
    // range: -2147483648..2147483647.
    DefMatchoption interface{}

    // Vendor action. The type is BaseAction. This attribute is mandatory.
    DefMatchaction interface{}
}

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption) GetFilter() yfilter.YFilter { return defOption.YFilter }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption) SetFilter(yf yfilter.YFilter) { defOption.YFilter = yf }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption) GetGoName(yname string) string {
    if yname == "def-matchoption" { return "DefMatchoption" }
    if yname == "def-matchaction" { return "DefMatchaction" }
    return ""
}

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption) GetSegmentPath() string {
    return "def-option" + "[def-matchoption='" + fmt.Sprintf("%v", defOption.DefMatchoption) + "']"
}

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["def-matchoption"] = defOption.DefMatchoption
    leafs["def-matchaction"] = defOption.DefMatchaction
    return leafs
}

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption) GetBundleName() string { return "cisco_ios_xr" }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption) GetYangName() string { return "def-option" }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption) SetParent(parent types.Entity) { defOption.parent = parent }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption) GetParent() types.Entity { return defOption.parent }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_Match_DefOptions_DefOption) GetParentYangName() string { return "def-options" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt
// Insert Relay Agent Information circuit ID
// and remote ID suboptions in client request
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enter remote-id value. The type is string with length: 1..256.
    RemoteId interface{}

    // Specify Relay Agent Information Option authenticate. The type is
    // interface{} with range: -2147483648..2147483647.
    Authenticate interface{}
}

func (baseRelayOpt *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt) GetFilter() yfilter.YFilter { return baseRelayOpt.YFilter }

func (baseRelayOpt *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt) SetFilter(yf yfilter.YFilter) { baseRelayOpt.YFilter = yf }

func (baseRelayOpt *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt) GetGoName(yname string) string {
    if yname == "remote-id" { return "RemoteId" }
    if yname == "authenticate" { return "Authenticate" }
    return ""
}

func (baseRelayOpt *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt) GetSegmentPath() string {
    return "base-relay-opt"
}

func (baseRelayOpt *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (baseRelayOpt *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (baseRelayOpt *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["remote-id"] = baseRelayOpt.RemoteId
    leafs["authenticate"] = baseRelayOpt.Authenticate
    return leafs
}

func (baseRelayOpt *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt) GetBundleName() string { return "cisco_ios_xr" }

func (baseRelayOpt *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt) GetYangName() string { return "base-relay-opt" }

func (baseRelayOpt *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (baseRelayOpt *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (baseRelayOpt *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (baseRelayOpt *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt) SetParent(parent types.Entity) { baseRelayOpt.parent = parent }

func (baseRelayOpt *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt) GetParent() types.Entity { return baseRelayOpt.parent }

func (baseRelayOpt *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseRelayOpt) GetParentYangName() string { return "base" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch
// Insert match keyword
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify match option.
    Options Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options
}

func (baseMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch) GetFilter() yfilter.YFilter { return baseMatch.YFilter }

func (baseMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch) SetFilter(yf yfilter.YFilter) { baseMatch.YFilter = yf }

func (baseMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch) GetGoName(yname string) string {
    if yname == "options" { return "Options" }
    return ""
}

func (baseMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch) GetSegmentPath() string {
    return "base-match"
}

func (baseMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "options" {
        return &baseMatch.Options
    }
    return nil
}

func (baseMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["options"] = &baseMatch.Options
    return children
}

func (baseMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (baseMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch) GetBundleName() string { return "cisco_ios_xr" }

func (baseMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch) GetYangName() string { return "base-match" }

func (baseMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (baseMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (baseMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (baseMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch) SetParent(parent types.Entity) { baseMatch.parent = parent }

func (baseMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch) GetParent() types.Entity { return baseMatch.parent }

func (baseMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch) GetParentYangName() string { return "base" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options
// Specify match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // none. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option.
    Option []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option
}

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options) GetFilter() yfilter.YFilter { return options.YFilter }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options) SetFilter(yf yfilter.YFilter) { options.YFilter = yf }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options) GetGoName(yname string) string {
    if yname == "option" { return "Option" }
    return ""
}

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options) GetSegmentPath() string {
    return "options"
}

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "option" {
        for _, c := range options.Option {
            if options.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option{}
        options.Option = append(options.Option, child)
        return &options.Option[len(options.Option)-1]
    }
    return nil
}

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range options.Option {
        children[options.Option[i].GetSegmentPath()] = &options.Option[i]
    }
    return children
}

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options) GetBundleName() string { return "cisco_ios_xr" }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options) GetYangName() string { return "options" }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options) SetParent(parent types.Entity) { options.parent = parent }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options) GetParent() types.Entity { return options.parent }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options) GetParentYangName() string { return "base-match" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option
// none
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. none. The type is interface{} with range:
    // -2147483648..2147483647.
    Opt60 interface{}

    // This attribute is a key. Enter hex pattern string. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Opt60HexStr interface{}

    // This attribute is a key. Set constant integer. The type is interface{} with
    // range: -2147483648..2147483647.
    Format interface{}

    // Enter a profile.
    OptionProfile Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option) GetFilter() yfilter.YFilter { return option.YFilter }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option) SetFilter(yf yfilter.YFilter) { option.YFilter = yf }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option) GetGoName(yname string) string {
    if yname == "opt60" { return "Opt60" }
    if yname == "opt60-hex-str" { return "Opt60HexStr" }
    if yname == "format" { return "Format" }
    if yname == "option-profile" { return "OptionProfile" }
    return ""
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option) GetSegmentPath() string {
    return "option" + "[opt60='" + fmt.Sprintf("%v", option.Opt60) + "']" + "[opt60-hex-str='" + fmt.Sprintf("%v", option.Opt60HexStr) + "']" + "[format='" + fmt.Sprintf("%v", option.Format) + "']"
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "option-profile" {
        return &option.OptionProfile
    }
    return nil
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["option-profile"] = &option.OptionProfile
    return children
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["opt60"] = option.Opt60
    leafs["opt60-hex-str"] = option.Opt60HexStr
    leafs["format"] = option.Format
    return leafs
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option) GetBundleName() string { return "cisco_ios_xr" }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option) GetYangName() string { return "option" }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option) SetParent(parent types.Entity) { option.parent = parent }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option) GetParent() types.Entity { return option.parent }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option) GetParentYangName() string { return "options" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile
// Enter a profile
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Profile name. The type is string with length: 1..64.
    ProfileName interface{}

    // none. The type is interface{} with range: -2147483648..2147483647.
    ProfileMode interface{}
}

func (optionProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile) GetFilter() yfilter.YFilter { return optionProfile.YFilter }

func (optionProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile) SetFilter(yf yfilter.YFilter) { optionProfile.YFilter = yf }

func (optionProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "profile-mode" { return "ProfileMode" }
    return ""
}

func (optionProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile) GetSegmentPath() string {
    return "option-profile"
}

func (optionProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (optionProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (optionProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = optionProfile.ProfileName
    leafs["profile-mode"] = optionProfile.ProfileMode
    return leafs
}

func (optionProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile) GetBundleName() string { return "cisco_ios_xr" }

func (optionProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile) GetYangName() string { return "option-profile" }

func (optionProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (optionProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (optionProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (optionProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile) SetParent(parent types.Entity) { optionProfile.parent = parent }

func (optionProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile) GetParent() types.Entity { return optionProfile.parent }

func (optionProfile *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Base_BaseMatch_Options_Option_OptionProfile) GetParentYangName() string { return "option" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server
// DHCP Server profile
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Allow dhcp subscriber move. The type is interface{}.
    ServerAllowMove interface{}

    // Configure Subnet Mask. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    SubnetMask interface{}

    // Specify the Pool name. The type is string with length: 1..64.
    Pool interface{}

    // Infinite lease. The type is interface{}.
    InfiniteLease interface{}

    // Domain name. The type is string with length: 1..256.
    DomainName interface{}

    // Enable Secure Arp. The type is interface{}.
    SecureArp interface{}

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

    // Table of OptionCode.
    OptionCodes Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes
}

func (server *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server) GetFilter() yfilter.YFilter { return server.YFilter }

func (server *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server) SetFilter(yf yfilter.YFilter) { server.YFilter = yf }

func (server *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server) GetGoName(yname string) string {
    if yname == "server-allow-move" { return "ServerAllowMove" }
    if yname == "subnet-mask" { return "SubnetMask" }
    if yname == "pool" { return "Pool" }
    if yname == "infinite-lease" { return "InfiniteLease" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "secure-arp" { return "SecureArp" }
    if yname == "boot-filename" { return "BootFilename" }
    if yname == "next-server" { return "NextServer" }
    if yname == "server-id-check" { return "ServerIdCheck" }
    if yname == "lease-limit" { return "LeaseLimit" }
    if yname == "requested-ip-address" { return "RequestedIpAddress" }
    if yname == "aaa-server" { return "AaaServer" }
    if yname == "default-routers" { return "DefaultRouters" }
    if yname == "net-bios-name-servers" { return "NetBiosNameServers" }
    if yname == "match" { return "Match" }
    if yname == "broadcast-flag" { return "BroadcastFlag" }
    if yname == "session" { return "Session" }
    if yname == "classes" { return "Classes" }
    if yname == "relay" { return "Relay" }
    if yname == "lease" { return "Lease" }
    if yname == "netbios-node-type" { return "NetbiosNodeType" }
    if yname == "dns-servers" { return "DnsServers" }
    if yname == "option-codes" { return "OptionCodes" }
    return ""
}

func (server *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server) GetSegmentPath() string {
    return "server"
}

func (server *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "server-id-check" {
        return &server.ServerIdCheck
    }
    if childYangName == "lease-limit" {
        return &server.LeaseLimit
    }
    if childYangName == "requested-ip-address" {
        return &server.RequestedIpAddress
    }
    if childYangName == "aaa-server" {
        return &server.AaaServer
    }
    if childYangName == "default-routers" {
        return &server.DefaultRouters
    }
    if childYangName == "net-bios-name-servers" {
        return &server.NetBiosNameServers
    }
    if childYangName == "match" {
        return &server.Match
    }
    if childYangName == "broadcast-flag" {
        return &server.BroadcastFlag
    }
    if childYangName == "session" {
        return &server.Session
    }
    if childYangName == "classes" {
        return &server.Classes
    }
    if childYangName == "relay" {
        return &server.Relay
    }
    if childYangName == "lease" {
        return &server.Lease
    }
    if childYangName == "netbios-node-type" {
        return &server.NetbiosNodeType
    }
    if childYangName == "dns-servers" {
        return &server.DnsServers
    }
    if childYangName == "option-codes" {
        return &server.OptionCodes
    }
    return nil
}

func (server *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["server-id-check"] = &server.ServerIdCheck
    children["lease-limit"] = &server.LeaseLimit
    children["requested-ip-address"] = &server.RequestedIpAddress
    children["aaa-server"] = &server.AaaServer
    children["default-routers"] = &server.DefaultRouters
    children["net-bios-name-servers"] = &server.NetBiosNameServers
    children["match"] = &server.Match
    children["broadcast-flag"] = &server.BroadcastFlag
    children["session"] = &server.Session
    children["classes"] = &server.Classes
    children["relay"] = &server.Relay
    children["lease"] = &server.Lease
    children["netbios-node-type"] = &server.NetbiosNodeType
    children["dns-servers"] = &server.DnsServers
    children["option-codes"] = &server.OptionCodes
    return children
}

func (server *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-allow-move"] = server.ServerAllowMove
    leafs["subnet-mask"] = server.SubnetMask
    leafs["pool"] = server.Pool
    leafs["infinite-lease"] = server.InfiniteLease
    leafs["domain-name"] = server.DomainName
    leafs["secure-arp"] = server.SecureArp
    leafs["boot-filename"] = server.BootFilename
    leafs["next-server"] = server.NextServer
    return leafs
}

func (server *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server) GetBundleName() string { return "cisco_ios_xr" }

func (server *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server) GetYangName() string { return "server" }

func (server *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (server *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (server *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (server *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server) SetParent(parent types.Entity) { server.parent = parent }

func (server *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server) GetParent() types.Entity { return server.parent }

func (server *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server) GetParentYangName() string { return "mode" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck
// Validate server ID check
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // specify server-id-check disable. The type is interface{}.
    Check interface{}
}

func (serverIdCheck *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck) GetFilter() yfilter.YFilter { return serverIdCheck.YFilter }

func (serverIdCheck *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck) SetFilter(yf yfilter.YFilter) { serverIdCheck.YFilter = yf }

func (serverIdCheck *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck) GetGoName(yname string) string {
    if yname == "check" { return "Check" }
    return ""
}

func (serverIdCheck *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck) GetSegmentPath() string {
    return "server-id-check"
}

func (serverIdCheck *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serverIdCheck *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serverIdCheck *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["check"] = serverIdCheck.Check
    return leafs
}

func (serverIdCheck *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck) GetBundleName() string { return "cisco_ios_xr" }

func (serverIdCheck *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck) GetYangName() string { return "server-id-check" }

func (serverIdCheck *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serverIdCheck *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serverIdCheck *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serverIdCheck *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck) SetParent(parent types.Entity) { serverIdCheck.parent = parent }

func (serverIdCheck *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck) GetParent() types.Entity { return serverIdCheck.parent }

func (serverIdCheck *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_ServerIdCheck) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit
// Specify limit lease
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure Lease limit value. The type is LeaseLimitValue.
    LeaseLimitValue interface{}

    // Value of limit lease count in Decimal. The type is interface{} with range:
    // 1..240000.
    Range interface{}
}

func (leaseLimit *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit) GetFilter() yfilter.YFilter { return leaseLimit.YFilter }

func (leaseLimit *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit) SetFilter(yf yfilter.YFilter) { leaseLimit.YFilter = yf }

func (leaseLimit *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit) GetGoName(yname string) string {
    if yname == "lease-limit-value" { return "LeaseLimitValue" }
    if yname == "range" { return "Range" }
    return ""
}

func (leaseLimit *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit) GetSegmentPath() string {
    return "lease-limit"
}

func (leaseLimit *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseLimit *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseLimit *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["lease-limit-value"] = leaseLimit.LeaseLimitValue
    leafs["range"] = leaseLimit.Range
    return leafs
}

func (leaseLimit *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit) GetBundleName() string { return "cisco_ios_xr" }

func (leaseLimit *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit) GetYangName() string { return "lease-limit" }

func (leaseLimit *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseLimit *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseLimit *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseLimit *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit) SetParent(parent types.Entity) { leaseLimit.parent = parent }

func (leaseLimit *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit) GetParent() types.Entity { return leaseLimit.parent }

func (leaseLimit *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_LeaseLimit) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress
// Validate Requested IP Address
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // specify requested-ip-address-check disable. The type is interface{}.
    Check interface{}
}

func (requestedIpAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress) GetFilter() yfilter.YFilter { return requestedIpAddress.YFilter }

func (requestedIpAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress) SetFilter(yf yfilter.YFilter) { requestedIpAddress.YFilter = yf }

func (requestedIpAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress) GetGoName(yname string) string {
    if yname == "check" { return "Check" }
    return ""
}

func (requestedIpAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress) GetSegmentPath() string {
    return "requested-ip-address"
}

func (requestedIpAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (requestedIpAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (requestedIpAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["check"] = requestedIpAddress.Check
    return leafs
}

func (requestedIpAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (requestedIpAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress) GetYangName() string { return "requested-ip-address" }

func (requestedIpAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (requestedIpAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (requestedIpAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (requestedIpAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress) SetParent(parent types.Entity) { requestedIpAddress.parent = parent }

func (requestedIpAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress) GetParent() types.Entity { return requestedIpAddress.parent }

func (requestedIpAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_RequestedIpAddress) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer
// Enable aaa dhcp option force-insert
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable aaa dhcp option force-insert.
    DhcpOption Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption
}

func (aaaServer *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer) GetFilter() yfilter.YFilter { return aaaServer.YFilter }

func (aaaServer *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer) SetFilter(yf yfilter.YFilter) { aaaServer.YFilter = yf }

func (aaaServer *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer) GetGoName(yname string) string {
    if yname == "dhcp-option" { return "DhcpOption" }
    return ""
}

func (aaaServer *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer) GetSegmentPath() string {
    return "aaa-server"
}

func (aaaServer *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dhcp-option" {
        return &aaaServer.DhcpOption
    }
    return nil
}

func (aaaServer *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dhcp-option"] = &aaaServer.DhcpOption
    return children
}

func (aaaServer *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (aaaServer *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer) GetBundleName() string { return "cisco_ios_xr" }

func (aaaServer *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer) GetYangName() string { return "aaa-server" }

func (aaaServer *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (aaaServer *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (aaaServer *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (aaaServer *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer) SetParent(parent types.Entity) { aaaServer.parent = parent }

func (aaaServer *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer) GetParent() types.Entity { return aaaServer.parent }

func (aaaServer *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption
// Enable aaa dhcp option force-insert
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable aaa dhcp option force-insert. The type is interface{}.
    ForceInsert interface{}
}

func (dhcpOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption) GetFilter() yfilter.YFilter { return dhcpOption.YFilter }

func (dhcpOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption) SetFilter(yf yfilter.YFilter) { dhcpOption.YFilter = yf }

func (dhcpOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption) GetGoName(yname string) string {
    if yname == "force-insert" { return "ForceInsert" }
    return ""
}

func (dhcpOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption) GetSegmentPath() string {
    return "dhcp-option"
}

func (dhcpOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dhcpOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dhcpOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["force-insert"] = dhcpOption.ForceInsert
    return leafs
}

func (dhcpOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption) GetBundleName() string { return "cisco_ios_xr" }

func (dhcpOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption) GetYangName() string { return "dhcp-option" }

func (dhcpOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcpOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcpOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcpOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption) SetParent(parent types.Entity) { dhcpOption.parent = parent }

func (dhcpOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption) GetParent() types.Entity { return dhcpOption.parent }

func (dhcpOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_AaaServer_DhcpOption) GetParentYangName() string { return "aaa-server" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters
// default routers
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Router's IP address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DefaultRouter []interface{}
}

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters) GetFilter() yfilter.YFilter { return defaultRouters.YFilter }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters) SetFilter(yf yfilter.YFilter) { defaultRouters.YFilter = yf }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters) GetGoName(yname string) string {
    if yname == "default-router" { return "DefaultRouter" }
    return ""
}

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters) GetSegmentPath() string {
    return "default-routers"
}

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["default-router"] = defaultRouters.DefaultRouter
    return leafs
}

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters) GetBundleName() string { return "cisco_ios_xr" }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters) GetYangName() string { return "default-routers" }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters) SetParent(parent types.Entity) { defaultRouters.parent = parent }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters) GetParent() types.Entity { return defaultRouters.parent }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DefaultRouters) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers
// NetBIOS name servers
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NetBIOSNameServer's IP address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NetBiosNameServer []interface{}
}

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers) GetFilter() yfilter.YFilter { return netBiosNameServers.YFilter }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers) SetFilter(yf yfilter.YFilter) { netBiosNameServers.YFilter = yf }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers) GetGoName(yname string) string {
    if yname == "net-bios-name-server" { return "NetBiosNameServer" }
    return ""
}

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers) GetSegmentPath() string {
    return "net-bios-name-servers"
}

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["net-bios-name-server"] = netBiosNameServers.NetBiosNameServer
    return leafs
}

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers) GetBundleName() string { return "cisco_ios_xr" }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers) GetYangName() string { return "net-bios-name-servers" }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers) SetParent(parent types.Entity) { netBiosNameServers.parent = parent }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers) GetParent() types.Entity { return netBiosNameServers.parent }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetBiosNameServers) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match
// Insert match keyword
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of OptionDefault.
    OptionDefaults Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults

    // Table of Option.
    Options Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match) GetFilter() yfilter.YFilter { return match.YFilter }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match) SetFilter(yf yfilter.YFilter) { match.YFilter = yf }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match) GetGoName(yname string) string {
    if yname == "option-defaults" { return "OptionDefaults" }
    if yname == "options" { return "Options" }
    return ""
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match) GetSegmentPath() string {
    return "match"
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "option-defaults" {
        return &match.OptionDefaults
    }
    if childYangName == "options" {
        return &match.Options
    }
    return nil
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["option-defaults"] = &match.OptionDefaults
    children["options"] = &match.Options
    return children
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match) GetBundleName() string { return "cisco_ios_xr" }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match) GetYangName() string { return "match" }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match) SetParent(parent types.Entity) { match.parent = parent }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match) GetParent() types.Entity { return match.parent }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults
// Table of OptionDefault
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify match option. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault.
    OptionDefault []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault
}

func (optionDefaults *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults) GetFilter() yfilter.YFilter { return optionDefaults.YFilter }

func (optionDefaults *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults) SetFilter(yf yfilter.YFilter) { optionDefaults.YFilter = yf }

func (optionDefaults *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults) GetGoName(yname string) string {
    if yname == "option-default" { return "OptionDefault" }
    return ""
}

func (optionDefaults *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults) GetSegmentPath() string {
    return "option-defaults"
}

func (optionDefaults *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "option-default" {
        for _, c := range optionDefaults.OptionDefault {
            if optionDefaults.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault{}
        optionDefaults.OptionDefault = append(optionDefaults.OptionDefault, child)
        return &optionDefaults.OptionDefault[len(optionDefaults.OptionDefault)-1]
    }
    return nil
}

func (optionDefaults *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range optionDefaults.OptionDefault {
        children[optionDefaults.OptionDefault[i].GetSegmentPath()] = &optionDefaults.OptionDefault[i]
    }
    return children
}

func (optionDefaults *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (optionDefaults *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults) GetBundleName() string { return "cisco_ios_xr" }

func (optionDefaults *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults) GetYangName() string { return "option-defaults" }

func (optionDefaults *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (optionDefaults *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (optionDefaults *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (optionDefaults *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults) SetParent(parent types.Entity) { optionDefaults.parent = parent }

func (optionDefaults *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults) GetParent() types.Entity { return optionDefaults.parent }

func (optionDefaults *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults) GetParentYangName() string { return "match" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault
// Specify match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Match option 60. The type is Matchoption.
    Matchoption interface{}

    // Vendor action. The type is Matchaction. This attribute is mandatory.
    Matchaction interface{}
}

func (optionDefault *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault) GetFilter() yfilter.YFilter { return optionDefault.YFilter }

func (optionDefault *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault) SetFilter(yf yfilter.YFilter) { optionDefault.YFilter = yf }

func (optionDefault *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault) GetGoName(yname string) string {
    if yname == "matchoption" { return "Matchoption" }
    if yname == "matchaction" { return "Matchaction" }
    return ""
}

func (optionDefault *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault) GetSegmentPath() string {
    return "option-default" + "[matchoption='" + fmt.Sprintf("%v", optionDefault.Matchoption) + "']"
}

func (optionDefault *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (optionDefault *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (optionDefault *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["matchoption"] = optionDefault.Matchoption
    leafs["matchaction"] = optionDefault.Matchaction
    return leafs
}

func (optionDefault *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault) GetBundleName() string { return "cisco_ios_xr" }

func (optionDefault *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault) GetYangName() string { return "option-default" }

func (optionDefault *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (optionDefault *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (optionDefault *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (optionDefault *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault) SetParent(parent types.Entity) { optionDefault.parent = parent }

func (optionDefault *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault) GetParent() types.Entity { return optionDefault.parent }

func (optionDefault *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_OptionDefaults_OptionDefault) GetParentYangName() string { return "option-defaults" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options
// Table of Option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify match option. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option.
    Option []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option
}

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options) GetFilter() yfilter.YFilter { return options.YFilter }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options) SetFilter(yf yfilter.YFilter) { options.YFilter = yf }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options) GetGoName(yname string) string {
    if yname == "option" { return "Option" }
    return ""
}

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options) GetSegmentPath() string {
    return "options"
}

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "option" {
        for _, c := range options.Option {
            if options.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option{}
        options.Option = append(options.Option, child)
        return &options.Option[len(options.Option)-1]
    }
    return nil
}

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range options.Option {
        children[options.Option[i].GetSegmentPath()] = &options.Option[i]
    }
    return children
}

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options) GetBundleName() string { return "cisco_ios_xr" }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options) GetYangName() string { return "options" }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options) SetParent(parent types.Entity) { options.parent = parent }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options) GetParent() types.Entity { return options.parent }

func (options *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options) GetParentYangName() string { return "match" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option
// Specify match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Match option 60. The type is Matchoption.
    Matchoption interface{}

    // This attribute is a key. Enter hex pattern string. The type is string with
    // length: 1..64.
    Pattern interface{}

    // This attribute is a key. Set constant integer. The type is interface{} with
    // range: -2147483648..2147483647.
    Format interface{}

    // Vendor action. The type is Matchaction. This attribute is mandatory.
    Matchaction interface{}
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option) GetFilter() yfilter.YFilter { return option.YFilter }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option) SetFilter(yf yfilter.YFilter) { option.YFilter = yf }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option) GetGoName(yname string) string {
    if yname == "matchoption" { return "Matchoption" }
    if yname == "pattern" { return "Pattern" }
    if yname == "format" { return "Format" }
    if yname == "matchaction" { return "Matchaction" }
    return ""
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option) GetSegmentPath() string {
    return "option" + "[matchoption='" + fmt.Sprintf("%v", option.Matchoption) + "']" + "[pattern='" + fmt.Sprintf("%v", option.Pattern) + "']" + "[format='" + fmt.Sprintf("%v", option.Format) + "']"
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["matchoption"] = option.Matchoption
    leafs["pattern"] = option.Pattern
    leafs["format"] = option.Format
    leafs["matchaction"] = option.Matchaction
    return leafs
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option) GetBundleName() string { return "cisco_ios_xr" }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option) GetYangName() string { return "option" }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option) SetParent(parent types.Entity) { option.parent = parent }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option) GetParent() types.Entity { return option.parent }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Match_Options_Option) GetParentYangName() string { return "options" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag
// None
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify broadcast flag policy. The type is Policy.
    Policy interface{}
}

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag) GetFilter() yfilter.YFilter { return broadcastFlag.YFilter }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag) SetFilter(yf yfilter.YFilter) { broadcastFlag.YFilter = yf }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag) GetGoName(yname string) string {
    if yname == "policy" { return "Policy" }
    return ""
}

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag) GetSegmentPath() string {
    return "broadcast-flag"
}

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy"] = broadcastFlag.Policy
    return leafs
}

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag) GetBundleName() string { return "cisco_ios_xr" }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag) GetYangName() string { return "broadcast-flag" }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag) SetParent(parent types.Entity) { broadcastFlag.parent = parent }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag) GetParent() types.Entity { return broadcastFlag.parent }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_BroadcastFlag) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session
// Change sessions configuration
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Throttle DHCP sessions based on MAC address.
    ThrottleType Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType
}

func (session *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session) GetFilter() yfilter.YFilter { return session.YFilter }

func (session *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session) SetFilter(yf yfilter.YFilter) { session.YFilter = yf }

func (session *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session) GetGoName(yname string) string {
    if yname == "throttle-type" { return "ThrottleType" }
    return ""
}

func (session *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session) GetSegmentPath() string {
    return "session"
}

func (session *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "throttle-type" {
        return &session.ThrottleType
    }
    return nil
}

func (session *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["throttle-type"] = &session.ThrottleType
    return children
}

func (session *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (session *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session) GetBundleName() string { return "cisco_ios_xr" }

func (session *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session) GetYangName() string { return "session" }

func (session *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (session *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (session *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (session *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session) SetParent(parent types.Entity) { session.parent = parent }

func (session *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session) GetParent() types.Entity { return session.parent }

func (session *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType
// Throttle DHCP sessions based on MAC
// address
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Throttle DHCP sessions from any one MAC address.
    MacThrottle Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle
}

func (throttleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType) GetFilter() yfilter.YFilter { return throttleType.YFilter }

func (throttleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType) SetFilter(yf yfilter.YFilter) { throttleType.YFilter = yf }

func (throttleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType) GetGoName(yname string) string {
    if yname == "mac-throttle" { return "MacThrottle" }
    return ""
}

func (throttleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType) GetSegmentPath() string {
    return "throttle-type"
}

func (throttleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac-throttle" {
        return &throttleType.MacThrottle
    }
    return nil
}

func (throttleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac-throttle"] = &throttleType.MacThrottle
    return children
}

func (throttleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (throttleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType) GetBundleName() string { return "cisco_ios_xr" }

func (throttleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType) GetYangName() string { return "throttle-type" }

func (throttleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (throttleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (throttleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (throttleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType) SetParent(parent types.Entity) { throttleType.parent = parent }

func (throttleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType) GetParent() types.Entity { return throttleType.parent }

func (throttleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType) GetParentYangName() string { return "session" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle
// Throttle DHCP sessions from any one MAC
// address
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle struct {
    parent types.Entity
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

func (macThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle) GetFilter() yfilter.YFilter { return macThrottle.YFilter }

func (macThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle) SetFilter(yf yfilter.YFilter) { macThrottle.YFilter = yf }

func (macThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle) GetGoName(yname string) string {
    if yname == "num-discover" { return "NumDiscover" }
    if yname == "num-request" { return "NumRequest" }
    if yname == "num-block" { return "NumBlock" }
    return ""
}

func (macThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle) GetSegmentPath() string {
    return "mac-throttle"
}

func (macThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (macThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (macThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-discover"] = macThrottle.NumDiscover
    leafs["num-request"] = macThrottle.NumRequest
    leafs["num-block"] = macThrottle.NumBlock
    return leafs
}

func (macThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle) GetBundleName() string { return "cisco_ios_xr" }

func (macThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle) GetYangName() string { return "mac-throttle" }

func (macThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (macThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (macThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (macThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle) SetParent(parent types.Entity) { macThrottle.parent = parent }

func (macThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle) GetParent() types.Entity { return macThrottle.parent }

func (macThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Session_ThrottleType_MacThrottle) GetParentYangName() string { return "throttle-type" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes
// Table of Class
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Create or enter server profile class. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class.
    Class []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class
}

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes) GetFilter() yfilter.YFilter { return classes.YFilter }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes) SetFilter(yf yfilter.YFilter) { classes.YFilter = yf }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes) GetGoName(yname string) string {
    if yname == "class" { return "Class" }
    return ""
}

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes) GetSegmentPath() string {
    return "classes"
}

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class" {
        for _, c := range classes.Class {
            if classes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class{}
        classes.Class = append(classes.Class, child)
        return &classes.Class[len(classes.Class)-1]
    }
    return nil
}

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classes.Class {
        children[classes.Class[i].GetSegmentPath()] = &classes.Class[i]
    }
    return children
}

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes) GetBundleName() string { return "cisco_ios_xr" }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes) GetYangName() string { return "classes" }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes) SetParent(parent types.Entity) { classes.parent = parent }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes) GetParent() types.Entity { return classes.parent }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class
// Create or enter server profile class
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

    // Infinite lease. The type is interface{}.
    InfiniteLease interface{}

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

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class) GetFilter() yfilter.YFilter { return class.YFilter }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class) SetFilter(yf yfilter.YFilter) { class.YFilter = yf }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class) GetGoName(yname string) string {
    if yname == "class-name" { return "ClassName" }
    if yname == "subnet-mask" { return "SubnetMask" }
    if yname == "pool" { return "Pool" }
    if yname == "enable" { return "Enable" }
    if yname == "infinite-lease" { return "InfiniteLease" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "boot-filename" { return "BootFilename" }
    if yname == "next-server" { return "NextServer" }
    if yname == "default-routers" { return "DefaultRouters" }
    if yname == "net-bios-name-servers" { return "NetBiosNameServers" }
    if yname == "class-match" { return "ClassMatch" }
    if yname == "lease" { return "Lease" }
    if yname == "netbios-node-type" { return "NetbiosNodeType" }
    if yname == "dns-servers" { return "DnsServers" }
    if yname == "option-codes" { return "OptionCodes" }
    return ""
}

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class) GetSegmentPath() string {
    return "class" + "[class-name='" + fmt.Sprintf("%v", class.ClassName) + "']"
}

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-routers" {
        return &class.DefaultRouters
    }
    if childYangName == "net-bios-name-servers" {
        return &class.NetBiosNameServers
    }
    if childYangName == "class-match" {
        return &class.ClassMatch
    }
    if childYangName == "lease" {
        return &class.Lease
    }
    if childYangName == "netbios-node-type" {
        return &class.NetbiosNodeType
    }
    if childYangName == "dns-servers" {
        return &class.DnsServers
    }
    if childYangName == "option-codes" {
        return &class.OptionCodes
    }
    return nil
}

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-routers"] = &class.DefaultRouters
    children["net-bios-name-servers"] = &class.NetBiosNameServers
    children["class-match"] = &class.ClassMatch
    children["lease"] = &class.Lease
    children["netbios-node-type"] = &class.NetbiosNodeType
    children["dns-servers"] = &class.DnsServers
    children["option-codes"] = &class.OptionCodes
    return children
}

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-name"] = class.ClassName
    leafs["subnet-mask"] = class.SubnetMask
    leafs["pool"] = class.Pool
    leafs["enable"] = class.Enable
    leafs["infinite-lease"] = class.InfiniteLease
    leafs["domain-name"] = class.DomainName
    leafs["boot-filename"] = class.BootFilename
    leafs["next-server"] = class.NextServer
    return leafs
}

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class) GetBundleName() string { return "cisco_ios_xr" }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class) GetYangName() string { return "class" }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class) SetParent(parent types.Entity) { class.parent = parent }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class) GetParent() types.Entity { return class.parent }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class) GetParentYangName() string { return "classes" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters
// default routers
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Router's IP address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DefaultRouter []interface{}
}

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters) GetFilter() yfilter.YFilter { return defaultRouters.YFilter }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters) SetFilter(yf yfilter.YFilter) { defaultRouters.YFilter = yf }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters) GetGoName(yname string) string {
    if yname == "default-router" { return "DefaultRouter" }
    return ""
}

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters) GetSegmentPath() string {
    return "default-routers"
}

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["default-router"] = defaultRouters.DefaultRouter
    return leafs
}

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters) GetBundleName() string { return "cisco_ios_xr" }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters) GetYangName() string { return "default-routers" }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters) SetParent(parent types.Entity) { defaultRouters.parent = parent }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters) GetParent() types.Entity { return defaultRouters.parent }

func (defaultRouters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DefaultRouters) GetParentYangName() string { return "class" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers
// NetBIOS name servers
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // NetBIOSNameServer's IP address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    NetBiosNameServer []interface{}
}

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers) GetFilter() yfilter.YFilter { return netBiosNameServers.YFilter }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers) SetFilter(yf yfilter.YFilter) { netBiosNameServers.YFilter = yf }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers) GetGoName(yname string) string {
    if yname == "net-bios-name-server" { return "NetBiosNameServer" }
    return ""
}

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers) GetSegmentPath() string {
    return "net-bios-name-servers"
}

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["net-bios-name-server"] = netBiosNameServers.NetBiosNameServer
    return leafs
}

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers) GetBundleName() string { return "cisco_ios_xr" }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers) GetYangName() string { return "net-bios-name-servers" }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers) SetParent(parent types.Entity) { netBiosNameServers.parent = parent }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers) GetParent() types.Entity { return netBiosNameServers.parent }

func (netBiosNameServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetBiosNameServers) GetParentYangName() string { return "class" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch
// Insert match keyword
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify match l2-interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    L2Interface interface{}

    // Specify match VRF. The type is string with length: 1..32.
    Vrf interface{}

    // Table of Class-Option.
    ClassOptions Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions
}

func (classMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch) GetFilter() yfilter.YFilter { return classMatch.YFilter }

func (classMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch) SetFilter(yf yfilter.YFilter) { classMatch.YFilter = yf }

func (classMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch) GetGoName(yname string) string {
    if yname == "l2-interface" { return "L2Interface" }
    if yname == "vrf" { return "Vrf" }
    if yname == "class-options" { return "ClassOptions" }
    return ""
}

func (classMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch) GetSegmentPath() string {
    return "class-match"
}

func (classMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class-options" {
        return &classMatch.ClassOptions
    }
    return nil
}

func (classMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["class-options"] = &classMatch.ClassOptions
    return children
}

func (classMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["l2-interface"] = classMatch.L2Interface
    leafs["vrf"] = classMatch.Vrf
    return leafs
}

func (classMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch) GetBundleName() string { return "cisco_ios_xr" }

func (classMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch) GetYangName() string { return "class-match" }

func (classMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch) SetParent(parent types.Entity) { classMatch.parent = parent }

func (classMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch) GetParent() types.Entity { return classMatch.parent }

func (classMatch *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch) GetParentYangName() string { return "class" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions
// Table of Class-Option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify match option. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption.
    ClassOption []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption
}

func (classOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions) GetFilter() yfilter.YFilter { return classOptions.YFilter }

func (classOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions) SetFilter(yf yfilter.YFilter) { classOptions.YFilter = yf }

func (classOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions) GetGoName(yname string) string {
    if yname == "class-option" { return "ClassOption" }
    return ""
}

func (classOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions) GetSegmentPath() string {
    return "class-options"
}

func (classOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class-option" {
        for _, c := range classOptions.ClassOption {
            if classOptions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption{}
        classOptions.ClassOption = append(classOptions.ClassOption, child)
        return &classOptions.ClassOption[len(classOptions.ClassOption)-1]
    }
    return nil
}

func (classOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classOptions.ClassOption {
        children[classOptions.ClassOption[i].GetSegmentPath()] = &classOptions.ClassOption[i]
    }
    return children
}

func (classOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (classOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions) GetBundleName() string { return "cisco_ios_xr" }

func (classOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions) GetYangName() string { return "class-options" }

func (classOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions) SetParent(parent types.Entity) { classOptions.parent = parent }

func (classOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions) GetParent() types.Entity { return classOptions.parent }

func (classOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions) GetParentYangName() string { return "class-match" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption
// Specify match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Match options. The type is Matchoption.
    Matchoption interface{}

    // Enter hex pattern string. The type is string with length: 1..64.
    Pattern interface{}

    // Enter bit mask pattern string. The type is string with length: 1..64.
    BitMask interface{}
}

func (classOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption) GetFilter() yfilter.YFilter { return classOption.YFilter }

func (classOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption) SetFilter(yf yfilter.YFilter) { classOption.YFilter = yf }

func (classOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption) GetGoName(yname string) string {
    if yname == "matchoption" { return "Matchoption" }
    if yname == "pattern" { return "Pattern" }
    if yname == "bit-mask" { return "BitMask" }
    return ""
}

func (classOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption) GetSegmentPath() string {
    return "class-option" + "[matchoption='" + fmt.Sprintf("%v", classOption.Matchoption) + "']"
}

func (classOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (classOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (classOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["matchoption"] = classOption.Matchoption
    leafs["pattern"] = classOption.Pattern
    leafs["bit-mask"] = classOption.BitMask
    return leafs
}

func (classOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption) GetBundleName() string { return "cisco_ios_xr" }

func (classOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption) GetYangName() string { return "class-option" }

func (classOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption) SetParent(parent types.Entity) { classOption.parent = parent }

func (classOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption) GetParent() types.Entity { return classOption.parent }

func (classOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_ClassMatch_ClassOptions_ClassOption) GetParentYangName() string { return "class-options" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease
// lease
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease struct {
    parent types.Entity
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

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease) GetFilter() yfilter.YFilter { return lease.YFilter }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease) SetFilter(yf yfilter.YFilter) { lease.YFilter = yf }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease) GetGoName(yname string) string {
    if yname == "infinite" { return "Infinite" }
    if yname == "days" { return "Days" }
    if yname == "hours" { return "Hours" }
    if yname == "minutes" { return "Minutes" }
    return ""
}

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease) GetSegmentPath() string {
    return "lease"
}

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["infinite"] = lease.Infinite
    leafs["days"] = lease.Days
    leafs["hours"] = lease.Hours
    leafs["minutes"] = lease.Minutes
    return leafs
}

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease) GetBundleName() string { return "cisco_ios_xr" }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease) GetYangName() string { return "lease" }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease) SetParent(parent types.Entity) { lease.parent = parent }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease) GetParent() types.Entity { return lease.parent }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_Lease) GetParentYangName() string { return "class" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType
// NetBIOS node type
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType struct {
    parent types.Entity
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

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType) GetFilter() yfilter.YFilter { return netbiosNodeType.YFilter }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType) SetFilter(yf yfilter.YFilter) { netbiosNodeType.YFilter = yf }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType) GetGoName(yname string) string {
    if yname == "broadcast-node" { return "BroadcastNode" }
    if yname == "hybrid-node" { return "HybridNode" }
    if yname == "mixed-node" { return "MixedNode" }
    if yname == "peer-to-peer-node" { return "PeerToPeerNode" }
    if yname == "hexadecimal" { return "Hexadecimal" }
    return ""
}

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType) GetSegmentPath() string {
    return "netbios-node-type"
}

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["broadcast-node"] = netbiosNodeType.BroadcastNode
    leafs["hybrid-node"] = netbiosNodeType.HybridNode
    leafs["mixed-node"] = netbiosNodeType.MixedNode
    leafs["peer-to-peer-node"] = netbiosNodeType.PeerToPeerNode
    leafs["hexadecimal"] = netbiosNodeType.Hexadecimal
    return leafs
}

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType) GetBundleName() string { return "cisco_ios_xr" }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType) GetYangName() string { return "netbios-node-type" }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType) SetParent(parent types.Entity) { netbiosNodeType.parent = parent }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType) GetParent() types.Entity { return netbiosNodeType.parent }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_NetbiosNodeType) GetParentYangName() string { return "class" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers
// DNS servers
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DNS Server's IP address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DnsServer []interface{}
}

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers) GetFilter() yfilter.YFilter { return dnsServers.YFilter }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers) SetFilter(yf yfilter.YFilter) { dnsServers.YFilter = yf }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers) GetGoName(yname string) string {
    if yname == "dns-server" { return "DnsServer" }
    return ""
}

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers) GetSegmentPath() string {
    return "dns-servers"
}

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dns-server"] = dnsServers.DnsServer
    return leafs
}

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers) GetBundleName() string { return "cisco_ios_xr" }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers) GetYangName() string { return "dns-servers" }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers) SetParent(parent types.Entity) { dnsServers.parent = parent }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers) GetParent() types.Entity { return dnsServers.parent }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_DnsServers) GetParentYangName() string { return "class" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes
// Table of OptionCode
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP option code. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode.
    OptionCode []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode
}

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes) GetFilter() yfilter.YFilter { return optionCodes.YFilter }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes) SetFilter(yf yfilter.YFilter) { optionCodes.YFilter = yf }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes) GetGoName(yname string) string {
    if yname == "option-code" { return "OptionCode" }
    return ""
}

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes) GetSegmentPath() string {
    return "option-codes"
}

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "option-code" {
        for _, c := range optionCodes.OptionCode {
            if optionCodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode{}
        optionCodes.OptionCode = append(optionCodes.OptionCode, child)
        return &optionCodes.OptionCode[len(optionCodes.OptionCode)-1]
    }
    return nil
}

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range optionCodes.OptionCode {
        children[optionCodes.OptionCode[i].GetSegmentPath()] = &optionCodes.OptionCode[i]
    }
    return children
}

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes) GetBundleName() string { return "cisco_ios_xr" }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes) GetYangName() string { return "option-codes" }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes) SetParent(parent types.Entity) { optionCodes.parent = parent }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes) GetParent() types.Entity { return optionCodes.parent }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes) GetParentYangName() string { return "class" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode
// DHCP option code
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. DHCP option code. The type is interface{} with
    // range: 0..255.
    OptionCode interface{}

    // ASCII string. The type is string.
    AsciiString interface{}

    // Hexadecimal string. The type is string.
    HexString interface{}

    // Set constant integer. The type is interface{} with range:
    // -2147483648..2147483647.
    ForceInsert interface{}

    // Server's IP address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress []interface{}
}

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode) GetFilter() yfilter.YFilter { return optionCode.YFilter }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode) SetFilter(yf yfilter.YFilter) { optionCode.YFilter = yf }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode) GetGoName(yname string) string {
    if yname == "option-code" { return "OptionCode" }
    if yname == "ascii-string" { return "AsciiString" }
    if yname == "hex-string" { return "HexString" }
    if yname == "force-insert" { return "ForceInsert" }
    if yname == "ip-address" { return "IpAddress" }
    return ""
}

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode) GetSegmentPath() string {
    return "option-code" + "[option-code='" + fmt.Sprintf("%v", optionCode.OptionCode) + "']"
}

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["option-code"] = optionCode.OptionCode
    leafs["ascii-string"] = optionCode.AsciiString
    leafs["hex-string"] = optionCode.HexString
    leafs["force-insert"] = optionCode.ForceInsert
    leafs["ip-address"] = optionCode.IpAddress
    return leafs
}

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode) GetBundleName() string { return "cisco_ios_xr" }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode) GetYangName() string { return "option-code" }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode) SetParent(parent types.Entity) { optionCode.parent = parent }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode) GetParent() types.Entity { return optionCode.parent }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Classes_Class_OptionCodes_OptionCode) GetParentYangName() string { return "option-codes" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay
// Specify Relay Agent Information Option
// configuration
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify Relay Agent Information Option authenticate. The type is
    // interface{} with range: -2147483648..2147483647.
    Authenticate interface{}
}

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay) GetFilter() yfilter.YFilter { return relay.YFilter }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay) SetFilter(yf yfilter.YFilter) { relay.YFilter = yf }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay) GetGoName(yname string) string {
    if yname == "authenticate" { return "Authenticate" }
    return ""
}

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay) GetSegmentPath() string {
    return "relay"
}

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["authenticate"] = relay.Authenticate
    return leafs
}

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay) GetBundleName() string { return "cisco_ios_xr" }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay) GetYangName() string { return "relay" }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay) SetParent(parent types.Entity) { relay.parent = parent }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay) GetParent() types.Entity { return relay.parent }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Relay) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease
// lease
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease struct {
    parent types.Entity
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

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease) GetFilter() yfilter.YFilter { return lease.YFilter }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease) SetFilter(yf yfilter.YFilter) { lease.YFilter = yf }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease) GetGoName(yname string) string {
    if yname == "infinite" { return "Infinite" }
    if yname == "days" { return "Days" }
    if yname == "hours" { return "Hours" }
    if yname == "minutes" { return "Minutes" }
    return ""
}

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease) GetSegmentPath() string {
    return "lease"
}

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["infinite"] = lease.Infinite
    leafs["days"] = lease.Days
    leafs["hours"] = lease.Hours
    leafs["minutes"] = lease.Minutes
    return leafs
}

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease) GetBundleName() string { return "cisco_ios_xr" }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease) GetYangName() string { return "lease" }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease) SetParent(parent types.Entity) { lease.parent = parent }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease) GetParent() types.Entity { return lease.parent }

func (lease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_Lease) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType
// NetBIOS node type
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType struct {
    parent types.Entity
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

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType) GetFilter() yfilter.YFilter { return netbiosNodeType.YFilter }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType) SetFilter(yf yfilter.YFilter) { netbiosNodeType.YFilter = yf }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType) GetGoName(yname string) string {
    if yname == "broadcast-node" { return "BroadcastNode" }
    if yname == "hybrid-node" { return "HybridNode" }
    if yname == "mixed-node" { return "MixedNode" }
    if yname == "peer-to-peer-node" { return "PeerToPeerNode" }
    if yname == "hexadecimal" { return "Hexadecimal" }
    return ""
}

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType) GetSegmentPath() string {
    return "netbios-node-type"
}

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["broadcast-node"] = netbiosNodeType.BroadcastNode
    leafs["hybrid-node"] = netbiosNodeType.HybridNode
    leafs["mixed-node"] = netbiosNodeType.MixedNode
    leafs["peer-to-peer-node"] = netbiosNodeType.PeerToPeerNode
    leafs["hexadecimal"] = netbiosNodeType.Hexadecimal
    return leafs
}

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType) GetBundleName() string { return "cisco_ios_xr" }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType) GetYangName() string { return "netbios-node-type" }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType) SetParent(parent types.Entity) { netbiosNodeType.parent = parent }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType) GetParent() types.Entity { return netbiosNodeType.parent }

func (netbiosNodeType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_NetbiosNodeType) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers
// DNS servers
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DNS Server's IP address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    DnsServer []interface{}
}

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers) GetFilter() yfilter.YFilter { return dnsServers.YFilter }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers) SetFilter(yf yfilter.YFilter) { dnsServers.YFilter = yf }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers) GetGoName(yname string) string {
    if yname == "dns-server" { return "DnsServer" }
    return ""
}

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers) GetSegmentPath() string {
    return "dns-servers"
}

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dns-server"] = dnsServers.DnsServer
    return leafs
}

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers) GetBundleName() string { return "cisco_ios_xr" }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers) GetYangName() string { return "dns-servers" }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers) SetParent(parent types.Entity) { dnsServers.parent = parent }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers) GetParent() types.Entity { return dnsServers.parent }

func (dnsServers *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_DnsServers) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes
// Table of OptionCode
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP option code. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode.
    OptionCode []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode
}

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes) GetFilter() yfilter.YFilter { return optionCodes.YFilter }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes) SetFilter(yf yfilter.YFilter) { optionCodes.YFilter = yf }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes) GetGoName(yname string) string {
    if yname == "option-code" { return "OptionCode" }
    return ""
}

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes) GetSegmentPath() string {
    return "option-codes"
}

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "option-code" {
        for _, c := range optionCodes.OptionCode {
            if optionCodes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode{}
        optionCodes.OptionCode = append(optionCodes.OptionCode, child)
        return &optionCodes.OptionCode[len(optionCodes.OptionCode)-1]
    }
    return nil
}

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range optionCodes.OptionCode {
        children[optionCodes.OptionCode[i].GetSegmentPath()] = &optionCodes.OptionCode[i]
    }
    return children
}

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes) GetBundleName() string { return "cisco_ios_xr" }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes) GetYangName() string { return "option-codes" }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes) SetParent(parent types.Entity) { optionCodes.parent = parent }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes) GetParent() types.Entity { return optionCodes.parent }

func (optionCodes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes) GetParentYangName() string { return "server" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode
// DHCP option code
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. DHCP option code. The type is interface{} with
    // range: 0..255.
    OptionCode interface{}

    // ASCII string. The type is string.
    AsciiString interface{}

    // Hexadecimal string. The type is string.
    HexString interface{}

    // Set constant integer. The type is interface{} with range:
    // -2147483648..2147483647.
    ForceInsert interface{}

    // Server's IP address. The type is slice of string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress []interface{}
}

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode) GetFilter() yfilter.YFilter { return optionCode.YFilter }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode) SetFilter(yf yfilter.YFilter) { optionCode.YFilter = yf }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode) GetGoName(yname string) string {
    if yname == "option-code" { return "OptionCode" }
    if yname == "ascii-string" { return "AsciiString" }
    if yname == "hex-string" { return "HexString" }
    if yname == "force-insert" { return "ForceInsert" }
    if yname == "ip-address" { return "IpAddress" }
    return ""
}

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode) GetSegmentPath() string {
    return "option-code" + "[option-code='" + fmt.Sprintf("%v", optionCode.OptionCode) + "']"
}

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["option-code"] = optionCode.OptionCode
    leafs["ascii-string"] = optionCode.AsciiString
    leafs["hex-string"] = optionCode.HexString
    leafs["force-insert"] = optionCode.ForceInsert
    leafs["ip-address"] = optionCode.IpAddress
    return leafs
}

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode) GetBundleName() string { return "cisco_ios_xr" }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode) GetYangName() string { return "option-code" }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode) SetParent(parent types.Entity) { optionCode.parent = parent }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode) GetParent() types.Entity { return optionCode.parent }

func (optionCode *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Server_OptionCodes_OptionCode) GetParentYangName() string { return "option-codes" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay
// DHCP Relay profile
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay struct {
    parent types.Entity
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

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay) GetFilter() yfilter.YFilter { return relay.YFilter }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay) SetFilter(yf yfilter.YFilter) { relay.YFilter = yf }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay) GetGoName(yname string) string {
    if yname == "mac-mismatch-action" { return "MacMismatchAction" }
    if yname == "gi-addr-policy" { return "GiAddrPolicy" }
    if yname == "vrfs" { return "Vrfs" }
    if yname == "relay-information-option" { return "RelayInformationOption" }
    if yname == "broadcast-policy" { return "BroadcastPolicy" }
    return ""
}

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay) GetSegmentPath() string {
    return "relay"
}

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "gi-addr-policy" {
        return &relay.GiAddrPolicy
    }
    if childYangName == "vrfs" {
        return &relay.Vrfs
    }
    if childYangName == "relay-information-option" {
        return &relay.RelayInformationOption
    }
    if childYangName == "broadcast-policy" {
        return &relay.BroadcastPolicy
    }
    return nil
}

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["gi-addr-policy"] = &relay.GiAddrPolicy
    children["vrfs"] = &relay.Vrfs
    children["relay-information-option"] = &relay.RelayInformationOption
    children["broadcast-policy"] = &relay.BroadcastPolicy
    return children
}

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-mismatch-action"] = relay.MacMismatchAction
    return leafs
}

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay) GetBundleName() string { return "cisco_ios_xr" }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay) GetYangName() string { return "relay" }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay) SetParent(parent types.Entity) { relay.parent = parent }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay) GetParent() types.Entity { return relay.parent }

func (relay *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay) GetParentYangName() string { return "mode" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy
// GIADDR policy
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // GIADDR policy. The type is Ipv4dhcpdGiaddrPolicy.
    Policy interface{}
}

func (giAddrPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy) GetFilter() yfilter.YFilter { return giAddrPolicy.YFilter }

func (giAddrPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy) SetFilter(yf yfilter.YFilter) { giAddrPolicy.YFilter = yf }

func (giAddrPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy) GetGoName(yname string) string {
    if yname == "policy" { return "Policy" }
    return ""
}

func (giAddrPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy) GetSegmentPath() string {
    return "gi-addr-policy"
}

func (giAddrPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (giAddrPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (giAddrPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy"] = giAddrPolicy.Policy
    return leafs
}

func (giAddrPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy) GetBundleName() string { return "cisco_ios_xr" }

func (giAddrPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy) GetYangName() string { return "gi-addr-policy" }

func (giAddrPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (giAddrPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (giAddrPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (giAddrPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy) SetParent(parent types.Entity) { giAddrPolicy.parent = parent }

func (giAddrPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy) GetParent() types.Entity { return giAddrPolicy.parent }

func (giAddrPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_GiAddrPolicy) GetParentYangName() string { return "relay" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs
// VRF Helper Addresses
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF Name. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf.
    Vrf []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs) GetParentYangName() string { return "relay" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf
// VRF Name
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF Name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Helper Addresses.
    HelperAddresses Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "helper-addresses" { return "HelperAddresses" }
    return ""
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "helper-addresses" {
        return &vrf.HelperAddresses
    }
    return nil
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["helper-addresses"] = &vrf.HelperAddresses
    return children
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses
// Helper Addresses
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Helper Address. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress.
    HelperAddress []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses) GetFilter() yfilter.YFilter { return helperAddresses.YFilter }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses) SetFilter(yf yfilter.YFilter) { helperAddresses.YFilter = yf }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses) GetGoName(yname string) string {
    if yname == "helper-address" { return "HelperAddress" }
    return ""
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses) GetSegmentPath() string {
    return "helper-addresses"
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "helper-address" {
        for _, c := range helperAddresses.HelperAddress {
            if helperAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress{}
        helperAddresses.HelperAddress = append(helperAddresses.HelperAddress, child)
        return &helperAddresses.HelperAddress[len(helperAddresses.HelperAddress)-1]
    }
    return nil
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range helperAddresses.HelperAddress {
        children[helperAddresses.HelperAddress[i].GetSegmentPath()] = &helperAddresses.HelperAddress[i]
    }
    return children
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses) GetYangName() string { return "helper-addresses" }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses) SetParent(parent types.Entity) { helperAddresses.parent = parent }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses) GetParent() types.Entity { return helperAddresses.parent }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses) GetParentYangName() string { return "vrf" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress
// Helper Address
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPV4 Address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // Enable helper - deprecated. The type is interface{}.
    Enable interface{}

    // GatewayAddress. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    GatewayAddress interface{}
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress) GetFilter() yfilter.YFilter { return helperAddress.YFilter }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress) SetFilter(yf yfilter.YFilter) { helperAddress.YFilter = yf }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "enable" { return "Enable" }
    if yname == "gateway-address" { return "GatewayAddress" }
    return ""
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress) GetSegmentPath() string {
    return "helper-address" + "[ip-address='" + fmt.Sprintf("%v", helperAddress.IpAddress) + "']"
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = helperAddress.IpAddress
    leafs["enable"] = helperAddress.Enable
    leafs["gateway-address"] = helperAddress.GatewayAddress
    return leafs
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress) GetBundleName() string { return "cisco_ios_xr" }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress) GetYangName() string { return "helper-address" }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress) SetParent(parent types.Entity) { helperAddress.parent = parent }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress) GetParent() types.Entity { return helperAddress.parent }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_Vrfs_Vrf_HelperAddresses_HelperAddress) GetParentYangName() string { return "helper-addresses" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption
// Relay agent information option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption struct {
    parent types.Entity
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

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption) GetFilter() yfilter.YFilter { return relayInformationOption.YFilter }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption) SetFilter(yf yfilter.YFilter) { relayInformationOption.YFilter = yf }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption) GetGoName(yname string) string {
    if yname == "vpn-mode" { return "VpnMode" }
    if yname == "subscriber-id" { return "SubscriberId" }
    if yname == "insert" { return "Insert" }
    if yname == "check" { return "Check" }
    if yname == "vpn" { return "Vpn" }
    if yname == "allow-untrusted" { return "AllowUntrusted" }
    if yname == "policy" { return "Policy" }
    return ""
}

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption) GetSegmentPath() string {
    return "relay-information-option"
}

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vpn-mode"] = relayInformationOption.VpnMode
    leafs["subscriber-id"] = relayInformationOption.SubscriberId
    leafs["insert"] = relayInformationOption.Insert
    leafs["check"] = relayInformationOption.Check
    leafs["vpn"] = relayInformationOption.Vpn
    leafs["allow-untrusted"] = relayInformationOption.AllowUntrusted
    leafs["policy"] = relayInformationOption.Policy
    return leafs
}

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption) GetBundleName() string { return "cisco_ios_xr" }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption) GetYangName() string { return "relay-information-option" }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption) SetParent(parent types.Entity) { relayInformationOption.parent = parent }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption) GetParent() types.Entity { return relayInformationOption.parent }

func (relayInformationOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_RelayInformationOption) GetParentYangName() string { return "relay" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy
// Broadcast Flag policy
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Broadcast flag policy. The type is Ipv4dhcpdBroadcastFlagPolicy.
    Policy interface{}
}

func (broadcastPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy) GetFilter() yfilter.YFilter { return broadcastPolicy.YFilter }

func (broadcastPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy) SetFilter(yf yfilter.YFilter) { broadcastPolicy.YFilter = yf }

func (broadcastPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy) GetGoName(yname string) string {
    if yname == "policy" { return "Policy" }
    return ""
}

func (broadcastPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy) GetSegmentPath() string {
    return "broadcast-policy"
}

func (broadcastPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (broadcastPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (broadcastPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy"] = broadcastPolicy.Policy
    return leafs
}

func (broadcastPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy) GetBundleName() string { return "cisco_ios_xr" }

func (broadcastPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy) GetYangName() string { return "broadcast-policy" }

func (broadcastPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (broadcastPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (broadcastPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (broadcastPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy) SetParent(parent types.Entity) { broadcastPolicy.parent = parent }

func (broadcastPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy) GetParent() types.Entity { return broadcastPolicy.parent }

func (broadcastPolicy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Relay_BroadcastPolicy) GetParentYangName() string { return "relay" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy
// DHCP proxy profile
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Allow dhcp subscriber move. The type is interface{}.
    ProxyAllowMove interface{}

    // DHCP IPV4 profile proxy secure-arp enable. The type is interface{}.
    SecureArp interface{}

    // For BNG session, delay the authentication. The type is interface{}.
    DelayedAuthenProxy interface{}

    // DHCP IPV4 profile mode enable. The type is interface{}.
    Enable interface{}

    // Specify gateway address policy.
    Giaddr Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr

    // DHCP class table.
    Classes Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes

    // Authentication Username formating.
    AuthUsername Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername

    // Relay agent information option.
    RelayInformation Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation

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

func (proxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy) GetFilter() yfilter.YFilter { return proxy.YFilter }

func (proxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy) SetFilter(yf yfilter.YFilter) { proxy.YFilter = yf }

func (proxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy) GetGoName(yname string) string {
    if yname == "proxy-allow-move" { return "ProxyAllowMove" }
    if yname == "secure-arp" { return "SecureArp" }
    if yname == "delayed-authen-proxy" { return "DelayedAuthenProxy" }
    if yname == "enable" { return "Enable" }
    if yname == "giaddr" { return "Giaddr" }
    if yname == "classes" { return "Classes" }
    if yname == "auth-username" { return "AuthUsername" }
    if yname == "relay-information" { return "RelayInformation" }
    if yname == "vrfs" { return "Vrfs" }
    if yname == "sessions" { return "Sessions" }
    if yname == "limit-lease" { return "LimitLease" }
    if yname == "lease-proxy" { return "LeaseProxy" }
    if yname == "broadcast-flag" { return "BroadcastFlag" }
    if yname == "match" { return "Match" }
    return ""
}

func (proxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy) GetSegmentPath() string {
    return "proxy"
}

func (proxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "giaddr" {
        return &proxy.Giaddr
    }
    if childYangName == "classes" {
        return &proxy.Classes
    }
    if childYangName == "auth-username" {
        return &proxy.AuthUsername
    }
    if childYangName == "relay-information" {
        return &proxy.RelayInformation
    }
    if childYangName == "vrfs" {
        return &proxy.Vrfs
    }
    if childYangName == "sessions" {
        return &proxy.Sessions
    }
    if childYangName == "limit-lease" {
        return &proxy.LimitLease
    }
    if childYangName == "lease-proxy" {
        return &proxy.LeaseProxy
    }
    if childYangName == "broadcast-flag" {
        return &proxy.BroadcastFlag
    }
    if childYangName == "match" {
        return &proxy.Match
    }
    return nil
}

func (proxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["giaddr"] = &proxy.Giaddr
    children["classes"] = &proxy.Classes
    children["auth-username"] = &proxy.AuthUsername
    children["relay-information"] = &proxy.RelayInformation
    children["vrfs"] = &proxy.Vrfs
    children["sessions"] = &proxy.Sessions
    children["limit-lease"] = &proxy.LimitLease
    children["lease-proxy"] = &proxy.LeaseProxy
    children["broadcast-flag"] = &proxy.BroadcastFlag
    children["match"] = &proxy.Match
    return children
}

func (proxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["proxy-allow-move"] = proxy.ProxyAllowMove
    leafs["secure-arp"] = proxy.SecureArp
    leafs["delayed-authen-proxy"] = proxy.DelayedAuthenProxy
    leafs["enable"] = proxy.Enable
    return leafs
}

func (proxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy) GetBundleName() string { return "cisco_ios_xr" }

func (proxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy) GetYangName() string { return "proxy" }

func (proxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (proxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (proxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (proxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy) SetParent(parent types.Entity) { proxy.parent = parent }

func (proxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy) GetParent() types.Entity { return proxy.parent }

func (proxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy) GetParentYangName() string { return "mode" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr
// Specify gateway address policy
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Gateway address policy. The type is Ipv4dhcpdGiaddrPolicy.
    Policy interface{}
}

func (giaddr *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr) GetFilter() yfilter.YFilter { return giaddr.YFilter }

func (giaddr *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr) SetFilter(yf yfilter.YFilter) { giaddr.YFilter = yf }

func (giaddr *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr) GetGoName(yname string) string {
    if yname == "policy" { return "Policy" }
    return ""
}

func (giaddr *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr) GetSegmentPath() string {
    return "giaddr"
}

func (giaddr *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (giaddr *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (giaddr *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy"] = giaddr.Policy
    return leafs
}

func (giaddr *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr) GetBundleName() string { return "cisco_ios_xr" }

func (giaddr *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr) GetYangName() string { return "giaddr" }

func (giaddr *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (giaddr *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (giaddr *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (giaddr *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr) SetParent(parent types.Entity) { giaddr.parent = parent }

func (giaddr *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr) GetParent() types.Entity { return giaddr.parent }

func (giaddr *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Giaddr) GetParentYangName() string { return "proxy" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes
// DHCP class table
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP class. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class.
    Class []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class
}

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes) GetFilter() yfilter.YFilter { return classes.YFilter }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes) SetFilter(yf yfilter.YFilter) { classes.YFilter = yf }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes) GetGoName(yname string) string {
    if yname == "class" { return "Class" }
    return ""
}

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes) GetSegmentPath() string {
    return "classes"
}

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class" {
        for _, c := range classes.Class {
            if classes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class{}
        classes.Class = append(classes.Class, child)
        return &classes.Class[len(classes.Class)-1]
    }
    return nil
}

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classes.Class {
        children[classes.Class[i].GetSegmentPath()] = &classes.Class[i]
    }
    return children
}

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes) GetBundleName() string { return "cisco_ios_xr" }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes) GetYangName() string { return "classes" }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes) SetParent(parent types.Entity) { classes.parent = parent }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes) GetParent() types.Entity { return classes.parent }

func (classes *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes) GetParentYangName() string { return "proxy" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class
// DHCP class
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Class name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ClassName interface{}

    // Enable the DHCP IPV4 proxy class. The type is interface{}.
    Enable interface{}

    // Match option.
    Match Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match

    // List of VRFs.
    Vrfs Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs
}

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class) GetFilter() yfilter.YFilter { return class.YFilter }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class) SetFilter(yf yfilter.YFilter) { class.YFilter = yf }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class) GetGoName(yname string) string {
    if yname == "class-name" { return "ClassName" }
    if yname == "enable" { return "Enable" }
    if yname == "match" { return "Match" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class) GetSegmentPath() string {
    return "class" + "[class-name='" + fmt.Sprintf("%v", class.ClassName) + "']"
}

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "match" {
        return &class.Match
    }
    if childYangName == "vrfs" {
        return &class.Vrfs
    }
    return nil
}

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["match"] = &class.Match
    children["vrfs"] = &class.Vrfs
    return children
}

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-name"] = class.ClassName
    leafs["enable"] = class.Enable
    return leafs
}

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class) GetBundleName() string { return "cisco_ios_xr" }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class) GetYangName() string { return "class" }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class) SetParent(parent types.Entity) { class.parent = parent }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class) GetParent() types.Entity { return class.parent }

func (class *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class) GetParentYangName() string { return "classes" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match
// Match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Match VRF name. The type is string.
    Vrf interface{}

    // Match option.
    Option Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match) GetFilter() yfilter.YFilter { return match.YFilter }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match) SetFilter(yf yfilter.YFilter) { match.YFilter = yf }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    if yname == "option" { return "Option" }
    return ""
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match) GetSegmentPath() string {
    return "match"
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "option" {
        return &match.Option
    }
    return nil
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["option"] = &match.Option
    return children
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf"] = match.Vrf
    return leafs
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match) GetBundleName() string { return "cisco_ios_xr" }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match) GetYangName() string { return "match" }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match) SetParent(parent types.Entity) { match.parent = parent }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match) GetParent() types.Entity { return match.parent }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match) GetParentYangName() string { return "class" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option
// Match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Match option. The type is Dhcpv4MatchOption.
    OptionType interface{}

    // Hex pattern string. The type is string.
    Pattern interface{}

    // Bit mask pattern. The type is string.
    BitMask interface{}
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option) GetFilter() yfilter.YFilter { return option.YFilter }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option) SetFilter(yf yfilter.YFilter) { option.YFilter = yf }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option) GetGoName(yname string) string {
    if yname == "option-type" { return "OptionType" }
    if yname == "pattern" { return "Pattern" }
    if yname == "bit-mask" { return "BitMask" }
    return ""
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option) GetSegmentPath() string {
    return "option"
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["option-type"] = option.OptionType
    leafs["pattern"] = option.Pattern
    leafs["bit-mask"] = option.BitMask
    return leafs
}

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option) GetBundleName() string { return "cisco_ios_xr" }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option) GetYangName() string { return "option" }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option) SetParent(parent types.Entity) { option.parent = parent }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option) GetParent() types.Entity { return option.parent }

func (option *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Match_Option) GetParentYangName() string { return "match" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs
// List of VRFs
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf.
    Vrf []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs) GetParentYangName() string { return "class" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf
// VRF name
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Helper addresses.
    HelperAddresses Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "helper-addresses" { return "HelperAddresses" }
    return ""
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "helper-addresses" {
        return &vrf.HelperAddresses
    }
    return nil
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["helper-addresses"] = &vrf.HelperAddresses
    return children
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses
// Helper addresses
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Helper address. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress.
    HelperAddress []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses) GetFilter() yfilter.YFilter { return helperAddresses.YFilter }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses) SetFilter(yf yfilter.YFilter) { helperAddresses.YFilter = yf }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses) GetGoName(yname string) string {
    if yname == "helper-address" { return "HelperAddress" }
    return ""
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses) GetSegmentPath() string {
    return "helper-addresses"
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "helper-address" {
        for _, c := range helperAddresses.HelperAddress {
            if helperAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress{}
        helperAddresses.HelperAddress = append(helperAddresses.HelperAddress, child)
        return &helperAddresses.HelperAddress[len(helperAddresses.HelperAddress)-1]
    }
    return nil
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range helperAddresses.HelperAddress {
        children[helperAddresses.HelperAddress[i].GetSegmentPath()] = &helperAddresses.HelperAddress[i]
    }
    return children
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses) GetYangName() string { return "helper-addresses" }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses) SetParent(parent types.Entity) { helperAddresses.parent = parent }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses) GetParent() types.Entity { return helperAddresses.parent }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses) GetParentYangName() string { return "vrf" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress
// Helper address
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ServerAddress interface{}

    // Gateway address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    GatewayAddress interface{}
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress) GetFilter() yfilter.YFilter { return helperAddress.YFilter }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress) SetFilter(yf yfilter.YFilter) { helperAddress.YFilter = yf }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress) GetGoName(yname string) string {
    if yname == "server-address" { return "ServerAddress" }
    if yname == "gateway-address" { return "GatewayAddress" }
    return ""
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress) GetSegmentPath() string {
    return "helper-address" + "[server-address='" + fmt.Sprintf("%v", helperAddress.ServerAddress) + "']"
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-address"] = helperAddress.ServerAddress
    leafs["gateway-address"] = helperAddress.GatewayAddress
    return leafs
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress) GetBundleName() string { return "cisco_ios_xr" }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress) GetYangName() string { return "helper-address" }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress) SetParent(parent types.Entity) { helperAddress.parent = parent }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress) GetParent() types.Entity { return helperAddress.parent }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Classes_Class_Vrfs_Vrf_HelperAddresses_HelperAddress) GetParentYangName() string { return "helper-addresses" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername
// Authentication Username formating
// This type is a presence type.
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Username Formatting first argument . The type is Dhcpv4AuthUsername. This
    // attribute is mandatory.
    Arg1 interface{}

    // Username Formatting second argument . The type is Dhcpv4AuthUsername.
    Arg2 interface{}
}

func (authUsername *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername) GetFilter() yfilter.YFilter { return authUsername.YFilter }

func (authUsername *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername) SetFilter(yf yfilter.YFilter) { authUsername.YFilter = yf }

func (authUsername *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername) GetGoName(yname string) string {
    if yname == "arg1" { return "Arg1" }
    if yname == "arg2" { return "Arg2" }
    return ""
}

func (authUsername *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername) GetSegmentPath() string {
    return "auth-username"
}

func (authUsername *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authUsername *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authUsername *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["arg1"] = authUsername.Arg1
    leafs["arg2"] = authUsername.Arg2
    return leafs
}

func (authUsername *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername) GetBundleName() string { return "cisco_ios_xr" }

func (authUsername *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername) GetYangName() string { return "auth-username" }

func (authUsername *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authUsername *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authUsername *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authUsername *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername) SetParent(parent types.Entity) { authUsername.parent = parent }

func (authUsername *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername) GetParent() types.Entity { return authUsername.parent }

func (authUsername *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_AuthUsername) GetParentYangName() string { return "proxy" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation
// Relay agent information option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation struct {
    parent types.Entity
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

func (relayInformation *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation) GetFilter() yfilter.YFilter { return relayInformation.YFilter }

func (relayInformation *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation) SetFilter(yf yfilter.YFilter) { relayInformation.YFilter = yf }

func (relayInformation *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation) GetGoName(yname string) string {
    if yname == "option" { return "Option" }
    if yname == "vpn" { return "Vpn" }
    if yname == "allow-untrusted" { return "AllowUntrusted" }
    if yname == "circuit-id" { return "CircuitId" }
    if yname == "policy" { return "Policy" }
    if yname == "vpn-mode" { return "VpnMode" }
    if yname == "remote-id-xr" { return "RemoteIdXr" }
    if yname == "remote-id-suppress" { return "RemoteIdSuppress" }
    if yname == "check" { return "Check" }
    if yname == "remote-id" { return "RemoteId" }
    if yname == "authenticate" { return "Authenticate" }
    return ""
}

func (relayInformation *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation) GetSegmentPath() string {
    return "relay-information"
}

func (relayInformation *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (relayInformation *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (relayInformation *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["option"] = relayInformation.Option
    leafs["vpn"] = relayInformation.Vpn
    leafs["allow-untrusted"] = relayInformation.AllowUntrusted
    leafs["circuit-id"] = relayInformation.CircuitId
    leafs["policy"] = relayInformation.Policy
    leafs["vpn-mode"] = relayInformation.VpnMode
    leafs["remote-id-xr"] = relayInformation.RemoteIdXr
    leafs["remote-id-suppress"] = relayInformation.RemoteIdSuppress
    leafs["check"] = relayInformation.Check
    leafs["remote-id"] = relayInformation.RemoteId
    leafs["authenticate"] = relayInformation.Authenticate
    return leafs
}

func (relayInformation *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation) GetBundleName() string { return "cisco_ios_xr" }

func (relayInformation *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation) GetYangName() string { return "relay-information" }

func (relayInformation *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relayInformation *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relayInformation *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relayInformation *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation) SetParent(parent types.Entity) { relayInformation.parent = parent }

func (relayInformation *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation) GetParent() types.Entity { return relayInformation.parent }

func (relayInformation *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_RelayInformation) GetParentYangName() string { return "proxy" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs
// List of VRFs
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf.
    Vrf []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs) GetParentYangName() string { return "proxy" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf
// VRF name
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Helper addresses.
    HelperAddresses Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "helper-addresses" { return "HelperAddresses" }
    return ""
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "helper-addresses" {
        return &vrf.HelperAddresses
    }
    return nil
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["helper-addresses"] = &vrf.HelperAddresses
    return children
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses
// Helper addresses
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Helper address. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress.
    HelperAddress []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses) GetFilter() yfilter.YFilter { return helperAddresses.YFilter }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses) SetFilter(yf yfilter.YFilter) { helperAddresses.YFilter = yf }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses) GetGoName(yname string) string {
    if yname == "helper-address" { return "HelperAddress" }
    return ""
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses) GetSegmentPath() string {
    return "helper-addresses"
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "helper-address" {
        for _, c := range helperAddresses.HelperAddress {
            if helperAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress{}
        helperAddresses.HelperAddress = append(helperAddresses.HelperAddress, child)
        return &helperAddresses.HelperAddress[len(helperAddresses.HelperAddress)-1]
    }
    return nil
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range helperAddresses.HelperAddress {
        children[helperAddresses.HelperAddress[i].GetSegmentPath()] = &helperAddresses.HelperAddress[i]
    }
    return children
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses) GetYangName() string { return "helper-addresses" }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses) SetParent(parent types.Entity) { helperAddresses.parent = parent }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses) GetParent() types.Entity { return helperAddresses.parent }

func (helperAddresses *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses) GetParentYangName() string { return "vrf" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress
// Helper address
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    ServerAddress interface{}

    // Gateway address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    GatewayAddress interface{}
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetFilter() yfilter.YFilter { return helperAddress.YFilter }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) SetFilter(yf yfilter.YFilter) { helperAddress.YFilter = yf }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetGoName(yname string) string {
    if yname == "server-address" { return "ServerAddress" }
    if yname == "gateway-address" { return "GatewayAddress" }
    return ""
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetSegmentPath() string {
    return "helper-address" + "[server-address='" + fmt.Sprintf("%v", helperAddress.ServerAddress) + "']"
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["server-address"] = helperAddress.ServerAddress
    leafs["gateway-address"] = helperAddress.GatewayAddress
    return leafs
}

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetBundleName() string { return "cisco_ios_xr" }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetYangName() string { return "helper-address" }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) SetParent(parent types.Entity) { helperAddress.parent = parent }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetParent() types.Entity { return helperAddress.parent }

func (helperAddress *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetParentYangName() string { return "helper-addresses" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions
// Change sessions configuration
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Throttle DHCP sessions based on MAC address.
    ProxyThrottleType Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType
}

func (sessions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions) GetGoName(yname string) string {
    if yname == "proxy-throttle-type" { return "ProxyThrottleType" }
    return ""
}

func (sessions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "proxy-throttle-type" {
        return &sessions.ProxyThrottleType
    }
    return nil
}

func (sessions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["proxy-throttle-type"] = &sessions.ProxyThrottleType
    return children
}

func (sessions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions) GetBundleName() string { return "cisco_ios_xr" }

func (sessions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions) GetYangName() string { return "sessions" }

func (sessions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions) GetParentYangName() string { return "proxy" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType
// Throttle DHCP sessions based on MAC
// address
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Throttle DHCP sessions from any one MAC address.
    ProxyMacThrottle Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle
}

func (proxyThrottleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType) GetFilter() yfilter.YFilter { return proxyThrottleType.YFilter }

func (proxyThrottleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType) SetFilter(yf yfilter.YFilter) { proxyThrottleType.YFilter = yf }

func (proxyThrottleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType) GetGoName(yname string) string {
    if yname == "proxy-mac-throttle" { return "ProxyMacThrottle" }
    return ""
}

func (proxyThrottleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType) GetSegmentPath() string {
    return "proxy-throttle-type"
}

func (proxyThrottleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "proxy-mac-throttle" {
        return &proxyThrottleType.ProxyMacThrottle
    }
    return nil
}

func (proxyThrottleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["proxy-mac-throttle"] = &proxyThrottleType.ProxyMacThrottle
    return children
}

func (proxyThrottleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (proxyThrottleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType) GetBundleName() string { return "cisco_ios_xr" }

func (proxyThrottleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType) GetYangName() string { return "proxy-throttle-type" }

func (proxyThrottleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (proxyThrottleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (proxyThrottleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (proxyThrottleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType) SetParent(parent types.Entity) { proxyThrottleType.parent = parent }

func (proxyThrottleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType) GetParent() types.Entity { return proxyThrottleType.parent }

func (proxyThrottleType *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType) GetParentYangName() string { return "sessions" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle
// Throttle DHCP sessions from any one MAC
// address
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle struct {
    parent types.Entity
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

func (proxyMacThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle) GetFilter() yfilter.YFilter { return proxyMacThrottle.YFilter }

func (proxyMacThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle) SetFilter(yf yfilter.YFilter) { proxyMacThrottle.YFilter = yf }

func (proxyMacThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle) GetGoName(yname string) string {
    if yname == "num-discover" { return "NumDiscover" }
    if yname == "num-request" { return "NumRequest" }
    if yname == "num-block" { return "NumBlock" }
    return ""
}

func (proxyMacThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle) GetSegmentPath() string {
    return "proxy-mac-throttle"
}

func (proxyMacThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (proxyMacThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (proxyMacThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-discover"] = proxyMacThrottle.NumDiscover
    leafs["num-request"] = proxyMacThrottle.NumRequest
    leafs["num-block"] = proxyMacThrottle.NumBlock
    return leafs
}

func (proxyMacThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle) GetBundleName() string { return "cisco_ios_xr" }

func (proxyMacThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle) GetYangName() string { return "proxy-mac-throttle" }

func (proxyMacThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (proxyMacThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (proxyMacThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (proxyMacThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle) SetParent(parent types.Entity) { proxyMacThrottle.parent = parent }

func (proxyMacThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle) GetParent() types.Entity { return proxyMacThrottle.parent }

func (proxyMacThrottle *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Sessions_ProxyThrottleType_ProxyMacThrottle) GetParentYangName() string { return "proxy-throttle-type" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease
// Proxy limit lease
// This type is a presence type.
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Lease limit type. The type is Dhcpv4LimitLease1. This attribute is
    // mandatory.
    LimitType interface{}

    // Limit lease count. The type is interface{} with range: 1..240000. This
    // attribute is mandatory.
    LimitLeaseCount interface{}
}

func (limitLease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease) GetFilter() yfilter.YFilter { return limitLease.YFilter }

func (limitLease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease) SetFilter(yf yfilter.YFilter) { limitLease.YFilter = yf }

func (limitLease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease) GetGoName(yname string) string {
    if yname == "limit-type" { return "LimitType" }
    if yname == "limit-lease-count" { return "LimitLeaseCount" }
    return ""
}

func (limitLease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease) GetSegmentPath() string {
    return "limit-lease"
}

func (limitLease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (limitLease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (limitLease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit-type"] = limitLease.LimitType
    leafs["limit-lease-count"] = limitLease.LimitLeaseCount
    return leafs
}

func (limitLease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease) GetBundleName() string { return "cisco_ios_xr" }

func (limitLease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease) GetYangName() string { return "limit-lease" }

func (limitLease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (limitLease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (limitLease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (limitLease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease) SetParent(parent types.Entity) { limitLease.parent = parent }

func (limitLease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease) GetParent() types.Entity { return limitLease.parent }

func (limitLease *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LimitLease) GetParentYangName() string { return "proxy" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy
// DHCPv4 lease proxy
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify client lease proxy time. The type is interface{} with range:
    // 300..4294967295.
    ClientLeaseTime interface{}

    // Set DHCP server sent options in lease proxy generating ACK. The type is
    // interface{}.
    SetServerOptions interface{}
}

func (leaseProxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy) GetFilter() yfilter.YFilter { return leaseProxy.YFilter }

func (leaseProxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy) SetFilter(yf yfilter.YFilter) { leaseProxy.YFilter = yf }

func (leaseProxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy) GetGoName(yname string) string {
    if yname == "client-lease-time" { return "ClientLeaseTime" }
    if yname == "set-server-options" { return "SetServerOptions" }
    return ""
}

func (leaseProxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy) GetSegmentPath() string {
    return "lease-proxy"
}

func (leaseProxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (leaseProxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (leaseProxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["client-lease-time"] = leaseProxy.ClientLeaseTime
    leafs["set-server-options"] = leaseProxy.SetServerOptions
    return leafs
}

func (leaseProxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy) GetBundleName() string { return "cisco_ios_xr" }

func (leaseProxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy) GetYangName() string { return "lease-proxy" }

func (leaseProxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (leaseProxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (leaseProxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (leaseProxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy) SetParent(parent types.Entity) { leaseProxy.parent = parent }

func (leaseProxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy) GetParent() types.Entity { return leaseProxy.parent }

func (leaseProxy *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_LeaseProxy) GetParentYangName() string { return "proxy" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag
// Specify broadcast flag
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Broadcast flag policy. The type is Ipv4dhcpdBroadcastFlagPolicy.
    Policy interface{}
}

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag) GetFilter() yfilter.YFilter { return broadcastFlag.YFilter }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag) SetFilter(yf yfilter.YFilter) { broadcastFlag.YFilter = yf }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag) GetGoName(yname string) string {
    if yname == "policy" { return "Policy" }
    return ""
}

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag) GetSegmentPath() string {
    return "broadcast-flag"
}

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["policy"] = broadcastFlag.Policy
    return leafs
}

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag) GetBundleName() string { return "cisco_ios_xr" }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag) GetYangName() string { return "broadcast-flag" }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag) SetParent(parent types.Entity) { broadcastFlag.parent = parent }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag) GetParent() types.Entity { return broadcastFlag.parent }

func (broadcastFlag *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_BroadcastFlag) GetParentYangName() string { return "proxy" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match
// Insert match keyword
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of Option.
    DefOptions Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions

    // Table of Option.
    OptionFilters Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match) GetFilter() yfilter.YFilter { return match.YFilter }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match) SetFilter(yf yfilter.YFilter) { match.YFilter = yf }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match) GetGoName(yname string) string {
    if yname == "def-options" { return "DefOptions" }
    if yname == "option-filters" { return "OptionFilters" }
    return ""
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match) GetSegmentPath() string {
    return "match"
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "def-options" {
        return &match.DefOptions
    }
    if childYangName == "option-filters" {
        return &match.OptionFilters
    }
    return nil
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["def-options"] = &match.DefOptions
    children["option-filters"] = &match.OptionFilters
    return children
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match) GetBundleName() string { return "cisco_ios_xr" }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match) GetYangName() string { return "match" }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match) SetParent(parent types.Entity) { match.parent = parent }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match) GetParent() types.Entity { return match.parent }

func (match *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match) GetParentYangName() string { return "proxy" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions
// Table of Option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify match option. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption.
    DefOption []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption
}

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions) GetFilter() yfilter.YFilter { return defOptions.YFilter }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions) SetFilter(yf yfilter.YFilter) { defOptions.YFilter = yf }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions) GetGoName(yname string) string {
    if yname == "def-option" { return "DefOption" }
    return ""
}

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions) GetSegmentPath() string {
    return "def-options"
}

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "def-option" {
        for _, c := range defOptions.DefOption {
            if defOptions.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption{}
        defOptions.DefOption = append(defOptions.DefOption, child)
        return &defOptions.DefOption[len(defOptions.DefOption)-1]
    }
    return nil
}

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range defOptions.DefOption {
        children[defOptions.DefOption[i].GetSegmentPath()] = &defOptions.DefOption[i]
    }
    return children
}

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions) GetBundleName() string { return "cisco_ios_xr" }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions) GetYangName() string { return "def-options" }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions) SetParent(parent types.Entity) { defOptions.parent = parent }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions) GetParent() types.Entity { return defOptions.parent }

func (defOptions *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions) GetParentYangName() string { return "match" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption
// Specify match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Match option 60. The type is interface{} with
    // range: -2147483648..2147483647.
    DefMatchoption interface{}

    // Vendor action. The type is ProxyAction. This attribute is mandatory.
    DefMatchaction interface{}
}

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption) GetFilter() yfilter.YFilter { return defOption.YFilter }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption) SetFilter(yf yfilter.YFilter) { defOption.YFilter = yf }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption) GetGoName(yname string) string {
    if yname == "def-matchoption" { return "DefMatchoption" }
    if yname == "def-matchaction" { return "DefMatchaction" }
    return ""
}

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption) GetSegmentPath() string {
    return "def-option" + "[def-matchoption='" + fmt.Sprintf("%v", defOption.DefMatchoption) + "']"
}

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["def-matchoption"] = defOption.DefMatchoption
    leafs["def-matchaction"] = defOption.DefMatchaction
    return leafs
}

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption) GetBundleName() string { return "cisco_ios_xr" }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption) GetYangName() string { return "def-option" }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption) SetParent(parent types.Entity) { defOption.parent = parent }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption) GetParent() types.Entity { return defOption.parent }

func (defOption *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_DefOptions_DefOption) GetParentYangName() string { return "def-options" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters
// Table of Option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify match option. The type is slice of
    // Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter.
    OptionFilter []Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter
}

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters) GetFilter() yfilter.YFilter { return optionFilters.YFilter }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters) SetFilter(yf yfilter.YFilter) { optionFilters.YFilter = yf }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters) GetGoName(yname string) string {
    if yname == "option-filter" { return "OptionFilter" }
    return ""
}

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters) GetSegmentPath() string {
    return "option-filters"
}

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "option-filter" {
        for _, c := range optionFilters.OptionFilter {
            if optionFilters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter{}
        optionFilters.OptionFilter = append(optionFilters.OptionFilter, child)
        return &optionFilters.OptionFilter[len(optionFilters.OptionFilter)-1]
    }
    return nil
}

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range optionFilters.OptionFilter {
        children[optionFilters.OptionFilter[i].GetSegmentPath()] = &optionFilters.OptionFilter[i]
    }
    return children
}

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters) GetBundleName() string { return "cisco_ios_xr" }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters) GetYangName() string { return "option-filters" }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters) SetParent(parent types.Entity) { optionFilters.parent = parent }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters) GetParent() types.Entity { return optionFilters.parent }

func (optionFilters *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters) GetParentYangName() string { return "match" }

// Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter
// Specify match option
type Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Match option 60. The type is interface{} with
    // range: -2147483648..2147483647.
    Matchoption interface{}

    // This attribute is a key. Enter hex pattern string. The type is string with
    // length: 1..64.
    Pattern interface{}

    // This attribute is a key. Set constant integer. The type is interface{} with
    // range: -2147483648..2147483647.
    Format interface{}

    // Vendor action. The type is ProxyAction. This attribute is mandatory.
    Matchaction interface{}
}

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter) GetFilter() yfilter.YFilter { return optionFilter.YFilter }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter) SetFilter(yf yfilter.YFilter) { optionFilter.YFilter = yf }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter) GetGoName(yname string) string {
    if yname == "matchoption" { return "Matchoption" }
    if yname == "pattern" { return "Pattern" }
    if yname == "format" { return "Format" }
    if yname == "matchaction" { return "Matchaction" }
    return ""
}

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter) GetSegmentPath() string {
    return "option-filter" + "[matchoption='" + fmt.Sprintf("%v", optionFilter.Matchoption) + "']" + "[pattern='" + fmt.Sprintf("%v", optionFilter.Pattern) + "']" + "[format='" + fmt.Sprintf("%v", optionFilter.Format) + "']"
}

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["matchoption"] = optionFilter.Matchoption
    leafs["pattern"] = optionFilter.Pattern
    leafs["format"] = optionFilter.Format
    leafs["matchaction"] = optionFilter.Matchaction
    return leafs
}

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter) GetBundleName() string { return "cisco_ios_xr" }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter) GetYangName() string { return "option-filter" }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter) SetParent(parent types.Entity) { optionFilter.parent = parent }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter) GetParent() types.Entity { return optionFilter.parent }

func (optionFilter *Ipv4Dhcpd_Profiles_Profile_Modes_Mode_Proxy_Match_OptionFilters_OptionFilter) GetParentYangName() string { return "option-filters" }

// Ipv4Dhcpd_Database
// Enable DHCP binding database storage to file
// system
type Ipv4Dhcpd_Database struct {
    parent types.Entity
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

func (database *Ipv4Dhcpd_Database) GetFilter() yfilter.YFilter { return database.YFilter }

func (database *Ipv4Dhcpd_Database) SetFilter(yf yfilter.YFilter) { database.YFilter = yf }

func (database *Ipv4Dhcpd_Database) GetGoName(yname string) string {
    if yname == "proxy" { return "Proxy" }
    if yname == "server" { return "Server" }
    if yname == "snoop" { return "Snoop" }
    if yname == "full-write-interval" { return "FullWriteInterval" }
    if yname == "incremental-write-interval" { return "IncrementalWriteInterval" }
    return ""
}

func (database *Ipv4Dhcpd_Database) GetSegmentPath() string {
    return "database"
}

func (database *Ipv4Dhcpd_Database) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (database *Ipv4Dhcpd_Database) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (database *Ipv4Dhcpd_Database) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["proxy"] = database.Proxy
    leafs["server"] = database.Server
    leafs["snoop"] = database.Snoop
    leafs["full-write-interval"] = database.FullWriteInterval
    leafs["incremental-write-interval"] = database.IncrementalWriteInterval
    return leafs
}

func (database *Ipv4Dhcpd_Database) GetBundleName() string { return "cisco_ios_xr" }

func (database *Ipv4Dhcpd_Database) GetYangName() string { return "database" }

func (database *Ipv4Dhcpd_Database) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (database *Ipv4Dhcpd_Database) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (database *Ipv4Dhcpd_Database) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (database *Ipv4Dhcpd_Database) SetParent(parent types.Entity) { database.parent = parent }

func (database *Ipv4Dhcpd_Database) GetParent() types.Entity { return database.parent }

func (database *Ipv4Dhcpd_Database) GetParentYangName() string { return "ipv4-dhcpd" }

// Ipv4Dhcpd_Interfaces
// DHCP IPV4 Interface Table
type Ipv4Dhcpd_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP IPV4 Interface. The type is slice of Ipv4Dhcpd_Interfaces_Interface.
    Interface []Ipv4Dhcpd_Interfaces_Interface
}

func (interfaces *Ipv4Dhcpd_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Ipv4Dhcpd_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Ipv4Dhcpd_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Ipv4Dhcpd_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Ipv4Dhcpd_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Ipv4Dhcpd_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Ipv4Dhcpd_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Ipv4Dhcpd_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Ipv4Dhcpd_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Ipv4Dhcpd_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Ipv4Dhcpd_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Ipv4Dhcpd_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Ipv4Dhcpd_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Ipv4Dhcpd_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Ipv4Dhcpd_Interfaces) GetParentYangName() string { return "ipv4-dhcpd" }

// Ipv4Dhcpd_Interfaces_Interface
// DHCP IPV4 Interface
type Ipv4Dhcpd_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
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

func (self *Ipv4Dhcpd_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Ipv4Dhcpd_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Ipv4Dhcpd_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "proxy-interface" { return "ProxyInterface" }
    if yname == "base-interface" { return "BaseInterface" }
    if yname == "relay-interface" { return "RelayInterface" }
    if yname == "static-mode" { return "StaticMode" }
    if yname == "profile" { return "Profile" }
    if yname == "server-interface" { return "ServerInterface" }
    if yname == "snoop-interface" { return "SnoopInterface" }
    return ""
}

func (self *Ipv4Dhcpd_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Ipv4Dhcpd_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "proxy-interface" {
        return &self.ProxyInterface
    }
    if childYangName == "base-interface" {
        return &self.BaseInterface
    }
    if childYangName == "relay-interface" {
        return &self.RelayInterface
    }
    if childYangName == "static-mode" {
        return &self.StaticMode
    }
    if childYangName == "profile" {
        return &self.Profile
    }
    if childYangName == "server-interface" {
        return &self.ServerInterface
    }
    if childYangName == "snoop-interface" {
        return &self.SnoopInterface
    }
    return nil
}

func (self *Ipv4Dhcpd_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["proxy-interface"] = &self.ProxyInterface
    children["base-interface"] = &self.BaseInterface
    children["relay-interface"] = &self.RelayInterface
    children["static-mode"] = &self.StaticMode
    children["profile"] = &self.Profile
    children["server-interface"] = &self.ServerInterface
    children["snoop-interface"] = &self.SnoopInterface
    return children
}

func (self *Ipv4Dhcpd_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *Ipv4Dhcpd_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Ipv4Dhcpd_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Ipv4Dhcpd_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Ipv4Dhcpd_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Ipv4Dhcpd_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Ipv4Dhcpd_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Ipv4Dhcpd_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Ipv4Dhcpd_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Ipv4Dhcpd_Interfaces_Interface_ProxyInterface
// DHCP IPv4 proxy information
type Ipv4Dhcpd_Interfaces_Interface_ProxyInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface profile name. The type is string.
    Profile interface{}

    // Circuit ID value.
    DhcpCircuitId Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId
}

func (proxyInterface *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface) GetFilter() yfilter.YFilter { return proxyInterface.YFilter }

func (proxyInterface *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface) SetFilter(yf yfilter.YFilter) { proxyInterface.YFilter = yf }

func (proxyInterface *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    if yname == "dhcp-circuit-id" { return "DhcpCircuitId" }
    return ""
}

func (proxyInterface *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface) GetSegmentPath() string {
    return "proxy-interface"
}

func (proxyInterface *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dhcp-circuit-id" {
        return &proxyInterface.DhcpCircuitId
    }
    return nil
}

func (proxyInterface *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dhcp-circuit-id"] = &proxyInterface.DhcpCircuitId
    return children
}

func (proxyInterface *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile"] = proxyInterface.Profile
    return leafs
}

func (proxyInterface *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface) GetBundleName() string { return "cisco_ios_xr" }

func (proxyInterface *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface) GetYangName() string { return "proxy-interface" }

func (proxyInterface *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (proxyInterface *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (proxyInterface *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (proxyInterface *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface) SetParent(parent types.Entity) { proxyInterface.parent = parent }

func (proxyInterface *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface) GetParent() types.Entity { return proxyInterface.parent }

func (proxyInterface *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface) GetParentYangName() string { return "interface" }

// Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId
// Circuit ID value
// This type is a presence type.
type Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (dhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId) GetFilter() yfilter.YFilter { return dhcpCircuitId.YFilter }

func (dhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId) SetFilter(yf yfilter.YFilter) { dhcpCircuitId.YFilter = yf }

func (dhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId) GetGoName(yname string) string {
    if yname == "circuit-id" { return "CircuitId" }
    if yname == "format" { return "Format" }
    if yname == "argument1" { return "Argument1" }
    if yname == "argument2" { return "Argument2" }
    if yname == "argument3" { return "Argument3" }
    if yname == "argument4" { return "Argument4" }
    if yname == "argument5" { return "Argument5" }
    if yname == "argument6" { return "Argument6" }
    if yname == "argument7" { return "Argument7" }
    if yname == "argument8" { return "Argument8" }
    if yname == "argument9" { return "Argument9" }
    if yname == "argument10" { return "Argument10" }
    if yname == "argument11" { return "Argument11" }
    if yname == "argument12" { return "Argument12" }
    if yname == "argument13" { return "Argument13" }
    if yname == "argument14" { return "Argument14" }
    if yname == "argument15" { return "Argument15" }
    if yname == "argument16" { return "Argument16" }
    return ""
}

func (dhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId) GetSegmentPath() string {
    return "dhcp-circuit-id"
}

func (dhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["circuit-id"] = dhcpCircuitId.CircuitId
    leafs["format"] = dhcpCircuitId.Format
    leafs["argument1"] = dhcpCircuitId.Argument1
    leafs["argument2"] = dhcpCircuitId.Argument2
    leafs["argument3"] = dhcpCircuitId.Argument3
    leafs["argument4"] = dhcpCircuitId.Argument4
    leafs["argument5"] = dhcpCircuitId.Argument5
    leafs["argument6"] = dhcpCircuitId.Argument6
    leafs["argument7"] = dhcpCircuitId.Argument7
    leafs["argument8"] = dhcpCircuitId.Argument8
    leafs["argument9"] = dhcpCircuitId.Argument9
    leafs["argument10"] = dhcpCircuitId.Argument10
    leafs["argument11"] = dhcpCircuitId.Argument11
    leafs["argument12"] = dhcpCircuitId.Argument12
    leafs["argument13"] = dhcpCircuitId.Argument13
    leafs["argument14"] = dhcpCircuitId.Argument14
    leafs["argument15"] = dhcpCircuitId.Argument15
    leafs["argument16"] = dhcpCircuitId.Argument16
    return leafs
}

func (dhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId) GetBundleName() string { return "cisco_ios_xr" }

func (dhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId) GetYangName() string { return "dhcp-circuit-id" }

func (dhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId) SetParent(parent types.Entity) { dhcpCircuitId.parent = parent }

func (dhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId) GetParent() types.Entity { return dhcpCircuitId.parent }

func (dhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ProxyInterface_DhcpCircuitId) GetParentYangName() string { return "proxy-interface" }

// Ipv4Dhcpd_Interfaces_Interface_BaseInterface
// DHCP IPv4 Base profile information
type Ipv4Dhcpd_Interfaces_Interface_BaseInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface profile name. The type is string.
    Profile interface{}

    // Circuit ID value.
    BaseDhcpCircuitId Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId
}

func (baseInterface *Ipv4Dhcpd_Interfaces_Interface_BaseInterface) GetFilter() yfilter.YFilter { return baseInterface.YFilter }

func (baseInterface *Ipv4Dhcpd_Interfaces_Interface_BaseInterface) SetFilter(yf yfilter.YFilter) { baseInterface.YFilter = yf }

func (baseInterface *Ipv4Dhcpd_Interfaces_Interface_BaseInterface) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    if yname == "base-dhcp-circuit-id" { return "BaseDhcpCircuitId" }
    return ""
}

func (baseInterface *Ipv4Dhcpd_Interfaces_Interface_BaseInterface) GetSegmentPath() string {
    return "base-interface"
}

func (baseInterface *Ipv4Dhcpd_Interfaces_Interface_BaseInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "base-dhcp-circuit-id" {
        return &baseInterface.BaseDhcpCircuitId
    }
    return nil
}

func (baseInterface *Ipv4Dhcpd_Interfaces_Interface_BaseInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["base-dhcp-circuit-id"] = &baseInterface.BaseDhcpCircuitId
    return children
}

func (baseInterface *Ipv4Dhcpd_Interfaces_Interface_BaseInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile"] = baseInterface.Profile
    return leafs
}

func (baseInterface *Ipv4Dhcpd_Interfaces_Interface_BaseInterface) GetBundleName() string { return "cisco_ios_xr" }

func (baseInterface *Ipv4Dhcpd_Interfaces_Interface_BaseInterface) GetYangName() string { return "base-interface" }

func (baseInterface *Ipv4Dhcpd_Interfaces_Interface_BaseInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (baseInterface *Ipv4Dhcpd_Interfaces_Interface_BaseInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (baseInterface *Ipv4Dhcpd_Interfaces_Interface_BaseInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (baseInterface *Ipv4Dhcpd_Interfaces_Interface_BaseInterface) SetParent(parent types.Entity) { baseInterface.parent = parent }

func (baseInterface *Ipv4Dhcpd_Interfaces_Interface_BaseInterface) GetParent() types.Entity { return baseInterface.parent }

func (baseInterface *Ipv4Dhcpd_Interfaces_Interface_BaseInterface) GetParentYangName() string { return "interface" }

// Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId
// Circuit ID value
// This type is a presence type.
type Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (baseDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId) GetFilter() yfilter.YFilter { return baseDhcpCircuitId.YFilter }

func (baseDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId) SetFilter(yf yfilter.YFilter) { baseDhcpCircuitId.YFilter = yf }

func (baseDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId) GetGoName(yname string) string {
    if yname == "circuit-id" { return "CircuitId" }
    if yname == "format" { return "Format" }
    if yname == "argument1" { return "Argument1" }
    if yname == "argument2" { return "Argument2" }
    if yname == "argument3" { return "Argument3" }
    if yname == "argument4" { return "Argument4" }
    if yname == "argument5" { return "Argument5" }
    if yname == "argument6" { return "Argument6" }
    if yname == "argument7" { return "Argument7" }
    if yname == "argument8" { return "Argument8" }
    if yname == "argument9" { return "Argument9" }
    if yname == "argument10" { return "Argument10" }
    if yname == "argument11" { return "Argument11" }
    if yname == "argument12" { return "Argument12" }
    if yname == "argument13" { return "Argument13" }
    if yname == "argument14" { return "Argument14" }
    if yname == "argument15" { return "Argument15" }
    if yname == "argument16" { return "Argument16" }
    return ""
}

func (baseDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId) GetSegmentPath() string {
    return "base-dhcp-circuit-id"
}

func (baseDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (baseDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (baseDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["circuit-id"] = baseDhcpCircuitId.CircuitId
    leafs["format"] = baseDhcpCircuitId.Format
    leafs["argument1"] = baseDhcpCircuitId.Argument1
    leafs["argument2"] = baseDhcpCircuitId.Argument2
    leafs["argument3"] = baseDhcpCircuitId.Argument3
    leafs["argument4"] = baseDhcpCircuitId.Argument4
    leafs["argument5"] = baseDhcpCircuitId.Argument5
    leafs["argument6"] = baseDhcpCircuitId.Argument6
    leafs["argument7"] = baseDhcpCircuitId.Argument7
    leafs["argument8"] = baseDhcpCircuitId.Argument8
    leafs["argument9"] = baseDhcpCircuitId.Argument9
    leafs["argument10"] = baseDhcpCircuitId.Argument10
    leafs["argument11"] = baseDhcpCircuitId.Argument11
    leafs["argument12"] = baseDhcpCircuitId.Argument12
    leafs["argument13"] = baseDhcpCircuitId.Argument13
    leafs["argument14"] = baseDhcpCircuitId.Argument14
    leafs["argument15"] = baseDhcpCircuitId.Argument15
    leafs["argument16"] = baseDhcpCircuitId.Argument16
    return leafs
}

func (baseDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId) GetBundleName() string { return "cisco_ios_xr" }

func (baseDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId) GetYangName() string { return "base-dhcp-circuit-id" }

func (baseDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (baseDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (baseDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (baseDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId) SetParent(parent types.Entity) { baseDhcpCircuitId.parent = parent }

func (baseDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId) GetParent() types.Entity { return baseDhcpCircuitId.parent }

func (baseDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_BaseInterface_BaseDhcpCircuitId) GetParentYangName() string { return "base-interface" }

// Ipv4Dhcpd_Interfaces_Interface_RelayInterface
// DHCP IPv4 relay information
type Ipv4Dhcpd_Interfaces_Interface_RelayInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Circuit ID value.
    RelayDhcpCircuitId Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId
}

func (relayInterface *Ipv4Dhcpd_Interfaces_Interface_RelayInterface) GetFilter() yfilter.YFilter { return relayInterface.YFilter }

func (relayInterface *Ipv4Dhcpd_Interfaces_Interface_RelayInterface) SetFilter(yf yfilter.YFilter) { relayInterface.YFilter = yf }

func (relayInterface *Ipv4Dhcpd_Interfaces_Interface_RelayInterface) GetGoName(yname string) string {
    if yname == "relay-dhcp-circuit-id" { return "RelayDhcpCircuitId" }
    return ""
}

func (relayInterface *Ipv4Dhcpd_Interfaces_Interface_RelayInterface) GetSegmentPath() string {
    return "relay-interface"
}

func (relayInterface *Ipv4Dhcpd_Interfaces_Interface_RelayInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "relay-dhcp-circuit-id" {
        return &relayInterface.RelayDhcpCircuitId
    }
    return nil
}

func (relayInterface *Ipv4Dhcpd_Interfaces_Interface_RelayInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["relay-dhcp-circuit-id"] = &relayInterface.RelayDhcpCircuitId
    return children
}

func (relayInterface *Ipv4Dhcpd_Interfaces_Interface_RelayInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (relayInterface *Ipv4Dhcpd_Interfaces_Interface_RelayInterface) GetBundleName() string { return "cisco_ios_xr" }

func (relayInterface *Ipv4Dhcpd_Interfaces_Interface_RelayInterface) GetYangName() string { return "relay-interface" }

func (relayInterface *Ipv4Dhcpd_Interfaces_Interface_RelayInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relayInterface *Ipv4Dhcpd_Interfaces_Interface_RelayInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relayInterface *Ipv4Dhcpd_Interfaces_Interface_RelayInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relayInterface *Ipv4Dhcpd_Interfaces_Interface_RelayInterface) SetParent(parent types.Entity) { relayInterface.parent = parent }

func (relayInterface *Ipv4Dhcpd_Interfaces_Interface_RelayInterface) GetParent() types.Entity { return relayInterface.parent }

func (relayInterface *Ipv4Dhcpd_Interfaces_Interface_RelayInterface) GetParentYangName() string { return "interface" }

// Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId
// Circuit ID value
// This type is a presence type.
type Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (relayDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId) GetFilter() yfilter.YFilter { return relayDhcpCircuitId.YFilter }

func (relayDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId) SetFilter(yf yfilter.YFilter) { relayDhcpCircuitId.YFilter = yf }

func (relayDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId) GetGoName(yname string) string {
    if yname == "circuit-id" { return "CircuitId" }
    if yname == "format" { return "Format" }
    if yname == "argument1" { return "Argument1" }
    if yname == "argument2" { return "Argument2" }
    if yname == "argument3" { return "Argument3" }
    if yname == "argument4" { return "Argument4" }
    if yname == "argument5" { return "Argument5" }
    if yname == "argument6" { return "Argument6" }
    if yname == "argument7" { return "Argument7" }
    if yname == "argument8" { return "Argument8" }
    if yname == "argument9" { return "Argument9" }
    if yname == "argument10" { return "Argument10" }
    if yname == "argument11" { return "Argument11" }
    if yname == "argument12" { return "Argument12" }
    if yname == "argument13" { return "Argument13" }
    if yname == "argument14" { return "Argument14" }
    if yname == "argument15" { return "Argument15" }
    if yname == "argument16" { return "Argument16" }
    return ""
}

func (relayDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId) GetSegmentPath() string {
    return "relay-dhcp-circuit-id"
}

func (relayDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (relayDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (relayDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["circuit-id"] = relayDhcpCircuitId.CircuitId
    leafs["format"] = relayDhcpCircuitId.Format
    leafs["argument1"] = relayDhcpCircuitId.Argument1
    leafs["argument2"] = relayDhcpCircuitId.Argument2
    leafs["argument3"] = relayDhcpCircuitId.Argument3
    leafs["argument4"] = relayDhcpCircuitId.Argument4
    leafs["argument5"] = relayDhcpCircuitId.Argument5
    leafs["argument6"] = relayDhcpCircuitId.Argument6
    leafs["argument7"] = relayDhcpCircuitId.Argument7
    leafs["argument8"] = relayDhcpCircuitId.Argument8
    leafs["argument9"] = relayDhcpCircuitId.Argument9
    leafs["argument10"] = relayDhcpCircuitId.Argument10
    leafs["argument11"] = relayDhcpCircuitId.Argument11
    leafs["argument12"] = relayDhcpCircuitId.Argument12
    leafs["argument13"] = relayDhcpCircuitId.Argument13
    leafs["argument14"] = relayDhcpCircuitId.Argument14
    leafs["argument15"] = relayDhcpCircuitId.Argument15
    leafs["argument16"] = relayDhcpCircuitId.Argument16
    return leafs
}

func (relayDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId) GetBundleName() string { return "cisco_ios_xr" }

func (relayDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId) GetYangName() string { return "relay-dhcp-circuit-id" }

func (relayDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relayDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relayDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relayDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId) SetParent(parent types.Entity) { relayDhcpCircuitId.parent = parent }

func (relayDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId) GetParent() types.Entity { return relayDhcpCircuitId.parent }

func (relayDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_RelayInterface_RelayDhcpCircuitId) GetParentYangName() string { return "relay-interface" }

// Ipv4Dhcpd_Interfaces_Interface_StaticMode
// Static Table Entries containing MAC address to
// IP address bindings
type Ipv4Dhcpd_Interfaces_Interface_StaticMode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Static Table Entries containing MAC address to IP address bindings.
    Statics Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics
}

func (staticMode *Ipv4Dhcpd_Interfaces_Interface_StaticMode) GetFilter() yfilter.YFilter { return staticMode.YFilter }

func (staticMode *Ipv4Dhcpd_Interfaces_Interface_StaticMode) SetFilter(yf yfilter.YFilter) { staticMode.YFilter = yf }

func (staticMode *Ipv4Dhcpd_Interfaces_Interface_StaticMode) GetGoName(yname string) string {
    if yname == "statics" { return "Statics" }
    return ""
}

func (staticMode *Ipv4Dhcpd_Interfaces_Interface_StaticMode) GetSegmentPath() string {
    return "static-mode"
}

func (staticMode *Ipv4Dhcpd_Interfaces_Interface_StaticMode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "statics" {
        return &staticMode.Statics
    }
    return nil
}

func (staticMode *Ipv4Dhcpd_Interfaces_Interface_StaticMode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["statics"] = &staticMode.Statics
    return children
}

func (staticMode *Ipv4Dhcpd_Interfaces_Interface_StaticMode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (staticMode *Ipv4Dhcpd_Interfaces_Interface_StaticMode) GetBundleName() string { return "cisco_ios_xr" }

func (staticMode *Ipv4Dhcpd_Interfaces_Interface_StaticMode) GetYangName() string { return "static-mode" }

func (staticMode *Ipv4Dhcpd_Interfaces_Interface_StaticMode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticMode *Ipv4Dhcpd_Interfaces_Interface_StaticMode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticMode *Ipv4Dhcpd_Interfaces_Interface_StaticMode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticMode *Ipv4Dhcpd_Interfaces_Interface_StaticMode) SetParent(parent types.Entity) { staticMode.parent = parent }

func (staticMode *Ipv4Dhcpd_Interfaces_Interface_StaticMode) GetParent() types.Entity { return staticMode.parent }

func (staticMode *Ipv4Dhcpd_Interfaces_Interface_StaticMode) GetParentYangName() string { return "interface" }

// Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics
// Static Table Entries containing MAC address
// to IP address bindings
type Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCP static binding of Mac address to IP address. The type is slice of
    // Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static.
    Static []Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static
}

func (statics *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics) GetFilter() yfilter.YFilter { return statics.YFilter }

func (statics *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics) SetFilter(yf yfilter.YFilter) { statics.YFilter = yf }

func (statics *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics) GetGoName(yname string) string {
    if yname == "static" { return "Static" }
    return ""
}

func (statics *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics) GetSegmentPath() string {
    return "statics"
}

func (statics *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static" {
        for _, c := range statics.Static {
            if statics.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static{}
        statics.Static = append(statics.Static, child)
        return &statics.Static[len(statics.Static)-1]
    }
    return nil
}

func (statics *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range statics.Static {
        children[statics.Static[i].GetSegmentPath()] = &statics.Static[i]
    }
    return children
}

func (statics *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (statics *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics) GetBundleName() string { return "cisco_ios_xr" }

func (statics *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics) GetYangName() string { return "statics" }

func (statics *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (statics *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (statics *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (statics *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics) SetParent(parent types.Entity) { statics.parent = parent }

func (statics *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics) GetParent() types.Entity { return statics.parent }

func (statics *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics) GetParentYangName() string { return "static-mode" }

// Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static
// DHCP static binding of Mac address to IP
// address
type Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. MACAddress. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // This attribute is a key. Client Id. The type is interface{} with range:
    // 1..4294967295.
    ClientId interface{}

    // This attribute is a key. DHCP IPV4 Static layer. The type is
    // Ipv4dhcpdLayer.
    Layer interface{}

    // IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    // This attribute is mandatory.
    StaticAddress interface{}
}

func (static *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static) GetFilter() yfilter.YFilter { return static.YFilter }

func (static *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static) SetFilter(yf yfilter.YFilter) { static.YFilter = yf }

func (static *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "client-id" { return "ClientId" }
    if yname == "layer" { return "Layer" }
    if yname == "static-address" { return "StaticAddress" }
    return ""
}

func (static *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static) GetSegmentPath() string {
    return "static" + "[mac-address='" + fmt.Sprintf("%v", static.MacAddress) + "']" + "[client-id='" + fmt.Sprintf("%v", static.ClientId) + "']" + "[layer='" + fmt.Sprintf("%v", static.Layer) + "']"
}

func (static *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (static *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (static *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = static.MacAddress
    leafs["client-id"] = static.ClientId
    leafs["layer"] = static.Layer
    leafs["static-address"] = static.StaticAddress
    return leafs
}

func (static *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static) GetBundleName() string { return "cisco_ios_xr" }

func (static *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static) GetYangName() string { return "static" }

func (static *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (static *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (static *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (static *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static) SetParent(parent types.Entity) { static.parent = parent }

func (static *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static) GetParent() types.Entity { return static.parent }

func (static *Ipv4Dhcpd_Interfaces_Interface_StaticMode_Statics_Static) GetParentYangName() string { return "statics" }

// Ipv4Dhcpd_Interfaces_Interface_Profile
// Profile name and mode
// This type is a presence type.
type Ipv4Dhcpd_Interfaces_Interface_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Profile name. The type is string. This attribute is mandatory.
    ProfileName interface{}

    // DHCP mode. The type is Ipv4dhcpdMode. This attribute is mandatory.
    Mode interface{}
}

func (profile *Ipv4Dhcpd_Interfaces_Interface_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Ipv4Dhcpd_Interfaces_Interface_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Ipv4Dhcpd_Interfaces_Interface_Profile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "mode" { return "Mode" }
    return ""
}

func (profile *Ipv4Dhcpd_Interfaces_Interface_Profile) GetSegmentPath() string {
    return "profile"
}

func (profile *Ipv4Dhcpd_Interfaces_Interface_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (profile *Ipv4Dhcpd_Interfaces_Interface_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (profile *Ipv4Dhcpd_Interfaces_Interface_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = profile.ProfileName
    leafs["mode"] = profile.Mode
    return leafs
}

func (profile *Ipv4Dhcpd_Interfaces_Interface_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Ipv4Dhcpd_Interfaces_Interface_Profile) GetYangName() string { return "profile" }

func (profile *Ipv4Dhcpd_Interfaces_Interface_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Ipv4Dhcpd_Interfaces_Interface_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Ipv4Dhcpd_Interfaces_Interface_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Ipv4Dhcpd_Interfaces_Interface_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Ipv4Dhcpd_Interfaces_Interface_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Ipv4Dhcpd_Interfaces_Interface_Profile) GetParentYangName() string { return "interface" }

// Ipv4Dhcpd_Interfaces_Interface_ServerInterface
// DHCP IPv4 Server information
type Ipv4Dhcpd_Interfaces_Interface_ServerInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface profile name. The type is string.
    Profile interface{}

    // Circuit ID value.
    ServerDhcpCircuitId Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId
}

func (serverInterface *Ipv4Dhcpd_Interfaces_Interface_ServerInterface) GetFilter() yfilter.YFilter { return serverInterface.YFilter }

func (serverInterface *Ipv4Dhcpd_Interfaces_Interface_ServerInterface) SetFilter(yf yfilter.YFilter) { serverInterface.YFilter = yf }

func (serverInterface *Ipv4Dhcpd_Interfaces_Interface_ServerInterface) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    if yname == "server-dhcp-circuit-id" { return "ServerDhcpCircuitId" }
    return ""
}

func (serverInterface *Ipv4Dhcpd_Interfaces_Interface_ServerInterface) GetSegmentPath() string {
    return "server-interface"
}

func (serverInterface *Ipv4Dhcpd_Interfaces_Interface_ServerInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "server-dhcp-circuit-id" {
        return &serverInterface.ServerDhcpCircuitId
    }
    return nil
}

func (serverInterface *Ipv4Dhcpd_Interfaces_Interface_ServerInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["server-dhcp-circuit-id"] = &serverInterface.ServerDhcpCircuitId
    return children
}

func (serverInterface *Ipv4Dhcpd_Interfaces_Interface_ServerInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile"] = serverInterface.Profile
    return leafs
}

func (serverInterface *Ipv4Dhcpd_Interfaces_Interface_ServerInterface) GetBundleName() string { return "cisco_ios_xr" }

func (serverInterface *Ipv4Dhcpd_Interfaces_Interface_ServerInterface) GetYangName() string { return "server-interface" }

func (serverInterface *Ipv4Dhcpd_Interfaces_Interface_ServerInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serverInterface *Ipv4Dhcpd_Interfaces_Interface_ServerInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serverInterface *Ipv4Dhcpd_Interfaces_Interface_ServerInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serverInterface *Ipv4Dhcpd_Interfaces_Interface_ServerInterface) SetParent(parent types.Entity) { serverInterface.parent = parent }

func (serverInterface *Ipv4Dhcpd_Interfaces_Interface_ServerInterface) GetParent() types.Entity { return serverInterface.parent }

func (serverInterface *Ipv4Dhcpd_Interfaces_Interface_ServerInterface) GetParentYangName() string { return "interface" }

// Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId
// Circuit ID value
// This type is a presence type.
type Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (serverDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId) GetFilter() yfilter.YFilter { return serverDhcpCircuitId.YFilter }

func (serverDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId) SetFilter(yf yfilter.YFilter) { serverDhcpCircuitId.YFilter = yf }

func (serverDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId) GetGoName(yname string) string {
    if yname == "circuit-id" { return "CircuitId" }
    if yname == "format" { return "Format" }
    if yname == "argument1" { return "Argument1" }
    if yname == "argument2" { return "Argument2" }
    if yname == "argument3" { return "Argument3" }
    if yname == "argument4" { return "Argument4" }
    if yname == "argument5" { return "Argument5" }
    if yname == "argument6" { return "Argument6" }
    if yname == "argument7" { return "Argument7" }
    if yname == "argument8" { return "Argument8" }
    if yname == "argument9" { return "Argument9" }
    if yname == "argument10" { return "Argument10" }
    if yname == "argument11" { return "Argument11" }
    if yname == "argument12" { return "Argument12" }
    if yname == "argument13" { return "Argument13" }
    if yname == "argument14" { return "Argument14" }
    if yname == "argument15" { return "Argument15" }
    if yname == "argument16" { return "Argument16" }
    return ""
}

func (serverDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId) GetSegmentPath() string {
    return "server-dhcp-circuit-id"
}

func (serverDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (serverDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (serverDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["circuit-id"] = serverDhcpCircuitId.CircuitId
    leafs["format"] = serverDhcpCircuitId.Format
    leafs["argument1"] = serverDhcpCircuitId.Argument1
    leafs["argument2"] = serverDhcpCircuitId.Argument2
    leafs["argument3"] = serverDhcpCircuitId.Argument3
    leafs["argument4"] = serverDhcpCircuitId.Argument4
    leafs["argument5"] = serverDhcpCircuitId.Argument5
    leafs["argument6"] = serverDhcpCircuitId.Argument6
    leafs["argument7"] = serverDhcpCircuitId.Argument7
    leafs["argument8"] = serverDhcpCircuitId.Argument8
    leafs["argument9"] = serverDhcpCircuitId.Argument9
    leafs["argument10"] = serverDhcpCircuitId.Argument10
    leafs["argument11"] = serverDhcpCircuitId.Argument11
    leafs["argument12"] = serverDhcpCircuitId.Argument12
    leafs["argument13"] = serverDhcpCircuitId.Argument13
    leafs["argument14"] = serverDhcpCircuitId.Argument14
    leafs["argument15"] = serverDhcpCircuitId.Argument15
    leafs["argument16"] = serverDhcpCircuitId.Argument16
    return leafs
}

func (serverDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId) GetBundleName() string { return "cisco_ios_xr" }

func (serverDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId) GetYangName() string { return "server-dhcp-circuit-id" }

func (serverDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (serverDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (serverDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (serverDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId) SetParent(parent types.Entity) { serverDhcpCircuitId.parent = parent }

func (serverDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId) GetParent() types.Entity { return serverDhcpCircuitId.parent }

func (serverDhcpCircuitId *Ipv4Dhcpd_Interfaces_Interface_ServerInterface_ServerDhcpCircuitId) GetParentYangName() string { return "server-interface" }

// Ipv4Dhcpd_Interfaces_Interface_SnoopInterface
// DHCP IPv4 snoop information
type Ipv4Dhcpd_Interfaces_Interface_SnoopInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure circuit ID for snoop 1. Hex 2. ASCII.
    SnoopCircuitId Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId
}

func (snoopInterface *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface) GetFilter() yfilter.YFilter { return snoopInterface.YFilter }

func (snoopInterface *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface) SetFilter(yf yfilter.YFilter) { snoopInterface.YFilter = yf }

func (snoopInterface *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface) GetGoName(yname string) string {
    if yname == "snoop-circuit-id" { return "SnoopCircuitId" }
    return ""
}

func (snoopInterface *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface) GetSegmentPath() string {
    return "snoop-interface"
}

func (snoopInterface *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "snoop-circuit-id" {
        return &snoopInterface.SnoopCircuitId
    }
    return nil
}

func (snoopInterface *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["snoop-circuit-id"] = &snoopInterface.SnoopCircuitId
    return children
}

func (snoopInterface *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (snoopInterface *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface) GetBundleName() string { return "cisco_ios_xr" }

func (snoopInterface *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface) GetYangName() string { return "snoop-interface" }

func (snoopInterface *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (snoopInterface *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (snoopInterface *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (snoopInterface *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface) SetParent(parent types.Entity) { snoopInterface.parent = parent }

func (snoopInterface *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface) GetParent() types.Entity { return snoopInterface.parent }

func (snoopInterface *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface) GetParentYangName() string { return "interface" }

// Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId
// Configure circuit ID for snoop 1. Hex 2.
// ASCII
type Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Format type, 1. Hex 2. ASCII. The type is interface{} with range: 1..2.
    FormatType interface{}

    // Enter circuit-id value. The type is string.
    CircuitIdValue interface{}
}

func (snoopCircuitId *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId) GetFilter() yfilter.YFilter { return snoopCircuitId.YFilter }

func (snoopCircuitId *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId) SetFilter(yf yfilter.YFilter) { snoopCircuitId.YFilter = yf }

func (snoopCircuitId *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId) GetGoName(yname string) string {
    if yname == "format-type" { return "FormatType" }
    if yname == "circuit-id-value" { return "CircuitIdValue" }
    return ""
}

func (snoopCircuitId *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId) GetSegmentPath() string {
    return "snoop-circuit-id"
}

func (snoopCircuitId *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (snoopCircuitId *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (snoopCircuitId *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["format-type"] = snoopCircuitId.FormatType
    leafs["circuit-id-value"] = snoopCircuitId.CircuitIdValue
    return leafs
}

func (snoopCircuitId *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId) GetBundleName() string { return "cisco_ios_xr" }

func (snoopCircuitId *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId) GetYangName() string { return "snoop-circuit-id" }

func (snoopCircuitId *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (snoopCircuitId *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (snoopCircuitId *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (snoopCircuitId *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId) SetParent(parent types.Entity) { snoopCircuitId.parent = parent }

func (snoopCircuitId *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId) GetParent() types.Entity { return snoopCircuitId.parent }

func (snoopCircuitId *Ipv4Dhcpd_Interfaces_Interface_SnoopInterface_SnoopCircuitId) GetParentYangName() string { return "snoop-interface" }

// Ipv4Dhcpd_DuplicateMacAllowed
// Allow Duplicate MAC Address
// This type is a presence type.
type Ipv4Dhcpd_DuplicateMacAllowed struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Duplicate mac is allowed. The type is interface{}. This attribute is
    // mandatory.
    DuplicateMac interface{}

    // Exclude vlan. The type is interface{}.
    ExcludeVlan interface{}

    // Include giaddr. The type is interface{}.
    IncludeGiaddr interface{}
}

func (duplicateMacAllowed *Ipv4Dhcpd_DuplicateMacAllowed) GetFilter() yfilter.YFilter { return duplicateMacAllowed.YFilter }

func (duplicateMacAllowed *Ipv4Dhcpd_DuplicateMacAllowed) SetFilter(yf yfilter.YFilter) { duplicateMacAllowed.YFilter = yf }

func (duplicateMacAllowed *Ipv4Dhcpd_DuplicateMacAllowed) GetGoName(yname string) string {
    if yname == "duplicate-mac" { return "DuplicateMac" }
    if yname == "exclude-vlan" { return "ExcludeVlan" }
    if yname == "include-giaddr" { return "IncludeGiaddr" }
    return ""
}

func (duplicateMacAllowed *Ipv4Dhcpd_DuplicateMacAllowed) GetSegmentPath() string {
    return "duplicate-mac-allowed"
}

func (duplicateMacAllowed *Ipv4Dhcpd_DuplicateMacAllowed) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (duplicateMacAllowed *Ipv4Dhcpd_DuplicateMacAllowed) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (duplicateMacAllowed *Ipv4Dhcpd_DuplicateMacAllowed) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["duplicate-mac"] = duplicateMacAllowed.DuplicateMac
    leafs["exclude-vlan"] = duplicateMacAllowed.ExcludeVlan
    leafs["include-giaddr"] = duplicateMacAllowed.IncludeGiaddr
    return leafs
}

func (duplicateMacAllowed *Ipv4Dhcpd_DuplicateMacAllowed) GetBundleName() string { return "cisco_ios_xr" }

func (duplicateMacAllowed *Ipv4Dhcpd_DuplicateMacAllowed) GetYangName() string { return "duplicate-mac-allowed" }

func (duplicateMacAllowed *Ipv4Dhcpd_DuplicateMacAllowed) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (duplicateMacAllowed *Ipv4Dhcpd_DuplicateMacAllowed) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (duplicateMacAllowed *Ipv4Dhcpd_DuplicateMacAllowed) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (duplicateMacAllowed *Ipv4Dhcpd_DuplicateMacAllowed) SetParent(parent types.Entity) { duplicateMacAllowed.parent = parent }

func (duplicateMacAllowed *Ipv4Dhcpd_DuplicateMacAllowed) GetParent() types.Entity { return duplicateMacAllowed.parent }

func (duplicateMacAllowed *Ipv4Dhcpd_DuplicateMacAllowed) GetParentYangName() string { return "ipv4-dhcpd" }

// Ipv4Dhcpd_RateLimit
// Rate limit ingress packets
type Ipv4Dhcpd_RateLimit struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Rate limiter period in msec (default: 200 msec). The type is interface{}
    // with range: 1..1000. The default value is 200.
    NumPeriod interface{}

    // Max DISCOVER packets per rate-limiter period (default 100). The type is
    // interface{} with range: 0..1000. The default value is 100.
    NumDiscover interface{}
}

func (rateLimit *Ipv4Dhcpd_RateLimit) GetFilter() yfilter.YFilter { return rateLimit.YFilter }

func (rateLimit *Ipv4Dhcpd_RateLimit) SetFilter(yf yfilter.YFilter) { rateLimit.YFilter = yf }

func (rateLimit *Ipv4Dhcpd_RateLimit) GetGoName(yname string) string {
    if yname == "num-period" { return "NumPeriod" }
    if yname == "num-discover" { return "NumDiscover" }
    return ""
}

func (rateLimit *Ipv4Dhcpd_RateLimit) GetSegmentPath() string {
    return "rate-limit"
}

func (rateLimit *Ipv4Dhcpd_RateLimit) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rateLimit *Ipv4Dhcpd_RateLimit) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rateLimit *Ipv4Dhcpd_RateLimit) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["num-period"] = rateLimit.NumPeriod
    leafs["num-discover"] = rateLimit.NumDiscover
    return leafs
}

func (rateLimit *Ipv4Dhcpd_RateLimit) GetBundleName() string { return "cisco_ios_xr" }

func (rateLimit *Ipv4Dhcpd_RateLimit) GetYangName() string { return "rate-limit" }

func (rateLimit *Ipv4Dhcpd_RateLimit) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rateLimit *Ipv4Dhcpd_RateLimit) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rateLimit *Ipv4Dhcpd_RateLimit) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rateLimit *Ipv4Dhcpd_RateLimit) SetParent(parent types.Entity) { rateLimit.parent = parent }

func (rateLimit *Ipv4Dhcpd_RateLimit) GetParent() types.Entity { return rateLimit.parent }

func (rateLimit *Ipv4Dhcpd_RateLimit) GetParentYangName() string { return "ipv4-dhcpd" }

