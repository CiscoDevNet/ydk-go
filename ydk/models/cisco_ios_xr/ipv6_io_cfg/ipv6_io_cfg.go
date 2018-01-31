// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-io package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ipv6-configuration: IPv6 Configuration Data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_io_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_io_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-io-cfg ipv6-configuration}", reflect.TypeOf(Ipv6Configuration{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-io-cfg:ipv6-configuration", reflect.TypeOf(Ipv6Configuration{}))
}

// Ipv6Configuration
// IPv6 Configuration Data
type Ipv6Configuration struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure IPv6 Path MTU timeout value in minutes. The type is interface{}
    // with range: 1..15. Units are minute.
    Ipv6PmtuTimeOut interface{}

    // TRUE if enabled, FALSE if disabled. The type is bool. The default value is
    // true.
    Ipv6SourceRoute interface{}

    // TRUE if enabled, FALSE if disabled. The type is bool. The default value is
    // false.
    Ipv6PmtuEnable interface{}

    // Configure IPv6 hop count limit. The type is interface{} with range: 1..255.
    Ipv6HopLimit interface{}

    // IPv6 fragmented packet assembler.
    Ipv6Assembler Ipv6Configuration_Ipv6Assembler

    // Configure IPv6 ICMP parameters.
    Ipv6Icmp Ipv6Configuration_Ipv6Icmp
}

func (ipv6Configuration *Ipv6Configuration) GetFilter() yfilter.YFilter { return ipv6Configuration.YFilter }

func (ipv6Configuration *Ipv6Configuration) SetFilter(yf yfilter.YFilter) { ipv6Configuration.YFilter = yf }

func (ipv6Configuration *Ipv6Configuration) GetGoName(yname string) string {
    if yname == "ipv6-pmtu-time-out" { return "Ipv6PmtuTimeOut" }
    if yname == "ipv6-source-route" { return "Ipv6SourceRoute" }
    if yname == "ipv6-pmtu-enable" { return "Ipv6PmtuEnable" }
    if yname == "ipv6-hop-limit" { return "Ipv6HopLimit" }
    if yname == "ipv6-assembler" { return "Ipv6Assembler" }
    if yname == "ipv6icmp" { return "Ipv6Icmp" }
    return ""
}

func (ipv6Configuration *Ipv6Configuration) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv6-io-cfg:ipv6-configuration"
}

func (ipv6Configuration *Ipv6Configuration) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipv6-assembler" {
        return &ipv6Configuration.Ipv6Assembler
    }
    if childYangName == "ipv6icmp" {
        return &ipv6Configuration.Ipv6Icmp
    }
    return nil
}

func (ipv6Configuration *Ipv6Configuration) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ipv6-assembler"] = &ipv6Configuration.Ipv6Assembler
    children["ipv6icmp"] = &ipv6Configuration.Ipv6Icmp
    return children
}

func (ipv6Configuration *Ipv6Configuration) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipv6-pmtu-time-out"] = ipv6Configuration.Ipv6PmtuTimeOut
    leafs["ipv6-source-route"] = ipv6Configuration.Ipv6SourceRoute
    leafs["ipv6-pmtu-enable"] = ipv6Configuration.Ipv6PmtuEnable
    leafs["ipv6-hop-limit"] = ipv6Configuration.Ipv6HopLimit
    return leafs
}

func (ipv6Configuration *Ipv6Configuration) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Configuration *Ipv6Configuration) GetYangName() string { return "ipv6-configuration" }

func (ipv6Configuration *Ipv6Configuration) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Configuration *Ipv6Configuration) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Configuration *Ipv6Configuration) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Configuration *Ipv6Configuration) SetParent(parent types.Entity) { ipv6Configuration.parent = parent }

func (ipv6Configuration *Ipv6Configuration) GetParent() types.Entity { return ipv6Configuration.parent }

func (ipv6Configuration *Ipv6Configuration) GetParentYangName() string { return "Cisco-IOS-XR-ipv6-io-cfg" }

// Ipv6Configuration_Ipv6Assembler
// IPv6 fragmented packet assembler
type Ipv6Configuration_Ipv6Assembler struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of seconds an assembly queue will hold before timeout. The type is
    // interface{} with range: 1..120. Units are second.
    Timeout interface{}

    // Maxinum packets allowed in assembly queues (in percent). The type is
    // interface{} with range: 1..50. Units are percentage.
    MaxPackets interface{}
}

