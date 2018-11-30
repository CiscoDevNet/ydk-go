// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-telnet-mgmt package configuration.
// 
// This module contains definitions
// for the following management objects:
//   telnet: Global Telnet configuration commands
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_telnet_mgmt_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_telnet_mgmt_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-telnet-mgmt-cfg telnet}", reflect.TypeOf(Telnet{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-telnet-mgmt-cfg:telnet", reflect.TypeOf(Telnet{}))
}

// Telnet
// Global Telnet configuration commands
type Telnet struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name for telnet service.
    Vrfs Telnet_Vrfs
}

func (telnet *Telnet) GetEntityData() *types.CommonEntityData {
    telnet.EntityData.YFilter = telnet.YFilter
    telnet.EntityData.YangName = "telnet"
    telnet.EntityData.BundleName = "cisco_ios_xr"
    telnet.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-telnet-mgmt-cfg"
    telnet.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-telnet-mgmt-cfg:telnet"
    telnet.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    telnet.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    telnet.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    telnet.EntityData.Children = types.NewOrderedMap()
    telnet.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &telnet.Vrfs})
    telnet.EntityData.Leafs = types.NewOrderedMap()

    telnet.EntityData.YListKeys = []string {}

    return &(telnet.EntityData)
}

// Telnet_Vrfs
// VRF name for telnet service
type Telnet_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name for telnet service. The type is slice of Telnet_Vrfs_Vrf.
    Vrf []*Telnet_Vrfs_Vrf
}

func (vrfs *Telnet_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "telnet"
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

// Telnet_Vrfs_Vrf
// VRF name for telnet service
type Telnet_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv4 configuration.
    Ipv4 Telnet_Vrfs_Vrf_Ipv4
}

func (vrf *Telnet_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &vrf.Ipv4})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Telnet_Vrfs_Vrf_Ipv4
// IPv4 configuration
type Telnet_Vrfs_Vrf_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the DSCP value. The type is interface{} with range: 0..63.
    Dscp interface{}
}

func (ipv4 *Telnet_Vrfs_Vrf_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "vrf"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("dscp", types.YLeaf{"Dscp", ipv4.Dscp})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

