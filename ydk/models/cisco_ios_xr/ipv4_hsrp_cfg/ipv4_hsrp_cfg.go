// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-hsrp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hsrp: HSRP configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-snmp-agent-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_hsrp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_hsrp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-hsrp-cfg hsrp}", reflect.TypeOf(Hsrp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-hsrp-cfg:hsrp", reflect.TypeOf(Hsrp{}))
}

// HsrpLinklocal represents Hsrp linklocal
type HsrpLinklocal string

const (
    // Manual Linklocal address configuration
    HsrpLinklocal_manual HsrpLinklocal = "manual"

    // Automatic Linklocal address configuration
    HsrpLinklocal_auto HsrpLinklocal = "auto"

    // Automatic legacy-compatible Linklocal address
    // configuration
    HsrpLinklocal_legacy HsrpLinklocal = "legacy"
)

// Hsrp
// HSRP configuration
type Hsrp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface Table for HSRP configuration.
    Interfaces Hsrp_Interfaces

    // HSRP logging options.
    Logging Hsrp_Logging
}

func (hsrp *Hsrp) GetFilter() yfilter.YFilter { return hsrp.YFilter }

func (hsrp *Hsrp) SetFilter(yf yfilter.YFilter) { hsrp.YFilter = yf }

func (hsrp *Hsrp) GetGoName(yname string) string {
    if yname == "interfaces" { return "Interfaces" }
    if yname == "logging" { return "Logging" }
    return ""
}

func (hsrp *Hsrp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-hsrp-cfg:hsrp"
}

func (hsrp *Hsrp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interfaces" {
        return &hsrp.Interfaces
    }
    if childYangName == "logging" {
        return &hsrp.Logging
    }
    return nil
}

func (hsrp *Hsrp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["interfaces"] = &hsrp.Interfaces
    children["logging"] = &hsrp.Logging
    return children
}

