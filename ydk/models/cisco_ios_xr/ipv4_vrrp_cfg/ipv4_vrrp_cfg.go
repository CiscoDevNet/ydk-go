// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-vrrp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   vrrp: VRRP configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-snmp-agent-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_vrrp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_vrrp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-vrrp-cfg vrrp}", reflect.TypeOf(Vrrp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp", reflect.TypeOf(Vrrp{}))
}

// Vrrp
// VRRP configuration
type Vrrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable vrrp process. The type is interface{}.
    Enable interface{}

    // VRRP logging options.
    Logging Vrrp_Logging

    // Interface configuration table.
    Interfaces Vrrp_Interfaces
}

func (vrrp *Vrrp) GetEntityData() *types.CommonEntityData {
    vrrp.EntityData.YFilter = vrrp.YFilter
    vrrp.EntityData.YangName = "vrrp"
    vrrp.EntityData.BundleName = "cisco_ios_xr"
    vrrp.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-vrrp-cfg"
    vrrp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp"
    vrrp.EntityData.AbsolutePath = vrrp.EntityData.SegmentPath
    vrrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrrp.EntityData.Children = types.NewOrderedMap()
    vrrp.EntityData.Children.Append("logging", types.YChild{"Logging", &vrrp.Logging})
    vrrp.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &vrrp.Interfaces})
    vrrp.EntityData.Leafs = types.NewOrderedMap()
    vrrp.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", vrrp.Enable})

    vrrp.EntityData.YListKeys = []string {}

    return &(vrrp.EntityData)
}

// Vrrp_Logging
// VRRP logging options
type Vrrp_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRRP state change IOS messages disable. The type is interface{}.
    StateChangeDisable interface{}
}

func (logging *Vrrp_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "vrrp"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/" + logging.EntityData.SegmentPath
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = types.NewOrderedMap()
    logging.EntityData.Leafs = types.NewOrderedMap()
    logging.EntityData.Leafs.Append("state-change-disable", types.YLeaf{"StateChangeDisable", logging.StateChangeDisable})

    logging.EntityData.YListKeys = []string {}

    return &(logging.EntityData)
}

// Vrrp_Interfaces
// Interface configuration table
type Vrrp_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The interface being configured. The type is slice of
    // Vrrp_Interfaces_Interface.
    Interface []*Vrrp_Interfaces_Interface
}

func (interfaces *Vrrp_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "vrrp"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Vrrp_Interfaces_Interface
// The interface being configured
type Vrrp_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface name to configure. The type is string
    // with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // VRRP Slave MAC-refresh rate in seconds. The type is interface{} with range:
    // 0..10000. Units are second. The default value is 60.
    MacRefresh interface{}

    // IPv6 VRRP configuration.
    Ipv6 Vrrp_Interfaces_Interface_Ipv6

    // Minimum and Reload Delay.
    Delay Vrrp_Interfaces_Interface_Delay

    // IPv4 VRRP configuration.
    Ipv4 Vrrp_Interfaces_Interface_Ipv4

    // BFD configuration.
    Bfd Vrrp_Interfaces_Interface_Bfd
}

func (self *Vrrp_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("ipv6", types.YChild{"Ipv6", &self.Ipv6})
    self.EntityData.Children.Append("delay", types.YChild{"Delay", &self.Delay})
    self.EntityData.Children.Append("ipv4", types.YChild{"Ipv4", &self.Ipv4})
    self.EntityData.Children.Append("bfd", types.YChild{"Bfd", &self.Bfd})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("mac-refresh", types.YLeaf{"MacRefresh", self.MacRefresh})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6
// IPv6 VRRP configuration
type Vrrp_Interfaces_Interface_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version 3 VRRP configuration.
    Version3 Vrrp_Interfaces_Interface_Ipv6_Version3

    // The VRRP slave group configuration table.
    SlaveVirtualRouters Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters
}

