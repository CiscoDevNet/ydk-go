// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-sbfd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   sbfd: Seamless BFD (S-BFD) operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_sbfd_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_sbfd_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-sbfd-oper sbfd}", reflect.TypeOf(Sbfd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-sbfd-oper:sbfd", reflect.TypeOf(Sbfd{}))
}

// SbfdAddressFamily represents Sbfd address family
type SbfdAddressFamily string

const (
    // ipv4
    SbfdAddressFamily_ipv4 SbfdAddressFamily = "ipv4"

    // ipv6
    SbfdAddressFamily_ipv6 SbfdAddressFamily = "ipv6"
)

// BfdAfId represents Bfd af id
type BfdAfId string

const (
    // No Address
    BfdAfId_bfd_af_id_none BfdAfId = "bfd-af-id-none"

    // IPv4 AFI
    BfdAfId_bfd_af_id_ipv4 BfdAfId = "bfd-af-id-ipv4"

    // IPv6 AFI
    BfdAfId_bfd_af_id_ipv6 BfdAfId = "bfd-af-id-ipv6"
)

// Sbfd
// Seamless BFD (S-BFD) operational data
type Sbfd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Target-identifier information.
    TargetIdentifier Sbfd_TargetIdentifier
}

func (sbfd *Sbfd) GetFilter() yfilter.YFilter { return sbfd.YFilter }

func (sbfd *Sbfd) SetFilter(yf yfilter.YFilter) { sbfd.YFilter = yf }

func (sbfd *Sbfd) GetGoName(yname string) string {
    if yname == "target-identifier" { return "TargetIdentifier" }
    return ""
}

func (sbfd *Sbfd) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-sbfd-oper:sbfd"
}

func (sbfd *Sbfd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "target-identifier" {
        return &sbfd.TargetIdentifier
    }
    return nil
}

func (sbfd *Sbfd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["target-identifier"] = &sbfd.TargetIdentifier
    return children
}

func (sbfd *Sbfd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (sbfd *Sbfd) GetBundleName() string { return "cisco_ios_xr" }

func (sbfd *Sbfd) GetYangName() string { return "sbfd" }

func (sbfd *Sbfd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (sbfd *Sbfd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (sbfd *Sbfd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (sbfd *Sbfd) SetParent(parent types.Entity) { sbfd.parent = parent }

func (sbfd *Sbfd) GetParent() types.Entity { return sbfd.parent }

func (sbfd *Sbfd) GetParentYangName() string { return "Cisco-IOS-XR-ip-sbfd-oper" }

// Sbfd_TargetIdentifier
// Target-identifier information
type Sbfd_TargetIdentifier struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // SBFD remote discriminator data.
    RemoteVrfs Sbfd_TargetIdentifier_RemoteVrfs

    // SBFD local discriminator  data.
    LocalVrfs Sbfd_TargetIdentifier_LocalVrfs
}

func (targetIdentifier *Sbfd_TargetIdentifier) GetFilter() yfilter.YFilter { return targetIdentifier.YFilter }

func (targetIdentifier *Sbfd_TargetIdentifier) SetFilter(yf yfilter.YFilter) { targetIdentifier.YFilter = yf }

func (targetIdentifier *Sbfd_TargetIdentifier) GetGoName(yname string) string {
    if yname == "remote-vrfs" { return "RemoteVrfs" }
    if yname == "local-vrfs" { return "LocalVrfs" }
    return ""
}

func (targetIdentifier *Sbfd_TargetIdentifier) GetSegmentPath() string {
    return "target-identifier"
}

func (targetIdentifier *Sbfd_TargetIdentifier) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-vrfs" {
        return &targetIdentifier.RemoteVrfs
    }
    if childYangName == "local-vrfs" {
        return &targetIdentifier.LocalVrfs
    }
    return nil
}

func (targetIdentifier *Sbfd_TargetIdentifier) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["remote-vrfs"] = &targetIdentifier.RemoteVrfs
    children["local-vrfs"] = &targetIdentifier.LocalVrfs
    return children
}

func (targetIdentifier *Sbfd_TargetIdentifier) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (targetIdentifier *Sbfd_TargetIdentifier) GetBundleName() string { return "cisco_ios_xr" }

func (targetIdentifier *Sbfd_TargetIdentifier) GetYangName() string { return "target-identifier" }

