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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Default Context.
    DefaultContext Mfwd_DefaultContext

    // VRF Table.
    Vrfs Mfwd_Vrfs
}

func (mfwd *Mfwd) GetEntityData() *types.CommonEntityData {
    mfwd.EntityData.YFilter = mfwd.YFilter
    mfwd.EntityData.YangName = "mfwd"
    mfwd.EntityData.BundleName = "cisco_ios_xr"
    mfwd.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-mfwd-cfg"
    mfwd.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-mfwd-cfg:mfwd"
    mfwd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mfwd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mfwd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mfwd.EntityData.Children = make(map[string]types.YChild)
    mfwd.EntityData.Children["default-context"] = types.YChild{"DefaultContext", &mfwd.DefaultContext}
    mfwd.EntityData.Children["vrfs"] = types.YChild{"Vrfs", &mfwd.Vrfs}
    mfwd.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(mfwd.EntityData)
}

// Mfwd_DefaultContext
// Default Context
// This type is a presence type.
type Mfwd_DefaultContext struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPV6 commands in the default context.
    Ipv6 Mfwd_DefaultContext_Ipv6

    // IPV4 commands in the default context.
    Ipv4 Mfwd_DefaultContext_Ipv4
}

func (defaultContext *Mfwd_DefaultContext) GetEntityData() *types.CommonEntityData {
    defaultContext.EntityData.YFilter = defaultContext.YFilter
    defaultContext.EntityData.YangName = "default-context"
    defaultContext.EntityData.BundleName = "cisco_ios_xr"
    defaultContext.EntityData.ParentYangName = "mfwd"
    defaultContext.EntityData.SegmentPath = "default-context"
    defaultContext.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultContext.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultContext.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultContext.EntityData.Children = make(map[string]types.YChild)
    defaultContext.EntityData.Children["ipv6"] = types.YChild{"Ipv6", &defaultContext.Ipv6}
    defaultContext.EntityData.Children["ipv4"] = types.YChild{"Ipv4", &defaultContext.Ipv4}
    defaultContext.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(defaultContext.EntityData)
}

// Mfwd_DefaultContext_Ipv6
// IPV6 commands in the default context
type Mfwd_DefaultContext_Ipv6 struct {
    EntityData types.CommonEntityData
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

func (ipv6 *Mfwd_DefaultContext_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "default-context"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Children["static-rpf-rules"] = types.YChild{"StaticRpfRules", &ipv6.StaticRpfRules}
    ipv6.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &ipv6.Interfaces}
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6.EntityData.Leafs["enable-on-all-interfaces"] = types.YLeaf{"EnableOnAllInterfaces", ipv6.EnableOnAllInterfaces}
    ipv6.EntityData.Leafs["maximum-checking-disable"] = types.YLeaf{"MaximumCheckingDisable", ipv6.MaximumCheckingDisable}
    ipv6.EntityData.Leafs["rate-per-route"] = types.YLeaf{"RatePerRoute", ipv6.RatePerRoute}
    ipv6.EntityData.Leafs["interface-inheritance-disable"] = types.YLeaf{"InterfaceInheritanceDisable", ipv6.InterfaceInheritanceDisable}
    ipv6.EntityData.Leafs["mofrr-lockout-timer-config"] = types.YLeaf{"MofrrLockoutTimerConfig", ipv6.MofrrLockoutTimerConfig}
    ipv6.EntityData.Leafs["forwarding-latency"] = types.YLeaf{"ForwardingLatency", ipv6.ForwardingLatency}
    ipv6.EntityData.Leafs["mofrr-loss-detection-timer-config"] = types.YLeaf{"MofrrLossDetectionTimerConfig", ipv6.MofrrLossDetectionTimerConfig}
    ipv6.EntityData.Leafs["multicast-forwarding"] = types.YLeaf{"MulticastForwarding", ipv6.MulticastForwarding}
    ipv6.EntityData.Leafs["log-traps"] = types.YLeaf{"LogTraps", ipv6.LogTraps}
    ipv6.EntityData.Leafs["accounting"] = types.YLeaf{"Accounting", ipv6.Accounting}
    return &(ipv6.EntityData)
}

// Mfwd_DefaultContext_Ipv6_StaticRpfRules
// Configure a static RPF rule for a given prefix
type Mfwd_DefaultContext_Ipv6_StaticRpfRules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RPF prefix address and mask. The type is slice of
    // Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule.
    StaticRpfRule []Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule
}

