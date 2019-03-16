// This module contains a collection of YANG definitions
// for Cisco IOS-XR kim-tpa package configuration.
// 
// This module contains definitions
// for the following management objects:
//   tpa: tpa configuration commands
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package kim_tpa_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package kim_tpa_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-kim-tpa-cfg tpa}", reflect.TypeOf(Tpa{}))
    ydk.RegisterEntity("Cisco-IOS-XR-kim-tpa-cfg:tpa", reflect.TypeOf(Tpa{}))
}

// IpProtocol represents Ip protocol
type IpProtocol string

const (
    // Transmission Control Protocol, RFC 793
    IpProtocol_tcp IpProtocol = "tcp"

    // User Datagram Protocol, RFC 768
    IpProtocol_udp IpProtocol = "udp"
)

// Tpa
// tpa configuration commands
type Tpa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF container.
    VrfNames Tpa_VrfNames

    // Third party app logging.
    Logging Tpa_Logging

    // Statistics.
    Statistics Tpa_Statistics
}

func (tpa *Tpa) GetEntityData() *types.CommonEntityData {
    tpa.EntityData.YFilter = tpa.YFilter
    tpa.EntityData.YangName = "tpa"
    tpa.EntityData.BundleName = "cisco_ios_xr"
    tpa.EntityData.ParentYangName = "Cisco-IOS-XR-kim-tpa-cfg"
    tpa.EntityData.SegmentPath = "Cisco-IOS-XR-kim-tpa-cfg:tpa"
    tpa.EntityData.AbsolutePath = tpa.EntityData.SegmentPath
    tpa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tpa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tpa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tpa.EntityData.Children = types.NewOrderedMap()
    tpa.EntityData.Children.Append("vrf-names", types.YChild{"VrfNames", &tpa.VrfNames})
    tpa.EntityData.Children.Append("logging", types.YChild{"Logging", &tpa.Logging})
    tpa.EntityData.Children.Append("statistics", types.YChild{"Statistics", &tpa.Statistics})
    tpa.EntityData.Leafs = types.NewOrderedMap()

    tpa.EntityData.YListKeys = []string {}

    return &(tpa.EntityData)
}

// Tpa_VrfNames
// VRF container
type Tpa_VrfNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF name. The type is slice of Tpa_VrfNames_VrfName.
    VrfName []*Tpa_VrfNames_VrfName
}

func (vrfNames *Tpa_VrfNames) GetEntityData() *types.CommonEntityData {
    vrfNames.EntityData.YFilter = vrfNames.YFilter
    vrfNames.EntityData.YangName = "vrf-names"
    vrfNames.EntityData.BundleName = "cisco_ios_xr"
    vrfNames.EntityData.ParentYangName = "tpa"
    vrfNames.EntityData.SegmentPath = "vrf-names"
    vrfNames.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/" + vrfNames.EntityData.SegmentPath
    vrfNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfNames.EntityData.Children = types.NewOrderedMap()
    vrfNames.EntityData.Children.Append("vrf-name", types.YChild{"VrfName", nil})
    for i := range vrfNames.VrfName {
        vrfNames.EntityData.Children.Append(types.GetSegmentPath(vrfNames.VrfName[i]), types.YChild{"VrfName", vrfNames.VrfName[i]})
    }
    vrfNames.EntityData.Leafs = types.NewOrderedMap()

    vrfNames.EntityData.YListKeys = []string {}

    return &(vrfNames.EntityData)
}

// Tpa_VrfNames_VrfName
// VRF name
type Tpa_VrfNames_VrfName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // Disable routes and interfaces. The type is interface{}.
    Disable interface{}

    // EastWest container.
    EastWestNames Tpa_VrfNames_VrfName_EastWestNames

    // Address family.
    AddressFamily Tpa_VrfNames_VrfName_AddressFamily
}

func (vrfName *Tpa_VrfNames_VrfName) GetEntityData() *types.CommonEntityData {
    vrfName.EntityData.YFilter = vrfName.YFilter
    vrfName.EntityData.YangName = "vrf-name"
    vrfName.EntityData.BundleName = "cisco_ios_xr"
    vrfName.EntityData.ParentYangName = "vrf-names"
    vrfName.EntityData.SegmentPath = "vrf-name" + types.AddKeyToken(vrfName.VrfName, "vrf-name")
    vrfName.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/" + vrfName.EntityData.SegmentPath
    vrfName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfName.EntityData.Children = types.NewOrderedMap()
    vrfName.EntityData.Children.Append("east-west-names", types.YChild{"EastWestNames", &vrfName.EastWestNames})
    vrfName.EntityData.Children.Append("address-family", types.YChild{"AddressFamily", &vrfName.AddressFamily})
    vrfName.EntityData.Leafs = types.NewOrderedMap()
    vrfName.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrfName.VrfName})
    vrfName.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", vrfName.Disable})

    vrfName.EntityData.YListKeys = []string {"VrfName"}

    return &(vrfName.EntityData)
}

// Tpa_VrfNames_VrfName_EastWestNames
// EastWest container
type Tpa_VrfNames_VrfName_EastWestNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // East West interface. The type is slice of
    // Tpa_VrfNames_VrfName_EastWestNames_EastWestName.
    EastWestName []*Tpa_VrfNames_VrfName_EastWestNames_EastWestName
}

func (eastWestNames *Tpa_VrfNames_VrfName_EastWestNames) GetEntityData() *types.CommonEntityData {
    eastWestNames.EntityData.YFilter = eastWestNames.YFilter
    eastWestNames.EntityData.YangName = "east-west-names"
    eastWestNames.EntityData.BundleName = "cisco_ios_xr"
    eastWestNames.EntityData.ParentYangName = "vrf-name"
    eastWestNames.EntityData.SegmentPath = "east-west-names"
    eastWestNames.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/" + eastWestNames.EntityData.SegmentPath
    eastWestNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eastWestNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eastWestNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eastWestNames.EntityData.Children = types.NewOrderedMap()
    eastWestNames.EntityData.Children.Append("east-west-name", types.YChild{"EastWestName", nil})
    for i := range eastWestNames.EastWestName {
        eastWestNames.EntityData.Children.Append(types.GetSegmentPath(eastWestNames.EastWestName[i]), types.YChild{"EastWestName", eastWestNames.EastWestName[i]})
    }
    eastWestNames.EntityData.Leafs = types.NewOrderedMap()

    eastWestNames.EntityData.YListKeys = []string {}

    return &(eastWestNames.EntityData)
}

// Tpa_VrfNames_VrfName_EastWestNames_EastWestName
// East West interface
type Tpa_VrfNames_VrfName_EastWestNames_EastWestName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    EastWestName interface{}

    // VRF name. The type is string.
    Vrf interface{}

    // Interface. The type is string.
    Interface interface{}
}

