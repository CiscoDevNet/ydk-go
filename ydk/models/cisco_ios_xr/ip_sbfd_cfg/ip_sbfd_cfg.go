// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-sbfd package configuration.
// 
// This module contains definitions
// for the following management objects:
//   sbfd: SBFD Configuration ,Seamless-BFD is method for detecting
//     faultsbetween two different nodes in a network
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_sbfd_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_sbfd_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-sbfd-cfg sbfd}", reflect.TypeOf(Sbfd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-sbfd-cfg:sbfd", reflect.TypeOf(Sbfd{}))
}

// Sbfd
// SBFD Configuration ,Seamless-BFD is method for
// detecting faultsbetween two different nodes in a
// network
type Sbfd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // configure remote target.
    RemoteTarget Sbfd_RemoteTarget

    // Configure local discriminator.
    LocalDiscriminator Sbfd_LocalDiscriminator
}

func (sbfd *Sbfd) GetFilter() yfilter.YFilter { return sbfd.YFilter }

func (sbfd *Sbfd) SetFilter(yf yfilter.YFilter) { sbfd.YFilter = yf }

func (sbfd *Sbfd) GetGoName(yname string) string {
    if yname == "remote-target" { return "RemoteTarget" }
    if yname == "local-discriminator" { return "LocalDiscriminator" }
    return ""
}

func (sbfd *Sbfd) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-sbfd-cfg:sbfd"
}

func (sbfd *Sbfd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-target" {
        return &sbfd.RemoteTarget
    }
    if childYangName == "local-discriminator" {
        return &sbfd.LocalDiscriminator
    }
    return nil
}

func (sbfd *Sbfd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["remote-target"] = &sbfd.RemoteTarget
    children["local-discriminator"] = &sbfd.LocalDiscriminator
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

func (sbfd *Sbfd) GetParentYangName() string { return "Cisco-IOS-XR-ip-sbfd-cfg" }

// Sbfd_RemoteTarget
// configure remote target
type Sbfd_RemoteTarget struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv4 address as target.
    Ipv4Addresses Sbfd_RemoteTarget_Ipv4Addresses

    // ipv6 address as target.
    Ipv6Addresses Sbfd_RemoteTarget_Ipv6Addresses
}

func (remoteTarget *Sbfd_RemoteTarget) GetFilter() yfilter.YFilter { return remoteTarget.YFilter }

func (remoteTarget *Sbfd_RemoteTarget) SetFilter(yf yfilter.YFilter) { remoteTarget.YFilter = yf }

func (remoteTarget *Sbfd_RemoteTarget) GetGoName(yname string) string {
    if yname == "ipv4-addresses" { return "Ipv4Addresses" }
    if yname == "ipv6-addresses" { return "Ipv6Addresses" }
    return ""
}

func (remoteTarget *Sbfd_RemoteTarget) GetSegmentPath() string {
    return "remote-target"
}

func (remoteTarget *Sbfd_RemoteTarget) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-addresses" {
        return &remoteTarget.Ipv4Addresses
    }
    if childYangName == "ipv6-addresses" {
        return &remoteTarget.Ipv6Addresses
    }
    return nil
}

func (remoteTarget *Sbfd_RemoteTarget) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv4-addresses"] = &remoteTarget.Ipv4Addresses
    children["ipv6-addresses"] = &remoteTarget.Ipv6Addresses
    return children
}

func (remoteTarget *Sbfd_RemoteTarget) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (remoteTarget *Sbfd_RemoteTarget) GetBundleName() string { return "cisco_ios_xr" }

func (remoteTarget *Sbfd_RemoteTarget) GetYangName() string { return "remote-target" }

func (remoteTarget *Sbfd_RemoteTarget) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteTarget *Sbfd_RemoteTarget) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteTarget *Sbfd_RemoteTarget) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteTarget *Sbfd_RemoteTarget) SetParent(parent types.Entity) { remoteTarget.parent = parent }

func (remoteTarget *Sbfd_RemoteTarget) GetParent() types.Entity { return remoteTarget.parent }

func (remoteTarget *Sbfd_RemoteTarget) GetParentYangName() string { return "sbfd" }

// Sbfd_RemoteTarget_Ipv4Addresses
// ipv4 address as target
type Sbfd_RemoteTarget_Ipv4Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP Address Value for RemoteDiscriminatorTable. The type is slice of
    // Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address.
    Ipv4Address []Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address
}

func (ipv4Addresses *Sbfd_RemoteTarget_Ipv4Addresses) GetFilter() yfilter.YFilter { return ipv4Addresses.YFilter }

