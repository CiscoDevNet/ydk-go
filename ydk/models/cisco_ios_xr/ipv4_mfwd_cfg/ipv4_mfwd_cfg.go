// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-mfwd package configuration.
// 
// This module contains definitions
// for the following management objects:
//   mfwd: Multicast routing configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_mfwd_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_mfwd_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-mfwd-cfg mfwd}", reflect.TypeOf(Mfwd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-mfwd-cfg:mfwd", reflect.TypeOf(Mfwd{}))
}

// AccountingMode represents Accounting mode
type AccountingMode string

const (
    // Enable per (S,G) accounting
    AccountingMode_enable AccountingMode = "enable"

    // Enable per (S,G) forward-only accounting
    AccountingMode_forward_only_enable AccountingMode = "forward-only-enable"
)

// Mfwd
// Multicast routing configuration
// This type is a presence type.
type Mfwd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Default Context.
    DefaultContext Mfwd_DefaultContext

    // VRF Table.
    Vrfs Mfwd_Vrfs
}

func (mfwd *Mfwd) GetFilter() yfilter.YFilter { return mfwd.YFilter }

func (mfwd *Mfwd) SetFilter(yf yfilter.YFilter) { mfwd.YFilter = yf }

func (mfwd *Mfwd) GetGoName(yname string) string {
    if yname == "default-context" { return "DefaultContext" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (mfwd *Mfwd) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-mfwd-cfg:mfwd"
}

func (mfwd *Mfwd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "default-context" {
        return &mfwd.DefaultContext
    }
    if childYangName == "vrfs" {
        return &mfwd.Vrfs
    }
    return nil
}

func (mfwd *Mfwd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["default-context"] = &mfwd.DefaultContext
    children["vrfs"] = &mfwd.Vrfs
    return children
}

func (mfwd *Mfwd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (mfwd *Mfwd) GetBundleName() string { return "cisco_ios_xr" }

func (mfwd *Mfwd) GetYangName() string { return "mfwd" }

func (mfwd *Mfwd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mfwd *Mfwd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mfwd *Mfwd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mfwd *Mfwd) SetParent(parent types.Entity) { mfwd.parent = parent }

func (mfwd *Mfwd) GetParent() types.Entity { return mfwd.parent }

func (mfwd *Mfwd) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-mfwd-cfg" }

// Mfwd_DefaultContext
// Default Context
// This type is a presence type.
type Mfwd_DefaultContext struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // IPV6 commands in the default context.
    Ipv6 Mfwd_DefaultContext_Ipv6

    // IPV4 commands in the default context.
    Ipv4 Mfwd_DefaultContext_Ipv4
}

func (defaultContext *Mfwd_DefaultContext) GetFilter() yfilter.YFilter { return defaultContext.YFilter }

func (defaultContext *Mfwd_DefaultContext) SetFilter(yf yfilter.YFilter) { defaultContext.YFilter = yf }

func (defaultContext *Mfwd_DefaultContext) GetGoName(yname string) string {
    if yname == "ipv6" { return "Ipv6" }
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (defaultContext *Mfwd_DefaultContext) GetSegmentPath() string {
    return "default-context"
}

func (defaultContext *Mfwd_DefaultContext) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6" {
        return &defaultContext.Ipv6
    }
    if childYangName == "ipv4" {
        return &defaultContext.Ipv4
    }
    return nil
}

func (defaultContext *Mfwd_DefaultContext) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6"] = &defaultContext.Ipv6
    children["ipv4"] = &defaultContext.Ipv4
    return children
}

func (defaultContext *Mfwd_DefaultContext) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (defaultContext *Mfwd_DefaultContext) GetBundleName() string { return "cisco_ios_xr" }

func (defaultContext *Mfwd_DefaultContext) GetYangName() string { return "default-context" }

func (defaultContext *Mfwd_DefaultContext) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (defaultContext *Mfwd_DefaultContext) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (defaultContext *Mfwd_DefaultContext) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (defaultContext *Mfwd_DefaultContext) SetParent(parent types.Entity) { defaultContext.parent = parent }

func (defaultContext *Mfwd_DefaultContext) GetParent() types.Entity { return defaultContext.parent }

func (defaultContext *Mfwd_DefaultContext) GetParentYangName() string { return "mfwd" }

// Mfwd_DefaultContext_Ipv6
// IPV6 commands in the default context
type Mfwd_DefaultContext_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure all interfaces for multicast routing and forwarding. The type is
    // interface{}.
    EnableOnAllInterfaces interface{}

    // Disable state limit maximum checking. The type is interface{}.
    MaximumCheckingDisable interface{}

    // Enable per (S,G) rate calculation. The type is interface{}.
    RatePerRoute interface{}

    // Disable interface inheritance configuration. The type is interface{}.
    InterfaceInheritanceDisable interface{}

    // Mofrr Lockout timer value. The type is interface{} with range: 1..3600.
    MofrrLockoutTimerConfig interface{}

    // Knob to delay traffic being forwarded on a route. The type is interface{}
    // with range: 5..500. Units are millisecond.
    ForwardingLatency interface{}

    // Mofrr Loss Detection timer value. The type is interface{} with range:
    // 1..3600.
    MofrrLossDetectionTimerConfig interface{}

    // Enable IP multicast routing and forwarding. The type is interface{}.
    MulticastForwarding interface{}

    // Enable logging trap event. The type is interface{}.
    LogTraps interface{}

    // Per-prefix accounting mode. The type is AccountingMode.
    Accounting interface{}

    // Configure a static RPF rule for a given prefix.
    StaticRpfRules Mfwd_DefaultContext_Ipv6_StaticRpfRules

    // Interface-level Configuration.
    Interfaces Mfwd_DefaultContext_Ipv6_Interfaces
}

