// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-bfd package configuration.
// 
// This module contains definitions
// for the following management objects:
//   bfd: BFD Configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-snmp-agent-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ip_bfd_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_bfd_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-bfd-cfg bfd}", reflect.TypeOf(Bfd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-bfd-cfg:bfd", reflect.TypeOf(Bfd{}))
}

// BfdEchoStartupValidate represents Bfd echo startup validate
type BfdEchoStartupValidate string

const (
    // Disable echo startup validation
    BfdEchoStartupValidate_off BfdEchoStartupValidate = "off"

    // Enable echo startup validation
    BfdEchoStartupValidate_on BfdEchoStartupValidate = "on"

    // Force echo startup validation
    BfdEchoStartupValidate_force BfdEchoStartupValidate = "force"
)

// BfdIfIpv6ChecksumUsage represents Bfd if ipv6 checksum usage
type BfdIfIpv6ChecksumUsage string

const (
    // Disable IPv6 checksum
    BfdIfIpv6ChecksumUsage_disable BfdIfIpv6ChecksumUsage = "disable"

    // Enable IPv6 checksum
    BfdIfIpv6ChecksumUsage_enable BfdIfIpv6ChecksumUsage = "enable"
)

// BfdIfEchoUsage represents Bfd if echo usage
type BfdIfEchoUsage string

const (
    // Enable echo
    BfdIfEchoUsage_enable BfdIfEchoUsage = "enable"

    // Disable echo
    BfdIfEchoUsage_disable BfdIfEchoUsage = "disable"
)

// BfdBundleCoexistenceBobBlb represents Bfd bundle coexistence bob blb
type BfdBundleCoexistenceBobBlb string

const (
    // Inherited coexistence mode
    BfdBundleCoexistenceBobBlb_inherited BfdBundleCoexistenceBobBlb = "inherited"

    // Logical coexistence mode
    BfdBundleCoexistenceBobBlb_logical BfdBundleCoexistenceBobBlb = "logical"
)

// Bfd
// BFD Configuration
type Bfd struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Echo configuration. The type is interface{}.
    GlobalEchoUsage interface{}

    // To disable IPv6 checksum configuration. The type is interface{}.
    Ipv6ChecksumDisable interface{}

    // Configure echo min-interval for bundle interface. The type is interface{}
    // with range: 15..2000. Units are millisecond. The default value is 15.
    GlobalEchoMinInterval interface{}

    // Multihop TTL Drop Threshold. The type is interface{} with range: 0..254.
    TtlDropThreshold interface{}

    // Single hop trap configuration. The type is interface{}.
    SingleHopTrap interface{}

    // IPv4 echo source address configuration. The type is one of the following
    // types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    GlobalIpv4EchoSource interface{}

    // Flapping class container.
    FlapDamp Bfd_FlapDamp

    // BFD echo latency feature class container.
    EchoLatency Bfd_EchoLatency

    // BFD echo startup feature class container.
    EchoStartup Bfd_EchoStartup

    // Interface configuration.
    Interfaces Bfd_Interfaces

    // Multipath Include configuration.
    MultiPathIncludes Bfd_MultiPathIncludes

    // Configuration related to BFD over Bundle.
    Bundle Bfd_Bundle
}

func (bfd *Bfd) GetFilter() yfilter.YFilter { return bfd.YFilter }

func (bfd *Bfd) SetFilter(yf yfilter.YFilter) { bfd.YFilter = yf }

func (bfd *Bfd) GetGoName(yname string) string {
    if yname == "global-echo-usage" { return "GlobalEchoUsage" }
    if yname == "ipv6-checksum-disable" { return "Ipv6ChecksumDisable" }
    if yname == "global-echo-min-interval" { return "GlobalEchoMinInterval" }
    if yname == "ttl-drop-threshold" { return "TtlDropThreshold" }
    if yname == "single-hop-trap" { return "SingleHopTrap" }
    if yname == "global-ipv4-echo-source" { return "GlobalIpv4EchoSource" }
    if yname == "flap-damp" { return "FlapDamp" }
    if yname == "echo-latency" { return "EchoLatency" }
    if yname == "echo-startup" { return "EchoStartup" }
    if yname == "interfaces" { return "Interfaces" }
    if yname == "multi-path-includes" { return "MultiPathIncludes" }
    if yname == "bundle" { return "Bundle" }
    return ""
}

