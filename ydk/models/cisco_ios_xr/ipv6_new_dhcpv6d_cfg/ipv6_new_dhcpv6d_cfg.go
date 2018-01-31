// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-new-dhcpv6d package configuration.
// 
// This module contains definitions
// for the following management objects:
//   dhcpv6: None
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_new_dhcpv6d_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_new_dhcpv6d_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-new-dhcpv6d-cfg dhcpv6}", reflect.TypeOf(Dhcpv6{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-new-dhcpv6d-cfg:dhcpv6", reflect.TypeOf(Dhcpv6{}))
}

// Insert represents Insert
type Insert string

const (
    // Insert locally generated/configured Interface
    // ID value
    Insert_local Insert = "local"

    // Insert received Interface ID value
    Insert_received Insert = "received"

    // Insert received Interface ID value from SADB
    Insert_pppoe Insert = "pppoe"
)

// LinkLayerAddr represents Link layer addr
type LinkLayerAddr string

const (
    // Insert Received LinkLayerAddr Value from SADB
    LinkLayerAddr_set LinkLayerAddr = "set"
)

// SubscriberId represents Subscriber id
type SubscriberId string

const (
    // Insert Received Subscriber-ID Value from SADB
    SubscriberId_pppoe SubscriberId = "pppoe"
)

// Dhcpv6
// None
// This type is a presence type.
type Dhcpv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable None. Deletion of this object also causes deletion of all associated
    // objects under DHCPv6. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // For BNG session, allow duid change for a client MAC. The type is
    // interface{}.
    AllowDuidChange interface{}

    // Enable DHCP binding database storage to file system.
    Database Dhcpv6_Database

    // Table of Profile.
    Profiles Dhcpv6_Profiles

    // Table of Interface.
    Interfaces Dhcpv6_Interfaces
}

func (dhcpv6 *Dhcpv6) GetFilter() yfilter.YFilter { return dhcpv6.YFilter }

func (dhcpv6 *Dhcpv6) SetFilter(yf yfilter.YFilter) { dhcpv6.YFilter = yf }

func (dhcpv6 *Dhcpv6) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "allow-duid-change" { return "AllowDuidChange" }
    if yname == "database" { return "Database" }
    if yname == "profiles" { return "Profiles" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (dhcpv6 *Dhcpv6) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-new-dhcpv6d-cfg:dhcpv6"
}

func (dhcpv6 *Dhcpv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "database" {
        return &dhcpv6.Database
    }
    if childYangName == "profiles" {
        return &dhcpv6.Profiles
    }
    if childYangName == "interfaces" {
        return &dhcpv6.Interfaces
    }
    return nil
}

func (dhcpv6 *Dhcpv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["database"] = &dhcpv6.Database
    children["profiles"] = &dhcpv6.Profiles
    children["interfaces"] = &dhcpv6.Interfaces
    return children
}

func (dhcpv6 *Dhcpv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = dhcpv6.Enable
    leafs["allow-duid-change"] = dhcpv6.AllowDuidChange
    return leafs
}

func (dhcpv6 *Dhcpv6) GetBundleName() string { return "cisco_ios_xr" }

func (dhcpv6 *Dhcpv6) GetYangName() string { return "dhcpv6" }

func (dhcpv6 *Dhcpv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcpv6 *Dhcpv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcpv6 *Dhcpv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcpv6 *Dhcpv6) SetParent(parent types.Entity) { dhcpv6.parent = parent }

func (dhcpv6 *Dhcpv6) GetParent() types.Entity { return dhcpv6.parent }

func (dhcpv6 *Dhcpv6) GetParentYangName() string { return "Cisco-IOS-XR-ipv6-new-dhcpv6d-cfg" }

// Dhcpv6_Database
// Enable DHCP binding database storage to file
// system
type Dhcpv6_Database struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable DHCP proxy binding database storage to file system. The type is
    // interface{}.
    Proxy interface{}

    // Enable DHCP server binding database storage to file system. The type is
    // interface{}.
    Server interface{}

    // Enable DHCP relay binding database storage to file system. The type is
    // interface{}.
    Relay interface{}

    // Full file write interval (default 10 minutes). The type is interface{} with
    // range: 1..1440. The default value is 10.
    FullWriteInterval interface{}

    // Incremental file write interval (default 1 minutes). The type is
    // interface{} with range: 1..1440. The default value is 1.
    IncrementalWriteInterval interface{}
}

func (database *Dhcpv6_Database) GetFilter() yfilter.YFilter { return database.YFilter }

func (database *Dhcpv6_Database) SetFilter(yf yfilter.YFilter) { database.YFilter = yf }

func (database *Dhcpv6_Database) GetGoName(yname string) string {
    if yname == "proxy" { return "Proxy" }
    if yname == "server" { return "Server" }
    if yname == "relay" { return "Relay" }
    if yname == "full-write-interval" { return "FullWriteInterval" }
    if yname == "incremental-write-interval" { return "IncrementalWriteInterval" }
    return ""
}

func (database *Dhcpv6_Database) GetSegmentPath() string {
    return "database"
}

func (database *Dhcpv6_Database) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (database *Dhcpv6_Database) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (database *Dhcpv6_Database) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["proxy"] = database.Proxy
    leafs["server"] = database.Server
    leafs["relay"] = database.Relay
    leafs["full-write-interval"] = database.FullWriteInterval
    leafs["incremental-write-interval"] = database.IncrementalWriteInterval
    return leafs
}

func (database *Dhcpv6_Database) GetBundleName() string { return "cisco_ios_xr" }

func (database *Dhcpv6_Database) GetYangName() string { return "database" }

func (database *Dhcpv6_Database) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (database *Dhcpv6_Database) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (database *Dhcpv6_Database) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (database *Dhcpv6_Database) SetParent(parent types.Entity) { database.parent = parent }

func (database *Dhcpv6_Database) GetParent() types.Entity { return database.parent }

func (database *Dhcpv6_Database) GetParentYangName() string { return "dhcpv6" }

// Dhcpv6_Profiles
// Table of Profile
type Dhcpv6_Profiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None. The type is slice of Dhcpv6_Profiles_Profile.
    Profile []Dhcpv6_Profiles_Profile
}

func (profiles *Dhcpv6_Profiles) GetFilter() yfilter.YFilter { return profiles.YFilter }

func (profiles *Dhcpv6_Profiles) SetFilter(yf yfilter.YFilter) { profiles.YFilter = yf }

func (profiles *Dhcpv6_Profiles) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (profiles *Dhcpv6_Profiles) GetSegmentPath() string {
    return "profiles"
}

func (profiles *Dhcpv6_Profiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        for _, c := range profiles.Profile {
            if profiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Profiles_Profile{}
        profiles.Profile = append(profiles.Profile, child)
        return &profiles.Profile[len(profiles.Profile)-1]
    }
    return nil
}

func (profiles *Dhcpv6_Profiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range profiles.Profile {
        children[profiles.Profile[i].GetSegmentPath()] = &profiles.Profile[i]
    }
    return children
}

func (profiles *Dhcpv6_Profiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (profiles *Dhcpv6_Profiles) GetBundleName() string { return "cisco_ios_xr" }

func (profiles *Dhcpv6_Profiles) GetYangName() string { return "profiles" }

func (profiles *Dhcpv6_Profiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profiles *Dhcpv6_Profiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profiles *Dhcpv6_Profiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profiles *Dhcpv6_Profiles) SetParent(parent types.Entity) { profiles.parent = parent }

func (profiles *Dhcpv6_Profiles) GetParent() types.Entity { return profiles.parent }

func (profiles *Dhcpv6_Profiles) GetParentYangName() string { return "dhcpv6" }

// Dhcpv6_Profiles_Profile
// None
type Dhcpv6_Profiles_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    ProfileName interface{}

    // None.
    Relay Dhcpv6_Profiles_Profile_Relay

    // None.
    Base Dhcpv6_Profiles_Profile_Base

    // None.
    Proxy Dhcpv6_Profiles_Profile_Proxy

    // None.
    Server Dhcpv6_Profiles_Profile_Server
}

func (profile *Dhcpv6_Profiles_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Dhcpv6_Profiles_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Dhcpv6_Profiles_Profile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "relay" { return "Relay" }
    if yname == "base" { return "Base" }
    if yname == "proxy" { return "Proxy" }
    if yname == "server" { return "Server" }
    return ""
}

func (profile *Dhcpv6_Profiles_Profile) GetSegmentPath() string {
    return "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
}

func (profile *Dhcpv6_Profiles_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "relay" {
        return &profile.Relay
    }
    if childYangName == "base" {
        return &profile.Base
    }
    if childYangName == "proxy" {
        return &profile.Proxy
    }
    if childYangName == "server" {
        return &profile.Server
    }
    return nil
}

func (profile *Dhcpv6_Profiles_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["relay"] = &profile.Relay
    children["base"] = &profile.Base
    children["proxy"] = &profile.Proxy
    children["server"] = &profile.Server
    return children
}

func (profile *Dhcpv6_Profiles_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = profile.ProfileName
    return leafs
}

func (profile *Dhcpv6_Profiles_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Dhcpv6_Profiles_Profile) GetYangName() string { return "profile" }

func (profile *Dhcpv6_Profiles_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Dhcpv6_Profiles_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Dhcpv6_Profiles_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Dhcpv6_Profiles_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Dhcpv6_Profiles_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Dhcpv6_Profiles_Profile) GetParentYangName() string { return "profiles" }

// Dhcpv6_Profiles_Profile_Relay
// None
// This type is a presence type.
type Dhcpv6_Profiles_Profile_Relay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable None. Deletion of this object also causes deletion of all associated
    // objects under Relay. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Enable route addition for IANA. The type is interface{}.
    IanaRouteAdd interface{}

    // Table of HelperAddress.
    HelperAddresses Dhcpv6_Profiles_Profile_Relay_HelperAddresses
}

func (relay *Dhcpv6_Profiles_Profile_Relay) GetFilter() yfilter.YFilter { return relay.YFilter }

func (relay *Dhcpv6_Profiles_Profile_Relay) SetFilter(yf yfilter.YFilter) { relay.YFilter = yf }

func (relay *Dhcpv6_Profiles_Profile_Relay) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "iana-route-add" { return "IanaRouteAdd" }
    if yname == "helper-addresses" { return "HelperAddresses" }
    return ""
}

func (relay *Dhcpv6_Profiles_Profile_Relay) GetSegmentPath() string {
    return "relay"
}

func (relay *Dhcpv6_Profiles_Profile_Relay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "helper-addresses" {
        return &relay.HelperAddresses
    }
    return nil
}

func (relay *Dhcpv6_Profiles_Profile_Relay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["helper-addresses"] = &relay.HelperAddresses
    return children
}

func (relay *Dhcpv6_Profiles_Profile_Relay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = relay.Enable
    leafs["iana-route-add"] = relay.IanaRouteAdd
    return leafs
}

func (relay *Dhcpv6_Profiles_Profile_Relay) GetBundleName() string { return "cisco_ios_xr" }

func (relay *Dhcpv6_Profiles_Profile_Relay) GetYangName() string { return "relay" }

