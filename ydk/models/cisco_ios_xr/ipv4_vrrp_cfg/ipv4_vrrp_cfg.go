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
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // VRRP logging options.
    Logging Vrrp_Logging

    // Interface configuration table.
    Interfaces Vrrp_Interfaces
}

func (vrrp *Vrrp) GetFilter() yfilter.YFilter { return vrrp.YFilter }

func (vrrp *Vrrp) SetFilter(yf yfilter.YFilter) { vrrp.YFilter = yf }

func (vrrp *Vrrp) GetGoName(yname string) string {
    if yname == "logging" { return "Logging" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (vrrp *Vrrp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-vrrp-cfg:vrrp"
}

func (vrrp *Vrrp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "logging" {
        return &vrrp.Logging
    }
    if childYangName == "interfaces" {
        return &vrrp.Interfaces
    }
    return nil
}

func (vrrp *Vrrp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["logging"] = &vrrp.Logging
    children["interfaces"] = &vrrp.Interfaces
    return children
}

func (vrrp *Vrrp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrrp *Vrrp) GetBundleName() string { return "cisco_ios_xr" }

func (vrrp *Vrrp) GetYangName() string { return "vrrp" }

func (vrrp *Vrrp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrrp *Vrrp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrrp *Vrrp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrrp *Vrrp) SetParent(parent types.Entity) { vrrp.parent = parent }

func (vrrp *Vrrp) GetParent() types.Entity { return vrrp.parent }

func (vrrp *Vrrp) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-vrrp-cfg" }

// Vrrp_Logging
// VRRP logging options
type Vrrp_Logging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRRP state change IOS messages disable. The type is interface{}.
    StateChangeDisable interface{}
}

func (logging *Vrrp_Logging) GetFilter() yfilter.YFilter { return logging.YFilter }

func (logging *Vrrp_Logging) SetFilter(yf yfilter.YFilter) { logging.YFilter = yf }

func (logging *Vrrp_Logging) GetGoName(yname string) string {
    if yname == "state-change-disable" { return "StateChangeDisable" }
    return ""
}

func (logging *Vrrp_Logging) GetSegmentPath() string {
    return "logging"
}

func (logging *Vrrp_Logging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (logging *Vrrp_Logging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (logging *Vrrp_Logging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["state-change-disable"] = logging.StateChangeDisable
    return leafs
}

func (logging *Vrrp_Logging) GetBundleName() string { return "cisco_ios_xr" }

func (logging *Vrrp_Logging) GetYangName() string { return "logging" }

func (logging *Vrrp_Logging) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logging *Vrrp_Logging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logging *Vrrp_Logging) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logging *Vrrp_Logging) SetParent(parent types.Entity) { logging.parent = parent }

func (logging *Vrrp_Logging) GetParent() types.Entity { return logging.parent }

func (logging *Vrrp_Logging) GetParentYangName() string { return "vrrp" }

// Vrrp_Interfaces
// Interface configuration table
type Vrrp_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The interface being configured. The type is slice of
    // Vrrp_Interfaces_Interface.
    Interface []Vrrp_Interfaces_Interface
}

func (interfaces *Vrrp_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Vrrp_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Vrrp_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Vrrp_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Vrrp_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Vrrp_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Vrrp_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Vrrp_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Vrrp_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Vrrp_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Vrrp_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Vrrp_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Vrrp_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Vrrp_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Vrrp_Interfaces) GetParentYangName() string { return "vrrp" }

// Vrrp_Interfaces_Interface
// The interface being configured
type Vrrp_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name to configure. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
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

func (self *Vrrp_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Vrrp_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Vrrp_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "mac-refresh" { return "MacRefresh" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "delay" { return "Delay" }
    if yname == "ipv4" { return "Ipv4" }
    if yname == "bfd" { return "Bfd" }
    return ""
}

func (self *Vrrp_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Vrrp_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6" {
        return &self.Ipv6
    }
    if childYangName == "delay" {
        return &self.Delay
    }
    if childYangName == "ipv4" {
        return &self.Ipv4
    }
    if childYangName == "bfd" {
        return &self.Bfd
    }
    return nil
}

func (self *Vrrp_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6"] = &self.Ipv6
    children["delay"] = &self.Delay
    children["ipv4"] = &self.Ipv4
    children["bfd"] = &self.Bfd
    return children
}

func (self *Vrrp_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["mac-refresh"] = self.MacRefresh
    return leafs
}

func (self *Vrrp_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Vrrp_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Vrrp_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Vrrp_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Vrrp_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Vrrp_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Vrrp_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Vrrp_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Vrrp_Interfaces_Interface_Ipv6
// IPv6 VRRP configuration
type Vrrp_Interfaces_Interface_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version 3 VRRP configuration.
    Version3 Vrrp_Interfaces_Interface_Ipv6_Version3

    // The VRRP slave group configuration table.
    SlaveVirtualRouters Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters
}

func (ipv6 *Vrrp_Interfaces_Interface_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Vrrp_Interfaces_Interface_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Vrrp_Interfaces_Interface_Ipv6) GetGoName(yname string) string {
    if yname == "version3" { return "Version3" }
    if yname == "slave-virtual-routers" { return "SlaveVirtualRouters" }
    return ""
}

func (ipv6 *Vrrp_Interfaces_Interface_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Vrrp_Interfaces_Interface_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "version3" {
        return &ipv6.Version3
    }
    if childYangName == "slave-virtual-routers" {
        return &ipv6.SlaveVirtualRouters
    }
    return nil
}

func (ipv6 *Vrrp_Interfaces_Interface_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["version3"] = &ipv6.Version3
    children["slave-virtual-routers"] = &ipv6.SlaveVirtualRouters
    return children
}

func (ipv6 *Vrrp_Interfaces_Interface_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *Vrrp_Interfaces_Interface_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Vrrp_Interfaces_Interface_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Vrrp_Interfaces_Interface_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Vrrp_Interfaces_Interface_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Vrrp_Interfaces_Interface_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Vrrp_Interfaces_Interface_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Vrrp_Interfaces_Interface_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Vrrp_Interfaces_Interface_Ipv6) GetParentYangName() string { return "interface" }

// Vrrp_Interfaces_Interface_Ipv6_Version3
// Version 3 VRRP configuration
type Vrrp_Interfaces_Interface_Ipv6_Version3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The VRRP virtual router configuration table.
    VirtualRouters Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters
}

func (version3 *Vrrp_Interfaces_Interface_Ipv6_Version3) GetFilter() yfilter.YFilter { return version3.YFilter }

func (version3 *Vrrp_Interfaces_Interface_Ipv6_Version3) SetFilter(yf yfilter.YFilter) { version3.YFilter = yf }

func (version3 *Vrrp_Interfaces_Interface_Ipv6_Version3) GetGoName(yname string) string {
    if yname == "virtual-routers" { return "VirtualRouters" }
    return ""
}

func (version3 *Vrrp_Interfaces_Interface_Ipv6_Version3) GetSegmentPath() string {
    return "version3"
}

func (version3 *Vrrp_Interfaces_Interface_Ipv6_Version3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "virtual-routers" {
        return &version3.VirtualRouters
    }
    return nil
}

func (version3 *Vrrp_Interfaces_Interface_Ipv6_Version3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["virtual-routers"] = &version3.VirtualRouters
    return children
}