func (bfd *Bfd) GetSegmentPath() string {
    return "Cisco-IOS-XR-ip-bfd-cfg:bfd"
}

func (bfd *Bfd) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "flap-damp" {
        return &bfd.FlapDamp
    }
    if childYangName == "echo-latency" {
        return &bfd.EchoLatency
    }
    if childYangName == "echo-startup" {
        return &bfd.EchoStartup
    }
    if childYangName == "interfaces" {
        return &bfd.Interfaces
    }
    if childYangName == "multi-path-includes" {
        return &bfd.MultiPathIncludes
    }
    if childYangName == "bundle" {
        return &bfd.Bundle
    }
    return nil
}

func (bfd *Bfd) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["flap-damp"] = &bfd.FlapDamp
    children["echo-latency"] = &bfd.EchoLatency
    children["echo-startup"] = &bfd.EchoStartup
    children["interfaces"] = &bfd.Interfaces
    children["multi-path-includes"] = &bfd.MultiPathIncludes
    children["bundle"] = &bfd.Bundle
    return children
}

func (bfd *Bfd) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["global-echo-usage"] = bfd.GlobalEchoUsage
    leafs["ipv6-checksum-disable"] = bfd.Ipv6ChecksumDisable
    leafs["global-echo-min-interval"] = bfd.GlobalEchoMinInterval
    leafs["ttl-drop-threshold"] = bfd.TtlDropThreshold
    leafs["single-hop-trap"] = bfd.SingleHopTrap
    leafs["global-ipv4-echo-source"] = bfd.GlobalIpv4EchoSource
    return leafs
}

func (bfd *Bfd) GetBundleName() string { return "cisco_ios_xr" }

func (bfd *Bfd) GetYangName() string { return "bfd" }

func (bfd *Bfd) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bfd *Bfd) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bfd *Bfd) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bfd *Bfd) SetParent(parent types.Entity) { bfd.parent = parent }

func (bfd *Bfd) GetParent() types.Entity { return bfd.parent }

func (bfd *Bfd) GetParentYangName() string { return "Cisco-IOS-XR-ip-bfd-cfg" }

// Bfd_FlapDamp
// Flapping class container
type Bfd_FlapDamp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Stability threshold to enable dampening. The type is interface{} with
    // range: 60000..3600000. Units are millisecond. The default value is 120000.
    Threshold interface{}

    // Initial delay before bringing up session. The type is interface{} with
    // range: 1..3600000. Units are millisecond. The default value is 2000.
    InitialDelay interface{}

    // Maximum delay before bringing up session. The type is interface{} with
    // range: 1..3600000. Units are millisecond. The default value is 120000.
    MaximumDelay interface{}

    // Dampening Enable/Disable. The type is interface{}.
    DampenDisable interface{}

    // Secondary delay before bringing up session. The type is interface{} with
    // range: 1..3600000. Units are millisecond. The default value is 5000.
    SecondaryDelay interface{}

    // Bundle per member feature class container.
    BundleMember Bfd_FlapDamp_BundleMember

    // Extensions to the BFD dampening feature.
    Extensions Bfd_FlapDamp_Extensions
}

func (flapDamp *Bfd_FlapDamp) GetFilter() yfilter.YFilter { return flapDamp.YFilter }

func (flapDamp *Bfd_FlapDamp) SetFilter(yf yfilter.YFilter) { flapDamp.YFilter = yf }

func (flapDamp *Bfd_FlapDamp) GetGoName(yname string) string {
    if yname == "threshold" { return "Threshold" }
    if yname == "initial-delay" { return "InitialDelay" }
    if yname == "maximum-delay" { return "MaximumDelay" }
    if yname == "dampen-disable" { return "DampenDisable" }
    if yname == "secondary-delay" { return "SecondaryDelay" }
    if yname == "bundle-member" { return "BundleMember" }
    if yname == "extensions" { return "Extensions" }
    return ""
}

