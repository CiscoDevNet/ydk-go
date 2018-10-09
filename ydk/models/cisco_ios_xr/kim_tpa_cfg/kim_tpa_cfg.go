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

    // Interface used for Source Address.
    UpdateSource Tpa_VrfNames_VrfName_AddressFamily_Ipv6_UpdateSource
}

func (ipv6 *Tpa_VrfNames_VrfName_AddressFamily_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "address-family"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("interface-names", types.YChild{"InterfaceNames", &ipv6.InterfaceNames})
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

    // Interface used for Source Address.
    UpdateSource Tpa_VrfNames_VrfName_AddressFamily_Ipv4_UpdateSource
}

func (ipv4 *Tpa_VrfNames_VrfName_AddressFamily_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "address-family"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("interface-names", types.YChild{"InterfaceNames", &ipv4.InterfaceNames})
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