func (ipv4Addresses *Sbfd_RemoteTarget_Ipv4Addresses) SetFilter(yf yfilter.YFilter) { ipv4Addresses.YFilter = yf }

func (ipv4Addresses *Sbfd_RemoteTarget_Ipv4Addresses) GetGoName(yname string) string {
    if yname == "ipv4-address" { return "Ipv4Address" }
    return ""
}

func (ipv4Addresses *Sbfd_RemoteTarget_Ipv4Addresses) GetSegmentPath() string {
    return "ipv4-addresses"
}

func (ipv4Addresses *Sbfd_RemoteTarget_Ipv4Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-address" {
        for _, c := range ipv4Addresses.Ipv4Address {
            if ipv4Addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address{}
        ipv4Addresses.Ipv4Address = append(ipv4Addresses.Ipv4Address, child)
        return &ipv4Addresses.Ipv4Address[len(ipv4Addresses.Ipv4Address)-1]
    }
    return nil
}

func (ipv4Addresses *Sbfd_RemoteTarget_Ipv4Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4Addresses.Ipv4Address {
        children[ipv4Addresses.Ipv4Address[i].GetSegmentPath()] = &ipv4Addresses.Ipv4Address[i]
    }
    return children
}

func (ipv4Addresses *Sbfd_RemoteTarget_Ipv4Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4Addresses *Sbfd_RemoteTarget_Ipv4Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Addresses *Sbfd_RemoteTarget_Ipv4Addresses) GetYangName() string { return "ipv4-addresses" }

func (ipv4Addresses *Sbfd_RemoteTarget_Ipv4Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Addresses *Sbfd_RemoteTarget_Ipv4Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Addresses *Sbfd_RemoteTarget_Ipv4Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Addresses *Sbfd_RemoteTarget_Ipv4Addresses) SetParent(parent types.Entity) { ipv4Addresses.parent = parent }

func (ipv4Addresses *Sbfd_RemoteTarget_Ipv4Addresses) GetParent() types.Entity { return ipv4Addresses.parent }

func (ipv4Addresses *Sbfd_RemoteTarget_Ipv4Addresses) GetParentYangName() string { return "remote-target" }

// Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address
// IP Address Value for RemoteDiscriminatorTable
type Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key.  IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Remote Discriminator value. The type is slice of
    // Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator.
    RemoteDiscriminator []Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator
}

func (ipv4Address *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address) GetFilter() yfilter.YFilter { return ipv4Address.YFilter }

func (ipv4Address *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address) SetFilter(yf yfilter.YFilter) { ipv4Address.YFilter = yf }

func (ipv4Address *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "remote-discriminator" { return "RemoteDiscriminator" }
    return ""
}

func (ipv4Address *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address) GetSegmentPath() string {
    return "ipv4-address" + "[address='" + fmt.Sprintf("%v", ipv4Address.Address) + "']"
}

func (ipv4Address *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-discriminator" {
        for _, c := range ipv4Address.RemoteDiscriminator {
            if ipv4Address.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator{}
        ipv4Address.RemoteDiscriminator = append(ipv4Address.RemoteDiscriminator, child)
        return &ipv4Address.RemoteDiscriminator[len(ipv4Address.RemoteDiscriminator)-1]
    }
    return nil
}

func (ipv4Address *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4Address.RemoteDiscriminator {
        children[ipv4Address.RemoteDiscriminator[i].GetSegmentPath()] = &ipv4Address.RemoteDiscriminator[i]
    }
    return children
}

func (ipv4Address *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = ipv4Address.Address
    return leafs
}

func (ipv4Address *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Address *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address) GetYangName() string { return "ipv4-address" }

func (ipv4Address *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Address *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Address *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Address *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address) SetParent(parent types.Entity) { ipv4Address.parent = parent }

func (ipv4Address *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address) GetParent() types.Entity { return ipv4Address.parent }

func (ipv4Address *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address) GetParentYangName() string { return "ipv4-addresses" }

// Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator
// Remote Discriminator value
type Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Remote Discriminator Value. The type is
    // interface{} with range: 1..4294967295.
    RemoteDiscriminator interface{}
}

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator) GetFilter() yfilter.YFilter { return remoteDiscriminator.YFilter }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator) SetFilter(yf yfilter.YFilter) { remoteDiscriminator.YFilter = yf }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator) GetGoName(yname string) string {
    if yname == "remote-discriminator" { return "RemoteDiscriminator" }
    return ""
}

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator) GetSegmentPath() string {
    return "remote-discriminator" + "[remote-discriminator='" + fmt.Sprintf("%v", remoteDiscriminator.RemoteDiscriminator) + "']"
}

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["remote-discriminator"] = remoteDiscriminator.RemoteDiscriminator
    return leafs
}

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator) GetBundleName() string { return "cisco_ios_xr" }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator) GetYangName() string { return "remote-discriminator" }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator) SetParent(parent types.Entity) { remoteDiscriminator.parent = parent }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator) GetParent() types.Entity { return remoteDiscriminator.parent }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator) GetParentYangName() string { return "ipv4-address" }

// Sbfd_RemoteTarget_Ipv6Addresses
// ipv6 address as target
type Sbfd_RemoteTarget_Ipv6Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IP Address Value for RemoteDiscriminatorTable. The type is slice of
    // Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address.
    Ipv6Address []Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address
}

func (ipv6Addresses *Sbfd_RemoteTarget_Ipv6Addresses) GetFilter() yfilter.YFilter { return ipv6Addresses.YFilter }

func (ipv6Addresses *Sbfd_RemoteTarget_Ipv6Addresses) SetFilter(yf yfilter.YFilter) { ipv6Addresses.YFilter = yf }

func (ipv6Addresses *Sbfd_RemoteTarget_Ipv6Addresses) GetGoName(yname string) string {
    if yname == "ipv6-address" { return "Ipv6Address" }
    return ""
}

func (ipv6Addresses *Sbfd_RemoteTarget_Ipv6Addresses) GetSegmentPath() string {
    return "ipv6-addresses"
}

func (ipv6Addresses *Sbfd_RemoteTarget_Ipv6Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-address" {
        for _, c := range ipv6Addresses.Ipv6Address {
            if ipv6Addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address{}
        ipv6Addresses.Ipv6Address = append(ipv6Addresses.Ipv6Address, child)
        return &ipv6Addresses.Ipv6Address[len(ipv6Addresses.Ipv6Address)-1]
    }
    return nil
}

func (ipv6Addresses *Sbfd_RemoteTarget_Ipv6Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6Addresses.Ipv6Address {
        children[ipv6Addresses.Ipv6Address[i].GetSegmentPath()] = &ipv6Addresses.Ipv6Address[i]
    }
    return children
}

func (ipv6Addresses *Sbfd_RemoteTarget_Ipv6Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6Addresses *Sbfd_RemoteTarget_Ipv6Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Addresses *Sbfd_RemoteTarget_Ipv6Addresses) GetYangName() string { return "ipv6-addresses" }

func (ipv6Addresses *Sbfd_RemoteTarget_Ipv6Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Addresses *Sbfd_RemoteTarget_Ipv6Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Addresses *Sbfd_RemoteTarget_Ipv6Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Addresses *Sbfd_RemoteTarget_Ipv6Addresses) SetParent(parent types.Entity) { ipv6Addresses.parent = parent }

func (ipv6Addresses *Sbfd_RemoteTarget_Ipv6Addresses) GetParent() types.Entity { return ipv6Addresses.parent }

func (ipv6Addresses *Sbfd_RemoteTarget_Ipv6Addresses) GetParentYangName() string { return "remote-target" }

// Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address
// IP Address Value for RemoteDiscriminatorTable
type Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key.  IPv6 adddress. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Remote Discriminator value. The type is slice of
    // Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator.
    RemoteDiscriminator []Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator
}

func (ipv6Address *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address) GetFilter() yfilter.YFilter { return ipv6Address.YFilter }

func (ipv6Address *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address) SetFilter(yf yfilter.YFilter) { ipv6Address.YFilter = yf }

func (ipv6Address *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "remote-discriminator" { return "RemoteDiscriminator" }
    return ""
}

func (ipv6Address *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address) GetSegmentPath() string {
    return "ipv6-address" + "[address='" + fmt.Sprintf("%v", ipv6Address.Address) + "']"
}