func (relay *Dhcpv6_Profiles_Profile_Relay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relay *Dhcpv6_Profiles_Profile_Relay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relay *Dhcpv6_Profiles_Profile_Relay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relay *Dhcpv6_Profiles_Profile_Relay) SetParent(parent types.Entity) { relay.parent = parent }

func (relay *Dhcpv6_Profiles_Profile_Relay) GetParent() types.Entity { return relay.parent }

func (relay *Dhcpv6_Profiles_Profile_Relay) GetParentYangName() string { return "profile" }

// Dhcpv6_Profiles_Profile_Relay_HelperAddresses
// Table of HelperAddress
type Dhcpv6_Profiles_Profile_Relay_HelperAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the server helper address. The type is slice of
    // Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress.
    HelperAddress []Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Relay_HelperAddresses) GetFilter() yfilter.YFilter { return helperAddresses.YFilter }

func (helperAddresses *Dhcpv6_Profiles_Profile_Relay_HelperAddresses) SetFilter(yf yfilter.YFilter) { helperAddresses.YFilter = yf }

func (helperAddresses *Dhcpv6_Profiles_Profile_Relay_HelperAddresses) GetGoName(yname string) string {
    if yname == "helper-address" { return "HelperAddress" }
    return ""
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Relay_HelperAddresses) GetSegmentPath() string {
    return "helper-addresses"
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Relay_HelperAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "helper-address" {
        for _, c := range helperAddresses.HelperAddress {
            if helperAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress{}
        helperAddresses.HelperAddress = append(helperAddresses.HelperAddress, child)
        return &helperAddresses.HelperAddress[len(helperAddresses.HelperAddress)-1]
    }
    return nil
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Relay_HelperAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range helperAddresses.HelperAddress {
        children[helperAddresses.HelperAddress[i].GetSegmentPath()] = &helperAddresses.HelperAddress[i]
    }
    return children
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Relay_HelperAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Relay_HelperAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (helperAddresses *Dhcpv6_Profiles_Profile_Relay_HelperAddresses) GetYangName() string { return "helper-addresses" }

func (helperAddresses *Dhcpv6_Profiles_Profile_Relay_HelperAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helperAddresses *Dhcpv6_Profiles_Profile_Relay_HelperAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helperAddresses *Dhcpv6_Profiles_Profile_Relay_HelperAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helperAddresses *Dhcpv6_Profiles_Profile_Relay_HelperAddresses) SetParent(parent types.Entity) { helperAddresses.parent = parent }

func (helperAddresses *Dhcpv6_Profiles_Profile_Relay_HelperAddresses) GetParent() types.Entity { return helperAddresses.parent }

func (helperAddresses *Dhcpv6_Profiles_Profile_Relay_HelperAddresses) GetParentYangName() string { return "relay" }

// Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress
// Specify the server helper address
type Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // This attribute is a key. Server Global unicast address. The type is string
    // with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    HelperAddress interface{}
}

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) GetFilter() yfilter.YFilter { return helperAddress.YFilter }

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) SetFilter(yf yfilter.YFilter) { helperAddress.YFilter = yf }

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "helper-address" { return "HelperAddress" }
    return ""
}

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) GetSegmentPath() string {
    return "helper-address" + "[vrf-name='" + fmt.Sprintf("%v", helperAddress.VrfName) + "']" + "[helper-address='" + fmt.Sprintf("%v", helperAddress.HelperAddress) + "']"
}

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = helperAddress.VrfName
    leafs["helper-address"] = helperAddress.HelperAddress
    return leafs
}

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) GetBundleName() string { return "cisco_ios_xr" }

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) GetYangName() string { return "helper-address" }

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) SetParent(parent types.Entity) { helperAddress.parent = parent }

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) GetParent() types.Entity { return helperAddress.parent }

func (helperAddress *Dhcpv6_Profiles_Profile_Relay_HelperAddresses_HelperAddress) GetParentYangName() string { return "helper-addresses" }

// Dhcpv6_Profiles_Profile_Base
// None
// This type is a presence type.
type Dhcpv6_Profiles_Profile_Base struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable None. Deletion of this object also causes deletion of all associated
    // objects under Base. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Default match option.
    Default Dhcpv6_Profiles_Profile_Base_Default

    // Enter match option.
    Match Dhcpv6_Profiles_Profile_Base_Match
}

func (base *Dhcpv6_Profiles_Profile_Base) GetFilter() yfilter.YFilter { return base.YFilter }

func (base *Dhcpv6_Profiles_Profile_Base) SetFilter(yf yfilter.YFilter) { base.YFilter = yf }

func (base *Dhcpv6_Profiles_Profile_Base) GetGoName(yname string) string {
    if yname == "enable" { return "Enable" }
    if yname == "default" { return "Default" }
    if yname == "match" { return "Match" }
    return ""
}

func (base *Dhcpv6_Profiles_Profile_Base) GetSegmentPath() string {
    return "base"
}

func (base *Dhcpv6_Profiles_Profile_Base) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default" {
        return &base.Default
    }
    if childYangName == "match" {
        return &base.Match
    }
    return nil
}

func (base *Dhcpv6_Profiles_Profile_Base) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default"] = &base.Default
    children["match"] = &base.Match
    return children
}

func (base *Dhcpv6_Profiles_Profile_Base) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable"] = base.Enable
    return leafs
}

func (base *Dhcpv6_Profiles_Profile_Base) GetBundleName() string { return "cisco_ios_xr" }

func (base *Dhcpv6_Profiles_Profile_Base) GetYangName() string { return "base" }

func (base *Dhcpv6_Profiles_Profile_Base) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (base *Dhcpv6_Profiles_Profile_Base) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (base *Dhcpv6_Profiles_Profile_Base) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (base *Dhcpv6_Profiles_Profile_Base) SetParent(parent types.Entity) { base.parent = parent }

func (base *Dhcpv6_Profiles_Profile_Base) GetParent() types.Entity { return base.parent }

func (base *Dhcpv6_Profiles_Profile_Base) GetParentYangName() string { return "profile" }

// Dhcpv6_Profiles_Profile_Base_Default
// Default match option
type Dhcpv6_Profiles_Profile_Base_Default struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enter proxy or server profile. The type is slice of
    // Dhcpv6_Profiles_Profile_Base_Default_Profile.
    Profile []Dhcpv6_Profiles_Profile_Base_Default_Profile
}

func (self *Dhcpv6_Profiles_Profile_Base_Default) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Dhcpv6_Profiles_Profile_Base_Default) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Dhcpv6_Profiles_Profile_Base_Default) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (self *Dhcpv6_Profiles_Profile_Base_Default) GetSegmentPath() string {
    return "default"
}

func (self *Dhcpv6_Profiles_Profile_Base_Default) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        for _, c := range self.Profile {
            if self.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Profiles_Profile_Base_Default_Profile{}
        self.Profile = append(self.Profile, child)
        return &self.Profile[len(self.Profile)-1]
    }
    return nil
}

func (self *Dhcpv6_Profiles_Profile_Base_Default) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range self.Profile {
        children[self.Profile[i].GetSegmentPath()] = &self.Profile[i]
    }
    return children
}

func (self *Dhcpv6_Profiles_Profile_Base_Default) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (self *Dhcpv6_Profiles_Profile_Base_Default) GetBundleName() string { return "cisco_ios_xr" }

func (self *Dhcpv6_Profiles_Profile_Base_Default) GetYangName() string { return "default" }

func (self *Dhcpv6_Profiles_Profile_Base_Default) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Dhcpv6_Profiles_Profile_Base_Default) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Dhcpv6_Profiles_Profile_Base_Default) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Dhcpv6_Profiles_Profile_Base_Default) SetParent(parent types.Entity) { self.parent = parent }

func (self *Dhcpv6_Profiles_Profile_Base_Default) GetParent() types.Entity { return self.parent }

func (self *Dhcpv6_Profiles_Profile_Base_Default) GetParentYangName() string { return "base" }

// Dhcpv6_Profiles_Profile_Base_Default_Profile
// Enter proxy or server profile
type Dhcpv6_Profiles_Profile_Base_Default_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with length:
    // 1..64.
    ProfileName interface{}

    // Specify mode-class based Server option. The type is interface{}.
    ServerMode interface{}

    // Specify mode-class based Proxy Option. The type is interface{}.
    ProxyMode interface{}
}

func (profile *Dhcpv6_Profiles_Profile_Base_Default_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Dhcpv6_Profiles_Profile_Base_Default_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Dhcpv6_Profiles_Profile_Base_Default_Profile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "server-mode" { return "ServerMode" }
    if yname == "proxy-mode" { return "ProxyMode" }
    return ""
}

func (profile *Dhcpv6_Profiles_Profile_Base_Default_Profile) GetSegmentPath() string {
    return "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
}

func (profile *Dhcpv6_Profiles_Profile_Base_Default_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (profile *Dhcpv6_Profiles_Profile_Base_Default_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (profile *Dhcpv6_Profiles_Profile_Base_Default_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = profile.ProfileName
    leafs["server-mode"] = profile.ServerMode
    leafs["proxy-mode"] = profile.ProxyMode
    return leafs
}

func (profile *Dhcpv6_Profiles_Profile_Base_Default_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Dhcpv6_Profiles_Profile_Base_Default_Profile) GetYangName() string { return "profile" }

func (profile *Dhcpv6_Profiles_Profile_Base_Default_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Dhcpv6_Profiles_Profile_Base_Default_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Dhcpv6_Profiles_Profile_Base_Default_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Dhcpv6_Profiles_Profile_Base_Default_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Dhcpv6_Profiles_Profile_Base_Default_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Dhcpv6_Profiles_Profile_Base_Default_Profile) GetParentYangName() string { return "default" }

// Dhcpv6_Profiles_Profile_Base_Match
// Enter match option
type Dhcpv6_Profiles_Profile_Base_Match struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of ModeClass.
    ModeClasses Dhcpv6_Profiles_Profile_Base_Match_ModeClasses
}

func (match *Dhcpv6_Profiles_Profile_Base_Match) GetFilter() yfilter.YFilter { return match.YFilter }

func (match *Dhcpv6_Profiles_Profile_Base_Match) SetFilter(yf yfilter.YFilter) { match.YFilter = yf }

func (match *Dhcpv6_Profiles_Profile_Base_Match) GetGoName(yname string) string {
    if yname == "mode-classes" { return "ModeClasses" }
    return ""
}

func (match *Dhcpv6_Profiles_Profile_Base_Match) GetSegmentPath() string {
    return "match"
}

func (match *Dhcpv6_Profiles_Profile_Base_Match) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mode-classes" {
        return &match.ModeClasses
    }
    return nil
}

func (match *Dhcpv6_Profiles_Profile_Base_Match) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mode-classes"] = &match.ModeClasses
    return children
}

func (match *Dhcpv6_Profiles_Profile_Base_Match) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (match *Dhcpv6_Profiles_Profile_Base_Match) GetBundleName() string { return "cisco_ios_xr" }

func (match *Dhcpv6_Profiles_Profile_Base_Match) GetYangName() string { return "match" }