func (ipv6 *Mfwd_DefaultContext_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Mfwd_DefaultContext_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Mfwd_DefaultContext_Ipv6) GetGoName(yname string) string {
    if yname == "enable-on-all-interfaces" { return "EnableOnAllInterfaces" }
    if yname == "maximum-checking-disable" { return "MaximumCheckingDisable" }
    if yname == "rate-per-route" { return "RatePerRoute" }
    if yname == "interface-inheritance-disable" { return "InterfaceInheritanceDisable" }
    if yname == "mofrr-lockout-timer-config" { return "MofrrLockoutTimerConfig" }
    if yname == "forwarding-latency" { return "ForwardingLatency" }
    if yname == "mofrr-loss-detection-timer-config" { return "MofrrLossDetectionTimerConfig" }
    if yname == "multicast-forwarding" { return "MulticastForwarding" }
    if yname == "log-traps" { return "LogTraps" }
    if yname == "accounting" { return "Accounting" }
    if yname == "static-rpf-rules" { return "StaticRpfRules" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (ipv6 *Mfwd_DefaultContext_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Mfwd_DefaultContext_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static-rpf-rules" {
        return &ipv6.StaticRpfRules
    }
    if childYangName == "interfaces" {
        return &ipv6.Interfaces
    }
    return nil
}

func (ipv6 *Mfwd_DefaultContext_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["static-rpf-rules"] = &ipv6.StaticRpfRules
    children["interfaces"] = &ipv6.Interfaces
    return children
}

func (ipv6 *Mfwd_DefaultContext_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable-on-all-interfaces"] = ipv6.EnableOnAllInterfaces
    leafs["maximum-checking-disable"] = ipv6.MaximumCheckingDisable
    leafs["rate-per-route"] = ipv6.RatePerRoute
    leafs["interface-inheritance-disable"] = ipv6.InterfaceInheritanceDisable
    leafs["mofrr-lockout-timer-config"] = ipv6.MofrrLockoutTimerConfig
    leafs["forwarding-latency"] = ipv6.ForwardingLatency
    leafs["mofrr-loss-detection-timer-config"] = ipv6.MofrrLossDetectionTimerConfig
    leafs["multicast-forwarding"] = ipv6.MulticastForwarding
    leafs["log-traps"] = ipv6.LogTraps
    leafs["accounting"] = ipv6.Accounting
    return leafs
}

func (ipv6 *Mfwd_DefaultContext_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Mfwd_DefaultContext_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Mfwd_DefaultContext_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Mfwd_DefaultContext_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Mfwd_DefaultContext_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Mfwd_DefaultContext_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Mfwd_DefaultContext_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Mfwd_DefaultContext_Ipv6) GetParentYangName() string { return "default-context" }

// Mfwd_DefaultContext_Ipv6_StaticRpfRules
// Configure a static RPF rule for a given prefix
type Mfwd_DefaultContext_Ipv6_StaticRpfRules struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RPF prefix address and mask. The type is slice of
    // Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule.
    StaticRpfRule []Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule
}

func (staticRpfRules *Mfwd_DefaultContext_Ipv6_StaticRpfRules) GetFilter() yfilter.YFilter { return staticRpfRules.YFilter }

func (staticRpfRules *Mfwd_DefaultContext_Ipv6_StaticRpfRules) SetFilter(yf yfilter.YFilter) { staticRpfRules.YFilter = yf }

func (staticRpfRules *Mfwd_DefaultContext_Ipv6_StaticRpfRules) GetGoName(yname string) string {
    if yname == "static-rpf-rule" { return "StaticRpfRule" }
    return ""
}

func (staticRpfRules *Mfwd_DefaultContext_Ipv6_StaticRpfRules) GetSegmentPath() string {
    return "static-rpf-rules"
}

func (staticRpfRules *Mfwd_DefaultContext_Ipv6_StaticRpfRules) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static-rpf-rule" {
        for _, c := range staticRpfRules.StaticRpfRule {
            if staticRpfRules.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule{}
        staticRpfRules.StaticRpfRule = append(staticRpfRules.StaticRpfRule, child)
        return &staticRpfRules.StaticRpfRule[len(staticRpfRules.StaticRpfRule)-1]
    }
    return nil
}

func (staticRpfRules *Mfwd_DefaultContext_Ipv6_StaticRpfRules) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range staticRpfRules.StaticRpfRule {
        children[staticRpfRules.StaticRpfRule[i].GetSegmentPath()] = &staticRpfRules.StaticRpfRule[i]
    }
    return children
}

func (staticRpfRules *Mfwd_DefaultContext_Ipv6_StaticRpfRules) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (staticRpfRules *Mfwd_DefaultContext_Ipv6_StaticRpfRules) GetBundleName() string { return "cisco_ios_xr" }

func (staticRpfRules *Mfwd_DefaultContext_Ipv6_StaticRpfRules) GetYangName() string { return "static-rpf-rules" }

func (staticRpfRules *Mfwd_DefaultContext_Ipv6_StaticRpfRules) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticRpfRules *Mfwd_DefaultContext_Ipv6_StaticRpfRules) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticRpfRules *Mfwd_DefaultContext_Ipv6_StaticRpfRules) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticRpfRules *Mfwd_DefaultContext_Ipv6_StaticRpfRules) SetParent(parent types.Entity) { staticRpfRules.parent = parent }

func (staticRpfRules *Mfwd_DefaultContext_Ipv6_StaticRpfRules) GetParent() types.Entity { return staticRpfRules.parent }

func (staticRpfRules *Mfwd_DefaultContext_Ipv6_StaticRpfRules) GetParentYangName() string { return "ipv6" }

// Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule
// RPF prefix address and mask
type Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the RPF prefix. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. Prefix mask of the RPF Prefix. The type is
    // interface{} with range: 0..128.
    PrefixMask interface{}

    // Neighbor address of the RPF Prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    NeighborAddress interface{}

    // The name of the RPF interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+. This attribute is mandatory.
    InterfaceName interface{}
}