func (ipv6Address *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "remote-discriminator" {
        for _, c := range ipv6Address.RemoteDiscriminator {
            if ipv6Address.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator{}
        ipv6Address.RemoteDiscriminator = append(ipv6Address.RemoteDiscriminator, child)
        return &ipv6Address.RemoteDiscriminator[len(ipv6Address.RemoteDiscriminator)-1]
    }
    return nil
}

func (ipv6Address *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv6Address.RemoteDiscriminator {
        children[ipv6Address.RemoteDiscriminator[i].GetSegmentPath()] = &ipv6Address.RemoteDiscriminator[i]
    }
    return children
}

func (ipv6Address *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = ipv6Address.Address
    return leafs
}

func (ipv6Address *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Address *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address) GetYangName() string { return "ipv6-address" }

func (ipv6Address *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Address *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Address *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Address *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address) SetParent(parent types.Entity) { ipv6Address.parent = parent }

func (ipv6Address *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address) GetParent() types.Entity { return ipv6Address.parent }

func (ipv6Address *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address) GetParentYangName() string { return "ipv6-addresses" }

// Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator
// Remote Discriminator value
type Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Remote Discriminator Value. The type is
    // interface{} with range: 1..4294967295.
    RemoteDiscriminator interface{}
}

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator) GetFilter() yfilter.YFilter { return remoteDiscriminator.YFilter }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator) SetFilter(yf yfilter.YFilter) { remoteDiscriminator.YFilter = yf }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator) GetGoName(yname string) string {
    if yname == "remote-discriminator" { return "RemoteDiscriminator" }
    return ""
}

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator) GetSegmentPath() string {
    return "remote-discriminator" + "[remote-discriminator='" + fmt.Sprintf("%v", remoteDiscriminator.RemoteDiscriminator) + "']"
}

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["remote-discriminator"] = remoteDiscriminator.RemoteDiscriminator
    return leafs
}

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator) GetBundleName() string { return "cisco_ios_xr" }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator) GetYangName() string { return "remote-discriminator" }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator) SetParent(parent types.Entity) { remoteDiscriminator.parent = parent }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator) GetParent() types.Entity { return remoteDiscriminator.parent }

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator) GetParentYangName() string { return "ipv6-address" }

// Sbfd_LocalDiscriminator
// Configure local discriminator
type Sbfd_LocalDiscriminator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure local discriminator from interface address.
    IntfDiscriminators Sbfd_LocalDiscriminator_IntfDiscriminators

    // Configure local discriminator dynamically.
    DynamicDiscriminators Sbfd_LocalDiscriminator_DynamicDiscriminators

    // Configure local discriminator as an ipv4 address.
    Ipv4Discriminators Sbfd_LocalDiscriminator_Ipv4Discriminators

    // Configure local discriminator as an integer.
    Val32Discriminators Sbfd_LocalDiscriminator_Val32Discriminators
}

func (localDiscriminator *Sbfd_LocalDiscriminator) GetFilter() yfilter.YFilter { return localDiscriminator.YFilter }

func (localDiscriminator *Sbfd_LocalDiscriminator) SetFilter(yf yfilter.YFilter) { localDiscriminator.YFilter = yf }

func (localDiscriminator *Sbfd_LocalDiscriminator) GetGoName(yname string) string {
    if yname == "intf-discriminators" { return "IntfDiscriminators" }
    if yname == "dynamic-discriminators" { return "DynamicDiscriminators" }
    if yname == "ipv4-discriminators" { return "Ipv4Discriminators" }
    if yname == "val32-discriminators" { return "Val32Discriminators" }
    return ""
}

func (localDiscriminator *Sbfd_LocalDiscriminator) GetSegmentPath() string {
    return "local-discriminator"
}

func (localDiscriminator *Sbfd_LocalDiscriminator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "intf-discriminators" {
        return &localDiscriminator.IntfDiscriminators
    }
    if childYangName == "dynamic-discriminators" {
        return &localDiscriminator.DynamicDiscriminators
    }
    if childYangName == "ipv4-discriminators" {
        return &localDiscriminator.Ipv4Discriminators
    }
    if childYangName == "val32-discriminators" {
        return &localDiscriminator.Val32Discriminators
    }
    return nil
}

func (localDiscriminator *Sbfd_LocalDiscriminator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["intf-discriminators"] = &localDiscriminator.IntfDiscriminators
    children["dynamic-discriminators"] = &localDiscriminator.DynamicDiscriminators
    children["ipv4-discriminators"] = &localDiscriminator.Ipv4Discriminators
    children["val32-discriminators"] = &localDiscriminator.Val32Discriminators
    return children
}