func (flapDamp *Bfd_FlapDamp) GetSegmentPath() string {
    return "flap-damp"
}

func (flapDamp *Bfd_FlapDamp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "bundle-member" {
        return &flapDamp.BundleMember
    }
    if childYangName == "extensions" {
        return &flapDamp.Extensions
    }
    return nil
}

func (flapDamp *Bfd_FlapDamp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["bundle-member"] = &flapDamp.BundleMember
    children["extensions"] = &flapDamp.Extensions
    return children
}

func (flapDamp *Bfd_FlapDamp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["threshold"] = flapDamp.Threshold
    leafs["initial-delay"] = flapDamp.InitialDelay
    leafs["maximum-delay"] = flapDamp.MaximumDelay
    leafs["dampen-disable"] = flapDamp.DampenDisable
    leafs["secondary-delay"] = flapDamp.SecondaryDelay
    return leafs
}

func (flapDamp *Bfd_FlapDamp) GetBundleName() string { return "cisco_ios_xr" }

func (flapDamp *Bfd_FlapDamp) GetYangName() string { return "flap-damp" }

func (flapDamp *Bfd_FlapDamp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (flapDamp *Bfd_FlapDamp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (flapDamp *Bfd_FlapDamp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (flapDamp *Bfd_FlapDamp) SetParent(parent types.Entity) { flapDamp.parent = parent }

func (flapDamp *Bfd_FlapDamp) GetParent() types.Entity { return flapDamp.parent }

func (flapDamp *Bfd_FlapDamp) GetParentYangName() string { return "bfd" }

// Bfd_FlapDamp_BundleMember
// Bundle per member feature class container
type Bfd_FlapDamp_BundleMember struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Initial delay before bringing up session. The type is interface{} with
    // range: 1..518400000. Units are millisecond. The default value is 16000.
    InitialDelay interface{}

    // Maximum delay before bringing up session. The type is interface{} with
    // range: 1..518400000. Units are millisecond. The default value is 600000.
    MaximumDelay interface{}

    // Secondary delay before bringing up session. The type is interface{} with
    // range: 1..518400000. Units are millisecond. The default value is 20000.
    SecondaryDelay interface{}

    // Apply immediate dampening and only when failure is L3 specific. The type is
    // interface{}.
    L3OnlyMode interface{}
}

func (bundleMember *Bfd_FlapDamp_BundleMember) GetFilter() yfilter.YFilter { return bundleMember.YFilter }

func (bundleMember *Bfd_FlapDamp_BundleMember) SetFilter(yf yfilter.YFilter) { bundleMember.YFilter = yf }

func (bundleMember *Bfd_FlapDamp_BundleMember) GetGoName(yname string) string {
    if yname == "initial-delay" { return "InitialDelay" }
    if yname == "maximum-delay" { return "MaximumDelay" }
    if yname == "secondary-delay" { return "SecondaryDelay" }
    if yname == "l3-only-mode" { return "L3OnlyMode" }
    return ""
}

func (bundleMember *Bfd_FlapDamp_BundleMember) GetSegmentPath() string {
    return "bundle-member"
}

func (bundleMember *Bfd_FlapDamp_BundleMember) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (bundleMember *Bfd_FlapDamp_BundleMember) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (bundleMember *Bfd_FlapDamp_BundleMember) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["initial-delay"] = bundleMember.InitialDelay
    leafs["maximum-delay"] = bundleMember.MaximumDelay
    leafs["secondary-delay"] = bundleMember.SecondaryDelay
    leafs["l3-only-mode"] = bundleMember.L3OnlyMode
    return leafs
}

func (bundleMember *Bfd_FlapDamp_BundleMember) GetBundleName() string { return "cisco_ios_xr" }

func (bundleMember *Bfd_FlapDamp_BundleMember) GetYangName() string { return "bundle-member" }

func (bundleMember *Bfd_FlapDamp_BundleMember) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bundleMember *Bfd_FlapDamp_BundleMember) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bundleMember *Bfd_FlapDamp_BundleMember) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bundleMember *Bfd_FlapDamp_BundleMember) SetParent(parent types.Entity) { bundleMember.parent = parent }

