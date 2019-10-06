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
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of SSRP Profiles.
    Profiles Ssrp_Profiles
}

func (ssrp *Ssrp) GetEntityData() *types.CommonEntityData {
    ssrp.EntityData.YFilter = ssrp.YFilter
    ssrp.EntityData.YangName = "ssrp"
    ssrp.EntityData.BundleName = "cisco_ios_xr"
    ssrp.EntityData.ParentYangName = "Cisco-IOS-XR-ppp-ma-ssrp-cfg"
    ssrp.EntityData.SegmentPath = "Cisco-IOS-XR-ppp-ma-ssrp-cfg:ssrp"
    ssrp.EntityData.AbsolutePath = ssrp.EntityData.SegmentPath
    ssrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssrp.EntityData.Children = types.NewOrderedMap()
    ssrp.EntityData.Children.Append("profiles", types.YChild{"Profiles", &ssrp.Profiles})
    ssrp.EntityData.Leafs = types.NewOrderedMap()

    ssrp.EntityData.YListKeys = []string {}

    return &(ssrp.EntityData)
}

// Ssrp_Profiles
// Table of SSRP Profiles
type Ssrp_Profiles struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SSRP Profile configuration. The type is slice of Ssrp_Profiles_Profile.
    Profile []*Ssrp_Profiles_Profile
}

func (profiles *Ssrp_Profiles) GetEntityData() *types.CommonEntityData {
    profiles.EntityData.YFilter = profiles.YFilter
    profiles.EntityData.YangName = "profiles"
    profiles.EntityData.BundleName = "cisco_ios_xr"
    profiles.EntityData.ParentYangName = "ssrp"
    profiles.EntityData.SegmentPath = "profiles"
    profiles.EntityData.AbsolutePath = "Cisco-IOS-XR-ppp-ma-ssrp-cfg:ssrp/" + profiles.EntityData.SegmentPath
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

// Ssrp_Profiles_Profile
// SSRP Profile configuration
type Ssrp_Profiles_Profile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

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

func (profile *Ssrp_Profiles_Profile) GetEntityData() *types.CommonEntityData {
    profile.EntityData.YFilter = profile.YFilter
    profile.EntityData.YangName = "profile"
    profile.EntityData.BundleName = "cisco_ios_xr"
    profile.EntityData.ParentYangName = "profiles"
    profile.EntityData.SegmentPath = "profile" + types.AddKeyToken(profile.Name, "name")
    profile.EntityData.AbsolutePath = "Cisco-IOS-XR-ppp-ma-ssrp-cfg:ssrp/profiles/" + profile.EntityData.SegmentPath
    profile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profile.EntityData.Children = types.NewOrderedMap()
    profile.EntityData.Leafs = types.NewOrderedMap()
    profile.EntityData.Leafs.Append("name", types.YLeaf{"Name", profile.Name})
    profile.EntityData.Leafs.Append("max-hops", types.YLeaf{"MaxHops", profile.MaxHops})
    profile.EntityData.Leafs.Append("peer-ipv4-address", types.YLeaf{"PeerIpv4Address", profile.PeerIpv4Address})

    profile.EntityData.YListKeys = []string {"Name"}

    return &(profile.EntityData)
}