func (staticRpfRules *Mfwd_DefaultContext_Ipv6_StaticRpfRules) GetEntityData() *types.CommonEntityData {
    staticRpfRules.EntityData.YFilter = staticRpfRules.YFilter
    staticRpfRules.EntityData.YangName = "static-rpf-rules"
    staticRpfRules.EntityData.BundleName = "cisco_ios_xr"
    staticRpfRules.EntityData.ParentYangName = "ipv6"
    staticRpfRules.EntityData.SegmentPath = "static-rpf-rules"
    staticRpfRules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticRpfRules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticRpfRules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticRpfRules.EntityData.Children = make(map[string]types.YChild)
    staticRpfRules.EntityData.Children["static-rpf-rule"] = types.YChild{"StaticRpfRule", nil}
    for i := range staticRpfRules.StaticRpfRule {
        staticRpfRules.EntityData.Children[types.GetSegmentPath(&staticRpfRules.StaticRpfRule[i])] = types.YChild{"StaticRpfRule", &staticRpfRules.StaticRpfRule[i]}
    }
    staticRpfRules.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(staticRpfRules.EntityData)
}

// Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule
// RPF prefix address and mask
type Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the RPF prefix. The type is one of
    // the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // This attribute is a key. Prefix mask of the RPF Prefix. The type is
    // interface{} with range: 0..128.
    PrefixMask interface{}

    // Neighbor address of the RPF Prefix. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory., or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory..
    NeighborAddress interface{}

    // The name of the RPF interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'. This attribute is mandatory.
    InterfaceName interface{}
}

func (staticRpfRule *Mfwd_DefaultContext_Ipv6_StaticRpfRules_StaticRpfRule) GetEntityData() *types.CommonEntityData {
    staticRpfRule.EntityData.YFilter = staticRpfRule.YFilter
    staticRpfRule.EntityData.YangName = "static-rpf-rule"
    staticRpfRule.EntityData.BundleName = "cisco_ios_xr"
    staticRpfRule.EntityData.ParentYangName = "static-rpf-rules"
    staticRpfRule.EntityData.SegmentPath = "static-rpf-rule" + "[address='" + fmt.Sprintf("%v", staticRpfRule.Address) + "']" + "[prefix-mask='" + fmt.Sprintf("%v", staticRpfRule.PrefixMask) + "']"
    staticRpfRule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticRpfRule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticRpfRule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticRpfRule.EntityData.Children = make(map[string]types.YChild)
    staticRpfRule.EntityData.Leafs = make(map[string]types.YLeaf)
    staticRpfRule.EntityData.Leafs["address"] = types.YLeaf{"Address", staticRpfRule.Address}
    staticRpfRule.EntityData.Leafs["prefix-mask"] = types.YLeaf{"PrefixMask", staticRpfRule.PrefixMask}
    staticRpfRule.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", staticRpfRule.NeighborAddress}
    staticRpfRule.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", staticRpfRule.InterfaceName}
    return &(staticRpfRule.EntityData)
}

// Mfwd_DefaultContext_Ipv6_Interfaces
// Interface-level Configuration
type Mfwd_DefaultContext_Ipv6_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Mfwd_DefaultContext_Ipv6_Interfaces_Interface_.
    Interface_ []Mfwd_DefaultContext_Ipv6_Interfaces_Interface
}