func (hsrp *Hsrp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (hsrp *Hsrp) GetBundleName() string { return "cisco_ios_xr" }

func (hsrp *Hsrp) GetYangName() string { return "hsrp" }

func (hsrp *Hsrp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (hsrp *Hsrp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (hsrp *Hsrp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (hsrp *Hsrp) SetParent(parent types.Entity) { hsrp.parent = parent }

func (hsrp *Hsrp) GetParent() types.Entity { return hsrp.parent }

func (hsrp *Hsrp) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-hsrp-cfg" }

// Hsrp_Interfaces
// Interface Table for HSRP configuration
type Hsrp_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Per-interface HSRP configuration. The type is slice of
    // Hsrp_Interfaces_Interface.
    Interface []Hsrp_Interfaces_Interface
}

func (interfaces *Hsrp_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Hsrp_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Hsrp_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Hsrp_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Hsrp_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Hsrp_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Hsrp_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Hsrp_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Hsrp_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Hsrp_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Hsrp_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Hsrp_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Hsrp_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Hsrp_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Hsrp_Interfaces) GetParentYangName() string { return "hsrp" }

// Hsrp_Interfaces_Interface
// Per-interface HSRP configuration
type Hsrp_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // HSRP MGO slave MAC refresh rate. The type is interface{} with range:
    // 0..10000. The default value is 60.
    MacRefresh interface{}

    // Use burned-in address. The type is interface{}.
    UseBia interface{}

    // Disable HSRP filtered ICMP redirects. The type is interface{}.
    RedirectsDisable interface{}

    // IPv6 HSRP configuration.
    Ipv6 Hsrp_Interfaces_Interface_Ipv6

    // BFD configuration.
    Bfd Hsrp_Interfaces_Interface_Bfd

    // Minimum and Reload Delay.
    Delay Hsrp_Interfaces_Interface_Delay

    // IPv4 HSRP configuration.
    Ipv4 Hsrp_Interfaces_Interface_Ipv4
}

func (self *Hsrp_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Hsrp_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Hsrp_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "mac-refresh" { return "MacRefresh" }
    if yname == "use-bia" { return "UseBia" }
    if yname == "redirects-disable" { return "RedirectsDisable" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "bfd" { return "Bfd" }
    if yname == "delay" { return "Delay" }
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (self *Hsrp_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Hsrp_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6" {
        return &self.Ipv6
    }
    if childYangName == "bfd" {
        return &self.Bfd
    }
    if childYangName == "delay" {
        return &self.Delay
    }
    if childYangName == "ipv4" {
        return &self.Ipv4
    }
    return nil
}

func (self *Hsrp_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6"] = &self.Ipv6
    children["bfd"] = &self.Bfd
    children["delay"] = &self.Delay
    children["ipv4"] = &self.Ipv4
    return children
}

func (self *Hsrp_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["mac-refresh"] = self.MacRefresh
    leafs["use-bia"] = self.UseBia
    leafs["redirects-disable"] = self.RedirectsDisable
    return leafs
}

func (self *Hsrp_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Hsrp_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Hsrp_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Hsrp_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Hsrp_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Hsrp_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Hsrp_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Hsrp_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Hsrp_Interfaces_Interface_Ipv6
// IPv6 HSRP configuration
type Hsrp_Interfaces_Interface_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Version 2 HSRP configuration.
    Version2 Hsrp_Interfaces_Interface_Ipv6_Version2

    // The HSRP slave group configuration table.
    SlaveGroups Hsrp_Interfaces_Interface_Ipv6_SlaveGroups
}

func (ipv6 *Hsrp_Interfaces_Interface_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Hsrp_Interfaces_Interface_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Hsrp_Interfaces_Interface_Ipv6) GetGoName(yname string) string {
    if yname == "version2" { return "Version2" }
    if yname == "slave-groups" { return "SlaveGroups" }
    return ""
}

func (ipv6 *Hsrp_Interfaces_Interface_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Hsrp_Interfaces_Interface_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "version2" {
        return &ipv6.Version2
    }
    if childYangName == "slave-groups" {
        return &ipv6.SlaveGroups
    }
    return nil
}

func (ipv6 *Hsrp_Interfaces_Interface_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["version2"] = &ipv6.Version2
    children["slave-groups"] = &ipv6.SlaveGroups
    return children
}

func (ipv6 *Hsrp_Interfaces_Interface_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv6 *Hsrp_Interfaces_Interface_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Hsrp_Interfaces_Interface_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Hsrp_Interfaces_Interface_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Hsrp_Interfaces_Interface_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Hsrp_Interfaces_Interface_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Hsrp_Interfaces_Interface_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Hsrp_Interfaces_Interface_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Hsrp_Interfaces_Interface_Ipv6) GetParentYangName() string { return "interface" }

// Hsrp_Interfaces_Interface_Ipv6_Version2
// Version 2 HSRP configuration
type Hsrp_Interfaces_Interface_Ipv6_Version2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The HSRP group configuration table.
    Groups Hsrp_Interfaces_Interface_Ipv6_Version2_Groups
}

func (version2 *Hsrp_Interfaces_Interface_Ipv6_Version2) GetFilter() yfilter.YFilter { return version2.YFilter }

func (version2 *Hsrp_Interfaces_Interface_Ipv6_Version2) SetFilter(yf yfilter.YFilter) { version2.YFilter = yf }

func (version2 *Hsrp_Interfaces_Interface_Ipv6_Version2) GetGoName(yname string) string {
    if yname == "groups" { return "Groups" }
    return ""
}

func (version2 *Hsrp_Interfaces_Interface_Ipv6_Version2) GetSegmentPath() string {
    return "version2"
}

func (version2 *Hsrp_Interfaces_Interface_Ipv6_Version2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "groups" {
        return &version2.Groups
    }
    return nil
}

func (version2 *Hsrp_Interfaces_Interface_Ipv6_Version2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["groups"] = &version2.Groups
    return children
}

func (version2 *Hsrp_Interfaces_Interface_Ipv6_Version2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (version2 *Hsrp_Interfaces_Interface_Ipv6_Version2) GetBundleName() string { return "cisco_ios_xr" }

func (version2 *Hsrp_Interfaces_Interface_Ipv6_Version2) GetYangName() string { return "version2" }

func (version2 *Hsrp_Interfaces_Interface_Ipv6_Version2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (version2 *Hsrp_Interfaces_Interface_Ipv6_Version2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (version2 *Hsrp_Interfaces_Interface_Ipv6_Version2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (version2 *Hsrp_Interfaces_Interface_Ipv6_Version2) SetParent(parent types.Entity) { version2.parent = parent }

func (version2 *Hsrp_Interfaces_Interface_Ipv6_Version2) GetParent() types.Entity { return version2.parent }

func (version2 *Hsrp_Interfaces_Interface_Ipv6_Version2) GetParentYangName() string { return "ipv6" }

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups
// The HSRP group configuration table
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The HSRP group being configured. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group.
    Group []Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group
}

func (groups *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups) GetYangName() string { return "groups" }

func (groups *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups) GetParent() types.Entity { return groups.parent }

func (groups *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups) GetParentYangName() string { return "version2" }

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group
// The HSRP group being configured
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP group number. The type is interface{} with
    // range: 0..4095.
    GroupNumber interface{}

    // Priority value. The type is interface{} with range: 0..255. The default
    // value is 100.
    Priority interface{}

    // Force active if higher priority. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 0.
    Preempt interface{}

    // HSRP Session name (for MGO). The type is string with length: 1..16.
    SessionName interface{}

    // HSRP MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    VirtualMacAddress interface{}

    // Enable use of Bidirectional Forwarding Detection.
    Bfd Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd

    // The HSRP tracked interface configuration table.
    TrackedInterfaces Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces

    // The HSRP tracked interface configuration table.
    TrackedObjects Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects

    // Hello and hold timers.
    Timers Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers

    // The HSRP IPv6 virtual linklocal address.
    LinkLocalIpv6Address Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address

    // The table of HSRP virtual global IPv6 addresses.
    GlobalIpv6Addresses Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses
}

func (group *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group) GetGoName(yname string) string {
    if yname == "group-number" { return "GroupNumber" }
    if yname == "priority" { return "Priority" }
    if yname == "preempt" { return "Preempt" }
    if yname == "session-name" { return "SessionName" }
    if yname == "virtual-mac-address" { return "VirtualMacAddress" }
    if yname == "bfd" { return "Bfd" }
    if yname == "tracked-interfaces" { return "TrackedInterfaces" }
    if yname == "tracked-objects" { return "TrackedObjects" }
    if yname == "timers" { return "Timers" }
    if yname == "link-local-ipv6-address" { return "LinkLocalIpv6Address" }
    if yname == "global-ipv6-addresses" { return "GlobalIpv6Addresses" }
    return ""
}

func (group *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group) GetSegmentPath() string {
    return "group" + "[group-number='" + fmt.Sprintf("%v", group.GroupNumber) + "']"
}

func (group *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bfd" {
        return &group.Bfd
    }
    if childYangName == "tracked-interfaces" {
        return &group.TrackedInterfaces
    }
    if childYangName == "tracked-objects" {
        return &group.TrackedObjects
    }
    if childYangName == "timers" {
        return &group.Timers
    }
    if childYangName == "link-local-ipv6-address" {
        return &group.LinkLocalIpv6Address
    }
    if childYangName == "global-ipv6-addresses" {
        return &group.GlobalIpv6Addresses
    }
    return nil
}

func (group *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bfd"] = &group.Bfd
    children["tracked-interfaces"] = &group.TrackedInterfaces
    children["tracked-objects"] = &group.TrackedObjects
    children["timers"] = &group.Timers
    children["link-local-ipv6-address"] = &group.LinkLocalIpv6Address
    children["global-ipv6-addresses"] = &group.GlobalIpv6Addresses
    return children
}

func (group *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-number"] = group.GroupNumber
    leafs["priority"] = group.Priority
    leafs["preempt"] = group.Preempt
    leafs["session-name"] = group.SessionName
    leafs["virtual-mac-address"] = group.VirtualMacAddress
    return leafs
}

func (group *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group) GetYangName() string { return "group" }

func (group *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group) GetParentYangName() string { return "groups" }

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd
// Enable use of Bidirectional Forwarding
// Detection
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable BFD for this remote IP. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Interface name to run BFD. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (bfd *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd) GetFilter() yfilter.YFilter { return bfd.YFilter }

func (bfd *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd) SetFilter(yf yfilter.YFilter) { bfd.YFilter = yf }

func (bfd *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (bfd *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd) GetSegmentPath() string {
    return "bfd"
}

func (bfd *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfd *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfd *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = bfd.Address
    leafs["interface-name"] = bfd.InterfaceName
    return leafs
}

func (bfd *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd) GetBundleName() string { return "cisco_ios_xr" }

func (bfd *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd) GetYangName() string { return "bfd" }

func (bfd *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfd *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfd *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfd *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd) SetParent(parent types.Entity) { bfd.parent = parent }

func (bfd *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd) GetParent() types.Entity { return bfd.parent }

func (bfd *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces
// The HSRP tracked interface configuration
// table
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface being tracked. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface.
    TrackedInterface []Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces) GetFilter() yfilter.YFilter { return trackedInterfaces.YFilter }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces) SetFilter(yf yfilter.YFilter) { trackedInterfaces.YFilter = yf }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces) GetGoName(yname string) string {
    if yname == "tracked-interface" { return "TrackedInterface" }
    return ""
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces) GetSegmentPath() string {
    return "tracked-interfaces"
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tracked-interface" {
        for _, c := range trackedInterfaces.TrackedInterface {
            if trackedInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface{}
        trackedInterfaces.TrackedInterface = append(trackedInterfaces.TrackedInterface, child)
        return &trackedInterfaces.TrackedInterface[len(trackedInterfaces.TrackedInterface)-1]
    }
    return nil
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackedInterfaces.TrackedInterface {
        children[trackedInterfaces.TrackedInterface[i].GetSegmentPath()] = &trackedInterfaces.TrackedInterface[i]
    }
    return children
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces) GetYangName() string { return "tracked-interfaces" }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces) SetParent(parent types.Entity) { trackedInterfaces.parent = parent }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces) GetParent() types.Entity { return trackedInterfaces.parent }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface
// Interface being tracked
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface being tracked. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Priority decrement. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetFilter() yfilter.YFilter { return trackedInterface.YFilter }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) SetFilter(yf yfilter.YFilter) { trackedInterface.YFilter = yf }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetSegmentPath() string {
    return "tracked-interface" + "[interface-name='" + fmt.Sprintf("%v", trackedInterface.InterfaceName) + "']"
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = trackedInterface.InterfaceName
    leafs["priority-decrement"] = trackedInterface.PriorityDecrement
    return leafs
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetBundleName() string { return "cisco_ios_xr" }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetYangName() string { return "tracked-interface" }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) SetParent(parent types.Entity) { trackedInterface.parent = parent }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetParent() types.Entity { return trackedInterface.parent }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetParentYangName() string { return "tracked-interfaces" }

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects
// The HSRP tracked interface configuration
// table
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object being tracked. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject.
    TrackedObject []Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects) GetFilter() yfilter.YFilter { return trackedObjects.YFilter }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects) SetFilter(yf yfilter.YFilter) { trackedObjects.YFilter = yf }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects) GetGoName(yname string) string {
    if yname == "tracked-object" { return "TrackedObject" }
    return ""
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects) GetSegmentPath() string {
    return "tracked-objects"
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tracked-object" {
        for _, c := range trackedObjects.TrackedObject {
            if trackedObjects.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject{}
        trackedObjects.TrackedObject = append(trackedObjects.TrackedObject, child)
        return &trackedObjects.TrackedObject[len(trackedObjects.TrackedObject)-1]
    }
    return nil
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackedObjects.TrackedObject {
        children[trackedObjects.TrackedObject[i].GetSegmentPath()] = &trackedObjects.TrackedObject[i]
    }
    return children
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects) GetBundleName() string { return "cisco_ios_xr" }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects) GetYangName() string { return "tracked-objects" }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects) SetParent(parent types.Entity) { trackedObjects.parent = parent }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects) GetParent() types.Entity { return trackedObjects.parent }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject
// Object being tracked
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface being tracked. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ObjectName interface{}

    // Priority decrement. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject) GetFilter() yfilter.YFilter { return trackedObject.YFilter }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject) SetFilter(yf yfilter.YFilter) { trackedObject.YFilter = yf }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject) GetSegmentPath() string {
    return "tracked-object" + "[object-name='" + fmt.Sprintf("%v", trackedObject.ObjectName) + "']"
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = trackedObject.ObjectName
    leafs["priority-decrement"] = trackedObject.PriorityDecrement
    return leafs
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject) GetBundleName() string { return "cisco_ios_xr" }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject) GetYangName() string { return "tracked-object" }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject) SetParent(parent types.Entity) { trackedObject.parent = parent }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject) GetParent() types.Entity { return trackedObject.parent }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject) GetParentYangName() string { return "tracked-objects" }

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers
// Hello and hold timers
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TRUE - Hello time configured in milliseconds, FALSE - Hello time configured
    // in seconds. The type is bool. The default value is false.
    HelloMsecFlag interface{}

    // Hello time in msecs. The type is interface{} with range: 100..3000. Units
    // are millisecond.
    HelloMsec interface{}

    // Hello time in seconds. The type is interface{} with range: 1..255. Units
    // are second. The default value is 3.
    HelloSec interface{}

    // TRUE - Hold time configured in milliseconds, FALSE - Hold time configured
    // in seconds. The type is bool. The default value is false.
    HoldMsecFlag interface{}

    // Hold time in msecs. The type is interface{} with range: 100..3000. Units
    // are millisecond.
    HoldMsec interface{}

    // Hold time in seconds. The type is interface{} with range: 1..255. Units are
    // second. The default value is 10.
    HoldSec interface{}
}

func (timers *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers) GetFilter() yfilter.YFilter { return timers.YFilter }

func (timers *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers) SetFilter(yf yfilter.YFilter) { timers.YFilter = yf }

func (timers *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers) GetGoName(yname string) string {
    if yname == "hello-msec-flag" { return "HelloMsecFlag" }
    if yname == "hello-msec" { return "HelloMsec" }
    if yname == "hello-sec" { return "HelloSec" }
    if yname == "hold-msec-flag" { return "HoldMsecFlag" }
    if yname == "hold-msec" { return "HoldMsec" }
    if yname == "hold-sec" { return "HoldSec" }
    return ""
}

func (timers *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers) GetSegmentPath() string {
    return "timers"
}

func (timers *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timers *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timers *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hello-msec-flag"] = timers.HelloMsecFlag
    leafs["hello-msec"] = timers.HelloMsec
    leafs["hello-sec"] = timers.HelloSec
    leafs["hold-msec-flag"] = timers.HoldMsecFlag
    leafs["hold-msec"] = timers.HoldMsec
    leafs["hold-sec"] = timers.HoldSec
    return leafs
}