func (eastWestName *Tpa_VrfNames_VrfName_EastWestNames_EastWestName) GetEntityData() *types.CommonEntityData {
    eastWestName.EntityData.YFilter = eastWestName.YFilter
    eastWestName.EntityData.YangName = "east-west-name"
    eastWestName.EntityData.BundleName = "cisco_ios_xr"
    eastWestName.EntityData.ParentYangName = "east-west-names"
    eastWestName.EntityData.SegmentPath = "east-west-name" + types.AddKeyToken(eastWestName.EastWestName, "east-west-name")
    eastWestName.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/east-west-names/" + eastWestName.EntityData.SegmentPath
    eastWestName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eastWestName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eastWestName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eastWestName.EntityData.Children = types.NewOrderedMap()
    eastWestName.EntityData.Leafs = types.NewOrderedMap()
    eastWestName.EntityData.Leafs.Append("east-west-name", types.YLeaf{"EastWestName", eastWestName.EastWestName})
    eastWestName.EntityData.Leafs.Append("vrf", types.YLeaf{"Vrf", eastWestName.Vrf})
    eastWestName.EntityData.Leafs.Append("interface", types.YLeaf{"Interface", eastWestName.Interface})

    eastWestName.EntityData.YListKeys = []string {"EastWestName"}

    return &(eastWestName.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily
// Address family
type Tpa_VrfNames_VrfName_AddressFamily struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 configuration.
    Ipv6 Tpa_VrfNames_VrfName_AddressFamily_Ipv6

    // IPv4 configuration.
    Ipv4 Tpa_VrfNames_VrfName_AddressFamily_Ipv4
}

func (addressFamily *Tpa_VrfNames_VrfName_AddressFamily) GetEntityData() *types.CommonEntityData {
    addressFamily.EntityData.YFilter = addressFamily.YFilter
    addressFamily.EntityData.YangName = "address-family"
    addressFamily.EntityData.BundleName = "cisco_ios_xr"
    addressFamily.EntityData.ParentYangName = "vrf-name"
    addressFamily.EntityData.SegmentPath = "address-family"
    addressFamily.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/" + addressFamily.EntityData.SegmentPath
    addressFamily.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressFamily.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressFamily.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressFamily.EntityData.Children = types.NewOrderedMap()
    addressFamily.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &addressFamily.Ipv6})
    addressFamily.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &addressFamily.Ipv4})
    addressFamily.EntityData.Leafs = types.NewOrderedMap()

    addressFamily.EntityData.YListKeys = []string {}

    return &(addressFamily.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6
// IPv6 configuration
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default interface used for routing. The type is string.
    DefaultRoute interface{}

    // Interface used for source address for egress interface.
    InterfaceNames Tpa_VrfNames_VrfName_AddressFamily_Ipv6_InterfaceNames

    // Traffic protection configuration.
    AllowEntries Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries

    // Interface used for Source Address.
    UpdateSource Tpa_VrfNames_VrfName_AddressFamily_Ipv6_UpdateSource
}

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "address-family"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/" + ipv6.EntityData.SegmentPath
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("interface-names", types.YChild{"InterfaceNames", &ipv6.InterfaceNames})
    ipv6.EntityData.Children.Append("allow-entries", types.YChild{"AllowEntries", &ipv6.AllowEntries})
    ipv6.EntityData.Children.Append("update-source", types.YChild{"UpdateSource", &ipv6.UpdateSource})
    ipv6.EntityData.Leafs = types.NewOrderedMap()
    ipv6.EntityData.Leafs.Append("default-route", types.YLeaf{"DefaultRoute", ipv6.DefaultRoute})

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6_InterfaceNames
// Interface used for source address for egress
// interface
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6_InterfaceNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Egress interface name. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv6_InterfaceNames_InterfaceName.
    InterfaceName []*Tpa_VrfNames_VrfName_AddressFamily_Ipv6_InterfaceNames_InterfaceName
}

func (interfaceNames *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_InterfaceNames) GetEntityData() *types.CommonEntityData {
    interfaceNames.EntityData.YFilter = interfaceNames.YFilter
    interfaceNames.EntityData.YangName = "interface-names"
    interfaceNames.EntityData.BundleName = "cisco_ios_xr"
    interfaceNames.EntityData.ParentYangName = "ipv6"
    interfaceNames.EntityData.SegmentPath = "interface-names"
    interfaceNames.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv6/" + interfaceNames.EntityData.SegmentPath
    interfaceNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceNames.EntityData.Children = types.NewOrderedMap()
    interfaceNames.EntityData.Children.Append("interface-name", types.YChild{"InterfaceName", nil})
    for i := range interfaceNames.InterfaceName {
        interfaceNames.EntityData.Children.Append(types.GetSegmentPath(interfaceNames.InterfaceName[i]), types.YChild{"InterfaceName", interfaceNames.InterfaceName[i]})
    }
    interfaceNames.EntityData.Leafs = types.NewOrderedMap()

    interfaceNames.EntityData.YListKeys = []string {}

    return &(interfaceNames.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6_InterfaceNames_InterfaceName
// Egress interface name
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6_InterfaceNames_InterfaceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    InterfaceName interface{}

    // Interface name for source address. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    EgressInterfaceSource interface{}
}

func (interfaceName *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_InterfaceNames_InterfaceName) GetEntityData() *types.CommonEntityData {
    interfaceName.EntityData.YFilter = interfaceName.YFilter
    interfaceName.EntityData.YangName = "interface-name"
    interfaceName.EntityData.BundleName = "cisco_ios_xr"
    interfaceName.EntityData.ParentYangName = "interface-names"
    interfaceName.EntityData.SegmentPath = "interface-name" + types.AddKeyToken(interfaceName.InterfaceName, "interface-name")
    interfaceName.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv6/interface-names/" + interfaceName.EntityData.SegmentPath
    interfaceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceName.EntityData.Children = types.NewOrderedMap()
    interfaceName.EntityData.Leafs = types.NewOrderedMap()
    interfaceName.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceName.InterfaceName})
    interfaceName.EntityData.Leafs.Append("egress-interface-source", types.YLeaf{"EgressInterfaceSource", interfaceName.EgressInterfaceSource})

    interfaceName.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceName.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries
// Traffic protection configuration
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allow traffic matching a filter. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntry.
    AllowEntry []*Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntry

    // Allow traffic matching a filter. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddress.
    AllowEntryLocalAddress []*Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddress

    // Allow traffic matching a filter. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryRemoteAddress.
    AllowEntryRemoteAddress []*Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryRemoteAddress

    // Allow traffic matching a filter. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddressRemoteAddress.
    AllowEntryLocalAddressRemoteAddress []*Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddressRemoteAddress

    // Allow traffic matching a filter. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryInterfaceName.
    AllowEntryInterfaceName []*Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryInterfaceName

    // Allow traffic matching a filter. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddressInterfaceName.
    AllowEntryLocalAddressInterfaceName []*Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddressInterfaceName

    // Allow traffic matching a filter. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryRemoteAddressInterfaceName.
    AllowEntryRemoteAddressInterfaceName []*Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryRemoteAddressInterfaceName

    // Allow traffic matching a filter. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddressRemoteAddressInterfaceName.
    AllowEntryLocalAddressRemoteAddressInterfaceName []*Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddressRemoteAddressInterfaceName
}