func (interfaces *Mfwd_DefaultContext_Ipv6_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "ipv6"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Mfwd_DefaultContext_Ipv6_Interfaces_Interface
// The name of the interface
type Mfwd_DefaultContext_Ipv6_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

func (self *Mfwd_DefaultContext_Ipv6_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["ttl-threshold"] = types.YLeaf{"TtlThreshold", self.TtlThreshold}
    self.EntityData.Leafs["enable-on-interface"] = types.YLeaf{"EnableOnInterface", self.EnableOnInterface}
    self.EntityData.Leafs["boundary"] = types.YLeaf{"Boundary", self.Boundary}
    return &(self.EntityData)
}

// Mfwd_DefaultContext_Ipv4
// IPV4 commands in the default context
type Mfwd_DefaultContext_Ipv4 struct {
    EntityData types.CommonEntityData
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

func (ipv4 *Mfwd_DefaultContext_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "default-context"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Children["static-rpf-rules"] = types.YChild{"StaticRpfRules", &ipv4.StaticRpfRules}
    ipv4.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &ipv4.Interfaces}
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4.EntityData.Leafs["out-of-memory-handling"] = types.YLeaf{"OutOfMemoryHandling", ipv4.OutOfMemoryHandling}
    ipv4.EntityData.Leafs["enable-on-all-interfaces"] = types.YLeaf{"EnableOnAllInterfaces", ipv4.EnableOnAllInterfaces}
    ipv4.EntityData.Leafs["maximum-checking-disable"] = types.YLeaf{"MaximumCheckingDisable", ipv4.MaximumCheckingDisable}
    ipv4.EntityData.Leafs["rate-per-route"] = types.YLeaf{"RatePerRoute", ipv4.RatePerRoute}
    ipv4.EntityData.Leafs["interface-inheritance-disable"] = types.YLeaf{"InterfaceInheritanceDisable", ipv4.InterfaceInheritanceDisable}
    ipv4.EntityData.Leafs["mofrr-lockout-timer-config"] = types.YLeaf{"MofrrLockoutTimerConfig", ipv4.MofrrLockoutTimerConfig}
    ipv4.EntityData.Leafs["forwarding-latency"] = types.YLeaf{"ForwardingLatency", ipv4.ForwardingLatency}
    ipv4.EntityData.Leafs["mofrr-loss-detection-timer-config"] = types.YLeaf{"MofrrLossDetectionTimerConfig", ipv4.MofrrLossDetectionTimerConfig}
    ipv4.EntityData.Leafs["multicast-forwarding"] = types.YLeaf{"MulticastForwarding", ipv4.MulticastForwarding}
    ipv4.EntityData.Leafs["log-traps"] = types.YLeaf{"LogTraps", ipv4.LogTraps}
    ipv4.EntityData.Leafs["accounting"] = types.YLeaf{"Accounting", ipv4.Accounting}
    return &(ipv4.EntityData)
}

// Mfwd_DefaultContext_Ipv4_StaticRpfRules
// Configure a static RPF rule for a given prefix
type Mfwd_DefaultContext_Ipv4_StaticRpfRules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RPF prefix address and mask. The type is slice of
    // Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule.
    StaticRpfRule []Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule
}

func (staticRpfRules *Mfwd_DefaultContext_Ipv4_StaticRpfRules) GetEntityData() *types.CommonEntityData {
    staticRpfRules.EntityData.YFilter = staticRpfRules.YFilter
    staticRpfRules.EntityData.YangName = "static-rpf-rules"
    staticRpfRules.EntityData.BundleName = "cisco_ios_xr"
    staticRpfRules.EntityData.ParentYangName = "ipv4"
    staticRpfRules.EntityData.SegmentPath = "static-rpf-rules"
    staticRpfRules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticRpfRules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticRpfRules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticRpfRules.EntityData.Children = make(map[string]types.YChild)
    staticRpfRules.EntityData.Children["static-rpf-rule"] = types.YChild{"StaticRpfRule", nil}
    for i := range staticRpfRules.StaticRpfRule {
        staticRpfRules.EntityData.Children[types.GetSegmentPath(&staticRpfRules.StaticRpfRule[i])] = types.YChild{"StaticRpfRule", &staticRpfRules.StaticRpfRule[i]}
    }
    staticRpfRules.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(staticRpfRules.EntityData)
}

// Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule
// RPF prefix address and mask
type Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the RPF prefix. The type is one of
    // the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // This attribute is a key. Prefix mask of the RPF Prefix. The type is
    // interface{} with range: 0..128.
    PrefixMask interface{}

    // Neighbor address of the RPF Prefix. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory., or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory..
    NeighborAddress interface{}

    // The name of the RPF interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'. This attribute is mandatory.
    InterfaceName interface{}
}

func (staticRpfRule *Mfwd_DefaultContext_Ipv4_StaticRpfRules_StaticRpfRule) GetEntityData() *types.CommonEntityData {
    staticRpfRule.EntityData.YFilter = staticRpfRule.YFilter
    staticRpfRule.EntityData.YangName = "static-rpf-rule"
    staticRpfRule.EntityData.BundleName = "cisco_ios_xr"
    staticRpfRule.EntityData.ParentYangName = "static-rpf-rules"
    staticRpfRule.EntityData.SegmentPath = "static-rpf-rule" + "[address='" + fmt.Sprintf("%v", staticRpfRule.Address) + "']" + "[prefix-mask='" + fmt.Sprintf("%v", staticRpfRule.PrefixMask) + "']"
    staticRpfRule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticRpfRule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticRpfRule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticRpfRule.EntityData.Children = make(map[string]types.YChild)
    staticRpfRule.EntityData.Leafs = make(map[string]types.YLeaf)
    staticRpfRule.EntityData.Leafs["address"] = types.YLeaf{"Address", staticRpfRule.Address}
    staticRpfRule.EntityData.Leafs["prefix-mask"] = types.YLeaf{"PrefixMask", staticRpfRule.PrefixMask}
    staticRpfRule.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", staticRpfRule.NeighborAddress}
    staticRpfRule.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", staticRpfRule.InterfaceName}
    return &(staticRpfRule.EntityData)
}