func (staticRpfRule *Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule) GetFilter() yfilter.YFilter { return staticRpfRule.YFilter }

func (staticRpfRule *Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule) SetFilter(yf yfilter.YFilter) { staticRpfRule.YFilter = yf }

func (staticRpfRule *Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-mask" { return "PrefixMask" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (staticRpfRule *Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule) GetSegmentPath() string {
    return "static-rpf-rule" + "[address='" + fmt.Sprintf("%v", staticRpfRule.Address) + "']" + "[prefix-mask='" + fmt.Sprintf("%v", staticRpfRule.PrefixMask) + "']"
}

func (staticRpfRule *Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticRpfRule *Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticRpfRule *Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = staticRpfRule.Address
    leafs["prefix-mask"] = staticRpfRule.PrefixMask
    leafs["neighbor-address"] = staticRpfRule.NeighborAddress
    leafs["interface-name"] = staticRpfRule.InterfaceName
    return leafs
}

func (staticRpfRule *Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule) GetBundleName() string { return "cisco_ios_xr" }

func (staticRpfRule *Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule) GetYangName() string { return "static-rpf-rule" }

func (staticRpfRule *Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticRpfRule *Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticRpfRule *Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticRpfRule *Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule) SetParent(parent types.Entity) { staticRpfRule.parent = parent }

func (staticRpfRule *Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule) GetParent() types.Entity { return staticRpfRule.parent }

func (staticRpfRule *Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule) GetParentYangName() string { return "static-rpf-rules" }

// Mfwd_DefaultContext_Ipv6_Interfaces
// Interface-level Configuration
type Mfwd_DefaultContext_Ipv6_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Mfwd_DefaultContext_Ipv6_Interfaces_Interface.
    Interface []Mfwd_DefaultContext_Ipv6_Interfaces_Interface
}

func (interfaces *Mfwd_DefaultContext_Ipv6_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Mfwd_DefaultContext_Ipv6_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Mfwd_DefaultContext_Ipv6_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Mfwd_DefaultContext_Ipv6_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Mfwd_DefaultContext_Ipv6_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mfwd_DefaultContext_Ipv6_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Mfwd_DefaultContext_Ipv6_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Mfwd_DefaultContext_Ipv6_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Mfwd_DefaultContext_Ipv6_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Mfwd_DefaultContext_Ipv6_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Mfwd_DefaultContext_Ipv6_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Mfwd_DefaultContext_Ipv6_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Mfwd_DefaultContext_Ipv6_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Mfwd_DefaultContext_Ipv6_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Mfwd_DefaultContext_Ipv6_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Mfwd_DefaultContext_Ipv6_Interfaces) GetParentYangName() string { return "ipv6" }

// Mfwd_DefaultContext_Ipv6_Interfaces_Interface
// The name of the interface
type Mfwd_DefaultContext_Ipv6_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // TTL threshold for multicast packets. The type is interface{} with range:
    // 1..255.
    TtlThreshold interface{}

    // Enable or disable IP multicast on the interface. The type is bool.
    EnableOnInterface interface{}

    // Boundary for administratively scoped multicast addresses. The type is
    // string with length: 1..64.
    Boundary interface{}
}

func (self *Mfwd_DefaultContext_Ipv6_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Mfwd_DefaultContext_Ipv6_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Mfwd_DefaultContext_Ipv6_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "ttl-threshold" { return "TtlThreshold" }
    if yname == "enable-on-interface" { return "EnableOnInterface" }
    if yname == "boundary" { return "Boundary" }
    return ""
}

func (self *Mfwd_DefaultContext_Ipv6_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Mfwd_DefaultContext_Ipv6_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Mfwd_DefaultContext_Ipv6_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Mfwd_DefaultContext_Ipv6_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["ttl-threshold"] = self.TtlThreshold
    leafs["enable-on-interface"] = self.EnableOnInterface
    leafs["boundary"] = self.Boundary
    return leafs
}

func (self *Mfwd_DefaultContext_Ipv6_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Mfwd_DefaultContext_Ipv6_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Mfwd_DefaultContext_Ipv6_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Mfwd_DefaultContext_Ipv6_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Mfwd_DefaultContext_Ipv6_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Mfwd_DefaultContext_Ipv6_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Mfwd_DefaultContext_Ipv6_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Mfwd_DefaultContext_Ipv6_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Mfwd_DefaultContext_Ipv4
// IPV4 commands in the default context
type Mfwd_DefaultContext_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable out-of-memory handling. The type is interface{}.
    OutOfMemoryHandling interface{}

    // Configure all interfaces for multicast routing and forwarding. The type is
    // interface{}.
    EnableOnAllInterfaces interface{}

    // Disable state limit maximum checking. The type is interface{}.
    MaximumCheckingDisable interface{}

    // Enable per (S,G) rate calculation. The type is interface{}.
    RatePerRoute interface{}

    // Disable interface inheritance configuration. The type is interface{}.
    InterfaceInheritanceDisable interface{}

    // Mofrr Lockout timer value. The type is interface{} with range: 1..3600.
    MofrrLockoutTimerConfig interface{}

    // Knob to delay traffic being forwarded on a route. The type is interface{}
    // with range: 5..500. Units are millisecond.
    ForwardingLatency interface{}

    // Mofrr Loss Detection timer value. The type is interface{} with range:
    // 1..3600.
    MofrrLossDetectionTimerConfig interface{}

    // Enable IP multicast routing and forwarding. The type is interface{}.
    MulticastForwarding interface{}

    // Enable logging trap event. The type is interface{}.
    LogTraps interface{}

    // Per-prefix accounting mode. The type is AccountingMode.
    Accounting interface{}

    // Configure a static RPF rule for a given prefix.
    StaticRpfRules Mfwd_DefaultContext_Ipv4_StaticRpfRules

    // Interface-level Configuration.
    Interfaces Mfwd_DefaultContext_Ipv4_Interfaces
}