func (localDiscriminator *Sbfd_LocalDiscriminator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (localDiscriminator *Sbfd_LocalDiscriminator) GetBundleName() string { return "cisco_ios_xr" }

func (localDiscriminator *Sbfd_LocalDiscriminator) GetYangName() string { return "local-discriminator" }

func (localDiscriminator *Sbfd_LocalDiscriminator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (localDiscriminator *Sbfd_LocalDiscriminator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (localDiscriminator *Sbfd_LocalDiscriminator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (localDiscriminator *Sbfd_LocalDiscriminator) SetParent(parent types.Entity) { localDiscriminator.parent = parent }

func (localDiscriminator *Sbfd_LocalDiscriminator) GetParent() types.Entity { return localDiscriminator.parent }

func (localDiscriminator *Sbfd_LocalDiscriminator) GetParentYangName() string { return "sbfd" }

// Sbfd_LocalDiscriminator_IntfDiscriminators
// Configure local discriminator from interface
// address
type Sbfd_LocalDiscriminator_IntfDiscriminators struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // interface address as discriminator. The type is slice of
    // Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator.
    IntfDiscriminator []Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator
}

func (intfDiscriminators *Sbfd_LocalDiscriminator_IntfDiscriminators) GetFilter() yfilter.YFilter { return intfDiscriminators.YFilter }

func (intfDiscriminators *Sbfd_LocalDiscriminator_IntfDiscriminators) SetFilter(yf yfilter.YFilter) { intfDiscriminators.YFilter = yf }

func (intfDiscriminators *Sbfd_LocalDiscriminator_IntfDiscriminators) GetGoName(yname string) string {
    if yname == "intf-discriminator" { return "IntfDiscriminator" }
    return ""
}

func (intfDiscriminators *Sbfd_LocalDiscriminator_IntfDiscriminators) GetSegmentPath() string {
    return "intf-discriminators"
}

func (intfDiscriminators *Sbfd_LocalDiscriminator_IntfDiscriminators) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "intf-discriminator" {
        for _, c := range intfDiscriminators.IntfDiscriminator {
            if intfDiscriminators.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator{}
        intfDiscriminators.IntfDiscriminator = append(intfDiscriminators.IntfDiscriminator, child)
        return &intfDiscriminators.IntfDiscriminator[len(intfDiscriminators.IntfDiscriminator)-1]
    }
    return nil
}

func (intfDiscriminators *Sbfd_LocalDiscriminator_IntfDiscriminators) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range intfDiscriminators.IntfDiscriminator {
        children[intfDiscriminators.IntfDiscriminator[i].GetSegmentPath()] = &intfDiscriminators.IntfDiscriminator[i]
    }
    return children
}

func (intfDiscriminators *Sbfd_LocalDiscriminator_IntfDiscriminators) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (intfDiscriminators *Sbfd_LocalDiscriminator_IntfDiscriminators) GetBundleName() string { return "cisco_ios_xr" }

func (intfDiscriminators *Sbfd_LocalDiscriminator_IntfDiscriminators) GetYangName() string { return "intf-discriminators" }

func (intfDiscriminators *Sbfd_LocalDiscriminator_IntfDiscriminators) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (intfDiscriminators *Sbfd_LocalDiscriminator_IntfDiscriminators) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (intfDiscriminators *Sbfd_LocalDiscriminator_IntfDiscriminators) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (intfDiscriminators *Sbfd_LocalDiscriminator_IntfDiscriminators) SetParent(parent types.Entity) { intfDiscriminators.parent = parent }

func (intfDiscriminators *Sbfd_LocalDiscriminator_IntfDiscriminators) GetParent() types.Entity { return intfDiscriminators.parent }

func (intfDiscriminators *Sbfd_LocalDiscriminator_IntfDiscriminators) GetParentYangName() string { return "local-discriminator" }

// Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator
// interface address as discriminator
type Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (intfDiscriminator *Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator) GetFilter() yfilter.YFilter { return intfDiscriminator.YFilter }

func (intfDiscriminator *Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator) SetFilter(yf yfilter.YFilter) { intfDiscriminator.YFilter = yf }

func (intfDiscriminator *Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (intfDiscriminator *Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator) GetSegmentPath() string {
    return "intf-discriminator" + "[interface-name='" + fmt.Sprintf("%v", intfDiscriminator.InterfaceName) + "']"
}

func (intfDiscriminator *Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (intfDiscriminator *Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (intfDiscriminator *Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = intfDiscriminator.InterfaceName
    return leafs
}

func (intfDiscriminator *Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator) GetBundleName() string { return "cisco_ios_xr" }

func (intfDiscriminator *Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator) GetYangName() string { return "intf-discriminator" }

func (intfDiscriminator *Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (intfDiscriminator *Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (intfDiscriminator *Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (intfDiscriminator *Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator) SetParent(parent types.Entity) { intfDiscriminator.parent = parent }

func (intfDiscriminator *Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator) GetParent() types.Entity { return intfDiscriminator.parent }

func (intfDiscriminator *Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator) GetParentYangName() string { return "intf-discriminators" }

// Sbfd_LocalDiscriminator_DynamicDiscriminators
// Configure local discriminator dynamically
type Sbfd_LocalDiscriminator_DynamicDiscriminators struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local discriminator value. The type is slice of
    // Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator.
    DynamicDiscriminator []Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator
}

func (dynamicDiscriminators *Sbfd_LocalDiscriminator_DynamicDiscriminators) GetFilter() yfilter.YFilter { return dynamicDiscriminators.YFilter }

func (dynamicDiscriminators *Sbfd_LocalDiscriminator_DynamicDiscriminators) SetFilter(yf yfilter.YFilter) { dynamicDiscriminators.YFilter = yf }

func (dynamicDiscriminators *Sbfd_LocalDiscriminator_DynamicDiscriminators) GetGoName(yname string) string {
    if yname == "dynamic-discriminator" { return "DynamicDiscriminator" }
    return ""
}

func (dynamicDiscriminators *Sbfd_LocalDiscriminator_DynamicDiscriminators) GetSegmentPath() string {
    return "dynamic-discriminators"
}

func (dynamicDiscriminators *Sbfd_LocalDiscriminator_DynamicDiscriminators) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "dynamic-discriminator" {
        for _, c := range dynamicDiscriminators.DynamicDiscriminator {
            if dynamicDiscriminators.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator{}
        dynamicDiscriminators.DynamicDiscriminator = append(dynamicDiscriminators.DynamicDiscriminator, child)
        return &dynamicDiscriminators.DynamicDiscriminator[len(dynamicDiscriminators.DynamicDiscriminator)-1]
    }
    return nil
}

func (dynamicDiscriminators *Sbfd_LocalDiscriminator_DynamicDiscriminators) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range dynamicDiscriminators.DynamicDiscriminator {
        children[dynamicDiscriminators.DynamicDiscriminator[i].GetSegmentPath()] = &dynamicDiscriminators.DynamicDiscriminator[i]
    }
    return children
}

func (dynamicDiscriminators *Sbfd_LocalDiscriminator_DynamicDiscriminators) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (dynamicDiscriminators *Sbfd_LocalDiscriminator_DynamicDiscriminators) GetBundleName() string { return "cisco_ios_xr" }

func (dynamicDiscriminators *Sbfd_LocalDiscriminator_DynamicDiscriminators) GetYangName() string { return "dynamic-discriminators" }

func (dynamicDiscriminators *Sbfd_LocalDiscriminator_DynamicDiscriminators) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dynamicDiscriminators *Sbfd_LocalDiscriminator_DynamicDiscriminators) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dynamicDiscriminators *Sbfd_LocalDiscriminator_DynamicDiscriminators) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dynamicDiscriminators *Sbfd_LocalDiscriminator_DynamicDiscriminators) SetParent(parent types.Entity) { dynamicDiscriminators.parent = parent }

func (dynamicDiscriminators *Sbfd_LocalDiscriminator_DynamicDiscriminators) GetParent() types.Entity { return dynamicDiscriminators.parent }

func (dynamicDiscriminators *Sbfd_LocalDiscriminator_DynamicDiscriminators) GetParentYangName() string { return "local-discriminator" }

// Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator
// Local discriminator value
type Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Dynamic discriminator value. The type is
    // interface{} with range: 0..1.
    Discriminator interface{}
}

func (dynamicDiscriminator *Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator) GetFilter() yfilter.YFilter { return dynamicDiscriminator.YFilter }

func (dynamicDiscriminator *Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator) SetFilter(yf yfilter.YFilter) { dynamicDiscriminator.YFilter = yf }

func (dynamicDiscriminator *Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator) GetGoName(yname string) string {
    if yname == "discriminator" { return "Discriminator" }
    return ""
}

func (dynamicDiscriminator *Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator) GetSegmentPath() string {
    return "dynamic-discriminator" + "[discriminator='" + fmt.Sprintf("%v", dynamicDiscriminator.Discriminator) + "']"
}

func (dynamicDiscriminator *Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (dynamicDiscriminator *Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (dynamicDiscriminator *Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["discriminator"] = dynamicDiscriminator.Discriminator
    return leafs
}

func (dynamicDiscriminator *Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator) GetBundleName() string { return "cisco_ios_xr" }

func (dynamicDiscriminator *Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator) GetYangName() string { return "dynamic-discriminator" }

func (dynamicDiscriminator *Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (dynamicDiscriminator *Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (dynamicDiscriminator *Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (dynamicDiscriminator *Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator) SetParent(parent types.Entity) { dynamicDiscriminator.parent = parent }

func (dynamicDiscriminator *Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator) GetParent() types.Entity { return dynamicDiscriminator.parent }

func (dynamicDiscriminator *Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator) GetParentYangName() string { return "dynamic-discriminators" }

// Sbfd_LocalDiscriminator_Ipv4Discriminators
// Configure local discriminator as an ipv4
// address
type Sbfd_LocalDiscriminator_Ipv4Discriminators struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // ipv4 address as discriminator. The type is slice of
    // Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator.
    Ipv4Discriminator []Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator
}

func (ipv4Discriminators *Sbfd_LocalDiscriminator_Ipv4Discriminators) GetFilter() yfilter.YFilter { return ipv4Discriminators.YFilter }

func (ipv4Discriminators *Sbfd_LocalDiscriminator_Ipv4Discriminators) SetFilter(yf yfilter.YFilter) { ipv4Discriminators.YFilter = yf }

func (ipv4Discriminators *Sbfd_LocalDiscriminator_Ipv4Discriminators) GetGoName(yname string) string {
    if yname == "ipv4-discriminator" { return "Ipv4Discriminator" }
    return ""
}

func (ipv4Discriminators *Sbfd_LocalDiscriminator_Ipv4Discriminators) GetSegmentPath() string {
    return "ipv4-discriminators"
}

func (ipv4Discriminators *Sbfd_LocalDiscriminator_Ipv4Discriminators) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv4-discriminator" {
        for _, c := range ipv4Discriminators.Ipv4Discriminator {
            if ipv4Discriminators.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator{}
        ipv4Discriminators.Ipv4Discriminator = append(ipv4Discriminators.Ipv4Discriminator, child)
        return &ipv4Discriminators.Ipv4Discriminator[len(ipv4Discriminators.Ipv4Discriminator)-1]
    }
    return nil
}

func (ipv4Discriminators *Sbfd_LocalDiscriminator_Ipv4Discriminators) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipv4Discriminators.Ipv4Discriminator {
        children[ipv4Discriminators.Ipv4Discriminator[i].GetSegmentPath()] = &ipv4Discriminators.Ipv4Discriminator[i]
    }
    return children
}

func (ipv4Discriminators *Sbfd_LocalDiscriminator_Ipv4Discriminators) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4Discriminators *Sbfd_LocalDiscriminator_Ipv4Discriminators) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Discriminators *Sbfd_LocalDiscriminator_Ipv4Discriminators) GetYangName() string { return "ipv4-discriminators" }

func (ipv4Discriminators *Sbfd_LocalDiscriminator_Ipv4Discriminators) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Discriminators *Sbfd_LocalDiscriminator_Ipv4Discriminators) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Discriminators *Sbfd_LocalDiscriminator_Ipv4Discriminators) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Discriminators *Sbfd_LocalDiscriminator_Ipv4Discriminators) SetParent(parent types.Entity) { ipv4Discriminators.parent = parent }