func (targetIdentifier *Sbfd_TargetIdentifier) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (targetIdentifier *Sbfd_TargetIdentifier) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (targetIdentifier *Sbfd_TargetIdentifier) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (targetIdentifier *Sbfd_TargetIdentifier) SetParent(parent types.Entity) { targetIdentifier.parent = parent }

func (targetIdentifier *Sbfd_TargetIdentifier) GetParent() types.Entity { return targetIdentifier.parent }

func (targetIdentifier *Sbfd_TargetIdentifier) GetParentYangName() string { return "sbfd" }

// Sbfd_TargetIdentifier_RemoteVrfs
// SBFD remote discriminator data
type Sbfd_TargetIdentifier_RemoteVrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of remote discriminator data per VRF. The type is slice of
    // Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf.
    RemoteVrf []Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf
}

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) GetFilter() yfilter.YFilter { return remoteVrfs.YFilter }

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) SetFilter(yf yfilter.YFilter) { remoteVrfs.YFilter = yf }

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) GetGoName(yname string) string {
    if yname == "remote-vrf" { return "RemoteVrf" }
    return ""
}

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) GetSegmentPath() string {
    return "remote-vrfs"
}

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-vrf" {
        for _, c := range remoteVrfs.RemoteVrf {
            if remoteVrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf{}
        remoteVrfs.RemoteVrf = append(remoteVrfs.RemoteVrf, child)
        return &remoteVrfs.RemoteVrf[len(remoteVrfs.RemoteVrf)-1]
    }
    return nil
}

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range remoteVrfs.RemoteVrf {
        children[remoteVrfs.RemoteVrf[i].GetSegmentPath()] = &remoteVrfs.RemoteVrf[i]
    }
    return children
}

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) GetBundleName() string { return "cisco_ios_xr" }

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) GetYangName() string { return "remote-vrfs" }

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) SetParent(parent types.Entity) { remoteVrfs.parent = parent }

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) GetParent() types.Entity { return remoteVrfs.parent }

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) GetParentYangName() string { return "target-identifier" }

// Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf
// Table of remote discriminator data per VRF
type Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // SBFD remote discriminator . The type is slice of
    // Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator.
    RemoteDiscriminator []Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator
}

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) GetFilter() yfilter.YFilter { return remoteVrf.YFilter }

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) SetFilter(yf yfilter.YFilter) { remoteVrf.YFilter = yf }

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "remote-discriminator" { return "RemoteDiscriminator" }
    return ""
}

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) GetSegmentPath() string {
    return "remote-vrf" + "[vrf-name='" + fmt.Sprintf("%v", remoteVrf.VrfName) + "']"
}

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-discriminator" {
        for _, c := range remoteVrf.RemoteDiscriminator {
            if remoteVrf.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator{}
        remoteVrf.RemoteDiscriminator = append(remoteVrf.RemoteDiscriminator, child)
        return &remoteVrf.RemoteDiscriminator[len(remoteVrf.RemoteDiscriminator)-1]
    }
    return nil
}

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range remoteVrf.RemoteDiscriminator {
        children[remoteVrf.RemoteDiscriminator[i].GetSegmentPath()] = &remoteVrf.RemoteDiscriminator[i]
    }
    return children
}

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = remoteVrf.VrfName
    return leafs
}

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) GetBundleName() string { return "cisco_ios_xr" }

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) GetYangName() string { return "remote-vrf" }

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) SetParent(parent types.Entity) { remoteVrf.parent = parent }

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) GetParent() types.Entity { return remoteVrf.parent }

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) GetParentYangName() string { return "remote-vrfs" }

// Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator
// SBFD remote discriminator 
type Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF Name . The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Remote Discriminator. The type is interface{} with range:
    // -2147483648..2147483647.
    RemoteDiscriminator interface{}

    // Address. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Target identifier for sbfd. The type is SbfdAddressFamily.
    TidType interface{}

    // Remote discriminator. The type is interface{} with range: 0..4294967295.
    Discr interface{}

    // VRF Name. The type is string.
    VrfNameXr interface{}

    // Status. The type is string.
    Status interface{}

    // Discriminator source name. The type is string.
    DiscrSrc interface{}

    // IP address.
    IpAddress Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress
}

func (remoteDiscriminator *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator) GetFilter() yfilter.YFilter { return remoteDiscriminator.YFilter }

