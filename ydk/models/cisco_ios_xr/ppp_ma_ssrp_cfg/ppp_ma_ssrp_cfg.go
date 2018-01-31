// This module contains a collection of YANG definitions
// for Cisco IOS-XR ppp-ma-ssrp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ssrp: Shared plane SSRP configuration data
// 
// This YANG module augments the
//   Cisco-IOS-XR-config-mda-cfg,
//   Cisco-IOS-XR-ifmgr-cfg
// modules with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ppp_ma_ssrp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ppp_ma_ssrp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ppp-ma-ssrp-cfg ssrp}", reflect.TypeOf(Ssrp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ppp-ma-ssrp-cfg:ssrp", reflect.TypeOf(Ssrp{}))
}

// Ssrp
// Shared plane SSRP configuration data
type Ssrp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of SSRP Profiles.
    Profiles Ssrp_Profiles
}

func (ssrp *Ssrp) GetFilter() yfilter.YFilter { return ssrp.YFilter }

func (ssrp *Ssrp) SetFilter(yf yfilter.YFilter) { ssrp.YFilter = yf }

func (ssrp *Ssrp) GetGoName(yname string) string {
    if yname == "profiles" { return "Profiles" }
    return ""
}

func (ssrp *Ssrp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ppp-ma-ssrp-cfg:ssrp"
}

func (ssrp *Ssrp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profiles" {
        return &ssrp.Profiles
    }
    return nil
}

func (ssrp *Ssrp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["profiles"] = &ssrp.Profiles
    return children
}

func (ssrp *Ssrp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ssrp *Ssrp) GetBundleName() string { return "cisco_ios_xr" }

func (ssrp *Ssrp) GetYangName() string { return "ssrp" }

func (ssrp *Ssrp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ssrp *Ssrp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ssrp *Ssrp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ssrp *Ssrp) SetParent(parent types.Entity) { ssrp.parent = parent }

func (ssrp *Ssrp) GetParent() types.Entity { return ssrp.parent }

func (ssrp *Ssrp) GetParentYangName() string { return "Cisco-IOS-XR-ppp-ma-ssrp-cfg" }

// Ssrp_Profiles
// Table of SSRP Profiles
type Ssrp_Profiles struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SSRP Profile configuration. The type is slice of Ssrp_Profiles_Profile.
    Profile []Ssrp_Profiles_Profile
}

func (profiles *Ssrp_Profiles) GetFilter() yfilter.YFilter { return profiles.YFilter }

func (profiles *Ssrp_Profiles) SetFilter(yf yfilter.YFilter) { profiles.YFilter = yf }

func (profiles *Ssrp_Profiles) GetGoName(yname string) string {
    if yname == "profile" { return "Profile" }
    return ""
}

func (profiles *Ssrp_Profiles) GetSegmentPath() string {
    return "profiles"
}

func (profiles *Ssrp_Profiles) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "profile" {
        for _, c := range profiles.Profile {
            if profiles.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Ssrp_Profiles_Profile{}
        profiles.Profile = append(profiles.Profile, child)
        return &profiles.Profile[len(profiles.Profile)-1]
    }
    return nil
}

func (profiles *Ssrp_Profiles) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range profiles.Profile {
        children[profiles.Profile[i].GetSegmentPath()] = &profiles.Profile[i]
    }
    return children
}

func (profiles *Ssrp_Profiles) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (profiles *Ssrp_Profiles) GetBundleName() string { return "cisco_ios_xr" }

func (profiles *Ssrp_Profiles) GetYangName() string { return "profiles" }

func (profiles *Ssrp_Profiles) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profiles *Ssrp_Profiles) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profiles *Ssrp_Profiles) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profiles *Ssrp_Profiles) SetParent(parent types.Entity) { profiles.parent = parent }

func (profiles *Ssrp_Profiles) GetParent() types.Entity { return profiles.parent }

func (profiles *Ssrp_Profiles) GetParentYangName() string { return "ssrp" }

// Ssrp_Profiles_Profile
// SSRP Profile configuration
type Ssrp_Profiles_Profile struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the profile. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    Name interface{}

    // This specifies the maximum number of hops for packets on the SSO channel.
    // The type is interface{} with range: 1..255.
    MaxHops interface{}

    // This specifies the remote end's IPv4-address for the SSO channel. The type
    // is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PeerIpv4Address interface{}
}

func (profile *Ssrp_Profiles_Profile) GetFilter() yfilter.YFilter { return profile.YFilter }

func (profile *Ssrp_Profiles_Profile) SetFilter(yf yfilter.YFilter) { profile.YFilter = yf }

func (profile *Ssrp_Profiles_Profile) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "max-hops" { return "MaxHops" }
    if yname == "peer-ipv4-address" { return "PeerIpv4Address" }
    return ""
}

func (profile *Ssrp_Profiles_Profile) GetSegmentPath() string {
    return "profile" + "[name='" + fmt.Sprintf("%v", profile.Name) + "']"
}

func (profile *Ssrp_Profiles_Profile) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (profile *Ssrp_Profiles_Profile) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (profile *Ssrp_Profiles_Profile) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = profile.Name
    leafs["max-hops"] = profile.MaxHops
    leafs["peer-ipv4-address"] = profile.PeerIpv4Address
    return leafs
}

func (profile *Ssrp_Profiles_Profile) GetBundleName() string { return "cisco_ios_xr" }

func (profile *Ssrp_Profiles_Profile) GetYangName() string { return "profile" }

func (profile *Ssrp_Profiles_Profile) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (profile *Ssrp_Profiles_Profile) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (profile *Ssrp_Profiles_Profile) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (profile *Ssrp_Profiles_Profile) SetParent(parent types.Entity) { profile.parent = parent }

func (profile *Ssrp_Profiles_Profile) GetParent() types.Entity { return profile.parent }

func (profile *Ssrp_Profiles_Profile) GetParentYangName() string { return "profiles" }