func (timers *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers) GetBundleName() string { return "cisco_ios_xr" }

func (timers *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers) GetYangName() string { return "timers" }

func (timers *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timers *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timers *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timers *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers) SetParent(parent types.Entity) { timers.parent = parent }

func (timers *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers) GetParent() types.Entity { return timers.parent }

func (timers *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address
// The HSRP IPv6 virtual linklocal address
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // HSRP IPv6 virtual linklocal address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Linklocal Configuration Type. The type is HsrpLinklocal. The default value
    // is manual.
    AutoConfigure interface{}
}

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address) GetFilter() yfilter.YFilter { return linkLocalIpv6Address.YFilter }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address) SetFilter(yf yfilter.YFilter) { linkLocalIpv6Address.YFilter = yf }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "auto-configure" { return "AutoConfigure" }
    return ""
}

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address) GetSegmentPath() string {
    return "link-local-ipv6-address"
}

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = linkLocalIpv6Address.Address
    leafs["auto-configure"] = linkLocalIpv6Address.AutoConfigure
    return leafs
}

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address) GetYangName() string { return "link-local-ipv6-address" }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address) SetParent(parent types.Entity) { linkLocalIpv6Address.parent = parent }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address) GetParent() types.Entity { return linkLocalIpv6Address.parent }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses
// The table of HSRP virtual global IPv6
// addresses
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A HSRP virtual global IPv6 IP address. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address.
    GlobalIpv6Address []Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address
}

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses) GetFilter() yfilter.YFilter { return globalIpv6Addresses.YFilter }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses) SetFilter(yf yfilter.YFilter) { globalIpv6Addresses.YFilter = yf }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses) GetGoName(yname string) string {
    if yname == "global-ipv6-address" { return "GlobalIpv6Address" }
    return ""
}

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses) GetSegmentPath() string {
    return "global-ipv6-addresses"
}

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global-ipv6-address" {
        for _, c := range globalIpv6Addresses.GlobalIpv6Address {
            if globalIpv6Addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address{}
        globalIpv6Addresses.GlobalIpv6Address = append(globalIpv6Addresses.GlobalIpv6Address, child)
        return &globalIpv6Addresses.GlobalIpv6Address[len(globalIpv6Addresses.GlobalIpv6Address)-1]
    }
    return nil
}

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range globalIpv6Addresses.GlobalIpv6Address {
        children[globalIpv6Addresses.GlobalIpv6Address[i].GetSegmentPath()] = &globalIpv6Addresses.GlobalIpv6Address[i]
    }
    return children
}

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses) GetYangName() string { return "global-ipv6-addresses" }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses) SetParent(parent types.Entity) { globalIpv6Addresses.parent = parent }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses) GetParent() types.Entity { return globalIpv6Addresses.parent }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address
// A HSRP virtual global IPv6 IP address
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP virtual global IPv6 address. The type is
    // string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address) GetFilter() yfilter.YFilter { return globalIpv6Address.YFilter }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address) SetFilter(yf yfilter.YFilter) { globalIpv6Address.YFilter = yf }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address) GetSegmentPath() string {
    return "global-ipv6-address" + "[address='" + fmt.Sprintf("%v", globalIpv6Address.Address) + "']"
}

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = globalIpv6Address.Address
    return leafs
}

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address) GetYangName() string { return "global-ipv6-address" }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address) SetParent(parent types.Entity) { globalIpv6Address.parent = parent }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address) GetParent() types.Entity { return globalIpv6Address.parent }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address) GetParentYangName() string { return "global-ipv6-addresses" }

// Hsrp_Interfaces_Interface_Ipv6_SlaveGroups
// The HSRP slave group configuration table
type Hsrp_Interfaces_Interface_Ipv6_SlaveGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The HSRP slave group being configured. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup.
    SlaveGroup []Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup
}

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups) GetFilter() yfilter.YFilter { return slaveGroups.YFilter }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups) SetFilter(yf yfilter.YFilter) { slaveGroups.YFilter = yf }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups) GetGoName(yname string) string {
    if yname == "slave-group" { return "SlaveGroup" }
    return ""
}

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups) GetSegmentPath() string {
    return "slave-groups"
}

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slave-group" {
        for _, c := range slaveGroups.SlaveGroup {
            if slaveGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup{}
        slaveGroups.SlaveGroup = append(slaveGroups.SlaveGroup, child)
        return &slaveGroups.SlaveGroup[len(slaveGroups.SlaveGroup)-1]
    }
    return nil
}

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slaveGroups.SlaveGroup {
        children[slaveGroups.SlaveGroup[i].GetSegmentPath()] = &slaveGroups.SlaveGroup[i]
    }
    return children
}

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups) GetBundleName() string { return "cisco_ios_xr" }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups) GetYangName() string { return "slave-groups" }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups) SetParent(parent types.Entity) { slaveGroups.parent = parent }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups) GetParent() types.Entity { return slaveGroups.parent }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups) GetParentYangName() string { return "ipv6" }

// Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup
// The HSRP slave group being configured
type Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP group number. The type is interface{} with
    // range: 0..4095.
    SlaveGroupNumber interface{}

    // HSRP Group name for this slave to follow. The type is string.
    Follow interface{}

    // HSRP MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    VirtualMacAddress interface{}

    // The HSRP IPv6 virtual linklocal address.
    LinkLocalIpv6Address Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address

    // The table of HSRP virtual global IPv6 addresses.
    GlobalIpv6Addresses Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses
}

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup) GetFilter() yfilter.YFilter { return slaveGroup.YFilter }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup) SetFilter(yf yfilter.YFilter) { slaveGroup.YFilter = yf }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup) GetGoName(yname string) string {
    if yname == "slave-group-number" { return "SlaveGroupNumber" }
    if yname == "follow" { return "Follow" }
    if yname == "virtual-mac-address" { return "VirtualMacAddress" }
    if yname == "link-local-ipv6-address" { return "LinkLocalIpv6Address" }
    if yname == "global-ipv6-addresses" { return "GlobalIpv6Addresses" }
    return ""
}

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup) GetSegmentPath() string {
    return "slave-group" + "[slave-group-number='" + fmt.Sprintf("%v", slaveGroup.SlaveGroupNumber) + "']"
}

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "link-local-ipv6-address" {
        return &slaveGroup.LinkLocalIpv6Address
    }
    if childYangName == "global-ipv6-addresses" {
        return &slaveGroup.GlobalIpv6Addresses
    }
    return nil
}

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["link-local-ipv6-address"] = &slaveGroup.LinkLocalIpv6Address
    children["global-ipv6-addresses"] = &slaveGroup.GlobalIpv6Addresses
    return children
}

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slave-group-number"] = slaveGroup.SlaveGroupNumber
    leafs["follow"] = slaveGroup.Follow
    leafs["virtual-mac-address"] = slaveGroup.VirtualMacAddress
    return leafs
}

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup) GetBundleName() string { return "cisco_ios_xr" }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup) GetYangName() string { return "slave-group" }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup) SetParent(parent types.Entity) { slaveGroup.parent = parent }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup) GetParent() types.Entity { return slaveGroup.parent }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup) GetParentYangName() string { return "slave-groups" }

// Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address
// The HSRP IPv6 virtual linklocal address
type Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // HSRP IPv6 virtual linklocal address. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Linklocal Configuration Type. The type is HsrpLinklocal. The default value
    // is manual.
    AutoConfigure interface{}
}

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address) GetFilter() yfilter.YFilter { return linkLocalIpv6Address.YFilter }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address) SetFilter(yf yfilter.YFilter) { linkLocalIpv6Address.YFilter = yf }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "auto-configure" { return "AutoConfigure" }
    return ""
}

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address) GetSegmentPath() string {
    return "link-local-ipv6-address"
}

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = linkLocalIpv6Address.Address
    leafs["auto-configure"] = linkLocalIpv6Address.AutoConfigure
    return leafs
}

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address) GetYangName() string { return "link-local-ipv6-address" }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address) SetParent(parent types.Entity) { linkLocalIpv6Address.parent = parent }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address) GetParent() types.Entity { return linkLocalIpv6Address.parent }

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address) GetParentYangName() string { return "slave-group" }

// Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses
// The table of HSRP virtual global IPv6
// addresses
type Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A HSRP virtual global IPv6 IP address. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address.
    GlobalIpv6Address []Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address
}

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses) GetFilter() yfilter.YFilter { return globalIpv6Addresses.YFilter }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses) SetFilter(yf yfilter.YFilter) { globalIpv6Addresses.YFilter = yf }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses) GetGoName(yname string) string {
    if yname == "global-ipv6-address" { return "GlobalIpv6Address" }
    return ""
}

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses) GetSegmentPath() string {
    return "global-ipv6-addresses"
}

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "global-ipv6-address" {
        for _, c := range globalIpv6Addresses.GlobalIpv6Address {
            if globalIpv6Addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address{}
        globalIpv6Addresses.GlobalIpv6Address = append(globalIpv6Addresses.GlobalIpv6Address, child)
        return &globalIpv6Addresses.GlobalIpv6Address[len(globalIpv6Addresses.GlobalIpv6Address)-1]
    }
    return nil
}

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range globalIpv6Addresses.GlobalIpv6Address {
        children[globalIpv6Addresses.GlobalIpv6Address[i].GetSegmentPath()] = &globalIpv6Addresses.GlobalIpv6Address[i]
    }
    return children
}

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses) GetYangName() string { return "global-ipv6-addresses" }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses) SetParent(parent types.Entity) { globalIpv6Addresses.parent = parent }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses) GetParent() types.Entity { return globalIpv6Addresses.parent }

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses) GetParentYangName() string { return "slave-group" }

// Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address
// A HSRP virtual global IPv6 IP address
type Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP virtual global IPv6 address. The type is
    // string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address) GetFilter() yfilter.YFilter { return globalIpv6Address.YFilter }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address) SetFilter(yf yfilter.YFilter) { globalIpv6Address.YFilter = yf }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address) GetSegmentPath() string {
    return "global-ipv6-address" + "[address='" + fmt.Sprintf("%v", globalIpv6Address.Address) + "']"
}

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = globalIpv6Address.Address
    return leafs
}

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address) GetBundleName() string { return "cisco_ios_xr" }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address) GetYangName() string { return "global-ipv6-address" }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address) SetParent(parent types.Entity) { globalIpv6Address.parent = parent }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address) GetParent() types.Entity { return globalIpv6Address.parent }

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address) GetParentYangName() string { return "global-ipv6-addresses" }

// Hsrp_Interfaces_Interface_Bfd
// BFD configuration
type Hsrp_Interfaces_Interface_Bfd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Detection multiplier for BFD sessions created by hsrp. The type is
    // interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Hello interval for BFD sessions created by hsrp. The type is interface{}
    // with range: 3..30000. Units are millisecond.
    Interval interface{}
}

func (bfd *Hsrp_Interfaces_Interface_Bfd) GetFilter() yfilter.YFilter { return bfd.YFilter }

func (bfd *Hsrp_Interfaces_Interface_Bfd) SetFilter(yf yfilter.YFilter) { bfd.YFilter = yf }

func (bfd *Hsrp_Interfaces_Interface_Bfd) GetGoName(yname string) string {
    if yname == "detection-multiplier" { return "DetectionMultiplier" }
    if yname == "interval" { return "Interval" }
    return ""
}

func (bfd *Hsrp_Interfaces_Interface_Bfd) GetSegmentPath() string {
    return "bfd"
}

func (bfd *Hsrp_Interfaces_Interface_Bfd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfd *Hsrp_Interfaces_Interface_Bfd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfd *Hsrp_Interfaces_Interface_Bfd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["detection-multiplier"] = bfd.DetectionMultiplier
    leafs["interval"] = bfd.Interval
    return leafs
}

func (bfd *Hsrp_Interfaces_Interface_Bfd) GetBundleName() string { return "cisco_ios_xr" }

func (bfd *Hsrp_Interfaces_Interface_Bfd) GetYangName() string { return "bfd" }

func (bfd *Hsrp_Interfaces_Interface_Bfd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfd *Hsrp_Interfaces_Interface_Bfd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfd *Hsrp_Interfaces_Interface_Bfd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfd *Hsrp_Interfaces_Interface_Bfd) SetParent(parent types.Entity) { bfd.parent = parent }

func (bfd *Hsrp_Interfaces_Interface_Bfd) GetParent() types.Entity { return bfd.parent }

func (bfd *Hsrp_Interfaces_Interface_Bfd) GetParentYangName() string { return "interface" }

// Hsrp_Interfaces_Interface_Delay
// Minimum and Reload Delay
type Hsrp_Interfaces_Interface_Delay struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Minimum delay in seconds. The type is interface{} with range: 0..10000.
    // Units are second. The default value is 1.
    MinimumDelay interface{}

    // Reload delay in seconds. The type is interface{} with range: 0..10000.
    // Units are second. The default value is 5.
    ReloadDelay interface{}
}

func (delay *Hsrp_Interfaces_Interface_Delay) GetFilter() yfilter.YFilter { return delay.YFilter }

func (delay *Hsrp_Interfaces_Interface_Delay) SetFilter(yf yfilter.YFilter) { delay.YFilter = yf }

func (delay *Hsrp_Interfaces_Interface_Delay) GetGoName(yname string) string {
    if yname == "minimum-delay" { return "MinimumDelay" }
    if yname == "reload-delay" { return "ReloadDelay" }
    return ""
}

func (delay *Hsrp_Interfaces_Interface_Delay) GetSegmentPath() string {
    return "delay"
}

func (delay *Hsrp_Interfaces_Interface_Delay) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (delay *Hsrp_Interfaces_Interface_Delay) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (delay *Hsrp_Interfaces_Interface_Delay) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["minimum-delay"] = delay.MinimumDelay
    leafs["reload-delay"] = delay.ReloadDelay
    return leafs
}

func (delay *Hsrp_Interfaces_Interface_Delay) GetBundleName() string { return "cisco_ios_xr" }

func (delay *Hsrp_Interfaces_Interface_Delay) GetYangName() string { return "delay" }

func (delay *Hsrp_Interfaces_Interface_Delay) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (delay *Hsrp_Interfaces_Interface_Delay) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (delay *Hsrp_Interfaces_Interface_Delay) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (delay *Hsrp_Interfaces_Interface_Delay) SetParent(parent types.Entity) { delay.parent = parent }

func (delay *Hsrp_Interfaces_Interface_Delay) GetParent() types.Entity { return delay.parent }

func (delay *Hsrp_Interfaces_Interface_Delay) GetParentYangName() string { return "interface" }

// Hsrp_Interfaces_Interface_Ipv4
// IPv4 HSRP configuration
type Hsrp_Interfaces_Interface_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The HSRP slave group configuration table.
    SlaveGroups Hsrp_Interfaces_Interface_Ipv4_SlaveGroups

    // Version 1 HSRP configuration.
    Version1 Hsrp_Interfaces_Interface_Ipv4_Version1

    // Version 2 HSRP configuration.
    Version2 Hsrp_Interfaces_Interface_Ipv4_Version2
}

func (ipv4 *Hsrp_Interfaces_Interface_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Hsrp_Interfaces_Interface_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Hsrp_Interfaces_Interface_Ipv4) GetGoName(yname string) string {
    if yname == "slave-groups" { return "SlaveGroups" }
    if yname == "version1" { return "Version1" }
    if yname == "version2" { return "Version2" }
    return ""
}

func (ipv4 *Hsrp_Interfaces_Interface_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Hsrp_Interfaces_Interface_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slave-groups" {
        return &ipv4.SlaveGroups
    }
    if childYangName == "version1" {
        return &ipv4.Version1
    }
    if childYangName == "version2" {
        return &ipv4.Version2
    }
    return nil
}

func (ipv4 *Hsrp_Interfaces_Interface_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["slave-groups"] = &ipv4.SlaveGroups
    children["version1"] = &ipv4.Version1
    children["version2"] = &ipv4.Version2
    return children
}

func (ipv4 *Hsrp_Interfaces_Interface_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipv4 *Hsrp_Interfaces_Interface_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Hsrp_Interfaces_Interface_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Hsrp_Interfaces_Interface_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Hsrp_Interfaces_Interface_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Hsrp_Interfaces_Interface_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Hsrp_Interfaces_Interface_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Hsrp_Interfaces_Interface_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Hsrp_Interfaces_Interface_Ipv4) GetParentYangName() string { return "interface" }

// Hsrp_Interfaces_Interface_Ipv4_SlaveGroups
// The HSRP slave group configuration table
type Hsrp_Interfaces_Interface_Ipv4_SlaveGroups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The HSRP slave group being configured. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup.
    SlaveGroup []Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup
}

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups) GetFilter() yfilter.YFilter { return slaveGroups.YFilter }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups) SetFilter(yf yfilter.YFilter) { slaveGroups.YFilter = yf }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups) GetGoName(yname string) string {
    if yname == "slave-group" { return "SlaveGroup" }
    return ""
}

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups) GetSegmentPath() string {
    return "slave-groups"
}

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "slave-group" {
        for _, c := range slaveGroups.SlaveGroup {
            if slaveGroups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup{}
        slaveGroups.SlaveGroup = append(slaveGroups.SlaveGroup, child)
        return &slaveGroups.SlaveGroup[len(slaveGroups.SlaveGroup)-1]
    }
    return nil
}

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range slaveGroups.SlaveGroup {
        children[slaveGroups.SlaveGroup[i].GetSegmentPath()] = &slaveGroups.SlaveGroup[i]
    }
    return children
}

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups) GetBundleName() string { return "cisco_ios_xr" }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups) GetYangName() string { return "slave-groups" }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups) SetParent(parent types.Entity) { slaveGroups.parent = parent }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups) GetParent() types.Entity { return slaveGroups.parent }

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups) GetParentYangName() string { return "ipv4" }

// Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup
// The HSRP slave group being configured
type Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP group number. The type is interface{} with
    // range: 0..4095.
    SlaveGroupNumber interface{}

    // HSRP Group name for this slave to follow. The type is string.
    Follow interface{}

    // HSRP MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    VirtualMacAddress interface{}

    // Primary HSRP IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    PrimaryIpv4Address interface{}

    // Secondary HSRP IP address Table.
    SecondaryIpv4Addresses Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses
}

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup) GetFilter() yfilter.YFilter { return slaveGroup.YFilter }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup) SetFilter(yf yfilter.YFilter) { slaveGroup.YFilter = yf }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup) GetGoName(yname string) string {
    if yname == "slave-group-number" { return "SlaveGroupNumber" }
    if yname == "follow" { return "Follow" }
    if yname == "virtual-mac-address" { return "VirtualMacAddress" }
    if yname == "primary-ipv4-address" { return "PrimaryIpv4Address" }
    if yname == "secondary-ipv4-addresses" { return "SecondaryIpv4Addresses" }
    return ""
}

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup) GetSegmentPath() string {
    return "slave-group" + "[slave-group-number='" + fmt.Sprintf("%v", slaveGroup.SlaveGroupNumber) + "']"
}

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "secondary-ipv4-addresses" {
        return &slaveGroup.SecondaryIpv4Addresses
    }
    return nil
}

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["secondary-ipv4-addresses"] = &slaveGroup.SecondaryIpv4Addresses
    return children
}

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["slave-group-number"] = slaveGroup.SlaveGroupNumber
    leafs["follow"] = slaveGroup.Follow
    leafs["virtual-mac-address"] = slaveGroup.VirtualMacAddress
    leafs["primary-ipv4-address"] = slaveGroup.PrimaryIpv4Address
    return leafs
}

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup) GetBundleName() string { return "cisco_ios_xr" }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup) GetYangName() string { return "slave-group" }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup) SetParent(parent types.Entity) { slaveGroup.parent = parent }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup) GetParent() types.Entity { return slaveGroup.parent }

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup) GetParentYangName() string { return "slave-groups" }

// Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses
// Secondary HSRP IP address Table
type Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Secondary HSRP IP address. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address.
    SecondaryIpv4Address []Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses) GetFilter() yfilter.YFilter { return secondaryIpv4Addresses.YFilter }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses) SetFilter(yf yfilter.YFilter) { secondaryIpv4Addresses.YFilter = yf }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses) GetGoName(yname string) string {
    if yname == "secondary-ipv4-address" { return "SecondaryIpv4Address" }
    return ""
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses) GetSegmentPath() string {
    return "secondary-ipv4-addresses"
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "secondary-ipv4-address" {
        for _, c := range secondaryIpv4Addresses.SecondaryIpv4Address {
            if secondaryIpv4Addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address{}
        secondaryIpv4Addresses.SecondaryIpv4Address = append(secondaryIpv4Addresses.SecondaryIpv4Address, child)
        return &secondaryIpv4Addresses.SecondaryIpv4Address[len(secondaryIpv4Addresses.SecondaryIpv4Address)-1]
    }
    return nil
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range secondaryIpv4Addresses.SecondaryIpv4Address {
        children[secondaryIpv4Addresses.SecondaryIpv4Address[i].GetSegmentPath()] = &secondaryIpv4Addresses.SecondaryIpv4Address[i]
    }
    return children
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses) GetYangName() string { return "secondary-ipv4-addresses" }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses) SetParent(parent types.Entity) { secondaryIpv4Addresses.parent = parent }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses) GetParent() types.Entity { return secondaryIpv4Addresses.parent }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses) GetParentYangName() string { return "slave-group" }

// Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address
// Secondary HSRP IP address
type Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address) GetFilter() yfilter.YFilter { return secondaryIpv4Address.YFilter }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address) SetFilter(yf yfilter.YFilter) { secondaryIpv4Address.YFilter = yf }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address) GetSegmentPath() string {
    return "secondary-ipv4-address" + "[address='" + fmt.Sprintf("%v", secondaryIpv4Address.Address) + "']"
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = secondaryIpv4Address.Address
    return leafs
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address) GetYangName() string { return "secondary-ipv4-address" }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address) SetParent(parent types.Entity) { secondaryIpv4Address.parent = parent }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address) GetParent() types.Entity { return secondaryIpv4Address.parent }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address) GetParentYangName() string { return "secondary-ipv4-addresses" }

// Hsrp_Interfaces_Interface_Ipv4_Version1
// Version 1 HSRP configuration
type Hsrp_Interfaces_Interface_Ipv4_Version1 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The HSRP group configuration table.
    Groups Hsrp_Interfaces_Interface_Ipv4_Version1_Groups
}

func (version1 *Hsrp_Interfaces_Interface_Ipv4_Version1) GetFilter() yfilter.YFilter { return version1.YFilter }

func (version1 *Hsrp_Interfaces_Interface_Ipv4_Version1) SetFilter(yf yfilter.YFilter) { version1.YFilter = yf }

func (version1 *Hsrp_Interfaces_Interface_Ipv4_Version1) GetGoName(yname string) string {
    if yname == "groups" { return "Groups" }
    return ""
}

func (version1 *Hsrp_Interfaces_Interface_Ipv4_Version1) GetSegmentPath() string {
    return "version1"
}

func (version1 *Hsrp_Interfaces_Interface_Ipv4_Version1) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "groups" {
        return &version1.Groups
    }
    return nil
}

func (version1 *Hsrp_Interfaces_Interface_Ipv4_Version1) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["groups"] = &version1.Groups
    return children
}

func (version1 *Hsrp_Interfaces_Interface_Ipv4_Version1) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (version1 *Hsrp_Interfaces_Interface_Ipv4_Version1) GetBundleName() string { return "cisco_ios_xr" }

func (version1 *Hsrp_Interfaces_Interface_Ipv4_Version1) GetYangName() string { return "version1" }

func (version1 *Hsrp_Interfaces_Interface_Ipv4_Version1) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (version1 *Hsrp_Interfaces_Interface_Ipv4_Version1) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (version1 *Hsrp_Interfaces_Interface_Ipv4_Version1) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (version1 *Hsrp_Interfaces_Interface_Ipv4_Version1) SetParent(parent types.Entity) { version1.parent = parent }

func (version1 *Hsrp_Interfaces_Interface_Ipv4_Version1) GetParent() types.Entity { return version1.parent }

func (version1 *Hsrp_Interfaces_Interface_Ipv4_Version1) GetParentYangName() string { return "ipv4" }

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups
// The HSRP group configuration table
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The HSRP group being configured. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group.
    Group []Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group
}

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups) GetYangName() string { return "groups" }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups) GetParent() types.Entity { return groups.parent }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups) GetParentYangName() string { return "version1" }

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group
// The HSRP group being configured
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP group number. The type is interface{} with
    // range: 0..255.
    GroupNumber interface{}

    // Authentication string. The type is string with length: 1..8. The default
    // value is cisco.
    Authentication interface{}

    // HSRP Session name (for MGO). The type is string with length: 1..16.
    SessionName interface{}

    // Priority value. The type is interface{} with range: 0..255. The default
    // value is 100.
    Priority interface{}

    // Force active if higher priority. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 0.
    Preempt interface{}

    // HSRP MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    VirtualMacAddress interface{}

    // The HSRP tracked interface configuration table.
    TrackedInterfaces Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces

    // Enable use of Bidirectional Forwarding Detection.
    Bfd Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd

    // The HSRP tracked interface configuration table.
    TrackedObjects Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects

    // Hello and hold timers.
    Timers Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers

    // Primary HSRP IP address.
    PrimaryIpv4Address Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address

    // Secondary HSRP IP address Table.
    SecondaryIpv4Addresses Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses
}

func (group *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group) GetGoName(yname string) string {
    if yname == "group-number" { return "GroupNumber" }
    if yname == "authentication" { return "Authentication" }
    if yname == "session-name" { return "SessionName" }
    if yname == "priority" { return "Priority" }
    if yname == "preempt" { return "Preempt" }
    if yname == "virtual-mac-address" { return "VirtualMacAddress" }
    if yname == "tracked-interfaces" { return "TrackedInterfaces" }
    if yname == "bfd" { return "Bfd" }
    if yname == "tracked-objects" { return "TrackedObjects" }
    if yname == "timers" { return "Timers" }
    if yname == "primary-ipv4-address" { return "PrimaryIpv4Address" }
    if yname == "secondary-ipv4-addresses" { return "SecondaryIpv4Addresses" }
    return ""
}

func (group *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group) GetSegmentPath() string {
    return "group" + "[group-number='" + fmt.Sprintf("%v", group.GroupNumber) + "']"
}

func (group *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tracked-interfaces" {
        return &group.TrackedInterfaces
    }
    if childYangName == "bfd" {
        return &group.Bfd
    }
    if childYangName == "tracked-objects" {
        return &group.TrackedObjects
    }
    if childYangName == "timers" {
        return &group.Timers
    }
    if childYangName == "primary-ipv4-address" {
        return &group.PrimaryIpv4Address
    }
    if childYangName == "secondary-ipv4-addresses" {
        return &group.SecondaryIpv4Addresses
    }
    return nil
}

func (group *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tracked-interfaces"] = &group.TrackedInterfaces
    children["bfd"] = &group.Bfd
    children["tracked-objects"] = &group.TrackedObjects
    children["timers"] = &group.Timers
    children["primary-ipv4-address"] = &group.PrimaryIpv4Address
    children["secondary-ipv4-addresses"] = &group.SecondaryIpv4Addresses
    return children
}

func (group *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-number"] = group.GroupNumber
    leafs["authentication"] = group.Authentication
    leafs["session-name"] = group.SessionName
    leafs["priority"] = group.Priority
    leafs["preempt"] = group.Preempt
    leafs["virtual-mac-address"] = group.VirtualMacAddress
    return leafs
}