func (bundleMember *Bfd_FlapDamp_BundleMember) GetParent() types.Entity { return bundleMember.parent }

func (bundleMember *Bfd_FlapDamp_BundleMember) GetParentYangName() string { return "flap-damp" }

// Bfd_FlapDamp_Extensions
// Extensions to the BFD dampening feature
type Bfd_FlapDamp_Extensions struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If set, DOWN state monitoring mode is enabled. The type is interface{}.
    DownMonitor interface{}
}

func (extensions *Bfd_FlapDamp_Extensions) GetFilter() yfilter.YFilter { return extensions.YFilter }

func (extensions *Bfd_FlapDamp_Extensions) SetFilter(yf yfilter.YFilter) { extensions.YFilter = yf }

func (extensions *Bfd_FlapDamp_Extensions) GetGoName(yname string) string {
    if yname == "down-monitor" { return "DownMonitor" }
    return ""
}

func (extensions *Bfd_FlapDamp_Extensions) GetSegmentPath() string {
    return "extensions"
}

func (extensions *Bfd_FlapDamp_Extensions) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (extensions *Bfd_FlapDamp_Extensions) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (extensions *Bfd_FlapDamp_Extensions) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["down-monitor"] = extensions.DownMonitor
    return leafs
}

func (extensions *Bfd_FlapDamp_Extensions) GetBundleName() string { return "cisco_ios_xr" }

func (extensions *Bfd_FlapDamp_Extensions) GetYangName() string { return "extensions" }

func (extensions *Bfd_FlapDamp_Extensions) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (extensions *Bfd_FlapDamp_Extensions) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (extensions *Bfd_FlapDamp_Extensions) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (extensions *Bfd_FlapDamp_Extensions) SetParent(parent types.Entity) { extensions.parent = parent }

func (extensions *Bfd_FlapDamp_Extensions) GetParent() types.Entity { return extensions.parent }

func (extensions *Bfd_FlapDamp_Extensions) GetParentYangName() string { return "flap-damp" }

// Bfd_EchoLatency
// BFD echo latency feature class container
type Bfd_EchoLatency struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BFD echo latency detection.
    Detect Bfd_EchoLatency_Detect
}

func (echoLatency *Bfd_EchoLatency) GetFilter() yfilter.YFilter { return echoLatency.YFilter }

func (echoLatency *Bfd_EchoLatency) SetFilter(yf yfilter.YFilter) { echoLatency.YFilter = yf }

func (echoLatency *Bfd_EchoLatency) GetGoName(yname string) string {
    if yname == "detect" { return "Detect" }
    return ""
}

func (echoLatency *Bfd_EchoLatency) GetSegmentPath() string {
    return "echo-latency"
}

func (echoLatency *Bfd_EchoLatency) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "detect" {
        return &echoLatency.Detect
    }
    return nil
}

func (echoLatency *Bfd_EchoLatency) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["detect"] = &echoLatency.Detect
    return children
}

func (echoLatency *Bfd_EchoLatency) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (echoLatency *Bfd_EchoLatency) GetBundleName() string { return "cisco_ios_xr" }

func (echoLatency *Bfd_EchoLatency) GetYangName() string { return "echo-latency" }

func (echoLatency *Bfd_EchoLatency) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (echoLatency *Bfd_EchoLatency) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (echoLatency *Bfd_EchoLatency) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (echoLatency *Bfd_EchoLatency) SetParent(parent types.Entity) { echoLatency.parent = parent }

func (echoLatency *Bfd_EchoLatency) GetParent() types.Entity { return echoLatency.parent }

func (echoLatency *Bfd_EchoLatency) GetParentYangName() string { return "bfd" }