func (allowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries) GetEntityData() *types.CommonEntityData {
    allowEntries.EntityData.YFilter = allowEntries.YFilter
    allowEntries.EntityData.YangName = "allow-entries"
    allowEntries.EntityData.BundleName = "cisco_ios_xr"
    allowEntries.EntityData.ParentYangName = "ipv6"
    allowEntries.EntityData.SegmentPath = "allow-entries"
    allowEntries.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv6/" + allowEntries.EntityData.SegmentPath
    allowEntries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntries.EntityData.Children = types.NewOrderedMap()
    allowEntries.EntityData.Children.Append("allow-entry", types.YChild{"AllowEntry", nil})
    for i := range allowEntries.AllowEntry {
        allowEntries.EntityData.Children.Append(types.GetSegmentPath(allowEntries.AllowEntry[i]), types.YChild{"AllowEntry", allowEntries.AllowEntry[i]})
    }
    allowEntries.EntityData.Children.Append("allow-entry-local-address", types.YChild{"AllowEntryLocalAddress", nil})
    for i := range allowEntries.AllowEntryLocalAddress {
        allowEntries.EntityData.Children.Append(types.GetSegmentPath(allowEntries.AllowEntryLocalAddress[i]), types.YChild{"AllowEntryLocalAddress", allowEntries.AllowEntryLocalAddress[i]})
    }
    allowEntries.EntityData.Children.Append("allow-entry-remote-address", types.YChild{"AllowEntryRemoteAddress", nil})
    for i := range allowEntries.AllowEntryRemoteAddress {
        allowEntries.EntityData.Children.Append(types.GetSegmentPath(allowEntries.AllowEntryRemoteAddress[i]), types.YChild{"AllowEntryRemoteAddress", allowEntries.AllowEntryRemoteAddress[i]})
    }
    allowEntries.EntityData.Children.Append("allow-entry-local-address-remote-address", types.YChild{"AllowEntryLocalAddressRemoteAddress", nil})
    for i := range allowEntries.AllowEntryLocalAddressRemoteAddress {
        allowEntries.EntityData.Children.Append(types.GetSegmentPath(allowEntries.AllowEntryLocalAddressRemoteAddress[i]), types.YChild{"AllowEntryLocalAddressRemoteAddress", allowEntries.AllowEntryLocalAddressRemoteAddress[i]})
    }
    allowEntries.EntityData.Children.Append("allow-entry-interface-name", types.YChild{"AllowEntryInterfaceName", nil})
    for i := range allowEntries.AllowEntryInterfaceName {
        allowEntries.EntityData.Children.Append(types.GetSegmentPath(allowEntries.AllowEntryInterfaceName[i]), types.YChild{"AllowEntryInterfaceName", allowEntries.AllowEntryInterfaceName[i]})
    }
    allowEntries.EntityData.Children.Append("allow-entry-local-address-interface-name", types.YChild{"AllowEntryLocalAddressInterfaceName", nil})
    for i := range allowEntries.AllowEntryLocalAddressInterfaceName {
        allowEntries.EntityData.Children.Append(types.GetSegmentPath(allowEntries.AllowEntryLocalAddressInterfaceName[i]), types.YChild{"AllowEntryLocalAddressInterfaceName", allowEntries.AllowEntryLocalAddressInterfaceName[i]})
    }
    allowEntries.EntityData.Children.Append("allow-entry-remote-address-interface-name", types.YChild{"AllowEntryRemoteAddressInterfaceName", nil})
    for i := range allowEntries.AllowEntryRemoteAddressInterfaceName {
        allowEntries.EntityData.Children.Append(types.GetSegmentPath(allowEntries.AllowEntryRemoteAddressInterfaceName[i]), types.YChild{"AllowEntryRemoteAddressInterfaceName", allowEntries.AllowEntryRemoteAddressInterfaceName[i]})
    }
    allowEntries.EntityData.Children.Append("allow-entry-local-address-remote-address-interface-name", types.YChild{"AllowEntryLocalAddressRemoteAddressInterfaceName", nil})
    for i := range allowEntries.AllowEntryLocalAddressRemoteAddressInterfaceName {
        allowEntries.EntityData.Children.Append(types.GetSegmentPath(allowEntries.AllowEntryLocalAddressRemoteAddressInterfaceName[i]), types.YChild{"AllowEntryLocalAddressRemoteAddressInterfaceName", allowEntries.AllowEntryLocalAddressRemoteAddressInterfaceName[i]})
    }
    allowEntries.EntityData.Leafs = types.NewOrderedMap()

    allowEntries.EntityData.YListKeys = []string {}

    return &(allowEntries.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntry
// Allow traffic matching a filter
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. L4 protocol. The type is IpProtocol.
    Protocol interface{}

    // This attribute is a key. Local port. The type is interface{} with range:
    // 1..65535.
    LocalPort interface{}
}

func (allowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntry) GetEntityData() *types.CommonEntityData {
    allowEntry.EntityData.YFilter = allowEntry.YFilter
    allowEntry.EntityData.YangName = "allow-entry"
    allowEntry.EntityData.BundleName = "cisco_ios_xr"
    allowEntry.EntityData.ParentYangName = "allow-entries"
    allowEntry.EntityData.SegmentPath = "allow-entry" + types.AddKeyToken(allowEntry.Protocol, "protocol") + types.AddKeyToken(allowEntry.LocalPort, "local-port")
    allowEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv6/allow-entries/" + allowEntry.EntityData.SegmentPath
    allowEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntry.EntityData.Children = types.NewOrderedMap()
    allowEntry.EntityData.Leafs = types.NewOrderedMap()
    allowEntry.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", allowEntry.Protocol})
    allowEntry.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", allowEntry.LocalPort})

    allowEntry.EntityData.YListKeys = []string {"Protocol", "LocalPort"}

    return &(allowEntry.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddress
// Allow traffic matching a filter
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. local prefix/length. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    LocalAddress interface{}

    // This attribute is a key. L4 protocol. The type is IpProtocol.
    Protocol interface{}

    // This attribute is a key. Local port. The type is interface{} with range:
    // 1..65535.
    LocalPort interface{}
}

func (allowEntryLocalAddress *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddress) GetEntityData() *types.CommonEntityData {
    allowEntryLocalAddress.EntityData.YFilter = allowEntryLocalAddress.YFilter
    allowEntryLocalAddress.EntityData.YangName = "allow-entry-local-address"
    allowEntryLocalAddress.EntityData.BundleName = "cisco_ios_xr"
    allowEntryLocalAddress.EntityData.ParentYangName = "allow-entries"
    allowEntryLocalAddress.EntityData.SegmentPath = "allow-entry-local-address" + types.AddKeyToken(allowEntryLocalAddress.LocalAddress, "local-address") + types.AddKeyToken(allowEntryLocalAddress.Protocol, "protocol") + types.AddKeyToken(allowEntryLocalAddress.LocalPort, "local-port")
    allowEntryLocalAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv6/allow-entries/" + allowEntryLocalAddress.EntityData.SegmentPath
    allowEntryLocalAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntryLocalAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntryLocalAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntryLocalAddress.EntityData.Children = types.NewOrderedMap()
    allowEntryLocalAddress.EntityData.Leafs = types.NewOrderedMap()
    allowEntryLocalAddress.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", allowEntryLocalAddress.LocalAddress})
    allowEntryLocalAddress.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", allowEntryLocalAddress.Protocol})
    allowEntryLocalAddress.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", allowEntryLocalAddress.LocalPort})

    allowEntryLocalAddress.EntityData.YListKeys = []string {"LocalAddress", "Protocol", "LocalPort"}

    return &(allowEntryLocalAddress.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryRemoteAddress
// Allow traffic matching a filter
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryRemoteAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. remote prefix/length. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    RemoteAddress interface{}

    // This attribute is a key. L4 protocol. The type is IpProtocol.
    Protocol interface{}

    // This attribute is a key. Local port. The type is interface{} with range:
    // 1..65535.
    LocalPort interface{}
}

func (allowEntryRemoteAddress *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryRemoteAddress) GetEntityData() *types.CommonEntityData {
    allowEntryRemoteAddress.EntityData.YFilter = allowEntryRemoteAddress.YFilter
    allowEntryRemoteAddress.EntityData.YangName = "allow-entry-remote-address"
    allowEntryRemoteAddress.EntityData.BundleName = "cisco_ios_xr"
    allowEntryRemoteAddress.EntityData.ParentYangName = "allow-entries"
    allowEntryRemoteAddress.EntityData.SegmentPath = "allow-entry-remote-address" + types.AddKeyToken(allowEntryRemoteAddress.RemoteAddress, "remote-address") + types.AddKeyToken(allowEntryRemoteAddress.Protocol, "protocol") + types.AddKeyToken(allowEntryRemoteAddress.LocalPort, "local-port")
    allowEntryRemoteAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv6/allow-entries/" + allowEntryRemoteAddress.EntityData.SegmentPath
    allowEntryRemoteAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntryRemoteAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntryRemoteAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntryRemoteAddress.EntityData.Children = types.NewOrderedMap()
    allowEntryRemoteAddress.EntityData.Leafs = types.NewOrderedMap()
    allowEntryRemoteAddress.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", allowEntryRemoteAddress.RemoteAddress})
    allowEntryRemoteAddress.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", allowEntryRemoteAddress.Protocol})
    allowEntryRemoteAddress.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", allowEntryRemoteAddress.LocalPort})

    allowEntryRemoteAddress.EntityData.YListKeys = []string {"RemoteAddress", "Protocol", "LocalPort"}

    return &(allowEntryRemoteAddress.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddressRemoteAddress
// Allow traffic matching a filter
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddressRemoteAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. local prefix/length. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    LocalAddress interface{}

    // This attribute is a key. remote prefix/length. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    RemoteAddress interface{}

    // This attribute is a key. L4 protocol. The type is IpProtocol.
    Protocol interface{}

    // This attribute is a key. Local port. The type is interface{} with range:
    // 1..65535.
    LocalPort interface{}
}

func (allowEntryLocalAddressRemoteAddress *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddressRemoteAddress) GetEntityData() *types.CommonEntityData {
    allowEntryLocalAddressRemoteAddress.EntityData.YFilter = allowEntryLocalAddressRemoteAddress.YFilter
    allowEntryLocalAddressRemoteAddress.EntityData.YangName = "allow-entry-local-address-remote-address"
    allowEntryLocalAddressRemoteAddress.EntityData.BundleName = "cisco_ios_xr"
    allowEntryLocalAddressRemoteAddress.EntityData.ParentYangName = "allow-entries"
    allowEntryLocalAddressRemoteAddress.EntityData.SegmentPath = "allow-entry-local-address-remote-address" + types.AddKeyToken(allowEntryLocalAddressRemoteAddress.LocalAddress, "local-address") + types.AddKeyToken(allowEntryLocalAddressRemoteAddress.RemoteAddress, "remote-address") + types.AddKeyToken(allowEntryLocalAddressRemoteAddress.Protocol, "protocol") + types.AddKeyToken(allowEntryLocalAddressRemoteAddress.LocalPort, "local-port")
    allowEntryLocalAddressRemoteAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv6/allow-entries/" + allowEntryLocalAddressRemoteAddress.EntityData.SegmentPath
    allowEntryLocalAddressRemoteAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntryLocalAddressRemoteAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntryLocalAddressRemoteAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntryLocalAddressRemoteAddress.EntityData.Children = types.NewOrderedMap()
    allowEntryLocalAddressRemoteAddress.EntityData.Leafs = types.NewOrderedMap()
    allowEntryLocalAddressRemoteAddress.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", allowEntryLocalAddressRemoteAddress.LocalAddress})
    allowEntryLocalAddressRemoteAddress.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", allowEntryLocalAddressRemoteAddress.RemoteAddress})
    allowEntryLocalAddressRemoteAddress.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", allowEntryLocalAddressRemoteAddress.Protocol})
    allowEntryLocalAddressRemoteAddress.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", allowEntryLocalAddressRemoteAddress.LocalPort})

    allowEntryLocalAddressRemoteAddress.EntityData.YListKeys = []string {"LocalAddress", "RemoteAddress", "Protocol", "LocalPort"}

    return &(allowEntryLocalAddressRemoteAddress.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryInterfaceName
// Allow traffic matching a filter
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryInterfaceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // This attribute is a key. L4 protocol. The type is IpProtocol.
    Protocol interface{}

    // This attribute is a key. Local port. The type is interface{} with range:
    // 1..65535.
    LocalPort interface{}
}

func (allowEntryInterfaceName *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryInterfaceName) GetEntityData() *types.CommonEntityData {
    allowEntryInterfaceName.EntityData.YFilter = allowEntryInterfaceName.YFilter
    allowEntryInterfaceName.EntityData.YangName = "allow-entry-interface-name"
    allowEntryInterfaceName.EntityData.BundleName = "cisco_ios_xr"
    allowEntryInterfaceName.EntityData.ParentYangName = "allow-entries"
    allowEntryInterfaceName.EntityData.SegmentPath = "allow-entry-interface-name" + types.AddKeyToken(allowEntryInterfaceName.InterfaceName, "interface-name") + types.AddKeyToken(allowEntryInterfaceName.Protocol, "protocol") + types.AddKeyToken(allowEntryInterfaceName.LocalPort, "local-port")
    allowEntryInterfaceName.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv6/allow-entries/" + allowEntryInterfaceName.EntityData.SegmentPath
    allowEntryInterfaceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntryInterfaceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntryInterfaceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntryInterfaceName.EntityData.Children = types.NewOrderedMap()
    allowEntryInterfaceName.EntityData.Leafs = types.NewOrderedMap()
    allowEntryInterfaceName.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", allowEntryInterfaceName.InterfaceName})
    allowEntryInterfaceName.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", allowEntryInterfaceName.Protocol})
    allowEntryInterfaceName.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", allowEntryInterfaceName.LocalPort})

    allowEntryInterfaceName.EntityData.YListKeys = []string {"InterfaceName", "Protocol", "LocalPort"}

    return &(allowEntryInterfaceName.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddressInterfaceName
// Allow traffic matching a filter
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddressInterfaceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. local prefix/length. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    LocalAddress interface{}

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // This attribute is a key. L4 protocol. The type is IpProtocol.
    Protocol interface{}

    // This attribute is a key. Local port. The type is interface{} with range:
    // 1..65535.
    LocalPort interface{}
}

func (allowEntryLocalAddressInterfaceName *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddressInterfaceName) GetEntityData() *types.CommonEntityData {
    allowEntryLocalAddressInterfaceName.EntityData.YFilter = allowEntryLocalAddressInterfaceName.YFilter
    allowEntryLocalAddressInterfaceName.EntityData.YangName = "allow-entry-local-address-interface-name"
    allowEntryLocalAddressInterfaceName.EntityData.BundleName = "cisco_ios_xr"
    allowEntryLocalAddressInterfaceName.EntityData.ParentYangName = "allow-entries"
    allowEntryLocalAddressInterfaceName.EntityData.SegmentPath = "allow-entry-local-address-interface-name" + types.AddKeyToken(allowEntryLocalAddressInterfaceName.LocalAddress, "local-address") + types.AddKeyToken(allowEntryLocalAddressInterfaceName.InterfaceName, "interface-name") + types.AddKeyToken(allowEntryLocalAddressInterfaceName.Protocol, "protocol") + types.AddKeyToken(allowEntryLocalAddressInterfaceName.LocalPort, "local-port")
    allowEntryLocalAddressInterfaceName.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv6/allow-entries/" + allowEntryLocalAddressInterfaceName.EntityData.SegmentPath
    allowEntryLocalAddressInterfaceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntryLocalAddressInterfaceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntryLocalAddressInterfaceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntryLocalAddressInterfaceName.EntityData.Children = types.NewOrderedMap()
    allowEntryLocalAddressInterfaceName.EntityData.Leafs = types.NewOrderedMap()
    allowEntryLocalAddressInterfaceName.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", allowEntryLocalAddressInterfaceName.LocalAddress})
    allowEntryLocalAddressInterfaceName.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", allowEntryLocalAddressInterfaceName.InterfaceName})
    allowEntryLocalAddressInterfaceName.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", allowEntryLocalAddressInterfaceName.Protocol})
    allowEntryLocalAddressInterfaceName.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", allowEntryLocalAddressInterfaceName.LocalPort})

    allowEntryLocalAddressInterfaceName.EntityData.YListKeys = []string {"LocalAddress", "InterfaceName", "Protocol", "LocalPort"}

    return &(allowEntryLocalAddressInterfaceName.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryRemoteAddressInterfaceName
// Allow traffic matching a filter
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryRemoteAddressInterfaceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. remote prefix/length. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    RemoteAddress interface{}

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // This attribute is a key. L4 protocol. The type is IpProtocol.
    Protocol interface{}

    // This attribute is a key. Local port. The type is interface{} with range:
    // 1..65535.
    LocalPort interface{}
}

func (allowEntryRemoteAddressInterfaceName *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryRemoteAddressInterfaceName) GetEntityData() *types.CommonEntityData {
    allowEntryRemoteAddressInterfaceName.EntityData.YFilter = allowEntryRemoteAddressInterfaceName.YFilter
    allowEntryRemoteAddressInterfaceName.EntityData.YangName = "allow-entry-remote-address-interface-name"
    allowEntryRemoteAddressInterfaceName.EntityData.BundleName = "cisco_ios_xr"
    allowEntryRemoteAddressInterfaceName.EntityData.ParentYangName = "allow-entries"
    allowEntryRemoteAddressInterfaceName.EntityData.SegmentPath = "allow-entry-remote-address-interface-name" + types.AddKeyToken(allowEntryRemoteAddressInterfaceName.RemoteAddress, "remote-address") + types.AddKeyToken(allowEntryRemoteAddressInterfaceName.InterfaceName, "interface-name") + types.AddKeyToken(allowEntryRemoteAddressInterfaceName.Protocol, "protocol") + types.AddKeyToken(allowEntryRemoteAddressInterfaceName.LocalPort, "local-port")
    allowEntryRemoteAddressInterfaceName.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv6/allow-entries/" + allowEntryRemoteAddressInterfaceName.EntityData.SegmentPath
    allowEntryRemoteAddressInterfaceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntryRemoteAddressInterfaceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntryRemoteAddressInterfaceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntryRemoteAddressInterfaceName.EntityData.Children = types.NewOrderedMap()
    allowEntryRemoteAddressInterfaceName.EntityData.Leafs = types.NewOrderedMap()
    allowEntryRemoteAddressInterfaceName.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", allowEntryRemoteAddressInterfaceName.RemoteAddress})
    allowEntryRemoteAddressInterfaceName.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", allowEntryRemoteAddressInterfaceName.InterfaceName})
    allowEntryRemoteAddressInterfaceName.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", allowEntryRemoteAddressInterfaceName.Protocol})
    allowEntryRemoteAddressInterfaceName.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", allowEntryRemoteAddressInterfaceName.LocalPort})

    allowEntryRemoteAddressInterfaceName.EntityData.YListKeys = []string {"RemoteAddress", "InterfaceName", "Protocol", "LocalPort"}

    return &(allowEntryRemoteAddressInterfaceName.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddressRemoteAddressInterfaceName
// Allow traffic matching a filter
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddressRemoteAddressInterfaceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. local prefix/length. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    LocalAddress interface{}

    // This attribute is a key. remote prefix/length. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    RemoteAddress interface{}

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // This attribute is a key. L4 protocol. The type is IpProtocol.
    Protocol interface{}

    // This attribute is a key. Local port. The type is interface{} with range:
    // 1..65535.
    LocalPort interface{}
}

func (allowEntryLocalAddressRemoteAddressInterfaceName *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_AllowEntries_AllowEntryLocalAddressRemoteAddressInterfaceName) GetEntityData() *types.CommonEntityData {
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.YFilter = allowEntryLocalAddressRemoteAddressInterfaceName.YFilter
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.YangName = "allow-entry-local-address-remote-address-interface-name"
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.BundleName = "cisco_ios_xr"
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.ParentYangName = "allow-entries"
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.SegmentPath = "allow-entry-local-address-remote-address-interface-name" + types.AddKeyToken(allowEntryLocalAddressRemoteAddressInterfaceName.LocalAddress, "local-address") + types.AddKeyToken(allowEntryLocalAddressRemoteAddressInterfaceName.RemoteAddress, "remote-address") + types.AddKeyToken(allowEntryLocalAddressRemoteAddressInterfaceName.InterfaceName, "interface-name") + types.AddKeyToken(allowEntryLocalAddressRemoteAddressInterfaceName.Protocol, "protocol") + types.AddKeyToken(allowEntryLocalAddressRemoteAddressInterfaceName.LocalPort, "local-port")
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv6/allow-entries/" + allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.SegmentPath
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.Children = types.NewOrderedMap()
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.Leafs = types.NewOrderedMap()
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", allowEntryLocalAddressRemoteAddressInterfaceName.LocalAddress})
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", allowEntryLocalAddressRemoteAddressInterfaceName.RemoteAddress})
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", allowEntryLocalAddressRemoteAddressInterfaceName.InterfaceName})
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", allowEntryLocalAddressRemoteAddressInterfaceName.Protocol})
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", allowEntryLocalAddressRemoteAddressInterfaceName.LocalPort})

    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.YListKeys = []string {"LocalAddress", "RemoteAddress", "InterfaceName", "Protocol", "LocalPort"}

    return &(allowEntryLocalAddressRemoteAddressInterfaceName.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv6_UpdateSource
// Interface used for Source Address
type Tpa_VrfNames_VrfName_AddressFamily_Ipv6_UpdateSource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name for source address. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Use the management port on the Active RP. The type is interface{}.
    ActiveManagement interface{}
}

func (updateSource *Tpa_VrfNames_VrfName_AddressFamily_Ipv6_UpdateSource) GetEntityData() *types.CommonEntityData {
    updateSource.EntityData.YFilter = updateSource.YFilter
    updateSource.EntityData.YangName = "update-source"
    updateSource.EntityData.BundleName = "cisco_ios_xr"
    updateSource.EntityData.ParentYangName = "ipv6"
    updateSource.EntityData.SegmentPath = "update-source"
    updateSource.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv6/" + updateSource.EntityData.SegmentPath
    updateSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    updateSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    updateSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    updateSource.EntityData.Children = types.NewOrderedMap()
    updateSource.EntityData.Leafs = types.NewOrderedMap()
    updateSource.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", updateSource.InterfaceName})
    updateSource.EntityData.Leafs.Append("active-management", types.YLeaf{"ActiveManagement", updateSource.ActiveManagement})

    updateSource.EntityData.YListKeys = []string {}

    return &(updateSource.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4
// IPv4 configuration
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default interface used for routing. The type is string.
    DefaultRoute interface{}

    // Interface used for source address for egress interface.
    InterfaceNames Tpa_VrfNames_VrfName_AddressFamily_Ipv4_InterfaceNames

    // Traffic protection configuration.
    AllowEntries Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries

    // Interface used for Source Address.
    UpdateSource Tpa_VrfNames_VrfName_AddressFamily_Ipv4_UpdateSource
}

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "address-family"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/" + ipv4.EntityData.SegmentPath
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("interface-names", types.YChild{"InterfaceNames", &ipv4.InterfaceNames})
    ipv4.EntityData.Children.Append("allow-entries", types.YChild{"AllowEntries", &ipv4.AllowEntries})
    ipv4.EntityData.Children.Append("update-source", types.YChild{"UpdateSource", &ipv4.UpdateSource})
    ipv4.EntityData.Leafs = types.NewOrderedMap()
    ipv4.EntityData.Leafs.Append("default-route", types.YLeaf{"DefaultRoute", ipv4.DefaultRoute})

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4_InterfaceNames
// Interface used for source address for egress
// interface
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4_InterfaceNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Egress interface name. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv4_InterfaceNames_InterfaceName.
    InterfaceName []*Tpa_VrfNames_VrfName_AddressFamily_Ipv4_InterfaceNames_InterfaceName
}

func (interfaceNames *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_InterfaceNames) GetEntityData() *types.CommonEntityData {
    interfaceNames.EntityData.YFilter = interfaceNames.YFilter
    interfaceNames.EntityData.YangName = "interface-names"
    interfaceNames.EntityData.BundleName = "cisco_ios_xr"
    interfaceNames.EntityData.ParentYangName = "ipv4"
    interfaceNames.EntityData.SegmentPath = "interface-names"
    interfaceNames.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv4/" + interfaceNames.EntityData.SegmentPath
    interfaceNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceNames.EntityData.Children = types.NewOrderedMap()
    interfaceNames.EntityData.Children.Append("interface-name", types.YChild{"InterfaceName", nil})
    for i := range interfaceNames.InterfaceName {
        interfaceNames.EntityData.Children.Append(types.GetSegmentPath(interfaceNames.InterfaceName[i]), types.YChild{"InterfaceName", interfaceNames.InterfaceName[i]})
    }
    interfaceNames.EntityData.Leafs = types.NewOrderedMap()

    interfaceNames.EntityData.YListKeys = []string {}

    return &(interfaceNames.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4_InterfaceNames_InterfaceName
// Egress interface name
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4_InterfaceNames_InterfaceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    InterfaceName interface{}

    // Interface name for source address. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    EgressInterfaceSource interface{}
}

func (interfaceName *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_InterfaceNames_InterfaceName) GetEntityData() *types.CommonEntityData {
    interfaceName.EntityData.YFilter = interfaceName.YFilter
    interfaceName.EntityData.YangName = "interface-name"
    interfaceName.EntityData.BundleName = "cisco_ios_xr"
    interfaceName.EntityData.ParentYangName = "interface-names"
    interfaceName.EntityData.SegmentPath = "interface-name" + types.AddKeyToken(interfaceName.InterfaceName, "interface-name")
    interfaceName.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv4/interface-names/" + interfaceName.EntityData.SegmentPath
    interfaceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaceName.EntityData.Children = types.NewOrderedMap()
    interfaceName.EntityData.Leafs = types.NewOrderedMap()
    interfaceName.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", interfaceName.InterfaceName})
    interfaceName.EntityData.Leafs.Append("egress-interface-source", types.YLeaf{"EgressInterfaceSource", interfaceName.EgressInterfaceSource})

    interfaceName.EntityData.YListKeys = []string {"InterfaceName"}

    return &(interfaceName.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries
// Traffic protection configuration
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Allow traffic matching a filter. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntry.
    AllowEntry []*Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntry

    // Allow traffic matching a filter. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddress.
    AllowEntryLocalAddress []*Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddress

    // Allow traffic matching a filter. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryRemoteAddress.
    AllowEntryRemoteAddress []*Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryRemoteAddress

    // Allow traffic matching a filter. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddressRemoteAddress.
    AllowEntryLocalAddressRemoteAddress []*Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddressRemoteAddress

    // Allow traffic matching a filter. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryInterfaceName.
    AllowEntryInterfaceName []*Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryInterfaceName

    // Allow traffic matching a filter. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddressInterfaceName.
    AllowEntryLocalAddressInterfaceName []*Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddressInterfaceName

    // Allow traffic matching a filter. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryRemoteAddressInterfaceName.
    AllowEntryRemoteAddressInterfaceName []*Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryRemoteAddressInterfaceName

    // Allow traffic matching a filter. The type is slice of
    // Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddressRemoteAddressInterfaceName.
    AllowEntryLocalAddressRemoteAddressInterfaceName []*Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddressRemoteAddressInterfaceName
}

func (allowEntries *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries) GetEntityData() *types.CommonEntityData {
    allowEntries.EntityData.YFilter = allowEntries.YFilter
    allowEntries.EntityData.YangName = "allow-entries"
    allowEntries.EntityData.BundleName = "cisco_ios_xr"
    allowEntries.EntityData.ParentYangName = "ipv4"
    allowEntries.EntityData.SegmentPath = "allow-entries"
    allowEntries.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv4/" + allowEntries.EntityData.SegmentPath
    allowEntries.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntries.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntries.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntries.EntityData.Children = types.NewOrderedMap()
    allowEntries.EntityData.Children.Append("allow-entry", types.YChild{"AllowEntry", nil})
    for i := range allowEntries.AllowEntry {
        allowEntries.EntityData.Children.Append(types.GetSegmentPath(allowEntries.AllowEntry[i]), types.YChild{"AllowEntry", allowEntries.AllowEntry[i]})
    }
    allowEntries.EntityData.Children.Append("allow-entry-local-address", types.YChild{"AllowEntryLocalAddress", nil})
    for i := range allowEntries.AllowEntryLocalAddress {
        allowEntries.EntityData.Children.Append(types.GetSegmentPath(allowEntries.AllowEntryLocalAddress[i]), types.YChild{"AllowEntryLocalAddress", allowEntries.AllowEntryLocalAddress[i]})
    }
    allowEntries.EntityData.Children.Append("allow-entry-remote-address", types.YChild{"AllowEntryRemoteAddress", nil})
    for i := range allowEntries.AllowEntryRemoteAddress {
        allowEntries.EntityData.Children.Append(types.GetSegmentPath(allowEntries.AllowEntryRemoteAddress[i]), types.YChild{"AllowEntryRemoteAddress", allowEntries.AllowEntryRemoteAddress[i]})
    }
    allowEntries.EntityData.Children.Append("allow-entry-local-address-remote-address", types.YChild{"AllowEntryLocalAddressRemoteAddress", nil})
    for i := range allowEntries.AllowEntryLocalAddressRemoteAddress {
        allowEntries.EntityData.Children.Append(types.GetSegmentPath(allowEntries.AllowEntryLocalAddressRemoteAddress[i]), types.YChild{"AllowEntryLocalAddressRemoteAddress", allowEntries.AllowEntryLocalAddressRemoteAddress[i]})
    }
    allowEntries.EntityData.Children.Append("allow-entry-interface-name", types.YChild{"AllowEntryInterfaceName", nil})
    for i := range allowEntries.AllowEntryInterfaceName {
        allowEntries.EntityData.Children.Append(types.GetSegmentPath(allowEntries.AllowEntryInterfaceName[i]), types.YChild{"AllowEntryInterfaceName", allowEntries.AllowEntryInterfaceName[i]})
    }
    allowEntries.EntityData.Children.Append("allow-entry-local-address-interface-name", types.YChild{"AllowEntryLocalAddressInterfaceName", nil})
    for i := range allowEntries.AllowEntryLocalAddressInterfaceName {
        allowEntries.EntityData.Children.Append(types.GetSegmentPath(allowEntries.AllowEntryLocalAddressInterfaceName[i]), types.YChild{"AllowEntryLocalAddressInterfaceName", allowEntries.AllowEntryLocalAddressInterfaceName[i]})
    }
    allowEntries.EntityData.Children.Append("allow-entry-remote-address-interface-name", types.YChild{"AllowEntryRemoteAddressInterfaceName", nil})
    for i := range allowEntries.AllowEntryRemoteAddressInterfaceName {
        allowEntries.EntityData.Children.Append(types.GetSegmentPath(allowEntries.AllowEntryRemoteAddressInterfaceName[i]), types.YChild{"AllowEntryRemoteAddressInterfaceName", allowEntries.AllowEntryRemoteAddressInterfaceName[i]})
    }
    allowEntries.EntityData.Children.Append("allow-entry-local-address-remote-address-interface-name", types.YChild{"AllowEntryLocalAddressRemoteAddressInterfaceName", nil})
    for i := range allowEntries.AllowEntryLocalAddressRemoteAddressInterfaceName {
        allowEntries.EntityData.Children.Append(types.GetSegmentPath(allowEntries.AllowEntryLocalAddressRemoteAddressInterfaceName[i]), types.YChild{"AllowEntryLocalAddressRemoteAddressInterfaceName", allowEntries.AllowEntryLocalAddressRemoteAddressInterfaceName[i]})
    }
    allowEntries.EntityData.Leafs = types.NewOrderedMap()

    allowEntries.EntityData.YListKeys = []string {}

    return &(allowEntries.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntry
// Allow traffic matching a filter
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. L4 protocol. The type is IpProtocol.
    Protocol interface{}

    // This attribute is a key. Local port. The type is interface{} with range:
    // 1..65535.
    LocalPort interface{}
}

func (allowEntry *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntry) GetEntityData() *types.CommonEntityData {
    allowEntry.EntityData.YFilter = allowEntry.YFilter
    allowEntry.EntityData.YangName = "allow-entry"
    allowEntry.EntityData.BundleName = "cisco_ios_xr"
    allowEntry.EntityData.ParentYangName = "allow-entries"
    allowEntry.EntityData.SegmentPath = "allow-entry" + types.AddKeyToken(allowEntry.Protocol, "protocol") + types.AddKeyToken(allowEntry.LocalPort, "local-port")
    allowEntry.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv4/allow-entries/" + allowEntry.EntityData.SegmentPath
    allowEntry.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntry.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntry.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntry.EntityData.Children = types.NewOrderedMap()
    allowEntry.EntityData.Leafs = types.NewOrderedMap()
    allowEntry.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", allowEntry.Protocol})
    allowEntry.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", allowEntry.LocalPort})

    allowEntry.EntityData.YListKeys = []string {"Protocol", "LocalPort"}

    return &(allowEntry.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddress
// Allow traffic matching a filter
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. local prefix/length. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    LocalAddress interface{}

    // This attribute is a key. L4 protocol. The type is IpProtocol.
    Protocol interface{}

    // This attribute is a key. Local port. The type is interface{} with range:
    // 1..65535.
    LocalPort interface{}
}

func (allowEntryLocalAddress *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddress) GetEntityData() *types.CommonEntityData {
    allowEntryLocalAddress.EntityData.YFilter = allowEntryLocalAddress.YFilter
    allowEntryLocalAddress.EntityData.YangName = "allow-entry-local-address"
    allowEntryLocalAddress.EntityData.BundleName = "cisco_ios_xr"
    allowEntryLocalAddress.EntityData.ParentYangName = "allow-entries"
    allowEntryLocalAddress.EntityData.SegmentPath = "allow-entry-local-address" + types.AddKeyToken(allowEntryLocalAddress.LocalAddress, "local-address") + types.AddKeyToken(allowEntryLocalAddress.Protocol, "protocol") + types.AddKeyToken(allowEntryLocalAddress.LocalPort, "local-port")
    allowEntryLocalAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv4/allow-entries/" + allowEntryLocalAddress.EntityData.SegmentPath
    allowEntryLocalAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntryLocalAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntryLocalAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntryLocalAddress.EntityData.Children = types.NewOrderedMap()
    allowEntryLocalAddress.EntityData.Leafs = types.NewOrderedMap()
    allowEntryLocalAddress.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", allowEntryLocalAddress.LocalAddress})
    allowEntryLocalAddress.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", allowEntryLocalAddress.Protocol})
    allowEntryLocalAddress.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", allowEntryLocalAddress.LocalPort})

    allowEntryLocalAddress.EntityData.YListKeys = []string {"LocalAddress", "Protocol", "LocalPort"}

    return &(allowEntryLocalAddress.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryRemoteAddress
// Allow traffic matching a filter
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryRemoteAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. remote prefix/length. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    RemoteAddress interface{}

    // This attribute is a key. L4 protocol. The type is IpProtocol.
    Protocol interface{}

    // This attribute is a key. Local port. The type is interface{} with range:
    // 1..65535.
    LocalPort interface{}
}

func (allowEntryRemoteAddress *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryRemoteAddress) GetEntityData() *types.CommonEntityData {
    allowEntryRemoteAddress.EntityData.YFilter = allowEntryRemoteAddress.YFilter
    allowEntryRemoteAddress.EntityData.YangName = "allow-entry-remote-address"
    allowEntryRemoteAddress.EntityData.BundleName = "cisco_ios_xr"
    allowEntryRemoteAddress.EntityData.ParentYangName = "allow-entries"
    allowEntryRemoteAddress.EntityData.SegmentPath = "allow-entry-remote-address" + types.AddKeyToken(allowEntryRemoteAddress.RemoteAddress, "remote-address") + types.AddKeyToken(allowEntryRemoteAddress.Protocol, "protocol") + types.AddKeyToken(allowEntryRemoteAddress.LocalPort, "local-port")
    allowEntryRemoteAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv4/allow-entries/" + allowEntryRemoteAddress.EntityData.SegmentPath
    allowEntryRemoteAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntryRemoteAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntryRemoteAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntryRemoteAddress.EntityData.Children = types.NewOrderedMap()
    allowEntryRemoteAddress.EntityData.Leafs = types.NewOrderedMap()
    allowEntryRemoteAddress.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", allowEntryRemoteAddress.RemoteAddress})
    allowEntryRemoteAddress.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", allowEntryRemoteAddress.Protocol})
    allowEntryRemoteAddress.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", allowEntryRemoteAddress.LocalPort})

    allowEntryRemoteAddress.EntityData.YListKeys = []string {"RemoteAddress", "Protocol", "LocalPort"}

    return &(allowEntryRemoteAddress.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddressRemoteAddress
// Allow traffic matching a filter
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddressRemoteAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. local prefix/length. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    LocalAddress interface{}

    // This attribute is a key. remote prefix/length. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    RemoteAddress interface{}

    // This attribute is a key. L4 protocol. The type is IpProtocol.
    Protocol interface{}

    // This attribute is a key. Local port. The type is interface{} with range:
    // 1..65535.
    LocalPort interface{}
}

func (allowEntryLocalAddressRemoteAddress *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddressRemoteAddress) GetEntityData() *types.CommonEntityData {
    allowEntryLocalAddressRemoteAddress.EntityData.YFilter = allowEntryLocalAddressRemoteAddress.YFilter
    allowEntryLocalAddressRemoteAddress.EntityData.YangName = "allow-entry-local-address-remote-address"
    allowEntryLocalAddressRemoteAddress.EntityData.BundleName = "cisco_ios_xr"
    allowEntryLocalAddressRemoteAddress.EntityData.ParentYangName = "allow-entries"
    allowEntryLocalAddressRemoteAddress.EntityData.SegmentPath = "allow-entry-local-address-remote-address" + types.AddKeyToken(allowEntryLocalAddressRemoteAddress.LocalAddress, "local-address") + types.AddKeyToken(allowEntryLocalAddressRemoteAddress.RemoteAddress, "remote-address") + types.AddKeyToken(allowEntryLocalAddressRemoteAddress.Protocol, "protocol") + types.AddKeyToken(allowEntryLocalAddressRemoteAddress.LocalPort, "local-port")
    allowEntryLocalAddressRemoteAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv4/allow-entries/" + allowEntryLocalAddressRemoteAddress.EntityData.SegmentPath
    allowEntryLocalAddressRemoteAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntryLocalAddressRemoteAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntryLocalAddressRemoteAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntryLocalAddressRemoteAddress.EntityData.Children = types.NewOrderedMap()
    allowEntryLocalAddressRemoteAddress.EntityData.Leafs = types.NewOrderedMap()
    allowEntryLocalAddressRemoteAddress.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", allowEntryLocalAddressRemoteAddress.LocalAddress})
    allowEntryLocalAddressRemoteAddress.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", allowEntryLocalAddressRemoteAddress.RemoteAddress})
    allowEntryLocalAddressRemoteAddress.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", allowEntryLocalAddressRemoteAddress.Protocol})
    allowEntryLocalAddressRemoteAddress.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", allowEntryLocalAddressRemoteAddress.LocalPort})

    allowEntryLocalAddressRemoteAddress.EntityData.YListKeys = []string {"LocalAddress", "RemoteAddress", "Protocol", "LocalPort"}

    return &(allowEntryLocalAddressRemoteAddress.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryInterfaceName
// Allow traffic matching a filter
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryInterfaceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // This attribute is a key. L4 protocol. The type is IpProtocol.
    Protocol interface{}

    // This attribute is a key. Local port. The type is interface{} with range:
    // 1..65535.
    LocalPort interface{}
}

func (allowEntryInterfaceName *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryInterfaceName) GetEntityData() *types.CommonEntityData {
    allowEntryInterfaceName.EntityData.YFilter = allowEntryInterfaceName.YFilter
    allowEntryInterfaceName.EntityData.YangName = "allow-entry-interface-name"
    allowEntryInterfaceName.EntityData.BundleName = "cisco_ios_xr"
    allowEntryInterfaceName.EntityData.ParentYangName = "allow-entries"
    allowEntryInterfaceName.EntityData.SegmentPath = "allow-entry-interface-name" + types.AddKeyToken(allowEntryInterfaceName.InterfaceName, "interface-name") + types.AddKeyToken(allowEntryInterfaceName.Protocol, "protocol") + types.AddKeyToken(allowEntryInterfaceName.LocalPort, "local-port")
    allowEntryInterfaceName.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv4/allow-entries/" + allowEntryInterfaceName.EntityData.SegmentPath
    allowEntryInterfaceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntryInterfaceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntryInterfaceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntryInterfaceName.EntityData.Children = types.NewOrderedMap()
    allowEntryInterfaceName.EntityData.Leafs = types.NewOrderedMap()
    allowEntryInterfaceName.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", allowEntryInterfaceName.InterfaceName})
    allowEntryInterfaceName.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", allowEntryInterfaceName.Protocol})
    allowEntryInterfaceName.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", allowEntryInterfaceName.LocalPort})

    allowEntryInterfaceName.EntityData.YListKeys = []string {"InterfaceName", "Protocol", "LocalPort"}

    return &(allowEntryInterfaceName.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddressInterfaceName
// Allow traffic matching a filter
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddressInterfaceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. local prefix/length. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    LocalAddress interface{}

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // This attribute is a key. L4 protocol. The type is IpProtocol.
    Protocol interface{}

    // This attribute is a key. Local port. The type is interface{} with range:
    // 1..65535.
    LocalPort interface{}
}

func (allowEntryLocalAddressInterfaceName *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddressInterfaceName) GetEntityData() *types.CommonEntityData {
    allowEntryLocalAddressInterfaceName.EntityData.YFilter = allowEntryLocalAddressInterfaceName.YFilter
    allowEntryLocalAddressInterfaceName.EntityData.YangName = "allow-entry-local-address-interface-name"
    allowEntryLocalAddressInterfaceName.EntityData.BundleName = "cisco_ios_xr"
    allowEntryLocalAddressInterfaceName.EntityData.ParentYangName = "allow-entries"
    allowEntryLocalAddressInterfaceName.EntityData.SegmentPath = "allow-entry-local-address-interface-name" + types.AddKeyToken(allowEntryLocalAddressInterfaceName.LocalAddress, "local-address") + types.AddKeyToken(allowEntryLocalAddressInterfaceName.InterfaceName, "interface-name") + types.AddKeyToken(allowEntryLocalAddressInterfaceName.Protocol, "protocol") + types.AddKeyToken(allowEntryLocalAddressInterfaceName.LocalPort, "local-port")
    allowEntryLocalAddressInterfaceName.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv4/allow-entries/" + allowEntryLocalAddressInterfaceName.EntityData.SegmentPath
    allowEntryLocalAddressInterfaceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntryLocalAddressInterfaceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntryLocalAddressInterfaceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntryLocalAddressInterfaceName.EntityData.Children = types.NewOrderedMap()
    allowEntryLocalAddressInterfaceName.EntityData.Leafs = types.NewOrderedMap()
    allowEntryLocalAddressInterfaceName.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", allowEntryLocalAddressInterfaceName.LocalAddress})
    allowEntryLocalAddressInterfaceName.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", allowEntryLocalAddressInterfaceName.InterfaceName})
    allowEntryLocalAddressInterfaceName.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", allowEntryLocalAddressInterfaceName.Protocol})
    allowEntryLocalAddressInterfaceName.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", allowEntryLocalAddressInterfaceName.LocalPort})

    allowEntryLocalAddressInterfaceName.EntityData.YListKeys = []string {"LocalAddress", "InterfaceName", "Protocol", "LocalPort"}

    return &(allowEntryLocalAddressInterfaceName.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryRemoteAddressInterfaceName
// Allow traffic matching a filter
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryRemoteAddressInterfaceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. remote prefix/length. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    RemoteAddress interface{}

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // This attribute is a key. L4 protocol. The type is IpProtocol.
    Protocol interface{}

    // This attribute is a key. Local port. The type is interface{} with range:
    // 1..65535.
    LocalPort interface{}
}

func (allowEntryRemoteAddressInterfaceName *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryRemoteAddressInterfaceName) GetEntityData() *types.CommonEntityData {
    allowEntryRemoteAddressInterfaceName.EntityData.YFilter = allowEntryRemoteAddressInterfaceName.YFilter
    allowEntryRemoteAddressInterfaceName.EntityData.YangName = "allow-entry-remote-address-interface-name"
    allowEntryRemoteAddressInterfaceName.EntityData.BundleName = "cisco_ios_xr"
    allowEntryRemoteAddressInterfaceName.EntityData.ParentYangName = "allow-entries"
    allowEntryRemoteAddressInterfaceName.EntityData.SegmentPath = "allow-entry-remote-address-interface-name" + types.AddKeyToken(allowEntryRemoteAddressInterfaceName.RemoteAddress, "remote-address") + types.AddKeyToken(allowEntryRemoteAddressInterfaceName.InterfaceName, "interface-name") + types.AddKeyToken(allowEntryRemoteAddressInterfaceName.Protocol, "protocol") + types.AddKeyToken(allowEntryRemoteAddressInterfaceName.LocalPort, "local-port")
    allowEntryRemoteAddressInterfaceName.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv4/allow-entries/" + allowEntryRemoteAddressInterfaceName.EntityData.SegmentPath
    allowEntryRemoteAddressInterfaceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntryRemoteAddressInterfaceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntryRemoteAddressInterfaceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntryRemoteAddressInterfaceName.EntityData.Children = types.NewOrderedMap()
    allowEntryRemoteAddressInterfaceName.EntityData.Leafs = types.NewOrderedMap()
    allowEntryRemoteAddressInterfaceName.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", allowEntryRemoteAddressInterfaceName.RemoteAddress})
    allowEntryRemoteAddressInterfaceName.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", allowEntryRemoteAddressInterfaceName.InterfaceName})
    allowEntryRemoteAddressInterfaceName.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", allowEntryRemoteAddressInterfaceName.Protocol})
    allowEntryRemoteAddressInterfaceName.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", allowEntryRemoteAddressInterfaceName.LocalPort})

    allowEntryRemoteAddressInterfaceName.EntityData.YListKeys = []string {"RemoteAddress", "InterfaceName", "Protocol", "LocalPort"}

    return &(allowEntryRemoteAddressInterfaceName.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddressRemoteAddressInterfaceName
// Allow traffic matching a filter
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddressRemoteAddressInterfaceName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. local prefix/length. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    LocalAddress interface{}

    // This attribute is a key. remote prefix/length. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2])),
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8]))).
    RemoteAddress interface{}

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // This attribute is a key. L4 protocol. The type is IpProtocol.
    Protocol interface{}

    // This attribute is a key. Local port. The type is interface{} with range:
    // 1..65535.
    LocalPort interface{}
}

func (allowEntryLocalAddressRemoteAddressInterfaceName *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_AllowEntries_AllowEntryLocalAddressRemoteAddressInterfaceName) GetEntityData() *types.CommonEntityData {
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.YFilter = allowEntryLocalAddressRemoteAddressInterfaceName.YFilter
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.YangName = "allow-entry-local-address-remote-address-interface-name"
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.BundleName = "cisco_ios_xr"
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.ParentYangName = "allow-entries"
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.SegmentPath = "allow-entry-local-address-remote-address-interface-name" + types.AddKeyToken(allowEntryLocalAddressRemoteAddressInterfaceName.LocalAddress, "local-address") + types.AddKeyToken(allowEntryLocalAddressRemoteAddressInterfaceName.RemoteAddress, "remote-address") + types.AddKeyToken(allowEntryLocalAddressRemoteAddressInterfaceName.InterfaceName, "interface-name") + types.AddKeyToken(allowEntryLocalAddressRemoteAddressInterfaceName.Protocol, "protocol") + types.AddKeyToken(allowEntryLocalAddressRemoteAddressInterfaceName.LocalPort, "local-port")
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv4/allow-entries/" + allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.SegmentPath
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.Children = types.NewOrderedMap()
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.Leafs = types.NewOrderedMap()
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.Leafs.Append("local-address", types.YLeaf{"LocalAddress", allowEntryLocalAddressRemoteAddressInterfaceName.LocalAddress})
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.Leafs.Append("remote-address", types.YLeaf{"RemoteAddress", allowEntryLocalAddressRemoteAddressInterfaceName.RemoteAddress})
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", allowEntryLocalAddressRemoteAddressInterfaceName.InterfaceName})
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.Leafs.Append("protocol", types.YLeaf{"Protocol", allowEntryLocalAddressRemoteAddressInterfaceName.Protocol})
    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.Leafs.Append("local-port", types.YLeaf{"LocalPort", allowEntryLocalAddressRemoteAddressInterfaceName.LocalPort})

    allowEntryLocalAddressRemoteAddressInterfaceName.EntityData.YListKeys = []string {"LocalAddress", "RemoteAddress", "InterfaceName", "Protocol", "LocalPort"}

    return &(allowEntryLocalAddressRemoteAddressInterfaceName.EntityData)
}

// Tpa_VrfNames_VrfName_AddressFamily_Ipv4_UpdateSource
// Interface used for Source Address
type Tpa_VrfNames_VrfName_AddressFamily_Ipv4_UpdateSource struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface name for source address. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Use the management port on the Active RP. The type is interface{}.
    ActiveManagement interface{}
}

func (updateSource *Tpa_VrfNames_VrfName_AddressFamily_Ipv4_UpdateSource) GetEntityData() *types.CommonEntityData {
    updateSource.EntityData.YFilter = updateSource.YFilter
    updateSource.EntityData.YangName = "update-source"
    updateSource.EntityData.BundleName = "cisco_ios_xr"
    updateSource.EntityData.ParentYangName = "ipv4"
    updateSource.EntityData.SegmentPath = "update-source"
    updateSource.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/vrf-names/vrf-name/address-family/ipv4/" + updateSource.EntityData.SegmentPath
    updateSource.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    updateSource.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    updateSource.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    updateSource.EntityData.Children = types.NewOrderedMap()
    updateSource.EntityData.Leafs = types.NewOrderedMap()
    updateSource.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", updateSource.InterfaceName})
    updateSource.EntityData.Leafs.Append("active-management", types.YLeaf{"ActiveManagement", updateSource.ActiveManagement})

    updateSource.EntityData.YListKeys = []string {}

    return &(updateSource.EntityData)
}

// Tpa_Logging
// Third party app logging
type Tpa_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // KIM logging.
    Kim Tpa_Logging_Kim
}

func (logging *Tpa_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "tpa"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/" + logging.EntityData.SegmentPath
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Children.Append("kim", types.YChild{"Kim", &logging.Kim})
    logging.EntityData.Leafs = types.NewOrderedMap()

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// Tpa_Logging_Kim
// KIM logging
type Tpa_Logging_Kim struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // How many log rotation files to keep. The type is interface{} with range:
    // 0..4294967295.
    RotationMax interface{}

    // Size in Kilobytes of the log file. The type is interface{} with range:
    // 0..4294967295. Units are kilobyte.
    FileSizeMaxKb interface{}
}

func (kim *Tpa_Logging_Kim) GetEntityData() *types.CommonEntityData {
    kim.EntityData.YFilter = kim.YFilter
    kim.EntityData.YangName = "kim"
    kim.EntityData.BundleName = "cisco_ios_xr"
    kim.EntityData.ParentYangName = "logging"
    kim.EntityData.SegmentPath = "kim"
    kim.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/logging/" + kim.EntityData.SegmentPath
    kim.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    kim.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    kim.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    kim.EntityData.Children = types.NewOrderedMap()
    kim.EntityData.Leafs = types.NewOrderedMap()
    kim.EntityData.Leafs.Append("rotation-max", types.YLeaf{"RotationMax", kim.RotationMax})
    kim.EntityData.Leafs.Append("file-size-max-kb", types.YLeaf{"FileSizeMaxKb", kim.FileSizeMaxKb})

    kim.EntityData.YListKeys = []string {}

    return &(kim.EntityData)
}

// Tpa_Statistics
// Statistics
type Tpa_Statistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // How many interface events to record. The type is interface{} with range:
    // 0..4294967295.
    MaxIntfEvents interface{}

    // How many LPTS events to record. The type is interface{} with range:
    // 0..4294967295.
    MaxLptsEvents interface{}

    // Statistics update frequency into Linux. The type is interface{} with range:
    // 0..4294967295. Units are second.
    StatisticsUpdateFrequency interface{}
}

func (statistics *Tpa_Statistics) GetEntityData() *types.CommonEntityData {
    statistics.EntityData.YFilter = statistics.YFilter
    statistics.EntityData.YangName = "statistics"
    statistics.EntityData.BundleName = "cisco_ios_xr"
    statistics.EntityData.ParentYangName = "tpa"
    statistics.EntityData.SegmentPath = "statistics"
    statistics.EntityData.AbsolutePath = "Cisco-IOS-XR-kim-tpa-cfg:tpa/" + statistics.EntityData.SegmentPath
    statistics.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    statistics.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    statistics.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    statistics.EntityData.Children = types.NewOrderedMap()
    statistics.EntityData.Leafs = types.NewOrderedMap()
    statistics.EntityData.Leafs.Append("max-intf-events", types.YLeaf{"MaxIntfEvents", statistics.MaxIntfEvents})
    statistics.EntityData.Leafs.Append("max-lpts-events", types.YLeaf{"MaxLptsEvents", statistics.MaxLptsEvents})
    statistics.EntityData.Leafs.Append("statistics-update-frequency", types.YLeaf{"StatisticsUpdateFrequency", statistics.StatisticsUpdateFrequency})

    statistics.EntityData.YListKeys = []string {}

    return &(statistics.EntityData)
}