func (ipv4 *Mfwd_DefaultContext_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Mfwd_DefaultContext_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Mfwd_DefaultContext_Ipv4) GetGoName(yname string) string {
    if yname == "out-of-memory-handling" { return "OutOfMemoryHandling" }
    if yname == "enable-on-all-interfaces" { return "EnableOnAllInterfaces" }
    if yname == "maximum-checking-disable" { return "MaximumCheckingDisable" }
    if yname == "rate-per-route" { return "RatePerRoute" }
    if yname == "interface-inheritance-disable" { return "InterfaceInheritanceDisable" }
    if yname == "mofrr-lockout-timer-config" { return "MofrrLockoutTimerConfig" }
    if yname == "forwarding-latency" { return "ForwardingLatency" }
    if yname == "mofrr-loss-detection-timer-config" { return "MofrrLossDetectionTimerConfig" }
    if yname == "multicast-forwarding" { return "MulticastForwarding" }
    if yname == "log-traps" { return "LogTraps" }
    if yname == "accounting" { return "Accounting" }
    if yname == "static-rpf-rules" { return "StaticRpfRules" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (ipv4 *Mfwd_DefaultContext_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Mfwd_DefaultContext_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static-rpf-rules" {
        return &ipv4.StaticRpfRules
    }
    if childYangName == "interfaces" {
        return &ipv4.Interfaces
    }
    return nil
}

func (ipv4 *Mfwd_DefaultContext_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["static-rpf-rules"] = &ipv4.StaticRpfRules
    children["interfaces"] = &ipv4.Interfaces
    return children
}

func (ipv4 *Mfwd_DefaultContext_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["out-of-memory-handling"] = ipv4.OutOfMemoryHandling
    leafs["enable-on-all-interfaces"] = ipv4.EnableOnAllInterfaces
    leafs["maximum-checking-disable"] = ipv4.MaximumCheckingDisable
    leafs["rate-per-route"] = ipv4.RatePerRoute
    leafs["interface-inheritance-disable"] = ipv4.InterfaceInheritanceDisable
    leafs["mofrr-lockout-timer-config"] = ipv4.MofrrLockoutTimerConfig
    leafs["forwarding-latency"] = ipv4.ForwardingLatency
    leafs["mofrr-loss-detection-timer-config"] = ipv4.MofrrLossDetectionTimerConfig
    leafs["multicast-forwarding"] = ipv4.MulticastForwarding
    leafs["log-traps"] = ipv4.LogTraps
    leafs["accounting"] = ipv4.Accounting
    return leafs
}

func (ipv4 *Mfwd_DefaultContext_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Mfwd_DefaultContext_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Mfwd_DefaultContext_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Mfwd_DefaultContext_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Mfwd_DefaultContext_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Mfwd_DefaultContext_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Mfwd_DefaultContext_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Mfwd_DefaultContext_Ipv4) GetParentYangName() string { return "default-context" }

// Mfwd_DefaultContext_Ipv4_StaticRpfRules
// Configure a static RPF rule for a given prefix
type Mfwd_DefaultContext_Ipv4_StaticRpfRules struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RPF prefix address and mask. The type is slice of
    // Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule.
    StaticRpfRule []Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule
}

func (staticRpfRules *Mfwd_DefaultContext_Ipv4_StaticRpfRules) GetFilter() yfilter.YFilter { return staticRpfRules.YFilter }

func (staticRpfRules *Mfwd_DefaultContext_Ipv4_StaticRpfRules) SetFilter(yf yfilter.YFilter) { staticRpfRules.YFilter = yf }

func (staticRpfRules *Mfwd_DefaultContext_Ipv4_StaticRpfRules) GetGoName(yname string) string {
    if yname == "static-rpf-rule" { return "StaticRpfRule" }
    return ""
}

func (staticRpfRules *Mfwd_DefaultContext_Ipv4_StaticRpfRules) GetSegmentPath() string {
    return "static-rpf-rules"
}

func (staticRpfRules *Mfwd_DefaultContext_Ipv4_StaticRpfRules) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static-rpf-rule" {
        for _, c := range staticRpfRules.StaticRpfRule {
            if staticRpfRules.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule{}
        staticRpfRules.StaticRpfRule = append(staticRpfRules.StaticRpfRule, child)
        return &staticRpfRules.StaticRpfRule[len(staticRpfRules.StaticRpfRule)-1]
    }
    return nil
}

func (staticRpfRules *Mfwd_DefaultContext_Ipv4_StaticRpfRules) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range staticRpfRules.StaticRpfRule {
        children[staticRpfRules.StaticRpfRule[i].GetSegmentPath()] = &staticRpfRules.StaticRpfRule[i]
    }
    return children
}

func (staticRpfRules *Mfwd_DefaultContext_Ipv4_StaticRpfRules) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (staticRpfRules *Mfwd_DefaultContext_Ipv4_StaticRpfRules) GetBundleName() string { return "cisco_ios_xr" }

func (staticRpfRules *Mfwd_DefaultContext_Ipv4_StaticRpfRules) GetYangName() string { return "static-rpf-rules" }

func (staticRpfRules *Mfwd_DefaultContext_Ipv4_StaticRpfRules) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticRpfRules *Mfwd_DefaultContext_Ipv4_StaticRpfRules) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticRpfRules *Mfwd_DefaultContext_Ipv4_StaticRpfRules) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticRpfRules *Mfwd_DefaultContext_Ipv4_StaticRpfRules) SetParent(parent types.Entity) { staticRpfRules.parent = parent }