// Bfd_EchoLatency_Detect
// BFD echo latency detection
type Bfd_EchoLatency_Detect struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // If set, echo latency detect is enabled. The type is bool. The default value
    // is false.
    LatencyDetectEnabled interface{}

    // Echo latency detect percentage. The type is interface{} with range:
    // 100..250. Units are percentage. The default value is 100.
    LatencyDetectPercentage interface{}

    // Echo latency detect count. The type is interface{} with range: 1..10. The
    // default value is 1.
    LatencyDetectCount interface{}
}

func (detect *Bfd_EchoLatency_Detect) GetFilter() yfilter.YFilter { return detect.YFilter }

func (detect *Bfd_EchoLatency_Detect) SetFilter(yf yfilter.YFilter) { detect.YFilter = yf }

func (detect *Bfd_EchoLatency_Detect) GetGoName(yname string) string {
    if yname == "latency-detect-enabled" { return "LatencyDetectEnabled" }
    if yname == "latency-detect-percentage" { return "LatencyDetectPercentage" }
    if yname == "latency-detect-count" { return "LatencyDetectCount" }
    return ""
}

func (detect *Bfd_EchoLatency_Detect) GetSegmentPath() string {
    return "detect"
}

func (detect *Bfd_EchoLatency_Detect) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (detect *Bfd_EchoLatency_Detect) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (detect *Bfd_EchoLatency_Detect) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["latency-detect-enabled"] = detect.LatencyDetectEnabled
    leafs["latency-detect-percentage"] = detect.LatencyDetectPercentage
    leafs["latency-detect-count"] = detect.LatencyDetectCount
    return leafs
}

func (detect *Bfd_EchoLatency_Detect) GetBundleName() string { return "cisco_ios_xr" }

func (detect *Bfd_EchoLatency_Detect) GetYangName() string { return "detect" }

func (detect *Bfd_EchoLatency_Detect) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (detect *Bfd_EchoLatency_Detect) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (detect *Bfd_EchoLatency_Detect) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (detect *Bfd_EchoLatency_Detect) SetParent(parent types.Entity) { detect.parent = parent }

func (detect *Bfd_EchoLatency_Detect) GetParent() types.Entity { return detect.parent }

func (detect *Bfd_EchoLatency_Detect) GetParentYangName() string { return "echo-latency" }

// Bfd_EchoStartup
// BFD echo startup feature class container
type Bfd_EchoStartup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // BFD echo validation prior to session bringup. The type is
    // BfdEchoStartupValidate. The default value is off.
    Validate interface{}
}

func (echoStartup *Bfd_EchoStartup) GetFilter() yfilter.YFilter { return echoStartup.YFilter }

func (echoStartup *Bfd_EchoStartup) SetFilter(yf yfilter.YFilter) { echoStartup.YFilter = yf }

func (echoStartup *Bfd_EchoStartup) GetGoName(yname string) string {
    if yname == "validate" { return "Validate" }
    return ""
}

func (echoStartup *Bfd_EchoStartup) GetSegmentPath() string {
    return "echo-startup"
}