func (ipv6 *Vrrp_Interfaces_Interface_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "interface"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/" + ipv6.EntityData.SegmentPath
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = types.NewOrderedMap()
    ipv6.EntityData.Children.Append("version3", types.YChild{"Version3", &ipv6.Version3})
    ipv6.EntityData.Children.Append("slave-virtual-routers", types.YChild{"SlaveVirtualRouters", &ipv6.SlaveVirtualRouters})
    ipv6.EntityData.Leafs = types.NewOrderedMap()

    ipv6.EntityData.YListKeys = []string {}

    return &(ipv6.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6_Version3
// Version 3 VRRP configuration
type Vrrp_Interfaces_Interface_Ipv6_Version3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The VRRP virtual router configuration table.
    VirtualRouters Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters
}

func (version3 *Vrrp_Interfaces_Interface_Ipv6_Version3) GetEntityData() *types.CommonEntityData {
    version3.EntityData.YFilter = version3.YFilter
    version3.EntityData.YangName = "version3"
    version3.EntityData.BundleName = "cisco_ios_xr"
    version3.EntityData.ParentYangName = "ipv6"
    version3.EntityData.SegmentPath = "version3"
    version3.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv6/" + version3.EntityData.SegmentPath
    version3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    version3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    version3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    version3.EntityData.Children = types.NewOrderedMap()
    version3.EntityData.Children.Append("virtual-routers", types.YChild{"VirtualRouters", &version3.VirtualRouters})
    version3.EntityData.Leafs = types.NewOrderedMap()

    version3.EntityData.YListKeys = []string {}

    return &(version3.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters
// The VRRP virtual router configuration table
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The VRRP virtual router being configured. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter.
    VirtualRouter []*Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters) GetEntityData() *types.CommonEntityData {
    virtualRouters.EntityData.YFilter = virtualRouters.YFilter
    virtualRouters.EntityData.YangName = "virtual-routers"
    virtualRouters.EntityData.BundleName = "cisco_ios_xr"
    virtualRouters.EntityData.ParentYangName = "version3"
    virtualRouters.EntityData.SegmentPath = "virtual-routers"
    virtualRouters.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv6/version3/" + virtualRouters.EntityData.SegmentPath
    virtualRouters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualRouters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualRouters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualRouters.EntityData.Children = types.NewOrderedMap()
    virtualRouters.EntityData.Children.Append("virtual-router", types.YChild{"VirtualRouter", nil})
    for i := range virtualRouters.VirtualRouter {
        virtualRouters.EntityData.Children.Append(types.GetSegmentPath(virtualRouters.VirtualRouter[i]), types.YChild{"VirtualRouter", virtualRouters.VirtualRouter[i]})
    }
    virtualRouters.EntityData.Leafs = types.NewOrderedMap()

    virtualRouters.EntityData.YListKeys = []string {}

    return &(virtualRouters.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter
// The VRRP virtual router being configured
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRID Virtual Router Identifier. The type is
    // interface{} with range: 1..255.
    VrId interface{}

    // Enable use of Bidirectional Forwarding Detection for this IP. The type is
    // one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Bfd interface{}

    // Priority value. The type is interface{} with range: 1..254. The default
    // value is 100.
    Priority interface{}

    // Disable Accept Mode for this virtual IPAddress. The type is interface{}.
    AcceptModeDisable interface{}

    // Preempt Master router if higher priority. The type is interface{} with
    // range: 0..3600. The default value is 0.
    Preempt interface{}

    // VRRP Session Name. The type is string with length: 1..16.
    SessionName interface{}

    // The table of VRRP virtual global IPv6 addresses.
    GlobalIpv6Addresses Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses

    // Track an item, reducing priority if it goes down.
    Tracks Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks

    // Set advertisement timer.
    Timer Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer

    // Track an object, reducing priority if it goes down.
    TrackedObjects Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects

    // The VRRP IPv6 virtual linklocal address.
    LinkLocalIpv6Address Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter) GetEntityData() *types.CommonEntityData {
    virtualRouter.EntityData.YFilter = virtualRouter.YFilter
    virtualRouter.EntityData.YangName = "virtual-router"
    virtualRouter.EntityData.BundleName = "cisco_ios_xr"
    virtualRouter.EntityData.ParentYangName = "virtual-routers"
    virtualRouter.EntityData.SegmentPath = "virtual-router" + types.AddKeyToken(virtualRouter.VrId, "vr-id")
    virtualRouter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv6/version3/virtual-routers/" + virtualRouter.EntityData.SegmentPath
    virtualRouter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualRouter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualRouter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualRouter.EntityData.Children = types.NewOrderedMap()
    virtualRouter.EntityData.Children.Append("global-ipv6-addresses", types.YChild{"GlobalIpv6Addresses", &virtualRouter.GlobalIpv6Addresses})
    virtualRouter.EntityData.Children.Append("tracks", types.YChild{"Tracks", &virtualRouter.Tracks})
    virtualRouter.EntityData.Children.Append("timer", types.YChild{"Timer", &virtualRouter.Timer})
    virtualRouter.EntityData.Children.Append("tracked-objects", types.YChild{"TrackedObjects", &virtualRouter.TrackedObjects})
    virtualRouter.EntityData.Children.Append("link-local-ipv6-address", types.YChild{"LinkLocalIpv6Address", &virtualRouter.LinkLocalIpv6Address})
    virtualRouter.EntityData.Leafs = types.NewOrderedMap()
    virtualRouter.EntityData.Leafs.Append("vr-id", types.YLeaf{"VrId", virtualRouter.VrId})
    virtualRouter.EntityData.Leafs.Append("bfd", types.YLeaf{"Bfd", virtualRouter.Bfd})
    virtualRouter.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", virtualRouter.Priority})
    virtualRouter.EntityData.Leafs.Append("accept-mode-disable", types.YLeaf{"AcceptModeDisable", virtualRouter.AcceptModeDisable})
    virtualRouter.EntityData.Leafs.Append("preempt", types.YLeaf{"Preempt", virtualRouter.Preempt})
    virtualRouter.EntityData.Leafs.Append("session-name", types.YLeaf{"SessionName", virtualRouter.SessionName})

    virtualRouter.EntityData.YListKeys = []string {"VrId"}

    return &(virtualRouter.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses
// The table of VRRP virtual global IPv6
// addresses
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VRRP virtual global IPv6 IP address. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address.
    GlobalIpv6Address []*Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address
}

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses) GetEntityData() *types.CommonEntityData {
    globalIpv6Addresses.EntityData.YFilter = globalIpv6Addresses.YFilter
    globalIpv6Addresses.EntityData.YangName = "global-ipv6-addresses"
    globalIpv6Addresses.EntityData.BundleName = "cisco_ios_xr"
    globalIpv6Addresses.EntityData.ParentYangName = "virtual-router"
    globalIpv6Addresses.EntityData.SegmentPath = "global-ipv6-addresses"
    globalIpv6Addresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv6/version3/virtual-routers/virtual-router/" + globalIpv6Addresses.EntityData.SegmentPath
    globalIpv6Addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalIpv6Addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalIpv6Addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalIpv6Addresses.EntityData.Children = types.NewOrderedMap()
    globalIpv6Addresses.EntityData.Children.Append("global-ipv6-address", types.YChild{"GlobalIpv6Address", nil})
    for i := range globalIpv6Addresses.GlobalIpv6Address {
        globalIpv6Addresses.EntityData.Children.Append(types.GetSegmentPath(globalIpv6Addresses.GlobalIpv6Address[i]), types.YChild{"GlobalIpv6Address", globalIpv6Addresses.GlobalIpv6Address[i]})
    }
    globalIpv6Addresses.EntityData.Leafs = types.NewOrderedMap()

    globalIpv6Addresses.EntityData.YListKeys = []string {}

    return &(globalIpv6Addresses.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address
// A VRRP virtual global IPv6 IP address
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRRP virtual global IPv6 address. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}
}

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetEntityData() *types.CommonEntityData {
    globalIpv6Address.EntityData.YFilter = globalIpv6Address.YFilter
    globalIpv6Address.EntityData.YangName = "global-ipv6-address"
    globalIpv6Address.EntityData.BundleName = "cisco_ios_xr"
    globalIpv6Address.EntityData.ParentYangName = "global-ipv6-addresses"
    globalIpv6Address.EntityData.SegmentPath = "global-ipv6-address" + types.AddKeyToken(globalIpv6Address.IpAddress, "ip-address")
    globalIpv6Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv6/version3/virtual-routers/virtual-router/global-ipv6-addresses/" + globalIpv6Address.EntityData.SegmentPath
    globalIpv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalIpv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalIpv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalIpv6Address.EntityData.Children = types.NewOrderedMap()
    globalIpv6Address.EntityData.Leafs = types.NewOrderedMap()
    globalIpv6Address.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", globalIpv6Address.IpAddress})

    globalIpv6Address.EntityData.YListKeys = []string {"IpAddress"}

    return &(globalIpv6Address.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks
// Track an item, reducing priority if it
// goes down
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object to be tracked. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track.
    Track []*Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track
}

func (tracks *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks) GetEntityData() *types.CommonEntityData {
    tracks.EntityData.YFilter = tracks.YFilter
    tracks.EntityData.YangName = "tracks"
    tracks.EntityData.BundleName = "cisco_ios_xr"
    tracks.EntityData.ParentYangName = "virtual-router"
    tracks.EntityData.SegmentPath = "tracks"
    tracks.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv6/version3/virtual-routers/virtual-router/" + tracks.EntityData.SegmentPath
    tracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tracks.EntityData.Children = types.NewOrderedMap()
    tracks.EntityData.Children.Append("track", types.YChild{"Track", nil})
    for i := range tracks.Track {
        tracks.EntityData.Children.Append(types.GetSegmentPath(tracks.Track[i]), types.YChild{"Track", tracks.Track[i]})
    }
    tracks.EntityData.Leafs = types.NewOrderedMap()

    tracks.EntityData.YListKeys = []string {}

    return &(tracks.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track
// Object to be tracked
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Object to be tracked, interface name for
    // interfaces. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Priority decrement. The type is interface{} with range: 1..254. This
    // attribute is mandatory.
    Priority interface{}
}

func (track *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetEntityData() *types.CommonEntityData {
    track.EntityData.YFilter = track.YFilter
    track.EntityData.YangName = "track"
    track.EntityData.BundleName = "cisco_ios_xr"
    track.EntityData.ParentYangName = "tracks"
    track.EntityData.SegmentPath = "track" + types.AddKeyToken(track.InterfaceName, "interface-name")
    track.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv6/version3/virtual-routers/virtual-router/tracks/" + track.EntityData.SegmentPath
    track.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    track.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    track.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    track.EntityData.Children = types.NewOrderedMap()
    track.EntityData.Leafs = types.NewOrderedMap()
    track.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", track.InterfaceName})
    track.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", track.Priority})

    track.EntityData.YListKeys = []string {"InterfaceName"}

    return &(track.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer
// Set advertisement timer
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Advertisement time in milliseconds. The type is interface{} with range:
    // 100..40950. Units are millisecond.
    AdvertisementTimeInMsec interface{}

    // Advertisement time in seconds. The type is interface{} with range: 1..40.
    // Units are second.
    AdvertisementTimeInSec interface{}

    // TRUE - Force configured timer values to be used, required when configured
    // in milliseconds. The type is interface{}. Units are millisecond.
    Forced interface{}
}

func (timer *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer) GetEntityData() *types.CommonEntityData {
    timer.EntityData.YFilter = timer.YFilter
    timer.EntityData.YangName = "timer"
    timer.EntityData.BundleName = "cisco_ios_xr"
    timer.EntityData.ParentYangName = "virtual-router"
    timer.EntityData.SegmentPath = "timer"
    timer.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv6/version3/virtual-routers/virtual-router/" + timer.EntityData.SegmentPath
    timer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timer.EntityData.Children = types.NewOrderedMap()
    timer.EntityData.Leafs = types.NewOrderedMap()
    timer.EntityData.Leafs.Append("advertisement-time-in-msec", types.YLeaf{"AdvertisementTimeInMsec", timer.AdvertisementTimeInMsec})
    timer.EntityData.Leafs.Append("advertisement-time-in-sec", types.YLeaf{"AdvertisementTimeInSec", timer.AdvertisementTimeInSec})
    timer.EntityData.Leafs.Append("forced", types.YLeaf{"Forced", timer.Forced})

    timer.EntityData.YListKeys = []string {}

    return &(timer.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects
// Track an object, reducing priority if it
// goes down
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object to be tracked. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject.
    TrackedObject []*Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetEntityData() *types.CommonEntityData {
    trackedObjects.EntityData.YFilter = trackedObjects.YFilter
    trackedObjects.EntityData.YangName = "tracked-objects"
    trackedObjects.EntityData.BundleName = "cisco_ios_xr"
    trackedObjects.EntityData.ParentYangName = "virtual-router"
    trackedObjects.EntityData.SegmentPath = "tracked-objects"
    trackedObjects.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv6/version3/virtual-routers/virtual-router/" + trackedObjects.EntityData.SegmentPath
    trackedObjects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedObjects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedObjects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedObjects.EntityData.Children = types.NewOrderedMap()
    trackedObjects.EntityData.Children.Append("tracked-object", types.YChild{"TrackedObject", nil})
    for i := range trackedObjects.TrackedObject {
        trackedObjects.EntityData.Children.Append(types.GetSegmentPath(trackedObjects.TrackedObject[i]), types.YChild{"TrackedObject", trackedObjects.TrackedObject[i]})
    }
    trackedObjects.EntityData.Leafs = types.NewOrderedMap()

    trackedObjects.EntityData.YListKeys = []string {}

    return &(trackedObjects.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject
// Object to be tracked
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Object to be tracked, interface name for
    // interfaces. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ObjectName interface{}

    // Priority decrement. The type is interface{} with range: 1..254. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetEntityData() *types.CommonEntityData {
    trackedObject.EntityData.YFilter = trackedObject.YFilter
    trackedObject.EntityData.YangName = "tracked-object"
    trackedObject.EntityData.BundleName = "cisco_ios_xr"
    trackedObject.EntityData.ParentYangName = "tracked-objects"
    trackedObject.EntityData.SegmentPath = "tracked-object" + types.AddKeyToken(trackedObject.ObjectName, "object-name")
    trackedObject.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv6/version3/virtual-routers/virtual-router/tracked-objects/" + trackedObject.EntityData.SegmentPath
    trackedObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedObject.EntityData.Children = types.NewOrderedMap()
    trackedObject.EntityData.Leafs = types.NewOrderedMap()
    trackedObject.EntityData.Leafs.Append("object-name", types.YLeaf{"ObjectName", trackedObject.ObjectName})
    trackedObject.EntityData.Leafs.Append("priority-decrement", types.YLeaf{"PriorityDecrement", trackedObject.PriorityDecrement})

    trackedObject.EntityData.YListKeys = []string {"ObjectName"}

    return &(trackedObject.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address
// The VRRP IPv6 virtual linklocal address
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRRP IPv6 virtual linklocal address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // TRUE if the virtual linklocal address is to be autoconfigured FALSE if an
    // IPv6 virtual linklocal address is configured. The type is bool. The default
    // value is false.
    AutoConfigure interface{}
}

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address) GetEntityData() *types.CommonEntityData {
    linkLocalIpv6Address.EntityData.YFilter = linkLocalIpv6Address.YFilter
    linkLocalIpv6Address.EntityData.YangName = "link-local-ipv6-address"
    linkLocalIpv6Address.EntityData.BundleName = "cisco_ios_xr"
    linkLocalIpv6Address.EntityData.ParentYangName = "virtual-router"
    linkLocalIpv6Address.EntityData.SegmentPath = "link-local-ipv6-address"
    linkLocalIpv6Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv6/version3/virtual-routers/virtual-router/" + linkLocalIpv6Address.EntityData.SegmentPath
    linkLocalIpv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkLocalIpv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkLocalIpv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkLocalIpv6Address.EntityData.Children = types.NewOrderedMap()
    linkLocalIpv6Address.EntityData.Leafs = types.NewOrderedMap()
    linkLocalIpv6Address.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", linkLocalIpv6Address.IpAddress})
    linkLocalIpv6Address.EntityData.Leafs.Append("auto-configure", types.YLeaf{"AutoConfigure", linkLocalIpv6Address.AutoConfigure})

    linkLocalIpv6Address.EntityData.YListKeys = []string {}

    return &(linkLocalIpv6Address.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters
// The VRRP slave group configuration table
type Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The VRRP slave being configured. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter.
    SlaveVirtualRouter []*Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter
}

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters) GetEntityData() *types.CommonEntityData {
    slaveVirtualRouters.EntityData.YFilter = slaveVirtualRouters.YFilter
    slaveVirtualRouters.EntityData.YangName = "slave-virtual-routers"
    slaveVirtualRouters.EntityData.BundleName = "cisco_ios_xr"
    slaveVirtualRouters.EntityData.ParentYangName = "ipv6"
    slaveVirtualRouters.EntityData.SegmentPath = "slave-virtual-routers"
    slaveVirtualRouters.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv6/" + slaveVirtualRouters.EntityData.SegmentPath
    slaveVirtualRouters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slaveVirtualRouters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slaveVirtualRouters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slaveVirtualRouters.EntityData.Children = types.NewOrderedMap()
    slaveVirtualRouters.EntityData.Children.Append("slave-virtual-router", types.YChild{"SlaveVirtualRouter", nil})
    for i := range slaveVirtualRouters.SlaveVirtualRouter {
        slaveVirtualRouters.EntityData.Children.Append(types.GetSegmentPath(slaveVirtualRouters.SlaveVirtualRouter[i]), types.YChild{"SlaveVirtualRouter", slaveVirtualRouters.SlaveVirtualRouter[i]})
    }
    slaveVirtualRouters.EntityData.Leafs = types.NewOrderedMap()

    slaveVirtualRouters.EntityData.YListKeys = []string {}

    return &(slaveVirtualRouters.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter
// The VRRP slave being configured
type Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Virtual Router ID. The type is interface{} with
    // range: 1..255.
    SlaveVirtualRouterId interface{}

    // VRRP Session name for this slave to follow. The type is string.
    Follow interface{}

    // Disable Accept Mode for this virtual IPAddress. The type is interface{}.
    AcceptModeDisable interface{}

    // The VRRP IPv6 virtual linklocal address.
    LinkLocalIpv6Address Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address

    // The table of VRRP virtual global IPv6 addresses.
    GlobalIpv6Addresses Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses
}

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter) GetEntityData() *types.CommonEntityData {
    slaveVirtualRouter.EntityData.YFilter = slaveVirtualRouter.YFilter
    slaveVirtualRouter.EntityData.YangName = "slave-virtual-router"
    slaveVirtualRouter.EntityData.BundleName = "cisco_ios_xr"
    slaveVirtualRouter.EntityData.ParentYangName = "slave-virtual-routers"
    slaveVirtualRouter.EntityData.SegmentPath = "slave-virtual-router" + types.AddKeyToken(slaveVirtualRouter.SlaveVirtualRouterId, "slave-virtual-router-id")
    slaveVirtualRouter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv6/slave-virtual-routers/" + slaveVirtualRouter.EntityData.SegmentPath
    slaveVirtualRouter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slaveVirtualRouter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slaveVirtualRouter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slaveVirtualRouter.EntityData.Children = types.NewOrderedMap()
    slaveVirtualRouter.EntityData.Children.Append("link-local-ipv6-address", types.YChild{"LinkLocalIpv6Address", &slaveVirtualRouter.LinkLocalIpv6Address})
    slaveVirtualRouter.EntityData.Children.Append("global-ipv6-addresses", types.YChild{"GlobalIpv6Addresses", &slaveVirtualRouter.GlobalIpv6Addresses})
    slaveVirtualRouter.EntityData.Leafs = types.NewOrderedMap()
    slaveVirtualRouter.EntityData.Leafs.Append("slave-virtual-router-id", types.YLeaf{"SlaveVirtualRouterId", slaveVirtualRouter.SlaveVirtualRouterId})
    slaveVirtualRouter.EntityData.Leafs.Append("follow", types.YLeaf{"Follow", slaveVirtualRouter.Follow})
    slaveVirtualRouter.EntityData.Leafs.Append("accept-mode-disable", types.YLeaf{"AcceptModeDisable", slaveVirtualRouter.AcceptModeDisable})

    slaveVirtualRouter.EntityData.YListKeys = []string {"SlaveVirtualRouterId"}

    return &(slaveVirtualRouter.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address
// The VRRP IPv6 virtual linklocal address
type Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRRP IPv6 virtual linklocal address. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}

    // TRUE if the virtual linklocal address is to be autoconfigured FALSE if an
    // IPv6 virtual linklocal address is configured. The type is bool. The default
    // value is false.
    AutoConfigure interface{}
}

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address) GetEntityData() *types.CommonEntityData {
    linkLocalIpv6Address.EntityData.YFilter = linkLocalIpv6Address.YFilter
    linkLocalIpv6Address.EntityData.YangName = "link-local-ipv6-address"
    linkLocalIpv6Address.EntityData.BundleName = "cisco_ios_xr"
    linkLocalIpv6Address.EntityData.ParentYangName = "slave-virtual-router"
    linkLocalIpv6Address.EntityData.SegmentPath = "link-local-ipv6-address"
    linkLocalIpv6Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv6/slave-virtual-routers/slave-virtual-router/" + linkLocalIpv6Address.EntityData.SegmentPath
    linkLocalIpv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkLocalIpv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkLocalIpv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkLocalIpv6Address.EntityData.Children = types.NewOrderedMap()
    linkLocalIpv6Address.EntityData.Leafs = types.NewOrderedMap()
    linkLocalIpv6Address.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", linkLocalIpv6Address.IpAddress})
    linkLocalIpv6Address.EntityData.Leafs.Append("auto-configure", types.YLeaf{"AutoConfigure", linkLocalIpv6Address.AutoConfigure})

    linkLocalIpv6Address.EntityData.YListKeys = []string {}

    return &(linkLocalIpv6Address.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses
// The table of VRRP virtual global IPv6
// addresses
type Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VRRP virtual global IPv6 IP address. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address.
    GlobalIpv6Address []*Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address
}

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses) GetEntityData() *types.CommonEntityData {
    globalIpv6Addresses.EntityData.YFilter = globalIpv6Addresses.YFilter
    globalIpv6Addresses.EntityData.YangName = "global-ipv6-addresses"
    globalIpv6Addresses.EntityData.BundleName = "cisco_ios_xr"
    globalIpv6Addresses.EntityData.ParentYangName = "slave-virtual-router"
    globalIpv6Addresses.EntityData.SegmentPath = "global-ipv6-addresses"
    globalIpv6Addresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv6/slave-virtual-routers/slave-virtual-router/" + globalIpv6Addresses.EntityData.SegmentPath
    globalIpv6Addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalIpv6Addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalIpv6Addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalIpv6Addresses.EntityData.Children = types.NewOrderedMap()
    globalIpv6Addresses.EntityData.Children.Append("global-ipv6-address", types.YChild{"GlobalIpv6Address", nil})
    for i := range globalIpv6Addresses.GlobalIpv6Address {
        globalIpv6Addresses.EntityData.Children.Append(types.GetSegmentPath(globalIpv6Addresses.GlobalIpv6Address[i]), types.YChild{"GlobalIpv6Address", globalIpv6Addresses.GlobalIpv6Address[i]})
    }
    globalIpv6Addresses.EntityData.Leafs = types.NewOrderedMap()

    globalIpv6Addresses.EntityData.YListKeys = []string {}

    return &(globalIpv6Addresses.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address
// A VRRP virtual global IPv6 IP address
type Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRRP virtual global IPv6 address. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}
}

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetEntityData() *types.CommonEntityData {
    globalIpv6Address.EntityData.YFilter = globalIpv6Address.YFilter
    globalIpv6Address.EntityData.YangName = "global-ipv6-address"
    globalIpv6Address.EntityData.BundleName = "cisco_ios_xr"
    globalIpv6Address.EntityData.ParentYangName = "global-ipv6-addresses"
    globalIpv6Address.EntityData.SegmentPath = "global-ipv6-address" + types.AddKeyToken(globalIpv6Address.IpAddress, "ip-address")
    globalIpv6Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv6/slave-virtual-routers/slave-virtual-router/global-ipv6-addresses/" + globalIpv6Address.EntityData.SegmentPath
    globalIpv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalIpv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalIpv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalIpv6Address.EntityData.Children = types.NewOrderedMap()
    globalIpv6Address.EntityData.Leafs = types.NewOrderedMap()
    globalIpv6Address.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", globalIpv6Address.IpAddress})

    globalIpv6Address.EntityData.YListKeys = []string {"IpAddress"}

    return &(globalIpv6Address.EntityData)
}

// Vrrp_Interfaces_Interface_Delay
// Minimum and Reload Delay
// This type is a presence type.
type Vrrp_Interfaces_Interface_Delay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Minimum delay in seconds. The type is interface{} with range: 0..10000.
    // This attribute is mandatory. Units are second.
    MinDelay interface{}

    // Reload delay in seconds. The type is interface{} with range: 0..10000. This
    // attribute is mandatory. Units are second.
    ReloadDelay interface{}
}

func (delay *Vrrp_Interfaces_Interface_Delay) GetEntityData() *types.CommonEntityData {
    delay.EntityData.YFilter = delay.YFilter
    delay.EntityData.YangName = "delay"
    delay.EntityData.BundleName = "cisco_ios_xr"
    delay.EntityData.ParentYangName = "interface"
    delay.EntityData.SegmentPath = "delay"
    delay.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/" + delay.EntityData.SegmentPath
    delay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delay.EntityData.Children = types.NewOrderedMap()
    delay.EntityData.Leafs = types.NewOrderedMap()
    delay.EntityData.Leafs.Append("min-delay", types.YLeaf{"MinDelay", delay.MinDelay})
    delay.EntityData.Leafs.Append("reload-delay", types.YLeaf{"ReloadDelay", delay.ReloadDelay})

    delay.EntityData.YListKeys = []string {}

    return &(delay.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4
// IPv4 VRRP configuration
type Vrrp_Interfaces_Interface_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version 3 VRRP configuration.
    Version3 Vrrp_Interfaces_Interface_Ipv4_Version3

    // The VRRP slave group configuration table.
    SlaveVirtualRouters Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters

    // Version 2 VRRP configuration.
    Version2 Vrrp_Interfaces_Interface_Ipv4_Version2
}

func (ipv4 *Vrrp_Interfaces_Interface_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "interface"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/" + ipv4.EntityData.SegmentPath
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = types.NewOrderedMap()
    ipv4.EntityData.Children.Append("version3", types.YChild{"Version3", &ipv4.Version3})
    ipv4.EntityData.Children.Append("slave-virtual-routers", types.YChild{"SlaveVirtualRouters", &ipv4.SlaveVirtualRouters})
    ipv4.EntityData.Children.Append("version2", types.YChild{"Version2", &ipv4.Version2})
    ipv4.EntityData.Leafs = types.NewOrderedMap()

    ipv4.EntityData.YListKeys = []string {}

    return &(ipv4.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version3
// Version 3 VRRP configuration
type Vrrp_Interfaces_Interface_Ipv4_Version3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The VRRP virtual router configuration table.
    VirtualRouters Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters
}

func (version3 *Vrrp_Interfaces_Interface_Ipv4_Version3) GetEntityData() *types.CommonEntityData {
    version3.EntityData.YFilter = version3.YFilter
    version3.EntityData.YangName = "version3"
    version3.EntityData.BundleName = "cisco_ios_xr"
    version3.EntityData.ParentYangName = "ipv4"
    version3.EntityData.SegmentPath = "version3"
    version3.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/" + version3.EntityData.SegmentPath
    version3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    version3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    version3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    version3.EntityData.Children = types.NewOrderedMap()
    version3.EntityData.Children.Append("virtual-routers", types.YChild{"VirtualRouters", &version3.VirtualRouters})
    version3.EntityData.Leafs = types.NewOrderedMap()

    version3.EntityData.YListKeys = []string {}

    return &(version3.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters
// The VRRP virtual router configuration table
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The VRRP virtual router being configured. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter.
    VirtualRouter []*Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters) GetEntityData() *types.CommonEntityData {
    virtualRouters.EntityData.YFilter = virtualRouters.YFilter
    virtualRouters.EntityData.YangName = "virtual-routers"
    virtualRouters.EntityData.BundleName = "cisco_ios_xr"
    virtualRouters.EntityData.ParentYangName = "version3"
    virtualRouters.EntityData.SegmentPath = "virtual-routers"
    virtualRouters.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version3/" + virtualRouters.EntityData.SegmentPath
    virtualRouters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualRouters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualRouters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualRouters.EntityData.Children = types.NewOrderedMap()
    virtualRouters.EntityData.Children.Append("virtual-router", types.YChild{"VirtualRouter", nil})
    for i := range virtualRouters.VirtualRouter {
        virtualRouters.EntityData.Children.Append(types.GetSegmentPath(virtualRouters.VirtualRouter[i]), types.YChild{"VirtualRouter", virtualRouters.VirtualRouter[i]})
    }
    virtualRouters.EntityData.Leafs = types.NewOrderedMap()

    virtualRouters.EntityData.YListKeys = []string {}

    return &(virtualRouters.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter
// The VRRP virtual router being configured
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRID Virtual Router Identifier. The type is
    // interface{} with range: 1..255.
    VrId interface{}

    // VRRP Session Name. The type is string with length: 1..16.
    SessionName interface{}

    // Enable use of Bidirectional Forwarding Detection for this IP. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Bfd interface{}

    // The Primary VRRP IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PrimaryIpv4Address interface{}

    // Preempt Master router if higher priority. The type is interface{} with
    // range: 0..3600. The default value is 0.
    Preempt interface{}

    // Disable Accept Mode for this virtual IPAddress. The type is interface{}.
    AcceptModeDisable interface{}

    // Priority value. The type is interface{} with range: 1..254. The default
    // value is 100.
    Priority interface{}

    // Set advertisement timer.
    Timer Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer

    // The table of VRRP secondary IPv4 addresses.
    SecondaryIpv4Addresses Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses

    // Track an object, reducing priority if it goes down.
    TrackedObjects Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects

    // Track an item, reducing priority if it goes down.
    Tracks Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter) GetEntityData() *types.CommonEntityData {
    virtualRouter.EntityData.YFilter = virtualRouter.YFilter
    virtualRouter.EntityData.YangName = "virtual-router"
    virtualRouter.EntityData.BundleName = "cisco_ios_xr"
    virtualRouter.EntityData.ParentYangName = "virtual-routers"
    virtualRouter.EntityData.SegmentPath = "virtual-router" + types.AddKeyToken(virtualRouter.VrId, "vr-id")
    virtualRouter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version3/virtual-routers/" + virtualRouter.EntityData.SegmentPath
    virtualRouter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualRouter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualRouter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualRouter.EntityData.Children = types.NewOrderedMap()
    virtualRouter.EntityData.Children.Append("timer", types.YChild{"Timer", &virtualRouter.Timer})
    virtualRouter.EntityData.Children.Append("secondary-ipv4-addresses", types.YChild{"SecondaryIpv4Addresses", &virtualRouter.SecondaryIpv4Addresses})
    virtualRouter.EntityData.Children.Append("tracked-objects", types.YChild{"TrackedObjects", &virtualRouter.TrackedObjects})
    virtualRouter.EntityData.Children.Append("tracks", types.YChild{"Tracks", &virtualRouter.Tracks})
    virtualRouter.EntityData.Leafs = types.NewOrderedMap()
    virtualRouter.EntityData.Leafs.Append("vr-id", types.YLeaf{"VrId", virtualRouter.VrId})
    virtualRouter.EntityData.Leafs.Append("session-name", types.YLeaf{"SessionName", virtualRouter.SessionName})
    virtualRouter.EntityData.Leafs.Append("bfd", types.YLeaf{"Bfd", virtualRouter.Bfd})
    virtualRouter.EntityData.Leafs.Append("primary-ipv4-address", types.YLeaf{"PrimaryIpv4Address", virtualRouter.PrimaryIpv4Address})
    virtualRouter.EntityData.Leafs.Append("preempt", types.YLeaf{"Preempt", virtualRouter.Preempt})
    virtualRouter.EntityData.Leafs.Append("accept-mode-disable", types.YLeaf{"AcceptModeDisable", virtualRouter.AcceptModeDisable})
    virtualRouter.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", virtualRouter.Priority})

    virtualRouter.EntityData.YListKeys = []string {"VrId"}

    return &(virtualRouter.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer
// Set advertisement timer
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Advertisement time in milliseconds. The type is interface{} with range:
    // 100..40950. Units are millisecond.
    AdvertisementTimeInMsec interface{}

    // Advertisement time in seconds. The type is interface{} with range: 1..40.
    // Units are second.
    AdvertisementTimeInSec interface{}

    // TRUE - Force configured timer values to be used, required when configured
    // in milliseconds. The type is interface{}. Units are millisecond.
    Forced interface{}
}

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer) GetEntityData() *types.CommonEntityData {
    timer.EntityData.YFilter = timer.YFilter
    timer.EntityData.YangName = "timer"
    timer.EntityData.BundleName = "cisco_ios_xr"
    timer.EntityData.ParentYangName = "virtual-router"
    timer.EntityData.SegmentPath = "timer"
    timer.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version3/virtual-routers/virtual-router/" + timer.EntityData.SegmentPath
    timer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timer.EntityData.Children = types.NewOrderedMap()
    timer.EntityData.Leafs = types.NewOrderedMap()
    timer.EntityData.Leafs.Append("advertisement-time-in-msec", types.YLeaf{"AdvertisementTimeInMsec", timer.AdvertisementTimeInMsec})
    timer.EntityData.Leafs.Append("advertisement-time-in-sec", types.YLeaf{"AdvertisementTimeInSec", timer.AdvertisementTimeInSec})
    timer.EntityData.Leafs.Append("forced", types.YLeaf{"Forced", timer.Forced})

    timer.EntityData.YListKeys = []string {}

    return &(timer.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses
// The table of VRRP secondary IPv4 addresses
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VRRP secondary IPv4 address. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address.
    SecondaryIpv4Address []*Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetEntityData() *types.CommonEntityData {
    secondaryIpv4Addresses.EntityData.YFilter = secondaryIpv4Addresses.YFilter
    secondaryIpv4Addresses.EntityData.YangName = "secondary-ipv4-addresses"
    secondaryIpv4Addresses.EntityData.BundleName = "cisco_ios_xr"
    secondaryIpv4Addresses.EntityData.ParentYangName = "virtual-router"
    secondaryIpv4Addresses.EntityData.SegmentPath = "secondary-ipv4-addresses"
    secondaryIpv4Addresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version3/virtual-routers/virtual-router/" + secondaryIpv4Addresses.EntityData.SegmentPath
    secondaryIpv4Addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryIpv4Addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryIpv4Addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryIpv4Addresses.EntityData.Children = types.NewOrderedMap()
    secondaryIpv4Addresses.EntityData.Children.Append("secondary-ipv4-address", types.YChild{"SecondaryIpv4Address", nil})
    for i := range secondaryIpv4Addresses.SecondaryIpv4Address {
        secondaryIpv4Addresses.EntityData.Children.Append(types.GetSegmentPath(secondaryIpv4Addresses.SecondaryIpv4Address[i]), types.YChild{"SecondaryIpv4Address", secondaryIpv4Addresses.SecondaryIpv4Address[i]})
    }
    secondaryIpv4Addresses.EntityData.Leafs = types.NewOrderedMap()

    secondaryIpv4Addresses.EntityData.YListKeys = []string {}

    return &(secondaryIpv4Addresses.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address
// A VRRP secondary IPv4 address
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRRP Secondary IPv4 address. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetEntityData() *types.CommonEntityData {
    secondaryIpv4Address.EntityData.YFilter = secondaryIpv4Address.YFilter
    secondaryIpv4Address.EntityData.YangName = "secondary-ipv4-address"
    secondaryIpv4Address.EntityData.BundleName = "cisco_ios_xr"
    secondaryIpv4Address.EntityData.ParentYangName = "secondary-ipv4-addresses"
    secondaryIpv4Address.EntityData.SegmentPath = "secondary-ipv4-address" + types.AddKeyToken(secondaryIpv4Address.IpAddress, "ip-address")
    secondaryIpv4Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version3/virtual-routers/virtual-router/secondary-ipv4-addresses/" + secondaryIpv4Address.EntityData.SegmentPath
    secondaryIpv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryIpv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryIpv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryIpv4Address.EntityData.Children = types.NewOrderedMap()
    secondaryIpv4Address.EntityData.Leafs = types.NewOrderedMap()
    secondaryIpv4Address.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", secondaryIpv4Address.IpAddress})

    secondaryIpv4Address.EntityData.YListKeys = []string {"IpAddress"}

    return &(secondaryIpv4Address.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects
// Track an object, reducing priority if it
// goes down
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object to be tracked. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject.
    TrackedObject []*Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetEntityData() *types.CommonEntityData {
    trackedObjects.EntityData.YFilter = trackedObjects.YFilter
    trackedObjects.EntityData.YangName = "tracked-objects"
    trackedObjects.EntityData.BundleName = "cisco_ios_xr"
    trackedObjects.EntityData.ParentYangName = "virtual-router"
    trackedObjects.EntityData.SegmentPath = "tracked-objects"
    trackedObjects.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version3/virtual-routers/virtual-router/" + trackedObjects.EntityData.SegmentPath
    trackedObjects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedObjects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedObjects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedObjects.EntityData.Children = types.NewOrderedMap()
    trackedObjects.EntityData.Children.Append("tracked-object", types.YChild{"TrackedObject", nil})
    for i := range trackedObjects.TrackedObject {
        trackedObjects.EntityData.Children.Append(types.GetSegmentPath(trackedObjects.TrackedObject[i]), types.YChild{"TrackedObject", trackedObjects.TrackedObject[i]})
    }
    trackedObjects.EntityData.Leafs = types.NewOrderedMap()

    trackedObjects.EntityData.YListKeys = []string {}

    return &(trackedObjects.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject
// Object to be tracked
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Object to be tracked, interface name for
    // interfaces. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ObjectName interface{}

    // Priority decrement. The type is interface{} with range: 1..254. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetEntityData() *types.CommonEntityData {
    trackedObject.EntityData.YFilter = trackedObject.YFilter
    trackedObject.EntityData.YangName = "tracked-object"
    trackedObject.EntityData.BundleName = "cisco_ios_xr"
    trackedObject.EntityData.ParentYangName = "tracked-objects"
    trackedObject.EntityData.SegmentPath = "tracked-object" + types.AddKeyToken(trackedObject.ObjectName, "object-name")
    trackedObject.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version3/virtual-routers/virtual-router/tracked-objects/" + trackedObject.EntityData.SegmentPath
    trackedObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedObject.EntityData.Children = types.NewOrderedMap()
    trackedObject.EntityData.Leafs = types.NewOrderedMap()
    trackedObject.EntityData.Leafs.Append("object-name", types.YLeaf{"ObjectName", trackedObject.ObjectName})
    trackedObject.EntityData.Leafs.Append("priority-decrement", types.YLeaf{"PriorityDecrement", trackedObject.PriorityDecrement})

    trackedObject.EntityData.YListKeys = []string {"ObjectName"}

    return &(trackedObject.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks
// Track an item, reducing priority if it
// goes down
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object to be tracked. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track.
    Track []*Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track
}

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks) GetEntityData() *types.CommonEntityData {
    tracks.EntityData.YFilter = tracks.YFilter
    tracks.EntityData.YangName = "tracks"
    tracks.EntityData.BundleName = "cisco_ios_xr"
    tracks.EntityData.ParentYangName = "virtual-router"
    tracks.EntityData.SegmentPath = "tracks"
    tracks.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version3/virtual-routers/virtual-router/" + tracks.EntityData.SegmentPath
    tracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tracks.EntityData.Children = types.NewOrderedMap()
    tracks.EntityData.Children.Append("track", types.YChild{"Track", nil})
    for i := range tracks.Track {
        tracks.EntityData.Children.Append(types.GetSegmentPath(tracks.Track[i]), types.YChild{"Track", tracks.Track[i]})
    }
    tracks.EntityData.Leafs = types.NewOrderedMap()

    tracks.EntityData.YListKeys = []string {}

    return &(tracks.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track
// Object to be tracked
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Object to be tracked, interface name for
    // interfaces. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Priority decrement. The type is interface{} with range: 1..254. This
    // attribute is mandatory.
    Priority interface{}
}

func (track *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetEntityData() *types.CommonEntityData {
    track.EntityData.YFilter = track.YFilter
    track.EntityData.YangName = "track"
    track.EntityData.BundleName = "cisco_ios_xr"
    track.EntityData.ParentYangName = "tracks"
    track.EntityData.SegmentPath = "track" + types.AddKeyToken(track.InterfaceName, "interface-name")
    track.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version3/virtual-routers/virtual-router/tracks/" + track.EntityData.SegmentPath
    track.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    track.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    track.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    track.EntityData.Children = types.NewOrderedMap()
    track.EntityData.Leafs = types.NewOrderedMap()
    track.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", track.InterfaceName})
    track.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", track.Priority})

    track.EntityData.YListKeys = []string {"InterfaceName"}

    return &(track.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters
// The VRRP slave group configuration table
type Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The VRRP slave being configured. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter.
    SlaveVirtualRouter []*Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter
}

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters) GetEntityData() *types.CommonEntityData {
    slaveVirtualRouters.EntityData.YFilter = slaveVirtualRouters.YFilter
    slaveVirtualRouters.EntityData.YangName = "slave-virtual-routers"
    slaveVirtualRouters.EntityData.BundleName = "cisco_ios_xr"
    slaveVirtualRouters.EntityData.ParentYangName = "ipv4"
    slaveVirtualRouters.EntityData.SegmentPath = "slave-virtual-routers"
    slaveVirtualRouters.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/" + slaveVirtualRouters.EntityData.SegmentPath
    slaveVirtualRouters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slaveVirtualRouters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slaveVirtualRouters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slaveVirtualRouters.EntityData.Children = types.NewOrderedMap()
    slaveVirtualRouters.EntityData.Children.Append("slave-virtual-router", types.YChild{"SlaveVirtualRouter", nil})
    for i := range slaveVirtualRouters.SlaveVirtualRouter {
        slaveVirtualRouters.EntityData.Children.Append(types.GetSegmentPath(slaveVirtualRouters.SlaveVirtualRouter[i]), types.YChild{"SlaveVirtualRouter", slaveVirtualRouters.SlaveVirtualRouter[i]})
    }
    slaveVirtualRouters.EntityData.Leafs = types.NewOrderedMap()

    slaveVirtualRouters.EntityData.YListKeys = []string {}

    return &(slaveVirtualRouters.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter
// The VRRP slave being configured
type Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Virtual Router ID. The type is interface{} with
    // range: 1..255.
    SlaveVirtualRouterId interface{}

    // VRRP Session name for this slave to follow. The type is string.
    Follow interface{}

    // Disable Accept Mode for this virtual IPAddress. The type is interface{}.
    AcceptModeDisable interface{}

    // The Primary VRRP IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PrimaryIpv4Address interface{}

    // The table of VRRP secondary IPv4 addresses.
    SecondaryIpv4Addresses Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses
}

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter) GetEntityData() *types.CommonEntityData {
    slaveVirtualRouter.EntityData.YFilter = slaveVirtualRouter.YFilter
    slaveVirtualRouter.EntityData.YangName = "slave-virtual-router"
    slaveVirtualRouter.EntityData.BundleName = "cisco_ios_xr"
    slaveVirtualRouter.EntityData.ParentYangName = "slave-virtual-routers"
    slaveVirtualRouter.EntityData.SegmentPath = "slave-virtual-router" + types.AddKeyToken(slaveVirtualRouter.SlaveVirtualRouterId, "slave-virtual-router-id")
    slaveVirtualRouter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/slave-virtual-routers/" + slaveVirtualRouter.EntityData.SegmentPath
    slaveVirtualRouter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slaveVirtualRouter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slaveVirtualRouter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slaveVirtualRouter.EntityData.Children = types.NewOrderedMap()
    slaveVirtualRouter.EntityData.Children.Append("secondary-ipv4-addresses", types.YChild{"SecondaryIpv4Addresses", &slaveVirtualRouter.SecondaryIpv4Addresses})
    slaveVirtualRouter.EntityData.Leafs = types.NewOrderedMap()
    slaveVirtualRouter.EntityData.Leafs.Append("slave-virtual-router-id", types.YLeaf{"SlaveVirtualRouterId", slaveVirtualRouter.SlaveVirtualRouterId})
    slaveVirtualRouter.EntityData.Leafs.Append("follow", types.YLeaf{"Follow", slaveVirtualRouter.Follow})
    slaveVirtualRouter.EntityData.Leafs.Append("accept-mode-disable", types.YLeaf{"AcceptModeDisable", slaveVirtualRouter.AcceptModeDisable})
    slaveVirtualRouter.EntityData.Leafs.Append("primary-ipv4-address", types.YLeaf{"PrimaryIpv4Address", slaveVirtualRouter.PrimaryIpv4Address})

    slaveVirtualRouter.EntityData.YListKeys = []string {"SlaveVirtualRouterId"}

    return &(slaveVirtualRouter.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses
// The table of VRRP secondary IPv4 addresses
type Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VRRP secondary IPv4 address. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address.
    SecondaryIpv4Address []*Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses) GetEntityData() *types.CommonEntityData {
    secondaryIpv4Addresses.EntityData.YFilter = secondaryIpv4Addresses.YFilter
    secondaryIpv4Addresses.EntityData.YangName = "secondary-ipv4-addresses"
    secondaryIpv4Addresses.EntityData.BundleName = "cisco_ios_xr"
    secondaryIpv4Addresses.EntityData.ParentYangName = "slave-virtual-router"
    secondaryIpv4Addresses.EntityData.SegmentPath = "secondary-ipv4-addresses"
    secondaryIpv4Addresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/slave-virtual-routers/slave-virtual-router/" + secondaryIpv4Addresses.EntityData.SegmentPath
    secondaryIpv4Addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryIpv4Addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryIpv4Addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryIpv4Addresses.EntityData.Children = types.NewOrderedMap()
    secondaryIpv4Addresses.EntityData.Children.Append("secondary-ipv4-address", types.YChild{"SecondaryIpv4Address", nil})
    for i := range secondaryIpv4Addresses.SecondaryIpv4Address {
        secondaryIpv4Addresses.EntityData.Children.Append(types.GetSegmentPath(secondaryIpv4Addresses.SecondaryIpv4Address[i]), types.YChild{"SecondaryIpv4Address", secondaryIpv4Addresses.SecondaryIpv4Address[i]})
    }
    secondaryIpv4Addresses.EntityData.Leafs = types.NewOrderedMap()

    secondaryIpv4Addresses.EntityData.YListKeys = []string {}

    return &(secondaryIpv4Addresses.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address
// A VRRP secondary IPv4 address
type Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRRP Secondary IPv4 address. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetEntityData() *types.CommonEntityData {
    secondaryIpv4Address.EntityData.YFilter = secondaryIpv4Address.YFilter
    secondaryIpv4Address.EntityData.YangName = "secondary-ipv4-address"
    secondaryIpv4Address.EntityData.BundleName = "cisco_ios_xr"
    secondaryIpv4Address.EntityData.ParentYangName = "secondary-ipv4-addresses"
    secondaryIpv4Address.EntityData.SegmentPath = "secondary-ipv4-address" + types.AddKeyToken(secondaryIpv4Address.IpAddress, "ip-address")
    secondaryIpv4Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/slave-virtual-routers/slave-virtual-router/secondary-ipv4-addresses/" + secondaryIpv4Address.EntityData.SegmentPath
    secondaryIpv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryIpv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryIpv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryIpv4Address.EntityData.Children = types.NewOrderedMap()
    secondaryIpv4Address.EntityData.Leafs = types.NewOrderedMap()
    secondaryIpv4Address.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", secondaryIpv4Address.IpAddress})

    secondaryIpv4Address.EntityData.YListKeys = []string {"IpAddress"}

    return &(secondaryIpv4Address.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version2
// Version 2 VRRP configuration
type Vrrp_Interfaces_Interface_Ipv4_Version2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The VRRP virtual router configuration table.
    VirtualRouters Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters
}

func (version2 *Vrrp_Interfaces_Interface_Ipv4_Version2) GetEntityData() *types.CommonEntityData {
    version2.EntityData.YFilter = version2.YFilter
    version2.EntityData.YangName = "version2"
    version2.EntityData.BundleName = "cisco_ios_xr"
    version2.EntityData.ParentYangName = "ipv4"
    version2.EntityData.SegmentPath = "version2"
    version2.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/" + version2.EntityData.SegmentPath
    version2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    version2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    version2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    version2.EntityData.Children = types.NewOrderedMap()
    version2.EntityData.Children.Append("virtual-routers", types.YChild{"VirtualRouters", &version2.VirtualRouters})
    version2.EntityData.Leafs = types.NewOrderedMap()

    version2.EntityData.YListKeys = []string {}

    return &(version2.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters
// The VRRP virtual router configuration table
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The VRRP virtual router being configured. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter.
    VirtualRouter []*Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters) GetEntityData() *types.CommonEntityData {
    virtualRouters.EntityData.YFilter = virtualRouters.YFilter
    virtualRouters.EntityData.YangName = "virtual-routers"
    virtualRouters.EntityData.BundleName = "cisco_ios_xr"
    virtualRouters.EntityData.ParentYangName = "version2"
    virtualRouters.EntityData.SegmentPath = "virtual-routers"
    virtualRouters.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version2/" + virtualRouters.EntityData.SegmentPath
    virtualRouters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualRouters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualRouters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualRouters.EntityData.Children = types.NewOrderedMap()
    virtualRouters.EntityData.Children.Append("virtual-router", types.YChild{"VirtualRouter", nil})
    for i := range virtualRouters.VirtualRouter {
        virtualRouters.EntityData.Children.Append(types.GetSegmentPath(virtualRouters.VirtualRouter[i]), types.YChild{"VirtualRouter", virtualRouters.VirtualRouter[i]})
    }
    virtualRouters.EntityData.Leafs = types.NewOrderedMap()

    virtualRouters.EntityData.YListKeys = []string {}

    return &(virtualRouters.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter
// The VRRP virtual router being configured
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRID Virtual Router Identifier. The type is
    // interface{} with range: 1..255.
    VrId interface{}

    // Priority value. The type is interface{} with range: 1..254. The default
    // value is 100.
    Priority interface{}

    // The Primary VRRP IPv4 address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PrimaryIpv4Address interface{}

    // Preempt Master router if higher priority. The type is interface{} with
    // range: 0..3600. The default value is 0.
    Preempt interface{}

    // Authentication password, 8 chars max. The type is string.
    TextPassword interface{}

    // Enable use of Bidirectional Forwarding Detection for this IP. The type is
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Bfd interface{}

    // VRRP Session Name. The type is string with length: 1..16.
    SessionName interface{}

    // Disable Accept Mode for this virtual IPAddress. The type is interface{}.
    AcceptModeDisable interface{}

    // Set advertisement timer.
    Timer Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer

    // The table of VRRP secondary IPv4 addresses.
    SecondaryIpv4Addresses Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses

    // Track an item, reducing priority if it goes down.
    Tracks Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks

    // Track an object, reducing priority if it goes down.
    TrackedObjects Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter) GetEntityData() *types.CommonEntityData {
    virtualRouter.EntityData.YFilter = virtualRouter.YFilter
    virtualRouter.EntityData.YangName = "virtual-router"
    virtualRouter.EntityData.BundleName = "cisco_ios_xr"
    virtualRouter.EntityData.ParentYangName = "virtual-routers"
    virtualRouter.EntityData.SegmentPath = "virtual-router" + types.AddKeyToken(virtualRouter.VrId, "vr-id")
    virtualRouter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version2/virtual-routers/" + virtualRouter.EntityData.SegmentPath
    virtualRouter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualRouter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualRouter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualRouter.EntityData.Children = types.NewOrderedMap()
    virtualRouter.EntityData.Children.Append("timer", types.YChild{"Timer", &virtualRouter.Timer})
    virtualRouter.EntityData.Children.Append("secondary-ipv4-addresses", types.YChild{"SecondaryIpv4Addresses", &virtualRouter.SecondaryIpv4Addresses})
    virtualRouter.EntityData.Children.Append("tracks", types.YChild{"Tracks", &virtualRouter.Tracks})
    virtualRouter.EntityData.Children.Append("tracked-objects", types.YChild{"TrackedObjects", &virtualRouter.TrackedObjects})
    virtualRouter.EntityData.Leafs = types.NewOrderedMap()
    virtualRouter.EntityData.Leafs.Append("vr-id", types.YLeaf{"VrId", virtualRouter.VrId})
    virtualRouter.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", virtualRouter.Priority})
    virtualRouter.EntityData.Leafs.Append("primary-ipv4-address", types.YLeaf{"PrimaryIpv4Address", virtualRouter.PrimaryIpv4Address})
    virtualRouter.EntityData.Leafs.Append("preempt", types.YLeaf{"Preempt", virtualRouter.Preempt})
    virtualRouter.EntityData.Leafs.Append("text-password", types.YLeaf{"TextPassword", virtualRouter.TextPassword})
    virtualRouter.EntityData.Leafs.Append("bfd", types.YLeaf{"Bfd", virtualRouter.Bfd})
    virtualRouter.EntityData.Leafs.Append("session-name", types.YLeaf{"SessionName", virtualRouter.SessionName})
    virtualRouter.EntityData.Leafs.Append("accept-mode-disable", types.YLeaf{"AcceptModeDisable", virtualRouter.AcceptModeDisable})

    virtualRouter.EntityData.YListKeys = []string {"VrId"}

    return &(virtualRouter.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer
// Set advertisement timer
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Advertisement time in milliseconds. The type is interface{} with range:
    // 100..40950. Units are millisecond.
    AdvertisementTimeInMsec interface{}

    // Advertisement time in seconds. The type is interface{} with range: 1..255.
    // Units are second.
    AdvertisementTimeInSec interface{}

    // TRUE - Force configured timer values to be used, required when configured
    // in milliseconds. The type is interface{}. Units are millisecond.
    Forced interface{}
}

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer) GetEntityData() *types.CommonEntityData {
    timer.EntityData.YFilter = timer.YFilter
    timer.EntityData.YangName = "timer"
    timer.EntityData.BundleName = "cisco_ios_xr"
    timer.EntityData.ParentYangName = "virtual-router"
    timer.EntityData.SegmentPath = "timer"
    timer.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version2/virtual-routers/virtual-router/" + timer.EntityData.SegmentPath
    timer.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timer.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timer.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timer.EntityData.Children = types.NewOrderedMap()
    timer.EntityData.Leafs = types.NewOrderedMap()
    timer.EntityData.Leafs.Append("advertisement-time-in-msec", types.YLeaf{"AdvertisementTimeInMsec", timer.AdvertisementTimeInMsec})
    timer.EntityData.Leafs.Append("advertisement-time-in-sec", types.YLeaf{"AdvertisementTimeInSec", timer.AdvertisementTimeInSec})
    timer.EntityData.Leafs.Append("forced", types.YLeaf{"Forced", timer.Forced})

    timer.EntityData.YListKeys = []string {}

    return &(timer.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses
// The table of VRRP secondary IPv4 addresses
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A VRRP secondary IPv4 address. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address.
    SecondaryIpv4Address []*Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetEntityData() *types.CommonEntityData {
    secondaryIpv4Addresses.EntityData.YFilter = secondaryIpv4Addresses.YFilter
    secondaryIpv4Addresses.EntityData.YangName = "secondary-ipv4-addresses"
    secondaryIpv4Addresses.EntityData.BundleName = "cisco_ios_xr"
    secondaryIpv4Addresses.EntityData.ParentYangName = "virtual-router"
    secondaryIpv4Addresses.EntityData.SegmentPath = "secondary-ipv4-addresses"
    secondaryIpv4Addresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version2/virtual-routers/virtual-router/" + secondaryIpv4Addresses.EntityData.SegmentPath
    secondaryIpv4Addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryIpv4Addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryIpv4Addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryIpv4Addresses.EntityData.Children = types.NewOrderedMap()
    secondaryIpv4Addresses.EntityData.Children.Append("secondary-ipv4-address", types.YChild{"SecondaryIpv4Address", nil})
    for i := range secondaryIpv4Addresses.SecondaryIpv4Address {
        secondaryIpv4Addresses.EntityData.Children.Append(types.GetSegmentPath(secondaryIpv4Addresses.SecondaryIpv4Address[i]), types.YChild{"SecondaryIpv4Address", secondaryIpv4Addresses.SecondaryIpv4Address[i]})
    }
    secondaryIpv4Addresses.EntityData.Leafs = types.NewOrderedMap()

    secondaryIpv4Addresses.EntityData.YListKeys = []string {}

    return &(secondaryIpv4Addresses.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address
// A VRRP secondary IPv4 address
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRRP Secondary IPv4 address. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetEntityData() *types.CommonEntityData {
    secondaryIpv4Address.EntityData.YFilter = secondaryIpv4Address.YFilter
    secondaryIpv4Address.EntityData.YangName = "secondary-ipv4-address"
    secondaryIpv4Address.EntityData.BundleName = "cisco_ios_xr"
    secondaryIpv4Address.EntityData.ParentYangName = "secondary-ipv4-addresses"
    secondaryIpv4Address.EntityData.SegmentPath = "secondary-ipv4-address" + types.AddKeyToken(secondaryIpv4Address.IpAddress, "ip-address")
    secondaryIpv4Address.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version2/virtual-routers/virtual-router/secondary-ipv4-addresses/" + secondaryIpv4Address.EntityData.SegmentPath
    secondaryIpv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryIpv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryIpv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryIpv4Address.EntityData.Children = types.NewOrderedMap()
    secondaryIpv4Address.EntityData.Leafs = types.NewOrderedMap()
    secondaryIpv4Address.EntityData.Leafs.Append("ip-address", types.YLeaf{"IpAddress", secondaryIpv4Address.IpAddress})

    secondaryIpv4Address.EntityData.YListKeys = []string {"IpAddress"}

    return &(secondaryIpv4Address.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks
// Track an item, reducing priority if it
// goes down
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object to be tracked. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track.
    Track []*Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track
}

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks) GetEntityData() *types.CommonEntityData {
    tracks.EntityData.YFilter = tracks.YFilter
    tracks.EntityData.YangName = "tracks"
    tracks.EntityData.BundleName = "cisco_ios_xr"
    tracks.EntityData.ParentYangName = "virtual-router"
    tracks.EntityData.SegmentPath = "tracks"
    tracks.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version2/virtual-routers/virtual-router/" + tracks.EntityData.SegmentPath
    tracks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tracks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tracks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tracks.EntityData.Children = types.NewOrderedMap()
    tracks.EntityData.Children.Append("track", types.YChild{"Track", nil})
    for i := range tracks.Track {
        tracks.EntityData.Children.Append(types.GetSegmentPath(tracks.Track[i]), types.YChild{"Track", tracks.Track[i]})
    }
    tracks.EntityData.Leafs = types.NewOrderedMap()

    tracks.EntityData.YListKeys = []string {}

    return &(tracks.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track
// Object to be tracked
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Object to be tracked, interface name for
    // interfaces. The type is string with pattern: [a-zA-Z0-9._/-]+.
    InterfaceName interface{}

    // Priority decrement. The type is interface{} with range: 1..254. This
    // attribute is mandatory.
    Priority interface{}
}

func (track *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track) GetEntityData() *types.CommonEntityData {
    track.EntityData.YFilter = track.YFilter
    track.EntityData.YangName = "track"
    track.EntityData.BundleName = "cisco_ios_xr"
    track.EntityData.ParentYangName = "tracks"
    track.EntityData.SegmentPath = "track" + types.AddKeyToken(track.InterfaceName, "interface-name")
    track.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version2/virtual-routers/virtual-router/tracks/" + track.EntityData.SegmentPath
    track.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    track.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    track.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    track.EntityData.Children = types.NewOrderedMap()
    track.EntityData.Leafs = types.NewOrderedMap()
    track.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", track.InterfaceName})
    track.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", track.Priority})

    track.EntityData.YListKeys = []string {"InterfaceName"}

    return &(track.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects
// Track an object, reducing priority if it
// goes down
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object to be tracked. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject.
    TrackedObject []*Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects) GetEntityData() *types.CommonEntityData {
    trackedObjects.EntityData.YFilter = trackedObjects.YFilter
    trackedObjects.EntityData.YangName = "tracked-objects"
    trackedObjects.EntityData.BundleName = "cisco_ios_xr"
    trackedObjects.EntityData.ParentYangName = "virtual-router"
    trackedObjects.EntityData.SegmentPath = "tracked-objects"
    trackedObjects.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version2/virtual-routers/virtual-router/" + trackedObjects.EntityData.SegmentPath
    trackedObjects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedObjects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedObjects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedObjects.EntityData.Children = types.NewOrderedMap()
    trackedObjects.EntityData.Children.Append("tracked-object", types.YChild{"TrackedObject", nil})
    for i := range trackedObjects.TrackedObject {
        trackedObjects.EntityData.Children.Append(types.GetSegmentPath(trackedObjects.TrackedObject[i]), types.YChild{"TrackedObject", trackedObjects.TrackedObject[i]})
    }
    trackedObjects.EntityData.Leafs = types.NewOrderedMap()

    trackedObjects.EntityData.YListKeys = []string {}

    return &(trackedObjects.EntityData)
}

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject
// Object to be tracked
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Object to be tracked, interface name for
    // interfaces. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ObjectName interface{}

    // Priority decrement. The type is interface{} with range: 1..254. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetEntityData() *types.CommonEntityData {
    trackedObject.EntityData.YFilter = trackedObject.YFilter
    trackedObject.EntityData.YangName = "tracked-object"
    trackedObject.EntityData.BundleName = "cisco_ios_xr"
    trackedObject.EntityData.ParentYangName = "tracked-objects"
    trackedObject.EntityData.SegmentPath = "tracked-object" + types.AddKeyToken(trackedObject.ObjectName, "object-name")
    trackedObject.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/ipv4/version2/virtual-routers/virtual-router/tracked-objects/" + trackedObject.EntityData.SegmentPath
    trackedObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedObject.EntityData.Children = types.NewOrderedMap()
    trackedObject.EntityData.Leafs = types.NewOrderedMap()
    trackedObject.EntityData.Leafs.Append("object-name", types.YLeaf{"ObjectName", trackedObject.ObjectName})
    trackedObject.EntityData.Leafs.Append("priority-decrement", types.YLeaf{"PriorityDecrement", trackedObject.PriorityDecrement})

    trackedObject.EntityData.YListKeys = []string {"ObjectName"}

    return &(trackedObject.EntityData)
}

// Vrrp_Interfaces_Interface_Bfd
// BFD configuration
type Vrrp_Interfaces_Interface_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hello interval for BFD sessions created by vrrp. The type is interface{}
    // with range: 3..30000. Units are millisecond.
    Interval interface{}

    // Detection multiplier for BFD sessions created by vrrp. The type is
    // interface{} with range: 2..50.
    DetectionMultiplier interface{}
}

func (bfd *Vrrp_Interfaces_Interface_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "interface"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp/interfaces/interface/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