func (staticRpfRules *Mfwd_DefaultContext_Ipv4_StaticRpfRules) GetParent() types.Entity { return staticRpfRules.parent }

func (staticRpfRules *Mfwd_DefaultContext_Ipv4_StaticRpfRules) GetParentYangName() string { return "ipv4" }

// Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule
// RPF prefix address and mask
type Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the RPF prefix. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. Prefix mask of the RPF Prefix. The type is
    // interface{} with range: 0..128.
    PrefixMask interface{}

    // Neighbor address of the RPF Prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    NeighborAddress interface{}

    // The name of the RPF interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+. This attribute is mandatory.
    InterfaceName interface{}
}

func (staticRpfRule *Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule) GetFilter() yfilter.YFilter { return staticRpfRule.YFilter }

func (staticRpfRule *Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule) SetFilter(yf yfilter.YFilter) { staticRpfRule.YFilter = yf }

func (staticRpfRule *Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-mask" { return "PrefixMask" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (staticRpfRule *Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule) GetSegmentPath() string {
    return "static-rpf-rule" + "[address='" + fmt.Sprintf("%v", staticRpfRule.Address) + "']" + "[prefix-mask='" + fmt.Sprintf("%v", staticRpfRule.PrefixMask) + "']"
}

func (staticRpfRule *Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticRpfRule *Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticRpfRule *Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = staticRpfRule.Address
    leafs["prefix-mask"] = staticRpfRule.PrefixMask
    leafs["neighbor-address"] = staticRpfRule.NeighborAddress
    leafs["interface-name"] = staticRpfRule.InterfaceName
    return leafs
}

func (staticRpfRule *Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule) GetBundleName() string { return "cisco_ios_xr" }

func (staticRpfRule *Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule) GetYangName() string { return "static-rpf-rule" }

func (staticRpfRule *Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticRpfRule *Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticRpfRule *Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticRpfRule *Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule) SetParent(parent types.Entity) { staticRpfRule.parent = parent }

func (staticRpfRule *Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule) GetParent() types.Entity { return staticRpfRule.parent }

func (staticRpfRule *Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule) GetParentYangName() string { return "static-rpf-rules" }

// Mfwd_DefaultContext_Ipv4_Interfaces
// Interface-level Configuration
type Mfwd_DefaultContext_Ipv4_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Mfwd_DefaultContext_Ipv4_Interfaces_Interface.
    Interface []Mfwd_DefaultContext_Ipv4_Interfaces_Interface
}

func (interfaces *Mfwd_DefaultContext_Ipv4_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Mfwd_DefaultContext_Ipv4_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Mfwd_DefaultContext_Ipv4_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Mfwd_DefaultContext_Ipv4_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Mfwd_DefaultContext_Ipv4_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mfwd_DefaultContext_Ipv4_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Mfwd_DefaultContext_Ipv4_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Mfwd_DefaultContext_Ipv4_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Mfwd_DefaultContext_Ipv4_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Mfwd_DefaultContext_Ipv4_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Mfwd_DefaultContext_Ipv4_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Mfwd_DefaultContext_Ipv4_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Mfwd_DefaultContext_Ipv4_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Mfwd_DefaultContext_Ipv4_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Mfwd_DefaultContext_Ipv4_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Mfwd_DefaultContext_Ipv4_Interfaces) GetParentYangName() string { return "ipv4" }

// Mfwd_DefaultContext_Ipv4_Interfaces_Interface
// The name of the interface
type Mfwd_DefaultContext_Ipv4_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // TTL threshold for multicast packets. The type is interface{} with range:
    // 1..255.
    TtlThreshold interface{}

    // Enable or disable IP multicast on the interface. The type is bool.
    EnableOnInterface interface{}

    // Boundary for administratively scoped multicast addresses. The type is
    // string with length: 1..64.
    Boundary interface{}
}

func (self *Mfwd_DefaultContext_Ipv4_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Mfwd_DefaultContext_Ipv4_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Mfwd_DefaultContext_Ipv4_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "ttl-threshold" { return "TtlThreshold" }
    if yname == "enable-on-interface" { return "EnableOnInterface" }
    if yname == "boundary" { return "Boundary" }
    return ""
}

func (self *Mfwd_DefaultContext_Ipv4_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Mfwd_DefaultContext_Ipv4_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Mfwd_DefaultContext_Ipv4_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Mfwd_DefaultContext_Ipv4_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["ttl-threshold"] = self.TtlThreshold
    leafs["enable-on-interface"] = self.EnableOnInterface
    leafs["boundary"] = self.Boundary
    return leafs
}

func (self *Mfwd_DefaultContext_Ipv4_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Mfwd_DefaultContext_Ipv4_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Mfwd_DefaultContext_Ipv4_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Mfwd_DefaultContext_Ipv4_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Mfwd_DefaultContext_Ipv4_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Mfwd_DefaultContext_Ipv4_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Mfwd_DefaultContext_Ipv4_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Mfwd_DefaultContext_Ipv4_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Mfwd_Vrfs
// VRF Table
type Mfwd_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF table names. The type is slice of Mfwd_Vrfs_Vrf.
    Vrf []Mfwd_Vrfs_Vrf
}

func (vrfs *Mfwd_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Mfwd_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Mfwd_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Mfwd_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Mfwd_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mfwd_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Mfwd_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Mfwd_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Mfwd_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Mfwd_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Mfwd_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Mfwd_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Mfwd_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Mfwd_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Mfwd_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Mfwd_Vrfs) GetParentYangName() string { return "mfwd" }

// Mfwd_Vrfs_Vrf
// VRF table names
type Mfwd_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. VRF Name. The type is string with length: 1..32.
    VrfName interface{}

    // VRF table for IPV6 commands.
    Ipv6 Mfwd_Vrfs_Vrf_Ipv6

    // VRF table for IPV4 commands.
    Ipv4 Mfwd_Vrfs_Vrf_Ipv4
}