func (group *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group) GetYangName() string { return "group" }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group) GetParentYangName() string { return "groups" }

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces
// The HSRP tracked interface configuration
// table
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface being tracked. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface.
    TrackedInterface []Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces) GetFilter() yfilter.YFilter { return trackedInterfaces.YFilter }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces) SetFilter(yf yfilter.YFilter) { trackedInterfaces.YFilter = yf }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces) GetGoName(yname string) string {
    if yname == "tracked-interface" { return "TrackedInterface" }
    return ""
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces) GetSegmentPath() string {
    return "tracked-interfaces"
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tracked-interface" {
        for _, c := range trackedInterfaces.TrackedInterface {
            if trackedInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface{}
        trackedInterfaces.TrackedInterface = append(trackedInterfaces.TrackedInterface, child)
        return &trackedInterfaces.TrackedInterface[len(trackedInterfaces.TrackedInterface)-1]
    }
    return nil
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackedInterfaces.TrackedInterface {
        children[trackedInterfaces.TrackedInterface[i].GetSegmentPath()] = &trackedInterfaces.TrackedInterface[i]
    }
    return children
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces) GetYangName() string { return "tracked-interfaces" }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces) SetParent(parent types.Entity) { trackedInterfaces.parent = parent }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces) GetParent() types.Entity { return trackedInterfaces.parent }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface
// Interface being tracked
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface being tracked. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Priority decrement. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface) GetFilter() yfilter.YFilter { return trackedInterface.YFilter }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface) SetFilter(yf yfilter.YFilter) { trackedInterface.YFilter = yf }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface) GetSegmentPath() string {
    return "tracked-interface" + "[interface-name='" + fmt.Sprintf("%v", trackedInterface.InterfaceName) + "']"
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = trackedInterface.InterfaceName
    leafs["priority-decrement"] = trackedInterface.PriorityDecrement
    return leafs
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface) GetBundleName() string { return "cisco_ios_xr" }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface) GetYangName() string { return "tracked-interface" }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface) SetParent(parent types.Entity) { trackedInterface.parent = parent }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface) GetParent() types.Entity { return trackedInterface.parent }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface) GetParentYangName() string { return "tracked-interfaces" }

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd
// Enable use of Bidirectional Forwarding
// Detection
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable BFD for this remote IP. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Interface name to run BFD. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd) GetFilter() yfilter.YFilter { return bfd.YFilter }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd) SetFilter(yf yfilter.YFilter) { bfd.YFilter = yf }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd) GetSegmentPath() string {
    return "bfd"
}

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = bfd.Address
    leafs["interface-name"] = bfd.InterfaceName
    return leafs
}

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd) GetBundleName() string { return "cisco_ios_xr" }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd) GetYangName() string { return "bfd" }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd) SetParent(parent types.Entity) { bfd.parent = parent }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd) GetParent() types.Entity { return bfd.parent }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects
// The HSRP tracked interface configuration
// table
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object being tracked. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject.
    TrackedObject []Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects) GetFilter() yfilter.YFilter { return trackedObjects.YFilter }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects) SetFilter(yf yfilter.YFilter) { trackedObjects.YFilter = yf }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects) GetGoName(yname string) string {
    if yname == "tracked-object" { return "TrackedObject" }
    return ""
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects) GetSegmentPath() string {
    return "tracked-objects"
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tracked-object" {
        for _, c := range trackedObjects.TrackedObject {
            if trackedObjects.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject{}
        trackedObjects.TrackedObject = append(trackedObjects.TrackedObject, child)
        return &trackedObjects.TrackedObject[len(trackedObjects.TrackedObject)-1]
    }
    return nil
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackedObjects.TrackedObject {
        children[trackedObjects.TrackedObject[i].GetSegmentPath()] = &trackedObjects.TrackedObject[i]
    }
    return children
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects) GetBundleName() string { return "cisco_ios_xr" }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects) GetYangName() string { return "tracked-objects" }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects) SetParent(parent types.Entity) { trackedObjects.parent = parent }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects) GetParent() types.Entity { return trackedObjects.parent }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject
// Object being tracked
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface being tracked. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ObjectName interface{}

    // Priority decrement. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject) GetFilter() yfilter.YFilter { return trackedObject.YFilter }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject) SetFilter(yf yfilter.YFilter) { trackedObject.YFilter = yf }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject) GetSegmentPath() string {
    return "tracked-object" + "[object-name='" + fmt.Sprintf("%v", trackedObject.ObjectName) + "']"
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = trackedObject.ObjectName
    leafs["priority-decrement"] = trackedObject.PriorityDecrement
    return leafs
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject) GetBundleName() string { return "cisco_ios_xr" }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject) GetYangName() string { return "tracked-object" }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject) SetParent(parent types.Entity) { trackedObject.parent = parent }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject) GetParent() types.Entity { return trackedObject.parent }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject) GetParentYangName() string { return "tracked-objects" }

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers
// Hello and hold timers
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TRUE - Hello time configured in milliseconds, FALSE - Hello time configured
    // in seconds. The type is bool. The default value is false.
    HelloMsecFlag interface{}

    // Hello time in msecs. The type is interface{} with range: 100..3000. Units
    // are millisecond.
    HelloMsec interface{}

    // Hello time in seconds. The type is interface{} with range: 1..255. Units
    // are second. The default value is 3.
    HelloSec interface{}

    // TRUE - Hold time configured in milliseconds, FALSE - Hold time configured
    // in seconds. The type is bool. The default value is false.
    HoldMsecFlag interface{}

    // Hold time in msecs. The type is interface{} with range: 100..3000. Units
    // are millisecond.
    HoldMsec interface{}

    // Hold time in seconds. The type is interface{} with range: 1..255. Units are
    // second. The default value is 10.
    HoldSec interface{}
}

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers) GetFilter() yfilter.YFilter { return timers.YFilter }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers) SetFilter(yf yfilter.YFilter) { timers.YFilter = yf }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers) GetGoName(yname string) string {
    if yname == "hello-msec-flag" { return "HelloMsecFlag" }
    if yname == "hello-msec" { return "HelloMsec" }
    if yname == "hello-sec" { return "HelloSec" }
    if yname == "hold-msec-flag" { return "HoldMsecFlag" }
    if yname == "hold-msec" { return "HoldMsec" }
    if yname == "hold-sec" { return "HoldSec" }
    return ""
}

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers) GetSegmentPath() string {
    return "timers"
}

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hello-msec-flag"] = timers.HelloMsecFlag
    leafs["hello-msec"] = timers.HelloMsec
    leafs["hello-sec"] = timers.HelloSec
    leafs["hold-msec-flag"] = timers.HoldMsecFlag
    leafs["hold-msec"] = timers.HoldMsec
    leafs["hold-sec"] = timers.HoldSec
    return leafs
}

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers) GetBundleName() string { return "cisco_ios_xr" }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers) GetYangName() string { return "timers" }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers) SetParent(parent types.Entity) { timers.parent = parent }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers) GetParent() types.Entity { return timers.parent }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address
// Primary HSRP IP address
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TRUE if the HSRP protocol is to learn the virtual IP address it is to use.
    // The type is bool.
    VirtualIpLearn interface{}

    // HSRP IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address) GetFilter() yfilter.YFilter { return primaryIpv4Address.YFilter }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address) SetFilter(yf yfilter.YFilter) { primaryIpv4Address.YFilter = yf }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address) GetGoName(yname string) string {
    if yname == "virtual-ip-learn" { return "VirtualIpLearn" }
    if yname == "address" { return "Address" }
    return ""
}

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address) GetSegmentPath() string {
    return "primary-ipv4-address"
}

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["virtual-ip-learn"] = primaryIpv4Address.VirtualIpLearn
    leafs["address"] = primaryIpv4Address.Address
    return leafs
}

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address) GetBundleName() string { return "cisco_ios_xr" }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address) GetYangName() string { return "primary-ipv4-address" }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address) SetParent(parent types.Entity) { primaryIpv4Address.parent = parent }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address) GetParent() types.Entity { return primaryIpv4Address.parent }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses
// Secondary HSRP IP address Table
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Secondary HSRP IP address. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address.
    SecondaryIpv4Address []Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses) GetFilter() yfilter.YFilter { return secondaryIpv4Addresses.YFilter }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses) SetFilter(yf yfilter.YFilter) { secondaryIpv4Addresses.YFilter = yf }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses) GetGoName(yname string) string {
    if yname == "secondary-ipv4-address" { return "SecondaryIpv4Address" }
    return ""
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses) GetSegmentPath() string {
    return "secondary-ipv4-addresses"
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "secondary-ipv4-address" {
        for _, c := range secondaryIpv4Addresses.SecondaryIpv4Address {
            if secondaryIpv4Addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address{}
        secondaryIpv4Addresses.SecondaryIpv4Address = append(secondaryIpv4Addresses.SecondaryIpv4Address, child)
        return &secondaryIpv4Addresses.SecondaryIpv4Address[len(secondaryIpv4Addresses.SecondaryIpv4Address)-1]
    }
    return nil
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range secondaryIpv4Addresses.SecondaryIpv4Address {
        children[secondaryIpv4Addresses.SecondaryIpv4Address[i].GetSegmentPath()] = &secondaryIpv4Addresses.SecondaryIpv4Address[i]
    }
    return children
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses) GetYangName() string { return "secondary-ipv4-addresses" }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses) SetParent(parent types.Entity) { secondaryIpv4Addresses.parent = parent }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses) GetParent() types.Entity { return secondaryIpv4Addresses.parent }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address
// Secondary HSRP IP address
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetFilter() yfilter.YFilter { return secondaryIpv4Address.YFilter }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) SetFilter(yf yfilter.YFilter) { secondaryIpv4Address.YFilter = yf }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetSegmentPath() string {
    return "secondary-ipv4-address" + "[address='" + fmt.Sprintf("%v", secondaryIpv4Address.Address) + "']"
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = secondaryIpv4Address.Address
    return leafs
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetYangName() string { return "secondary-ipv4-address" }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) SetParent(parent types.Entity) { secondaryIpv4Address.parent = parent }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetParent() types.Entity { return secondaryIpv4Address.parent }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetParentYangName() string { return "secondary-ipv4-addresses" }