func (ipv4Discriminators *Sbfd_LocalDiscriminator_Ipv4Discriminators) GetParent() types.Entity { return ipv4Discriminators.parent }

func (ipv4Discriminators *Sbfd_LocalDiscriminator_Ipv4Discriminators) GetParentYangName() string { return "local-discriminator" }

// Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator
// ipv4 address as discriminator
type Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key.  IPv4 address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (ipv4Discriminator *Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator) GetFilter() yfilter.YFilter { return ipv4Discriminator.YFilter }

func (ipv4Discriminator *Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator) SetFilter(yf yfilter.YFilter) { ipv4Discriminator.YFilter = yf }

func (ipv4Discriminator *Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (ipv4Discriminator *Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator) GetSegmentPath() string {
    return "ipv4-discriminator" + "[address='" + fmt.Sprintf("%v", ipv4Discriminator.Address) + "']"
}

func (ipv4Discriminator *Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv4Discriminator *Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv4Discriminator *Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = ipv4Discriminator.Address
    return leafs
}

func (ipv4Discriminator *Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4Discriminator *Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator) GetYangName() string { return "ipv4-discriminator" }

func (ipv4Discriminator *Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4Discriminator *Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4Discriminator *Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4Discriminator *Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator) SetParent(parent types.Entity) { ipv4Discriminator.parent = parent }

func (ipv4Discriminator *Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator) GetParent() types.Entity { return ipv4Discriminator.parent }

func (ipv4Discriminator *Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator) GetParentYangName() string { return "ipv4-discriminators" }

// Sbfd_LocalDiscriminator_Val32Discriminators
// Configure local discriminator as an integer
type Sbfd_LocalDiscriminator_Val32Discriminators struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Local discriminator value. The type is slice of
    // Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator.
    Val32Discriminator []Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator
}

