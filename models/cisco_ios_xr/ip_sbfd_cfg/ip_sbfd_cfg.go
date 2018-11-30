// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-sbfd package configuration.
// 
// This module contains definitions
// for the following management objects:
//   sbfd: SBFD Configuration ,Seamless-BFD is method for detecting
//     faultsbetween two different nodes in a network
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // configure remote target.
    RemoteTarget Sbfd_RemoteTarget

    // Configure local discriminator.
    LocalDiscriminator Sbfd_LocalDiscriminator
}

func (sbfd *Sbfd) GetEntityData() *types.CommonEntityData {
    sbfd.EntityData.YFilter = sbfd.YFilter
    sbfd.EntityData.YangName = "sbfd"
    sbfd.EntityData.BundleName = "cisco_ios_xr"
    sbfd.EntityData.ParentYangName = "Cisco-IOS-XR-ip-sbfd-cfg"
    sbfd.EntityData.SegmentPath = "Cisco-IOS-XR-ip-sbfd-cfg:sbfd"
    sbfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sbfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sbfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sbfd.EntityData.Children = types.NewOrderedMap()
    sbfd.EntityData.Children.Append("remote-target", types.YChild{"RemoteTarget", &sbfd.RemoteTarget})
    sbfd.EntityData.Children.Append("local-discriminator", types.YChild{"LocalDiscriminator", &sbfd.LocalDiscriminator})
    sbfd.EntityData.Leafs = types.NewOrderedMap()

    sbfd.EntityData.YListKeys = []string {}

    return &(sbfd.EntityData)
}

// Sbfd_RemoteTarget
// configure remote target
type Sbfd_RemoteTarget struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv4 address as target.
    Ipv4Addresses Sbfd_RemoteTarget_Ipv4Addresses

    // ipv6 address as target.
    Ipv6Addresses Sbfd_RemoteTarget_Ipv6Addresses
}

func (remoteTarget *Sbfd_RemoteTarget) GetEntityData() *types.CommonEntityData {
    remoteTarget.EntityData.YFilter = remoteTarget.YFilter
    remoteTarget.EntityData.YangName = "remote-target"
    remoteTarget.EntityData.BundleName = "cisco_ios_xr"
    remoteTarget.EntityData.ParentYangName = "sbfd"
    remoteTarget.EntityData.SegmentPath = "remote-target"
    remoteTarget.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteTarget.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteTarget.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteTarget.EntityData.Children = types.NewOrderedMap()
    remoteTarget.EntityData.Children.Append("ipv4-addresses", types.YChild{"Ipv4Addresses", &remoteTarget.Ipv4Addresses})
    remoteTarget.EntityData.Children.Append("ipv6-addresses", types.YChild{"Ipv6Addresses", &remoteTarget.Ipv6Addresses})
    remoteTarget.EntityData.Leafs = types.NewOrderedMap()

    remoteTarget.EntityData.YListKeys = []string {}

    return &(remoteTarget.EntityData)
}

// Sbfd_RemoteTarget_Ipv4Addresses
// ipv4 address as target
type Sbfd_RemoteTarget_Ipv4Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP Address Value for RemoteDiscriminatorTable. The type is slice of
    // Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address.
    Ipv4Address []*Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address
}

func (ipv4Addresses *Sbfd_RemoteTarget_Ipv4Addresses) GetEntityData() *types.CommonEntityData {
    ipv4Addresses.EntityData.YFilter = ipv4Addresses.YFilter
    ipv4Addresses.EntityData.YangName = "ipv4-addresses"
    ipv4Addresses.EntityData.BundleName = "cisco_ios_xr"
    ipv4Addresses.EntityData.ParentYangName = "remote-target"
    ipv4Addresses.EntityData.SegmentPath = "ipv4-addresses"
    ipv4Addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Addresses.EntityData.Children = types.NewOrderedMap()
    ipv4Addresses.EntityData.Children.Append("ipv4-address", types.YChild{"Ipv4Address", nil})
    for i := range ipv4Addresses.Ipv4Address {
        ipv4Addresses.EntityData.Children.Append(types.GetSegmentPath(ipv4Addresses.Ipv4Address[i]), types.YChild{"Ipv4Address", ipv4Addresses.Ipv4Address[i]})
    }
    ipv4Addresses.EntityData.Leafs = types.NewOrderedMap()

    ipv4Addresses.EntityData.YListKeys = []string {}

    return &(ipv4Addresses.EntityData)
}

// Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address
// IP Address Value for RemoteDiscriminatorTable
type Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key.  IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Remote Discriminator value. The type is slice of
    // Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator.
    RemoteDiscriminator []*Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator
}

func (ipv4Address *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address) GetEntityData() *types.CommonEntityData {
    ipv4Address.EntityData.YFilter = ipv4Address.YFilter
    ipv4Address.EntityData.YangName = "ipv4-address"
    ipv4Address.EntityData.BundleName = "cisco_ios_xr"
    ipv4Address.EntityData.ParentYangName = "ipv4-addresses"
    ipv4Address.EntityData.SegmentPath = "ipv4-address" + types.AddKeyToken(ipv4Address.Address, "address")
    ipv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Address.EntityData.Children = types.NewOrderedMap()
    ipv4Address.EntityData.Children.Append("remote-discriminator", types.YChild{"RemoteDiscriminator", nil})
    for i := range ipv4Address.RemoteDiscriminator {
        ipv4Address.EntityData.Children.Append(types.GetSegmentPath(ipv4Address.RemoteDiscriminator[i]), types.YChild{"RemoteDiscriminator", ipv4Address.RemoteDiscriminator[i]})
    }
    ipv4Address.EntityData.Leafs = types.NewOrderedMap()
    ipv4Address.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4Address.Address})

    ipv4Address.EntityData.YListKeys = []string {"Address"}

    return &(ipv4Address.EntityData)
}

// Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator
// Remote Discriminator value
type Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Remote Discriminator Value. The type is
    // interface{} with range: 1..4294967295.
    RemoteDiscriminator interface{}
}

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv4Addresses_Ipv4Address_RemoteDiscriminator) GetEntityData() *types.CommonEntityData {
    remoteDiscriminator.EntityData.YFilter = remoteDiscriminator.YFilter
    remoteDiscriminator.EntityData.YangName = "remote-discriminator"
    remoteDiscriminator.EntityData.BundleName = "cisco_ios_xr"
    remoteDiscriminator.EntityData.ParentYangName = "ipv4-address"
    remoteDiscriminator.EntityData.SegmentPath = "remote-discriminator" + types.AddKeyToken(remoteDiscriminator.RemoteDiscriminator, "remote-discriminator")
    remoteDiscriminator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteDiscriminator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteDiscriminator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteDiscriminator.EntityData.Children = types.NewOrderedMap()
    remoteDiscriminator.EntityData.Leafs = types.NewOrderedMap()
    remoteDiscriminator.EntityData.Leafs.Append("remote-discriminator", types.YLeaf{"RemoteDiscriminator", remoteDiscriminator.RemoteDiscriminator})

    remoteDiscriminator.EntityData.YListKeys = []string {"RemoteDiscriminator"}

    return &(remoteDiscriminator.EntityData)
}

// Sbfd_RemoteTarget_Ipv6Addresses
// ipv6 address as target
type Sbfd_RemoteTarget_Ipv6Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP Address Value for RemoteDiscriminatorTable. The type is slice of
    // Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address.
    Ipv6Address []*Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address
}

func (ipv6Addresses *Sbfd_RemoteTarget_Ipv6Addresses) GetEntityData() *types.CommonEntityData {
    ipv6Addresses.EntityData.YFilter = ipv6Addresses.YFilter
    ipv6Addresses.EntityData.YangName = "ipv6-addresses"
    ipv6Addresses.EntityData.BundleName = "cisco_ios_xr"
    ipv6Addresses.EntityData.ParentYangName = "remote-target"
    ipv6Addresses.EntityData.SegmentPath = "ipv6-addresses"
    ipv6Addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Addresses.EntityData.Children = types.NewOrderedMap()
    ipv6Addresses.EntityData.Children.Append("ipv6-address", types.YChild{"Ipv6Address", nil})
    for i := range ipv6Addresses.Ipv6Address {
        ipv6Addresses.EntityData.Children.Append(types.GetSegmentPath(ipv6Addresses.Ipv6Address[i]), types.YChild{"Ipv6Address", ipv6Addresses.Ipv6Address[i]})
    }
    ipv6Addresses.EntityData.Leafs = types.NewOrderedMap()

    ipv6Addresses.EntityData.YListKeys = []string {}

    return &(ipv6Addresses.EntityData)
}

// Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address
// IP Address Value for RemoteDiscriminatorTable
type Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key.  IPv6 adddress. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Remote Discriminator value. The type is slice of
    // Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator.
    RemoteDiscriminator []*Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator
}

func (ipv6Address *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address) GetEntityData() *types.CommonEntityData {
    ipv6Address.EntityData.YFilter = ipv6Address.YFilter
    ipv6Address.EntityData.YangName = "ipv6-address"
    ipv6Address.EntityData.BundleName = "cisco_ios_xr"
    ipv6Address.EntityData.ParentYangName = "ipv6-addresses"
    ipv6Address.EntityData.SegmentPath = "ipv6-address" + types.AddKeyToken(ipv6Address.Address, "address")
    ipv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Address.EntityData.Children = types.NewOrderedMap()
    ipv6Address.EntityData.Children.Append("remote-discriminator", types.YChild{"RemoteDiscriminator", nil})
    for i := range ipv6Address.RemoteDiscriminator {
        ipv6Address.EntityData.Children.Append(types.GetSegmentPath(ipv6Address.RemoteDiscriminator[i]), types.YChild{"RemoteDiscriminator", ipv6Address.RemoteDiscriminator[i]})
    }
    ipv6Address.EntityData.Leafs = types.NewOrderedMap()
    ipv6Address.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv6Address.Address})

    ipv6Address.EntityData.YListKeys = []string {"Address"}

    return &(ipv6Address.EntityData)
}

// Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator
// Remote Discriminator value
type Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Remote Discriminator Value. The type is
    // interface{} with range: 1..4294967295.
    RemoteDiscriminator interface{}
}

func (remoteDiscriminator *Sbfd_RemoteTarget_Ipv6Addresses_Ipv6Address_RemoteDiscriminator) GetEntityData() *types.CommonEntityData {
    remoteDiscriminator.EntityData.YFilter = remoteDiscriminator.YFilter
    remoteDiscriminator.EntityData.YangName = "remote-discriminator"
    remoteDiscriminator.EntityData.BundleName = "cisco_ios_xr"
    remoteDiscriminator.EntityData.ParentYangName = "ipv6-address"
    remoteDiscriminator.EntityData.SegmentPath = "remote-discriminator" + types.AddKeyToken(remoteDiscriminator.RemoteDiscriminator, "remote-discriminator")
    remoteDiscriminator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteDiscriminator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteDiscriminator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteDiscriminator.EntityData.Children = types.NewOrderedMap()
    remoteDiscriminator.EntityData.Leafs = types.NewOrderedMap()
    remoteDiscriminator.EntityData.Leafs.Append("remote-discriminator", types.YLeaf{"RemoteDiscriminator", remoteDiscriminator.RemoteDiscriminator})

    remoteDiscriminator.EntityData.YListKeys = []string {"RemoteDiscriminator"}

    return &(remoteDiscriminator.EntityData)
}