func (match *Dhcpv6_Profiles_Profile_Base_Match) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (match *Dhcpv6_Profiles_Profile_Base_Match) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (match *Dhcpv6_Profiles_Profile_Base_Match) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (match *Dhcpv6_Profiles_Profile_Base_Match) SetParent(parent types.Entity) { match.parent = parent }

func (match *Dhcpv6_Profiles_Profile_Base_Match) GetParent() types.Entity { return match.parent }

func (match *Dhcpv6_Profiles_Profile_Base_Match) GetParentYangName() string { return "base" }

// Dhcpv6_Profiles_Profile_Base_Match_ModeClasses
// Table of ModeClass
type Dhcpv6_Profiles_Profile_Base_Match_ModeClasses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify PPP/IPoE class option. The type is slice of
    // Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass.
    ModeClass []Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass
}

func (modeClasses *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses) GetFilter() yfilter.YFilter { return modeClasses.YFilter }

func (modeClasses *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses) SetFilter(yf yfilter.YFilter) { modeClasses.YFilter = yf }

func (modeClasses *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses) GetGoName(yname string) string {
    if yname == "mode-class" { return "ModeClass" }
    return ""
}

func (modeClasses *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses) GetSegmentPath() string {
    return "mode-classes"
}

func (modeClasses *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mode-class" {
        for _, c := range modeClasses.ModeClass {
            if modeClasses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass{}
        modeClasses.ModeClass = append(modeClasses.ModeClass, child)
        return &modeClasses.ModeClass[len(modeClasses.ModeClass)-1]
    }
    return nil
}

func (modeClasses *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range modeClasses.ModeClass {
        children[modeClasses.ModeClass[i].GetSegmentPath()] = &modeClasses.ModeClass[i]
    }
    return children
}

func (modeClasses *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (modeClasses *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses) GetBundleName() string { return "cisco_ios_xr" }

func (modeClasses *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses) GetYangName() string { return "mode-classes" }

func (modeClasses *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (modeClasses *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (modeClasses *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (modeClasses *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses) SetParent(parent types.Entity) { modeClasses.parent = parent }

func (modeClasses *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses) GetParent() types.Entity { return modeClasses.parent }

func (modeClasses *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses) GetParentYangName() string { return "match" }

// Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass
// Specify PPP/IPoE class option
type Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Class name. The type is string with length:
    // 1..128.
    ClassName interface{}

    // Enter proxy or server profile. The type is slice of
    // Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile.
    Profile []Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile
}

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) GetFilter() yfilter.YFilter { return modeClass.YFilter }

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) SetFilter(yf yfilter.YFilter) { modeClass.YFilter = yf }

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) GetGoName(yname string) string {
    if yname == "class-name" { return "ClassName" }
    if yname == "profile" { return "Profile" }
    return ""
}

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) GetSegmentPath() string {
    return "mode-class" + "[class-name='" + fmt.Sprintf("%v", modeClass.ClassName) + "']"
}

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        for _, c := range modeClass.Profile {
            if modeClass.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile{}
        modeClass.Profile = append(modeClass.Profile, child)
        return &modeClass.Profile[len(modeClass.Profile)-1]
    }
    return nil
}

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range modeClass.Profile {
        children[modeClass.Profile[i].GetSegmentPath()] = &modeClass.Profile[i]
    }
    return children
}

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-name"] = modeClass.ClassName
    return leafs
}

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) GetBundleName() string { return "cisco_ios_xr" }

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) GetYangName() string { return "mode-class" }

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) SetParent(parent types.Entity) { modeClass.parent = parent }

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) GetParent() types.Entity { return modeClass.parent }

func (modeClass *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass) GetParentYangName() string { return "mode-classes" }

// Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile
// Enter proxy or server profile
type Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Profile name. The type is string with length:
    // 1..64.
    ProfileName interface{}

    // Specify mode-class based Server option. The type is interface{}.
    ServerMode interface{}

    // Specify mode-class based Proxy Option. The type is interface{}.
    ProxyMode interface{}
}

func (profile *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile) GetGoName(yname string) string {
    if yname == "profile-name" { return "ProfileName" }
    if yname == "server-mode" { return "ServerMode" }
    if yname == "proxy-mode" { return "ProxyMode" }
    return ""
}

func (profile *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile) GetSegmentPath() string {
    return "profile" + "[profile-name='" + fmt.Sprintf("%v", profile.ProfileName) + "']"
}

func (profile *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (profile *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (profile *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile-name"] = profile.ProfileName
    leafs["server-mode"] = profile.ServerMode
    leafs["proxy-mode"] = profile.ProxyMode
    return leafs
}

func (profile *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile) GetYangName() string { return "profile" }

func (profile *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Dhcpv6_Profiles_Profile_Base_Match_ModeClasses_ModeClass_Profile) GetParentYangName() string { return "mode-class" }

// Dhcpv6_Profiles_Profile_Proxy
// None
// This type is a presence type.
type Dhcpv6_Profiles_Profile_Proxy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Fill linkaddress in Relay fwd msg with Prefix from Router Advertisement for
    // PPPoE sessions. The type is interface{}.
    LinkaddressFromRaEnable interface{}

    // RouteDisable. The type is interface{}.
    RouteAddDisable interface{}

    // IPv6 address to be filled in link-address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LinkAddress interface{}

    // Create or enter proxy profile Source Interface Name. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    SrcIntfName interface{}

    // Enable None. Deletion of this object also causes deletion of all associated
    // objects under Proxy. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Table of Interface.
    Interfaces Dhcpv6_Profiles_Profile_Proxy_Interfaces

    // Specify relay configuration.
    Relay Dhcpv6_Profiles_Profile_Proxy_Relay

    // VRF related configuration.
    Vrfs Dhcpv6_Profiles_Profile_Proxy_Vrfs

    // Authentication username format.
    Authentication Dhcpv6_Profiles_Profile_Proxy_Authentication

    // Table of Class.
    Classes Dhcpv6_Profiles_Profile_Proxy_Classes

    // Change sessions configuration.
    Sessions Dhcpv6_Profiles_Profile_Proxy_Sessions
}

func (proxy *Dhcpv6_Profiles_Profile_Proxy) GetFilter() yfilter.YFilter { return proxy.YFilter }

func (proxy *Dhcpv6_Profiles_Profile_Proxy) SetFilter(yf yfilter.YFilter) { proxy.YFilter = yf }

func (proxy *Dhcpv6_Profiles_Profile_Proxy) GetGoName(yname string) string {
    if yname == "linkaddress-from-ra-enable" { return "LinkaddressFromRaEnable" }
    if yname == "route-add-disable" { return "RouteAddDisable" }
    if yname == "link-address" { return "LinkAddress" }
    if yname == "src-intf-name" { return "SrcIntfName" }
    if yname == "enable" { return "Enable" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "relay" { return "Relay" }
    if yname == "vrfs" { return "Vrfs" }
    if yname == "authentication" { return "Authentication" }
    if yname == "classes" { return "Classes" }
    if yname == "sessions" { return "Sessions" }
    return ""
}

func (proxy *Dhcpv6_Profiles_Profile_Proxy) GetSegmentPath() string {
    return "proxy"
}

func (proxy *Dhcpv6_Profiles_Profile_Proxy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &proxy.Interfaces
    }
    if childYangName == "relay" {
        return &proxy.Relay
    }
    if childYangName == "vrfs" {
        return &proxy.Vrfs
    }
    if childYangName == "authentication" {
        return &proxy.Authentication
    }
    if childYangName == "classes" {
        return &proxy.Classes
    }
    if childYangName == "sessions" {
        return &proxy.Sessions
    }
    return nil
}

func (proxy *Dhcpv6_Profiles_Profile_Proxy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &proxy.Interfaces
    children["relay"] = &proxy.Relay
    children["vrfs"] = &proxy.Vrfs
    children["authentication"] = &proxy.Authentication
    children["classes"] = &proxy.Classes
    children["sessions"] = &proxy.Sessions
    return children
}

func (proxy *Dhcpv6_Profiles_Profile_Proxy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["linkaddress-from-ra-enable"] = proxy.LinkaddressFromRaEnable
    leafs["route-add-disable"] = proxy.RouteAddDisable
    leafs["link-address"] = proxy.LinkAddress
    leafs["src-intf-name"] = proxy.SrcIntfName
    leafs["enable"] = proxy.Enable
    return leafs
}

func (proxy *Dhcpv6_Profiles_Profile_Proxy) GetBundleName() string { return "cisco_ios_xr" }

func (proxy *Dhcpv6_Profiles_Profile_Proxy) GetYangName() string { return "proxy" }

func (proxy *Dhcpv6_Profiles_Profile_Proxy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (proxy *Dhcpv6_Profiles_Profile_Proxy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (proxy *Dhcpv6_Profiles_Profile_Proxy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (proxy *Dhcpv6_Profiles_Profile_Proxy) SetParent(parent types.Entity) { proxy.parent = parent }

func (proxy *Dhcpv6_Profiles_Profile_Proxy) GetParent() types.Entity { return proxy.parent }

func (proxy *Dhcpv6_Profiles_Profile_Proxy) GetParentYangName() string { return "profile" }

// Dhcpv6_Profiles_Profile_Proxy_Interfaces
// Table of Interface
type Dhcpv6_Profiles_Profile_Proxy_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None. The type is slice of
    // Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface.
    Interface []Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface
}

func (interfaces *Dhcpv6_Profiles_Profile_Proxy_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Dhcpv6_Profiles_Profile_Proxy_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Dhcpv6_Profiles_Profile_Proxy_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Dhcpv6_Profiles_Profile_Proxy_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Dhcpv6_Profiles_Profile_Proxy_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Dhcpv6_Profiles_Profile_Proxy_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Dhcpv6_Profiles_Profile_Proxy_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Dhcpv6_Profiles_Profile_Proxy_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Dhcpv6_Profiles_Profile_Proxy_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Dhcpv6_Profiles_Profile_Proxy_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Dhcpv6_Profiles_Profile_Proxy_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Dhcpv6_Profiles_Profile_Proxy_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Dhcpv6_Profiles_Profile_Proxy_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Dhcpv6_Profiles_Profile_Proxy_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Dhcpv6_Profiles_Profile_Proxy_Interfaces) GetParentYangName() string { return "proxy" }

// Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface
// None
type Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface to configure. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Physical interface ID. The type is string.
    InterfaceId interface{}
}

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-id" { return "InterfaceId" }
    return ""
}

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-id"] = self.InterfaceId
    return leafs
}

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Dhcpv6_Profiles_Profile_Proxy_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Dhcpv6_Profiles_Profile_Proxy_Relay
// Specify relay configuration
type Dhcpv6_Profiles_Profile_Proxy_Relay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify relay option configuration.
    Option Dhcpv6_Profiles_Profile_Proxy_Relay_Option
}

func (relay *Dhcpv6_Profiles_Profile_Proxy_Relay) GetFilter() yfilter.YFilter { return relay.YFilter }

