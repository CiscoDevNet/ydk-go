// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-telnet-mgmt package configuration.
// 
// This module contains definitions
// for the following management objects:
//   telnet: Global Telnet configuration commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name for telnet service.
    Vrfs Telnet_Vrfs
}

func (telnet *Telnet) GetFilter() yfilter.YFilter { return telnet.YFilter }

func (telnet *Telnet) SetFilter(yf yfilter.YFilter) { telnet.YFilter = yf }

func (telnet *Telnet) GetGoName(yname string) string {
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (telnet *Telnet) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-telnet-mgmt-cfg:telnet"
}

func (telnet *Telnet) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &telnet.Vrfs
    }
    return nil
}

func (telnet *Telnet) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &telnet.Vrfs
    return children
}

func (telnet *Telnet) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (telnet *Telnet) GetBundleName() string { return "cisco_ios_xr" }

func (telnet *Telnet) GetYangName() string { return "telnet" }

func (telnet *Telnet) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (telnet *Telnet) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (telnet *Telnet) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (telnet *Telnet) SetParent(parent types.Entity) { telnet.parent = parent }

func (telnet *Telnet) GetParent() types.Entity { return telnet.parent }

func (telnet *Telnet) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-telnet-mgmt-cfg" }

// Telnet_Vrfs
// VRF name for telnet service
type Telnet_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF name for telnet service. The type is slice of Telnet_Vrfs_Vrf.
    Vrf []Telnet_Vrfs_Vrf
}

func (vrfs *Telnet_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Telnet_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Telnet_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Telnet_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Telnet_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Telnet_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Telnet_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Telnet_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Telnet_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Telnet_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Telnet_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Telnet_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Telnet_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Telnet_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Telnet_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Telnet_Vrfs) GetParentYangName() string { return "telnet" }

// Telnet_Vrfs_Vrf
// VRF name for telnet service
type Telnet_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // IPv4 configuration.
    Ipv4 Telnet_Vrfs_Vrf_Ipv4
}

func (vrf *Telnet_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Telnet_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Telnet_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (vrf *Telnet_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Telnet_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4" {
        return &vrf.Ipv4
    }
    return nil
}

func (vrf *Telnet_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4"] = &vrf.Ipv4
    return children
}

func (vrf *Telnet_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Telnet_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Telnet_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Telnet_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Telnet_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Telnet_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Telnet_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Telnet_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Telnet_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Telnet_Vrfs_Vrf_Ipv4
// IPv4 configuration
type Telnet_Vrfs_Vrf_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the DSCP value. The type is interface{} with range: 0..63.
    Dscp interface{}
}

func (ipv4 *Telnet_Vrfs_Vrf_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Telnet_Vrfs_Vrf_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Telnet_Vrfs_Vrf_Ipv4) GetGoName(yname string) string {
    if yname == "dscp" { return "Dscp" }
    return ""
}

func (ipv4 *Telnet_Vrfs_Vrf_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Telnet_Vrfs_Vrf_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4 *Telnet_Vrfs_Vrf_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4 *Telnet_Vrfs_Vrf_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["dscp"] = ipv4.Dscp
    return leafs
}

func (ipv4 *Telnet_Vrfs_Vrf_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Telnet_Vrfs_Vrf_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Telnet_Vrfs_Vrf_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Telnet_Vrfs_Vrf_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Telnet_Vrfs_Vrf_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Telnet_Vrfs_Vrf_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Telnet_Vrfs_Vrf_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Telnet_Vrfs_Vrf_Ipv4) GetParentYangName() string { return "vrf" }