// Sbfd_LocalDiscriminator
// Configure local discriminator
type Sbfd_LocalDiscriminator struct {
    EntityData types.CommonEntityData
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

func (localDiscriminator *Sbfd_LocalDiscriminator) GetEntityData() *types.CommonEntityData {
    localDiscriminator.EntityData.YFilter = localDiscriminator.YFilter
    localDiscriminator.EntityData.YangName = "local-discriminator"
    localDiscriminator.EntityData.BundleName = "cisco_ios_xr"
    localDiscriminator.EntityData.ParentYangName = "sbfd"
    localDiscriminator.EntityData.SegmentPath = "local-discriminator"
    localDiscriminator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localDiscriminator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localDiscriminator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localDiscriminator.EntityData.Children = types.NewOrderedMap()
    localDiscriminator.EntityData.Children.Append("intf-discriminators", types.YChild{"IntfDiscriminators", &localDiscriminator.IntfDiscriminators})
    localDiscriminator.EntityData.Children.Append("dynamic-discriminators", types.YChild{"DynamicDiscriminators", &localDiscriminator.DynamicDiscriminators})
    localDiscriminator.EntityData.Children.Append("ipv4-discriminators", types.YChild{"Ipv4Discriminators", &localDiscriminator.Ipv4Discriminators})
    localDiscriminator.EntityData.Children.Append("val32-discriminators", types.YChild{"Val32Discriminators", &localDiscriminator.Val32Discriminators})
    localDiscriminator.EntityData.Leafs = types.NewOrderedMap()

    localDiscriminator.EntityData.YListKeys = []string {}

    return &(localDiscriminator.EntityData)
}

// Sbfd_LocalDiscriminator_IntfDiscriminators
// Configure local discriminator from interface
// address
type Sbfd_LocalDiscriminator_IntfDiscriminators struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // interface address as discriminator. The type is slice of
    // Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator.
    IntfDiscriminator []*Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator
}

func (intfDiscriminators *Sbfd_LocalDiscriminator_IntfDiscriminators) GetEntityData() *types.CommonEntityData {
    intfDiscriminators.EntityData.YFilter = intfDiscriminators.YFilter
    intfDiscriminators.EntityData.YangName = "intf-discriminators"
    intfDiscriminators.EntityData.BundleName = "cisco_ios_xr"
    intfDiscriminators.EntityData.ParentYangName = "local-discriminator"
    intfDiscriminators.EntityData.SegmentPath = "intf-discriminators"
    intfDiscriminators.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    intfDiscriminators.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    intfDiscriminators.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    intfDiscriminators.EntityData.Children = types.NewOrderedMap()
    intfDiscriminators.EntityData.Children.Append("intf-discriminator", types.YChild{"IntfDiscriminator", nil})
    for i := range intfDiscriminators.IntfDiscriminator {
        intfDiscriminators.EntityData.Children.Append(types.GetSegmentPath(intfDiscriminators.IntfDiscriminator[i]), types.YChild{"IntfDiscriminator", intfDiscriminators.IntfDiscriminator[i]})
    }
    intfDiscriminators.EntityData.Leafs = types.NewOrderedMap()

    intfDiscriminators.EntityData.YListKeys = []string {}

    return &(intfDiscriminators.EntityData)
}

// Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator
// interface address as discriminator
type Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}
}

func (intfDiscriminator *Sbfd_LocalDiscriminator_IntfDiscriminators_IntfDiscriminator) GetEntityData() *types.CommonEntityData {
    intfDiscriminator.EntityData.YFilter = intfDiscriminator.YFilter
    intfDiscriminator.EntityData.YangName = "intf-discriminator"
    intfDiscriminator.EntityData.BundleName = "cisco_ios_xr"
    intfDiscriminator.EntityData.ParentYangName = "intf-discriminators"
    intfDiscriminator.EntityData.SegmentPath = "intf-discriminator" + types.AddKeyToken(intfDiscriminator.InterfaceName, "interface-name")
    intfDiscriminator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    intfDiscriminator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    intfDiscriminator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    intfDiscriminator.EntityData.Children = types.NewOrderedMap()
    intfDiscriminator.EntityData.Leafs = types.NewOrderedMap()
    intfDiscriminator.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", intfDiscriminator.InterfaceName})

    intfDiscriminator.EntityData.YListKeys = []string {"InterfaceName"}

    return &(intfDiscriminator.EntityData)
}

// Sbfd_LocalDiscriminator_DynamicDiscriminators
// Configure local discriminator dynamically
type Sbfd_LocalDiscriminator_DynamicDiscriminators struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local discriminator value. The type is slice of
    // Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator.
    DynamicDiscriminator []*Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator
}