func (vrf *Mfwd_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Mfwd_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Mfwd_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "ipv6" { return "Ipv6" }
    if yname == "ipv4" { return "Ipv4" }
    return ""
}

func (vrf *Mfwd_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Mfwd_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6" {
        return &vrf.Ipv6
    }
    if childYangName == "ipv4" {
        return &vrf.Ipv4
    }
    return nil
}

func (vrf *Mfwd_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6"] = &vrf.Ipv6
    children["ipv4"] = &vrf.Ipv4
    return children
}

func (vrf *Mfwd_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    return leafs
}

func (vrf *Mfwd_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Mfwd_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Mfwd_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Mfwd_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Mfwd_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Mfwd_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Mfwd_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Mfwd_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

// Mfwd_Vrfs_Vrf_Ipv6
// VRF table for IPV6 commands
type Mfwd_Vrfs_Vrf_Ipv6 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure all interfaces for multicast routing and forwarding. The type is
    // interface{}.
    EnableOnAllInterfaces interface{}

    // Enable per (S,G) rate calculation. The type is interface{}.
    RatePerRoute interface{}

    // Enable IP multicast routing and forwarding. The type is interface{}.
    MulticastForwarding interface{}

    // Enable logging trap event. The type is interface{}.
    LogTraps interface{}

    // Per-prefix accounting mode. The type is AccountingMode.
    Accounting interface{}

    // Configure a static RPF rule for a given prefix.
    StaticRpfRules Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules

    // Interface-level Configuration.
    Interfaces Mfwd_Vrfs_Vrf_Ipv6_Interfaces
}

func (ipv6 *Mfwd_Vrfs_Vrf_Ipv6) GetFilter() yfilter.YFilter { return ipv6.YFilter }

func (ipv6 *Mfwd_Vrfs_Vrf_Ipv6) SetFilter(yf yfilter.YFilter) { ipv6.YFilter = yf }

func (ipv6 *Mfwd_Vrfs_Vrf_Ipv6) GetGoName(yname string) string {
    if yname == "enable-on-all-interfaces" { return "EnableOnAllInterfaces" }
    if yname == "rate-per-route" { return "RatePerRoute" }
    if yname == "multicast-forwarding" { return "MulticastForwarding" }
    if yname == "log-traps" { return "LogTraps" }
    if yname == "accounting" { return "Accounting" }
    if yname == "static-rpf-rules" { return "StaticRpfRules" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (ipv6 *Mfwd_Vrfs_Vrf_Ipv6) GetSegmentPath() string {
    return "ipv6"
}

func (ipv6 *Mfwd_Vrfs_Vrf_Ipv6) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static-rpf-rules" {
        return &ipv6.StaticRpfRules
    }
    if childYangName == "interfaces" {
        return &ipv6.Interfaces
    }
    return nil
}

func (ipv6 *Mfwd_Vrfs_Vrf_Ipv6) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["static-rpf-rules"] = &ipv6.StaticRpfRules
    children["interfaces"] = &ipv6.Interfaces
    return children
}

func (ipv6 *Mfwd_Vrfs_Vrf_Ipv6) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable-on-all-interfaces"] = ipv6.EnableOnAllInterfaces
    leafs["rate-per-route"] = ipv6.RatePerRoute
    leafs["multicast-forwarding"] = ipv6.MulticastForwarding
    leafs["log-traps"] = ipv6.LogTraps
    leafs["accounting"] = ipv6.Accounting
    return leafs
}

func (ipv6 *Mfwd_Vrfs_Vrf_Ipv6) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6 *Mfwd_Vrfs_Vrf_Ipv6) GetYangName() string { return "ipv6" }

func (ipv6 *Mfwd_Vrfs_Vrf_Ipv6) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6 *Mfwd_Vrfs_Vrf_Ipv6) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6 *Mfwd_Vrfs_Vrf_Ipv6) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6 *Mfwd_Vrfs_Vrf_Ipv6) SetParent(parent types.Entity) { ipv6.parent = parent }

func (ipv6 *Mfwd_Vrfs_Vrf_Ipv6) GetParent() types.Entity { return ipv6.parent }

func (ipv6 *Mfwd_Vrfs_Vrf_Ipv6) GetParentYangName() string { return "vrf" }

// Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules
// Configure a static RPF rule for a given prefix
type Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RPF prefix address and mask. The type is slice of
    // Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule.
    StaticRpfRule []Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule
}

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules) GetFilter() yfilter.YFilter { return staticRpfRules.YFilter }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules) SetFilter(yf yfilter.YFilter) { staticRpfRules.YFilter = yf }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules) GetGoName(yname string) string {
    if yname == "static-rpf-rule" { return "StaticRpfRule" }
    return ""
}

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules) GetSegmentPath() string {
    return "static-rpf-rules"
}

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static-rpf-rule" {
        for _, c := range staticRpfRules.StaticRpfRule {
            if staticRpfRules.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule{}
        staticRpfRules.StaticRpfRule = append(staticRpfRules.StaticRpfRule, child)
        return &staticRpfRules.StaticRpfRule[len(staticRpfRules.StaticRpfRule)-1]
    }
    return nil
}

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range staticRpfRules.StaticRpfRule {
        children[staticRpfRules.StaticRpfRule[i].GetSegmentPath()] = &staticRpfRules.StaticRpfRule[i]
    }
    return children
}

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules) GetBundleName() string { return "cisco_ios_xr" }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules) GetYangName() string { return "static-rpf-rules" }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules) SetParent(parent types.Entity) { staticRpfRules.parent = parent }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules) GetParent() types.Entity { return staticRpfRules.parent }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules) GetParentYangName() string { return "ipv6" }

// Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule
// RPF prefix address and mask
type Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the RPF prefix. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. Prefix mask of the RPF Prefix. The type is
    // interface{} with range: 0..128.
    PrefixMask interface{}

    // Neighbor address of the RPF Prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    NeighborAddress interface{}

    // The name of the RPF interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+. This attribute is mandatory.
    InterfaceName interface{}
}

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule) GetFilter() yfilter.YFilter { return staticRpfRule.YFilter }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule) SetFilter(yf yfilter.YFilter) { staticRpfRule.YFilter = yf }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-mask" { return "PrefixMask" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule) GetSegmentPath() string {
    return "static-rpf-rule" + "[address='" + fmt.Sprintf("%v", staticRpfRule.Address) + "']" + "[prefix-mask='" + fmt.Sprintf("%v", staticRpfRule.PrefixMask) + "']"
}

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = staticRpfRule.Address
    leafs["prefix-mask"] = staticRpfRule.PrefixMask
    leafs["neighbor-address"] = staticRpfRule.NeighborAddress
    leafs["interface-name"] = staticRpfRule.InterfaceName
    return leafs
}

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule) GetBundleName() string { return "cisco_ios_xr" }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule) GetYangName() string { return "static-rpf-rule" }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule) SetParent(parent types.Entity) { staticRpfRule.parent = parent }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule) GetParent() types.Entity { return staticRpfRule.parent }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule) GetParentYangName() string { return "static-rpf-rules" }

// Mfwd_Vrfs_Vrf_Ipv6_Interfaces
// Interface-level Configuration
type Mfwd_Vrfs_Vrf_Ipv6_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface.
    Interface []Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface
}

func (interfaces *Mfwd_Vrfs_Vrf_Ipv6_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv6_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv6_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Mfwd_Vrfs_Vrf_Ipv6_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Mfwd_Vrfs_Vrf_Ipv6_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Mfwd_Vrfs_Vrf_Ipv6_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Mfwd_Vrfs_Vrf_Ipv6_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Mfwd_Vrfs_Vrf_Ipv6_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv6_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv6_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv6_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv6_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv6_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv6_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv6_Interfaces) GetParentYangName() string { return "ipv6" }

// Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface
// The name of the interface
type Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // TTL threshold for multicast packets. The type is interface{} with range:
    // 1..255.
    TtlThreshold interface{}

    // Enable or disable IP multicast on the interface. The type is bool.
    EnableOnInterface interface{}

    // Boundary for administratively scoped multicast addresses. The type is
    // string with length: 1..64.
    Boundary interface{}
}

func (self *Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "ttl-threshold" { return "TtlThreshold" }
    if yname == "enable-on-interface" { return "EnableOnInterface" }
    if yname == "boundary" { return "Boundary" }
    return ""
}

func (self *Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["ttl-threshold"] = self.TtlThreshold
    leafs["enable-on-interface"] = self.EnableOnInterface
    leafs["boundary"] = self.Boundary
    return leafs
}

func (self *Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Mfwd_Vrfs_Vrf_Ipv4
// VRF table for IPV4 commands
type Mfwd_Vrfs_Vrf_Ipv4 struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure all interfaces for multicast routing and forwarding. The type is
    // interface{}.
    EnableOnAllInterfaces interface{}

    // Enable per (S,G) rate calculation. The type is interface{}.
    RatePerRoute interface{}

    // Enable IP multicast routing and forwarding. The type is interface{}.
    MulticastForwarding interface{}

    // Enable logging trap event. The type is interface{}.
    LogTraps interface{}

    // Per-prefix accounting mode. The type is AccountingMode.
    Accounting interface{}

    // Configure a static RPF rule for a given prefix.
    StaticRpfRules Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules

    // Interface-level Configuration.
    Interfaces Mfwd_Vrfs_Vrf_Ipv4_Interfaces
}

func (ipv4 *Mfwd_Vrfs_Vrf_Ipv4) GetFilter() yfilter.YFilter { return ipv4.YFilter }

func (ipv4 *Mfwd_Vrfs_Vrf_Ipv4) SetFilter(yf yfilter.YFilter) { ipv4.YFilter = yf }

func (ipv4 *Mfwd_Vrfs_Vrf_Ipv4) GetGoName(yname string) string {
    if yname == "enable-on-all-interfaces" { return "EnableOnAllInterfaces" }
    if yname == "rate-per-route" { return "RatePerRoute" }
    if yname == "multicast-forwarding" { return "MulticastForwarding" }
    if yname == "log-traps" { return "LogTraps" }
    if yname == "accounting" { return "Accounting" }
    if yname == "static-rpf-rules" { return "StaticRpfRules" }
    if yname == "interfaces" { return "Interfaces" }
    return ""
}

func (ipv4 *Mfwd_Vrfs_Vrf_Ipv4) GetSegmentPath() string {
    return "ipv4"
}

func (ipv4 *Mfwd_Vrfs_Vrf_Ipv4) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static-rpf-rules" {
        return &ipv4.StaticRpfRules
    }
    if childYangName == "interfaces" {
        return &ipv4.Interfaces
    }
    return nil
}

func (ipv4 *Mfwd_Vrfs_Vrf_Ipv4) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["static-rpf-rules"] = &ipv4.StaticRpfRules
    children["interfaces"] = &ipv4.Interfaces
    return children
}

func (ipv4 *Mfwd_Vrfs_Vrf_Ipv4) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable-on-all-interfaces"] = ipv4.EnableOnAllInterfaces
    leafs["rate-per-route"] = ipv4.RatePerRoute
    leafs["multicast-forwarding"] = ipv4.MulticastForwarding
    leafs["log-traps"] = ipv4.LogTraps
    leafs["accounting"] = ipv4.Accounting
    return leafs
}