func (relay *Dhcpv6_Profiles_Profile_Proxy_Relay) SetFilter(yf yfilter.YFilter) { relay.YFilter = yf }

func (relay *Dhcpv6_Profiles_Profile_Proxy_Relay) GetGoName(yname string) string {
    if yname == "option" { return "Option" }
    return ""
}

func (relay *Dhcpv6_Profiles_Profile_Proxy_Relay) GetSegmentPath() string {
    return "relay"
}

func (relay *Dhcpv6_Profiles_Profile_Proxy_Relay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "option" {
        return &relay.Option
    }
    return nil
}

func (relay *Dhcpv6_Profiles_Profile_Proxy_Relay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["option"] = &relay.Option
    return children
}

func (relay *Dhcpv6_Profiles_Profile_Proxy_Relay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (relay *Dhcpv6_Profiles_Profile_Proxy_Relay) GetBundleName() string { return "cisco_ios_xr" }

func (relay *Dhcpv6_Profiles_Profile_Proxy_Relay) GetYangName() string { return "relay" }

func (relay *Dhcpv6_Profiles_Profile_Proxy_Relay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relay *Dhcpv6_Profiles_Profile_Proxy_Relay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relay *Dhcpv6_Profiles_Profile_Proxy_Relay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relay *Dhcpv6_Profiles_Profile_Proxy_Relay) SetParent(parent types.Entity) { relay.parent = parent }

func (relay *Dhcpv6_Profiles_Profile_Proxy_Relay) GetParent() types.Entity { return relay.parent }

func (relay *Dhcpv6_Profiles_Profile_Proxy_Relay) GetParentYangName() string { return "proxy" }

// Dhcpv6_Profiles_Profile_Proxy_Relay_Option
// Specify relay option configuration
type Dhcpv6_Profiles_Profile_Proxy_Relay_Option struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure Received SubscriberID. The type is SubscriberId.
    SubscriberId interface{}

    // Configure Received link-layer-Addr. The type is LinkLayerAddr.
    LinkLayerAddr interface{}

    // Set remote-id value from SADB. The type is interface{} with range:
    // -2147483648..2147483647.
    RemoteIDreceived interface{}

    // Enter remote-id value. The type is string with length: 1..256.
    RemoteId interface{}

    // Interface Id option.
    InterfaceId Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId
}

func (option *Dhcpv6_Profiles_Profile_Proxy_Relay_Option) GetFilter() yfilter.YFilter { return option.YFilter }

func (option *Dhcpv6_Profiles_Profile_Proxy_Relay_Option) SetFilter(yf yfilter.YFilter) { option.YFilter = yf }

func (option *Dhcpv6_Profiles_Profile_Proxy_Relay_Option) GetGoName(yname string) string {
    if yname == "subscriber-id" { return "SubscriberId" }
    if yname == "link-layer-addr" { return "LinkLayerAddr" }
    if yname == "remote-i-dreceived" { return "RemoteIDreceived" }
    if yname == "remote-id" { return "RemoteId" }
    if yname == "interface-id" { return "InterfaceId" }
    return ""
}

func (option *Dhcpv6_Profiles_Profile_Proxy_Relay_Option) GetSegmentPath() string {
    return "option"
}

func (option *Dhcpv6_Profiles_Profile_Proxy_Relay_Option) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface-id" {
        return &option.InterfaceId
    }
    return nil
}

func (option *Dhcpv6_Profiles_Profile_Proxy_Relay_Option) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interface-id"] = &option.InterfaceId
    return children
}

func (option *Dhcpv6_Profiles_Profile_Proxy_Relay_Option) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["subscriber-id"] = option.SubscriberId
    leafs["link-layer-addr"] = option.LinkLayerAddr
    leafs["remote-i-dreceived"] = option.RemoteIDreceived
    leafs["remote-id"] = option.RemoteId
    return leafs
}

func (option *Dhcpv6_Profiles_Profile_Proxy_Relay_Option) GetBundleName() string { return "cisco_ios_xr" }

func (option *Dhcpv6_Profiles_Profile_Proxy_Relay_Option) GetYangName() string { return "option" }

func (option *Dhcpv6_Profiles_Profile_Proxy_Relay_Option) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (option *Dhcpv6_Profiles_Profile_Proxy_Relay_Option) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (option *Dhcpv6_Profiles_Profile_Proxy_Relay_Option) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (option *Dhcpv6_Profiles_Profile_Proxy_Relay_Option) SetParent(parent types.Entity) { option.parent = parent }

func (option *Dhcpv6_Profiles_Profile_Proxy_Relay_Option) GetParent() types.Entity { return option.parent }

func (option *Dhcpv6_Profiles_Profile_Proxy_Relay_Option) GetParentYangName() string { return "relay" }

// Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId
// Interface Id option
type Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure InterfaceID insert type. The type is Insert.
    Insert interface{}
}

func (interfaceId *Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId) GetFilter() yfilter.YFilter { return interfaceId.YFilter }

func (interfaceId *Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId) SetFilter(yf yfilter.YFilter) { interfaceId.YFilter = yf }

func (interfaceId *Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId) GetGoName(yname string) string {
    if yname == "insert" { return "Insert" }
    return ""
}

func (interfaceId *Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId) GetSegmentPath() string {
    return "interface-id"
}

func (interfaceId *Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (interfaceId *Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (interfaceId *Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["insert"] = interfaceId.Insert
    return leafs
}

func (interfaceId *Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId) GetBundleName() string { return "cisco_ios_xr" }

func (interfaceId *Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId) GetYangName() string { return "interface-id" }

func (interfaceId *Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaceId *Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaceId *Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaceId *Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId) SetParent(parent types.Entity) { interfaceId.parent = parent }

func (interfaceId *Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId) GetParent() types.Entity { return interfaceId.parent }

func (interfaceId *Dhcpv6_Profiles_Profile_Proxy_Relay_Option_InterfaceId) GetParentYangName() string { return "option" }

// Dhcpv6_Profiles_Profile_Proxy_Vrfs
// VRF related configuration
type Dhcpv6_Profiles_Profile_Proxy_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPv6 DHCP proxy VRF name. The type is slice of
    // Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf.
    Vrf []Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf
}

func (vrfs *Dhcpv6_Profiles_Profile_Proxy_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Dhcpv6_Profiles_Profile_Proxy_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Dhcpv6_Profiles_Profile_Proxy_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Dhcpv6_Profiles_Profile_Proxy_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Dhcpv6_Profiles_Profile_Proxy_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Dhcpv6_Profiles_Profile_Proxy_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Dhcpv6_Profiles_Profile_Proxy_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Dhcpv6_Profiles_Profile_Proxy_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Dhcpv6_Profiles_Profile_Proxy_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Dhcpv6_Profiles_Profile_Proxy_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Dhcpv6_Profiles_Profile_Proxy_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Dhcpv6_Profiles_Profile_Proxy_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Dhcpv6_Profiles_Profile_Proxy_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Dhcpv6_Profiles_Profile_Proxy_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Dhcpv6_Profiles_Profile_Proxy_Vrfs) GetParentYangName() string { return "proxy" }

// Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf
// IPv6 DHCP proxy VRF name
type Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Table of HelperAddress.
    HelperAddresses Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses
}

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "helper-addresses" { return "HelperAddresses" }
    return ""
}

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "helper-addresses" {
        return &vrf.HelperAddresses
    }
    return nil
}

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["helper-addresses"] = &vrf.HelperAddresses
    return children
}

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses
// Table of HelperAddress
type Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // DHCPv6 Helper Address. The type is slice of
    // Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress.
    HelperAddress []Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses) GetFilter() yfilter.YFilter { return helperAddresses.YFilter }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses) SetFilter(yf yfilter.YFilter) { helperAddresses.YFilter = yf }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses) GetGoName(yname string) string {
    if yname == "helper-address" { return "HelperAddress" }
    return ""
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses) GetSegmentPath() string {
    return "helper-addresses"
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "helper-address" {
        for _, c := range helperAddresses.HelperAddress {
            if helperAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress{}
        helperAddresses.HelperAddress = append(helperAddresses.HelperAddress, child)
        return &helperAddresses.HelperAddress[len(helperAddresses.HelperAddress)-1]
    }
    return nil
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range helperAddresses.HelperAddress {
        children[helperAddresses.HelperAddress[i].GetSegmentPath()] = &helperAddresses.HelperAddress[i]
    }
    return children
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses) GetYangName() string { return "helper-addresses" }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses) SetParent(parent types.Entity) { helperAddresses.parent = parent }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses) GetParent() types.Entity { return helperAddresses.parent }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses) GetParentYangName() string { return "vrf" }

// Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress
// DHCPv6 Helper Address
type Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. DHCPv6 Helper Address. The type is string with
    // pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    HelperAddress interface{}

    // DHCPv6 HelperAddress Specific Output Interface. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    OutInterface interface{}

    // DHCPv6 HelperAddress Output Interface. The type is interface{}.
    AnyOutInterface interface{}
}

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetFilter() yfilter.YFilter { return helperAddress.YFilter }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) SetFilter(yf yfilter.YFilter) { helperAddress.YFilter = yf }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetGoName(yname string) string {
    if yname == "helper-address" { return "HelperAddress" }
    if yname == "out-interface" { return "OutInterface" }
    if yname == "any-out-interface" { return "AnyOutInterface" }
    return ""
}

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetSegmentPath() string {
    return "helper-address" + "[helper-address='" + fmt.Sprintf("%v", helperAddress.HelperAddress) + "']"
}

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["helper-address"] = helperAddress.HelperAddress
    leafs["out-interface"] = helperAddress.OutInterface
    leafs["any-out-interface"] = helperAddress.AnyOutInterface
    return leafs
}

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetBundleName() string { return "cisco_ios_xr" }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetYangName() string { return "helper-address" }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) SetParent(parent types.Entity) { helperAddress.parent = parent }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetParent() types.Entity { return helperAddress.parent }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Vrfs_Vrf_HelperAddresses_HelperAddress) GetParentYangName() string { return "helper-addresses" }

// Dhcpv6_Profiles_Profile_Proxy_Authentication
// Authentication username format
type Dhcpv6_Profiles_Profile_Proxy_Authentication struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set username as DUID. The type is interface{}.
    Username interface{}
}

func (authentication *Dhcpv6_Profiles_Profile_Proxy_Authentication) GetFilter() yfilter.YFilter { return authentication.YFilter }

func (authentication *Dhcpv6_Profiles_Profile_Proxy_Authentication) SetFilter(yf yfilter.YFilter) { authentication.YFilter = yf }

func (authentication *Dhcpv6_Profiles_Profile_Proxy_Authentication) GetGoName(yname string) string {
    if yname == "username" { return "Username" }
    return ""
}

func (authentication *Dhcpv6_Profiles_Profile_Proxy_Authentication) GetSegmentPath() string {
    return "authentication"
}

func (authentication *Dhcpv6_Profiles_Profile_Proxy_Authentication) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (authentication *Dhcpv6_Profiles_Profile_Proxy_Authentication) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (authentication *Dhcpv6_Profiles_Profile_Proxy_Authentication) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["username"] = authentication.Username
    return leafs
}