// Hsrp_Interfaces_Interface_Ipv4_Version2
// Version 2 HSRP configuration
type Hsrp_Interfaces_Interface_Ipv4_Version2 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The HSRP group configuration table.
    Groups Hsrp_Interfaces_Interface_Ipv4_Version2_Groups
}

func (version2 *Hsrp_Interfaces_Interface_Ipv4_Version2) GetFilter() yfilter.YFilter { return version2.YFilter }

func (version2 *Hsrp_Interfaces_Interface_Ipv4_Version2) SetFilter(yf yfilter.YFilter) { version2.YFilter = yf }

func (version2 *Hsrp_Interfaces_Interface_Ipv4_Version2) GetGoName(yname string) string {
    if yname == "groups" { return "Groups" }
    return ""
}

func (version2 *Hsrp_Interfaces_Interface_Ipv4_Version2) GetSegmentPath() string {
    return "version2"
}

func (version2 *Hsrp_Interfaces_Interface_Ipv4_Version2) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "groups" {
        return &version2.Groups
    }
    return nil
}

func (version2 *Hsrp_Interfaces_Interface_Ipv4_Version2) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["groups"] = &version2.Groups
    return children
}

func (version2 *Hsrp_Interfaces_Interface_Ipv4_Version2) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (version2 *Hsrp_Interfaces_Interface_Ipv4_Version2) GetBundleName() string { return "cisco_ios_xr" }

func (version2 *Hsrp_Interfaces_Interface_Ipv4_Version2) GetYangName() string { return "version2" }

func (version2 *Hsrp_Interfaces_Interface_Ipv4_Version2) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (version2 *Hsrp_Interfaces_Interface_Ipv4_Version2) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (version2 *Hsrp_Interfaces_Interface_Ipv4_Version2) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (version2 *Hsrp_Interfaces_Interface_Ipv4_Version2) SetParent(parent types.Entity) { version2.parent = parent }

func (version2 *Hsrp_Interfaces_Interface_Ipv4_Version2) GetParent() types.Entity { return version2.parent }

func (version2 *Hsrp_Interfaces_Interface_Ipv4_Version2) GetParentYangName() string { return "ipv4" }

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups
// The HSRP group configuration table
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The HSRP group being configured. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group.
    Group []Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group
}

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups) GetFilter() yfilter.YFilter { return groups.YFilter }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups) SetFilter(yf yfilter.YFilter) { groups.YFilter = yf }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups) GetGoName(yname string) string {
    if yname == "group" { return "Group" }
    return ""
}

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups) GetSegmentPath() string {
    return "groups"
}

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "group" {
        for _, c := range groups.Group {
            if groups.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group{}
        groups.Group = append(groups.Group, child)
        return &groups.Group[len(groups.Group)-1]
    }
    return nil
}

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range groups.Group {
        children[groups.Group[i].GetSegmentPath()] = &groups.Group[i]
    }
    return children
}

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups) GetBundleName() string { return "cisco_ios_xr" }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups) GetYangName() string { return "groups" }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups) SetParent(parent types.Entity) { groups.parent = parent }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups) GetParent() types.Entity { return groups.parent }

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups) GetParentYangName() string { return "version2" }

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group
// The HSRP group being configured
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP group number. The type is interface{} with
    // range: 0..4095.
    GroupNumber interface{}

    // Force active if higher priority. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 0.
    Preempt interface{}

    // Priority value. The type is interface{} with range: 0..255. The default
    // value is 100.
    Priority interface{}

    // HSRP MAC address. The type is string with pattern:
    // [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    VirtualMacAddress interface{}

    // HSRP Session name (for MGO). The type is string with length: 1..16.
    SessionName interface{}

    // Secondary HSRP IP address Table.
    SecondaryIpv4Addresses Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses

    // Enable use of Bidirectional Forwarding Detection.
    Bfd Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd

    // Primary HSRP IP address.
    PrimaryIpv4Address Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address

    // The HSRP tracked interface configuration table.
    TrackedObjects Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects

    // The HSRP tracked interface configuration table.
    TrackedInterfaces Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces

    // Hello and hold timers.
    Timers Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers
}

func (group *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group) GetFilter() yfilter.YFilter { return group.YFilter }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group) SetFilter(yf yfilter.YFilter) { group.YFilter = yf }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group) GetGoName(yname string) string {
    if yname == "group-number" { return "GroupNumber" }
    if yname == "preempt" { return "Preempt" }
    if yname == "priority" { return "Priority" }
    if yname == "virtual-mac-address" { return "VirtualMacAddress" }
    if yname == "session-name" { return "SessionName" }
    if yname == "secondary-ipv4-addresses" { return "SecondaryIpv4Addresses" }
    if yname == "bfd" { return "Bfd" }
    if yname == "primary-ipv4-address" { return "PrimaryIpv4Address" }
    if yname == "tracked-objects" { return "TrackedObjects" }
    if yname == "tracked-interfaces" { return "TrackedInterfaces" }
    if yname == "timers" { return "Timers" }
    return ""
}

func (group *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group) GetSegmentPath() string {
    return "group" + "[group-number='" + fmt.Sprintf("%v", group.GroupNumber) + "']"
}

func (group *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "secondary-ipv4-addresses" {
        return &group.SecondaryIpv4Addresses
    }
    if childYangName == "bfd" {
        return &group.Bfd
    }
    if childYangName == "primary-ipv4-address" {
        return &group.PrimaryIpv4Address
    }
    if childYangName == "tracked-objects" {
        return &group.TrackedObjects
    }
    if childYangName == "tracked-interfaces" {
        return &group.TrackedInterfaces
    }
    if childYangName == "timers" {
        return &group.Timers
    }
    return nil
}

func (group *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["secondary-ipv4-addresses"] = &group.SecondaryIpv4Addresses
    children["bfd"] = &group.Bfd
    children["primary-ipv4-address"] = &group.PrimaryIpv4Address
    children["tracked-objects"] = &group.TrackedObjects
    children["tracked-interfaces"] = &group.TrackedInterfaces
    children["timers"] = &group.Timers
    return children
}

func (group *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["group-number"] = group.GroupNumber
    leafs["preempt"] = group.Preempt
    leafs["priority"] = group.Priority
    leafs["virtual-mac-address"] = group.VirtualMacAddress
    leafs["session-name"] = group.SessionName
    return leafs
}

func (group *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group) GetBundleName() string { return "cisco_ios_xr" }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group) GetYangName() string { return "group" }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group) SetParent(parent types.Entity) { group.parent = parent }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group) GetParent() types.Entity { return group.parent }