func (dynamicDiscriminators *Sbfd_LocalDiscriminator_DynamicDiscriminators) GetEntityData() *types.CommonEntityData {
    dynamicDiscriminators.EntityData.YFilter = dynamicDiscriminators.YFilter
    dynamicDiscriminators.EntityData.YangName = "dynamic-discriminators"
    dynamicDiscriminators.EntityData.BundleName = "cisco_ios_xr"
    dynamicDiscriminators.EntityData.ParentYangName = "local-discriminator"
    dynamicDiscriminators.EntityData.SegmentPath = "dynamic-discriminators"
    dynamicDiscriminators.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamicDiscriminators.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamicDiscriminators.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamicDiscriminators.EntityData.Children = types.NewOrderedMap()
    dynamicDiscriminators.EntityData.Children.Append("dynamic-discriminator", types.YChild{"DynamicDiscriminator", nil})
    for i := range dynamicDiscriminators.DynamicDiscriminator {
        dynamicDiscriminators.EntityData.Children.Append(types.GetSegmentPath(dynamicDiscriminators.DynamicDiscriminator[i]), types.YChild{"DynamicDiscriminator", dynamicDiscriminators.DynamicDiscriminator[i]})
    }
    dynamicDiscriminators.EntityData.Leafs = types.NewOrderedMap()

    dynamicDiscriminators.EntityData.YListKeys = []string {}

    return &(dynamicDiscriminators.EntityData)
}

// Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator
// Local discriminator value
type Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Dynamic discriminator value. The type is
    // interface{} with range: 0..1.
    Discriminator interface{}
}

func (dynamicDiscriminator *Sbfd_LocalDiscriminator_DynamicDiscriminators_DynamicDiscriminator) GetEntityData() *types.CommonEntityData {
    dynamicDiscriminator.EntityData.YFilter = dynamicDiscriminator.YFilter
    dynamicDiscriminator.EntityData.YangName = "dynamic-discriminator"
    dynamicDiscriminator.EntityData.BundleName = "cisco_ios_xr"
    dynamicDiscriminator.EntityData.ParentYangName = "dynamic-discriminators"
    dynamicDiscriminator.EntityData.SegmentPath = "dynamic-discriminator" + types.AddKeyToken(dynamicDiscriminator.Discriminator, "discriminator")
    dynamicDiscriminator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dynamicDiscriminator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dynamicDiscriminator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dynamicDiscriminator.EntityData.Children = types.NewOrderedMap()
    dynamicDiscriminator.EntityData.Leafs = types.NewOrderedMap()
    dynamicDiscriminator.EntityData.Leafs.Append("discriminator", types.YLeaf{"Discriminator", dynamicDiscriminator.Discriminator})

    dynamicDiscriminator.EntityData.YListKeys = []string {"Discriminator"}

    return &(dynamicDiscriminator.EntityData)
}

// Sbfd_LocalDiscriminator_Ipv4Discriminators
// Configure local discriminator as an ipv4
// address
type Sbfd_LocalDiscriminator_Ipv4Discriminators struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ipv4 address as discriminator. The type is slice of
    // Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator.
    Ipv4Discriminator []*Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator
}

func (ipv4Discriminators *Sbfd_LocalDiscriminator_Ipv4Discriminators) GetEntityData() *types.CommonEntityData {
    ipv4Discriminators.EntityData.YFilter = ipv4Discriminators.YFilter
    ipv4Discriminators.EntityData.YangName = "ipv4-discriminators"
    ipv4Discriminators.EntityData.BundleName = "cisco_ios_xr"
    ipv4Discriminators.EntityData.ParentYangName = "local-discriminator"
    ipv4Discriminators.EntityData.SegmentPath = "ipv4-discriminators"
    ipv4Discriminators.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Discriminators.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Discriminators.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Discriminators.EntityData.Children = types.NewOrderedMap()
    ipv4Discriminators.EntityData.Children.Append("ipv4-discriminator", types.YChild{"Ipv4Discriminator", nil})
    for i := range ipv4Discriminators.Ipv4Discriminator {
        ipv4Discriminators.EntityData.Children.Append(types.GetSegmentPath(ipv4Discriminators.Ipv4Discriminator[i]), types.YChild{"Ipv4Discriminator", ipv4Discriminators.Ipv4Discriminator[i]})
    }
    ipv4Discriminators.EntityData.Leafs = types.NewOrderedMap()

    ipv4Discriminators.EntityData.YListKeys = []string {}

    return &(ipv4Discriminators.EntityData)
}

// Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator
// ipv4 address as discriminator
type Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key.  IPv4 address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (ipv4Discriminator *Sbfd_LocalDiscriminator_Ipv4Discriminators_Ipv4Discriminator) GetEntityData() *types.CommonEntityData {
    ipv4Discriminator.EntityData.YFilter = ipv4Discriminator.YFilter
    ipv4Discriminator.EntityData.YangName = "ipv4-discriminator"
    ipv4Discriminator.EntityData.BundleName = "cisco_ios_xr"
    ipv4Discriminator.EntityData.ParentYangName = "ipv4-discriminators"
    ipv4Discriminator.EntityData.SegmentPath = "ipv4-discriminator" + types.AddKeyToken(ipv4Discriminator.Address, "address")
    ipv4Discriminator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4Discriminator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4Discriminator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4Discriminator.EntityData.Children = types.NewOrderedMap()
    ipv4Discriminator.EntityData.Leafs = types.NewOrderedMap()
    ipv4Discriminator.EntityData.Leafs.Append("address", types.YLeaf{"Address", ipv4Discriminator.Address})

    ipv4Discriminator.EntityData.YListKeys = []string {"Address"}

    return &(ipv4Discriminator.EntityData)
}

// Sbfd_LocalDiscriminator_Val32Discriminators
// Configure local discriminator as an integer
type Sbfd_LocalDiscriminator_Val32Discriminators struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local discriminator value. The type is slice of
    // Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator.
    Val32Discriminator []*Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator
}

func (val32Discriminators *Sbfd_LocalDiscriminator_Val32Discriminators) GetEntityData() *types.CommonEntityData {
    val32Discriminators.EntityData.YFilter = val32Discriminators.YFilter
    val32Discriminators.EntityData.YangName = "val32-discriminators"
    val32Discriminators.EntityData.BundleName = "cisco_ios_xr"
    val32Discriminators.EntityData.ParentYangName = "local-discriminator"
    val32Discriminators.EntityData.SegmentPath = "val32-discriminators"
    val32Discriminators.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    val32Discriminators.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    val32Discriminators.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    val32Discriminators.EntityData.Children = types.NewOrderedMap()
    val32Discriminators.EntityData.Children.Append("val32-discriminator", types.YChild{"Val32Discriminator", nil})
    for i := range val32Discriminators.Val32Discriminator {
        val32Discriminators.EntityData.Children.Append(types.GetSegmentPath(val32Discriminators.Val32Discriminator[i]), types.YChild{"Val32Discriminator", val32Discriminators.Val32Discriminator[i]})
    }
    val32Discriminators.EntityData.Leafs = types.NewOrderedMap()

    val32Discriminators.EntityData.YListKeys = []string {}

    return &(val32Discriminators.EntityData)
}

// Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator
// Local discriminator value
type Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Local discriminator value. The type is interface{}
    // with range: 1..4294967295.
    Discriminator interface{}
}

func (val32Discriminator *Sbfd_LocalDiscriminator_Val32Discriminators_Val32Discriminator) GetEntityData() *types.CommonEntityData {
    val32Discriminator.EntityData.YFilter = val32Discriminator.YFilter
    val32Discriminator.EntityData.YangName = "val32-discriminator"
    val32Discriminator.EntityData.BundleName = "cisco_ios_xr"
    val32Discriminator.EntityData.ParentYangName = "val32-discriminators"
    val32Discriminator.EntityData.SegmentPath = "val32-discriminator" + types.AddKeyToken(val32Discriminator.Discriminator, "discriminator")
    val32Discriminator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    val32Discriminator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    val32Discriminator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    val32Discriminator.EntityData.Children = types.NewOrderedMap()
    val32Discriminator.EntityData.Leafs = types.NewOrderedMap()
    val32Discriminator.EntityData.Leafs.Append("discriminator", types.YLeaf{"Discriminator", val32Discriminator.Discriminator})

    val32Discriminator.EntityData.YListKeys = []string {"Discriminator"}

    return &(val32Discriminator.EntityData)
}