func (authentication *Dhcpv6_Profiles_Profile_Proxy_Authentication) GetBundleName() string { return "cisco_ios_xr" }

func (authentication *Dhcpv6_Profiles_Profile_Proxy_Authentication) GetYangName() string { return "authentication" }

func (authentication *Dhcpv6_Profiles_Profile_Proxy_Authentication) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (authentication *Dhcpv6_Profiles_Profile_Proxy_Authentication) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (authentication *Dhcpv6_Profiles_Profile_Proxy_Authentication) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (authentication *Dhcpv6_Profiles_Profile_Proxy_Authentication) SetParent(parent types.Entity) { authentication.parent = parent }

func (authentication *Dhcpv6_Profiles_Profile_Proxy_Authentication) GetParent() types.Entity { return authentication.parent }

func (authentication *Dhcpv6_Profiles_Profile_Proxy_Authentication) GetParentYangName() string { return "proxy" }

// Dhcpv6_Profiles_Profile_Proxy_Classes
// Table of Class
type Dhcpv6_Profiles_Profile_Proxy_Classes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None. The type is slice of Dhcpv6_Profiles_Profile_Proxy_Classes_Class.
    Class []Dhcpv6_Profiles_Profile_Proxy_Classes_Class
}

func (classes *Dhcpv6_Profiles_Profile_Proxy_Classes) GetFilter() yfilter.YFilter { return classes.YFilter }

func (classes *Dhcpv6_Profiles_Profile_Proxy_Classes) SetFilter(yf yfilter.YFilter) { classes.YFilter = yf }

func (classes *Dhcpv6_Profiles_Profile_Proxy_Classes) GetGoName(yname string) string {
    if yname == "class" { return "Class" }
    return ""
}

func (classes *Dhcpv6_Profiles_Profile_Proxy_Classes) GetSegmentPath() string {
    return "classes"
}

func (classes *Dhcpv6_Profiles_Profile_Proxy_Classes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class" {
        for _, c := range classes.Class {
            if classes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Profiles_Profile_Proxy_Classes_Class{}
        classes.Class = append(classes.Class, child)
        return &classes.Class[len(classes.Class)-1]
    }
    return nil
}

func (classes *Dhcpv6_Profiles_Profile_Proxy_Classes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classes.Class {
        children[classes.Class[i].GetSegmentPath()] = &classes.Class[i]
    }
    return children
}

func (classes *Dhcpv6_Profiles_Profile_Proxy_Classes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (classes *Dhcpv6_Profiles_Profile_Proxy_Classes) GetBundleName() string { return "cisco_ios_xr" }

func (classes *Dhcpv6_Profiles_Profile_Proxy_Classes) GetYangName() string { return "classes" }

func (classes *Dhcpv6_Profiles_Profile_Proxy_Classes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classes *Dhcpv6_Profiles_Profile_Proxy_Classes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classes *Dhcpv6_Profiles_Profile_Proxy_Classes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classes *Dhcpv6_Profiles_Profile_Proxy_Classes) SetParent(parent types.Entity) { classes.parent = parent }

func (classes *Dhcpv6_Profiles_Profile_Proxy_Classes) GetParent() types.Entity { return classes.parent }

func (classes *Dhcpv6_Profiles_Profile_Proxy_Classes) GetParentYangName() string { return "proxy" }

// Dhcpv6_Profiles_Profile_Proxy_Classes_Class
// None
type Dhcpv6_Profiles_Profile_Proxy_Classes_Class struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Class name. The type is string with length:
    // 1..128.
    ClassName interface{}

    // IPv6 address to be filled in link-address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    LinkAddress interface{}

    // Table of HelperAddress.
    HelperAddresses Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses
}

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) GetFilter() yfilter.YFilter { return class.YFilter }

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) SetFilter(yf yfilter.YFilter) { class.YFilter = yf }

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) GetGoName(yname string) string {
    if yname == "class-name" { return "ClassName" }
    if yname == "link-address" { return "LinkAddress" }
    if yname == "helper-addresses" { return "HelperAddresses" }
    return ""
}

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) GetSegmentPath() string {
    return "class" + "[class-name='" + fmt.Sprintf("%v", class.ClassName) + "']"
}

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "helper-addresses" {
        return &class.HelperAddresses
    }
    return nil
}

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["helper-addresses"] = &class.HelperAddresses
    return children
}

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-name"] = class.ClassName
    leafs["link-address"] = class.LinkAddress
    return leafs
}

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) GetBundleName() string { return "cisco_ios_xr" }

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) GetYangName() string { return "class" }

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) SetParent(parent types.Entity) { class.parent = parent }

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) GetParent() types.Entity { return class.parent }

func (class *Dhcpv6_Profiles_Profile_Proxy_Classes_Class) GetParentYangName() string { return "classes" }

// Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses
// Table of HelperAddress
type Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the server helper address. The type is slice of
    // Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress.
    HelperAddress []Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses) GetFilter() yfilter.YFilter { return helperAddresses.YFilter }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses) SetFilter(yf yfilter.YFilter) { helperAddresses.YFilter = yf }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses) GetGoName(yname string) string {
    if yname == "helper-address" { return "HelperAddress" }
    return ""
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses) GetSegmentPath() string {
    return "helper-addresses"
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "helper-address" {
        for _, c := range helperAddresses.HelperAddress {
            if helperAddresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress{}
        helperAddresses.HelperAddress = append(helperAddresses.HelperAddress, child)
        return &helperAddresses.HelperAddress[len(helperAddresses.HelperAddress)-1]
    }
    return nil
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range helperAddresses.HelperAddress {
        children[helperAddresses.HelperAddress[i].GetSegmentPath()] = &helperAddresses.HelperAddress[i]
    }
    return children
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses) GetBundleName() string { return "cisco_ios_xr" }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses) GetYangName() string { return "helper-addresses" }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses) SetParent(parent types.Entity) { helperAddresses.parent = parent }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses) GetParent() types.Entity { return helperAddresses.parent }

func (helperAddresses *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses) GetParentYangName() string { return "class" }

// Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress
// Specify the server helper address
type Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // This attribute is a key. Server address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    HelperAddress interface{}
}

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) GetFilter() yfilter.YFilter { return helperAddress.YFilter }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) SetFilter(yf yfilter.YFilter) { helperAddress.YFilter = yf }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "helper-address" { return "HelperAddress" }
    return ""
}

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) GetSegmentPath() string {
    return "helper-address" + "[vrf-name='" + fmt.Sprintf("%v", helperAddress.VrfName) + "']" + "[helper-address='" + fmt.Sprintf("%v", helperAddress.HelperAddress) + "']"
}

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = helperAddress.VrfName
    leafs["helper-address"] = helperAddress.HelperAddress
    return leafs
}

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) GetBundleName() string { return "cisco_ios_xr" }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) GetYangName() string { return "helper-address" }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) SetParent(parent types.Entity) { helperAddress.parent = parent }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) GetParent() types.Entity { return helperAddress.parent }

func (helperAddress *Dhcpv6_Profiles_Profile_Proxy_Classes_Class_HelperAddresses_HelperAddress) GetParentYangName() string { return "helper-addresses" }

// Dhcpv6_Profiles_Profile_Proxy_Sessions
// Change sessions configuration
type Dhcpv6_Profiles_Profile_Proxy_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Throttle DHCP sessions based on MAC address.
    Mac Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac
}

func (sessions *Dhcpv6_Profiles_Profile_Proxy_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *Dhcpv6_Profiles_Profile_Proxy_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *Dhcpv6_Profiles_Profile_Proxy_Sessions) GetGoName(yname string) string {
    if yname == "mac" { return "Mac" }
    return ""
}

func (sessions *Dhcpv6_Profiles_Profile_Proxy_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *Dhcpv6_Profiles_Profile_Proxy_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac" {
        return &sessions.Mac
    }
    return nil
}

func (sessions *Dhcpv6_Profiles_Profile_Proxy_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac"] = &sessions.Mac
    return children
}

func (sessions *Dhcpv6_Profiles_Profile_Proxy_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *Dhcpv6_Profiles_Profile_Proxy_Sessions) GetBundleName() string { return "cisco_ios_xr" }

func (sessions *Dhcpv6_Profiles_Profile_Proxy_Sessions) GetYangName() string { return "sessions" }

func (sessions *Dhcpv6_Profiles_Profile_Proxy_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessions *Dhcpv6_Profiles_Profile_Proxy_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessions *Dhcpv6_Profiles_Profile_Proxy_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessions *Dhcpv6_Profiles_Profile_Proxy_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *Dhcpv6_Profiles_Profile_Proxy_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *Dhcpv6_Profiles_Profile_Proxy_Sessions) GetParentYangName() string { return "proxy" }

// Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac
// Throttle DHCP sessions based on MAC address
type Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Throttle DHCP sessions from any one MAC address.
    Throttle Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle
}

func (mac *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac) GetFilter() yfilter.YFilter { return mac.YFilter }

func (mac *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac) SetFilter(yf yfilter.YFilter) { mac.YFilter = yf }

func (mac *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac) GetGoName(yname string) string {
    if yname == "throttle" { return "Throttle" }
    return ""
}

func (mac *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac) GetSegmentPath() string {
    return "mac"
}

func (mac *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "throttle" {
        return &mac.Throttle
    }
    return nil
}

func (mac *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["throttle"] = &mac.Throttle
    return children
}

func (mac *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mac *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac) GetBundleName() string { return "cisco_ios_xr" }

func (mac *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac) GetYangName() string { return "mac" }

func (mac *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mac *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mac *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mac *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac) SetParent(parent types.Entity) { mac.parent = parent }

func (mac *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac) GetParent() types.Entity { return mac.parent }

func (mac *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac) GetParentYangName() string { return "sessions" }

// Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle
// Throttle DHCP sessions from any one MAC
// address
type Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of solicits at which to throttle. The type is interface{} with
    // range: 1..65535.
    Limit interface{}

    // Throttle request period (in secs). The type is interface{} with range:
    // 1..100. Units are second.
    Request interface{}

    // Throttle blocking period (in secs). The type is interface{} with range:
    // 1..100. Units are second.
    Block interface{}
}

func (throttle *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle) GetFilter() yfilter.YFilter { return throttle.YFilter }

func (throttle *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle) SetFilter(yf yfilter.YFilter) { throttle.YFilter = yf }

func (throttle *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "request" { return "Request" }
    if yname == "block" { return "Block" }
    return ""
}

func (throttle *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle) GetSegmentPath() string {
    return "throttle"
}