func (version3 *Vrrp_Interfaces_Interface_Ipv6_Version3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (version3 *Vrrp_Interfaces_Interface_Ipv6_Version3) GetBundleName() string { return "cisco_ios_xr" }

func (version3 *Vrrp_Interfaces_Interface_Ipv6_Version3) GetYangName() string { return "version3" }

func (version3 *Vrrp_Interfaces_Interface_Ipv6_Version3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (version3 *Vrrp_Interfaces_Interface_Ipv6_Version3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (version3 *Vrrp_Interfaces_Interface_Ipv6_Version3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (version3 *Vrrp_Interfaces_Interface_Ipv6_Version3) SetParent(parent types.Entity) { version3.parent = parent }

func (version3 *Vrrp_Interfaces_Interface_Ipv6_Version3) GetParent() types.Entity { return version3.parent }

func (version3 *Vrrp_Interfaces_Interface_Ipv6_Version3) GetParentYangName() string { return "ipv6" }

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters
// The VRRP virtual router configuration table
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The VRRP virtual router being configured. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter.
    VirtualRouter []Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters) GetFilter() yfilter.YFilter { return virtualRouters.YFilter }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters) SetFilter(yf yfilter.YFilter) { virtualRouters.YFilter = yf }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters) GetGoName(yname string) string {
    if yname == "virtual-router" { return "VirtualRouter" }
    return ""
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters) GetSegmentPath() string {
    return "virtual-routers"
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "virtual-router" {
        for _, c := range virtualRouters.VirtualRouter {
            if virtualRouters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter{}
        virtualRouters.VirtualRouter = append(virtualRouters.VirtualRouter, child)
        return &virtualRouters.VirtualRouter[len(virtualRouters.VirtualRouter)-1]
    }
    return nil
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range virtualRouters.VirtualRouter {
        children[virtualRouters.VirtualRouter[i].GetSegmentPath()] = &virtualRouters.VirtualRouter[i]
    }
    return children
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters) GetBundleName() string { return "cisco_ios_xr" }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters) GetYangName() string { return "virtual-routers" }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters) SetParent(parent types.Entity) { virtualRouters.parent = parent }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters) GetParent() types.Entity { return virtualRouters.parent }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters) GetParentYangName() string { return "version3" }

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter
// The VRRP virtual router being configured
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter) GetFilter() yfilter.YFilter { return virtualRouter.YFilter }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter) SetFilter(yf yfilter.YFilter) { virtualRouter.YFilter = yf }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter) GetGoName(yname string) string {
    if yname == "vr-id" { return "VrId" }
    if yname == "bfd" { return "Bfd" }
    if yname == "priority" { return "Priority" }
    if yname == "accept-mode-disable" { return "AcceptModeDisable" }
    if yname == "preempt" { return "Preempt" }
    if yname == "session-name" { return "SessionName" }
    if yname == "global-ipv6-addresses" { return "GlobalIpv6Addresses" }
    if yname == "tracks" { return "Tracks" }
    if yname == "timer" { return "Timer" }
    if yname == "tracked-objects" { return "TrackedObjects" }
    if yname == "link-local-ipv6-address" { return "LinkLocalIpv6Address" }
    return ""
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter) GetSegmentPath() string {
    return "virtual-router" + "[vr-id='" + fmt.Sprintf("%v", virtualRouter.VrId) + "']"
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global-ipv6-addresses" {
        return &virtualRouter.GlobalIpv6Addresses
    }
    if childYangName == "tracks" {
        return &virtualRouter.Tracks
    }
    if childYangName == "timer" {
        return &virtualRouter.Timer
    }
    if childYangName == "tracked-objects" {
        return &virtualRouter.TrackedObjects
    }
    if childYangName == "link-local-ipv6-address" {
        return &virtualRouter.LinkLocalIpv6Address
    }
    return nil
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["global-ipv6-addresses"] = &virtualRouter.GlobalIpv6Addresses
    children["tracks"] = &virtualRouter.Tracks
    children["timer"] = &virtualRouter.Timer
    children["tracked-objects"] = &virtualRouter.TrackedObjects
    children["link-local-ipv6-address"] = &virtualRouter.LinkLocalIpv6Address
    return children
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vr-id"] = virtualRouter.VrId
    leafs["bfd"] = virtualRouter.Bfd
    leafs["priority"] = virtualRouter.Priority
    leafs["accept-mode-disable"] = virtualRouter.AcceptModeDisable
    leafs["preempt"] = virtualRouter.Preempt
    leafs["session-name"] = virtualRouter.SessionName
    return leafs
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter) GetBundleName() string { return "cisco_ios_xr" }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter) GetYangName() string { return "virtual-router" }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter) SetParent(parent types.Entity) { virtualRouter.parent = parent }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter) GetParent() types.Entity { return virtualRouter.parent }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter) GetParentYangName() string { return "virtual-routers" }

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses
// The table of VRRP virtual global IPv6
// addresses
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A VRRP virtual global IPv6 IP address. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address.
    GlobalIpv6Address []Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address
}

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses) GetFilter() yfilter.YFilter { return globalIpv6Addresses.YFilter }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses) SetFilter(yf yfilter.YFilter) { globalIpv6Addresses.YFilter = yf }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses) GetGoName(yname string) string {
    if yname == "global-ipv6-address" { return "GlobalIpv6Address" }
    return ""
}

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses) GetSegmentPath() string {
    return "global-ipv6-addresses"
}

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global-ipv6-address" {
        for _, c := range globalIpv6Addresses.GlobalIpv6Address {
            if globalIpv6Addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address{}
        globalIpv6Addresses.GlobalIpv6Address = append(globalIpv6Addresses.GlobalIpv6Address, child)
        return &globalIpv6Addresses.GlobalIpv6Address[len(globalIpv6Addresses.GlobalIpv6Address)-1]
    }
    return nil
}

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range globalIpv6Addresses.GlobalIpv6Address {
        children[globalIpv6Addresses.GlobalIpv6Address[i].GetSegmentPath()] = &globalIpv6Addresses.GlobalIpv6Address[i]
    }
    return children
}

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses) GetYangName() string { return "global-ipv6-addresses" }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses) SetParent(parent types.Entity) { globalIpv6Addresses.parent = parent }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses) GetParent() types.Entity { return globalIpv6Addresses.parent }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses) GetParentYangName() string { return "virtual-router" }

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address
// A VRRP virtual global IPv6 IP address
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRRP virtual global IPv6 address. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}
}

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetFilter() yfilter.YFilter { return globalIpv6Address.YFilter }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) SetFilter(yf yfilter.YFilter) { globalIpv6Address.YFilter = yf }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    return ""
}

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetSegmentPath() string {
    return "global-ipv6-address" + "[ip-address='" + fmt.Sprintf("%v", globalIpv6Address.IpAddress) + "']"
}

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = globalIpv6Address.IpAddress
    return leafs
}

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetYangName() string { return "global-ipv6-address" }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) SetParent(parent types.Entity) { globalIpv6Address.parent = parent }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetParent() types.Entity { return globalIpv6Address.parent }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetParentYangName() string { return "global-ipv6-addresses" }

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks
// Track an item, reducing priority if it
// goes down
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object to be tracked. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track.
    Track []Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track
}

func (tracks *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks) GetFilter() yfilter.YFilter { return tracks.YFilter }

func (tracks *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks) SetFilter(yf yfilter.YFilter) { tracks.YFilter = yf }

func (tracks *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks) GetGoName(yname string) string {
    if yname == "track" { return "Track" }
    return ""
}

func (tracks *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks) GetSegmentPath() string {
    return "tracks"
}

func (tracks *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track" {
        for _, c := range tracks.Track {
            if tracks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track{}
        tracks.Track = append(tracks.Track, child)
        return &tracks.Track[len(tracks.Track)-1]
    }
    return nil
}

func (tracks *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tracks.Track {
        children[tracks.Track[i].GetSegmentPath()] = &tracks.Track[i]
    }
    return children
}

func (tracks *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tracks *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks) GetBundleName() string { return "cisco_ios_xr" }

func (tracks *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks) GetYangName() string { return "tracks" }

func (tracks *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tracks *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tracks *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tracks *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks) SetParent(parent types.Entity) { tracks.parent = parent }

func (tracks *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks) GetParent() types.Entity { return tracks.parent }

func (tracks *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks) GetParentYangName() string { return "virtual-router" }

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track
// Object to be tracked
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object to be tracked, interface name for
    // interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Priority decrement. The type is interface{} with range: 1..254. This
    // attribute is mandatory.
    Priority interface{}
}