func (val32Discriminators *Sbfd_LocalDiscriminator_Val32Discriminators) GetFilter() yfilter.YFilter { return val32Discriminators.YFilter }

func (val32Discriminators *Sbfd_LocalDiscriminator_Val32Discriminators) SetFilter(yf yfilter.YFilter) { val32Discriminators.YFilter = yf }

func (val32Discriminators *Sbfd_LocalDiscriminator_Val32Discriminators) GetGoName(yname string) string {
    if yname == "val32-discriminator" { return "Val32Discriminator" }
    return ""
}

func (val32Discriminators *Sbfd_LocalDiscriminator_Val32Discriminators) GetSegmentPath() string {
    return "val32-discriminators"
}

func (val32Discriminators *Sbfd_LocalDiscriminator_Val32Discriminators) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "val32-discriminator" {
        for _, c := range val32Discriminators.Val32Discriminator {
            if val32Discriminators.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator{}
        val32Discriminators.Val32Discriminator = append(val32Discriminators.Val32Discriminator, child)
        return &val32Discriminators.Val32Discriminator[len(val32Discriminators.Val32Discriminator)-1]
    }
    return nil
}

func (val32Discriminators *Sbfd_LocalDiscriminator_Val32Discriminators) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range val32Discriminators.Val32Discriminator {
        children[val32Discriminators.Val32Discriminator[i].GetSegmentPath()] = &val32Discriminators.Val32Discriminator[i]
    }
    return children
}