func (ipv6Assembler *Ipv6Configuration_Ipv6Assembler) GetFilter() yfilter.YFilter { return ipv6Assembler.YFilter }

func (ipv6Assembler *Ipv6Configuration_Ipv6Assembler) SetFilter(yf yfilter.YFilter) { ipv6Assembler.YFilter = yf }

func (ipv6Assembler *Ipv6Configuration_Ipv6Assembler) GetGoName(yname string) string {
    if yname == "timeout" { return "Timeout" }
    if yname == "max-packets" { return "MaxPackets" }
    return ""
}

func (ipv6Assembler *Ipv6Configuration_Ipv6Assembler) GetSegmentPath() string {
    return "ipv6-assembler"
}

func (ipv6Assembler *Ipv6Configuration_Ipv6Assembler) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Assembler *Ipv6Configuration_Ipv6Assembler) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Assembler *Ipv6Configuration_Ipv6Assembler) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["timeout"] = ipv6Assembler.Timeout
    leafs["max-packets"] = ipv6Assembler.MaxPackets
    return leafs
}

func (ipv6Assembler *Ipv6Configuration_Ipv6Assembler) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Assembler *Ipv6Configuration_Ipv6Assembler) GetYangName() string { return "ipv6-assembler" }

func (ipv6Assembler *Ipv6Configuration_Ipv6Assembler) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Assembler *Ipv6Configuration_Ipv6Assembler) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Assembler *Ipv6Configuration_Ipv6Assembler) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Assembler *Ipv6Configuration_Ipv6Assembler) SetParent(parent types.Entity) { ipv6Assembler.parent = parent }

func (ipv6Assembler *Ipv6Configuration_Ipv6Assembler) GetParent() types.Entity { return ipv6Assembler.parent }

func (ipv6Assembler *Ipv6Configuration_Ipv6Assembler) GetParentYangName() string { return "ipv6-configuration" }

// Ipv6Configuration_Ipv6Icmp
// Configure IPv6 ICMP parameters
// This type is a presence type.
type Ipv6Configuration_Ipv6Icmp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Interval between tokens in milliseconds. The type is interface{} with
    // range: 0..2147483647. This attribute is mandatory. Units are millisecond.
    ErrorInterval interface{}

    // Bucket size. The type is interface{} with range: 1..200. The default value
    // is 10.
    BucketSize interface{}
}

func (ipv6Icmp *Ipv6Configuration_Ipv6Icmp) GetFilter() yfilter.YFilter { return ipv6Icmp.YFilter }

func (ipv6Icmp *Ipv6Configuration_Ipv6Icmp) SetFilter(yf yfilter.YFilter) { ipv6Icmp.YFilter = yf }

func (ipv6Icmp *Ipv6Configuration_Ipv6Icmp) GetGoName(yname string) string {
    if yname == "error-interval" { return "ErrorInterval" }
    if yname == "bucket-size" { return "BucketSize" }
    return ""
}

func (ipv6Icmp *Ipv6Configuration_Ipv6Icmp) GetSegmentPath() string {
    return "ipv6icmp"
}

func (ipv6Icmp *Ipv6Configuration_Ipv6Icmp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipv6Icmp *Ipv6Configuration_Ipv6Icmp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipv6Icmp *Ipv6Configuration_Ipv6Icmp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["error-interval"] = ipv6Icmp.ErrorInterval
    leafs["bucket-size"] = ipv6Icmp.BucketSize
    return leafs
}

func (ipv6Icmp *Ipv6Configuration_Ipv6Icmp) GetBundleName() string { return "cisco_ios_xr" }

func (ipv6Icmp *Ipv6Configuration_Ipv6Icmp) GetYangName() string { return "ipv6icmp" }

func (ipv6Icmp *Ipv6Configuration_Ipv6Icmp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ipv6Icmp *Ipv6Configuration_Ipv6Icmp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ipv6Icmp *Ipv6Configuration_Ipv6Icmp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ipv6Icmp *Ipv6Configuration_Ipv6Icmp) SetParent(parent types.Entity) { ipv6Icmp.parent = parent }

func (ipv6Icmp *Ipv6Configuration_Ipv6Icmp) GetParent() types.Entity { return ipv6Icmp.parent }

func (ipv6Icmp *Ipv6Configuration_Ipv6Icmp) GetParentYangName() string { return "ipv6-configuration" }