func (ipv4 *Mfwd_Vrfs_Vrf_Ipv4) GetBundleName() string { return "cisco_ios_xr" }

func (ipv4 *Mfwd_Vrfs_Vrf_Ipv4) GetYangName() string { return "ipv4" }

func (ipv4 *Mfwd_Vrfs_Vrf_Ipv4) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv4 *Mfwd_Vrfs_Vrf_Ipv4) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv4 *Mfwd_Vrfs_Vrf_Ipv4) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv4 *Mfwd_Vrfs_Vrf_Ipv4) SetParent(parent types.Entity) { ipv4.parent = parent }

func (ipv4 *Mfwd_Vrfs_Vrf_Ipv4) GetParent() types.Entity { return ipv4.parent }

func (ipv4 *Mfwd_Vrfs_Vrf_Ipv4) GetParentYangName() string { return "vrf" }

// Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules
// Configure a static RPF rule for a given prefix
type Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // RPF prefix address and mask. The type is slice of
    // Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule.
    StaticRpfRule []Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule
}

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules) GetFilter() yfilter.YFilter { return staticRpfRules.YFilter }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules) SetFilter(yf yfilter.YFilter) { staticRpfRules.YFilter = yf }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules) GetGoName(yname string) string {
    if yname == "static-rpf-rule" { return "StaticRpfRule" }
    return ""
}

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules) GetSegmentPath() string {
    return "static-rpf-rules"
}

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "static-rpf-rule" {
        for _, c := range staticRpfRules.StaticRpfRule {
            if staticRpfRules.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule{}
        staticRpfRules.StaticRpfRule = append(staticRpfRules.StaticRpfRule, child)
        return &staticRpfRules.StaticRpfRule[len(staticRpfRules.StaticRpfRule)-1]
    }
    return nil
}

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range staticRpfRules.StaticRpfRule {
        children[staticRpfRules.StaticRpfRule[i].GetSegmentPath()] = &staticRpfRules.StaticRpfRule[i]
    }
    return children
}

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules) GetBundleName() string { return "cisco_ios_xr" }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules) GetYangName() string { return "static-rpf-rules" }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules) SetParent(parent types.Entity) { staticRpfRules.parent = parent }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules) GetParent() types.Entity { return staticRpfRules.parent }

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules) GetParentYangName() string { return "ipv4" }

// Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule
// RPF prefix address and mask
type Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the RPF prefix. The type is one of
    // the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // This attribute is a key. Prefix mask of the RPF Prefix. The type is
    // interface{} with range: 0..128.
    PrefixMask interface{}

    // Neighbor address of the RPF Prefix. The type is one of the following types:
    // string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?
    // This attribute is mandatory., or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?
    // This attribute is mandatory..
    NeighborAddress interface{}

    // The name of the RPF interface. The type is string with pattern:
    // [a-zA-Z0-9./-]+. This attribute is mandatory.
    InterfaceName interface{}
}

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule) GetFilter() yfilter.YFilter { return staticRpfRule.YFilter }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule) SetFilter(yf yfilter.YFilter) { staticRpfRule.YFilter = yf }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule) GetGoName(yname string) string {
    if yname == "address" { return "Address" }
    if yname == "prefix-mask" { return "PrefixMask" }
    if yname == "neighbor-address" { return "NeighborAddress" }
    if yname == "interface-name" { return "InterfaceName" }
    return ""
}

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule) GetSegmentPath() string {
    return "static-rpf-rule" + "[address='" + fmt.Sprintf("%v", staticRpfRule.Address) + "']" + "[prefix-mask='" + fmt.Sprintf("%v", staticRpfRule.PrefixMask) + "']"
}

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["address"] = staticRpfRule.Address
    leafs["prefix-mask"] = staticRpfRule.PrefixMask
    leafs["neighbor-address"] = staticRpfRule.NeighborAddress
    leafs["interface-name"] = staticRpfRule.InterfaceName
    return leafs
}

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule) GetBundleName() string { return "cisco_ios_xr" }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule) GetYangName() string { return "static-rpf-rule" }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule) SetParent(parent types.Entity) { staticRpfRule.parent = parent }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule) GetParent() types.Entity { return staticRpfRule.parent }

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule) GetParentYangName() string { return "static-rpf-rules" }

// Mfwd_Vrfs_Vrf_Ipv4_Interfaces
// Interface-level Configuration
type Mfwd_Vrfs_Vrf_Ipv4_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface.
    Interface []Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface
}

func (interfaces *Mfwd_Vrfs_Vrf_Ipv4_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv4_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv4_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Mfwd_Vrfs_Vrf_Ipv4_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Mfwd_Vrfs_Vrf_Ipv4_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Mfwd_Vrfs_Vrf_Ipv4_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Mfwd_Vrfs_Vrf_Ipv4_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Mfwd_Vrfs_Vrf_Ipv4_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv4_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv4_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv4_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv4_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv4_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv4_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Mfwd_Vrfs_Vrf_Ipv4_Interfaces) GetParentYangName() string { return "ipv4" }

// Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface
// The name of the interface
type Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // TTL threshold for multicast packets. The type is interface{} with range:
    // 1..255.
    TtlThreshold interface{}

    // Enable or disable IP multicast on the interface. The type is bool.
    EnableOnInterface interface{}

    // Boundary for administratively scoped multicast addresses. The type is
    // string with length: 1..64.
    Boundary interface{}
}

func (self *Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "ttl-threshold" { return "TtlThreshold" }
    if yname == "enable-on-interface" { return "EnableOnInterface" }
    if yname == "boundary" { return "Boundary" }
    return ""
}

func (self *Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["ttl-threshold"] = self.TtlThreshold
    leafs["enable-on-interface"] = self.EnableOnInterface
    leafs["boundary"] = self.Boundary
    return leafs
}

func (self *Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