func (track *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetFilter() yfilter.YFilter { return track.YFilter }

func (track *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track) SetFilter(yf yfilter.YFilter) { track.YFilter = yf }

func (track *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (track *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetSegmentPath() string {
    return "track" + "[interface-name='" + fmt.Sprintf("%v", track.InterfaceName) + "']"
}

func (track *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (track *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (track *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = track.InterfaceName
    leafs["priority"] = track.Priority
    return leafs
}

func (track *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetBundleName() string { return "cisco_ios_xr" }

func (track *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetYangName() string { return "track" }

func (track *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (track *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (track *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (track *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track) SetParent(parent types.Entity) { track.parent = parent }

func (track *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetParent() types.Entity { return track.parent }

func (track *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetParentYangName() string { return "tracks" }

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer
// Set advertisement timer
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TRUE - Advertise time configured in milliseconds, FALSE - Advertise time
    // configured in seconds. The type is bool. The default value is false.
    InMsec interface{}

    // Advertisement time in milliseconds. The type is interface{} with range:
    // 100..40950. Units are millisecond.
    AdvertisementTimeInMsec interface{}

    // Advertisement time in seconds. The type is interface{} with range: 1..40.
    // Units are second.
    AdvertisementTimeInSec interface{}

    // TRUE - Force configured timer values to be used, required when configured
    // in milliseconds. The type is bool. The default value is false.
    Forced interface{}
}

func (timer *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer) GetFilter() yfilter.YFilter { return timer.YFilter }

func (timer *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer) SetFilter(yf yfilter.YFilter) { timer.YFilter = yf }

func (timer *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer) GetGoName(yname string) string {
    if yname == "in-msec" { return "InMsec" }
    if yname == "advertisement-time-in-msec" { return "AdvertisementTimeInMsec" }
    if yname == "advertisement-time-in-sec" { return "AdvertisementTimeInSec" }
    if yname == "forced" { return "Forced" }
    return ""
}

func (timer *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer) GetSegmentPath() string {
    return "timer"
}

func (timer *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timer *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timer *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-msec"] = timer.InMsec
    leafs["advertisement-time-in-msec"] = timer.AdvertisementTimeInMsec
    leafs["advertisement-time-in-sec"] = timer.AdvertisementTimeInSec
    leafs["forced"] = timer.Forced
    return leafs
}

func (timer *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer) GetBundleName() string { return "cisco_ios_xr" }

func (timer *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer) GetYangName() string { return "timer" }

func (timer *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timer *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timer *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timer *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer) SetParent(parent types.Entity) { timer.parent = parent }

func (timer *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer) GetParent() types.Entity { return timer.parent }

func (timer *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_Timer) GetParentYangName() string { return "virtual-router" }

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects
// Track an object, reducing priority if it
// goes down
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object to be tracked. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject.
    TrackedObject []Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetFilter() yfilter.YFilter { return trackedObjects.YFilter }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects) SetFilter(yf yfilter.YFilter) { trackedObjects.YFilter = yf }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetGoName(yname string) string {
    if yname == "tracked-object" { return "TrackedObject" }
    return ""
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetSegmentPath() string {
    return "tracked-objects"
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tracked-object" {
        for _, c := range trackedObjects.TrackedObject {
            if trackedObjects.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject{}
        trackedObjects.TrackedObject = append(trackedObjects.TrackedObject, child)
        return &trackedObjects.TrackedObject[len(trackedObjects.TrackedObject)-1]
    }
    return nil
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackedObjects.TrackedObject {
        children[trackedObjects.TrackedObject[i].GetSegmentPath()] = &trackedObjects.TrackedObject[i]
    }
    return children
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetBundleName() string { return "cisco_ios_xr" }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetYangName() string { return "tracked-objects" }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects) SetParent(parent types.Entity) { trackedObjects.parent = parent }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetParent() types.Entity { return trackedObjects.parent }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetParentYangName() string { return "virtual-router" }

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject
// Object to be tracked
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object to be tracked, interface name for
    // interfaces. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ObjectName interface{}

    // Priority decrement. The type is interface{} with range: 1..254. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetFilter() yfilter.YFilter { return trackedObject.YFilter }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) SetFilter(yf yfilter.YFilter) { trackedObject.YFilter = yf }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetSegmentPath() string {
    return "tracked-object" + "[object-name='" + fmt.Sprintf("%v", trackedObject.ObjectName) + "']"
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = trackedObject.ObjectName
    leafs["priority-decrement"] = trackedObject.PriorityDecrement
    return leafs
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetBundleName() string { return "cisco_ios_xr" }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetYangName() string { return "tracked-object" }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) SetParent(parent types.Entity) { trackedObject.parent = parent }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetParent() types.Entity { return trackedObject.parent }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetParentYangName() string { return "tracked-objects" }

// Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address
// The VRRP IPv6 virtual linklocal address
type Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address struct {
    parent types.Entity
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

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address) GetFilter() yfilter.YFilter { return linkLocalIpv6Address.YFilter }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address) SetFilter(yf yfilter.YFilter) { linkLocalIpv6Address.YFilter = yf }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "auto-configure" { return "AutoConfigure" }
    return ""
}

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address) GetSegmentPath() string {
    return "link-local-ipv6-address"
}

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = linkLocalIpv6Address.IpAddress
    leafs["auto-configure"] = linkLocalIpv6Address.AutoConfigure
    return leafs
}

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address) GetYangName() string { return "link-local-ipv6-address" }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address) SetParent(parent types.Entity) { linkLocalIpv6Address.parent = parent }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address) GetParent() types.Entity { return linkLocalIpv6Address.parent }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_Version3_VirtualRouters_VirtualRouter_LinkLocalIpv6Address) GetParentYangName() string { return "virtual-router" }

// Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters
// The VRRP slave group configuration table
type Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The VRRP slave being configured. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter.
    SlaveVirtualRouter []Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter
}

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters) GetFilter() yfilter.YFilter { return slaveVirtualRouters.YFilter }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters) SetFilter(yf yfilter.YFilter) { slaveVirtualRouters.YFilter = yf }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters) GetGoName(yname string) string {
    if yname == "slave-virtual-router" { return "SlaveVirtualRouter" }
    return ""
}

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters) GetSegmentPath() string {
    return "slave-virtual-routers"
}

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slave-virtual-router" {
        for _, c := range slaveVirtualRouters.SlaveVirtualRouter {
            if slaveVirtualRouters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter{}
        slaveVirtualRouters.SlaveVirtualRouter = append(slaveVirtualRouters.SlaveVirtualRouter, child)
        return &slaveVirtualRouters.SlaveVirtualRouter[len(slaveVirtualRouters.SlaveVirtualRouter)-1]
    }
    return nil
}

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slaveVirtualRouters.SlaveVirtualRouter {
        children[slaveVirtualRouters.SlaveVirtualRouter[i].GetSegmentPath()] = &slaveVirtualRouters.SlaveVirtualRouter[i]
    }
    return children
}

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters) GetBundleName() string { return "cisco_ios_xr" }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters) GetYangName() string { return "slave-virtual-routers" }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters) SetParent(parent types.Entity) { slaveVirtualRouters.parent = parent }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters) GetParent() types.Entity { return slaveVirtualRouters.parent }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters) GetParentYangName() string { return "ipv6" }

// Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter
// The VRRP slave being configured
type Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter) GetFilter() yfilter.YFilter { return slaveVirtualRouter.YFilter }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter) SetFilter(yf yfilter.YFilter) { slaveVirtualRouter.YFilter = yf }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter) GetGoName(yname string) string {
    if yname == "slave-virtual-router-id" { return "SlaveVirtualRouterId" }
    if yname == "follow" { return "Follow" }
    if yname == "accept-mode-disable" { return "AcceptModeDisable" }
    if yname == "link-local-ipv6-address" { return "LinkLocalIpv6Address" }
    if yname == "global-ipv6-addresses" { return "GlobalIpv6Addresses" }
    return ""
}

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter) GetSegmentPath() string {
    return "slave-virtual-router" + "[slave-virtual-router-id='" + fmt.Sprintf("%v", slaveVirtualRouter.SlaveVirtualRouterId) + "']"
}

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-local-ipv6-address" {
        return &slaveVirtualRouter.LinkLocalIpv6Address
    }
    if childYangName == "global-ipv6-addresses" {
        return &slaveVirtualRouter.GlobalIpv6Addresses
    }
    return nil
}

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["link-local-ipv6-address"] = &slaveVirtualRouter.LinkLocalIpv6Address
    children["global-ipv6-addresses"] = &slaveVirtualRouter.GlobalIpv6Addresses
    return children
}

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slave-virtual-router-id"] = slaveVirtualRouter.SlaveVirtualRouterId
    leafs["follow"] = slaveVirtualRouter.Follow
    leafs["accept-mode-disable"] = slaveVirtualRouter.AcceptModeDisable
    return leafs
}

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter) GetBundleName() string { return "cisco_ios_xr" }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter) GetYangName() string { return "slave-virtual-router" }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter) SetParent(parent types.Entity) { slaveVirtualRouter.parent = parent }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter) GetParent() types.Entity { return slaveVirtualRouter.parent }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter) GetParentYangName() string { return "slave-virtual-routers" }

// Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address
// The VRRP IPv6 virtual linklocal address
type Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address struct {
    parent types.Entity
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

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address) GetFilter() yfilter.YFilter { return linkLocalIpv6Address.YFilter }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address) SetFilter(yf yfilter.YFilter) { linkLocalIpv6Address.YFilter = yf }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    if yname == "auto-configure" { return "AutoConfigure" }
    return ""
}

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address) GetSegmentPath() string {
    return "link-local-ipv6-address"
}

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = linkLocalIpv6Address.IpAddress
    leafs["auto-configure"] = linkLocalIpv6Address.AutoConfigure
    return leafs
}

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address) GetYangName() string { return "link-local-ipv6-address" }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address) SetParent(parent types.Entity) { linkLocalIpv6Address.parent = parent }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address) GetParent() types.Entity { return linkLocalIpv6Address.parent }

func (linkLocalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_LinkLocalIpv6Address) GetParentYangName() string { return "slave-virtual-router" }

// Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses
// The table of VRRP virtual global IPv6
// addresses
type Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A VRRP virtual global IPv6 IP address. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address.
    GlobalIpv6Address []Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address
}

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses) GetFilter() yfilter.YFilter { return globalIpv6Addresses.YFilter }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses) SetFilter(yf yfilter.YFilter) { globalIpv6Addresses.YFilter = yf }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses) GetGoName(yname string) string {
    if yname == "global-ipv6-address" { return "GlobalIpv6Address" }
    return ""
}

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses) GetSegmentPath() string {
    return "global-ipv6-addresses"
}

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global-ipv6-address" {
        for _, c := range globalIpv6Addresses.GlobalIpv6Address {
            if globalIpv6Addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address{}
        globalIpv6Addresses.GlobalIpv6Address = append(globalIpv6Addresses.GlobalIpv6Address, child)
        return &globalIpv6Addresses.GlobalIpv6Address[len(globalIpv6Addresses.GlobalIpv6Address)-1]
    }
    return nil
}

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range globalIpv6Addresses.GlobalIpv6Address {
        children[globalIpv6Addresses.GlobalIpv6Address[i].GetSegmentPath()] = &globalIpv6Addresses.GlobalIpv6Address[i]
    }
    return children
}

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses) GetYangName() string { return "global-ipv6-addresses" }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses) SetParent(parent types.Entity) { globalIpv6Addresses.parent = parent }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses) GetParent() types.Entity { return globalIpv6Addresses.parent }

func (globalIpv6Addresses *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses) GetParentYangName() string { return "slave-virtual-router" }

// Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address
// A VRRP virtual global IPv6 IP address
type Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRRP virtual global IPv6 address. The type is one
    // of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    IpAddress interface{}
}

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetFilter() yfilter.YFilter { return globalIpv6Address.YFilter }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) SetFilter(yf yfilter.YFilter) { globalIpv6Address.YFilter = yf }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    return ""
}

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetSegmentPath() string {
    return "global-ipv6-address" + "[ip-address='" + fmt.Sprintf("%v", globalIpv6Address.IpAddress) + "']"
}

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = globalIpv6Address.IpAddress
    return leafs
}

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetYangName() string { return "global-ipv6-address" }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) SetParent(parent types.Entity) { globalIpv6Address.parent = parent }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetParent() types.Entity { return globalIpv6Address.parent }

func (globalIpv6Address *Vrrp_Interfaces_Interface_Ipv6_SlaveVirtualRouters_SlaveVirtualRouter_GlobalIpv6Addresses_GlobalIpv6Address) GetParentYangName() string { return "global-ipv6-addresses" }

// Vrrp_Interfaces_Interface_Delay
// Minimum and Reload Delay
// This type is a presence type.
type Vrrp_Interfaces_Interface_Delay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minimum delay in seconds. The type is interface{} with range: 0..10000.
    // This attribute is mandatory. Units are second.
    MinDelay interface{}

    // Reload delay in seconds. The type is interface{} with range: 0..10000. This
    // attribute is mandatory. Units are second.
    ReloadDelay interface{}
}

func (delay *Vrrp_Interfaces_Interface_Delay) GetFilter() yfilter.YFilter { return delay.YFilter }

func (delay *Vrrp_Interfaces_Interface_Delay) SetFilter(yf yfilter.YFilter) { delay.YFilter = yf }

func (delay *Vrrp_Interfaces_Interface_Delay) GetGoName(yname string) string {
    if yname == "min-delay" { return "MinDelay" }
    if yname == "reload-delay" { return "ReloadDelay" }
    return ""
}

func (delay *Vrrp_Interfaces_Interface_Delay) GetSegmentPath() string {
    return "delay"
}

func (delay *Vrrp_Interfaces_Interface_Delay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delay *Vrrp_Interfaces_Interface_Delay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delay *Vrrp_Interfaces_Interface_Delay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["min-delay"] = delay.MinDelay
    leafs["reload-delay"] = delay.ReloadDelay
    return leafs
}

func (delay *Vrrp_Interfaces_Interface_Delay) GetBundleName() string { return "cisco_ios_xr" }

func (delay *Vrrp_Interfaces_Interface_Delay) GetYangName() string { return "delay" }

func (delay *Vrrp_Interfaces_Interface_Delay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delay *Vrrp_Interfaces_Interface_Delay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delay *Vrrp_Interfaces_Interface_Delay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delay *Vrrp_Interfaces_Interface_Delay) SetParent(parent types.Entity) { delay.parent = parent }

func (delay *Vrrp_Interfaces_Interface_Delay) GetParent() types.Entity { return delay.parent }

func (delay *Vrrp_Interfaces_Interface_Delay) GetParentYangName() string { return "interface" }

// Vrrp_Interfaces_Interface_Ipv4
// IPv4 VRRP configuration
type Vrrp_Interfaces_Interface_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version 3 VRRP configuration.
    Version3 Vrrp_Interfaces_Interface_Ipv4_Version3

    // The VRRP slave group configuration table.
    SlaveVirtualRouters Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters

    // Version 2 VRRP configuration.
    Version2 Vrrp_Interfaces_Interface_Ipv4_Version2
}

func (ipv4 *Vrrp_Interfaces_Interface_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Vrrp_Interfaces_Interface_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Vrrp_Interfaces_Interface_Ipv4) GetGoName(yname string) string {
    if yname == "version3" { return "Version3" }
    if yname == "slave-virtual-routers" { return "SlaveVirtualRouters" }
    if yname == "version2" { return "Version2" }
    return ""
}

func (ipv4 *Vrrp_Interfaces_Interface_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Vrrp_Interfaces_Interface_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "version3" {
        return &ipv4.Version3
    }
    if childYangName == "slave-virtual-routers" {
        return &ipv4.SlaveVirtualRouters
    }
    if childYangName == "version2" {
        return &ipv4.Version2
    }
    return nil
}

func (ipv4 *Vrrp_Interfaces_Interface_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["version3"] = &ipv4.Version3
    children["slave-virtual-routers"] = &ipv4.SlaveVirtualRouters
    children["version2"] = &ipv4.Version2
    return children
}

func (ipv4 *Vrrp_Interfaces_Interface_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *Vrrp_Interfaces_Interface_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Vrrp_Interfaces_Interface_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Vrrp_Interfaces_Interface_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Vrrp_Interfaces_Interface_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Vrrp_Interfaces_Interface_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Vrrp_Interfaces_Interface_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Vrrp_Interfaces_Interface_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Vrrp_Interfaces_Interface_Ipv4) GetParentYangName() string { return "interface" }

// Vrrp_Interfaces_Interface_Ipv4_Version3
// Version 3 VRRP configuration
type Vrrp_Interfaces_Interface_Ipv4_Version3 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The VRRP virtual router configuration table.
    VirtualRouters Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters
}

func (version3 *Vrrp_Interfaces_Interface_Ipv4_Version3) GetFilter() yfilter.YFilter { return version3.YFilter }

func (version3 *Vrrp_Interfaces_Interface_Ipv4_Version3) SetFilter(yf yfilter.YFilter) { version3.YFilter = yf }

func (version3 *Vrrp_Interfaces_Interface_Ipv4_Version3) GetGoName(yname string) string {
    if yname == "virtual-routers" { return "VirtualRouters" }
    return ""
}

func (version3 *Vrrp_Interfaces_Interface_Ipv4_Version3) GetSegmentPath() string {
    return "version3"
}

func (version3 *Vrrp_Interfaces_Interface_Ipv4_Version3) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "virtual-routers" {
        return &version3.VirtualRouters
    }
    return nil
}

func (version3 *Vrrp_Interfaces_Interface_Ipv4_Version3) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["virtual-routers"] = &version3.VirtualRouters
    return children
}

func (version3 *Vrrp_Interfaces_Interface_Ipv4_Version3) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (version3 *Vrrp_Interfaces_Interface_Ipv4_Version3) GetBundleName() string { return "cisco_ios_xr" }

func (version3 *Vrrp_Interfaces_Interface_Ipv4_Version3) GetYangName() string { return "version3" }

func (version3 *Vrrp_Interfaces_Interface_Ipv4_Version3) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (version3 *Vrrp_Interfaces_Interface_Ipv4_Version3) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (version3 *Vrrp_Interfaces_Interface_Ipv4_Version3) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (version3 *Vrrp_Interfaces_Interface_Ipv4_Version3) SetParent(parent types.Entity) { version3.parent = parent }

func (version3 *Vrrp_Interfaces_Interface_Ipv4_Version3) GetParent() types.Entity { return version3.parent }