func (remoteDiscriminator *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator) SetFilter(yf yfilter.YFilter) { remoteDiscriminator.YFilter = yf }

func (remoteDiscriminator *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "remote-discriminator" { return "RemoteDiscriminator" }
    if yname == "address" { return "Address" }
    if yname == "tid-type" { return "TidType" }
    if yname == "discr" { return "Discr" }
    if yname == "vrf-name-xr" { return "VrfNameXr" }
    if yname == "status" { return "Status" }
    if yname == "discr-src" { return "DiscrSrc" }
    if yname == "ip-address" { return "IpAddress" }
    return ""
}

func (remoteDiscriminator *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator) GetSegmentPath() string {
    return "remote-discriminator"
}

func (remoteDiscriminator *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ip-address" {
        return &remoteDiscriminator.IpAddress
    }
    return nil
}

func (remoteDiscriminator *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ip-address"] = &remoteDiscriminator.IpAddress
    return children
}

func (remoteDiscriminator *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = remoteDiscriminator.VrfName
    leafs["remote-discriminator"] = remoteDiscriminator.RemoteDiscriminator
    leafs["address"] = remoteDiscriminator.Address
    leafs["tid-type"] = remoteDiscriminator.TidType
    leafs["discr"] = remoteDiscriminator.Discr
    leafs["vrf-name-xr"] = remoteDiscriminator.VrfNameXr
    leafs["status"] = remoteDiscriminator.Status
    leafs["discr-src"] = remoteDiscriminator.DiscrSrc
    return leafs
}

func (remoteDiscriminator *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator) GetBundleName() string { return "cisco_ios_xr" }

func (remoteDiscriminator *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator) GetYangName() string { return "remote-discriminator" }

func (remoteDiscriminator *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteDiscriminator *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteDiscriminator *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteDiscriminator *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator) SetParent(parent types.Entity) { remoteDiscriminator.parent = parent }

func (remoteDiscriminator *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator) GetParent() types.Entity { return remoteDiscriminator.parent }

func (remoteDiscriminator *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator) GetParentYangName() string { return "remote-vrf" }

// Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress
// IP address
type Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) GetFilter() yfilter.YFilter { return ipAddress.YFilter }

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) SetFilter(yf yfilter.YFilter) { ipAddress.YFilter = yf }

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) GetGoName(yname string) string {
    if yname == "afi" { return "Afi" }
    if yname == "dummy" { return "Dummy" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "ipv6" { return "Ipv6" }
    return ""
}

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) GetSegmentPath() string {
    return "ip-address"
}

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["afi"] = ipAddress.Afi
    leafs["dummy"] = ipAddress.Dummy
    leafs["ipv4"] = ipAddress.Ipv4
    leafs["ipv6"] = ipAddress.Ipv6
    return leafs
}

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) GetBundleName() string { return "cisco_ios_xr" }

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) GetYangName() string { return "ip-address" }

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) SetParent(parent types.Entity) { ipAddress.parent = parent }

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) GetParent() types.Entity { return ipAddress.parent }

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) GetParentYangName() string { return "remote-discriminator" }

// Sbfd_TargetIdentifier_LocalVrfs
// SBFD local discriminator  data
type Sbfd_TargetIdentifier_LocalVrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Table of local discriminator data per VRF. The type is slice of
    // Sbfd_TargetIdentifier_LocalVrfs_LocalVrf.
    LocalVrf []Sbfd_TargetIdentifier_LocalVrfs_LocalVrf
}

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) GetFilter() yfilter.YFilter { return localVrfs.YFilter }

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) SetFilter(yf yfilter.YFilter) { localVrfs.YFilter = yf }

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) GetGoName(yname string) string {
    if yname == "local-vrf" { return "LocalVrf" }
    return ""
}

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) GetSegmentPath() string {
    return "local-vrfs"
}

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-vrf" {
        for _, c := range localVrfs.LocalVrf {
            if localVrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sbfd_TargetIdentifier_LocalVrfs_LocalVrf{}
        localVrfs.LocalVrf = append(localVrfs.LocalVrf, child)
        return &localVrfs.LocalVrf[len(localVrfs.LocalVrf)-1]
    }
    return nil
}

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range localVrfs.LocalVrf {
        children[localVrfs.LocalVrf[i].GetSegmentPath()] = &localVrfs.LocalVrf[i]
    }
    return children
}

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) GetBundleName() string { return "cisco_ios_xr" }

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) GetYangName() string { return "local-vrfs" }

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) SetParent(parent types.Entity) { localVrfs.parent = parent }

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) GetParent() types.Entity { return localVrfs.parent }

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) GetParentYangName() string { return "target-identifier" }