// Mfwd_DefaultContext_Ipv4_Interfaces
// Interface-level Configuration
type Mfwd_DefaultContext_Ipv4_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Mfwd_DefaultContext_Ipv4_Interfaces_Interface_.
    Interface_ []Mfwd_DefaultContext_Ipv4_Interfaces_Interface
}

func (interfaces *Mfwd_DefaultContext_Ipv4_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "ipv4"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Mfwd_DefaultContext_Ipv4_Interfaces_Interface
// The name of the interface
type Mfwd_DefaultContext_Ipv4_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

func (self *Mfwd_DefaultContext_Ipv4_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["ttl-threshold"] = types.YLeaf{"TtlThreshold", self.TtlThreshold}
    self.EntityData.Leafs["enable-on-interface"] = types.YLeaf{"EnableOnInterface", self.EnableOnInterface}
    self.EntityData.Leafs["boundary"] = types.YLeaf{"Boundary", self.Boundary}
    return &(self.EntityData)
}

// Mfwd_Vrfs
// VRF Table
type Mfwd_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF table names. The type is slice of Mfwd_Vrfs_Vrf.
    Vrf []Mfwd_Vrfs_Vrf
}

func (vrfs *Mfwd_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "mfwd"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = make(map[string]types.YChild)
    vrfs.EntityData.Children["vrf"] = types.YChild{"Vrf", nil}
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children[types.GetSegmentPath(&vrfs.Vrf[i])] = types.YChild{"Vrf", &vrfs.Vrf[i]}
    }
    vrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(vrfs.EntityData)
}

// Mfwd_Vrfs_Vrf
// VRF table names
type Mfwd_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF Name. The type is string with length: 1..32.
    VrfName interface{}

    // VRF table for IPV6 commands.
    Ipv6 Mfwd_Vrfs_Vrf_Ipv6

    // VRF table for IPV4 commands.
    Ipv4 Mfwd_Vrfs_Vrf_Ipv4
}

func (vrf *Mfwd_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = make(map[string]types.YChild)
    vrf.EntityData.Children["ipv6"] = types.YChild{"Ipv6", &vrf.Ipv6}
    vrf.EntityData.Children["ipv4"] = types.YChild{"Ipv4", &vrf.Ipv4}
    vrf.EntityData.Leafs = make(map[string]types.YLeaf)
    vrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", vrf.VrfName}
    return &(vrf.EntityData)
}

// Mfwd_Vrfs_Vrf_Ipv6
// VRF table for IPV6 commands
type Mfwd_Vrfs_Vrf_Ipv6 struct {
    EntityData types.CommonEntityData
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

func (ipv6 *Mfwd_Vrfs_Vrf_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "vrf"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Children["static-rpf-rules"] = types.YChild{"StaticRpfRules", &ipv6.StaticRpfRules}
    ipv6.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &ipv6.Interfaces}
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6.EntityData.Leafs["enable-on-all-interfaces"] = types.YLeaf{"EnableOnAllInterfaces", ipv6.EnableOnAllInterfaces}
    ipv6.EntityData.Leafs["rate-per-route"] = types.YLeaf{"RatePerRoute", ipv6.RatePerRoute}
    ipv6.EntityData.Leafs["multicast-forwarding"] = types.YLeaf{"MulticastForwarding", ipv6.MulticastForwarding}
    ipv6.EntityData.Leafs["log-traps"] = types.YLeaf{"LogTraps", ipv6.LogTraps}
    ipv6.EntityData.Leafs["accounting"] = types.YLeaf{"Accounting", ipv6.Accounting}
    return &(ipv6.EntityData)
}

// Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules
// Configure a static RPF rule for a given prefix
type Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RPF prefix address and mask. The type is slice of
    // Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule.
    StaticRpfRule []Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule
}

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules) GetEntityData() *types.CommonEntityData {
    staticRpfRules.EntityData.YFilter = staticRpfRules.YFilter
    staticRpfRules.EntityData.YangName = "static-rpf-rules"
    staticRpfRules.EntityData.BundleName = "cisco_ios_xr"
    staticRpfRules.EntityData.ParentYangName = "ipv6"
    staticRpfRules.EntityData.SegmentPath = "static-rpf-rules"
    staticRpfRules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticRpfRules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticRpfRules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticRpfRules.EntityData.Children = make(map[string]types.YChild)
    staticRpfRules.EntityData.Children["static-rpf-rule"] = types.YChild{"StaticRpfRule", nil}
    for i := range staticRpfRules.StaticRpfRule {
        staticRpfRules.EntityData.Children[types.GetSegmentPath(&staticRpfRules.StaticRpfRule[i])] = types.YChild{"StaticRpfRule", &staticRpfRules.StaticRpfRule[i]}
    }
    staticRpfRules.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(staticRpfRules.EntityData)
}

// Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule
// RPF prefix address and mask
type Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the RPF prefix. The type is one of
    // the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // This attribute is a key. Prefix mask of the RPF Prefix. The type is
    // interface{} with range: 0..128.
    PrefixMask interface{}

    // Neighbor address of the RPF Prefix. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory., or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory..
    NeighborAddress interface{}

    // The name of the RPF interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'. This attribute is mandatory.
    InterfaceName interface{}
}

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv6_StaticRpfRules_StaticRpfRule) GetEntityData() *types.CommonEntityData {
    staticRpfRule.EntityData.YFilter = staticRpfRule.YFilter
    staticRpfRule.EntityData.YangName = "static-rpf-rule"
    staticRpfRule.EntityData.BundleName = "cisco_ios_xr"
    staticRpfRule.EntityData.ParentYangName = "static-rpf-rules"
    staticRpfRule.EntityData.SegmentPath = "static-rpf-rule" + "[address='" + fmt.Sprintf("%v", staticRpfRule.Address) + "']" + "[prefix-mask='" + fmt.Sprintf("%v", staticRpfRule.PrefixMask) + "']"
    staticRpfRule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticRpfRule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticRpfRule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticRpfRule.EntityData.Children = make(map[string]types.YChild)
    staticRpfRule.EntityData.Leafs = make(map[string]types.YLeaf)
    staticRpfRule.EntityData.Leafs["address"] = types.YLeaf{"Address", staticRpfRule.Address}
    staticRpfRule.EntityData.Leafs["prefix-mask"] = types.YLeaf{"PrefixMask", staticRpfRule.PrefixMask}
    staticRpfRule.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", staticRpfRule.NeighborAddress}
    staticRpfRule.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", staticRpfRule.InterfaceName}
    return &(staticRpfRule.EntityData)
}

// Mfwd_Vrfs_Vrf_Ipv6_Interfaces
// Interface-level Configuration
type Mfwd_Vrfs_Vrf_Ipv6_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface_.
    Interface_ []Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface
}

func (interfaces *Mfwd_Vrfs_Vrf_Ipv6_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "ipv6"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface
// The name of the interface
type Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

func (self *Mfwd_Vrfs_Vrf_Ipv6_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["ttl-threshold"] = types.YLeaf{"TtlThreshold", self.TtlThreshold}
    self.EntityData.Leafs["enable-on-interface"] = types.YLeaf{"EnableOnInterface", self.EnableOnInterface}
    self.EntityData.Leafs["boundary"] = types.YLeaf{"Boundary", self.Boundary}
    return &(self.EntityData)
}

// Mfwd_Vrfs_Vrf_Ipv4
// VRF table for IPV4 commands
type Mfwd_Vrfs_Vrf_Ipv4 struct {
    EntityData types.CommonEntityData
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

func (ipv4 *Mfwd_Vrfs_Vrf_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "vrf"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Children["static-rpf-rules"] = types.YChild{"StaticRpfRules", &ipv4.StaticRpfRules}
    ipv4.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &ipv4.Interfaces}
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv4.EntityData.Leafs["enable-on-all-interfaces"] = types.YLeaf{"EnableOnAllInterfaces", ipv4.EnableOnAllInterfaces}
    ipv4.EntityData.Leafs["rate-per-route"] = types.YLeaf{"RatePerRoute", ipv4.RatePerRoute}
    ipv4.EntityData.Leafs["multicast-forwarding"] = types.YLeaf{"MulticastForwarding", ipv4.MulticastForwarding}
    ipv4.EntityData.Leafs["log-traps"] = types.YLeaf{"LogTraps", ipv4.LogTraps}
    ipv4.EntityData.Leafs["accounting"] = types.YLeaf{"Accounting", ipv4.Accounting}
    return &(ipv4.EntityData)
}

// Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules
// Configure a static RPF rule for a given prefix
type Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RPF prefix address and mask. The type is slice of
    // Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule.
    StaticRpfRule []Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule
}