func (echoStartup *Bfd_EchoStartup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (echoStartup *Bfd_EchoStartup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (echoStartup *Bfd_EchoStartup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["validate"] = echoStartup.Validate
    return leafs
}

func (echoStartup *Bfd_EchoStartup) GetBundleName() string { return "cisco_ios_xr" }

func (echoStartup *Bfd_EchoStartup) GetYangName() string { return "echo-startup" }

func (echoStartup *Bfd_EchoStartup) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (echoStartup *Bfd_EchoStartup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (echoStartup *Bfd_EchoStartup) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (echoStartup *Bfd_EchoStartup) SetParent(parent types.Entity) { echoStartup.parent = parent }

func (echoStartup *Bfd_EchoStartup) GetParent() types.Entity { return echoStartup.parent }

func (echoStartup *Bfd_EchoStartup) GetParentYangName() string { return "bfd" }

// Bfd_Interfaces
// Interface configuration
type Bfd_Interfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Single interface configuration. The type is slice of
    // Bfd_Interfaces_Interface.
    Interface []Bfd_Interfaces_Interface
}

func (interfaces *Bfd_Interfaces) GetFilter() yfilter.YFilter { return interfaces.YFilter }

func (interfaces *Bfd_Interfaces) SetFilter(yf yfilter.YFilter) { interfaces.YFilter = yf }

func (interfaces *Bfd_Interfaces) GetGoName(yname string) string {
    if yname == "interface" { return "Interface" }
    return ""
}

func (interfaces *Bfd_Interfaces) GetSegmentPath() string {
    return "interfaces"
}

func (interfaces *Bfd_Interfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "interface" {
        for _, c := range interfaces.Interface {
            if interfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Bfd_Interfaces_Interface{}
        interfaces.Interface = append(interfaces.Interface, child)
        return &interfaces.Interface[len(interfaces.Interface)-1]
    }
    return nil
}

func (interfaces *Bfd_Interfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range interfaces.Interface {
        children[interfaces.Interface[i].GetSegmentPath()] = &interfaces.Interface[i]
    }
    return children
}

func (interfaces *Bfd_Interfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (interfaces *Bfd_Interfaces) GetBundleName() string { return "cisco_ios_xr" }

func (interfaces *Bfd_Interfaces) GetYangName() string { return "interfaces" }

func (interfaces *Bfd_Interfaces) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (interfaces *Bfd_Interfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (interfaces *Bfd_Interfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (interfaces *Bfd_Interfaces) SetParent(parent types.Entity) { interfaces.parent = parent }

func (interfaces *Bfd_Interfaces) GetParent() types.Entity { return interfaces.parent }

func (interfaces *Bfd_Interfaces) GetParentYangName() string { return "bfd" }

// Bfd_Interfaces_Interface
// Single interface configuration
type Bfd_Interfaces_Interface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Interface Name. The type is string with pattern:
    // [a-zA-Z0-9./-]+.
    InterfaceName interface{}

    // Echo usage for this interface. The type is BfdIfEchoUsage. The default
    // value is enable.
    InterfaceEchoUsage interface{}

    // IPv6 checksum usage for this interface - Interface config will always take
    // precedence over global config. The type is BfdIfIpv6ChecksumUsage. The
    // default value is enable.
    Ipv6Checksum interface{}

    // Interface IPv4 echo source address configuration. The type is one of the
    // following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    InterfaceIpv4EchoSource interface{}
}

func (self *Bfd_Interfaces_Interface) GetFilter() yfilter.YFilter { return self.YFilter }

func (self *Bfd_Interfaces_Interface) SetFilter(yf yfilter.YFilter) { self.YFilter = yf }

func (self *Bfd_Interfaces_Interface) GetGoName(yname string) string {
    if yname == "interface-name" { return "InterfaceName" }
    if yname == "interface-echo-usage" { return "InterfaceEchoUsage" }
    if yname == "ipv6-checksum" { return "Ipv6Checksum" }
    if yname == "interface-ipv4-echo-source" { return "InterfaceIpv4EchoSource" }
    return ""
}

func (self *Bfd_Interfaces_Interface) GetSegmentPath() string {
    return "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
}

func (self *Bfd_Interfaces_Interface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (self *Bfd_Interfaces_Interface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (self *Bfd_Interfaces_Interface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["interface-name"] = self.InterfaceName
    leafs["interface-echo-usage"] = self.InterfaceEchoUsage
    leafs["ipv6-checksum"] = self.Ipv6Checksum
    leafs["interface-ipv4-echo-source"] = self.InterfaceIpv4EchoSource
    return leafs
}

func (self *Bfd_Interfaces_Interface) GetBundleName() string { return "cisco_ios_xr" }

func (self *Bfd_Interfaces_Interface) GetYangName() string { return "interface" }

func (self *Bfd_Interfaces_Interface) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (self *Bfd_Interfaces_Interface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (self *Bfd_Interfaces_Interface) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (self *Bfd_Interfaces_Interface) SetParent(parent types.Entity) { self.parent = parent }

func (self *Bfd_Interfaces_Interface) GetParent() types.Entity { return self.parent }

func (self *Bfd_Interfaces_Interface) GetParentYangName() string { return "interfaces" }

// Bfd_MultiPathIncludes
// Multipath Include configuration
type Bfd_MultiPathIncludes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Location configuration. The type is slice of
    // Bfd_MultiPathIncludes_MultiPathInclude.
    MultiPathInclude []Bfd_MultiPathIncludes_MultiPathInclude
}

func (multiPathIncludes *Bfd_MultiPathIncludes) GetFilter() yfilter.YFilter { return multiPathIncludes.YFilter }

func (multiPathIncludes *Bfd_MultiPathIncludes) SetFilter(yf yfilter.YFilter) { multiPathIncludes.YFilter = yf }

func (multiPathIncludes *Bfd_MultiPathIncludes) GetGoName(yname string) string {
    if yname == "multi-path-include" { return "MultiPathInclude" }
    return ""
}

func (multiPathIncludes *Bfd_MultiPathIncludes) GetSegmentPath() string {
    return "multi-path-includes"
}

func (multiPathIncludes *Bfd_MultiPathIncludes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "multi-path-include" {
        for _, c := range multiPathIncludes.MultiPathInclude {
            if multiPathIncludes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Bfd_MultiPathIncludes_MultiPathInclude{}
        multiPathIncludes.MultiPathInclude = append(multiPathIncludes.MultiPathInclude, child)
        return &multiPathIncludes.MultiPathInclude[len(multiPathIncludes.MultiPathInclude)-1]
    }
    return nil
}

func (multiPathIncludes *Bfd_MultiPathIncludes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range multiPathIncludes.MultiPathInclude {
        children[multiPathIncludes.MultiPathInclude[i].GetSegmentPath()] = &multiPathIncludes.MultiPathInclude[i]
    }
    return children
}

func (multiPathIncludes *Bfd_MultiPathIncludes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (multiPathIncludes *Bfd_MultiPathIncludes) GetBundleName() string { return "cisco_ios_xr" }

func (multiPathIncludes *Bfd_MultiPathIncludes) GetYangName() string { return "multi-path-includes" }

func (multiPathIncludes *Bfd_MultiPathIncludes) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multiPathIncludes *Bfd_MultiPathIncludes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multiPathIncludes *Bfd_MultiPathIncludes) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multiPathIncludes *Bfd_MultiPathIncludes) SetParent(parent types.Entity) { multiPathIncludes.parent = parent }

func (multiPathIncludes *Bfd_MultiPathIncludes) GetParent() types.Entity { return multiPathIncludes.parent }

func (multiPathIncludes *Bfd_MultiPathIncludes) GetParentYangName() string { return "bfd" }

// Bfd_MultiPathIncludes_MultiPathInclude
// Location configuration
type Bfd_MultiPathIncludes_MultiPathInclude struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Location. The type is string with pattern:
    // ([a-zA-Z0-9_]*\d+/){1,2}([a-zA-Z0-9_]*\d+).
    Location interface{}
}

func (multiPathInclude *Bfd_MultiPathIncludes_MultiPathInclude) GetFilter() yfilter.YFilter { return multiPathInclude.YFilter }

func (multiPathInclude *Bfd_MultiPathIncludes_MultiPathInclude) SetFilter(yf yfilter.YFilter) { multiPathInclude.YFilter = yf }

func (multiPathInclude *Bfd_MultiPathIncludes_MultiPathInclude) GetGoName(yname string) string {
    if yname == "location" { return "Location" }
    return ""
}

func (multiPathInclude *Bfd_MultiPathIncludes_MultiPathInclude) GetSegmentPath() string {
    return "multi-path-include" + "[location='" + fmt.Sprintf("%v", multiPathInclude.Location) + "']"
}

func (multiPathInclude *Bfd_MultiPathIncludes_MultiPathInclude) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (multiPathInclude *Bfd_MultiPathIncludes_MultiPathInclude) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (multiPathInclude *Bfd_MultiPathIncludes_MultiPathInclude) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["location"] = multiPathInclude.Location
    return leafs
}

func (multiPathInclude *Bfd_MultiPathIncludes_MultiPathInclude) GetBundleName() string { return "cisco_ios_xr" }

func (multiPathInclude *Bfd_MultiPathIncludes_MultiPathInclude) GetYangName() string { return "multi-path-include" }

func (multiPathInclude *Bfd_MultiPathIncludes_MultiPathInclude) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (multiPathInclude *Bfd_MultiPathIncludes_MultiPathInclude) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (multiPathInclude *Bfd_MultiPathIncludes_MultiPathInclude) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (multiPathInclude *Bfd_MultiPathIncludes_MultiPathInclude) SetParent(parent types.Entity) { multiPathInclude.parent = parent }

func (multiPathInclude *Bfd_MultiPathIncludes_MultiPathInclude) GetParent() types.Entity { return multiPathInclude.parent }

func (multiPathInclude *Bfd_MultiPathIncludes_MultiPathInclude) GetParentYangName() string { return "multi-path-includes" }

// Bfd_Bundle
// Configuration related to BFD over Bundle
type Bfd_Bundle struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Coexistence mode for BFD on bundle feature.
    Coexistence Bfd_Bundle_Coexistence
}

func (bundle *Bfd_Bundle) GetFilter() yfilter.YFilter { return bundle.YFilter }

func (bundle *Bfd_Bundle) SetFilter(yf yfilter.YFilter) { bundle.YFilter = yf }

func (bundle *Bfd_Bundle) GetGoName(yname string) string {
    if yname == "coexistence" { return "Coexistence" }
    return ""
}

func (bundle *Bfd_Bundle) GetSegmentPath() string {
    return "bundle"
}

func (bundle *Bfd_Bundle) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "coexistence" {
        return &bundle.Coexistence
    }
    return nil
}

func (bundle *Bfd_Bundle) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["coexistence"] = &bundle.Coexistence
    return children
}

func (bundle *Bfd_Bundle) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (bundle *Bfd_Bundle) GetBundleName() string { return "cisco_ios_xr" }

func (bundle *Bfd_Bundle) GetYangName() string { return "bundle" }

func (bundle *Bfd_Bundle) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (bundle *Bfd_Bundle) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (bundle *Bfd_Bundle) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (bundle *Bfd_Bundle) SetParent(parent types.Entity) { bundle.parent = parent }

func (bundle *Bfd_Bundle) GetParent() types.Entity { return bundle.parent }

func (bundle *Bfd_Bundle) GetParentYangName() string { return "bfd" }

// Bfd_Bundle_Coexistence
// Coexistence mode for BFD on bundle feature
type Bfd_Bundle_Coexistence struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Coexistence mode for BoB and BLB feature. The type is
    // BfdBundleCoexistenceBobBlb.
    BobBlb interface{}
}

func (coexistence *Bfd_Bundle_Coexistence) GetFilter() yfilter.YFilter { return coexistence.YFilter }

func (coexistence *Bfd_Bundle_Coexistence) SetFilter(yf yfilter.YFilter) { coexistence.YFilter = yf }

func (coexistence *Bfd_Bundle_Coexistence) GetGoName(yname string) string {
    if yname == "bob-blb" { return "BobBlb" }
    return ""
}

func (coexistence *Bfd_Bundle_Coexistence) GetSegmentPath() string {
    return "coexistence"
}

func (coexistence *Bfd_Bundle_Coexistence) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (coexistence *Bfd_Bundle_Coexistence) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (coexistence *Bfd_Bundle_Coexistence) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["bob-blb"] = coexistence.BobBlb
    return leafs
}

func (coexistence *Bfd_Bundle_Coexistence) GetBundleName() string { return "cisco_ios_xr" }

func (coexistence *Bfd_Bundle_Coexistence) GetYangName() string { return "coexistence" }

func (coexistence *Bfd_Bundle_Coexistence) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (coexistence *Bfd_Bundle_Coexistence) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (coexistence *Bfd_Bundle_Coexistence) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (coexistence *Bfd_Bundle_Coexistence) SetParent(parent types.Entity) { coexistence.parent = parent }

func (coexistence *Bfd_Bundle_Coexistence) GetParent() types.Entity { return coexistence.parent }

func (coexistence *Bfd_Bundle_Coexistence) GetParentYangName() string { return "bundle" }