func (version3 *Vrrp_Interfaces_Interface_Ipv4_Version3) GetParentYangName() string { return "ipv4" }

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters
// The VRRP virtual router configuration table
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The VRRP virtual router being configured. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter.
    VirtualRouter []Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters) GetFilter() yfilter.YFilter { return virtualRouters.YFilter }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters) SetFilter(yf yfilter.YFilter) { virtualRouters.YFilter = yf }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters) GetGoName(yname string) string {
    if yname == "virtual-router" { return "VirtualRouter" }
    return ""
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters) GetSegmentPath() string {
    return "virtual-routers"
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "virtual-router" {
        for _, c := range virtualRouters.VirtualRouter {
            if virtualRouters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter{}
        virtualRouters.VirtualRouter = append(virtualRouters.VirtualRouter, child)
        return &virtualRouters.VirtualRouter[len(virtualRouters.VirtualRouter)-1]
    }
    return nil
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range virtualRouters.VirtualRouter {
        children[virtualRouters.VirtualRouter[i].GetSegmentPath()] = &virtualRouters.VirtualRouter[i]
    }
    return children
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters) GetBundleName() string { return "cisco_ios_xr" }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters) GetYangName() string { return "virtual-routers" }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters) SetParent(parent types.Entity) { virtualRouters.parent = parent }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters) GetParent() types.Entity { return virtualRouters.parent }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters) GetParentYangName() string { return "version3" }

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter
// The VRRP virtual router being configured
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter) GetFilter() yfilter.YFilter { return virtualRouter.YFilter }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter) SetFilter(yf yfilter.YFilter) { virtualRouter.YFilter = yf }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter) GetGoName(yname string) string {
    if yname == "vr-id" { return "VrId" }
    if yname == "session-name" { return "SessionName" }
    if yname == "bfd" { return "Bfd" }
    if yname == "primary-ipv4-address" { return "PrimaryIpv4Address" }
    if yname == "preempt" { return "Preempt" }
    if yname == "accept-mode-disable" { return "AcceptModeDisable" }
    if yname == "priority" { return "Priority" }
    if yname == "timer" { return "Timer" }
    if yname == "secondary-ipv4-addresses" { return "SecondaryIpv4Addresses" }
    if yname == "tracked-objects" { return "TrackedObjects" }
    if yname == "tracks" { return "Tracks" }
    return ""
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter) GetSegmentPath() string {
    return "virtual-router" + "[vr-id='" + fmt.Sprintf("%v", virtualRouter.VrId) + "']"
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "timer" {
        return &virtualRouter.Timer
    }
    if childYangName == "secondary-ipv4-addresses" {
        return &virtualRouter.SecondaryIpv4Addresses
    }
    if childYangName == "tracked-objects" {
        return &virtualRouter.TrackedObjects
    }
    if childYangName == "tracks" {
        return &virtualRouter.Tracks
    }
    return nil
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["timer"] = &virtualRouter.Timer
    children["secondary-ipv4-addresses"] = &virtualRouter.SecondaryIpv4Addresses
    children["tracked-objects"] = &virtualRouter.TrackedObjects
    children["tracks"] = &virtualRouter.Tracks
    return children
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vr-id"] = virtualRouter.VrId
    leafs["session-name"] = virtualRouter.SessionName
    leafs["bfd"] = virtualRouter.Bfd
    leafs["primary-ipv4-address"] = virtualRouter.PrimaryIpv4Address
    leafs["preempt"] = virtualRouter.Preempt
    leafs["accept-mode-disable"] = virtualRouter.AcceptModeDisable
    leafs["priority"] = virtualRouter.Priority
    return leafs
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter) GetBundleName() string { return "cisco_ios_xr" }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter) GetYangName() string { return "virtual-router" }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter) SetParent(parent types.Entity) { virtualRouter.parent = parent }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter) GetParent() types.Entity { return virtualRouter.parent }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter) GetParentYangName() string { return "virtual-routers" }

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer
// Set advertisement timer
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TRUE - Advertise time configured in milliseconds, FALSE - Advertise time
    // configured in seconds. The type is bool. The default value is false.
    InMsec interface{}

    // Advertisement time in milliseconds. The type is interface{} with range:
    // 100..40950. Units are millisecond.
    AdvertisementTimeInMsec interface{}

    // Advertisement time in seconds. The type is interface{} with range: 1..40.
    // Units are second.
    AdvertisementTimeInSec interface{}

    // TRUE - Force configured timer values to be used, required when configured
    // in milliseconds. The type is bool. The default value is false.
    Forced interface{}
}

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer) GetFilter() yfilter.YFilter { return timer.YFilter }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer) SetFilter(yf yfilter.YFilter) { timer.YFilter = yf }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer) GetGoName(yname string) string {
    if yname == "in-msec" { return "InMsec" }
    if yname == "advertisement-time-in-msec" { return "AdvertisementTimeInMsec" }
    if yname == "advertisement-time-in-sec" { return "AdvertisementTimeInSec" }
    if yname == "forced" { return "Forced" }
    return ""
}

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer) GetSegmentPath() string {
    return "timer"
}

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-msec"] = timer.InMsec
    leafs["advertisement-time-in-msec"] = timer.AdvertisementTimeInMsec
    leafs["advertisement-time-in-sec"] = timer.AdvertisementTimeInSec
    leafs["forced"] = timer.Forced
    return leafs
}

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer) GetBundleName() string { return "cisco_ios_xr" }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer) GetYangName() string { return "timer" }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer) SetParent(parent types.Entity) { timer.parent = parent }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer) GetParent() types.Entity { return timer.parent }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Timer) GetParentYangName() string { return "virtual-router" }

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses
// The table of VRRP secondary IPv4 addresses
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A VRRP secondary IPv4 address. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address.
    SecondaryIpv4Address []Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetFilter() yfilter.YFilter { return secondaryIpv4Addresses.YFilter }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) SetFilter(yf yfilter.YFilter) { secondaryIpv4Addresses.YFilter = yf }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetGoName(yname string) string {
    if yname == "secondary-ipv4-address" { return "SecondaryIpv4Address" }
    return ""
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetSegmentPath() string {
    return "secondary-ipv4-addresses"
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "secondary-ipv4-address" {
        for _, c := range secondaryIpv4Addresses.SecondaryIpv4Address {
            if secondaryIpv4Addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address{}
        secondaryIpv4Addresses.SecondaryIpv4Address = append(secondaryIpv4Addresses.SecondaryIpv4Address, child)
        return &secondaryIpv4Addresses.SecondaryIpv4Address[len(secondaryIpv4Addresses.SecondaryIpv4Address)-1]
    }
    return nil
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range secondaryIpv4Addresses.SecondaryIpv4Address {
        children[secondaryIpv4Addresses.SecondaryIpv4Address[i].GetSegmentPath()] = &secondaryIpv4Addresses.SecondaryIpv4Address[i]
    }
    return children
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetYangName() string { return "secondary-ipv4-addresses" }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) SetParent(parent types.Entity) { secondaryIpv4Addresses.parent = parent }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetParent() types.Entity { return secondaryIpv4Addresses.parent }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetParentYangName() string { return "virtual-router" }

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address
// A VRRP secondary IPv4 address
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRRP Secondary IPv4 address. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetFilter() yfilter.YFilter { return secondaryIpv4Address.YFilter }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) SetFilter(yf yfilter.YFilter) { secondaryIpv4Address.YFilter = yf }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    return ""
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetSegmentPath() string {
    return "secondary-ipv4-address" + "[ip-address='" + fmt.Sprintf("%v", secondaryIpv4Address.IpAddress) + "']"
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = secondaryIpv4Address.IpAddress
    return leafs
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetYangName() string { return "secondary-ipv4-address" }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) SetParent(parent types.Entity) { secondaryIpv4Address.parent = parent }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetParent() types.Entity { return secondaryIpv4Address.parent }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetParentYangName() string { return "secondary-ipv4-addresses" }

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects
// Track an object, reducing priority if it
// goes down
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object to be tracked. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject.
    TrackedObject []Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetFilter() yfilter.YFilter { return trackedObjects.YFilter }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects) SetFilter(yf yfilter.YFilter) { trackedObjects.YFilter = yf }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetGoName(yname string) string {
    if yname == "tracked-object" { return "TrackedObject" }
    return ""
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetSegmentPath() string {
    return "tracked-objects"
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tracked-object" {
        for _, c := range trackedObjects.TrackedObject {
            if trackedObjects.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject{}
        trackedObjects.TrackedObject = append(trackedObjects.TrackedObject, child)
        return &trackedObjects.TrackedObject[len(trackedObjects.TrackedObject)-1]
    }
    return nil
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackedObjects.TrackedObject {
        children[trackedObjects.TrackedObject[i].GetSegmentPath()] = &trackedObjects.TrackedObject[i]
    }
    return children
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetBundleName() string { return "cisco_ios_xr" }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetYangName() string { return "tracked-objects" }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects) SetParent(parent types.Entity) { trackedObjects.parent = parent }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetParent() types.Entity { return trackedObjects.parent }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects) GetParentYangName() string { return "virtual-router" }

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject
// Object to be tracked
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object to be tracked, interface name for
    // interfaces. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ObjectName interface{}

    // Priority decrement. The type is interface{} with range: 1..254. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetFilter() yfilter.YFilter { return trackedObject.YFilter }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) SetFilter(yf yfilter.YFilter) { trackedObject.YFilter = yf }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetSegmentPath() string {
    return "tracked-object" + "[object-name='" + fmt.Sprintf("%v", trackedObject.ObjectName) + "']"
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = trackedObject.ObjectName
    leafs["priority-decrement"] = trackedObject.PriorityDecrement
    return leafs
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetBundleName() string { return "cisco_ios_xr" }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetYangName() string { return "tracked-object" }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) SetParent(parent types.Entity) { trackedObject.parent = parent }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetParent() types.Entity { return trackedObject.parent }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetParentYangName() string { return "tracked-objects" }

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks
// Track an item, reducing priority if it
// goes down
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object to be tracked. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track.
    Track []Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track
}

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks) GetFilter() yfilter.YFilter { return tracks.YFilter }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks) SetFilter(yf yfilter.YFilter) { tracks.YFilter = yf }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks) GetGoName(yname string) string {
    if yname == "track" { return "Track" }
    return ""
}

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks) GetSegmentPath() string {
    return "tracks"
}

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track" {
        for _, c := range tracks.Track {
            if tracks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track{}
        tracks.Track = append(tracks.Track, child)
        return &tracks.Track[len(tracks.Track)-1]
    }
    return nil
}

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tracks.Track {
        children[tracks.Track[i].GetSegmentPath()] = &tracks.Track[i]
    }
    return children
}

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks) GetBundleName() string { return "cisco_ios_xr" }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks) GetYangName() string { return "tracks" }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks) SetParent(parent types.Entity) { tracks.parent = parent }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks) GetParent() types.Entity { return tracks.parent }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks) GetParentYangName() string { return "virtual-router" }

// Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track
// Object to be tracked
type Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object to be tracked, interface name for
    // interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Priority decrement. The type is interface{} with range: 1..254. This
    // attribute is mandatory.
    Priority interface{}
}

func (track *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetFilter() yfilter.YFilter { return track.YFilter }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track) SetFilter(yf yfilter.YFilter) { track.YFilter = yf }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (track *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetSegmentPath() string {
    return "track" + "[interface-name='" + fmt.Sprintf("%v", track.InterfaceName) + "']"
}

func (track *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (track *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (track *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = track.InterfaceName
    leafs["priority"] = track.Priority
    return leafs
}

func (track *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetBundleName() string { return "cisco_ios_xr" }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetYangName() string { return "track" }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track) SetParent(parent types.Entity) { track.parent = parent }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetParent() types.Entity { return track.parent }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version3_VirtualRouters_VirtualRouter_Tracks_Track) GetParentYangName() string { return "tracks" }

// Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters
// The VRRP slave group configuration table
type Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The VRRP slave being configured. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter.
    SlaveVirtualRouter []Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter
}

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters) GetFilter() yfilter.YFilter { return slaveVirtualRouters.YFilter }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters) SetFilter(yf yfilter.YFilter) { slaveVirtualRouters.YFilter = yf }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters) GetGoName(yname string) string {
    if yname == "slave-virtual-router" { return "SlaveVirtualRouter" }
    return ""
}

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters) GetSegmentPath() string {
    return "slave-virtual-routers"
}

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slave-virtual-router" {
        for _, c := range slaveVirtualRouters.SlaveVirtualRouter {
            if slaveVirtualRouters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter{}
        slaveVirtualRouters.SlaveVirtualRouter = append(slaveVirtualRouters.SlaveVirtualRouter, child)
        return &slaveVirtualRouters.SlaveVirtualRouter[len(slaveVirtualRouters.SlaveVirtualRouter)-1]
    }
    return nil
}

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slaveVirtualRouters.SlaveVirtualRouter {
        children[slaveVirtualRouters.SlaveVirtualRouter[i].GetSegmentPath()] = &slaveVirtualRouters.SlaveVirtualRouter[i]
    }
    return children
}

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters) GetBundleName() string { return "cisco_ios_xr" }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters) GetYangName() string { return "slave-virtual-routers" }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters) SetParent(parent types.Entity) { slaveVirtualRouters.parent = parent }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters) GetParent() types.Entity { return slaveVirtualRouters.parent }

func (slaveVirtualRouters *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters) GetParentYangName() string { return "ipv4" }

// Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter
// The VRRP slave being configured
type Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter) GetFilter() yfilter.YFilter { return slaveVirtualRouter.YFilter }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter) SetFilter(yf yfilter.YFilter) { slaveVirtualRouter.YFilter = yf }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter) GetGoName(yname string) string {
    if yname == "slave-virtual-router-id" { return "SlaveVirtualRouterId" }
    if yname == "follow" { return "Follow" }
    if yname == "accept-mode-disable" { return "AcceptModeDisable" }
    if yname == "primary-ipv4-address" { return "PrimaryIpv4Address" }
    if yname == "secondary-ipv4-addresses" { return "SecondaryIpv4Addresses" }
    return ""
}

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter) GetSegmentPath() string {
    return "slave-virtual-router" + "[slave-virtual-router-id='" + fmt.Sprintf("%v", slaveVirtualRouter.SlaveVirtualRouterId) + "']"
}

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "secondary-ipv4-addresses" {
        return &slaveVirtualRouter.SecondaryIpv4Addresses
    }
    return nil
}

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["secondary-ipv4-addresses"] = &slaveVirtualRouter.SecondaryIpv4Addresses
    return children
}

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slave-virtual-router-id"] = slaveVirtualRouter.SlaveVirtualRouterId
    leafs["follow"] = slaveVirtualRouter.Follow
    leafs["accept-mode-disable"] = slaveVirtualRouter.AcceptModeDisable
    leafs["primary-ipv4-address"] = slaveVirtualRouter.PrimaryIpv4Address
    return leafs
}

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter) GetBundleName() string { return "cisco_ios_xr" }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter) GetYangName() string { return "slave-virtual-router" }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter) SetParent(parent types.Entity) { slaveVirtualRouter.parent = parent }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter) GetParent() types.Entity { return slaveVirtualRouter.parent }

func (slaveVirtualRouter *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter) GetParentYangName() string { return "slave-virtual-routers" }

// Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses
// The table of VRRP secondary IPv4 addresses
type Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A VRRP secondary IPv4 address. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address.
    SecondaryIpv4Address []Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses) GetFilter() yfilter.YFilter { return secondaryIpv4Addresses.YFilter }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses) SetFilter(yf yfilter.YFilter) { secondaryIpv4Addresses.YFilter = yf }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses) GetGoName(yname string) string {
    if yname == "secondary-ipv4-address" { return "SecondaryIpv4Address" }
    return ""
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses) GetSegmentPath() string {
    return "secondary-ipv4-addresses"
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "secondary-ipv4-address" {
        for _, c := range secondaryIpv4Addresses.SecondaryIpv4Address {
            if secondaryIpv4Addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address{}
        secondaryIpv4Addresses.SecondaryIpv4Address = append(secondaryIpv4Addresses.SecondaryIpv4Address, child)
        return &secondaryIpv4Addresses.SecondaryIpv4Address[len(secondaryIpv4Addresses.SecondaryIpv4Address)-1]
    }
    return nil
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range secondaryIpv4Addresses.SecondaryIpv4Address {
        children[secondaryIpv4Addresses.SecondaryIpv4Address[i].GetSegmentPath()] = &secondaryIpv4Addresses.SecondaryIpv4Address[i]
    }
    return children
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses) GetYangName() string { return "secondary-ipv4-addresses" }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses) SetParent(parent types.Entity) { secondaryIpv4Addresses.parent = parent }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses) GetParent() types.Entity { return secondaryIpv4Addresses.parent }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses) GetParentYangName() string { return "slave-virtual-router" }

// Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address
// A VRRP secondary IPv4 address
type Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRRP Secondary IPv4 address. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetFilter() yfilter.YFilter { return secondaryIpv4Address.YFilter }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) SetFilter(yf yfilter.YFilter) { secondaryIpv4Address.YFilter = yf }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    return ""
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetSegmentPath() string {
    return "secondary-ipv4-address" + "[ip-address='" + fmt.Sprintf("%v", secondaryIpv4Address.IpAddress) + "']"
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = secondaryIpv4Address.IpAddress
    return leafs
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetYangName() string { return "secondary-ipv4-address" }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) SetParent(parent types.Entity) { secondaryIpv4Address.parent = parent }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetParent() types.Entity { return secondaryIpv4Address.parent }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_SlaveVirtualRouters_SlaveVirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetParentYangName() string { return "secondary-ipv4-addresses" }