func (staticRpfRules *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules) GetEntityData() *types.CommonEntityData {
    staticRpfRules.EntityData.YFilter = staticRpfRules.YFilter
    staticRpfRules.EntityData.YangName = "static-rpf-rules"
    staticRpfRules.EntityData.BundleName = "cisco_ios_xr"
    staticRpfRules.EntityData.ParentYangName = "ipv4"
    staticRpfRules.EntityData.SegmentPath = "static-rpf-rules"
    staticRpfRules.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticRpfRules.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticRpfRules.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticRpfRules.EntityData.Children = make(map[string]types.YChild)
    staticRpfRules.EntityData.Children["static-rpf-rule"] = types.YChild{"StaticRpfRule", nil}
    for i := range staticRpfRules.StaticRpfRule {
        staticRpfRules.EntityData.Children[types.GetSegmentPath(&staticRpfRules.StaticRpfRule[i])] = types.YChild{"StaticRpfRule", &staticRpfRules.StaticRpfRule[i]}
    }
    staticRpfRules.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(staticRpfRules.EntityData)
}

// Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule
// RPF prefix address and mask
type Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. IP address of the RPF prefix. The type is one of
    // the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // This attribute is a key. Prefix mask of the RPF Prefix. The type is
    // interface{} with range: 0..128.
    PrefixMask interface{}

    // Neighbor address of the RPF Prefix. The type is one of the following types:
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory., or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'
    // This attribute is mandatory..
    NeighborAddress interface{}

    // The name of the RPF interface. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'. This attribute is mandatory.
    InterfaceName interface{}
}

func (staticRpfRule *Mfwd_Vrfs_Vrf_Ipv4_StaticRpfRules_StaticRpfRule) GetEntityData() *types.CommonEntityData {
    staticRpfRule.EntityData.YFilter = staticRpfRule.YFilter
    staticRpfRule.EntityData.YangName = "static-rpf-rule"
    staticRpfRule.EntityData.BundleName = "cisco_ios_xr"
    staticRpfRule.EntityData.ParentYangName = "static-rpf-rules"
    staticRpfRule.EntityData.SegmentPath = "static-rpf-rule" + "[address='" + fmt.Sprintf("%v", staticRpfRule.Address) + "']" + "[prefix-mask='" + fmt.Sprintf("%v", staticRpfRule.PrefixMask) + "']"
    staticRpfRule.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    staticRpfRule.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    staticRpfRule.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    staticRpfRule.EntityData.Children = make(map[string]types.YChild)
    staticRpfRule.EntityData.Leafs = make(map[string]types.YLeaf)
    staticRpfRule.EntityData.Leafs["address"] = types.YLeaf{"Address", staticRpfRule.Address}
    staticRpfRule.EntityData.Leafs["prefix-mask"] = types.YLeaf{"PrefixMask", staticRpfRule.PrefixMask}
    staticRpfRule.EntityData.Leafs["neighbor-address"] = types.YLeaf{"NeighborAddress", staticRpfRule.NeighborAddress}
    staticRpfRule.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", staticRpfRule.InterfaceName}
    return &(staticRpfRule.EntityData)
}

// Mfwd_Vrfs_Vrf_Ipv4_Interfaces
// Interface-level Configuration
type Mfwd_Vrfs_Vrf_Ipv4_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The name of the interface. The type is slice of
    // Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface_.
    Interface_ []Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface
}

func (interfaces *Mfwd_Vrfs_Vrf_Ipv4_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "ipv4"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface
// The name of the interface
type Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
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

func (self *Mfwd_Vrfs_Vrf_Ipv4_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["ttl-threshold"] = types.YLeaf{"TtlThreshold", self.TtlThreshold}
    self.EntityData.Leafs["enable-on-interface"] = types.YLeaf{"EnableOnInterface", self.EnableOnInterface}
    self.EntityData.Leafs["boundary"] = types.YLeaf{"Boundary", self.Boundary}
    return &(self.EntityData)
}