// Sbfd_TargetIdentifier_LocalVrfs_LocalVrf
// Table of local discriminator data per VRF
type Sbfd_TargetIdentifier_LocalVrfs_LocalVrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // SBFD local discriminator . The type is slice of
    // Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator.
    LocalDiscriminator []Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator
}

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) GetFilter() yfilter.YFilter { return localVrf.YFilter }

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) SetFilter(yf yfilter.YFilter) { localVrf.YFilter = yf }

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "local-discriminator" { return "LocalDiscriminator" }
    return ""
}

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) GetSegmentPath() string {
    return "local-vrf" + "[vrf-name='" + fmt.Sprintf("%v", localVrf.VrfName) + "']"
}

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "local-discriminator" {
        for _, c := range localVrf.LocalDiscriminator {
            if localVrf.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator{}
        localVrf.LocalDiscriminator = append(localVrf.LocalDiscriminator, child)
        return &localVrf.LocalDiscriminator[len(localVrf.LocalDiscriminator)-1]
    }
    return nil
}

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range localVrf.LocalDiscriminator {
        children[localVrf.LocalDiscriminator[i].GetSegmentPath()] = &localVrf.LocalDiscriminator[i]
    }
    return children
}

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = localVrf.VrfName
    return leafs
}

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) GetBundleName() string { return "cisco_ios_xr" }

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) GetYangName() string { return "local-vrf" }

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) SetParent(parent types.Entity) { localVrf.parent = parent }

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) GetParent() types.Entity { return localVrf.parent }

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) GetParentYangName() string { return "local-vrfs" }

// Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator
// SBFD local discriminator 
type Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local discriminator. The type is interface{} with range:
    // -2147483648..2147483647.
    LocalDiscriminator interface{}

    // VRF Name . The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Local discriminator. The type is interface{} with range: 0..4294967295.
    Discr interface{}

    // VRF Name. The type is string.
    VrfNameXr interface{}

    // MODE name. The type is string.
    Flags interface{}

    // Status. The type is string.
    Status interface{}

    // Discriminator source name. The type is string.
    DiscrSrc interface{}
}

func (localDiscriminator *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator) GetFilter() yfilter.YFilter { return localDiscriminator.YFilter }

func (localDiscriminator *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator) SetFilter(yf yfilter.YFilter) { localDiscriminator.YFilter = yf }

func (localDiscriminator *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator) GetGoName(yname string) string {
    if yname == "local-discriminator" { return "LocalDiscriminator" }
    if yname == "vrf-name" { return "VrfName" }
    if yname == "discr" { return "Discr" }
    if yname == "vrf-name-xr" { return "VrfNameXr" }
    if yname == "flags" { return "Flags" }
    if yname == "status" { return "Status" }
    if yname == "discr-src" { return "DiscrSrc" }
    return ""
}

func (localDiscriminator *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator) GetSegmentPath() string {
    return "local-discriminator"
}

func (localDiscriminator *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (localDiscriminator *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (localDiscriminator *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["local-discriminator"] = localDiscriminator.LocalDiscriminator
    leafs["vrf-name"] = localDiscriminator.VrfName
    leafs["discr"] = localDiscriminator.Discr
    leafs["vrf-name-xr"] = localDiscriminator.VrfNameXr
    leafs["flags"] = localDiscriminator.Flags
    leafs["status"] = localDiscriminator.Status
    leafs["discr-src"] = localDiscriminator.DiscrSrc
    return leafs
}

func (localDiscriminator *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator) GetBundleName() string { return "cisco_ios_xr" }

func (localDiscriminator *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator) GetYangName() string { return "local-discriminator" }

func (localDiscriminator *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localDiscriminator *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localDiscriminator *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localDiscriminator *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator) SetParent(parent types.Entity) { localDiscriminator.parent = parent }

func (localDiscriminator *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator) GetParent() types.Entity { return localDiscriminator.parent }

func (localDiscriminator *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator) GetParentYangName() string { return "local-vrf" }