func (group *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group) GetParentYangName() string { return "groups" }

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses
// Secondary HSRP IP address Table
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Secondary HSRP IP address. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address.
    SecondaryIpv4Address []Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses) GetFilter() yfilter.YFilter { return secondaryIpv4Addresses.YFilter }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses) SetFilter(yf yfilter.YFilter) { secondaryIpv4Addresses.YFilter = yf }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses) GetGoName(yname string) string {
    if yname == "secondary-ipv4-address" { return "SecondaryIpv4Address" }
    return ""
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses) GetSegmentPath() string {
    return "secondary-ipv4-addresses"
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "secondary-ipv4-address" {
        for _, c := range secondaryIpv4Addresses.SecondaryIpv4Address {
            if secondaryIpv4Addresses.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address{}
        secondaryIpv4Addresses.SecondaryIpv4Address = append(secondaryIpv4Addresses.SecondaryIpv4Address, child)
        return &secondaryIpv4Addresses.SecondaryIpv4Address[len(secondaryIpv4Addresses.SecondaryIpv4Address)-1]
    }
    return nil
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range secondaryIpv4Addresses.SecondaryIpv4Address {
        children[secondaryIpv4Addresses.SecondaryIpv4Address[i].GetSegmentPath()] = &secondaryIpv4Addresses.SecondaryIpv4Address[i]
    }
    return children
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses) GetYangName() string { return "secondary-ipv4-addresses" }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses) SetParent(parent types.Entity) { secondaryIpv4Addresses.parent = parent }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses) GetParent() types.Entity { return secondaryIpv4Addresses.parent }

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address
// Secondary HSRP IP address
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetFilter() yfilter.YFilter { return secondaryIpv4Address.YFilter }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) SetFilter(yf yfilter.YFilter) { secondaryIpv4Address.YFilter = yf }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    return ""
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetSegmentPath() string {
    return "secondary-ipv4-address" + "[address='" + fmt.Sprintf("%v", secondaryIpv4Address.Address) + "']"
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = secondaryIpv4Address.Address
    return leafs
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetBundleName() string { return "cisco_ios_xr" }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetYangName() string { return "secondary-ipv4-address" }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) SetParent(parent types.Entity) { secondaryIpv4Address.parent = parent }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetParent() types.Entity { return secondaryIpv4Address.parent }

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetParentYangName() string { return "secondary-ipv4-addresses" }

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd
// Enable use of Bidirectional Forwarding
// Detection
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable BFD for this remote IP. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Interface name to run BFD. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}
}

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd) GetFilter() yfilter.YFilter { return bfd.YFilter }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd) SetFilter(yf yfilter.YFilter) { bfd.YFilter = yf }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd) GetSegmentPath() string {
    return "bfd"
}

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = bfd.Address
    leafs["interface-name"] = bfd.InterfaceName
    return leafs
}

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd) GetBundleName() string { return "cisco_ios_xr" }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd) GetYangName() string { return "bfd" }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd) SetParent(parent types.Entity) { bfd.parent = parent }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd) GetParent() types.Entity { return bfd.parent }

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address
// Primary HSRP IP address
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TRUE if the HSRP protocol is to learn the virtual IP address it is to use.
    // The type is bool.
    VirtualIpLearn interface{}

    // HSRP IP address. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Address interface{}
}

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address) GetFilter() yfilter.YFilter { return primaryIpv4Address.YFilter }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address) SetFilter(yf yfilter.YFilter) { primaryIpv4Address.YFilter = yf }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address) GetGoName(yname string) string {
    if yname == "virtual-ip-learn" { return "VirtualIpLearn" }
    if yname == "address" { return "Address" }
    return ""
}

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address) GetSegmentPath() string {
    return "primary-ipv4-address"
}

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["virtual-ip-learn"] = primaryIpv4Address.VirtualIpLearn
    leafs["address"] = primaryIpv4Address.Address
    return leafs
}

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address) GetBundleName() string { return "cisco_ios_xr" }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address) GetYangName() string { return "primary-ipv4-address" }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address) SetParent(parent types.Entity) { primaryIpv4Address.parent = parent }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address) GetParent() types.Entity { return primaryIpv4Address.parent }

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects
// The HSRP tracked interface configuration
// table
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Object being tracked. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject.
    TrackedObject []Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects) GetFilter() yfilter.YFilter { return trackedObjects.YFilter }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects) SetFilter(yf yfilter.YFilter) { trackedObjects.YFilter = yf }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects) GetGoName(yname string) string {
    if yname == "tracked-object" { return "TrackedObject" }
    return ""
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects) GetSegmentPath() string {
    return "tracked-objects"
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tracked-object" {
        for _, c := range trackedObjects.TrackedObject {
            if trackedObjects.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject{}
        trackedObjects.TrackedObject = append(trackedObjects.TrackedObject, child)
        return &trackedObjects.TrackedObject[len(trackedObjects.TrackedObject)-1]
    }
    return nil
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackedObjects.TrackedObject {
        children[trackedObjects.TrackedObject[i].GetSegmentPath()] = &trackedObjects.TrackedObject[i]
    }
    return children
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects) GetBundleName() string { return "cisco_ios_xr" }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects) GetYangName() string { return "tracked-objects" }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects) SetParent(parent types.Entity) { trackedObjects.parent = parent }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects) GetParent() types.Entity { return trackedObjects.parent }

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject
// Object being tracked
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface being tracked. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    ObjectName interface{}

    // Priority decrement. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject) GetFilter() yfilter.YFilter { return trackedObject.YFilter }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject) SetFilter(yf yfilter.YFilter) { trackedObject.YFilter = yf }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject) GetGoName(yname string) string {
    if yname == "object-name" { return "ObjectName" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject) GetSegmentPath() string {
    return "tracked-object" + "[object-name='" + fmt.Sprintf("%v", trackedObject.ObjectName) + "']"
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["object-name"] = trackedObject.ObjectName
    leafs["priority-decrement"] = trackedObject.PriorityDecrement
    return leafs
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject) GetBundleName() string { return "cisco_ios_xr" }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject) GetYangName() string { return "tracked-object" }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject) SetParent(parent types.Entity) { trackedObject.parent = parent }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject) GetParent() types.Entity { return trackedObject.parent }

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject) GetParentYangName() string { return "tracked-objects" }

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces
// The HSRP tracked interface configuration
// table
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interface being tracked. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface.
    TrackedInterface []Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces) GetFilter() yfilter.YFilter { return trackedInterfaces.YFilter }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces) SetFilter(yf yfilter.YFilter) { trackedInterfaces.YFilter = yf }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces) GetGoName(yname string) string {
    if yname == "tracked-interface" { return "TrackedInterface" }
    return ""
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces) GetSegmentPath() string {
    return "tracked-interfaces"
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tracked-interface" {
        for _, c := range trackedInterfaces.TrackedInterface {
            if trackedInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface{}
        trackedInterfaces.TrackedInterface = append(trackedInterfaces.TrackedInterface, child)
        return &trackedInterfaces.TrackedInterface[len(trackedInterfaces.TrackedInterface)-1]
    }
    return nil
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range trackedInterfaces.TrackedInterface {
        children[trackedInterfaces.TrackedInterface[i].GetSegmentPath()] = &trackedInterfaces.TrackedInterface[i]
    }
    return children
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces) GetBundleName() string { return "cisco_ios_xr" }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces) GetYangName() string { return "tracked-interfaces" }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces) SetParent(parent types.Entity) { trackedInterfaces.parent = parent }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces) GetParent() types.Entity { return trackedInterfaces.parent }

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces) GetParentYangName() string { return "group" }

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface
// Interface being tracked
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface being tracked. The type is string with
    // pattern: [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Priority decrement. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetFilter() yfilter.YFilter { return trackedInterface.YFilter }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) SetFilter(yf yfilter.YFilter) { trackedInterface.YFilter = yf }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "priority-decrement" { return "PriorityDecrement" }
    return ""
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetSegmentPath() string {
    return "tracked-interface" + "[interface-name='" + fmt.Sprintf("%v", trackedInterface.InterfaceName) + "']"
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = trackedInterface.InterfaceName
    leafs["priority-decrement"] = trackedInterface.PriorityDecrement
    return leafs
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetBundleName() string { return "cisco_ios_xr" }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetYangName() string { return "tracked-interface" }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) SetParent(parent types.Entity) { trackedInterface.parent = parent }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetParent() types.Entity { return trackedInterface.parent }

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetParentYangName() string { return "tracked-interfaces" }

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers
// Hello and hold timers
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TRUE - Hello time configured in milliseconds, FALSE - Hello time configured
    // in seconds. The type is bool. The default value is false.
    HelloMsecFlag interface{}

    // Hello time in msecs. The type is interface{} with range: 100..3000. Units
    // are millisecond.
    HelloMsec interface{}

    // Hello time in seconds. The type is interface{} with range: 1..255. Units
    // are second. The default value is 3.
    HelloSec interface{}

    // TRUE - Hold time configured in milliseconds, FALSE - Hold time configured
    // in seconds. The type is bool. The default value is false.
    HoldMsecFlag interface{}

    // Hold time in msecs. The type is interface{} with range: 100..3000. Units
    // are millisecond.
    HoldMsec interface{}

    // Hold time in seconds. The type is interface{} with range: 1..255. Units are
    // second. The default value is 10.
    HoldSec interface{}
}

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers) GetFilter() yfilter.YFilter { return timers.YFilter }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers) SetFilter(yf yfilter.YFilter) { timers.YFilter = yf }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers) GetGoName(yname string) string {
    if yname == "hello-msec-flag" { return "HelloMsecFlag" }
    if yname == "hello-msec" { return "HelloMsec" }
    if yname == "hello-sec" { return "HelloSec" }
    if yname == "hold-msec-flag" { return "HoldMsecFlag" }
    if yname == "hold-msec" { return "HoldMsec" }
    if yname == "hold-sec" { return "HoldSec" }
    return ""
}

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers) GetSegmentPath() string {
    return "timers"
}

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hello-msec-flag"] = timers.HelloMsecFlag
    leafs["hello-msec"] = timers.HelloMsec
    leafs["hello-sec"] = timers.HelloSec
    leafs["hold-msec-flag"] = timers.HoldMsecFlag
    leafs["hold-msec"] = timers.HoldMsec
    leafs["hold-sec"] = timers.HoldSec
    return leafs
}

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers) GetBundleName() string { return "cisco_ios_xr" }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers) GetYangName() string { return "timers" }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers) SetParent(parent types.Entity) { timers.parent = parent }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers) GetParent() types.Entity { return timers.parent }

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers) GetParentYangName() string { return "group" }

// Hsrp_Logging
// HSRP logging options
type Hsrp_Logging struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // HSRP state change IOS messages disable. The type is interface{}.
    StateChangeDisable interface{}
}

func (logging *Hsrp_Logging) GetFilter() yfilter.YFilter { return logging.YFilter }

func (logging *Hsrp_Logging) SetFilter(yf yfilter.YFilter) { logging.YFilter = yf }

func (logging *Hsrp_Logging) GetGoName(yname string) string {
    if yname == "state-change-disable" { return "StateChangeDisable" }
    return ""
}

func (logging *Hsrp_Logging) GetSegmentPath() string {
    return "logging"
}

func (logging *Hsrp_Logging) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (logging *Hsrp_Logging) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (logging *Hsrp_Logging) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["state-change-disable"] = logging.StateChangeDisable
    return leafs
}

func (logging *Hsrp_Logging) GetBundleName() string { return "cisco_ios_xr" }

func (logging *Hsrp_Logging) GetYangName() string { return "logging" }

func (logging *Hsrp_Logging) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (logging *Hsrp_Logging) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (logging *Hsrp_Logging) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (logging *Hsrp_Logging) SetParent(parent types.Entity) { logging.parent = parent }

func (logging *Hsrp_Logging) GetParent() types.Entity { return logging.parent }

func (logging *Hsrp_Logging) GetParentYangName() string { return "hsrp" }