func (val32Discriminators *Sbfd_LocalDiscriminator_Val32Discriminators) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (val32Discriminators *Sbfd_LocalDiscriminator_Val32Discriminators) GetBundleName() string { return "cisco_ios_xr" }

func (val32Discriminators *Sbfd_LocalDiscriminator_Val32Discriminators) GetYangName() string { return "val32-discriminators" }

func (val32Discriminators *Sbfd_LocalDiscriminator_Val32Discriminators) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (val32Discriminators *Sbfd_LocalDiscriminator_Val32Discriminators) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (val32Discriminators *Sbfd_LocalDiscriminator_Val32Discriminators) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (val32Discriminators *Sbfd_LocalDiscriminator_Val32Discriminators) SetParent(parent types.Entity) { val32Discriminators.parent = parent }

func (val32Discriminators *Sbfd_LocalDiscriminator_Val32Discriminators) GetParent() types.Entity { return val32Discriminators.parent }

func (val32Discriminators *Sbfd_LocalDiscriminator_Val32Discriminators) GetParentYangName() string { return "local-discriminator" }

// Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator
// Local discriminator value
type Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Local discriminator value. The type is interface{}
    // with range: 1..4294967295.
    Discriminator interface{}
}

func (val32Discriminator *Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator) GetFilter() yfilter.YFilter { return val32Discriminator.YFilter }

func (val32Discriminator *Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator) SetFilter(yf yfilter.YFilter) { val32Discriminator.YFilter = yf }

func (val32Discriminator *Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator) GetGoName(yname string) string {
    if yname == "discriminator" { return "Discriminator" }
    return ""
}

func (val32Discriminator *Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator) GetSegmentPath() string {
    return "val32-discriminator" + "[discriminator='" + fmt.Sprintf("%v", val32Discriminator.Discriminator) + "']"
}

func (val32Discriminator *Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (val32Discriminator *Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (val32Discriminator *Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["discriminator"] = val32Discriminator.Discriminator
    return leafs
}

func (val32Discriminator *Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator) GetBundleName() string { return "cisco_ios_xr" }

func (val32Discriminator *Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator) GetYangName() string { return "val32-discriminator" }

func (val32Discriminator *Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (val32Discriminator *Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (val32Discriminator *Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (val32Discriminator *Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator) SetParent(parent types.Entity) { val32Discriminator.parent = parent }

func (val32Discriminator *Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator) GetParent() types.Entity { return val32Discriminator.parent }

func (val32Discriminator *Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator) GetParentYangName() string { return "val32-discriminators" }