func (throttle *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (throttle *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (throttle *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = throttle.Limit
    leafs["request"] = throttle.Request
    leafs["block"] = throttle.Block
    return leafs
}

func (throttle *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle) GetBundleName() string { return "cisco_ios_xr" }

func (throttle *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle) GetYangName() string { return "throttle" }

func (throttle *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (throttle *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (throttle *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (throttle *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle) SetParent(parent types.Entity) { throttle.parent = parent }

func (throttle *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle) GetParent() types.Entity { return throttle.parent }

func (throttle *Dhcpv6_Profiles_Profile_Proxy_Sessions_Mac_Throttle) GetParentYangName() string { return "mac" }

// Dhcpv6_Profiles_Profile_Server
// None
// This type is a presence type.
type Dhcpv6_Profiles_Profile_Server struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Address pool name. The type is string with length: 1..64.
    AddressPool interface{}

    // AFTR name. The type is string with length: 1..64.
    AftrName interface{}

    // Domain name. The type is string with length: 1..64.
    DomainName interface{}

    // DHCP Server Preference. The type is interface{} with range: 0..255.
    Preference interface{}

    // Allow RAPID Commit. The type is interface{}.
    RapidCommit interface{}

    // Enable None. Deletion of this object also causes deletion of all associated
    // objects under Server. The type is interface{}. This attribute is mandatory.
    Enable interface{}

    // Prefix pool name. The type is string with length: 1..64.
    PrefixPool interface{}

    // Change sessions configuration.
    Sessions Dhcpv6_Profiles_Profile_Server_Sessions

    // DNS servers.
    DnsServers Dhcpv6_Profiles_Profile_Server_DnsServers

    // Table of Class.
    Classes Dhcpv6_Profiles_Profile_Server_Classes

    // lease.
    Lease Dhcpv6_Profiles_Profile_Server_Lease

    // DHCPv6 options.
    Dhcpv6Options Dhcpv6_Profiles_Profile_Server_Dhcpv6Options
}

func (server *Dhcpv6_Profiles_Profile_Server) GetFilter() yfilter.YFilter { return server.YFilter }

func (server *Dhcpv6_Profiles_Profile_Server) SetFilter(yf yfilter.YFilter) { server.YFilter = yf }

func (server *Dhcpv6_Profiles_Profile_Server) GetGoName(yname string) string {
    if yname == "address-pool" { return "AddressPool" }
    if yname == "aftr-name" { return "AftrName" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "preference" { return "Preference" }
    if yname == "rapid-commit" { return "RapidCommit" }
    if yname == "enable" { return "Enable" }
    if yname == "prefix-pool" { return "PrefixPool" }
    if yname == "sessions" { return "Sessions" }
    if yname == "dns-servers" { return "DnsServers" }
    if yname == "classes" { return "Classes" }
    if yname == "lease" { return "Lease" }
    if yname == "dhcpv6-options" { return "Dhcpv6Options" }
    return ""
}

func (server *Dhcpv6_Profiles_Profile_Server) GetSegmentPath() string {
    return "server"
}

func (server *Dhcpv6_Profiles_Profile_Server) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "sessions" {
        return &server.Sessions
    }
    if childYangName == "dns-servers" {
        return &server.DnsServers
    }
    if childYangName == "classes" {
        return &server.Classes
    }
    if childYangName == "lease" {
        return &server.Lease
    }
    if childYangName == "dhcpv6-options" {
        return &server.Dhcpv6Options
    }
    return nil
}

func (server *Dhcpv6_Profiles_Profile_Server) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["sessions"] = &server.Sessions
    children["dns-servers"] = &server.DnsServers
    children["classes"] = &server.Classes
    children["lease"] = &server.Lease
    children["dhcpv6-options"] = &server.Dhcpv6Options
    return children
}

func (server *Dhcpv6_Profiles_Profile_Server) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address-pool"] = server.AddressPool
    leafs["aftr-name"] = server.AftrName
    leafs["domain-name"] = server.DomainName
    leafs["preference"] = server.Preference
    leafs["rapid-commit"] = server.RapidCommit
    leafs["enable"] = server.Enable
    leafs["prefix-pool"] = server.PrefixPool
    return leafs
}

func (server *Dhcpv6_Profiles_Profile_Server) GetBundleName() string { return "cisco_ios_xr" }

func (server *Dhcpv6_Profiles_Profile_Server) GetYangName() string { return "server" }

func (server *Dhcpv6_Profiles_Profile_Server) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (server *Dhcpv6_Profiles_Profile_Server) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (server *Dhcpv6_Profiles_Profile_Server) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (server *Dhcpv6_Profiles_Profile_Server) SetParent(parent types.Entity) { server.parent = parent }

func (server *Dhcpv6_Profiles_Profile_Server) GetParent() types.Entity { return server.parent }

func (server *Dhcpv6_Profiles_Profile_Server) GetParentYangName() string { return "profile" }

// Dhcpv6_Profiles_Profile_Server_Sessions
// Change sessions configuration
type Dhcpv6_Profiles_Profile_Server_Sessions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Throttle DHCP sessions based on MAC address.
    Mac Dhcpv6_Profiles_Profile_Server_Sessions_Mac
}

func (sessions *Dhcpv6_Profiles_Profile_Server_Sessions) GetFilter() yfilter.YFilter { return sessions.YFilter }

func (sessions *Dhcpv6_Profiles_Profile_Server_Sessions) SetFilter(yf yfilter.YFilter) { sessions.YFilter = yf }

func (sessions *Dhcpv6_Profiles_Profile_Server_Sessions) GetGoName(yname string) string {
    if yname == "mac" { return "Mac" }
    return ""
}

func (sessions *Dhcpv6_Profiles_Profile_Server_Sessions) GetSegmentPath() string {
    return "sessions"
}

func (sessions *Dhcpv6_Profiles_Profile_Server_Sessions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "mac" {
        return &sessions.Mac
    }
    return nil
}

func (sessions *Dhcpv6_Profiles_Profile_Server_Sessions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["mac"] = &sessions.Mac
    return children
}

func (sessions *Dhcpv6_Profiles_Profile_Server_Sessions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sessions *Dhcpv6_Profiles_Profile_Server_Sessions) GetBundleName() string { return "cisco_ios_xr" }

func (sessions *Dhcpv6_Profiles_Profile_Server_Sessions) GetYangName() string { return "sessions" }

func (sessions *Dhcpv6_Profiles_Profile_Server_Sessions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sessions *Dhcpv6_Profiles_Profile_Server_Sessions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sessions *Dhcpv6_Profiles_Profile_Server_Sessions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sessions *Dhcpv6_Profiles_Profile_Server_Sessions) SetParent(parent types.Entity) { sessions.parent = parent }

func (sessions *Dhcpv6_Profiles_Profile_Server_Sessions) GetParent() types.Entity { return sessions.parent }

func (sessions *Dhcpv6_Profiles_Profile_Server_Sessions) GetParentYangName() string { return "server" }

// Dhcpv6_Profiles_Profile_Server_Sessions_Mac
// Throttle DHCP sessions based on MAC address
type Dhcpv6_Profiles_Profile_Server_Sessions_Mac struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Throttle DHCP sessions from any one MAC address.
    Throttle Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle
}

func (mac *Dhcpv6_Profiles_Profile_Server_Sessions_Mac) GetFilter() yfilter.YFilter { return mac.YFilter }

func (mac *Dhcpv6_Profiles_Profile_Server_Sessions_Mac) SetFilter(yf yfilter.YFilter) { mac.YFilter = yf }

func (mac *Dhcpv6_Profiles_Profile_Server_Sessions_Mac) GetGoName(yname string) string {
    if yname == "throttle" { return "Throttle" }
    return ""
}

func (mac *Dhcpv6_Profiles_Profile_Server_Sessions_Mac) GetSegmentPath() string {
    return "mac"
}

func (mac *Dhcpv6_Profiles_Profile_Server_Sessions_Mac) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "throttle" {
        return &mac.Throttle
    }
    return nil
}

func (mac *Dhcpv6_Profiles_Profile_Server_Sessions_Mac) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["throttle"] = &mac.Throttle
    return children
}

func (mac *Dhcpv6_Profiles_Profile_Server_Sessions_Mac) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mac *Dhcpv6_Profiles_Profile_Server_Sessions_Mac) GetBundleName() string { return "cisco_ios_xr" }

func (mac *Dhcpv6_Profiles_Profile_Server_Sessions_Mac) GetYangName() string { return "mac" }

func (mac *Dhcpv6_Profiles_Profile_Server_Sessions_Mac) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mac *Dhcpv6_Profiles_Profile_Server_Sessions_Mac) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mac *Dhcpv6_Profiles_Profile_Server_Sessions_Mac) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mac *Dhcpv6_Profiles_Profile_Server_Sessions_Mac) SetParent(parent types.Entity) { mac.parent = parent }

func (mac *Dhcpv6_Profiles_Profile_Server_Sessions_Mac) GetParent() types.Entity { return mac.parent }

func (mac *Dhcpv6_Profiles_Profile_Server_Sessions_Mac) GetParentYangName() string { return "sessions" }

// Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle
// Throttle DHCP sessions from any one MAC
// address
type Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of solicits at which to throttle. The type is interface{} with
    // range: 1..65535.
    Limit interface{}

    // Throttle request period (in secs). The type is interface{} with range:
    // 1..100. Units are second.
    Request interface{}

    // Throttle blocking period (in secs). The type is interface{} with range:
    // 1..100. Units are second.
    Block interface{}
}

func (throttle *Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle) GetFilter() yfilter.YFilter { return throttle.YFilter }

func (throttle *Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle) SetFilter(yf yfilter.YFilter) { throttle.YFilter = yf }

func (throttle *Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle) GetGoName(yname string) string {
    if yname == "limit" { return "Limit" }
    if yname == "request" { return "Request" }
    if yname == "block" { return "Block" }
    return ""
}

func (throttle *Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle) GetSegmentPath() string {
    return "throttle"
}

func (throttle *Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (throttle *Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (throttle *Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["limit"] = throttle.Limit
    leafs["request"] = throttle.Request
    leafs["block"] = throttle.Block
    return leafs
}

func (throttle *Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle) GetBundleName() string { return "cisco_ios_xr" }

func (throttle *Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle) GetYangName() string { return "throttle" }

func (throttle *Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (throttle *Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (throttle *Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (throttle *Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle) SetParent(parent types.Entity) { throttle.parent = parent }

func (throttle *Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle) GetParent() types.Entity { return throttle.parent }

func (throttle *Dhcpv6_Profiles_Profile_Server_Sessions_Mac_Throttle) GetParentYangName() string { return "mac" }

// Dhcpv6_Profiles_Profile_Server_DnsServers
// DNS servers
type Dhcpv6_Profiles_Profile_Server_DnsServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Server's IPv6 address. The type is one of the following types: slice of
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DnsServer []interface{}
}

func (dnsServers *Dhcpv6_Profiles_Profile_Server_DnsServers) GetFilter() yfilter.YFilter { return dnsServers.YFilter }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_DnsServers) SetFilter(yf yfilter.YFilter) { dnsServers.YFilter = yf }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_DnsServers) GetGoName(yname string) string {
    if yname == "dns-server" { return "DnsServer" }
    return ""
}

func (dnsServers *Dhcpv6_Profiles_Profile_Server_DnsServers) GetSegmentPath() string {
    return "dns-servers"
}

func (dnsServers *Dhcpv6_Profiles_Profile_Server_DnsServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dnsServers *Dhcpv6_Profiles_Profile_Server_DnsServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dnsServers *Dhcpv6_Profiles_Profile_Server_DnsServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dns-server"] = dnsServers.DnsServer
    return leafs
}

func (dnsServers *Dhcpv6_Profiles_Profile_Server_DnsServers) GetBundleName() string { return "cisco_ios_xr" }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_DnsServers) GetYangName() string { return "dns-servers" }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_DnsServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_DnsServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_DnsServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_DnsServers) SetParent(parent types.Entity) { dnsServers.parent = parent }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_DnsServers) GetParent() types.Entity { return dnsServers.parent }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_DnsServers) GetParentYangName() string { return "server" }