// Vrrp_Interfaces_Interface_Ipv4_Version2
// Version 2 VRRP configuration
type Vrrp_Interfaces_Interface_Ipv4_Version2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The VRRP virtual router configuration table.
    VirtualRouters Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters
}

func (version2 *Vrrp_Interfaces_Interface_Ipv4_Version2) GetFilter() yfilter.YFilter { return version2.YFilter }

func (version2 *Vrrp_Interfaces_Interface_Ipv4_Version2) SetFilter(yf yfilter.YFilter) { version2.YFilter = yf }

func (version2 *Vrrp_Interfaces_Interface_Ipv4_Version2) GetGoName(yname string) string {
    if yname == "virtual-routers" { return "VirtualRouters" }
    return ""
}

func (version2 *Vrrp_Interfaces_Interface_Ipv4_Version2) GetSegmentPath() string {
    return "version2"
}

func (version2 *Vrrp_Interfaces_Interface_Ipv4_Version2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "virtual-routers" {
        return &version2.VirtualRouters
    }
    return nil
}

func (version2 *Vrrp_Interfaces_Interface_Ipv4_Version2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["virtual-routers"] = &version2.VirtualRouters
    return children
}

func (version2 *Vrrp_Interfaces_Interface_Ipv4_Version2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (version2 *Vrrp_Interfaces_Interface_Ipv4_Version2) GetBundleName() string { return "cisco_ios_xr" }

func (version2 *Vrrp_Interfaces_Interface_Ipv4_Version2) GetYangName() string { return "version2" }

func (version2 *Vrrp_Interfaces_Interface_Ipv4_Version2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (version2 *Vrrp_Interfaces_Interface_Ipv4_Version2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (version2 *Vrrp_Interfaces_Interface_Ipv4_Version2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (version2 *Vrrp_Interfaces_Interface_Ipv4_Version2) SetParent(parent types.Entity) { version2.parent = parent }

func (version2 *Vrrp_Interfaces_Interface_Ipv4_Version2) GetParent() types.Entity { return version2.parent }

func (version2 *Vrrp_Interfaces_Interface_Ipv4_Version2) GetParentYangName() string { return "ipv4" }

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters
// The VRRP virtual router configuration table
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The VRRP virtual router being configured. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter.
    VirtualRouter []Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters) GetFilter() yfilter.YFilter { return virtualRouters.YFilter }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters) SetFilter(yf yfilter.YFilter) { virtualRouters.YFilter = yf }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters) GetGoName(yname string) string {
    if yname == "virtual-router" { return "VirtualRouter" }
    return ""
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters) GetSegmentPath() string {
    return "virtual-routers"
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "virtual-router" {
        for _, c := range virtualRouters.VirtualRouter {
            if virtualRouters.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter{}
        virtualRouters.VirtualRouter = append(virtualRouters.VirtualRouter, child)
        return &virtualRouters.VirtualRouter[len(virtualRouters.VirtualRouter)-1]
    }
    return nil
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range virtualRouters.VirtualRouter {
        children[virtualRouters.VirtualRouter[i].GetSegmentPath()] = &virtualRouters.VirtualRouter[i]
    }
    return children
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters) GetBundleName() string { return "cisco_ios_xr" }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters) GetYangName() string { return "virtual-routers" }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters) SetParent(parent types.Entity) { virtualRouters.parent = parent }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters) GetParent() types.Entity { return virtualRouters.parent }

func (virtualRouters *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters) GetParentYangName() string { return "version2" }

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter
// The VRRP virtual router being configured
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter) GetFilter() yfilter.YFilter { return virtualRouter.YFilter }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter) SetFilter(yf yfilter.YFilter) { virtualRouter.YFilter = yf }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter) GetGoName(yname string) string {
    if yname == "vr-id" { return "VrId" }
    if yname == "priority" { return "Priority" }
    if yname == "primary-ipv4-address" { return "PrimaryIpv4Address" }
    if yname == "preempt" { return "Preempt" }
    if yname == "text-password" { return "TextPassword" }
    if yname == "bfd" { return "Bfd" }
    if yname == "session-name" { return "SessionName" }
    if yname == "accept-mode-disable" { return "AcceptModeDisable" }
    if yname == "timer" { return "Timer" }
    if yname == "secondary-ipv4-addresses" { return "SecondaryIpv4Addresses" }
    if yname == "tracks" { return "Tracks" }
    if yname == "tracked-objects" { return "TrackedObjects" }
    return ""
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter) GetSegmentPath() string {
    return "virtual-router" + "[vr-id='" + fmt.Sprintf("%v", virtualRouter.VrId) + "']"
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "timer" {
        return &virtualRouter.Timer
    }
    if childYangName == "secondary-ipv4-addresses" {
        return &virtualRouter.SecondaryIpv4Addresses
    }
    if childYangName == "tracks" {
        return &virtualRouter.Tracks
    }
    if childYangName == "tracked-objects" {
        return &virtualRouter.TrackedObjects
    }
    return nil
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["timer"] = &virtualRouter.Timer
    children["secondary-ipv4-addresses"] = &virtualRouter.SecondaryIpv4Addresses
    children["tracks"] = &virtualRouter.Tracks
    children["tracked-objects"] = &virtualRouter.TrackedObjects
    return children
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vr-id"] = virtualRouter.VrId
    leafs["priority"] = virtualRouter.Priority
    leafs["primary-ipv4-address"] = virtualRouter.PrimaryIpv4Address
    leafs["preempt"] = virtualRouter.Preempt
    leafs["text-password"] = virtualRouter.TextPassword
    leafs["bfd"] = virtualRouter.Bfd
    leafs["session-name"] = virtualRouter.SessionName
    leafs["accept-mode-disable"] = virtualRouter.AcceptModeDisable
    return leafs
}

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter) GetBundleName() string { return "cisco_ios_xr" }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter) GetYangName() string { return "virtual-router" }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter) SetParent(parent types.Entity) { virtualRouter.parent = parent }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter) GetParent() types.Entity { return virtualRouter.parent }

func (virtualRouter *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter) GetParentYangName() string { return "virtual-routers" }

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer
// Set advertisement timer
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TRUE - Advertise time configured in milliseconds, FALSE - Advertise time
    // configured in seconds. The type is bool. The default value is false.
    InMsec interface{}

    // Advertisement time in milliseconds. The type is interface{} with range:
    // 100..40950. Units are millisecond.
    AdvertisementTimeInMsec interface{}

    // Advertisement time in seconds. The type is interface{} with range: 1..255.
    // Units are second.
    AdvertisementTimeInSec interface{}

    // TRUE - Force configured timer values to be used, required when configured
    // in milliseconds. The type is bool. The default value is false.
    Forced interface{}
}

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer) GetFilter() yfilter.YFilter { return timer.YFilter }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer) SetFilter(yf yfilter.YFilter) { timer.YFilter = yf }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer) GetGoName(yname string) string {
    if yname == "in-msec" { return "InMsec" }
    if yname == "advertisement-time-in-msec" { return "AdvertisementTimeInMsec" }
    if yname == "advertisement-time-in-sec" { return "AdvertisementTimeInSec" }
    if yname == "forced" { return "Forced" }
    return ""
}

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer) GetSegmentPath() string {
    return "timer"
}

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["in-msec"] = timer.InMsec
    leafs["advertisement-time-in-msec"] = timer.AdvertisementTimeInMsec
    leafs["advertisement-time-in-sec"] = timer.AdvertisementTimeInSec
    leafs["forced"] = timer.Forced
    return leafs
}

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer) GetBundleName() string { return "cisco_ios_xr" }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer) GetYangName() string { return "timer" }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer) SetParent(parent types.Entity) { timer.parent = parent }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer) GetParent() types.Entity { return timer.parent }

