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
    EntityData types.CommonEntityData
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

func (ipv6Configuration *Ipv6Configuration) GetEntityData() *types.CommonEntityData {
    ipv6Configuration.EntityData.YFilter = ipv6Configuration.YFilter
    ipv6Configuration.EntityData.YangName = "ipv6-configuration"
    ipv6Configuration.EntityData.BundleName = "cisco_ios_xr"
    ipv6Configuration.EntityData.ParentYangName = "Cisco-IOS-XR-ipv6-io-cfg"
    ipv6Configuration.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-io-cfg:ipv6-configuration"
    ipv6Configuration.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Configuration.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Configuration.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Configuration.EntityData.Children = make(map[string]types.YChild)
    ipv6Configuration.EntityData.Children["ipv6-assembler"] = types.YChild{"Ipv6Assembler", &ipv6Configuration.Ipv6Assembler}
    ipv6Configuration.EntityData.Children["ipv6icmp"] = types.YChild{"Ipv6Icmp", &ipv6Configuration.Ipv6Icmp}
    ipv6Configuration.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Configuration.EntityData.Leafs["ipv6-pmtu-time-out"] = types.YLeaf{"Ipv6PmtuTimeOut", ipv6Configuration.Ipv6PmtuTimeOut}
    ipv6Configuration.EntityData.Leafs["ipv6-source-route"] = types.YLeaf{"Ipv6SourceRoute", ipv6Configuration.Ipv6SourceRoute}
    ipv6Configuration.EntityData.Leafs["ipv6-pmtu-enable"] = types.YLeaf{"Ipv6PmtuEnable", ipv6Configuration.Ipv6PmtuEnable}
    ipv6Configuration.EntityData.Leafs["ipv6-hop-limit"] = types.YLeaf{"Ipv6HopLimit", ipv6Configuration.Ipv6HopLimit}
    return &(ipv6Configuration.EntityData)
}

// Ipv6Configuration_Ipv6Assembler
// IPv6 fragmented packet assembler
type Ipv6Configuration_Ipv6Assembler struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of seconds an assembly queue will hold before timeout. The type is
    // interface{} with range: 1..120. Units are second.
    Timeout interface{}

    // Maxinum packets allowed in assembly queues (in percent). The type is
    // interface{} with range: 1..50. Units are percentage.
    MaxPackets interface{}
}

func (ipv6Assembler *Ipv6Configuration_Ipv6Assembler) GetEntityData() *types.CommonEntityData {
    ipv6Assembler.EntityData.YFilter = ipv6Assembler.YFilter
    ipv6Assembler.EntityData.YangName = "ipv6-assembler"
    ipv6Assembler.EntityData.BundleName = "cisco_ios_xr"
    ipv6Assembler.EntityData.ParentYangName = "ipv6-configuration"
    ipv6Assembler.EntityData.SegmentPath = "ipv6-assembler"
    ipv6Assembler.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Assembler.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Assembler.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Assembler.EntityData.Children = make(map[string]types.YChild)
    ipv6Assembler.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Assembler.EntityData.Leafs["timeout"] = types.YLeaf{"Timeout", ipv6Assembler.Timeout}
    ipv6Assembler.EntityData.Leafs["max-packets"] = types.YLeaf{"MaxPackets", ipv6Assembler.MaxPackets}
    return &(ipv6Assembler.EntityData)
}

// Ipv6Configuration_Ipv6Icmp
// Configure IPv6 ICMP parameters
// This type is a presence type.
type Ipv6Configuration_Ipv6Icmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interval between tokens in milliseconds. The type is interface{} with
    // range: 0..2147483647. This attribute is mandatory. Units are millisecond.
    ErrorInterval interface{}

    // Bucket size. The type is interface{} with range: 1..200. The default value
    // is 10.
    BucketSize interface{}
}

func (ipv6Icmp *Ipv6Configuration_Ipv6Icmp) GetEntityData() *types.CommonEntityData {
    ipv6Icmp.EntityData.YFilter = ipv6Icmp.YFilter
    ipv6Icmp.EntityData.YangName = "ipv6icmp"
    ipv6Icmp.EntityData.BundleName = "cisco_ios_xr"
    ipv6Icmp.EntityData.ParentYangName = "ipv6-configuration"
    ipv6Icmp.EntityData.SegmentPath = "ipv6icmp"
    ipv6Icmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6Icmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6Icmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6Icmp.EntityData.Children = make(map[string]types.YChild)
    ipv6Icmp.EntityData.Leafs = make(map[string]types.YLeaf)
    ipv6Icmp.EntityData.Leafs["error-interval"] = types.YLeaf{"ErrorInterval", ipv6Icmp.ErrorInterval}
    ipv6Icmp.EntityData.Leafs["bucket-size"] = types.YLeaf{"BucketSize", ipv6Icmp.BucketSize}
    return &(ipv6Icmp.EntityData)
}