// Dhcpv6_Profiles_Profile_Server_Classes
// Table of Class
type Dhcpv6_Profiles_Profile_Server_Classes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None. The type is slice of Dhcpv6_Profiles_Profile_Server_Classes_Class.
    Class []Dhcpv6_Profiles_Profile_Server_Classes_Class
}

func (classes *Dhcpv6_Profiles_Profile_Server_Classes) GetFilter() yfilter.YFilter { return classes.YFilter }

func (classes *Dhcpv6_Profiles_Profile_Server_Classes) SetFilter(yf yfilter.YFilter) { classes.YFilter = yf }

func (classes *Dhcpv6_Profiles_Profile_Server_Classes) GetGoName(yname string) string {
    if yname == "class" { return "Class" }
    return ""
}

func (classes *Dhcpv6_Profiles_Profile_Server_Classes) GetSegmentPath() string {
    return "classes"
}

func (classes *Dhcpv6_Profiles_Profile_Server_Classes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "class" {
        for _, c := range classes.Class {
            if classes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Profiles_Profile_Server_Classes_Class{}
        classes.Class = append(classes.Class, child)
        return &classes.Class[len(classes.Class)-1]
    }
    return nil
}

func (classes *Dhcpv6_Profiles_Profile_Server_Classes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range classes.Class {
        children[classes.Class[i].GetSegmentPath()] = &classes.Class[i]
    }
    return children
}

func (classes *Dhcpv6_Profiles_Profile_Server_Classes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (classes *Dhcpv6_Profiles_Profile_Server_Classes) GetBundleName() string { return "cisco_ios_xr" }

func (classes *Dhcpv6_Profiles_Profile_Server_Classes) GetYangName() string { return "classes" }

func (classes *Dhcpv6_Profiles_Profile_Server_Classes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (classes *Dhcpv6_Profiles_Profile_Server_Classes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (classes *Dhcpv6_Profiles_Profile_Server_Classes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (classes *Dhcpv6_Profiles_Profile_Server_Classes) SetParent(parent types.Entity) { classes.parent = parent }

func (classes *Dhcpv6_Profiles_Profile_Server_Classes) GetParent() types.Entity { return classes.parent }

func (classes *Dhcpv6_Profiles_Profile_Server_Classes) GetParentYangName() string { return "server" }

// Dhcpv6_Profiles_Profile_Server_Classes_Class
// None
type Dhcpv6_Profiles_Profile_Server_Classes_Class struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. class name. The type is string with length:
    // 1..128.
    ClassName interface{}

    // Address pool name. The type is string with length: 1..64.
    AddressPool interface{}

    // Domain name. The type is string with length: 1..64.
    DomainName interface{}

    // DHCP Server Preference. The type is interface{} with range: 1..255.
    Preference interface{}

    // Prefix pool name. The type is string with length: 1..64.
    PrefixPool interface{}

    // DNS servers.
    DnsServers Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers
}

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) GetFilter() yfilter.YFilter { return class.YFilter }

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) SetFilter(yf yfilter.YFilter) { class.YFilter = yf }

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) GetGoName(yname string) string {
    if yname == "class-name" { return "ClassName" }
    if yname == "address-pool" { return "AddressPool" }
    if yname == "domain-name" { return "DomainName" }
    if yname == "preference" { return "Preference" }
    if yname == "prefix-pool" { return "PrefixPool" }
    if yname == "dns-servers" { return "DnsServers" }
    return ""
}

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) GetSegmentPath() string {
    return "class" + "[class-name='" + fmt.Sprintf("%v", class.ClassName) + "']"
}

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dns-servers" {
        return &class.DnsServers
    }
    return nil
}

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["dns-servers"] = &class.DnsServers
    return children
}

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["class-name"] = class.ClassName
    leafs["address-pool"] = class.AddressPool
    leafs["domain-name"] = class.DomainName
    leafs["preference"] = class.Preference
    leafs["prefix-pool"] = class.PrefixPool
    return leafs
}

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) GetBundleName() string { return "cisco_ios_xr" }

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) GetYangName() string { return "class" }

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) SetParent(parent types.Entity) { class.parent = parent }

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) GetParent() types.Entity { return class.parent }

func (class *Dhcpv6_Profiles_Profile_Server_Classes_Class) GetParentYangName() string { return "classes" }

// Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers
// DNS servers
type Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Server's IPv6 address. The type is one of the following types: slice of
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or slice of string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    DnsServer []interface{}
}

func (dnsServers *Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers) GetFilter() yfilter.YFilter { return dnsServers.YFilter }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers) SetFilter(yf yfilter.YFilter) { dnsServers.YFilter = yf }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers) GetGoName(yname string) string {
    if yname == "dns-server" { return "DnsServer" }
    return ""
}

func (dnsServers *Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers) GetSegmentPath() string {
    return "dns-servers"
}

func (dnsServers *Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dnsServers *Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dnsServers *Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dns-server"] = dnsServers.DnsServer
    return leafs
}

func (dnsServers *Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers) GetBundleName() string { return "cisco_ios_xr" }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers) GetYangName() string { return "dns-servers" }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers) SetParent(parent types.Entity) { dnsServers.parent = parent }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers) GetParent() types.Entity { return dnsServers.parent }

func (dnsServers *Dhcpv6_Profiles_Profile_Server_Classes_Class_DnsServers) GetParentYangName() string { return "class" }

// Dhcpv6_Profiles_Profile_Server_Lease
// lease
type Dhcpv6_Profiles_Profile_Server_Lease struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Days. The type is interface{} with range: 0..365. Units are day.
    Days interface{}

    // Hours. The type is interface{} with range: 0..23. Units are hour.
    Hours interface{}

    // Minutes. The type is interface{} with range: 1..59. Units are minute.
    Minutes interface{}

    // Set string. The type is string.
    Infinite interface{}
}

func (lease *Dhcpv6_Profiles_Profile_Server_Lease) GetFilter() yfilter.YFilter { return lease.YFilter }

func (lease *Dhcpv6_Profiles_Profile_Server_Lease) SetFilter(yf yfilter.YFilter) { lease.YFilter = yf }

func (lease *Dhcpv6_Profiles_Profile_Server_Lease) GetGoName(yname string) string {
    if yname == "days" { return "Days" }
    if yname == "hours" { return "Hours" }
    if yname == "minutes" { return "Minutes" }
    if yname == "infinite" { return "Infinite" }
    return ""
}

func (lease *Dhcpv6_Profiles_Profile_Server_Lease) GetSegmentPath() string {
    return "lease"
}

func (lease *Dhcpv6_Profiles_Profile_Server_Lease) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (lease *Dhcpv6_Profiles_Profile_Server_Lease) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (lease *Dhcpv6_Profiles_Profile_Server_Lease) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["days"] = lease.Days
    leafs["hours"] = lease.Hours
    leafs["minutes"] = lease.Minutes
    leafs["infinite"] = lease.Infinite
    return leafs
}

func (lease *Dhcpv6_Profiles_Profile_Server_Lease) GetBundleName() string { return "cisco_ios_xr" }

func (lease *Dhcpv6_Profiles_Profile_Server_Lease) GetYangName() string { return "lease" }

func (lease *Dhcpv6_Profiles_Profile_Server_Lease) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (lease *Dhcpv6_Profiles_Profile_Server_Lease) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (lease *Dhcpv6_Profiles_Profile_Server_Lease) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (lease *Dhcpv6_Profiles_Profile_Server_Lease) SetParent(parent types.Entity) { lease.parent = parent }

func (lease *Dhcpv6_Profiles_Profile_Server_Lease) GetParent() types.Entity { return lease.parent }

func (lease *Dhcpv6_Profiles_Profile_Server_Lease) GetParentYangName() string { return "server" }

// Dhcpv6_Profiles_Profile_Server_Dhcpv6Options
// DHCPv6 options
type Dhcpv6_Profiles_Profile_Server_Dhcpv6Options struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Vendor options.
    VendorOptions Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions
}

func (dhcpv6Options *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options) GetFilter() yfilter.YFilter { return dhcpv6Options.YFilter }

func (dhcpv6Options *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options) SetFilter(yf yfilter.YFilter) { dhcpv6Options.YFilter = yf }

func (dhcpv6Options *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options) GetGoName(yname string) string {
    if yname == "vendor-options" { return "VendorOptions" }
    return ""
}

func (dhcpv6Options *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options) GetSegmentPath() string {
    return "dhcpv6-options"
}

func (dhcpv6Options *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vendor-options" {
        return &dhcpv6Options.VendorOptions
    }
    return nil
}

func (dhcpv6Options *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vendor-options"] = &dhcpv6Options.VendorOptions
    return children
}

func (dhcpv6Options *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dhcpv6Options *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options) GetBundleName() string { return "cisco_ios_xr" }

func (dhcpv6Options *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options) GetYangName() string { return "dhcpv6-options" }

func (dhcpv6Options *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dhcpv6Options *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dhcpv6Options *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dhcpv6Options *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options) SetParent(parent types.Entity) { dhcpv6Options.parent = parent }

func (dhcpv6Options *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options) GetParent() types.Entity { return dhcpv6Options.parent }

func (dhcpv6Options *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options) GetParentYangName() string { return "server" }

// Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions
// Vendor options
type Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Set string. The type is string.
    Type interface{}

    // Vendor options. The type is string with length: 1..512.
    VendorOptions interface{}
}

func (vendorOptions *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions) GetFilter() yfilter.YFilter { return vendorOptions.YFilter }

func (vendorOptions *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions) SetFilter(yf yfilter.YFilter) { vendorOptions.YFilter = yf }

func (vendorOptions *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions) GetGoName(yname string) string {
    if yname == "type" { return "Type" }
    if yname == "vendor-options" { return "VendorOptions" }
    return ""
}

func (vendorOptions *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions) GetSegmentPath() string {
    return "vendor-options"
}

func (vendorOptions *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vendorOptions *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vendorOptions *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["type"] = vendorOptions.Type
    leafs["vendor-options"] = vendorOptions.VendorOptions
    return leafs
}

func (vendorOptions *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions) GetBundleName() string { return "cisco_ios_xr" }

func (vendorOptions *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions) GetYangName() string { return "vendor-options" }

func (vendorOptions *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vendorOptions *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vendorOptions *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vendorOptions *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions) SetParent(parent types.Entity) { vendorOptions.parent = parent }

func (vendorOptions *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions) GetParent() types.Entity { return vendorOptions.parent }

func (vendorOptions *Dhcpv6_Profiles_Profile_Server_Dhcpv6Options_VendorOptions) GetParentYangName() string { return "dhcpv6-options" }