func (timer *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Timer) GetParentYangName() string { return "virtual-router" }

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses
// The table of VRRP secondary IPv4 addresses
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A VRRP secondary IPv4 address. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address.
    SecondaryIpv4Address []Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetFilter() yfilter.YFilter { return secondaryIpv4Addresses.YFilter }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) SetFilter(yf yfilter.YFilter) { secondaryIpv4Addresses.YFilter = yf }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetGoName(yname string) string {
    if yname == "secondary-ipv4-address" { return "SecondaryIpv4Address" }
    return ""
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetSegmentPath() string {
    return "secondary-ipv4-addresses"
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "secondary-ipv4-address" {
        for _, c := range secondaryIpv4Addresses.SecondaryIpv4Address {
            if secondaryIpv4Addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address{}
        secondaryIpv4Addresses.SecondaryIpv4Address = append(secondaryIpv4Addresses.SecondaryIpv4Address, child)
        return &secondaryIpv4Addresses.SecondaryIpv4Address[len(secondaryIpv4Addresses.SecondaryIpv4Address)-1]
    }
    return nil
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range secondaryIpv4Addresses.SecondaryIpv4Address {
        children[secondaryIpv4Addresses.SecondaryIpv4Address[i].GetSegmentPath()] = &secondaryIpv4Addresses.SecondaryIpv4Address[i]
    }
    return children
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetYangName() string { return "secondary-ipv4-addresses" }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) SetParent(parent types.Entity) { secondaryIpv4Addresses.parent = parent }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetParent() types.Entity { return secondaryIpv4Addresses.parent }

func (secondaryIpv4Addresses *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses) GetParentYangName() string { return "virtual-router" }

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address
// A VRRP secondary IPv4 address
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRRP Secondary IPv4 address. The type is string
    // with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    IpAddress interface{}
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetFilter() yfilter.YFilter { return secondaryIpv4Address.YFilter }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) SetFilter(yf yfilter.YFilter) { secondaryIpv4Address.YFilter = yf }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetGoName(yname string) string {
    if yname == "ip-address" { return "IpAddress" }
    return ""
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetSegmentPath() string {
    return "secondary-ipv4-address" + "[ip-address='" + fmt.Sprintf("%v", secondaryIpv4Address.IpAddress) + "']"
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ip-address"] = secondaryIpv4Address.IpAddress
    return leafs
}

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetYangName() string { return "secondary-ipv4-address" }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) SetParent(parent types.Entity) { secondaryIpv4Address.parent = parent }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetParent() types.Entity { return secondaryIpv4Address.parent }

func (secondaryIpv4Address *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_SecondaryIpv4Addresses_SecondaryIpv4Address) GetParentYangName() string { return "secondary-ipv4-addresses" }

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks
// Track an item, reducing priority if it
// goes down
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object to be tracked. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track.
    Track []Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track
}

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks) GetFilter() yfilter.YFilter { return tracks.YFilter }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks) SetFilter(yf yfilter.YFilter) { tracks.YFilter = yf }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks) GetGoName(yname string) string {
    if yname == "track" { return "Track" }
    return ""
}

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks) GetSegmentPath() string {
    return "tracks"
}

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "track" {
        for _, c := range tracks.Track {
            if tracks.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track{}
        tracks.Track = append(tracks.Track, child)
        return &tracks.Track[len(tracks.Track)-1]
    }
    return nil
}

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range tracks.Track {
        children[tracks.Track[i].GetSegmentPath()] = &tracks.Track[i]
    }
    return children
}

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks) GetBundleName() string { return "cisco_ios_xr" }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks) GetYangName() string { return "tracks" }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks) SetParent(parent types.Entity) { tracks.parent = parent }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks) GetParent() types.Entity { return tracks.parent }

func (tracks *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks) GetParentYangName() string { return "virtual-router" }

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track
// Object to be tracked
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object to be tracked, interface name for
    // interfaces. The type is string with pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Priority decrement. The type is interface{} with range: 1..254. This
    // attribute is mandatory.
    Priority interface{}
}

func (track *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track) GetFilter() yfilter.YFilter { return track.YFilter }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track) SetFilter(yf yfilter.YFilter) { track.YFilter = yf }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "priority" { return "Priority" }
    return ""
}

func (track *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track) GetSegmentPath() string {
    return "track" + "[interface-name='" + fmt.Sprintf("%v", track.InterfaceName) + "']"
}

func (track *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (track *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (track *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = track.InterfaceName
    leafs["priority"] = track.Priority
    return leafs
}

func (track *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track) GetBundleName() string { return "cisco_ios_xr" }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track) GetYangName() string { return "track" }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track) SetParent(parent types.Entity) { track.parent = parent }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track) GetParent() types.Entity { return track.parent }

func (track *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_Tracks_Track) GetParentYangName() string { return "tracks" }

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects
// Track an object, reducing priority if it
// goes down
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object to be tracked. The type is slice of
    // Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject.
    TrackedObject []Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects) GetFilter() yfilter.YFilter { return trackedObjects.YFilter }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects) SetFilter(yf yfilter.YFilter) { trackedObjects.YFilter = yf }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects) GetGoName(yname string) string {
    if yname == "tracked-object" { return "TrackedObject" }
    return ""
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects) GetSegmentPath() string {
    return "tracked-objects"
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tracked-object" {
        for _, c := range trackedObjects.TrackedObject {
            if trackedObjects.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject{}
        trackedObjects.TrackedObject = append(trackedObjects.TrackedObject, child)
        return &trackedObjects.TrackedObject[len(trackedObjects.TrackedObject)-1]
    }
    return nil
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackedObjects.TrackedObject {
        children[trackedObjects.TrackedObject[i].GetSegmentPath()] = &trackedObjects.TrackedObject[i]
    }
    return children
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects) GetBundleName() string { return "cisco_ios_xr" }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects) GetYangName() string { return "tracked-objects" }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects) SetParent(parent types.Entity) { trackedObjects.parent = parent }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects) GetParent() types.Entity { return trackedObjects.parent }

func (trackedObjects *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects) GetParentYangName() string { return "virtual-router" }

// Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject
// Object to be tracked
type Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object to be tracked, interface name for
    // interfaces. The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ObjectName interface{}

    // Priority decrement. The type is interface{} with range: 1..254. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetFilter() yfilter.YFilter { return trackedObject.YFilter }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) SetFilter(yf yfilter.YFilter) { trackedObject.YFilter = yf }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetSegmentPath() string {
    return "tracked-object" + "[object-name='" + fmt.Sprintf("%v", trackedObject.ObjectName) + "']"
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = trackedObject.ObjectName
    leafs["priority-decrement"] = trackedObject.PriorityDecrement
    return leafs
}

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetBundleName() string { return "cisco_ios_xr" }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetYangName() string { return "tracked-object" }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) SetParent(parent types.Entity) { trackedObject.parent = parent }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetParent() types.Entity { return trackedObject.parent }

func (trackedObject *Vrrp_Interfaces_Interface_Ipv4_Version2_VirtualRouters_VirtualRouter_TrackedObjects_TrackedObject) GetParentYangName() string { return "tracked-objects" }

// Vrrp_Interfaces_Interface_Bfd
// BFD configuration
type Vrrp_Interfaces_Interface_Bfd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Hello interval for BFD sessions created by vrrp. The type is interface{}
    // with range: 3..30000. Units are millisecond.
    Interval interface{}

    // Detection multiplier for BFD sessions created by vrrp. The type is
    // interface{} with range: 2..50.
    DetectionMultiplier interface{}
}

func (bfd *Vrrp_Interfaces_Interface_Bfd) GetFilter() yfilter.YFilter { return bfd.YFilter }

func (bfd *Vrrp_Interfaces_Interface_Bfd) SetFilter(yf yfilter.YFilter) { bfd.YFilter = yf }

func (bfd *Vrrp_Interfaces_Interface_Bfd) GetGoName(yname string) string {
    if yname == "interval" { return "Interval" }
    if yname == "detection-multiplier" { return "DetectionMultiplier" }
    return ""
}

func (bfd *Vrrp_Interfaces_Interface_Bfd) GetSegmentPath() string {
    return "bfd"
}

func (bfd *Vrrp_Interfaces_Interface_Bfd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfd *Vrrp_Interfaces_Interface_Bfd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfd *Vrrp_Interfaces_Interface_Bfd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interval"] = bfd.Interval
    leafs["detection-multiplier"] = bfd.DetectionMultiplier
    return leafs
}

func (bfd *Vrrp_Interfaces_Interface_Bfd) GetBundleName() string { return "cisco_ios_xr" }

func (bfd *Vrrp_Interfaces_Interface_Bfd) GetYangName() string { return "bfd" }

func (bfd *Vrrp_Interfaces_Interface_Bfd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfd *Vrrp_Interfaces_Interface_Bfd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfd *Vrrp_Interfaces_Interface_Bfd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfd *Vrrp_Interfaces_Interface_Bfd) SetParent(parent types.Entity) { bfd.parent = parent }

func (bfd *Vrrp_Interfaces_Interface_Bfd) GetParent() types.Entity { return bfd.parent }

func (bfd *Vrrp_Interfaces_Interface_Bfd) GetParentYangName() string { return "interface" }