// Dhcpv6_Interfaces
// Table of Interface
type Dhcpv6_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // None. The type is slice of Dhcpv6_Interfaces_Interface.
    Interface []Dhcpv6_Interfaces_Interface
}

func (interfaces *Dhcpv6_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Dhcpv6_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Dhcpv6_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Dhcpv6_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Dhcpv6_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Dhcpv6_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Dhcpv6_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Dhcpv6_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Dhcpv6_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Dhcpv6_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Dhcpv6_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Dhcpv6_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Dhcpv6_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Dhcpv6_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Dhcpv6_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Dhcpv6_Interfaces) GetParentYangName() string { return "dhcpv6" }

// Dhcpv6_Interfaces_Interface
// None
type Dhcpv6_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface to configure. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // PPPoE subscriber interface.
    Pppoe Dhcpv6_Interfaces_Interface_Pppoe

    // Assign a proxy profile to interface.
    Proxy Dhcpv6_Interfaces_Interface_Proxy

    // Assign a base profile to interface.
    Base Dhcpv6_Interfaces_Interface_Base

    // Assign a server profile to interface.
    Server Dhcpv6_Interfaces_Interface_Server

    // Assign a relay profile to interface.
    Relay Dhcpv6_Interfaces_Interface_Relay
}

func (self *Dhcpv6_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Dhcpv6_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Dhcpv6_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "pppoe" { return "Pppoe" }
    if yname == "proxy" { return "Proxy" }
    if yname == "base" { return "Base" }
    if yname == "server" { return "Server" }
    if yname == "relay" { return "Relay" }
    return ""
}

func (self *Dhcpv6_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Dhcpv6_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "pppoe" {
        return &self.Pppoe
    }
    if childYangName == "proxy" {
        return &self.Proxy
    }
    if childYangName == "base" {
        return &self.Base
    }
    if childYangName == "server" {
        return &self.Server
    }
    if childYangName == "relay" {
        return &self.Relay
    }
    return nil
}

func (self *Dhcpv6_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["pppoe"] = &self.Pppoe
    children["proxy"] = &self.Proxy
    children["base"] = &self.Base
    children["server"] = &self.Server
    children["relay"] = &self.Relay
    return children
}

func (self *Dhcpv6_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    return leafs
}

func (self *Dhcpv6_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Dhcpv6_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Dhcpv6_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Dhcpv6_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Dhcpv6_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Dhcpv6_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Dhcpv6_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Dhcpv6_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Dhcpv6_Interfaces_Interface_Pppoe
// PPPoE subscriber interface
type Dhcpv6_Interfaces_Interface_Pppoe struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enter profile name. The type is string with length: 1..64.
    Profile interface{}
}

func (pppoe *Dhcpv6_Interfaces_Interface_Pppoe) GetFilter() yfilter.YFilter { return pppoe.YFilter }

func (pppoe *Dhcpv6_Interfaces_Interface_Pppoe) SetFilter(yf yfilter.YFilter) { pppoe.YFilter = yf }

func (pppoe *Dhcpv6_Interfaces_Interface_Pppoe) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (pppoe *Dhcpv6_Interfaces_Interface_Pppoe) GetSegmentPath() string {
    return "pppoe"
}

func (pppoe *Dhcpv6_Interfaces_Interface_Pppoe) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (pppoe *Dhcpv6_Interfaces_Interface_Pppoe) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (pppoe *Dhcpv6_Interfaces_Interface_Pppoe) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile"] = pppoe.Profile
    return leafs
}

func (pppoe *Dhcpv6_Interfaces_Interface_Pppoe) GetBundleName() string { return "cisco_ios_xr" }

func (pppoe *Dhcpv6_Interfaces_Interface_Pppoe) GetYangName() string { return "pppoe" }

func (pppoe *Dhcpv6_Interfaces_Interface_Pppoe) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (pppoe *Dhcpv6_Interfaces_Interface_Pppoe) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (pppoe *Dhcpv6_Interfaces_Interface_Pppoe) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (pppoe *Dhcpv6_Interfaces_Interface_Pppoe) SetParent(parent types.Entity) { pppoe.parent = parent }

func (pppoe *Dhcpv6_Interfaces_Interface_Pppoe) GetParent() types.Entity { return pppoe.parent }

func (pppoe *Dhcpv6_Interfaces_Interface_Pppoe) GetParentYangName() string { return "interface" }

// Dhcpv6_Interfaces_Interface_Proxy
// Assign a proxy profile to interface
type Dhcpv6_Interfaces_Interface_Proxy struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enter profile name. The type is string with length: 1..64.
    Profile interface{}
}

func (proxy *Dhcpv6_Interfaces_Interface_Proxy) GetFilter() yfilter.YFilter { return proxy.YFilter }

func (proxy *Dhcpv6_Interfaces_Interface_Proxy) SetFilter(yf yfilter.YFilter) { proxy.YFilter = yf }

func (proxy *Dhcpv6_Interfaces_Interface_Proxy) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (proxy *Dhcpv6_Interfaces_Interface_Proxy) GetSegmentPath() string {
    return "proxy"
}

func (proxy *Dhcpv6_Interfaces_Interface_Proxy) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (proxy *Dhcpv6_Interfaces_Interface_Proxy) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (proxy *Dhcpv6_Interfaces_Interface_Proxy) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile"] = proxy.Profile
    return leafs
}

func (proxy *Dhcpv6_Interfaces_Interface_Proxy) GetBundleName() string { return "cisco_ios_xr" }

func (proxy *Dhcpv6_Interfaces_Interface_Proxy) GetYangName() string { return "proxy" }

func (proxy *Dhcpv6_Interfaces_Interface_Proxy) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (proxy *Dhcpv6_Interfaces_Interface_Proxy) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (proxy *Dhcpv6_Interfaces_Interface_Proxy) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (proxy *Dhcpv6_Interfaces_Interface_Proxy) SetParent(parent types.Entity) { proxy.parent = parent }

func (proxy *Dhcpv6_Interfaces_Interface_Proxy) GetParent() types.Entity { return proxy.parent }

func (proxy *Dhcpv6_Interfaces_Interface_Proxy) GetParentYangName() string { return "interface" }

// Dhcpv6_Interfaces_Interface_Base
// Assign a base profile to interface
type Dhcpv6_Interfaces_Interface_Base struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enter profile name. The type is string with length: 1..64.
    Profile interface{}
}

func (base *Dhcpv6_Interfaces_Interface_Base) GetFilter() yfilter.YFilter { return base.YFilter }

func (base *Dhcpv6_Interfaces_Interface_Base) SetFilter(yf yfilter.YFilter) { base.YFilter = yf }

func (base *Dhcpv6_Interfaces_Interface_Base) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (base *Dhcpv6_Interfaces_Interface_Base) GetSegmentPath() string {
    return "base"
}

func (base *Dhcpv6_Interfaces_Interface_Base) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (base *Dhcpv6_Interfaces_Interface_Base) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (base *Dhcpv6_Interfaces_Interface_Base) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile"] = base.Profile
    return leafs
}

func (base *Dhcpv6_Interfaces_Interface_Base) GetBundleName() string { return "cisco_ios_xr" }

func (base *Dhcpv6_Interfaces_Interface_Base) GetYangName() string { return "base" }

func (base *Dhcpv6_Interfaces_Interface_Base) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (base *Dhcpv6_Interfaces_Interface_Base) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (base *Dhcpv6_Interfaces_Interface_Base) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (base *Dhcpv6_Interfaces_Interface_Base) SetParent(parent types.Entity) { base.parent = parent }

func (base *Dhcpv6_Interfaces_Interface_Base) GetParent() types.Entity { return base.parent }

func (base *Dhcpv6_Interfaces_Interface_Base) GetParentYangName() string { return "interface" }

// Dhcpv6_Interfaces_Interface_Server
// Assign a server profile to interface
type Dhcpv6_Interfaces_Interface_Server struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enter profile name. The type is string with length: 1..64.
    Profile interface{}
}

func (server *Dhcpv6_Interfaces_Interface_Server) GetFilter() yfilter.YFilter { return server.YFilter }

func (server *Dhcpv6_Interfaces_Interface_Server) SetFilter(yf yfilter.YFilter) { server.YFilter = yf }

func (server *Dhcpv6_Interfaces_Interface_Server) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (server *Dhcpv6_Interfaces_Interface_Server) GetSegmentPath() string {
    return "server"
}

func (server *Dhcpv6_Interfaces_Interface_Server) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (server *Dhcpv6_Interfaces_Interface_Server) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (server *Dhcpv6_Interfaces_Interface_Server) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile"] = server.Profile
    return leafs
}

func (server *Dhcpv6_Interfaces_Interface_Server) GetBundleName() string { return "cisco_ios_xr" }

func (server *Dhcpv6_Interfaces_Interface_Server) GetYangName() string { return "server" }

func (server *Dhcpv6_Interfaces_Interface_Server) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (server *Dhcpv6_Interfaces_Interface_Server) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (server *Dhcpv6_Interfaces_Interface_Server) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (server *Dhcpv6_Interfaces_Interface_Server) SetParent(parent types.Entity) { server.parent = parent }

func (server *Dhcpv6_Interfaces_Interface_Server) GetParent() types.Entity { return server.parent }

func (server *Dhcpv6_Interfaces_Interface_Server) GetParentYangName() string { return "interface" }

// Dhcpv6_Interfaces_Interface_Relay
// Assign a relay profile to interface
type Dhcpv6_Interfaces_Interface_Relay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enter profile name. The type is string with length: 1..64.
    Profile interface{}
}

func (relay *Dhcpv6_Interfaces_Interface_Relay) GetFilter() yfilter.YFilter { return relay.YFilter }

func (relay *Dhcpv6_Interfaces_Interface_Relay) SetFilter(yf yfilter.YFilter) { relay.YFilter = yf }

func (relay *Dhcpv6_Interfaces_Interface_Relay) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (relay *Dhcpv6_Interfaces_Interface_Relay) GetSegmentPath() string {
    return "relay"
}

func (relay *Dhcpv6_Interfaces_Interface_Relay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (relay *Dhcpv6_Interfaces_Interface_Relay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (relay *Dhcpv6_Interfaces_Interface_Relay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["profile"] = relay.Profile
    return leafs
}

func (relay *Dhcpv6_Interfaces_Interface_Relay) GetBundleName() string { return "cisco_ios_xr" }

func (relay *Dhcpv6_Interfaces_Interface_Relay) GetYangName() string { return "relay" }

func (relay *Dhcpv6_Interfaces_Interface_Relay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (relay *Dhcpv6_Interfaces_Interface_Relay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (relay *Dhcpv6_Interfaces_Interface_Relay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (relay *Dhcpv6_Interfaces_Interface_Relay) SetParent(parent types.Entity) { relay.parent = parent }

func (relay *Dhcpv6_Interfaces_Interface_Relay) GetParent() types.Entity { return relay.parent }

func (relay *Dhcpv6_Interfaces_Interface_Relay) GetParentYangName() string { return "interface" }

